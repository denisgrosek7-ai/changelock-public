package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	developerEcosystemValBStatusSchema = "point8.developer_ecosystem.valb.status.v1"
	developerEcosystemValBProofsSchema = "point8.developer_ecosystem.valb.proofs.v1"
)

type developerEcosystemValBStatusResponse struct {
	SchemaVersion string                                        `json:"schema_version"`
	GeneratedAt   time.Time                                     `json:"generated_at"`
	CurrentState  string                                        `json:"current_state"`
	Model         operability.DeveloperEcosystemValBIntegration `json:"model"`
	RouteRefs     []string                                      `json:"route_refs,omitempty"`
	Limitations   []string                                      `json:"limitations,omitempty"`
}

type developerEcosystemValBProofsResponse struct {
	SchemaVersion               string    `json:"schema_version"`
	GeneratedAt                 time.Time `json:"generated_at"`
	CurrentState                string    `json:"current_state"`
	ValECurrentState            string    `json:"vale_current_state"`
	ValEPoint7State             string    `json:"vale_point_7_state"`
	ValEPassRuleState           string    `json:"vale_pass_rule_state"`
	ValENoOverclaimState        string    `json:"vale_no_overclaim_state"`
	ValEProofSurfaceState       string    `json:"vale_proof_surface_state"`
	ValEEvidenceQualityState    string    `json:"vale_evidence_quality_state"`
	ValEPoint7PassAllowed       bool      `json:"vale_point_7_pass_allowed"`
	ValEPoint7PassReason        string    `json:"vale_point_7_pass_reason"`
	ValAState                   string    `json:"vala_current_state"`
	ValAPoint8State             string    `json:"vala_point_8_state"`
	ValADependencyState         string    `json:"vala_dependency_state"`
	IDEBaselineState            string    `json:"ide_baseline_state"`
	TrustFeedbackState          string    `json:"trust_feedback_state"`
	CAVIVEXContextState         string    `json:"cavi_vex_context_state"`
	LocalAdvisoryState          string    `json:"local_advisory_state"`
	ValidationHarnessState      string    `json:"local_validation_harness_state"`
	MockVerificationState       string    `json:"mock_verification_server_state"`
	InspectExplainState         string    `json:"inspect_explain_state"`
	DegradedModeState           string    `json:"degraded_mode_state"`
	ValANoOverclaimState        string    `json:"vala_no_overclaim_state"`
	ValECompatibilityState      string    `json:"vale_compatibility_state"`
	DependencyState             string    `json:"dependency_state"`
	RepoConfigSchemaState       string    `json:"repo_config_schema_state"`
	RepoConfigValidationState   string    `json:"repo_config_validation_state"`
	PolicyPreviewState          string    `json:"policy_preview_state"`
	LocalCIContinuityState      string    `json:"local_ci_continuity_state"`
	APISDKSurfaceState          string    `json:"api_sdk_surface_state"`
	ExamplesTemplatesState      string    `json:"examples_templates_state"`
	APIVersioningState          string    `json:"api_versioning_state"`
	NoOverclaimState            string    `json:"no_overclaim_state"`
	Point8State                 string    `json:"point_8_state"`
	SupportedRepoSchemaVersions []string  `json:"supported_repo_schema_versions,omitempty"`
	SupportedSDKVersions        []string  `json:"supported_sdk_versions,omitempty"`
	SurfaceRefs                 []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                []string  `json:"evidence_refs,omitempty"`
	BlockingReasons             []string  `json:"blocking_reasons,omitempty"`
	WhyPoint8NotPass            []string  `json:"why_point_8_not_pass,omitempty"`
	Limitations                 []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer        string    `json:"projection_disclaimer"`
	IntegrationSummary          []string  `json:"integration_summary,omitempty"`
}

func developerEcosystemValBAllSurfaceRefs() []string {
	return operability.DeveloperEcosystemValBProofSurfaceRefs()
}

func developerEcosystemValBProjectionDisclaimer() string {
	return "projection_only not_canonical_truth developer_ecosystem_valb advisory_projection repo_sdk_integration"
}

func developerEcosystemValBEvidenceRefs() []string {
	return operability.DeveloperEcosystemValBProofEvidenceRefs()
}

