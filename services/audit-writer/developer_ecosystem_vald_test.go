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

func TestDeveloperEcosystemValDHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/developer-ecosystem/vald/status?tenant_id=acme&environment=prod",
		"/v1/developer-ecosystem/vald/proofs?tenant_id=acme&environment=prod",
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

	req := httptest.NewRequest(http.MethodPost, "/v1/developer-ecosystem/vald/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status route to remain read-only, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestDeveloperEcosystemValDProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/developer-ecosystem/vald/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response developerEcosystemValDProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.DeveloperEcosystemValDStateActive ||
		response.ValECompatibilityState != operability.DeveloperEcosystemValDValECompatibilityStateActive ||
		response.Val0FoundationState != operability.DeveloperEcosystemValDVal0FoundationStateActive ||
		response.ValAReadinessState != operability.DeveloperEcosystemValDValAReadinessStateActive ||
		response.ValBReadinessState != operability.DeveloperEcosystemValDValBReadinessStateActive ||
		response.ValCReadinessState != operability.DeveloperEcosystemValDValCReadinessStateActive ||
		response.VerifyPolicyCICompatibilityState != operability.DeveloperEcosystemValDVerifyPolicyCICompatibilityStateActive ||
		response.FinalDeveloperEcosystemGateState != operability.DeveloperEcosystemValDFinalGateStateActive ||
		response.NoOverclaimState != operability.DeveloperEcosystemValDNoOverclaimStateActive ||
		response.Point8State != operability.DeveloperEcosystemPoint8StateNotComplete {
		t.Fatalf("expected active Val D proofs state with point 8 not complete, got %#v", response)
	}
	if response.ValEPoint7PassReason != operability.VerifierEcosystemValEPoint7PassReasonAllowed {
		t.Fatalf("expected canonical Val E pass reason, got %#v", response)
	}
	if response.Val0PerformanceBudgetDisciplineID != operability.DeveloperEcosystemVal0PerformanceBudgetDisciplineID ||
		response.ValBCompatibilityBehavior != operability.DeveloperEcosystemValBRepoConfigCompatibilityBehaviorBounded ||
		response.ValBAPIVersionIdentity != operability.DeveloperEcosystemValBAPIVersionIdentity ||
		response.ValBAPICompatibilityWindow != operability.DeveloperEcosystemValBAPICompatibilityWindow ||
		response.ValCSandboxDisciplineID != operability.DeveloperEcosystemValCSandboxIsolationDisciplineID ||
		response.ValCSandboxVersion != operability.DeveloperEcosystemValCSandboxIsolationVersion ||
		response.ValCPluginExecutionBudgetRef != operability.DeveloperEcosystemVal0PerformanceBudgetDisciplineID {
		t.Fatalf("expected canonical prior-wave compatibility values, got %#v", response)
	}
	if response.VerifyPolicyClassifierPath != "scripts/ci/collect_verify_policy_inputs.sh" ||
		response.VerifyPolicyActionPath != ".github/actions/changelock-shift-left/action.yml" ||
		response.VerifyPolicyKyvernoVersion != operability.DeveloperEcosystemValDVerifyPolicyKyvernoVersion {
		t.Fatalf("expected canonical verify-policy compatibility metadata, got %#v", response)
	}
	if len(response.TriggerOnlyPrefixes) != len(operability.DeveloperEcosystemValDVerifyPolicyCICompatibilityModel().TriggerOnlyPrefixes) ||
		len(response.ManifestResourcePrefixes) != len(operability.DeveloperEcosystemValDVerifyPolicyCICompatibilityModel().ManifestResourcePrefixes) ||
		len(response.OptionOnlyArgs) != len(operability.DeveloperEcosystemValDVerifyPolicyCICompatibilityModel().OptionOnlyArgs) ||
		len(response.SurfaceRefs) != len(operability.DeveloperEcosystemValDProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.DeveloperEcosystemValDProofEvidenceRefs()) ||
		len(response.WhyPoint8NotPass) == 0 ||
		len(response.Limitations) == 0 ||
		len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact proof/evidence refs and read-only summary fields, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "developer_ecosystem_vald") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

func TestDeveloperEcosystemValDStatusExposesCanonicalReadinessContracts(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/developer-ecosystem/vald/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response developerEcosystemValDStatusResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode status: %v", err)
	}
	if response.Model.VerifyPolicyCICompatibility.KyvernoVersion != operability.DeveloperEcosystemValDVerifyPolicyKyvernoVersion ||
		response.Model.VerifyPolicyCICompatibility.ClassifierScriptPath != "scripts/ci/collect_verify_policy_inputs.sh" ||
		response.Model.VerifyPolicyCICompatibility.NoInputBehavior != operability.DeveloperEcosystemValDVerifyPolicyNoInputBehavior {
		t.Fatalf("expected canonical verify-policy compatibility contract, got %#v", response.Model.VerifyPolicyCICompatibility)
	}
	if response.Model.CleanRoomIPGuardrail.CurrentState != operability.DeveloperEcosystemValDCleanRoomIPGuardrailStateActive ||
		response.Model.CleanRoomIPGuardrail.LegalCertificationClaim {
		t.Fatalf("expected clean-room guardrail active without legal certification claim, got %#v", response.Model.CleanRoomIPGuardrail)
	}
}
