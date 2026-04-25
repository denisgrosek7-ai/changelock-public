package operability

import "testing"

func TestProductionUsabilityValDStateRequiresActiveVal0(t *testing.T) {
	if got := EvaluateProductionUsabilityValDState(
		ProductionUsabilityVal0StateSubstantial,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateActive,
		ProductionUsabilityValDConfigReviewStateActive,
		ProductionUsabilityValDExplainabilityReviewStateActive,
		ProductionUsabilityValDDryRunReviewStateActive,
		ProductionUsabilityValDRedactionReviewStateActive,
		ProductionUsabilityValDDegradedBehaviorReviewStateActive,
		ProductionUsabilityValDUIWindowingReviewStateActive,
		ProductionUsabilityValDCommandNoiseReviewStateActive,
		ProductionUsabilityValDAPIProtectionReviewStateActive,
		ProductionUsabilityValDCLIResilienceReviewStateActive,
		ProductionUsabilityValDSupportabilityReviewStateActive,
		ProductionUsabilityValDRecoveryReviewStateActive,
		ProductionUsabilityValDUpgradeRollbackReviewStateActive,
		ProductionUsabilityValDScaleEnvelopeReviewStateActive,
		ProductionUsabilityValDGovernanceBoundaryReviewStateActive,
		ProductionUsabilityValDRegressionGateStateActive,
	); got == ProductionUsabilityValDStateActive {
		t.Fatalf("expected non-active Val D state without active Val 0, got %q", got)
	}
}

func TestProductionUsabilityValDStateRequiresActiveValA(t *testing.T) {
	if got := EvaluateProductionUsabilityValDState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateSubstantial,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateActive,
		ProductionUsabilityValDConfigReviewStateActive,
		ProductionUsabilityValDExplainabilityReviewStateActive,
		ProductionUsabilityValDDryRunReviewStateActive,
		ProductionUsabilityValDRedactionReviewStateActive,
		ProductionUsabilityValDDegradedBehaviorReviewStateActive,
		ProductionUsabilityValDUIWindowingReviewStateActive,
		ProductionUsabilityValDCommandNoiseReviewStateActive,
		ProductionUsabilityValDAPIProtectionReviewStateActive,
		ProductionUsabilityValDCLIResilienceReviewStateActive,
		ProductionUsabilityValDSupportabilityReviewStateActive,
		ProductionUsabilityValDRecoveryReviewStateActive,
		ProductionUsabilityValDUpgradeRollbackReviewStateActive,
		ProductionUsabilityValDScaleEnvelopeReviewStateActive,
		ProductionUsabilityValDGovernanceBoundaryReviewStateActive,
		ProductionUsabilityValDRegressionGateStateActive,
	); got == ProductionUsabilityValDStateActive {
		t.Fatalf("expected non-active Val D state without active Val A, got %q", got)
	}
}

func TestProductionUsabilityValDStateRequiresActiveValB(t *testing.T) {
	if got := EvaluateProductionUsabilityValDState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateSubstantial,
		ProductionUsabilityValCStateActive,
		ProductionUsabilityValDConfigReviewStateActive,
		ProductionUsabilityValDExplainabilityReviewStateActive,
		ProductionUsabilityValDDryRunReviewStateActive,
		ProductionUsabilityValDRedactionReviewStateActive,
		ProductionUsabilityValDDegradedBehaviorReviewStateActive,
		ProductionUsabilityValDUIWindowingReviewStateActive,
		ProductionUsabilityValDCommandNoiseReviewStateActive,
		ProductionUsabilityValDAPIProtectionReviewStateActive,
		ProductionUsabilityValDCLIResilienceReviewStateActive,
		ProductionUsabilityValDSupportabilityReviewStateActive,
		ProductionUsabilityValDRecoveryReviewStateActive,
		ProductionUsabilityValDUpgradeRollbackReviewStateActive,
		ProductionUsabilityValDScaleEnvelopeReviewStateActive,
		ProductionUsabilityValDGovernanceBoundaryReviewStateActive,
		ProductionUsabilityValDRegressionGateStateActive,
	); got == ProductionUsabilityValDStateActive {
		t.Fatalf("expected non-active Val D state without active Val B, got %q", got)
	}
}

