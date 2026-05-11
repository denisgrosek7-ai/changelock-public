package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	developerEcosystemValCStatusSchema = "point8.developer_ecosystem.valc.status.v1"
	developerEcosystemValCProofsSchema = "point8.developer_ecosystem.valc.proofs.v1"
)

var developerEcosystemValCModelCache cachedJSONValue

type developerEcosystemValCStatusResponse struct {
	SchemaVersion string                                        `json:"schema_version"`
	GeneratedAt   time.Time                                     `json:"generated_at"`
	CurrentState  string                                        `json:"current_state"`
	Model         operability.DeveloperEcosystemValCIntegration `json:"model"`
	RouteRefs     []string                                      `json:"route_refs,omitempty"`
	Limitations   []string                                      `json:"limitations,omitempty"`
}

type developerEcosystemValCProofsResponse struct {
	SchemaVersion                string    `json:"schema_version"`
	GeneratedAt                  time.Time `json:"generated_at"`
	CurrentState                 string    `json:"current_state"`
	ValECurrentState             string    `json:"vale_current_state"`
	ValEPoint7State              string    `json:"vale_point_7_state"`
	ValEPassRuleState            string    `json:"vale_pass_rule_state"`
	ValENoOverclaimState         string    `json:"vale_no_overclaim_state"`
	ValEProofSurfaceState        string    `json:"vale_proof_surface_state"`
	ValEEvidenceQualityState     string    `json:"vale_evidence_quality_state"`
	ValEPoint7PassAllowed        bool      `json:"vale_point_7_pass_allowed"`
	ValEPoint7PassReason         string    `json:"vale_point_7_pass_reason"`
	ValBCurrentState             string    `json:"valb_current_state"`
	ValBPoint8State              string    `json:"valb_point_8_state"`
	ValBValECompatibilityState   string    `json:"valb_vale_compatibility_state"`
	ValBDependencyState          string    `json:"valb_dependency_state"`
	ValBRepoConfigSchemaState    string    `json:"valb_repo_config_schema_state"`
	ValBAPIVersioningState       string    `json:"valb_api_versioning_state"`
	ValBNoOverclaimState         string    `json:"valb_no_overclaim_state"`
	ValBCompatibilityState       string    `json:"valb_compatibility_state"`
	DependencyState              string    `json:"dependency_state"`
	PluginManifestState          string    `json:"plugin_manifest_state"`
	PluginLifecycleState         string    `json:"plugin_lifecycle_state"`
	CapabilityDeclarationState   string    `json:"capability_declaration_state"`
	SandboxIsolationState        string    `json:"sandbox_isolation_state"`
	BoundedCustomChecksState     string    `json:"bounded_custom_checks_state"`
	PluginDiagnosticsState       string    `json:"plugin_diagnostics_state"`
	PluginPerformanceState       string    `json:"plugin_performance_state"`
	PluginTrustBoundaryState     string    `json:"plugin_trust_boundary_state"`
	SamplePluginDescriptorState  string    `json:"sample_plugin_descriptor_state"`
	ExtensionCompatibilityState  string    `json:"extension_compatibility_state"`
	NoOverclaimState             string    `json:"no_overclaim_state"`
	Point8State                  string    `json:"point_8_state"`
	DeclaredCapabilities         []string  `json:"declared_capabilities,omitempty"`
	SupportedPluginVersions      []string  `json:"supported_plugin_versions,omitempty"`
	SandboxIsolationDisciplineID string    `json:"sandbox_isolation_discipline_id"`
	SandboxIsolationVersion      string    `json:"sandbox_isolation_version"`
	PluginExecutionBudgetRef     string    `json:"plugin_execution_budget_ref"`
	PluginAPIVersionIdentity     string    `json:"plugin_api_version_identity"`
	PluginAPICompatibilityWindow string    `json:"plugin_api_compatibility_window"`
	SurfaceRefs                  []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                 []string  `json:"evidence_refs,omitempty"`
	BlockingReasons              []string  `json:"blocking_reasons,omitempty"`
	WhyPoint8NotPass             []string  `json:"why_point_8_not_pass,omitempty"`
	Limitations                  []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer         string    `json:"projection_disclaimer"`
	IntegrationSummary           []string  `json:"integration_summary,omitempty"`
}

