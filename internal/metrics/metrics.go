package metrics

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	registerOnce sync.Once

	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_http_requests_total",
			Help: "Total HTTP requests handled by a ChangeLock service.",
		},
		[]string{"component", "route", "method", "status"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "changelock_http_request_duration_seconds",
			Help:    "HTTP request duration in seconds for ChangeLock services.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"component", "route", "method"},
	)
	decisionAllowTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_decision_allow_total",
			Help: "Total allowed security decisions emitted by ChangeLock services.",
		},
		[]string{"component", "event_type"},
	)
	decisionDenyTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_decision_deny_total",
			Help: "Total denied security decisions emitted by ChangeLock services.",
		},
		[]string{"component", "event_type"},
	)
	decisionErrorTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_decision_error_total",
			Help: "Total errored security decisions emitted by ChangeLock services.",
		},
		[]string{"component", "event_type"},
	)
	artifactVerificationSuccessTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_artifact_verification_success_total",
			Help: "Total successful artifact verification decisions.",
		},
		[]string{"component"},
	)
	artifactVerificationFailureTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_artifact_verification_failure_total",
			Help: "Total failed or errored artifact verification decisions.",
		},
		[]string{"component"},
	)
	runtimeDriftTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_runtime_drift_total",
			Help: "Total runtime drift detections grouped by drift result.",
		},
		[]string{"component", "drift_result"},
	)
	runtimeNoDriftTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_runtime_no_drift_total",
			Help: "Total runtime scans that found no drift.",
		},
		[]string{"component"},
	)
	auditForwardingFailureTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_audit_forwarding_failure_total",
			Help: "Total audit forwarding failures when services push events to audit-writer.",
		},
		[]string{"component"},
	)
	auditStoreWriteSuccessTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_audit_store_write_success_total",
			Help: "Total successful persistent audit store writes.",
		},
		[]string{"component", "backend"},
	)
	auditStoreWriteFailureTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_audit_store_write_failure_total",
			Help: "Total failed persistent audit store writes.",
		},
		[]string{"component", "backend"},
	)
	executionAsyncTaskTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_execution_async_task_total",
			Help: "Total Phase 1 async task transitions by task type and state.",
		},
		[]string{"component", "task_type", "state"},
	)
	executionAsyncReplayTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "changelock_execution_async_replay_total",
			Help: "Total Phase 1 async task replay requests.",
		},
		[]string{"component", "task_type", "outcome"},
	)
)

func Handler() http.Handler {
	ensureRegistered()
	return promhttp.Handler()
}

func InstrumentHTTP(component string, next http.Handler) http.Handler {
	ensureRegistered()

	component = normalizeLabel(component, "unknown")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(recorder, r)

		route := normalizeRoute(r.URL.Path)
		method := normalizeLabel(r.Method, http.MethodGet)
		status := strconv.Itoa(recorder.status)

		httpRequestsTotal.WithLabelValues(component, route, method, status).Inc()
		httpRequestDuration.WithLabelValues(component, route, method).Observe(time.Since(start).Seconds())
	})
}

func IncDecision(component, decision, eventType string) {
	ensureRegistered()

	component = normalizeLabel(component, "unknown")
	eventType = normalizeLabel(eventType, "unknown")

	switch decision {
	case "ALLOW":
		decisionAllowTotal.WithLabelValues(component, eventType).Inc()
	case "DENY":
		decisionDenyTotal.WithLabelValues(component, eventType).Inc()
	case "ERROR":
		decisionErrorTotal.WithLabelValues(component, eventType).Inc()
	}
}

func IncArtifactVerificationSuccess(component string) {
	ensureRegistered()
	artifactVerificationSuccessTotal.WithLabelValues(normalizeLabel(component, "unknown")).Inc()
}

func IncArtifactVerificationFailure(component string) {
	ensureRegistered()
	artifactVerificationFailureTotal.WithLabelValues(normalizeLabel(component, "unknown")).Inc()
}

func IncRuntimeDrift(component, driftResult string) {
	ensureRegistered()
	runtimeDriftTotal.WithLabelValues(
		normalizeLabel(component, "unknown"),
		normalizeLabel(driftResult, "unknown"),
	).Inc()
}

func IncRuntimeNoDrift(component string) {
	ensureRegistered()
	runtimeNoDriftTotal.WithLabelValues(normalizeLabel(component, "unknown")).Inc()
}

func IncAuditForwardingFailure(component string) {
	ensureRegistered()
	auditForwardingFailureTotal.WithLabelValues(normalizeLabel(component, "unknown")).Inc()
}

func IncAuditStoreWriteSuccess(component, backend string) {
	ensureRegistered()
	auditStoreWriteSuccessTotal.WithLabelValues(
		normalizeLabel(component, "unknown"),
		normalizeLabel(backend, "unknown"),
	).Inc()
}

func IncAuditStoreWriteFailure(component, backend string) {
	ensureRegistered()
	auditStoreWriteFailureTotal.WithLabelValues(
		normalizeLabel(component, "unknown"),
		normalizeLabel(backend, "unknown"),
	).Inc()
}

func IncExecutionAsyncTask(component, taskType, state string) {
	ensureRegistered()
	executionAsyncTaskTotal.WithLabelValues(
		normalizeLabel(component, "unknown"),
		normalizeLabel(taskType, "unknown"),
		normalizeLabel(state, "unknown"),
	).Inc()
}

func IncExecutionAsyncReplay(component, taskType, outcome string) {
	ensureRegistered()
	executionAsyncReplayTotal.WithLabelValues(
		normalizeLabel(component, "unknown"),
		normalizeLabel(taskType, "unknown"),
		normalizeLabel(outcome, "unknown"),
	).Inc()
}

func ensureRegistered() {
	registerOnce.Do(func() {
		prometheus.MustRegister(
			httpRequestsTotal,
			httpRequestDuration,
			decisionAllowTotal,
			decisionDenyTotal,
			decisionErrorTotal,
			artifactVerificationSuccessTotal,
			artifactVerificationFailureTotal,
			runtimeDriftTotal,
			runtimeNoDriftTotal,
			auditForwardingFailureTotal,
			auditStoreWriteSuccessTotal,
			auditStoreWriteFailureTotal,
			executionAsyncTaskTotal,
			executionAsyncReplayTotal,
		)
	})
}

func normalizeRoute(path string) string {
	if path == "" {
		return "/"
	}
	return path
}

func normalizeLabel(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
