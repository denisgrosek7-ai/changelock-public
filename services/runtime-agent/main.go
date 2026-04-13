package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/metrics"
	runtimestate "github.com/denisgrosek/changelock/internal/runtime"
)

var auditWriter = audit.NewDefaultWriter()
var stateReader = newStateReader()

type scanRequest struct {
	ScanID   string                              `json:"scan_id,omitempty"`
	Approved runtimestate.ApprovedWorkloadState  `json:"approved"`
	Observed *runtimestate.ObservedWorkloadState `json:"observed,omitempty"`
}

type scanResponse struct {
	ScanID string                        `json:"scan_id"`
	Result runtimestate.ComparisonResult `json:"result"`
}

func main() {
	addr := ":" + envOrDefault("PORT", "8093")
	log.Printf("runtime-agent listening on %s", addr)
	log.Fatal((&http.Server{
		Addr:              addr,
		Handler:           newHandler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
	}).ListenAndServe())
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.Handle("/metrics", metrics.Handler())
	mux.HandleFunc("/scan", scanHandler)
	return metrics.InstrumentHTTP("runtime-agent", mux)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	httpjson.Write(w, http.StatusOK, map[string]string{"status": "ok"})
}

func scanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var request scanRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	scanID := firstNonEmpty(request.ScanID, r.Header.Get("X-Request-Id"), audit.NewRequestID())
	observed, err := resolveObservedState(r.Context(), request)
	if err != nil {
		event := audit.Event{
			RequestID:   scanID,
			Component:   "runtime-agent",
			EventType:   audit.EventTypeRuntimeDriftResult,
			TenantID:    audit.TenantFromNamespace(request.Approved.Namespace),
			Environment: audit.EnvironmentFromNamespace(request.Approved.Namespace),
			Namespace:   request.Approved.Namespace,
			Workload:    request.Approved.Workload,
			Decision:    audit.DecisionError,
			Reasons:     []string{"runtime state unavailable: " + err.Error()},
		}
		writeAuditEvent(r.Context(), event)
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	result := runtimestate.Compare(request.Approved, observed)
	result.ScanID = scanID

	writeAuditEvent(r.Context(), buildAuditEvent(scanID, result))
	httpjson.Write(w, http.StatusOK, scanResponse{
		ScanID: scanID,
		Result: result,
	})
}

func resolveObservedState(ctx context.Context, request scanRequest) (runtimestate.ObservedWorkloadState, error) {
	if request.Observed != nil {
		return *request.Observed, nil
	}
	if request.Approved.Namespace == "" || request.Approved.Workload == "" {
		return runtimestate.ObservedWorkloadState{}, errors.New("approved namespace and workload are required when observed runtime state is omitted")
	}
	return stateReader.ReadObservedWorkload(ctx, runtimestate.WorkloadTarget{
		Namespace: request.Approved.Namespace,
		Workload:  request.Approved.Workload,
	})
}

func buildAuditEvent(scanID string, result runtimestate.ComparisonResult) audit.Event {
	decision := audit.DecisionAllow
	if result.HasDrift() {
		decision = audit.DecisionDeny
	}

	return audit.Event{
		RequestID:    scanID,
		Component:    "runtime-agent",
		EventType:    audit.EventTypeRuntimeDriftResult,
		TenantID:     audit.TenantFromNamespace(result.Namespace),
		Environment:  audit.EnvironmentFromNamespace(result.Namespace),
		Namespace:    result.Namespace,
		Workload:     result.Workload,
		Image:        result.Image,
		Digest:       firstNonEmpty(result.RunningDigest, result.ApprovedDigest),
		Decision:     decision,
		Reasons:      result.Reasons,
		DriftResult:  result.Result,
		DriftClasses: append([]string(nil), result.Classes...),
		Evidence:     audit.FromRuntimeComparison(&result),
	}
}

func writeAuditEvent(ctx context.Context, event audit.Event) {
	metrics.IncDecision("runtime-agent", event.Decision, event.EventType)
	if event.DriftResult == string(runtimestate.DriftClassNoDrift) {
		metrics.IncRuntimeNoDrift("runtime-agent")
	} else if event.DriftResult != "" {
		metrics.IncRuntimeDrift("runtime-agent", event.DriftResult)
	}
	if err := auditWriter.Write(ctx, event); err != nil {
		log.Printf("runtime-agent audit write failed: %v", err)
	}
}

func newStateReader() runtimestate.StateReader {
	path := os.Getenv("CHANGELOCK_RUNTIME_FIXTURE")
	if path == "" {
		return runtimestate.NoopReader{}
	}

	reader, err := runtimestate.NewFixtureReader(path)
	if err != nil {
		log.Printf("runtime-agent fixture reader unavailable: %v", err)
		return runtimestate.NoopReader{}
	}
	return reader
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
