package main

import (
	"context"
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	productionUsabilityValEDependencyClosureSchema       = "point4.production_usability.vale.dependency_closure.v1"
	productionUsabilityValECoherenceReviewSchema         = "point4.production_usability.vale.coherence_review.v1"
	productionUsabilityValEPassRuleSchema                = "point4.production_usability.vale.pass_rule.v1"
	productionUsabilityValECanonicalBoundaryReviewSchema = "point4.production_usability.vale.canonical_truth_boundary_review.v1"
	productionUsabilityValERedactionExportReviewSchema   = "point4.production_usability.vale.redaction_export_review.v1"
	productionUsabilityValESupportabilityRecoverySchema  = "point4.production_usability.vale.supportability_recovery_review.v1"
	productionUsabilityValERegressionClosureSchema       = "point4.production_usability.vale.regression_closure.v1"
	productionUsabilityValEProofsSchema                  = "point4.production_usability.vale.proofs.v1"
)

type productionUsabilityValEDependencyClosureResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	GeneratedAt   time.Time                               `json:"generated_at"`
	CurrentState  string                                  `json:"current_state"`
	Model         operability.IntegratedDependencyClosure `json:"model"`
	RouteRefs     []string                                `json:"route_refs,omitempty"`
	Limitations   []string                                `json:"limitations,omitempty"`
}

type productionUsabilityValECoherenceReviewResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.CrossValCoherenceReview `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type productionUsabilityValEPassRuleResponse struct {
	SchemaVersion string                               `json:"schema_version"`
	GeneratedAt   time.Time                            `json:"generated_at"`
	CurrentState  string                               `json:"current_state"`
	Model         operability.Point4IntegratedPassRule `json:"model"`
	RouteRefs     []string                             `json:"route_refs,omitempty"`
	Limitations   []string                             `json:"limitations,omitempty"`
}

type productionUsabilityValECanonicalBoundaryReviewResponse struct {
	SchemaVersion string                                             `json:"schema_version"`
	GeneratedAt   time.Time                                          `json:"generated_at"`
	CurrentState  string                                             `json:"current_state"`
	Model         operability.IntegratedCanonicalTruthBoundaryReview `json:"model"`
	RouteRefs     []string                                           `json:"route_refs,omitempty"`
	Limitations   []string                                           `json:"limitations,omitempty"`
}

type productionUsabilityValERedactionExportReviewResponse struct {
	SchemaVersion string                                                `json:"schema_version"`
	GeneratedAt   time.Time                                             `json:"generated_at"`
	CurrentState  string                                                `json:"current_state"`
	Model         operability.IntegratedPermissionRedactionExportReview `json:"model"`
	RouteRefs     []string                                              `json:"route_refs,omitempty"`
	Limitations   []string                                              `json:"limitations,omitempty"`
}

type productionUsabilityValESupportabilityRecoveryResponse struct {
	SchemaVersion string                                             `json:"schema_version"`
	GeneratedAt   time.Time                                          `json:"generated_at"`
	CurrentState  string                                             `json:"current_state"`
	Model         operability.IntegratedSupportabilityRecoveryReview `json:"model"`
	RouteRefs     []string                                           `json:"route_refs,omitempty"`
	Limitations   []string                                           `json:"limitations,omitempty"`
}

type productionUsabilityValERegressionClosureResponse struct {
	SchemaVersion string                                           `json:"schema_version"`
	GeneratedAt   time.Time                                        `json:"generated_at"`
	CurrentState  string                                           `json:"current_state"`
	Model         operability.IntegratedUsabilityRegressionClosure `json:"model"`
	RouteRefs     []string                                         `json:"route_refs,omitempty"`
	Limitations   []string                                         `json:"limitations,omitempty"`
}

type productionUsabilityValEProofsResponse struct {
	SchemaVersion               string    `json:"schema_version"`
	GeneratedAt                 time.Time `json:"generated_at"`
	CurrentState                string    `json:"current_state"`
	Val0DependencyState         string    `json:"val_0_dependency_state"`
	Val0FoundationState         string    `json:"val_0_foundation_state"`
	ValADependencyState         string    `json:"val_a_dependency_state"`
	ValACoreState               string    `json:"val_a_core_state"`
	ValBDependencyState         string    `json:"val_b_dependency_state"`
	ValBResilienceState         string    `json:"val_b_resilience_state"`
	ValCDependencyState         string    `json:"val_c_dependency_state"`
	ValCSupportabilityState     string    `json:"val_c_supportability_state"`
	ValDDependencyState         string    `json:"val_d_dependency_state"`
	ValDFinalGateState          string    `json:"val_d_final_gate_state"`
	ValEState                   string    `json:"val_e_state"`
	DependencyClosureState      string    `json:"dependency_closure_state"`
	CoherenceReviewState        string    `json:"coherence_review_state"`
	Point4State                 string    `json:"point_4_state"`
	PassCriteriaMet             bool      `json:"pass_criteria_met"`
	PassBlockers                []string  `json:"pass_blockers,omitempty"`
	PassWarnings                []string  `json:"pass_warnings,omitempty"`
	PassLimitations             []string  `json:"pass_limitations,omitempty"`
	CanonicalTruthBoundaryState string    `json:"canonical_truth_boundary_state"`
	RedactionExportReviewState  string    `json:"redaction_export_review_state"`
	SupportabilityRecoveryState string    `json:"supportability_recovery_state"`
	RegressionClosureState      string    `json:"regression_closure_state"`
	EvidenceRefs                []string  `json:"evidence_refs,omitempty"`
	SurfaceRefs                 []string  `json:"surface_refs,omitempty"`
	Limitations                 []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer        string    `json:"projection_disclaimer"`
	IntegrationSummary          []string  `json:"integration_summary,omitempty"`
}

type productionUsabilityValEModels struct {
	val0 productionUsabilityVal0ProofsResponse
	valA productionUsabilityValAProofsResponse
	valB productionUsabilityValBProofsResponse
	valC productionUsabilityValCProofsResponse
	valD productionUsabilityValDProofsResponse

	dependencyClosure     operability.IntegratedDependencyClosure
	coherenceReview       operability.CrossValCoherenceReview
	passRule              operability.Point4IntegratedPassRule
	canonicalBoundary     operability.IntegratedCanonicalTruthBoundaryReview
	redactionExportReview operability.IntegratedPermissionRedactionExportReview
	supportabilityReview  operability.IntegratedSupportabilityRecoveryReview
	regressionClosure     operability.IntegratedUsabilityRegressionClosure

	dependencyClosureState     string
	coherenceReviewState       string
	passRuleState              string
	canonicalBoundaryState     string
	redactionExportReviewState string
	supportabilityReviewState  string
	regressionClosureState     string
	valEState                  string
	point4State                string

	surfaceRefs  []string
	evidenceRefs []string
	limitations  []string
}

func productionUsabilityValECheckedScopes() []string {
	return []string{
		operability.ProductionUsabilityVisibilityInternalAdmin,
		operability.ProductionUsabilityVisibilityOperator,
		operability.ProductionUsabilityVisibilityDeveloper,
		operability.ProductionUsabilityVisibilityPartner,
		operability.ProductionUsabilityVisibilityPublicSafe,
	}
}

func productionUsabilityValERequiredVals() []string {
	return []string{"val_0", "val_a", "val_b", "val_c", "val_d", "val_e"}
}

func productionUsabilityValEAllSurfaceRefs() []string {
	return []string{
		"/v1/production/usability-operability-recovery/val0/proofs",
		"/v1/production/usability-operability-recovery/vala/proofs",
		"/v1/production/usability-operability-recovery/valb/proofs",
		"/v1/production/usability-operability-recovery/valc/proofs",
		"/v1/production/usability-operability-recovery/vald/proofs",
		"/v1/production/usability-operability-recovery/vale/dependency-closure",
		"/v1/production/usability-operability-recovery/vale/coherence-review",
		"/v1/production/usability-operability-recovery/vale/pass-rule",
		"/v1/production/usability-operability-recovery/vale/canonical-truth-boundary",
		"/v1/production/usability-operability-recovery/vale/redaction-export-review",
		"/v1/production/usability-operability-recovery/vale/supportability-recovery-review",
		"/v1/production/usability-operability-recovery/vale/regression-closure",
		"/v1/production/usability-operability-recovery/vale/proofs",
	}
}

func productionUsabilityValEEvidenceRefs() []string {
	return []string{
		"val0_proofs",
		"vala_proofs",
		"valb_proofs",
		"valc_proofs",
		"vald_proofs",
		"dependency_closure_contract",
		"coherence_review_contract",
		"point4_pass_rule_contract",
		"canonical_truth_boundary_review_contract",
		"redaction_export_review_contract",
		"supportability_recovery_review_contract",
		"regression_closure_contract",
		"evidence_spine",
	}
}

func productionUsabilityValEProjectionDisclaimer() string {
	return "projection_only not_canonical_truth integrated_closure_summary"
}

func appendPrefixedLimitations(dst []string, prefix string, items []string) []string {
	for _, item := range items {
		if item == "" {
			continue
		}
		dst = append(dst, prefix+": "+item)
	}
	return dst
}

func collectIntegratedLimitations(val0 productionUsabilityVal0ProofsResponse, valA productionUsabilityValAProofsResponse, valB productionUsabilityValBProofsResponse, valC productionUsabilityValCProofsResponse, valD productionUsabilityValDProofsResponse) []string {
	limitations := []string{
		"Integrated closure remains a projection-only summary and does not replace canonical truth.",
		"Point 4 pass is allowed only through the active Val E integrated closure and remains fail-closed on prior val proof state.",
	}
	limitations = appendPrefixedLimitations(limitations, "val0", val0.Limitations)
	limitations = appendPrefixedLimitations(limitations, "vala", valA.Limitations)
	limitations = appendPrefixedLimitations(limitations, "valb", valB.Limitations)
	limitations = appendPrefixedLimitations(limitations, "valc", valC.Limitations)
	limitations = appendPrefixedLimitations(limitations, "vald", valD.Limitations)
	return limitations
}

func activeValNames(val0State, valAState, valBState, valCState, valDState, valEState string) []string {
	active := []string{}
	if val0State == operability.ProductionUsabilityVal0StateActive {
		active = append(active, "val_0")
	}
	if valAState == operability.ProductionUsabilityValAStateActive {
		active = append(active, "val_a")
	}
	if valBState == operability.ProductionUsabilityValBStateActive {
		active = append(active, "val_b")
	}
	if valCState == operability.ProductionUsabilityValCStateActive {
		active = append(active, "val_c")
	}
	if valDState == operability.ProductionUsabilityValDStateActive {
		active = append(active, "val_d")
	}
	if valEState == operability.ProductionUsabilityValEStateActive {
		active = append(active, "val_e")
	}
	return active
}

func classifyRequiredVals(val0State, valAState, valBState, valCState, valDState, valEState string) (missingVals, partialVals, unsupportedVals []string) {
	for _, item := range []struct {
		name  string
		state string
		want  string
	}{
		{name: "val_0", state: val0State, want: operability.ProductionUsabilityVal0StateActive},
		{name: "val_a", state: valAState, want: operability.ProductionUsabilityValAStateActive},
		{name: "val_b", state: valBState, want: operability.ProductionUsabilityValBStateActive},
		{name: "val_c", state: valCState, want: operability.ProductionUsabilityValCStateActive},
		{name: "val_d", state: valDState, want: operability.ProductionUsabilityValDStateActive},
		{name: "val_e", state: valEState, want: operability.ProductionUsabilityValEStateActive},
	} {
		switch {
		case item.state == "":
			missingVals = append(missingVals, item.name)
		case item.state == item.want:
		case item.state == operability.ProductionUsabilityDependencyUnsupported:
			unsupportedVals = append(unsupportedVals, item.name)
		default:
			partialVals = append(partialVals, item.name)
		}
	}
	return missingVals, partialVals, unsupportedVals
}

func buildProductionUsabilityValEDependencyClosureModel(val0 productionUsabilityVal0ProofsResponse, valA productionUsabilityValAProofsResponse, valB productionUsabilityValBProofsResponse, valC productionUsabilityValCProofsResponse, valD productionUsabilityValDProofsResponse) operability.IntegratedDependencyClosure {
	model := operability.ProductionUsabilityValEDependencyClosure()
	model.Val0State = val0.Val0State
	model.ValAState = valA.ValAState
	model.ValBState = valB.ValBState
	model.ValCState = valC.ValCState
	model.ValDState = valD.ValDState
	model.DependencyEvidenceRefs = []string{"val0_proofs", "vala_proofs", "valb_proofs", "valc_proofs", "vald_proofs"}
	model.DependencySurfaceRefs = []string{
		"/v1/production/usability-operability-recovery/val0/proofs",
		"/v1/production/usability-operability-recovery/vala/proofs",
		"/v1/production/usability-operability-recovery/valb/proofs",
		"/v1/production/usability-operability-recovery/valc/proofs",
		"/v1/production/usability-operability-recovery/vald/proofs",
		"/v1/production/usability-operability-recovery/vale/dependency-closure",
		"/v1/production/usability-operability-recovery/vale/proofs",
	}
	model.ProofStatesObserved = val0.CurrentState != "" && valA.CurrentState != "" && valB.CurrentState != "" && valC.CurrentState != "" && valD.CurrentState != ""

	if val0.CurrentState == "" || val0.Val0State == "" {
		model.MissingVals = append(model.MissingVals, "val_0")
	}
	if valA.CurrentState == "" || valA.ValAState == "" {
		model.MissingVals = append(model.MissingVals, "val_a")
	}
	if valB.CurrentState == "" || valB.ValBState == "" {
		model.MissingVals = append(model.MissingVals, "val_b")
	}
	if valC.CurrentState == "" || valC.ValCState == "" {
		model.MissingVals = append(model.MissingVals, "val_c")
	}
	if valD.CurrentState == "" || valD.ValDState == "" {
		model.MissingVals = append(model.MissingVals, "val_d")
	}

	if val0.Val0State != operability.ProductionUsabilityVal0StateActive && val0.Val0State != "" {
		model.InactiveVals = append(model.InactiveVals, "val_0")
	}
	if valA.ValAState != operability.ProductionUsabilityValAStateActive && valA.ValAState != "" {
		model.InactiveVals = append(model.InactiveVals, "val_a")
	}
	if valB.ValBState != operability.ProductionUsabilityValBStateActive && valB.ValBState != "" {
		model.InactiveVals = append(model.InactiveVals, "val_b")
	}
	if valC.ValCState != operability.ProductionUsabilityValCStateActive && valC.ValCState != "" {
		model.InactiveVals = append(model.InactiveVals, "val_c")
	}
	if valD.ValDState != operability.ProductionUsabilityValDStateActive && valD.ValDState != "" {
		model.InactiveVals = append(model.InactiveVals, "val_d")
	}

	if val0.CurrentState != "" && val0.CurrentState != operability.ProductionUsabilityVal0StateActive {
		model.InconsistentVals = append(model.InconsistentVals, "val_0.proofs_state")
	}
	if valA.CurrentState != "" && valA.CurrentState != operability.ProductionUsabilityValAStateActive {
		model.InconsistentVals = append(model.InconsistentVals, "val_a.proofs_state")
	}
	if valB.CurrentState != "" && valB.CurrentState != operability.ProductionUsabilityValBStateActive {
		model.InconsistentVals = append(model.InconsistentVals, "val_b.proofs_state")
	}
	if valC.CurrentState != "" && valC.CurrentState != operability.ProductionUsabilityValCStateActive {
		model.InconsistentVals = append(model.InconsistentVals, "val_c.proofs_state")
	}
	if valD.CurrentState != "" && valD.CurrentState != operability.ProductionUsabilityValDStateActive {
		model.InconsistentVals = append(model.InconsistentVals, "val_d.proofs_state")
	}
	if val0.Point4State == operability.ProductionUsabilityPoint4StatePass {
		model.InconsistentVals = append(model.InconsistentVals, "val_0.claims_point_4_pass")
	}
	if valA.Point4State == operability.ProductionUsabilityPoint4StatePass {
		model.InconsistentVals = append(model.InconsistentVals, "val_a.claims_point_4_pass")
	}
	if valB.Point4State == operability.ProductionUsabilityPoint4StatePass {
		model.InconsistentVals = append(model.InconsistentVals, "val_b.claims_point_4_pass")
	}
	if valC.Point4State == operability.ProductionUsabilityPoint4StatePass {
		model.InconsistentVals = append(model.InconsistentVals, "val_c.claims_point_4_pass")
	}
	if valD.Point4State == operability.ProductionUsabilityPoint4StatePass {
		model.InconsistentVals = append(model.InconsistentVals, "val_d.claims_point_4_pass")
	}

	switch {
	case len(model.MissingVals) > 0:
		model.DependencyStatus = operability.ProductionUsabilityDependencyIncomplete
	case len(model.InactiveVals) > 0:
		model.DependencyStatus = operability.ProductionUsabilityDependencyFail
	case len(model.InconsistentVals) > 0:
		model.DependencyStatus = operability.ProductionUsabilityDependencyPartial
	default:
		model.DependencyStatus = operability.ProductionUsabilityDependencyPass
	}
	return model
}

func buildProductionUsabilityValECoherenceReviewModel(val0 productionUsabilityVal0ProofsResponse, valA productionUsabilityValAProofsResponse, valB productionUsabilityValBProofsResponse, valC productionUsabilityValCProofsResponse, valD productionUsabilityValDProofsResponse) operability.CrossValCoherenceReview {
	model := operability.ProductionUsabilityValECrossValCoherenceReview()
	model.CarriedForwardLimitations = collectIntegratedLimitations(val0, valA, valB, valC, valD)
	model.EvidenceRefs = []string{"val0_proofs", "vala_proofs", "valb_proofs", "valc_proofs", "vald_proofs"}
	model.SurfaceRefs = []string{
		"/v1/production/usability-operability-recovery/val0/proofs",
		"/v1/production/usability-operability-recovery/vala/proofs",
		"/v1/production/usability-operability-recovery/valb/proofs",
		"/v1/production/usability-operability-recovery/valc/proofs",
		"/v1/production/usability-operability-recovery/vald/proofs",
		"/v1/production/usability-operability-recovery/vale/coherence-review",
		"/v1/production/usability-operability-recovery/vale/proofs",
	}
	model.Val0ContractsUsedByLaterVals = valA.Val0DependencyState == operability.ProductionUsabilityVal0StateActive &&
		valB.Val0DependencyState == operability.ProductionUsabilityVal0StateActive
	model.ValARespectedByLaterVals = valB.ValADependencyState == operability.ProductionUsabilityValAStateActive &&
		valC.ValADependencyState == operability.ProductionUsabilityValAStateActive &&
		valD.ValADependencyState == operability.ProductionUsabilityValAStateActive
	model.ValBRespectedByLaterVals = valC.ValBDependencyState == operability.ProductionUsabilityValBStateActive &&
		valD.ValBDependencyState == operability.ProductionUsabilityValBStateActive
	model.ValCRespectedByValD = valD.ValCDependencyState == operability.ProductionUsabilityValCStateActive
	model.ValDFinalGateCoversPriorVals = valD.ValDState == operability.ProductionUsabilityValDStateActive
	model.NoPriorValClaimsPoint4Pass = val0.Point4State != operability.ProductionUsabilityPoint4StatePass &&
		valA.Point4State != operability.ProductionUsabilityPoint4StatePass &&
		valB.Point4State != operability.ProductionUsabilityPoint4StatePass &&
		valC.Point4State != operability.ProductionUsabilityPoint4StatePass &&
		valD.Point4State != operability.ProductionUsabilityPoint4StatePass
	model.LimitationsCarriedForward = len(model.CarriedForwardLimitations) > 0

	if !model.Val0ContractsUsedByLaterVals {
		model.MissingLinks = append(model.MissingLinks,
			"val0.contracts->vala.config_explainability_core",
			"val0.contracts->valb.resilience_action_modes",
		)
	}
	if !model.ValARespectedByLaterVals {
		model.MissingLinks = append(model.MissingLinks,
			"vala.config_explainability->valb.ui_api_cli_resilience",
			"vala.config_explainability->valc.supportability_lifecycle",
		)
	}
	if !model.ValBRespectedByLaterVals {
		model.MissingLinks = append(model.MissingLinks, "valb.resilience->valc.supportability_lifecycle")
	}
	if !model.ValCRespectedByValD {
		model.MissingLinks = append(model.MissingLinks, "valc.supportability->vald.final_usability_gate")
	}
	if !model.ValDFinalGateCoversPriorVals {
		model.MissingLinks = append(model.MissingLinks, "vald.final_gate->vale.integrated_closure")
	}
	if !model.NoPriorValClaimsPoint4Pass {
		model.InconsistentLinks = append(model.InconsistentLinks, "prior_vals->point4_not_complete_until_vale")
	}
	if !model.LimitationsCarriedForward {
		model.MissingLinks = append(model.MissingLinks, "limitations->integrated_closure")
	}

	if len(model.MissingLinks) > 0 || len(model.InconsistentLinks) > 0 {
		model.CoherenceState = operability.ProductionUsabilityFinalGateBlocked
	} else {
		model.CoherenceState = operability.ProductionUsabilityFinalGatePass
	}
	return model
}

func buildProductionUsabilityValECanonicalBoundaryReviewModel(valA productionUsabilityValAProofsResponse, valB productionUsabilityValBProofsResponse, valC productionUsabilityValCProofsResponse, valD productionUsabilityValDProofsResponse) operability.IntegratedCanonicalTruthBoundaryReview {
	model := operability.ProductionUsabilityValECanonicalTruthBoundaryReview()
	model.EffectiveConfigProjectionOnly = valA.EffectiveConfigState == operability.ProductionUsabilityValAEffectiveConfigStateActive
	model.DryRunAuditProjectionOnly = valA.DryRunState == operability.ProductionUsabilityValADryRunStateActive
	model.UIAndCacheProjectionOnly = valB.UIDataResilienceState == operability.ProductionUsabilityValBUIDataResilienceStateActive &&
		valB.WindowingState == operability.ProductionUsabilityValBWindowingStateActive &&
		valB.ResultSemanticsState == operability.ProductionUsabilityValBResultSemanticsStateActive
	model.SupportProjectionOnly = valC.SupportBundleState == operability.ProductionUsabilityValCSupportBundleStateActive &&
		valC.DiagnosticsState == operability.ProductionUsabilityValCDiagnosticsStateActive &&
		valC.HealthSnapshotState == operability.ProductionUsabilityValCHealthSnapshotStateActive
	model.ValDGateProjectionOnly = valD.GovernanceBoundaryReviewState == operability.ProductionUsabilityValDGovernanceBoundaryReviewStateActive
	model.IntegratedSummaryProjectionOnly = true
	model.CheckedSurfaces = []string{
		"/v1/production/usability-operability-recovery/vala/effective-config",
		"/v1/production/usability-operability-recovery/vala/policy-dry-run",
		"/v1/production/usability-operability-recovery/valb/ui-data-resilience",
		"/v1/production/usability-operability-recovery/valb/windowing",
		"/v1/production/usability-operability-recovery/valc/support-bundle",
		"/v1/production/usability-operability-recovery/valc/diagnostics",
		"/v1/production/usability-operability-recovery/valc/health-snapshot",
		"/v1/production/usability-operability-recovery/vald/governance-boundary-review",
		"/v1/production/usability-operability-recovery/vale/proofs",
	}
	model.EvidenceRefs = []string{"evidence_spine", "vala_proofs", "valb_proofs", "valc_proofs", "vald_proofs"}
	if !model.EffectiveConfigProjectionOnly {
		model.ViolationSurfaces = append(model.ViolationSurfaces, "/v1/production/usability-operability-recovery/vala/effective-config")
	}
	if !model.DryRunAuditProjectionOnly {
		model.ViolationSurfaces = append(model.ViolationSurfaces, "/v1/production/usability-operability-recovery/vala/policy-dry-run")
	}
	if !model.UIAndCacheProjectionOnly {
		model.ViolationSurfaces = append(model.ViolationSurfaces, "/v1/production/usability-operability-recovery/valb/ui-data-resilience")
	}
	if !model.SupportProjectionOnly {
		model.ViolationSurfaces = append(model.ViolationSurfaces, "/v1/production/usability-operability-recovery/valc/support-bundle")
	}
	if !model.ValDGateProjectionOnly {
		model.ViolationSurfaces = append(model.ViolationSurfaces, "/v1/production/usability-operability-recovery/vald/governance-boundary-review")
	}
	if len(model.ViolationSurfaces) > 0 {
		model.BoundaryState = operability.ProductionUsabilityFinalGateBlocked
	} else {
		model.BoundaryState = operability.ProductionUsabilityFinalGatePass
	}
	return model
}

func buildProductionUsabilityValERedactionExportReviewModel(valA productionUsabilityValAProofsResponse, valC productionUsabilityValCProofsResponse, valD productionUsabilityValDProofsResponse) operability.IntegratedPermissionRedactionExportReview {
	model := operability.ProductionUsabilityValEIntegratedRedactionExportReview()
	model.CheckedScopes = productionUsabilityValECheckedScopes()
	model.ExportSafetyState = valC.RedactionExportSafetyState
	if valA.ExplainState != operability.ProductionUsabilityValAExplainStateActive ||
		valC.PermissionSupportFlowState != operability.ProductionUsabilityValCPermissionSupportStateActive ||
		valC.RedactionExportSafetyState != operability.ProductionUsabilityValCExportSafetyStateActive ||
		valD.RedactionReviewState != operability.ProductionUsabilityValDRedactionReviewStateActive {
		model.MissingScopes = append(model.MissingScopes, productionUsabilityValECheckedScopes()...)
	}
	model.HiddenMetadataRepresented = valC.PermissionSupportFlowState == operability.ProductionUsabilityValCPermissionSupportStateActive
	model.RedactedMetadataRepresented = valC.RedactionExportSafetyState == operability.ProductionUsabilityValCExportSafetyStateActive
	if len(model.MissingScopes) > 0 || len(model.UnsafeScopes) > 0 || !model.HiddenMetadataRepresented || !model.RedactedMetadataRepresented {
		model.RedactionState = operability.ProductionUsabilityFinalGateBlocked
	} else {
		model.RedactionState = operability.ProductionUsabilityFinalGatePass
	}
	return model
}

func buildProductionUsabilityValESupportabilityRecoveryReviewModel(val0 productionUsabilityVal0ProofsResponse, valA productionUsabilityValAProofsResponse, valB productionUsabilityValBProofsResponse, valC productionUsabilityValCProofsResponse, valD productionUsabilityValDProofsResponse) operability.IntegratedSupportabilityRecoveryReview {
	model := operability.ProductionUsabilityValESupportabilityRecoveryReview()
	model.ReadinessState = valC.ReadinessState
	model.DiagnosticsState = valC.DiagnosticsState
	model.RecoveryState = valD.RecoveryReviewState
	model.UpgradeAdvisoryState = valD.UpgradeRollbackReviewState
	model.FailedProofStateObserved = val0.CurrentState != operability.ProductionUsabilityVal0StateActive ||
		valA.CurrentState != operability.ProductionUsabilityValAStateActive ||
		valB.CurrentState != operability.ProductionUsabilityValBStateActive ||
		valC.CurrentState != operability.ProductionUsabilityValCStateActive ||
		valD.CurrentState != operability.ProductionUsabilityValDStateActive
	if model.FailedProofStateObserved {
		model.Blockers = append(model.Blockers, "supportability cannot override failed prior proof state")
	}
	if model.ReadinessState != operability.ProductionUsabilityValCReadinessStateActive {
		model.Blockers = append(model.Blockers, "readiness review is not active")
	}
	if model.DiagnosticsState != operability.ProductionUsabilityValCDiagnosticsStateActive {
		model.Blockers = append(model.Blockers, "diagnostics review is not active")
	}
	if model.RecoveryState != operability.ProductionUsabilityValDRecoveryReviewStateActive {
		model.Blockers = append(model.Blockers, "recovery review is not active")
	}
	if model.UpgradeAdvisoryState != operability.ProductionUsabilityValDUpgradeRollbackReviewStateActive {
		model.Blockers = append(model.Blockers, "upgrade advisory review is not active")
	}
	if len(model.Blockers) > 0 {
		model.SupportabilityState = operability.ProductionUsabilityFinalGateBlocked
	} else {
		model.SupportabilityState = operability.ProductionUsabilityFinalGatePass
	}
	model.Limitations = collectIntegratedLimitations(val0, valA, valB, valC, valD)
	return model
}

func buildProductionUsabilityValERegressionClosureModel() operability.IntegratedUsabilityRegressionClosure {
	return operability.ProductionUsabilityValERegressionClosure()
}

func buildProductionUsabilityValEPassRuleModel(val0 productionUsabilityVal0ProofsResponse, valA productionUsabilityValAProofsResponse, valB productionUsabilityValBProofsResponse, valC productionUsabilityValCProofsResponse, valD productionUsabilityValDProofsResponse, prereqValEState string, dependencyClosure operability.IntegratedDependencyClosure, coherenceReview operability.CrossValCoherenceReview, canonicalBoundary operability.IntegratedCanonicalTruthBoundaryReview, redactionExportReview operability.IntegratedPermissionRedactionExportReview, supportabilityReview operability.IntegratedSupportabilityRecoveryReview, regressionClosure operability.IntegratedUsabilityRegressionClosure) operability.Point4IntegratedPassRule {
	model := operability.ProductionUsabilityValEPoint4PassRule()
	model.RequiredVals = productionUsabilityValERequiredVals()
	model.PassLimitations = collectIntegratedLimitations(val0, valA, valB, valC, valD)
	model.ValEState = prereqValEState
	model.ActiveVals = activeValNames(val0.Val0State, valA.ValAState, valB.ValBState, valC.ValCState, valD.ValDState, prereqValEState)
	model.MissingVals, model.PartialVals, model.UnsupportedVals = classifyRequiredVals(val0.Val0State, valA.ValAState, valB.ValBState, valC.ValCState, valD.ValDState, prereqValEState)
	if len(dependencyClosure.MissingVals) > 0 {
		model.PassBlockers = append(model.PassBlockers, "dependency closure missing required vals")
	}
	if len(dependencyClosure.InactiveVals) > 0 {
		model.PassBlockers = append(model.PassBlockers, "dependency closure has inactive vals")
	}
	if len(dependencyClosure.InconsistentVals) > 0 {
		model.PassBlockers = append(model.PassBlockers, "dependency closure has inconsistent vals")
	}
	if len(coherenceReview.MissingLinks) > 0 {
		model.PassBlockers = append(model.PassBlockers, "cross-val coherence review has missing critical links")
	}
	if len(coherenceReview.InconsistentLinks) > 0 {
		model.PassBlockers = append(model.PassBlockers, "cross-val coherence review has inconsistent links")
	}
	if len(canonicalBoundary.ViolationSurfaces) > 0 {
		model.PassBlockers = append(model.PassBlockers, "canonical-truth boundary review has violations")
	}
	if len(redactionExportReview.MissingScopes) > 0 || len(redactionExportReview.UnsafeScopes) > 0 {
		model.PassBlockers = append(model.PassBlockers, "permission/redaction/export review is incomplete or unsafe")
	}
	if len(supportabilityReview.Blockers) > 0 {
		model.PassBlockers = append(model.PassBlockers, "supportability/recovery review has blockers")
	}
	if len(regressionClosure.CriticalMissingCategories) > 0 {
		model.PassBlockers = append(model.PassBlockers, "regression closure has critical missing categories")
	}
	if prereqValEState != operability.ProductionUsabilityValEStateActive {
		model.PassBlockers = append(model.PassBlockers, "val_e prerequisites are not fully active")
	}
	model.PassWarnings = append(model.PassWarnings,
		"Integrated closure remains projection-only and does not replace canonical truth.",
	)
	model.Point4State = operability.ProductionUsabilityPoint4StateNotComplete
	model.PassCriteriaMet = false
	if len(model.PassBlockers) == 0 &&
		len(model.MissingVals) == 0 &&
		len(model.PartialVals) == 0 &&
		len(model.UnsupportedVals) == 0 &&
		prereqValEState == operability.ProductionUsabilityValEStateActive {
		model.Point4State = operability.ProductionUsabilityPoint4StatePass
		model.PassCriteriaMet = true
	}
	return model
}

func buildProductionUsabilityValEModelsCurrentState(models productionUsabilityValEModels) string {
	return operability.EvaluateProductionUsabilityValEProofsState(
		models.val0.Val0State,
		models.valA.ValAState,
		models.valB.ValBState,
		models.valC.ValCState,
		models.valD.ValDState,
		models.dependencyClosureState,
		models.coherenceReviewState,
		models.passRuleState,
		models.canonicalBoundaryState,
		models.redactionExportReviewState,
		models.supportabilityReviewState,
		models.regressionClosureState,
		models.point4State,
		models.surfaceRefs,
		models.evidenceRefs,
		models.limitations,
	)
}

func (s server) buildProductionUsabilityValEModels(ctx context.Context, filter phase4EnterpriseFilter) (productionUsabilityValEModels, error) {
	val0, err := s.buildProductionUsabilityVal0Proofs(ctx, filter)
	if err != nil {
		return productionUsabilityValEModels{}, err
	}
	valA, err := s.buildProductionUsabilityValAProofs(ctx, filter)
	if err != nil {
		return productionUsabilityValEModels{}, err
	}
	valB, err := s.buildProductionUsabilityValBProofs(ctx, filter)
	if err != nil {
		return productionUsabilityValEModels{}, err
	}
	valC, err := s.buildProductionUsabilityValCProofs(ctx, filter)
	if err != nil {
		return productionUsabilityValEModels{}, err
	}
	valD, err := s.buildProductionUsabilityValDProofs(ctx, filter)
	if err != nil {
		return productionUsabilityValEModels{}, err
	}

	dependencyClosure := buildProductionUsabilityValEDependencyClosureModel(val0, valA, valB, valC, valD)
	dependencyClosureState := operability.EvaluateProductionUsabilityValEDependencyClosureState(dependencyClosure)
	coherenceReview := buildProductionUsabilityValECoherenceReviewModel(val0, valA, valB, valC, valD)
	coherenceReviewState := operability.EvaluateProductionUsabilityValECoherenceReviewState(coherenceReview)
	canonicalBoundary := buildProductionUsabilityValECanonicalBoundaryReviewModel(valA, valB, valC, valD)
	canonicalBoundaryState := operability.EvaluateProductionUsabilityValECanonicalBoundaryReviewState(canonicalBoundary)
	redactionExportReview := buildProductionUsabilityValERedactionExportReviewModel(valA, valC, valD)
	redactionExportReviewState := operability.EvaluateProductionUsabilityValERedactionExportReviewState(redactionExportReview)
	supportabilityReview := buildProductionUsabilityValESupportabilityRecoveryReviewModel(val0, valA, valB, valC, valD)
	supportabilityReviewState := operability.EvaluateProductionUsabilityValESupportabilityRecoveryReviewState(supportabilityReview)
	regressionClosure := buildProductionUsabilityValERegressionClosureModel()
	regressionClosureState := operability.EvaluateProductionUsabilityValERegressionClosureState(regressionClosure)

	prereqValEState := operability.EvaluateProductionUsabilityValEPrerequisiteState(
		val0.Val0State,
		valA.ValAState,
		valB.ValBState,
		valC.ValCState,
		valD.ValDState,
		dependencyClosureState,
		coherenceReviewState,
		canonicalBoundaryState,
		redactionExportReviewState,
		supportabilityReviewState,
		regressionClosureState,
	)

	passRule := buildProductionUsabilityValEPassRuleModel(
		val0,
		valA,
		valB,
		valC,
		valD,
		prereqValEState,
		dependencyClosure,
		coherenceReview,
		canonicalBoundary,
		redactionExportReview,
		supportabilityReview,
		regressionClosure,
	)
	passRuleState := operability.EvaluateProductionUsabilityValEPassRuleState(passRule)
	valEState := operability.EvaluateProductionUsabilityValEState(
		val0.Val0State,
		valA.ValAState,
		valB.ValBState,
		valC.ValCState,
		valD.ValDState,
		dependencyClosureState,
		coherenceReviewState,
		passRuleState,
		canonicalBoundaryState,
		redactionExportReviewState,
		supportabilityReviewState,
		regressionClosureState,
	)

	passRule.ValEState = valEState
	passRule.ActiveVals = activeValNames(val0.Val0State, valA.ValAState, valB.ValBState, valC.ValCState, valD.ValDState, valEState)
	passRule.MissingVals, passRule.PartialVals, passRule.UnsupportedVals = classifyRequiredVals(val0.Val0State, valA.ValAState, valB.ValBState, valC.ValCState, valD.ValDState, valEState)
	passRule.Point4State = operability.ProductionUsabilityPoint4StateNotComplete
	passRule.PassCriteriaMet = false
	if len(passRule.PassBlockers) == 0 &&
		len(passRule.MissingVals) == 0 &&
		len(passRule.PartialVals) == 0 &&
		len(passRule.UnsupportedVals) == 0 &&
		valEState == operability.ProductionUsabilityValEStateActive {
		passRule.Point4State = operability.ProductionUsabilityPoint4StatePass
		passRule.PassCriteriaMet = true
	}
	passRuleState = operability.EvaluateProductionUsabilityValEPassRuleState(passRule)
	valEState = operability.EvaluateProductionUsabilityValEState(
		val0.Val0State,
		valA.ValAState,
		valB.ValBState,
		valC.ValCState,
		valD.ValDState,
		dependencyClosureState,
		coherenceReviewState,
		passRuleState,
		canonicalBoundaryState,
		redactionExportReviewState,
		supportabilityReviewState,
		regressionClosureState,
	)

	dependencyClosure.ValEState = valEState
	passRule.ValEState = valEState
	point4State := operability.ProductionUsabilityPoint4StateNotComplete
	if valEState == operability.ProductionUsabilityValEStateActive && passRuleState == operability.ProductionUsabilityValEPassRuleStateActive {
		point4State = operability.ProductionUsabilityPoint4StatePass
	}
	passRule.Point4State = point4State
	passRule.PassCriteriaMet = point4State == operability.ProductionUsabilityPoint4StatePass
	passRuleState = operability.EvaluateProductionUsabilityValEPassRuleState(passRule)

	limitations := collectIntegratedLimitations(val0, valA, valB, valC, valD)
	surfaceRefs := productionUsabilityValEAllSurfaceRefs()
	evidenceRefs := productionUsabilityValEEvidenceRefs()

	return productionUsabilityValEModels{
		val0:                       val0,
		valA:                       valA,
		valB:                       valB,
		valC:                       valC,
		valD:                       valD,
		dependencyClosure:          dependencyClosure,
		coherenceReview:            coherenceReview,
		passRule:                   passRule,
		canonicalBoundary:          canonicalBoundary,
		redactionExportReview:      redactionExportReview,
		supportabilityReview:       supportabilityReview,
		regressionClosure:          regressionClosure,
		dependencyClosureState:     dependencyClosureState,
		coherenceReviewState:       coherenceReviewState,
		passRuleState:              passRuleState,
		canonicalBoundaryState:     canonicalBoundaryState,
		redactionExportReviewState: redactionExportReviewState,
		supportabilityReviewState:  supportabilityReviewState,
		regressionClosureState:     regressionClosureState,
		valEState:                  valEState,
		point4State:                point4State,
		surfaceRefs:                surfaceRefs,
		evidenceRefs:               evidenceRefs,
		limitations:                limitations,
	}, nil
}

func (s server) productionUsabilityValEModelsFromRequest(w http.ResponseWriter, r *http.Request) (productionUsabilityValEModels, bool) {
	req, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r)
	if !ok {
		return productionUsabilityValEModels{}, false
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return productionUsabilityValEModels{}, false
	}
	filter, err := parsePhase4EnterpriseFilter(req)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return productionUsabilityValEModels{}, false
	}
	ctx, cancel := context.WithTimeout(req.Context(), s.requestTimeout)
	defer cancel()
	models, err := s.buildProductionUsabilityValEModels(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return productionUsabilityValEModels{}, false
	}
	return models, true
}

func (s server) productionUsabilityValEDependencyClosureHandler(w http.ResponseWriter, r *http.Request) {
	models, ok := s.productionUsabilityValEModelsFromRequest(w, r)
	if !ok {
		return
	}
	httpjson.Write(w, http.StatusOK, productionUsabilityValEDependencyClosureResponse{
		SchemaVersion: productionUsabilityValEDependencyClosureSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.dependencyClosureState,
		Model:         models.dependencyClosure,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vale/coherence-review",
			"/v1/production/usability-operability-recovery/vale/proofs",
		},
		Limitations: models.limitations,
	})
}

func (s server) productionUsabilityValECoherenceReviewHandler(w http.ResponseWriter, r *http.Request) {
	models, ok := s.productionUsabilityValEModelsFromRequest(w, r)
	if !ok {
		return
	}
	httpjson.Write(w, http.StatusOK, productionUsabilityValECoherenceReviewResponse{
		SchemaVersion: productionUsabilityValECoherenceReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.coherenceReviewState,
		Model:         models.coherenceReview,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vale/pass-rule",
			"/v1/production/usability-operability-recovery/vale/proofs",
		},
		Limitations: models.limitations,
	})
}

func (s server) productionUsabilityValEPassRuleHandler(w http.ResponseWriter, r *http.Request) {
	models, ok := s.productionUsabilityValEModelsFromRequest(w, r)
	if !ok {
		return
	}
	httpjson.Write(w, http.StatusOK, productionUsabilityValEPassRuleResponse{
		SchemaVersion: productionUsabilityValEPassRuleSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.passRuleState,
		Model:         models.passRule,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vale/proofs",
		},
		Limitations: models.limitations,
	})
}

func (s server) productionUsabilityValECanonicalTruthBoundaryHandler(w http.ResponseWriter, r *http.Request) {
	models, ok := s.productionUsabilityValEModelsFromRequest(w, r)
	if !ok {
		return
	}
	httpjson.Write(w, http.StatusOK, productionUsabilityValECanonicalBoundaryReviewResponse{
		SchemaVersion: productionUsabilityValECanonicalBoundaryReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.canonicalBoundaryState,
		Model:         models.canonicalBoundary,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vale/redaction-export-review",
			"/v1/production/usability-operability-recovery/vale/proofs",
		},
		Limitations: models.limitations,
	})
}

func (s server) productionUsabilityValERedactionExportReviewHandler(w http.ResponseWriter, r *http.Request) {
	models, ok := s.productionUsabilityValEModelsFromRequest(w, r)
	if !ok {
		return
	}
	httpjson.Write(w, http.StatusOK, productionUsabilityValERedactionExportReviewResponse{
		SchemaVersion: productionUsabilityValERedactionExportReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.redactionExportReviewState,
		Model:         models.redactionExportReview,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vale/supportability-recovery-review",
			"/v1/production/usability-operability-recovery/vale/proofs",
		},
		Limitations: models.limitations,
	})
}

func (s server) productionUsabilityValESupportabilityRecoveryReviewHandler(w http.ResponseWriter, r *http.Request) {
	models, ok := s.productionUsabilityValEModelsFromRequest(w, r)
	if !ok {
		return
	}
	httpjson.Write(w, http.StatusOK, productionUsabilityValESupportabilityRecoveryResponse{
		SchemaVersion: productionUsabilityValESupportabilityRecoverySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.supportabilityReviewState,
		Model:         models.supportabilityReview,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vale/regression-closure",
			"/v1/production/usability-operability-recovery/vale/proofs",
		},
		Limitations: models.limitations,
	})
}

func (s server) productionUsabilityValERegressionClosureHandler(w http.ResponseWriter, r *http.Request) {
	models, ok := s.productionUsabilityValEModelsFromRequest(w, r)
	if !ok {
		return
	}
	httpjson.Write(w, http.StatusOK, productionUsabilityValERegressionClosureResponse{
		SchemaVersion: productionUsabilityValERegressionClosureSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.regressionClosureState,
		Model:         models.regressionClosure,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vale/proofs",
		},
		Limitations: models.limitations,
	})
}

func (s server) productionUsabilityValEProofsHandler(w http.ResponseWriter, r *http.Request) {
	models, ok := s.productionUsabilityValEModelsFromRequest(w, r)
	if !ok {
		return
	}
	httpjson.Write(w, http.StatusOK, productionUsabilityValEProofsResponse{
		SchemaVersion:               productionUsabilityValEProofsSchema,
		GeneratedAt:                 publicSampleTime(),
		CurrentState:                buildProductionUsabilityValEModelsCurrentState(models),
		Val0DependencyState:         models.val0.CurrentState,
		Val0FoundationState:         models.val0.Val0State,
		ValADependencyState:         models.valA.CurrentState,
		ValACoreState:               models.valA.ValAState,
		ValBDependencyState:         models.valB.CurrentState,
		ValBResilienceState:         models.valB.ValBState,
		ValCDependencyState:         models.valC.CurrentState,
		ValCSupportabilityState:     models.valC.ValCState,
		ValDDependencyState:         models.valD.CurrentState,
		ValDFinalGateState:          models.valD.ValDState,
		ValEState:                   models.valEState,
		DependencyClosureState:      models.dependencyClosureState,
		CoherenceReviewState:        models.coherenceReviewState,
		Point4State:                 models.point4State,
		PassCriteriaMet:             models.passRule.PassCriteriaMet,
		PassBlockers:                models.passRule.PassBlockers,
		PassWarnings:                models.passRule.PassWarnings,
		PassLimitations:             models.passRule.PassLimitations,
		CanonicalTruthBoundaryState: models.canonicalBoundaryState,
		RedactionExportReviewState:  models.redactionExportReviewState,
		SupportabilityRecoveryState: models.supportabilityReviewState,
		RegressionClosureState:      models.regressionClosureState,
		EvidenceRefs:                models.evidenceRefs,
		SurfaceRefs:                 models.surfaceRefs,
		Limitations:                 models.limitations,
		ProjectionDisclaimer:        productionUsabilityValEProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val E is the only Točka 4 slice that can raise point_4_state to production_usability_point_4_pass.",
			"Integrated closure remains fail-closed on active Val 0, Val A, Val B, Val C, and Val D proof states.",
			"Integrated closure output is projection-only and does not replace the canonical evidence spine.",
		},
	})
}
