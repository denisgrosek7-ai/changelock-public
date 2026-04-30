package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

func TestOSSTrustNetworkValEHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/oss-trust-network/vale/closure?tenant_id=acme&environment=prod",
		"/v1/oss-trust-network/vale/proofs?tenant_id=acme&environment=prod",
	}
	for _, path := range paths {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200 for %s, got %d: %s", path, rec.Code, rec.Body.String())
		}
	}

	req := httptest.NewRequest(http.MethodPost, "/v1/oss-trust-network/vale/closure?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected closure route to remain read-only, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestOSSTrustNetworkValEProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/oss-trust-network/vale/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "point_9_pass") {
		t.Fatalf("expected Val E proofs response to expose point_9_pass fields, got %s", rec.Body.String())
	}

	var response ossTrustNetworkValEProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.OSSTrustNetworkValEStatePass ||
		response.Point9State != operability.OSSTrustNetworkPoint9StatePass ||
		!response.Point9PassAllowed ||
		response.Point9PassReason != operability.OSSTrustNetworkValEPoint9PassReasonAllowed {
		t.Fatalf("expected Val E pass proofs response, got %#v", response)
	}
	if response.DependencyState != operability.OSSTrustNetworkValEDependencyStateActive ||
		response.Val0SourceState != operability.OSSTrustNetworkValESourceStateActive ||
		response.ValASourceState != operability.OSSTrustNetworkValESourceStateActive ||
		response.ValBSourceState != operability.OSSTrustNetworkValESourceStateActive ||
		response.ValCSourceState != operability.OSSTrustNetworkValESourceStateActive ||
		response.ValDSourceState != operability.OSSTrustNetworkValESourceStateActive ||
		response.IntegratedClosureState != operability.OSSTrustNetworkValEIntegratedClosureStateActive ||
		response.CanonicalBoundaryState != operability.OSSTrustNetworkValECanonicalBoundaryStateActive ||
		response.EvidenceQualityState != operability.OSSTrustNetworkValEEvidenceQualityStateActive ||
		response.NoOverclaimState != operability.OSSTrustNetworkValENoOverclaimStateActive ||
		response.FinalPassRuleState != operability.OSSTrustNetworkValEFinalPassRuleStateActive {
		t.Fatalf("expected active Val E gate states, got %#v", response)
	}
	if response.ValDCurrentState != operability.OSSTrustNetworkValDStateActive ||
		response.ValDPoint9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.ValDDependencyState != operability.OSSTrustNetworkValDDependencyStateActive ||
		response.ValDSignalCorrectnessState != operability.OSSTrustNetworkValDSignalCorrectnessStateActive ||
		response.ValDReleaseFoundationState != operability.OSSTrustNetworkValDReleaseFoundationStateActive ||
		response.ValDReviewedIntelligenceState != operability.OSSTrustNetworkValDReviewedIntelligenceStateActive ||
		response.ValDPropagationSafetyState != operability.OSSTrustNetworkValDPropagationSafetyStateActive ||
		response.ValDRemediationPRSafetyState != operability.OSSTrustNetworkValDRemediationPRSafetyStateActive ||
		response.ValDEcosystemVisibilityConsistencyState != operability.OSSTrustNetworkValDEcosystemVisibilityConsistencyStateActive ||
		response.ValDEvidenceQualityState != operability.OSSTrustNetworkValDEvidenceQualityStateActive ||
		response.ValDNoOverclaimState != operability.OSSTrustNetworkValDNoOverclaimStateActive {
		t.Fatalf("expected canonical Val D readiness summary, got %#v", response)
	}
	if len(response.SurfaceRefs) != len(operability.OSSTrustNetworkValEProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.OSSTrustNetworkValEProofEvidenceRefs()) ||
		len(response.WhyPoint9Pass) == 0 ||
		len(response.Limitations) == 0 ||
		len(response.IntegratedClosureSummary) == 0 {
		t.Fatalf("expected exact proof/evidence refs and summary fields, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "oss_trust_network_vale") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

func TestOSSTrustNetworkValEClosureHandlerAndDocsStayBounded(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/oss-trust-network/vale/closure?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected closure 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "point_9_pass") {
		t.Fatalf("expected closure response to expose point_9_pass fields, got %s", rec.Body.String())
	}

	var response ossTrustNetworkValEClosureResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode closure: %v", err)
	}
	if response.Model.CurrentState != operability.OSSTrustNetworkValEStatePass ||
		response.Model.Point9State != operability.OSSTrustNetworkPoint9StatePass ||
		!response.Model.Point9PassAllowed {
		t.Fatalf("expected active Val E closure model, got %#v", response.Model)
	}
	if response.Model.Val0Source.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.Model.ValASource.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.Model.ValBSource.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.Model.ValCSource.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.Model.ValDSource.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete {
		t.Fatalf("expected only Val E to promote point 9 pass, got %#v", response.Model)
	}

	content, err := os.ReadFile("../../docs/oss-trust-network-vale-core.md")
	if err != nil {
		t.Fatalf("read vale core doc: %v", err)
	}
	lower := strings.ToLower(string(content))
	for _, blocked := range []string{
		"certified package",
		"regulator-approved",
		"production approved",
		"deployment approved",
		"public badge",
		"official oss authority",
		"global truth",
	} {
		if strings.Contains(lower, blocked) {
			t.Fatalf("expected vale core doc to avoid blocked wording %q", blocked)
		}
	}
	if !strings.Contains(lower, "point_9_pass") {
		t.Fatalf("expected vale core doc to state that Val E may emit point_9_pass")
	}
}
