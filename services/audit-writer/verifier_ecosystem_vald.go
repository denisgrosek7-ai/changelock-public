package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	verifierEcosystemValDCorrectnessGateSchema            = "point7.verifier_ecosystem.vald.correctness_gate.v1"
	verifierEcosystemValDToolingGateSchema                = "point7.verifier_ecosystem.vald.tooling_gate.v1"
	verifierEcosystemValDSchemaCompatibilityGateSchema    = "point7.verifier_ecosystem.vald.schema_compatibility_gate.v1"
	verifierEcosystemValDDiagnosticsConformanceGateSchema = "point7.verifier_ecosystem.vald.diagnostics_conformance_gate.v1"
	verifierEcosystemValDTrustKeyRotationGateSchema       = "point7.verifier_ecosystem.vald.trust_key_rotation_gate.v1"
	verifierEcosystemValDNegativeDiagnosticsGateSchema    = "point7.verifier_ecosystem.vald.negative_diagnostics_gate.v1"
	verifierEcosystemValDRedactionGateSchema              = "point7.verifier_ecosystem.vald.redaction_gate.v1"
	verifierEcosystemValDPublisherArtifactGateSchema      = "point7.verifier_ecosystem.vald.publisher_artifact_gate.v1"
	verifierEcosystemValDNoOverclaimGateSchema            = "point7.verifier_ecosystem.vald.no_overclaim_gate.v1"
	verifierEcosystemValDProofsSchema                     = "point7.verifier_ecosystem.vald.proofs.v1"
)

type verifierEcosystemValDModelResponse struct {
	SchemaVersion string    `json:"schema_version"`
	GeneratedAt   time.Time `json:"generated_at"`
	CurrentState  string    `json:"current_state"`
	Model         any       `json:"model"`
	RouteRefs     []string  `json:"route_refs,omitempty"`
	Limitations   []string  `json:"limitations,omitempty"`
}

type verifierEcosystemValDProofsResponse struct {
	SchemaVersion                   string    `json:"schema_version"`
	GeneratedAt                     time.Time `json:"generated_at"`
	CurrentState                    string    `json:"current_state"`
	Point5State                     string    `json:"point_5_state"`
	Point5DependencyState           string    `json:"point_5_dependency_state"`
	Point6State                     string    `json:"point_6_state"`
	Point6ClosureState              string    `json:"point_6_closure_state"`
	Point6ClosurePrerequisite       string    `json:"point_6_closure_prerequisite_state"`
	Point6ClosureInvariant          string    `json:"point_6_closure_invariant_state"`
	Point6ProofSurfaceState         string    `json:"point_6_proof_surface_state"`
	Point6PassRuleState             string    `json:"point_6_pass_rule_state"`
	Point6PassAllowed               bool      `json:"point_6_pass_allowed"`
	Val0CurrentState                string    `json:"val_0_current_state"`
	Val0State                       string    `json:"val_0_state"`
	ValACurrentState                string    `json:"val_a_current_state"`
	ValAState                       string    `json:"val_a_state"`
	ValBCurrentState                string    `json:"val_b_current_state"`
	ValBState                       string    `json:"val_b_state"`
	ValCCurrentState                string    `json:"val_c_current_state"`
	ValCState                       string    `json:"val_c_state"`
	ValDState                       string    `json:"val_d_state"`
	Point7State                     string    `json:"point_7_state"`
	CorrectnessGateState            string    `json:"correctness_gate_state"`
	ToolingGateState                string    `json:"tooling_gate_state"`
	SchemaCompatibilityGateState    string    `json:"schema_compatibility_gate_state"`
	DiagnosticsConformanceGateState string    `json:"diagnostics_conformance_gate_state"`
	TrustKeyRotationGateState       string    `json:"trust_key_rotation_gate_state"`
	NegativeDiagnosticsGateState    string    `json:"negative_diagnostics_gate_state"`
	RedactionGateState              string    `json:"redaction_gate_state"`
	PublisherArtifactGateState      string    `json:"publisher_artifact_gate_state"`
	NoOverclaimGateState            string    `json:"no_overclaim_gate_state"`
	DerivedPublicOutputClass        string    `json:"derived_public_output_class"`
	DerivedPartnerOutputClass       string    `json:"derived_partner_output_class"`
	TrustDistributionMode           string    `json:"trust_distribution_mode"`
	OfflineDistributionScope        string    `json:"offline_distribution_scope"`
	PublisherType                   string    `json:"publisher_type"`
	WhyPoint7NotPass                []string  `json:"why_point_7_not_pass,omitempty"`
	SurfaceRefs                     []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                    []string  `json:"evidence_refs,omitempty"`
	Limitations                     []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer            string    `json:"projection_disclaimer"`
	IntegrationSummary              []string  `json:"integration_summary,omitempty"`
}

