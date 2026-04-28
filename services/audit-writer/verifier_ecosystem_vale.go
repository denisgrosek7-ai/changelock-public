package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	verifierEcosystemValEClosureSchema = "point7.verifier_ecosystem.vale.closure.v1"
	verifierEcosystemValEProofsSchema  = "point7.verifier_ecosystem.vale.proofs.v1"
)

type verifierEcosystemValEClosureResponse struct {
	SchemaVersion string                                         `json:"schema_version"`
	GeneratedAt   time.Time                                      `json:"generated_at"`
	CurrentState  string                                         `json:"current_state"`
	Model         operability.VerifierEcosystemIntegratedClosure `json:"model"`
	RouteRefs     []string                                       `json:"route_refs,omitempty"`
	Limitations   []string                                       `json:"limitations,omitempty"`
}

type verifierEcosystemValEProofsResponse struct {
	SchemaVersion             string                                          `json:"schema_version"`
	GeneratedAt               time.Time                                       `json:"generated_at"`
	CurrentState              string                                          `json:"current_state"`
	Point5State               string                                          `json:"point_5_state"`
	Point5DependencyState     string                                          `json:"point_5_dependency_state"`
	Point6State               string                                          `json:"point_6_state"`
	Point6ClosureState        string                                          `json:"point_6_closure_state"`
	Point6ClosurePrerequisite string                                          `json:"point_6_closure_prerequisite_state"`
	Point6ClosureInvariant    string                                          `json:"point_6_closure_invariant_state"`
	Point6ProofSurfaceState   string                                          `json:"point_6_proof_surface_state"`
	Point6PassRuleState       string                                          `json:"point_6_pass_rule_state"`
	Point6PassAllowed         bool                                            `json:"point_6_pass_allowed"`
	Val0CurrentState          string                                          `json:"val_0_current_state"`
	Val0State                 string                                          `json:"val_0_state"`
	ValACurrentState          string                                          `json:"val_a_current_state"`
	ValAState                 string                                          `json:"val_a_state"`
	ValBCurrentState          string                                          `json:"val_b_current_state"`
	ValBState                 string                                          `json:"val_b_state"`
	ValCCurrentState          string                                          `json:"val_c_current_state"`
	ValCState                 string                                          `json:"val_c_state"`
	ValDCurrentState          string                                          `json:"val_d_current_state"`
	ValDState                 string                                          `json:"val_d_state"`
	ValDFinalGateState        string                                          `json:"val_d_final_gate_state"`
	ClosurePrerequisiteState  string                                          `json:"closure_prerequisite_state"`
	ClosureInvariantState     string                                          `json:"closure_invariant_state"`
	ProofSurfaceState         string                                          `json:"proof_surface_state"`
	EvidenceQualityState      string                                          `json:"evidence_quality_state"`
	NoOverclaimState          string                                          `json:"no_overclaim_state"`
	PassRuleState             string                                          `json:"pass_rule_state"`
	ValEState                 string                                          `json:"val_e_state"`
	Point7State               string                                          `json:"point_7_state"`
	Point7PassAllowed         bool                                            `json:"point_7_pass_allowed"`
	Point7PassReason          string                                          `json:"point_7_pass_reason"`
	ClosureInvariants         []operability.VerifierEcosystemClosureInvariant `json:"closure_invariants,omitempty"`
	BlockingReasons           []string                                        `json:"blocking_reasons,omitempty"`
	Caveats                   []string                                        `json:"caveats,omitempty"`
	Limitations               []string                                        `json:"limitations,omitempty"`
	SurfaceRefs               []string                                        `json:"surface_refs,omitempty"`
	EvidenceRefs              []string                                        `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer      string                                          `json:"projection_disclaimer"`
	IntegrationSummary        []string                                        `json:"integration_summary,omitempty"`
}

func verifierEcosystemValEAllSurfaceRefs() []string {
	return operability.VerifierEcosystemValEProofSurfaceRefs()
}

func verifierEcosystemValEProjectionDisclaimer() string {
	return "projection_only not_canonical_truth integrated_verifier_ecosystem_closure evidence_linked_verifier_closure"
}

func verifierEcosystemValEUniqueRefs(groups ...[]string) []string {
	seen := map[string]struct{}{}
	refs := []string{}
	for _, group := range groups {
		for _, ref := range group {
			trimmed := strings.TrimSpace(ref)
			if trimmed == "" {
				continue
			}
			if _, ok := seen[trimmed]; ok {
				continue
			}
			seen[trimmed] = struct{}{}
			refs = append(refs, trimmed)
		}
	}
	return refs
}

func verifierEcosystemValEUniqueTrimmedCount(values []string) (int, bool) {
	unique := make(map[string]struct{}, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			return 0, false
		}
		unique[trimmed] = struct{}{}
	}
	if len(unique) == 0 {
		return 0, false
	}
	return len(unique), true
}

func verifierEcosystemValEAudienceBreadthFacts(model operability.VerifierEcosystemValCAudienceSurfaceCatalog) (bool, bool) {
	publicCount := 0
	partnerCount := 0
	publicFound := false
	partnerFound := false
	for _, item := range model.Surfaces {
		switch strings.TrimSpace(item.AudienceType) {
		case operability.VerifierEcosystemValCAudiencePublic:
			count, ok := verifierEcosystemValEUniqueTrimmedCount(item.AllowedOutputClasses)
			if !ok {
				return false, false
			}
			publicCount = count
			publicFound = true
		case operability.VerifierEcosystemValCAudiencePartner:
			count, ok := verifierEcosystemValEUniqueTrimmedCount(item.AllowedOutputClasses)
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

func verifierEcosystemValEEvidenceFresh(states ...string) bool {
	supportedActiveStates := []string{
		operability.VerifierEcosystemVal0StateActive,
		operability.VerifierEcosystemValAStateActive,
		operability.VerifierEcosystemValBStateActive,
		operability.VerifierEcosystemValCStateActive,
		operability.VerifierEcosystemValDStateActive,
		operability.VerifierEcosystemValDStateActive,
	}
	for _, state := range states {
		matched := false
		for _, supported := range supportedActiveStates {
			if strings.TrimSpace(state) == strings.TrimSpace(supported) {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

func verifierEcosystemValEObservedClaims(
	publicOutput operability.VerifierEcosystemValCPublicOutputContract,
	partnerOutput operability.VerifierEcosystemValCPartnerOutputContract,
	requestContract operability.VerifierEcosystemValCRequestContract,
	publisherProfile operability.VerifierEcosystemValCPublisherCompatibilityProfile,
	artifactRules operability.VerifierEcosystemValCArtifactPublishingRuleCatalog,
	correctness operability.VerifierEcosystemValDCorrectnessGate,
	tooling operability.VerifierEcosystemValDToolingGate,
	diagnosticsConformance operability.VerifierEcosystemValDDiagnosticsConformanceGate,
	publisherArtifact operability.VerifierEcosystemValDPublisherArtifactGate,
	noOverclaim operability.VerifierEcosystemValDNoOverclaimGate,
) []string {
	claims := append([]string{}, publisherProfile.ObservedClaims...)
	claims = append(claims, noOverclaim.ObservedClaims...)
	for _, rule := range artifactRules.Rules {
		claims = append(claims, rule.ObservedClaims...)
	}
	if publicOutput.CertificationClaim {
		claims = append(claims, "verifier certification")
	}
	if publicOutput.UniversalTruthClaim {
		claims = append(claims, "mathematically proves total truth")
	}
	if publicOutput.RegulatorApprovalClaim {
		claims = append(claims, "regulator-approved verifier")
	}
	if publicOutput.DeploymentApprovalClaim || partnerOutput.ApprovesDeployment {
		claims = append(claims, "deployment approved")
	}
	if requestContract.ApprovesPublication {
		claims = append(claims, "publication authority")
	}
	if publisherProfile.ApprovedVendorClaim {
		claims = append(claims, "approved vendor")
	}
	if publisherProfile.AutomaticallyTrustedClaim || publisherArtifact.AutomaticallyTrustedClaim {
		claims = append(claims, "automatically trusted verifier-compatible artifact")
	}
	if correctness.CertificationClaim || tooling.CertificationClaim || diagnosticsConformance.CertificationClaim {
		claims = append(claims, "verifier certification")
	}
	if diagnosticsConformance.IntegrityRatingClaim {
		claims = append(claims, "integrity rating")
	}
	return verifierEcosystemValEUniqueRefs(claims)
}

func (s server) verifierEcosystemValEClosureHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValEClosure())
}

func (s server) verifierEcosystemValEProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValEProofs())
}

func buildVerifierEcosystemValEClosureModel() operability.VerifierEcosystemIntegratedClosure {
	val0Proofs := buildVerifierEcosystemVal0Proofs()
	valAProofs := buildVerifierEcosystemValAProofs()
	valBProofs := buildVerifierEcosystemValBProofs()
	valCProofs := buildVerifierEcosystemValCProofs()
	valDProofs := buildVerifierEcosystemValDProofs()

	_, _, _, _, _, trustModel, _, _, contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState, val0State := buildVerifierEcosystemVal0SharedStates()
	_, _, engine, _, _, _, _, inputState, engineState, resultState, diagnosticsMappingState, commandState, sdkState, valAState := buildVerifierEcosystemValASharedStates()
	_, _, schemaProofCompatibility, _, _, _, _, conformanceSuite, _, matrixState, schemaCompatibilityState, mixedVersionState, diagnosticPrecedenceState, fixtureDescriptorState, conformanceCaseState, conformanceSuiteState, outputClassState, valBState := buildVerifierEcosystemValBSharedStates()
	_, audiences, publicOutput, partnerOutput, _, requestContract, publisherProfile, artifactRules, trustDistribution, audienceState, publicOutputState, partnerOutputState, auditorFlowState, requestState, publisherState, artifactRuleState, trustDistributionState, _ := buildVerifierEcosystemValCSharedStates()
	_, correctness, tooling, _, diagnosticsConformance, trustKeyRotation, negativeDiagnostics, _, publisherArtifact, noOverclaim, correctnessState, toolingState, schemaCompatibilityGateState, diagnosticsConformanceState, trustKeyRotationState, negativeDiagnosticsState, redactionState, publisherArtifactState, noOverclaimState, _ := buildVerifierEcosystemValDSharedStates()

	audienceOrderIndependent, uniqueBreadthValid := verifierEcosystemValEAudienceBreadthFacts(audiences)
	observedClaims := verifierEcosystemValEObservedClaims(
		publicOutput,
		partnerOutput,
		requestContract,
		publisherProfile,
		artifactRules,
		correctness,
		tooling,
		diagnosticsConformance,
		publisherArtifact,
		noOverclaim,
	)

	model := operability.VerifierEcosystemIntegratedClosure{
		ClosureID:        "verifier-ecosystem-point-7-closure",
		Version:          "2026.04",
		Point:            "point_7",
		ClosureVal:       "val_e",
		Point7PassReason: operability.VerifierEcosystemValEPoint7PassReasonAllowed,
		SourceValStates: operability.VerifierEcosystemValESourceValStates{
			Val0State: val0Proofs.Val0State,
			ValAState: valAProofs.ValAState,
			ValBState: valBProofs.ValBState,
			ValCState: valCProofs.ValCState,
			ValDState: valDProofs.ValDState,
		},
		SourceCurrentStates: operability.VerifierEcosystemValESourceCurrentStates{
			Val0CurrentState: val0Proofs.CurrentState,
			ValACurrentState: valAProofs.CurrentState,
			ValBCurrentState: valBProofs.CurrentState,
			ValCCurrentState: valCProofs.CurrentState,
			ValDCurrentState: valDProofs.CurrentState,
		},
		DependencyStates: operability.VerifierEcosystemValEDependencyStates{
			Point5State:                    valDProofs.Point5State,
			Point5DependencyState:          valDProofs.Point5DependencyState,
			Point6State:                    valDProofs.Point6State,
			Point6ClosureState:             valDProofs.Point6ClosureState,
			Point6ClosurePrerequisiteState: valDProofs.Point6ClosurePrerequisite,
			Point6ClosureInvariantState:    valDProofs.Point6ClosureInvariant,
			Point6ProofSurfaceState:        valDProofs.Point6ProofSurfaceState,
			Point6PassRuleState:            valDProofs.Point6PassRuleState,
			Point6PassAllowed:              valDProofs.Point6PassAllowed,
			ValDFinalGateState:             valDProofs.ValDState,
			PreClosurePoint7State:          valDProofs.Point7State,
		},
		Val0: operability.VerifierEcosystemValEVal0ProofSnapshot{
			CurrentState:             val0Proofs.CurrentState,
			Val0State:                val0Proofs.Val0State,
			Point7State:              val0Proofs.Point7State,
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
			SurfaceRefs:              val0Proofs.SurfaceRefs,
			EvidenceRefs:             val0Proofs.EvidenceRefs,
			ProjectionDisclaimer:     val0Proofs.ProjectionDisclaimer,
			WorstSeverityPrecedence:  val0State == operability.VerifierEcosystemVal0StateActive,
		},
		ValA: operability.VerifierEcosystemValEValAProofSnapshot{
			CurrentState:                 valAProofs.CurrentState,
			ValAState:                    valAProofs.ValAState,
			Point7State:                  valAProofs.Point7State,
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
			SurfaceRefs:                  valAProofs.SurfaceRefs,
			EvidenceRefs:                 valAProofs.EvidenceRefs,
			ProjectionDisclaimer:         valAProofs.ProjectionDisclaimer,
			WorstSeverityPrecedence:      valAState == operability.VerifierEcosystemValAStateActive,
		},
		ValB: operability.VerifierEcosystemValEValBProofSnapshot{
			CurrentState:                  valBProofs.CurrentState,
			ValBState:                     valBProofs.ValBState,
			Point7State:                   valBProofs.Point7State,
			CompatibilityMatrixState:      matrixState,
			SchemaProofCompatibilityState: schemaCompatibilityState,
			MixedVersionDiagnosticState:   mixedVersionState,
			DiagnosticPrecedenceState:     diagnosticPrecedenceState,
			FixtureDescriptorState:        fixtureDescriptorState,
			ConformanceCaseState:          conformanceCaseState,
			ConformanceSuiteState:         conformanceSuiteState,
			OutputClassState:              outputClassState,
			CompatibilityState:            schemaProofCompatibility.CompatibilityState,
			DerivedDiagnosticClass:        valBProofs.DerivedDiagnosticClass,
			DerivedOutputClass:            valBProofs.DerivedOutputClass,
			NegativeCasesPreserved:        fixtureDescriptorState == operability.VerifierEcosystemValBFixtureDescriptorStateActive && conformanceCaseState == operability.VerifierEcosystemValBConformanceCaseStateActive,
			ConformanceCertificationClaim: conformanceSuite.CertificationClaim,
			IntegrityRatingClaim:          false,
			SurfaceRefs:                   valBProofs.SurfaceRefs,
			EvidenceRefs:                  valBProofs.EvidenceRefs,
			ProjectionDisclaimer:          valBProofs.ProjectionDisclaimer,
			WorstSeverityPrecedence:       valBState == operability.VerifierEcosystemValBStateActive,
		},
		ValC: operability.VerifierEcosystemValEValCProofSnapshot{
			CurrentState:                          valCProofs.CurrentState,
			ValCState:                             valCProofs.ValCState,
			Point7State:                           valCProofs.Point7State,
			AudienceSurfaceState:                  audienceState,
			PublicOutputState:                     publicOutputState,
			PartnerOutputState:                    partnerOutputState,
			AuditorFlowState:                      auditorFlowState,
			RequestContractState:                  requestState,
			PublisherProfileState:                 publisherState,
			ArtifactRuleState:                     artifactRuleState,
			TrustDistributionState:                trustDistributionState,
			PublicOutputClass:                     valCProofs.PublicOutputClass,
			PartnerOutputClass:                    valCProofs.PartnerOutputClass,
			RequestMode:                           valCProofs.RequestMode,
			PublisherType:                         valCProofs.PublisherType,
			TrustDistributionMode:                 valCProofs.TrustDistributionMode,
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
			SurfaceRefs:                           valCProofs.SurfaceRefs,
			EvidenceRefs:                          valCProofs.EvidenceRefs,
			ProjectionDisclaimer:                  valCProofs.ProjectionDisclaimer,
		},
		ValD: operability.VerifierEcosystemValEValDProofSnapshot{
			CurrentState:                            valDProofs.CurrentState,
			ValDState:                               valDProofs.ValDState,
			Point7State:                             valDProofs.Point7State,
			CorrectnessGateState:                    correctnessState,
			ToolingGateState:                        toolingState,
			SchemaCompatibilityGateState:            schemaCompatibilityGateState,
			DiagnosticsConformanceGateState:         diagnosticsConformanceState,
			TrustKeyRotationGateState:               trustKeyRotationState,
			NegativeDiagnosticsGateState:            negativeDiagnosticsState,
			RedactionGateState:                      redactionState,
			PublisherArtifactGateState:              publisherArtifactState,
			NoOverclaimGateState:                    noOverclaimState,
			TrustDistributionMode:                   valDProofs.TrustDistributionMode,
			OfflineDistributionScope:                valDProofs.OfflineDistributionScope,
			TrustDistributionModeUsesActualValCMode: trustKeyRotation.TrustDistributionMode == trustDistribution.TrustRootDistributionMode && trustKeyRotation.TrustDistributionMode != trustKeyRotation.OfflineDistributionScope,
			ClaimsIntegratedClosure:                 false,
			SurfaceRefs:                             valDProofs.SurfaceRefs,
			EvidenceRefs:                            valDProofs.EvidenceRefs,
			ProjectionDisclaimer:                    valDProofs.ProjectionDisclaimer,
		},
		ProofSurfaceRefs: verifierEcosystemValEAllSurfaceRefs(),
		EvidenceRefs:     operability.VerifierEcosystemValEProofEvidenceRefs(),
		ObservedClaims:   observedClaims,
		Caveats: []string{
			"Integrated verifier ecosystem closure remains an advisory projection over the canonical execution, audit, and evidence spine.",
		},
		Limitations: []string{
			"Val E closes Točka 7 only and does not start Točka 8 or create verifier, publisher, vendor, or deployment authority.",
			"Integrated closure remains fail-closed and blocks on stale, revoked, expired, unsupported, malformed, unknown, partial, blocked, degraded, or superseded source states.",
		},
		ProjectionDisclaimer: verifierEcosystemValEProjectionDisclaimer(),
		CreatedAt:            publicSampleTime().Format(time.RFC3339),
		UpdatedAt:            publicSampleTime().Format(time.RFC3339),
		EvidenceFresh: verifierEcosystemValEEvidenceFresh(
			val0Proofs.CurrentState,
			valAProofs.CurrentState,
			valBProofs.CurrentState,
			valCProofs.CurrentState,
			valDProofs.CurrentState,
			valDProofs.ValDState,
		),
		StaleEvidenceDetected:         false,
		RedactionKeepsFailuresVisible: negativeDiagnostics.RedactionBoundaryPreserved && negativeDiagnostics.PublicPreservesNonVerified && negativeDiagnostics.AuditorRepeatable && negativeDiagnostics.AuditorEvidenceLinked,
		MutatesCanonicalEvidence:      tooling.MutatesEvidence || requestContract.IngestsCanonicalEvidence,
		ApprovesDeployment:            tooling.ApprovesDeployment || publicOutput.DeploymentApprovalClaim || partnerOutput.ApprovesDeployment,
		SuppressesFailures:            tooling.SuppressesFailures || partnerOutput.SuppressesFailures,
	}
	model = operability.ComputeVerifierEcosystemValEClosure(model)
	if model.Point7PassAllowed {
		model.Point7PassReason = operability.VerifierEcosystemValEPoint7PassReasonAllowed
	} else {
		model.Point7PassReason = operability.VerifierEcosystemValEPoint7PassReasonBlocked
	}
	return operability.ComputeVerifierEcosystemValEClosure(model)
}

func buildVerifierEcosystemValEClosure() verifierEcosystemValEClosureResponse {
	model := buildVerifierEcosystemValEClosureModel()
	return verifierEcosystemValEClosureResponse{
		SchemaVersion: verifierEcosystemValEClosureSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     verifierEcosystemValEAllSurfaceRefs(),
		Limitations:   model.Limitations,
	}
}

func buildVerifierEcosystemValEProofs() verifierEcosystemValEProofsResponse {
	model := buildVerifierEcosystemValEClosureModel()
	return verifierEcosystemValEProofsResponse{
		SchemaVersion:             verifierEcosystemValEProofsSchema,
		GeneratedAt:               publicSampleTime(),
		CurrentState:              model.CurrentState,
		Point5State:               model.DependencyStates.Point5State,
		Point5DependencyState:     model.DependencyStates.Point5DependencyState,
		Point6State:               model.DependencyStates.Point6State,
		Point6ClosureState:        model.DependencyStates.Point6ClosureState,
		Point6ClosurePrerequisite: model.DependencyStates.Point6ClosurePrerequisiteState,
		Point6ClosureInvariant:    model.DependencyStates.Point6ClosureInvariantState,
		Point6ProofSurfaceState:   model.DependencyStates.Point6ProofSurfaceState,
		Point6PassRuleState:       model.DependencyStates.Point6PassRuleState,
		Point6PassAllowed:         model.DependencyStates.Point6PassAllowed,
		Val0CurrentState:          model.SourceCurrentStates.Val0CurrentState,
		Val0State:                 model.SourceValStates.Val0State,
		ValACurrentState:          model.SourceCurrentStates.ValACurrentState,
		ValAState:                 model.SourceValStates.ValAState,
		ValBCurrentState:          model.SourceCurrentStates.ValBCurrentState,
		ValBState:                 model.SourceValStates.ValBState,
		ValCCurrentState:          model.SourceCurrentStates.ValCCurrentState,
		ValCState:                 model.SourceValStates.ValCState,
		ValDCurrentState:          model.SourceCurrentStates.ValDCurrentState,
		ValDState:                 model.SourceValStates.ValDState,
		ValDFinalGateState:        model.DependencyStates.ValDFinalGateState,
		ClosurePrerequisiteState:  model.ClosurePrerequisiteState,
		ClosureInvariantState:     model.ClosureInvariantState,
		ProofSurfaceState:         model.ProofSurfaceState,
		EvidenceQualityState:      model.EvidenceQualityState,
		NoOverclaimState:          model.NoOverclaimState,
		PassRuleState:             model.PassRuleState,
		ValEState:                 model.CurrentState,
		Point7State:               model.Point7State,
		Point7PassAllowed:         model.Point7PassAllowed,
		Point7PassReason:          model.Point7PassReason,
		ClosureInvariants:         model.ClosureInvariants,
		BlockingReasons:           model.BlockingReasons,
		Caveats:                   model.Caveats,
		Limitations:               model.Limitations,
		SurfaceRefs:               model.ProofSurfaceRefs,
		EvidenceRefs:              model.EvidenceRefs,
		ProjectionDisclaimer:      model.ProjectionDisclaimer,
		IntegrationSummary: []string{
			"Val E integrates actual Val 0 through Val D verifier ecosystem states and is the only layer that can return point_7_pass.",
			"Integrated closure remains advisory and projection-only, does not create certification or authority, and does not start Točka 8.",
		},
	}
}
