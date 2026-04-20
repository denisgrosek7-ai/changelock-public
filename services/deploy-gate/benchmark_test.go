package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/verify"
)

var deployGateBenchmarkProfiles = []struct {
	name           string
	containerCount int
}{
	{name: "small", containerCount: 1},
	{name: "medium", containerCount: 10},
	{name: "large", containerCount: 100},
}

func BenchmarkAdmissionReview(b *testing.B) {
	b.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")

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
	auditWriter = audit.NewWriter(audit.NoopSink{})
	defer func() {
		artifactVerifier = previousVerifier
		auditWriter = previousWriter
	}()

	for _, profile := range deployGateBenchmarkProfiles {
		profile := profile
		payload, err := json.Marshal(buildAdmissionBenchmarkReview(profile.containerCount))
		if err != nil {
			b.Fatalf("Marshal(review) error = %v", err)
		}
		b.Run(profile.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				req := httptest.NewRequest(http.MethodPost, "/admission/review", bytes.NewReader(payload))
				rec := httptest.NewRecorder()
				admissionReviewHandler(rec, req)
				if rec.Code != http.StatusOK {
					b.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
				}
			}
		})
	}
}

func buildAdmissionBenchmarkReview(containerCount int) admissionReview {
	readOnly := true
	noPrivEsc := false
	runAsNonRoot := true

	containers := make([]container, 0, containerCount)
	for i := 0; i < containerCount; i++ {
		containers = append(containers, container{
			Name:  fmt.Sprintf("app-%03d", i),
			Image: "ghcr.io/my-org/acme-app@sha256:abc123",
			SecurityContext: &securityContext{
				ReadOnlyRootFilesystem:   &readOnly,
				AllowPrivilegeEscalation: &noPrivEsc,
			},
		})
	}

	return admissionReview{
		Request: &admissionRequest{
			UID:       "bench-admission",
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
					Containers:      containers,
				},
			},
		},
	}
}
