package operability

import "testing"

func activeVerifierEcosystemValDInputs() (
	VerifierEcosystemValDDependencySnapshot,
	VerifierEcosystemValDCorrectnessGate,
	VerifierEcosystemValDToolingGate,
	VerifierEcosystemValDSchemaCompatibilityGate,
	VerifierEcosystemValDDiagnosticsConformanceGate,
	VerifierEcosystemValDTrustKeyRotationGate,
	VerifierEcosystemValDNegativeDiagnosticsGate,
	VerifierEcosystemValDRedactionGate,
	VerifierEcosystemValDPublisherArtifactGate,
	VerifierEcosystemValDNoOverclaimGate,
) {
	valCDependency, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, valCState := activeVerifierEcosystemValCStates()
	dependency := VerifierEcosystemValDDependencySnapshot{
		Point5State:                    valCDependency.Point5State,
		Point5DependencyState:          valCDependency.Point5DependencyState,
		Point6State:                    valCDependency.Point6State,
		Point6ClosureState:             valCDependency.Point6ClosureState,
		Point6ClosurePrerequisiteState: valCDependency.Point6ClosurePrerequisiteState,
		Point6ClosureInvariantState:    valCDependency.Point6ClosureInvariantState,
		Point6ProofSurfaceState:        valCDependency.Point6ProofSurfaceState,
		Point6PassRuleState:            valCDependency.Point6PassRuleState,
		Point6PassAllowed:              valCDependency.Point6PassAllowed,
		Val0CurrentState:               valCDependency.Val0CurrentState,
		Val0State:                      valCDependency.Val0State,
		ValACurrentState:               valCDependency.ValACurrentState,
		ValAState:                      valCDependency.ValAState,
		ValBCurrentState:               valCDependency.ValBCurrentState,
		ValBState:                      valCDependency.ValBState,
		ValCCurrentState:               VerifierEcosystemValCStateActive,
		ValCState:                      valCState,
		Point7State:                    valCDependency.Point7State,
	}
	return dependency,
		VerifierEcosystemValDCorrectnessGateModel(),
		VerifierEcosystemValDToolingGateModel(),
		VerifierEcosystemValDSchemaCompatibilityGateModel(),
		VerifierEcosystemValDDiagnosticsConformanceGateModel(),
		VerifierEcosystemValDTrustKeyRotationGateModel(),
		VerifierEcosystemValDNegativeDiagnosticsGateModel(),
		VerifierEcosystemValDRedactionGateModel(),
		VerifierEcosystemValDPublisherArtifactGateModel(),
		VerifierEcosystemValDNoOverclaimGateModel()
}

