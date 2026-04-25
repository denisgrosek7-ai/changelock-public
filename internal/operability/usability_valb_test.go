package operability

import "testing"

func TestProductionUsabilityValBStateRequiresActiveVal0(t *testing.T) {
	got := EvaluateProductionUsabilityValBState(
		ProductionUsabilityVal0StateSubstantial,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBUIDataResilienceStateActive,
		ProductionUsabilityValBWindowingStateActive,
		ProductionUsabilityValBResultSemanticsStateActive,
		ProductionUsabilityValBCommandCenterStateActive,
		ProductionUsabilityValBNoiseBudgetStateActive,
		ProductionUsabilityValBAPIProtectionStateActive,
		ProductionUsabilityValBCLIResilienceStateActive,
		ProductionUsabilityValBScaleEnvelopeStateActive,
		ProductionUsabilityValBActionModeEnforcementStateActive,
	)
	if got != ProductionUsabilityValBStateIncomplete {
		t.Fatalf("expected incomplete Val B state without active Val 0, got %q", got)
	}
}

func TestProductionUsabilityValBStateRequiresActiveValA(t *testing.T) {
	got := EvaluateProductionUsabilityValBState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateSubstantial,
		ProductionUsabilityValBUIDataResilienceStateActive,
		ProductionUsabilityValBWindowingStateActive,
		ProductionUsabilityValBResultSemanticsStateActive,
		ProductionUsabilityValBCommandCenterStateActive,
		ProductionUsabilityValBNoiseBudgetStateActive,
		ProductionUsabilityValBAPIProtectionStateActive,
		ProductionUsabilityValBCLIResilienceStateActive,
		ProductionUsabilityValBScaleEnvelopeStateActive,
		ProductionUsabilityValBActionModeEnforcementStateActive,
	)
	if got != ProductionUsabilityValBStateIncomplete {
		t.Fatalf("expected incomplete Val B state without active Val A, got %q", got)
	}
}

func TestProductionUsabilityValBUIDataResilienceRequiresFreshnessAndCanonicalDisclaimer(t *testing.T) {
	model := ProductionUsabilityValBUIDataResilience()
	model.Items[0].FreshnessState = ""
	if got := EvaluateProductionUsabilityValBUIDataResilienceState(model); got == ProductionUsabilityValBUIDataResilienceStateActive {
		t.Fatalf("expected non-active UI data resilience without freshness state, got %q", got)
	}

	model = ProductionUsabilityValBUIDataResilience()
	model.Items[0].CanonicalTruthDisclaimer = ""
	if got := EvaluateProductionUsabilityValBUIDataResilienceState(model); got == ProductionUsabilityValBUIDataResilienceStateActive {
		t.Fatalf("expected non-active UI data resilience without canonical disclaimer, got %q", got)
	}
}

func TestProductionUsabilityValBUIDataStatesRemainDistinct(t *testing.T) {
	model := ProductionUsabilityValBUIDataResilience()
	model.Items[1].FreshnessState = ProductionUsabilityStatusFresh
	if got := EvaluateProductionUsabilityValBUIDataResilienceState(model); got == ProductionUsabilityValBUIDataResilienceStateActive {
		t.Fatalf("expected non-active UI data resilience when freshness states collapse, got %q", got)
	}
}

func TestProductionUsabilityValBCachedOrPartialUIStateCannotClaimCanonicalTruth(t *testing.T) {
	model := ProductionUsabilityValBUIDataResilience()
	model.Items[2].ClaimsCanonicalTruth = true
	if got := EvaluateProductionUsabilityValBUIDataResilienceState(model); got == ProductionUsabilityValBUIDataResilienceStateActive {
		t.Fatalf("expected non-active UI data resilience when partial projection claims canonical truth, got %q", got)
	}
}

func TestProductionUsabilityValBWindowingRequiresBoundedMaxWindowSize(t *testing.T) {
	model := ProductionUsabilityValBWindowing()
	model.Limit = 600
	if got := EvaluateProductionUsabilityValBWindowingState(model); got == ProductionUsabilityValBWindowingStateActive {
		t.Fatalf("expected non-active windowing when limit exceeds max window size, got %q", got)
	}
}

func TestProductionUsabilityValBWindowedPartialResultsExposeLimitations(t *testing.T) {
	model := ProductionUsabilityValBWindowing()
	model.TruncationWarning = ""
	if got := EvaluateProductionUsabilityValBWindowingState(model); got == ProductionUsabilityValBWindowingStateActive {
		t.Fatalf("expected non-active windowing without truncation warning for partial result, got %q", got)
	}
}

