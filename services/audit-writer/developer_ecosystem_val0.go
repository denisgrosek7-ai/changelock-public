package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	developerEcosystemVal0StatusSchema = "point8.developer_ecosystem.val0.status.v1"
	developerEcosystemVal0ProofsSchema = "point8.developer_ecosystem.val0.proofs.v1"
)

type developerEcosystemVal0StatusResponse struct {
	SchemaVersion string                                       `json:"schema_version"`
	GeneratedAt   time.Time                                    `json:"generated_at"`
	CurrentState  string                                       `json:"current_state"`
	Model         operability.DeveloperEcosystemVal0Foundation `json:"model"`
	RouteRefs     []string                                     `json:"route_refs,omitempty"`
	Limitations   []string                                     `json:"limitations,omitempty"`
}

type developerEcosystemVal0ProofsResponse struct {
	SchemaVersion              string    `json:"schema_version"`
	GeneratedAt                time.Time `json:"generated_at"`
	CurrentState               string    `json:"current_state"`
	Point6State                string    `json:"point_6_state"`
	Point7State                string    `json:"point_7_state"`
	Point7ClosureState         string    `json:"point_7_closure_state"`
	Point7PrerequisiteState    string    `json:"point_7_prerequisite_state"`
	Point7InvariantState       string    `json:"point_7_invariant_state"`
	Point7ProofSurfaceState    string    `json:"point_7_proof_surface_state"`
	Point7EvidenceQualityState string    `json:"point_7_evidence_quality_state"`
	Point7NoOverclaimState     string    `json:"point_7_no_overclaim_state"`
	Point7PassRuleState        string    `json:"point_7_pass_rule_state"`
	Point7PassAllowed          bool      `json:"point_7_pass_allowed"`
	DependencyState            string    `json:"dependency_state"`
	OutputClassificationState  string    `json:"output_classification_state"`
	IDEAdvisoryState           string    `json:"ide_advisory_state"`
	LocalProductionState       string    `json:"local_production_state"`
	RepoPolicyBoundaryState    string    `json:"repo_policy_boundary_state"`
	PluginSafetyState          string    `json:"plugin_safety_state"`
	PerformanceBudgetState     string    `json:"performance_budget_state"`
	DXMetricsState             string    `json:"dx_metrics_state"`
	NoOverclaimState           string    `json:"no_overclaim_state"`
	Point8State                string    `json:"point_8_state"`
	ClassifiedSurfaces         []string  `json:"classified_surfaces,omitempty"`
	OutputClasses              []string  `json:"output_classes,omitempty"`
	DXMetricNames              []string  `json:"dx_metric_names,omitempty"`
	SurfaceRefs                []string  `json:"surface_refs,omitempty"`
	EvidenceRefs               []string  `json:"evidence_refs,omitempty"`
	BlockingReasons            []string  `json:"blocking_reasons,omitempty"`
	WhyPoint8NotPass           []string  `json:"why_point_8_not_pass,omitempty"`
	Limitations                []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer       string    `json:"projection_disclaimer"`
	IntegrationSummary         []string  `json:"integration_summary,omitempty"`
}

func developerEcosystemVal0AllSurfaceRefs() []string {
	return operability.DeveloperEcosystemVal0ProofSurfaceRefs()
}

func developerEcosystemVal0ProjectionDisclaimer() string {
	return "projection_only not_canonical_truth developer_ecosystem_discipline_foundation advisory_projection"
}

func developerEcosystemVal0EvidenceRefs() []string {
	return operability.DeveloperEcosystemVal0ProofEvidenceRefs()
}

func buildDeveloperEcosystemVal0DependencySnapshot() operability.DeveloperEcosystemVal0DependencySnapshot {
	valE := buildVerifierEcosystemValEProofs()
	return operability.DeveloperEcosystemVal0DependencySnapshot{
		Point6State:                valE.Point6State,
		Point7State:                valE.Point7State,
		Point7ClosureState:         valE.ValEState,
		Point7PrerequisiteState:    valE.ClosurePrerequisiteState,
		Point7InvariantState:       valE.ClosureInvariantState,
		Point7ProofSurfaceState:    valE.ProofSurfaceState,
		Point7EvidenceQualityState: valE.EvidenceQualityState,
		Point7NoOverclaimState:     valE.NoOverclaimState,
		Point7PassRuleState:        valE.PassRuleState,
		Point7PassAllowed:          valE.Point7PassAllowed,
	}
}

