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

func TestArtifactHandlerWritesAllowAuditEvent(t *testing.T) {
	auditPath := filepath.Join(t.TempDir(), "audit.jsonl")
	previousVerifier := artifactVerifier
	previousWriter := auditWriter
	artifactVerifier = fakeArtifactVerifier{
		result: verify.ArtifactVerification{
			SignatureValid:    true,
			AttestationValid:  true,
			VerifiedRepo:      "my-org/acme-app",
			VerifiedWorkflow:  ".github/workflows/build-sign-attest.yml",
			VerifiedRef:       "refs/heads/main",
			VerifiedCommitSHA: "abc123",
			VerifiedDigest:    "sha256:abc123",
			VerifiedIdentity:  "https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main",
			VerifiedIssuer:    "https://token.actions.githubusercontent.com",
		},
	}
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	defer func() {
		artifactVerifier = previousVerifier
		auditWriter = previousWriter
	}()

	payload, err := json.Marshal(verify.ArtifactVerificationRequest{
		Image:                   "ghcr.io/my-org/acme-app@sha256:abc123",
		ExpectedRepository:      "my-org/acme-app",
		ExpectedRef:             "refs/heads/main",
		ExpectedCommitSHA:       "abc123",
		AllowedSignerIdentities: []string{"https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main"},
		AllowedOIDCIssuers:      []string{"https://token.actions.githubusercontent.com"},
	})
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	request := newJSONRequest(t, "POST", "/verify/artifact", payload)
	recorder := newRecorder()
	newHandler().ServeHTTP(recorder, request)

	events := readAuditEvents(t, auditPath)
	if len(events) != 1 {
		t.Fatalf("expected 1 audit event, got %d", len(events))
	}
	if events[0].Decision != audit.DecisionAllow || events[0].EventType != audit.EventTypeArtifactVerificationResult {
		t.Fatalf("unexpected audit event %#v", events[0])
	}
}

func TestArtifactHandlerWritesDenyAuditEvent(t *testing.T) {
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

	payload, err := json.Marshal(verify.ArtifactVerificationRequest{
		Image: "ghcr.io/my-org/acme-app@sha256:abc123",
	})
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	request := newJSONRequest(t, "POST", "/verify/artifact", payload)
	recorder := newRecorder()
	newHandler().ServeHTTP(recorder, request)

	events := readAuditEvents(t, auditPath)
	if len(events) != 1 {
		t.Fatalf("expected 1 audit event, got %d", len(events))
	}
	if events[0].Decision != audit.DecisionDeny {
		t.Fatalf("expected DENY event, got %#v", events[0])
	}
	if len(events[0].Reasons) == 0 {
		t.Fatalf("expected deny reasons in audit event")
	}
}

func newJSONRequest(t *testing.T, method, path string, payload []byte) *http.Request {
	t.Helper()
	request, err := http.NewRequest(method, path, bytes.NewReader(payload))
	if err != nil {
		t.Fatalf("http.NewRequest() error = %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	return request
}

func newRecorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
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
