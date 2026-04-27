package operability

import "testing"

func activeVerifierEcosystemValAResult() VerifierEcosystemValAVerificationResult {
	input := VerifierEcosystemValAReferenceVerifierInputModel()
	return VerifierEcosystemValAVerificationResult{
		CurrentState:           "verifier_ecosystem_vala_result_ready",
		VerificationResultID:   "reference-verifier-result-" + input.VerificationRequestID,
		RequestID:              input.VerificationRequestID,
		VerifierVersion:        "reference-verifier/vala-2026.04",
		ProofType:              input.ProofType,
		SchemaVersion:          input.SchemaVersion,
		Scope:                  input.RequestedScope,
		OutputBoundary:         input.ExpectedOutputBoundary,
		OverallResult:          VerifierEcosystemValAOverallResultVerified,
		DiagnosticClass:        VerifierEcosystemDiagnosticVerified,
		DigestResult:           input.DigestVerificationState,
		SignatureResult:        input.SignatureVerificationState,
		SchemaResult:           input.SchemaVerificationState,
		ScopeResult:            input.ScopeVerificationState,
		FreshnessResult:        input.FreshnessVerificationState,
		TrustRootResult:        input.TrustRootVerificationState,
		IssuerResult:           input.IssuerVerificationState,
		CompatibilityResult:    input.CompatibilityEvaluationState,
		RevocationResult:       input.RevocationEvaluationState,
		SupersessionResult:     input.SupersessionEvaluationState,
		LineageResult:          input.LineageVerificationState,
		OutputBoundaryResult:   input.OutputBoundaryVerificationState,
		EvidenceRefs:           []string{"evidence:verifier-input-vala-001", "evidence:verifier-engine-vala-001", "evidence:verifier-report-vala-001"},
		Caveats:                input.Caveats,
		Limitations:            []string{"Val A verification reports remain advisory and scope-bounded."},
		ProjectionDisclaimer:   verifierEcosystemValAProjectionDisclaimer(),
		VerifiedAt:             input.VerificationTime,
		TruthOutsideScopeClaim: false,
	}
}

func activeVerifierEcosystemValAInputs() (
	VerifierEcosystemValADependencySnapshot,
	VerifierEcosystemValAReferenceVerifierInput,
	VerifierEcosystemValAReferenceVerifierEngine,
	VerifierEcosystemValAVerificationResult,
	VerifierEcosystemValADiagnosticsMapping,
	VerifierEcosystemValACommandContract,
	VerifierEcosystemValASDKEntrypoint,
) {
	_, _, _, _, _, _, _, _, contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState, val0State := activeVerifierEcosystemVal0States()
	dependency := VerifierEcosystemValADependencySnapshot{
		Point5State:                    IntelligenceCalibrationPoint5StatePass,
		Point5DependencyState:          IntelligenceCalibrationValEStateActive,
		Point6State:                    ReferenceArchitecturePoint6StatePass,
		Point6ClosureState:             ReferenceArchitectureValEStateActive,
		Point6ClosurePrerequisiteState: ReferenceArchitectureValEPrerequisiteStateActive,
		Point6ClosureInvariantState:    ReferenceArchitectureValEInvariantStateActive,
		Point6ProofSurfaceState:        ReferenceArchitectureValEProofSurfaceStateActive,
		Point6PassRuleState:            ReferenceArchitectureValEPassRuleStateActive,
		Point6PassAllowed:              true,
		Val0CurrentState:               VerifierEcosystemVal0StateActive,
		Val0State:                      val0State,
		Point7State:                    VerifierEcosystemPoint7StateNotComplete,
	}
	input := VerifierEcosystemValAReferenceVerifierInputModel()
	engine := VerifierEcosystemValAReferenceVerifierEngineModel(contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState)
	result := activeVerifierEcosystemValAResult()
	diagnosticsMapping := VerifierEcosystemValADiagnosticsMappingModel(result)
	command := VerifierEcosystemValACommandContractModel()
	sdk := VerifierEcosystemValASDKEntrypointModel()
	return dependency, input, engine, result, diagnosticsMapping, command, sdk
}