func TestProductionUsabilityValDStateRequiresActiveValC(t *testing.T) {
	if got := EvaluateProductionUsabilityValDState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateSubstantial,
		ProductionUsabilityValDConfigReviewStateActive,
		ProductionUsabilityValDExplainabilityReviewStateActive,
		ProductionUsabilityValDDryRunReviewStateActive,
		ProductionUsabilityValDRedactionReviewStateActive,
		ProductionUsabilityValDDegradedBehaviorReviewStateActive,
		ProductionUsabilityValDUIWindowingReviewStateActive,
		ProductionUsabilityValDCommandNoiseReviewStateActive,
		ProductionUsabilityValDAPIProtectionReviewStateActive,
		ProductionUsabilityValDCLIResilienceReviewStateActive,
		ProductionUsabilityValDSupportabilityReviewStateActive,
		ProductionUsabilityValDRecoveryReviewStateActive,
		ProductionUsabilityValDUpgradeRollbackReviewStateActive,
		ProductionUsabilityValDScaleEnvelopeReviewStateActive,
		ProductionUsabilityValDGovernanceBoundaryReviewStateActive,
		ProductionUsabilityValDRegressionGateStateActive,
	); got == ProductionUsabilityValDStateActive {
		t.Fatalf("expected non-active Val D state without active Val C, got %q", got)
	}
}

func TestProductionUsabilityValDConfigReviewBlocksUnsupportedUnknownSchemaState(t *testing.T) {
	model := ProductionUsabilityValDConfigCorrectnessReview()
	model.SchemaCompatibilityStatus = ProductionUsabilityCompatibilityUnknown
	if got := EvaluateProductionUsabilityValDConfigReviewState(model); got == ProductionUsabilityValDConfigReviewStateActive {
		t.Fatalf("expected non-active config review state for unknown schema status, got %q", got)
	}
}

func TestProductionUsabilityValDConfigReviewBlocksCanonicalTruthClaim(t *testing.T) {
	model := ProductionUsabilityValDConfigCorrectnessReview()
	model.EffectiveConfigClaimsCanonical = true
	if got := EvaluateProductionUsabilityValDConfigReviewState(model); got == ProductionUsabilityValDConfigReviewStateActive {
		t.Fatalf("expected non-active config review state when effective config claims canonical truth, got %q", got)
	}
}

func TestProductionUsabilityValDConfigReviewBlocksSecretsExposure(t *testing.T) {
	model := ProductionUsabilityValDConfigCorrectnessReview()
	model.SecretsExposedInEffectiveConfig = true
	if got := EvaluateProductionUsabilityValDConfigReviewState(model); got == ProductionUsabilityValDConfigReviewStateActive {
		t.Fatalf("expected non-active config review state when secrets are exposed, got %q", got)
	}
}

func TestProductionUsabilityValDExplainabilityReviewBlocksVagueExplanation(t *testing.T) {
	model := ProductionUsabilityValDExplainabilityClarityReview()
	model.ReasonCodesPresent = false
	if got := EvaluateProductionUsabilityValDExplainabilityReviewState(model); got == ProductionUsabilityValDExplainabilityReviewStateActive {
		t.Fatalf("expected non-active explainability review when reason code is missing, got %q", got)
	}
}

func TestProductionUsabilityValDRedactionReviewBlocksPartnerPublicFullEvidence(t *testing.T) {
	model := ProductionUsabilityValDPermissionRedactionReview()
	model.PartnerOrPublicExposeFullEvidence = true
	if got := EvaluateProductionUsabilityValDRedactionReviewState(model); got == ProductionUsabilityValDRedactionReviewStateActive {
		t.Fatalf("expected non-active redaction review when partner/public exposes full evidence, got %q", got)
	}
}

