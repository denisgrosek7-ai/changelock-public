package operability

import "testing"

func activeVerifierEcosystemVal0Inputs() (
	VerifierEcosystemVal0DependencySnapshot,
	VerifierEcosystemVal0VerifierContract,
	VerifierEcosystemVal0ProofEnvelope,
	VerifierEcosystemVal0VerificationScopeCatalog,
	VerifierEcosystemVal0SchemaCompatibilityBaseline,
	VerifierEcosystemVal0TrustIssuerDiscipline,
	VerifierEcosystemVal0DiagnosticsModel,
	VerifierEcosystemVal0OutputBoundaryCollection,
) {
	return VerifierEcosystemVal0DependencySnapshot{
			Point5State:                    IntelligenceCalibrationPoint5StatePass,
			Point5DependencyState:          IntelligenceCalibrationValEStateActive,
			Point6State:                    ReferenceArchitecturePoint6StatePass,
			Point6ClosureState:             ReferenceArchitectureValEStateActive,
			Point6ClosurePrerequisiteState: ReferenceArchitectureValEPrerequisiteStateActive,
			Point6ClosureInvariantState:    ReferenceArchitectureValEInvariantStateActive,
			Point6ProofSurfaceState:        ReferenceArchitectureValEProofSurfaceStateActive,
			Point6PassRuleState:            ReferenceArchitectureValEPassRuleStateActive,
			Point6PassAllowed:              true,
		},
		VerifierEcosystemVal0VerifierContractModel(),
		VerifierEcosystemVal0ProofEnvelopeModel(),
		VerifierEcosystemVal0VerificationScopeCatalogModel(),
		VerifierEcosystemVal0SchemaCompatibilityBaselineModel(),
		VerifierEcosystemVal0TrustIssuerDisciplineModel(),
		VerifierEcosystemVal0DiagnosticsCatalogModel(),
		VerifierEcosystemVal0OutputBoundaryCollectionModel()
}

func activeVerifierEcosystemVal0States() (
	VerifierEcosystemVal0DependencySnapshot,
	VerifierEcosystemVal0VerifierContract,
	VerifierEcosystemVal0ProofEnvelope,
	VerifierEcosystemVal0VerificationScopeCatalog,
	VerifierEcosystemVal0SchemaCompatibilityBaseline,
	VerifierEcosystemVal0TrustIssuerDiscipline,
	VerifierEcosystemVal0DiagnosticsModel,
	VerifierEcosystemVal0OutputBoundaryCollection,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
) {
	dependency, contract, envelope, scopeCatalog, compatibility, trust, diagnostics, outputBoundaries := activeVerifierEcosystemVal0Inputs()
	contractState := EvaluateVerifierEcosystemVal0VerifierContractState(contract)
	envelopeState := EvaluateVerifierEcosystemVal0ProofEnvelopeState(envelope)
	scopeState := EvaluateVerifierEcosystemVal0VerificationScopeState(scopeCatalog)
	compatibilityState := EvaluateVerifierEcosystemVal0SchemaCompatibilityBaselineState(compatibility)
	trustState := EvaluateVerifierEcosystemVal0TrustIssuerDisciplineState(trust)
	diagnosticsState := EvaluateVerifierEcosystemVal0DiagnosticsState(diagnostics)
	outputBoundaryState := EvaluateVerifierEcosystemVal0OutputBoundaryState(outputBoundaries)
	val0State := EvaluateVerifierEcosystemVal0State(
		dependency,
		contractState,
		envelopeState,
		scopeState,
		compatibilityState,
		trustState,
		diagnosticsState,
		outputBoundaryState,
	)
	return dependency, contract, envelope, scopeCatalog, compatibility, trust, diagnostics, outputBoundaries, contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState, val0State
}

