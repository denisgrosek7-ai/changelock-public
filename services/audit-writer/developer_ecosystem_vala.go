package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	developerEcosystemValAStatusSchema = "point8.developer_ecosystem.vala.status.v1"
	developerEcosystemValAProofsSchema = "point8.developer_ecosystem.vala.proofs.v1"
)

type developerEcosystemValAStatusResponse struct {
	SchemaVersion string                                 `json:"schema_version"`
	GeneratedAt   time.Time                              `json:"generated_at"`
	CurrentState  string                                 `json:"current_state"`
	Model         operability.DeveloperEcosystemValACore `json:"model"`
	RouteRefs     []string                               `json:"route_refs,omitempty"`
	Limitations   []string                               `json:"limitations,omitempty"`
}

type developerEcosystemValAProofsResponse struct {
	SchemaVersion               string    `json:"schema_version"`
	GeneratedAt                 time.Time `json:"generated_at"`
	CurrentState                string    `json:"current_state"`
	Val0CurrentState            string    `json:"val0_current_state"`
	Val0Point8State             string    `json:"val0_point_8_state"`
	Val0OutputClassification    string    `json:"val0_output_classification_state"`
	Val0IDEAdvisoryState        string    `json:"val0_ide_advisory_state"`
	Val0LocalProductionState    string    `json:"val0_local_production_state"`
	Val0RepoPolicyBoundaryState string    `json:"val0_repo_policy_boundary_state"`
	Val0PluginSafetyState       string    `json:"val0_plugin_safety_state"`
	Val0PerformanceBudgetState  string    `json:"val0_performance_budget_state"`
	Val0DXMetricsState          string    `json:"val0_dx_metrics_state"`
	Val0NoOverclaimState        string    `json:"val0_no_overclaim_state"`
	DependencyState             string    `json:"dependency_state"`
	IDEBaselineState            string    `json:"ide_baseline_state"`
	TrustFeedbackState          string    `json:"trust_feedback_state"`
	CAVIVEXContextState         string    `json:"cavi_vex_context_state"`
	LocalAdvisoryState          string    `json:"local_advisory_state"`
	ValidationHarnessState      string    `json:"local_validation_harness_state"`
	MockVerificationState       string    `json:"mock_verification_server_state"`
	InspectExplainState         string    `json:"inspect_explain_state"`
	DegradedModeState           string    `json:"degraded_mode_state"`
	NoOverclaimState            string    `json:"no_overclaim_state"`
	Point8State                 string    `json:"point_8_state"`
	SupportedEditors            []string  `json:"supported_editors,omitempty"`
	TrustSignalClasses          []string  `json:"trust_signal_classes,omitempty"`
	ValidationClasses           []string  `json:"validation_classes,omitempty"`
	SurfaceRefs                 []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                []string  `json:"evidence_refs,omitempty"`
	BlockingReasons             []string  `json:"blocking_reasons,omitempty"`
	WhyPoint8NotPass            []string  `json:"why_point_8_not_pass,omitempty"`
	Limitations                 []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer        string    `json:"projection_disclaimer"`
	IntegrationSummary          []string  `json:"integration_summary,omitempty"`
}

func developerEcosystemValAAllSurfaceRefs() []string {
	return operability.DeveloperEcosystemValAProofSurfaceRefs()
}

func developerEcosystemValAProjectionDisclaimer() string {
	return "projection_only not_canonical_truth developer_ecosystem_vala advisory_projection ide_local_tooling_core"
}

func developerEcosystemValAEvidenceRefs() []string {
	return operability.DeveloperEcosystemValAProofEvidenceRefs()
}

func buildDeveloperEcosystemValADependencySnapshot() operability.DeveloperEcosystemValADependencySnapshot {
	val0 := buildDeveloperEcosystemVal0Proofs()
	return operability.DeveloperEcosystemValADependencySnapshot{
		Val0CurrentState:           val0.CurrentState,
		Val0Point8State:            val0.Point8State,
		Val0OutputClassification:   val0.OutputClassificationState,
		Val0IDEAdvisoryState:       val0.IDEAdvisoryState,
		Val0LocalProductionState:   val0.LocalProductionState,
		Val0RepoPolicyBoundary:     val0.RepoPolicyBoundaryState,
		Val0PluginSafetyState:      val0.PluginSafetyState,
		Val0PerformanceBudgetState: val0.PerformanceBudgetState,
		Val0DXMetricsState:         val0.DXMetricsState,
		Val0NoOverclaimState:       val0.NoOverclaimState,
		Val0ProofSurfaceRefs:       val0.SurfaceRefs,
		Val0EvidenceRefs:           val0.EvidenceRefs,
		Val0ProjectionDisclaimer:   val0.ProjectionDisclaimer,
	}
}

