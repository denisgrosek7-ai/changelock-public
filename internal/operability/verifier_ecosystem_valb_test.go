package operability

import "testing"

func activeVerifierEcosystemValBInputs() (
	VerifierEcosystemValBDependencySnapshot,
	VerifierEcosystemValBCompatibilityMatrix,
	VerifierEcosystemValBSchemaProofCompatibility,
	VerifierEcosystemValBMixedVersionDiagnosticsCatalog,
	VerifierEcosystemValBDiagnosticPrecedence,
	VerifierEcosystemValBFixtureCatalog,
	VerifierEcosystemValBConformanceCaseCatalog,
	VerifierEcosystemValBConformanceSuite,
	VerifierEcosystemValBOutputClassCatalog,
) {
	valADependency, _, _, result, _, _, _, _, _, _, _, _, _, valAState := activeVerifierEcosystemValAStates()
	dependency := VerifierEcosystemValBDependencySnapshot{
		Point5State:                    valADependency.Point5State,
		Point5DependencyState:          valADependency.Point5DependencyState,
		Point6State:                    valADependency.Point6State,
		Point6ClosureState:             valADependency.Point6ClosureState,
		Point6ClosurePrerequisiteState: valADependency.Point6ClosurePrerequisiteState,
		Point6ClosureInvariantState:    valADependency.Point6ClosureInvariantState,
		Point6ProofSurfaceState:        valADependency.Point6ProofSurfaceState,
		Point6PassRuleState:            valADependency.Point6PassRuleState,
		Point6PassAllowed:              valADependency.Point6PassAllowed,
		Val0CurrentState:               valADependency.Val0CurrentState,
		Val0State:                      valADependency.Val0State,
		ValACurrentState:               VerifierEcosystemValAStateActive,
		ValAState:                      valAState,
		Point7State:                    valADependency.Point7State,
	}
	matrix := VerifierEcosystemValBCompatibilityMatrixModel()
	compatibility := VerifierEcosystemValBSchemaProofCompatibilityModel()
	compatibility.SchemaVersion = result.SchemaVersion
	compatibility.ProofType = result.ProofType
	compatibility.VerifierVersion = result.VerifierVersion
	mixed := VerifierEcosystemValBMixedVersionDiagnosticsCatalogModel()
	precedence := VerifierEcosystemValBDiagnosticPrecedenceModel()
	precedence.ObservedDiagnostics = []string{VerifierEcosystemDiagnosticVerified}
	precedence.DerivedDiagnosticClass = DeriveVerifierEcosystemValBDiagnostic(precedence.ObservedDiagnostics, precedence.Caveats)
	fixtures := VerifierEcosystemValBFixtureCatalogModel()
	cases := VerifierEcosystemValBConformanceCaseCatalogModel()
	suite := VerifierEcosystemValBConformanceSuiteModel()
	outputs := VerifierEcosystemValBOutputClassCatalogModel()
	return dependency, matrix, compatibility, mixed, precedence, fixtures, cases, suite, outputs
}

func activeVerifierEcosystemValBStates() (
	VerifierEcosystemValBDependencySnapshot,
	VerifierEcosystemValBCompatibilityMatrix,
	VerifierEcosystemValBSchemaProofCompatibility,
	VerifierEcosystemValBMixedVersionDiagnosticsCatalog,
	VerifierEcosystemValBDiagnosticPrecedence,
	VerifierEcosystemValBFixtureCatalog,
	VerifierEcosystemValBConformanceCaseCatalog,
	VerifierEcosystemValBConformanceSuite,
	VerifierEcosystemValBOutputClassCatalog,
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
	dependency, matrix, compatibility, mixed, precedence, fixtures, cases, suite, outputs := activeVerifierEcosystemValBInputs()
	matrixState := EvaluateVerifierEcosystemValBCompatibilityMatrixState(matrix)
	compatibilityState := EvaluateVerifierEcosystemValBSchemaProofCompatibilityState(compatibility, matrix)
	mixedState := EvaluateVerifierEcosystemValBMixedVersionDiagnosticsState(mixed)
	precedenceState := EvaluateVerifierEcosystemValBDiagnosticPrecedenceState(precedence)
	fixtureState := EvaluateVerifierEcosystemValBFixtureDescriptorState(fixtures)
	caseState := EvaluateVerifierEcosystemValBConformanceCaseState(cases, fixtures, outputs)
	suiteState := EvaluateVerifierEcosystemValBConformanceSuiteState(suite, cases, fixtures, outputs)
	outputClassState := EvaluateVerifierEcosystemValBOutputClassState(outputs)
	valBState := EvaluateVerifierEcosystemValBState(
		dependency,
		matrixState,
		compatibilityState,
		mixedState,
		precedenceState,
		fixtureState,
		caseState,
		suiteState,
		outputClassState,
	)
	return dependency, matrix, compatibility, mixed, precedence, fixtures, cases, suite, outputs, matrixState, compatibilityState, mixedState, precedenceState, fixtureState, caseState, suiteState, outputClassState, valBState
}

