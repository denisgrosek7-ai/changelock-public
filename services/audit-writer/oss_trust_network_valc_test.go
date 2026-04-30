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

func TestOSSTrustNetworkValCHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/oss-trust-network/valc/status?tenant_id=acme&environment=prod",
		"/v1/oss-trust-network/valc/proofs?tenant_id=acme&environment=prod",
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

	req := httptest.NewRequest(http.MethodPost, "/v1/oss-trust-network/valc/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status route to remain read-only, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestOSSTrustNetworkValCProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/oss-trust-network/valc/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if strings.Contains(rec.Body.String(), "point_9_pass") {
		t.Fatalf("expected proofs response to omit forbidden final pass field, got %s", rec.Body.String())
	}

	var response ossTrustNetworkValCProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.OSSTrustNetworkValCStateActive ||
		response.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.DependencyState != operability.OSSTrustNetworkValCDependencyStateActive {
		t.Fatalf("expected active OSTN Val C proofs response with point 9 not complete, got %#v", response)
	}
	if response.ValBCurrentState != operability.OSSTrustNetworkValBStateActive ||
		response.ValBPoint9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.ValBDependencyState != operability.OSSTrustNetworkValBDependencyStateActive ||
		response.ValBCandidateSignalIntakeState != operability.OSSTrustNetworkValBCandidateSignalIntakeStateActive ||
		response.ValBReviewWorkflowState != operability.OSSTrustNetworkValBReviewWorkflowStateActive ||
		response.ValBSharedVEXTriageState != operability.OSSTrustNetworkValBSharedVEXTriageStateActive ||
		response.ValBSourceWeightingState != operability.OSSTrustNetworkValBSourceWeightingStateActive ||
		response.ValBLocalApplicabilityState != operability.OSSTrustNetworkValBLocalApplicabilityStateActive ||
		response.ValBPropagationExchangeState != operability.OSSTrustNetworkValBPropagationExchangeStateActive ||
		response.ValBSupersessionRevocationState != operability.OSSTrustNetworkValBSupersessionRevocationStateActive ||
		response.ValBReviewerAuditabilityState != operability.OSSTrustNetworkValBReviewerAuditabilityStateActive ||
		response.ValBNoOverclaimState != operability.OSSTrustNetworkValBNoOverclaimStateActive {
		t.Fatalf("expected exact Val B dependency closure in Val C proofs response, got %#v", response)
	}
	if response.TrustVisibilityState != operability.OSSTrustNetworkValCTrustVisibilityStateActive ||
		response.PackageTrustStatusState != operability.OSSTrustNetworkValCPackageTrustStatusStateActive ||
		response.ExportBoundaryState != operability.OSSTrustNetworkValCExportBoundaryStateActive ||
		response.RemediationSuggestionState != operability.OSSTrustNetworkValCRemediationSuggestionStateActive ||
		response.PRProposalState != operability.OSSTrustNetworkValCPRProposalStateActive ||
		response.LocalOverrideState != operability.OSSTrustNetworkValCLocalOverrideStateActive ||
		response.RemediationSafetyState != operability.OSSTrustNetworkValCRemediationSafetyStateActive ||
		response.EcosystemConsistencyState != operability.OSSTrustNetworkValCEcosystemConsistencyStateActive ||
		response.NoOverclaimState != operability.OSSTrustNetworkValCNoOverclaimStateActive {
		t.Fatalf("expected active Val C discipline states, got %#v", response)
	}
	if response.VisibilityState != operability.OSSTrustNetworkValCVisibilityVisible ||
		response.PackageStatusClass != operability.OSSTrustNetworkValCPackageStatusReviewedSignalAvailable ||
		response.ExportClass != operability.OSSTrustNetworkValCExportClassEnterpriseCustomerView ||
		response.SuggestionClass != operability.OSSTrustNetworkValCSuggestionClassVersionUpgrade ||
		response.SuggestionConfidenceClass != operability.OSSTrustNetworkConfidenceBounded ||
		response.ProposalState != operability.OSSTrustNetworkValCProposalStateProposalReady ||
		response.LocalOverrideVisibilityState != operability.OSSTrustNetworkValCOverrideStateNoOverride ||
		response.RemediationRiskClass != operability.OSSTrustNetworkValCRiskClassMedium {
		t.Fatalf("expected canonical Val C signal metadata, got %#v", response)
	}
	if len(response.SupportedVisibilityStates) != 6 ||
		len(response.SupportedPackageStatusClasses) != 7 ||
		len(response.SupportedExportClasses) != 6 ||
		len(response.SupportedSuggestionClasses) != 7 ||
		len(response.SupportedProposalStates) != 5 ||
		len(response.SupportedLocalOverrideStates) != 6 ||
		len(response.SupportedRiskClasses) != 3 ||
		len(response.SurfaceRefs) != len(operability.OSSTrustNetworkValCProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.OSSTrustNetworkValCProofEvidenceRefs()) ||
		len(response.WhyPoint9NotComplete) == 0 ||
		len(response.Limitations) == 0 ||
		len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact proof/evidence refs and summary fields, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "oss_trust_network_valc") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

func TestOSSTrustNetworkValCDocsAndStatusStayBounded(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/oss-trust-network/valc/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if strings.Contains(rec.Body.String(), "point_9_pass") {
		t.Fatalf("expected status response to omit forbidden final pass field, got %s", rec.Body.String())
	}

	var response ossTrustNetworkValCStatusResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode status: %v", err)
	}
	if response.Model.CurrentState != operability.OSSTrustNetworkValCStateActive ||
		response.Model.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete {
		t.Fatalf("expected active OSTN Val C status model with point 9 not complete, got %#v", response)
	}

	content, err := os.ReadFile("../../docs/oss-trust-network-valc-core.md")
	if err != nil {
		t.Fatalf("read valc core doc: %v", err)
	}
	lower := strings.ToLower(string(content))
	for _, blocked := range []string{
		"de-facto standard",
		"immune system for open source",
		"certified package",
		"regulator-approved",
		"production approved",
		"audit passed",
		"auto-remediated",
		"auto-merged",
		"point_9_pass",
	} {
		if strings.Contains(lower, blocked) {
			t.Fatalf("expected valc core doc to avoid blocked wording %q", blocked)
		}
	}
}