func activeVerifierEcosystemValAStates() (
	VerifierEcosystemValADependencySnapshot,
	VerifierEcosystemValAReferenceVerifierInput,
	VerifierEcosystemValAReferenceVerifierEngine,
	VerifierEcosystemValAVerificationResult,
	VerifierEcosystemValADiagnosticsMapping,
	VerifierEcosystemValACommandContract,
	VerifierEcosystemValASDKEntrypoint,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
) {
	dependency, input, engine, result, diagnosticsMapping, command, sdk := activeVerifierEcosystemValAInputs()
	inputState := EvaluateVerifierEcosystemValAReferenceVerifierInputState(input)
	engineState := EvaluateVerifierEcosystemValAReferenceVerifierEngineState(engine)
	resultState := EvaluateVerifierEcosystemValAVerificationResultState(result)
	diagnosticsMappingState := EvaluateVerifierEcosystemValADiagnosticsMappingState(diagnosticsMapping)
	commandState := EvaluateVerifierEcosystemValACommandContractState(command)
	sdkState := EvaluateVerifierEcosystemValASDKEntrypointState(sdk)
	valAState := EvaluateVerifierEcosystemValAState(dependency, inputState, engineState, resultState, diagnosticsMappingState, commandState, sdkState)
	return dependency, input, engine, result, diagnosticsMapping, command, sdk, inputState, engineState, resultState, diagnosticsMappingState, commandState, sdkState, valAState
}

