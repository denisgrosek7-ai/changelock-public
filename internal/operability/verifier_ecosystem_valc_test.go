package operability

import "testing"

func activeVerifierEcosystemValCInputs() (
	VerifierEcosystemValCDependencySnapshot,
	VerifierEcosystemValCAudienceSurfaceCatalog,
	VerifierEcosystemValCPublicOutputContract,
	VerifierEcosystemValCPartnerOutputContract,
	VerifierEcosystemValCAuditorFlowContract,
	VerifierEcosystemValCRequestContract,
	VerifierEcosystemValCPublisherCompatibilityProfile,
	VerifierEcosystemValCArtifactPublishingRuleCatalog,
	VerifierEcosystemValCTrustDistributionVisibility,
) {
	valBDependency, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, valBState := activeVerifierEcosystemValBStates()
	dependency := VerifierEcosystemValCDependencySnapshot{
		Point5State:                    valBDependency.Point5State,
		Point5DependencyState:          valBDependency.Point5DependencyState,
		Point6State:                    valBDependency.Point6State,
		Point6ClosureState:             valBDependency.Point6ClosureState,
		Point6ClosurePrerequisiteState: valBDependency.Point6ClosurePrerequisiteState,
		Point6ClosureInvariantState:    valBDependency.Point6ClosureInvariantState,
		Point6ProofSurfaceState:        valBDependency.Point6ProofSurfaceState,
		Point6PassRuleState:            valBDependency.Point6PassRuleState,
		Point6PassAllowed:              valBDependency.Point6PassAllowed,
		Val0CurrentState:               valBDependency.Val0CurrentState,
		Val0State:                      valBDependency.Val0State,
		ValACurrentState:               valBDependency.ValACurrentState,
		ValAState:                      valBDependency.ValAState,
		ValBCurrentState:               VerifierEcosystemValBStateActive,
		ValBState:                      valBState,
		Point7State:                    valBDependency.Point7State,
	}
	audiences := verifierEcosystemValCAudienceSurfaceCatalogModel()
	publicOutput := VerifierEcosystemValCPublicOutputContractModel()
	partnerOutput := VerifierEcosystemValCPartnerOutputContractModel()
	auditorFlow := VerifierEcosystemValCAuditorFlowContractModel()
	requestContract := VerifierEcosystemValCRequestContractModel()
	publisherProfile := VerifierEcosystemValCPublisherCompatibilityProfileModel()
	artifactRules := VerifierEcosystemValCArtifactPublishingRuleCatalogModel()
	trustDistribution := VerifierEcosystemValCTrustDistributionVisibilityModel()
	return dependency, audiences, publicOutput, partnerOutput, auditorFlow, requestContract, publisherProfile, artifactRules, trustDistribution
}

func activeVerifierEcosystemValCStates() (
	VerifierEcosystemValCDependencySnapshot,
	VerifierEcosystemValCAudienceSurfaceCatalog,
	VerifierEcosystemValCPublicOutputContract,
	VerifierEcosystemValCPartnerOutputContract,
	VerifierEcosystemValCAuditorFlowContract,
	VerifierEcosystemValCRequestContract,
	VerifierEcosystemValCPublisherCompatibilityProfile,
	VerifierEcosystemValCArtifactPublishingRuleCatalog,
	VerifierEcosystemValCTrustDistributionVisibility,
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
	dependency, audiences, publicOutput, partnerOutput, auditorFlow, requestContract, publisherProfile, artifactRules, trustDistribution := activeVerifierEcosystemValCInputs()
	audienceState := EvaluateVerifierEcosystemValCAudienceSurfaceState(audiences)
	publicState := EvaluateVerifierEcosystemValCPublicOutputState(publicOutput, audiences)
	partnerState := EvaluateVerifierEcosystemValCPartnerOutputState(partnerOutput, audiences)
	auditorState := EvaluateVerifierEcosystemValCAuditorFlowState(auditorFlow)
	requestState := EvaluateVerifierEcosystemValCRequestContractState(requestContract)
	publisherState := EvaluateVerifierEcosystemValCPublisherProfileState(publisherProfile)
	artifactRuleState := EvaluateVerifierEcosystemValCArtifactRuleState(artifactRules)
	trustDistributionState := EvaluateVerifierEcosystemValCTrustDistributionState(trustDistribution)
	valCState := EvaluateVerifierEcosystemValCState(
		dependency,
		audienceState,
		publicState,
		partnerState,
		auditorState,
		requestState,
		publisherState,
		artifactRuleState,
		trustDistributionState,
	)
	return dependency, audiences, publicOutput, partnerOutput, auditorFlow, requestContract, publisherProfile, artifactRules, trustDistribution, audienceState, publicState, partnerState, auditorState, requestState, publisherState, artifactRuleState, trustDistributionState, valCState
}