func buildDeveloperEcosystemValBValECompatibilityGate() operability.DeveloperEcosystemValBValECompatibilityGate {
	model := operability.DeveloperEcosystemValBValECompatibilityGateModel()
	valE := buildVerifierEcosystemValEProofs()
	model.ValECurrentState = valE.ValEState
	model.Point7State = valE.Point7State
	model.PassRuleState = valE.PassRuleState
	model.NoOverclaimState = valE.NoOverclaimState
	model.ProofSurfaceState = valE.ProofSurfaceState
	model.EvidenceQualityState = valE.EvidenceQualityState
	model.Point7PassAllowed = valE.Point7PassAllowed
	model.Point7PassReason = valE.Point7PassReason
	model.SurfaceRefs = valE.SurfaceRefs
	model.EvidenceRefs = valE.EvidenceRefs
	model.ProjectionDisclaimer = valE.ProjectionDisclaimer
	return model
}

func buildDeveloperEcosystemValBDependencySnapshot() operability.DeveloperEcosystemValBDependencySnapshot {
	valA := buildDeveloperEcosystemValAProofs()
	return operability.DeveloperEcosystemValBDependencySnapshot{
		ValACurrentState:         valA.CurrentState,
		ValAPoint8State:          valA.Point8State,
		ValADependencyState:      valA.DependencyState,
		IDEBaselineState:         valA.IDEBaselineState,
		TrustFeedbackState:       valA.TrustFeedbackState,
		CAVIVEXContextState:      valA.CAVIVEXContextState,
		LocalAdvisoryState:       valA.LocalAdvisoryState,
		ValidationHarnessState:   valA.ValidationHarnessState,
		MockVerificationState:    valA.MockVerificationState,
		InspectExplainState:      valA.InspectExplainState,
		DegradedModeState:        valA.DegradedModeState,
		NoOverclaimState:         valA.NoOverclaimState,
		ValAProofSurfaceRefs:     valA.SurfaceRefs,
		ValAEvidenceRefs:         valA.EvidenceRefs,
		ValAProjectionDisclaimer: valA.ProjectionDisclaimer,
	}
}

func buildDeveloperEcosystemValBModel() operability.DeveloperEcosystemValBIntegration {
	model := operability.DeveloperEcosystemValBIntegrationModel()
	model.ValECompatibility = buildDeveloperEcosystemValBValECompatibilityGate()
	model.Dependency = buildDeveloperEcosystemValBDependencySnapshot()
	model = operability.ComputeDeveloperEcosystemValBIntegration(model)
	model.ValECompatibility.CurrentState = model.ValECompatibilityState
	model.RepoConfigSchema.CurrentState = model.RepoConfigSchemaState
	model.RepoConfigValidation.CurrentState = model.RepoConfigValidationState
	model.PolicyPreview.CurrentState = model.PolicyPreviewState
	model.LocalCIContinuity.CurrentState = model.LocalCIContinuityState
	model.DeveloperAPISDK.CurrentState = model.APISDKSurfaceState
	model.ExamplesTemplates.CurrentState = model.ExamplesTemplatesState
	model.APIVersioning.CurrentState = model.APIVersioningState
	model.NoOverclaim.CurrentState = model.NoOverclaimState
	return model
}

func (s server) developerEcosystemValBStatusHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildDeveloperEcosystemValBStatus())
}

func (s server) developerEcosystemValBProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildDeveloperEcosystemValBProofs())
}

func buildDeveloperEcosystemValBStatus() developerEcosystemValBStatusResponse {
	model := buildDeveloperEcosystemValBModel()
	limitations := []string{
		"Val B implements repo and SDK integration contracts only and does not ship a production SDK runtime, repo config parser/runtime, plugin runtime, marketplace publishing, or Točka 9 work.",
		"Repo config, policy preview, SDK/API, examples/templates, and local-to-CI continuity outputs remain advisory projections and do not approve deployment or mutate canonical evidence.",
	}
	return developerEcosystemValBStatusResponse{
		SchemaVersion: developerEcosystemValBStatusSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     developerEcosystemValBAllSurfaceRefs(),
		Limitations:   limitations,
	}
}

