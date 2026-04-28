package operability

import (
	"os"
	"strings"
	"testing"
)

func activeVerifierEcosystemValEAudienceFacts(model VerifierEcosystemValCAudienceSurfaceCatalog) (bool, bool) {
	publicFound := false
	partnerFound := false
	publicCount := 0
	partnerCount := 0
	for _, item := range model.Surfaces {
		switch strings.TrimSpace(item.AudienceType) {
		case VerifierEcosystemValCAudiencePublic:
			count, ok := verifierEcosystemValCNormalizedUniqueOutputClassCount(item.AllowedOutputClasses)
			if !ok {
				return false, false
			}
			publicCount = count
			publicFound = true
		case VerifierEcosystemValCAudiencePartner:
			count, ok := verifierEcosystemValCNormalizedUniqueOutputClassCount(item.AllowedOutputClasses)
			if !ok {
				return false, false
			}
			partnerCount = count
			partnerFound = true
		}
	}
	if !publicFound || !partnerFound {
		return false, false
	}
	return true, partnerCount > publicCount
}

func activeVerifierEcosystemValEModel() VerifierEcosystemIntegratedClosure {
	_, _, _, _, _, trustModel, _, _, contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState, val0State := activeVerifierEcosystemVal0States()
	_, _, engine, _, _, _, _, inputState, engineState, resultState, diagnosticsMappingState, commandState, sdkState, valAState := activeVerifierEcosystemValAStates()
	_, _, schemaProofCompatibility, _, _, _, _, conformanceSuite, _, matrixState, schemaCompatibilityState, mixedState, precedenceState, fixtureState, caseState, suiteState, outputClassState, valBState := activeVerifierEcosystemValBStates()
	valCDependency, audiences, publicOutput, partnerOutput, _, requestContract, publisherProfile, _, trustDistribution, audienceState, publicState, partnerState, auditorState, requestState, publisherState, artifactRuleState, trustDistributionState, valCState := activeVerifierEcosystemValCStates()
	valDDependency, _, tooling, _, _, trustKeyRotation, negativeDiagnostics, _, _, _, correctnessState, toolingState, schemaGateState, diagnosticsConformanceState, trustGateState, negativeDiagnosticsState, redactionState, publisherArtifactState, noOverclaimState, valDState := activeVerifierEcosystemValDStates()
	audienceOrderIndependent, uniqueBreadthValid := activeVerifierEcosystemValEAudienceFacts(audiences)

	model := VerifierEcosystemIntegratedClosure{
		ClosureID:        "verifier-ecosystem-point-7-closure",
		Version:          "2026.04",
		Point:            "point_7",
		ClosureVal:       "val_e",
		Point7PassReason: "point_7_pass through Val E only after actual Val 0 through Val D proof states, exact proof surfaces, exact evidence refs, and cross-val closure invariants all remain active and fail-closed.",
		SourceValStates: VerifierEcosystemValESourceValStates{
			Val0State: val0State,
			ValAState: valAState,
			ValBState: valBState,
			ValCState: valCState,
			ValDState: valDState,
		},
		SourceCurrentStates: VerifierEcosystemValESourceCurrentStates{
			Val0CurrentState: val0State,
			ValACurrentState: valAState,
			ValBCurrentState: valBState,
			ValCCurrentState: valCState,
			ValDCurrentState: valDState,
		},
		DependencyStates: VerifierEcosystemValEDependencyStates{
			Point5State:                    valDDependency.Point5State,
			Point5DependencyState:          valDDependency.Point5DependencyState,
			Point6State:                    valDDependency.Point6State,
			Point6ClosureState:             valDDependency.Point6ClosureState,
			Point6ClosurePrerequisiteState: valDDependency.Point6ClosurePrerequisiteState,
			Point6ClosureInvariantState:    valDDependency.Point6ClosureInvariantState,
			Point6ProofSurfaceState:        valDDependency.Point6ProofSurfaceState,
			Point6PassRuleState:            valDDependency.Point6PassRuleState,
			Point6PassAllowed:              valDDependency.Point6PassAllowed,
			ValDFinalGateState:             valDState,
			PreClosurePoint7State:          VerifierEcosystemPoint7StateNotComplete,
		},
		Val0: VerifierEcosystemValEVal0ProofSnapshot{
			CurrentState:             val0State,
			Val0State:                val0State,
			Point7State:              VerifierEcosystemPoint7StateNotComplete,
			VerifierContractState:    contractState,
			ProofEnvelopeState:       envelopeState,
			VerificationScopeState:   scopeState,
			SchemaCompatibilityState: compatibilityState,
			TrustRootIssuerState:     trustState,
			DiagnosticsState:         diagnosticsState,
			OutputBoundaryState:      outputBoundaryState,
			TrustRootState:           trustModel.TrustRootState,
			RevocationState:          trustModel.RevocationState,
			KeyRotationState:         trustModel.KeyRotationState,
			RolloverMetadataRef:      trustModel.RolloverMetadataRef,
			SurfaceRefs:              VerifierEcosystemVal0ProofSurfaceRefs(),
			EvidenceRefs:             VerifierEcosystemVal0ProofEvidenceRefs(),
			ProjectionDisclaimer:     verifierEcosystemVal0ProjectionDisclaimer(),
			WorstSeverityPrecedence:  true,
		},
		ValA: VerifierEcosystemValEValAProofSnapshot{
			CurrentState:                 valAState,
			ValAState:                    valAState,
			Point7State:                  VerifierEcosystemPoint7StateNotComplete,
			InputModelState:              inputState,
			VerifierEngineState:          engineState,
			VerificationResultState:      resultState,
			DiagnosticsMappingState:      diagnosticsMappingState,
			CommandContractState:         commandState,
			SDKEntrypointState:           sdkState,
			DeterministicOutput:          engine.DeterministicOutput,
			HiddenMainInstanceDependency: false,
			NetworkDependency:            engine.NetworkDependency,
			MutatesEvidence:              engine.MutatesEvidence,
			ApprovesDeployment:           engine.ClaimsDeploymentApproval,
			SuppressesFailures:           false,
			ClaimsActualCryptoValidity:   engine.ClaimsActualCryptoValidity,
			UsesRealCryptoPrimitives:     engine.UsesRealCryptoPrimitives,
			SurfaceRefs:                  VerifierEcosystemValAProofSurfaceRefs(),
			EvidenceRefs:                 verifierEcosystemValAExpectedProofEvidenceRefs(),
			ProjectionDisclaimer:         verifierEcosystemValAProjectionDisclaimer(),
			WorstSeverityPrecedence:      true,
		},
		ValB: VerifierEcosystemValEValBProofSnapshot{
			CurrentState:                  valBState,
			ValBState:                     valBState,
			Point7State:                   VerifierEcosystemPoint7StateNotComplete,
			CompatibilityMatrixState:      matrixState,
			SchemaProofCompatibilityState: schemaCompatibilityState,
			MixedVersionDiagnosticState:   mixedState,
			DiagnosticPrecedenceState:     precedenceState,
			FixtureDescriptorState:        fixtureState,
			ConformanceCaseState:          caseState,
			ConformanceSuiteState:         suiteState,
			OutputClassState:              outputClassState,
			CompatibilityState:            schemaProofCompatibility.CompatibilityState,
			DerivedDiagnosticClass:        VerifierEcosystemDiagnosticVerified,
			DerivedOutputClass:            VerifierEcosystemValBOutputClassVerified,
			NegativeCasesPreserved:        true,
			ConformanceCertificationClaim: conformanceSuite.CertificationClaim,
			IntegrityRatingClaim:          false,
			SurfaceRefs:                   VerifierEcosystemValBProofSurfaceRefs(),
			EvidenceRefs:                  VerifierEcosystemValBProofEvidenceRefs(),
			ProjectionDisclaimer:          verifierEcosystemValBProjectionDisclaimer(),
			WorstSeverityPrecedence:       true,
		},
		ValC: VerifierEcosystemValEValCProofSnapshot{
			CurrentState:                          valCState,
			ValCState:                             valCState,
			Point7State:                           valCDependency.Point7State,
			AudienceSurfaceState:                  audienceState,
			PublicOutputState:                     publicState,
			PartnerOutputState:                    partnerState,
			AuditorFlowState:                      auditorState,
			RequestContractState:                  requestState,
			PublisherProfileState:                 publisherState,
			ArtifactRuleState:                     artifactRuleState,
			TrustDistributionState:                trustDistributionState,
			PublicOutputClass:                     publicOutput.OutputClass,
			PartnerOutputClass:                    partnerOutput.OutputClass,
			RequestMode:                           requestContract.RequestMode,
			PublisherType:                         publisherProfile.PublisherType,
			TrustDistributionMode:                 trustDistribution.TrustRootDistributionMode,
			AudienceUniqueBreadthValid:            uniqueBreadthValid,
			AudienceBreadthOrderIndependent:       audienceOrderIndependent,
			PublisherApprovedVendorClaim:          publisherProfile.ApprovedVendorClaim,
			PublisherCertificationClaim:           false,
			PublisherAutomaticallyTrustedClaim:    publisherProfile.AutomaticallyTrustedClaim,
			TrustDistributionGlobalDirectoryClaim: trustDistribution.GlobalKeyDirectoryClaim,
			TrustDistributionKeyRotationState:     trustDistribution.KeyRotationState,
			TrustDistributionRolloverMetadataRef:  trustDistribution.RolloverMetadataRef,
			TrustDistributionTrustRootState:       trustDistribution.TrustRootState,
			TrustDistributionRevocationState:      trustDistribution.RevocationState,
			SurfaceRefs:                           VerifierEcosystemValCProofSurfaceRefs(),
			EvidenceRefs:                          VerifierEcosystemValCProofEvidenceRefs(),
			ProjectionDisclaimer:                  verifierEcosystemValCProjectionDisclaimer(),
		},
		ValD: VerifierEcosystemValEValDProofSnapshot{
			CurrentState:                            valDState,
			ValDState:                               valDState,
			Point7State:                             valDDependency.Point7State,
			CorrectnessGateState:                    correctnessState,
			ToolingGateState:                        toolingState,
			SchemaCompatibilityGateState:            schemaGateState,
			DiagnosticsConformanceGateState:         diagnosticsConformanceState,
			TrustKeyRotationGateState:               trustGateState,
			NegativeDiagnosticsGateState:            negativeDiagnosticsState,
			RedactionGateState:                      redactionState,
			PublisherArtifactGateState:              publisherArtifactState,
			NoOverclaimGateState:                    noOverclaimState,
			TrustDistributionMode:                   trustKeyRotation.TrustDistributionMode,
			OfflineDistributionScope:                trustKeyRotation.OfflineDistributionScope,
			TrustDistributionModeUsesActualValCMode: trustKeyRotation.TrustDistributionMode == trustDistribution.TrustRootDistributionMode && trustKeyRotation.TrustDistributionMode != trustKeyRotation.OfflineDistributionScope,
			ClaimsIntegratedClosure:                 false,
			SurfaceRefs:                             VerifierEcosystemValDProofSurfaceRefs(),
			EvidenceRefs:                            VerifierEcosystemValDProofEvidenceRefs(),
			ProjectionDisclaimer:                    verifierEcosystemValDProjectionDisclaimer(),
		},
		ProofSurfaceRefs:              VerifierEcosystemValEProofSurfaceRefs(),
		EvidenceRefs:                  VerifierEcosystemValEProofEvidenceRefs(),
		ProjectionDisclaimer:          verifierEcosystemValEProjectionDisclaimer(),
		ObservedClaims:                []string{},
		Caveats:                       []string{"Integrated verifier ecosystem closure remains advisory only."},
		Limitations:                   []string{"Val E closes Točka 7 only and does not start Točka 8."},
		CreatedAt:                     "2026-04-28T00:20:00Z",
		UpdatedAt:                     "2026-04-28T00:20:00Z",
		EvidenceFresh:                 true,
		StaleEvidenceDetected:         false,
		RedactionKeepsFailuresVisible: negativeDiagnostics.RedactionBoundaryPreserved && negativeDiagnostics.PublicPreservesNonVerified && negativeDiagnostics.AuditorRepeatable && negativeDiagnostics.AuditorEvidenceLinked,
		MutatesCanonicalEvidence:      tooling.MutatesEvidence || requestContract.IngestsCanonicalEvidence,
		ApprovesDeployment:            tooling.ApprovesDeployment || publicOutput.DeploymentApprovalClaim || partnerOutput.ApprovesDeployment,
		SuppressesFailures:            partnerOutput.SuppressesFailures,
	}
	return ComputeVerifierEcosystemValEClosure(model)
}