func TestProductionUsabilityValDRedactionReviewPreservesHiddenMetadataRepresentation(t *testing.T) {
	model := ProductionUsabilityValDPermissionRedactionReview()
	model.HiddenMetadataRepresented = false
	if got := EvaluateProductionUsabilityValDRedactionReviewState(model); got == ProductionUsabilityValDRedactionReviewStateActive {
		t.Fatalf("expected non-active redaction review when hidden metadata is omitted, got %q", got)
	}
}

func TestProductionUsabilityValDDryRunReviewBlocksMutation(t *testing.T) {
	model := ProductionUsabilityValDDryRunAuditReview()
	model.DryRunMutatesCanonicalState = true
	if got := EvaluateProductionUsabilityValDDryRunReviewState(model); got == ProductionUsabilityValDDryRunReviewStateActive {
		t.Fatalf("expected non-active dry-run review when dry-run mutates, got %q", got)
	}
}

func TestProductionUsabilityValDDryRunSuccessCannotEqualActivation(t *testing.T) {
	model := ProductionUsabilityValDDryRunAuditReview()
	model.PreviewSuccessImpliesActivate = true
	if got := EvaluateProductionUsabilityValDDryRunReviewState(model); got == ProductionUsabilityValDDryRunReviewStateActive {
		t.Fatalf("expected non-active dry-run review when preview implies activation, got %q", got)
	}
}

func TestProductionUsabilityValDAuditOnlyCannotEqualApproval(t *testing.T) {
	model := ProductionUsabilityValDDryRunAuditReview()
	model.AuditOnlyImpliesApproval = true
	if got := EvaluateProductionUsabilityValDDryRunReviewState(model); got == ProductionUsabilityValDDryRunReviewStateActive {
		t.Fatalf("expected non-active dry-run review when audit-only implies approval, got %q", got)
	}
}

func TestProductionUsabilityValDDegradedReviewBlocksStaleAsFresh(t *testing.T) {
	model := ProductionUsabilityValDDegradedBehaviorReview()
	model.StaleReportedAsFresh = true
	if got := EvaluateProductionUsabilityValDDegradedBehaviorReviewState(model); got == ProductionUsabilityValDDegradedBehaviorReviewStateActive {
		t.Fatalf("expected non-active degraded review when stale is reported as fresh, got %q", got)
	}
}

func TestProductionUsabilityValDDegradedReviewBlocksPartialAsComplete(t *testing.T) {
	model := ProductionUsabilityValDDegradedBehaviorReview()
	model.PartialReportedAsComplete = true
	if got := EvaluateProductionUsabilityValDDegradedBehaviorReviewState(model); got == ProductionUsabilityValDDegradedBehaviorReviewStateActive {
		t.Fatalf("expected non-active degraded review when partial is reported as complete, got %q", got)
	}
}

func TestProductionUsabilityValDDegradedReviewBlocksDegradedAsHealthy(t *testing.T) {
	model := ProductionUsabilityValDDegradedBehaviorReview()
	model.DegradedReportedAsHealthy = true
	if got := EvaluateProductionUsabilityValDDegradedBehaviorReviewState(model); got == ProductionUsabilityValDDegradedBehaviorReviewStateActive {
		t.Fatalf("expected non-active degraded review when degraded is reported as healthy, got %q", got)
	}
}

func TestProductionUsabilityValDUIWindowingReviewBlocksUnknownTotalClaimingCompleteness(t *testing.T) {
	model := ProductionUsabilityValDUIWindowingResultReview()
	model.UnknownTotalClaimsComplete = true
	if got := EvaluateProductionUsabilityValDUIWindowingReviewState(model); got == ProductionUsabilityValDUIWindowingReviewStateActive {
		t.Fatalf("expected non-active UI/windowing review when unknown total claims completeness, got %q", got)
	}
}