func activeVerifierEcosystemValDStates() (
	VerifierEcosystemValDDependencySnapshot,
	VerifierEcosystemValDCorrectnessGate,
	VerifierEcosystemValDToolingGate,
	VerifierEcosystemValDSchemaCompatibilityGate,
	VerifierEcosystemValDDiagnosticsConformanceGate,
	VerifierEcosystemValDTrustKeyRotationGate,
	VerifierEcosystemValDNegativeDiagnosticsGate,
	VerifierEcosystemValDRedactionGate,
	VerifierEcosystemValDPublisherArtifactGate,
	VerifierEcosystemValDNoOverclaimGate,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
) {
	dependency, correctness, tooling, schemaCompatibility, diagnosticsConformance, trustKeyRotation, negativeDiagnostics, redaction, publisherArtifact, noOverclaim := activeVerifierEcosystemValDInputs()
	correctnessState := EvaluateVerifierEcosystemValDCorrectnessGateState(correctness)
	toolingState := EvaluateVerifierEcosystemValDToolingGateState(tooling)
	schemaCompatibilityState := EvaluateVerifierEcosystemValDSchemaCompatibilityGateState(schemaCompatibility)
	diagnosticsConformanceState := EvaluateVerifierEcosystemValDDiagnosticsConformanceGateState(diagnosticsConformance)
	trustKeyRotationState := EvaluateVerifierEcosystemValDTrustKeyRotationGateState(trustKeyRotation)
	negativeDiagnosticsState := EvaluateVerifierEcosystemValDNegativeDiagnosticsGateState(negativeDiagnostics)
	redactionState := EvaluateVerifierEcosystemValDRedactionGateState(redaction)
	publisherArtifactState := EvaluateVerifierEcosystemValDPublisherArtifactGateState(publisherArtifact)
	noOverclaimState := EvaluateVerifierEcosystemValDNoOverclaimGateState(noOverclaim)
	valDState := EvaluateVerifierEcosystemValDState(
		dependency,
		correctnessState,
		toolingState,
		schemaCompatibilityState,
		diagnosticsConformanceState,
		trustKeyRotationState,
		negativeDiagnosticsState,
		redactionState,
		publisherArtifactState,
		noOverclaimState,
	)
	return dependency, correctness, tooling, schemaCompatibility, diagnosticsConformance, trustKeyRotation, negativeDiagnostics, redaction, publisherArtifact, noOverclaim, correctnessState, toolingState, schemaCompatibilityState, diagnosticsConformanceState, trustKeyRotationState, negativeDiagnosticsState, redactionState, publisherArtifactState, noOverclaimState, valDState
}

func TestVerifierEcosystemValDDependencyGates(t *testing.T) {
	dependency, _, _, _, _, _, _, _, _, _, correctnessState, toolingState, schemaCompatibilityState, diagnosticsConformanceState, trustKeyRotationState, negativeDiagnosticsState, redactionState, publisherArtifactState, noOverclaimState, valDState := activeVerifierEcosystemValDStates()
	if valDState != VerifierEcosystemValDStateActive {
		t.Fatalf("expected active Val D state with active Točka 6, Val 0, Val A, Val B, and Val C dependencies, got %q", valDState)
	}
	if got := EvaluateVerifierEcosystemValDPoint7State(valDState); got != VerifierEcosystemPoint7StateNotComplete {
		t.Fatalf("expected point 7 to remain not complete in Val D, got %q", got)
	}

	testCases := []struct {
		name   string
		mutate func(*VerifierEcosystemValDDependencySnapshot)
	}{
		{name: "missing val0 blocks active", mutate: func(model *VerifierEcosystemValDDependencySnapshot) {
			model.Val0State = VerifierEcosystemVal0StatePartial
		}},
		{name: "missing vala blocks active", mutate: func(model *VerifierEcosystemValDDependencySnapshot) {
			model.ValAState = VerifierEcosystemValAStatePartial
		}},
		{name: "missing valb blocks active", mutate: func(model *VerifierEcosystemValDDependencySnapshot) {
			model.ValBState = VerifierEcosystemValBStatePartial
		}},
		{name: "missing valc blocks active", mutate: func(model *VerifierEcosystemValDDependencySnapshot) {
			model.ValCState = VerifierEcosystemValCStatePartial
		}},
		{name: "missing point6 closure blocks active", mutate: func(model *VerifierEcosystemValDDependencySnapshot) {
			model.Point6ClosureState = ReferenceArchitectureValEStatePartial
		}},
		{name: "point7 other than not complete blocks active", mutate: func(model *VerifierEcosystemValDDependencySnapshot) {
			model.Point7State = VerifierEcosystemPoint7StatePass
		}},
	}

	for _, tc := range testCases {
		snapshot := dependency
		tc.mutate(&snapshot)
		if got := EvaluateVerifierEcosystemValDState(snapshot, correctnessState, toolingState, schemaCompatibilityState, diagnosticsConformanceState, trustKeyRotationState, negativeDiagnosticsState, redactionState, publisherArtifactState, noOverclaimState); got != VerifierEcosystemValDStateBlocked {
			t.Fatalf("expected %s to return %q, got %q", tc.name, VerifierEcosystemValDStateBlocked, got)
		}
	}
}

