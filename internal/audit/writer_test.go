package audit

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	runtimestate "github.com/denisgrosek/changelock/internal/runtime"
	"github.com/denisgrosek/changelock/internal/verify"
)

func TestWriterWritesJSONL(t *testing.T) {
	path := filepath.Join(t.TempDir(), "events.jsonl")
	writer := NewWriter(NewFileSink(path))
	writer.now = func() time.Time {
		return time.Date(2026, 4, 13, 12, 0, 0, 0, time.UTC)
	}

	summary, evidence := FromArtifactVerification(&verify.ArtifactVerification{
		SignatureValid:   true,
		AttestationValid: true,
		VerifiedIdentity: "identity",
		VerifiedIssuer:   "issuer",
		VerifiedWorkflow: ".github/workflows/build.yml",
		VerifiedDigest:   "sha256:abc123",
	})

	err := writer.Write(context.Background(), Event{
		RequestID:       "req-1",
		Component:       "deploy-gate",
		EventType:       EventTypeDeployGateDecision,
		Repo:            "my-org/acme-app",
		Image:           "ghcr.io/my-org/acme-app@sha256:abc123",
		Digest:          "sha256:abc123",
		Decision:        DecisionAllow,
		VerifierSummary: summary,
		Evidence:        evidence,
	})
	if err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(lines))
	}

	var event Event
	if err := json.Unmarshal([]byte(lines[0]), &event); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if event.RequestID != "req-1" || event.Decision != DecisionAllow {
		t.Fatalf("unexpected event %#v", event)
	}
	if event.Timestamp.IsZero() {
		t.Fatalf("expected timestamp to be set")
	}
}

func TestFromArtifactVerificationBuildsEvidence(t *testing.T) {
	summary, evidence := FromArtifactVerification(&verify.ArtifactVerification{
		SignatureValid:    false,
		AttestationValid:  true,
		VerifiedIdentity:  "identity",
		VerifiedIssuer:    "issuer",
		VerifiedSubject:   "repo:my-org/acme-app",
		VerifiedRepo:      "my-org/acme-app",
		VerifiedWorkflow:  ".github/workflows/build-sign-attest.yml",
		VerifiedRef:       "refs/heads/main",
		VerifiedCommitSHA: "abc123",
		VerifiedDigest:    "sha256:abc123",
		Evidence: verify.VerificationEvidence{
			SupplyChain: &verify.SupplyChainEvidence{
				SBOMFormat:              "spdx-json",
				SBOMDigestRef:           "ghcr.io/my-org/acme-app@sha256:abc123",
				SBOMHash:                "sha256:sbom",
				SBOMArtifactRef:         "sbom.spdx.json",
				VulnerabilityScanStatus: "passed",
				VulnerabilityScanTool:   "trivy",
				VulnerabilitySummary: &verify.VulnerabilitySummary{
					Critical: 1,
					Total:    1,
				},
				VulnerabilityReportRef: "trivy-report.json",
			},
		},
	})

	if summary == nil || evidence == nil {
		t.Fatalf("expected summary and evidence")
	}
	if evidence.Artifact == nil {
		t.Fatalf("expected artifact evidence")
	}
	if evidence.Artifact.Repository != "my-org/acme-app" || evidence.Artifact.CommitSHA != "abc123" {
		t.Fatalf("unexpected evidence %#v", evidence)
	}
	if evidence.Artifact.SBOMHash != "sha256:sbom" || evidence.Artifact.VulnerabilityReportRef != "trivy-report.json" {
		t.Fatalf("expected supply-chain evidence to be preserved, got %#v", evidence.Artifact)
	}
}

func TestFromRuntimeComparisonBuildsEvidence(t *testing.T) {
	evidence := FromRuntimeComparison(&runtimestate.ComparisonResult{
		ApprovedDigest: "sha256:approved",
		RunningDigest:  "sha256:running",
		Evidence: &runtimestate.DriftEvidence{
			ConfigExpectation: "cfg-1",
			ConfigObserved:    "cfg-2",
			ImageMismatches: []runtimestate.ImageMismatch{
				{
					Container:      "app",
					ApprovedDigest: "sha256:approved",
					RunningDigest:  "sha256:running",
				},
			},
		},
	})

	if evidence == nil || evidence.Runtime == nil {
		t.Fatalf("expected runtime evidence")
	}
	if evidence.Runtime.ApprovedDigest != "sha256:approved" || evidence.Runtime.ActualConfigHash != "cfg-2" {
		t.Fatalf("unexpected runtime evidence %#v", evidence.Runtime)
	}
}

func TestNormalizeEventAddsDecisionHash(t *testing.T) {
	event := NormalizeEvent(Event{
		RequestID:        "req-1",
		Component:        "deploy-gate",
		EventType:        EventTypePolicyDecision,
		Repo:             "my-org/acme-app",
		Environment:      "prod",
		Image:            "ghcr.io/my-org/acme-app@sha256:abc123",
		PolicyBundleHash: "sha256:bundle",
		Decision:         DecisionDeny,
	}, nil)

	if event.DecisionHash == "" {
		t.Fatal("expected decision hash to be populated")
	}
}
