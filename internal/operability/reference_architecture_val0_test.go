package operability

import "testing"

func TestReferenceArchitectureVal0ValidFamilyPassesBlueprintValidation(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	if got := EvaluateReferenceArchitectureVal0BlueprintDisciplineState(model); got != ReferenceArchitectureVal0BlueprintDisciplineStateActive {
		t.Fatalf("expected active blueprint discipline state for valid family, got %q", got)
	}
}

func TestReferenceArchitectureVal0UnknownFamilyFailsClosed(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.Family = "unknown_family"
	if got := EvaluateReferenceArchitectureVal0BlueprintDisciplineState(model); got == ReferenceArchitectureVal0BlueprintDisciplineStateActive {
		t.Fatalf("expected non-active blueprint discipline state for unknown family, got %q", got)
	}
}

func TestReferenceArchitectureVal0TypoLifecycleFailsClosed(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.LifecycleState = "acitve"
	if got := EvaluateReferenceArchitectureVal0BlueprintDisciplineState(model); got == ReferenceArchitectureVal0BlueprintDisciplineStateActive {
		t.Fatalf("expected non-active blueprint discipline state for typo lifecycle, got %q", got)
	}
}

func TestReferenceArchitectureVal0MatchedRequiresFieldsCapabilitiesAssumptionsAndEvidence(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected matched conformance state for valid reference blueprint, got %q", got)
	}

	model = ReferenceArchitectureVal0BlueprintContract()
	model.SupportAssumptions = nil
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected non-matched conformance state without support assumptions, got %q", got)
	}
}

func TestReferenceArchitectureVal0MissingRequiredCapabilityCannotReturnMatched(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.ObservedCapabilities = []string{
		ReferenceArchitectureCapabilitySigning,
		ReferenceArchitectureCapabilityAuditWriter,
		ReferenceArchitectureCapabilityEvidenceStorage,
		ReferenceArchitectureCapabilityPolicyDist,
		ReferenceArchitectureCapabilityVerifierAccess,
	}
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected non-matched conformance state without required recovery capability, got %q", got)
	}
}

func TestReferenceArchitectureVal0UnsupportedConditionReturnsUnsupported(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.TriggeredUnsupportedConditions = []string{ReferenceArchitectureUnsupportedTrustMismatch}
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformanceUnsupported {
		t.Fatalf("expected unsupported conformance state for triggered unsupported condition, got %q", got)
	}
}

func TestReferenceArchitectureVal0DegradedConditionReturnsDegraded(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.TriggeredDegradedConditions = []string{ReferenceArchitectureDegradedAuditPathReduced}
	model.DegradedReasons = []string{"audit path replication is reduced under the declared environment fit"}
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformanceDegraded {
		t.Fatalf("expected degraded conformance state for degraded condition, got %q", got)
	}
}

func TestReferenceArchitectureVal0UnknownStateDoesNotBecomeMatched(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.ConformanceState = "macthed"
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected non-matched conformance state for typo conformance enum, got %q", got)
	}
}

func TestReferenceArchitectureVal0IncompatibleTrustAnchorBlocksMatched(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.ObservedEnvironment.TrustAnchorMode = ReferenceArchitectureTrustCentralized
	if got := EvaluateReferenceArchitectureVal0EnvironmentFitState(model); got == ReferenceArchitectureVal0EnvironmentFitStateActive {
		t.Fatalf("expected non-active environment fit state for trust anchor mismatch, got %q", got)
	}
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected non-matched conformance state for trust anchor mismatch, got %q", got)
	}
}

func TestReferenceArchitectureVal0IncompatibleAuditPathBlocksMatched(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.ObservedEnvironment.AuditPathMode = ReferenceArchitectureAuditCentralized
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected non-matched conformance state for audit path mismatch, got %q", got)
	}
}

func TestReferenceArchitectureVal0AirGappedConnectivityMismatchBlocksMatched(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.TargetEnvironment.ConnectivityMode = ReferenceArchitectureConnectivityAirGapped
	model.ObservedEnvironment.ConnectivityMode = ReferenceArchitectureConnectivityRestricted
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformanceUnsupported {
		t.Fatalf("expected unsupported conformance state for air-gapped mismatch, got %q", got)
	}
}