func verifierEcosystemValDAllSurfaceRefs() []string {
	return operability.VerifierEcosystemValDProofSurfaceRefs()
}

func verifierEcosystemValDProjectionDisclaimer() string {
	return "projection_only not_canonical_truth final_verifier_ecosystem_gate advisory_projection"
}

func verifierEcosystemValDEvidenceRefs() []string {
	return operability.VerifierEcosystemValDProofEvidenceRefs()
}

func buildVerifierEcosystemValDDependencySnapshot() operability.VerifierEcosystemValDDependencySnapshot {
	valE := buildReferenceArchitectureValEProofs()
	val0 := buildVerifierEcosystemVal0Proofs()
	valA := buildVerifierEcosystemValAProofs()
	valB := buildVerifierEcosystemValBProofs()
	valC := buildVerifierEcosystemValCProofs()
	return operability.VerifierEcosystemValDDependencySnapshot{
		Point5State:                    valE.Point5State,
		Point5DependencyState:          valE.Point5DependencyState,
		Point6State:                    valE.Point6State,
		Point6ClosureState:             valE.ValEState,
		Point6ClosurePrerequisiteState: valE.ClosurePrerequisiteState,
		Point6ClosureInvariantState:    valE.ClosureInvariantState,
		Point6ProofSurfaceState:        valE.ProofSurfaceState,
		Point6PassRuleState:            valE.PassRuleState,
		Point6PassAllowed:              valE.Point6PassAllowed,
		Val0CurrentState:               val0.CurrentState,
		Val0State:                      val0.Val0State,
		ValACurrentState:               valA.CurrentState,
		ValAState:                      valA.ValAState,
		ValBCurrentState:               valB.CurrentState,
		ValBState:                      valB.ValBState,
		ValCCurrentState:               valC.CurrentState,
		ValCState:                      valC.ValCState,
		Point7State:                    valC.Point7State,
	}
}

func verifierEcosystemValDAudienceSurfaceSummary(model operability.VerifierEcosystemValCAudienceSurfaceCatalog) (int, int, string, string, string, bool) {
	publicCount := 0
	partnerCount := 0
	redactionPolicy := ""
	evidenceVisibility := ""
	trustVisibility := ""
	internalSeparated := false
	for _, item := range model.Surfaces {
		switch item.AudienceType {
		case operability.VerifierEcosystemValCAudiencePublic:
			publicCount = len(item.AllowedOutputClasses)
			redactionPolicy = item.RedactionPolicyRef
			evidenceVisibility = item.EvidenceVisibilityPolicy
			trustVisibility = item.TrustMaterialVisibility
		case operability.VerifierEcosystemValCAudiencePartner:
			partnerCount = len(item.AllowedOutputClasses)
		case operability.VerifierEcosystemValCAudienceInternal:
			internalSeparated = !item.PublicReuseAllowed
		}
	}
	return publicCount, partnerCount, redactionPolicy, evidenceVisibility, trustVisibility, internalSeparated
}