func TestProductionUsabilityValBWindowingCannotClaimFullDatasetWithoutKnownTotal(t *testing.T) {
	model := ProductionUsabilityValBWindowing()
	model.ClaimsCompleteData = true
	if got := EvaluateProductionUsabilityValBWindowingState(model); got == ProductionUsabilityValBWindowingStateActive {
		t.Fatalf("expected non-active windowing when full dataset is claimed without known total, got %q", got)
	}
}

func TestProductionUsabilityValBPartialResultCannotReportFullSuccess(t *testing.T) {
	model := ProductionUsabilityValBResultSemantics()
	for idx := range model.Items {
		if model.Items[idx].ResultHealth == ProductionUsabilityResultPartial {
			model.Items[idx].ReportedAsFullSuccess = true
		}
	}
	if got := EvaluateProductionUsabilityValBResultSemanticsState(model); got == ProductionUsabilityValBResultSemanticsStateActive {
		t.Fatalf("expected non-active result semantics when partial result reports full success, got %q", got)
	}
}

func TestProductionUsabilityValBDegradedResponseCannotReportHealthy(t *testing.T) {
	model := ProductionUsabilityValBResultSemantics()
	for idx := range model.Items {
		if model.Items[idx].ResultHealth == ProductionUsabilityResultDegraded {
			model.Items[idx].ReportedAsHealthy = true
		}
	}
	if got := EvaluateProductionUsabilityValBResultSemanticsState(model); got == ProductionUsabilityValBResultSemanticsStateActive {
		t.Fatalf("expected non-active result semantics when degraded response reports healthy, got %q", got)
	}
}

func TestProductionUsabilityValBUnsupportedComponentsCannotBeSilentlyOmitted(t *testing.T) {
	model := ProductionUsabilityValBResultSemantics()
	for idx := range model.Items {
		if model.Items[idx].ResultHealth == ProductionUsabilityResultUnsupported {
			model.Items[idx].UnsupportedSilentlyOmitted = true
		}
	}
	if got := EvaluateProductionUsabilityValBResultSemanticsState(model); got == ProductionUsabilityValBResultSemanticsStateActive {
		t.Fatalf("expected non-active result semantics when unsupported components are silently omitted, got %q", got)
	}
}

func TestProductionUsabilityValBCommandCenterIsDecisionSupportOnly(t *testing.T) {
	model := ProductionUsabilityValBCommandCenterTasks()
	model.Items[0].DecisionSupportOnly = false
	if got := EvaluateProductionUsabilityValBCommandCenterState(model); got == ProductionUsabilityValBCommandCenterStateActive {
		t.Fatalf("expected non-active command center state when tasks are not decision support only, got %q", got)
	}
}

func TestProductionUsabilityValBTaskAcknowledgementDoesNotEqualRemediation(t *testing.T) {
	model := ProductionUsabilityValBCommandCenterTasks()
	model.Items[0].AcknowledgementEqualsRemediation = true
	if got := EvaluateProductionUsabilityValBCommandCenterState(model); got == ProductionUsabilityValBCommandCenterStateActive {
		t.Fatalf("expected non-active command center state when acknowledgement equals remediation, got %q", got)
	}
}

func TestProductionUsabilityValBTaskResolvedDoesNotEqualCanonicalClosureWithoutWorkflowEvidence(t *testing.T) {
	model := ProductionUsabilityValBCommandCenterTasks()
	model.Items[0].ResolvedEqualsCanonicalClosure = true
	if got := EvaluateProductionUsabilityValBCommandCenterState(model); got == ProductionUsabilityValBCommandCenterStateActive {
		t.Fatalf("expected non-active command center state when resolution equals canonical closure, got %q", got)
	}
}

func TestProductionUsabilityValBCommandCenterFailsClosedForTypoVisibilityScope(t *testing.T) {
	model := ProductionUsabilityValBCommandCenterTasks()
	model.Items[0].VisibilityScope = "partnr"
	if got := EvaluateProductionUsabilityValBCommandCenterState(model); got == ProductionUsabilityValBCommandCenterStateActive {
		t.Fatalf("expected non-active command center state for typo visibility scope, got %q", got)
	}
}

