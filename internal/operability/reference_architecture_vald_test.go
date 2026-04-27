package operability

import "testing"

func activeReferenceArchitectureValDPrereqs() (string, string, string, string, string, string, string, string, string, string, string) {
	return IntelligenceCalibrationPoint5StatePass,
		IntelligenceCalibrationValEStateActive,
		ReferenceArchitectureVal0StateActive,
		ReferenceArchitectureVal0StateActive,
		ReferenceArchitectureValAStateActive,
		ReferenceArchitectureValAStateActive,
		ReferenceArchitectureValBStateActive,
		ReferenceArchitectureValBStateActive,
		ReferenceArchitectureValCStateActive,
		ReferenceArchitectureValCStateActive,
		ReferenceArchitecturePoint6StateNotComplete
}

func activeReferenceArchitectureValDComponents() (
	ReferenceArchitectureOperationalVisibilityCollection,
	ReferenceArchitectureBlueprintAlignmentCollection,
	ReferenceArchitectureDeviationAlertCollection,
	ReferenceArchitectureSupportBoundaryCollection,
	ReferenceArchitectureMigrationUpgradeCollection,
	ReferenceArchitectureTopologyGateCollection,
	ReferenceArchitectureSecurityBoundaryCollection,
	ReferenceArchitectureOperabilityGateCollection,
	ReferenceArchitectureCompatibilityGateCollection,
	ReferenceArchitectureFinalGateCollection,
) {
	return ReferenceArchitectureValDOperationalVisibilityCollection(),
		ReferenceArchitectureValDAlignmentSummaryCollection(),
		ReferenceArchitectureValDDeviationAlertCollection(),
		ReferenceArchitectureValDSupportBoundaryCollection(),
		ReferenceArchitectureValDMigrationUpgradeCollection(),
		ReferenceArchitectureValDTopologyGateCollection(),
		ReferenceArchitectureValDSecurityBoundaryCollection(),
		ReferenceArchitectureValDOperabilityGateCollection(),
		ReferenceArchitectureValDCompatibilityGateCollection(),
		ReferenceArchitectureValDFinalGateCollection()
}

func validReferenceArchitectureValDAlert() ReferenceArchitectureDeviationAlert {
	return ReferenceArchitectureDeviationAlert{
		AlertID:                "alert-enterprise-default",
		BlueprintFamily:        ReferenceArchitectureFamilyEnterpriseDefault,
		SourceLayer:            "valb",
		DeviationCategory:      ReferenceArchitectureValDDeviationReadinessGap,
		Severity:               ReferenceArchitectureValBSeverityMedium,
		AffectedScope:          "enterprise_default/readiness",
		EvidenceRef:            "evidence-enterprise-default",
		BlocksAlignment:        false,
		OperatorActionRequired: "review bounded readiness gap",
		SupportBoundaryRef:     "support-boundary/enterprise_default",
		Timestamp:              "2026-04-26T10:00:00Z",
		FreshnessState:         IntelligenceCalibrationFreshnessFresh,
	}
}

