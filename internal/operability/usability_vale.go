package operability

import "strings"

const (
	ProductionUsabilityValEDependencyClosureStateActive     = "production_usability_vale_dependency_closure_active"
	ProductionUsabilityValEDependencyClosureStatePartial    = "production_usability_vale_dependency_closure_partial"
	ProductionUsabilityValEDependencyClosureStateIncomplete = "production_usability_vale_dependency_closure_incomplete"

	ProductionUsabilityValECoherenceReviewStateActive     = "production_usability_vale_coherence_review_active"
	ProductionUsabilityValECoherenceReviewStatePartial    = "production_usability_vale_coherence_review_partial"
	ProductionUsabilityValECoherenceReviewStateIncomplete = "production_usability_vale_coherence_review_incomplete"

	ProductionUsabilityValEPassRuleStateActive     = "production_usability_vale_pass_rule_active"
	ProductionUsabilityValEPassRuleStatePartial    = "production_usability_vale_pass_rule_partial"
	ProductionUsabilityValEPassRuleStateIncomplete = "production_usability_vale_pass_rule_incomplete"

	ProductionUsabilityValECanonicalBoundaryReviewStateActive     = "production_usability_vale_canonical_boundary_review_active"
	ProductionUsabilityValECanonicalBoundaryReviewStatePartial    = "production_usability_vale_canonical_boundary_review_partial"
	ProductionUsabilityValECanonicalBoundaryReviewStateIncomplete = "production_usability_vale_canonical_boundary_review_incomplete"

	ProductionUsabilityValERedactionExportReviewStateActive     = "production_usability_vale_redaction_export_review_active"
	ProductionUsabilityValERedactionExportReviewStatePartial    = "production_usability_vale_redaction_export_review_partial"
	ProductionUsabilityValERedactionExportReviewStateIncomplete = "production_usability_vale_redaction_export_review_incomplete"

	ProductionUsabilityValESupportabilityRecoveryReviewStateActive     = "production_usability_vale_supportability_recovery_review_active"
	ProductionUsabilityValESupportabilityRecoveryReviewStatePartial    = "production_usability_vale_supportability_recovery_review_partial"
	ProductionUsabilityValESupportabilityRecoveryReviewStateIncomplete = "production_usability_vale_supportability_recovery_review_incomplete"

	ProductionUsabilityValERegressionClosureStateActive     = "production_usability_vale_regression_closure_active"
	ProductionUsabilityValERegressionClosureStatePartial    = "production_usability_vale_regression_closure_partial"
	ProductionUsabilityValERegressionClosureStateIncomplete = "production_usability_vale_regression_closure_incomplete"

	ProductionUsabilityValEStateIncomplete  = "production_usability_vale_incomplete"
	ProductionUsabilityValEStateSubstantial = "production_usability_vale_substantially_ready"
	ProductionUsabilityValEStateActive      = "production_usability_vale_active"

	ProductionUsabilityPoint4StatePass = "production_usability_point_4_pass"

	ProductionUsabilityDependencyPass        = "pass"
	ProductionUsabilityDependencyFail        = "fail"
	ProductionUsabilityDependencyIncomplete  = "incomplete"
	ProductionUsabilityDependencyPartial     = "partial"
	ProductionUsabilityDependencyUnsupported = "unsupported"
)

type IntegratedDependencyClosure struct {
	CurrentState           string   `json:"current_state"`
	Val0State              string   `json:"val_0_state"`
	ValAState              string   `json:"val_a_state"`
	ValBState              string   `json:"val_b_state"`
	ValCState              string   `json:"val_c_state"`
	ValDState              string   `json:"val_d_state"`
	ValEState              string   `json:"val_e_state"`
	DependencyStatus       string   `json:"dependency_status"`
	MissingVals            []string `json:"missing_vals,omitempty"`
	InactiveVals           []string `json:"inactive_vals,omitempty"`
	InconsistentVals       []string `json:"inconsistent_vals,omitempty"`
	DependencyEvidenceRefs []string `json:"dependency_evidence_refs,omitempty"`
	DependencySurfaceRefs  []string `json:"dependency_surface_refs,omitempty"`
	ClosureGeneratedAt     string   `json:"closure_generated_at"`
	ProofStatesObserved    bool     `json:"proof_states_observed"`
	ProjectionDisclaimer   string   `json:"projection_disclaimer"`
}

