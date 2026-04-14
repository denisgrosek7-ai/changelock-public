package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/verify"
)

type fakeArtifactVerifier struct {
	result verify.ArtifactVerification
	err    error
}

func (f fakeArtifactVerifier) VerifyArtifact(_ context.Context, _ verify.ArtifactVerificationRequest) (verify.ArtifactVerification, error) {
	return f.result, f.err
}

type fakeExceptionValidator struct {
	result audit.ExceptionValidationResult
	err    error
}

func (f fakeExceptionValidator) Validate(_ context.Context, _ audit.ExceptionValidationRequest) (audit.ExceptionValidationResult, error) {
	return f.result, f.err
}

func TestAdmissionReviewAllowsTrustedWorkload(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")
	auditPath := filepath.Join(t.TempDir(), "audit.jsonl")
	previousVerifier := artifactVerifier
	previousWriter := auditWriter
	artifactVerifier = fakeArtifactVerifier{
		result: verify.ArtifactVerification{
			SignatureValid:   true,
			AttestationValid: true,
			VerifiedIdentity: "https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main",
			VerifiedRepo:     "my-org/acme-app",
			VerifiedWorkflow: ".github/workflows/build-sign-attest.yml",
			VerifiedSubject:  "repo:my-org/acme-app",
			VerifiedDigest:   "sha256:abc123",
		},
	}
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	defer func() {
		artifactVerifier = previousVerifier
		auditWriter = previousWriter
	}()

	readOnly := true
	noPrivEsc := false
	runAsNonRoot := true

	review := admissionReview{
		Request: &admissionRequest{
			UID:       "allow-1",
			Namespace: "acme-prod",
			Kind:      objectReference{Kind: "Pod"},
			Object: pod{
				Metadata: objectMeta{
					Annotations: map[string]string{
						"changelock.io/tenant":       "acme",
						"changelock.io/repository":   "my-org/acme-app",
						"changelock.io/subject":      "repo:my-org/acme-app",
						"changelock.io/workflow-sha": "abc123",
					},
				},
				Spec: podSpec{
					SecurityContext: &podSecurityContext{RunAsNonRoot: &runAsNonRoot},
					Containers: []container{
						{
							Name:  "app",
							Image: "ghcr.io/my-org/acme-app@sha256:abc123",
							SecurityContext: &securityContext{
								ReadOnlyRootFilesystem:   &readOnly,
								AllowPrivilegeEscalation: &noPrivEsc,
							},
						},
					},
				},
			},
		},
	}

	response := executeAdmissionRequest(t, review)
	if !response.Response.Allowed {
		t.Fatalf("expected admission to allow, got %#v", response.Response)
	}

	events := readAuditEvents(t, auditPath)
	if !hasDecisionEvent(events, audit.EventTypePolicyDecision, audit.DecisionAllow) {
		t.Fatalf("expected ALLOW policy decision event, got %#v", events)
	}
	if !hasDecisionEvent(events, audit.EventTypeDeployGateDecision, audit.DecisionAllow) {
		t.Fatalf("expected ALLOW deploy gate event, got %#v", events)
	}
}

func TestAdmissionReviewDeniesMutableAndPrivilegedWorkload(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")
	auditPath := filepath.Join(t.TempDir(), "audit.jsonl")
	previousVerifier := artifactVerifier
	previousWriter := auditWriter
	artifactVerifier = fakeArtifactVerifier{
		result: verify.ArtifactVerification{
			SignatureValid:   false,
			AttestationValid: false,
			Reasons:          []string{"signature verification failed", "attestation verification failed"},
		},
	}
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	defer func() {
		artifactVerifier = previousVerifier
		auditWriter = previousWriter
	}()

	review := admissionReview{
		Request: &admissionRequest{
			UID:       "deny-1",
			Namespace: "acme-prod",
			Kind:      objectReference{Kind: "Pod"},
			Object: pod{
				Metadata: objectMeta{
					Annotations: map[string]string{
						"changelock.io/tenant":     "acme",
						"changelock.io/repository": "my-org/acme-app",
						"changelock.io/subject":    "repo:my-org/acme-app",
					},
				},
				Spec: podSpec{
					Containers: []container{
						{
							Name:  "app",
							Image: "ghcr.io/my-org/acme-app:latest",
						},
					},
				},
			},
		},
	}

	response := executeAdmissionRequest(t, review)
	if response.Response.Allowed {
		t.Fatalf("expected admission to deny, got %#v", response.Response)
	}
	if response.Response.Status == nil || response.Response.Status.Message == "" {
		t.Fatalf("expected denial message")
	}

	events := readAuditEvents(t, auditPath)
	if !hasDecisionEvent(events, audit.EventTypePolicyDecision, audit.DecisionDeny) {
		t.Fatalf("expected DENY policy decision event, got %#v", events)
	}
	deployEvent := findDecisionEvent(events, audit.EventTypeDeployGateDecision, audit.DecisionDeny)
	if deployEvent == nil || len(deployEvent.Reasons) == 0 {
		t.Fatalf("expected explainable DENY deploy gate event, got %#v", events)
	}
}