func developerEcosystemValCAllSurfaceRefs() []string {
	return operability.DeveloperEcosystemValCProofSurfaceRefs()
}

func developerEcosystemValCProjectionDisclaimer() string {
	return "projection_only not_canonical_truth developer_ecosystem_valc advisory_projection plugin_extensibility"
}

func buildDeveloperEcosystemValCValECompatibilityGateFromValB(valB operability.DeveloperEcosystemValBIntegration) operability.DeveloperEcosystemValCValECompatibilityGate {
	model := operability.DeveloperEcosystemValCValECompatibilityGateModel()
	model.ValECurrentState = valB.ValECompatibility.ValECurrentState
	model.Point7State = valB.ValECompatibility.Point7State
	model.PassRuleState = valB.ValECompatibility.PassRuleState
	model.NoOverclaimState = valB.ValECompatibility.NoOverclaimState
	model.ProofSurfaceState = valB.ValECompatibility.ProofSurfaceState
	model.EvidenceQualityState = valB.ValECompatibility.EvidenceQualityState
	model.Point7PassAllowed = valB.ValECompatibility.Point7PassAllowed
	model.Point7PassReason = valB.ValECompatibility.Point7PassReason
	model.SurfaceRefs = append([]string{}, valB.ValECompatibility.SurfaceRefs...)
	model.EvidenceRefs = append([]string{}, valB.ValECompatibility.EvidenceRefs...)
	model.ProjectionDisclaimer = valB.ValECompatibility.ProjectionDisclaimer
	return model
}

func buildDeveloperEcosystemValCValBCompatibilityGateFromValB(valB operability.DeveloperEcosystemValBIntegration) operability.DeveloperEcosystemValCValBCompatibilityGate {
	model := operability.DeveloperEcosystemValCValBCompatibilityGateModel()
	model.ValBCurrentState = valB.CurrentState
	model.Point8State = valB.Point8State
	model.ValECompatibilityState = valB.ValECompatibilityState
	model.RepoConfigSchemaState = valB.RepoConfigSchemaState
	model.APIVersioningState = valB.APIVersioningState
	model.NoOverclaimState = valB.NoOverclaimState
	model.RepoConfigCompatibilityBehavior = valB.RepoConfigSchema.CompatibilityBehavior
	model.APIVersionIdentity = valB.APIVersioning.VersionIdentity
	model.APICompatibilityWindow = valB.APIVersioning.CompatibilityWindow
	model.SurfaceRefs = valB.ProofSurfaceRefs
	model.EvidenceRefs = valB.EvidenceRefs
	model.ProjectionDisclaimer = valB.ProjectionDisclaimer
	return model
}

func buildDeveloperEcosystemValCDependencySnapshotFromValB(valB operability.DeveloperEcosystemValBIntegration) operability.DeveloperEcosystemValCDependencySnapshot {
	return operability.DeveloperEcosystemValCDependencySnapshot{
		ValBCurrentState:          valB.CurrentState,
		ValBPoint8State:           valB.Point8State,
		ValECompatibilityState:    valB.ValECompatibilityState,
		DependencyState:           valB.DependencyState,
		RepoConfigSchemaState:     valB.RepoConfigSchemaState,
		RepoConfigValidationState: valB.RepoConfigValidationState,
		PolicyPreviewState:        valB.PolicyPreviewState,
		LocalCIContinuityState:    valB.LocalCIContinuityState,
		APISDKSurfaceState:        valB.APISDKSurfaceState,
		ExamplesTemplatesState:    valB.ExamplesTemplatesState,
		APIVersioningState:        valB.APIVersioningState,
		NoOverclaimState:          valB.NoOverclaimState,
		ValBProofSurfaceRefs:      valB.ProofSurfaceRefs,
		ValBEvidenceRefs:          valB.EvidenceRefs,
		ValBProjectionDisclaimer:  valB.ProjectionDisclaimer,
	}
}