type CrossValCoherenceReview struct {
	CurrentState                 string   `json:"current_state"`
	CoherenceState               string   `json:"coherence_state"`
	SupportedCoherenceStates     []string `json:"supported_coherence_states,omitempty"`
	CheckedLinks                 []string `json:"checked_links,omitempty"`
	MissingLinks                 []string `json:"missing_links,omitempty"`
	InconsistentLinks            []string `json:"inconsistent_links,omitempty"`
	CarriedForwardLimitations    []string `json:"carried_forward_limitations,omitempty"`
	EvidenceRefs                 []string `json:"evidence_refs,omitempty"`
	SurfaceRefs                  []string `json:"surface_refs,omitempty"`
	Val0ContractsUsedByLaterVals bool     `json:"val_0_contracts_used_by_later_vals"`
	ValARespectedByLaterVals     bool     `json:"val_a_respected_by_later_vals"`
	ValBRespectedByLaterVals     bool     `json:"val_b_respected_by_later_vals"`
	ValCRespectedByValD          bool     `json:"val_c_respected_by_val_d"`
	ValDFinalGateCoversPriorVals bool     `json:"val_d_final_gate_covers_prior_vals"`
	NoPriorValClaimsPoint4Pass   bool     `json:"no_prior_val_claims_point_4_completion"`
	LimitationsCarriedForward    bool     `json:"limitations_carried_forward"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
}

type Point4IntegratedPassRule struct {
	CurrentState         string   `json:"current_state"`
	Point4State          string   `json:"point_4_state"`
	PassCriteriaMet      bool     `json:"pass_criteria_met"`
	PassBlockers         []string `json:"pass_blockers,omitempty"`
	PassWarnings         []string `json:"pass_warnings,omitempty"`
	PassLimitations      []string `json:"pass_limitations,omitempty"`
	RequiredVals         []string `json:"required_vals,omitempty"`
	ActiveVals           []string `json:"active_vals,omitempty"`
	MissingVals          []string `json:"missing_vals,omitempty"`
	PartialVals          []string `json:"partial_vals,omitempty"`
	UnsupportedVals      []string `json:"unsupported_vals,omitempty"`
	ValEState            string   `json:"val_e_state"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type IntegratedCanonicalTruthBoundaryReview struct {
	CurrentState                      string   `json:"current_state"`
	BoundaryState                     string   `json:"boundary_state"`
	SupportedBoundaryStates           []string `json:"supported_boundary_states,omitempty"`
	CheckedSurfaces                   []string `json:"checked_surfaces,omitempty"`
	ProjectionSurfaces                []string `json:"projection_surfaces,omitempty"`
	ViolationSurfaces                 []string `json:"violation_surfaces,omitempty"`
	GovernanceRefs                    []string `json:"governance_refs,omitempty"`
	EvidenceRefs                      []string `json:"evidence_refs,omitempty"`
	LimitationMessage                 string   `json:"limitation_message"`
	EvidenceSpineRemainsCanonical     bool     `json:"evidence_spine_remains_canonical"`
	EffectiveConfigProjectionOnly     bool     `json:"effective_config_projection_only"`
	UIAndCacheProjectionOnly          bool     `json:"ui_and_cache_projection_only"`
	DryRunAuditProjectionOnly         bool     `json:"dry_run_audit_projection_only"`
	SupportProjectionOnly             bool     `json:"support_projection_only"`
	ValDGateProjectionOnly            bool     `json:"val_d_gate_projection_only"`
	IntegratedSummaryProjectionOnly   bool     `json:"integrated_summary_projection_only"`
	ProjectionClaimsCanonicalTruth    bool     `json:"projection_claims_canonical_truth"`
	AdvisoryMutationWithoutGovernance bool     `json:"advisory_mutation_without_governance"`
	PublicOrPartnerRawEvidenceLeaked  bool     `json:"public_or_partner_raw_internal_evidence_leaked"`
}

type IntegratedPermissionRedactionExportReview struct {
	CurrentState                      string   `json:"current_state"`
	RedactionState                    string   `json:"redaction_state"`
	CheckedScopes                     []string `json:"checked_scopes,omitempty"`
	MissingScopes                     []string `json:"missing_scopes,omitempty"`
	UnsafeScopes                      []string `json:"unsafe_scopes,omitempty"`
	ExportSafetyState                 string   `json:"export_safety_state"`
	EvidenceVisibilitySummary         []string `json:"evidence_visibility_summary,omitempty"`
	LimitationMessage                 string   `json:"limitation_message"`
	HiddenMetadataRepresented         bool     `json:"hidden_metadata_represented"`
	RedactedMetadataRepresented       bool     `json:"redacted_metadata_represented"`
	PartnerOrPublicExposeFullEvidence bool     `json:"partner_or_public_expose_full_raw_evidence"`
	RawSecretsOrTokensDetected        bool     `json:"raw_secrets_or_tokens_detected"`
	RedactionConvertedFailToPass      bool     `json:"redaction_converted_fail_to_pass"`
	AuditorImpliesPublicSafe          bool     `json:"auditor_safe_implies_public_safe"`
}

type IntegratedSupportabilityRecoveryReview struct {
	CurrentState                       string   `json:"current_state"`
	SupportabilityState                string   `json:"supportability_state"`
	ReadinessState                     string   `json:"readiness_state"`
	DiagnosticsState                   string   `json:"diagnostics_state"`
	RecoveryState                      string   `json:"recovery_state"`
	UpgradeAdvisoryState               string   `json:"upgrade_advisory_state"`
	Blockers                           []string `json:"blockers,omitempty"`
	Warnings                           []string `json:"warnings,omitempty"`
	Limitations                        []string `json:"limitations,omitempty"`
	SupportabilityOverridesFailedProof bool     `json:"supportability_overrides_failed_proof_state"`
	FailedProofStateObserved           bool     `json:"failed_proof_state_observed"`
	RecoveryPolicyBypassRecommended    bool     `json:"recovery_policy_bypass_recommended"`
	RecoveryUnsafeRetryRecommended     bool     `json:"recovery_unsafe_retry_recommended"`
	UpgradeAdvisoryMutatesState        bool     `json:"upgrade_advisory_mutates_state"`
	HealthSnapshotOverridesProofState  bool     `json:"health_snapshot_overrides_proof_state"`
	ProjectionDisclaimer               string   `json:"projection_disclaimer"`
}

type IntegratedUsabilityRegressionClosure struct {
	CurrentState                 string   `json:"current_state"`
	RegressionState              string   `json:"regression_state"`
	CoveredCategories            []string `json:"covered_categories,omitempty"`
	MissingCategories            []string `json:"missing_categories,omitempty"`
	CriticalMissingCategories    []string `json:"critical_missing_categories,omitempty"`
	LimitationMessage            string   `json:"limitation_message"`
	Val0Coverage                 bool     `json:"val_0_coverage"`
	ValACoverage                 bool     `json:"val_a_coverage"`
	ValBCoverage                 bool     `json:"val_b_coverage"`
	ValCCoverage                 bool     `json:"val_c_coverage"`
	ValDCoverage                 bool     `json:"val_d_coverage"`
	IntegratedDependencyCoverage bool     `json:"integrated_dependency_closure_coverage"`
}

func productionUsabilityValERequiredVals() []string {
	return []string{"val_0", "val_a", "val_b", "val_c", "val_d", "val_e"}
}

func productionUsabilityValEDependencyStatuses() []string {
	return []string{
		ProductionUsabilityDependencyPass,
		ProductionUsabilityDependencyFail,
		ProductionUsabilityDependencyIncomplete,
		ProductionUsabilityDependencyPartial,
		ProductionUsabilityDependencyUnsupported,
	}
}

func productionUsabilityValEReviewStates() []string {
	return []string{
		ProductionUsabilityFinalGatePass,
		ProductionUsabilityFinalGateFail,
		ProductionUsabilityFinalGateWarning,
		ProductionUsabilityFinalGateBlocked,
		ProductionUsabilityFinalGateUnsupported,
		ProductionUsabilityFinalGateNotRun,
	}
}

func productionUsabilityValECoherenceLinks() []string {
	return []string{
		"val0.contracts->vala.config_explainability_core",
		"val0.contracts->valb.resilience_action_modes",
		"vala.config_explainability->valb.ui_api_cli_resilience",
		"vala.config_explainability->valc.supportability_lifecycle",
		"valb.resilience->valc.supportability_lifecycle",
		"valc.supportability->vald.final_usability_gate",
		"vald.final_gate->vale.integrated_closure",
		"prior_vals->point4_not_complete_until_vale",
	}
}

func productionUsabilityValECheckedScopes() []string {
	return []string{
		ProductionUsabilityVisibilityInternalAdmin,
		ProductionUsabilityVisibilityOperator,
		ProductionUsabilityVisibilityDeveloper,
		ProductionUsabilityVisibilityPartner,
		ProductionUsabilityVisibilityPublicSafe,
	}
}

func productionUsabilityValERegressionCategories() []string {
	return []string{
		"config_failure",
		"unknown_fields",
		"rejection_explainability",
		"stale_degraded_partial",
		"cli_retry_idempotency",
		"api_backpressure_fairness",
		"support_bundle_redaction",
		"upgrade_rollback_advisory",
		"canonical_truth_boundary",
		"integrated_dependency_closure",
	}
}

func ProductionUsabilityValEDependencyClosure() IntegratedDependencyClosure {
	return IntegratedDependencyClosure{
		CurrentState:     "production_usability_vale_dependency_closure_ready",
		Val0State:        ProductionUsabilityVal0StateActive,
		ValAState:        ProductionUsabilityValAStateActive,
		ValBState:        ProductionUsabilityValBStateActive,
		ValCState:        ProductionUsabilityValCStateActive,
		ValDState:        ProductionUsabilityValDStateActive,
		ValEState:        ProductionUsabilityValEStateActive,
		DependencyStatus: ProductionUsabilityDependencyPass,
		DependencyEvidenceRefs: []string{
			"val0_proofs",
			"vala_proofs",
			"valb_proofs",
			"valc_proofs",
			"vald_proofs",
		},
		DependencySurfaceRefs: []string{
			"/v1/production/usability-operability-recovery/val0/proofs",
			"/v1/production/usability-operability-recovery/vala/proofs",
			"/v1/production/usability-operability-recovery/valb/proofs",
			"/v1/production/usability-operability-recovery/valc/proofs",
			"/v1/production/usability-operability-recovery/vald/proofs",
			"/v1/production/usability-operability-recovery/vale/dependency-closure",
			"/v1/production/usability-operability-recovery/vale/proofs",
		},
		ClosureGeneratedAt:   "generated_at_present",
		ProofStatesObserved:  true,
		ProjectionDisclaimer: "projection_only not_canonical_truth",
	}
}

func ProductionUsabilityValECrossValCoherenceReview() CrossValCoherenceReview {
	return CrossValCoherenceReview{
		CurrentState:             "production_usability_vale_coherence_review_ready",
		CoherenceState:           ProductionUsabilityFinalGatePass,
		SupportedCoherenceStates: productionUsabilityValEReviewStates(),
		CheckedLinks:             productionUsabilityValECoherenceLinks(),
		CarriedForwardLimitations: []string{
			"val0: foundation contracts remain fail-closed and projection-only.",
			"vala: config and explainability core remain bounded and non-canonical.",
			"valb: resilience surfaces remain advisory-only and projection-only.",
			"valc: supportability surfaces remain bounded, redaction-safe, and projection-only.",
			"vald: final usability gate remains review-only and non-mutating.",
		},
		EvidenceRefs: []string{
			"val0_proofs",
			"vala_proofs",
			"valb_proofs",
			"valc_proofs",
			"vald_proofs",
		},
		SurfaceRefs: []string{
			"/v1/production/usability-operability-recovery/val0/proofs",
			"/v1/production/usability-operability-recovery/vala/proofs",
			"/v1/production/usability-operability-recovery/valb/proofs",
			"/v1/production/usability-operability-recovery/valc/proofs",
			"/v1/production/usability-operability-recovery/vald/proofs",
			"/v1/production/usability-operability-recovery/vale/coherence-review",
			"/v1/production/usability-operability-recovery/vale/proofs",
		},
		Val0ContractsUsedByLaterVals: true,
		ValARespectedByLaterVals:     true,
		ValBRespectedByLaterVals:     true,
		ValCRespectedByValD:          true,
		ValDFinalGateCoversPriorVals: true,
		NoPriorValClaimsPoint4Pass:   true,
		LimitationsCarriedForward:    true,
		ProjectionDisclaimer:         "projection_only not_canonical_truth",
	}
}

func ProductionUsabilityValEPoint4PassRule() Point4IntegratedPassRule {
	return Point4IntegratedPassRule{
		CurrentState:    "production_usability_vale_pass_rule_ready",
		Point4State:     ProductionUsabilityPoint4StatePass,
		PassCriteriaMet: true,
		PassWarnings: []string{
			"Integrated closure remains projection-only and does not replace canonical truth.",
		},
		PassLimitations: []string{
			"Point 4 pass is an integrated closure summary over accepted Val 0 through Val E slices.",
		},
		RequiredVals:         productionUsabilityValERequiredVals(),
		ActiveVals:           productionUsabilityValERequiredVals(),
		ValEState:            ProductionUsabilityValEStateActive,
		ProjectionDisclaimer: "projection_only not_canonical_truth",
	}
}

func ProductionUsabilityValECanonicalTruthBoundaryReview() IntegratedCanonicalTruthBoundaryReview {
	return IntegratedCanonicalTruthBoundaryReview{
		CurrentState:            "production_usability_vale_canonical_boundary_review_ready",
		BoundaryState:           ProductionUsabilityFinalGatePass,
		SupportedBoundaryStates: productionUsabilityValEReviewStates(),
		CheckedSurfaces: []string{
			"/v1/production/usability-operability-recovery/vala/effective-config",
			"/v1/production/usability-operability-recovery/vala/policy-dry-run",
			"/v1/production/usability-operability-recovery/valb/ui-data-resilience",
			"/v1/production/usability-operability-recovery/valb/windowing",
			"/v1/production/usability-operability-recovery/valc/support-bundle",
			"/v1/production/usability-operability-recovery/valc/diagnostics",
			"/v1/production/usability-operability-recovery/valc/health-snapshot",
			"/v1/production/usability-operability-recovery/vald/proofs",
			"/v1/production/usability-operability-recovery/vale/proofs",
		},
		ProjectionSurfaces: []string{
			"effective_config",
			"dry_run",
			"ui_data",
			"windowed_results",
			"support_bundle",
			"diagnostics",
			"health_snapshot",
			"final_gate_review",
			"integrated_closure_summary",
		},
		GovernanceRefs: []string{
			"governance_boundary_review_contract",
			"canonical_truth_rule",
		},
		EvidenceRefs: []string{
			"evidence_spine",
			"vala_proofs",
			"valb_proofs",
			"valc_proofs",
			"vald_proofs",
		},
		LimitationMessage:               "projection_only not_canonical_truth integrated closure summary",
		EvidenceSpineRemainsCanonical:   true,
		EffectiveConfigProjectionOnly:   true,
		UIAndCacheProjectionOnly:        true,
		DryRunAuditProjectionOnly:       true,
		SupportProjectionOnly:           true,
		ValDGateProjectionOnly:          true,
		IntegratedSummaryProjectionOnly: true,
	}
}

func ProductionUsabilityValEIntegratedRedactionExportReview() IntegratedPermissionRedactionExportReview {
	return IntegratedPermissionRedactionExportReview{
		CurrentState:      "production_usability_vale_redaction_export_review_ready",
		RedactionState:    ProductionUsabilityFinalGatePass,
		CheckedScopes:     productionUsabilityValECheckedScopes(),
		ExportSafetyState: ProductionUsabilityValCExportSafetyStateActive,
		EvidenceVisibilitySummary: []string{
			"internal_admin:full",
			"operator:metadata_only",
			"developer:redacted",
			"partner:redacted",
			"public_safe:hidden",
		},
		LimitationMessage:           "projection_only not_canonical_truth redaction_export_review",
		HiddenMetadataRepresented:   true,
		RedactedMetadataRepresented: true,
	}
}

func ProductionUsabilityValESupportabilityRecoveryReview() IntegratedSupportabilityRecoveryReview {
	return IntegratedSupportabilityRecoveryReview{
		CurrentState:         "production_usability_vale_supportability_recovery_review_ready",
		SupportabilityState:  ProductionUsabilityFinalGatePass,
		ReadinessState:       ProductionUsabilityValCReadinessStateActive,
		DiagnosticsState:     ProductionUsabilityValCDiagnosticsStateActive,
		RecoveryState:        ProductionUsabilityValDRecoveryReviewStateActive,
		UpgradeAdvisoryState: ProductionUsabilityValDUpgradeRollbackReviewStateActive,
		Warnings: []string{
			"Health snapshots remain bounded point-in-time projections and do not override proof state.",
		},
		Limitations: []string{
			"Supportability and recovery closure remain projection-only and do not authorize bypasses.",
		},
		ProjectionDisclaimer: "projection_only not_canonical_truth supportability_recovery_review",
	}
}

func ProductionUsabilityValERegressionClosure() IntegratedUsabilityRegressionClosure {
	return IntegratedUsabilityRegressionClosure{
		CurrentState:                 "production_usability_vale_regression_closure_ready",
		RegressionState:              ProductionUsabilityFinalGatePass,
		CoveredCategories:            productionUsabilityValERegressionCategories(),
		LimitationMessage:            "Regression closure is bounded proof metadata and does not claim exhaustive testing.",
		Val0Coverage:                 true,
		ValACoverage:                 true,
		ValBCoverage:                 true,
		ValCCoverage:                 true,
		ValDCoverage:                 true,
		IntegratedDependencyCoverage: true,
	}
}

func EvaluateProductionUsabilityValEDependencyClosureState(model IntegratedDependencyClosure) string {
	if strings.TrimSpace(model.CurrentState) == "" ||
		strings.TrimSpace(model.Val0State) == "" ||
		strings.TrimSpace(model.ValAState) == "" ||
		strings.TrimSpace(model.ValBState) == "" ||
		strings.TrimSpace(model.ValCState) == "" ||
		strings.TrimSpace(model.ValDState) == "" ||
		strings.TrimSpace(model.ValEState) == "" ||
		strings.TrimSpace(model.DependencyStatus) == "" ||
		strings.TrimSpace(model.ClosureGeneratedAt) == "" {
		return ProductionUsabilityValEDependencyClosureStateIncomplete
	}
	if !containsTrimmedString(productionUsabilityValEDependencyStatuses(), model.DependencyStatus) ||
		len(model.DependencyEvidenceRefs) < 5 ||
		len(model.DependencySurfaceRefs) < 6 ||
		!model.ProofStatesObserved ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValEDependencyClosureStatePartial
	}
	if strings.TrimSpace(model.Val0State) != ProductionUsabilityVal0StateActive ||
		strings.TrimSpace(model.ValAState) != ProductionUsabilityValAStateActive ||
		strings.TrimSpace(model.ValBState) != ProductionUsabilityValBStateActive ||
		strings.TrimSpace(model.ValCState) != ProductionUsabilityValCStateActive ||
		strings.TrimSpace(model.ValDState) != ProductionUsabilityValDStateActive ||
		strings.TrimSpace(model.DependencyStatus) != ProductionUsabilityDependencyPass ||
		len(model.MissingVals) > 0 ||
		len(model.InactiveVals) > 0 ||
		len(model.InconsistentVals) > 0 {
		return ProductionUsabilityValEDependencyClosureStatePartial
	}
	return ProductionUsabilityValEDependencyClosureStateActive
}

