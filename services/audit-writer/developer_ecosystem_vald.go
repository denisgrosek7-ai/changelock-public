package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	developerEcosystemValDStatusSchema = "point8.developer_ecosystem.vald.status.v1"
	developerEcosystemValDProofsSchema = "point8.developer_ecosystem.vald.proofs.v1"
)

type developerEcosystemValDStatusResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         operability.DeveloperEcosystemValDFinalGate `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type developerEcosystemValDProofsResponse struct {
	SchemaVersion                     string    `json:"schema_version"`
	GeneratedAt                       time.Time `json:"generated_at"`
	CurrentState                      string    `json:"current_state"`
	ValECompatibilityState            string    `json:"vale_compatibility_state"`
	ValEPoint7PassReason              string    `json:"vale_point_7_pass_reason"`
	Val0FoundationState               string    `json:"val0_foundation_state"`
	Val0CurrentState                  string    `json:"val0_current_state"`
	Val0Point8State                   string    `json:"val0_point_8_state"`
	Val0PerformanceBudgetDisciplineID string    `json:"val0_performance_budget_discipline_id"`
	ValAReadinessState                string    `json:"vala_readiness_state"`
	ValACurrentState                  string    `json:"vala_current_state"`
	ValAPoint8State                   string    `json:"vala_point_8_state"`
	ValADegradedModeState             string    `json:"vala_degraded_mode_state"`
	ValBReadinessState                string    `json:"valb_readiness_state"`
	ValBCurrentState                  string    `json:"valb_current_state"`
	ValBPoint8State                   string    `json:"valb_point_8_state"`
	ValBCompatibilityBehavior         string    `json:"valb_repo_config_compatibility_behavior"`
	ValBAPIVersionIdentity            string    `json:"valb_api_version_identity"`
	ValBAPICompatibilityWindow        string    `json:"valb_api_compatibility_window"`
	ValCReadinessState                string    `json:"valc_readiness_state"`
	ValCCurrentState                  string    `json:"valc_current_state"`
	ValCPoint8State                   string    `json:"valc_point_8_state"`
	ValCSandboxDisciplineID           string    `json:"valc_sandbox_discipline_id"`
	ValCSandboxVersion                string    `json:"valc_sandbox_version"`
	ValCPluginExecutionBudgetRef      string    `json:"valc_plugin_execution_budget_ref"`
	VerifyPolicyCICompatibilityState  string    `json:"verify_policy_ci_compatibility_state"`
	VerifyPolicyClassifierPath        string    `json:"verify_policy_classifier_path"`
	VerifyPolicyActionPath            string    `json:"verify_policy_action_path"`
	VerifyPolicyKyvernoVersion        string    `json:"verify_policy_kyverno_version"`
	IDELocalReadinessState            string    `json:"ide_local_readiness_state"`
	RepoSDKReadinessState             string    `json:"repo_sdk_readiness_state"`
	PluginExtensibilityReadinessState string    `json:"plugin_extensibility_readiness_state"`
	AdvisoryBoundaryState             string    `json:"advisory_boundary_state"`
	LocalMockNonEquivalenceState      string    `json:"local_mock_non_equivalence_state"`
	GovernanceNoBypassState           string    `json:"governance_no_bypass_state"`
	PerformanceVisibilityState        string    `json:"performance_visibility_state"`
	ExamplesNoCertificationState      string    `json:"examples_no_certification_state"`
	CleanRoomIPGuardrailState         string    `json:"clean_room_ip_guardrail_state"`
	NoOverclaimState                  string    `json:"no_overclaim_state"`
	FinalDeveloperEcosystemGateState  string    `json:"final_developer_ecosystem_gate_state"`
	Point8State                       string    `json:"point_8_state"`
	TriggerOnlyPrefixes               []string  `json:"trigger_only_prefixes,omitempty"`
	ManifestResourcePrefixes          []string  `json:"manifest_resource_prefixes,omitempty"`
	OptionOnlyArgs                    []string  `json:"option_only_args,omitempty"`
	SurfaceRefs                       []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                      []string  `json:"evidence_refs,omitempty"`
	BlockingReasons                   []string  `json:"blocking_reasons,omitempty"`
	WhyPoint8NotPass                  []string  `json:"why_point_8_not_pass,omitempty"`
	Limitations                       []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer              string    `json:"projection_disclaimer"`
	IntegrationSummary                []string  `json:"integration_summary,omitempty"`
}

func developerEcosystemValDAllSurfaceRefs() []string {
	return operability.DeveloperEcosystemValDProofSurfaceRefs()
}

func developerEcosystemValDProjectionDisclaimer() string {
	return "projection_only not_canonical_truth developer_ecosystem_vald advisory_projection final_developer_ecosystem_gate"
}

func buildDeveloperEcosystemValDValECompatibilityGate() operability.DeveloperEcosystemValDValECompatibilityGate {
	return operability.DeveloperEcosystemValDValECompatibilityGateModel()
}

func buildDeveloperEcosystemValDVal0FoundationSnapshot() operability.DeveloperEcosystemValDVal0FoundationSnapshot {
	return operability.DeveloperEcosystemValDVal0FoundationSnapshotModel()
}

func buildDeveloperEcosystemValDValAReadinessSnapshot() operability.DeveloperEcosystemValDValAReadinessSnapshot {
	return operability.DeveloperEcosystemValDValAReadinessSnapshotModel()
}

func buildDeveloperEcosystemValDValBReadinessSnapshot() operability.DeveloperEcosystemValDValBReadinessSnapshot {
	return operability.DeveloperEcosystemValDValBReadinessSnapshotModel()
}

func buildDeveloperEcosystemValDValCReadinessSnapshot() operability.DeveloperEcosystemValDValCReadinessSnapshot {
	return operability.DeveloperEcosystemValDValCReadinessSnapshotModel()
}

func buildDeveloperEcosystemValDVerifyPolicyCICompatibility() operability.DeveloperEcosystemValDVerifyPolicyCICompatibility {
	return operability.DeveloperEcosystemValDVerifyPolicyCICompatibilityModel()
}

func buildDeveloperEcosystemValDIDELocalReadinessGate() operability.DeveloperEcosystemValDIDELocalReadinessGate {
	return operability.DeveloperEcosystemValDIDELocalReadinessGateModel()
}

func buildDeveloperEcosystemValDRepoSDKReadinessGate() operability.DeveloperEcosystemValDRepoSDKReadinessGate {
	return operability.DeveloperEcosystemValDRepoSDKReadinessGateModel()
}

func buildDeveloperEcosystemValDPluginExtensibilityReadinessGate() operability.DeveloperEcosystemValDPluginExtensibilityReadinessGate {
	return operability.DeveloperEcosystemValDPluginExtensibilityReadinessGateModel()
}

func buildDeveloperEcosystemValDAdvisoryBoundaryGate() operability.DeveloperEcosystemValDAdvisoryBoundaryGate {
	return operability.DeveloperEcosystemValDAdvisoryBoundaryGateModel()
}

func buildDeveloperEcosystemValDLocalMockNonEquivalenceGate() operability.DeveloperEcosystemValDLocalMockNonEquivalenceGate {
	return operability.DeveloperEcosystemValDLocalMockNonEquivalenceGateModel()
}

func buildDeveloperEcosystemValDGovernanceNoBypassGate() operability.DeveloperEcosystemValDGovernanceNoBypassGate {
	return operability.DeveloperEcosystemValDGovernanceNoBypassGateModel()
}

func buildDeveloperEcosystemValDPerformanceVisibilityGate() operability.DeveloperEcosystemValDPerformanceVisibilityGate {
	return operability.DeveloperEcosystemValDPerformanceVisibilityGateModel()
}

func buildDeveloperEcosystemValDExamplesNoCertificationGate() operability.DeveloperEcosystemValDExamplesNoCertificationGate {
	return operability.DeveloperEcosystemValDExamplesNoCertificationGateModel()
}

func buildDeveloperEcosystemValDCleanRoomIPGuardrailEvidenceGate() operability.DeveloperEcosystemValDCleanRoomIPGuardrailEvidenceGate {
	return operability.DeveloperEcosystemValDCleanRoomIPGuardrailEvidenceGateModel()
}

func buildDeveloperEcosystemValDNoOverclaimGate() operability.DeveloperEcosystemValDNoOverclaimGate {
	return operability.DeveloperEcosystemValDNoOverclaimGateModel()
}

func buildDeveloperEcosystemValDModel() operability.DeveloperEcosystemValDFinalGate {
	model := operability.DeveloperEcosystemValDFinalGateModel()
	model.ValECompatibility = buildDeveloperEcosystemValDValECompatibilityGate()
	model.Val0Foundation = buildDeveloperEcosystemValDVal0FoundationSnapshot()
	model.ValAReadiness = buildDeveloperEcosystemValDValAReadinessSnapshot()
	model.ValBReadiness = buildDeveloperEcosystemValDValBReadinessSnapshot()
	model.ValCReadiness = buildDeveloperEcosystemValDValCReadinessSnapshot()
	model.VerifyPolicyCICompatibility = buildDeveloperEcosystemValDVerifyPolicyCICompatibility()
	model.IDELocalReadiness = buildDeveloperEcosystemValDIDELocalReadinessGate()
	model.RepoSDKReadiness = buildDeveloperEcosystemValDRepoSDKReadinessGate()
	model.PluginExtensibilityReadiness = buildDeveloperEcosystemValDPluginExtensibilityReadinessGate()
	model.AdvisoryBoundary = buildDeveloperEcosystemValDAdvisoryBoundaryGate()
	model.LocalMockNonEquivalence = buildDeveloperEcosystemValDLocalMockNonEquivalenceGate()
	model.GovernanceNoBypass = buildDeveloperEcosystemValDGovernanceNoBypassGate()
	model.PerformanceVisibility = buildDeveloperEcosystemValDPerformanceVisibilityGate()
	model.ExamplesNoCertification = buildDeveloperEcosystemValDExamplesNoCertificationGate()
	model.CleanRoomIPGuardrail = buildDeveloperEcosystemValDCleanRoomIPGuardrailEvidenceGate()
	model.NoOverclaim = buildDeveloperEcosystemValDNoOverclaimGate()
	model = operability.ComputeDeveloperEcosystemValDFinalGate(model)

	model.ValECompatibility.CurrentState = model.ValECompatibilityState
	model.VerifyPolicyCICompatibility.CurrentState = model.VerifyPolicyCICompatibilityState
	model.IDELocalReadiness.CurrentState = model.IDELocalReadinessState
	model.RepoSDKReadiness.CurrentState = model.RepoSDKReadinessState
	model.PluginExtensibilityReadiness.CurrentState = model.PluginExtensibilityReadinessState
	model.AdvisoryBoundary.CurrentState = model.AdvisoryBoundaryState
	model.LocalMockNonEquivalence.CurrentState = model.LocalMockNonEquivalenceState
	model.GovernanceNoBypass.CurrentState = model.GovernanceNoBypassState
	model.PerformanceVisibility.CurrentState = model.PerformanceVisibilityState
	model.ExamplesNoCertification.CurrentState = model.ExamplesNoCertificationState
	model.CleanRoomIPGuardrail.CurrentState = model.CleanRoomIPGuardrailState
	model.NoOverclaim.CurrentState = model.NoOverclaimState
	model.FinalDeveloperEcosystemGate.CurrentState = model.FinalDeveloperEcosystemGateState
	return model
}

func (s server) developerEcosystemValDStatusHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildDeveloperEcosystemValDStatus())
}

func (s server) developerEcosystemValDProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildDeveloperEcosystemValDProofs())
}

func buildDeveloperEcosystemValDStatus() developerEcosystemValDStatusResponse {
	model := buildDeveloperEcosystemValDModel()
	limitations := []string{
		"Val D implements the final developer ecosystem gate only and does not mark Točka 8 complete, return point_8_pass, or implement Točka 9.",
		"IDE, local, repo, SDK, plugin, examples, diagnostics, and developer metrics remain advisory or projection-only surfaces and do not approve deployment, certify trust, create canonical truth, or mutate canonical evidence.",
		"Clean-room and IP evidence remains a bounded static repo guardrail only and is not legal, patent, regulator, or formal opinion certification.",
	}
	return developerEcosystemValDStatusResponse{
		SchemaVersion: developerEcosystemValDStatusSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     developerEcosystemValDAllSurfaceRefs(),
		Limitations:   limitations,
	}
}

func buildDeveloperEcosystemValDProofs() developerEcosystemValDProofsResponse {
	model := buildDeveloperEcosystemValDModel()
	limitations := []string{
		"Val D is the final developer ecosystem gate only and cannot return point_8_pass or make Točka 8 complete.",
		"Integrated closure still requires Val E; Val D active is readiness consistency, not deployment approval, certification, production approval, or canonical truth.",
		"Clean-room and IP guardrail evidence is a static bounded repo check and does not claim legal certification, patent clearance, regulator approval, or formal legal opinion.",
	}
	currentState := operability.EvaluateDeveloperEcosystemValDProofsState(model, limitations)
	return developerEcosystemValDProofsResponse{
		SchemaVersion:                     developerEcosystemValDProofsSchema,
		GeneratedAt:                       publicSampleTime(),
		CurrentState:                      currentState,
		ValECompatibilityState:            model.ValECompatibilityState,
		ValEPoint7PassReason:              model.ValECompatibility.Point7PassReason,
		Val0FoundationState:               model.Val0FoundationState,
		Val0CurrentState:                  model.Val0Foundation.CurrentState,
		Val0Point8State:                   model.Val0Foundation.Point8State,
		Val0PerformanceBudgetDisciplineID: model.Val0Foundation.PerformanceBudgetDiscipline,
		ValAReadinessState:                model.ValAReadinessState,
		ValACurrentState:                  model.ValAReadiness.CurrentState,
		ValAPoint8State:                   model.ValAReadiness.Point8State,
		ValADegradedModeState:             model.ValAReadiness.DegradedModeState,
		ValBReadinessState:                model.ValBReadinessState,
		ValBCurrentState:                  model.ValBReadiness.CurrentState,
		ValBPoint8State:                   model.ValBReadiness.Point8State,
		ValBCompatibilityBehavior:         model.ValBReadiness.RepoConfigCompatibilityBehavior,
		ValBAPIVersionIdentity:            model.ValBReadiness.APIVersionIdentity,
		ValBAPICompatibilityWindow:        model.ValBReadiness.APICompatibilityWindow,
		ValCReadinessState:                model.ValCReadinessState,
		ValCCurrentState:                  model.ValCReadiness.CurrentState,
		ValCPoint8State:                   model.ValCReadiness.Point8State,
		ValCSandboxDisciplineID:           model.ValCReadiness.SandboxDisciplineID,
		ValCSandboxVersion:                model.ValCReadiness.SandboxVersion,
		ValCPluginExecutionBudgetRef:      model.ValCReadiness.PluginExecutionBudgetRef,
		VerifyPolicyCICompatibilityState:  model.VerifyPolicyCICompatibilityState,
		VerifyPolicyClassifierPath:        model.VerifyPolicyCICompatibility.ClassifierScriptPath,
		VerifyPolicyActionPath:            model.VerifyPolicyCICompatibility.ShiftLeftActionPath,
		VerifyPolicyKyvernoVersion:        model.VerifyPolicyCICompatibility.KyvernoVersion,
		IDELocalReadinessState:            model.IDELocalReadinessState,
		RepoSDKReadinessState:             model.RepoSDKReadinessState,
		PluginExtensibilityReadinessState: model.PluginExtensibilityReadinessState,
		AdvisoryBoundaryState:             model.AdvisoryBoundaryState,
		LocalMockNonEquivalenceState:      model.LocalMockNonEquivalenceState,
		GovernanceNoBypassState:           model.GovernanceNoBypassState,
		PerformanceVisibilityState:        model.PerformanceVisibilityState,
		ExamplesNoCertificationState:      model.ExamplesNoCertificationState,
		CleanRoomIPGuardrailState:         model.CleanRoomIPGuardrailState,
		NoOverclaimState:                  model.NoOverclaimState,
		FinalDeveloperEcosystemGateState:  model.FinalDeveloperEcosystemGateState,
		Point8State:                       model.Point8State,
		TriggerOnlyPrefixes:               model.VerifyPolicyCICompatibility.TriggerOnlyPrefixes,
		ManifestResourcePrefixes:          model.VerifyPolicyCICompatibility.ManifestResourcePrefixes,
		OptionOnlyArgs:                    model.VerifyPolicyCICompatibility.OptionOnlyArgs,
		SurfaceRefs:                       model.ProofSurfaceRefs,
		EvidenceRefs:                      model.EvidenceRefs,
		BlockingReasons:                   model.BlockingReasons,
		WhyPoint8NotPass: []string{
			"Val D is the final developer ecosystem gate only and cannot return point_8_pass.",
			"Integrated closure still explicitly requires Val E even when Val D final gate is active.",
			"Developer ecosystem outputs remain advisory, projection-only, and governance-bound rather than deployment approval, certification, or canonical truth.",
		},
		Limitations:          limitations,
		ProjectionDisclaimer: developerEcosystemValDProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val D verifies that patched Val E compatibility, accepted Val 0-A-B-C waves, and verify-policy or shift-left CI behavior remain aligned without declaring Točka 8 complete.",
			"Val D final gate active means the developer ecosystem expansion is internally consistent and bounded, but integrated closure still requires Val E and point_8_pass remains unavailable here.",
		},
	}
}