func buildDeveloperEcosystemValBProofs() developerEcosystemValBProofsResponse {
	model := buildDeveloperEcosystemValBModel()
	limitations := []string{
		"Val B implements repo and SDK integration contracts only and does not implement a production SDK runtime, repo parser/runtime, plugin runtime, marketplace publishing, or Točka 9 work.",
		"Točka 8 remains not complete because later developer ecosystem waves are still required before any integrated closure can exist.",
		"Repo config, policy preview, local-to-CI continuity, SDK/API, and examples/templates remain advisory only and cannot approve deployment, certify trust, or create canonical evidence.",
	}
	currentState := operability.EvaluateDeveloperEcosystemValBProofsState(model, limitations)
	return developerEcosystemValBProofsResponse{
		SchemaVersion:               developerEcosystemValBProofsSchema,
		GeneratedAt:                 publicSampleTime(),
		CurrentState:                currentState,
		ValECurrentState:            model.ValECompatibility.ValECurrentState,
		ValEPoint7State:             model.ValECompatibility.Point7State,
		ValEPassRuleState:           model.ValECompatibility.PassRuleState,
		ValENoOverclaimState:        model.ValECompatibility.NoOverclaimState,
		ValEProofSurfaceState:       model.ValECompatibility.ProofSurfaceState,
		ValEEvidenceQualityState:    model.ValECompatibility.EvidenceQualityState,
		ValEPoint7PassAllowed:       model.ValECompatibility.Point7PassAllowed,
		ValEPoint7PassReason:        model.ValECompatibility.Point7PassReason,
		ValAState:                   model.Dependency.ValACurrentState,
		ValAPoint8State:             model.Dependency.ValAPoint8State,
		ValADependencyState:         model.Dependency.ValADependencyState,
		IDEBaselineState:            model.Dependency.IDEBaselineState,
		TrustFeedbackState:          model.Dependency.TrustFeedbackState,
		CAVIVEXContextState:         model.Dependency.CAVIVEXContextState,
		LocalAdvisoryState:          model.Dependency.LocalAdvisoryState,
		ValidationHarnessState:      model.Dependency.ValidationHarnessState,
		MockVerificationState:       model.Dependency.MockVerificationState,
		InspectExplainState:         model.Dependency.InspectExplainState,
		DegradedModeState:           model.Dependency.DegradedModeState,
		ValANoOverclaimState:        model.Dependency.NoOverclaimState,
		ValECompatibilityState:      model.ValECompatibilityState,
		DependencyState:             model.DependencyState,
		RepoConfigSchemaState:       model.RepoConfigSchemaState,
		RepoConfigValidationState:   model.RepoConfigValidationState,
		PolicyPreviewState:          model.PolicyPreviewState,
		LocalCIContinuityState:      model.LocalCIContinuityState,
		APISDKSurfaceState:          model.APISDKSurfaceState,
		ExamplesTemplatesState:      model.ExamplesTemplatesState,
		APIVersioningState:          model.APIVersioningState,
		NoOverclaimState:            model.NoOverclaimState,
		Point8State:                 model.Point8State,
		SupportedRepoSchemaVersions: model.RepoConfigSchema.SupportedSchemaVersions,
		SupportedSDKVersions:        model.DeveloperAPISDK.SupportedVersions,
		SurfaceRefs:                 model.ProofSurfaceRefs,
		EvidenceRefs:                model.EvidenceRefs,
		BlockingReasons:             model.BlockingReasons,
		WhyPoint8NotPass: []string{
			"Val B is the repo and SDK integration contract layer only and cannot return point_8_pass.",
			"Later Točka 8 waves remain required before any integrated developer ecosystem closure can exist.",
			"Repo config, policy preview, SDK/API, and examples/templates remain advisory or projection-only and cannot become approval, certification, policy override, or canonical truth.",
		},
		Limitations:          limitations,
		ProjectionDisclaimer: developerEcosystemValBProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val B adds bounded repo-local schema and validation contracts, policy preview, local-to-CI continuity, developer API and SDK contracts, examples/templates, and API versioning on top of accepted Val A outputs and the patched Val E compatibility gate.",
			"Repo and SDK integration remain fail-closed, versioned, governance-safe, and explicitly separated from deployment approval, enterprise policy authority, and canonical evidence mutation.",
		},
	}
}