func EvaluateProductionUsabilityValECoherenceReviewState(model CrossValCoherenceReview) string {
	if strings.TrimSpace(model.CurrentState) == "" ||
		strings.TrimSpace(model.CoherenceState) == "" {
		return ProductionUsabilityValECoherenceReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedCoherenceStates, productionUsabilityValEReviewStates()...) ||
		len(model.CheckedLinks) == 0 ||
		len(model.CarriedForwardLimitations) == 0 ||
		len(model.EvidenceRefs) == 0 ||
		len(model.SurfaceRefs) == 0 ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValECoherenceReviewStatePartial
	}
	if !containsAllTrimmedStrings(model.CheckedLinks, productionUsabilityValECoherenceLinks()...) ||
		len(model.MissingLinks) > 0 ||
		len(model.InconsistentLinks) > 0 ||
		!model.Val0ContractsUsedByLaterVals ||
		!model.ValARespectedByLaterVals ||
		!model.ValBRespectedByLaterVals ||
		!model.ValCRespectedByValD ||
		!model.ValDFinalGateCoversPriorVals ||
		!model.NoPriorValClaimsPoint4Pass ||
		!model.LimitationsCarriedForward ||
		strings.TrimSpace(model.CoherenceState) != ProductionUsabilityFinalGatePass {
		return ProductionUsabilityValECoherenceReviewStatePartial
	}
	return ProductionUsabilityValECoherenceReviewStateActive
}