func TestVerifierEcosystemValADependencyGates(t *testing.T) {
	dependency, _, _, _, _, _, _, inputState, engineState, resultState, diagnosticsMappingState, commandState, sdkState, valAState := activeVerifierEcosystemValAStates()
	if valAState != VerifierEcosystemValAStateActive {
		t.Fatalf("expected active Val A state with active Val 0 and Točka 6 closure dependency, got %q", valAState)
	}
	if got := EvaluateVerifierEcosystemValAPoint7State(valAState); got != VerifierEcosystemPoint7StateNotComplete {
		t.Fatalf("expected point 7 to remain not complete in Val A, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValADependencySnapshot)
		expected string
	}{
		{name: "missing val0 blocks active", mutate: func(s *VerifierEcosystemValADependencySnapshot) {
			s.Val0State = VerifierEcosystemVal0StatePartial
		}, expected: VerifierEcosystemValAStateBlocked},
		{name: "missing point6 closure blocks active", mutate: func(s *VerifierEcosystemValADependencySnapshot) {
			s.Point6ClosureState = ReferenceArchitectureValEStatePartial
		}, expected: VerifierEcosystemValAStateBlocked},
		{name: "point7 state other than not complete blocks active", mutate: func(s *VerifierEcosystemValADependencySnapshot) {
			s.Point7State = VerifierEcosystemPoint7StatePass
		}, expected: VerifierEcosystemValAStateBlocked},
	}

	for _, tc := range testCases {
		snapshot := dependency
		tc.mutate(&snapshot)
		if got := EvaluateVerifierEcosystemValAState(snapshot, inputState, engineState, resultState, diagnosticsMappingState, commandState, sdkState); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValAAggregateStatePrecedence(t *testing.T) {
	dependency, _, _, _, _, _, _, inputState, engineState, resultState, diagnosticsMappingState, commandState, sdkState, _ := activeVerifierEcosystemValAStates()

	testCases := []struct {
		name        string
		input       string
		engine      string
		result      string
		diagnostics string
		command     string
		sdk         string
		expected    string
	}{
		{
			name:        "partial before blocked returns blocked",
			input:       VerifierEcosystemValAInputStatePartial,
			engine:      engineState,
			result:      resultState,
			diagnostics: diagnosticsMappingState,
			command:     VerifierEcosystemValACommandContractStateBlocked,
			sdk:         sdkState,
			expected:    VerifierEcosystemValAStateBlocked,
		},
		{
			name:        "blocked before partial returns blocked",
			input:       VerifierEcosystemValAInputStateBlocked,
			engine:      VerifierEcosystemValAEngineStatePartial,
			result:      resultState,
			diagnostics: diagnosticsMappingState,
			command:     commandState,
			sdk:         sdkState,
			expected:    VerifierEcosystemValAStateBlocked,
		},
		{
			name:        "partial plus unknown returns unknown",
			input:       VerifierEcosystemValAInputStatePartial,
			engine:      engineState,
			result:      VerifierEcosystemValAResultStateUnknown,
			diagnostics: diagnosticsMappingState,
			command:     commandState,
			sdk:         sdkState,
			expected:    VerifierEcosystemValAStateUnknown,
		},
		{
			name:        "incomplete plus partial returns incomplete",
			input:       inputState,
			engine:      VerifierEcosystemValAEngineStateIncomplete,
			result:      VerifierEcosystemValAResultStatePartial,
			diagnostics: diagnosticsMappingState,
			command:     commandState,
			sdk:         sdkState,
			expected:    VerifierEcosystemValAStateIncomplete,
		},
		{
			name:        "all active returns active",
			input:       inputState,
			engine:      engineState,
			result:      resultState,
			diagnostics: diagnosticsMappingState,
			command:     commandState,
			sdk:         sdkState,
			expected:    VerifierEcosystemValAStateActive,
		},
		{
			name:        "fake component state fails closed",
			input:       "not_a_real_vala_input_state",
			engine:      engineState,
			result:      resultState,
			diagnostics: diagnosticsMappingState,
			command:     commandState,
			sdk:         sdkState,
			expected:    VerifierEcosystemValAStateUnknown,
		},
	}

	for _, tc := range testCases {
		if got := EvaluateVerifierEcosystemValAState(
			dependency,
			tc.input,
			tc.engine,
			tc.result,
			tc.diagnostics,
			tc.command,
			tc.sdk,
		); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValAReferenceVerifierInputValidation(t *testing.T) {
	input := VerifierEcosystemValAReferenceVerifierInputModel()
	if got := EvaluateVerifierEcosystemValAReferenceVerifierInputState(input); got != VerifierEcosystemValAInputStateActive {
		t.Fatalf("expected valid reference verifier input to be active, got %q", got)
	}

	testCases := []struct {
		name   string
		mutate func(*VerifierEcosystemValAReferenceVerifierInput)
	}{
		{name: "missing verification request id", mutate: func(model *VerifierEcosystemValAReferenceVerifierInput) { model.VerificationRequestID = "" }},
		{name: "missing proof envelope ref", mutate: func(model *VerifierEcosystemValAReferenceVerifierInput) { model.ProofEnvelopeRef = "" }},
		{name: "missing artifact digest", mutate: func(model *VerifierEcosystemValAReferenceVerifierInput) { model.ArtifactDigest = "" }},
		{name: "unknown digest algorithm", mutate: func(model *VerifierEcosystemValAReferenceVerifierInput) { model.ArtifactDigestAlgorithm = "sha3" }},
		{name: "missing signature ref", mutate: func(model *VerifierEcosystemValAReferenceVerifierInput) { model.SignatureRef = "" }},
		{name: "missing issuer ref", mutate: func(model *VerifierEcosystemValAReferenceVerifierInput) { model.IssuerRef = "" }},
		{name: "missing trust root ref", mutate: func(model *VerifierEcosystemValAReferenceVerifierInput) { model.TrustRootRef = "" }},
		{name: "unknown requested scope", mutate: func(model *VerifierEcosystemValAReferenceVerifierInput) { model.RequestedScope = "global_public" }},
		{name: "missing projection disclaimer", mutate: func(model *VerifierEcosystemValAReferenceVerifierInput) { model.ProjectionDisclaimer = "" }},
	}

	for _, tc := range testCases {
		model := VerifierEcosystemValAReferenceVerifierInputModel()
		tc.mutate(&model)
		if got := EvaluateVerifierEcosystemValAReferenceVerifierInputState(model); got == VerifierEcosystemValAInputStateActive {
			t.Fatalf("expected %s to fail closed, got %q", tc.name, got)
		}
	}
}

func TestVerifierEcosystemValAReferenceVerifierEngineGuardrails(t *testing.T) {
	_, _, engine, _, _, _, _ := activeVerifierEcosystemValAInputs()

	testCases := []struct {
		name   string
		mutate func(*VerifierEcosystemValAReferenceVerifierEngine)
	}{
		{
			name: "mutates evidence blocks even after partial component",
			mutate: func(model *VerifierEcosystemValAReferenceVerifierEngine) {
				model.ProofEnvelopeState = VerifierEcosystemVal0EnvelopeStatePartial
				model.MutatesEvidence = true
			},
		},
		{
			name: "deployment approval blocks even after partial component",
			mutate: func(model *VerifierEcosystemValAReferenceVerifierEngine) {
				model.RequestedScopeState = VerifierEcosystemVal0ScopeStatePartial
				model.ClaimsDeploymentApproval = true
			},
		},
		{
			name: "canonical truth claim blocks",
			mutate: func(model *VerifierEcosystemValAReferenceVerifierEngine) {
				model.DiagnosticsState = VerifierEcosystemVal0DiagnosticsStatePartial
				model.ClaimsCanonicalTruth = true
			},
		},
		{
			name: "certification claim blocks",
			mutate: func(model *VerifierEcosystemValAReferenceVerifierEngine) {
				model.OutputBoundaryState = VerifierEcosystemVal0OutputBoundaryStatePartial
				model.ClaimsCertification = true
			},
		},
		{
			name: "network dependency blocks",
			mutate: func(model *VerifierEcosystemValAReferenceVerifierEngine) {
				model.VerifierContractState = VerifierEcosystemVal0ContractStatePartial
				model.NetworkDependency = true
			},
		},
		{
			name: "non deterministic output blocks",
			mutate: func(model *VerifierEcosystemValAReferenceVerifierEngine) {
				model.TrustRootIssuerState = VerifierEcosystemVal0TrustStatePartial
				model.DeterministicOutput = false
			},
		},
		{
			name: "missing explicit fixture semantics blocks",
			mutate: func(model *VerifierEcosystemValAReferenceVerifierEngine) {
				model.SchemaCompatibilityState = VerifierEcosystemVal0CompatibilityStatePartial
				model.ExplicitFixtureSemantics = false
			},
		},
		{
			name: "claims actual crypto validity without primitives blocks",
			mutate: func(model *VerifierEcosystemValAReferenceVerifierEngine) {
				model.ProofEnvelopeState = VerifierEcosystemVal0EnvelopeStatePartial
				model.ClaimsActualCryptoValidity = true
			},
		},
	}

	for _, tc := range testCases {
		model := engine
		tc.mutate(&model)
		if got := EvaluateVerifierEcosystemValAReferenceVerifierEngineState(model); got != VerifierEcosystemValAEngineStateBlocked {
			t.Fatalf("expected %s to return %q, got %q", tc.name, VerifierEcosystemValAEngineStateBlocked, got)
		}
	}

	model := engine
	model.VerifierContractState = "not_a_real_val0_component_state"
	if got := EvaluateVerifierEcosystemValAReferenceVerifierEngineState(model); got == VerifierEcosystemValAEngineStateActive {
		t.Fatalf("expected fake engine component state to fail closed, got %q", got)
	}
}

func TestVerifierEcosystemValAVerificationResultValidation(t *testing.T) {
	result := activeVerifierEcosystemValAResult()
	if got := EvaluateVerifierEcosystemValAVerificationResultState(result); got != VerifierEcosystemValAResultStateActive {
		t.Fatalf("expected fully valid verification result to be active, got %q", got)
	}

	testCases := []struct {
		name               string
		mutate             func(*VerifierEcosystemValAVerificationResult)
		expectedDiagnostic string
	}{
		{name: "invalid signature", mutate: func(model *VerifierEcosystemValAVerificationResult) {
			model.SignatureResult = VerifierEcosystemValASignatureInvalid
			model.DiagnosticClass = VerifierEcosystemDiagnosticInvalidSignature
			model.OverallResult = VerifierEcosystemValAOverallResultInvalid
		}, expectedDiagnostic: VerifierEcosystemDiagnosticInvalidSignature},
		{name: "digest mismatch", mutate: func(model *VerifierEcosystemValAVerificationResult) {
			model.DigestResult = VerifierEcosystemValADigestMismatch
			model.DiagnosticClass = VerifierEcosystemDiagnosticDigestMismatch
			model.OverallResult = VerifierEcosystemValAOverallResultInvalid
		}, expectedDiagnostic: VerifierEcosystemDiagnosticDigestMismatch},
		{name: "schema mismatch", mutate: func(model *VerifierEcosystemValAVerificationResult) {
			model.SchemaResult = VerifierEcosystemValASchemaMismatch
			model.DiagnosticClass = VerifierEcosystemDiagnosticSchemaMismatch
			model.OverallResult = VerifierEcosystemValAOverallResultInvalid
		}, expectedDiagnostic: VerifierEcosystemDiagnosticSchemaMismatch},
		{name: "unsupported proof type", mutate: func(model *VerifierEcosystemValAVerificationResult) {
			model.ProofType = "publisher_profile_v9"
			model.DiagnosticClass = VerifierEcosystemDiagnosticUnsupportedProofType
			model.OverallResult = VerifierEcosystemValAOverallResultUnsupported
		}, expectedDiagnostic: VerifierEcosystemDiagnosticUnsupportedProofType},
		{name: "stale artifact", mutate: func(model *VerifierEcosystemValAVerificationResult) {
			model.FreshnessResult = IntelligenceCalibrationFreshnessStale
			model.DiagnosticClass = VerifierEcosystemDiagnosticStaleArtifact
			model.OverallResult = VerifierEcosystemValAOverallResultStale
		}, expectedDiagnostic: VerifierEcosystemDiagnosticStaleArtifact},
		{name: "expired artifact", mutate: func(model *VerifierEcosystemValAVerificationResult) {
			model.FreshnessResult = IntelligenceCalibrationFreshnessExpired
			model.DiagnosticClass = VerifierEcosystemDiagnosticExpiredArtifact
			model.OverallResult = VerifierEcosystemValAOverallResultStale
		}, expectedDiagnostic: VerifierEcosystemDiagnosticExpiredArtifact},
		{name: "revoked issuer", mutate: func(model *VerifierEcosystemValAVerificationResult) {
			model.IssuerResult = VerifierEcosystemValAIssuerRevoked
			model.DiagnosticClass = VerifierEcosystemDiagnosticRevokedIssuer
			model.OverallResult = VerifierEcosystemValAOverallResultRevoked
		}, expectedDiagnostic: VerifierEcosystemDiagnosticRevokedIssuer},
		{name: "superseded proof", mutate: func(model *VerifierEcosystemValAVerificationResult) {
			model.SupersessionResult = VerifierEcosystemValASupersessionSuperseded
			model.DiagnosticClass = VerifierEcosystemDiagnosticSupersededProof
			model.OverallResult = VerifierEcosystemValAOverallResultSuperseded
		}, expectedDiagnostic: VerifierEcosystemDiagnosticSupersededProof},
		{name: "insufficient trust material", mutate: func(model *VerifierEcosystemValAVerificationResult) {
			model.TrustRootResult = VerifierEcosystemTrustRootUnsupported
			model.DiagnosticClass = VerifierEcosystemDiagnosticInsufficientTrustMaterial
			model.OverallResult = VerifierEcosystemValAOverallResultIncomplete
		}, expectedDiagnostic: VerifierEcosystemDiagnosticInsufficientTrustMaterial},
		{name: "unknown component state", mutate: func(model *VerifierEcosystemValAVerificationResult) {
			model.RevocationResult = VerifierEcosystemValARevocationUnknown
			model.DiagnosticClass = VerifierEcosystemDiagnosticUnknown
			model.OverallResult = VerifierEcosystemValAOverallResultUnknown
		}, expectedDiagnostic: VerifierEcosystemDiagnosticUnknown},
	}

	for _, tc := range testCases {
		model := activeVerifierEcosystemValAResult()
		tc.mutate(&model)
		if got := DeriveVerifierEcosystemValADiagnosticClass(model); got != tc.expectedDiagnostic {
			t.Fatalf("expected %s diagnostic %q, got %q", tc.name, tc.expectedDiagnostic, got)
		}
		if got := EvaluateVerifierEcosystemValAVerificationResultState(model); got == VerifierEcosystemValAResultStateActive {
			t.Fatalf("expected %s not to remain active, got %q", tc.name, got)
		}
	}
}

func TestVerifierEcosystemValADiagnosticPrecedence(t *testing.T) {
	result := activeVerifierEcosystemValAResult()
	result.SignatureResult = VerifierEcosystemValASignatureInvalid
	result.CompatibilityResult = ReferenceArchitectureCompatibilityCompatibleWithWarning
	if got := DeriveVerifierEcosystemValADiagnosticClass(result); got != VerifierEcosystemDiagnosticInvalidSignature {
		t.Fatalf("expected invalid_signature to win over compatibility warning, got %q", got)
	}

	result = activeVerifierEcosystemValAResult()
	result.DigestResult = VerifierEcosystemValADigestMismatch
	result.CompatibilityResult = ReferenceArchitectureCompatibilityCompatibleWithWarning
	if got := DeriveVerifierEcosystemValADiagnosticClass(result); got != VerifierEcosystemDiagnosticDigestMismatch {
		t.Fatalf("expected digest_mismatch to win over compatibility warning, got %q", got)
	}

	result = activeVerifierEcosystemValAResult()
	result.IssuerResult = VerifierEcosystemValAIssuerRevoked
	result.DiagnosticClass = VerifierEcosystemDiagnosticRevokedIssuer
	result.OverallResult = VerifierEcosystemValAOverallResultRevoked
	if got := EvaluateVerifierEcosystemValAVerificationResultState(result); got == VerifierEcosystemValAResultStateActive {
		t.Fatalf("expected revoked issuer to block verified result, got %q", got)
	}
}

func TestVerifierEcosystemValACommandContractAndSDKValidation(t *testing.T) {
	command := VerifierEcosystemValACommandContractModel()
	if got := EvaluateVerifierEcosystemValACommandContractState(command); got != VerifierEcosystemValACommandContractStateActive {
		t.Fatalf("expected valid command contract to be active, got %q", got)
	}

	command = VerifierEcosystemValACommandContractModel()
	command.ReportFormat = VerifierEcosystemValACommandReportFormatUnknown
	if got := EvaluateVerifierEcosystemValACommandContractState(command); got == VerifierEcosystemValACommandContractStateActive {
		t.Fatalf("expected unknown report format to fail closed, got %q", got)
	}

	command = VerifierEcosystemValACommandContractModel()
	command.RequestedScope = ""
	if got := EvaluateVerifierEcosystemValACommandContractState(command); got == VerifierEcosystemValACommandContractStateActive {
		t.Fatalf("expected missing requested scope to fail closed, got %q", got)
	}

	command = VerifierEcosystemValACommandContractModel()
	command.ApprovesDeployment = true
	if got := EvaluateVerifierEcosystemValACommandContractState(command); got == VerifierEcosystemValACommandContractStateActive {
		t.Fatalf("expected approving command contract to fail closed, got %q", got)
	}

	sdk := VerifierEcosystemValASDKEntrypointModel()
	if got := EvaluateVerifierEcosystemValASDKEntrypointState(sdk); got != VerifierEcosystemValASDKEntrypointStateActive {
		t.Fatalf("expected valid sdk entrypoint to be active, got %q", got)
	}

	sdk = VerifierEcosystemValASDKEntrypointModel()
	sdk.HiddenMainInstanceDependency = true
	if got := EvaluateVerifierEcosystemValASDKEntrypointState(sdk); got == VerifierEcosystemValASDKEntrypointStateActive {
		t.Fatalf("expected hidden main-instance dependency to fail closed, got %q", got)
	}
}

func TestVerifierEcosystemValAProofSurfaceCompletenessAndPoint7PassImpossibility(t *testing.T) {
	_, _, _, _, _, _, _, _, _, _, _, _, _, valAState := activeVerifierEcosystemValAStates()
	if got := EvaluateVerifierEcosystemValAProofsState(
		valAState,
		VerifierEcosystemPoint7StateNotComplete,
		VerifierEcosystemVal0StateActive,
		VerifierEcosystemValAProofSurfaceRefs(),
		[]string{"e1", "e2", "e3", "e4"},
		[]string{"Val A remains not complete."},
		[]string{"Val A cannot return point_7_pass."},
		verifierEcosystemValAProjectionDisclaimer(),
	); got != VerifierEcosystemValAStateActive {
		t.Fatalf("expected exact Val A surface set to stay active, got %q", got)
	}

	testCases := []struct {
		name     string
		surfaces []string
	}{
		{name: "missing required surface", surfaces: VerifierEcosystemValAProofSurfaceRefs()[1:]},
		{name: "duplicate surface does not compensate", surfaces: append(VerifierEcosystemValAProofSurfaceRefs()[1:], VerifierEcosystemValAProofSurfaceRefs()[1])},
		{name: "unknown extra surface", surfaces: append(VerifierEcosystemValAProofSurfaceRefs(), "/v1/verifier-ecosystem/vala/unknown")},
		{name: "whitespace surface", surfaces: append(VerifierEcosystemValAProofSurfaceRefs()[1:], " ")},
	}
	for _, tc := range testCases {
		if got := EvaluateVerifierEcosystemValAProofsState(
			VerifierEcosystemValAStateActive,
			VerifierEcosystemPoint7StateNotComplete,
			VerifierEcosystemVal0StateActive,
			tc.surfaces,
			[]string{"e1", "e2", "e3", "e4"},
			[]string{"Val A remains not complete."},
			[]string{"Val A cannot return point_7_pass."},
			verifierEcosystemValAProjectionDisclaimer(),
		); got == VerifierEcosystemValAStateActive {
			t.Fatalf("expected %s to fail closed, got %q", tc.name, got)
		}
	}

	if got := EvaluateVerifierEcosystemValAPoint7State(VerifierEcosystemValAStateActive); got != VerifierEcosystemPoint7StateNotComplete {
		t.Fatalf("expected point_7_pass to remain impossible in Val A, got %q", got)
	}
}

func TestVerifierEcosystemValANoOverclaimBehavior(t *testing.T) {
	result := activeVerifierEcosystemValAResult()
	result.TruthOutsideScopeClaim = true
	if got := EvaluateVerifierEcosystemValAVerificationResultState(result); got == VerifierEcosystemValAResultStateActive {
		t.Fatalf("expected truth-outside-scope claim to block active verification report, got %q", got)
	}
}
