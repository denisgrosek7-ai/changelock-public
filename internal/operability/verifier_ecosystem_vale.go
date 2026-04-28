package operability

import "strings"

const (
	VerifierEcosystemValEPrerequisiteStateActive     = "verifier_ecosystem_vale_prerequisites_active"
	VerifierEcosystemValEPrerequisiteStatePartial    = "verifier_ecosystem_vale_prerequisites_partial"
	VerifierEcosystemValEPrerequisiteStateIncomplete = "verifier_ecosystem_vale_prerequisites_incomplete"
	VerifierEcosystemValEPrerequisiteStateBlocked    = "verifier_ecosystem_vale_prerequisites_blocked"
	VerifierEcosystemValEPrerequisiteStateUnknown    = "verifier_ecosystem_vale_prerequisites_unknown"

	VerifierEcosystemValEInvariantStateActive     = "verifier_ecosystem_vale_invariants_active"
	VerifierEcosystemValEInvariantStatePartial    = "verifier_ecosystem_vale_invariants_partial"
	VerifierEcosystemValEInvariantStateIncomplete = "verifier_ecosystem_vale_invariants_incomplete"
	VerifierEcosystemValEInvariantStateBlocked    = "verifier_ecosystem_vale_invariants_blocked"
	VerifierEcosystemValEInvariantStateUnknown    = "verifier_ecosystem_vale_invariants_unknown"

	VerifierEcosystemValEProofSurfaceStateActive     = "verifier_ecosystem_vale_proof_surface_active"
	VerifierEcosystemValEProofSurfaceStatePartial    = "verifier_ecosystem_vale_proof_surface_partial"
	VerifierEcosystemValEProofSurfaceStateIncomplete = "verifier_ecosystem_vale_proof_surface_incomplete"
	VerifierEcosystemValEProofSurfaceStateBlocked    = "verifier_ecosystem_vale_proof_surface_blocked"
	VerifierEcosystemValEProofSurfaceStateUnknown    = "verifier_ecosystem_vale_proof_surface_unknown"

	VerifierEcosystemValEEvidenceQualityStateActive     = "verifier_ecosystem_vale_evidence_quality_active"
	VerifierEcosystemValEEvidenceQualityStatePartial    = "verifier_ecosystem_vale_evidence_quality_partial"
	VerifierEcosystemValEEvidenceQualityStateIncomplete = "verifier_ecosystem_vale_evidence_quality_incomplete"
	VerifierEcosystemValEEvidenceQualityStateBlocked    = "verifier_ecosystem_vale_evidence_quality_blocked"
	VerifierEcosystemValEEvidenceQualityStateUnknown    = "verifier_ecosystem_vale_evidence_quality_unknown"

	VerifierEcosystemValENoOverclaimStateActive     = "verifier_ecosystem_vale_no_overclaim_active"
	VerifierEcosystemValENoOverclaimStatePartial    = "verifier_ecosystem_vale_no_overclaim_partial"
	VerifierEcosystemValENoOverclaimStateIncomplete = "verifier_ecosystem_vale_no_overclaim_incomplete"
	VerifierEcosystemValENoOverclaimStateBlocked    = "verifier_ecosystem_vale_no_overclaim_blocked"
	VerifierEcosystemValENoOverclaimStateUnknown    = "verifier_ecosystem_vale_no_overclaim_unknown"

	VerifierEcosystemValEPassRuleStateActive     = "verifier_ecosystem_vale_pass_rule_active"
	VerifierEcosystemValEPassRuleStatePartial    = "verifier_ecosystem_vale_pass_rule_partial"
	VerifierEcosystemValEPassRuleStateIncomplete = "verifier_ecosystem_vale_pass_rule_incomplete"
	VerifierEcosystemValEPassRuleStateBlocked    = "verifier_ecosystem_vale_pass_rule_blocked"
	VerifierEcosystemValEPassRuleStateUnknown    = "verifier_ecosystem_vale_pass_rule_unknown"

	VerifierEcosystemValEStatePass       = "verifier_ecosystem_vale_pass"
	VerifierEcosystemValEStateActive     = "verifier_ecosystem_vale_active"
	VerifierEcosystemValEStatePartial    = "verifier_ecosystem_vale_partial"
	VerifierEcosystemValEStateIncomplete = "verifier_ecosystem_vale_incomplete"
	VerifierEcosystemValEStateBlocked    = "verifier_ecosystem_vale_blocked"
	VerifierEcosystemValEStateUnknown    = "verifier_ecosystem_vale_unknown"

	VerifierEcosystemValEPoint7PassReasonAllowed                  = "point_7_pass through Val E only after actual Val 0 through Val D proof states, exact proof surfaces, exact evidence refs, fresh evidence quality, and fail-closed closure invariants all remain active."
	VerifierEcosystemValEPoint7PassReasonBlocked                  = "point_7_pass remains blocked until actual Val 0 through Val D proof states, exact proof surfaces, exact evidence refs, fresh evidence quality, and fail-closed closure invariants all remain active."
	VerifierEcosystemValEPoint7PassSafeDiagnosticVal0CannotReturn = "Val 0 cannot return point_7_pass before integrated closure."
	VerifierEcosystemValEPoint7PassSafeDiagnosticValACannotReturn = "Val A cannot return point_7_pass before integrated closure."
	VerifierEcosystemValEPoint7PassSafeDiagnosticValBCannotReturn = "Val B cannot return point_7_pass before integrated closure."
	VerifierEcosystemValEPoint7PassSafeDiagnosticValCCannotReturn = "Val C cannot return point_7_pass before integrated closure."
	VerifierEcosystemValEPoint7PassSafeDiagnosticValDCannotReturn = "Val D cannot return point_7_pass before integrated closure."

	VerifierEcosystemValEClosureInvariantVal0Discipline     = "val0_verifier_discipline_invariant"
	VerifierEcosystemValEClosureInvariantValATooling        = "vala_reference_verifier_tooling_invariant"
	VerifierEcosystemValEClosureInvariantValBCompatibility  = "valb_compatibility_diagnostics_conformance_invariant"
	VerifierEcosystemValEClosureInvariantValCEcosystem      = "valc_public_partner_auditor_publisher_invariant"
	VerifierEcosystemValEClosureInvariantValDFinalGate      = "vald_final_verifier_ecosystem_gate_invariant"
	VerifierEcosystemValEClosureInvariantAdvisoryProjection = "advisory_projection_invariant"
)

const (
	verifierEcosystemValEPoint7PassReasonStateAllowed = "allowed"
	verifierEcosystemValEPoint7PassReasonStateBlocked = "blocked"
	verifierEcosystemValEPoint7PassReasonStateUnknown = "unknown"
)