func TestVerifierEcosystemValCDependencyGates(t *testing.T) {
	dependency, _, _, _, _, _, _, _, _, audienceState, publicState, partnerState, auditorState, requestState, publisherState, artifactRuleState, trustDistributionState, valCState := activeVerifierEcosystemValCStates()
	if valCState != VerifierEcosystemValCStateActive {
		t.Fatalf("expected active Val C state with active Točka 6, Val 0, Val A, and Val B dependencies, got %q", valCState)
	}
	if got := EvaluateVerifierEcosystemValCPoint7State(valCState); got != VerifierEcosystemPoint7StateNotComplete {
		t.Fatalf("expected point 7 to remain not complete in Val C, got %q", got)
	}

	testCases := []struct {
		name   string
		mutate func(*VerifierEcosystemValCDependencySnapshot)
	}{
		{name: "missing val0 blocks active", mutate: func(model *VerifierEcosystemValCDependencySnapshot) {
			model.Val0State = VerifierEcosystemVal0StatePartial
		}},
		{name: "missing vala blocks active", mutate: func(model *VerifierEcosystemValCDependencySnapshot) {
			model.ValAState = VerifierEcosystemValAStatePartial
		}},
		{name: "missing valb blocks active", mutate: func(model *VerifierEcosystemValCDependencySnapshot) {
			model.ValBState = VerifierEcosystemValBStatePartial
		}},
		{name: "missing point6 closure blocks active", mutate: func(model *VerifierEcosystemValCDependencySnapshot) {
			model.Point6ClosureState = ReferenceArchitectureValEStatePartial
		}},
		{name: "point7 other than not complete blocks active", mutate: func(model *VerifierEcosystemValCDependencySnapshot) {
			model.Point7State = VerifierEcosystemPoint7StatePass
		}},
	}

	for _, tc := range testCases {
		snapshot := dependency
		tc.mutate(&snapshot)
		if got := EvaluateVerifierEcosystemValCState(snapshot, audienceState, publicState, partnerState, auditorState, requestState, publisherState, artifactRuleState, trustDistributionState); got != VerifierEcosystemValCStateBlocked {
			t.Fatalf("expected %s to return %q, got %q", tc.name, VerifierEcosystemValCStateBlocked, got)
		}
	}
}