func TestVerifierEcosystemValDCorrectnessGate(t *testing.T) {
	model := VerifierEcosystemValDCorrectnessGateModel()
	if got := EvaluateVerifierEcosystemValDCorrectnessGateState(model); got != VerifierEcosystemValDCorrectnessGateStateActive {
		t.Fatalf("expected valid correctness gate to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValDCorrectnessGate)
		expected string
	}{
		{name: "fake active like component fails closed", mutate: func(model *VerifierEcosystemValDCorrectnessGate) {
			model.SourceValStates[0] = "verifier_ecosystem_val0_active_like"
		}, expected: VerifierEcosystemValDCorrectnessGateStateUnknown},
		{name: "suffix active string is not accepted", mutate: func(model *VerifierEcosystemValDCorrectnessGate) {
			model.ReferenceEngineState = "engine_active"
		}, expected: VerifierEcosystemValDCorrectnessGateStateUnknown},
		{name: "hard policy violation blocks active", mutate: func(model *VerifierEcosystemValDCorrectnessGate) {
			model.ApprovalClaim = true
		}, expected: VerifierEcosystemValDCorrectnessGateStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValDCorrectnessGateModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValDCorrectnessGateState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValDToolingGate(t *testing.T) {
	model := VerifierEcosystemValDToolingGateModel()
	if got := EvaluateVerifierEcosystemValDToolingGateState(model); got != VerifierEcosystemValDToolingGateStateActive {
		t.Fatalf("expected valid tooling gate to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValDToolingGate)
		expected string
	}{
		{name: "engine mutation claim blocks active", mutate: func(model *VerifierEcosystemValDToolingGate) { model.MutatesEvidence = true }, expected: VerifierEcosystemValDToolingGateStateBlocked},
		{name: "deployment approval claim blocks active", mutate: func(model *VerifierEcosystemValDToolingGate) { model.ApprovesDeployment = true }, expected: VerifierEcosystemValDToolingGateStateBlocked},
		{name: "fake crypto validity without primitives blocks active", mutate: func(model *VerifierEcosystemValDToolingGate) { model.ClaimsRealCryptoWithoutPrimitive = true }, expected: VerifierEcosystemValDToolingGateStateBlocked},
		{name: "hidden instance dependency blocks active", mutate: func(model *VerifierEcosystemValDToolingGate) { model.HiddenMainInstanceDependency = true }, expected: VerifierEcosystemValDToolingGateStateBlocked},
		{name: "sdk entrypoint remains advisory", mutate: func(model *VerifierEcosystemValDToolingGate) {
			model.SDKEntrypointState = VerifierEcosystemValASDKEntrypointStatePartial
		}, expected: VerifierEcosystemValDToolingGateStatePartial},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValDToolingGateModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValDToolingGateState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValDSchemaCompatibilityGate(t *testing.T) {
	model := VerifierEcosystemValDSchemaCompatibilityGateModel()
	if got := EvaluateVerifierEcosystemValDSchemaCompatibilityGateState(model); got != VerifierEcosystemValDSchemaCompatibilityGateStateActive {
		t.Fatalf("expected valid schema compatibility gate to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValDSchemaCompatibilityGate)
		expected string
	}{
		{name: "unsupported schema blocks active", mutate: func(model *VerifierEcosystemValDSchemaCompatibilityGate) {
			model.CompatibilityState = ReferenceArchitectureCompatibilityUnsupported
		}, expected: VerifierEcosystemValDSchemaCompatibilityGateStateBlocked},
		{name: "unknown compatibility fails closed", mutate: func(model *VerifierEcosystemValDSchemaCompatibilityGate) {
			model.CompatibilityState = ReferenceArchitectureCompatibilityUnknown
		}, expected: VerifierEcosystemValDSchemaCompatibilityGateStateUnknown},
		{name: "deprecated without migration visibility fails closed", mutate: func(model *VerifierEcosystemValDSchemaCompatibilityGate) {
			model.CompatibilityState = ReferenceArchitectureCompatibilityDeprecated
			model.DeprecatedMigrationVisible = false
		}, expected: VerifierEcosystemValDSchemaCompatibilityGateStateBlocked},
		{name: "compatible with warnings without caveat fails closed", mutate: func(model *VerifierEcosystemValDSchemaCompatibilityGate) {
			model.CompatibilityState = ReferenceArchitectureCompatibilityCompatibleWithWarning
			model.Caveats = nil
		}, expected: VerifierEcosystemValDSchemaCompatibilityGateStateIncomplete},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValDSchemaCompatibilityGateModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValDSchemaCompatibilityGateState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValDDiagnosticsConformanceGate(t *testing.T) {
	model := VerifierEcosystemValDDiagnosticsConformanceGateModel()
	if got := EvaluateVerifierEcosystemValDDiagnosticsConformanceGateState(model); got != VerifierEcosystemValDDiagnosticsConformanceGateStateActive {
		t.Fatalf("expected valid diagnostics and conformance gate to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValDDiagnosticsConformanceGate)
		expected string
	}{
		{name: "invalid signature outranks warnings", mutate: func(model *VerifierEcosystemValDDiagnosticsConformanceGate) {
			model.ObservedDiagnostics = []string{VerifierEcosystemDiagnosticCompatibilityWarning, VerifierEcosystemDiagnosticInvalidSignature}
			model.DerivedDiagnosticClass = VerifierEcosystemDiagnosticCompatibilityWarning
		}, expected: VerifierEcosystemValDDiagnosticsConformanceGateStatePartial},
		{name: "digest mismatch outranks warnings", mutate: func(model *VerifierEcosystemValDDiagnosticsConformanceGate) {
			model.ObservedDiagnostics = []string{VerifierEcosystemDiagnosticCompatibilityWarning, VerifierEcosystemDiagnosticDigestMismatch}
			model.DerivedDiagnosticClass = VerifierEcosystemDiagnosticCompatibilityWarning
		}, expected: VerifierEcosystemValDDiagnosticsConformanceGateStatePartial},
		{name: "conformance suite certification claim blocks active", mutate: func(model *VerifierEcosystemValDDiagnosticsConformanceGate) {
			model.CertificationClaim = true
		}, expected: VerifierEcosystemValDDiagnosticsConformanceGateStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValDDiagnosticsConformanceGateModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValDDiagnosticsConformanceGateState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValDTrustKeyRotationGate(t *testing.T) {
	model := VerifierEcosystemValDTrustKeyRotationGateModel()
	if got := EvaluateVerifierEcosystemValDTrustKeyRotationGateState(model); got != VerifierEcosystemValDTrustKeyRotationGateStateActive {
		t.Fatalf("expected valid trust key rotation gate to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValDTrustKeyRotationGate)
		expected string
	}{
		{name: "revoked issuer blocks active", mutate: func(model *VerifierEcosystemValDTrustKeyRotationGate) {
			model.IssuerState = verifierEcosystemValDIssuerStateRevoked
		}, expected: VerifierEcosystemValDTrustKeyRotationGateStateBlocked},
		{name: "revoked trust root blocks active", mutate: func(model *VerifierEcosystemValDTrustKeyRotationGate) {
			model.TrustRootState = VerifierEcosystemTrustRootRevoked
		}, expected: VerifierEcosystemValDTrustKeyRotationGateStateBlocked},
		{name: "unknown trust root fails closed", mutate: func(model *VerifierEcosystemValDTrustKeyRotationGate) {
			model.TrustRootState = VerifierEcosystemTrustRootUnknown
		}, expected: VerifierEcosystemValDTrustKeyRotationGateStateUnknown},
		{name: "rollover without metadata blocks active", mutate: func(model *VerifierEcosystemValDTrustKeyRotationGate) {
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			model.RolloverMetadataRef = ""
		}, expected: VerifierEcosystemValDTrustKeyRotationGateStateBlocked},
		{name: "rollover metadata does not downgrade hard failure", mutate: func(model *VerifierEcosystemValDTrustKeyRotationGate) {
			model.TrustRootState = VerifierEcosystemTrustRootRevoked
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
		}, expected: VerifierEcosystemValDTrustKeyRotationGateStateBlocked},
		{name: "global key directory claim blocks active", mutate: func(model *VerifierEcosystemValDTrustKeyRotationGate) { model.GlobalKeyDirectoryClaim = true }, expected: VerifierEcosystemValDTrustKeyRotationGateStateBlocked},
		{name: "unknown distribution mode fails closed", mutate: func(model *VerifierEcosystemValDTrustKeyRotationGate) {
			model.TrustDistributionMode = VerifierEcosystemValCDistributionModeUnknown
		}, expected: VerifierEcosystemValDTrustKeyRotationGateStateUnknown},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValDTrustKeyRotationGateModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValDTrustKeyRotationGateState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValDNegativeDiagnosticsGate(t *testing.T) {
	model := VerifierEcosystemValDNegativeDiagnosticsGateModel()
	if got := EvaluateVerifierEcosystemValDNegativeDiagnosticsGateState(model); got != VerifierEcosystemValDNegativeDiagnosticsGateStateActive {
		t.Fatalf("expected valid negative diagnostics gate to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValDNegativeDiagnosticsGate)
		expected string
	}{
		{name: "stale artifact remains non verified", mutate: func(model *VerifierEcosystemValDNegativeDiagnosticsGate) {
			model.PublicOverallResult = VerifierEcosystemValAOverallResultStale
			model.PublicDiagnosticClass = VerifierEcosystemDiagnosticStaleArtifact
			model.PublicOutputClass = VerifierEcosystemValBOutputClassNonVerifiedStale
		}, expected: VerifierEcosystemValDNegativeDiagnosticsGateStateActive},
		{name: "redaction cannot convert non verified into verified", mutate: func(model *VerifierEcosystemValDNegativeDiagnosticsGate) {
			model.PublicOverallResult = VerifierEcosystemValAOverallResultInvalid
			model.PublicDiagnosticClass = VerifierEcosystemDiagnosticInvalidSignature
			model.PublicOutputClass = VerifierEcosystemValBOutputClassVerified
		}, expected: VerifierEcosystemValDNegativeDiagnosticsGateStateBlocked},
		{name: "auditor output preserves repeatability and evidence traceability", mutate: func(model *VerifierEcosystemValDNegativeDiagnosticsGate) {
			model.AuditorRepeatable = false
		}, expected: VerifierEcosystemValDNegativeDiagnosticsGateStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValDNegativeDiagnosticsGateModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValDNegativeDiagnosticsGateState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValDRedactionGate(t *testing.T) {
	model := VerifierEcosystemValDRedactionGateModel()
	if got := EvaluateVerifierEcosystemValDRedactionGateState(model); got != VerifierEcosystemValDRedactionGateStateActive {
		t.Fatalf("expected valid redaction gate to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValDRedactionGate)
		expected string
	}{
		{name: "public without redaction policy fails closed", mutate: func(model *VerifierEcosystemValDRedactionGate) { model.RedactionPolicyRef = "" }, expected: VerifierEcosystemValDRedactionGateStateIncomplete},
		{name: "partner public breadth validation stays non active when invalid", mutate: func(model *VerifierEcosystemValDRedactionGate) { model.PartnerBroaderThanPublic = false }, expected: VerifierEcosystemValDRedactionGateStatePartial},
		{name: "internal diagnostic reused as public blocks", mutate: func(model *VerifierEcosystemValDRedactionGate) { model.InternalDiagnosticSeparated = false }, expected: VerifierEcosystemValDRedactionGateStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValDRedactionGateModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValDRedactionGateState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValDPublisherArtifactGate(t *testing.T) {
	model := VerifierEcosystemValDPublisherArtifactGateModel()
	if got := EvaluateVerifierEcosystemValDPublisherArtifactGateState(model); got != VerifierEcosystemValDPublisherArtifactGateStateActive {
		t.Fatalf("expected valid publisher artifact gate to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValDPublisherArtifactGate)
		expected string
	}{
		{name: "approved vendor claim blocks active", mutate: func(model *VerifierEcosystemValDPublisherArtifactGate) {
			model.ObservedClaims = []string{"approved vendor"}
		}, expected: VerifierEcosystemValDPublisherArtifactGateStateBlocked},
		{name: "certification claim blocks active", mutate: func(model *VerifierEcosystemValDPublisherArtifactGate) {
			model.ObservedClaims = []string{"certified publisher"}
		}, expected: VerifierEcosystemValDPublisherArtifactGateStateBlocked},
		{name: "unsupported proof type blocks active", mutate: func(model *VerifierEcosystemValDPublisherArtifactGate) {
			model.SupportedArtifactTypes = []string{"unknown"}
		}, expected: VerifierEcosystemValDPublisherArtifactGateStateUnknown},
		{name: "verifier compatible does not mean automatically trusted", mutate: func(model *VerifierEcosystemValDPublisherArtifactGate) { model.AutomaticallyTrustedClaim = true }, expected: VerifierEcosystemValDPublisherArtifactGateStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValDPublisherArtifactGateModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValDPublisherArtifactGateState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValDNoOverclaimGate(t *testing.T) {
	model := VerifierEcosystemValDNoOverclaimGateModel()
	if got := EvaluateVerifierEcosystemValDNoOverclaimGateState(model); got != VerifierEcosystemValDNoOverclaimGateStateActive {
		t.Fatalf("expected valid no-overclaim gate to be active, got %q", got)
	}

	for _, claim := range []string{
		"verifier certification",
		"integrity rating",
		"universal trust protocol",
		"global key registry for all instances",
		"deployment approved",
		"point_7_pass",
	} {
		mutated := VerifierEcosystemValDNoOverclaimGateModel()
		mutated.ObservedClaims = []string{claim}
		if got := EvaluateVerifierEcosystemValDNoOverclaimGateState(mutated); got != VerifierEcosystemValDNoOverclaimGateStateBlocked {
			t.Fatalf("expected claim %q to block active, got %q", claim, got)
		}
	}
}

func TestVerifierEcosystemValDAggregateStatePrecedence(t *testing.T) {
	dependency, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := activeVerifierEcosystemValDStates()
	testCases := []struct {
		name        string
		correctness string
		tooling     string
		expected    string
	}{
		{name: "partial before blocked returns blocked", correctness: VerifierEcosystemValDCorrectnessGateStatePartial, tooling: VerifierEcosystemValDToolingGateStateBlocked, expected: VerifierEcosystemValDStateBlocked},
		{name: "blocked before partial returns blocked", correctness: VerifierEcosystemValDCorrectnessGateStateBlocked, tooling: VerifierEcosystemValDToolingGateStatePartial, expected: VerifierEcosystemValDStateBlocked},
		{name: "partial plus unknown returns unknown", correctness: VerifierEcosystemValDCorrectnessGateStatePartial, tooling: VerifierEcosystemValDToolingGateStateUnknown, expected: VerifierEcosystemValDStateUnknown},
		{name: "incomplete plus partial returns incomplete", correctness: VerifierEcosystemValDCorrectnessGateStateIncomplete, tooling: VerifierEcosystemValDToolingGateStatePartial, expected: VerifierEcosystemValDStateIncomplete},
		{name: "all active returns active", correctness: VerifierEcosystemValDCorrectnessGateStateActive, tooling: VerifierEcosystemValDToolingGateStateActive, expected: VerifierEcosystemValDStateActive},
		{name: "fake component state fails closed", correctness: "fake_active", tooling: VerifierEcosystemValDToolingGateStateActive, expected: VerifierEcosystemValDStateUnknown},
	}
	for _, tc := range testCases {
		if got := EvaluateVerifierEcosystemValDState(
			dependency,
			tc.correctness,
			tc.tooling,
			VerifierEcosystemValDSchemaCompatibilityGateStateActive,
			VerifierEcosystemValDDiagnosticsConformanceGateStateActive,
			VerifierEcosystemValDTrustKeyRotationGateStateActive,
			VerifierEcosystemValDNegativeDiagnosticsGateStateActive,
			VerifierEcosystemValDRedactionGateStateActive,
			VerifierEcosystemValDPublisherArtifactGateStateActive,
			VerifierEcosystemValDNoOverclaimGateStateActive,
		); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValDProofSurfaceCompletenessAndNoFinalPass(t *testing.T) {
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, valDState := activeVerifierEcosystemValDStates()
	point7State := EvaluateVerifierEcosystemValDPoint7State(valDState)
	if point7State != VerifierEcosystemPoint7StateNotComplete {
		t.Fatalf("expected point 7 to remain not complete, got %q", point7State)
	}
	if got := EvaluateVerifierEcosystemValDProofsState(
		valDState,
		point7State,
		VerifierEcosystemVal0StateActive,
		VerifierEcosystemValAStateActive,
		VerifierEcosystemValBStateActive,
		VerifierEcosystemValCStateActive,
		VerifierEcosystemValDProofSurfaceRefs(),
		VerifierEcosystemValDProofEvidenceRefs(),
		[]string{"Val D does not close Točka 7."},
		[]string{"Val D cannot return point_7_pass."},
		verifierEcosystemValDProjectionDisclaimer(),
	); got != VerifierEcosystemValDStateActive {
		t.Fatalf("expected exact proof surface and evidence set to remain active, got %q", got)
	}

	testCases := []struct {
		name        string
		surfaceRefs []string
		expected    string
	}{
		{name: "missing required surface fails closed", surfaceRefs: VerifierEcosystemValDProofSurfaceRefs()[:len(VerifierEcosystemValDProofSurfaceRefs())-1], expected: VerifierEcosystemValDStatePartial},
		{name: "duplicate surface does not compensate for missing required surface", surfaceRefs: append(append([]string{}, VerifierEcosystemValDProofSurfaceRefs()[:len(VerifierEcosystemValDProofSurfaceRefs())-1]...), VerifierEcosystemValDProofSurfaceRefs()[0]), expected: VerifierEcosystemValDStatePartial},
		{name: "unknown extra surface does not compensate for missing required surface", surfaceRefs: append(append([]string{}, VerifierEcosystemValDProofSurfaceRefs()[:len(VerifierEcosystemValDProofSurfaceRefs())-1]...), "/v1/verifier-ecosystem/vald/extra"), expected: VerifierEcosystemValDStatePartial},
		{name: "whitespace surface ref fails closed", surfaceRefs: append([]string{}, append(VerifierEcosystemValDProofSurfaceRefs()[:len(VerifierEcosystemValDProofSurfaceRefs())-1], " ")...), expected: VerifierEcosystemValDStatePartial},
	}

	for _, tc := range testCases {
		if got := EvaluateVerifierEcosystemValDProofsState(
			valDState,
			point7State,
			VerifierEcosystemVal0StateActive,
			VerifierEcosystemValAStateActive,
			VerifierEcosystemValBStateActive,
			VerifierEcosystemValCStateActive,
			tc.surfaceRefs,
			VerifierEcosystemValDProofEvidenceRefs(),
			[]string{"Val D does not close Točka 7."},
			[]string{"Val D cannot return point_7_pass."},
			verifierEcosystemValDProjectionDisclaimer(),
		); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}
