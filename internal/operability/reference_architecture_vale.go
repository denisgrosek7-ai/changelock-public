package operability

import "strings"

const (
	ReferenceArchitectureValEPrerequisiteStateActive     = "reference_architecture_vale_prerequisites_active"
	ReferenceArchitectureValEPrerequisiteStatePartial    = "reference_architecture_vale_prerequisites_partial"
	ReferenceArchitectureValEPrerequisiteStateIncomplete = "reference_architecture_vale_prerequisites_incomplete"
	ReferenceArchitectureValEPrerequisiteStateBlocked    = "reference_architecture_vale_prerequisites_blocked"
	ReferenceArchitectureValEPrerequisiteStateUnknown    = "reference_architecture_vale_prerequisites_unknown"

	ReferenceArchitectureValEInvariantStateActive     = "reference_architecture_vale_invariants_active"
	ReferenceArchitectureValEInvariantStatePartial    = "reference_architecture_vale_invariants_partial"
	ReferenceArchitectureValEInvariantStateIncomplete = "reference_architecture_vale_invariants_incomplete"
	ReferenceArchitectureValEInvariantStateBlocked    = "reference_architecture_vale_invariants_blocked"
	ReferenceArchitectureValEInvariantStateUnknown    = "reference_architecture_vale_invariants_unknown"

	ReferenceArchitectureValEProofSurfaceStateActive     = "reference_architecture_vale_proof_surface_active"
	ReferenceArchitectureValEProofSurfaceStatePartial    = "reference_architecture_vale_proof_surface_partial"
	ReferenceArchitectureValEProofSurfaceStateIncomplete = "reference_architecture_vale_proof_surface_incomplete"
	ReferenceArchitectureValEProofSurfaceStateBlocked    = "reference_architecture_vale_proof_surface_blocked"
	ReferenceArchitectureValEProofSurfaceStateUnknown    = "reference_architecture_vale_proof_surface_unknown"

	ReferenceArchitectureValEPassRuleStateActive     = "reference_architecture_vale_pass_rule_active"
	ReferenceArchitectureValEPassRuleStatePartial    = "reference_architecture_vale_pass_rule_partial"
	ReferenceArchitectureValEPassRuleStateIncomplete = "reference_architecture_vale_pass_rule_incomplete"
	ReferenceArchitectureValEPassRuleStateBlocked    = "reference_architecture_vale_pass_rule_blocked"
	ReferenceArchitectureValEPassRuleStateUnknown    = "reference_architecture_vale_pass_rule_unknown"

	ReferenceArchitectureValEStateActive     = "reference_architecture_vale_active"
	ReferenceArchitectureValEStatePartial    = "reference_architecture_vale_partial"
	ReferenceArchitectureValEStateIncomplete = "reference_architecture_vale_incomplete"
	ReferenceArchitectureValEStateBlocked    = "reference_architecture_vale_blocked"
	ReferenceArchitectureValEStateUnknown    = "reference_architecture_vale_unknown"

	ReferenceArchitectureValEClosureInvariantBlueprintDiscipline = "blueprint_discipline_invariant"
	ReferenceArchitectureValEClosureInvariantFamilyProfiles      = "family_profile_invariant"
	ReferenceArchitectureValEClosureInvariantBlueprintAsCode     = "blueprint_as_code_validation_invariant"
	ReferenceArchitectureValEClosureInvariantResilienceScaling   = "resilience_scaling_invariant"
	ReferenceArchitectureValEClosureInvariantOperationalGate     = "operational_visibility_final_gate_invariant"
	ReferenceArchitectureValEClosureInvariantAdvisoryProjection  = "advisory_projection_invariant"
)