func TestProductionUsabilityValDUIWindowingReviewBlocksLimitGreaterThanMax(t *testing.T) {
	model := ProductionUsabilityValDUIWindowingResultReview()
	model.LimitExceedsMaxWindow = true
	if got := EvaluateProductionUsabilityValDUIWindowingReviewState(model); got == ProductionUsabilityValDUIWindowingReviewStateActive {
		t.Fatalf("expected non-active UI/windowing review when limit exceeds max, got %q", got)
	}
}

func TestProductionUsabilityValDCommandNoiseReviewBlocksUngovernedTaskMutation(t *testing.T) {
	model := ProductionUsabilityValDCommandNoiseReview()
	model.UngovernedTaskMutation = true
	if got := EvaluateProductionUsabilityValDCommandNoiseReviewState(model); got == ProductionUsabilityValDCommandNoiseReviewStateActive {
		t.Fatalf("expected non-active command/noise review when task mutation is unguided, got %q", got)
	}
}

func TestProductionUsabilityValDCommandNoiseReviewBlocksInvisibleCriticalSuppression(t *testing.T) {
	model := ProductionUsabilityValDCommandNoiseReview()
	model.CriticalSuppressionInvisible = true
	if got := EvaluateProductionUsabilityValDCommandNoiseReviewState(model); got == ProductionUsabilityValDCommandNoiseReviewStateActive {
		t.Fatalf("expected non-active command/noise review when critical suppression is invisible, got %q", got)
	}
}

func TestProductionUsabilityValDAPIProtectionReviewBlocksPriorityBypassOfGovernance(t *testing.T) {
	model := ProductionUsabilityValDAPIProtectionReview()
	model.PriorityBypassesGovernance = true
	if got := EvaluateProductionUsabilityValDAPIProtectionReviewState(model); got == ProductionUsabilityValDAPIProtectionReviewStateActive {
		t.Fatalf("expected non-active API protection review when priority bypasses governance, got %q", got)
	}
}

func TestProductionUsabilityValDAPIProtectionReviewBlocksPolicyDenialHiddenAsThrottling(t *testing.T) {
	model := ProductionUsabilityValDAPIProtectionReview()
	model.PolicyDenialHiddenThrottle = true
	if got := EvaluateProductionUsabilityValDAPIProtectionReviewState(model); got == ProductionUsabilityValDAPIProtectionReviewStateActive {
		t.Fatalf("expected non-active API protection review when policy denial is hidden as throttling, got %q", got)
	}
}

func TestProductionUsabilityValDCLIReviewBlocksRetryUnsafeCommandWithoutReason(t *testing.T) {
	model := ProductionUsabilityValDCLIResilienceReview()
	model.RetryUnsafeMissingReason = true
	if got := EvaluateProductionUsabilityValDCLIResilienceReviewState(model); got == ProductionUsabilityValDCLIResilienceReviewStateActive {
		t.Fatalf("expected non-active CLI review when retry-unsafe reason is missing, got %q", got)
	}
}

func TestProductionUsabilityValDCLIReviewBlocksMissingRequiredIdempotencyKey(t *testing.T) {
	model := ProductionUsabilityValDCLIResilienceReview()
	model.MissingRequiredKey = true
	if got := EvaluateProductionUsabilityValDCLIResilienceReviewState(model); got == ProductionUsabilityValDCLIResilienceReviewStateActive {
		t.Fatalf("expected non-active CLI review when required idempotency key is missing, got %q", got)
	}
}

func TestProductionUsabilityValDSupportabilityReviewBlocksSupportBundleWithoutManifest(t *testing.T) {
	model := ProductionUsabilityValDSupportabilityReview()
	model.SupportBundleManifestMissing = true
	if got := EvaluateProductionUsabilityValDSupportabilityReviewState(model); got == ProductionUsabilityValDSupportabilityReviewStateActive {
		t.Fatalf("expected non-active supportability review when support bundle manifest is missing, got %q", got)
	}
}