func EvaluateProductionUsabilityValEPassRuleState(model Point4IntegratedPassRule) string {
	if strings.TrimSpace(model.CurrentState) == "" ||
		strings.TrimSpace(model.Point4State) == "" ||
		strings.TrimSpace(model.ValEState) == "" {
		return ProductionUsabilityValEPassRuleStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.RequiredVals, productionUsabilityValERequiredVals()...) ||
		len(model.PassLimitations) == 0 ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValEPassRuleStatePartial
	}
	if strings.TrimSpace(model.Point4State) != ProductionUsabilityPoint4StatePass ||
		!model.PassCriteriaMet ||
		len(model.PassBlockers) > 0 ||
		len(model.MissingVals) > 0 ||
		len(model.PartialVals) > 0 ||
		len(model.UnsupportedVals) > 0 ||
		!containsExactTrimmedStringSet(model.ActiveVals, productionUsabilityValERequiredVals()...) ||
		strings.TrimSpace(model.ValEState) != ProductionUsabilityValEStateActive {
		return ProductionUsabilityValEPassRuleStatePartial
	}
	return ProductionUsabilityValEPassRuleStateActive
}

func EvaluateProductionUsabilityValECanonicalBoundaryReviewState(model IntegratedCanonicalTruthBoundaryReview) string {
	if strings.TrimSpace(model.CurrentState) == "" ||
		strings.TrimSpace(model.BoundaryState) == "" ||
		strings.TrimSpace(model.LimitationMessage) == "" {
		return ProductionUsabilityValECanonicalBoundaryReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedBoundaryStates, productionUsabilityValEReviewStates()...) ||
		len(model.CheckedSurfaces) == 0 ||
		len(model.ProjectionSurfaces) == 0 ||
		len(model.GovernanceRefs) == 0 ||
		len(model.EvidenceRefs) == 0 ||
		!strings.Contains(strings.TrimSpace(model.LimitationMessage), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.LimitationMessage), "not_canonical_truth") {
		return ProductionUsabilityValECanonicalBoundaryReviewStatePartial
	}
	if !model.EvidenceSpineRemainsCanonical ||
		!model.EffectiveConfigProjectionOnly ||
		!model.UIAndCacheProjectionOnly ||
		!model.DryRunAuditProjectionOnly ||
		!model.SupportProjectionOnly ||
		!model.ValDGateProjectionOnly ||
		!model.IntegratedSummaryProjectionOnly ||
		model.ProjectionClaimsCanonicalTruth ||
		model.AdvisoryMutationWithoutGovernance ||
		model.PublicOrPartnerRawEvidenceLeaked ||
		len(model.ViolationSurfaces) > 0 ||
		strings.TrimSpace(model.BoundaryState) != ProductionUsabilityFinalGatePass {
		return ProductionUsabilityValECanonicalBoundaryReviewStatePartial
	}
	return ProductionUsabilityValECanonicalBoundaryReviewStateActive
}