func buildVerifierEcosystemValDSharedStates() (
	operability.VerifierEcosystemValDDependencySnapshot,
	operability.VerifierEcosystemValDCorrectnessGate,
	operability.VerifierEcosystemValDToolingGate,
	operability.VerifierEcosystemValDSchemaCompatibilityGate,
	operability.VerifierEcosystemValDDiagnosticsConformanceGate,
	operability.VerifierEcosystemValDTrustKeyRotationGate,
	operability.VerifierEcosystemValDNegativeDiagnosticsGate,
	operability.VerifierEcosystemValDRedactionGate,
	operability.VerifierEcosystemValDPublisherArtifactGate,
	operability.VerifierEcosystemValDNoOverclaimGate,
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
	dependency := buildVerifierEcosystemValDDependencySnapshot()
	_, _, _, _, _, trustModel, _, _, contractState, _, _, _, trustState, _, outputBoundaryState, _ := buildVerifierEcosystemVal0SharedStates()
	_, _, engine, result, _, command, sdk, inputState, engineState, resultState, diagnosticsMappingState, commandState, sdkState, _ := buildVerifierEcosystemValASharedStates()
	_, matrix, schemaProofCompatibility, mixedVersion, diagnosticPrecedence, fixtures, conformanceCases, conformanceSuite, _, matrixState, schemaCompatibilityState, mixedVersionState, diagnosticPrecedenceState, fixtureDescriptorState, conformanceCaseState, conformanceSuiteState, outputClassState, _ := buildVerifierEcosystemValBSharedStates()
	_, audiences, publicOutput, partnerOutput, auditorFlow, _, publisherProfile, artifactRules, trustDistribution, audienceState, publicOutputState, partnerOutputState, auditorFlowState, _, publisherProfileState, artifactRuleState, trustDistributionState, valCState := buildVerifierEcosystemValCSharedStates()

	correctness := operability.VerifierEcosystemValDCorrectnessGateModel()
	correctness.SourceValStates = []string{dependency.Val0State, dependency.ValAState, dependency.ValBState, dependency.ValCState}
	correctness.VerifierContractState = contractState
	correctness.ReferenceEngineState = engineState
	correctness.CompatibilityMatrixState = matrixState
	correctness.DiagnosticsState = diagnosticPrecedenceState
	correctness.ConformanceSuiteState = conformanceSuiteState
	correctness.EcosystemSurfaceState = valCState

	tooling := operability.VerifierEcosystemValDToolingGateModel()
	tooling.VerifierInputModelState = inputState
	tooling.ReferenceVerifierEngineState = engineState
	tooling.VerificationResultModelState = resultState
	tooling.DiagnosticsMappingState = diagnosticsMappingState
	tooling.CommandSurfaceState = commandState
	tooling.SDKEntrypointState = sdkState
	tooling.DeterministicOutput = engine.DeterministicOutput
	tooling.NetworkDependency = engine.NetworkDependency
	tooling.MutatesEvidence = engine.MutatesEvidence
	tooling.ApprovesDeployment = engine.ClaimsDeploymentApproval
	tooling.ClaimsRealCryptoWithoutPrimitive = engine.ClaimsActualCryptoValidity && !engine.UsesRealCryptoPrimitives
	_ = command
	_ = result

	schemaCompatibility := operability.VerifierEcosystemValDSchemaCompatibilityGateModel()
	schemaCompatibility.CompatibilityMatrixState = matrixState
	schemaCompatibility.SchemaProofCompatibilityState = schemaCompatibilityState
	schemaCompatibility.MixedVersionDiagnosticsState = mixedVersionState
	schemaCompatibility.SupportedSchemaVersions = matrix.SupportedSchemaVersions
	schemaCompatibility.SupportedProofTypes = matrix.SupportedProofTypes
	schemaCompatibility.SupportedVerifierVersions = matrix.SupportedVerifierVersions
	schemaCompatibility.SupportedTrustRootVersions = matrix.SupportedTrustRootVersions
	schemaCompatibility.MixedVersionRuleCoverage = matrix.MixedVersionRules
	schemaCompatibility.CompatibilityState = schemaProofCompatibility.CompatibilityState
	if len(schemaProofCompatibility.Caveats) > 0 {
		schemaCompatibility.Caveats = schemaProofCompatibility.Caveats
	}
	schemaCompatibility.DeprecatedMigrationVisible = schemaProofCompatibility.MigrationOrSupersessionRef != ""
	schemaCompatibility.SupersessionVisible = schemaProofCompatibility.MigrationOrSupersessionRef != ""

	diagnosticsConformance := operability.VerifierEcosystemValDDiagnosticsConformanceGateModel()
	diagnosticsConformance.DiagnosticPrecedenceState = diagnosticPrecedenceState
	diagnosticsConformance.FixtureDescriptorState = fixtureDescriptorState
	diagnosticsConformance.ConformanceCaseState = conformanceCaseState
	diagnosticsConformance.ConformanceSuiteState = conformanceSuiteState
	diagnosticsConformance.OutputClassState = outputClassState
	diagnosticsConformance.ObservedDiagnostics = diagnosticPrecedence.ObservedDiagnostics
	diagnosticsConformance.DerivedDiagnosticClass = diagnosticPrecedence.DerivedDiagnosticClass
	diagnosticsConformance.StaleCoverageVisible = true
	diagnosticsConformance.RevokedCoverageVisible = true
	diagnosticsConformance.SupersededCoverageVisible = true
	diagnosticsConformance.UnsupportedCoverageVisible = true
	diagnosticsConformance.MalformedCoverageVisible = true
	_ = mixedVersion
	_ = fixtures
	_ = conformanceCases
	_ = conformanceSuite
	_ = sdk

	trustKeyRotation := operability.VerifierEcosystemValDTrustKeyRotationGateModel()
	trustKeyRotation.TrustState = trustState
	trustKeyRotation.TrustRootState = trustModel.TrustRootState
	if trustModel.RevocationState == operability.VerifierEcosystemRevocationRevoked {
		trustKeyRotation.IssuerState = "issuer_revoked"
	} else if trustModel.TrustRootState == operability.VerifierEcosystemTrustRootTrustedWithWarnings {
		trustKeyRotation.IssuerState = "issuer_warning"
	} else {
		trustKeyRotation.IssuerState = "issuer_bound"
	}
	trustKeyRotation.RevocationState = trustModel.RevocationState
	trustKeyRotation.KeyRotationState = trustModel.KeyRotationState
	trustKeyRotation.RolloverMetadataRef = trustModel.RolloverMetadataRef
	trustKeyRotation.TrustDistributionState = trustDistributionState
	trustKeyRotation.TrustDistributionMode = trustDistribution.TrustRootDistributionMode
	trustKeyRotation.OfflineDistributionScope = trustDistribution.AudienceVisibilityScope
	trustKeyRotation.GlobalKeyDirectoryClaim = trustDistribution.GlobalKeyDirectoryClaim
	trustKeyRotation.SensitivePublicKeyExposure = trustDistribution.SensitiveKeyMaterialExposed && trustDistribution.AudienceVisibilityScope == operability.VerifierEcosystemScopePublicSafe

	negativeDiagnostics := operability.VerifierEcosystemValDNegativeDiagnosticsGateModel()
	negativeDiagnostics.PublicOutputState = publicOutputState
	negativeDiagnostics.PartnerOutputState = partnerOutputState
	negativeDiagnostics.AuditorFlowState = auditorFlowState
	negativeDiagnostics.PublicOverallResult = publicOutput.OverallResult
	negativeDiagnostics.PublicDiagnosticClass = publicOutput.DiagnosticClass
	negativeDiagnostics.PublicOutputClass = publicOutput.OutputClass
	negativeDiagnostics.PartnerOverallResult = partnerOutput.OverallResult
	negativeDiagnostics.PartnerDiagnosticClass = partnerOutput.DiagnosticClass
	negativeDiagnostics.PartnerOutputClass = partnerOutput.OutputClass
	negativeDiagnostics.PartnerInternalDiagnosticsExposed = partnerOutput.InternalOnlyDiagnosticsExposed
	negativeDiagnostics.AuditorRepeatable = auditorFlow.Repeatable
	negativeDiagnostics.AuditorEvidenceLinked = auditorFlow.EvidenceLinked

	publicCount, partnerCount, redactionPolicy, evidenceVisibility, trustVisibility, internalSeparated := verifierEcosystemValDAudienceSurfaceSummary(audiences)

	redaction := operability.VerifierEcosystemValDRedactionGateModel()
	redaction.AudienceSurfaceState = audienceState
	redaction.PublicOutputState = publicOutputState
	redaction.PartnerOutputState = partnerOutputState
	redaction.AuditorFlowState = auditorFlowState
	redaction.OutputBoundaryState = outputBoundaryState
	redaction.RedactionPolicyRef = redactionPolicy
	redaction.EvidenceVisibilityPolicy = evidenceVisibility
	redaction.TrustMaterialVisibility = trustVisibility
	redaction.InternalDiagnosticSeparated = internalSeparated
	redaction.PartnerBroaderThanPublic = partnerCount > publicCount

	publisherArtifact := operability.VerifierEcosystemValDPublisherArtifactGateModel()
	publisherArtifact.PublisherProfileState = publisherProfileState
	publisherArtifact.ArtifactRuleState = artifactRuleState
	publisherArtifact.PublisherType = publisherProfile.PublisherType
	publisherArtifact.RequiredSchemaPolicy = "required_schema_versions_present"
	publisherArtifact.RequiredSignaturePolicy = publisherProfile.RequiredSignaturePolicy
	publisherArtifact.RequiredDigestPolicy = publisherProfile.RequiredDigestPolicy
	publisherArtifact.TrustRootPolicy = publisherProfile.TrustRootPolicy
	publisherArtifact.SupportedArtifactTypes = publisherProfile.SupportedArtifactTypes
	publisherArtifact.ConformanceCaseRefs = publisherProfile.ConformanceCaseRefs
	publisherArtifact.ObservedClaims = publisherProfile.ObservedClaims
	publisherArtifact.OutputBoundaryCompatible = len(artifactRules.Rules) > 0
	publisherArtifact.AutomaticallyTrustedClaim = publisherProfile.AutomaticallyTrustedClaim

	noOverclaim := operability.VerifierEcosystemValDNoOverclaimGateModel()
	noOverclaim.ObservedClaims = append(noOverclaim.ObservedClaims, publisherProfile.ObservedClaims...)

	correctnessState := operability.EvaluateVerifierEcosystemValDCorrectnessGateState(correctness)
	toolingState := operability.EvaluateVerifierEcosystemValDToolingGateState(tooling)
	schemaCompatibilityStateGate := operability.EvaluateVerifierEcosystemValDSchemaCompatibilityGateState(schemaCompatibility)
	diagnosticsConformanceState := operability.EvaluateVerifierEcosystemValDDiagnosticsConformanceGateState(diagnosticsConformance)
	trustKeyRotationState := operability.EvaluateVerifierEcosystemValDTrustKeyRotationGateState(trustKeyRotation)
	negativeDiagnosticsState := operability.EvaluateVerifierEcosystemValDNegativeDiagnosticsGateState(negativeDiagnostics)
	redactionState := operability.EvaluateVerifierEcosystemValDRedactionGateState(redaction)
	publisherArtifactState := operability.EvaluateVerifierEcosystemValDPublisherArtifactGateState(publisherArtifact)
	noOverclaimState := operability.EvaluateVerifierEcosystemValDNoOverclaimGateState(noOverclaim)
	valDState := operability.EvaluateVerifierEcosystemValDState(
		dependency,
		correctnessState,
		toolingState,
		schemaCompatibilityStateGate,
		diagnosticsConformanceState,
		trustKeyRotationState,
		negativeDiagnosticsState,
		redactionState,
		publisherArtifactState,
		noOverclaimState,
	)
	return dependency, correctness, tooling, schemaCompatibility, diagnosticsConformance, trustKeyRotation, negativeDiagnostics, redaction, publisherArtifact, noOverclaim, correctnessState, toolingState, schemaCompatibilityStateGate, diagnosticsConformanceState, trustKeyRotationState, negativeDiagnosticsState, redactionState, publisherArtifactState, noOverclaimState, valDState
}