func removeTrimmedString(values []string, target string) []string {
	filtered := make([]string, 0, len(values))
	for _, value := range values {
		if value == target {
			continue
		}
		filtered = append(filtered, value)
	}
	return filtered
}

func fixtureByID(t *testing.T, catalog VerifierEcosystemValBFixtureCatalog, fixtureID string) VerifierEcosystemValBFixtureDescriptor {
	t.Helper()
	for _, fixture := range catalog.Fixtures {
		if fixture.FixtureID == fixtureID {
			return fixture
		}
	}
	t.Fatalf("fixture %q not found", fixtureID)
	return VerifierEcosystemValBFixtureDescriptor{}
}

func TestVerifierEcosystemValBDependencyGates(t *testing.T) {
	dependency, _, _, _, _, _, _, _, _, matrixState, compatibilityState, mixedState, precedenceState, fixtureState, caseState, suiteState, outputClassState, valBState := activeVerifierEcosystemValBStates()
	if valBState != VerifierEcosystemValBStateActive {
		t.Fatalf("expected active Val B state with active Točka 6, Val 0, and Val A dependencies, got %q", valBState)
	}
	if got := EvaluateVerifierEcosystemValBPoint7State(valBState); got != VerifierEcosystemPoint7StateNotComplete {
		t.Fatalf("expected point 7 to remain not complete in Val B, got %q", got)
	}

	testCases := []struct {
		name   string
		mutate func(*VerifierEcosystemValBDependencySnapshot)
	}{
		{name: "missing val0 blocks active", mutate: func(model *VerifierEcosystemValBDependencySnapshot) {
			model.Val0State = VerifierEcosystemVal0StatePartial
		}},
		{name: "missing vala blocks active", mutate: func(model *VerifierEcosystemValBDependencySnapshot) {
			model.ValAState = VerifierEcosystemValAStatePartial
		}},
		{name: "missing point6 closure blocks active", mutate: func(model *VerifierEcosystemValBDependencySnapshot) {
			model.Point6ClosureState = ReferenceArchitectureValEStatePartial
		}},
		{name: "point7 state other than not complete blocks active", mutate: func(model *VerifierEcosystemValBDependencySnapshot) {
			model.Point7State = VerifierEcosystemPoint7StatePass
		}},
	}

	for _, tc := range testCases {
		snapshot := dependency
		tc.mutate(&snapshot)
		if got := EvaluateVerifierEcosystemValBState(snapshot, matrixState, compatibilityState, mixedState, precedenceState, fixtureState, caseState, suiteState, outputClassState); got != VerifierEcosystemValBStateBlocked {
			t.Fatalf("expected %s to return %q, got %q", tc.name, VerifierEcosystemValBStateBlocked, got)
		}
	}
}