func buildDeveloperEcosystemValAModel() operability.DeveloperEcosystemValACore {
	model := operability.DeveloperEcosystemValACoreModel()
	model.Dependency = buildDeveloperEcosystemValADependencySnapshot()
	model = operability.ComputeDeveloperEcosystemValACore(model)
	model.IDEBaseline.CurrentState = model.IDEBaselineState
	model.TrustFeedback.CurrentState = model.TrustFeedbackState
	model.CAVIVEXContext.CurrentState = model.CAVIVEXContextState
	model.LocalAdvisory.CurrentState = model.LocalAdvisoryState
	model.ValidationHarness.CurrentState = model.ValidationHarnessState
	model.MockVerificationServer.CurrentState = model.MockVerificationState
	model.InspectExplain.CurrentState = model.InspectExplainState
	model.DegradedMode.CurrentState = model.DegradedModeState
	model.NoOverclaim.CurrentState = model.NoOverclaimState
	return model
}

func (s server) developerEcosystemValAStatusHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildDeveloperEcosystemValAStatus())
}

func (s server) developerEcosystemValAProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildDeveloperEcosystemValAProofs())
}

func buildDeveloperEcosystemValAStatus() developerEcosystemValAStatusResponse {
	model := buildDeveloperEcosystemValAModel()
	limitations := []string{
		"Val A implements IDE and local tooling core contracts only and does not ship marketplace IDE extensions, SDK runtime, repo config runtime, mock runtime, or plugin runtime.",
		"IDE, local advisory, validation harness, mock verification, and inspect/explain outputs remain advisory projections and do not approve deployment or mutate canonical evidence.",
	}
	return developerEcosystemValAStatusResponse{
		SchemaVersion: developerEcosystemValAStatusSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     developerEcosystemValAAllSurfaceRefs(),
		Limitations:   limitations,
	}
}

func buildDeveloperEcosystemValAProofs() developerEcosystemValAProofsResponse {
	model := buildDeveloperEcosystemValAModel()
	limitations := []string{
		"Val A defines IDE and local tooling core contracts only and does not implement production IDE marketplace publishing, SDK runtime, repo config runtime, plugin runtime, or Točka 9 work.",
		"Točka 8 remains not complete because later developer ecosystem waves are still required before any integrated closure can exist.",
		"IDE signals, local advisory, validation harness, mock verification, and inspect/explain outputs remain advisory only and cannot approve deployment, certify trust, or create canonical evidence.",
	}
	currentState := operability.EvaluateDeveloperEcosystemValAProofsState(model, limitations)
	return developerEcosystemValAProofsResponse{
		SchemaVersion:               developerEcosystemValAProofsSchema,
		GeneratedAt:                 publicSampleTime(),
		CurrentState:                currentState,
		Val0CurrentState:            model.Dependency.Val0CurrentState,
		Val0Point8State:             model.Dependency.Val0Point8State,
		Val0OutputClassification:    model.Dependency.Val0OutputClassification,
		Val0IDEAdvisoryState:        model.Dependency.Val0IDEAdvisoryState,
		Val0LocalProductionState:    model.Dependency.Val0LocalProductionState,
		Val0RepoPolicyBoundaryState: model.Dependency.Val0RepoPolicyBoundary,
		Val0PluginSafetyState:       model.Dependency.Val0PluginSafetyState,
		Val0PerformanceBudgetState:  model.Dependency.Val0PerformanceBudgetState,
		Val0DXMetricsState:          model.Dependency.Val0DXMetricsState,
		Val0NoOverclaimState:        model.Dependency.Val0NoOverclaimState,
		DependencyState:             model.DependencyState,
		IDEBaselineState:            model.IDEBaselineState,
		TrustFeedbackState:          model.TrustFeedbackState,
		CAVIVEXContextState:         model.CAVIVEXContextState,
		LocalAdvisoryState:          model.LocalAdvisoryState,
		ValidationHarnessState:      model.ValidationHarnessState,
		MockVerificationState:       model.MockVerificationState,
		InspectExplainState:         model.InspectExplainState,
		DegradedModeState:           model.DegradedModeState,
		NoOverclaimState:            model.NoOverclaimState,
		Point8State:                 model.Point8State,
		SupportedEditors:            model.IDEBaseline.SupportedEditors,
		TrustSignalClasses:          model.TrustFeedback.SignalClasses,
		ValidationClasses:           model.ValidationHarness.SupportedValidationClasses,
		SurfaceRefs:                 model.ProofSurfaceRefs,
		EvidenceRefs:                model.EvidenceRefs,
		BlockingReasons:             model.BlockingReasons,
		WhyPoint8NotPass: []string{
			"Val A is the IDE and local tooling core only and cannot return point_8_pass.",
			"Later Točka 8 waves remain required before any integrated developer ecosystem closure can exist.",
			"IDE, local advisory, validation harness, mock verification, and inspect/explain outputs remain advisory or projection-only and cannot become approval, certification, or canonical truth.",
		},
		Limitations:          limitations,
		ProjectionDisclaimer: developerEcosystemValAProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val A adds bounded VS Code and IntelliJ contracts, in-editor trust feedback, CAVI/VEX context, local advisory, validation harness, mock verification, inspect/explain, and degraded-mode contracts on top of accepted Val 0 discipline.",
			"Local and IDE tooling remain fail-closed, review-aware, freshness-aware, and explicitly separated from canonical truth, deployment approval, and production verifier equivalence.",
		},
	}
}
