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

func TestDeveloperEcosystemValEHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/developer-ecosystem/vale/closure?tenant_id=acme&environment=prod",
		"/v1/developer-ecosystem/vale/proofs?tenant_id=acme&environment=prod",
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

	req := httptest.NewRequest(http.MethodPost, "/v1/developer-ecosystem/vale/closure?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected closure route to remain read-only, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestDeveloperEcosystemValEProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/developer-ecosystem/vale/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response developerEcosystemValEProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.DeveloperEcosystemValEStatePass ||
		response.Point8State != operability.DeveloperEcosystemPoint8StatePass ||
		!response.Point8PassAllowed ||
		response.Point8PassReason != operability.DeveloperEcosystemValEPoint8PassReasonAllowed {
		t.Fatalf("expected Val E pass proofs response, got %#v", response)
	}
	if response.ValECompatibilityState != operability.DeveloperEcosystemValEValECompatibilityStateActive ||
		response.Val0SourceState != operability.DeveloperEcosystemValEVal0SourceStateActive ||
		response.ValASourceState != operability.DeveloperEcosystemValEValASourceStateActive ||
		response.ValBSourceState != operability.DeveloperEcosystemValEValBSourceStateActive ||
		response.ValCSourceState != operability.DeveloperEcosystemValEValCSourceStateActive ||
		response.ValDSourceState != operability.DeveloperEcosystemValEValDSourceStateActive ||
		response.FinalPassRuleState != operability.DeveloperEcosystemValEFinalPassRuleStateActive {
		t.Fatalf("expected active source and pass rule states, got %#v", response)
	}
	if response.ValDFinalDeveloperEcosystemState != operability.DeveloperEcosystemValDFinalGateStateActive ||
		response.ValBCompatibilityBehavior != operability.DeveloperEcosystemValBRepoConfigCompatibilityBehaviorBounded ||
		response.ValBAPIVersionIdentity != operability.DeveloperEcosystemValBAPIVersionIdentity ||
		response.ValBAPICompatibilityWindow != operability.DeveloperEcosystemValBAPICompatibilityWindow ||
		response.ValCSandboxDisciplineID != operability.DeveloperEcosystemValCSandboxIsolationDisciplineID ||
		response.ValCSandboxVersion != operability.DeveloperEcosystemValCSandboxIsolationVersion ||
		response.ValCPluginExecutionBudgetRef != operability.DeveloperEcosystemVal0PerformanceBudgetDisciplineID ||
		response.Val0PerformanceBudgetDisciplineID != operability.DeveloperEcosystemVal0PerformanceBudgetDisciplineID {
		t.Fatalf("expected canonical prior-wave closure values, got %#v", response)
	}
	if response.VerifyPolicyClassifierPath != "scripts/ci/collect_verify_policy_inputs.sh" ||
		response.VerifyPolicyActionPath != ".github/actions/changelock-shift-left/action.yml" ||
		response.VerifyPolicyKyvernoVersion != operability.DeveloperEcosystemValDVerifyPolicyKyvernoVersion {
		t.Fatalf("expected canonical verify-policy metadata, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "developer_ecosystem_vale") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
	if len(response.SurfaceRefs) != len(operability.DeveloperEcosystemValEProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.DeveloperEcosystemValEProofEvidenceRefs()) ||
		len(response.WhyPoint8Pass) == 0 ||
		len(response.Limitations) == 0 {
		t.Fatalf("expected exact proof/evidence refs and summary fields, got %#v", response)
	}
}

func TestDeveloperEcosystemValEClosureHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/developer-ecosystem/vale/closure?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected closure 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response developerEcosystemValEClosureResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode closure: %v", err)
	}
	if response.Model.CurrentState != operability.DeveloperEcosystemValEStatePass ||
		response.Model.Point8State != operability.DeveloperEcosystemPoint8StatePass ||
		!response.Model.Point8PassAllowed ||
		response.Model.ValDSource.Point8PassAvailable {
		t.Fatalf("expected active closure model with no prior-wave point_8_pass, got %#v", response.Model)
	}
	if response.Model.CleanRoomIPGuardrailState != operability.DeveloperEcosystemValECleanRoomIPGuardrailStateActive ||
		response.Model.CleanRoomIPGuardrail.LegalCertificationClaim {
		t.Fatalf("expected non-certifying clean-room guardrail, got %#v", response.Model.CleanRoomIPGuardrail)
	}
}
