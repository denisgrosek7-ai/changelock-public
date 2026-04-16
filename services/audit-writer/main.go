package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/metrics"
)

type server struct {
	store          audit.Store
	backend        string
	allowedOrigins map[string]struct{}
	requestTimeout time.Duration
	authConfig     auth.Config
	vulnOps        *vulnOpsRuntime
	syncRuntime    *syncRuntime
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

type exceptionsResponse struct {
	Exceptions []audit.PolicyException `json:"exceptions"`
}

type exceptionResponse struct {
	Status    string                `json:"status"`
	Exception audit.PolicyException `json:"exception"`
}

type exceptionActionResponse struct {
	Status    string                `json:"status"`
	Exception audit.PolicyException `json:"exception"`
}

type authInfoResponse struct {
	Authenticated bool   `json:"authenticated"`
	AuthMode      string `json:"auth_mode"`
	Subject       string `json:"subject,omitempty"`
	Role          string `json:"role,omitempty"`
	TokenID       string `json:"token_id,omitempty"`
	IdentityType  string `json:"identity_type,omitempty"`
	Email         string `json:"email,omitempty"`
	TenantID      string `json:"tenant_id,omitempty"`
	GlobalScope   bool   `json:"global_scope,omitempty"`
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

	authConfig, err := loadAuthConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	vulnOps, err := loadVulnOpsRuntimeFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	syncRuntime, err := loadSyncRuntimeFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	if *migrateOnly {
		log.Printf("audit-writer migrations applied using %s backend", backend)
		return
	}
	if vulnOps != nil {
		vulnOps.start(context.Background(), store)
	}
	if syncRuntime != nil {
		syncRuntime.start(context.Background(), store)
	}

	addr := ":" + envOrDefault("PORT", "8094")
	log.Printf("audit-writer listening on %s using %s backend", addr, backend)
	httpServer := &http.Server{
		Addr:              addr,
		Handler:           newHandlerWithDeps(store, backend, authConfig, vulnOps),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Fatal(httpServer.ListenAndServe())
}

func newHandler(store audit.Store, backend string) http.Handler {
	authConfig, err := loadAuthConfigFromEnv()
	if err != nil {
		panic(err)
	}
	vulnOps, err := loadVulnOpsRuntimeFromEnv()
	if err != nil {
		panic(err)
	}
	return newHandlerWithDeps(store, backend, authConfig, vulnOps)
}

func newHandlerWithAuth(store audit.Store, backend string, authConfig auth.Config) http.Handler {
	vulnOps, err := loadVulnOpsRuntimeFromEnv()
	if err != nil {
		panic(err)
	}
	syncRuntime, err := loadSyncRuntimeFromEnv()
	if err != nil {
		panic(err)
	}
	return newHandlerWithRuntimes(store, backend, authConfig, vulnOps, syncRuntime)
}

func newHandlerWithDeps(store audit.Store, backend string, authConfig auth.Config, vulnOps *vulnOpsRuntime) http.Handler {
	syncRuntime, err := loadSyncRuntimeFromEnv()
	if err != nil {
		panic(err)
	}
	return newHandlerWithRuntimes(store, backend, authConfig, vulnOps, syncRuntime)
}

func newHandlerWithRuntimes(store audit.Store, backend string, authConfig auth.Config, vulnOps *vulnOpsRuntime, syncRuntime *syncRuntime) http.Handler {
	srv := server{
		store:          store,
		backend:        backend,
		allowedOrigins: allowedOriginsFromEnv(),
		requestTimeout: envDurationOrDefault("CHANGELOCK_REPORTS_TIMEOUT", 5*time.Second),
		authConfig:     authConfig,
		vulnOps:        vulnOps,
		syncRuntime:    syncRuntime,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", srv.healthHandler)
	mux.HandleFunc("/ready", srv.readyHandler)
	mux.HandleFunc("/v1/sync/status", srv.syncStatusHandler)
	mux.HandleFunc("/v1/sync/exceptions", srv.syncExceptionsHandler)
	mux.Handle("/metrics", metrics.Handler())
	mux.HandleFunc("/v1/ingest", srv.ingestHandler)
	mux.HandleFunc("/v1/auth/me", srv.authMeHandler)
	mux.HandleFunc("/v1/sbom/ingest", srv.sbomIngestHandler)
	mux.HandleFunc("/v1/sbom/images/", srv.sbomImageHandler)
	mux.HandleFunc("/v1/sbom/components/search", srv.sbomComponentsSearchHandler)
	mux.HandleFunc("/v1/exceptions", srv.exceptionsHandler)
	mux.HandleFunc("/v1/exceptions/request", srv.requestExceptionHandler)
	mux.HandleFunc("/v1/exceptions/", srv.exceptionByIDHandler)
	mux.HandleFunc("/v1/exceptions/validate", srv.validateExceptionHandler)
	mux.HandleFunc("/v1/analytics/trends", srv.trendsHandler)
	mux.HandleFunc("/v1/analytics/top-violators", srv.topViolatorsHandler)
	mux.HandleFunc("/v1/analytics/drift-stats", srv.driftStatsHandler)
	mux.HandleFunc("/v1/vulnerabilities/active", srv.activeVulnerabilitiesHandler)
	mux.HandleFunc("/v1/vulnerabilities/blast-radius", srv.vulnerabilityBlastRadiusHandler)
	mux.HandleFunc("/v1/vulnerabilities/timeline", srv.vulnerabilityTimelineHandler)
	mux.HandleFunc("/v1/vulnerabilities/rescan", srv.vulnerabilityRescanHandler)
	mux.HandleFunc("/v1/vulnerabilities/decisions", srv.vulnerabilityDecisionsHandler)
	mux.HandleFunc("/v1/vulnerabilities/decisions/", srv.vulnerabilityDecisionByIDHandler)
	mux.HandleFunc("/v1/reports/events", srv.eventsHandler)
	mux.HandleFunc("/v1/reports/summary", srv.summaryHandler)
	mux.HandleFunc("/v1/reports/denies", srv.deniesHandler)
	mux.HandleFunc("/v1/reports/runtime-drift", srv.runtimeDriftHandler)
	mux.HandleFunc("/v1/reports/exceptions", srv.exceptionsReportHandler)
	return metrics.InstrumentHTTP("audit-writer", srv.wrap(mux))
}

func loadAuthConfigFromEnv() (auth.Config, error) {
	return auth.ParseEnvConfig(os.Getenv)
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
	httpjson.Write(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"backend": s.backend,
	})
}

func (s server) readyHandler(w http.ResponseWriter, _ *http.Request) {
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
	clusterID, err := s.resolveInboundClusterID(r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if clusterID != "" {
		event.ClusterID = clusterID
	} else if s.syncRuntime != nil && s.syncRuntime.config.Mode == audit.SyncModeSpoke && strings.TrimSpace(event.ClusterID) == "" {
		event.ClusterID = s.syncRuntime.config.ClusterID
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
	if err := s.syncRuntime.forwardEvent(ctx, event); err != nil {
		log.Printf("audit-writer sync forward failed: %v", err)
	}

	httpjson.Write(w, http.StatusCreated, ingestResponse{
		Status:     "stored",
		ID:         record.ID,
		RequestID:  record.RequestID,
		ReceivedAt: record.ReceivedAt,
	})
}

func (s server) authMeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	principal, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	httpjson.Write(w, http.StatusOK, authInfoResponse{
		Authenticated: principal.Authenticated,
		AuthMode:      principal.AuthMode,
		Subject:       principal.Subject,
		Role:          principal.Role,
		TokenID:       principal.TokenID,
		IdentityType:  principal.IdentityType,
		Email:         principal.Email,
		TenantID:      principal.TenantID,
		GlobalScope:   principal.GlobalScope,
	})
}

func (s server) eventsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
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

func (s server) exceptionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		r, err := applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parseExceptionFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()

		exceptions, err := s.store.ListExceptions(ctx, filter)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, audit.ErrInvalidException) {
				status = http.StatusBadRequest
			} else if errors.Is(err, context.DeadlineExceeded) {
				status = http.StatusGatewayTimeout
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}

		httpjson.Write(w, http.StatusOK, exceptionsResponse{Exceptions: exceptions})
	case http.MethodPost:
		if reason := s.exceptionMutationBlockedReason(); reason != "" {
			httpjson.Write(w, http.StatusConflict, map[string]string{"error": reason})
			return
		}
		principal, r, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		var request audit.ExceptionCreateRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request, err := applyPrincipalTenantToExceptionRequest(principal, request)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()

		if strings.TrimSpace(request.ApprovedBy) == "" {
			request.ApprovedBy = principal.Subject
		}
		exception, err := s.store.CreateException(ctx, request)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, audit.ErrInvalidException) {
				status = http.StatusBadRequest
			} else if errors.Is(err, context.DeadlineExceeded) {
				status = http.StatusGatewayTimeout
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}

		s.writeLifecycleAuditEvent(ctx, r, principal.Subject, audit.EventTypeExceptionApproved, audit.DecisionAllow, exception, "direct emergency exception created as approved")
		httpjson.Write(w, http.StatusCreated, exceptionResponse{
			Status:    "created",
			Exception: exception,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) requestExceptionHandler(w http.ResponseWriter, r *http.Request) {
	if reason := s.exceptionMutationBlockedReason(); reason != "" {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": reason})
		return
	}
	principal, r, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var request audit.ExceptionCreateRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	request, err := applyPrincipalTenantToExceptionRequest(principal, request)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	exception, err := s.store.RequestException(ctx, request, principal.Subject, principal.Role)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidException) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	s.writeLifecycleAuditEvent(ctx, r, principal.Subject, audit.EventTypeExceptionRequested, audit.DecisionAllow, exception, "exception requested for approval")
	httpjson.Write(w, http.StatusCreated, exceptionActionResponse{Status: "requested", Exception: exception})
}

func (s server) exceptionByIDHandler(w http.ResponseWriter, r *http.Request) {
	exceptionID, action, err := exceptionActionFromPath(r.URL.Path)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	switch {
	case action == "" && r.Method == http.MethodDelete:
		if reason := s.exceptionMutationBlockedReason(); reason != "" {
			httpjson.Write(w, http.StatusConflict, map[string]string{"error": reason})
			return
		}
		principal, r, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
		if !ok {
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		existing, err := s.store.GetException(ctx, exceptionID)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, audit.ErrExceptionNotFound) {
				status = http.StatusNotFound
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}
		if err := ensureExceptionTenantAccess(principal, existing); err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}

		exception, err := s.store.RevokeException(ctx, exceptionID)
		if err != nil {
			status := http.StatusInternalServerError
			switch {
			case errors.Is(err, audit.ErrInvalidException):
				status = http.StatusBadRequest
			case errors.Is(err, audit.ErrExceptionNotFound):
				status = http.StatusNotFound
			case errors.Is(err, context.DeadlineExceeded):
				status = http.StatusGatewayTimeout
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}

		s.writeLifecycleAuditEvent(ctx, r, principal.Subject, audit.EventTypeExceptionRevoked, audit.DecisionAllow, exception, "exception revoked")
		httpjson.Write(w, http.StatusOK, exceptionResponse{
			Status:    "revoked",
			Exception: exception,
		})
	case action == "approve" && r.Method == http.MethodPost:
		s.approveExceptionHandler(w, r, exceptionID)
	case action == "reject" && r.Method == http.MethodPost:
		s.rejectExceptionHandler(w, r, exceptionID)
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) approveExceptionHandler(w http.ResponseWriter, r *http.Request, exceptionID string) {
	if reason := s.exceptionMutationBlockedReason(); reason != "" {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": reason})
		return
	}
	principal, r, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
	if !ok {
		return
	}

	var request audit.ExceptionActionRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	existing, err := s.store.GetException(ctx, exceptionID)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrExceptionNotFound) {
			status = http.StatusNotFound
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	if err := ensureExceptionTenantAccess(principal, existing); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	exception, err := s.store.ApproveException(ctx, exceptionID, principal.Subject, principal.Role)
	if err != nil {
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, audit.ErrInvalidException):
			status = http.StatusBadRequest
		case errors.Is(err, audit.ErrExceptionNotFound):
			status = http.StatusNotFound
		case errors.Is(err, context.DeadlineExceeded):
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	reason := "exception approved"
	if normalized := audit.NormalizeExceptionActionRequest(request); normalized.Reason != "" {
		reason = normalized.Reason
	}
	s.writeLifecycleAuditEvent(ctx, r, principal.Subject, audit.EventTypeExceptionApproved, audit.DecisionAllow, exception, reason)
	httpjson.Write(w, http.StatusOK, exceptionActionResponse{Status: "approved", Exception: exception})
}

func (s server) rejectExceptionHandler(w http.ResponseWriter, r *http.Request, exceptionID string) {
	if reason := s.exceptionMutationBlockedReason(); reason != "" {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": reason})
		return
	}
	principal, r, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
	if !ok {
		return
	}

	var request audit.ExceptionActionRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	request = audit.NormalizeExceptionActionRequest(request)
	if request.Reason == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "reason is required"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	existing, err := s.store.GetException(ctx, exceptionID)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrExceptionNotFound) {
			status = http.StatusNotFound
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	if err := ensureExceptionTenantAccess(principal, existing); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	exception, err := s.store.RejectException(ctx, exceptionID, request.Reason, principal.Subject, principal.Role)
	if err != nil {
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, audit.ErrInvalidException):
			status = http.StatusBadRequest
		case errors.Is(err, audit.ErrExceptionNotFound):
			status = http.StatusNotFound
		case errors.Is(err, context.DeadlineExceeded):
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	s.writeLifecycleAuditEvent(ctx, r, principal.Subject, audit.EventTypeExceptionRejected, audit.DecisionDeny, exception, request.Reason)
	httpjson.Write(w, http.StatusOK, exceptionActionResponse{Status: "rejected", Exception: exception})
}

func (s server) validateExceptionHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleService, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var request audit.ExceptionValidationRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if principal.Role != auth.RoleService {
		updatedRequest, err := applyPrincipalTenantToExceptionValidation(principal, request)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		request = updatedRequest
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	if reason := s.exceptionValidationBlockedReason(); reason != "" {
		httpjson.Write(w, http.StatusOK, audit.ExceptionValidationResult{Valid: false, Reason: reason})
		return
	}

	result, err := s.store.ValidateException(ctx, request)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidException) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	httpjson.Write(w, http.StatusOK, result)
}

func (s server) trendsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseTrendsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	response, err := s.store.Trends(ctx, filter)
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

	httpjson.Write(w, http.StatusOK, response)
}

