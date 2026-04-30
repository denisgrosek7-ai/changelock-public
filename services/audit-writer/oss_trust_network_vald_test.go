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

func TestOSSTrustNetworkValDHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/oss-trust-network/vald/status?tenant_id=acme&environment=prod",
		"/v1/oss-trust-network/vald/proofs?tenant_id=acme&environment=prod",
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

	req := httptest.NewRequest(http.MethodPost, "/v1/oss-trust-network/vald/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status route to remain read-only, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestOSSTrustNetworkValDProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/oss-trust-network/vald/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if strings.Contains(rec.Body.String(), "point_9_pass") {
		t.Fatalf("expected proofs response to omit forbidden final pass field, got %s", rec.Body.String())
	}

	var response ossTrustNetworkValDProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.OSSTrustNetworkValDStateActive ||
		response.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.DependencyState != operability.OSSTrustNetworkValDDependencyStateActive {
		t.Fatalf("expected active OSTN Val D proofs response with point 9 not complete, got %#v", response)
	}
	if response.ValCCurrentState != operability.OSSTrustNetworkValCStateActive ||
		response.ValCPoint9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.ValCDependencyState != operability.OSSTrustNetworkValCDependencyStateActive ||
		response.ValCTrustVisibilityState != operability.OSSTrustNetworkValCTrustVisibilityStateActive ||
		response.ValCPackageTrustStatusState != operability.OSSTrustNetworkValCPackageTrustStatusStateActive ||
		response.ValCExportBoundaryState != operability.OSSTrustNetworkValCExportBoundaryStateActive ||
		response.ValCRemediationSuggestionState != operability.OSSTrustNetworkValCRemediationSuggestionStateActive ||
		response.ValCPRProposalState != operability.OSSTrustNetworkValCPRProposalStateActive ||
		response.ValCLocalOverrideState != operability.OSSTrustNetworkValCLocalOverrideStateActive ||
		response.ValCRemediationSafetyState != operability.OSSTrustNetworkValCRemediationSafetyStateActive ||
		response.ValCEcosystemConsistencyState != operability.OSSTrustNetworkValCEcosystemConsistencyStateActive ||
		response.ValCNoOverclaimState != operability.OSSTrustNetworkValCNoOverclaimStateActive {
		t.Fatalf("expected exact Val C dependency closure in Val D proofs response, got %#v", response)
	}
	if response.SignalCorrectnessState != operability.OSSTrustNetworkValDSignalCorrectnessStateActive ||
		response.ReleaseFoundationState != operability.OSSTrustNetworkValDReleaseFoundationStateActive ||
		response.ReviewedIntelligenceState != operability.OSSTrustNetworkValDReviewedIntelligenceStateActive ||
		response.PropagationSafetyState != operability.OSSTrustNetworkValDPropagationSafetyStateActive ||
		response.RemediationPRSafetyState != operability.OSSTrustNetworkValDRemediationPRSafetyStateActive ||
		response.EcosystemVisibilityConsistencyState != operability.OSSTrustNetworkValDEcosystemVisibilityConsistencyStateActive ||
		response.EvidenceQualityState != operability.OSSTrustNetworkValDEvidenceQualityStateActive ||
		response.NoOverclaimState != operability.OSSTrustNetworkValDNoOverclaimStateActive {
		t.Fatalf("expected active Val D discipline states, got %#v", response)
	}
	if response.SignalLifecycleState != operability.OSSTrustNetworkValDSignalLifecycleReviewed ||
		response.ReviewState != operability.OSSTrustNetworkValBReviewStateReviewed ||
		response.ReviewerDecisionState != operability.OSSTrustNetworkValBReviewerDecisionStateAccepted ||
		response.PropagationState != operability.OSSTrustNetworkValBPropagationStateReviewedExchange ||
		response.PackageStatusClass != operability.OSSTrustNetworkValCPackageStatusReviewedSignalAvailable ||
		response.ExportClass != operability.OSSTrustNetworkValCExportClassEnterpriseCustomerView ||
		response.SuggestionClass != operability.OSSTrustNetworkValCSuggestionClassVersionUpgrade ||
		response.ProposalState != operability.OSSTrustNetworkValCProposalStateProposalReady {
		t.Fatalf("expected canonical Val D readiness metadata, got %#v", response)
	}
	if len(response.SupportedSignalLifecycleStates) != 7 ||
		len(response.SupportedSourceClasses) != 7 ||
		len(response.SupportedSourceWeightClasses) != 4 ||
		len(response.SupportedExportClasses) != 6 ||
		len(response.SupportedSuggestionClasses) != 7 ||
		len(response.SupportedProposalStates) != 5 ||
		len(response.SurfaceRefs) != len(operability.OSSTrustNetworkValDProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.OSSTrustNetworkValDProofEvidenceRefs()) ||
		len(response.WhyPoint9NotComplete) == 0 ||
		len(response.Limitations) == 0 ||
		len(response.FinalReadinessSummary) == 0 {
		t.Fatalf("expected exact proof/evidence refs and summary fields, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "oss_trust_network_vald") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

func TestOSSTrustNetworkValDDocsAndStatusStayBounded(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/oss-trust-network/vald/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if strings.Contains(rec.Body.String(), "point_9_pass") {
		t.Fatalf("expected status response to omit forbidden final pass field, got %s", rec.Body.String())
	}

	var response ossTrustNetworkValDStatusResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode status: %v", err)
	}
	if response.Model.CurrentState != operability.OSSTrustNetworkValDStateActive ||
		response.Model.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete {
		t.Fatalf("expected active OSTN Val D status model with point 9 not complete, got %#v", response)
	}

	content, err := os.ReadFile("../../docs/oss-trust-network-vald-core.md")
	if err != nil {
		t.Fatalf("read vald core doc: %v", err)
	}
	lower := strings.ToLower(string(content))
	for _, blocked := range []string{
		"certified package",
		"regulator-approved",
		"production approved",
		"public badge",
		"official oss authority",
		"point_9_pass",
	} {
		if strings.Contains(lower, blocked) {
			t.Fatalf("expected vald core doc to avoid blocked wording %q", blocked)
		}
	}
}