func buildDeveloperEcosystemVal0FoundationModel() operability.DeveloperEcosystemVal0Foundation {
	model := operability.DeveloperEcosystemVal0FoundationModel()
	model.Dependency = buildDeveloperEcosystemVal0DependencySnapshot()
	model = operability.ComputeDeveloperEcosystemVal0Foundation(model)
	model.OutputClassification.CurrentState = model.OutputClassificationState
	model.IDEAdvisory.CurrentState = model.IDEAdvisoryState
	model.LocalProduction.CurrentState = model.LocalProductionState
	model.RepoPolicyBoundary.CurrentState = model.RepoPolicyBoundaryState
	model.PluginSafety.CurrentState = model.PluginSafetyState
	model.PerformanceBudget.CurrentState = model.PerformanceBudgetState
	model.DXMetrics.CurrentState = model.DXMetricsState
	model.NoOverclaim.CurrentState = model.NoOverclaimState
	return model
}

func (s server) developerEcosystemVal0StatusHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildDeveloperEcosystemVal0Status())
}

func (s server) developerEcosystemVal0ProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildDeveloperEcosystemVal0Proofs())
}

func buildDeveloperEcosystemVal0Status() developerEcosystemVal0StatusResponse {
	model := buildDeveloperEcosystemVal0FoundationModel()
	limitations := []string{
		"Val 0 defines developer discipline foundation only and does not implement IDE extensions, SDK packages, repo config parsers, mock verification runtime, or plugin execution runtime.",
		"Developer tooling remains advisory or projection-only and cannot approve deployment, certify trust, mutate canonical evidence, or create policy authority.",
	}
	return developerEcosystemVal0StatusResponse{
		SchemaVersion: developerEcosystemVal0StatusSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     developerEcosystemVal0AllSurfaceRefs(),
		Limitations:   limitations,
	}
}

func buildDeveloperEcosystemVal0Proofs() developerEcosystemVal0ProofsResponse {
	model := buildDeveloperEcosystemVal0FoundationModel()
	limitations := []string{
		"Val 0 is only the developer discipline foundation and does not implement actual IDE, SDK, repo config runtime, mock runtime, or plugin runtime execution.",
		"Točka 8 remains not complete because later waves are still required before any integrated developer ecosystem closure can exist.",
		"Developer tooling outputs remain advisory and cannot approve deployment, certify trust, override enterprise governance, or mutate canonical evidence.",
	}
	currentState := operability.EvaluateDeveloperEcosystemVal0ProofsState(model, limitations)
	return developerEcosystemVal0ProofsResponse{
		SchemaVersion:              developerEcosystemVal0ProofsSchema,
		GeneratedAt:                publicSampleTime(),
		CurrentState:               currentState,
		Point6State:                model.Dependency.Point6State,
		Point7State:                model.Dependency.Point7State,
		Point7ClosureState:         model.Dependency.Point7ClosureState,
		Point7PrerequisiteState:    model.Dependency.Point7PrerequisiteState,
		Point7InvariantState:       model.Dependency.Point7InvariantState,
		Point7ProofSurfaceState:    model.Dependency.Point7ProofSurfaceState,
		Point7EvidenceQualityState: model.Dependency.Point7EvidenceQualityState,
		Point7NoOverclaimState:     model.Dependency.Point7NoOverclaimState,
		Point7PassRuleState:        model.Dependency.Point7PassRuleState,
		Point7PassAllowed:          model.Dependency.Point7PassAllowed,
		DependencyState:            model.DependencyState,
		OutputClassificationState:  model.OutputClassificationState,
		IDEAdvisoryState:           model.IDEAdvisoryState,
		LocalProductionState:       model.LocalProductionState,
		RepoPolicyBoundaryState:    model.RepoPolicyBoundaryState,
		PluginSafetyState:          model.PluginSafetyState,
		PerformanceBudgetState:     model.PerformanceBudgetState,
		DXMetricsState:             model.DXMetricsState,
		NoOverclaimState:           model.NoOverclaimState,
		Point8State:                model.Point8State,
		ClassifiedSurfaces:         model.OutputClassification.ClassifiedSurfaceKinds,
		OutputClasses:              model.OutputClassification.AllowedOutputClasses,
		DXMetricNames:              model.DXMetrics.MetricNames,
		SurfaceRefs:                model.ProofSurfaceRefs,
		EvidenceRefs:               model.EvidenceRefs,
		BlockingReasons:            model.BlockingReasons,
		WhyPoint8NotPass: []string{
			"Val 0 is the developer discipline foundation only and cannot return point_8_pass.",
			"Later Točka 8 waves remain required before any integrated developer ecosystem closure can exist.",
			"Developer tooling outputs remain advisory or projection-only and cannot become approval, certification, or canonical truth.",
		},
		Limitations:          limitations,
		ProjectionDisclaimer: developerEcosystemVal0ProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val 0 establishes output classification, IDE advisory, local-vs-production, repo policy boundary, plugin safety, performance budget, DX metrics, and no-overclaim discipline for the developer ecosystem expansion.",
			"Developer tooling remains fail-closed, governance-safe, evidence-linked where applicable, and bounded away from canonical evidence mutation or deployment approval authority.",
		},
	}
}