func (s server) verifierEcosystemValDCorrectnessGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValDCorrectnessGate())
}

func (s server) verifierEcosystemValDToolingGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValDToolingGate())
}

func (s server) verifierEcosystemValDSchemaCompatibilityGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValDSchemaCompatibilityGate())
}

func (s server) verifierEcosystemValDDiagnosticsConformanceGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValDDiagnosticsConformanceGate())
}

func (s server) verifierEcosystemValDTrustKeyRotationGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValDTrustKeyRotationGate())
}

func (s server) verifierEcosystemValDNegativeDiagnosticsGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValDNegativeDiagnosticsGate())
}

func (s server) verifierEcosystemValDRedactionGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValDRedactionGate())
}

func (s server) verifierEcosystemValDPublisherArtifactGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValDPublisherArtifactGate())
}

func (s server) verifierEcosystemValDNoOverclaimGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValDNoOverclaimGate())
}

func (s server) verifierEcosystemValDProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValDProofs())
}

func buildVerifierEcosystemValDCorrectnessGate() verifierEcosystemValDModelResponse {
	_, correctness, _, _, _, _, _, _, _, _, correctnessState, _, _, _, _, _, _, _, _, _ := buildVerifierEcosystemValDSharedStates()
	return verifierEcosystemValDModelResponse{
		SchemaVersion: verifierEcosystemValDCorrectnessGateSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  correctnessState,
		Model:         correctness,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vald/tooling-gate",
			"/v1/verifier-ecosystem/vald/proofs",
		},
		Limitations: []string{
			"Correctness gate remains a bounded final verifier ecosystem gate and does not certify or approve a verifier.",
		},
	}
}

