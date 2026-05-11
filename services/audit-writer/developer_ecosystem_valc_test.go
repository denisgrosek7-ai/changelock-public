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
	if response.ProjectionDisclaimer != developerEcosystemValCProjectionDisclaimer() {
		t.Fatalf("expected exact projection disclaimer %q, got %#v", developerEcosystemValCProjectionDisclaimer(), response)
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

func TestDeveloperEcosystemValCModelUsesCanonicalValBSnapshot(t *testing.T) {
	valB := buildDeveloperEcosystemValBModel()
	valC := buildDeveloperEcosystemValCModel()

	if valC.ValECompatibility.ValECurrentState != valB.ValECompatibility.ValECurrentState ||
		valC.ValECompatibility.Point7State != valB.ValECompatibility.Point7State ||
		valC.ValECompatibility.PassRuleState != valB.ValECompatibility.PassRuleState ||
		valC.ValECompatibility.NoOverclaimState != valB.ValECompatibility.NoOverclaimState ||
		valC.ValECompatibility.ProofSurfaceState != valB.ValECompatibility.ProofSurfaceState ||
		valC.ValECompatibility.EvidenceQualityState != valB.ValECompatibility.EvidenceQualityState ||
		valC.ValECompatibility.Point7PassAllowed != valB.ValECompatibility.Point7PassAllowed ||
		valC.ValECompatibility.Point7PassReason != valB.ValECompatibility.Point7PassReason ||
		strings.Join(valC.ValECompatibility.SurfaceRefs, "\n") != strings.Join(valB.ValECompatibility.SurfaceRefs, "\n") ||
		strings.Join(valC.ValECompatibility.EvidenceRefs, "\n") != strings.Join(valB.ValECompatibility.EvidenceRefs, "\n") ||
		valC.ValECompatibility.ProjectionDisclaimer != valB.ValECompatibility.ProjectionDisclaimer {
		t.Fatalf("expected Val C Val E compatibility snapshot to match canonical Val B snapshot, got %#v vs %#v", valC.ValECompatibility, valB.ValECompatibility)
	}

	if valC.ValBCompatibility.ValBCurrentState != valB.CurrentState ||
		valC.ValBCompatibility.Point8State != valB.Point8State ||
		valC.ValBCompatibility.ValECompatibilityState != valB.ValECompatibilityState ||
		valC.ValBCompatibility.RepoConfigSchemaState != valB.RepoConfigSchemaState ||
		valC.ValBCompatibility.APIVersioningState != valB.APIVersioningState ||
		valC.ValBCompatibility.NoOverclaimState != valB.NoOverclaimState ||
		valC.ValBCompatibility.RepoConfigCompatibilityBehavior != valB.RepoConfigSchema.CompatibilityBehavior ||
		valC.ValBCompatibility.APIVersionIdentity != valB.APIVersioning.VersionIdentity ||
		valC.ValBCompatibility.APICompatibilityWindow != valB.APIVersioning.CompatibilityWindow ||
		strings.Join(valC.ValBCompatibility.SurfaceRefs, "\n") != strings.Join(valB.ProofSurfaceRefs, "\n") ||
		strings.Join(valC.ValBCompatibility.EvidenceRefs, "\n") != strings.Join(valB.EvidenceRefs, "\n") ||
		valC.ValBCompatibility.ProjectionDisclaimer != valB.ProjectionDisclaimer {
		t.Fatalf("expected Val C Val B compatibility snapshot to match canonical Val B model, got %#v vs %#v", valC.ValBCompatibility, valB)
	}

	if valC.Dependency.ValBCurrentState != valB.CurrentState ||
		valC.Dependency.ValBPoint8State != valB.Point8State ||
		valC.Dependency.ValECompatibilityState != valB.ValECompatibilityState ||
		valC.Dependency.DependencyState != valB.DependencyState ||
		valC.Dependency.RepoConfigSchemaState != valB.RepoConfigSchemaState ||
		valC.Dependency.RepoConfigValidationState != valB.RepoConfigValidationState ||
		valC.Dependency.PolicyPreviewState != valB.PolicyPreviewState ||
		valC.Dependency.LocalCIContinuityState != valB.LocalCIContinuityState ||
		valC.Dependency.APISDKSurfaceState != valB.APISDKSurfaceState ||
		valC.Dependency.ExamplesTemplatesState != valB.ExamplesTemplatesState ||
		valC.Dependency.APIVersioningState != valB.APIVersioningState ||
		valC.Dependency.NoOverclaimState != valB.NoOverclaimState ||
		strings.Join(valC.Dependency.ValBProofSurfaceRefs, "\n") != strings.Join(valB.ProofSurfaceRefs, "\n") ||
		strings.Join(valC.Dependency.ValBEvidenceRefs, "\n") != strings.Join(valB.EvidenceRefs, "\n") ||
		valC.Dependency.ValBProjectionDisclaimer != valB.ProjectionDisclaimer {
		t.Fatalf("expected Val C dependency snapshot to match canonical Val B model, got %#v vs %#v", valC.Dependency, valB)
	}
}