func TestReferenceArchitectureVal0MissingSupportBoundaryBlocksMatched(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.SupportBoundaryRef = ""
	if got := EvaluateReferenceArchitectureVal0EnvironmentFitState(model); got == ReferenceArchitectureVal0EnvironmentFitStateActive {
		t.Fatalf("expected non-active environment fit state without support boundary, got %q", got)
	}
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected non-matched conformance state without support boundary, got %q", got)
	}
}

func TestReferenceArchitectureVal0MissingEvidenceFailsClosed(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.EvidenceRefs = nil
	if got := EvaluateReferenceArchitectureVal0EvidenceDisciplineState(model); got == ReferenceArchitectureVal0EvidenceDisciplineStateActive {
		t.Fatalf("expected non-active evidence discipline state without evidence refs, got %q", got)
	}
}

func TestReferenceArchitectureVal0MalformedTimestampFailsClosed(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.EvidenceRefs[0].Timestamp = "2026-04-26 08:00:00Z"
	if got := EvaluateReferenceArchitectureVal0EvidenceDisciplineState(model); got == ReferenceArchitectureVal0EvidenceDisciplineStateActive {
		t.Fatalf("expected non-active evidence discipline state for malformed timestamp, got %q", got)
	}
}

func TestReferenceArchitectureVal0StaleEvidenceCannotReturnMatched(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.EvidenceRefs[0].FreshnessState = IntelligenceCalibrationFreshnessStale
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformanceDrifted {
		t.Fatalf("expected drifted conformance state for stale evidence, got %q", got)
	}
}

func TestReferenceArchitectureVal0FreshEvidenceSupportsMatched(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	if got := EvaluateReferenceArchitectureVal0EvidenceDisciplineState(model); got != ReferenceArchitectureVal0EvidenceDisciplineStateActive {
		t.Fatalf("expected active evidence discipline state for fresh evidence, got %q", got)
	}
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected matched conformance state for fresh evidence, got %q", got)
	}
}

func TestReferenceArchitectureVal0DeprecatedBlueprintCannotReturnCleanMatched(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.LifecycleState = ReferenceArchitectureLifecycleDeprecated
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformancePartiallyMatched {
		t.Fatalf("expected partially matched conformance state for deprecated lifecycle, got %q", got)
	}
}

func TestReferenceArchitectureVal0DeprecatedCompatibilityCannotReturnCleanMatched(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.CompatibilityState = ReferenceArchitectureCompatibilityDeprecated
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformancePartiallyMatched {
		t.Fatalf("expected partially matched conformance state for deprecated compatibility, got %q", got)
	}
}

func TestReferenceArchitectureVal0SupersededBlueprintReturnsSupersededReference(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.LifecycleState = ReferenceArchitectureLifecycleSuperseded
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformanceSupersededRef {
		t.Fatalf("expected superseded_reference conformance state, got %q", got)
	}
}

func TestReferenceArchitectureVal0SupersededCompatibilityReturnsSupersededReference(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.CompatibilityState = ReferenceArchitectureCompatibilitySuperseded
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformanceSupersededRef {
		t.Fatalf("expected superseded_reference conformance state for superseded compatibility, got %q", got)
	}
}

func TestReferenceArchitectureVal0UnsupportedLifecycleFailsClosed(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.LifecycleState = ReferenceArchitectureLifecycleUnsupported
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformanceUnsupported {
		t.Fatalf("expected unsupported conformance state for unsupported lifecycle, got %q", got)
	}
}

func TestReferenceArchitectureVal0UnsupportedCompatibilityFailsClosed(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.CompatibilityState = ReferenceArchitectureCompatibilityUnsupported
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformanceUnsupported {
		t.Fatalf("expected unsupported conformance state for unsupported compatibility, got %q", got)
	}
}

func TestReferenceArchitectureVal0UnknownLifecycleAndCompatibilityFailClosed(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.LifecycleState = ReferenceArchitectureLifecycleUnknown
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformanceUnknown {
		t.Fatalf("expected unknown conformance state for unknown lifecycle, got %q", got)
	}

	model = ReferenceArchitectureVal0BlueprintContract()
	model.CompatibilityState = ReferenceArchitectureCompatibilityUnknown
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got != ReferenceArchitectureConformanceUnknown {
		t.Fatalf("expected unknown conformance state for unknown compatibility, got %q", got)
	}
}