func EvaluateProductionUsabilityValERedactionExportReviewState(model IntegratedPermissionRedactionExportReview) string {
	if strings.TrimSpace(model.CurrentState) == "" ||
		strings.TrimSpace(model.RedactionState) == "" ||
		strings.TrimSpace(model.ExportSafetyState) == "" ||
		strings.TrimSpace(model.LimitationMessage) == "" {
		return ProductionUsabilityValERedactionExportReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.CheckedScopes, productionUsabilityValECheckedScopes()...) ||
		len(model.EvidenceVisibilitySummary) == 0 ||
		!strings.Contains(strings.TrimSpace(model.LimitationMessage), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.LimitationMessage), "not_canonical_truth") {
		return ProductionUsabilityValERedactionExportReviewStatePartial
	}
	if len(model.MissingScopes) > 0 ||
		len(model.UnsafeScopes) > 0 ||
		!model.HiddenMetadataRepresented ||
		!model.RedactedMetadataRepresented ||
		model.PartnerOrPublicExposeFullEvidence ||
		model.RawSecretsOrTokensDetected ||
		model.RedactionConvertedFailToPass ||
		model.AuditorImpliesPublicSafe ||
		strings.TrimSpace(model.RedactionState) != ProductionUsabilityFinalGatePass ||
		strings.TrimSpace(model.ExportSafetyState) != ProductionUsabilityValCExportSafetyStateActive {
		return ProductionUsabilityValERedactionExportReviewStatePartial
	}
	return ProductionUsabilityValERedactionExportReviewStateActive
}