func TestProductionUsabilityValDSupportabilityReviewBlocksRawSecretsTokens(t *testing.T) {
	model := ProductionUsabilityValDSupportabilityReview()
	model.RawSecretsOrTokensDetected = true
	if got := EvaluateProductionUsabilityValDSupportabilityReviewState(model); got == ProductionUsabilityValDSupportabilityReviewStateActive {
		t.Fatalf("expected non-active supportability review when raw secrets or tokens are detected, got %q", got)
	}
}

func TestProductionUsabilityValDSupportabilityReviewBlocksHealthyWithBlockingDegraded(t *testing.T) {
	model := ProductionUsabilityValDSupportabilityReview()
	model.HealthyWithBlockingDegraded = true
	if got := EvaluateProductionUsabilityValDSupportabilityReviewState(model); got == ProductionUsabilityValDSupportabilityReviewStateActive {
		t.Fatalf("expected non-active supportability review when healthy coexists with blocking degraded state, got %q", got)
	}
}

func TestProductionUsabilityValDRecoveryReviewBlocksPolicyBypassRecommendation(t *testing.T) {
	model := ProductionUsabilityValDRecoveryUXReview()
	model.PolicyBypassRecommended = true
	if got := EvaluateProductionUsabilityValDRecoveryReviewState(model); got == ProductionUsabilityValDRecoveryReviewStateActive {
		t.Fatalf("expected non-active recovery review when policy bypass is recommended, got %q", got)
	}
}

func TestProductionUsabilityValDRecoveryReviewBlocksUnsafeRetryRecommendation(t *testing.T) {
	model := ProductionUsabilityValDRecoveryUXReview()
	model.UnsafeRetryRecommended = true
	if got := EvaluateProductionUsabilityValDRecoveryReviewState(model); got == ProductionUsabilityValDRecoveryReviewStateActive {
		t.Fatalf("expected non-active recovery review when unsafe retry is recommended, got %q", got)
	}
}

func TestProductionUsabilityValDUpgradeRollbackReviewBlocksUnknownTargetVersion(t *testing.T) {
	model := ProductionUsabilityValDUpgradeRollbackReview()
	model.TargetVersionKnown = false
	if got := EvaluateProductionUsabilityValDUpgradeRollbackReviewState(model); got == ProductionUsabilityValDUpgradeRollbackReviewStateActive {
		t.Fatalf("expected non-active upgrade review when target version is unknown, got %q", got)
	}
}

func TestProductionUsabilityValDUpgradeRollbackReviewBlocksMutatingAdvisory(t *testing.T) {
	model := ProductionUsabilityValDUpgradeRollbackReview()
	model.AdvisoryMutatesState = true
	if got := EvaluateProductionUsabilityValDUpgradeRollbackReviewState(model); got == ProductionUsabilityValDUpgradeRollbackReviewStateActive {
		t.Fatalf("expected non-active upgrade review when advisory mutates state, got %q", got)
	}
}

func TestProductionUsabilityValDScaleEnvelopeReviewMarksUnknownUnmeasuredAsLimitation(t *testing.T) {
	model := ProductionUsabilityValDScaleEnvelopeReview()
	model.UnknownOrUnmeasuredMarkedLimit = false
	if got := EvaluateProductionUsabilityValDScaleEnvelopeReviewState(model); got == ProductionUsabilityValDScaleEnvelopeReviewStateActive {
		t.Fatalf("expected non-active scale review when unknown/unmeasured scale is not marked as limitation, got %q", got)
	}
}

func TestProductionUsabilityValDGovernanceBoundaryReviewBlocksProjectionClaimingCanonicalTruth(t *testing.T) {
	model := ProductionUsabilityValDGovernanceBoundaryReview()
	model.ProjectionClaimsCanonicalTruth = true
	if got := EvaluateProductionUsabilityValDGovernanceBoundaryReviewState(model); got == ProductionUsabilityValDGovernanceBoundaryReviewStateActive {
		t.Fatalf("expected non-active governance review when projection claims canonical truth, got %q", got)
	}
}