func TestVerifierEcosystemValEDependencyGates(t *testing.T) {
	model := activeVerifierEcosystemValEModel()
	if model.CurrentState != VerifierEcosystemValEStatePass || model.Point7State != VerifierEcosystemPoint7StatePass || !model.Point7PassAllowed {
		t.Fatalf("expected active Val E closure to return point_7_pass, got %#v", model)
	}

	testCases := []struct {
		name   string
		mutate func(*VerifierEcosystemIntegratedClosure)
	}{
		{name: "missing val0 blocks vale", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.SourceValStates.Val0State = VerifierEcosystemVal0StatePartial
		}},
		{name: "missing vala blocks vale", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.SourceValStates.ValAState = VerifierEcosystemValAStatePartial
		}},
		{name: "missing valb blocks vale", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.SourceValStates.ValBState = VerifierEcosystemValBStatePartial
		}},
		{name: "missing valc blocks vale", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.SourceValStates.ValCState = VerifierEcosystemValCStatePartial
		}},
		{name: "missing vald blocks vale", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.SourceValStates.ValDState = VerifierEcosystemValDStatePartial
		}},
		{name: "missing vald final gate blocks pass", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.DependencyStates.ValDFinalGateState = VerifierEcosystemValDStatePartial
		}},
		{name: "missing point6 closure blocks pass", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.DependencyStates.Point6ClosureState = ReferenceArchitectureValEStatePartial
		}},
		{name: "pre closure point7 already pass blocks pass", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.DependencyStates.PreClosurePoint7State = VerifierEcosystemPoint7StatePass
		}},
	}

	for _, tc := range testCases {
		mutated := activeVerifierEcosystemValEModel()
		tc.mutate(&mutated)
		mutated = ComputeVerifierEcosystemValEClosure(mutated)
		if mutated.CurrentState == VerifierEcosystemValEStatePass || mutated.Point7State == VerifierEcosystemPoint7StatePass || mutated.Point7PassAllowed {
			t.Fatalf("expected %s to block Val E pass, got %#v", tc.name, mutated)
		}
	}
}