func TestReferenceArchitectureValDDependencyGates(t *testing.T) {
	point5State, point5DependencyState, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, valCCurrentState, valCState, point6State := activeReferenceArchitectureValDPrereqs()
	visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate := activeReferenceArchitectureValDComponents()

	visibilityState := EvaluateReferenceArchitectureValDOperationalVisibilityCollectionState(visibility)
	alignmentState := EvaluateReferenceArchitectureValDAlignmentSummaryCollectionState(alignment)
	alertState := EvaluateReferenceArchitectureValDDeviationAlertCollectionState(alerts)
	supportState := EvaluateReferenceArchitectureValDSupportBoundaryCollectionState(support)
	migrationState := EvaluateReferenceArchitectureValDMigrationUpgradeCollectionState(migration)
	topologyState := EvaluateReferenceArchitectureValDTopologyGateCollectionState(topology)
	securityState := EvaluateReferenceArchitectureValDSecurityBoundaryCollectionState(security)
	operabilityState := EvaluateReferenceArchitectureValDOperabilityGateCollectionState(operabilityCollection)
	compatibilityState := EvaluateReferenceArchitectureValDCompatibilityGateCollectionState(compatibility)
	finalGateState := EvaluateReferenceArchitectureValDFinalGateCollectionState(finalGate)

	if got := EvaluateReferenceArchitectureValDState(point5State, point5DependencyState, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, valCCurrentState, valCState, point6State, visibilityState, alignmentState, alertState, supportState, migrationState, topologyState, securityState, operabilityState, compatibilityState, finalGateState); got != ReferenceArchitectureValDStateActive {
		t.Fatalf("expected active Val D state with valid dependencies and components, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValDState(point5State, IntelligenceCalibrationValEStateSubstantial, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, valCCurrentState, valCState, point6State, visibilityState, alignmentState, alertState, supportState, migrationState, topologyState, securityState, operabilityState, compatibilityState, finalGateState); got != ReferenceArchitectureValDStateBlocked {
		t.Fatalf("expected blocked Val D state when point 5 dependency health regresses, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValDState(point5State, point5DependencyState, ReferenceArchitectureVal0StateSubstantial, val0State, valACurrentState, valAState, valBCurrentState, valBState, valCCurrentState, valCState, point6State, visibilityState, alignmentState, alertState, supportState, migrationState, topologyState, securityState, operabilityState, compatibilityState, finalGateState); got != ReferenceArchitectureValDStateBlocked {
		t.Fatalf("expected blocked Val D state when Val 0 dependency is missing, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValDState(point5State, point5DependencyState, val0CurrentState, val0State, ReferenceArchitectureValAStatePartial, valAState, valBCurrentState, valBState, valCCurrentState, valCState, point6State, visibilityState, alignmentState, alertState, supportState, migrationState, topologyState, securityState, operabilityState, compatibilityState, finalGateState); got != ReferenceArchitectureValDStateBlocked {
		t.Fatalf("expected blocked Val D state when Val A dependency is missing, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValDState(point5State, point5DependencyState, val0CurrentState, val0State, valACurrentState, valAState, ReferenceArchitectureValBStatePartial, valBState, valCCurrentState, valCState, point6State, visibilityState, alignmentState, alertState, supportState, migrationState, topologyState, securityState, operabilityState, compatibilityState, finalGateState); got != ReferenceArchitectureValDStateBlocked {
		t.Fatalf("expected blocked Val D state when Val B dependency is missing, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValDState(point5State, point5DependencyState, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, ReferenceArchitectureValCStatePartial, valCState, point6State, visibilityState, alignmentState, alertState, supportState, migrationState, topologyState, securityState, operabilityState, compatibilityState, finalGateState); got != ReferenceArchitectureValDStateBlocked {
		t.Fatalf("expected blocked Val D state when Val C dependency is missing, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValDState(point5State, point5DependencyState, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, valCCurrentState, valCState, ReferenceArchitecturePoint6StatePass, visibilityState, alignmentState, alertState, supportState, migrationState, topologyState, securityState, operabilityState, compatibilityState, finalGateState); got != ReferenceArchitectureValDStateBlocked {
		t.Fatalf("expected blocked Val D state when point 6 is not not_complete, got %q", got)
	}
}

func TestReferenceArchitectureValDOperationalVisibilityValidation(t *testing.T) {
	report := ReferenceArchitectureValDOperationalVisibilityCollection().Reports[0]
	if got := EvaluateReferenceArchitectureValDOperationalVisibilityReportState(report); got != ReferenceArchitectureValDVisibilityStateActive {
		t.Fatalf("expected active operational visibility state, got %q", got)
	}
	report = ReferenceArchitectureValDOperationalVisibilityCollection().Reports[0]
	report.VisibilityReportID = ""
	if got := EvaluateReferenceArchitectureValDOperationalVisibilityReportState(report); got == ReferenceArchitectureValDVisibilityStateActive {
		t.Fatalf("expected non-active visibility state for missing report id, got %q", got)
	}
	report = ReferenceArchitectureValDOperationalVisibilityCollection().Reports[0]
	report.BlueprintFamily = "enterprise-defualt"
	if got := EvaluateReferenceArchitectureValDOperationalVisibilityReportState(report); got == ReferenceArchitectureValDVisibilityStateActive {
		t.Fatalf("expected non-active visibility state for unknown family, got %q", got)
	}
	report = ReferenceArchitectureValDOperationalVisibilityCollection().Reports[0]
	report.AlignmentStatus = "macthed"
	if got := EvaluateReferenceArchitectureValDOperationalVisibilityReportState(report); got == ReferenceArchitectureValDVisibilityStateActive {
		t.Fatalf("expected non-active visibility state for typo alignment status, got %q", got)
	}
	report = ReferenceArchitectureValDOperationalVisibilityCollection().Reports[0]
	report.ProjectionDisclaimer = ""
	if got := EvaluateReferenceArchitectureValDOperationalVisibilityReportState(report); got == ReferenceArchitectureValDVisibilityStateActive {
		t.Fatalf("expected non-active visibility state without projection disclaimer, got %q", got)
	}
	report = ReferenceArchitectureValDOperationalVisibilityCollection().Reports[0]
	report.EvidenceRefs[0].FreshnessState = IntelligenceCalibrationFreshnessStale
	if got := EvaluateReferenceArchitectureValDOperationalVisibilityReportState(report); got == ReferenceArchitectureValDVisibilityStateActive {
		t.Fatalf("expected stale evidence to prevent active visibility state, got %q", got)
	}
	report = ReferenceArchitectureValDOperationalVisibilityCollection().Reports[0]
	report.GuaranteedSecurityClaim = true
	if got := EvaluateReferenceArchitectureValDOperationalVisibilityReportState(report); got != ReferenceArchitectureValDVisibilityStateBlocked {
		t.Fatalf("expected overclaim language to block visibility state, got %q", got)
	}
}

func TestReferenceArchitectureValDAlignmentSummaryValidation(t *testing.T) {
	summary := ReferenceArchitectureValDAlignmentSummaryCollection().Summaries[0]
	if got := EvaluateReferenceArchitectureValDAlignmentSummaryState(summary); got != ReferenceArchitectureValDAlignmentStateActive {
		t.Fatalf("expected active alignment summary state, got %q", got)
	}
	summary = ReferenceArchitectureValDAlignmentSummaryCollection().Summaries[0]
	summary.ValCState = ReferenceArchitectureValCStatePartial
	if got := EvaluateReferenceArchitectureValDAlignmentSummaryState(summary); got == ReferenceArchitectureValDAlignmentStateActive {
		t.Fatalf("expected degraded source state to prevent matched alignment, got %q", got)
	}
	summary = ReferenceArchitectureValDAlignmentSummaryCollection().Summaries[0]
	summary.AlignmentStatus = ReferenceArchitectureConformanceUnsupported
	if got := EvaluateReferenceArchitectureValDAlignmentSummaryState(summary); got == ReferenceArchitectureValDAlignmentStateActive {
		t.Fatalf("expected unsupported alignment status to prevent active summary, got %q", got)
	}
	summary = ReferenceArchitectureValDAlignmentSummaryCollection().Summaries[0]
	summary.StaleEvidenceRefs = []string{"stale-evidence"}
	if got := EvaluateReferenceArchitectureValDAlignmentSummaryState(summary); got == ReferenceArchitectureValDAlignmentStateActive {
		t.Fatalf("expected stale source evidence to prevent active summary, got %q", got)
	}
	summary = ReferenceArchitectureValDAlignmentSummaryCollection().Summaries[0]
	summary.RedactionKeepsBlockingVisible = false
	if got := EvaluateReferenceArchitectureValDAlignmentSummaryState(summary); got == ReferenceArchitectureValDAlignmentStateActive {
		t.Fatalf("expected redaction hiding blocking deviations to fail closed, got %q", got)
	}
}

func TestReferenceArchitectureValDDeviationAlertsValidation(t *testing.T) {
	alert := validReferenceArchitectureValDAlert()
	if got := EvaluateReferenceArchitectureValDDeviationAlertState(alert); got != ReferenceArchitectureValDAlertStateActive {
		t.Fatalf("expected active alert state, got %q", got)
	}
	alert = validReferenceArchitectureValDAlert()
	alert.DeviationCategory = "readines_gap"
	if got := EvaluateReferenceArchitectureValDDeviationAlertState(alert); got == ReferenceArchitectureValDAlertStateActive {
		t.Fatalf("expected unknown category to fail closed, got %q", got)
	}
	alert = validReferenceArchitectureValDAlert()
	alert.Severity = "sev0"
	if got := EvaluateReferenceArchitectureValDDeviationAlertState(alert); got == ReferenceArchitectureValDAlertStateActive {
		t.Fatalf("expected unknown severity to fail closed, got %q", got)
	}
	alert = validReferenceArchitectureValDAlert()
	alert.BlocksAlignment = true
	alert.EvidenceRef = ""
	if got := EvaluateReferenceArchitectureValDDeviationAlertState(alert); got == ReferenceArchitectureValDAlertStateActive {
		t.Fatalf("expected blocking alert without evidence to fail closed, got %q", got)
	}
	alert = validReferenceArchitectureValDAlert()
	alert.FreshnessState = IntelligenceCalibrationFreshnessStale
	if got := EvaluateReferenceArchitectureValDDeviationAlertState(alert); got == ReferenceArchitectureValDAlertStateActive {
		t.Fatalf("expected stale alert evidence to prevent active state, got %q", got)
	}
	alert = validReferenceArchitectureValDAlert()
	alert.AdvisoryOnly = true
	alert.BlocksAlignment = true
	if got := EvaluateReferenceArchitectureValDDeviationAlertState(alert); got == ReferenceArchitectureValDAlertStateActive {
		t.Fatalf("expected advisory_only not to override blocking semantics, got %q", got)
	}
	alert = validReferenceArchitectureValDAlert()
	alert.DeviationCategory = ReferenceArchitectureValDDeviationOverclaimDetected
	alert.BlocksAlignment = true
	if got := EvaluateReferenceArchitectureValDDeviationAlertState(alert); got != ReferenceArchitectureValDAlertStateBlocked {
		t.Fatalf("expected overclaim alert to block clean gate state, got %q", got)
	}
}

func TestReferenceArchitectureValDSupportBoundaryValidation(t *testing.T) {
	view := ReferenceArchitectureValDSupportBoundaryCollection().Views[0]
	if got := EvaluateReferenceArchitectureValDSupportBoundaryViewState(view); got != ReferenceArchitectureValDSupportBoundaryStateActive {
		t.Fatalf("expected active support boundary state, got %q", got)
	}
	view = ReferenceArchitectureValDSupportBoundaryCollection().Views[0]
	view.SupportBoundaryRef = ""
	if got := EvaluateReferenceArchitectureValDSupportBoundaryViewState(view); got == ReferenceArchitectureValDSupportBoundaryStateActive {
		t.Fatalf("expected missing support boundary to fail closed, got %q", got)
	}
	view = ReferenceArchitectureValDSupportBoundaryCollection().Views[0]
	view.PartnerCanonicalAuthority = true
	if got := EvaluateReferenceArchitectureValDSupportBoundaryViewState(view); got == ReferenceArchitectureValDSupportBoundaryStateActive {
		t.Fatalf("expected partner or MSP boundary to avoid canonical authority, got %q", got)
	}
	view = ReferenceArchitectureValDSupportBoundaryCollection().Views[0]
	view.RedactionKeepsUnsupportedVisible = false
	if got := EvaluateReferenceArchitectureValDSupportBoundaryViewState(view); got == ReferenceArchitectureValDSupportBoundaryStateActive {
		t.Fatalf("expected redaction hiding unsupported support conditions to fail closed, got %q", got)
	}
}

func TestReferenceArchitectureValDMigrationUpgradeValidation(t *testing.T) {
	view := ReferenceArchitectureValDMigrationUpgradeCollection().Views[0]
	if got := EvaluateReferenceArchitectureValDMigrationVisibilityState(view); got != ReferenceArchitectureValDMigrationStateActive {
		t.Fatalf("expected active migration visibility state, got %q", got)
	}
	view = ReferenceArchitectureValDMigrationUpgradeCollection().Views[0]
	view.DeprecationState = ReferenceArchitectureLifecycleSuperseded
	view.TargetBlueprintVersion = "v2"
	if got := EvaluateReferenceArchitectureValDMigrationVisibilityState(view); got == ReferenceArchitectureValDMigrationStateActive {
		t.Fatalf("expected superseded blueprint to avoid clean matched migration state, got %q", got)
	}
	view = ReferenceArchitectureValDMigrationUpgradeCollection().Views[0]
	view.MigrationPathRef = ""
	if got := EvaluateReferenceArchitectureValDMigrationVisibilityState(view); got == ReferenceArchitectureValDMigrationStateActive {
		t.Fatalf("expected missing migration path to fail closed, got %q", got)
	}
	view = ReferenceArchitectureValDMigrationUpgradeCollection().Views[0]
	view.RollbackBoundaryRef = ""
	if got := EvaluateReferenceArchitectureValDMigrationVisibilityState(view); got == ReferenceArchitectureValDMigrationStateActive {
		t.Fatalf("expected missing rollback boundary to fail closed, got %q", got)
	}
	view = ReferenceArchitectureValDMigrationUpgradeCollection().Views[0]
	view.EvidenceRefs[0].FreshnessState = IntelligenceCalibrationFreshnessStale
	if got := EvaluateReferenceArchitectureValDMigrationVisibilityState(view); got == ReferenceArchitectureValDMigrationStateActive {
		t.Fatalf("expected stale migration evidence to prevent active state, got %q", got)
	}
}

func TestReferenceArchitectureValDTopologyGateValidation(t *testing.T) {
	check := ReferenceArchitectureValDTopologyGateCollection().Checks[0]
	if got := EvaluateReferenceArchitectureValDTopologyGateState(check); got != ReferenceArchitectureValDTopologyGateStateActive {
		t.Fatalf("expected active topology gate state, got %q", got)
	}
	check = ReferenceArchitectureValDTopologyGateCollection().Checks[0]
	check.SupportedTopology = false
	if got := EvaluateReferenceArchitectureValDTopologyGateState(check); got == ReferenceArchitectureValDTopologyGateStateActive {
		t.Fatalf("expected unsupported topology to block clean alignment, got %q", got)
	}
	check = ReferenceArchitectureValDTopologyGateCollection().Checks[0]
	check.TrustAnchorMode = "shared-anchor"
	if got := EvaluateReferenceArchitectureValDTopologyGateState(check); got == ReferenceArchitectureValDTopologyGateStateActive {
		t.Fatalf("expected unknown trust anchor topology to fail closed, got %q", got)
	}
	check = ReferenceArchitectureValDTopologyGateCollection().Checks[0]
	check.RedactionKeepsMismatchVisible = false
	if got := EvaluateReferenceArchitectureValDTopologyGateState(check); got == ReferenceArchitectureValDTopologyGateStateActive {
		t.Fatalf("expected topology mismatch not to be redacted into matched, got %q", got)
	}
}

func TestReferenceArchitectureValDSecurityBoundaryValidation(t *testing.T) {
	check := ReferenceArchitectureValDSecurityBoundaryCollection().Checks[0]
	if got := EvaluateReferenceArchitectureValDSecurityBoundaryGateState(check); got != ReferenceArchitectureValDSecurityGateStateActive {
		t.Fatalf("expected active security boundary gate state, got %q", got)
	}
	check = ReferenceArchitectureValDSecurityBoundaryCollection().Checks[0]
	check.TrustAnchorBoundary = ""
	if got := EvaluateReferenceArchitectureValDSecurityBoundaryGateState(check); got == ReferenceArchitectureValDSecurityGateStateActive {
		t.Fatalf("expected missing trust or custody boundary to block active state, got %q", got)
	}
	check = ReferenceArchitectureValDSecurityBoundaryCollection().Checks[0]
	check.RedactionExportBoundary = ""
	if got := EvaluateReferenceArchitectureValDSecurityBoundaryGateState(check); got == ReferenceArchitectureValDSecurityGateStateActive {
		t.Fatalf("expected missing redaction or export boundary to block active state, got %q", got)
	}
	check = ReferenceArchitectureValDSecurityBoundaryCollection().Checks[0]
	check.ApprovalAuthorityBlocked = false
	if got := EvaluateReferenceArchitectureValDSecurityBoundaryGateState(check); got != ReferenceArchitectureValDSecurityGateStateBlocked {
		t.Fatalf("expected partner or verifier authority escalation to block active state, got %q", got)
	}
	check = ReferenceArchitectureValDSecurityBoundaryCollection().Checks[0]
	check.NoShadowTruthBoundary = ""
	if got := EvaluateReferenceArchitectureValDSecurityBoundaryGateState(check); got == ReferenceArchitectureValDSecurityGateStateActive {
		t.Fatalf("expected explicit no-shadow-truth rule, got %q", got)
	}
}

func TestReferenceArchitectureValDOperabilityGateValidation(t *testing.T) {
	check := ReferenceArchitectureValDOperabilityGateCollection().Checks[0]
	if got := EvaluateReferenceArchitectureValDOperabilityGateState(check); got != ReferenceArchitectureValDOperabilityGateStateActive {
		t.Fatalf("expected active operability gate state, got %q", got)
	}
	check = ReferenceArchitectureValDOperabilityGateCollection().Checks[0]
	check.ReadinessState = ReferenceArchitectureValBReadinessStatePartial
	if got := EvaluateReferenceArchitectureValDOperabilityGateState(check); got == ReferenceArchitectureValDOperabilityGateStateActive {
		t.Fatalf("expected missing readiness link to block active state, got %q", got)
	}
	check = ReferenceArchitectureValDOperabilityGateCollection().Checks[0]
	check.RecoveryExpectationState = ReferenceArchitectureValCRecoveryExpectationStatePartial
	if got := EvaluateReferenceArchitectureValDOperabilityGateState(check); got == ReferenceArchitectureValDOperabilityGateStateActive {
		t.Fatalf("expected missing recovery expectation to block active state, got %q", got)
	}
	check = ReferenceArchitectureValDOperabilityGateCollection().Checks[0]
	check.OperatorActionGuidance = ""
	if got := EvaluateReferenceArchitectureValDOperabilityGateState(check); got == ReferenceArchitectureValDOperabilityGateStateActive {
		t.Fatalf("expected missing operator action guidance to block active state, got %q", got)
	}
	check = ReferenceArchitectureValDOperabilityGateCollection().Checks[0]
	check.EvidenceRefs[0].FreshnessState = IntelligenceCalibrationFreshnessStale
	if got := EvaluateReferenceArchitectureValDOperabilityGateState(check); got == ReferenceArchitectureValDOperabilityGateStateActive {
		t.Fatalf("expected stale evidence to block active state, got %q", got)
	}
}

func TestReferenceArchitectureValDCompatibilityGateValidation(t *testing.T) {
	check := ReferenceArchitectureValDCompatibilityGateCollection().Checks[0]
	if got := EvaluateReferenceArchitectureValDCompatibilityGateState(check); got != ReferenceArchitectureValDCompatibilityGateStateActive {
		t.Fatalf("expected active compatibility gate state, got %q", got)
	}
	check = ReferenceArchitectureValDCompatibilityGateCollection().Checks[0]
	check.CompatibilityState = ReferenceArchitectureCompatibilityUnsupported
	if got := EvaluateReferenceArchitectureValDCompatibilityGateState(check); got != ReferenceArchitectureValDCompatibilityGateStateBlocked {
		t.Fatalf("expected unsupported compatibility to return blocked or unsupported state, got %q", got)
	}
	check = ReferenceArchitectureValDCompatibilityGateCollection().Checks[0]
	check.LifecycleState = ReferenceArchitectureLifecycleSuperseded
	if got := EvaluateReferenceArchitectureValDCompatibilityGateState(check); got == ReferenceArchitectureValDCompatibilityGateStateActive {
		t.Fatalf("expected superseded lifecycle without migration closure to avoid clean active state, got %q", got)
	}
	check = ReferenceArchitectureValDCompatibilityGateCollection().Checks[0]
	check.CompatibilityState = "universally_supported"
	if got := EvaluateReferenceArchitectureValDCompatibilityGateState(check); got == ReferenceArchitectureValDCompatibilityGateStateActive {
		t.Fatalf("expected unknown compatibility to fail closed, got %q", got)
	}
	check = ReferenceArchitectureValDCompatibilityGateCollection().Checks[0]
	check.UniversalSupportClaim = true
	if got := EvaluateReferenceArchitectureValDCompatibilityGateState(check); got != ReferenceArchitectureValDCompatibilityGateStateBlocked {
		t.Fatalf("expected universal support overclaim to block active state, got %q", got)
	}
}

func TestReferenceArchitectureValDFinalGateReportRequiresExactComponentStates(t *testing.T) {
	report := ReferenceArchitectureValDFinalGateCollection().Reports[0]
	if got := EvaluateReferenceArchitectureValDFinalGateReportState(report); got != ReferenceArchitectureValDFinalGateStateActive {
		t.Fatalf("expected exact active component states to produce active final gate, got %q", got)
	}

	report = ReferenceArchitectureValDFinalGateCollection().Reports[0]
	report.OperationalVisibilityState = "not_a_real_state_active"
	if got := EvaluateReferenceArchitectureValDFinalGateReportState(report); got == ReferenceArchitectureValDFinalGateStateActive {
		t.Fatalf("expected fake active-like visibility state to fail closed, got %q", got)
	}

	report = ReferenceArchitectureValDFinalGateCollection().Reports[0]
	report.AlignmentSummaryState = "typo_alignment_active"
	if got := EvaluateReferenceArchitectureValDFinalGateReportState(report); got == ReferenceArchitectureValDFinalGateStateActive {
		t.Fatalf("expected typo active-like alignment state to fail closed, got %q", got)
	}

	report = ReferenceArchitectureValDFinalGateCollection().Reports[0]
	report.SecurityBoundaryGateState = ReferenceArchitectureValDAlignmentStateActive
	if got := EvaluateReferenceArchitectureValDFinalGateReportState(report); got == ReferenceArchitectureValDFinalGateStateActive {
		t.Fatalf("expected wrong-enum active-like state to fail closed, got %q", got)
	}

	report = ReferenceArchitectureValDFinalGateCollection().Reports[0]
	report.CompatibilityGateState = ReferenceArchitectureValDCompatibilityGateStatePartial
	if got := EvaluateReferenceArchitectureValDFinalGateReportState(report); got == ReferenceArchitectureValDFinalGateStateActive {
		t.Fatalf("expected partial component state to prevent final_gate_active, got %q", got)
	}

	report = ReferenceArchitectureValDFinalGateCollection().Reports[0]
	report.OperabilityGateState = ReferenceArchitectureValDOperabilityGateStateUnknown
	if got := EvaluateReferenceArchitectureValDFinalGateReportState(report); got == ReferenceArchitectureValDFinalGateStateActive {
		t.Fatalf("expected unknown component state to prevent final_gate_active, got %q", got)
	}
}

func TestReferenceArchitectureValDCollectionsNormalizeFamilyDuplicates(t *testing.T) {
	visibility := ReferenceArchitectureValDOperationalVisibilityCollection()
	visibility.Reports[1].BlueprintFamily = " " + visibility.Reports[0].BlueprintFamily + " "
	if got := EvaluateReferenceArchitectureValDOperationalVisibilityCollectionState(visibility); got == ReferenceArchitectureValDVisibilityStateActive {
		t.Fatalf("expected whitespace-variant duplicate family to fail closed in visibility collection, got %q", got)
	}

	visibility = ReferenceArchitectureValDOperationalVisibilityCollection()
	visibility.Reports[0].BlueprintFamily = "   "
	if got := EvaluateReferenceArchitectureValDOperationalVisibilityCollectionState(visibility); got == ReferenceArchitectureValDVisibilityStateActive {
		t.Fatalf("expected empty or whitespace-only family to fail closed in visibility collection, got %q", got)
	}

	finalGate := ReferenceArchitectureValDFinalGateCollection()
	finalGate.Reports[1].BlueprintFamily = " " + finalGate.Reports[0].BlueprintFamily + " "
	if got := EvaluateReferenceArchitectureValDFinalGateCollectionState(finalGate); got == ReferenceArchitectureValDFinalGateStateActive {
		t.Fatalf("expected whitespace-variant duplicate family to fail closed in final gate collection, got %q", got)
	}

	finalGate = ReferenceArchitectureValDFinalGateCollection()
	finalGate.Reports[0].BlueprintFamily = "   "
	if got := EvaluateReferenceArchitectureValDFinalGateCollectionState(finalGate); got == ReferenceArchitectureValDFinalGateStateActive {
		t.Fatalf("expected empty or whitespace-only family to fail closed in final gate collection, got %q", got)
	}
}

func TestReferenceArchitectureValDProofSurfaceCompleteness(t *testing.T) {
	visibility := ReferenceArchitectureValDOperationalVisibilityCollection()
	supportedFamilies := visibility.SupportedFamilies
	surfaceRefs := referenceArchitectureValDProofSurfaceRefs()
	evidenceRefs := []string{"point5", "val0", "vala", "valb", "valc", "visibility", "alignment", "alerts", "support", "migration", "topology", "security", "operability", "compatibility"}
	limitations := []string{"Val D keeps point 6 not complete."}
	disclaimer := referenceArchitectureValDProjectionDisclaimer()

	if got := EvaluateReferenceArchitectureValDProofsState(ReferenceArchitectureValDStateActive, IntelligenceCalibrationValEStateActive, ReferenceArchitecturePoint6StateNotComplete, supportedFamilies, surfaceRefs, evidenceRefs, limitations, disclaimer); got != ReferenceArchitectureValDStateActive {
		t.Fatalf("expected active Val D proofs state with exact surface set, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValDProofsState(ReferenceArchitectureValDStateActive, IntelligenceCalibrationValEStateSubstantial, ReferenceArchitecturePoint6StateNotComplete, supportedFamilies, surfaceRefs, evidenceRefs, limitations, disclaimer); got == ReferenceArchitectureValDStateActive {
		t.Fatalf("expected point 5 dependency health regression to block active proof surface, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValDProofsState(ReferenceArchitectureValDStateActive, IntelligenceCalibrationValEStateActive, ReferenceArchitecturePoint6StateNotComplete, supportedFamilies, surfaceRefs[1:], evidenceRefs, limitations, disclaimer); got == ReferenceArchitectureValDStateActive {
		t.Fatalf("expected missing required surface to fail closed, got %q", got)
	}
	duplicateSurfaceRefs := append([]string{}, surfaceRefs[1:]...)
	duplicateSurfaceRefs = append(duplicateSurfaceRefs, surfaceRefs[0])
	duplicateSurfaceRefs[0] = surfaceRefs[len(surfaceRefs)-1]
	if got := EvaluateReferenceArchitectureValDProofsState(ReferenceArchitectureValDStateActive, IntelligenceCalibrationValEStateActive, ReferenceArchitecturePoint6StateNotComplete, supportedFamilies, duplicateSurfaceRefs, evidenceRefs, limitations, disclaimer); got == ReferenceArchitectureValDStateActive {
		t.Fatalf("expected duplicate surface not to compensate for a missing required surface, got %q", got)
	}
	extraSurfaceRefs := append([]string{}, surfaceRefs...)
	extraSurfaceRefs[len(extraSurfaceRefs)-1] = "/v1/reference-architecture/vald/unknown"
	if got := EvaluateReferenceArchitectureValDProofsState(ReferenceArchitectureValDStateActive, IntelligenceCalibrationValEStateActive, ReferenceArchitecturePoint6StateNotComplete, supportedFamilies, extraSurfaceRefs, evidenceRefs, limitations, disclaimer); got == ReferenceArchitectureValDStateActive {
		t.Fatalf("expected unknown extra surface not to compensate for a missing required surface, got %q", got)
	}
	whitespaceSurfaceRefs := append([]string{}, surfaceRefs...)
	whitespaceSurfaceRefs[0] = "   "
	if got := EvaluateReferenceArchitectureValDProofsState(ReferenceArchitectureValDStateActive, IntelligenceCalibrationValEStateActive, ReferenceArchitecturePoint6StateNotComplete, supportedFamilies, whitespaceSurfaceRefs, evidenceRefs, limitations, disclaimer); got == ReferenceArchitectureValDStateActive {
		t.Fatalf("expected empty or whitespace surface ref to fail closed, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValDProofsState(ReferenceArchitectureValDStateActive, IntelligenceCalibrationValEStateActive, ReferenceArchitecturePoint6StatePass, supportedFamilies, surfaceRefs, evidenceRefs, limitations, disclaimer); got == ReferenceArchitectureValDStateActive {
		t.Fatalf("expected point_6_pass to remain impossible in Val D, got %q", got)
	}
}