func EvaluateProductionUsabilityValESupportabilityRecoveryReviewState(model IntegratedSupportabilityRecoveryReview) string {
	if strings.TrimSpace(model.CurrentState) == "" ||
		strings.TrimSpace(model.SupportabilityState) == "" ||
		strings.TrimSpace(model.ReadinessState) == "" ||
		strings.TrimSpace(model.DiagnosticsState) == "" ||
		strings.TrimSpace(model.RecoveryState) == "" ||
		strings.TrimSpace(model.UpgradeAdvisoryState) == "" {
		return ProductionUsabilityValESupportabilityRecoveryReviewStateIncomplete
	}
	if len(model.Limitations) == 0 ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValESupportabilityRecoveryReviewStatePartial
	}
	if strings.TrimSpace(model.SupportabilityState) != ProductionUsabilityFinalGatePass ||
		strings.TrimSpace(model.ReadinessState) != ProductionUsabilityValCReadinessStateActive ||
		strings.TrimSpace(model.DiagnosticsState) != ProductionUsabilityValCDiagnosticsStateActive ||
		strings.TrimSpace(model.RecoveryState) != ProductionUsabilityValDRecoveryReviewStateActive ||
		strings.TrimSpace(model.UpgradeAdvisoryState) != ProductionUsabilityValDUpgradeRollbackReviewStateActive ||
		len(model.Blockers) > 0 ||
		model.SupportabilityOverridesFailedProof ||
		model.FailedProofStateObserved ||
		model.RecoveryPolicyBypassRecommended ||
		model.RecoveryUnsafeRetryRecommended ||
		model.UpgradeAdvisoryMutatesState ||
		model.HealthSnapshotOverridesProofState {
		return ProductionUsabilityValESupportabilityRecoveryReviewStatePartial
	}
	return ProductionUsabilityValESupportabilityRecoveryReviewStateActive
}

