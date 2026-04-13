package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/metrics"
)

type server struct {
	store          audit.Store
	backend        string
	allowedOrigins map[string]struct{}
	requestTimeout time.Duration
}

type ingestResponse struct {
	Status     string    `json:"status"`
	ID         int64     `json:"id"`
	RequestID  string    `json:"request_id"`
	ReceivedAt time.Time `json:"received_at"`
}

type eventsResponse struct {
	Events []audit.StoredEvent `json:"events"`
}

func main() {
	migrateOnly := flag.Bool("migrate-only", false, "apply database migrations and exit")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	store, backend, err := newStoreFromEnv(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	if *migrateOnly {
		log.Printf("audit-writer migrations applied using %s backend", backend)
		return
	}

	addr := ":" + envOrDefault("PORT", "8094")
	log.Printf("audit-writer listening on %s using %s backend", addr, backend)
	httpServer := &http.Server{
		Addr:              addr,
		Handler:           newHandler(store, backend),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Fatal(httpServer.ListenAndServe())
}

func newHandler(store audit.Store, backend string) http.Handler {
	srv := server{
		store:          store,
		backend:        backend,
		allowedOrigins: allowedOriginsFromEnv(),
		requestTimeout: envDurationOrDefault("CHANGELOCK_REPORTS_TIMEOUT", 5*time.Second),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", srv.healthHandler)
	mux.Handle("/metrics", metrics.Handler())
	mux.HandleFunc("/v1/ingest", srv.ingestHandler)
	mux.HandleFunc("/v1/reports/events", srv.eventsHandler)
	mux.HandleFunc("/v1/reports/summary", srv.summaryHandler)
	mux.HandleFunc("/v1/reports/denies", srv.deniesHandler)
	mux.HandleFunc("/v1/reports/runtime-drift", srv.runtimeDriftHandler)
	return metrics.InstrumentHTTP("audit-writer", srv.wrap(mux))
}

func newStoreFromEnv(ctx context.Context) (audit.Store, string, error) {
	storeKind := strings.ToLower(strings.TrimSpace(os.Getenv("CHANGELOCK_AUDIT_STORE")))
	dsn := strings.TrimSpace(firstNonEmpty(os.Getenv("CHANGELOCK_POSTGRES_DSN"), os.Getenv("DATABASE_URL")))

	switch storeKind {
	case "", "auto":
		if dsn == "" {
			return audit.NewMemoryStore(), "memory", nil
		}
		return newPostgresStore(ctx, dsn)
	case "memory":
		return audit.NewMemoryStore(), "memory", nil
	case "postgres":
		if dsn == "" {
			return nil, "", errors.New("CHANGELOCK_POSTGRES_DSN is required when CHANGELOCK_AUDIT_STORE=postgres")
		}
		return newPostgresStore(ctx, dsn)
	default:
		return nil, "", errors.New("unsupported CHANGELOCK_AUDIT_STORE: " + storeKind)
	}
}

func newPostgresStore(ctx context.Context, dsn string) (audit.Store, string, error) {
	store, err := audit.NewPostgresStore(ctx, dsn)
	if err != nil {
		return nil, "", err
	}
	if err := store.Migrate(ctx); err != nil {
		store.Close()
		return nil, "", err
	}
	return store, "postgres", nil
}

func (s server) healthHandler(w http.ResponseWriter, _ *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := s.store.Ping(ctx); err != nil {
		httpjson.Write(w, http.StatusServiceUnavailable, map[string]string{
			"status":  "error",
			"backend": s.backend,
			"error":   err.Error(),
		})
		return
	}

	httpjson.Write(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"backend": s.backend,
	})
}

func (s server) ingestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var event audit.Event
	if err := httpjson.Decode(r, &event); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if event.RequestID == "" {
		event.RequestID = requestIDFromHeader(r)
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	record, err := s.store.Ingest(ctx, event)
	if err != nil {
		metrics.IncAuditStoreWriteFailure("audit-writer", s.backend)
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidEvent) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	metrics.IncAuditStoreWriteSuccess("audit-writer", s.backend)

	httpjson.Write(w, http.StatusCreated, ingestResponse{
		Status:     "stored",
		ID:         record.ID,
		RequestID:  record.RequestID,
		ReceivedAt: record.ReceivedAt,
	})
}

func (s server) eventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	events, err := s.store.ListEvents(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	httpjson.Write(w, http.StatusOK, eventsResponse{Events: events})
}

func (s server) summaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	summary, err := s.store.Summary(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	httpjson.Write(w, http.StatusOK, summary)
}

func (s server) deniesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	query.Set("decision", audit.DecisionDeny)
	r.URL.RawQuery = query.Encode()
	s.eventsHandler(w, r)
}

func (s server) runtimeDriftHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	query.Set("event_type", audit.EventTypeRuntimeDriftResult)
	r.URL.RawQuery = query.Encode()
	s.eventsHandler(w, r)
}