func buildDeveloperEcosystemValCModelUncached() operability.DeveloperEcosystemValCIntegration {
	valB := buildDeveloperEcosystemValBModel()
	model := operability.DeveloperEcosystemValCIntegrationModel()
	model.ValECompatibility = buildDeveloperEcosystemValCValECompatibilityGateFromValB(valB)
	model.ValBCompatibility = buildDeveloperEcosystemValCValBCompatibilityGateFromValB(valB)
	model.Dependency = buildDeveloperEcosystemValCDependencySnapshotFromValB(valB)
	model = operability.ComputeDeveloperEcosystemValCIntegration(model)
	model.ValECompatibility.CurrentState = model.ValECompatibilityState
	model.ValBCompatibility.CurrentState = model.ValBCompatibilityState
	model.PluginManifest.CurrentState = model.PluginManifestState
	model.PluginLifecycle.CurrentState = model.PluginLifecycleState
	model.CapabilityDeclaration.CurrentState = model.CapabilityDeclarationState
	model.SandboxIsolation.CurrentState = model.SandboxIsolationState
	model.BoundedCustomChecks.CurrentState = model.BoundedCustomChecksState
	model.PluginDiagnostics.CurrentState = model.PluginDiagnosticsState
	model.PluginPerformance.CurrentState = model.PluginPerformanceState
	model.PluginTrustBoundary.CurrentState = model.PluginTrustBoundaryState
	model.SamplePluginDescriptors.CurrentState = model.SamplePluginDescriptorState
	model.ExtensionCompatibility.CurrentState = model.ExtensionCompatibilityState
	model.NoOverclaim.CurrentState = model.NoOverclaimState
	return model
}

func buildDeveloperEcosystemValCModel() operability.DeveloperEcosystemValCIntegration {
	return loadCachedJSON(&developerEcosystemValCModelCache, buildDeveloperEcosystemValCModelUncached)
}

func (s server) developerEcosystemValCStatusHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildDeveloperEcosystemValCStatus())
}

func (s server) developerEcosystemValCProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildDeveloperEcosystemValCProofs())
}

func buildDeveloperEcosystemValCStatus() developerEcosystemValCStatusResponse {
	model := buildDeveloperEcosystemValCModel()
	limitations := []string{
		"Val C implements plugin and extensibility contracts only and does not ship a plugin runtime, marketplace, external registry, remote installation flow, production SDK runtime, or Točka 9 work.",
		"Plugin manifests, diagnostics, custom checks, and sample descriptors remain advisory projections and do not approve deployment, certify trust, or mutate canonical evidence.",
	}
	return developerEcosystemValCStatusResponse{
		SchemaVersion: developerEcosystemValCStatusSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     developerEcosystemValCAllSurfaceRefs(),
		Limitations:   limitations,
	}
}