func TestVerifierEcosystemValEPassRuleAndInvariants(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*VerifierEcosystemIntegratedClosure)
	}{
		{name: "val0 revoked issuer blocks closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.Val0.RevocationState = VerifierEcosystemRevocationRevoked
			model.Val0.TrustRootIssuerState = VerifierEcosystemVal0TrustStateActive
		}},
		{name: "val0 expired trust root blocks closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.Val0.TrustRootState = VerifierEcosystemTrustRootExpired
			model.Val0.TrustRootIssuerState = VerifierEcosystemVal0TrustStatePartial
		}},
		{name: "val0 rollover missing metadata blocks closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.Val0.KeyRotationState = VerifierEcosystemKeyRotationRollover
			model.Val0.RolloverMetadataRef = ""
		}},
		{name: "val0 count only evidence fails closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.Val0.EvidenceRefs = append(removeTrimmedString(model.Val0.EvidenceRefs, VerifierEcosystemVal0ProofEvidenceRefs()[0]), VerifierEcosystemVal0ProofEvidenceRefs()[1])
		}},
		{name: "vala engine guardrail blocks closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ValA.VerifierEngineState = VerifierEcosystemValAEngineStateBlocked
		}},
		{name: "vala fake crypto claim blocks closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ValA.ClaimsActualCryptoValidity = true
			model.ValA.UsesRealCryptoPrimitives = false
		}},
		{name: "valb unsupported schema blocks closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ValB.CompatibilityState = ReferenceArchitectureCompatibilityUnsupported
		}},
		{name: "valb conformance certification claim blocks closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ValB.ConformanceCertificationClaim = true
		}},
		{name: "valc publisher approved vendor claim blocks closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ValC.PublisherApprovedVendorClaim = true
		}},
		{name: "valc duplicate breadth cannot satisfy closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ValC.AudienceUniqueBreadthValid = false
		}},
		{name: "valc trust distribution missing rollover metadata blocks closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ValC.TrustDistributionKeyRotationState = VerifierEcosystemKeyRotationRollover
			model.ValC.TrustDistributionRolloverMetadataRef = ""
		}},
		{name: "vald trust distribution mode from scope blocks closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ValD.TrustDistributionModeUsesActualValCMode = false
			model.ValD.TrustDistributionMode = VerifierEcosystemScopePartnerSafe
		}},
		{name: "vald no overclaim gate blocked blocks closure", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ValD.NoOverclaimGateState = VerifierEcosystemValDNoOverclaimGateStateBlocked
		}},
		{name: "partial inputs cannot reach pass", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ValD.SchemaCompatibilityGateState = VerifierEcosystemValDSchemaCompatibilityGateStatePartial
		}},
		{name: "stale evidence cannot reach pass", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.EvidenceFresh = false
			model.StaleEvidenceDetected = true
		}},
		{name: "unknown state cannot reach pass", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ValB.CompatibilityState = ReferenceArchitectureCompatibilityUnknown
		}},
	}

	for _, tc := range testCases {
		mutated := activeVerifierEcosystemValEModel()
		tc.mutate(&mutated)
		mutated = ComputeVerifierEcosystemValEClosure(mutated)
		if mutated.CurrentState == VerifierEcosystemValEStatePass || mutated.Point7State == VerifierEcosystemPoint7StatePass {
			t.Fatalf("expected %s to block point_7_pass, got %#v", tc.name, mutated)
		}
		if len(mutated.BlockingReasons) == 0 {
			t.Fatalf("expected %s to surface blocking reasons", tc.name)
		}
	}
}