func buildVerifierEcosystemValDToolingGate() verifierEcosystemValDModelResponse {
	_, _, tooling, _, _, _, _, _, _, _, _, toolingState, _, _, _, _, _, _, _, _ := buildVerifierEcosystemValDSharedStates()
	return verifierEcosystemValDModelResponse{
		SchemaVersion: verifierEcosystemValDToolingGateSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  toolingState,
		Model:         tooling,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vald/schema-compatibility-gate",
			"/v1/verifier-ecosystem/vald/proofs",
		},
		Limitations: []string{
			"Reference verifier tooling remains advisory, deterministic, and non-authoritative.",
		},
	}
}

func buildVerifierEcosystemValDSchemaCompatibilityGate() verifierEcosystemValDModelResponse {
	_, _, _, schemaCompatibility, _, _, _, _, _, _, _, _, schemaCompatibilityState, _, _, _, _, _, _, _ := buildVerifierEcosystemValDSharedStates()
	return verifierEcosystemValDModelResponse{
		SchemaVersion: verifierEcosystemValDSchemaCompatibilityGateSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  schemaCompatibilityState,
		Model:         schemaCompatibility,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vald/diagnostics-conformance-gate",
			"/v1/verifier-ecosystem/vald/proofs",
		},
		Limitations: []string{
			"Schema compatibility gate remains version-bound and warning-bearing where migration or supersession visibility is required.",
		},
	}
}

