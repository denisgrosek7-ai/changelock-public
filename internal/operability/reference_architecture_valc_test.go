package operability

import "testing"

func activeReferenceArchitectureValCPrereqs() (string, string, string, string, string, string, string, string) {
	return IntelligenceCalibrationPoint5StatePass,
		ReferenceArchitectureVal0StateActive,
		ReferenceArchitectureVal0StateActive,
		ReferenceArchitectureValAStateActive,
		ReferenceArchitectureValAStateActive,
		ReferenceArchitectureValBStateActive,
		ReferenceArchitectureValBStateActive,
		ReferenceArchitecturePoint6StateNotComplete
}

func activeReferenceArchitectureValCComponents() (
	ReferenceArchitectureResilienceScenarioPackRegistry,
	ReferenceArchitectureFailureModeTaxonomy,
	ReferenceArchitectureScenarioDescriptorCollection,
	ReferenceArchitectureDegradedModeCollection,
	ReferenceArchitectureRecoveryExpectationCollection,
	ReferenceArchitectureScalingScenarioCollection,
	ReferenceArchitectureTrustPathCollection,
	ReferenceArchitectureAuditPathCollection,
	ReferenceArchitectureControlPlaneSafetyCollection,
) {
	return ReferenceArchitectureValCScenarioPackRegistry(),
		ReferenceArchitectureValCFailureModeTaxonomy(),
		ReferenceArchitectureValCScenarioDescriptorCollection(),
		ReferenceArchitectureValCDegradedModeCollection(),
		ReferenceArchitectureValCRecoveryExpectationCollection(),
		ReferenceArchitectureValCScalingScenarioCollection(),
		ReferenceArchitectureValCTrustPathCollection(),
		ReferenceArchitectureValCAuditPathCollection(),
		ReferenceArchitectureValCControlPlaneSafetyCollection()
}