func TestVerifierEcosystemValEProofSurfaceExactSet(t *testing.T) {
	model := activeVerifierEcosystemValEModel()
	if got := EvaluateVerifierEcosystemValEProofSurfaceState(model); got != VerifierEcosystemValEProofSurfaceStateActive {
		t.Fatalf("expected exact Val E proof surfaces to be active, got %q", got)
	}

	testCases := []struct {
		name   string
		mutate func(*VerifierEcosystemIntegratedClosure)
	}{
		{name: "missing val0 proofs blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/val0/proofs")
		}},
		{name: "missing vala proofs blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/vala/proofs")
		}},
		{name: "missing valb proofs blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/valb/proofs")
		}},
		{name: "missing valc proofs blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/valc/proofs")
		}},
		{name: "missing vald proofs blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/vald/proofs")
		}},
		{name: "missing vald gate surface blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/vald/no-overclaim-gate")
		}},
		{name: "missing vale closure blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/vale/closure")
		}},
		{name: "missing vale proofs blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/vale/proofs")
		}},
		{name: "duplicate does not compensate", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/vale/proofs")
			model.ProofSurfaceRefs = append(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/vald/proofs")
		}},
		{name: "unknown extra does not compensate", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ProofSurfaceRefs = append(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/vale/unexpected")
		}},
		{name: "whitespace surface ref fails closed", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.ProofSurfaceRefs[0] = " "
		}},
	}

	for _, tc := range testCases {
		mutated := activeVerifierEcosystemValEModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValEProofSurfaceState(mutated); got == VerifierEcosystemValEProofSurfaceStateActive {
			t.Fatalf("expected %s to fail closed, got %q", tc.name, got)
		}
	}
}