func TestProductionUsabilityValDGovernanceBoundaryReviewBlocksAdvisoryMutationWithoutGovernance(t *testing.T) {
	model := ProductionUsabilityValDGovernanceBoundaryReview()
	model.AdvisoryMutatesWithoutGovernance = true
	if got := EvaluateProductionUsabilityValDGovernanceBoundaryReviewState(model); got == ProductionUsabilityValDGovernanceBoundaryReviewStateActive {
		t.Fatalf("expected non-active governance review when advisory mutates without governance, got %q", got)
	}
}

func TestProductionUsabilityValDRegressionGateBlocksMissingCriticalFixtureCoverage(t *testing.T) {
	model := ProductionUsabilityValDUsabilityRegressionGate()
	model.MissingCriticalFixtureCoverage = true
	if got := EvaluateProductionUsabilityValDRegressionGateState(model); got == ProductionUsabilityValDRegressionGateStateActive {
		t.Fatalf("expected non-active regression gate when critical fixture coverage is missing, got %q", got)
	}
}

func TestProductionUsabilityValDProofsCanBecomeActiveOnlyAsFinalGateWhilePoint4RemainsNotComplete(t *testing.T) {
	configReview := ProductionUsabilityValDConfigCorrectnessReview()
	explainReview := ProductionUsabilityValDExplainabilityClarityReview()
	dryRunReview := ProductionUsabilityValDDryRunAuditReview()
	redactionReview := ProductionUsabilityValDPermissionRedactionReview()
	degradedReview := ProductionUsabilityValDDegradedBehaviorReview()
	uiWindowingReview := ProductionUsabilityValDUIWindowingResultReview()
	commandNoiseReview := ProductionUsabilityValDCommandNoiseReview()
	apiReview := ProductionUsabilityValDAPIProtectionReview()
	cliReview := ProductionUsabilityValDCLIResilienceReview()
	supportabilityReview := ProductionUsabilityValDSupportabilityReview()
	recoveryReview := ProductionUsabilityValDRecoveryUXReview()
	upgradeReview := ProductionUsabilityValDUpgradeRollbackReview()
	scaleReview := ProductionUsabilityValDScaleEnvelopeReview()
	governanceReview := ProductionUsabilityValDGovernanceBoundaryReview()
	regressionGate := ProductionUsabilityValDUsabilityRegressionGate()

	if got := EvaluateProductionUsabilityValDProofsState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateActive,
		EvaluateProductionUsabilityValDConfigReviewState(configReview),
		EvaluateProductionUsabilityValDExplainabilityReviewState(explainReview),
		EvaluateProductionUsabilityValDDryRunReviewState(dryRunReview),
		EvaluateProductionUsabilityValDRedactionReviewState(redactionReview),
		EvaluateProductionUsabilityValDDegradedBehaviorReviewState(degradedReview),
		EvaluateProductionUsabilityValDUIWindowingReviewState(uiWindowingReview),
		EvaluateProductionUsabilityValDCommandNoiseReviewState(commandNoiseReview),
		EvaluateProductionUsabilityValDAPIProtectionReviewState(apiReview),
		EvaluateProductionUsabilityValDCLIResilienceReviewState(cliReview),
		EvaluateProductionUsabilityValDSupportabilityReviewState(supportabilityReview),
		EvaluateProductionUsabilityValDRecoveryReviewState(recoveryReview),
		EvaluateProductionUsabilityValDUpgradeRollbackReviewState(upgradeReview),
		EvaluateProductionUsabilityValDScaleEnvelopeReviewState(scaleReview),
		EvaluateProductionUsabilityValDGovernanceBoundaryReviewState(governanceReview),
		EvaluateProductionUsabilityValDRegressionGateState(regressionGate),
		[]string{"val0", "vala", "valb", "valc", "vald/config-review", "vald/explainability-review", "vald/dry-run-review", "vald/redaction-review", "vald/degraded-state-review", "vald/ui-windowing-review", "vald/command-noise-review", "vald/api-protection-review", "vald/cli-resilience-review", "vald/supportability-review", "vald/recovery-review", "vald/upgrade-rollback-review", "vald/scale-envelope-review", "vald/governance-boundary-review", "vald/regression-gate", "vald/proofs"},
		[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
		[]string{"projection_only"},
		[]string{"Val E integrated closure still outstanding"},
	); got != ProductionUsabilityValDStateActive {
		t.Fatalf("expected active Val D proofs state, got %q", got)
	}
}

