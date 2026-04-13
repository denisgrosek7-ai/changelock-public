package metrics

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMetricsHandlerExposesRegisteredMetrics(t *testing.T) {
	IncDecision("policy-engine", "ALLOW", "policy_decision")
	IncDecision("policy-engine", "DENY", "policy_decision")
	IncDecision("policy-engine", "ERROR", "policy_decision")
	IncArtifactVerificationSuccess("attestation-verifier")
	IncArtifactVerificationFailure("attestation-verifier")
	IncRuntimeNoDrift("runtime-agent")
	IncRuntimeDrift("runtime-agent", "image_drift")
	IncAuditForwardingFailure("deploy-gate")
	IncAuditStoreWriteSuccess("audit-writer", "postgres")
	IncAuditStoreWriteFailure("audit-writer", "postgres")

	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	rec := httptest.NewRecorder()

	Handler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	expected := []string{
		"changelock_decision_allow_total",
		"changelock_decision_deny_total",
		"changelock_decision_error_total",
		"changelock_artifact_verification_success_total",
		"changelock_artifact_verification_failure_total",
		"changelock_runtime_drift_total",
		"changelock_runtime_no_drift_total",
		"changelock_audit_forwarding_failure_total",
		"changelock_audit_store_write_success_total",
		"changelock_audit_store_write_failure_total",
	}
	for _, metricName := range expected {
		if !strings.Contains(body, metricName) {
			t.Fatalf("expected metrics output to contain %s", metricName)
		}
	}
}

func TestInstrumentHTTPRecordsRequests(t *testing.T) {
	handler := InstrumentHTTP("policy-engine", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	}))

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	metricsReq := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	metricsRec := httptest.NewRecorder()
	Handler().ServeHTTP(metricsRec, metricsReq)

	body := metricsRec.Body.String()
	if !strings.Contains(body, `changelock_http_requests_total{component="policy-engine",method="GET",route="/health",status="202"}`) {
		t.Fatalf("expected instrumented request metric in output, got %s", body)
	}
}