func TestVerifierEcosystemValEEvidenceQualityChecks(t *testing.T) {
	model := activeVerifierEcosystemValEModel()
	if got := EvaluateVerifierEcosystemValEEvidenceQualityState(model); got != VerifierEcosystemValEEvidenceQualityStateActive {
		t.Fatalf("expected exact Val E evidence refs to be active, got %q", got)
	}

	testCases := []struct {
		name   string
		mutate func(*VerifierEcosystemIntegratedClosure)
	}{
		{name: "missing val0 evidence blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "point7_verifier_discipline_foundation")
		}},
		{name: "missing vala evidence blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "point7_reference_verifier_tooling")
		}},
		{name: "missing valb evidence blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "point7_compatibility_diagnostics_conformance")
		}},
		{name: "missing valc evidence blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "point7_public_partner_auditor_publisher_ecosystem")
		}},
		{name: "missing vald evidence blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "point7_final_verifier_ecosystem_gate")
		}},
		{name: "missing no overclaim evidence blocks", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "point7_no_overclaim_governance")
		}},
		{name: "duplicate does not compensate", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "evidence:vale-point7-governance-001")
			model.EvidenceRefs = append(model.EvidenceRefs, "evidence:vale-closure-001")
		}},
		{name: "unknown extra does not compensate", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.EvidenceRefs = append(model.EvidenceRefs, "evidence:vale-unexpected-001")
		}},
		{name: "whitespace evidence ref fails closed", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.EvidenceRefs[0] = " "
		}},
		{name: "stale evidence blocks point7 pass", mutate: func(model *VerifierEcosystemIntegratedClosure) {
			model.EvidenceFresh = false
			model.StaleEvidenceDetected = true
		}},
	}

	for _, tc := range testCases {
		mutated := activeVerifierEcosystemValEModel()
		tc.mutate(&mutated)
		if got := EvaluateVerifierEcosystemValEEvidenceQualityState(mutated); got == VerifierEcosystemValEEvidenceQualityStateActive {
			t.Fatalf("expected %s to fail closed, got %q", tc.name, got)
		}
		mutated = ComputeVerifierEcosystemValEClosure(mutated)
		if mutated.Point7State == VerifierEcosystemPoint7StatePass {
			t.Fatalf("expected %s to block point_7_pass", tc.name)
		}
	}
}