func buildVerifierEcosystemValDDiagnosticsConformanceGate() verifierEcosystemValDModelResponse {
	_, _, _, _, diagnosticsConformance, _, _, _, _, _, _, _, _, diagnosticsConformanceState, _, _, _, _, _, _ := buildVerifierEcosystemValDSharedStates()
	return verifierEcosystemValDModelResponse{
		SchemaVersion: verifierEcosystemValDDiagnosticsConformanceGateSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  diagnosticsConformanceState,
		Model:         diagnosticsConformance,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vald/trust-key-rotation-gate",
			"/v1/verifier-ecosystem/vald/proofs",
		},
		Limitations: []string{
			"Diagnostics and conformance gate remains a bounded verifier behavior gate and does not certify a verifier.",
		},
	}
}

func buildVerifierEcosystemValDTrustKeyRotationGate() verifierEcosystemValDModelResponse {
	_, _, _, _, _, trustKeyRotation, _, _, _, _, _, _, _, _, trustKeyRotationState, _, _, _, _, _ := buildVerifierEcosystemValDSharedStates()
	return verifierEcosystemValDModelResponse{
		SchemaVersion: verifierEcosystemValDTrustKeyRotationGateSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  trustKeyRotationState,
		Model:         trustKeyRotation,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vald/negative-diagnostics-gate",
			"/v1/verifier-ecosystem/vald/proofs",
		},
		Limitations: []string{
			"Trust-root and key-rotation gate remains scoped, bounded, and not a global key directory or universal trust protocol.",
		},
	}
}

