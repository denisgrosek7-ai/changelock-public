package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/policy"
	"github.com/denisgrosek/changelock/internal/verify"
)

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
