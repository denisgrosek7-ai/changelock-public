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

func TestDeveloperEcosystemValCHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/developer-ecosystem/valc/status?tenant_id=acme&environment=prod",
		"/v1/developer-ecosystem/valc/proofs?tenant_id=acme&environment=prod",
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

	req := httptest.NewRequest(http.MethodPost, "/v1/developer-ecosystem/valc/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status route to remain read-only, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestDeveloperEcosystemValCProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/developer-ecosystem/valc/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response developerEcosystemValCProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.DeveloperEcosystemValCStateActive ||
		response.ValECurrentState != operability.VerifierEcosystemValEStatePass ||
		response.ValEPassRuleState != operability.VerifierEcosystemValEPassRuleStateActive ||
		response.ValENoOverclaimState != operability.VerifierEcosystemValENoOverclaimStateActive ||
		response.ValEPoint7PassReason != operability.VerifierEcosystemValEPoint7PassReasonAllowed {
		t.Fatalf("expected active Val E compatibility gate for developer Val C, got %#v", response)
	}
	if response.ValBCurrentState != operability.DeveloperEcosystemValBStateActive ||
		response.ValBPoint8State != operability.DeveloperEcosystemPoint8StateNotComplete ||
		response.ValBValECompatibilityState != operability.DeveloperEcosystemValBValECompatibilityStateActive ||
		response.ValBDependencyState != operability.DeveloperEcosystemValBDependencyStateActive ||
		response.ValBRepoConfigSchemaState != operability.DeveloperEcosystemValBRepoConfigSchemaStateActive ||
		response.ValBAPIVersioningState != operability.DeveloperEcosystemValBAPIVersioningStateActive ||
		response.ValBNoOverclaimState != operability.DeveloperEcosystemValBNoOverclaimStateActive ||
		response.ValBCompatibilityState != operability.DeveloperEcosystemValCValBCompatibilityStateActive ||
		response.DependencyState != operability.DeveloperEcosystemValCDependencyStateActive {
		t.Fatalf("expected active Val B compatibility and dependency gates for developer Val C, got %#v", response)
	}
	if response.Point8State != operability.DeveloperEcosystemPoint8StateNotComplete ||
		response.PluginManifestState != operability.DeveloperEcosystemValCPluginManifestStateActive ||
		response.PluginLifecycleState != operability.DeveloperEcosystemValCPluginLifecycleStateActive ||
		response.CapabilityDeclarationState != operability.DeveloperEcosystemValCCapabilityStateActive ||
		response.SandboxIsolationState != operability.DeveloperEcosystemValCSandboxIsolationStateActive ||
		response.BoundedCustomChecksState != operability.DeveloperEcosystemValCCustomChecksStateActive ||
		response.PluginDiagnosticsState != operability.DeveloperEcosystemValCPluginDiagnosticsStateActive ||
		response.PluginPerformanceState != operability.DeveloperEcosystemValCPluginPerformanceStateActive ||
		response.PluginTrustBoundaryState != operability.DeveloperEcosystemValCPluginTrustBoundaryStateActive ||
		response.SamplePluginDescriptorState != operability.DeveloperEcosystemValCSamplePluginDescriptorStateActive ||
		response.ExtensionCompatibilityState != operability.DeveloperEcosystemValCExtensionCompatibilityStateActive ||
		response.NoOverclaimState != operability.DeveloperEcosystemValCNoOverclaimStateActive {
		t.Fatalf("expected active developer Val C contract states with point 8 not complete, got %#v", response)
	}
	if response.SandboxIsolationDisciplineID != operability.DeveloperEcosystemValCSandboxIsolationDisciplineID ||
		response.SandboxIsolationVersion != operability.DeveloperEcosystemValCSandboxIsolationVersion {
		t.Fatalf("expected canonical sandbox identity metadata, got %#v", response)
	}
	if response.PluginExecutionBudgetRef != operability.DeveloperEcosystemVal0PerformanceBudgetDisciplineID {
		t.Fatalf("expected canonical plugin execution budget ref, got %#v", response)
	}
	if response.PluginAPIVersionIdentity != operability.DeveloperEcosystemValCPluginAPIVersionIdentity ||
		response.PluginAPICompatibilityWindow != operability.DeveloperEcosystemValCPluginAPICompatibilityWindow {
		t.Fatalf("expected canonical plugin API identity/window, got %#v", response)
	}
	if len(response.DeclaredCapabilities) != len(operability.DeveloperEcosystemValCCapabilityDeclarationDisciplineModel().DeclaredCapabilities) ||
		len(response.SupportedPluginVersions) != len(operability.DeveloperEcosystemValCExtensionCompatibilityDisciplineModel().SupportedVersions) ||
		len(response.SurfaceRefs) != len(operability.DeveloperEcosystemValCProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.DeveloperEcosystemValCProofEvidenceRefs()) ||
		len(response.WhyPoint8NotPass) == 0 ||
		len(response.Limitations) == 0 ||
		len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact proof/evidence refs and read-only summary fields, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "developer_ecosystem_valc") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

func TestDeveloperEcosystemValCStatusExposesCanonicalPluginContracts(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/developer-ecosystem/valc/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response developerEcosystemValCStatusResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode status: %v", err)
	}
	if response.Model.PluginPerformance.PluginExecutionBudgetRef != operability.DeveloperEcosystemVal0PerformanceBudgetDisciplineID {
		t.Fatalf("expected canonical plugin performance budget ref, got %#v", response.Model.PluginPerformance)
	}
	if response.Model.SandboxIsolation.DisciplineID != operability.DeveloperEcosystemValCSandboxIsolationDisciplineID ||
		response.Model.SandboxIsolation.Version != operability.DeveloperEcosystemValCSandboxIsolationVersion {
		t.Fatalf("expected canonical sandbox discipline identity metadata, got %#v", response.Model.SandboxIsolation)
	}
	if response.Model.ExtensionCompatibility.PluginAPIVersionIdentity != operability.DeveloperEcosystemValCPluginAPIVersionIdentity ||
		response.Model.ExtensionCompatibility.CompatibilityWindow != operability.DeveloperEcosystemValCPluginAPICompatibilityWindow {
		t.Fatalf("expected canonical plugin API identity/window, got %#v", response.Model.ExtensionCompatibility)
	}
}