func buildDeveloperEcosystemValCProofs() developerEcosystemValCProofsResponse {
	model := buildDeveloperEcosystemValCModel()
	limitations := []string{
		"Val C implements plugin and extensibility contracts only and does not implement a plugin runtime, marketplace, external registry, remote installation, production SDK runtime, or Točka 9 work.",
		"Točka 8 remains not complete because later developer ecosystem waves are still required before any integrated closure can exist.",
		"Plugin manifests, diagnostics, custom checks, compatibility descriptors, and samples remain advisory only and cannot approve deployment, certify trust, or create canonical evidence.",
	}
	currentState := operability.EvaluateDeveloperEcosystemValCProofsState(model, limitations)
	return developerEcosystemValCProofsResponse{
		SchemaVersion:                developerEcosystemValCProofsSchema,
		GeneratedAt:                  publicSampleTime(),
		CurrentState:                 currentState,
		ValECurrentState:             model.ValECompatibility.ValECurrentState,
		ValEPoint7State:              model.ValECompatibility.Point7State,
		ValEPassRuleState:            model.ValECompatibility.PassRuleState,
		ValENoOverclaimState:         model.ValECompatibility.NoOverclaimState,
		ValEProofSurfaceState:        model.ValECompatibility.ProofSurfaceState,
		ValEEvidenceQualityState:     model.ValECompatibility.EvidenceQualityState,
		ValEPoint7PassAllowed:        model.ValECompatibility.Point7PassAllowed,
		ValEPoint7PassReason:         model.ValECompatibility.Point7PassReason,
		ValBCurrentState:             model.Dependency.ValBCurrentState,
		ValBPoint8State:              model.Dependency.ValBPoint8State,
		ValBValECompatibilityState:   model.Dependency.ValECompatibilityState,
		ValBDependencyState:          model.Dependency.DependencyState,
		ValBRepoConfigSchemaState:    model.Dependency.RepoConfigSchemaState,
		ValBAPIVersioningState:       model.Dependency.APIVersioningState,
		ValBNoOverclaimState:         model.Dependency.NoOverclaimState,
		ValBCompatibilityState:       model.ValBCompatibilityState,
		DependencyState:              model.DependencyState,
		PluginManifestState:          model.PluginManifestState,
		PluginLifecycleState:         model.PluginLifecycleState,
		CapabilityDeclarationState:   model.CapabilityDeclarationState,
		SandboxIsolationState:        model.SandboxIsolationState,
		BoundedCustomChecksState:     model.BoundedCustomChecksState,
		PluginDiagnosticsState:       model.PluginDiagnosticsState,
		PluginPerformanceState:       model.PluginPerformanceState,
		PluginTrustBoundaryState:     model.PluginTrustBoundaryState,
		SamplePluginDescriptorState:  model.SamplePluginDescriptorState,
		ExtensionCompatibilityState:  model.ExtensionCompatibilityState,
		NoOverclaimState:             model.NoOverclaimState,
		Point8State:                  model.Point8State,
		DeclaredCapabilities:         model.CapabilityDeclaration.DeclaredCapabilities,
		SupportedPluginVersions:      model.ExtensionCompatibility.SupportedVersions,
		SandboxIsolationDisciplineID: model.SandboxIsolation.DisciplineID,
		SandboxIsolationVersion:      model.SandboxIsolation.Version,
		PluginExecutionBudgetRef:     model.PluginPerformance.PluginExecutionBudgetRef,
		PluginAPIVersionIdentity:     model.ExtensionCompatibility.PluginAPIVersionIdentity,
		PluginAPICompatibilityWindow: model.ExtensionCompatibility.CompatibilityWindow,
		SurfaceRefs:                  model.ProofSurfaceRefs,
		EvidenceRefs:                 model.EvidenceRefs,
		BlockingReasons:              model.BlockingReasons,
		WhyPoint8NotPass: []string{
			"Val C is the plugin and extensibility contract layer only and cannot return point_8_pass.",
			"Later Točka 8 waves remain required before any integrated developer ecosystem closure can exist.",
			"Plugin manifests, diagnostics, custom checks, and sample descriptors remain advisory or projection-only and cannot become approval, certification, or canonical truth.",
		},
		Limitations:          limitations,
		ProjectionDisclaimer: developerEcosystemValCProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val C adds bounded plugin manifest, lifecycle, capability, sandbox, diagnostics, performance, trust-boundary, sample descriptor, and extension compatibility contracts on top of patched Val E and patched Val B gates.",
			"Plugin and extensibility integration remain fail-closed, governance-safe, and explicitly separated from deployment approval, certification, production authorization, and canonical evidence mutation.",
		},
	}
}