func (s server) topViolatorsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseTopViolatorsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	response, err := s.store.TopViolators(ctx, filter)
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

	httpjson.Write(w, http.StatusOK, response)
}

func (s server) driftStatsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseDriftStatsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	response, err := s.store.DriftStats(ctx, filter)
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

	httpjson.Write(w, http.StatusOK, response)
}

func (s server) summaryHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
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
	if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
		return
	}
	query := r.URL.Query()
	query.Set("decision", audit.DecisionDeny)
	r.URL.RawQuery = query.Encode()
	s.eventsHandler(w, r)
}

func (s server) runtimeDriftHandler(w http.ResponseWriter, r *http.Request) {
	if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
		return
	}
	query := r.URL.Query()
	query.Set("event_type", audit.EventTypeRuntimeDriftResult)
	r.URL.RawQuery = query.Encode()
	s.eventsHandler(w, r)
}

func (s server) exceptionsReportHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseExceptionFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	report, err := s.store.ExceptionReport(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidException) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	httpjson.Write(w, http.StatusOK, report)
}

func (s server) syncStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
		return
	}

	status := deriveSyncStatus(audit.SyncStatus{
		SyncMode: audit.SyncModeDisabled,
		Mode:     audit.SyncModeDisabled,
		Health:   audit.SyncHealthDisabled,
	}, syncConfig{}, time.Now().UTC())
	if s.syncRuntime != nil {
		status = s.syncRuntime.statusSnapshot()
		if s.syncRuntime.config.Mode == audit.SyncModeHub {
			revision, err := s.currentHubExceptionRevision(r.Context())
			if err == nil {
				status.CurrentRevision = revision
				status.RevisionETag = quotedETag(revision)
			}
		}
		status = deriveSyncStatus(status, s.syncRuntime.config, time.Now().UTC())
	}
	httpjson.Write(w, http.StatusOK, status)
}