func buildVerifierEcosystemValDNegativeDiagnosticsGate() verifierEcosystemValDModelResponse {
	_, _, _, _, _, _, negativeDiagnostics, _, _, _, _, _, _, _, _, negativeDiagnosticsState, _, _, _, _ := buildVerifierEcosystemValDSharedStates()
	return verifierEcosystemValDModelResponse{
		SchemaVersion: verifierEcosystemValDNegativeDiagnosticsGateSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  negativeDiagnosticsState,
		Model:         negativeDiagnostics,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vald/redaction-gate",
			"/v1/verifier-ecosystem/vald/proofs",
		},
		Limitations: []string{
			"Negative diagnostics remain visible as non-verified outcomes and cannot be redacted into verified status.",
		},
	}
}

func buildVerifierEcosystemValDRedactionGate() verifierEcosystemValDModelResponse {
	_, _, _, _, _, _, _, redaction, _, _, _, _, _, _, _, _, redactionState, _, _, _ := buildVerifierEcosystemValDSharedStates()
	return verifierEcosystemValDModelResponse{
		SchemaVersion: verifierEcosystemValDRedactionGateSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  redactionState,
		Model:         redaction,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vald/publisher-artifact-gate",
			"/v1/verifier-ecosystem/vald/proofs",
		},
		Limitations: []string{
			"Redaction gate preserves audience separation and output-boundary discipline without creating canonical truth.",
		},
	}
}

func buildVerifierEcosystemValDPublisherArtifactGate() verifierEcosystemValDModelResponse {
	_, _, _, _, _, _, _, _, publisherArtifact, _, _, _, _, _, _, _, _, publisherArtifactState, _, _ := buildVerifierEcosystemValDSharedStates()
	return verifierEcosystemValDModelResponse{
		SchemaVersion: verifierEcosystemValDPublisherArtifactGateSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  publisherArtifactState,
		Model:         publisherArtifact,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vald/no-overclaim-gate",
			"/v1/verifier-ecosystem/vald/proofs",
		},
		Limitations: []string{
			"Publisher and artifact compatibility remain bounded guidance only and do not certify publishers or artifacts.",
		},
	}
}

func buildVerifierEcosystemValDNoOverclaimGate() verifierEcosystemValDModelResponse {
	_, _, _, _, _, _, _, _, _, noOverclaim, _, _, _, _, _, _, _, _, noOverclaimState, _ := buildVerifierEcosystemValDSharedStates()
	return verifierEcosystemValDModelResponse{
		SchemaVersion: verifierEcosystemValDNoOverclaimGateSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  noOverclaimState,
		Model:         noOverclaim,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vald/proofs",
		},
		Limitations: []string{
			"No-overclaim gate blocks certification, approval, universal authority, and point_7_pass claims across Val D outputs.",
		},
	}
}

