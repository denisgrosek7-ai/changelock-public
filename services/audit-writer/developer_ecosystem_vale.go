package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	developerEcosystemValEClosureSchema = "point8.developer_ecosystem.vale.closure.v1"
	developerEcosystemValEProofsSchema  = "point8.developer_ecosystem.vale.proofs.v1"
)

type developerEcosystemValEClosureResponse struct {
	SchemaVersion string                                              `json:"schema_version"`
	GeneratedAt   time.Time                                           `json:"generated_at"`
	CurrentState  string                                              `json:"current_state"`
	Model         operability.DeveloperEcosystemValEIntegratedClosure `json:"model"`
	RouteRefs     []string                                            `json:"route_refs,omitempty"`
	Limitations   []string                                            `json:"limitations,omitempty"`
}

type developerEcosystemValEProofsResponse struct {
	SchemaVersion                     string    `json:"schema_version"`
	GeneratedAt                       time.Time `json:"generated_at"`
	CurrentState                      string    `json:"current_state"`
	Point8State                       string    `json:"point_8_state"`
	Point8PassAllowed                 bool      `json:"point_8_pass_allowed"`
	Point8PassReason                  string    `json:"point_8_pass_reason"`
	ClosureState                      string    `json:"closure_state"`
	ValECompatibilityState            string    `json:"tocka7_vale_compatibility_state"`
	Val0SourceState                   string    `json:"val0_source_state"`
	ValASourceState                   string    `json:"vala_source_state"`
	ValBSourceState                   string    `json:"valb_source_state"`
	ValCSourceState                   string    `json:"valc_source_state"`
	ValDSourceState                   string    `json:"vald_source_state"`
	DependencyClosureState            string    `json:"dependency_closure_state"`
	CrossWaveInvariantState           string    `json:"cross_wave_invariant_state"`
	ProofSurfaceState                 string    `json:"proof_surface_state"`
	EvidenceQualityState              string    `json:"evidence_quality_state"`
	AdvisoryBoundaryState             string    `json:"advisory_boundary_state"`
	LocalMockNonEquivalenceState      string    `json:"local_mock_non_equivalence_state"`
	RepoSDKGovernanceBoundaryState    string    `json:"repo_sdk_governance_boundary_state"`
	PluginExtensibilityBoundaryState  string    `json:"plugin_extensibility_boundary_state"`
	VerifyPolicyCICompatibilityState  string    `json:"verify_policy_ci_compatibility_state"`
	CleanRoomIPGuardrailState         string    `json:"clean_room_ip_guardrail_state"`
	NoOverclaimState                  string    `json:"no_overclaim_state"`
	FinalPassRuleState                string    `json:"final_pass_rule_state"`
	Tocka7Point7PassReason            string    `json:"tocka7_point_7_pass_reason"`
	ValDFinalDeveloperEcosystemState  string    `json:"vald_final_developer_ecosystem_gate_state"`
	ValBCompatibilityBehavior         string    `json:"valb_repo_config_compatibility_behavior"`
	ValBAPIVersionIdentity            string    `json:"valb_api_version_identity"`
	ValBAPICompatibilityWindow        string    `json:"valb_api_compatibility_window"`
	ValCSandboxDisciplineID           string    `json:"valc_sandbox_discipline_id"`
	ValCSandboxVersion                string    `json:"valc_sandbox_version"`
	ValCPluginExecutionBudgetRef      string    `json:"valc_plugin_execution_budget_ref"`
	Val0PerformanceBudgetDisciplineID string    `json:"val0_performance_budget_discipline_id"`
	VerifyPolicyClassifierPath        string    `json:"verify_policy_classifier_path"`
	VerifyPolicyActionPath            string    `json:"verify_policy_action_path"`
	VerifyPolicyKyvernoVersion        string    `json:"verify_policy_kyverno_version"`
	SurfaceRefs                       []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                      []string  `json:"evidence_refs,omitempty"`
	BlockingReasons                   []string  `json:"blocking_reasons,omitempty"`
	WhyPoint8Pass                     []string  `json:"why_point_8_pass,omitempty"`
	Limitations                       []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer              string    `json:"projection_disclaimer"`
}

func developerEcosystemValEAllSurfaceRefs() []string {
	return operability.DeveloperEcosystemValEProofSurfaceRefs()
}

func developerEcosystemValEProjectionDisclaimer() string {
	return "projection_only not_canonical_truth developer_ecosystem_vale advisory_projection integrated_closure"
}

func buildDeveloperEcosystemValEModel() operability.DeveloperEcosystemValEIntegratedClosure {
	model := operability.DeveloperEcosystemValEIntegratedClosureModel()
	model = operability.ComputeDeveloperEcosystemValEClosure(model)
	return model
}

func (s server) developerEcosystemValEClosureHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildDeveloperEcosystemValEClosure())
}

func (s server) developerEcosystemValEProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildDeveloperEcosystemValEProofs())
}