func (s server) syncExceptionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	if s.syncRuntime == nil || s.syncRuntime.config.Mode != audit.SyncModeHub {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "sync hub endpoint is disabled"})
		return
	}
	principal, _, ok := s.authorize(w, r, auth.RoleService)
	if !ok {
		return
	}
	clusterID := strings.TrimSpace(r.Header.Get(syncClusterHeader))
	binding, err := s.syncRuntime.authorizeClusterPrincipal(principal, clusterID)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	filtered, revision, err := s.currentHubSyncedExceptions(r.Context(), binding)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	etag := fmt.Sprintf("%q", revision)
	if strings.TrimSpace(r.Header.Get("If-None-Match")) == etag {
		w.Header().Set("ETag", etag)
		w.WriteHeader(http.StatusNotModified)
		return
	}

	response := audit.ExceptionSyncSnapshot{
		ClusterID:   clusterID,
		Revision:    revision,
		GeneratedAt: time.Now().UTC(),
		Exceptions:  filtered,
	}
	w.Header().Set("ETag", etag)
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) authorize(w http.ResponseWriter, r *http.Request, roles ...string) (auth.Principal, *http.Request, bool) {
	principal, err := s.authConfig.Require(r, roles...)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return auth.Principal{}, r, false
	}
	r = r.WithContext(auth.WithPrincipal(r.Context(), principal))
	return principal, r, true
}