func TestVerifierEcosystemValBCompatibilityMatrixValidation(t *testing.T) {
	matrix := VerifierEcosystemValBCompatibilityMatrixModel()
	if got := EvaluateVerifierEcosystemValBCompatibilityMatrixState(matrix); got != VerifierEcosystemValBCompatibilityMatrixStateActive {
		t.Fatalf("expected valid compatibility matrix to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValBCompatibilityMatrix)
		expected string
	}{
		{name: "missing matrix id", mutate: func(model *VerifierEcosystemValBCompatibilityMatrix) { model.CompatibilityMatrixID = "" }, expected: VerifierEcosystemValBCompatibilityMatrixStateIncomplete},
		{name: "unsupported schema fails closed", mutate: func(model *VerifierEcosystemValBCompatibilityMatrix) {
			model.CompatibilityEntries[0].CompatibilityState = ReferenceArchitectureCompatibilityUnsupported
			model.CompatibilityEntries[0].RequiredDiagnostics = []string{VerifierEcosystemDiagnosticUnsupportedSchema}
		}, expected: VerifierEcosystemValBCompatibilityMatrixStateBlocked},
		{name: "compatible with warnings requires caveat", mutate: func(model *VerifierEcosystemValBCompatibilityMatrix) {
			model.CompatibilityEntries[4].Caveats = nil
		}, expected: VerifierEcosystemValBCompatibilityMatrixStateBlocked},
		{name: "duplicate entry does not compensate for missing required entry", mutate: func(model *VerifierEcosystemValBCompatibilityMatrix) {
			model.CompatibilityEntries[5] = model.CompatibilityEntries[0]
		}, expected: VerifierEcosystemValBCompatibilityMatrixStatePartial},
		{name: "unknown extra entry does not compensate for missing required entry", mutate: func(model *VerifierEcosystemValBCompatibilityMatrix) {
			model.CompatibilityEntries = append(model.CompatibilityEntries, VerifierEcosystemValBCompatibilityEntry{
				CurrentState:         "compatibility_entry_ready",
				EntryID:              "matrix-entry:extra",
				SchemaVersion:        "changelock.verifier.proof_envelope.v1",
				ProofType:            VerifierEcosystemProofTypeSignedAttestation,
				VerifierVersion:      "reference-verifier/vala-2026.04",
				TrustRootVersion:     "2026.04",
				CompatibilityState:   ReferenceArchitectureCompatibilityCompatible,
				RequiredDiagnostics:  []string{VerifierEcosystemDiagnosticVerified},
				ProjectionDisclaimer: verifierEcosystemValBProjectionDisclaimer(),
			})
		}, expected: VerifierEcosystemValBCompatibilityMatrixStatePartial},
	}

	for _, tc := range testCases {
		model := VerifierEcosystemValBCompatibilityMatrixModel()
		tc.mutate(&model)
		if got := EvaluateVerifierEcosystemValBCompatibilityMatrixState(model); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValBSchemaProofCompatibilityEvaluator(t *testing.T) {
	matrix := VerifierEcosystemValBCompatibilityMatrixModel()
	model := VerifierEcosystemValBSchemaProofCompatibilityModel()
	if got := EvaluateVerifierEcosystemValBSchemaProofCompatibilityState(model, matrix); got != VerifierEcosystemValBSchemaProofCompatibilityStateActive {
		t.Fatalf("expected compatible schema and proof pair to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValBSchemaProofCompatibility)
		expected string
	}{
		{name: "unsupported proof fails closed", mutate: func(model *VerifierEcosystemValBSchemaProofCompatibility) { model.ProofType = "unknown_proof_type" }, expected: VerifierEcosystemValBSchemaProofCompatibilityStateUnknown},
		{name: "unsupported schema fails closed", mutate: func(model *VerifierEcosystemValBSchemaProofCompatibility) {
			model.SchemaVersion = "changelock.verifier.proof_envelope.v9"
		}, expected: VerifierEcosystemValBSchemaProofCompatibilityStateUnknown},
		{name: "deprecated schema returns warning state", mutate: func(model *VerifierEcosystemValBSchemaProofCompatibility) {
			model.SchemaVersion = "changelock.verifier.proof_envelope.v0"
			model.CompatibilityState = ReferenceArchitectureCompatibilityDeprecated
			model.DerivedDiagnosticClass = VerifierEcosystemDiagnosticCompatibilityWarning
			model.RequiredDiagnostics = []string{VerifierEcosystemDiagnosticCompatibilityWarning}
			model.Caveats = []string{"deprecated schema remains warning-bearing"}
			model.MigrationOrSupersessionRef = "migration:proof-envelope-v1"
		}, expected: VerifierEcosystemValBSchemaProofCompatibilityStatePartial},
		{name: "superseded proof returns superseded state", mutate: func(model *VerifierEcosystemValBSchemaProofCompatibility) {
			model.SchemaVersion = "changelock.verifier.proof_envelope.v1"
			model.ProofType = VerifierEcosystemProofTypeLineageBundle
			model.VerifierVersion = "reference-verifier/vala-2025.12"
			model.TrustRootVersion = "2025.12"
			model.CompatibilityState = ReferenceArchitectureCompatibilitySuperseded
			model.DerivedDiagnosticClass = VerifierEcosystemDiagnosticSupersededProof
			model.RequiredDiagnostics = []string{VerifierEcosystemDiagnosticSupersededProof}
			model.Caveats = []string{"superseded proof remains explicit"}
			model.MigrationOrSupersessionRef = "supersession:lineage-bundle-current"
		}, expected: VerifierEcosystemValBSchemaProofCompatibilityStatePartial},
		{name: "unknown compatibility fails closed", mutate: func(model *VerifierEcosystemValBSchemaProofCompatibility) {
			model.CompatibilityState = ReferenceArchitectureCompatibilityUnknown
			model.DerivedDiagnosticClass = VerifierEcosystemDiagnosticUnknown
			model.RequiredDiagnostics = []string{VerifierEcosystemDiagnosticUnknown}
		}, expected: VerifierEcosystemValBSchemaProofCompatibilityStateUnknown},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValBSchemaProofCompatibilityModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValBSchemaProofCompatibilityState(mutated, matrix); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValBMixedVersionDiagnostics(t *testing.T) {
	model := VerifierEcosystemValBMixedVersionDiagnosticsCatalogModel()
	if got := EvaluateVerifierEcosystemValBMixedVersionDiagnosticsState(model); got != VerifierEcosystemValBMixedVersionStateActive {
		t.Fatalf("expected valid mixed-version catalog to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValBMixedVersionDiagnosticsCatalog)
		expected string
	}{
		{name: "missing expected diagnostic fails closed", mutate: func(model *VerifierEcosystemValBMixedVersionDiagnosticsCatalog) {
			model.Cases[0].ExpectedDiagnosticClass = ""
		}, expected: VerifierEcosystemValBMixedVersionStateIncomplete},
		{name: "unknown diagnostic class fails closed", mutate: func(model *VerifierEcosystemValBMixedVersionDiagnosticsCatalog) {
			model.Cases[0].ExpectedDiagnosticClass = "diagnostic:not_real"
		}, expected: VerifierEcosystemValBMixedVersionStateUnknown},
		{name: "deprecated mixed-version case requires caveat", mutate: func(model *VerifierEcosystemValBMixedVersionDiagnosticsCatalog) { model.Cases[0].Caveats = nil }, expected: VerifierEcosystemValBMixedVersionStateBlocked},
		{name: "superseded mixed-version case requires supersession ref", mutate: func(model *VerifierEcosystemValBMixedVersionDiagnosticsCatalog) {
			model.Cases[2].MigrationOrSupersessionRef = ""
		}, expected: VerifierEcosystemValBMixedVersionStateBlocked},
		{name: "mixed-version output cannot claim universal compatibility", mutate: func(model *VerifierEcosystemValBMixedVersionDiagnosticsCatalog) {
			model.Cases[1].UniversalCompatibilityClaim = true
		}, expected: VerifierEcosystemValBMixedVersionStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValBMixedVersionDiagnosticsCatalogModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValBMixedVersionDiagnosticsState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValBDiagnosticPrecedence(t *testing.T) {
	testCases := []struct {
		name     string
		observed []string
		caveats  []string
		expected string
	}{
		{name: "invalid signature outranks compatibility warning", observed: []string{VerifierEcosystemDiagnosticCompatibilityWarning, VerifierEcosystemDiagnosticInvalidSignature}, caveats: []string{"warning"}, expected: VerifierEcosystemDiagnosticInvalidSignature},
		{name: "digest mismatch outranks compatibility warning", observed: []string{VerifierEcosystemDiagnosticCompatibilityWarning, VerifierEcosystemDiagnosticDigestMismatch}, caveats: []string{"warning"}, expected: VerifierEcosystemDiagnosticDigestMismatch},
		{name: "revoked issuer blocks verified", observed: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticRevokedIssuer}, expected: VerifierEcosystemDiagnosticRevokedIssuer},
		{name: "unknown diagnostic fails closed", observed: []string{"not_a_real_diagnostic"}, expected: VerifierEcosystemDiagnosticUnknown},
		{name: "compatibility warning requires caveat", observed: []string{VerifierEcosystemDiagnosticCompatibilityWarning}, expected: VerifierEcosystemDiagnosticUnknown},
	}
	for _, tc := range testCases {
		if got := DeriveVerifierEcosystemValBDiagnostic(tc.observed, tc.caveats); got != tc.expected {
			t.Fatalf("expected %s to derive %q, got %q", tc.name, tc.expected, got)
		}
	}

	model := VerifierEcosystemValBDiagnosticPrecedenceModel()
	if got := EvaluateVerifierEcosystemValBDiagnosticPrecedenceState(model); got != VerifierEcosystemValBDiagnosticPrecedenceStateActive {
		t.Fatalf("expected valid precedence model to be active, got %q", got)
	}
	model.DerivedDiagnosticClass = "not_real"
	if got := EvaluateVerifierEcosystemValBDiagnosticPrecedenceState(model); got != VerifierEcosystemValBDiagnosticPrecedenceStateUnknown {
		t.Fatalf("expected fake diagnostic class to fail closed, got %q", got)
	}
}

func TestVerifierEcosystemValBFixtureDescriptors(t *testing.T) {
	model := VerifierEcosystemValBFixtureCatalogModel()
	if got := EvaluateVerifierEcosystemValBFixtureDescriptorState(model); got != VerifierEcosystemValBFixtureDescriptorStateActive {
		t.Fatalf("expected valid fixture catalog to be active, got %q", got)
	}

	if stale := fixtureByID(t, model, "fixture:stale-proof-envelope"); stale.ExpectedDiagnostic != VerifierEcosystemDiagnosticStaleArtifact {
		t.Fatalf("expected stale fixture to require stale diagnostic, got %#v", stale)
	}
	if revoked := fixtureByID(t, model, "fixture:revoked-issuer"); revoked.ExpectedDiagnostic != VerifierEcosystemDiagnosticRevokedIssuer {
		t.Fatalf("expected revoked fixture to require revoked diagnostic, got %#v", revoked)
	}
	if unsupported := fixtureByID(t, model, "fixture:unsupported-schema"); unsupported.ExpectedDiagnostic != VerifierEcosystemDiagnosticUnsupportedSchema {
		t.Fatalf("expected unsupported schema fixture to require unsupported_schema, got %#v", unsupported)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValBFixtureCatalog)
		expected string
	}{
		{name: "fixture descriptor cannot claim production evidence", mutate: func(model *VerifierEcosystemValBFixtureCatalog) { model.Fixtures[0].ProductionEvidenceClaim = true }, expected: VerifierEcosystemValBFixtureDescriptorStateBlocked},
		{name: "malformed timestamp fixture cannot verify", mutate: func(model *VerifierEcosystemValBFixtureCatalog) {
			for i := range model.Fixtures {
				if model.Fixtures[i].FixtureID == "fixture:malformed-timestamp" {
					model.Fixtures[i].ExpectedResult = VerifierEcosystemValAOverallResultVerified
				}
			}
		}, expected: VerifierEcosystemValBFixtureDescriptorStatePartial},
		{name: "missing signature fixture cannot verify", mutate: func(model *VerifierEcosystemValBFixtureCatalog) {
			for i := range model.Fixtures {
				if model.Fixtures[i].FixtureID == "fixture:missing-signature" {
					model.Fixtures[i].ExpectedResult = VerifierEcosystemValAOverallResultVerified
				}
			}
		}, expected: VerifierEcosystemValBFixtureDescriptorStatePartial},
		{name: "digest mismatch fixture cannot verify", mutate: func(model *VerifierEcosystemValBFixtureCatalog) {
			for i := range model.Fixtures {
				if model.Fixtures[i].FixtureID == "fixture:digest-mismatch" {
					model.Fixtures[i].ExpectedDiagnostic = VerifierEcosystemDiagnosticVerified
				}
			}
		}, expected: VerifierEcosystemValBFixtureDescriptorStatePartial},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValBFixtureCatalogModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValBFixtureDescriptorState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValBConformanceCasesAndSuite(t *testing.T) {
	_, _, _, _, _, fixtures, cases, suite, outputs, _, _, _, _, fixtureState, caseState, suiteState, outputClassState, _ := activeVerifierEcosystemValBStates()
	if fixtureState != VerifierEcosystemValBFixtureDescriptorStateActive ||
		caseState != VerifierEcosystemValBConformanceCaseStateActive ||
		suiteState != VerifierEcosystemValBConformanceSuiteStateActive ||
		outputClassState != VerifierEcosystemValBOutputClassStateActive {
		t.Fatalf("expected active conformance and output states, got fixture=%q case=%q suite=%q output=%q", fixtureState, caseState, suiteState, outputClassState)
	}

	missingCaseCatalog := VerifierEcosystemValBConformanceCaseCatalogModel()
	missingCaseCatalog.Cases = missingCaseCatalog.Cases[1:]
	if got := EvaluateVerifierEcosystemValBConformanceSuiteState(suite, missingCaseCatalog, fixtures, outputs); got != VerifierEcosystemValBConformanceSuiteStateIncomplete {
		t.Fatalf("expected missing required case to return %q, got %q", VerifierEcosystemValBConformanceSuiteStateIncomplete, got)
	}

	duplicateCaseCatalog := VerifierEcosystemValBConformanceCaseCatalogModel()
	duplicateCaseCatalog.Cases = append(duplicateCaseCatalog.Cases[:1], duplicateCaseCatalog.Cases[0])
	if got := EvaluateVerifierEcosystemValBConformanceSuiteState(suite, duplicateCaseCatalog, fixtures, outputs); got != VerifierEcosystemValBConformanceSuiteStateIncomplete {
		t.Fatalf("expected duplicate case not to compensate for missing required case, got %q", got)
	}

	extraCaseCatalog := VerifierEcosystemValBConformanceCaseCatalogModel()
	extraCaseCatalog.Cases = extraCaseCatalog.Cases[1:]
	extraCaseCatalog.Cases = append(extraCaseCatalog.Cases, VerifierEcosystemValBConformanceCase{
		CurrentState:            "conformance_case_ready",
		ConformanceCaseID:       "conformance:unknown-extra",
		FixtureRef:              "fixture:valid-proof-envelope",
		VerifierContractRef:     "verifier-contract-ref/val0",
		InputRef:                "input-ref/extra",
		ExpectedOverallResult:   VerifierEcosystemValAOverallResultVerified,
		ExpectedDiagnosticClass: VerifierEcosystemDiagnosticVerified,
		ExpectedOutputClass:     VerifierEcosystemValBOutputClassVerified,
		RequiredFields:          verifierEcosystemValBRequiredConformanceFields(),
		ForbiddenClaims:         verifierEcosystemValBRequiredForbiddenClaims(),
		ProjectionDisclaimer:    verifierEcosystemValBProjectionDisclaimer(),
	})
	if got := EvaluateVerifierEcosystemValBConformanceSuiteState(suite, extraCaseCatalog, fixtures, outputs); got != VerifierEcosystemValBConformanceSuiteStateIncomplete {
		t.Fatalf("expected unknown extra case not to compensate for missing required case, got %q", got)
	}

	unknownOutputCatalog := VerifierEcosystemValBConformanceCaseCatalogModel()
	unknownOutputCatalog.Cases[0].ExpectedOutputClass = "not_a_real_output_class"
	if got := EvaluateVerifierEcosystemValBConformanceCaseState(unknownOutputCatalog, fixtures, outputs); got != VerifierEcosystemValBConformanceCaseStateUnknown {
		t.Fatalf("expected unknown expected output class to fail closed, got %q", got)
	}

	blockedClaimsCatalog := VerifierEcosystemValBConformanceCaseCatalogModel()
	blockedClaimsCatalog.Cases[0].ObservedClaims = append(blockedClaimsCatalog.Cases[0].ObservedClaims, "verifier certification")
	if got := EvaluateVerifierEcosystemValBConformanceCaseState(blockedClaimsCatalog, fixtures, outputs); got != VerifierEcosystemValBConformanceCaseStateBlocked {
		t.Fatalf("expected forbidden claim to block conformance, got %q", got)
	}

	blockedSuite := VerifierEcosystemValBConformanceSuiteModel()
	blockedSuite.CertificationClaim = true
	if got := EvaluateVerifierEcosystemValBConformanceSuiteState(blockedSuite, cases, fixtures, outputs); got != VerifierEcosystemValBConformanceSuiteStateBlocked {
		t.Fatalf("expected suite certification claim to block conformance, got %q", got)
	}
}

func TestVerifierEcosystemValBOutputClasses(t *testing.T) {
	model := VerifierEcosystemValBOutputClassCatalogModel()
	if got := EvaluateVerifierEcosystemValBOutputClassState(model); got != VerifierEcosystemValBOutputClassStateActive {
		t.Fatalf("expected valid output class catalog to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValBOutputClassCatalog)
		expected string
	}{
		{name: "verified output class requires verified result and diagnostic", mutate: func(model *VerifierEcosystemValBOutputClassCatalog) {
			model.Mappings[0].OverallResult = VerifierEcosystemValAOverallResultInvalid
		}, expected: VerifierEcosystemValBOutputClassStatePartial},
		{name: "verified with warnings requires caveat", mutate: func(model *VerifierEcosystemValBOutputClassCatalog) {
			model.Mappings[1].Caveats = nil
		}, expected: VerifierEcosystemValBOutputClassStatePartial},
		{name: "redaction blocked cannot become verified", mutate: func(model *VerifierEcosystemValBOutputClassCatalog) {
			for i := range model.Mappings {
				if model.Mappings[i].DiagnosticClass == VerifierEcosystemDiagnosticRedactionViolation {
					model.Mappings[i].OutputClass = VerifierEcosystemValBOutputClassVerified
				}
			}
		}, expected: VerifierEcosystemValBOutputClassStatePartial},
		{name: "unknown output class fails closed", mutate: func(model *VerifierEcosystemValBOutputClassCatalog) {
			model.Mappings[0].OutputClass = "not_a_real_output_class"
		}, expected: VerifierEcosystemValBOutputClassStateUnknown},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValBOutputClassCatalogModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValBOutputClassState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValBAggregateStatePrecedence(t *testing.T) {
	dependency, _, _, _, _, _, _, _, _, matrixState, compatibilityState, mixedState, precedenceState, fixtureState, caseState, suiteState, outputClassState, _ := activeVerifierEcosystemValBStates()
	testCases := []struct {
		name          string
		matrix        string
		compatibility string
		mixed         string
		precedence    string
		fixture       string
		cases         string
		suite         string
		output        string
		expected      string
	}{
		{name: "partial before blocked returns blocked", matrix: VerifierEcosystemValBCompatibilityMatrixStatePartial, compatibility: compatibilityState, mixed: mixedState, precedence: precedenceState, fixture: fixtureState, cases: caseState, suite: VerifierEcosystemValBConformanceSuiteStateBlocked, output: outputClassState, expected: VerifierEcosystemValBStateBlocked},
		{name: "blocked before partial returns blocked", matrix: VerifierEcosystemValBCompatibilityMatrixStateBlocked, compatibility: VerifierEcosystemValBSchemaProofCompatibilityStatePartial, mixed: mixedState, precedence: precedenceState, fixture: fixtureState, cases: caseState, suite: suiteState, output: outputClassState, expected: VerifierEcosystemValBStateBlocked},
		{name: "partial plus unknown returns unknown", matrix: VerifierEcosystemValBCompatibilityMatrixStatePartial, compatibility: VerifierEcosystemValBSchemaProofCompatibilityStateUnknown, mixed: mixedState, precedence: precedenceState, fixture: fixtureState, cases: caseState, suite: suiteState, output: outputClassState, expected: VerifierEcosystemValBStateUnknown},
		{name: "incomplete plus partial returns incomplete", matrix: matrixState, compatibility: VerifierEcosystemValBSchemaProofCompatibilityStateIncomplete, mixed: VerifierEcosystemValBMixedVersionStatePartial, precedence: precedenceState, fixture: fixtureState, cases: caseState, suite: suiteState, output: outputClassState, expected: VerifierEcosystemValBStateIncomplete},
		{name: "all active returns active", matrix: matrixState, compatibility: compatibilityState, mixed: mixedState, precedence: precedenceState, fixture: fixtureState, cases: caseState, suite: suiteState, output: outputClassState, expected: VerifierEcosystemValBStateActive},
		{name: "fake component state fails closed", matrix: "not_a_real_valb_matrix_state", compatibility: compatibilityState, mixed: mixedState, precedence: precedenceState, fixture: fixtureState, cases: caseState, suite: suiteState, output: outputClassState, expected: VerifierEcosystemValBStateUnknown},
	}
	for _, tc := range testCases {
		if got := EvaluateVerifierEcosystemValBState(dependency, tc.matrix, tc.compatibility, tc.mixed, tc.precedence, tc.fixture, tc.cases, tc.suite, tc.output); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}

	dependency.Point6PassAllowed = false
	if got := EvaluateVerifierEcosystemValBState(dependency, matrixState, compatibilityState, mixedState, precedenceState, fixtureState, caseState, suiteState, outputClassState); got != VerifierEcosystemValBStateBlocked {
		t.Fatalf("expected missing dependency to return %q, got %q", VerifierEcosystemValBStateBlocked, got)
	}
}

func TestVerifierEcosystemValBProofSurfaceCompletenessAndPoint7PassImpossibility(t *testing.T) {
	_, _, _, _, _, _, _, _, _, matrixState, compatibilityState, mixedState, precedenceState, fixtureState, caseState, suiteState, outputClassState, valBState := activeVerifierEcosystemValBStates()
	if valBState != VerifierEcosystemValBStateActive {
		t.Fatalf("expected active Val B state, got %q", valBState)
	}
	if got := EvaluateVerifierEcosystemValBPoint7State(valBState); got != VerifierEcosystemPoint7StateNotComplete {
		t.Fatalf("expected point_7_pass to remain impossible in Val B, got %q", got)
	}

	surfaces := VerifierEcosystemValBProofSurfaceRefs()
	evidenceRefs := VerifierEcosystemValBProofEvidenceRefs()
	limitations := []string{"Val B remains advisory."}
	reasons := []string{"Val B cannot return point_7_pass."}
	if got := EvaluateVerifierEcosystemValBProofsState(
		valBState,
		VerifierEcosystemPoint7StateNotComplete,
		VerifierEcosystemVal0StateActive,
		VerifierEcosystemValAStateActive,
		surfaces,
		evidenceRefs,
		limitations,
		reasons,
		verifierEcosystemValBProjectionDisclaimer(),
	); got != VerifierEcosystemValBStateActive {
		t.Fatalf("expected exact Val B proof surface set to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		surfaces []string
		evidence []string
		expected string
	}{
		{name: "missing surface fails closed", surfaces: removeTrimmedString(surfaces, "/v1/verifier-ecosystem/valb/proofs"), evidence: evidenceRefs, expected: VerifierEcosystemValBStatePartial},
		{name: "duplicate surface does not compensate", surfaces: append(removeTrimmedString(surfaces, "/v1/verifier-ecosystem/valb/proofs"), "/v1/verifier-ecosystem/valb/output-classes"), evidence: evidenceRefs, expected: VerifierEcosystemValBStatePartial},
		{name: "unknown extra surface does not compensate", surfaces: append(append([]string{}, surfaces...), "/v1/verifier-ecosystem/valb/extra"), evidence: evidenceRefs, expected: VerifierEcosystemValBStatePartial},
		{name: "whitespace surface ref fails closed", surfaces: append(removeTrimmedString(surfaces, "/v1/verifier-ecosystem/valb/proofs"), " "), evidence: evidenceRefs, expected: VerifierEcosystemValBStatePartial},
		{name: "missing evidence ref fails closed", surfaces: surfaces, evidence: removeTrimmedString(evidenceRefs, "evidence:compatibility-matrix-001"), expected: VerifierEcosystemValBStatePartial},
		{name: "duplicate evidence ref does not compensate", surfaces: surfaces, evidence: append(removeTrimmedString(evidenceRefs, "evidence:compatibility-matrix-001"), "evidence:output-classes-001"), expected: VerifierEcosystemValBStatePartial},
		{name: "unknown extra evidence ref does not compensate", surfaces: surfaces, evidence: append(append([]string{}, evidenceRefs...), "evidence:unknown-extra"), expected: VerifierEcosystemValBStatePartial},
		{name: "whitespace evidence ref fails closed", surfaces: surfaces, evidence: append(removeTrimmedString(evidenceRefs, "evidence:compatibility-matrix-001"), " "), expected: VerifierEcosystemValBStatePartial},
	}

	for _, tc := range testCases {
		if got := EvaluateVerifierEcosystemValBProofsState(
			valBState,
			VerifierEcosystemPoint7StateNotComplete,
			VerifierEcosystemVal0StateActive,
			VerifierEcosystemValAStateActive,
			tc.surfaces,
			tc.evidence,
			limitations,
			reasons,
			verifierEcosystemValBProjectionDisclaimer(),
		); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}

	_ = matrixState
	_ = compatibilityState
	_ = mixedState
	_ = precedenceState
	_ = fixtureState
	_ = caseState
	_ = suiteState
	_ = outputClassState
}