func TestAdmissionReviewAllowsValidBreakGlassException(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")
	auditPath := filepath.Join(t.TempDir(), "audit.jsonl")
	previousVerifier := artifactVerifier
	previousWriter := auditWriter
	previousValidator := exceptionValidator
	artifactVerifier = fakeArtifactVerifier{}
	expiresAt := time.Date(2026, 4, 14, 12, 0, 0, 0, time.UTC)
	exceptionValidator = fakeExceptionValidator{
		result: audit.ExceptionValidationResult{
			Valid: true,
			Exception: &audit.PolicyException{
				ExceptionID:   "EX-2026-001",
				ExceptionType: audit.ExceptionTypeBreakGlass,
				Reason:        "P0 production fix",
				TicketID:      "INC-1234",
				ApprovedBy:    "oncall@example.com",
				ExpiresAt:     expiresAt,
				Active:        true,
			},
		},
	}
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	defer func() {
		artifactVerifier = previousVerifier
		auditWriter = previousWriter
		exceptionValidator = previousValidator
	}()

	review := admissionReview{
		Request: &admissionRequest{
			UID:       "allow-break-glass",
			Namespace: "acme-prod",
			Kind:      objectReference{Kind: "Pod"},
			Object: pod{
				Metadata: objectMeta{
					Name: "booking-api",
					Annotations: map[string]string{
						"changelock.io/tenant":       "acme",
						"changelock.io/repository":   "my-org/acme-app",
						"changelock.io/break-glass":  "true",
						"changelock.io/exception-id": "EX-2026-001",
						"changelock.io/reason":       "P0 production fix",
						"changelock.io/ticket-id":    "INC-1234",
						"changelock.io/environment":  "prod",
					},
				},
				Spec: podSpec{
					Containers: []container{{
						Name:  "app",
						Image: "ghcr.io/my-org/acme-app:latest",
					}},
				},
			},
		},
	}

	response := executeAdmissionRequest(t, review)
	if !response.Response.Allowed {
		t.Fatalf("expected break-glass admission to allow, got %#v", response.Response)
	}

	events := readAuditEvents(t, auditPath)
	if !hasDecisionEvent(events, audit.EventTypeExceptionUsed, audit.DecisionAllow) {
		t.Fatalf("expected exception_used event, got %#v", events)
	}
	if !hasDecisionEvent(events, audit.EventTypeDeployGateDecision, audit.DecisionAllow) {
		t.Fatalf("expected ALLOW deploy gate event, got %#v", events)
	}
}

func TestAdmissionReviewDeniesInvalidBreakGlassException(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")
	auditPath := filepath.Join(t.TempDir(), "audit.jsonl")
	previousVerifier := artifactVerifier
	previousWriter := auditWriter
	previousValidator := exceptionValidator
	artifactVerifier = fakeArtifactVerifier{}
	exceptionValidator = fakeExceptionValidator{
		result: audit.ExceptionValidationResult{
			Valid:  false,
			Reason: "exception is expired",
		},
	}
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	defer func() {
		artifactVerifier = previousVerifier
		auditWriter = previousWriter
		exceptionValidator = previousValidator
	}()

	review := admissionReview{
		Request: &admissionRequest{
			UID:       "deny-break-glass",
			Namespace: "acme-prod",
			Kind:      objectReference{Kind: "Pod"},
			Object: pod{
				Metadata: objectMeta{
					Name: "booking-api",
					Annotations: map[string]string{
						"changelock.io/tenant":       "acme",
						"changelock.io/repository":   "my-org/acme-app",
						"changelock.io/break-glass":  "true",
						"changelock.io/exception-id": "EX-2026-002",
						"changelock.io/reason":       "hotfix",
						"changelock.io/ticket-id":    "INC-2000",
						"changelock.io/environment":  "prod",
					},
				},
				Spec: podSpec{
					Containers: []container{{
						Name:  "app",
						Image: "ghcr.io/my-org/acme-app:latest",
					}},
				},
			},
		},
	}

	response := executeAdmissionRequest(t, review)
	if response.Response.Allowed {
		t.Fatalf("expected invalid break-glass admission to deny, got %#v", response.Response)
	}
	if response.Response.Status == nil || !strings.Contains(response.Response.Status.Message, "exception is expired") {
		t.Fatalf("expected exception deny message, got %#v", response.Response)
	}

	events := readAuditEvents(t, auditPath)
	if !hasDecisionEvent(events, audit.EventTypeExceptionValidationFailed, audit.DecisionDeny) {
		t.Fatalf("expected exception_validation_failed event, got %#v", events)
	}
}

func executeAdmissionRequest(t *testing.T, review admissionReview) admissionReview {
	t.Helper()

	payload, err := json.Marshal(review)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/admission/review", bytes.NewReader(payload))
	recorder := httptest.NewRecorder()

	newHandler().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("unexpected status code %d", recorder.Code)
	}

	var response admissionReview
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	return response
}

func readAuditEvents(t *testing.T, path string) []audit.Event {
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

func hasDecisionEvent(events []audit.Event, eventType, decision string) bool {
	return findDecisionEvent(events, eventType, decision) != nil
}

func findDecisionEvent(events []audit.Event, eventType, decision string) *audit.Event {
	for _, event := range events {
		if event.EventType == eventType && event.Decision == decision {
			eventCopy := event
			return &eventCopy
		}
	}
	return nil
}
