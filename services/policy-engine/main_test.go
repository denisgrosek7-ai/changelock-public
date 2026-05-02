package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/policy"
	"github.com/denisgrosek/changelock/internal/verify"
)

type fakeExceptionValidator struct {
	result audit.ExceptionValidationResult
	err    error
}

func (f fakeExceptionValidator) Validate(_ context.Context, _ audit.ExceptionValidationRequest) (audit.ExceptionValidationResult, error) {
	return f.result, f.err
}

func TestEvaluateArtifactWritesPolicyAuditEvent(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")
	auditPath := filepath.Join(t.TempDir(), "audit.jsonl")
	previousWriter := auditWriter
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	defer func() { auditWriter = previousWriter }()

	payload, err := json.Marshal(policy.ArtifactEvaluationRequest{
		Tenant:     "acme",
		Repository: "my-org/acme-app",
		Image:      "ghcr.io/my-org/acme-app@sha256:abc123",
		Verification: &verify.ArtifactVerification{
			SignatureValid:   true,
			AttestationValid: true,
			VerifiedIdentity: "https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main",
			VerifiedRepo:     "my-org/acme-app",
			VerifiedWorkflow: ".github/workflows/build-sign-attest.yml",
			VerifiedSubject:  "repo:my-org/acme-app",
			VerifiedDigest:   "sha256:abc123",
		},
	})
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/evaluate/artifact", bytes.NewReader(payload))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	newHandler().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("unexpected status code %d", recorder.Code)
	}

	events := readPolicyAuditEvents(t, auditPath)
	if len(events) != 1 {
		t.Fatalf("expected 1 audit event, got %d", len(events))
	}
	if events[0].EventType != audit.EventTypePolicyDecision || events[0].Decision != audit.DecisionAllow {
		t.Fatalf("unexpected audit event %#v", events[0])
	}
	if events[0].PolicyVersion != "global-artifact-policy" {
		t.Fatalf("expected policy version, got %#v", events[0])
	}
}

func TestEvaluateArtifactAllowsValidExceptionBypass(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")
	auditPath := filepath.Join(t.TempDir(), "audit.jsonl")
	previousWriter := auditWriter
	previousValidator := exceptionValidator
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	expiresAt := time.Date(2026, 4, 14, 12, 0, 0, 0, time.UTC)
	exceptionValidator = fakeExceptionValidator{
		result: audit.ExceptionValidationResult{
			Valid: true,
			Exception: &audit.PolicyException{
				ExceptionID:   "EX-2026-001",
				ExceptionType: audit.ExceptionTypeBreakGlass,
				Reason:        "CL-B0 production fix",
				TicketID:      "INC-1234",
				ApprovedBy:    "oncall@example.com",
				ExpiresAt:     expiresAt,
				Active:        true,
			},
		},
	}
	defer func() {
		auditWriter = previousWriter
		exceptionValidator = previousValidator
	}()

	payload, err := json.Marshal(policy.ArtifactEvaluationRequest{
		Tenant:      "acme",
		Repository:  "my-org/acme-app",
		Environment: "prod",
		Namespace:   "acme-prod",
		Image:       "ghcr.io/my-org/acme-app:latest",
		Exception: &policy.ExceptionIntent{
			BreakGlass:  true,
			ExceptionID: "EX-2026-001",
			Reason:      "CL-B0 production fix",
			TicketID:    "INC-1234",
		},
	})
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/evaluate/artifact", bytes.NewReader(payload))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	newHandler().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("unexpected status code %d", recorder.Code)
	}

	var decision policy.Decision
	if err := json.NewDecoder(recorder.Body).Decode(&decision); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if decision.Decision != audit.DecisionAllow || !decision.IsException || decision.ExceptionID != "EX-2026-001" {
		t.Fatalf("unexpected decision %#v", decision)
	}

	events := readPolicyAuditEvents(t, auditPath)
	if len(events) != 2 {
		t.Fatalf("expected 2 audit events, got %d", len(events))
	}
	if events[0].EventType != audit.EventTypeExceptionUsed && events[1].EventType != audit.EventTypeExceptionUsed {
		t.Fatalf("expected exception_used event, got %#v", events)
	}
}

func TestEvaluateArtifactDeniesInvalidExceptionBypass(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")
	auditPath := filepath.Join(t.TempDir(), "audit.jsonl")
	previousWriter := auditWriter
	previousValidator := exceptionValidator
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	exceptionValidator = fakeExceptionValidator{
		result: audit.ExceptionValidationResult{
			Valid:  false,
			Reason: "exception is expired",
		},
	}
	defer func() {
		auditWriter = previousWriter
		exceptionValidator = previousValidator
	}()

	payload, err := json.Marshal(policy.ArtifactEvaluationRequest{
		Tenant:      "acme",
		Repository:  "my-org/acme-app",
		Environment: "prod",
		Namespace:   "acme-prod",
		Image:       "ghcr.io/my-org/acme-app:latest",
		Exception: &policy.ExceptionIntent{
			BreakGlass:  true,
			ExceptionID: "EX-2026-002",
			Reason:      "hotfix",
			TicketID:    "INC-2000",
		},
	})
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/evaluate/artifact", bytes.NewReader(payload))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	newHandler().ServeHTTP(recorder, request)

	var decision policy.Decision
	if err := json.NewDecoder(recorder.Body).Decode(&decision); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if decision.Decision != audit.DecisionDeny {
		t.Fatalf("expected DENY, got %#v", decision)
	}
	if len(decision.Reasons) == 0 || decision.Reasons[0] != "exception is expired" {
		t.Fatalf("unexpected deny reasons %#v", decision)
	}

	events := readPolicyAuditEvents(t, auditPath)
	if len(events) != 2 {
		t.Fatalf("expected 2 audit events, got %#v", events)
	}
}

func readPolicyAuditEvents(t *testing.T, path string) []audit.Event {
	t.Helper()

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	lines := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	events := make([]audit.Event, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var event audit.Event
		if err := json.Unmarshal(line, &event); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		events = append(events, event)
	}

	return events
}

func TestValidateExceptionValidatorConfigRequiresServiceTokenWhenStaticAuthIsEnabled(t *testing.T) {
	t.Setenv("AUDIT_WRITER_URL", "http://audit-writer:8094")
	t.Setenv("CHANGELOCK_AUTH_MODE", "static-token")
	t.Setenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN", "")

	if err := validateExceptionValidatorConfig(); err == nil {
		t.Fatal("expected missing service token error")
	}

	t.Setenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN", "service-internal-demo-token")
	if err := validateExceptionValidatorConfig(); err != nil {
		t.Fatalf("expected valid config, got %v", err)
	}
}