func TestProductionUsabilityValBCommandCenterFailsClosedForTypoRedactionTier(t *testing.T) {
	model := ProductionUsabilityValBCommandCenterTasks()
	model.Items[0].RedactionTier = "medum"
	if got := EvaluateProductionUsabilityValBCommandCenterState(model); got == ProductionUsabilityValBCommandCenterStateActive {
		t.Fatalf("expected non-active command center state for typo redaction tier, got %q", got)
	}
}

func TestProductionUsabilityValBCommandCenterPassesForSupportedVisibilityAndRedactionValues(t *testing.T) {
	model := ProductionUsabilityValBCommandCenterTasks()
	if got := EvaluateProductionUsabilityValBCommandCenterState(model); got != ProductionUsabilityValBCommandCenterStateActive {
		t.Fatalf("expected active command center state for supported visibility and redaction values, got %q", got)
	}
}

func TestProductionUsabilityValBNoiseSuppressionDoesNotHideCriticalBlockers(t *testing.T) {
	model := ProductionUsabilityValBNoiseBudget()
	model.Items[0].AcknowledgementState = ProductionUsabilityAckSuppressedDup
	model.Items[0].SuppressionReason = "duplicate"
	if got := EvaluateProductionUsabilityValBNoiseBudgetState(model); got == ProductionUsabilityValBNoiseBudgetStateActive {
		t.Fatalf("expected non-active noise budget when critical blocker is suppressed, got %q", got)
	}
}

func TestProductionUsabilityValBGroupingPreservesHighestSeverity(t *testing.T) {
	model := ProductionUsabilityValBNoiseBudget()
	model.Items[1].HighestSeverityPreserved = false
	if got := EvaluateProductionUsabilityValBNoiseBudgetState(model); got == ProductionUsabilityValBNoiseBudgetStateActive {
		t.Fatalf("expected non-active noise budget when highest severity is not preserved, got %q", got)
	}
}

func TestProductionUsabilityValBSuppressedDuplicatesRemainAuditable(t *testing.T) {
	model := ProductionUsabilityValBNoiseBudget()
	model.Items[1].SuppressedDuplicatesAuditable = false
	if got := EvaluateProductionUsabilityValBNoiseBudgetState(model); got == ProductionUsabilityValBNoiseBudgetStateActive {
		t.Fatalf("expected non-active noise budget when suppressed duplicates are not auditable, got %q", got)
	}
}

func TestProductionUsabilityValBAPIProtectionRequiresPriorityFairnessAndBackpressure(t *testing.T) {
	model := ProductionUsabilityValBAPIProtection()
	model.Items[0].FairnessScope = ""
	if got := EvaluateProductionUsabilityValBAPIProtectionState(model); got == ProductionUsabilityValBAPIProtectionStateActive {
		t.Fatalf("expected non-active API protection without fairness scope, got %q", got)
	}
}

func TestProductionUsabilityValBMutationRequestsDoNotBypassGovernanceViaPriority(t *testing.T) {
	model := ProductionUsabilityValBAPIProtection()
	for idx := range model.Items {
		if model.Items[idx].RequestClass == ProductionUsabilityRequestClassMutation {
			model.Items[idx].PriorityBypassesGovernance = true
		}
	}
	if got := EvaluateProductionUsabilityValBAPIProtectionState(model); got == ProductionUsabilityValBAPIProtectionStateActive {
		t.Fatalf("expected non-active API protection when mutation bypasses governance, got %q", got)
	}
}

func TestProductionUsabilityValBAPIProtectionDoesNotHidePolicyDenialAsThrottling(t *testing.T) {
	model := ProductionUsabilityValBAPIProtection()
	model.Items[0].PolicyDenialHiddenAsThrottling = true
	if got := EvaluateProductionUsabilityValBAPIProtectionState(model); got == ProductionUsabilityValBAPIProtectionStateActive {
		t.Fatalf("expected non-active API protection when policy denial is hidden as throttling, got %q", got)
	}
}

func TestProductionUsabilityValBCLIReadPreviewExplainDryRunAuditDoNotMutate(t *testing.T) {
	model := ProductionUsabilityValBCLIResilience()
	model.Items[0].MutatesCanonicalState = true
	if got := EvaluateProductionUsabilityValBCLIResilienceState(model); got == ProductionUsabilityValBCLIResilienceStateActive {
		t.Fatalf("expected non-active CLI resilience when non-mutating command mutates, got %q", got)
	}
}