func parseFilter(r *http.Request) (audit.EventFilter, error) {
	query := r.URL.Query()
	filter := audit.EventFilter{
		Decision:    query.Get("decision"),
		EventType:   query.Get("event_type"),
		Component:   query.Get("component"),
		ClusterID:   query.Get("cluster_id"),
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

func parseExceptionFilter(r *http.Request) (audit.ExceptionFilter, error) {
	query := r.URL.Query()
	filter := audit.ExceptionFilter{
		Status:        query.Get("status"),
		ExceptionType: query.Get("exception_type"),
		TenantID:      query.Get("tenant_id"),
		Environment:   query.Get("environment"),
		Namespace:     query.Get("namespace"),
		Repo:          query.Get("repo"),
		ImageDigest:   query.Get("image_digest"),
		CVEID:         query.Get("cve_id"),
	}

	if rawActive := strings.TrimSpace(query.Get("active")); rawActive != "" {
		active, err := strconv.ParseBool(rawActive)
		if err != nil {
			return audit.ExceptionFilter{}, errors.New("active must be a boolean")
		}
		filter.Active = &active
	}

	if rawLimit := strings.TrimSpace(query.Get("limit")); rawLimit != "" {
		limit, err := strconv.Atoi(rawLimit)
		if err != nil {
			return audit.ExceptionFilter{}, errors.New("limit must be an integer")
		}
		filter.Limit = limit
	}

	return audit.NormalizeExceptionFilter(filter)
}

func (s server) resolveInboundClusterID(r *http.Request) (string, error) {
	clusterID := strings.TrimSpace(r.Header.Get(syncClusterHeader))
	if clusterID == "" {
		return "", nil
	}
	if s.syncRuntime == nil || s.syncRuntime.config.Mode != audit.SyncModeHub {
		return clusterID, nil
	}
	principal, err := s.authConfig.Require(r, auth.RoleService)
	if err != nil {
		return "", err
	}
	if _, err := s.syncRuntime.authorizeClusterPrincipal(principal, clusterID); err != nil {
		return "", err
	}
	return clusterID, nil
}

func (s server) exceptionMutationBlockedReason() string {
	if s.syncRuntime == nil {
		return ""
	}
	return s.syncRuntime.mutationBlockedReason()
}

func (s server) exceptionValidationBlockedReason() string {
	if s.syncRuntime == nil {
		return ""
	}
	return s.syncRuntime.exceptionValidationBlockReason()
}

func (s server) currentHubSyncedExceptions(ctx context.Context, binding clusterBinding) ([]audit.SyncedException, string, error) {
	filter := audit.ExceptionFilter{
		Status: audit.ExceptionStatusApproved,
		Limit:  500,
	}
	var exceptions []audit.PolicyException
	if len(binding.Tenants) == 0 {
		listed, err := s.store.ListExceptions(ctx, filter)
		if err != nil {
			return nil, "", err
		}
		exceptions = listed
	} else {
		byID := map[string]audit.PolicyException{}
		for _, tenantID := range binding.Tenants {
			filter.TenantID = tenantID
			listed, err := s.store.ListExceptions(ctx, filter)
			if err != nil {
				return nil, "", err
			}
			for _, exception := range listed {
				byID[exception.ExceptionID] = exception
			}
		}
		exceptions = make([]audit.PolicyException, 0, len(byID))
		for _, exception := range byID {
			exceptions = append(exceptions, exception)
		}
	}

	filtered := filterSyncedExceptionsForBinding(exceptions, binding)
	revision := audit.ComputeExceptionSyncRevision(filtered)
	return filtered, revision, nil
}

func (s server) currentHubExceptionRevision(ctx context.Context) (string, error) {
	exceptions, err := s.store.ListExceptions(ctx, audit.ExceptionFilter{
		Status: audit.ExceptionStatusApproved,
		Limit:  500,
	})
	if err != nil {
		return "", err
	}
	synced := make([]audit.SyncedException, 0, len(exceptions))
	for _, exception := range exceptions {
		synced = append(synced, audit.SyncedExceptionFromPolicyException(exception))
	}
	return audit.ComputeExceptionSyncRevision(synced), nil
}

func parseTrendsFilter(r *http.Request) (audit.TrendsFilter, error) {
	query := r.URL.Query()
	filter := audit.TrendsFilter{
		WindowDays:  30,
		Granularity: query.Get("granularity"),
		ClusterID:   query.Get("cluster_id"),
		TenantID:    query.Get("tenant_id"),
		Environment: query.Get("environment"),
		Repo:        query.Get("repo"),
		EventType:   query.Get("event_type"),
	}
	if raw := strings.TrimSpace(query.Get("window_days")); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil {
			return audit.TrendsFilter{}, errors.New("window_days must be an integer")
		}
		filter.WindowDays = value
	}
	return audit.NormalizeTrendsFilter(filter)
}

func parseTopViolatorsFilter(r *http.Request) (audit.TopViolatorsFilter, error) {
	query := r.URL.Query()
	filter := audit.TopViolatorsFilter{
		WindowDays:  30,
		Limit:       5,
		Dimension:   query.Get("dimension"),
		ClusterID:   query.Get("cluster_id"),
		TenantID:    query.Get("tenant_id"),
		Environment: query.Get("environment"),
		Repo:        query.Get("repo"),
	}
	if raw := strings.TrimSpace(query.Get("window_days")); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil {
			return audit.TopViolatorsFilter{}, errors.New("window_days must be an integer")
		}
		filter.WindowDays = value
	}
	if raw := strings.TrimSpace(query.Get("limit")); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil {
			return audit.TopViolatorsFilter{}, errors.New("limit must be an integer")
		}
		filter.Limit = value
	}
	return audit.NormalizeTopViolatorsFilter(filter)
}