func TestReferenceArchitectureValCDependencyGates(t *testing.T) {
	point5State, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, point6State := activeReferenceArchitectureValCPrereqs()
	registry, taxonomy, descriptors, degraded, recovery, scaling, trust, audit, control := activeReferenceArchitectureValCComponents()
	scenarioPackState := EvaluateReferenceArchitectureValCScenarioPackRegistryState(registry)
	failureTaxonomyState := EvaluateReferenceArchitectureValCFailureModeTaxonomyState(taxonomy)
	scenarioDescriptorState := EvaluateReferenceArchitectureValCScenarioDescriptorCollectionState(descriptors)
	degradedModeState := EvaluateReferenceArchitectureValCDegradedModeCollectionState(degraded)
	recoveryState := EvaluateReferenceArchitectureValCRecoveryExpectationCollectionState(recovery)
	scalingState := EvaluateReferenceArchitectureValCScalingScenarioCollectionState(scaling)
	trustState := EvaluateReferenceArchitectureValCTrustPathCollectionState(trust)
	auditState := EvaluateReferenceArchitectureValCAuditPathCollectionState(audit)
	controlState := EvaluateReferenceArchitectureValCControlPlaneCollectionState(control)

	if got := EvaluateReferenceArchitectureValCState(point5State, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, point6State, scenarioPackState, failureTaxonomyState, scenarioDescriptorState, degradedModeState, recoveryState, scalingState, trustState, auditState, controlState); got != ReferenceArchitectureValCStateActive {
		t.Fatalf("expected active Val C state with valid dependencies and components, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValCState(point5State, ReferenceArchitectureVal0StateSubstantial, val0State, valACurrentState, valAState, valBCurrentState, valBState, point6State, scenarioPackState, failureTaxonomyState, scenarioDescriptorState, degradedModeState, recoveryState, scalingState, trustState, auditState, controlState); got != ReferenceArchitectureValCStateBlocked {
		t.Fatalf("expected blocked Val C state when Val 0 dependency is missing, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValCState(point5State, val0CurrentState, val0State, ReferenceArchitectureValAStatePartial, valAState, valBCurrentState, valBState, point6State, scenarioPackState, failureTaxonomyState, scenarioDescriptorState, degradedModeState, recoveryState, scalingState, trustState, auditState, controlState); got != ReferenceArchitectureValCStateBlocked {
		t.Fatalf("expected blocked Val C state when Val A dependency is missing, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValCState(point5State, val0CurrentState, val0State, valACurrentState, valAState, ReferenceArchitectureValBStatePartial, valBState, point6State, scenarioPackState, failureTaxonomyState, scenarioDescriptorState, degradedModeState, recoveryState, scalingState, trustState, auditState, controlState); got != ReferenceArchitectureValCStateBlocked {
		t.Fatalf("expected blocked Val C state when Val B dependency is missing, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValCState(point5State, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, ReferenceArchitecturePoint6StatePass, scenarioPackState, failureTaxonomyState, scenarioDescriptorState, degradedModeState, recoveryState, scalingState, trustState, auditState, controlState); got != ReferenceArchitectureValCStateBlocked {
		t.Fatalf("expected blocked Val C state when point 6 is not not_complete, got %q", got)
	}
}

func TestReferenceArchitectureValCScenarioPackValidation(t *testing.T) {
	registry := ReferenceArchitectureValCScenarioPackRegistry()
	pack := registry.ScenarioPacks[0]
	if got := EvaluateReferenceArchitectureValCScenarioPackState(pack); got != ReferenceArchitectureValCScenarioPackStateActive {
		t.Fatalf("expected active scenario pack state for valid fixture, got %q", got)
	}
	pack = registry.ScenarioPacks[0]
	pack.ScenarioPackID = ""
	if got := EvaluateReferenceArchitectureValCScenarioPackState(pack); got == ReferenceArchitectureValCScenarioPackStateActive {
		t.Fatalf("expected non-active scenario pack state for missing scenario_pack_id, got %q", got)
	}
	pack = registry.ScenarioPacks[0]
	pack.BlueprintFamily = "enterprise-defualt"
	if got := EvaluateReferenceArchitectureValCScenarioPackState(pack); got == ReferenceArchitectureValCScenarioPackStateActive {
		t.Fatalf("expected non-active scenario pack state for unknown family, got %q", got)
	}
	pack = registry.ScenarioPacks[0]
	pack.LifecycleState = "acitve"
	if got := EvaluateReferenceArchitectureValCScenarioPackState(pack); got == ReferenceArchitectureValCScenarioPackStateActive {
		t.Fatalf("expected non-active scenario pack state for typo lifecycle, got %q", got)
	}
	pack = registry.ScenarioPacks[0]
	pack.ProjectionDisclaimer = ""
	if got := EvaluateReferenceArchitectureValCScenarioPackState(pack); got == ReferenceArchitectureValCScenarioPackStateActive {
		t.Fatalf("expected non-active scenario pack state without projection disclaimer, got %q", got)
	}
	pack = registry.ScenarioPacks[0]
	pack.GuaranteedResilience = true
	if got := EvaluateReferenceArchitectureValCScenarioPackState(pack); got != ReferenceArchitectureValCScenarioPackStateBlocked {
		t.Fatalf("expected blocked scenario pack state for guaranteed resilience language, got %q", got)
	}
}

func TestReferenceArchitectureValCFailureModeTaxonomyValidation(t *testing.T) {
	taxonomy := ReferenceArchitectureValCFailureModeTaxonomy()
	if got := EvaluateReferenceArchitectureValCFailureModeTaxonomyState(taxonomy); got != ReferenceArchitectureValCFailureTaxonomyStateActive {
		t.Fatalf("expected active failure taxonomy state, got %q", got)
	}
	taxonomy = ReferenceArchitectureValCFailureModeTaxonomy()
	taxonomy.SupportedCategories = append(taxonomy.SupportedCategories[1:], taxonomy.SupportedCategories[0])
	taxonomy.SupportedCategories[0] = ReferenceArchitectureValCFailureTrustAnchorUnavailable
	if got := EvaluateReferenceArchitectureValCFailureModeTaxonomyState(taxonomy); got == ReferenceArchitectureValCFailureTaxonomyStateActive {
		t.Fatalf("expected duplicate category to fail closed, got %q", got)
	}
	taxonomy = ReferenceArchitectureValCFailureModeTaxonomy()
	taxonomy.SupportedCategories[0] = "trust_anchor_unavailble"
	if got := EvaluateReferenceArchitectureValCFailureModeTaxonomyState(taxonomy); got == ReferenceArchitectureValCFailureTaxonomyStateActive {
		t.Fatalf("expected typo category to fail closed, got %q", got)
	}
	taxonomy = ReferenceArchitectureValCFailureModeTaxonomy()
	taxonomy.SupportedCategories = append(taxonomy.SupportedCategories, "memory_pressure")
	if got := EvaluateReferenceArchitectureValCFailureModeTaxonomyState(taxonomy); got == ReferenceArchitectureValCFailureTaxonomyStateActive {
		t.Fatalf("expected extra unsupported category to fail closed, got %q", got)
	}
}

func TestReferenceArchitectureValCScenarioDescriptorValidation(t *testing.T) {
	collection := ReferenceArchitectureValCScenarioDescriptorCollection()
	pack := collection.Packs[0]
	if got := EvaluateReferenceArchitectureValCScenarioDescriptorPackState(pack); got != ReferenceArchitectureValCScenarioDescriptorStateActive {
		t.Fatalf("expected active scenario descriptor state, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Scenarios[0].ExpectedState = "readyish"
	if got := EvaluateReferenceArchitectureValCScenarioDescriptorPackState(pack); got == ReferenceArchitectureValCScenarioDescriptorStateActive {
		t.Fatalf("expected unknown expected state to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Scenarios[0].Severity = "sev0"
	if got := EvaluateReferenceArchitectureValCScenarioDescriptorPackState(pack); got == ReferenceArchitectureValCScenarioDescriptorStateActive {
		t.Fatalf("expected unknown severity to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Scenarios[0].RequiredEvidenceTypes = nil
	if got := EvaluateReferenceArchitectureValCScenarioDescriptorPackState(pack); got == ReferenceArchitectureValCScenarioDescriptorStateActive {
		t.Fatalf("expected missing evidence requirement to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Scenarios[0].BlocksMatched = true
	pack.Scenarios[0].ExpectedState = ReferenceArchitectureValCScenarioReady
	if got := EvaluateReferenceArchitectureValCScenarioDescriptorPackState(pack); got == ReferenceArchitectureValCScenarioDescriptorStateActive {
		t.Fatalf("expected blocking scenario with ready state to fail closed, got %q", got)
	}
}

func TestReferenceArchitectureValCDegradedModeValidation(t *testing.T) {
	collection := ReferenceArchitectureValCDegradedModeCollection()
	pack := collection.Packs[0]
	if got := EvaluateReferenceArchitectureValCDegradedModePackState(pack); got != ReferenceArchitectureValCDegradedModeStateActive {
		t.Fatalf("expected active degraded mode state, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Modes[0].BlockedOperations = nil
	if got := EvaluateReferenceArchitectureValCDegradedModePackState(pack); got == ReferenceArchitectureValCDegradedModeStateActive {
		t.Fatalf("expected missing blocked operations to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Modes[0].RequiredOperatorAction = ""
	if got := EvaluateReferenceArchitectureValCDegradedModePackState(pack); got == ReferenceArchitectureValCDegradedModeStateActive {
		t.Fatalf("expected missing operator action to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Modes[0].EvidenceRequired = nil
	if got := EvaluateReferenceArchitectureValCDegradedModePackState(pack); got == ReferenceArchitectureValCDegradedModeStateActive {
		t.Fatalf("expected missing evidence requirement to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Modes[0].UnsupportedBehavior = true
	if got := EvaluateReferenceArchitectureValCDegradedModePackState(pack); got == ReferenceArchitectureValCDegradedModeStateActive {
		t.Fatalf("expected unsupported degraded behavior to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Modes[0].RedactionKeepsCaveats = false
	if got := EvaluateReferenceArchitectureValCDegradedModePackState(pack); got == ReferenceArchitectureValCDegradedModeStateActive {
		t.Fatalf("expected caveat redaction bypass to fail closed, got %q", got)
	}
}

func TestReferenceArchitectureValCRecoveryExpectationValidation(t *testing.T) {
	collection := ReferenceArchitectureValCRecoveryExpectationCollection()
	pack := collection.Packs[0]
	if got := EvaluateReferenceArchitectureValCRecoveryExpectationPackState(pack); got != ReferenceArchitectureValCRecoveryExpectationStateActive {
		t.Fatalf("expected active recovery expectation state, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Expectations[0].ExpectedRecoveryPath = ""
	if got := EvaluateReferenceArchitectureValCRecoveryExpectationPackState(pack); got == ReferenceArchitectureValCRecoveryExpectationStateActive {
		t.Fatalf("expected missing recovery path to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Expectations[0].RequiredEvidenceTypes = nil
	if got := EvaluateReferenceArchitectureValCRecoveryExpectationPackState(pack); got == ReferenceArchitectureValCRecoveryExpectationStateActive {
		t.Fatalf("expected missing recovery evidence to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Expectations[0].Timestamp = "2026/04/26"
	if got := EvaluateReferenceArchitectureValCRecoveryExpectationPackState(pack); got == ReferenceArchitectureValCRecoveryExpectationStateActive {
		t.Fatalf("expected malformed recovery timestamp to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Expectations[0].FreshnessState = IntelligenceCalibrationFreshnessStale
	if got := EvaluateReferenceArchitectureValCRecoveryExpectationPackState(pack); got == ReferenceArchitectureValCRecoveryExpectationStateActive {
		t.Fatalf("expected stale recovery evidence to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Expectations[0].CertifiedRecoveryClaim = true
	if got := EvaluateReferenceArchitectureValCRecoveryExpectationPackState(pack); got == ReferenceArchitectureValCRecoveryExpectationStateActive {
		t.Fatalf("expected certified recovery language to fail closed, got %q", got)
	}
}

func TestReferenceArchitectureValCScalingScenarioValidation(t *testing.T) {
	collection := ReferenceArchitectureValCScalingScenarioCollection()
	pack := collection.Packs[0]
	if got := EvaluateReferenceArchitectureValCScalingScenarioPackState(pack); got != ReferenceArchitectureValCScalingScenarioStateActive {
		t.Fatalf("expected active scaling scenario state, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Scenarios[0].Category = "queue_baklog_behavior"
	if got := EvaluateReferenceArchitectureValCScalingScenarioPackState(pack); got == ReferenceArchitectureValCScalingScenarioStateActive {
		t.Fatalf("expected unknown scaling category to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Scenarios[0].DegradationThreshold = 0
	if got := EvaluateReferenceArchitectureValCScalingScenarioPackState(pack); got == ReferenceArchitectureValCScalingScenarioStateActive {
		t.Fatalf("expected missing threshold to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Scenarios[0].Timestamp = "2026/04/26"
	if got := EvaluateReferenceArchitectureValCScalingScenarioPackState(pack); got == ReferenceArchitectureValCScalingScenarioStateActive {
		t.Fatalf("expected malformed timestamp to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Scenarios[0].FreshnessState = IntelligenceCalibrationFreshnessStale
	if got := EvaluateReferenceArchitectureValCScalingScenarioPackState(pack); got == ReferenceArchitectureValCScalingScenarioStateActive {
		t.Fatalf("expected stale scaling evidence to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Scenarios[0].PerformanceGuarantee = true
	if got := EvaluateReferenceArchitectureValCScalingScenarioPackState(pack); got == ReferenceArchitectureValCScalingScenarioStateActive {
		t.Fatalf("expected performance guarantee claim to fail closed, got %q", got)
	}
	pack = collection.Packs[0]
	pack.Scenarios[0].DegradationThreshold = 200
	pack.Scenarios[0].FailClosedThreshold = 100
	if got := EvaluateReferenceArchitectureValCScalingScenarioPackState(pack); got == ReferenceArchitectureValCScalingScenarioStateActive {
		t.Fatalf("expected invalid threshold ordering to fail closed, got %q", got)
	}
}

func TestReferenceArchitectureValCTrustPathContinuity(t *testing.T) {
	collection := ReferenceArchitectureValCTrustPathCollection()
	if got := EvaluateReferenceArchitectureValCTrustPathCollectionState(collection); got != ReferenceArchitectureValCTrustPathStateActive {
		t.Fatalf("expected active trust-path collection state, got %q", got)
	}
	enterprise := collection.Checks[0]
	highAssurance := collection.Checks[1]
	if !highAssurance.StrictCustodyBoundaryExpected || enterprise.StrictCustodyBoundaryExpected {
		t.Fatalf("expected high assurance to be stricter than enterprise default")
	}
	check := collection.Checks[0]
	check.EvidenceRefs = nil
	if got := EvaluateReferenceArchitectureValCTrustPathCheckState(check); got == ReferenceArchitectureValCTrustPathStateActive {
		t.Fatalf("expected missing trust evidence to fail closed, got %q", got)
	}
	check = collection.Checks[0]
	check.EvidenceRefs[0].FreshnessState = IntelligenceCalibrationFreshnessStale
	if got := EvaluateReferenceArchitectureValCTrustPathCheckState(check); got == ReferenceArchitectureValCTrustPathStateActive {
		t.Fatalf("expected stale trust evidence to fail closed, got %q", got)
	}
	var sovereignCheck ReferenceArchitectureTrustPathContinuityCheck
	for _, candidate := range collection.Checks {
		if candidate.BlueprintFamily == ReferenceArchitectureFamilySovereignAirGapped {
			sovereignCheck = candidate
			break
		}
	}
	if sovereignCheck.RequiresLiveExternalTrustDependency {
		t.Fatalf("expected sovereign air-gapped trust path not to require live external dependency by default")
	}
	sovereignCheck.RequiresLiveExternalTrustDependency = true
	if got := EvaluateReferenceArchitectureValCTrustPathCheckState(sovereignCheck); got != ReferenceArchitectureValCTrustPathStateBlocked {
		t.Fatalf("expected unsupported sovereign trust path to be blocked, got %q", got)
	}
}

func TestReferenceArchitectureValCAuditPathDegradation(t *testing.T) {
	collection := ReferenceArchitectureValCAuditPathCollection()
	check := collection.Checks[0]
	if got := EvaluateReferenceArchitectureValCAuditPathCheckState(check); got != ReferenceArchitectureValCAuditPathStateActive {
		t.Fatalf("expected active audit-path state, got %q", got)
	}
	check = collection.Checks[0]
	check.AuditWriterAvailabilityExpected = false
	if got := EvaluateReferenceArchitectureValCAuditPathCheckState(check); got == ReferenceArchitectureValCAuditPathStateActive {
		t.Fatalf("expected unavailable audit writer to fail closed, got %q", got)
	}
	check = collection.Checks[0]
	check.AuditLatencyDegradedBehavior = ""
	if got := EvaluateReferenceArchitectureValCAuditPathCheckState(check); got == ReferenceArchitectureValCAuditPathStateActive {
		t.Fatalf("expected missing audit degradation semantics to fail closed, got %q", got)
	}
	check = collection.Checks[0]
	check.EvidenceCustodyPath = ""
	if got := EvaluateReferenceArchitectureValCAuditPathCheckState(check); got == ReferenceArchitectureValCAuditPathStateActive {
		t.Fatalf("expected missing evidence custody path to fail closed, got %q", got)
	}
	check = collection.Checks[0]
	check.CanonicalFailureStatePreserved = false
	if got := EvaluateReferenceArchitectureValCAuditPathCheckState(check); got == ReferenceArchitectureValCAuditPathStateActive {
		t.Fatalf("expected audit degradation not to suppress canonical failure state, got %q", got)
	}
}

func TestReferenceArchitectureValCControlPlaneSafety(t *testing.T) {
	collection := ReferenceArchitectureValCControlPlaneSafetyCollection()
	check := collection.Checks[0]
	if got := EvaluateReferenceArchitectureValCControlPlaneCheckState(check); got != ReferenceArchitectureValCControlPlaneStateActive {
		t.Fatalf("expected active control-plane safety state, got %q", got)
	}
	check = collection.Checks[0]
	check.BackpressureSemantics = ""
	if got := EvaluateReferenceArchitectureValCControlPlaneCheckState(check); got == ReferenceArchitectureValCControlPlaneStateActive {
		t.Fatalf("expected missing backpressure semantics to fail closed, got %q", got)
	}
	check = collection.Checks[0]
	check.DependencyTimeoutBehavior = "active"
	if got := EvaluateReferenceArchitectureValCControlPlaneCheckState(check); got == ReferenceArchitectureValCControlPlaneStateActive {
		t.Fatalf("expected timeout behavior not to be active, got %q", got)
	}
	check = collection.Checks[0]
	check.OperatorActionRequired = ""
	if got := EvaluateReferenceArchitectureValCControlPlaneCheckState(check); got == ReferenceArchitectureValCControlPlaneStateActive {
		t.Fatalf("expected missing operator action requirement to fail closed, got %q", got)
	}
	check = collection.Checks[0]
	check.AutomaticApproval = true
	if got := EvaluateReferenceArchitectureValCControlPlaneCheckState(check); got != ReferenceArchitectureValCControlPlaneStateBlocked {
		t.Fatalf("expected automatic approval to block control-plane state, got %q", got)
	}
}

func TestReferenceArchitectureValCProofsRequireExactSurfaceSet(t *testing.T) {
	baseState := ReferenceArchitectureValCStateActive
	point6State := ReferenceArchitecturePoint6StateNotComplete
	supportedFamilies := referenceArchitectureVal0Families()
	evidenceRefs := []string{"point5", "val0", "vala", "valb", "registry", "taxonomy", "descriptors", "recovery", "trust", "audit", "control", "scenario-pack"}
	limitations := []string{"Val C keeps point 6 not complete."}
	disclaimer := referenceArchitectureValCProjectionDisclaimer()

	if got := EvaluateReferenceArchitectureValCProofsState(baseState, point6State, supportedFamilies, referenceArchitectureValCProofSurfaceRefs(), evidenceRefs, limitations, disclaimer); got != ReferenceArchitectureValCStateActive {
		t.Fatalf("expected active proofs state with exact surface set, got %q", got)
	}

	testCases := []struct {
		name        string
		surfaceRefs []string
	}{
		{name: "missing val0 proofs", surfaceRefs: []string{"/v1/reference-architecture/vala/proofs", "/v1/reference-architecture/valb/proofs", "/v1/reference-architecture/valc/scenario-packs", "/v1/reference-architecture/valc/failure-taxonomy", "/v1/reference-architecture/valc/scenario-descriptors", "/v1/reference-architecture/valc/degraded-modes", "/v1/reference-architecture/valc/recovery-expectations", "/v1/reference-architecture/valc/scaling-scenarios", "/v1/reference-architecture/valc/trust-path", "/v1/reference-architecture/valc/audit-path", "/v1/reference-architecture/valc/control-plane-safety", "/v1/reference-architecture/valc/proofs"}},
		{name: "duplicate does not compensate", surfaceRefs: []string{"/v1/reference-architecture/val0/proofs", "/v1/reference-architecture/val0/proofs", "/v1/reference-architecture/valb/proofs", "/v1/reference-architecture/valc/scenario-packs", "/v1/reference-architecture/valc/failure-taxonomy", "/v1/reference-architecture/valc/scenario-descriptors", "/v1/reference-architecture/valc/degraded-modes", "/v1/reference-architecture/valc/recovery-expectations", "/v1/reference-architecture/valc/scaling-scenarios", "/v1/reference-architecture/valc/trust-path", "/v1/reference-architecture/valc/audit-path", "/v1/reference-architecture/valc/control-plane-safety", "/v1/reference-architecture/valc/proofs"}},
		{name: "unknown extra does not compensate", surfaceRefs: []string{"/v1/reference-architecture/val0/proofs", "/v1/reference-architecture/vala/proofs", "/v1/reference-architecture/valb/proofs", "/v1/reference-architecture/valc/failure-taxonomy", "/v1/reference-architecture/valc/scenario-descriptors", "/v1/reference-architecture/valc/degraded-modes", "/v1/reference-architecture/valc/recovery-expectations", "/v1/reference-architecture/valc/scaling-scenarios", "/v1/reference-architecture/valc/trust-path", "/v1/reference-architecture/valc/audit-path", "/v1/reference-architecture/valc/control-plane-safety", "/v1/reference-architecture/valc/proofs", "/v1/reference-architecture/valc/extra"}},
		{name: "whitespace ref fails closed", surfaceRefs: []string{"/v1/reference-architecture/val0/proofs", "/v1/reference-architecture/vala/proofs", "/v1/reference-architecture/valb/proofs", "/v1/reference-architecture/valc/scenario-packs", "/v1/reference-architecture/valc/failure-taxonomy", "/v1/reference-architecture/valc/scenario-descriptors", "/v1/reference-architecture/valc/degraded-modes", "/v1/reference-architecture/valc/recovery-expectations", "/v1/reference-architecture/valc/scaling-scenarios", "/v1/reference-architecture/valc/trust-path", "/v1/reference-architecture/valc/audit-path", "/v1/reference-architecture/valc/control-plane-safety", "   "}},
	}
	for _, tc := range testCases {
		if got := EvaluateReferenceArchitectureValCProofsState(baseState, point6State, supportedFamilies, tc.surfaceRefs, evidenceRefs, limitations, disclaimer); got == ReferenceArchitectureValCStateActive {
			t.Fatalf("expected proofs surface test %q to fail closed, got %q", tc.name, got)
		}
	}
}

func TestReferenceArchitectureValCNoOverclaimAndPoint6NotComplete(t *testing.T) {
	point5State, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, point6State := activeReferenceArchitectureValCPrereqs()
	registry, taxonomy, descriptors, degraded, recovery, scaling, trust, audit, control := activeReferenceArchitectureValCComponents()
	valCState := EvaluateReferenceArchitectureValCState(
		point5State,
		val0CurrentState,
		val0State,
		valACurrentState,
		valAState,
		valBCurrentState,
		valBState,
		point6State,
		EvaluateReferenceArchitectureValCScenarioPackRegistryState(registry),
		EvaluateReferenceArchitectureValCFailureModeTaxonomyState(taxonomy),
		EvaluateReferenceArchitectureValCScenarioDescriptorCollectionState(descriptors),
		EvaluateReferenceArchitectureValCDegradedModeCollectionState(degraded),
		EvaluateReferenceArchitectureValCRecoveryExpectationCollectionState(recovery),
		EvaluateReferenceArchitectureValCScalingScenarioCollectionState(scaling),
		EvaluateReferenceArchitectureValCTrustPathCollectionState(trust),
		EvaluateReferenceArchitectureValCAuditPathCollectionState(audit),
		EvaluateReferenceArchitectureValCControlPlaneCollectionState(control),
	)
	if valCState != ReferenceArchitectureValCStateActive {
		t.Fatalf("expected active Val C state with valid fixtures, got %q", valCState)
	}
	if got := EvaluateReferenceArchitectureValCProofsState(
		valCState,
		ReferenceArchitecturePoint6StatePass,
		referenceArchitectureVal0Families(),
		referenceArchitectureValCProofSurfaceRefs(),
		[]string{"point5", "val0", "vala", "valb", "registry", "taxonomy", "descriptors", "degraded", "recovery", "scaling", "trust", "audit"},
		[]string{"Val C keeps point 6 not complete."},
		referenceArchitectureValCProjectionDisclaimer(),
	); got == ReferenceArchitectureValCStateActive {
		t.Fatalf("expected non-active proofs state when point 6 pass is claimed in Val C, got %q", got)
	}
}