func TestProductionUsabilityValDMissingRequiredComponentKeepsValDInactive(t *testing.T) {
	if got := EvaluateProductionUsabilityValDState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateActive,
		ProductionUsabilityValDConfigReviewStatePartial,
		ProductionUsabilityValDExplainabilityReviewStateActive,
		ProductionUsabilityValDDryRunReviewStateActive,
		ProductionUsabilityValDRedactionReviewStateActive,
		ProductionUsabilityValDDegradedBehaviorReviewStateActive,
		ProductionUsabilityValDUIWindowingReviewStateActive,
		ProductionUsabilityValDCommandNoiseReviewStateActive,
		ProductionUsabilityValDAPIProtectionReviewStateActive,
		ProductionUsabilityValDCLIResilienceReviewStateActive,
		ProductionUsabilityValDSupportabilityReviewStateActive,
		ProductionUsabilityValDRecoveryReviewStateActive,
		ProductionUsabilityValDUpgradeRollbackReviewStateActive,
		ProductionUsabilityValDScaleEnvelopeReviewStateActive,
		ProductionUsabilityValDGovernanceBoundaryReviewStateActive,
		ProductionUsabilityValDRegressionGateStateActive,
	); got == ProductionUsabilityValDStateActive {
		t.Fatalf("expected non-active Val D state when a required component is partial, got %q", got)
	}
}