func TestProductionUsabilityValBCLIMutatingOperationsAreNotAssumedRetrySafe(t *testing.T) {
	model := ProductionUsabilityValBCLIResilience()
	for idx := range model.Items {
		if model.Items[idx].CommandName == "changelock governed rotate" {
			model.Items[idx].RetrySafety = ProductionUsabilityRetrySafe
		}
	}
	if got := EvaluateProductionUsabilityValBCLIResilienceState(model); got == ProductionUsabilityValBCLIResilienceStateActive {
		t.Fatalf("expected non-active CLI resilience when side-effecting command is retry-safe, got %q", got)
	}
}

func TestProductionUsabilityValBCLIIdempotencyKeyRequiredFailsClosedWhenMissing(t *testing.T) {
	model := ProductionUsabilityValBCLIResilience()
	for idx := range model.Items {
		if model.Items[idx].CommandName == "changelock governed apply" {
			model.Items[idx].IdempotencyKeyPresent = false
		}
	}
	if got := EvaluateProductionUsabilityValBCLIResilienceState(model); got == ProductionUsabilityValBCLIResilienceStateActive {
		t.Fatalf("expected non-active CLI resilience when idempotency key is missing, got %q", got)
	}
}

func TestProductionUsabilityValBCLIRetryUnsafeExposesDoNotRetryReason(t *testing.T) {
	model := ProductionUsabilityValBCLIResilience()
	for idx := range model.Items {
		if model.Items[idx].CommandName == "changelock governed rotate" {
			model.Items[idx].DoNotRetryReason = ""
		}
	}
	if got := EvaluateProductionUsabilityValBCLIResilienceState(model); got == ProductionUsabilityValBCLIResilienceStateActive {
		t.Fatalf("expected non-active CLI resilience when retry-unsafe command lacks do-not-retry reason, got %q", got)
	}
}

func TestProductionUsabilityValBCLIPartialFailureIsExplicit(t *testing.T) {
	model := ProductionUsabilityValBCLIResilience()
	model.Items[0].PartialFailureBehavior = ""
	if got := EvaluateProductionUsabilityValBCLIResilienceState(model); got == ProductionUsabilityValBCLIResilienceStateActive {
		t.Fatalf("expected non-active CLI resilience when partial failure behavior is missing, got %q", got)
	}
}

func TestProductionUsabilityValBScaleEnvelopeMarksUnknownOrUnmeasuredAsLimitation(t *testing.T) {
	model := ProductionUsabilityValBScaleEnvelope()
	model.LimitationDisclaimer = "bounded_projection_only"
	if got := EvaluateProductionUsabilityValBScaleEnvelopeState(model); got == ProductionUsabilityValBScaleEnvelopeStateActive {
		t.Fatalf("expected non-active scale envelope without explicit unmeasured limitation, got %q", got)
	}
}

func TestProductionUsabilityValBUnsupportedScaleConditionsAreExplicit(t *testing.T) {
	model := ProductionUsabilityValBScaleEnvelope()
	model.UnsupportedScaleConditions = nil
	if got := EvaluateProductionUsabilityValBScaleEnvelopeState(model); got == ProductionUsabilityValBScaleEnvelopeStateActive {
		t.Fatalf("expected non-active scale envelope without unsupported scale conditions, got %q", got)
	}
}

func TestProductionUsabilityValBSafeActionModeEnforcementPreventsPreviewOrExplainMutation(t *testing.T) {
	model := ProductionUsabilityValBActionModeEnforcement()
	for idx := range model.Items {
		if model.Items[idx].ActionMode == ProductionUsabilityActionModePreview {
			model.Items[idx].ExecutesMutation = true
		}
	}
	if got := EvaluateProductionUsabilityValBActionModeEnforcementState(model); got == ProductionUsabilityValBActionModeEnforcementStateActive {
		t.Fatalf("expected non-active action mode enforcement when preview executes mutation, got %q", got)
	}
}

func TestProductionUsabilityValBActionModeEnforcementDoesNotReferenceMissingExplainRoute(t *testing.T) {
	model := ProductionUsabilityValBActionModeEnforcement()
	for _, item := range model.Items {
		if item.SurfaceRef == "/v1/production/usability-operability-recovery/valb/explain" {
			t.Fatalf("expected Val B action-mode enforcement to avoid missing /valb/explain route")
		}
	}
}