type VerifierEcosystemValEVal0ProofSnapshot struct {
	CurrentState             string   `json:"current_state"`
	Val0State                string   `json:"val_0_state"`
	Point7State              string   `json:"point_7_state"`
	VerifierContractState    string   `json:"verifier_contract_state"`
	ProofEnvelopeState       string   `json:"proof_envelope_state"`
	VerificationScopeState   string   `json:"verification_scope_state"`
	SchemaCompatibilityState string   `json:"schema_compatibility_state"`
	TrustRootIssuerState     string   `json:"trust_root_issuer_state"`
	DiagnosticsState         string   `json:"diagnostics_state"`
	OutputBoundaryState      string   `json:"output_boundary_state"`
	TrustRootState           string   `json:"trust_root_state"`
	RevocationState          string   `json:"revocation_state"`
	KeyRotationState         string   `json:"key_rotation_state"`
	RolloverMetadataRef      string   `json:"rollover_metadata_ref"`
	SurfaceRefs              []string `json:"surface_refs,omitempty"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
	WorstSeverityPrecedence  bool     `json:"worst_severity_precedence"`
}

type VerifierEcosystemValEValAProofSnapshot struct {
	CurrentState                 string   `json:"current_state"`
	ValAState                    string   `json:"val_a_state"`
	Point7State                  string   `json:"point_7_state"`
	InputModelState              string   `json:"input_model_state"`
	VerifierEngineState          string   `json:"verifier_engine_state"`
	VerificationResultState      string   `json:"verification_result_state"`
	DiagnosticsMappingState      string   `json:"diagnostics_mapping_state"`
	CommandContractState         string   `json:"command_contract_state"`
	SDKEntrypointState           string   `json:"sdk_entrypoint_state"`
	DeterministicOutput          bool     `json:"deterministic_output"`
	HiddenMainInstanceDependency bool     `json:"hidden_main_instance_dependency"`
	NetworkDependency            bool     `json:"network_dependency"`
	MutatesEvidence              bool     `json:"mutates_evidence"`
	ApprovesDeployment           bool     `json:"approves_deployment"`
	SuppressesFailures           bool     `json:"suppresses_failures"`
	ClaimsActualCryptoValidity   bool     `json:"claims_actual_crypto_validity"`
	UsesRealCryptoPrimitives     bool     `json:"uses_real_crypto_primitives"`
	SurfaceRefs                  []string `json:"surface_refs,omitempty"`
	EvidenceRefs                 []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
	WorstSeverityPrecedence      bool     `json:"worst_severity_precedence"`
}

type VerifierEcosystemValEValBProofSnapshot struct {
	CurrentState                  string   `json:"current_state"`
	ValBState                     string   `json:"val_b_state"`
	Point7State                   string   `json:"point_7_state"`
	CompatibilityMatrixState      string   `json:"compatibility_matrix_state"`
	SchemaProofCompatibilityState string   `json:"schema_proof_compatibility_state"`
	MixedVersionDiagnosticState   string   `json:"mixed_version_diagnostic_state"`
	DiagnosticPrecedenceState     string   `json:"diagnostic_precedence_state"`
	FixtureDescriptorState        string   `json:"fixture_descriptor_state"`
	ConformanceCaseState          string   `json:"conformance_case_state"`
	ConformanceSuiteState         string   `json:"conformance_suite_state"`
	OutputClassState              string   `json:"output_class_state"`
	CompatibilityState            string   `json:"compatibility_state"`
	DerivedDiagnosticClass        string   `json:"derived_diagnostic_class"`
	DerivedOutputClass            string   `json:"derived_output_class"`
	NegativeCasesPreserved        bool     `json:"negative_cases_preserved"`
	ConformanceCertificationClaim bool     `json:"conformance_certification_claim"`
	IntegrityRatingClaim          bool     `json:"integrity_rating_claim"`
	SurfaceRefs                   []string `json:"surface_refs,omitempty"`
	EvidenceRefs                  []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	WorstSeverityPrecedence       bool     `json:"worst_severity_precedence"`
}

type VerifierEcosystemValEValCProofSnapshot struct {
	CurrentState                          string   `json:"current_state"`
	ValCState                             string   `json:"val_c_state"`
	Point7State                           string   `json:"point_7_state"`
	AudienceSurfaceState                  string   `json:"audience_surface_state"`
	PublicOutputState                     string   `json:"public_output_state"`
	PartnerOutputState                    string   `json:"partner_output_state"`
	AuditorFlowState                      string   `json:"auditor_flow_state"`
	RequestContractState                  string   `json:"request_contract_state"`
	PublisherProfileState                 string   `json:"publisher_profile_state"`
	ArtifactRuleState                     string   `json:"artifact_rule_state"`
	TrustDistributionState                string   `json:"trust_distribution_state"`
	PublicOutputClass                     string   `json:"public_output_class"`
	PartnerOutputClass                    string   `json:"partner_output_class"`
	RequestMode                           string   `json:"request_mode"`
	PublisherType                         string   `json:"publisher_type"`
	TrustDistributionMode                 string   `json:"trust_distribution_mode"`
	AudienceUniqueBreadthValid            bool     `json:"audience_unique_breadth_valid"`
	AudienceBreadthOrderIndependent       bool     `json:"audience_breadth_order_independent"`
	PublisherApprovedVendorClaim          bool     `json:"publisher_approved_vendor_claim"`
	PublisherCertificationClaim           bool     `json:"publisher_certification_claim"`
	PublisherAutomaticallyTrustedClaim    bool     `json:"publisher_automatically_trusted_claim"`
	TrustDistributionGlobalDirectoryClaim bool     `json:"trust_distribution_global_directory_claim"`
	TrustDistributionKeyRotationState     string   `json:"trust_distribution_key_rotation_state"`
	TrustDistributionRolloverMetadataRef  string   `json:"trust_distribution_rollover_metadata_ref"`
	TrustDistributionTrustRootState       string   `json:"trust_distribution_trust_root_state"`
	TrustDistributionRevocationState      string   `json:"trust_distribution_revocation_state"`
	SurfaceRefs                           []string `json:"surface_refs,omitempty"`
	EvidenceRefs                          []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer                  string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValEValDProofSnapshot struct {
	CurrentState                            string   `json:"current_state"`
	ValDState                               string   `json:"val_d_state"`
	Point7State                             string   `json:"point_7_state"`
	CorrectnessGateState                    string   `json:"correctness_gate_state"`
	ToolingGateState                        string   `json:"tooling_gate_state"`
	SchemaCompatibilityGateState            string   `json:"schema_compatibility_gate_state"`
	DiagnosticsConformanceGateState         string   `json:"diagnostics_conformance_gate_state"`
	TrustKeyRotationGateState               string   `json:"trust_key_rotation_gate_state"`
	NegativeDiagnosticsGateState            string   `json:"negative_diagnostics_gate_state"`
	RedactionGateState                      string   `json:"redaction_gate_state"`
	PublisherArtifactGateState              string   `json:"publisher_artifact_gate_state"`
	NoOverclaimGateState                    string   `json:"no_overclaim_gate_state"`
	TrustDistributionMode                   string   `json:"trust_distribution_mode"`
	OfflineDistributionScope                string   `json:"offline_distribution_scope"`
	TrustDistributionModeUsesActualValCMode bool     `json:"trust_distribution_mode_uses_actual_valc_mode"`
	ClaimsIntegratedClosure                 bool     `json:"claims_integrated_closure"`
	SurfaceRefs                             []string `json:"surface_refs,omitempty"`
	EvidenceRefs                            []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer                    string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValESourceValStates struct {
	Val0State string `json:"val_0_state"`
	ValAState string `json:"val_a_state"`
	ValBState string `json:"val_b_state"`
	ValCState string `json:"val_c_state"`
	ValDState string `json:"val_d_state"`
}

type VerifierEcosystemValESourceCurrentStates struct {
	Val0CurrentState string `json:"val_0_current_state"`
	ValACurrentState string `json:"val_a_current_state"`
	ValBCurrentState string `json:"val_b_current_state"`
	ValCCurrentState string `json:"val_c_current_state"`
	ValDCurrentState string `json:"val_d_current_state"`
}

type VerifierEcosystemValEDependencyStates struct {
	Point5State                    string `json:"point_5_state"`
	Point5DependencyState          string `json:"point_5_dependency_state"`
	Point6State                    string `json:"point_6_state"`
	Point6ClosureState             string `json:"point_6_closure_state"`
	Point6ClosurePrerequisiteState string `json:"point_6_closure_prerequisite_state"`
	Point6ClosureInvariantState    string `json:"point_6_closure_invariant_state"`
	Point6ProofSurfaceState        string `json:"point_6_proof_surface_state"`
	Point6PassRuleState            string `json:"point_6_pass_rule_state"`
	Point6PassAllowed              bool   `json:"point_6_pass_allowed"`
	ValDFinalGateState             string `json:"val_d_final_gate_state"`
	PreClosurePoint7State          string `json:"pre_closure_point_7_state"`
}

type VerifierEcosystemClosureInvariant struct {
	CurrentState         string   `json:"current_state"`
	InvariantID          string   `json:"invariant_id"`
	Title                string   `json:"title"`
	BlockingReasons      []string `json:"blocking_reasons,omitempty"`
	Caveats              []string `json:"caveats,omitempty"`
	EvidenceRefs         []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type VerifierEcosystemIntegratedClosure struct {
	CurrentState                  string                                   `json:"current_state"`
	ClosureID                     string                                   `json:"closure_id"`
	Version                       string                                   `json:"version"`
	Point                         string                                   `json:"point"`
	ClosureVal                    string                                   `json:"closure_val"`
	Point7State                   string                                   `json:"point_7_state"`
	Point7PassAllowed             bool                                     `json:"point_7_pass_allowed"`
	Point7PassReason              string                                   `json:"point_7_pass_reason"`
	SourceValStates               VerifierEcosystemValESourceValStates     `json:"source_val_states"`
	SourceCurrentStates           VerifierEcosystemValESourceCurrentStates `json:"source_current_states"`
	DependencyStates              VerifierEcosystemValEDependencyStates    `json:"dependency_states"`
	Val0                          VerifierEcosystemValEVal0ProofSnapshot   `json:"val_0"`
	ValA                          VerifierEcosystemValEValAProofSnapshot   `json:"val_a"`
	ValB                          VerifierEcosystemValEValBProofSnapshot   `json:"val_b"`
	ValC                          VerifierEcosystemValEValCProofSnapshot   `json:"val_c"`
	ValD                          VerifierEcosystemValEValDProofSnapshot   `json:"val_d"`
	ClosurePrerequisiteState      string                                   `json:"closure_prerequisite_state"`
	ClosureInvariantState         string                                   `json:"closure_invariant_state"`
	ProofSurfaceState             string                                   `json:"proof_surface_state"`
	EvidenceQualityState          string                                   `json:"evidence_quality_state"`
	NoOverclaimState              string                                   `json:"no_overclaim_state"`
	PassRuleState                 string                                   `json:"pass_rule_state"`
	ProofSurfaceRefs              []string                                 `json:"proof_surface_refs,omitempty"`
	EvidenceRefs                  []string                                 `json:"evidence_refs,omitempty"`
	ClosureInvariants             []VerifierEcosystemClosureInvariant      `json:"closure_invariants,omitempty"`
	BlockingReasons               []string                                 `json:"blocking_reasons,omitempty"`
	ObservedClaims                []string                                 `json:"observed_claims,omitempty"`
	Caveats                       []string                                 `json:"caveats,omitempty"`
	Limitations                   []string                                 `json:"limitations,omitempty"`
	ProjectionDisclaimer          string                                   `json:"projection_disclaimer"`
	CreatedAt                     string                                   `json:"created_at"`
	UpdatedAt                     string                                   `json:"updated_at"`
	EvidenceFresh                 bool                                     `json:"evidence_fresh"`
	StaleEvidenceDetected         bool                                     `json:"stale_evidence_detected"`
	RedactionKeepsFailuresVisible bool                                     `json:"redaction_keeps_failures_visible"`
	MutatesCanonicalEvidence      bool                                     `json:"mutates_canonical_evidence"`
	ApprovesDeployment            bool                                     `json:"approves_deployment"`
	SuppressesFailures            bool                                     `json:"suppresses_failures"`
}

func verifierEcosystemValEProjectionDisclaimer() string {
	return "projection_only not_canonical_truth integrated_verifier_ecosystem_closure evidence_linked_verifier_closure"
}

func verifierEcosystemValEHasProjectionDisclaimer(value string) bool {
	return strings.Contains(strings.TrimSpace(value), "projection_only") &&
		strings.Contains(strings.TrimSpace(value), "not_canonical_truth")
}

func VerifierEcosystemValEProofSurfaceRefs() []string {
	return []string{
		"/v1/verifier-ecosystem/val0/proofs",
		"/v1/verifier-ecosystem/vala/proofs",
		"/v1/verifier-ecosystem/valb/proofs",
		"/v1/verifier-ecosystem/valc/proofs",
		"/v1/verifier-ecosystem/vald/proofs",
		"/v1/verifier-ecosystem/vald/correctness-gate",
		"/v1/verifier-ecosystem/vald/tooling-gate",
		"/v1/verifier-ecosystem/vald/schema-compatibility-gate",
		"/v1/verifier-ecosystem/vald/diagnostics-conformance-gate",
		"/v1/verifier-ecosystem/vald/trust-key-rotation-gate",
		"/v1/verifier-ecosystem/vald/negative-diagnostics-gate",
		"/v1/verifier-ecosystem/vald/redaction-gate",
		"/v1/verifier-ecosystem/vald/publisher-artifact-gate",
		"/v1/verifier-ecosystem/vald/no-overclaim-gate",
		"/v1/verifier-ecosystem/vale/closure",
		"/v1/verifier-ecosystem/vale/proofs",
	}
}

func verifierEcosystemValERequiredEvidenceIDs() []string {
	return []string{
		"evidence:vale-closure-001",
		"evidence:vale-prerequisites-001",
		"evidence:vale-invariants-001",
		"evidence:vale-proof-surface-001",
		"evidence:vale-evidence-quality-001",
		"evidence:vale-no-overclaim-001",
		"evidence:vale-point7-governance-001",
	}
}

func verifierEcosystemValERequiredEvidenceScopes() []string {
	return []string{
		"integrated_closure",
		"closure_prerequisites",
		"closure_invariants",
		"proof_surface",
		"evidence_quality",
		"no_overclaim",
		"point7_governance",
	}
}

func VerifierEcosystemValEVerifierEvidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:vale-closure-001", EvidenceType: "integrated_closure", Source: "verifier/vale/closure", Timestamp: "2026-04-28T00:10:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "integrated_closure", Caveats: []string{"integrated verifier ecosystem closure remains bounded to actual Val 0 through Val D proof states only"}},
		{EvidenceID: "evidence:vale-prerequisites-001", EvidenceType: "closure_prerequisites", Source: "verifier/vale/prerequisites", Timestamp: "2026-04-28T00:11:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "closure_prerequisites", Caveats: []string{"point_7_pass requires exact active prerequisite states and fail-closed dependency health"}},
		{EvidenceID: "evidence:vale-invariants-001", EvidenceType: "closure_invariants", Source: "verifier/vale/invariants", Timestamp: "2026-04-28T00:12:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "closure_invariants", Caveats: []string{"cross-val invariants remain advisory, evidence-linked, and fail-closed"}},
		{EvidenceID: "evidence:vale-proof-surface-001", EvidenceType: "proof_surface", Source: "verifier/vale/proof-surface", Timestamp: "2026-04-28T00:13:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "proof_surface", Caveats: []string{"integrated closure requires the exact declared proof surface set"}},
		{EvidenceID: "evidence:vale-evidence-quality-001", EvidenceType: "evidence_quality", Source: "verifier/vale/evidence-quality", Timestamp: "2026-04-28T00:14:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "evidence_quality", Caveats: []string{"integrated closure evidence remains exact-set, fresh, and non-synthetic"}},
		{EvidenceID: "evidence:vale-no-overclaim-001", EvidenceType: "no_overclaim", Source: "verifier/vale/no-overclaim", Timestamp: "2026-04-28T00:15:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "no_overclaim", Caveats: []string{"integrated closure blocks certification, approval, rating, and authority claims"}},
		{EvidenceID: "evidence:vale-point7-governance-001", EvidenceType: "state_governance", Source: "verifier/point7-governance", Timestamp: "2026-04-28T00:16:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point7_governance", Caveats: []string{"point_7_pass is returned through Val E only after exact prerequisite, invariant, and evidence checks"}},
	}
}

func verifierEcosystemValAExpectedProofEvidenceRefs() []string {
	refs := append([]string{}, VerifierEcosystemVal0ProofEvidenceRefs()...)
	refs = append(refs, "point7_verifier_discipline_foundation")
	refs = append(refs, "point7_reference_verifier_tooling")
	for _, evidence := range VerifierEcosystemValAVerifierEvidence() {
		if strings.TrimSpace(evidence.EvidenceID) != "" {
			refs = append(refs, evidence.EvidenceID)
		}
	}
	return refs
}

func VerifierEcosystemValEProofEvidenceRefs() []string {
	return []string{
		"point6_integrated_closure",
		"point7_verifier_discipline_foundation",
		"point7_reference_verifier_tooling",
		"point7_compatibility_diagnostics_conformance",
		"point7_public_partner_auditor_publisher_ecosystem",
		"point7_final_verifier_ecosystem_gate",
		"point7_trust_root_issuer_discipline",
		"point7_diagnostics_conformance",
		"point7_redaction_output_boundary",
		"point7_no_overclaim_governance",
		"point7_governance",
		"point7_integrated_verifier_ecosystem_closure",
		"evidence:vale-closure-001",
		"evidence:vale-prerequisites-001",
		"evidence:vale-invariants-001",
		"evidence:vale-proof-surface-001",
		"evidence:vale-evidence-quality-001",
		"evidence:vale-no-overclaim-001",
		"evidence:vale-point7-governance-001",
	}
}

func verifierEcosystemValEPassAllowedClaim(values ...string) bool {
	for _, value := range values {
		if verifierEcosystemValEPoint7PassReasonState(value) == verifierEcosystemValEPoint7PassReasonStateAllowed {
			return true
		}
	}
	return false
}

func verifierEcosystemValENormalizeText(value string) string {
	return strings.Join(strings.Fields(strings.ToLower(strings.TrimSpace(value))), " ")
}

func verifierEcosystemValEPoint7PassReasonState(value string) string {
	switch verifierEcosystemValENormalizeText(value) {
	case verifierEcosystemValENormalizeText(VerifierEcosystemValEPoint7PassReasonAllowed):
		return verifierEcosystemValEPoint7PassReasonStateAllowed
	case verifierEcosystemValENormalizeText(VerifierEcosystemValEPoint7PassReasonBlocked):
		return verifierEcosystemValEPoint7PassReasonStateBlocked
	default:
		return verifierEcosystemValEPoint7PassReasonStateUnknown
	}
}

func verifierEcosystemValEExactSafePoint7PassDiagnostic(value string) bool {
	normalized := verifierEcosystemValENormalizeText(value)
	if normalized == "" {
		return false
	}
	safe := []string{
		VerifierEcosystemValEPoint7PassReasonAllowed,
		VerifierEcosystemValEPoint7PassReasonBlocked,
		VerifierEcosystemValEPoint7PassSafeDiagnosticVal0CannotReturn,
		VerifierEcosystemValEPoint7PassSafeDiagnosticValACannotReturn,
		VerifierEcosystemValEPoint7PassSafeDiagnosticValBCannotReturn,
		VerifierEcosystemValEPoint7PassSafeDiagnosticValCCannotReturn,
		VerifierEcosystemValEPoint7PassSafeDiagnosticValDCannotReturn,
	}
	for _, candidate := range safe {
		if normalized == verifierEcosystemValENormalizeText(candidate) {
			return true
		}
	}
	return false
}

func verifierEcosystemValEContainsDisallowedClaim(values ...string) bool {
	disallowed := []string{
		"verifier certification",
		"certified verifier",
		"certified publisher",
		"certified vendor",
		"approved vendor",
		"integrity rating",
		"marketplace rating",
		"anyone can verify everything",
		"mathematically proves total truth",
		"universal trust protocol",
		"regulator-approved verifier",
		"global key registry for all instances",
		"formal certification",
		"absolute proof",
		"universal authority",
		"deployment approved",
		"production approved",
	}
	for _, value := range values {
		normalized := verifierEcosystemValENormalizeText(value)
		if normalized == "" {
			continue
		}
		for _, claim := range disallowed {
			if strings.Contains(normalized, claim) {
				return true
			}
		}
		if strings.Contains(normalized, "point_7_pass") && !verifierEcosystemValEExactSafePoint7PassDiagnostic(value) {
			return true
		}
	}
	return false
}

func verifierEcosystemValEPoint6DependencySeverity(snapshot VerifierEcosystemValEDependencyStates) int {
	if strings.TrimSpace(snapshot.Point6State) == "" ||
		strings.TrimSpace(snapshot.Point6ClosureState) == "" ||
		strings.TrimSpace(snapshot.Point6ClosurePrerequisiteState) == "" ||
		strings.TrimSpace(snapshot.Point6ClosureInvariantState) == "" ||
		strings.TrimSpace(snapshot.Point6ProofSurfaceState) == "" ||
		strings.TrimSpace(snapshot.Point6PassRuleState) == "" {
		return 2
	}
	if strings.TrimSpace(snapshot.Point6State) != ReferenceArchitecturePoint6StatePass || !snapshot.Point6PassAllowed {
		return 4
	}
	return verifierEcosystemValDHighestSeverity(
		verifierEcosystemValDStateSeverity(snapshot.Point6ClosureState, ReferenceArchitectureValEStateActive, ReferenceArchitectureValEStatePartial, ReferenceArchitectureValEStateIncomplete, ReferenceArchitectureValEStateBlocked, ReferenceArchitectureValEStateUnknown),
		verifierEcosystemValDStateSeverity(snapshot.Point6ClosurePrerequisiteState, ReferenceArchitectureValEPrerequisiteStateActive, ReferenceArchitectureValEPrerequisiteStatePartial, ReferenceArchitectureValEPrerequisiteStateIncomplete, ReferenceArchitectureValEPrerequisiteStateBlocked, ReferenceArchitectureValEPrerequisiteStateUnknown),
		verifierEcosystemValDStateSeverity(snapshot.Point6ClosureInvariantState, ReferenceArchitectureValEInvariantStateActive, ReferenceArchitectureValEInvariantStatePartial, ReferenceArchitectureValEInvariantStateIncomplete, ReferenceArchitectureValEInvariantStateBlocked, ReferenceArchitectureValEInvariantStateUnknown),
		verifierEcosystemValDStateSeverity(snapshot.Point6ProofSurfaceState, ReferenceArchitectureValEProofSurfaceStateActive, ReferenceArchitectureValEProofSurfaceStatePartial, ReferenceArchitectureValEProofSurfaceStateIncomplete, ReferenceArchitectureValEProofSurfaceStateBlocked, ReferenceArchitectureValEProofSurfaceStateUnknown),
		verifierEcosystemValDStateSeverity(snapshot.Point6PassRuleState, ReferenceArchitectureValEPassRuleStateActive, ReferenceArchitectureValEPassRuleStatePartial, ReferenceArchitectureValEPassRuleStateIncomplete, ReferenceArchitectureValEPassRuleStateBlocked, ReferenceArchitectureValEPassRuleStateUnknown),
	)
}

func verifierEcosystemValEInvariant(id, title, currentState string, blockingReasons, caveats, evidenceRefs []string) VerifierEcosystemClosureInvariant {
	return VerifierEcosystemClosureInvariant{
		CurrentState:         currentState,
		InvariantID:          id,
		Title:                title,
		BlockingReasons:      blockingReasons,
		Caveats:              caveats,
		EvidenceRefs:         evidenceRefs,
		ProjectionDisclaimer: verifierEcosystemValEProjectionDisclaimer(),
	}
}

func EvaluateVerifierEcosystemValEPrerequisiteState(model VerifierEcosystemIntegratedClosure) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ClosureID,
		model.Point,
		model.ClosureVal,
		model.Point7PassReason,
		model.ProjectionDisclaimer,
	) {
		return VerifierEcosystemValEPrerequisiteStateIncomplete
	}
	if strings.TrimSpace(model.Point) != "point_7" ||
		strings.TrimSpace(model.ClosureVal) != "val_e" ||
		!verifierEcosystemValEHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValEPrerequisiteStateUnknown
	}
	point6Severity := verifierEcosystemValEPoint6DependencySeverity(model.DependencyStates)
	val0CurrentSeverity, ok := verifierEcosystemValDPriorValStateSeverity(model.SourceCurrentStates.Val0CurrentState)
	if !ok {
		return VerifierEcosystemValEPrerequisiteStateUnknown
	}
	val0ValSeverity, ok := verifierEcosystemValDPriorValStateSeverity(model.SourceValStates.Val0State)
	if !ok {
		return VerifierEcosystemValEPrerequisiteStateUnknown
	}
	valACurrentSeverity, ok := verifierEcosystemValDPriorValStateSeverity(model.SourceCurrentStates.ValACurrentState)
	if !ok {
		return VerifierEcosystemValEPrerequisiteStateUnknown
	}
	valAValSeverity, ok := verifierEcosystemValDPriorValStateSeverity(model.SourceValStates.ValAState)
	if !ok {
		return VerifierEcosystemValEPrerequisiteStateUnknown
	}
	valBCurrentSeverity, ok := verifierEcosystemValDPriorValStateSeverity(model.SourceCurrentStates.ValBCurrentState)
	if !ok {
		return VerifierEcosystemValEPrerequisiteStateUnknown
	}
	valBValSeverity, ok := verifierEcosystemValDPriorValStateSeverity(model.SourceValStates.ValBState)
	if !ok {
		return VerifierEcosystemValEPrerequisiteStateUnknown
	}
	valCCurrentSeverity, ok := verifierEcosystemValDPriorValStateSeverity(model.SourceCurrentStates.ValCCurrentState)
	if !ok {
		return VerifierEcosystemValEPrerequisiteStateUnknown
	}
	valCValSeverity, ok := verifierEcosystemValDPriorValStateSeverity(model.SourceValStates.ValCState)
	if !ok {
		return VerifierEcosystemValEPrerequisiteStateUnknown
	}
	valDCurrentSeverity := verifierEcosystemValDStateSeverity(model.SourceCurrentStates.ValDCurrentState, VerifierEcosystemValDStateActive, VerifierEcosystemValDStatePartial, VerifierEcosystemValDStateIncomplete, VerifierEcosystemValDStateBlocked, VerifierEcosystemValDStateUnknown)
	valDValSeverity := verifierEcosystemValDStateSeverity(model.SourceValStates.ValDState, VerifierEcosystemValDStateActive, VerifierEcosystemValDStatePartial, VerifierEcosystemValDStateIncomplete, VerifierEcosystemValDStateBlocked, VerifierEcosystemValDStateUnknown)
	valDFinalGateSeverity := verifierEcosystemValDStateSeverity(model.DependencyStates.ValDFinalGateState, VerifierEcosystemValDStateActive, VerifierEcosystemValDStatePartial, VerifierEcosystemValDStateIncomplete, VerifierEcosystemValDStateBlocked, VerifierEcosystemValDStateUnknown)
	preClosurePoint7Severity := 3
	switch strings.TrimSpace(model.DependencyStates.PreClosurePoint7State) {
	case VerifierEcosystemPoint7StateNotComplete:
		preClosurePoint7Severity = 0
	case VerifierEcosystemPoint7StatePass:
		preClosurePoint7Severity = 4
	}
	highest := verifierEcosystemValDHighestSeverity(
		point6Severity,
		val0CurrentSeverity,
		val0ValSeverity,
		valACurrentSeverity,
		valAValSeverity,
		valBCurrentSeverity,
		valBValSeverity,
		valCCurrentSeverity,
		valCValSeverity,
		valDCurrentSeverity,
		valDValSeverity,
		valDFinalGateSeverity,
		preClosurePoint7Severity,
	)
	return verifierEcosystemValDSeverityToState(
		highest,
		VerifierEcosystemValEPrerequisiteStateActive,
		VerifierEcosystemValEPrerequisiteStatePartial,
		VerifierEcosystemValEPrerequisiteStateIncomplete,
		VerifierEcosystemValEPrerequisiteStateBlocked,
		VerifierEcosystemValEPrerequisiteStateUnknown,
	)
}

func evaluateVerifierEcosystemValEVal0Invariant(model VerifierEcosystemIntegratedClosure) VerifierEcosystemClosureInvariant {
	blockingReasons := []string{}
	val0 := model.Val0
	if val0.CurrentState != VerifierEcosystemVal0StateActive || val0.Val0State != VerifierEcosystemVal0StateActive {
		blockingReasons = append(blockingReasons, "Val 0 current and val states must remain active.")
	}
	if val0.Point7State != VerifierEcosystemPoint7StateNotComplete {
		blockingReasons = append(blockingReasons, VerifierEcosystemValEPoint7PassSafeDiagnosticVal0CannotReturn)
	}
	if val0.VerifierContractState != VerifierEcosystemVal0ContractStateActive ||
		val0.ProofEnvelopeState != VerifierEcosystemVal0EnvelopeStateActive ||
		val0.VerificationScopeState != VerifierEcosystemVal0ScopeStateActive ||
		val0.SchemaCompatibilityState != VerifierEcosystemVal0CompatibilityStateActive ||
		val0.TrustRootIssuerState != VerifierEcosystemVal0TrustStateActive ||
		val0.DiagnosticsState != VerifierEcosystemVal0DiagnosticsStateActive ||
		val0.OutputBoundaryState != VerifierEcosystemVal0OutputBoundaryStateActive {
		blockingReasons = append(blockingReasons, "Val 0 verifier discipline, envelope, scope, compatibility, trust, diagnostics, and output boundary states must all remain active.")
	}
	if (val0.RevocationState == VerifierEcosystemRevocationRevoked ||
		val0.RevocationState == VerifierEcosystemRevocationExpired ||
		val0.RevocationState == VerifierEcosystemRevocationUnsupported) &&
		val0.TrustRootIssuerState == VerifierEcosystemVal0TrustStateActive {
		blockingReasons = append(blockingReasons, "Revoked, expired, or unsupported issuer material cannot remain active.")
	}
	if (val0.TrustRootState == VerifierEcosystemTrustRootRevoked ||
		val0.TrustRootState == VerifierEcosystemTrustRootExpired ||
		val0.TrustRootState == VerifierEcosystemTrustRootUnsupported) &&
		(val0.TrustRootIssuerState == VerifierEcosystemVal0TrustStateActive || val0.TrustRootIssuerState == VerifierEcosystemVal0TrustStatePartial) {
		blockingReasons = append(blockingReasons, "Revoked, expired, or unsupported trust-root material cannot be downgraded into active or partial trust state.")
	}
	if strings.TrimSpace(val0.KeyRotationState) == VerifierEcosystemKeyRotationRollover && strings.TrimSpace(val0.RolloverMetadataRef) == "" {
		blockingReasons = append(blockingReasons, "Val 0 rollover-in-progress requires explicit rollover metadata and cannot fail open.")
	}
	if !val0.WorstSeverityPrecedence {
		blockingReasons = append(blockingReasons, "Val 0 aggregate severity precedence must remain worst-state rather than first-non-active.")
	}
	val0EvidenceFresh, val0EvidenceStale, val0EvidenceOK := verifierEcosystemVal0EvidenceValid(VerifierEcosystemVal0VerifierEvidence())
	if !containsExactTrimmedStringSet(val0.SurfaceRefs, VerifierEcosystemVal0ProofSurfaceRefs()...) ||
		!containsExactTrimmedStringSet(val0.EvidenceRefs, VerifierEcosystemVal0ProofEvidenceRefs()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(val0.ProjectionDisclaimer) ||
		!val0EvidenceFresh || val0EvidenceStale || !val0EvidenceOK {
		blockingReasons = append(blockingReasons, "Val 0 proof surfaces, evidence refs, and advisory boundary must remain exact and quality-checked.")
	}
	state := VerifierEcosystemValEInvariantStateActive
	if len(blockingReasons) > 0 {
		state = VerifierEcosystemValEInvariantStateBlocked
	}
	return verifierEcosystemValEInvariant(
		VerifierEcosystemValEClosureInvariantVal0Discipline,
		"Val 0 Verifier Discipline Invariant",
		state,
		blockingReasons,
		nil,
		val0.EvidenceRefs,
	)
}

func evaluateVerifierEcosystemValEValAInvariant(model VerifierEcosystemIntegratedClosure) VerifierEcosystemClosureInvariant {
	blockingReasons := []string{}
	valA := model.ValA
	if valA.CurrentState != VerifierEcosystemValAStateActive || valA.ValAState != VerifierEcosystemValAStateActive {
		blockingReasons = append(blockingReasons, "Val A current and val states must remain active.")
	}
	if valA.Point7State != VerifierEcosystemPoint7StateNotComplete {
		blockingReasons = append(blockingReasons, VerifierEcosystemValEPoint7PassSafeDiagnosticValACannotReturn)
	}
	if valA.InputModelState != VerifierEcosystemValAInputStateActive ||
		valA.VerifierEngineState != VerifierEcosystemValAEngineStateActive ||
		valA.VerificationResultState != VerifierEcosystemValAResultStateActive ||
		valA.DiagnosticsMappingState != VerifierEcosystemValADiagnosticsMappingStateActive ||
		valA.CommandContractState != VerifierEcosystemValACommandContractStateActive ||
		valA.SDKEntrypointState != VerifierEcosystemValASDKEntrypointStateActive {
		blockingReasons = append(blockingReasons, "Val A input, engine, result, diagnostics mapping, command, and SDK states must all remain active.")
	}
	if !valA.DeterministicOutput {
		blockingReasons = append(blockingReasons, "Val A deterministic verifier output must remain enabled.")
	}
	if valA.HiddenMainInstanceDependency || valA.NetworkDependency || valA.MutatesEvidence || valA.ApprovesDeployment || valA.SuppressesFailures {
		blockingReasons = append(blockingReasons, "Val A tooling must not depend on hidden instance state, network, mutation, approval, or suppression behavior.")
	}
	if valA.ClaimsActualCryptoValidity && !valA.UsesRealCryptoPrimitives {
		blockingReasons = append(blockingReasons, "Val A must not claim real cryptographic verification when real primitives are not used.")
	}
	if !valA.WorstSeverityPrecedence {
		blockingReasons = append(blockingReasons, "Val A aggregate severity precedence must remain worst-state and keep hard guardrails visible.")
	}
	valAEvidenceFresh, valAEvidenceStale, valAEvidenceOK := verifierEcosystemVal0EvidenceValid(VerifierEcosystemValAVerifierEvidence())
	if !containsExactTrimmedStringSet(valA.SurfaceRefs, VerifierEcosystemValAProofSurfaceRefs()...) ||
		!containsExactTrimmedStringSet(valA.EvidenceRefs, verifierEcosystemValAExpectedProofEvidenceRefs()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(valA.ProjectionDisclaimer) ||
		!valAEvidenceFresh || valAEvidenceStale || !valAEvidenceOK {
		blockingReasons = append(blockingReasons, "Val A proof surfaces, evidence refs, and advisory boundary must remain exact and evidence-backed.")
	}
	state := VerifierEcosystemValEInvariantStateActive
	if len(blockingReasons) > 0 {
		state = VerifierEcosystemValEInvariantStateBlocked
	}
	return verifierEcosystemValEInvariant(
		VerifierEcosystemValEClosureInvariantValATooling,
		"Val A Reference Verifier Tooling Invariant",
		state,
		blockingReasons,
		nil,
		valA.EvidenceRefs,
	)
}

func evaluateVerifierEcosystemValEValBInvariant(model VerifierEcosystemIntegratedClosure) VerifierEcosystemClosureInvariant {
	blockingReasons := []string{}
	valB := model.ValB
	if valB.CurrentState != VerifierEcosystemValBStateActive || valB.ValBState != VerifierEcosystemValBStateActive {
		blockingReasons = append(blockingReasons, "Val B current and val states must remain active.")
	}
	if valB.Point7State != VerifierEcosystemPoint7StateNotComplete {
		blockingReasons = append(blockingReasons, VerifierEcosystemValEPoint7PassSafeDiagnosticValBCannotReturn)
	}
	if valB.CompatibilityMatrixState != VerifierEcosystemValBCompatibilityMatrixStateActive ||
		valB.SchemaProofCompatibilityState != VerifierEcosystemValBSchemaProofCompatibilityStateActive ||
		valB.MixedVersionDiagnosticState != VerifierEcosystemValBMixedVersionStateActive ||
		valB.DiagnosticPrecedenceState != VerifierEcosystemValBDiagnosticPrecedenceStateActive ||
		valB.FixtureDescriptorState != VerifierEcosystemValBFixtureDescriptorStateActive ||
		valB.ConformanceCaseState != VerifierEcosystemValBConformanceCaseStateActive ||
		valB.ConformanceSuiteState != VerifierEcosystemValBConformanceSuiteStateActive ||
		valB.OutputClassState != VerifierEcosystemValBOutputClassStateActive {
		blockingReasons = append(blockingReasons, "Val B compatibility, mixed-version, diagnostics, fixtures, conformance, and output class states must all remain active.")
	}
	if valB.CompatibilityState == ReferenceArchitectureCompatibilityUnsupported || valB.CompatibilityState == ReferenceArchitectureCompatibilityUnknown {
		blockingReasons = append(blockingReasons, "Unsupported or unknown compatibility state cannot pass integrated closure.")
	}
	if valB.ConformanceCertificationClaim || valB.IntegrityRatingClaim {
		blockingReasons = append(blockingReasons, "Val B conformance outputs must not create verifier certification or ratings.")
	}
	if !valB.NegativeCasesPreserved {
		blockingReasons = append(blockingReasons, "Val B unsupported, stale, revoked, superseded, and malformed cases must remain non-verified.")
	}
	if !valB.WorstSeverityPrecedence {
		blockingReasons = append(blockingReasons, "Val B aggregate severity precedence must remain worst-state.")
	}
	if !containsExactTrimmedStringSet(valB.SurfaceRefs, VerifierEcosystemValBProofSurfaceRefs()...) ||
		!verifierEcosystemValBProofEvidenceQualityValid(VerifierEcosystemValBVerifierEvidence(), valB.EvidenceRefs) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(valB.ProjectionDisclaimer) {
		blockingReasons = append(blockingReasons, "Val B proof surfaces, evidence refs, and advisory boundary must remain exact and quality-checked.")
	}
	state := VerifierEcosystemValEInvariantStateActive
	if len(blockingReasons) > 0 {
		state = VerifierEcosystemValEInvariantStateBlocked
	}
	return verifierEcosystemValEInvariant(
		VerifierEcosystemValEClosureInvariantValBCompatibility,
		"Val B Compatibility, Diagnostics, and Conformance Invariant",
		state,
		blockingReasons,
		nil,
		valB.EvidenceRefs,
	)
}

func evaluateVerifierEcosystemValEValCInvariant(model VerifierEcosystemIntegratedClosure) VerifierEcosystemClosureInvariant {
	blockingReasons := []string{}
	valC := model.ValC
	if valC.CurrentState != VerifierEcosystemValCStateActive || valC.ValCState != VerifierEcosystemValCStateActive {
		blockingReasons = append(blockingReasons, "Val C current and val states must remain active.")
	}
	if valC.Point7State != VerifierEcosystemPoint7StateNotComplete {
		blockingReasons = append(blockingReasons, VerifierEcosystemValEPoint7PassSafeDiagnosticValCCannotReturn)
	}
	if valC.AudienceSurfaceState != VerifierEcosystemValCAudienceSurfaceStateActive ||
		valC.PublicOutputState != VerifierEcosystemValCPublicOutputStateActive ||
		valC.PartnerOutputState != VerifierEcosystemValCPartnerOutputStateActive ||
		valC.AuditorFlowState != VerifierEcosystemValCAuditorFlowStateActive ||
		valC.RequestContractState != VerifierEcosystemValCRequestContractStateActive ||
		valC.PublisherProfileState != VerifierEcosystemValCPublisherProfileStateActive ||
		valC.ArtifactRuleState != VerifierEcosystemValCArtifactRuleStateActive ||
		valC.TrustDistributionState != VerifierEcosystemValCTrustDistributionStateActive {
		blockingReasons = append(blockingReasons, "Val C audience, output, auditor, request, publisher, artifact, and trust distribution states must all remain active.")
	}
	if !valC.AudienceBreadthOrderIndependent {
		blockingReasons = append(blockingReasons, "Val C partner/public breadth validation must remain order-independent.")
	}
	if !valC.AudienceUniqueBreadthValid {
		blockingReasons = append(blockingReasons, "Val C partner/public breadth validation must use unique normalized output classes.")
	}
	if strings.TrimSpace(valC.TrustDistributionKeyRotationState) == VerifierEcosystemKeyRotationRollover &&
		strings.TrimSpace(valC.TrustDistributionRolloverMetadataRef) == "" {
		blockingReasons = append(blockingReasons, "Val C trust distribution rollover requires explicit metadata before partial trust can be returned.")
	}
	if valC.TrustDistributionTrustRootState == VerifierEcosystemTrustRootRevoked ||
		valC.TrustDistributionTrustRootState == VerifierEcosystemTrustRootExpired ||
		valC.TrustDistributionTrustRootState == VerifierEcosystemTrustRootUnsupported ||
		valC.TrustDistributionRevocationState == VerifierEcosystemRevocationRevoked ||
		valC.TrustDistributionRevocationState == VerifierEcosystemRevocationExpired ||
		valC.TrustDistributionRevocationState == VerifierEcosystemRevocationUnsupported {
		blockingReasons = append(blockingReasons, "Val C trust distribution must not hide revoked, expired, or unsupported trust material.")
	}
	if valC.PublisherApprovedVendorClaim || valC.PublisherCertificationClaim || valC.PublisherAutomaticallyTrustedClaim {
		blockingReasons = append(blockingReasons, "Val C publisher compatibility must not imply approved vendor, certification, or automatic trust.")
	}
	if valC.TrustDistributionGlobalDirectoryClaim {
		blockingReasons = append(blockingReasons, "Val C trust distribution must remain scoped and cannot claim a global key directory.")
	}
	if !containsExactTrimmedStringSet(valC.SurfaceRefs, VerifierEcosystemValCProofSurfaceRefs()...) ||
		!verifierEcosystemValCProofEvidenceQualityValid(VerifierEcosystemValCVerifierEvidence(), valC.EvidenceRefs) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(valC.ProjectionDisclaimer) {
		blockingReasons = append(blockingReasons, "Val C proof surfaces, evidence refs, and advisory boundary must remain exact and quality-checked.")
	}
	state := VerifierEcosystemValEInvariantStateActive
	if len(blockingReasons) > 0 {
		state = VerifierEcosystemValEInvariantStateBlocked
	}
	return verifierEcosystemValEInvariant(
		VerifierEcosystemValEClosureInvariantValCEcosystem,
		"Val C Public, Partner, Auditor, and Publisher Invariant",
		state,
		blockingReasons,
		nil,
		valC.EvidenceRefs,
	)
}

func evaluateVerifierEcosystemValEValDInvariant(model VerifierEcosystemIntegratedClosure) VerifierEcosystemClosureInvariant {
	blockingReasons := []string{}
	valD := model.ValD
	if valD.CurrentState != VerifierEcosystemValDStateActive || valD.ValDState != VerifierEcosystemValDStateActive {
		blockingReasons = append(blockingReasons, "Val D current and val states must remain active.")
	}
	if valD.Point7State != VerifierEcosystemPoint7StateNotComplete {
		blockingReasons = append(blockingReasons, VerifierEcosystemValEPoint7PassSafeDiagnosticValDCannotReturn)
	}
	if valD.CorrectnessGateState != VerifierEcosystemValDCorrectnessGateStateActive ||
		valD.ToolingGateState != VerifierEcosystemValDToolingGateStateActive ||
		valD.SchemaCompatibilityGateState != VerifierEcosystemValDSchemaCompatibilityGateStateActive ||
		valD.DiagnosticsConformanceGateState != VerifierEcosystemValDDiagnosticsConformanceGateStateActive ||
		valD.TrustKeyRotationGateState != VerifierEcosystemValDTrustKeyRotationGateStateActive ||
		valD.NegativeDiagnosticsGateState != VerifierEcosystemValDNegativeDiagnosticsGateStateActive ||
		valD.RedactionGateState != VerifierEcosystemValDRedactionGateStateActive ||
		valD.PublisherArtifactGateState != VerifierEcosystemValDPublisherArtifactGateStateActive ||
		valD.NoOverclaimGateState != VerifierEcosystemValDNoOverclaimGateStateActive {
		blockingReasons = append(blockingReasons, "Val D correctness, tooling, compatibility, diagnostics, trust, redaction, publisher, and no-overclaim gates must all remain active.")
	}
	if !valD.TrustDistributionModeUsesActualValCMode || strings.TrimSpace(valD.TrustDistributionMode) == strings.TrimSpace(valD.OfflineDistributionScope) {
		blockingReasons = append(blockingReasons, "Val D trust distribution mode must be populated from the actual Val C distribution mode and remain separate from offline scope.")
	}
	if valD.ClaimsIntegratedClosure {
		blockingReasons = append(blockingReasons, "Val D must not claim integrated closure.")
	}
	if !containsExactTrimmedStringSet(valD.SurfaceRefs, VerifierEcosystemValDProofSurfaceRefs()...) ||
		!verifierEcosystemValDProofEvidenceQualityValid(VerifierEcosystemValDVerifierEvidence(), valD.EvidenceRefs) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(valD.ProjectionDisclaimer) {
		blockingReasons = append(blockingReasons, "Val D proof surfaces, evidence refs, and advisory boundary must remain exact and quality-checked.")
	}
	state := VerifierEcosystemValEInvariantStateActive
	if len(blockingReasons) > 0 {
		state = VerifierEcosystemValEInvariantStateBlocked
	}
	return verifierEcosystemValEInvariant(
		VerifierEcosystemValEClosureInvariantValDFinalGate,
		"Val D Final Verifier Ecosystem Gate Invariant",
		state,
		blockingReasons,
		nil,
		valD.EvidenceRefs,
	)
}

func evaluateVerifierEcosystemValEAdvisoryInvariant(model VerifierEcosystemIntegratedClosure) VerifierEcosystemClosureInvariant {
	blockingReasons := []string{}
	if !verifierEcosystemValEHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!verifierEcosystemValEHasProjectionDisclaimer(model.Val0.ProjectionDisclaimer) ||
		!verifierEcosystemValEHasProjectionDisclaimer(model.ValA.ProjectionDisclaimer) ||
		!verifierEcosystemValEHasProjectionDisclaimer(model.ValB.ProjectionDisclaimer) ||
		!verifierEcosystemValEHasProjectionDisclaimer(model.ValC.ProjectionDisclaimer) ||
		!verifierEcosystemValEHasProjectionDisclaimer(model.ValD.ProjectionDisclaimer) {
		blockingReasons = append(blockingReasons, "All Val 0 through Val E surfaces must remain advisory projections and not canonical truth.")
	}
	if model.MutatesCanonicalEvidence || model.ApprovesDeployment || model.SuppressesFailures {
		blockingReasons = append(blockingReasons, "Integrated closure must not mutate canonical evidence, approve deployment, or suppress failures.")
	}
	if !model.RedactionKeepsFailuresVisible {
		blockingReasons = append(blockingReasons, "Redaction and caveats must not convert non-verified outcomes into pass.")
	}
	if verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), strings.Join(model.Limitations, " "), model.ProjectionDisclaimer) ||
		verifierEcosystemValEContainsDisallowedClaim(model.ObservedClaims...) {
		blockingReasons = append(blockingReasons, "Integrated closure outputs must not claim certification, rating, universal authority, or out-of-scope pass semantics.")
	}
	state := VerifierEcosystemValEInvariantStateActive
	if len(blockingReasons) > 0 {
		state = VerifierEcosystemValEInvariantStateBlocked
	}
	return verifierEcosystemValEInvariant(
		VerifierEcosystemValEClosureInvariantAdvisoryProjection,
		"Advisory Projection Invariant",
		state,
		blockingReasons,
		nil,
		model.EvidenceRefs,
	)
}

func VerifierEcosystemValEClosureInvariants(model VerifierEcosystemIntegratedClosure) []VerifierEcosystemClosureInvariant {
	return []VerifierEcosystemClosureInvariant{
		evaluateVerifierEcosystemValEVal0Invariant(model),
		evaluateVerifierEcosystemValEValAInvariant(model),
		evaluateVerifierEcosystemValEValBInvariant(model),
		evaluateVerifierEcosystemValEValCInvariant(model),
		evaluateVerifierEcosystemValEValDInvariant(model),
		evaluateVerifierEcosystemValEAdvisoryInvariant(model),
	}
}

func EvaluateVerifierEcosystemValEInvariantState(model VerifierEcosystemIntegratedClosure) string {
	invariants := VerifierEcosystemValEClosureInvariants(model)
	if len(invariants) == 0 {
		return VerifierEcosystemValEInvariantStateIncomplete
	}
	allActive := true
	for _, invariant := range invariants {
		if strings.TrimSpace(invariant.CurrentState) == "" {
			return VerifierEcosystemValEInvariantStateIncomplete
		}
		if !verifierEcosystemValEHasProjectionDisclaimer(invariant.ProjectionDisclaimer) {
			return VerifierEcosystemValEInvariantStatePartial
		}
		if invariant.CurrentState != VerifierEcosystemValEInvariantStateActive {
			allActive = false
		}
	}
	if allActive {
		return VerifierEcosystemValEInvariantStateActive
	}
	for _, invariant := range invariants {
		switch invariant.CurrentState {
		case VerifierEcosystemValEInvariantStateBlocked:
			return VerifierEcosystemValEInvariantStateBlocked
		case VerifierEcosystemValEInvariantStateUnknown:
			return VerifierEcosystemValEInvariantStateUnknown
		case VerifierEcosystemValEInvariantStateIncomplete:
			return VerifierEcosystemValEInvariantStateIncomplete
		}
	}
	return VerifierEcosystemValEInvariantStatePartial
}

func EvaluateVerifierEcosystemValEProofSurfaceState(model VerifierEcosystemIntegratedClosure) string {
	if strings.TrimSpace(model.ProjectionDisclaimer) == "" || len(model.ProofSurfaceRefs) == 0 {
		return VerifierEcosystemValEProofSurfaceStateIncomplete
	}
	if !verifierEcosystemValEHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValEProofSurfaceStateUnknown
	}
	if !containsExactTrimmedStringSet(model.ProofSurfaceRefs, VerifierEcosystemValEProofSurfaceRefs()...) {
		return VerifierEcosystemValEProofSurfaceStatePartial
	}
	if model.StaleEvidenceDetected || !model.EvidenceFresh {
		return VerifierEcosystemValEProofSurfaceStateBlocked
	}
	return VerifierEcosystemValEProofSurfaceStateActive
}

func EvaluateVerifierEcosystemValEEvidenceQualityState(model VerifierEcosystemIntegratedClosure) string {
	if len(model.EvidenceRefs) == 0 || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return VerifierEcosystemValEEvidenceQualityStateIncomplete
	}
	if !verifierEcosystemValEHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValEEvidenceQualityStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, VerifierEcosystemValEProofEvidenceRefs()...) {
		return VerifierEcosystemValEEvidenceQualityStatePartial
	}
	allFresh, stale, ok := verifierEcosystemVal0EvidenceValid(VerifierEcosystemValEVerifierEvidence())
	if !ok {
		return VerifierEcosystemValEEvidenceQualityStateUnknown
	}
	if !allFresh || stale || !model.EvidenceFresh || model.StaleEvidenceDetected {
		return VerifierEcosystemValEEvidenceQualityStateBlocked
	}
	evidenceIDs := make([]string, 0, len(VerifierEcosystemValEVerifierEvidence()))
	evidenceScopes := make([]string, 0, len(VerifierEcosystemValEVerifierEvidence()))
	for _, item := range VerifierEcosystemValEVerifierEvidence() {
		evidenceIDs = append(evidenceIDs, item.EvidenceID)
		evidenceScopes = append(evidenceScopes, item.Scope)
	}
	if !containsExactTrimmedStringSet(evidenceIDs, verifierEcosystemValERequiredEvidenceIDs()...) ||
		!containsExactTrimmedStringSet(evidenceScopes, verifierEcosystemValERequiredEvidenceScopes()...) {
		return VerifierEcosystemValEEvidenceQualityStateBlocked
	}
	return VerifierEcosystemValEEvidenceQualityStateActive
}

func EvaluateVerifierEcosystemValENoOverclaimState(model VerifierEcosystemIntegratedClosure) string {
	if strings.TrimSpace(model.ProjectionDisclaimer) == "" || strings.TrimSpace(model.Point7PassReason) == "" {
		return VerifierEcosystemValENoOverclaimStateIncomplete
	}
	if !verifierEcosystemValEHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValENoOverclaimStateUnknown
	}
	if model.Val0.Point7State == VerifierEcosystemPoint7StatePass ||
		model.ValA.Point7State == VerifierEcosystemPoint7StatePass ||
		model.ValB.Point7State == VerifierEcosystemPoint7StatePass ||
		model.ValC.Point7State == VerifierEcosystemPoint7StatePass ||
		model.ValD.Point7State == VerifierEcosystemPoint7StatePass {
		return VerifierEcosystemValENoOverclaimStateBlocked
	}
	if verifierEcosystemValEContainsDisallowedClaim(model.ObservedClaims...) {
		return VerifierEcosystemValENoOverclaimStateBlocked
	}
	if verifierEcosystemValEContainsDisallowedClaim(model.Point7PassReason) {
		return VerifierEcosystemValENoOverclaimStateBlocked
	}
	if verifierEcosystemValDStateSeverity(model.ValD.NoOverclaimGateState, VerifierEcosystemValDNoOverclaimGateStateActive, VerifierEcosystemValDNoOverclaimGateStatePartial, VerifierEcosystemValDNoOverclaimGateStateIncomplete, VerifierEcosystemValDNoOverclaimGateStateBlocked, VerifierEcosystemValDNoOverclaimGateStateUnknown) > 0 {
		return VerifierEcosystemValENoOverclaimStateBlocked
	}
	return VerifierEcosystemValENoOverclaimStateActive
}

func EvaluateVerifierEcosystemValEPassRuleState(model VerifierEcosystemIntegratedClosure) string {
	if strings.TrimSpace(model.Point7PassReason) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return VerifierEcosystemValEPassRuleStateIncomplete
	}
	reasonState := verifierEcosystemValEPoint7PassReasonState(model.Point7PassReason)
	if reasonState == verifierEcosystemValEPoint7PassReasonStateUnknown &&
		verifierEcosystemValEContainsDisallowedClaim(model.Point7PassReason) {
		return VerifierEcosystemValEPassRuleStateBlocked
	}
	if verifierEcosystemValEContainsDisallowedClaim(strings.Join(model.Caveats, " "), strings.Join(model.Limitations, " "), strings.Join(model.BlockingReasons, " ")) {
		return VerifierEcosystemValEPassRuleStateBlocked
	}
	prereqState := EvaluateVerifierEcosystemValEPrerequisiteState(model)
	invariantState := EvaluateVerifierEcosystemValEInvariantState(model)
	proofSurfaceState := EvaluateVerifierEcosystemValEProofSurfaceState(model)
	evidenceState := EvaluateVerifierEcosystemValEEvidenceQualityState(model)
	noOverclaimState := EvaluateVerifierEcosystemValENoOverclaimState(model)
	if prereqState == VerifierEcosystemValEPrerequisiteStateActive &&
		invariantState == VerifierEcosystemValEInvariantStateActive &&
		proofSurfaceState == VerifierEcosystemValEProofSurfaceStateActive &&
		evidenceState == VerifierEcosystemValEEvidenceQualityStateActive &&
		noOverclaimState == VerifierEcosystemValENoOverclaimStateActive &&
		model.RedactionKeepsFailuresVisible &&
		!model.MutatesCanonicalEvidence &&
		!model.ApprovesDeployment &&
		!model.SuppressesFailures {
		if verifierEcosystemValEPassAllowedClaim(model.Point7PassReason) {
			return VerifierEcosystemValEPassRuleStateActive
		}
		if reasonState == verifierEcosystemValEPoint7PassReasonStateBlocked {
			return VerifierEcosystemValEPassRuleStatePartial
		}
		return VerifierEcosystemValEPassRuleStateUnknown
	}
	if prereqState == VerifierEcosystemValEPrerequisiteStateBlocked ||
		invariantState == VerifierEcosystemValEInvariantStateBlocked ||
		proofSurfaceState == VerifierEcosystemValEProofSurfaceStateBlocked ||
		evidenceState == VerifierEcosystemValEEvidenceQualityStateBlocked ||
		noOverclaimState == VerifierEcosystemValENoOverclaimStateBlocked ||
		!model.RedactionKeepsFailuresVisible ||
		model.MutatesCanonicalEvidence ||
		model.ApprovesDeployment ||
		model.SuppressesFailures {
		return VerifierEcosystemValEPassRuleStateBlocked
	}
	if prereqState == VerifierEcosystemValEPrerequisiteStateIncomplete ||
		invariantState == VerifierEcosystemValEInvariantStateIncomplete ||
		proofSurfaceState == VerifierEcosystemValEProofSurfaceStateIncomplete ||
		evidenceState == VerifierEcosystemValEEvidenceQualityStateIncomplete ||
		noOverclaimState == VerifierEcosystemValENoOverclaimStateIncomplete {
		return VerifierEcosystemValEPassRuleStateIncomplete
	}
	if prereqState == VerifierEcosystemValEPrerequisiteStateUnknown ||
		invariantState == VerifierEcosystemValEInvariantStateUnknown ||
		proofSurfaceState == VerifierEcosystemValEProofSurfaceStateUnknown ||
		evidenceState == VerifierEcosystemValEEvidenceQualityStateUnknown ||
		noOverclaimState == VerifierEcosystemValENoOverclaimStateUnknown {
		return VerifierEcosystemValEPassRuleStateUnknown
	}
	if reasonState == verifierEcosystemValEPoint7PassReasonStateUnknown {
		return VerifierEcosystemValEPassRuleStateUnknown
	}
	return VerifierEcosystemValEPassRuleStatePartial
}

func EvaluateVerifierEcosystemValEPoint7State(model VerifierEcosystemIntegratedClosure) string {
	if EvaluateVerifierEcosystemValEPassRuleState(model) == VerifierEcosystemValEPassRuleStateActive {
		return VerifierEcosystemPoint7StatePass
	}
	return VerifierEcosystemPoint7StateNotComplete
}

func EvaluateVerifierEcosystemValEState(model VerifierEcosystemIntegratedClosure) string {
	passRuleState := EvaluateVerifierEcosystemValEPassRuleState(model)
	if passRuleState == VerifierEcosystemValEPassRuleStateActive && EvaluateVerifierEcosystemValEPoint7State(model) == VerifierEcosystemPoint7StatePass {
		return VerifierEcosystemValEStatePass
	}
	switch passRuleState {
	case VerifierEcosystemValEPassRuleStateBlocked:
		return VerifierEcosystemValEStateBlocked
	case VerifierEcosystemValEPassRuleStateIncomplete:
		return VerifierEcosystemValEStateIncomplete
	case VerifierEcosystemValEPassRuleStateUnknown:
		return VerifierEcosystemValEStateUnknown
	default:
		return VerifierEcosystemValEStatePartial
	}
}

func verifierEcosystemValECollectText(values []string) []string {
	collected := make([]string, 0, len(values))
	seen := make(map[string]struct{}, len(values))
	for _, value := range values {
		normalized := strings.TrimSpace(value)
		if normalized == "" {
			continue
		}
		if _, ok := seen[normalized]; ok {
			continue
		}
		seen[normalized] = struct{}{}
		collected = append(collected, normalized)
	}
	return collected
}

func verifierEcosystemValEBlockingReasons(model VerifierEcosystemIntegratedClosure, invariants []VerifierEcosystemClosureInvariant) []string {
	reasons := []string{}
	if prereqState := EvaluateVerifierEcosystemValEPrerequisiteState(model); prereqState != VerifierEcosystemValEPrerequisiteStateActive {
		reasons = append(reasons, "Integrated closure prerequisites are not fully active and healthy.")
	}
	if proofSurfaceState := EvaluateVerifierEcosystemValEProofSurfaceState(model); proofSurfaceState != VerifierEcosystemValEProofSurfaceStateActive {
		reasons = append(reasons, "Integrated closure proof surfaces are not exact or evidence freshness is degraded.")
	}
	if evidenceState := EvaluateVerifierEcosystemValEEvidenceQualityState(model); evidenceState != VerifierEcosystemValEEvidenceQualityStateActive {
		reasons = append(reasons, "Integrated closure evidence refs or evidence quality are not exact and fresh.")
	}
	if noOverclaimState := EvaluateVerifierEcosystemValENoOverclaimState(model); noOverclaimState != VerifierEcosystemValENoOverclaimStateActive {
		reasons = append(reasons, "Integrated closure no-overclaim discipline is not fully active.")
	}
	if passRuleState := EvaluateVerifierEcosystemValEPassRuleState(model); passRuleState != VerifierEcosystemValEPassRuleStateActive {
		reasons = append(reasons, "Val E pass rule remains fail-closed until all prerequisites, invariants, surfaces, evidence, and no-overclaim checks are active.")
	}
	for _, invariant := range invariants {
		if invariant.CurrentState != VerifierEcosystemValEInvariantStateActive {
			reasons = append(reasons, invariant.BlockingReasons...)
		}
	}
	return verifierEcosystemValECollectText(reasons)
}

func ComputeVerifierEcosystemValEClosure(model VerifierEcosystemIntegratedClosure) VerifierEcosystemIntegratedClosure {
	invariants := VerifierEcosystemValEClosureInvariants(model)
	model.ClosureInvariants = invariants
	model.ClosurePrerequisiteState = EvaluateVerifierEcosystemValEPrerequisiteState(model)
	model.ClosureInvariantState = EvaluateVerifierEcosystemValEInvariantState(model)
	model.ProofSurfaceState = EvaluateVerifierEcosystemValEProofSurfaceState(model)
	model.EvidenceQualityState = EvaluateVerifierEcosystemValEEvidenceQualityState(model)
	model.NoOverclaimState = EvaluateVerifierEcosystemValENoOverclaimState(model)
	model.PassRuleState = EvaluateVerifierEcosystemValEPassRuleState(model)
	model.Point7State = EvaluateVerifierEcosystemValEPoint7State(model)
	model.Point7PassAllowed = model.Point7State == VerifierEcosystemPoint7StatePass
	model.CurrentState = EvaluateVerifierEcosystemValEState(model)
	model.BlockingReasons = verifierEcosystemValEBlockingReasons(model, invariants)
	return model
}