func EvaluateProductionUsabilityValERegressionClosureState(model IntegratedUsabilityRegressionClosure) string {
	if strings.TrimSpace(model.CurrentState) == "" ||
		strings.TrimSpace(model.RegressionState) == "" ||
		strings.TrimSpace(model.LimitationMessage) == "" {
		return ProductionUsabilityValERegressionClosureStateIncomplete
	}
	if !model.Val0Coverage ||
		!model.ValACoverage ||
		!model.ValBCoverage ||
		!model.ValCCoverage ||
		!model.ValDCoverage ||
		!model.IntegratedDependencyCoverage ||
		!containsAllTrimmedStrings(model.CoveredCategories, productionUsabilityValERegressionCategories()...) ||
		len(model.CriticalMissingCategories) > 0 ||
		strings.TrimSpace(model.RegressionState) != ProductionUsabilityFinalGatePass {
		return ProductionUsabilityValERegressionClosureStatePartial
	}
	return ProductionUsabilityValERegressionClosureStateActive
}

func EvaluateProductionUsabilityValEPrerequisiteState(val0State, valAState, valBState, valCState, valDState, dependencyClosureState, coherenceReviewState, canonicalBoundaryReviewState, redactionExportReviewState, supportabilityRecoveryReviewState, regressionClosureState string) string {
	if strings.TrimSpace(val0State) != ProductionUsabilityVal0StateActive ||
		strings.TrimSpace(valAState) != ProductionUsabilityValAStateActive ||
		strings.TrimSpace(valBState) != ProductionUsabilityValBStateActive ||
		strings.TrimSpace(valCState) != ProductionUsabilityValCStateActive ||
		strings.TrimSpace(valDState) != ProductionUsabilityValDStateActive {
		return ProductionUsabilityValEStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		dependencyClosureState,
		coherenceReviewState,
		canonicalBoundaryReviewState,
		redactionExportReviewState,
		supportabilityRecoveryReviewState,
		regressionClosureState,
	} {
		switch strings.TrimSpace(state) {
		case ProductionUsabilityValEDependencyClosureStateActive,
			ProductionUsabilityValECoherenceReviewStateActive,
			ProductionUsabilityValECanonicalBoundaryReviewStateActive,
			ProductionUsabilityValERedactionExportReviewStateActive,
			ProductionUsabilityValESupportabilityRecoveryReviewStateActive,
			ProductionUsabilityValERegressionClosureStateActive:
		case ProductionUsabilityValEDependencyClosureStatePartial,
			ProductionUsabilityValECoherenceReviewStatePartial,
			ProductionUsabilityValECanonicalBoundaryReviewStatePartial,
			ProductionUsabilityValERedactionExportReviewStatePartial,
			ProductionUsabilityValESupportabilityRecoveryReviewStatePartial,
			ProductionUsabilityValERegressionClosureStatePartial:
			hasPartial = true
		default:
			return ProductionUsabilityValEStateIncomplete
		}
	}
	if hasPartial {
		return ProductionUsabilityValEStateSubstantial
	}
	return ProductionUsabilityValEStateActive
}

