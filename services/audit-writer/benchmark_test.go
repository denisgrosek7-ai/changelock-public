package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
)

var auditWriterBenchmarkProfiles = []struct {
	name                string
	backgroundWorkloads int
}{
	{name: "small", backgroundWorkloads: 10},
	{name: "medium", backgroundWorkloads: 100},
	{name: "large", backgroundWorkloads: 1000},
}

func BenchmarkAuditWriterTopologyBlastRadius(b *testing.B) {
	for _, profile := range auditWriterBenchmarkProfiles {
		profile := profile
		fixture := benchmarkAuditWriterFixture(b, profile.backgroundWorkloads)
		req := httptest.NewRequest(http.MethodGet, "/v1/topology/blast-radius?tenant_id=acme&environment=prod&service=edge-gateway", nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")

		b.Run(profile.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				rec := httptest.NewRecorder()
				fixture.handler.ServeHTTP(rec, req)
				if rec.Code != http.StatusOK {
					b.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
				}
			}
		})
	}
}

func BenchmarkAuditWriterForensicsState(b *testing.B) {
	for _, profile := range auditWriterBenchmarkProfiles {
		profile := profile
		fixture := benchmarkAuditWriterFixture(b, profile.backgroundWorkloads)
		req := httptest.NewRequest(http.MethodGet, "/v1/forensics/state?tenant_id=acme&environment=prod&service=edge-gateway&timestamp="+fixture.historicalTimestamp.Format(time.RFC3339), nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")

		b.Run(profile.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				rec := httptest.NewRecorder()
				fixture.handler.ServeHTTP(rec, req)
				if rec.Code != http.StatusOK {
					b.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
				}
			}
		})
	}
}

func BenchmarkAuditWriterRuntimeFindings(b *testing.B) {
	for _, profile := range auditWriterBenchmarkProfiles {
		profile := profile
		fixture := benchmarkAuditWriterFixture(b, profile.backgroundWorkloads)
		req := httptest.NewRequest(http.MethodGet, "/v1/runtime/findings?tenant_id=acme&environment=prod", nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")

		b.Run(profile.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				rec := httptest.NewRecorder()
				fixture.handler.ServeHTTP(rec, req)
				if rec.Code != http.StatusOK {
					b.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
				}
			}
		})
	}
}

func BenchmarkAuditWriterHandoffSeal(b *testing.B) {
	for _, profile := range auditWriterBenchmarkProfiles {
		profile := profile
		b.Run(profile.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				fixture := benchmarkAuditWriterFixture(b, profile.backgroundWorkloads)
				req := httptest.NewRequest(http.MethodPost, "/v1/handoff/seal?tenant_id=acme&environment=prod", bytes.NewBufferString(`{"audience":"auditor_safe","include_forensics":true,"co_sign_mode":"system_only"}`))
				req.Header.Set("Authorization", "Bearer operator-demo-token")
				req.Header.Set("Content-Type", "application/json")
				rec := httptest.NewRecorder()
				b.StartTimer()
				fixture.handler.ServeHTTP(rec, req)
				b.StopTimer()
				if rec.Code != http.StatusOK {
					b.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
				}
			}
		})
	}
}

func BenchmarkAuditWriterHandoffVerify(b *testing.B) {
	for _, profile := range auditWriterBenchmarkProfiles {
		profile := profile
		fixture := benchmarkAuditWriterFixture(b, profile.backgroundWorkloads)
		sealed := sealFederationHandoffForTest(b, fixture.handler, incidentAudienceAuditorSafe)
		bundle := downloadHandoffBundleForTest(b, fixture.handler, sealed.PackageID)
		reqBody := []byte(`{"bundle_base64":"` + base64.StdEncoding.EncodeToString(bundle) + `"}`)

		b.Run(profile.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				req := httptest.NewRequest(http.MethodPost, "/v1/handoff/verify", bytes.NewReader(reqBody))
				req.Header.Set("Authorization", "Bearer viewer-demo-token")
				req.Header.Set("Content-Type", "application/json")
				rec := httptest.NewRecorder()
				fixture.handler.ServeHTTP(rec, req)
				if rec.Code != http.StatusOK {
					b.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
				}
			}
		})
	}
}