func TestVerifierEcosystemVal0DependencyGates(t *testing.T) {
	dependency, _, _, _, _, _, _, _, contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState, val0State := activeVerifierEcosystemVal0States()
	if val0State != VerifierEcosystemVal0StateActive {
		t.Fatalf("expected active Val 0 state with active Točka 6 closure dependency, got %q", val0State)
	}
	if got := EvaluateVerifierEcosystemPoint7State(val0State); got != VerifierEcosystemPoint7StateNotComplete {
		t.Fatalf("expected point 7 to remain not complete in Val 0, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemVal0DependencySnapshot)
		expected string
	}{
		{name: "missing point 6 pass blocks active", mutate: func(s *VerifierEcosystemVal0DependencySnapshot) {
			s.Point6State = ReferenceArchitecturePoint6StateNotComplete
		}, expected: VerifierEcosystemVal0StateBlocked},
		{name: "point 5 not pass blocks active", mutate: func(s *VerifierEcosystemVal0DependencySnapshot) {
			s.Point5State = IntelligenceCalibrationPoint5StateNotComplete
		}, expected: VerifierEcosystemVal0StateBlocked},
		{name: "stale closure dependency blocks active", mutate: func(s *VerifierEcosystemVal0DependencySnapshot) {
			s.Point6ClosureState = ReferenceArchitectureValEStatePartial
		}, expected: VerifierEcosystemVal0StateBlocked},
		{name: "closure proof surface regression blocks active", mutate: func(s *VerifierEcosystemVal0DependencySnapshot) {
			s.Point6ProofSurfaceState = ReferenceArchitectureValEProofSurfaceStatePartial
		}, expected: VerifierEcosystemVal0StateBlocked},
	}

	for _, tc := range testCases {
		snapshot := dependency
		tc.mutate(&snapshot)
		if got := EvaluateVerifierEcosystemVal0State(snapshot, contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemVal0AggregateStatePrecedence(t *testing.T) {
	dependency, _, _, _, _, _, _, _, contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState, _ := activeVerifierEcosystemVal0States()

	testCases := []struct {
		name        string
		contract    string
		envelope    string
		scope       string
		compat      string
		trust       string
		diagnostics string
		output      string
		expected    string
	}{
		{
			name:        "partial before blocked returns blocked",
			contract:    VerifierEcosystemVal0ContractStatePartial,
			envelope:    envelopeState,
			scope:       scopeState,
			compat:      compatibilityState,
			trust:       trustState,
			diagnostics: VerifierEcosystemVal0DiagnosticsStateBlocked,
			output:      outputBoundaryState,
			expected:    VerifierEcosystemVal0StateBlocked,
		},
		{
			name:        "blocked before partial returns blocked",
			contract:    VerifierEcosystemVal0ContractStateBlocked,
			envelope:    envelopeState,
			scope:       VerifierEcosystemVal0ScopeStatePartial,
			compat:      compatibilityState,
			trust:       trustState,
			diagnostics: diagnosticsState,
			output:      outputBoundaryState,
			expected:    VerifierEcosystemVal0StateBlocked,
		},
		{
			name:        "partial plus unknown returns unknown",
			contract:    VerifierEcosystemVal0ContractStatePartial,
			envelope:    envelopeState,
			scope:       scopeState,
			compat:      VerifierEcosystemVal0CompatibilityStateUnknown,
			trust:       trustState,
			diagnostics: diagnosticsState,
			output:      outputBoundaryState,
			expected:    VerifierEcosystemVal0StateUnknown,
		},
		{
			name:        "incomplete plus partial returns incomplete",
			contract:    contractState,
			envelope:    VerifierEcosystemVal0EnvelopeStateIncomplete,
			scope:       VerifierEcosystemVal0ScopeStatePartial,
			compat:      compatibilityState,
			trust:       trustState,
			diagnostics: diagnosticsState,
			output:      outputBoundaryState,
			expected:    VerifierEcosystemVal0StateIncomplete,
		},
		{
			name:        "all active returns active",
			contract:    contractState,
			envelope:    envelopeState,
			scope:       scopeState,
			compat:      compatibilityState,
			trust:       trustState,
			diagnostics: diagnosticsState,
			output:      outputBoundaryState,
			expected:    VerifierEcosystemVal0StateActive,
		},
		{
			name:        "fake component state fails closed",
			contract:    "not_a_real_val0_contract_state",
			envelope:    envelopeState,
			scope:       scopeState,
			compat:      compatibilityState,
			trust:       trustState,
			diagnostics: diagnosticsState,
			output:      outputBoundaryState,
			expected:    VerifierEcosystemVal0StateUnknown,
		},
	}

	for _, tc := range testCases {
		if got := EvaluateVerifierEcosystemVal0State(
			dependency,
			tc.contract,
			tc.envelope,
			tc.scope,
			tc.compat,
			tc.trust,
			tc.diagnostics,
			tc.output,
		); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemVal0VerifierContractValidation(t *testing.T) {
	contract := VerifierEcosystemVal0VerifierContractModel()
	if got := EvaluateVerifierEcosystemVal0VerifierContractState(contract); got != VerifierEcosystemVal0ContractStateActive {
		t.Fatalf("expected valid verifier contract to be active, got %q", got)
	}

	contract = VerifierEcosystemVal0VerifierContractModel()
	contract.VerifierContractID = ""
	if got := EvaluateVerifierEcosystemVal0VerifierContractState(contract); got != VerifierEcosystemVal0ContractStateIncomplete {
		t.Fatalf("expected missing verifier_contract_id to fail closed, got %q", got)
	}

	contract = VerifierEcosystemVal0VerifierContractModel()
	contract.VerifierProfile = "global_public"
	if got := EvaluateVerifierEcosystemVal0VerifierContractState(contract); got == VerifierEcosystemVal0ContractStateActive {
		t.Fatalf("expected unknown verifier profile to fail closed, got %q", got)
	}

	contract = VerifierEcosystemVal0VerifierContractModel()
	contract.VerifierMode = VerifierEcosystemModeUnknown
	if got := EvaluateVerifierEcosystemVal0VerifierContractState(contract); got == VerifierEcosystemVal0ContractStateActive {
		t.Fatalf("expected unknown verifier mode to fail closed, got %q", got)
	}

	contract = VerifierEcosystemVal0VerifierContractModel()
	contract.LifecycleState = "acitve"
	if got := EvaluateVerifierEcosystemVal0VerifierContractState(contract); got == VerifierEcosystemVal0ContractStateActive {
		t.Fatalf("expected typo lifecycle to fail closed, got %q", got)
	}

	contract = VerifierEcosystemVal0VerifierContractModel()
	contract.CompatibilityState = ReferenceArchitectureCompatibilityUnsupported
	if got := EvaluateVerifierEcosystemVal0VerifierContractState(contract); got != VerifierEcosystemVal0ContractStateBlocked {
		t.Fatalf("expected unsupported compatibility to block active state, got %q", got)
	}

	contract = VerifierEcosystemVal0VerifierContractModel()
	contract.ProjectionDisclaimer = ""
	if got := EvaluateVerifierEcosystemVal0VerifierContractState(contract); got == VerifierEcosystemVal0ContractStateActive {
		t.Fatalf("expected missing projection disclaimer to fail closed, got %q", got)
	}

	contract = VerifierEcosystemVal0VerifierContractModel()
	contract.RegulatorApprovedClaim = true
	if got := EvaluateVerifierEcosystemVal0VerifierContractState(contract); got != VerifierEcosystemVal0ContractStateBlocked {
		t.Fatalf("expected overclaim language to block active state, got %q", got)
	}
}

func TestVerifierEcosystemVal0ProofEnvelopeBoundary(t *testing.T) {
	envelope := VerifierEcosystemVal0ProofEnvelopeModel()
	if got := EvaluateVerifierEcosystemVal0ProofEnvelopeState(envelope); got != VerifierEcosystemVal0EnvelopeStateActive {
		t.Fatalf("expected valid proof envelope boundary to be active, got %q", got)
	}

	testCases := []struct {
		name   string
		mutate func(*VerifierEcosystemVal0ProofEnvelope)
	}{
		{name: "missing schema version", mutate: func(model *VerifierEcosystemVal0ProofEnvelope) { model.SchemaVersion = "" }},
		{name: "missing artifact digest ref", mutate: func(model *VerifierEcosystemVal0ProofEnvelope) { model.ArtifactDigestRef = "" }},
		{name: "missing signature ref", mutate: func(model *VerifierEcosystemVal0ProofEnvelope) { model.SignatureRef = "" }},
		{name: "missing issuer ref", mutate: func(model *VerifierEcosystemVal0ProofEnvelope) { model.IssuerRef = "" }},
		{name: "missing trust root ref", mutate: func(model *VerifierEcosystemVal0ProofEnvelope) { model.TrustRootRef = "" }},
		{name: "unknown proof type", mutate: func(model *VerifierEcosystemVal0ProofEnvelope) { model.ProofType = "signature_bundle_v9" }},
		{name: "missing scope", mutate: func(model *VerifierEcosystemVal0ProofEnvelope) { model.Scope = "" }},
		{name: "malformed timestamp", mutate: func(model *VerifierEcosystemVal0ProofEnvelope) { model.IssuedAt = "2026-04-27 07:05:00Z" }},
		{name: "stale expiry", mutate: func(model *VerifierEcosystemVal0ProofEnvelope) { model.ExpiresAt = "2026-04-26T07:05:00Z" }},
		{name: "claims truth outside scope", mutate: func(model *VerifierEcosystemVal0ProofEnvelope) { model.ClaimsTruthOutsideScope = true }},
	}

	for _, tc := range testCases {
		model := VerifierEcosystemVal0ProofEnvelopeModel()
		tc.mutate(&model)
		if got := EvaluateVerifierEcosystemVal0ProofEnvelopeState(model); got == VerifierEcosystemVal0EnvelopeStateActive {
			t.Fatalf("expected %s to fail closed, got %q", tc.name, got)
		}
	}
}

func TestVerifierEcosystemVal0VerificationScopeValidation(t *testing.T) {
	catalog := VerifierEcosystemVal0VerificationScopeCatalogModel()
	if got := EvaluateVerifierEcosystemVal0VerificationScopeState(catalog); got != VerifierEcosystemVal0ScopeStateActive {
		t.Fatalf("expected supported verification scopes to be active, got %q", got)
	}

	catalog = VerifierEcosystemVal0VerificationScopeCatalogModel()
	catalog.Scopes[0].ScopeClass = VerifierEcosystemModeUnknown
	if got := EvaluateVerifierEcosystemVal0VerificationScopeState(catalog); got == VerifierEcosystemVal0ScopeStateActive {
		t.Fatalf("expected unknown scope to fail closed, got %q", got)
	}

	catalog = VerifierEcosystemVal0VerificationScopeCatalogModel()
	catalog.Scopes[0].RedactionAware = false
	if got := EvaluateVerifierEcosystemVal0VerificationScopeState(catalog); got == VerifierEcosystemVal0ScopeStateActive {
		t.Fatalf("expected public_safe without redaction discipline to fail closed, got %q", got)
	}

	catalog = VerifierEcosystemVal0VerificationScopeCatalogModel()
	catalog.Scopes[2].EvidenceTraceable = false
	if got := EvaluateVerifierEcosystemVal0VerificationScopeState(catalog); got == VerifierEcosystemVal0ScopeStateActive {
		t.Fatalf("expected auditor_safe without evidence traceability to fail closed, got %q", got)
	}

	catalog = VerifierEcosystemVal0VerificationScopeCatalogModel()
	catalog.Scopes[3].InternalOnly = false
	if got := EvaluateVerifierEcosystemVal0VerificationScopeState(catalog); got == VerifierEcosystemVal0ScopeStateActive {
		t.Fatalf("expected internal_diagnostic reused as public-safe to fail closed, got %q", got)
	}
}

func TestVerifierEcosystemVal0SchemaCompatibilityBaseline(t *testing.T) {
	baseline := VerifierEcosystemVal0SchemaCompatibilityBaselineModel()
	if got := EvaluateVerifierEcosystemVal0SchemaCompatibilityBaselineState(baseline); got != VerifierEcosystemVal0CompatibilityStateActive {
		t.Fatalf("expected compatible baseline to be active, got %q", got)
	}
	if len(baseline.MixedVersionDiagnostics) == 0 {
		t.Fatalf("expected mixed-version diagnostic representation in baseline, got %#v", baseline)
	}

	baseline = VerifierEcosystemVal0SchemaCompatibilityBaselineModel()
	baseline.CompatibilityState = ReferenceArchitectureCompatibilityCompatibleWithWarning
	if got := EvaluateVerifierEcosystemVal0SchemaCompatibilityBaselineState(baseline); got != VerifierEcosystemVal0CompatibilityStatePartial {
		t.Fatalf("expected compatible_with_warnings to require caveat handling, got %q", got)
	}

	baseline = VerifierEcosystemVal0SchemaCompatibilityBaselineModel()
	baseline.CompatibilityState = ReferenceArchitectureCompatibilityDeprecated
	if got := EvaluateVerifierEcosystemVal0SchemaCompatibilityBaselineState(baseline); got == VerifierEcosystemVal0CompatibilityStateActive {
		t.Fatalf("expected deprecated compatibility not to return clean verified state, got %q", got)
	}

	baseline = VerifierEcosystemVal0SchemaCompatibilityBaselineModel()
	baseline.CompatibilityState = ReferenceArchitectureCompatibilitySuperseded
	if got := EvaluateVerifierEcosystemVal0SchemaCompatibilityBaselineState(baseline); got == VerifierEcosystemVal0CompatibilityStateActive {
		t.Fatalf("expected superseded compatibility not to return clean verified state, got %q", got)
	}

	baseline = VerifierEcosystemVal0SchemaCompatibilityBaselineModel()
	baseline.CompatibilityState = ReferenceArchitectureCompatibilityUnsupported
	if got := EvaluateVerifierEcosystemVal0SchemaCompatibilityBaselineState(baseline); got != VerifierEcosystemVal0CompatibilityStateBlocked {
		t.Fatalf("expected unsupported compatibility to fail closed, got %q", got)
	}

	baseline = VerifierEcosystemVal0SchemaCompatibilityBaselineModel()
	baseline.CompatibilityState = ReferenceArchitectureCompatibilityUnknown
	if got := EvaluateVerifierEcosystemVal0SchemaCompatibilityBaselineState(baseline); got != VerifierEcosystemVal0CompatibilityStateUnknown {
		t.Fatalf("expected unknown compatibility to fail closed, got %q", got)
	}
}

func TestVerifierEcosystemVal0TrustRootIssuerDiscipline(t *testing.T) {
	model := VerifierEcosystemVal0TrustIssuerDisciplineModel()
	if got := EvaluateVerifierEcosystemVal0TrustIssuerDisciplineState(model); got != VerifierEcosystemVal0TrustStateActive {
		t.Fatalf("expected trusted trust-root to be active, got %q", got)
	}

	testCases := []struct {
		name   string
		mutate func(*VerifierEcosystemVal0TrustIssuerDiscipline)
	}{
		{name: "revoked trust root", mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
			model.TrustRootState = VerifierEcosystemTrustRootRevoked
		}},
		{name: "expired trust root", mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
			model.TrustRootState = VerifierEcosystemTrustRootExpired
		}},
		{name: "unsupported trust root", mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
			model.TrustRootState = VerifierEcosystemTrustRootUnsupported
		}},
		{name: "unknown trust root", mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
			model.TrustRootState = VerifierEcosystemTrustRootUnknown
		}},
		{name: "rotated without rollover metadata", mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
			model.TrustRootState = VerifierEcosystemTrustRootRotated
			model.RolloverMetadataRef = ""
		}},
		{name: "offline distribution unscoped", mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) { model.OfflineDistributionScope = "" }},
		{name: "global key directory claim", mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) { model.GlobalKeyDirectoryClaim = true }},
	}

	for _, tc := range testCases {
		model := VerifierEcosystemVal0TrustIssuerDisciplineModel()
		tc.mutate(&model)
		if got := EvaluateVerifierEcosystemVal0TrustIssuerDisciplineState(model); got == VerifierEcosystemVal0TrustStateActive {
			t.Fatalf("expected %s to fail closed, got %q", tc.name, got)
		}
	}

	model = VerifierEcosystemVal0TrustIssuerDisciplineModel()
	model.RevocationState = VerifierEcosystemRevocationRevoked
	if got := EvaluateVerifierEcosystemVal0TrustIssuerDisciplineState(model); got == VerifierEcosystemVal0TrustStateActive {
		t.Fatalf("expected revoked issuer not to be overridden by trusted trust-root state, got %q", got)
	}

	model = VerifierEcosystemVal0TrustIssuerDisciplineModel()
	model.RevocationState = VerifierEcosystemRevocationUnknown
	if got := EvaluateVerifierEcosystemVal0TrustIssuerDisciplineState(model); got == VerifierEcosystemVal0TrustStateActive {
		t.Fatalf("expected unknown revocation state to fail closed, got %q", got)
	}

	model = VerifierEcosystemVal0TrustIssuerDisciplineModel()
	model.RevocationState = "issuer_not_revked"
	if got := EvaluateVerifierEcosystemVal0TrustIssuerDisciplineState(model); got == VerifierEcosystemVal0TrustStateActive {
		t.Fatalf("expected typo revocation state to fail closed, got %q", got)
	}

	model = VerifierEcosystemVal0TrustIssuerDisciplineModel()
	model.KeyRotationState = VerifierEcosystemKeyRotationRollover
	model.RolloverMetadataRef = ""
	if got := EvaluateVerifierEcosystemVal0TrustIssuerDisciplineState(model); got == VerifierEcosystemVal0TrustStateActive {
		t.Fatalf("expected rollover without metadata to fail closed, got %q", got)
	}

	model = VerifierEcosystemVal0TrustIssuerDisciplineModel()
	model.KeyRotationState = VerifierEcosystemKeyRotationRollover
	if got := EvaluateVerifierEcosystemVal0TrustIssuerDisciplineState(model); got != VerifierEcosystemVal0TrustStatePartial {
		t.Fatalf("expected healthy rollover metadata to remain non-active but explicit, got %q", got)
	}

	precedenceCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemVal0TrustIssuerDiscipline)
		expected string
	}{
		{
			name: "revoked trust root with rollover metadata remains blocked",
			mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
				model.TrustRootState = VerifierEcosystemTrustRootRevoked
				model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			},
			expected: VerifierEcosystemVal0TrustStateBlocked,
		},
		{
			name: "expired trust root with rollover metadata remains blocked",
			mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
				model.TrustRootState = VerifierEcosystemTrustRootExpired
				model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			},
			expected: VerifierEcosystemVal0TrustStateBlocked,
		},
		{
			name: "unsupported trust root with rollover metadata remains blocked",
			mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
				model.TrustRootState = VerifierEcosystemTrustRootUnsupported
				model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			},
			expected: VerifierEcosystemVal0TrustStateBlocked,
		},
		{
			name: "unknown trust root with rollover metadata remains unknown",
			mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
				model.TrustRootState = VerifierEcosystemTrustRootUnknown
				model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			},
			expected: VerifierEcosystemVal0TrustStateUnknown,
		},
		{
			name: "typo trust root with rollover metadata fails closed",
			mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
				model.TrustRootState = "trustd"
				model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			},
			expected: VerifierEcosystemVal0TrustStateUnknown,
		},
		{
			name: "revoked issuer with rollover metadata remains blocked",
			mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
				model.RevocationState = VerifierEcosystemRevocationRevoked
				model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			},
			expected: VerifierEcosystemVal0TrustStateBlocked,
		},
		{
			name: "expired revocation with rollover metadata remains blocked",
			mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
				model.RevocationState = VerifierEcosystemRevocationExpired
				model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			},
			expected: VerifierEcosystemVal0TrustStateBlocked,
		},
		{
			name: "unsupported revocation with rollover metadata remains blocked",
			mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
				model.RevocationState = VerifierEcosystemRevocationUnsupported
				model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			},
			expected: VerifierEcosystemVal0TrustStateBlocked,
		},
		{
			name: "healthy rollover with metadata remains partial",
			mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
				model.TrustRootState = VerifierEcosystemTrustRootTrusted
				model.RevocationState = VerifierEcosystemRevocationNotRevoked
				model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			},
			expected: VerifierEcosystemVal0TrustStatePartial,
		},
		{
			name: "healthy rollover without metadata remains blocked",
			mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
				model.TrustRootState = VerifierEcosystemTrustRootTrusted
				model.RevocationState = VerifierEcosystemRevocationNotRevoked
				model.KeyRotationState = VerifierEcosystemKeyRotationRollover
				model.RolloverMetadataRef = ""
			},
			expected: VerifierEcosystemVal0TrustStateBlocked,
		},
		{
			name: "healthy non-rollover remains active",
			mutate: func(model *VerifierEcosystemVal0TrustIssuerDiscipline) {
				model.TrustRootState = VerifierEcosystemTrustRootTrusted
				model.RevocationState = VerifierEcosystemRevocationNotRevoked
				model.KeyRotationState = VerifierEcosystemKeyRotationCurrent
			},
			expected: VerifierEcosystemVal0TrustStateActive,
		},
	}

	for _, tc := range precedenceCases {
		model := VerifierEcosystemVal0TrustIssuerDisciplineModel()
		tc.mutate(&model)
		if got := EvaluateVerifierEcosystemVal0TrustIssuerDisciplineState(model); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemVal0DiagnosticsModel(t *testing.T) {
	model := VerifierEcosystemVal0DiagnosticsCatalogModel()
	if got := EvaluateVerifierEcosystemVal0DiagnosticsState(model); got != VerifierEcosystemVal0DiagnosticsStateActive {
		t.Fatalf("expected verified diagnostics to be active, got %q", got)
	}

	model = VerifierEcosystemVal0DiagnosticsCatalogModel()
	model.SupportedDiagnosticClasses = append(model.SupportedDiagnosticClasses[:0], model.SupportedDiagnosticClasses[1:]...)
	if got := EvaluateVerifierEcosystemVal0DiagnosticsState(model); got == VerifierEcosystemVal0DiagnosticsStateActive {
		t.Fatalf("expected missing diagnostic class to fail closed, got %q", got)
	}

	model = VerifierEcosystemVal0DiagnosticsCatalogModel()
	model.ObservedDiagnosticClass = VerifierEcosystemDiagnosticUnknown
	if got := EvaluateVerifierEcosystemVal0DiagnosticsState(model); got == VerifierEcosystemVal0DiagnosticsStateActive {
		t.Fatalf("expected unknown diagnostic class not to become verified, got %q", got)
	}

	model = VerifierEcosystemVal0DiagnosticsCatalogModel()
	model.ObservedDiagnosticClass = VerifierEcosystemDiagnosticStaleArtifact
	if got := EvaluateVerifierEcosystemVal0DiagnosticsState(model); got == VerifierEcosystemVal0DiagnosticsStateActive {
		t.Fatalf("expected stale artifact diagnostic not to become verified, got %q", got)
	}

	model = VerifierEcosystemVal0DiagnosticsCatalogModel()
	model.RedactionKeepsFailuresVisible = false
	if got := EvaluateVerifierEcosystemVal0DiagnosticsState(model); got == VerifierEcosystemVal0DiagnosticsStateActive {
		t.Fatalf("expected redaction hiding failures to fail closed, got %q", got)
	}
}

func TestVerifierEcosystemVal0OutputBoundaryDiscipline(t *testing.T) {
	collection := VerifierEcosystemVal0OutputBoundaryCollectionModel()
	if got := EvaluateVerifierEcosystemVal0OutputBoundaryState(collection); got != VerifierEcosystemVal0OutputBoundaryStateActive {
		t.Fatalf("expected valid output boundaries to be active, got %q", got)
	}

	collection = VerifierEcosystemVal0OutputBoundaryCollectionModel()
	if len(collection.Boundaries[0].RedactedFields) == 0 {
		t.Fatalf("expected public boundary test fixture to include redactions")
	}
	collection.Boundaries[0].RedactedFields = nil
	if got := EvaluateVerifierEcosystemVal0OutputBoundaryState(collection); got == VerifierEcosystemVal0OutputBoundaryStateActive {
		t.Fatalf("expected public output without redaction to fail closed, got %q", got)
	}

	collection = VerifierEcosystemVal0OutputBoundaryCollectionModel()
	collection.Boundaries[2].EvidenceRefPolicy = "public_caveated"
	if got := EvaluateVerifierEcosystemVal0OutputBoundaryState(collection); got == VerifierEcosystemVal0OutputBoundaryStateActive {
		t.Fatalf("expected auditor output without evidence traceability to fail closed, got %q", got)
	}

	collection = VerifierEcosystemVal0OutputBoundaryCollectionModel()
	collection.Boundaries[3].PublicReuseAllowed = true
	if got := EvaluateVerifierEcosystemVal0OutputBoundaryState(collection); got == VerifierEcosystemVal0OutputBoundaryStateActive {
		t.Fatalf("expected internal diagnostic output reused as public output to fail closed, got %q", got)
	}

	collection = VerifierEcosystemVal0OutputBoundaryCollectionModel()
	collection.Boundaries[1].RequiredCaveats = nil
	if got := EvaluateVerifierEcosystemVal0OutputBoundaryState(collection); got == VerifierEcosystemVal0OutputBoundaryStateActive {
		t.Fatalf("expected missing caveats to fail closed, got %q", got)
	}

	collection = VerifierEcosystemVal0OutputBoundaryCollectionModel()
	collection.Boundaries[2].PreservesInvalidDiagnostics = false
	if got := EvaluateVerifierEcosystemVal0OutputBoundaryState(collection); got == VerifierEcosystemVal0OutputBoundaryStateActive {
		t.Fatalf("expected output boundary not to convert invalid into verified, got %q", got)
	}
}

func TestVerifierEcosystemVal0NoOverclaimAndPoint7PassImpossibility(t *testing.T) {
	dependency, contract, envelope, scopeCatalog, compatibility, trust, diagnostics, outputBoundaries := activeVerifierEcosystemVal0Inputs()
	contract.CertifiedLanguagePresent = true
	contractState := EvaluateVerifierEcosystemVal0VerifierContractState(contract)
	val0State := EvaluateVerifierEcosystemVal0State(
		dependency,
		contractState,
		EvaluateVerifierEcosystemVal0ProofEnvelopeState(envelope),
		EvaluateVerifierEcosystemVal0VerificationScopeState(scopeCatalog),
		EvaluateVerifierEcosystemVal0SchemaCompatibilityBaselineState(compatibility),
		EvaluateVerifierEcosystemVal0TrustIssuerDisciplineState(trust),
		EvaluateVerifierEcosystemVal0DiagnosticsState(diagnostics),
		EvaluateVerifierEcosystemVal0OutputBoundaryState(outputBoundaries),
	)
	if val0State == VerifierEcosystemVal0StateActive {
		t.Fatalf("expected overclaim language to block active verifier discipline state, got %q", val0State)
	}
	if got := EvaluateVerifierEcosystemPoint7State(val0State); got != VerifierEcosystemPoint7StateNotComplete {
		t.Fatalf("expected point_7_pass to remain impossible in Val 0, got %q", got)
	}
	if got := EvaluateVerifierEcosystemVal0ProofsState(
		VerifierEcosystemVal0StateActive,
		VerifierEcosystemPoint7StatePass,
		verifierEcosystemVal0SupportedProfiles(),
		verifierEcosystemVal0SupportedModes(),
		VerifierEcosystemVal0ProofSurfaceRefs(),
		VerifierEcosystemVal0ProofEvidenceRefs(),
		[]string{"Val 0 cannot return point_7_pass."},
		verifierEcosystemVal0ProjectionDisclaimer(),
	); got == VerifierEcosystemVal0StateActive {
		t.Fatalf("expected point_7_pass to remain impossible in Val 0 proofs, got %q", got)
	}
}

func TestVerifierEcosystemVal0ProofEvidenceQualityValidation(t *testing.T) {
	evidence := VerifierEcosystemVal0VerifierEvidence()
	if !verifierEcosystemVal0ProofEvidenceQualityValid(evidence, VerifierEcosystemVal0ProofEvidenceRefs()) {
		t.Fatalf("expected healthy exact proof evidence set to validate")
	}

	evidence = VerifierEcosystemVal0VerifierEvidence()
	evidence[0].FreshnessState = IntelligenceCalibrationFreshnessStale
	if verifierEcosystemVal0ProofEvidenceQualityValid(evidence, VerifierEcosystemVal0ProofEvidenceRefs()) {
		t.Fatalf("expected stale modeled evidence to fail closed")
	}

	evidence = VerifierEcosystemVal0VerifierEvidence()
	evidence[0].EvidenceID = " "
	if verifierEcosystemVal0ProofEvidenceQualityValid(evidence, VerifierEcosystemVal0ProofEvidenceRefs()) {
		t.Fatalf("expected malformed modeled evidence ref to fail closed")
	}
}

func TestVerifierEcosystemVal0ProofsRequireExactEvidenceRefs(t *testing.T) {
	validEvidenceRefs := VerifierEcosystemVal0ProofEvidenceRefs()
	if got := EvaluateVerifierEcosystemVal0ProofsState(
		VerifierEcosystemVal0StateActive,
		VerifierEcosystemPoint7StateNotComplete,
		verifierEcosystemVal0SupportedProfiles(),
		verifierEcosystemVal0SupportedModes(),
		VerifierEcosystemVal0ProofSurfaceRefs(),
		validEvidenceRefs,
		[]string{"Val 0 cannot return point_7_pass."},
		verifierEcosystemVal0ProjectionDisclaimer(),
	); got != VerifierEcosystemVal0StateActive {
		t.Fatalf("expected exact required Val 0 proof evidence refs to keep proofs active, got %q", got)
	}

	testCases := []struct {
		name         string
		evidenceRefs []string
	}{
		{name: "five arbitrary strings do not keep proofs active", evidenceRefs: []string{"e1", "e2", "e3", "e4", "e5"}},
		{name: "missing verifier contract evidence blocks active", evidenceRefs: []string{
			"point6_integrated_closure",
			"verifier_discipline_foundation",
			"evidence:proof-envelope-001",
			"evidence:verification-scope-001",
			"evidence:trust-root-001",
			"evidence:revocation-001",
			"evidence:compatibility-001",
			"evidence:diagnostics-001",
			"evidence:output-boundary-001",
			"evidence:point7-governance-001",
		}},
		{name: "missing trust root evidence blocks active", evidenceRefs: []string{
			"point6_integrated_closure",
			"verifier_discipline_foundation",
			"evidence:verifier-contract-001",
			"evidence:proof-envelope-001",
			"evidence:verification-scope-001",
			"evidence:revocation-001",
			"evidence:compatibility-001",
			"evidence:diagnostics-001",
			"evidence:output-boundary-001",
			"evidence:point7-governance-001",
		}},
		{name: "duplicate evidence ref does not compensate for missing required evidence", evidenceRefs: []string{
			"point6_integrated_closure",
			"verifier_discipline_foundation",
			"evidence:verifier-contract-001",
			"evidence:proof-envelope-001",
			"evidence:verification-scope-001",
			"evidence:verification-scope-001",
			"evidence:revocation-001",
			"evidence:compatibility-001",
			"evidence:diagnostics-001",
			"evidence:output-boundary-001",
			"evidence:point7-governance-001",
		}},
		{name: "unknown extra evidence ref does not compensate for missing required evidence", evidenceRefs: []string{
			"point6_integrated_closure",
			"verifier_discipline_foundation",
			"evidence:verifier-contract-001",
			"evidence:proof-envelope-001",
			"evidence:verification-scope-001",
			"evidence:trust-root-001",
			"evidence:revocation-001",
			"evidence:compatibility-001",
			"evidence:diagnostics-001",
			"evidence:output-boundary-001",
			"evidence:unknown-extra-001",
		}},
		{name: "whitespace evidence ref fails closed", evidenceRefs: []string{
			"point6_integrated_closure",
			"verifier_discipline_foundation",
			"evidence:verifier-contract-001",
			"evidence:proof-envelope-001",
			"evidence:verification-scope-001",
			"evidence:trust-root-001",
			" ",
			"evidence:compatibility-001",
			"evidence:diagnostics-001",
			"evidence:output-boundary-001",
			"evidence:point7-governance-001",
		}},
	}

	for _, tc := range testCases {
		if got := EvaluateVerifierEcosystemVal0ProofsState(
			VerifierEcosystemVal0StateActive,
			VerifierEcosystemPoint7StateNotComplete,
			verifierEcosystemVal0SupportedProfiles(),
			verifierEcosystemVal0SupportedModes(),
			VerifierEcosystemVal0ProofSurfaceRefs(),
			tc.evidenceRefs,
			[]string{"Val 0 cannot return point_7_pass."},
			verifierEcosystemVal0ProjectionDisclaimer(),
		); got == VerifierEcosystemVal0StateActive {
			t.Fatalf("expected %s to fail closed, got %q", tc.name, got)
		}
	}
}