func TestVerifierEcosystemValCAudienceSurfaces(t *testing.T) {
	model := verifierEcosystemValCAudienceSurfaceCatalogModel()
	if got := EvaluateVerifierEcosystemValCAudienceSurfaceState(model); got != VerifierEcosystemValCAudienceSurfaceStateActive {
		t.Fatalf("expected valid audience surfaces to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValCAudienceSurfaceCatalog)
		expected string
	}{
		{name: "partner before public still validates actual public breadth", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) {
			model.Surfaces[0], model.Surfaces[1] = model.Surfaces[1], model.Surfaces[0]
		}, expected: VerifierEcosystemValCAudienceSurfaceStateActive},
		{name: "partner before public with insufficient breadth fails closed", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) {
			model.Surfaces[0], model.Surfaces[1] = model.Surfaces[1], model.Surfaces[0]
			model.Surfaces[0].AllowedOutputClasses = append([]string{}, model.Surfaces[1].AllowedOutputClasses...)
		}, expected: VerifierEcosystemValCAudienceSurfaceStatePartial},
		{name: "partner duplicate classes do not count as extra breadth", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) {
			model.Surfaces[1].AllowedOutputClasses = append(append([]string{}, model.Surfaces[0].AllowedOutputClasses...), model.Surfaces[0].AllowedOutputClasses[0], model.Surfaces[0].AllowedOutputClasses[1])
		}, expected: VerifierEcosystemValCAudienceSurfaceStatePartial},
		{name: "partner same unique classes plus duplicates fails closed", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) {
			model.Surfaces[1].AllowedOutputClasses = append(append([]string{}, model.Surfaces[0].AllowedOutputClasses...), model.Surfaces[0].AllowedOutputClasses[0])
		}, expected: VerifierEcosystemValCAudienceSurfaceStatePartial},
		{name: "whitespace output class fails closed", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) {
			model.Surfaces[1].AllowedOutputClasses[0] = " "
		}, expected: VerifierEcosystemValCAudienceSurfaceStateUnknown},
		{name: "unknown output class fails closed", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) {
			model.Surfaces[1].AllowedOutputClasses[0] = "non_verified_unknownish"
		}, expected: VerifierEcosystemValCAudienceSurfaceStateUnknown},
		{name: "duplicate audience type fails closed", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) {
			model.Surfaces[0].AudienceType = VerifierEcosystemValCAudiencePartner
		}, expected: VerifierEcosystemValCAudienceSurfaceStatePartial},
		{name: "unknown audience type fails closed", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) {
			model.Surfaces[0].AudienceType = VerifierEcosystemValCAudienceUnknown
		}, expected: VerifierEcosystemValCAudienceSurfaceStateUnknown},
		{name: "whitespace audience type fails closed", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) {
			model.Surfaces[0].AudienceType = " "
		}, expected: VerifierEcosystemValCAudienceSurfaceStateIncomplete},
		{name: "missing public surface fails closed", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) {
			model.Surfaces = append([]VerifierEcosystemValCAudienceSurface{}, model.Surfaces[1:]...)
		}, expected: VerifierEcosystemValCAudienceSurfaceStatePartial},
		{name: "missing partner surface fails closed", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) {
			model.Surfaces = append(append([]VerifierEcosystemValCAudienceSurface{}, model.Surfaces[:1]...), model.Surfaces[2:]...)
		}, expected: VerifierEcosystemValCAudienceSurfaceStatePartial},
		{name: "public surface without redaction policy fails closed", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) { model.Surfaces[0].RedactionPolicyRef = "" }, expected: VerifierEcosystemValCAudienceSurfaceStateIncomplete},
		{name: "auditor surface without repeatability fails closed", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) {
			model.Surfaces[2].RepeatabilityRequired = false
		}, expected: VerifierEcosystemValCAudienceSurfaceStateBlocked},
		{name: "internal diagnostic cannot be public reusable", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) { model.Surfaces[3].PublicReuseAllowed = true }, expected: VerifierEcosystemValCAudienceSurfaceStateBlocked},
		{name: "publisher self check cannot imply certification", mutate: func(model *VerifierEcosystemValCAudienceSurfaceCatalog) { model.Surfaces[4].CertificationClaim = true }, expected: VerifierEcosystemValCAudienceSurfaceStateBlocked},
	}

	for _, tc := range testCases {
		mutated := verifierEcosystemValCAudienceSurfaceCatalogModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValCAudienceSurfaceState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValCPublicOutput(t *testing.T) {
	audiences := verifierEcosystemValCAudienceSurfaceCatalogModel()
	model := VerifierEcosystemValCPublicOutputContractModel()
	if got := EvaluateVerifierEcosystemValCPublicOutputState(model, audiences); got != VerifierEcosystemValCPublicOutputStateActive {
		t.Fatalf("expected valid public-safe output to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValCPublicOutputContract)
		expected string
	}{
		{name: "restricted trust material in public output fails closed", mutate: func(model *VerifierEcosystemValCPublicOutputContract) { model.SensitiveTrustMaterialExposed = true }, expected: VerifierEcosystemValCPublicOutputStateBlocked},
		{name: "stale state remains visible", mutate: func(model *VerifierEcosystemValCPublicOutputContract) {
			model.OverallResult = VerifierEcosystemValAOverallResultStale
			model.DiagnosticClass = VerifierEcosystemDiagnosticStaleArtifact
			model.OutputClass = VerifierEcosystemValBOutputClassNonVerifiedStale
			model.FreshnessState = IntelligenceCalibrationFreshnessStale
		}, expected: VerifierEcosystemValCPublicOutputStateActive},
		{name: "redaction cannot convert invalid into verified", mutate: func(model *VerifierEcosystemValCPublicOutputContract) {
			model.OverallResult = VerifierEcosystemValAOverallResultInvalid
			model.DiagnosticClass = VerifierEcosystemDiagnosticInvalidSignature
			model.OutputClass = VerifierEcosystemValBOutputClassVerified
		}, expected: VerifierEcosystemValCPublicOutputStateBlocked},
		{name: "certification claim blocks active", mutate: func(model *VerifierEcosystemValCPublicOutputContract) { model.CertificationClaim = true }, expected: VerifierEcosystemValCPublicOutputStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValCPublicOutputContractModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValCPublicOutputState(mutated, audiences); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValCPartnerOutput(t *testing.T) {
	audiences := verifierEcosystemValCAudienceSurfaceCatalogModel()
	model := VerifierEcosystemValCPartnerOutputContractModel()
	if got := EvaluateVerifierEcosystemValCPartnerOutputState(model, audiences); got != VerifierEcosystemValCPartnerOutputStateActive {
		t.Fatalf("expected valid partner-safe output to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValCPartnerOutputContract)
		expected string
	}{
		{name: "missing evidence policy fails closed", mutate: func(model *VerifierEcosystemValCPartnerOutputContract) { model.EvidenceRefPolicy = "" }, expected: VerifierEcosystemValCPartnerOutputStateIncomplete},
		{name: "internal only diagnostics exposed to partner fails closed", mutate: func(model *VerifierEcosystemValCPartnerOutputContract) { model.InternalOnlyDiagnosticsExposed = true }, expected: VerifierEcosystemValCPartnerOutputStateBlocked},
		{name: "partner output cannot mutate or approve", mutate: func(model *VerifierEcosystemValCPartnerOutputContract) { model.MutatesEvidence = true }, expected: VerifierEcosystemValCPartnerOutputStateBlocked},
		{name: "partner output cannot become canonical truth", mutate: func(model *VerifierEcosystemValCPartnerOutputContract) { model.CanonicalTruthClaim = true }, expected: VerifierEcosystemValCPartnerOutputStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValCPartnerOutputContractModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValCPartnerOutputState(mutated, audiences); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValCAuditorFlow(t *testing.T) {
	model := VerifierEcosystemValCAuditorFlowContractModel()
	if got := EvaluateVerifierEcosystemValCAuditorFlowState(model); got != VerifierEcosystemValCAuditorFlowStateActive {
		t.Fatalf("expected valid auditor flow to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValCAuditorFlowContract)
		expected string
	}{
		{name: "missing deterministic report ref fails closed", mutate: func(model *VerifierEcosystemValCAuditorFlowContract) { model.DeterministicReportRef = "" }, expected: VerifierEcosystemValCAuditorFlowStateIncomplete},
		{name: "missing trust material policy fails closed", mutate: func(model *VerifierEcosystemValCAuditorFlowContract) { model.TrustRootMaterialRef = "" }, expected: VerifierEcosystemValCAuditorFlowStateIncomplete},
		{name: "missing evidence refs fails closed", mutate: func(model *VerifierEcosystemValCAuditorFlowContract) { model.RequiredEvidenceRefs = nil }, expected: VerifierEcosystemValCAuditorFlowStateIncomplete},
		{name: "hidden dependency blocks active", mutate: func(model *VerifierEcosystemValCAuditorFlowContract) { model.HiddenMainInstanceDependency = true }, expected: VerifierEcosystemValCAuditorFlowStateBlocked},
		{name: "stale incomplete unsupported diagnostics remain visible", mutate: func(model *VerifierEcosystemValCAuditorFlowContract) {
			model.PreservedDiagnostics = []string{VerifierEcosystemDiagnosticStaleArtifact}
		}, expected: VerifierEcosystemValCAuditorFlowStatePartial},
		{name: "auditor flow does not certify organization", mutate: func(model *VerifierEcosystemValCAuditorFlowContract) { model.CertifiesOrganization = true }, expected: VerifierEcosystemValCAuditorFlowStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValCAuditorFlowContractModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValCAuditorFlowState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValCRequestContract(t *testing.T) {
	testModes := []string{
		VerifierEcosystemValCRequestModeUploadDescriptor,
		VerifierEcosystemValCRequestModeReferenceDescriptor,
		VerifierEcosystemValCRequestModeOfflineBundleDescriptor,
		VerifierEcosystemValCRequestModeAPIReferenceDescriptor,
	}
	for _, mode := range testModes {
		model := VerifierEcosystemValCRequestContractModel()
		model.RequestMode = mode
		if got := EvaluateVerifierEcosystemValCRequestContractState(model); got != VerifierEcosystemValCRequestContractStateActive {
			t.Fatalf("expected request mode %q to be active, got %q", mode, got)
		}
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValCRequestContract)
		expected string
	}{
		{name: "unknown request mode fails closed", mutate: func(model *VerifierEcosystemValCRequestContract) {
			model.RequestMode = VerifierEcosystemValCRequestModeUnknown
		}, expected: VerifierEcosystemValCRequestContractStateUnknown},
		{name: "unsupported artifact type fails closed", mutate: func(model *VerifierEcosystemValCRequestContract) {
			model.AllowedArtifactTypes = []string{"unknown_artifact"}
		}, expected: VerifierEcosystemValCRequestContractStateUnknown},
		{name: "missing required metadata fails closed", mutate: func(model *VerifierEcosystemValCRequestContract) { model.RequiredMetadata = nil }, expected: VerifierEcosystemValCRequestContractStateIncomplete},
		{name: "internal only artifact cannot be accepted for public output", mutate: func(model *VerifierEcosystemValCRequestContract) { model.InternalArtifactForPublic = true }, expected: VerifierEcosystemValCRequestContractStateBlocked},
		{name: "request contract cannot ingest canonical evidence", mutate: func(model *VerifierEcosystemValCRequestContract) { model.IngestsCanonicalEvidence = true }, expected: VerifierEcosystemValCRequestContractStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValCRequestContractModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValCRequestContractState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValCPublisherProfile(t *testing.T) {
	model := VerifierEcosystemValCPublisherCompatibilityProfileModel()
	if got := EvaluateVerifierEcosystemValCPublisherProfileState(model); got != VerifierEcosystemValCPublisherProfileStateActive {
		t.Fatalf("expected valid publisher profile to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValCPublisherCompatibilityProfile)
		expected string
	}{
		{name: "unknown publisher type fails closed", mutate: func(model *VerifierEcosystemValCPublisherCompatibilityProfile) {
			model.PublisherType = VerifierEcosystemValCPublisherTypeUnknown
		}, expected: VerifierEcosystemValCPublisherProfileStateUnknown},
		{name: "missing signature policy fails closed", mutate: func(model *VerifierEcosystemValCPublisherCompatibilityProfile) { model.RequiredSignaturePolicy = "" }, expected: VerifierEcosystemValCPublisherProfileStateIncomplete},
		{name: "approved vendor claim blocks active", mutate: func(model *VerifierEcosystemValCPublisherCompatibilityProfile) { model.ApprovedVendorClaim = true }, expected: VerifierEcosystemValCPublisherProfileStateBlocked},
		{name: "forbidden claims block active", mutate: func(model *VerifierEcosystemValCPublisherCompatibilityProfile) {
			model.ObservedClaims = []string{"approved vendor"}
		}, expected: VerifierEcosystemValCPublisherProfileStateBlocked},
		{name: "verifier compatible does not mean automatically trusted", mutate: func(model *VerifierEcosystemValCPublisherCompatibilityProfile) {
			model.AutomaticallyTrustedClaim = true
		}, expected: VerifierEcosystemValCPublisherProfileStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValCPublisherCompatibilityProfileModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValCPublisherProfileState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValCArtifactRules(t *testing.T) {
	model := VerifierEcosystemValCArtifactPublishingRuleCatalogModel()
	if got := EvaluateVerifierEcosystemValCArtifactRuleState(model); got != VerifierEcosystemValCArtifactRuleStateActive {
		t.Fatalf("expected valid artifact rules to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValCArtifactPublishingRuleCatalog)
		expected string
	}{
		{name: "missing required fields fails closed", mutate: func(model *VerifierEcosystemValCArtifactPublishingRuleCatalog) {
			model.Rules[0].ObservedFields = []string{"proof_type"}
		}, expected: VerifierEcosystemValCArtifactRuleStatePartial},
		{name: "unknown artifact type fails closed", mutate: func(model *VerifierEcosystemValCArtifactPublishingRuleCatalog) {
			model.Rules[0].ArtifactType = "unknown_artifact"
		}, expected: VerifierEcosystemValCArtifactRuleStateUnknown},
		{name: "unsupported schema fails closed", mutate: func(model *VerifierEcosystemValCArtifactPublishingRuleCatalog) {
			model.Rules[0].SchemaVersion = "changelock.verifier.proof_envelope.v9"
		}, expected: VerifierEcosystemValCArtifactRuleStateBlocked},
		{name: "incompatible output boundary fails closed", mutate: func(model *VerifierEcosystemValCArtifactPublishingRuleCatalog) {
			model.Rules[0].SelectedOutputBoundary = VerifierEcosystemScopeAuditorSafe
		}, expected: VerifierEcosystemValCArtifactRuleStateBlocked},
		{name: "forbidden claims block active", mutate: func(model *VerifierEcosystemValCArtifactPublishingRuleCatalog) {
			model.Rules[0].ObservedClaims = []string{"approved_vendor_claim"}
		}, expected: VerifierEcosystemValCArtifactRuleStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValCArtifactPublishingRuleCatalogModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValCArtifactRuleState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValCTrustDistribution(t *testing.T) {
	model := VerifierEcosystemValCTrustDistributionVisibilityModel()
	if got := EvaluateVerifierEcosystemValCTrustDistributionState(model); got != VerifierEcosystemValCTrustDistributionStateActive {
		t.Fatalf("expected valid scoped trust distribution to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*VerifierEcosystemValCTrustDistributionVisibility)
		expected string
	}{
		{name: "unknown distribution mode fails closed", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.TrustRootDistributionMode = VerifierEcosystemValCDistributionModeUnknown
		}, expected: VerifierEcosystemValCTrustDistributionStateUnknown},
		{name: "trusted with warnings plus rollover without metadata remains blocked", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.TrustRootState = VerifierEcosystemTrustRootTrustedWithWarnings
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "trusted with warnings plus rollover metadata remains partial", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.TrustRootState = VerifierEcosystemTrustRootTrustedWithWarnings
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			model.RolloverMetadataRef = "rollover:trust-root-2026.04"
		}, expected: VerifierEcosystemValCTrustDistributionStatePartial},
		{name: "revoked trust material blocks active", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.TrustRootState = VerifierEcosystemTrustRootRevoked
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "trusted rollover without metadata blocks active", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.TrustRootState = VerifierEcosystemTrustRootTrusted
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "revoked trust root with rollover metadata remains blocked", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.TrustRootState = VerifierEcosystemTrustRootRevoked
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			model.RolloverMetadataRef = "rollover:trust-root-2026.04"
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "expired trust root with rollover metadata remains blocked", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.TrustRootState = VerifierEcosystemTrustRootExpired
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			model.RolloverMetadataRef = "rollover:trust-root-2026.04"
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "unsupported trust root with rollover metadata remains blocked", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.TrustRootState = VerifierEcosystemTrustRootUnsupported
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			model.RolloverMetadataRef = "rollover:trust-root-2026.04"
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "unknown trust root with rollover metadata fails closed", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.TrustRootState = VerifierEcosystemTrustRootUnknown
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			model.RolloverMetadataRef = "rollover:trust-root-2026.04"
		}, expected: VerifierEcosystemValCTrustDistributionStateUnknown},
		{name: "typo trust root with rollover metadata fails closed", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.TrustRootState = "trustd"
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			model.RolloverMetadataRef = "rollover:trust-root-2026.04"
		}, expected: VerifierEcosystemValCTrustDistributionStateUnknown},
		{name: "expired trust material blocks active", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.RevocationState = VerifierEcosystemRevocationExpired
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "revoked revocation with rollover metadata remains blocked", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.RevocationState = VerifierEcosystemRevocationRevoked
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			model.RolloverMetadataRef = "rollover:trust-root-2026.04"
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "expired revocation with rollover metadata remains blocked", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.RevocationState = VerifierEcosystemRevocationExpired
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			model.RolloverMetadataRef = "rollover:trust-root-2026.04"
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "unsupported revocation with rollover metadata remains blocked", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.RevocationState = VerifierEcosystemRevocationUnsupported
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
			model.RolloverMetadataRef = "rollover:trust-root-2026.04"
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "rollover without metadata blocks active", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.KeyRotationState = VerifierEcosystemKeyRotationRollover
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "offline distribution must be scoped", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.TrustRootDistributionMode = VerifierEcosystemValCDistributionModeOfflineBundle
			model.AudienceVisibilityScope = VerifierEcosystemScopePublicSafe
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "global all instance key directory claim blocks active", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) { model.GlobalKeyDirectoryClaim = true }, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
		{name: "public output cannot expose sensitive key material", mutate: func(model *VerifierEcosystemValCTrustDistributionVisibility) {
			model.AudienceVisibilityScope = VerifierEcosystemScopePublicSafe
			model.SensitiveKeyMaterialExposed = true
		}, expected: VerifierEcosystemValCTrustDistributionStateBlocked},
	}

	for _, tc := range testCases {
		mutated := VerifierEcosystemValCTrustDistributionVisibilityModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValCTrustDistributionState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestVerifierEcosystemValCAggregateStatePrecedence(t *testing.T) {
	dependency, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, valCState := activeVerifierEcosystemValCStates()
	if valCState != VerifierEcosystemValCStateActive {
		t.Fatalf("expected active baseline Val C state, got %q", valCState)
	}

	testCases := []struct {
		name      string
		audience  string
		public    string
		partner   string
		auditor   string
		request   string
		publisher string
		rules     string
		trust     string
		expected  string
	}{
		{name: "partial before blocked returns blocked", audience: VerifierEcosystemValCAudienceSurfaceStatePartial, rules: VerifierEcosystemValCArtifactRuleStateBlocked, expected: VerifierEcosystemValCStateBlocked},
		{name: "blocked before partial returns blocked", audience: VerifierEcosystemValCAudienceSurfaceStateBlocked, public: VerifierEcosystemValCPublicOutputStatePartial, expected: VerifierEcosystemValCStateBlocked},
		{name: "partial plus unknown returns unknown", audience: VerifierEcosystemValCAudienceSurfaceStatePartial, request: VerifierEcosystemValCRequestContractStateUnknown, expected: VerifierEcosystemValCStateUnknown},
		{name: "incomplete plus partial returns incomplete", audience: VerifierEcosystemValCAudienceSurfaceStateIncomplete, public: VerifierEcosystemValCPublicOutputStatePartial, expected: VerifierEcosystemValCStateIncomplete},
		{name: "fake component state fails closed", audience: "not_a_real_valc_audience_state", expected: VerifierEcosystemValCStateUnknown},
	}

	for _, tc := range testCases {
		audience := VerifierEcosystemValCAudienceSurfaceStateActive
		public := VerifierEcosystemValCPublicOutputStateActive
		partner := VerifierEcosystemValCPartnerOutputStateActive
		auditor := VerifierEcosystemValCAuditorFlowStateActive
		request := VerifierEcosystemValCRequestContractStateActive
		publisher := VerifierEcosystemValCPublisherProfileStateActive
		rules := VerifierEcosystemValCArtifactRuleStateActive
		trust := VerifierEcosystemValCTrustDistributionStateActive
		if tc.audience != "" {
			audience = tc.audience
		}
		if tc.public != "" {
			public = tc.public
		}
		if tc.partner != "" {
			partner = tc.partner
		}
		if tc.auditor != "" {
			auditor = tc.auditor
		}
		if tc.request != "" {
			request = tc.request
		}
		if tc.publisher != "" {
			publisher = tc.publisher
		}
		if tc.rules != "" {
			rules = tc.rules
		}
		if tc.trust != "" {
			trust = tc.trust
		}
		if got := EvaluateVerifierEcosystemValCState(dependency, audience, public, partner, auditor, request, publisher, rules, trust); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}

	missingDependency := dependency
	missingDependency.ValBState = VerifierEcosystemValBStatePartial
	if got := EvaluateVerifierEcosystemValCState(missingDependency, VerifierEcosystemValCAudienceSurfaceStateActive, VerifierEcosystemValCPublicOutputStateActive, VerifierEcosystemValCPartnerOutputStateActive, VerifierEcosystemValCAuditorFlowStateActive, VerifierEcosystemValCRequestContractStateActive, VerifierEcosystemValCPublisherProfileStateActive, VerifierEcosystemValCArtifactRuleStateActive, VerifierEcosystemValCTrustDistributionStateActive); got != VerifierEcosystemValCStateBlocked {
		t.Fatalf("expected missing dependency to return %q, got %q", VerifierEcosystemValCStateBlocked, got)
	}
}

func TestVerifierEcosystemValCProofSurfaceCompletenessAndPoint7PassImpossibility(t *testing.T) {
	dependency, _, _, _, _, _, _, _, _, audienceState, publicState, partnerState, auditorState, requestState, publisherState, artifactRuleState, trustDistributionState, valCState := activeVerifierEcosystemValCStates()
	point7State := EvaluateVerifierEcosystemValCPoint7State(valCState)
	if point7State != VerifierEcosystemPoint7StateNotComplete {
		t.Fatalf("expected point 7 to remain not complete in Val C, got %q", point7State)
	}
	if valCState != VerifierEcosystemValCStateActive ||
		audienceState != VerifierEcosystemValCAudienceSurfaceStateActive ||
		publicState != VerifierEcosystemValCPublicOutputStateActive ||
		partnerState != VerifierEcosystemValCPartnerOutputStateActive ||
		auditorState != VerifierEcosystemValCAuditorFlowStateActive ||
		requestState != VerifierEcosystemValCRequestContractStateActive ||
		publisherState != VerifierEcosystemValCPublisherProfileStateActive ||
		artifactRuleState != VerifierEcosystemValCArtifactRuleStateActive ||
		trustDistributionState != VerifierEcosystemValCTrustDistributionStateActive {
		t.Fatalf("expected active Val C baseline component states")
	}

	surfaces := VerifierEcosystemValCProofSurfaceRefs()
	evidenceRefs := VerifierEcosystemValCProofEvidenceRefs()
	limitations := []string{
		"Val C implements bounded public, partner, auditor, publisher, request, and trust distribution contracts only.",
	}
	reasons := []string{
		"Val C cannot return point_7_pass.",
		"Val D and Val E remain required for final Točka 7 closure.",
	}
	if got := EvaluateVerifierEcosystemValCProofsState(valCState, VerifierEcosystemPoint7StateNotComplete, dependency.Val0CurrentState, dependency.ValACurrentState, dependency.ValBCurrentState, surfaces, evidenceRefs, limitations, reasons, verifierEcosystemValCProjectionDisclaimer()); got != VerifierEcosystemValCStateActive {
		t.Fatalf("expected complete Val C proofs to remain active, got %q", got)
	}

	testCases := []struct {
		name     string
		surfaces []string
		evidence []string
		expected string
	}{
		{name: "missing surface fails closed", surfaces: removeTrimmedString(surfaces, "/v1/verifier-ecosystem/valc/proofs"), evidence: evidenceRefs, expected: VerifierEcosystemValCStatePartial},
		{name: "duplicate surface does not compensate", surfaces: append(removeTrimmedString(surfaces, "/v1/verifier-ecosystem/valc/proofs"), "/v1/verifier-ecosystem/valc/trust-distribution"), evidence: evidenceRefs, expected: VerifierEcosystemValCStatePartial},
		{name: "unknown extra surface does not compensate", surfaces: append(append([]string{}, surfaces...), "/v1/verifier-ecosystem/valc/extra"), evidence: evidenceRefs, expected: VerifierEcosystemValCStatePartial},
		{name: "whitespace surface ref fails closed", surfaces: append(removeTrimmedString(surfaces, "/v1/verifier-ecosystem/valc/proofs"), " "), evidence: evidenceRefs, expected: VerifierEcosystemValCStatePartial},
		{name: "missing evidence ref fails closed", surfaces: surfaces, evidence: removeTrimmedString(evidenceRefs, "evidence:audience-surfaces-001"), expected: VerifierEcosystemValCStatePartial},
		{name: "duplicate evidence ref does not compensate", surfaces: surfaces, evidence: append(removeTrimmedString(evidenceRefs, "evidence:audience-surfaces-001"), "evidence:trust-distribution-001"), expected: VerifierEcosystemValCStatePartial},
		{name: "unknown extra evidence ref does not compensate", surfaces: surfaces, evidence: append(append([]string{}, evidenceRefs...), "evidence:unknown-extra"), expected: VerifierEcosystemValCStatePartial},
		{name: "whitespace evidence ref fails closed", surfaces: surfaces, evidence: append(removeTrimmedString(evidenceRefs, "evidence:audience-surfaces-001"), " "), expected: VerifierEcosystemValCStatePartial},
	}

	for _, tc := range testCases {
		if got := EvaluateVerifierEcosystemValCProofsState(
			valCState,
			VerifierEcosystemPoint7StateNotComplete,
			dependency.Val0CurrentState,
			dependency.ValACurrentState,
			dependency.ValBCurrentState,
			tc.surfaces,
			tc.evidence,
			limitations,
			reasons,
			verifierEcosystemValCProjectionDisclaimer(),
		); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}