func parseDriftStatsFilter(r *http.Request) (audit.DriftStatsFilter, error) {
	query := r.URL.Query()
	filter := audit.DriftStatsFilter{
		WindowDays:  30,
		ClusterID:   query.Get("cluster_id"),
		TenantID:    query.Get("tenant_id"),
		Environment: query.Get("environment"),
		Repo:        query.Get("repo"),
		Namespace:   query.Get("namespace"),
		Workload:    query.Get("workload"),
	}
	if raw := strings.TrimSpace(query.Get("window_days")); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil {
			return audit.DriftStatsFilter{}, errors.New("window_days must be an integer")
		}
		filter.WindowDays = value
	}
	return audit.NormalizeDriftStatsFilter(filter)
}

func exceptionActionFromPath(path string) (string, string, error) {
	raw := strings.TrimPrefix(path, "/v1/exceptions/")
	raw = strings.TrimSpace(strings.Trim(raw, "/"))
	if raw == "" || raw == "validate" {
		return "", "", errors.New("exception_id path segment is required")
	}
	parts := strings.Split(raw, "/")
	if len(parts) > 2 {
		return "", "", errors.New("invalid exception path")
	}
	value, err := url.PathUnescape(parts[0])
	if err != nil {
		return "", "", errors.New("invalid exception_id path segment")
	}
	if strings.TrimSpace(value) == "" {
		return "", "", errors.New("exception_id path segment is required")
	}
	action := ""
	if len(parts) == 2 {
		action = strings.TrimSpace(parts[1])
	}
	return value, action, nil
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
	if r.URL.Path == "/health" || r.URL.Path == "/ready" || strings.HasPrefix(r.URL.Path, "/v1/") {
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
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Request-Id")
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

func (s server) writeLifecycleAuditEvent(ctx context.Context, r *http.Request, actor, eventType, decision string, exception audit.PolicyException, reason string) {
	if actor == "" {
		actor = firstNonEmpty(exception.ApprovedBy, exception.RequestedBy, exception.RejectedBy)
	}
	_, _ = s.store.Ingest(ctx, audit.Event{
		RequestID:                requestIDFromHeader(r),
		Component:                "audit-writer",
		EventType:                eventType,
		Actor:                    actor,
		TenantID:                 exception.TenantID,
		Repo:                     exception.Repo,
		Environment:              exception.Environment,
		Namespace:                exception.Namespace,
		Digest:                   exception.ImageDigest,
		CVEID:                    exception.CVEID,
		Decision:                 decision,
		Reasons:                  []string{reason},
		IsException:              true,
		ExceptionID:              exception.ExceptionID,
		ExceptionType:            exception.ExceptionType,
		ExceptionStatus:          exception.Status,
		ExceptionReason:          exception.Reason,
		ExceptionTicketID:        exception.TicketID,
		ExceptionRequestedBy:     exception.RequestedBy,
		ExceptionRequestedAt:     exception.RequestedAt,
		ExceptionApprovedBy:      exception.ApprovedBy,
		ExceptionApprovedAt:      exception.ApprovedAt,
		ExceptionRejectedBy:      exception.RejectedBy,
		ExceptionRejectedAt:      exception.RejectedAt,
		ExceptionRejectionReason: exception.RejectionReason,
		ExceptionExpiresAt:       &exception.ExpiresAt,
	})
}