func buildVerifierEcosystemValDProofs() verifierEcosystemValDProofsResponse {
	dependency, _, _, _, _, trustKeyRotation, negativeDiagnostics, _, publisherArtifact, _, correctnessState, toolingState, schemaCompatibilityState, diagnosticsConformanceState, trustKeyRotationState, negativeDiagnosticsState, redactionState, publisherArtifactState, noOverclaimState, valDState := buildVerifierEcosystemValDSharedStates()
	point7State := operability.EvaluateVerifierEcosystemValDPoint7State(valDState)
	limitations := []string{
		"Val D implements the final verifier ecosystem gate only and does not close Točka 7 or create point_7_pass.",
		"Val D remains advisory and does not create certification authority, approval authority, or canonical truth.",
		"Val E remains required for integrated closure.",
	}
	currentState := operability.EvaluateVerifierEcosystemValDProofsState(
		valDState,
		point7State,
		dependency.Val0CurrentState,
		dependency.ValACurrentState,
		dependency.ValBCurrentState,
		dependency.ValCCurrentState,
		verifierEcosystemValDAllSurfaceRefs(),
		verifierEcosystemValDEvidenceRefs(),
		limitations,
		[]string{
			"Val D cannot return point_7_pass and remains a final verifier ecosystem gate only.",
			"Točka 7 remains not complete until Val E integrated closure.",
			"Verifier ecosystem gate outputs remain advisory projections over the canonical execution, audit, and evidence spine.",
		},
		verifierEcosystemValDProjectionDisclaimer(),
	)
	return verifierEcosystemValDProofsResponse{
		SchemaVersion:                   verifierEcosystemValDProofsSchema,
		GeneratedAt:                     publicSampleTime(),
		CurrentState:                    currentState,
		Point5State:                     dependency.Point5State,
		Point5DependencyState:           dependency.Point5DependencyState,
		Point6State:                     dependency.Point6State,
		Point6ClosureState:              dependency.Point6ClosureState,
		Point6ClosurePrerequisite:       dependency.Point6ClosurePrerequisiteState,
		Point6ClosureInvariant:          dependency.Point6ClosureInvariantState,
		Point6ProofSurfaceState:         dependency.Point6ProofSurfaceState,
		Point6PassRuleState:             dependency.Point6PassRuleState,
		Point6PassAllowed:               dependency.Point6PassAllowed,
		Val0CurrentState:                dependency.Val0CurrentState,
		Val0State:                       dependency.Val0State,
		ValACurrentState:                dependency.ValACurrentState,
		ValAState:                       dependency.ValAState,
		ValBCurrentState:                dependency.ValBCurrentState,
		ValBState:                       dependency.ValBState,
		ValCCurrentState:                dependency.ValCCurrentState,
		ValCState:                       dependency.ValCState,
		ValDState:                       valDState,
		Point7State:                     point7State,
		CorrectnessGateState:            correctnessState,
		ToolingGateState:                toolingState,
		SchemaCompatibilityGateState:    schemaCompatibilityState,
		DiagnosticsConformanceGateState: diagnosticsConformanceState,
		TrustKeyRotationGateState:       trustKeyRotationState,
		NegativeDiagnosticsGateState:    negativeDiagnosticsState,
		RedactionGateState:              redactionState,
		PublisherArtifactGateState:      publisherArtifactState,
		NoOverclaimGateState:            noOverclaimState,
		DerivedPublicOutputClass:        negativeDiagnostics.PublicOutputClass,
		DerivedPartnerOutputClass:       negativeDiagnostics.PartnerOutputClass,
		TrustDistributionMode:           trustKeyRotation.TrustDistributionMode,
		OfflineDistributionScope:        trustKeyRotation.OfflineDistributionScope,
		PublisherType:                   publisherArtifact.PublisherType,
		WhyPoint7NotPass: []string{
			"Val D is the final verifier ecosystem gate only and cannot return point_7_pass.",
			"Val E remains the only integrated closure point for Točka 7.",
			"Verifier ecosystem gate outputs remain advisory projections rather than canonical truth or approval authority.",
		},
		SurfaceRefs:          verifierEcosystemValDAllSurfaceRefs(),
		EvidenceRefs:         verifierEcosystemValDEvidenceRefs(),
		Limitations:          limitations,
		ProjectionDisclaimer: verifierEcosystemValDProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val D finalizes verifier ecosystem gate consistency across correctness, tooling, compatibility, diagnostics, trust, redaction, publisher compatibility, and no-overclaim semantics.",
			"Val D builds on actual Točka 6 closure plus active Točka 7 Val 0, Val A, Val B, and Val C proof surfaces while keeping point_7_state not complete.",
		},
	}
}