func TestProductionUsabilityValDFoundationIsActive(t *testing.T) {
	configReview := ProductionUsabilityValDConfigCorrectnessReview()
	if got := EvaluateProductionUsabilityValDConfigReviewState(configReview); got != ProductionUsabilityValDConfigReviewStateActive {
		t.Fatalf("expected active config review state, got %q", got)
	}

	explainReview := ProductionUsabilityValDExplainabilityClarityReview()
	if got := EvaluateProductionUsabilityValDExplainabilityReviewState(explainReview); got != ProductionUsabilityValDExplainabilityReviewStateActive {
		t.Fatalf("expected active explainability review state, got %q", got)
	}

	dryRunReview := ProductionUsabilityValDDryRunAuditReview()
	if got := EvaluateProductionUsabilityValDDryRunReviewState(dryRunReview); got != ProductionUsabilityValDDryRunReviewStateActive {
		t.Fatalf("expected active dry-run review state, got %q", got)
	}

	redactionReview := ProductionUsabilityValDPermissionRedactionReview()
	if got := EvaluateProductionUsabilityValDRedactionReviewState(redactionReview); got != ProductionUsabilityValDRedactionReviewStateActive {
		t.Fatalf("expected active redaction review state, got %q", got)
	}

	degradedReview := ProductionUsabilityValDDegradedBehaviorReview()
	if got := EvaluateProductionUsabilityValDDegradedBehaviorReviewState(degradedReview); got != ProductionUsabilityValDDegradedBehaviorReviewStateActive {
		t.Fatalf("expected active degraded review state, got %q", got)
	}

	uiWindowingReview := ProductionUsabilityValDUIWindowingResultReview()
	if got := EvaluateProductionUsabilityValDUIWindowingReviewState(uiWindowingReview); got != ProductionUsabilityValDUIWindowingReviewStateActive {
		t.Fatalf("expected active UI/windowing review state, got %q", got)
	}

	commandNoiseReview := ProductionUsabilityValDCommandNoiseReview()
	if got := EvaluateProductionUsabilityValDCommandNoiseReviewState(commandNoiseReview); got != ProductionUsabilityValDCommandNoiseReviewStateActive {
		t.Fatalf("expected active command/noise review state, got %q", got)
	}

	apiReview := ProductionUsabilityValDAPIProtectionReview()
	if got := EvaluateProductionUsabilityValDAPIProtectionReviewState(apiReview); got != ProductionUsabilityValDAPIProtectionReviewStateActive {
		t.Fatalf("expected active API protection review state, got %q", got)
	}

	cliReview := ProductionUsabilityValDCLIResilienceReview()
	if got := EvaluateProductionUsabilityValDCLIResilienceReviewState(cliReview); got != ProductionUsabilityValDCLIResilienceReviewStateActive {
		t.Fatalf("expected active CLI review state, got %q", got)
	}

	supportabilityReview := ProductionUsabilityValDSupportabilityReview()
	if got := EvaluateProductionUsabilityValDSupportabilityReviewState(supportabilityReview); got != ProductionUsabilityValDSupportabilityReviewStateActive {
		t.Fatalf("expected active supportability review state, got %q", got)
	}

	recoveryReview := ProductionUsabilityValDRecoveryUXReview()
	if got := EvaluateProductionUsabilityValDRecoveryReviewState(recoveryReview); got != ProductionUsabilityValDRecoveryReviewStateActive {
		t.Fatalf("expected active recovery review state, got %q", got)
	}

	upgradeReview := ProductionUsabilityValDUpgradeRollbackReview()
	if got := EvaluateProductionUsabilityValDUpgradeRollbackReviewState(upgradeReview); got != ProductionUsabilityValDUpgradeRollbackReviewStateActive {
		t.Fatalf("expected active upgrade/rollback review state, got %q", got)
	}

	scaleReview := ProductionUsabilityValDScaleEnvelopeReview()
	if got := EvaluateProductionUsabilityValDScaleEnvelopeReviewState(scaleReview); got != ProductionUsabilityValDScaleEnvelopeReviewStateActive {
		t.Fatalf("expected active scale review state, got %q", got)
	}

	governanceReview := ProductionUsabilityValDGovernanceBoundaryReview()
	if got := EvaluateProductionUsabilityValDGovernanceBoundaryReviewState(governanceReview); got != ProductionUsabilityValDGovernanceBoundaryReviewStateActive {
		t.Fatalf("expected active governance review state, got %q", got)
	}

	regressionGate := ProductionUsabilityValDUsabilityRegressionGate()
	if got := EvaluateProductionUsabilityValDRegressionGateState(regressionGate); got != ProductionUsabilityValDRegressionGateStateActive {
		t.Fatalf("expected active regression gate state, got %q", got)
	}

	if got := EvaluateProductionUsabilityValDState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateActive,
		EvaluateProductionUsabilityValDConfigReviewState(configReview),
		EvaluateProductionUsabilityValDExplainabilityReviewState(explainReview),
		EvaluateProductionUsabilityValDDryRunReviewState(dryRunReview),
		EvaluateProductionUsabilityValDRedactionReviewState(redactionReview),
		EvaluateProductionUsabilityValDDegradedBehaviorReviewState(degradedReview),
		EvaluateProductionUsabilityValDUIWindowingReviewState(uiWindowingReview),
		EvaluateProductionUsabilityValDCommandNoiseReviewState(commandNoiseReview),
		EvaluateProductionUsabilityValDAPIProtectionReviewState(apiReview),
		EvaluateProductionUsabilityValDCLIResilienceReviewState(cliReview),
		EvaluateProductionUsabilityValDSupportabilityReviewState(supportabilityReview),
		EvaluateProductionUsabilityValDRecoveryReviewState(recoveryReview),
		EvaluateProductionUsabilityValDUpgradeRollbackReviewState(upgradeReview),
		EvaluateProductionUsabilityValDScaleEnvelopeReviewState(scaleReview),
		EvaluateProductionUsabilityValDGovernanceBoundaryReviewState(governanceReview),
		EvaluateProductionUsabilityValDRegressionGateState(regressionGate),
	); got != ProductionUsabilityValDStateActive {
		t.Fatalf("expected active overall Val D state, got %q", got)
	}
}