func buildDeveloperEcosystemValEClosure() developerEcosystemValEClosureResponse {
	model := buildDeveloperEcosystemValEModel()
	limitations := []string{
		"Val E is the integrated closure for Točka 8 only and does not implement Točka 9 or any new runtime functionality beyond closure proof and status wiring.",
		"Developer ecosystem outputs remain advisory, projection-only, or developer-assist surfaces unless backed by canonical evidence.",
		"Clean-room and IP guardrail evidence remains a static bounded repo check only and does not claim legal certification, patent clearance, regulator approval, or formal legal opinion.",
	}
	return developerEcosystemValEClosureResponse{
		SchemaVersion: developerEcosystemValEClosureSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     developerEcosystemValEAllSurfaceRefs(),
		Limitations:   limitations,
	}
}

func buildDeveloperEcosystemValEProofs() developerEcosystemValEProofsResponse {
	model := buildDeveloperEcosystemValEModel()
	limitations := []string{
		"Only Val E may return point_8_pass and Točka 8 becomes complete only when the Val E final pass rule is active.",
		"Val 0 through Val D remain prerequisites and cannot close Točka 8 on their own.",
		"verify-policy / shift-left compatibility and clean-room/IP evidence are operational or static guardrail evidence only, not deployment approval, certification, production approval, or legal/IP certification.",
	}
	whyPoint8Pass := []string{
		"Val E returns point_8_pass only after actual Val 0 through Val D source states, exact proof surfaces, exact evidence refs, fail-closed cross-wave invariants, and no-overclaim closure all remain active.",
		"Local, mock, repo, SDK, plugin, examples, diagnostics, and developer-facing outputs remain advisory and non-canonical even when point_8_pass is active.",
	}
	return developerEcosystemValEProofsResponse{
		SchemaVersion:                     developerEcosystemValEProofsSchema,
		GeneratedAt:                       publicSampleTime(),
		CurrentState:                      model.CurrentState,
		Point8State:                       model.Point8State,
		Point8PassAllowed:                 model.Point8PassAllowed,
		Point8PassReason:                  model.Point8PassReason,
		ClosureState:                      model.ClosureState,
		ValECompatibilityState:            model.ValECompatibilityState,
		Val0SourceState:                   model.Val0SourceState,
		ValASourceState:                   model.ValASourceState,
		ValBSourceState:                   model.ValBSourceState,
		ValCSourceState:                   model.ValCSourceState,
		ValDSourceState:                   model.ValDSourceState,
		DependencyClosureState:            model.DependencyClosureState,
		CrossWaveInvariantState:           model.CrossWaveInvariantState,
		ProofSurfaceState:                 model.ProofSurfaceState,
		EvidenceQualityState:              model.EvidenceQualityState,
		AdvisoryBoundaryState:             model.AdvisoryBoundaryState,
		LocalMockNonEquivalenceState:      model.LocalMockNonEquivalenceState,
		RepoSDKGovernanceBoundaryState:    model.RepoSDKGovernanceBoundaryState,
		PluginExtensibilityBoundaryState:  model.PluginExtensibilityBoundaryState,
		VerifyPolicyCICompatibilityState:  model.VerifyPolicyCICompatibilityState,
		CleanRoomIPGuardrailState:         model.CleanRoomIPGuardrailState,
		NoOverclaimState:                  model.NoOverclaimState,
		FinalPassRuleState:                model.FinalPassRuleState,
		Tocka7Point7PassReason:            model.Tocka7ValECompatibility.Point7PassReason,
		ValDFinalDeveloperEcosystemState:  model.ValDSource.FinalDeveloperEcosystemGateState,
		ValBCompatibilityBehavior:         model.ValBSource.RepoConfigCompatibilityBehavior,
		ValBAPIVersionIdentity:            model.ValBSource.APIVersionIdentity,
		ValBAPICompatibilityWindow:        model.ValBSource.APICompatibilityWindow,
		ValCSandboxDisciplineID:           model.ValCSource.SandboxDisciplineID,
		ValCSandboxVersion:                model.ValCSource.SandboxVersion,
		ValCPluginExecutionBudgetRef:      model.ValCSource.PluginExecutionBudgetRef,
		Val0PerformanceBudgetDisciplineID: model.Val0Source.PerformanceBudgetDiscipline,
		VerifyPolicyClassifierPath:        model.VerifyPolicyCICompatibility.ClassifierScriptPath,
		VerifyPolicyActionPath:            model.VerifyPolicyCICompatibility.ShiftLeftActionPath,
		VerifyPolicyKyvernoVersion:        model.VerifyPolicyCICompatibility.KyvernoVersion,
		SurfaceRefs:                       model.ProofSurfaceRefs,
		EvidenceRefs:                      model.EvidenceRefs,
		BlockingReasons:                   model.BlockingReasons,
		WhyPoint8Pass:                     whyPoint8Pass,
		Limitations:                       limitations,
		ProjectionDisclaimer:              developerEcosystemValEProjectionDisclaimer(),
	}
}