type ReferenceArchitectureVal0ProofSnapshot struct {
	CurrentState               string   `json:"current_state"`
	Point5DependencyState      string   `json:"point_5_dependency_state"`
	Point5State                string   `json:"point_5_state"`
	Val0State                  string   `json:"val_0_state"`
	Point6State                string   `json:"point_6_state"`
	BlueprintDisciplineState   string   `json:"blueprint_discipline_state"`
	TaxonomyState              string   `json:"taxonomy_state"`
	EnvironmentFitState        string   `json:"environment_fit_state"`
	EvidenceDisciplineState    string   `json:"evidence_discipline_state"`
	CompatibilityBaselineState string   `json:"compatibility_baseline_state"`
	ConformanceState           string   `json:"conformance_state"`
	SupportedFamilies          []string `json:"supported_families,omitempty"`
	SupportedConformanceStates []string `json:"supported_conformance_states,omitempty"`
	SupportedCompatibility     []string `json:"supported_compatibility,omitempty"`
	SupportedLifecycle         []string `json:"supported_lifecycle,omitempty"`
	SurfaceRefs                []string `json:"surface_refs,omitempty"`
	EvidenceRefs               []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureValAProofSnapshot struct {
	CurrentState          string   `json:"current_state"`
	Point5DependencyState string   `json:"point_5_dependency_state"`
	Point5State           string   `json:"point_5_state"`
	Val0DependencyState   string   `json:"val_0_dependency_state"`
	Val0State             string   `json:"val_0_state"`
	ValAState             string   `json:"val_a_state"`
	Point6State           string   `json:"point_6_state"`
	RegistryState         string   `json:"registry_state"`
	SupportedFamilies     []string `json:"supported_families,omitempty"`
	SurfaceRefs           []string `json:"surface_refs,omitempty"`
	EvidenceRefs          []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer  string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureValBProofSnapshot struct {
	CurrentState          string   `json:"current_state"`
	Point5DependencyState string   `json:"point_5_dependency_state"`
	Point5State           string   `json:"point_5_state"`
	Val0DependencyState   string   `json:"val_0_dependency_state"`
	Val0State             string   `json:"val_0_state"`
	ValADependencyState   string   `json:"val_a_dependency_state"`
	ValAState             string   `json:"val_a_state"`
	ValBState             string   `json:"val_b_state"`
	Point6State           string   `json:"point_6_state"`
	PackRegistryState     string   `json:"pack_registry_state"`
	ArtifactManifestState string   `json:"artifact_manifest_state"`
	BundleState           string   `json:"bundle_state"`
	ReadinessState        string   `json:"readiness_state"`
	ValidationHookState   string   `json:"validation_hook_state"`
	ConformanceKitState   string   `json:"conformance_kit_state"`
	DeviationState        string   `json:"deviation_state"`
	SupportedFamilies     []string `json:"supported_families,omitempty"`
	SurfaceRefs           []string `json:"surface_refs,omitempty"`
	EvidenceRefs          []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer  string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureValCProofSnapshot struct {
	CurrentState             string   `json:"current_state"`
	Point5DependencyState    string   `json:"point_5_dependency_state"`
	Point5State              string   `json:"point_5_state"`
	Val0DependencyState      string   `json:"val_0_dependency_state"`
	Val0State                string   `json:"val_0_state"`
	ValADependencyState      string   `json:"val_a_dependency_state"`
	ValAState                string   `json:"val_a_state"`
	ValBDependencyState      string   `json:"val_b_dependency_state"`
	ValBState                string   `json:"val_b_state"`
	ValCState                string   `json:"val_c_state"`
	Point6State              string   `json:"point_6_state"`
	ScenarioPackState        string   `json:"scenario_pack_state"`
	FailureTaxonomyState     string   `json:"failure_taxonomy_state"`
	ScenarioDescriptorState  string   `json:"scenario_descriptor_state"`
	DegradedModeState        string   `json:"degraded_mode_state"`
	RecoveryExpectationState string   `json:"recovery_expectation_state"`
	ScalingScenarioState     string   `json:"scaling_scenario_state"`
	TrustPathState           string   `json:"trust_path_state"`
	AuditPathState           string   `json:"audit_path_state"`
	ControlPlaneSafetyState  string   `json:"control_plane_safety_state"`
	SupportedFamilies        []string `json:"supported_families,omitempty"`
	SurfaceRefs              []string `json:"surface_refs,omitempty"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureValDProofSnapshot struct {
	CurrentState               string   `json:"current_state"`
	Point5DependencyState      string   `json:"point_5_dependency_state"`
	Point5State                string   `json:"point_5_state"`
	Val0DependencyState        string   `json:"val_0_dependency_state"`
	Val0State                  string   `json:"val_0_state"`
	ValADependencyState        string   `json:"val_a_dependency_state"`
	ValAState                  string   `json:"val_a_state"`
	ValBDependencyState        string   `json:"val_b_dependency_state"`
	ValBState                  string   `json:"val_b_state"`
	ValCDependencyState        string   `json:"val_c_dependency_state"`
	ValCState                  string   `json:"val_c_state"`
	ValDState                  string   `json:"val_d_state"`
	Point6State                string   `json:"point_6_state"`
	OperationalVisibilityState string   `json:"operational_visibility_state"`
	AlignmentSummaryState      string   `json:"alignment_summary_state"`
	DeviationAlertState        string   `json:"deviation_alert_state"`
	SupportBoundaryState       string   `json:"support_boundary_state"`
	MigrationUpgradeState      string   `json:"migration_upgrade_state"`
	TopologyGateState          string   `json:"topology_gate_state"`
	SecurityBoundaryGateState  string   `json:"security_boundary_gate_state"`
	OperabilityGateState       string   `json:"operability_gate_state"`
	CompatibilityGateState     string   `json:"compatibility_gate_state"`
	FinalGateState             string   `json:"final_gate_state"`
	SupportedFamilies          []string `json:"supported_families,omitempty"`
	SurfaceRefs                []string `json:"surface_refs,omitempty"`
	EvidenceRefs               []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureValESourceValStates struct {
	Val0State string `json:"val_0_state"`
	ValAState string `json:"val_a_state"`
	ValBState string `json:"val_b_state"`
	ValCState string `json:"val_c_state"`
	ValDState string `json:"val_d_state"`
}

type ReferenceArchitectureValESourceCurrentStates struct {
	Val0CurrentState string `json:"val_0_current_state"`
	ValACurrentState string `json:"val_a_current_state"`
	ValBCurrentState string `json:"val_b_current_state"`
	ValCCurrentState string `json:"val_c_current_state"`
	ValDCurrentState string `json:"val_d_current_state"`
}

type ReferenceArchitectureValEDependencyStates struct {
	Point5State         string `json:"point_5_state"`
	Point5Dependency    string `json:"point_5_dependency_state"`
	Val0DependencyState string `json:"val_0_dependency_state"`
	ValADependencyState string `json:"val_a_dependency_state"`
	ValBDependencyState string `json:"val_b_dependency_state"`
	ValCDependencyState string `json:"val_c_dependency_state"`
	ValDFinalGateState  string `json:"val_d_final_gate_state"`
	PreClosurePoint6    string `json:"pre_closure_point_6_state"`
}

type ReferenceArchitectureClosureInvariant struct {
	CurrentState         string   `json:"current_state"`
	InvariantID          string   `json:"invariant_id"`
	Title                string   `json:"title"`
	BlockingReasons      []string `json:"blocking_reasons,omitempty"`
	Caveats              []string `json:"caveats,omitempty"`
	EvidenceRefs         []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureIntegratedClosure struct {
	CurrentState                  string                                       `json:"current_state"`
	ClosureID                     string                                       `json:"closure_id"`
	Version                       string                                       `json:"version"`
	Point                         string                                       `json:"point"`
	ClosureVal                    string                                       `json:"closure_val"`
	Point6State                   string                                       `json:"point_6_state"`
	Point6PassAllowed             bool                                         `json:"point_6_pass_allowed"`
	Point6PassReason              string                                       `json:"point_6_pass_reason"`
	SourceValStates               ReferenceArchitectureValESourceValStates     `json:"source_val_states"`
	SourceCurrentStates           ReferenceArchitectureValESourceCurrentStates `json:"source_current_states"`
	DependencyStates              ReferenceArchitectureValEDependencyStates    `json:"dependency_states"`
	Val0                          ReferenceArchitectureVal0ProofSnapshot       `json:"val_0"`
	ValA                          ReferenceArchitectureValAProofSnapshot       `json:"val_a"`
	ValB                          ReferenceArchitectureValBProofSnapshot       `json:"val_b"`
	ValC                          ReferenceArchitectureValCProofSnapshot       `json:"val_c"`
	ValD                          ReferenceArchitectureValDProofSnapshot       `json:"val_d"`
	ClosurePrerequisiteState      string                                       `json:"closure_prerequisite_state"`
	ClosureInvariantState         string                                       `json:"closure_invariant_state"`
	ProofSurfaceState             string                                       `json:"proof_surface_state"`
	PassRuleState                 string                                       `json:"pass_rule_state"`
	ProofSurfaceRefs              []string                                     `json:"proof_surface_refs,omitempty"`
	EvidenceRefs                  []string                                     `json:"evidence_refs,omitempty"`
	ClosureInvariants             []ReferenceArchitectureClosureInvariant      `json:"closure_invariants,omitempty"`
	BlockingReasons               []string                                     `json:"blocking_reasons,omitempty"`
	Caveats                       []string                                     `json:"caveats,omitempty"`
	Limitations                   []string                                     `json:"limitations,omitempty"`
	ProjectionDisclaimer          string                                       `json:"projection_disclaimer"`
	CreatedAt                     string                                       `json:"created_at"`
	UpdatedAt                     string                                       `json:"updated_at"`
	EvidenceFresh                 bool                                         `json:"evidence_fresh"`
	StaleEvidenceDetected         bool                                         `json:"stale_evidence_detected"`
	RedactionKeepsFailuresVisible bool                                         `json:"redaction_keeps_failures_visible"`
}

func referenceArchitectureValEProjectionDisclaimer() string {
	return "projection_only not_canonical_truth integrated_reference_architecture_closure evidence_linked_closure"
}

func referenceArchitectureValEHasProjectionDisclaimer(value string) bool {
	return strings.Contains(strings.TrimSpace(value), "projection_only") &&
		strings.Contains(strings.TrimSpace(value), "not_canonical_truth")
}

func referenceArchitectureValEProofSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/proofs",
		"/v1/reference-architecture/valc/proofs",
		"/v1/reference-architecture/vald/proofs",
		"/v1/reference-architecture/vald/final-gate",
		"/v1/reference-architecture/vale/closure",
		"/v1/reference-architecture/vale/proofs",
	}
}

func referenceArchitectureValEVal0ProofSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/blueprint-discipline",
		"/v1/reference-architecture/val0/environment-fit",
		"/v1/reference-architecture/val0/conformance-evidence",
		"/v1/reference-architecture/val0/compatibility-baseline",
		"/v1/reference-architecture/val0/proofs",
	}
}

func referenceArchitectureValEValAProofSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/family-registry",
		"/v1/reference-architecture/vala/family-profiles",
		"/v1/reference-architecture/vala/proofs",
	}
}

func referenceArchitectureValEValBProofSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/pack-registry",
		"/v1/reference-architecture/valb/bundles",
		"/v1/reference-architecture/valb/artifact-manifests",
		"/v1/reference-architecture/valb/readiness-checks",
		"/v1/reference-architecture/valb/validation-hooks",
		"/v1/reference-architecture/valb/conformance-kit",
		"/v1/reference-architecture/valb/deviations",
		"/v1/reference-architecture/valb/proofs",
	}
}

func referenceArchitectureValEValCProofSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/proofs",
		"/v1/reference-architecture/valc/scenario-packs",
		"/v1/reference-architecture/valc/failure-taxonomy",
		"/v1/reference-architecture/valc/scenario-descriptors",
		"/v1/reference-architecture/valc/degraded-modes",
		"/v1/reference-architecture/valc/recovery-expectations",
		"/v1/reference-architecture/valc/scaling-scenarios",
		"/v1/reference-architecture/valc/trust-path",
		"/v1/reference-architecture/valc/audit-path",
		"/v1/reference-architecture/valc/control-plane-safety",
		"/v1/reference-architecture/valc/proofs",
	}
}

func referenceArchitectureValEValDProofSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/proofs",
		"/v1/reference-architecture/valc/proofs",
		"/v1/reference-architecture/vald/operational-visibility",
		"/v1/reference-architecture/vald/alignment-summary",
		"/v1/reference-architecture/vald/deviation-alerts",
		"/v1/reference-architecture/vald/support-boundaries",
		"/v1/reference-architecture/vald/migration-upgrade",
		"/v1/reference-architecture/vald/topology-gate",
		"/v1/reference-architecture/vald/security-boundary-gate",
		"/v1/reference-architecture/vald/operability-gate",
		"/v1/reference-architecture/vald/compatibility-gate",
		"/v1/reference-architecture/vald/final-gate",
		"/v1/reference-architecture/vald/proofs",
	}
}

func referenceArchitectureValEHasOverclaim(values ...string) bool {
	disallowed := []string{
		"certified architecture",
		"guaranteed secure architecture",
		"absolute security",
		"engineering guarantee",
		"universal gold image",
		"regulator-approved architecture",
		"production approved",
		"deployment approved",
		"automatic approval",
		"legal certification",
		"formal certification",
	}
	for _, value := range values {
		normalized := strings.ToLower(strings.TrimSpace(value))
		for _, item := range disallowed {
			if strings.Contains(normalized, item) {
				return true
			}
		}
	}
	return false
}

func referenceArchitectureValECollectText(values []string) []string {
	filtered := make([]string, 0, len(values))
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func referenceArchitectureValEInvariant(id, title, state string, blockingReasons, caveats, evidenceRefs []string) ReferenceArchitectureClosureInvariant {
	return ReferenceArchitectureClosureInvariant{
		CurrentState:         state,
		InvariantID:          id,
		Title:                title,
		BlockingReasons:      blockingReasons,
		Caveats:              caveats,
		EvidenceRefs:         evidenceRefs,
		ProjectionDisclaimer: referenceArchitectureValEProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValEIntegratedClosureModel() ReferenceArchitectureIntegratedClosure {
	model := ReferenceArchitectureIntegratedClosure{
		ClosureID:        "reference-architecture-point-6-closure",
		Version:          "2026.04",
		Point:            "point_6",
		ClosureVal:       "val_e",
		Point6PassReason: "point_6_pass through Val E only after evidence-linked fail-closed closure across Val 0 through Val D.",
		ProofSurfaceRefs: referenceArchitectureValEProofSurfaceRefs(),
		EvidenceRefs: []string{
			"reference_architecture_val0_proofs",
			"reference_architecture_vala_proofs",
			"reference_architecture_valb_proofs",
			"reference_architecture_valc_proofs",
			"reference_architecture_vald_proofs",
			"reference_architecture_vald_final_gate",
		},
		Caveats: []string{
			"Integrated reference architecture closure remains advisory and evidence-linked.",
		},
		Limitations: []string{
			"Val E closes Točka 6 only and does not start Točka 7 or create deployment approval authority.",
			"Closure remains fail-closed and blocks on stale, partial, degraded, unknown, unsupported, drifted, or blocked source states.",
		},
		ProjectionDisclaimer:          referenceArchitectureValEProjectionDisclaimer(),
		CreatedAt:                     "2026-04-27T00:00:00Z",
		UpdatedAt:                     "2026-04-27T00:00:00Z",
		EvidenceFresh:                 true,
		RedactionKeepsFailuresVisible: true,
	}
	model.SourceValStates = ReferenceArchitectureValESourceValStates{
		Val0State: ReferenceArchitectureVal0StateActive,
		ValAState: ReferenceArchitectureValAStateActive,
		ValBState: ReferenceArchitectureValBStateActive,
		ValCState: ReferenceArchitectureValCStateActive,
		ValDState: ReferenceArchitectureValDStateActive,
	}
	model.SourceCurrentStates = ReferenceArchitectureValESourceCurrentStates{
		Val0CurrentState: ReferenceArchitectureVal0StateActive,
		ValACurrentState: ReferenceArchitectureValAStateActive,
		ValBCurrentState: ReferenceArchitectureValBStateActive,
		ValCCurrentState: ReferenceArchitectureValCStateActive,
		ValDCurrentState: ReferenceArchitectureValDStateActive,
	}
	model.DependencyStates = ReferenceArchitectureValEDependencyStates{
		Point5State:         IntelligenceCalibrationPoint5StatePass,
		Point5Dependency:    IntelligenceCalibrationValEStateActive,
		Val0DependencyState: ReferenceArchitectureVal0StateActive,
		ValADependencyState: ReferenceArchitectureValAStateActive,
		ValBDependencyState: ReferenceArchitectureValBStateActive,
		ValCDependencyState: ReferenceArchitectureValCStateActive,
		ValDFinalGateState:  ReferenceArchitectureValDFinalGateStateActive,
		PreClosurePoint6:    ReferenceArchitecturePoint6StateNotComplete,
	}
	model.Val0 = ReferenceArchitectureVal0ProofSnapshot{
		CurrentState:               ReferenceArchitectureVal0StateActive,
		Point5DependencyState:      IntelligenceCalibrationValEStateActive,
		Point5State:                IntelligenceCalibrationPoint5StatePass,
		Val0State:                  ReferenceArchitectureVal0StateActive,
		Point6State:                ReferenceArchitecturePoint6StateNotComplete,
		BlueprintDisciplineState:   ReferenceArchitectureVal0BlueprintDisciplineStateActive,
		TaxonomyState:              ReferenceArchitectureVal0TaxonomyStateActive,
		EnvironmentFitState:        ReferenceArchitectureVal0EnvironmentFitStateActive,
		EvidenceDisciplineState:    ReferenceArchitectureVal0EvidenceDisciplineStateActive,
		CompatibilityBaselineState: ReferenceArchitectureVal0CompatibilityBaselineStateActive,
		ConformanceState:           ReferenceArchitectureConformanceMatched,
		SupportedFamilies:          referenceArchitectureVal0Families(),
		SupportedConformanceStates: referenceArchitectureVal0ConformanceStates(),
		SupportedCompatibility:     referenceArchitectureVal0CompatibilityStates(),
		SupportedLifecycle:         referenceArchitectureVal0LifecycleStates(),
		SurfaceRefs:                referenceArchitectureValEVal0ProofSurfaceRefs(),
		EvidenceRefs:               []string{"val0-blueprint-discipline", "val0-environment-fit", "val0-compatibility-baseline"},
		ProjectionDisclaimer:       "projection_only not_canonical_truth bounded_blueprint_discipline",
	}
	model.ValA = ReferenceArchitectureValAProofSnapshot{
		CurrentState:          ReferenceArchitectureValAStateActive,
		Point5DependencyState: IntelligenceCalibrationValEStateActive,
		Point5State:           IntelligenceCalibrationPoint5StatePass,
		Val0DependencyState:   ReferenceArchitectureVal0StateActive,
		Val0State:             ReferenceArchitectureVal0StateActive,
		ValAState:             ReferenceArchitectureValAStateActive,
		Point6State:           ReferenceArchitecturePoint6StateNotComplete,
		RegistryState:         ReferenceArchitectureValAFamilyRegistryStateActive,
		SupportedFamilies:     referenceArchitectureVal0Families(),
		SurfaceRefs:           referenceArchitectureValEValAProofSurfaceRefs(),
		EvidenceRefs:          []string{"vala-family-registry", "vala-family-profiles"},
		ProjectionDisclaimer:  "projection_only not_canonical_truth validated_reference_blueprint_profile advisory_only",
	}
	model.ValB = ReferenceArchitectureValBProofSnapshot{
		CurrentState:          ReferenceArchitectureValBStateActive,
		Point5DependencyState: IntelligenceCalibrationValEStateActive,
		Point5State:           IntelligenceCalibrationPoint5StatePass,
		Val0DependencyState:   ReferenceArchitectureVal0StateActive,
		Val0State:             ReferenceArchitectureVal0StateActive,
		ValADependencyState:   ReferenceArchitectureValAStateActive,
		ValAState:             ReferenceArchitectureValAStateActive,
		ValBState:             ReferenceArchitectureValBStateActive,
		Point6State:           ReferenceArchitecturePoint6StateNotComplete,
		PackRegistryState:     ReferenceArchitectureValBPackStateActive,
		ArtifactManifestState: ReferenceArchitectureValBManifestStateActive,
		BundleState:           ReferenceArchitectureValBBundleStateActive,
		ReadinessState:        ReferenceArchitectureValBReadinessStateActive,
		ValidationHookState:   ReferenceArchitectureValBHookStateActive,
		ConformanceKitState:   ReferenceArchitectureValBConformanceKitStateActive,
		DeviationState:        ReferenceArchitectureValBDeviationStateActive,
		SupportedFamilies:     referenceArchitectureVal0Families(),
		SurfaceRefs:           referenceArchitectureValEValBProofSurfaceRefs(),
		EvidenceRefs:          []string{"valb-pack-registry", "valb-conformance-kit", "valb-validation-hooks"},
		ProjectionDisclaimer:  "projection_only not_canonical_truth bounded_blueprint_as_code_pack advisory_projection",
	}
	model.ValC = ReferenceArchitectureValCProofSnapshot{
		CurrentState:             ReferenceArchitectureValCStateActive,
		Point5DependencyState:    IntelligenceCalibrationValEStateActive,
		Point5State:              IntelligenceCalibrationPoint5StatePass,
		Val0DependencyState:      ReferenceArchitectureVal0StateActive,
		Val0State:                ReferenceArchitectureVal0StateActive,
		ValADependencyState:      ReferenceArchitectureValAStateActive,
		ValAState:                ReferenceArchitectureValAStateActive,
		ValBDependencyState:      ReferenceArchitectureValBStateActive,
		ValBState:                ReferenceArchitectureValBStateActive,
		ValCState:                ReferenceArchitectureValCStateActive,
		Point6State:              ReferenceArchitecturePoint6StateNotComplete,
		ScenarioPackState:        ReferenceArchitectureValCScenarioPackStateActive,
		FailureTaxonomyState:     ReferenceArchitectureValCFailureTaxonomyStateActive,
		ScenarioDescriptorState:  ReferenceArchitectureValCScenarioDescriptorStateActive,
		DegradedModeState:        ReferenceArchitectureValCDegradedModeStateActive,
		RecoveryExpectationState: ReferenceArchitectureValCRecoveryExpectationStateActive,
		ScalingScenarioState:     ReferenceArchitectureValCScalingScenarioStateActive,
		TrustPathState:           ReferenceArchitectureValCTrustPathStateActive,
		AuditPathState:           ReferenceArchitectureValCAuditPathStateActive,
		ControlPlaneSafetyState:  ReferenceArchitectureValCControlPlaneStateActive,
		SupportedFamilies:        referenceArchitectureVal0Families(),
		SurfaceRefs:              referenceArchitectureValEValCProofSurfaceRefs(),
		EvidenceRefs:             []string{"valc-scenario-pack", "valc-trust-path", "valc-control-plane-safety"},
		ProjectionDisclaimer:     "projection_only not_canonical_truth bounded_resilience_scaling_hardening",
	}
	model.ValD = ReferenceArchitectureValDProofSnapshot{
		CurrentState:               ReferenceArchitectureValDStateActive,
		Point5DependencyState:      IntelligenceCalibrationValEStateActive,
		Point5State:                IntelligenceCalibrationPoint5StatePass,
		Val0DependencyState:        ReferenceArchitectureVal0StateActive,
		Val0State:                  ReferenceArchitectureVal0StateActive,
		ValADependencyState:        ReferenceArchitectureValAStateActive,
		ValAState:                  ReferenceArchitectureValAStateActive,
		ValBDependencyState:        ReferenceArchitectureValBStateActive,
		ValBState:                  ReferenceArchitectureValBStateActive,
		ValCDependencyState:        ReferenceArchitectureValCStateActive,
		ValCState:                  ReferenceArchitectureValCStateActive,
		ValDState:                  ReferenceArchitectureValDStateActive,
		Point6State:                ReferenceArchitecturePoint6StateNotComplete,
		OperationalVisibilityState: ReferenceArchitectureValDVisibilityStateActive,
		AlignmentSummaryState:      ReferenceArchitectureValDAlignmentStateActive,
		DeviationAlertState:        ReferenceArchitectureValDAlertStateActive,
		SupportBoundaryState:       ReferenceArchitectureValDSupportBoundaryStateActive,
		MigrationUpgradeState:      ReferenceArchitectureValDMigrationStateActive,
		TopologyGateState:          ReferenceArchitectureValDTopologyGateStateActive,
		SecurityBoundaryGateState:  ReferenceArchitectureValDSecurityGateStateActive,
		OperabilityGateState:       ReferenceArchitectureValDOperabilityGateStateActive,
		CompatibilityGateState:     ReferenceArchitectureValDCompatibilityGateStateActive,
		FinalGateState:             ReferenceArchitectureValDFinalGateStateActive,
		SupportedFamilies:          referenceArchitectureVal0Families(),
		SurfaceRefs:                referenceArchitectureValEValDProofSurfaceRefs(),
		EvidenceRefs:               []string{"vald-operational-visibility", "vald-final-gate", "vald-compatibility-gate"},
		ProjectionDisclaimer:       "projection_only not_canonical_truth bounded_operational_visibility_final_reference_gate",
	}
	return ComputeReferenceArchitectureValEClosure(model)
}

func EvaluateReferenceArchitectureValEPrerequisiteState(model ReferenceArchitectureIntegratedClosure) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ClosureID,
		model.Point,
		model.ClosureVal,
		model.Point6PassReason,
		model.ProjectionDisclaimer,
	) {
		return ReferenceArchitectureValEPrerequisiteStateIncomplete
	}
	if strings.TrimSpace(model.Point) != "point_6" ||
		strings.TrimSpace(model.ClosureVal) != "val_e" ||
		!referenceArchitectureValEHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return ReferenceArchitectureValEPrerequisiteStatePartial
	}
	if strings.TrimSpace(model.DependencyStates.Point5State) != IntelligenceCalibrationPoint5StatePass ||
		!referenceArchitectureValDPoint5DependencyHealthy(model.DependencyStates.Point5Dependency) ||
		strings.TrimSpace(model.DependencyStates.Val0DependencyState) != ReferenceArchitectureVal0StateActive ||
		strings.TrimSpace(model.DependencyStates.ValADependencyState) != ReferenceArchitectureValAStateActive ||
		strings.TrimSpace(model.DependencyStates.ValBDependencyState) != ReferenceArchitectureValBStateActive ||
		strings.TrimSpace(model.DependencyStates.ValCDependencyState) != ReferenceArchitectureValCStateActive ||
		strings.TrimSpace(model.SourceCurrentStates.Val0CurrentState) != ReferenceArchitectureVal0StateActive ||
		strings.TrimSpace(model.SourceValStates.Val0State) != ReferenceArchitectureVal0StateActive ||
		strings.TrimSpace(model.SourceCurrentStates.ValACurrentState) != ReferenceArchitectureValAStateActive ||
		strings.TrimSpace(model.SourceValStates.ValAState) != ReferenceArchitectureValAStateActive ||
		strings.TrimSpace(model.SourceCurrentStates.ValBCurrentState) != ReferenceArchitectureValBStateActive ||
		strings.TrimSpace(model.SourceValStates.ValBState) != ReferenceArchitectureValBStateActive ||
		strings.TrimSpace(model.SourceCurrentStates.ValCCurrentState) != ReferenceArchitectureValCStateActive ||
		strings.TrimSpace(model.SourceValStates.ValCState) != ReferenceArchitectureValCStateActive ||
		strings.TrimSpace(model.SourceCurrentStates.ValDCurrentState) != ReferenceArchitectureValDStateActive ||
		strings.TrimSpace(model.SourceValStates.ValDState) != ReferenceArchitectureValDStateActive ||
		strings.TrimSpace(model.DependencyStates.ValDFinalGateState) != ReferenceArchitectureValDFinalGateStateActive ||
		strings.TrimSpace(model.DependencyStates.PreClosurePoint6) != ReferenceArchitecturePoint6StateNotComplete {
		return ReferenceArchitectureValEPrerequisiteStateBlocked
	}
	return ReferenceArchitectureValEPrerequisiteStateActive
}

func evaluateReferenceArchitectureValEBlueprintDisciplineInvariant(model ReferenceArchitectureIntegratedClosure) ReferenceArchitectureClosureInvariant {
	blockingReasons := []string{}
	if model.Val0.CurrentState != ReferenceArchitectureVal0StateActive || model.Val0.Val0State != ReferenceArchitectureVal0StateActive {
		blockingReasons = append(blockingReasons, "Val 0 current and val states must remain active for integrated closure.")
	}
	if model.Val0.Point6State != ReferenceArchitecturePoint6StateNotComplete {
		blockingReasons = append(blockingReasons, "Val 0 must not claim point_6_pass before Val E integrated closure.")
	}
	if model.Val0.BlueprintDisciplineState != ReferenceArchitectureVal0BlueprintDisciplineStateActive ||
		model.Val0.TaxonomyState != ReferenceArchitectureVal0TaxonomyStateActive ||
		model.Val0.EnvironmentFitState != ReferenceArchitectureVal0EnvironmentFitStateActive ||
		model.Val0.EvidenceDisciplineState != ReferenceArchitectureVal0EvidenceDisciplineStateActive ||
		model.Val0.CompatibilityBaselineState != ReferenceArchitectureVal0CompatibilityBaselineStateActive ||
		model.Val0.ConformanceState != ReferenceArchitectureConformanceMatched {
		blockingReasons = append(blockingReasons, "Val 0 blueprint discipline, taxonomy, evidence, compatibility, and conformance states must all remain active and matched.")
	}
	if !containsExactTrimmedStringSet(model.Val0.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(model.Val0.SupportedConformanceStates, referenceArchitectureVal0ConformanceStates()...) ||
		!containsExactTrimmedStringSet(model.Val0.SupportedCompatibility, referenceArchitectureVal0CompatibilityStates()...) ||
		!containsExactTrimmedStringSet(model.Val0.SupportedLifecycle, referenceArchitectureVal0LifecycleStates()...) ||
		!containsExactTrimmedStringSet(model.Val0.SurfaceRefs, referenceArchitectureValEVal0ProofSurfaceRefs()...) ||
		!referenceArchitectureValEHasProjectionDisclaimer(model.Val0.ProjectionDisclaimer) {
		blockingReasons = append(blockingReasons, "Val 0 supported enums, proof surfaces, and advisory disclaimer must remain exact and fail closed.")
	}
	state := ReferenceArchitectureValEInvariantStateActive
	if len(blockingReasons) > 0 {
		state = ReferenceArchitectureValEInvariantStateBlocked
	}
	return referenceArchitectureValEInvariant(
		ReferenceArchitectureValEClosureInvariantBlueprintDiscipline,
		"Blueprint Discipline Invariant",
		state,
		blockingReasons,
		nil,
		model.Val0.EvidenceRefs,
	)
}

func evaluateReferenceArchitectureValEFamilyProfileInvariant(model ReferenceArchitectureIntegratedClosure) ReferenceArchitectureClosureInvariant {
	blockingReasons := []string{}
	if model.ValA.CurrentState != ReferenceArchitectureValAStateActive || model.ValA.ValAState != ReferenceArchitectureValAStateActive || model.ValA.RegistryState != ReferenceArchitectureValAFamilyRegistryStateActive {
		blockingReasons = append(blockingReasons, "Val A current state, val state, and family registry state must remain active.")
	}
	if model.ValA.Point6State != ReferenceArchitecturePoint6StateNotComplete {
		blockingReasons = append(blockingReasons, "Val A must not close point 6 before Val E.")
	}
	if !containsExactTrimmedStringSet(model.ValA.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(model.ValA.SurfaceRefs, referenceArchitectureValEValAProofSurfaceRefs()...) ||
		!referenceArchitectureValEHasProjectionDisclaimer(model.ValA.ProjectionDisclaimer) {
		blockingReasons = append(blockingReasons, "Val A family coverage, proof surfaces, and advisory boundary must remain exact.")
	}
	state := ReferenceArchitectureValEInvariantStateActive
	if len(blockingReasons) > 0 {
		state = ReferenceArchitectureValEInvariantStateBlocked
	}
	return referenceArchitectureValEInvariant(
		ReferenceArchitectureValEClosureInvariantFamilyProfiles,
		"Family Profile Invariant",
		state,
		blockingReasons,
		nil,
		model.ValA.EvidenceRefs,
	)
}

func evaluateReferenceArchitectureValEBlueprintAsCodeInvariant(model ReferenceArchitectureIntegratedClosure) ReferenceArchitectureClosureInvariant {
	blockingReasons := []string{}
	if model.ValB.CurrentState != ReferenceArchitectureValBStateActive || model.ValB.ValBState != ReferenceArchitectureValBStateActive {
		blockingReasons = append(blockingReasons, "Val B current state and val state must remain active.")
	}
	if model.ValB.Point6State != ReferenceArchitecturePoint6StateNotComplete {
		blockingReasons = append(blockingReasons, "Val B must not return point_6_pass before integrated closure.")
	}
	if model.ValB.PackRegistryState != ReferenceArchitectureValBPackStateActive ||
		model.ValB.ArtifactManifestState != ReferenceArchitectureValBManifestStateActive ||
		model.ValB.BundleState != ReferenceArchitectureValBBundleStateActive ||
		model.ValB.ReadinessState != ReferenceArchitectureValBReadinessStateActive ||
		model.ValB.ValidationHookState != ReferenceArchitectureValBHookStateActive ||
		model.ValB.ConformanceKitState != ReferenceArchitectureValBConformanceKitStateActive ||
		model.ValB.DeviationState != ReferenceArchitectureValBDeviationStateActive {
		blockingReasons = append(blockingReasons, "Val B pack registry, artifact manifests, bundles, readiness, hooks, conformance kits, and deviations must all remain active.")
	}
	if !containsExactTrimmedStringSet(model.ValB.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(model.ValB.SurfaceRefs, referenceArchitectureValEValBProofSurfaceRefs()...) ||
		!referenceArchitectureValEHasProjectionDisclaimer(model.ValB.ProjectionDisclaimer) {
		blockingReasons = append(blockingReasons, "Val B proof surfaces, family coverage, and advisory boundary must remain exact.")
	}
	state := ReferenceArchitectureValEInvariantStateActive
	if len(blockingReasons) > 0 {
		state = ReferenceArchitectureValEInvariantStateBlocked
	}
	return referenceArchitectureValEInvariant(
		ReferenceArchitectureValEClosureInvariantBlueprintAsCode,
		"Blueprint-as-Code Validation Invariant",
		state,
		blockingReasons,
		nil,
		model.ValB.EvidenceRefs,
	)
}

func evaluateReferenceArchitectureValEResilienceInvariant(model ReferenceArchitectureIntegratedClosure) ReferenceArchitectureClosureInvariant {
	blockingReasons := []string{}
	if model.ValC.CurrentState != ReferenceArchitectureValCStateActive || model.ValC.ValCState != ReferenceArchitectureValCStateActive {
		blockingReasons = append(blockingReasons, "Val C current state and val state must remain active.")
	}
	if model.ValC.Point6State != ReferenceArchitecturePoint6StateNotComplete {
		blockingReasons = append(blockingReasons, "Val C must not claim point_6_pass before integrated closure.")
	}
	if !referenceArchitectureValDPoint5DependencyHealthy(model.ValC.Point5DependencyState) || model.ValC.Point5DependencyState == model.ValC.Point5State {
		blockingReasons = append(blockingReasons, "Val C must report Point 5 dependency health separately from point_5_state and keep it healthy.")
	}
	if model.ValC.ScenarioPackState != ReferenceArchitectureValCScenarioPackStateActive ||
		model.ValC.FailureTaxonomyState != ReferenceArchitectureValCFailureTaxonomyStateActive ||
		model.ValC.ScenarioDescriptorState != ReferenceArchitectureValCScenarioDescriptorStateActive ||
		model.ValC.DegradedModeState != ReferenceArchitectureValCDegradedModeStateActive ||
		model.ValC.RecoveryExpectationState != ReferenceArchitectureValCRecoveryExpectationStateActive ||
		model.ValC.ScalingScenarioState != ReferenceArchitectureValCScalingScenarioStateActive ||
		model.ValC.TrustPathState != ReferenceArchitectureValCTrustPathStateActive ||
		model.ValC.AuditPathState != ReferenceArchitectureValCAuditPathStateActive ||
		model.ValC.ControlPlaneSafetyState != ReferenceArchitectureValCControlPlaneStateActive {
		blockingReasons = append(blockingReasons, "Val C scenario, failure-mode, recovery, scaling, trust-path, audit-path, and control-plane states must all remain active.")
	}
	if !containsExactTrimmedStringSet(model.ValC.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(model.ValC.SurfaceRefs, referenceArchitectureValEValCProofSurfaceRefs()...) ||
		!referenceArchitectureValEHasProjectionDisclaimer(model.ValC.ProjectionDisclaimer) {
		blockingReasons = append(blockingReasons, "Val C proof surfaces, family coverage, and advisory boundary must remain exact.")
	}
	state := ReferenceArchitectureValEInvariantStateActive
	if len(blockingReasons) > 0 {
		state = ReferenceArchitectureValEInvariantStateBlocked
	}
	return referenceArchitectureValEInvariant(
		ReferenceArchitectureValEClosureInvariantResilienceScaling,
		"Resilience and Scaling Invariant",
		state,
		blockingReasons,
		nil,
		model.ValC.EvidenceRefs,
	)
}

func evaluateReferenceArchitectureValEOperationalGateInvariant(model ReferenceArchitectureIntegratedClosure) ReferenceArchitectureClosureInvariant {
	blockingReasons := []string{}
	if model.ValD.CurrentState != ReferenceArchitectureValDStateActive || model.ValD.ValDState != ReferenceArchitectureValDStateActive {
		blockingReasons = append(blockingReasons, "Val D current state and val state must remain active.")
	}
	if model.ValD.Point6State != ReferenceArchitecturePoint6StateNotComplete {
		blockingReasons = append(blockingReasons, "Val D must not claim integrated closure.")
	}
	if model.ValD.OperationalVisibilityState != ReferenceArchitectureValDVisibilityStateActive ||
		model.ValD.AlignmentSummaryState != ReferenceArchitectureValDAlignmentStateActive ||
		model.ValD.DeviationAlertState != ReferenceArchitectureValDAlertStateActive ||
		model.ValD.SupportBoundaryState != ReferenceArchitectureValDSupportBoundaryStateActive ||
		model.ValD.MigrationUpgradeState != ReferenceArchitectureValDMigrationStateActive ||
		model.ValD.TopologyGateState != ReferenceArchitectureValDTopologyGateStateActive ||
		model.ValD.SecurityBoundaryGateState != ReferenceArchitectureValDSecurityGateStateActive ||
		model.ValD.OperabilityGateState != ReferenceArchitectureValDOperabilityGateStateActive ||
		model.ValD.CompatibilityGateState != ReferenceArchitectureValDCompatibilityGateStateActive ||
		model.ValD.FinalGateState != ReferenceArchitectureValDFinalGateStateActive {
		blockingReasons = append(blockingReasons, "Val D operational visibility, gate states, and final gate must all remain exact active constants.")
	}
	if !containsExactTrimmedStringSet(model.ValD.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(model.ValD.SurfaceRefs, referenceArchitectureValEValDProofSurfaceRefs()...) ||
		!referenceArchitectureValEHasProjectionDisclaimer(model.ValD.ProjectionDisclaimer) {
		blockingReasons = append(blockingReasons, "Val D family normalization, proof surface set, and advisory boundary must remain exact.")
	}
	state := ReferenceArchitectureValEInvariantStateActive
	if len(blockingReasons) > 0 {
		state = ReferenceArchitectureValEInvariantStateBlocked
	}
	return referenceArchitectureValEInvariant(
		ReferenceArchitectureValEClosureInvariantOperationalGate,
		"Operational Visibility and Final Gate Invariant",
		state,
		blockingReasons,
		nil,
		model.ValD.EvidenceRefs,
	)
}

func evaluateReferenceArchitectureValEAdvisoryInvariant(model ReferenceArchitectureIntegratedClosure) ReferenceArchitectureClosureInvariant {
	texts := append([]string{model.Point6PassReason, model.ProjectionDisclaimer}, model.Caveats...)
	texts = append(texts, model.Limitations...)
	blockingReasons := []string{}
	if !referenceArchitectureValEHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!referenceArchitectureValEHasProjectionDisclaimer(model.Val0.ProjectionDisclaimer) ||
		!referenceArchitectureValEHasProjectionDisclaimer(model.ValA.ProjectionDisclaimer) ||
		!referenceArchitectureValEHasProjectionDisclaimer(model.ValB.ProjectionDisclaimer) ||
		!referenceArchitectureValEHasProjectionDisclaimer(model.ValC.ProjectionDisclaimer) ||
		!referenceArchitectureValEHasProjectionDisclaimer(model.ValD.ProjectionDisclaimer) {
		blockingReasons = append(blockingReasons, "All Val 0 through Val E surfaces must remain advisory and projection-only.")
	}
	if !model.RedactionKeepsFailuresVisible {
		blockingReasons = append(blockingReasons, "Redaction or caveat handling must keep failures visible and fail closed.")
	}
	if referenceArchitectureValEHasOverclaim(texts...) {
		blockingReasons = append(blockingReasons, "Closure language must not claim certification, guaranteed security, or deployment approval.")
	}
	state := ReferenceArchitectureValEInvariantStateActive
	if len(blockingReasons) > 0 {
		state = ReferenceArchitectureValEInvariantStateBlocked
	}
	return referenceArchitectureValEInvariant(
		ReferenceArchitectureValEClosureInvariantAdvisoryProjection,
		"Advisory Projection Invariant",
		state,
		blockingReasons,
		nil,
		model.EvidenceRefs,
	)
}

func ReferenceArchitectureValEClosureInvariants(model ReferenceArchitectureIntegratedClosure) []ReferenceArchitectureClosureInvariant {
	return []ReferenceArchitectureClosureInvariant{
		evaluateReferenceArchitectureValEBlueprintDisciplineInvariant(model),
		evaluateReferenceArchitectureValEFamilyProfileInvariant(model),
		evaluateReferenceArchitectureValEBlueprintAsCodeInvariant(model),
		evaluateReferenceArchitectureValEResilienceInvariant(model),
		evaluateReferenceArchitectureValEOperationalGateInvariant(model),
		evaluateReferenceArchitectureValEAdvisoryInvariant(model),
	}
}

func EvaluateReferenceArchitectureValEInvariantState(model ReferenceArchitectureIntegratedClosure) string {
	invariants := ReferenceArchitectureValEClosureInvariants(model)
	if len(invariants) == 0 {
		return ReferenceArchitectureValEInvariantStateIncomplete
	}
	allActive := true
	for _, invariant := range invariants {
		if strings.TrimSpace(invariant.CurrentState) == "" {
			return ReferenceArchitectureValEInvariantStateIncomplete
		}
		if !referenceArchitectureValEHasProjectionDisclaimer(invariant.ProjectionDisclaimer) {
			return ReferenceArchitectureValEInvariantStatePartial
		}
		if invariant.CurrentState != ReferenceArchitectureValEInvariantStateActive {
			allActive = false
		}
	}
	if allActive {
		return ReferenceArchitectureValEInvariantStateActive
	}
	for _, invariant := range invariants {
		switch invariant.CurrentState {
		case ReferenceArchitectureValEInvariantStateBlocked:
			return ReferenceArchitectureValEInvariantStateBlocked
		case ReferenceArchitectureValEInvariantStateUnknown:
			return ReferenceArchitectureValEInvariantStateUnknown
		case ReferenceArchitectureValEInvariantStateIncomplete:
			return ReferenceArchitectureValEInvariantStateIncomplete
		}
	}
	return ReferenceArchitectureValEInvariantStatePartial
}

func EvaluateReferenceArchitectureValEProofSurfaceState(model ReferenceArchitectureIntegratedClosure) string {
	if strings.TrimSpace(model.ProjectionDisclaimer) == "" || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 {
		return ReferenceArchitectureValEProofSurfaceStateIncomplete
	}
	if !referenceArchitectureValEHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return ReferenceArchitectureValEProofSurfaceStatePartial
	}
	if !containsExactTrimmedStringSet(model.ProofSurfaceRefs, referenceArchitectureValEProofSurfaceRefs()...) {
		return ReferenceArchitectureValEProofSurfaceStatePartial
	}
	if model.StaleEvidenceDetected || !model.EvidenceFresh {
		return ReferenceArchitectureValEProofSurfaceStateBlocked
	}
	return ReferenceArchitectureValEProofSurfaceStateActive
}

func EvaluateReferenceArchitectureValEPassRuleState(model ReferenceArchitectureIntegratedClosure) string {
	if strings.TrimSpace(model.Point6PassReason) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return ReferenceArchitectureValEPassRuleStateIncomplete
	}
	if referenceArchitectureValEHasOverclaim(model.Point6PassReason, strings.Join(model.Caveats, " "), strings.Join(model.Limitations, " "), strings.Join(model.BlockingReasons, " ")) {
		return ReferenceArchitectureValEPassRuleStateBlocked
	}
	prereqState := EvaluateReferenceArchitectureValEPrerequisiteState(model)
	invariantState := EvaluateReferenceArchitectureValEInvariantState(model)
	proofSurfaceState := EvaluateReferenceArchitectureValEProofSurfaceState(model)
	if prereqState == ReferenceArchitectureValEPrerequisiteStateActive &&
		invariantState == ReferenceArchitectureValEInvariantStateActive &&
		proofSurfaceState == ReferenceArchitectureValEProofSurfaceStateActive &&
		model.RedactionKeepsFailuresVisible {
		return ReferenceArchitectureValEPassRuleStateActive
	}
	if prereqState == ReferenceArchitectureValEPrerequisiteStateBlocked ||
		invariantState == ReferenceArchitectureValEInvariantStateBlocked ||
		proofSurfaceState == ReferenceArchitectureValEProofSurfaceStateBlocked ||
		!model.RedactionKeepsFailuresVisible {
		return ReferenceArchitectureValEPassRuleStateBlocked
	}
	if prereqState == ReferenceArchitectureValEPrerequisiteStateIncomplete ||
		invariantState == ReferenceArchitectureValEInvariantStateIncomplete ||
		proofSurfaceState == ReferenceArchitectureValEProofSurfaceStateIncomplete {
		return ReferenceArchitectureValEPassRuleStateIncomplete
	}
	if prereqState == ReferenceArchitectureValEPrerequisiteStateUnknown ||
		invariantState == ReferenceArchitectureValEInvariantStateUnknown ||
		proofSurfaceState == ReferenceArchitectureValEProofSurfaceStateUnknown {
		return ReferenceArchitectureValEPassRuleStateUnknown
	}
	return ReferenceArchitectureValEPassRuleStatePartial
}

func EvaluateReferenceArchitecturePoint6FinalState(model ReferenceArchitectureIntegratedClosure) string {
	if EvaluateReferenceArchitectureValEPassRuleState(model) == ReferenceArchitectureValEPassRuleStateActive {
		return ReferenceArchitecturePoint6StatePass
	}
	return ReferenceArchitecturePoint6StateNotComplete
}

func EvaluateReferenceArchitectureValEState(model ReferenceArchitectureIntegratedClosure) string {
	passRuleState := EvaluateReferenceArchitectureValEPassRuleState(model)
	if passRuleState == ReferenceArchitectureValEPassRuleStateActive && EvaluateReferenceArchitecturePoint6FinalState(model) == ReferenceArchitecturePoint6StatePass {
		return ReferenceArchitectureValEStateActive
	}
	switch passRuleState {
	case ReferenceArchitectureValEPassRuleStateBlocked:
		return ReferenceArchitectureValEStateBlocked
	case ReferenceArchitectureValEPassRuleStateIncomplete:
		return ReferenceArchitectureValEStateIncomplete
	case ReferenceArchitectureValEPassRuleStateUnknown:
		return ReferenceArchitectureValEStateUnknown
	default:
		return ReferenceArchitectureValEStatePartial
	}
}

func referenceArchitectureValEBlockingReasons(model ReferenceArchitectureIntegratedClosure, invariants []ReferenceArchitectureClosureInvariant) []string {
	reasons := []string{}
	if prereqState := EvaluateReferenceArchitectureValEPrerequisiteState(model); prereqState != ReferenceArchitectureValEPrerequisiteStateActive {
		reasons = append(reasons, "Integrated closure prerequisites are not fully active and healthy.")
	}
	if proofSurfaceState := EvaluateReferenceArchitectureValEProofSurfaceState(model); proofSurfaceState != ReferenceArchitectureValEProofSurfaceStateActive {
		reasons = append(reasons, "Integrated closure proof surfaces or aggregated evidence refs are not exact and fresh.")
	}
	if passRuleState := EvaluateReferenceArchitectureValEPassRuleState(model); passRuleState != ReferenceArchitectureValEPassRuleStateActive {
		reasons = append(reasons, "Val E pass rule remains fail-closed until all prerequisites and invariants are active.")
	}
	for _, invariant := range invariants {
		if invariant.CurrentState != ReferenceArchitectureValEInvariantStateActive {
			reasons = append(reasons, invariant.BlockingReasons...)
		}
	}
	return referenceArchitectureValECollectText(reasons)
}

func ComputeReferenceArchitectureValEClosure(model ReferenceArchitectureIntegratedClosure) ReferenceArchitectureIntegratedClosure {
	invariants := ReferenceArchitectureValEClosureInvariants(model)
	model.ClosureInvariants = invariants
	model.ClosurePrerequisiteState = EvaluateReferenceArchitectureValEPrerequisiteState(model)
	model.ClosureInvariantState = EvaluateReferenceArchitectureValEInvariantState(model)
	model.ProofSurfaceState = EvaluateReferenceArchitectureValEProofSurfaceState(model)
	model.PassRuleState = EvaluateReferenceArchitectureValEPassRuleState(model)
	model.Point6State = EvaluateReferenceArchitecturePoint6FinalState(model)
	model.Point6PassAllowed = model.Point6State == ReferenceArchitecturePoint6StatePass
	model.CurrentState = EvaluateReferenceArchitectureValEState(model)
	model.BlockingReasons = referenceArchitectureValEBlockingReasons(model, invariants)
	return model
}
