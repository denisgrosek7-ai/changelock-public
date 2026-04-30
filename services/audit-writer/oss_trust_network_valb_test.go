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

func TestOSSTrustNetworkValBHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/oss-trust-network/valb/status?tenant_id=acme&environment=prod",
		"/v1/oss-trust-network/valb/proofs?tenant_id=acme&environment=prod",
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

	req := httptest.NewRequest(http.MethodPost, "/v1/oss-trust-network/valb/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status route to remain read-only, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestOSSTrustNetworkValBProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/oss-trust-network/valb/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if strings.Contains(rec.Body.String(), "point_9_pass") {
		t.Fatalf("expected proofs response to omit forbidden final pass field, got %s", rec.Body.String())
	}

	var response ossTrustNetworkValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.OSSTrustNetworkValBStateActive ||
		response.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.DependencyState != operability.OSSTrustNetworkValBDependencyStateActive {
		t.Fatalf("expected active OSTN Val B proofs response with point 9 not complete, got %#v", response)
	}
	if response.ValACurrentState != operability.OSSTrustNetworkValAStateActive ||
		response.ValAPoint9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.ValADependencyState != operability.OSSTrustNetworkValADependencyStateActive ||
		response.ValAReleaseTrustIntakeState != operability.OSSTrustNetworkValAReleaseTrustIntakeStateActive ||
		response.ValASigningSignalState != operability.OSSTrustNetworkValASigningSignalStateActive ||
		response.ValAMaintainerState != operability.OSSTrustNetworkValAMaintainerAttestationStateActive ||
		response.ValAProvenanceState != operability.OSSTrustNetworkValAProvenanceMaterialStateActive ||
		response.ValARegistryDescriptorState != operability.OSSTrustNetworkValARegistryDescriptorStateActive ||
		response.ValARegistryMetadataState != operability.OSSTrustNetworkValARegistryMetadataStateActive ||
		response.ValATypoWarningState != operability.OSSTrustNetworkValATypoSquattingWarningStateActive ||
		response.ValADriftSignalState != operability.OSSTrustNetworkValADriftSignalStateActive ||
		response.ValANoOverclaimState != operability.OSSTrustNetworkValANoOverclaimStateActive {
		t.Fatalf("expected exact Val A dependency closure in Val B proofs response, got %#v", response)
	}
	if response.CandidateSignalIntakeState != operability.OSSTrustNetworkValBCandidateSignalIntakeStateActive ||
		response.ReviewWorkflowState != operability.OSSTrustNetworkValBReviewWorkflowStateActive ||
		response.SharedVEXTriageState != operability.OSSTrustNetworkValBSharedVEXTriageStateActive ||
		response.SourceWeightingState != operability.OSSTrustNetworkValBSourceWeightingStateActive ||
		response.LocalApplicabilityState != operability.OSSTrustNetworkValBLocalApplicabilityStateActive ||
		response.PropagationExchangeState != operability.OSSTrustNetworkValBPropagationExchangeStateActive ||
		response.SupersessionRevocationState != operability.OSSTrustNetworkValBSupersessionRevocationStateActive ||
		response.ReviewerAuditabilityState != operability.OSSTrustNetworkValBReviewerAuditabilityStateActive ||
		response.NoOverclaimState != operability.OSSTrustNetworkValBNoOverclaimStateActive {
		t.Fatalf("expected active Val B discipline states, got %#v", response)
	}
	if response.CandidateIntakeState != operability.OSSTrustNetworkValBCandidateIntakeStateNormalized ||
		response.CandidateSourceClass != operability.OSSTrustNetworkValBCandidateSourceClassVerifier ||
		response.CandidateFreshness != operability.IntelligenceCalibrationFreshnessFresh ||
		response.ReviewState != operability.OSSTrustNetworkValBReviewStateReviewed ||
		response.ReviewerDecisionState != operability.OSSTrustNetworkValBReviewerDecisionStateAccepted ||
		response.SharedVEXState != operability.OSSTrustNetworkValBSharedVEXStateReviewed ||
		response.SourceClass != operability.OSSTrustNetworkValBCandidateSourceClassVerifier ||
		response.SourceWeightClass != operability.OSSTrustNetworkValBSourceWeightClassMedium ||
		response.LocalApplicabilityStatus != operability.OSSTrustNetworkValBLocalApplicabilityStatusApplicable ||
		response.PropagationState != operability.OSSTrustNetworkValBPropagationStateReviewedExchange ||
		response.LifecycleState != operability.OSSTrustNetworkValBLifecycleStateActive {
		t.Fatalf("expected canonical Val B signal metadata, got %#v", response)
	}
	if len(response.SupportedCandidateSources) != 7 ||
		len(response.SupportedCandidateIntakeStates) != 7 ||
		len(response.SupportedReviewStates) != 6 ||
		len(response.SupportedReviewerDecisionStates) != 6 ||
		len(response.SupportedSharedVEXStates) != 7 ||
		len(response.SupportedSourceWeightClasses) != 4 ||
		len(response.SupportedLocalApplicabilityStates) != 5 ||
		len(response.SupportedPropagationStates) != 8 ||
		len(response.SurfaceRefs) != len(operability.OSSTrustNetworkValBProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.OSSTrustNetworkValBProofEvidenceRefs()) ||
		len(response.WhyPoint9NotComplete) == 0 ||
		len(response.Limitations) == 0 ||
		len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact proof/evidence refs and summary fields, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "oss_trust_network_valb") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

func TestOSSTrustNetworkValBDocsAndStatusStayBounded(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/oss-trust-network/valb/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if strings.Contains(rec.Body.String(), "point_9_pass") {
		t.Fatalf("expected status response to omit forbidden final pass field, got %s", rec.Body.String())
	}

	var response ossTrustNetworkValBStatusResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode status: %v", err)
	}
	if response.Model.CurrentState != operability.OSSTrustNetworkValBStateActive ||
		response.Model.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete {
		t.Fatalf("expected active OSTN Val B status model with point 9 not complete, got %#v", response)
	}

	content, err := os.ReadFile("../../docs/oss-trust-network-valb-core.md")
	if err != nil {
		t.Fatalf("read valb core doc: %v", err)
	}
	lower := strings.ToLower(string(content))
	for _, blocked := range []string{
		"de-facto standard",
		"immune system for open source",
		"certified package",
		"regulator-approved",
		"production approved",
		"audit passed",
		"point_9_pass",
	} {
		if strings.Contains(lower, blocked) {
			t.Fatalf("expected valb core doc to avoid blocked wording %q", blocked)
		}
	}
}