func TestProductionUsabilityValBActionModeEnforcementOnlyUsesRegisteredSurfaceRefs(t *testing.T) {
	model := ProductionUsabilityValBActionModeEnforcement()
	expected := productionUsabilityValBRegisteredSurfaceRefs()
	for _, item := range model.Items {
		if !containsTrimmedString(expected, item.SurfaceRef) {
			t.Fatalf("expected %q to be a registered Val B surface ref", item.SurfaceRef)
		}
	}

	model.Items[0].SurfaceRef = "/v1/production/usability-operability-recovery/valb/explain"
	if got := EvaluateProductionUsabilityValBActionModeEnforcementState(model); got == ProductionUsabilityValBActionModeEnforcementStateActive {
		t.Fatalf("expected non-active action mode enforcement for unregistered surface ref, got %q", got)
	}
}

func TestProductionUsabilityValBActionModeEnforcementExplainUsesRegisteredAPISurface(t *testing.T) {
	model := ProductionUsabilityValBActionModeEnforcement()
	for _, item := range model.Items {
		if item.ActionMode != ProductionUsabilityActionModeExplain {
			continue
		}
		if item.SurfaceRef != "/v1/production/usability-operability-recovery/valb/api-protection" {
			t.Fatalf("expected explain action mode to use api-protection surface, got %q", item.SurfaceRef)
		}
		if item.OperationType != ProductionUsabilityOperationReadOnly || !item.NonMutating || item.ExecutesMutation {
			t.Fatalf("expected explain action mode to remain read-only and non-mutating, got %#v", item)
		}
		return
	}
	t.Fatalf("expected explain action mode item to be present")
}

func TestProductionUsabilityValBEnforceOrMutateRequireGovernanceOrRemainUnavailable(t *testing.T) {
	model := ProductionUsabilityValBActionModeEnforcement()
	for idx := range model.Items {
		if model.Items[idx].ActionMode == ProductionUsabilityActionModeMutate {
			model.Items[idx].GovernanceRef = ""
		}
	}
	if got := EvaluateProductionUsabilityValBActionModeEnforcementState(model); got == ProductionUsabilityValBActionModeEnforcementStateActive {
		t.Fatalf("expected non-active action mode enforcement when mutate surface lacks governance ref, got %q", got)
	}
}

func TestProductionUsabilityValBProofsPassOnlyAsResilienceWhilePoint4RemainsNotComplete(t *testing.T) {
	got := EvaluateProductionUsabilityValBProofsState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBUIDataResilienceStateActive,
		ProductionUsabilityValBWindowingStateActive,
		ProductionUsabilityValBResultSemanticsStateActive,
		ProductionUsabilityValBCommandCenterStateActive,
		ProductionUsabilityValBNoiseBudgetStateActive,
		ProductionUsabilityValBAPIProtectionStateActive,
		ProductionUsabilityValBCLIResilienceStateActive,
		ProductionUsabilityValBScaleEnvelopeStateActive,
		ProductionUsabilityValBActionModeEnforcementStateActive,
		[]string{
			"/v1/production/usability-operability-recovery/val0/proofs",
			"/v1/production/usability-operability-recovery/vala/proofs",
			"/v1/production/usability-operability-recovery/valb/ui-data-resilience",
			"/v1/production/usability-operability-recovery/valb/windowing",
			"/v1/production/usability-operability-recovery/valb/result-semantics",
			"/v1/production/usability-operability-recovery/valb/command-center-tasks",
			"/v1/production/usability-operability-recovery/valb/noise-budget",
			"/v1/production/usability-operability-recovery/valb/api-protection",
			"/v1/production/usability-operability-recovery/valb/cli-resilience",
			"/v1/production/usability-operability-recovery/valb/scale-envelope",
			"/v1/production/usability-operability-recovery/valb/action-mode-enforcement",
			"/v1/production/usability-operability-recovery/valb/proofs",
		},
		[]string{"val0_proofs", "vala_proofs", "ui_data", "windowing", "result_semantics", "command_center", "noise_budget", "api_protection"},
		[]string{"Val B proves resilience only."},
		[]string{"Point 4 full PASS remains deferred."},
	)
	if got != ProductionUsabilityValBStateActive {
		t.Fatalf("expected active Val B proofs state for complete resilience slice, got %q", got)
	}
	if ProductionUsabilityPoint4StateNotComplete == ProductionUsabilityValBStateActive {
		t.Fatalf("point 4 state must remain distinct from active Val B state")
	}
}