func TestVerifierEcosystemValENoOverclaim(t *testing.T) {
	testCases := []string{
		"verifier certification",
		"certified publisher",
		"approved vendor",
		"integrity rating",
		"global key registry for all instances",
		"deployment approved",
		"production approved",
	}

	for _, claim := range testCases {
		mutated := activeVerifierEcosystemValEModel()
		mutated.ObservedClaims = append(mutated.ObservedClaims, claim)
		mutated = ComputeVerifierEcosystemValEClosure(mutated)
		if mutated.NoOverclaimState != VerifierEcosystemValENoOverclaimStateBlocked || mutated.Point7State == VerifierEcosystemPoint7StatePass {
			t.Fatalf("expected claim %q to block no-overclaim state, got %#v", claim, mutated)
		}
	}
}

func TestVerifierEcosystemValEFinalStateBehavior(t *testing.T) {
	model := activeVerifierEcosystemValEModel()
	if model.CurrentState != VerifierEcosystemValEStatePass || model.Point7State != VerifierEcosystemPoint7StatePass || !model.Point7PassAllowed {
		t.Fatalf("expected Val E pass state, got %#v", model)
	}

	mutated := activeVerifierEcosystemValEModel()
	mutated.ValD.ToolingGateState = VerifierEcosystemValDToolingGateStatePartial
	mutated = ComputeVerifierEcosystemValEClosure(mutated)
	if mutated.CurrentState == VerifierEcosystemValEStatePass || mutated.Point7State == VerifierEcosystemPoint7StatePass {
		t.Fatalf("expected partial closure to remain non-pass, got %#v", mutated)
	}
	if len(mutated.BlockingReasons) == 0 || len(mutated.Caveats) == 0 || len(mutated.Limitations) == 0 {
		t.Fatalf("expected non-pass closure to expose blocking reasons, caveats, and limitations, got %#v", mutated)
	}
}

func TestVerifierEcosystemValEDocLanguage(t *testing.T) {
	content, err := os.ReadFile("../../docs/verifier-ecosystem-vale-core.md")
	if err != nil {
		t.Fatalf("read doc: %v", err)
	}
	lower := strings.ToLower(string(content))
	disallowed := []string{
		"certified verifier",
		"certified publisher",
		"certified vendor",
		"approved vendor",
		"integrity rating",
		"absolute proof",
		"universal authority",
		"točka 8 started",
	}
	for _, phrase := range disallowed {
		if strings.Contains(lower, phrase) {
			t.Fatalf("expected doc to avoid disallowed phrase %q", phrase)
		}
	}
}