func BenchmarkAuditWriterFederationProofVerify(b *testing.B) {
	for _, profile := range auditWriterBenchmarkProfiles {
		profile := profile
		b.Run(profile.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				fixture := benchmarkAuditWriterFixture(b, profile.backgroundWorkloads)
				peer := registerFederationPeerForTest(b, fixture.handler, federationPeerRequest{
					PeerID:            "peer-bench",
					Organization:      "Bench Supplier",
					Region:            "us-east",
					Cluster:           "cluster-bench",
					TrustDomain:       "bench.example",
					Endpoint:          "https://bench.example.invalid",
					PublicKeys:        []string{"pub-bench-1"},
					Capabilities:      []string{"sealed_handoff"},
					PolicyRole:        federationPolicyRoleSupplier,
					AcceptedAudiences: []string{incidentAudienceAuditorSafe},
					LastSeen:          timePointer(time.Date(2026, time.April, 20, 10, 0, 0, 0, time.UTC)),
				})
				sealed := sealFederationHandoffForTest(b, fixture.handler, incidentAudienceAuditorSafe)
				bundle := downloadHandoffBundleForTest(b, fixture.handler, sealed.PackageID)
				req := httptest.NewRequest(http.MethodPost, "/v1/federation/proof-verify", bytes.NewBufferString(`{"peer_id":"`+peer.PeerID+`","bundle_base64":"`+base64.StdEncoding.EncodeToString(bundle)+`","requested_scope":{"tenant_id":"acme","environment":"prod","audience":"auditor_safe"}}`))
				req.Header.Set("Authorization", "Bearer operator-demo-token")
				req.Header.Set("Content-Type", "application/json")
				rec := httptest.NewRecorder()
				b.StartTimer()
				fixture.handler.ServeHTTP(rec, req)
				b.StopTimer()
				if rec.Code != http.StatusOK {
					b.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
				}
			}
		})
	}
}

func BenchmarkAuditWriterValidationExecute(b *testing.B) {
	for _, profile := range auditWriterBenchmarkProfiles {
		profile := profile
		b.Run(profile.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				fixture := benchmarkAuditWriterFixture(b, profile.backgroundWorkloads)
				req := httptest.NewRequest(http.MethodPost, "/v1/validation/execute?tenant_id=acme&environment=prod", bytes.NewBufferString(`{"mode":"policy_dry_run"}`))
				req.Header.Set("Authorization", "Bearer operator-demo-token")
				req.Header.Set("Content-Type", "application/json")
				rec := httptest.NewRecorder()
				b.StartTimer()
				fixture.handler.ServeHTTP(rec, req)
				b.StopTimer()
				if rec.Code != http.StatusOK {
					b.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
				}
			}
		})
	}
}

func benchmarkAuditWriterFixture(tb testing.TB, backgroundWorkloads int) forensicsFixtureData {
	tb.Helper()
	tb.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")
	tb.Setenv("CHANGELOCK_FEDERATION_SIGNING_SEED", "federation-seed")
	fixture := forensicsTestFixture(tb)
	seedBenchmarkBackgroundEvents(tb, fixture.store, backgroundWorkloads)
	return fixture
}

func seedBenchmarkBackgroundEvents(tb testing.TB, store audit.Store, workloads int) {
	tb.Helper()
	base := time.Date(2026, time.April, 20, 8, 0, 0, 0, time.UTC)
	for i := 0; i < workloads; i++ {
		workload := fmt.Sprintf("bench-workload-%04d", i)
		digest := fmt.Sprintf("sha256:bench-%04d", i)
		events := []audit.Event{
			{
				RequestID:      fmt.Sprintf("bench-deploy-%04d", i),
				Timestamp:      base.Add(time.Duration(i) * time.Second),
				Component:      "deploy-gate",
				EventType:      audit.EventTypeDeployGateDecision,
				Decision:       audit.DecisionAllow,
				TenantID:       "acme",
				Repo:           "acme/bench",
				Environment:    "prod",
				Namespace:      "acme-prod",
				Workload:       workload,
				ServiceAccount: "bench-sa",
				Digest:         digest,
				Reasons:        []string{"benchmark background deploy"},
			},
			{
				RequestID:                fmt.Sprintf("bench-runtime-%04d", i),
				Timestamp:                base.Add(time.Duration(i) * time.Second).Add(500 * time.Millisecond),
				Component:                "runtime-agent",
				EventType:                audit.EventTypeRuntimeActiveStateObserved,
				Decision:                 audit.DecisionAllow,
				TenantID:                 "acme",
				ClusterID:                "local",
				Repo:                     "acme/bench",
				Environment:              "prod",
				Namespace:                "acme-prod",
				WorkloadKind:             "Deployment",
				Workload:                 workload,
				ServiceAccount:           "bench-sa",
				Digest:                   digest,
				ReconciliationStatus:     "in_sync",
				DesiredStateSourceRef:    "deploy:" + workload,
				DesiredStateApprovalID:   "bench-approval",
				DesiredStateVerification: "verified",
				Reasons:                  []string{"benchmark background runtime"},
			},
		}
		for _, event := range events {
			if _, err := store.Ingest(context.Background(), event); err != nil {
				tb.Fatalf("Ingest() error = %v", err)
			}
		}
	}
}