func TestProductionUsabilityValBMissingRequiredComponentKeepsValBNonActive(t *testing.T) {
	got := EvaluateProductionUsabilityValBState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBUIDataResilienceStateActive,
		ProductionUsabilityValBWindowingStateActive,
		ProductionUsabilityValBResultSemanticsStateActive,
		ProductionUsabilityValBCommandCenterStateActive,
		ProductionUsabilityValBNoiseBudgetStateIncomplete,
		ProductionUsabilityValBAPIProtectionStateActive,
		ProductionUsabilityValBCLIResilienceStateActive,
		ProductionUsabilityValBScaleEnvelopeStateActive,
		ProductionUsabilityValBActionModeEnforcementStateActive,
	)
	if got == ProductionUsabilityValBStateActive {
		t.Fatalf("expected non-active Val B state with missing required component, got %q", got)
	}
}

func TestProductionUsabilityValBFoundationIsActive(t *testing.T) {
	uiData := ProductionUsabilityValBUIDataResilience()
	if got := EvaluateProductionUsabilityValBUIDataResilienceState(uiData); got != ProductionUsabilityValBUIDataResilienceStateActive {
		t.Fatalf("expected active UI data resilience state, got %q", got)
	}

	windowing := ProductionUsabilityValBWindowing()
	if got := EvaluateProductionUsabilityValBWindowingState(windowing); got != ProductionUsabilityValBWindowingStateActive {
		t.Fatalf("expected active windowing state, got %q", got)
	}

	resultSemantics := ProductionUsabilityValBResultSemantics()
	if got := EvaluateProductionUsabilityValBResultSemanticsState(resultSemantics); got != ProductionUsabilityValBResultSemanticsStateActive {
		t.Fatalf("expected active result semantics state, got %q", got)
	}

	commandCenter := ProductionUsabilityValBCommandCenterTasks()
	if got := EvaluateProductionUsabilityValBCommandCenterState(commandCenter); got != ProductionUsabilityValBCommandCenterStateActive {
		t.Fatalf("expected active command center state, got %q", got)
	}

	noiseBudget := ProductionUsabilityValBNoiseBudget()
	if got := EvaluateProductionUsabilityValBNoiseBudgetState(noiseBudget); got != ProductionUsabilityValBNoiseBudgetStateActive {
		t.Fatalf("expected active noise budget state, got %q", got)
	}

	apiProtection := ProductionUsabilityValBAPIProtection()
	if got := EvaluateProductionUsabilityValBAPIProtectionState(apiProtection); got != ProductionUsabilityValBAPIProtectionStateActive {
		t.Fatalf("expected active API protection state, got %q", got)
	}

	cliResilience := ProductionUsabilityValBCLIResilience()
	if got := EvaluateProductionUsabilityValBCLIResilienceState(cliResilience); got != ProductionUsabilityValBCLIResilienceStateActive {
		t.Fatalf("expected active CLI resilience state, got %q", got)
	}

	scaleEnvelope := ProductionUsabilityValBScaleEnvelope()
	if got := EvaluateProductionUsabilityValBScaleEnvelopeState(scaleEnvelope); got != ProductionUsabilityValBScaleEnvelopeStateActive {
		t.Fatalf("expected active scale envelope state, got %q", got)
	}

	actionModes := ProductionUsabilityValBActionModeEnforcement()
	if got := EvaluateProductionUsabilityValBActionModeEnforcementState(actionModes); got != ProductionUsabilityValBActionModeEnforcementStateActive {
		t.Fatalf("expected active action mode enforcement state, got %q", got)
	}

	if got := EvaluateProductionUsabilityValBState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		EvaluateProductionUsabilityValBUIDataResilienceState(uiData),
		EvaluateProductionUsabilityValBWindowingState(windowing),
		EvaluateProductionUsabilityValBResultSemanticsState(resultSemantics),
		EvaluateProductionUsabilityValBCommandCenterState(commandCenter),
		EvaluateProductionUsabilityValBNoiseBudgetState(noiseBudget),
		EvaluateProductionUsabilityValBAPIProtectionState(apiProtection),
		EvaluateProductionUsabilityValBCLIResilienceState(cliResilience),
		EvaluateProductionUsabilityValBScaleEnvelopeState(scaleEnvelope),
		EvaluateProductionUsabilityValBActionModeEnforcementState(actionModes),
	); got != ProductionUsabilityValBStateActive {
		t.Fatalf("expected active overall Val B state, got %q", got)
	}
}