func parseFilter(r *http.Request) (audit.EventFilter, error) {
	query := r.URL.Query()
	filter := audit.EventFilter{
		Decision:    query.Get("decision"),
		EventType:   query.Get("event_type"),
		Component:   query.Get("component"),
		Repo:        query.Get("repo"),
		Environment: query.Get("environment"),
		TenantID:    query.Get("tenant_id"),
	}
	if rawLimit := strings.TrimSpace(query.Get("limit")); rawLimit != "" {
		limit, err := strconv.Atoi(rawLimit)
		if err != nil {
			return audit.EventFilter{}, errors.New("limit must be an integer")
		}
		filter.Limit = limit
	}
	return audit.NormalizeFilter(filter)
}

func requestIDFromHeader(r *http.Request) string {
	if requestID := strings.TrimSpace(r.Header.Get("X-Request-Id")); requestID != "" {
		return requestID
	}
	return audit.NewRequestID()
}

func envOrDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}

func envDurationOrDefault(key string, fallback time.Duration) time.Duration {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	parsed, err := time.ParseDuration(value)
	if err != nil || parsed <= 0 {
		return fallback
	}
	return parsed
}

func firstNonEmpty(values ...string) string {
	return audit.FirstNonEmpty(values...)
}

func allowedOriginsFromEnv() map[string]struct{} {
	raw := strings.TrimSpace(os.Getenv("CHANGELOCK_CORS_ALLOW_ORIGINS"))
	if raw == "" {
		raw = strings.Join([]string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
			"http://localhost:3000",
			"http://127.0.0.1:3000",
		}, ",")
	}

	allowed := map[string]struct{}{}
	for _, origin := range strings.Split(raw, ",") {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			allowed[origin] = struct{}{}
		}
	}
	return allowed
}

func (s server) wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.applySecurityHeaders(w, r)
		if s.handleCORS(w, r) {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s server) applySecurityHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if r.URL.Path == "/health" || strings.HasPrefix(r.URL.Path, "/v1/") {
		w.Header().Set("Cache-Control", "no-store, max-age=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
	}
}

func (s server) handleCORS(w http.ResponseWriter, r *http.Request) bool {
	origin := strings.TrimSpace(r.Header.Get("Origin"))
	if origin != "" {
		w.Header().Add("Vary", "Origin")
	}
	w.Header().Add("Vary", "Access-Control-Request-Method")
	w.Header().Add("Vary", "Access-Control-Request-Headers")

	if origin != "" {
		if _, ok := s.allowedOrigins[origin]; ok {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Request-Id")
			w.Header().Set("Access-Control-Max-Age", "600")
		} else if r.Method == http.MethodOptions {
			httpjson.Write(w, http.StatusForbidden, map[string]string{"error": "origin not allowed"})
			return true
		}
	}

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return true
	}
	return false
}
