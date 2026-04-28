package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

func TestDeveloperEcosystemValBHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/developer-ecosystem/valb/status?tenant_id=acme&environment=prod",
		"/v1/developer-ecosystem/valb/proofs?tenant_id=acme&environment=prod",
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

	req := httptest.NewRequest(http.MethodPost, "/v1/developer-ecosystem/valb/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status route to remain read-only, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestDeveloperEcosystemValBProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/developer-ecosystem/valb/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response developerEcosystemValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.DeveloperEcosystemValBStateActive ||
		response.ValECompatibilityState != operability.DeveloperEcosystemValBValECompatibilityStateActive ||
		response.DependencyState != operability.DeveloperEcosystemValBDependencyStateActive ||
		response.ValECurrentState != operability.VerifierEcosystemValEStatePass ||
		response.ValEPoint7State != operability.VerifierEcosystemPoint7StatePass ||
		response.ValEPassRuleState != operability.VerifierEcosystemValEPassRuleStateActive ||
		response.ValENoOverclaimState != operability.VerifierEcosystemValENoOverclaimStateActive ||
		response.ValEPoint7PassReason != operability.VerifierEcosystemValEPoint7PassReasonAllowed {
		t.Fatalf("expected active Val E compatibility gate for developer Val B, got %#v", response)
	}
	if response.ValAState != operability.DeveloperEcosystemValAStateActive ||
		response.ValAPoint8State != operability.DeveloperEcosystemPoint8StateNotComplete ||
		response.ValADependencyState != operability.DeveloperEcosystemValADependencyStateActive ||
		response.IDEBaselineState != operability.DeveloperEcosystemValAIDEBaselineStateActive ||
		response.TrustFeedbackState != operability.DeveloperEcosystemValATrustFeedbackStateActive ||
		response.CAVIVEXContextState != operability.DeveloperEcosystemValACAVIVEXStateActive ||
		response.LocalAdvisoryState != operability.DeveloperEcosystemValALocalAdvisoryStateActive ||
		response.ValidationHarnessState != operability.DeveloperEcosystemValAValidationHarnessStateActive ||
		response.MockVerificationState != operability.DeveloperEcosystemValAMockVerificationStateActive ||
		response.InspectExplainState != operability.DeveloperEcosystemValAInspectExplainStateActive ||
		response.DegradedModeState != operability.DeveloperEcosystemValADegradedModeStateActive ||
		response.ValANoOverclaimState != operability.DeveloperEcosystemValANoOverclaimStateActive {
		t.Fatalf("expected active Val A dependency for developer Val B, got %#v", response)
	}
	if response.Point8State != operability.DeveloperEcosystemPoint8StateNotComplete ||
		response.RepoConfigSchemaState != operability.DeveloperEcosystemValBRepoConfigSchemaStateActive ||
		response.RepoConfigValidationState != operability.DeveloperEcosystemValBRepoConfigValidationStateActive ||
		response.PolicyPreviewState != operability.DeveloperEcosystemValBPolicyPreviewStateActive ||
		response.LocalCIContinuityState != operability.DeveloperEcosystemValBLocalCIContinuityStateActive ||
		response.APISDKSurfaceState != operability.DeveloperEcosystemValBAPISDKSurfaceStateActive ||
		response.ExamplesTemplatesState != operability.DeveloperEcosystemValBExamplesTemplatesStateActive ||
		response.APIVersioningState != operability.DeveloperEcosystemValBAPIVersioningStateActive ||
		response.NoOverclaimState != operability.DeveloperEcosystemValBNoOverclaimStateActive {
		t.Fatalf("expected active developer Val B contract states with point 8 not complete, got %#v", response)
	}
	if len(response.SupportedRepoSchemaVersions) != 1 ||
		len(response.SupportedSDKVersions) != 2 ||
		len(response.SurfaceRefs) != len(operability.DeveloperEcosystemValBProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.DeveloperEcosystemValBProofEvidenceRefs()) ||
		len(response.WhyPoint8NotPass) == 0 ||
		len(response.Limitations) == 0 ||
		len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact proof/evidence refs and read-only summary fields, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "developer_ecosystem_valb") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

func TestDeveloperEcosystemValBStatusExposesCanonicalCompatibilityContracts(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/developer-ecosystem/valb/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response developerEcosystemValBStatusResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode status: %v", err)
	}
	if response.Model.RepoConfigSchema.CompatibilityBehavior != operability.DeveloperEcosystemValBRepoConfigCompatibilityBehaviorBounded {
		t.Fatalf("expected canonical repo schema compatibility behavior, got %#v", response.Model.RepoConfigSchema)
	}
	if response.Model.APIVersioning.VersionIdentity != operability.DeveloperEcosystemValBAPIVersionIdentity ||
		response.Model.APIVersioning.CompatibilityWindow != operability.DeveloperEcosystemValBAPICompatibilityWindow {
		t.Fatalf("expected canonical API version identity/window, got %#v", response.Model.APIVersioning)
	}
}
