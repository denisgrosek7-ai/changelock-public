package audit

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

var auditBenchmarkProfiles = []struct {
	name        string
	reasons     int
	withVEX     bool
	withRuntime bool
}{
	{name: "small", reasons: 1, withVEX: false, withRuntime: false},
	{name: "medium", reasons: 4, withVEX: true, withRuntime: false},
	{name: "large", reasons: 8, withVEX: true, withRuntime: true},
}

func BenchmarkMemoryStoreIngest(b *testing.B) {
	for _, profile := range auditBenchmarkProfiles {
		profile := profile
		b.Run(profile.name, func(b *testing.B) {
			store := NewMemoryStore()
			var counter int64

			b.ReportAllocs()
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					index := atomic.AddInt64(&counter, 1)
					if _, err := store.Ingest(context.Background(), benchmarkAuditEvent(index, profile.reasons, profile.withVEX, profile.withRuntime)); err != nil {
						b.Fatalf("Ingest() error = %v", err)
					}
				}
			})
		})
	}
}

func benchmarkAuditEvent(index int64, reasonCount int, withVEX, withRuntime bool) Event {
	reasons := make([]string, 0, reasonCount)
	for i := 0; i < reasonCount; i++ {
		reasons = append(reasons, fmt.Sprintf("reason-%02d", i))
	}

	event := Event{
		RequestID:   fmt.Sprintf("bench-%d", index),
		Timestamp:   time.Date(2026, time.April, 20, 12, 0, 0, 0, time.UTC),
		Component:   "deploy-gate",
		EventType:   EventTypeDeployGateDecision,
		Decision:    DecisionDeny,
		TenantID:    "acme",
		Repo:        "my-org/acme-app",
		Environment: "prod",
		Namespace:   "acme-prod",
		Workload:    "edge-gateway",
		Image:       "ghcr.io/my-org/acme-app@sha256:abc123",
		Digest:      "sha256:abc123",
		Reasons:     reasons,
	}

	if withVEX {
		event.CVEID = "CVE-2026-9000"
		event.Evidence = &Evidence{
			Artifact: &ArtifactEvidence{
				SignerIdentity: "signer-a",
				Issuer:         "issuer-a",
				SBOMHash:       "sbom-a",
				VulnerabilitySummary: &VulnerabilitySummary{
					Critical: 1,
					Total:    3,
				},
			},
		}
	}
	if withRuntime {
		if event.Evidence == nil {
			event.Evidence = &Evidence{}
		}
		event.Evidence.Runtime = &RuntimeEvidence{
			ApprovedDigest:         "sha256:abc123",
			RunningDigest:          "sha256:def456",
			ExpectedConfigHash:     "cfg-a",
			ActualConfigHash:       "cfg-b",
			ServiceAccountExpected: "svc-a",
			ServiceAccountObserved: "svc-b",
		}
	}

	return event
}
