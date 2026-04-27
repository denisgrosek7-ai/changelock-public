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

func TestVerifierEcosystemValAHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/verifier-ecosystem/vala/input-model?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vala/verifier-engine?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vala/verification-report?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vala/diagnostics-mapping?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vala/command-contract?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vala/sdk-entrypoint?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vala/proofs?tenant_id=acme&environment=prod",
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
}

func TestVerifierEcosystemValAProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/verifier-ecosystem/vala/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response verifierEcosystemValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.VerifierEcosystemValAStateActive ||
		response.ValAState != operability.VerifierEcosystemValAStateActive ||
		response.Point5State != operability.IntelligenceCalibrationPoint5StatePass ||
		response.Point5DependencyState != operability.IntelligenceCalibrationValEStateActive ||
		response.Point6State != operability.ReferenceArchitecturePoint6StatePass ||
		response.Point6ClosureState != operability.ReferenceArchitectureValEStateActive ||
		!response.Point6PassAllowed ||
		response.Val0CurrentState != operability.VerifierEcosystemVal0StateActive ||
		response.Val0State != operability.VerifierEcosystemVal0StateActive {
		t.Fatalf("expected active Val A proofs with healthy Točka 6 and Val 0 dependencies, got %#v", response)
	}
	if response.Point7State != operability.VerifierEcosystemPoint7StateNotComplete ||
		response.InputModelState != operability.VerifierEcosystemValAInputStateActive ||
		response.VerifierEngineState != operability.VerifierEcosystemValAEngineStateActive ||
		response.VerificationResultState != operability.VerifierEcosystemValAResultStateActive ||
		response.DiagnosticsMappingState != operability.VerifierEcosystemValADiagnosticsMappingStateActive ||
		response.CommandContractState != operability.VerifierEcosystemValACommandContractStateActive ||
		response.SDKEntrypointState != operability.VerifierEcosystemValASDKEntrypointStateActive ||
		response.VerificationOverallResult != operability.VerifierEcosystemValAOverallResultVerified ||
		response.VerificationDiagnosticClass != operability.VerifierEcosystemDiagnosticVerified {
		t.Fatalf("expected active Val A component and verification states, got %#v", response)
	}
	if len(response.SurfaceRefs) != 8 || len(response.EvidenceRefs) < 8 || len(response.WhyPoint7NotPass) == 0 || len(response.Limitations) == 0 || len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact surface refs, merged evidence refs, limitations, and summary, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}