func TestReferenceArchitectureVal0TypoCompatibilityFailsClosed(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.CompatibilityState = "compatibel"
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected non-matched conformance state for typo compatibility, got %q", got)
	}
}

func TestReferenceArchitectureVal0ConformanceOutputMustNotClaimCertification(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.CertifiedLanguagePresent = true
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected non-matched conformance state when certification language is present, got %q", got)
	}
}

func TestReferenceArchitectureVal0CaveatOmissionCannotTurnUnsupportedIntoMatched(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	model.RedactionKeepsCaveats = false
	if got := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected non-matched conformance state when redaction omits caveats, got %q", got)
	}
}

func TestReferenceArchitectureVal0Point6PassIsImpossibleInVal0(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	currentState := EvaluateReferenceArchitectureVal0ProofsState(
		IntelligenceCalibrationPoint5StatePass,
		ReferenceArchitectureVal0StateActive,
		EvaluateReferenceArchitectureVal0BlueprintDisciplineState(model),
		EvaluateReferenceArchitectureVal0TaxonomyState(model),
		EvaluateReferenceArchitectureVal0EnvironmentFitState(model),
		EvaluateReferenceArchitectureVal0EvidenceDisciplineState(model),
		EvaluateReferenceArchitectureVal0CompatibilityBaselineState(model),
		EvaluateReferenceArchitectureVal0ReferenceConformanceState(model),
		ReferenceArchitecturePoint6StatePass,
		model.SupportedFamilies,
		model.SupportedConformanceStates,
		model.SupportedCompatibilityStates,
		[]string{
			"/v1/reference-architecture/val0/blueprint-discipline",
			"/v1/reference-architecture/val0/environment-fit",
			"/v1/reference-architecture/val0/conformance-evidence",
			"/v1/reference-architecture/val0/compatibility-baseline",
			"/v1/reference-architecture/val0/proofs",
		},
		[]string{"point5_integrated_closure", "evidence_spine", "evidence:1", "evidence:2", "evidence:3", "evidence:4"},
		[]string{"Val 0 remains not complete for Točka 6."},
		"projection_only not_canonical_truth bounded_reference_architecture_foundation",
	)
	if currentState == ReferenceArchitectureVal0StateActive {
		t.Fatalf("expected non-active proofs state when point_6_pass is claimed in Val 0, got %q", currentState)
	}
}

func TestReferenceArchitectureVal0StateRequiresPoint5PassAndKeepsPoint6NotComplete(t *testing.T) {
	model := ReferenceArchitectureVal0BlueprintContract()
	blueprintDisciplineState := EvaluateReferenceArchitectureVal0BlueprintDisciplineState(model)
	taxonomyState := EvaluateReferenceArchitectureVal0TaxonomyState(model)
	environmentFitState := EvaluateReferenceArchitectureVal0EnvironmentFitState(model)
	evidenceDisciplineState := EvaluateReferenceArchitectureVal0EvidenceDisciplineState(model)
	compatibilityBaselineState := EvaluateReferenceArchitectureVal0CompatibilityBaselineState(model)
	conformanceState := EvaluateReferenceArchitectureVal0ReferenceConformanceState(model)

	if got := EvaluateReferenceArchitectureVal0State(
		IntelligenceCalibrationPoint5StateNotComplete,
		blueprintDisciplineState,
		taxonomyState,
		environmentFitState,
		evidenceDisciplineState,
		compatibilityBaselineState,
		conformanceState,
	); got == ReferenceArchitectureVal0StateActive {
		t.Fatalf("expected non-active Val 0 state without point 5 pass dependency, got %q", got)
	}

	if got := EvaluateReferenceArchitectureVal0State(
		IntelligenceCalibrationPoint5StatePass,
		blueprintDisciplineState,
		taxonomyState,
		environmentFitState,
		evidenceDisciplineState,
		compatibilityBaselineState,
		conformanceState,
	); got != ReferenceArchitectureVal0StateActive {
		t.Fatalf("expected active Val 0 state when point 5 passes and all reference checks pass, got %q", got)
	}
}