func EvaluateProductionUsabilityValEState(val0State, valAState, valBState, valCState, valDState, dependencyClosureState, coherenceReviewState, passRuleState, canonicalBoundaryReviewState, redactionExportReviewState, supportabilityRecoveryReviewState, regressionClosureState string) string {
	baseState := EvaluateProductionUsabilityValEPrerequisiteState(val0State, valAState, valBState, valCState, valDState, dependencyClosureState, coherenceReviewState, canonicalBoundaryReviewState, redactionExportReviewState, supportabilityRecoveryReviewState, regressionClosureState)
	if baseState != ProductionUsabilityValEStateActive {
		return baseState
	}
	switch strings.TrimSpace(passRuleState) {
	case ProductionUsabilityValEPassRuleStateActive:
		return ProductionUsabilityValEStateActive
	case ProductionUsabilityValEPassRuleStatePartial:
		return ProductionUsabilityValEStateSubstantial
	default:
		return ProductionUsabilityValEStateIncomplete
	}
}

func EvaluateProductionUsabilityValEProofsState(val0State, valAState, valBState, valCState, valDState, dependencyClosureState, coherenceReviewState, passRuleState, canonicalBoundaryReviewState, redactionExportReviewState, supportabilityRecoveryReviewState, regressionClosureState, point4State string, surfaceRefs, evidenceRefs, limitations []string) string {
	baseState := EvaluateProductionUsabilityValEState(val0State, valAState, valBState, valCState, valDState, dependencyClosureState, coherenceReviewState, passRuleState, canonicalBoundaryReviewState, redactionExportReviewState, supportabilityRecoveryReviewState, regressionClosureState)
	if len(surfaceRefs) < 13 || len(evidenceRefs) < 12 || len(limitations) == 0 {
		if baseState == ProductionUsabilityValEStateActive {
			return ProductionUsabilityValEStateSubstantial
		}
		return baseState
	}
	if baseState == ProductionUsabilityValEStateActive && strings.TrimSpace(point4State) != ProductionUsabilityPoint4StatePass {
		return ProductionUsabilityValEStateSubstantial
	}
	return baseState
}
