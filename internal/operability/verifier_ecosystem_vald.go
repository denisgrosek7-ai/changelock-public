package operability

import "strings"

const (
	VerifierEcosystemValDCorrectnessGateStateActive     = "verifier_ecosystem_vald_correctness_gate_active"
	VerifierEcosystemValDCorrectnessGateStatePartial    = "verifier_ecosystem_vald_correctness_gate_partial"
	VerifierEcosystemValDCorrectnessGateStateIncomplete = "verifier_ecosystem_vald_correctness_gate_incomplete"
	VerifierEcosystemValDCorrectnessGateStateBlocked    = "verifier_ecosystem_vald_correctness_gate_blocked"
	VerifierEcosystemValDCorrectnessGateStateUnknown    = "verifier_ecosystem_vald_correctness_gate_unknown"

	VerifierEcosystemValDToolingGateStateActive     = "verifier_ecosystem_vald_tooling_gate_active"
	VerifierEcosystemValDToolingGateStatePartial    = "verifier_ecosystem_vald_tooling_gate_partial"
	VerifierEcosystemValDToolingGateStateIncomplete = "verifier_ecosystem_vald_tooling_gate_incomplete"
	VerifierEcosystemValDToolingGateStateBlocked    = "verifier_ecosystem_vald_tooling_gate_blocked"
	VerifierEcosystemValDToolingGateStateUnknown    = "verifier_ecosystem_vald_tooling_gate_unknown"

	VerifierEcosystemValDSchemaCompatibilityGateStateActive     = "verifier_ecosystem_vald_schema_compatibility_gate_active"
	VerifierEcosystemValDSchemaCompatibilityGateStatePartial    = "verifier_ecosystem_vald_schema_compatibility_gate_partial"
	VerifierEcosystemValDSchemaCompatibilityGateStateIncomplete = "verifier_ecosystem_vald_schema_compatibility_gate_incomplete"
	VerifierEcosystemValDSchemaCompatibilityGateStateBlocked    = "verifier_ecosystem_vald_schema_compatibility_gate_blocked"
	VerifierEcosystemValDSchemaCompatibilityGateStateUnknown    = "verifier_ecosystem_vald_schema_compatibility_gate_unknown"

	VerifierEcosystemValDDiagnosticsConformanceGateStateActive     = "verifier_ecosystem_vald_diagnostics_conformance_gate_active"
	VerifierEcosystemValDDiagnosticsConformanceGateStatePartial    = "verifier_ecosystem_vald_diagnostics_conformance_gate_partial"
	VerifierEcosystemValDDiagnosticsConformanceGateStateIncomplete = "verifier_ecosystem_vald_diagnostics_conformance_gate_incomplete"
	VerifierEcosystemValDDiagnosticsConformanceGateStateBlocked    = "verifier_ecosystem_vald_diagnostics_conformance_gate_blocked"
	VerifierEcosystemValDDiagnosticsConformanceGateStateUnknown    = "verifier_ecosystem_vald_diagnostics_conformance_gate_unknown"

	VerifierEcosystemValDTrustKeyRotationGateStateActive     = "verifier_ecosystem_vald_trust_key_rotation_gate_active"
	VerifierEcosystemValDTrustKeyRotationGateStatePartial    = "verifier_ecosystem_vald_trust_key_rotation_gate_partial"
	VerifierEcosystemValDTrustKeyRotationGateStateIncomplete = "verifier_ecosystem_vald_trust_key_rotation_gate_incomplete"
	VerifierEcosystemValDTrustKeyRotationGateStateBlocked    = "verifier_ecosystem_vald_trust_key_rotation_gate_blocked"
	VerifierEcosystemValDTrustKeyRotationGateStateUnknown    = "verifier_ecosystem_vald_trust_key_rotation_gate_unknown"

	VerifierEcosystemValDNegativeDiagnosticsGateStateActive     = "verifier_ecosystem_vald_negative_diagnostics_gate_active"
	VerifierEcosystemValDNegativeDiagnosticsGateStatePartial    = "verifier_ecosystem_vald_negative_diagnostics_gate_partial"
	VerifierEcosystemValDNegativeDiagnosticsGateStateIncomplete = "verifier_ecosystem_vald_negative_diagnostics_gate_incomplete"
	VerifierEcosystemValDNegativeDiagnosticsGateStateBlocked    = "verifier_ecosystem_vald_negative_diagnostics_gate_blocked"
	VerifierEcosystemValDNegativeDiagnosticsGateStateUnknown    = "verifier_ecosystem_vald_negative_diagnostics_gate_unknown"

	VerifierEcosystemValDRedactionGateStateActive     = "verifier_ecosystem_vald_redaction_gate_active"
	VerifierEcosystemValDRedactionGateStatePartial    = "verifier_ecosystem_vald_redaction_gate_partial"
	VerifierEcosystemValDRedactionGateStateIncomplete = "verifier_ecosystem_vald_redaction_gate_incomplete"
	VerifierEcosystemValDRedactionGateStateBlocked    = "verifier_ecosystem_vald_redaction_gate_blocked"
	VerifierEcosystemValDRedactionGateStateUnknown    = "verifier_ecosystem_vald_redaction_gate_unknown"

	VerifierEcosystemValDPublisherArtifactGateStateActive     = "verifier_ecosystem_vald_publisher_artifact_gate_active"
	VerifierEcosystemValDPublisherArtifactGateStatePartial    = "verifier_ecosystem_vald_publisher_artifact_gate_partial"
	VerifierEcosystemValDPublisherArtifactGateStateIncomplete = "verifier_ecosystem_vald_publisher_artifact_gate_incomplete"
	VerifierEcosystemValDPublisherArtifactGateStateBlocked    = "verifier_ecosystem_vald_publisher_artifact_gate_blocked"
	VerifierEcosystemValDPublisherArtifactGateStateUnknown    = "verifier_ecosystem_vald_publisher_artifact_gate_unknown"

	VerifierEcosystemValDNoOverclaimGateStateActive     = "verifier_ecosystem_vald_no_overclaim_gate_active"
	VerifierEcosystemValDNoOverclaimGateStatePartial    = "verifier_ecosystem_vald_no_overclaim_gate_partial"
	VerifierEcosystemValDNoOverclaimGateStateIncomplete = "verifier_ecosystem_vald_no_overclaim_gate_incomplete"
	VerifierEcosystemValDNoOverclaimGateStateBlocked    = "verifier_ecosystem_vald_no_overclaim_gate_blocked"
	VerifierEcosystemValDNoOverclaimGateStateUnknown    = "verifier_ecosystem_vald_no_overclaim_gate_unknown"

	VerifierEcosystemValDStateActive     = "verifier_ecosystem_vald_active"
	VerifierEcosystemValDStatePartial    = "verifier_ecosystem_vald_partial"
	VerifierEcosystemValDStateIncomplete = "verifier_ecosystem_vald_incomplete"
	VerifierEcosystemValDStateBlocked    = "verifier_ecosystem_vald_blocked"
	VerifierEcosystemValDStateUnknown    = "verifier_ecosystem_vald_unknown"

	verifierEcosystemValDIssuerStateBound   = "issuer_bound"
	verifierEcosystemValDIssuerStateWarning = "issuer_warning"
	verifierEcosystemValDIssuerStateRevoked = "issuer_revoked"
)

type VerifierEcosystemValDDependencySnapshot struct {
	Point5State                    string `json:"point_5_state"`
	Point5DependencyState          string `json:"point_5_dependency_state"`
	Point6State                    string `json:"point_6_state"`
	Point6ClosureState             string `json:"point_6_closure_state"`
	Point6ClosurePrerequisiteState string `json:"point_6_closure_prerequisite_state"`
	Point6ClosureInvariantState    string `json:"point_6_closure_invariant_state"`
	Point6ProofSurfaceState        string `json:"point_6_proof_surface_state"`
	Point6PassRuleState            string `json:"point_6_pass_rule_state"`
	Point6PassAllowed              bool   `json:"point_6_pass_allowed"`
	Val0CurrentState               string `json:"val_0_current_state"`
	Val0State                      string `json:"val_0_state"`
	ValACurrentState               string `json:"val_a_current_state"`
	ValAState                      string `json:"val_a_state"`
	ValBCurrentState               string `json:"val_b_current_state"`
	ValBState                      string `json:"val_b_state"`
	ValCCurrentState               string `json:"val_c_current_state"`
	ValCState                      string `json:"val_c_state"`
	Point7State                    string `json:"point_7_state"`
}

type VerifierEcosystemValDCorrectnessGate struct {
	CurrentState             string   `json:"current_state"`
	CorrectnessGateID        string   `json:"correctness_gate_id"`
	Version                  string   `json:"version"`
	SourceValStates          []string `json:"source_val_states,omitempty"`
	VerifierContractState    string   `json:"verifier_contract_state"`
	ReferenceEngineState     string   `json:"reference_engine_state"`
	CompatibilityMatrixState string   `json:"compatibility_matrix_state"`
	DiagnosticsState         string   `json:"diagnostics_state"`
	ConformanceSuiteState    string   `json:"conformance_suite_state"`
	EcosystemSurfaceState    string   `json:"ecosystem_surface_state"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	BlockingReasons          []string `json:"blocking_reasons,omitempty"`
	Caveats                  []string `json:"caveats,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
	LifecycleState           string   `json:"lifecycle_state"`
	CompatibilityState       string   `json:"compatibility_state"`
	CreatedAt                string   `json:"created_at"`
	UpdatedAt                string   `json:"updated_at"`
	CertificationClaim       bool     `json:"certification_claim"`
	ApprovalClaim            bool     `json:"approval_claim"`
}

type VerifierEcosystemValDToolingGate struct {
	CurrentState                     string   `json:"current_state"`
	ToolingGateID                    string   `json:"tooling_gate_id"`
	VerifierInputModelState          string   `json:"verifier_input_model_state"`
	ReferenceVerifierEngineState     string   `json:"reference_verifier_engine_state"`
	VerificationResultModelState     string   `json:"verification_result_model_state"`
	DiagnosticsMappingState          string   `json:"diagnostics_mapping_state"`
	CommandSurfaceState              string   `json:"command_surface_state"`
	SDKEntrypointState               string   `json:"sdk_entrypoint_state"`
	DeterministicOutput              bool     `json:"deterministic_output"`
	HiddenMainInstanceDependency     bool     `json:"hidden_main_instance_dependency"`
	NetworkDependency                bool     `json:"network_dependency"`
	MutatesEvidence                  bool     `json:"mutates_evidence"`
	ApprovesDeployment               bool     `json:"approves_deployment"`
	SuppressesFailures               bool     `json:"suppresses_failures"`
	CertificationClaim               bool     `json:"certification_claim"`
	ClaimsRealCryptoWithoutPrimitive bool     `json:"claims_real_crypto_without_primitive"`
	Caveats                          []string `json:"caveats,omitempty"`
	ProjectionDisclaimer             string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValDSchemaCompatibilityGate struct {
	CurrentState                  string   `json:"current_state"`
	SchemaCompatibilityGateID     string   `json:"schema_compatibility_gate_id"`
	CompatibilityMatrixState      string   `json:"compatibility_matrix_state"`
	SchemaProofCompatibilityState string   `json:"schema_proof_compatibility_state"`
	MixedVersionDiagnosticsState  string   `json:"mixed_version_diagnostics_state"`
	SupportedSchemaVersions       []string `json:"supported_schema_versions,omitempty"`
	SupportedProofTypes           []string `json:"supported_proof_types,omitempty"`
	SupportedVerifierVersions     []string `json:"supported_verifier_versions,omitempty"`
	SupportedTrustRootVersions    []string `json:"supported_trust_root_versions,omitempty"`
	CompatibilityEntryKeys        []string `json:"compatibility_entry_keys,omitempty"`
	MixedVersionRuleCoverage      []string `json:"mixed_version_rule_coverage,omitempty"`
	CompatibilityState            string   `json:"compatibility_state"`
	DeprecatedMigrationVisible    bool     `json:"deprecated_migration_visible"`
	SupersessionVisible           bool     `json:"supersession_visible"`
	EvidenceRefs                  []string `json:"evidence_refs,omitempty"`
	Caveats                       []string `json:"caveats,omitempty"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValDDiagnosticsConformanceGate struct {
	CurrentState                 string   `json:"current_state"`
	DiagnosticsConformanceGateID string   `json:"diagnostics_conformance_gate_id"`
	DiagnosticPrecedenceState    string   `json:"diagnostic_precedence_state"`
	FixtureDescriptorState       string   `json:"fixture_descriptor_state"`
	ConformanceCaseState         string   `json:"conformance_case_state"`
	ConformanceSuiteState        string   `json:"conformance_suite_state"`
	OutputClassState             string   `json:"output_class_state"`
	ObservedDiagnostics          []string `json:"observed_diagnostics,omitempty"`
	DerivedDiagnosticClass       string   `json:"derived_diagnostic_class"`
	StaleCoverageVisible         bool     `json:"stale_coverage_visible"`
	RevokedCoverageVisible       bool     `json:"revoked_coverage_visible"`
	SupersededCoverageVisible    bool     `json:"superseded_coverage_visible"`
	UnsupportedCoverageVisible   bool     `json:"unsupported_coverage_visible"`
	MalformedCoverageVisible     bool     `json:"malformed_coverage_visible"`
	CertificationClaim           bool     `json:"certification_claim"`
	IntegrityRatingClaim         bool     `json:"integrity_rating_claim"`
	Caveats                      []string `json:"caveats,omitempty"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValDTrustKeyRotationGate struct {
	CurrentState               string   `json:"current_state"`
	TrustKeyRotationGateID     string   `json:"trust_key_rotation_gate_id"`
	TrustState                 string   `json:"trust_state"`
	TrustRootState             string   `json:"trust_root_state"`
	IssuerState                string   `json:"issuer_state"`
	RevocationState            string   `json:"revocation_state"`
	KeyRotationState           string   `json:"key_rotation_state"`
	RolloverMetadataRef        string   `json:"rollover_metadata_ref"`
	TrustDistributionState     string   `json:"trust_distribution_state"`
	TrustDistributionMode      string   `json:"trust_distribution_mode"`
	OfflineDistributionScope   string   `json:"offline_distribution_scope"`
	EvidenceFreshnessState     string   `json:"evidence_freshness_state"`
	EvidenceRefs               []string `json:"evidence_refs,omitempty"`
	Caveats                    []string `json:"caveats,omitempty"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
	GlobalKeyDirectoryClaim    bool     `json:"global_key_directory_claim"`
	SensitivePublicKeyExposure bool     `json:"sensitive_public_key_exposure"`
}

type VerifierEcosystemValDNegativeDiagnosticsGate struct {
	CurrentState                      string   `json:"current_state"`
	NegativeDiagnosticsGateID         string   `json:"negative_diagnostics_gate_id"`
	PublicOutputState                 string   `json:"public_output_state"`
	PartnerOutputState                string   `json:"partner_output_state"`
	AuditorFlowState                  string   `json:"auditor_flow_state"`
	PublicOverallResult               string   `json:"public_overall_result"`
	PublicDiagnosticClass             string   `json:"public_diagnostic_class"`
	PublicOutputClass                 string   `json:"public_output_class"`
	PartnerOverallResult              string   `json:"partner_overall_result"`
	PartnerDiagnosticClass            string   `json:"partner_diagnostic_class"`
	PartnerOutputClass                string   `json:"partner_output_class"`
	StaleArtifactVisible              bool     `json:"stale_artifact_visible"`
	ExpiredArtifactVisible            bool     `json:"expired_artifact_visible"`
	RevokedIssuerVisible              bool     `json:"revoked_issuer_visible"`
	SupersededProofVisible            bool     `json:"superseded_proof_visible"`
	UnsupportedSchemaVisible          bool     `json:"unsupported_schema_visible"`
	UnsupportedProofTypeVisible       bool     `json:"unsupported_proof_type_visible"`
	InsufficientTrustVisible          bool     `json:"insufficient_trust_visible"`
	RedactionBoundaryPreserved        bool     `json:"redaction_boundary_preserved"`
	PublicPreservesNonVerified        bool     `json:"public_preserves_non_verified"`
	PartnerInternalDiagnosticsExposed bool     `json:"partner_internal_diagnostics_exposed"`
	AuditorRepeatable                 bool     `json:"auditor_repeatable"`
	AuditorEvidenceLinked             bool     `json:"auditor_evidence_linked"`
	Caveats                           []string `json:"caveats,omitempty"`
	ProjectionDisclaimer              string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValDRedactionGate struct {
	CurrentState                string   `json:"current_state"`
	RedactionGateID             string   `json:"redaction_gate_id"`
	AudienceSurfaceState        string   `json:"audience_surface_state"`
	PublicOutputState           string   `json:"public_output_state"`
	PartnerOutputState          string   `json:"partner_output_state"`
	AuditorFlowState            string   `json:"auditor_flow_state"`
	OutputBoundaryState         string   `json:"output_boundary_state"`
	RedactionPolicyRef          string   `json:"redaction_policy_ref"`
	EvidenceVisibilityPolicy    string   `json:"evidence_visibility_policy"`
	TrustMaterialVisibility     string   `json:"trust_material_visibility"`
	InternalDiagnosticSeparated bool     `json:"internal_diagnostic_separated"`
	PartnerBroaderThanPublic    bool     `json:"partner_broader_than_public"`
	Caveats                     []string `json:"caveats,omitempty"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValDPublisherArtifactGate struct {
	CurrentState              string   `json:"current_state"`
	PublisherArtifactGateID   string   `json:"publisher_artifact_gate_id"`
	PublisherProfileState     string   `json:"publisher_profile_state"`
	ArtifactRuleState         string   `json:"artifact_rule_state"`
	PublisherType             string   `json:"publisher_type"`
	RequiredSchemaPolicy      string   `json:"required_schema_policy"`
	RequiredSignaturePolicy   string   `json:"required_signature_policy"`
	RequiredDigestPolicy      string   `json:"required_digest_policy"`
	TrustRootPolicy           string   `json:"trust_root_policy"`
	SupportedArtifactTypes    []string `json:"supported_artifact_types,omitempty"`
	OutputBoundaryCompatible  bool     `json:"output_boundary_compatible"`
	ConformanceCaseRefs       []string `json:"conformance_case_refs,omitempty"`
	ObservedClaims            []string `json:"observed_claims,omitempty"`
	Caveats                   []string `json:"caveats,omitempty"`
	ProjectionDisclaimer      string   `json:"projection_disclaimer"`
	AutomaticallyTrustedClaim bool     `json:"automatically_trusted_claim"`
}

type VerifierEcosystemValDNoOverclaimGate struct {
	CurrentState         string   `json:"current_state"`
	NoOverclaimGateID    string   `json:"no_overclaim_gate_id"`
	OutputRefs           []string `json:"output_refs,omitempty"`
	ObservedClaims       []string `json:"observed_claims,omitempty"`
	Caveats              []string `json:"caveats,omitempty"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

func verifierEcosystemValDProjectionDisclaimer() string {
	return "projection_only not_canonical_truth final_verifier_ecosystem_gate advisory_projection"
}

func verifierEcosystemValDRequiredSourceValStates() []string {
	return []string{
		VerifierEcosystemVal0StateActive,
		VerifierEcosystemValAStateActive,
		VerifierEcosystemValBStateActive,
		VerifierEcosystemValCStateActive,
	}
}

func verifierEcosystemValDIssuerStates() []string {
	return []string{
		verifierEcosystemValDIssuerStateBound,
		verifierEcosystemValDIssuerStateWarning,
		verifierEcosystemValDIssuerStateRevoked,
	}
}

func verifierEcosystemValDRequiredNoOverclaimRefs() []string {
	return []string{
		"vald-output:correctness",
		"vald-output:tooling",
		"vald-output:schema-compatibility",
		"vald-output:diagnostics-conformance",
		"vald-output:trust-key-rotation",
		"vald-output:negative-diagnostics",
		"vald-output:redaction",
		"vald-output:publisher-artifact",
		"vald-output:proofs",
	}
}

func verifierEcosystemValDDisallowedClaims() []string {
	return []string{
		"verifier certification",
		"certified verifier",
		"certified publisher",
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
		"point_7_pass",
	}
}

func verifierEcosystemValDContainsDisallowedClaim(values ...string) bool {
	for _, value := range values {
		normalized := strings.ToLower(strings.TrimSpace(value))
		for _, claim := range verifierEcosystemValDDisallowedClaims() {
			if strings.Contains(normalized, claim) {
				return true
			}
		}
	}
	return false
}

func verifierEcosystemValDClaimsBlocked(values []string) bool {
	for _, value := range values {
		if verifierEcosystemValDContainsDisallowedClaim(value) {
			return true
		}
	}
	return false
}

func VerifierEcosystemValDCorrectnessGateModel() VerifierEcosystemValDCorrectnessGate {
	return VerifierEcosystemValDCorrectnessGate{
		CurrentState:             "verifier_ecosystem_vald_correctness_gate_ready",
		CorrectnessGateID:        "verifier-correctness-gate-vald",
		Version:                  "2026.04",
		SourceValStates:          verifierEcosystemValDRequiredSourceValStates(),
		VerifierContractState:    VerifierEcosystemVal0ContractStateActive,
		ReferenceEngineState:     VerifierEcosystemValAEngineStateActive,
		CompatibilityMatrixState: VerifierEcosystemValBCompatibilityMatrixStateActive,
		DiagnosticsState:         VerifierEcosystemValBDiagnosticPrecedenceStateActive,
		ConformanceSuiteState:    VerifierEcosystemValBConformanceSuiteStateActive,
		EcosystemSurfaceState:    VerifierEcosystemValCStateActive,
		EvidenceRefs:             []string{"evidence:correctness-gate-001", "evidence:point7-governance-004"},
		Caveats:                  []string{"final verifier ecosystem gate remains advisory and bounded to exact Val 0, Val A, Val B, and Val C states"},
		Limitations:              []string{"Val D does not close Točka 7 and does not create point_7_pass"},
		ProjectionDisclaimer:     verifierEcosystemValDProjectionDisclaimer(),
		LifecycleState:           "active",
		CompatibilityState:       ReferenceArchitectureCompatibilityCompatible,
		CreatedAt:                "2026-04-27T20:20:00Z",
		UpdatedAt:                "2026-04-27T20:20:00Z",
	}
}

func VerifierEcosystemValDToolingGateModel() VerifierEcosystemValDToolingGate {
	return VerifierEcosystemValDToolingGate{
		CurrentState:                 "verifier_ecosystem_vald_tooling_gate_ready",
		ToolingGateID:                "reference-verifier-tooling-gate-vald",
		VerifierInputModelState:      VerifierEcosystemValAInputStateActive,
		ReferenceVerifierEngineState: VerifierEcosystemValAEngineStateActive,
		VerificationResultModelState: VerifierEcosystemValAResultStateActive,
		DiagnosticsMappingState:      VerifierEcosystemValADiagnosticsMappingStateActive,
		CommandSurfaceState:          VerifierEcosystemValACommandContractStateActive,
		SDKEntrypointState:           VerifierEcosystemValASDKEntrypointStateActive,
		DeterministicOutput:          true,
		Caveats:                      []string{"reference verifier tooling gate remains deterministic and advisory only"},
		ProjectionDisclaimer:         verifierEcosystemValDProjectionDisclaimer(),
	}
}

func VerifierEcosystemValDSchemaCompatibilityGateModel() VerifierEcosystemValDSchemaCompatibilityGate {
	matrix := VerifierEcosystemValBCompatibilityMatrixModel()
	entryKeys := make([]string, 0, len(matrix.CompatibilityEntries))
	for _, entry := range matrix.CompatibilityEntries {
		entryKeys = append(entryKeys, verifierEcosystemValBCompatibilityEntryKey(entry))
	}
	return VerifierEcosystemValDSchemaCompatibilityGate{
		CurrentState:                  "verifier_ecosystem_vald_schema_compatibility_gate_ready",
		SchemaCompatibilityGateID:     "schema-compatibility-gate-vald",
		CompatibilityMatrixState:      VerifierEcosystemValBCompatibilityMatrixStateActive,
		SchemaProofCompatibilityState: VerifierEcosystemValBSchemaProofCompatibilityStateActive,
		MixedVersionDiagnosticsState:  VerifierEcosystemValBMixedVersionStateActive,
		SupportedSchemaVersions:       matrix.SupportedSchemaVersions,
		SupportedProofTypes:           matrix.SupportedProofTypes,
		SupportedVerifierVersions:     matrix.SupportedVerifierVersions,
		SupportedTrustRootVersions:    matrix.SupportedTrustRootVersions,
		CompatibilityEntryKeys:        entryKeys,
		MixedVersionRuleCoverage:      matrix.MixedVersionRules,
		CompatibilityState:            ReferenceArchitectureCompatibilityCompatible,
		DeprecatedMigrationVisible:    true,
		SupersessionVisible:           true,
		EvidenceRefs:                  []string{"evidence:schema-compatibility-gate-001"},
		Caveats:                       []string{"compatibility gate remains version-bound and warning-bearing where migration or supersession visibility is required"},
		ProjectionDisclaimer:          verifierEcosystemValDProjectionDisclaimer(),
	}
}

func VerifierEcosystemValDDiagnosticsConformanceGateModel() VerifierEcosystemValDDiagnosticsConformanceGate {
	return VerifierEcosystemValDDiagnosticsConformanceGate{
		CurrentState:                 "verifier_ecosystem_vald_diagnostics_conformance_gate_ready",
		DiagnosticsConformanceGateID: "diagnostics-conformance-gate-vald",
		DiagnosticPrecedenceState:    VerifierEcosystemValBDiagnosticPrecedenceStateActive,
		FixtureDescriptorState:       VerifierEcosystemValBFixtureDescriptorStateActive,
		ConformanceCaseState:         VerifierEcosystemValBConformanceCaseStateActive,
		ConformanceSuiteState:        VerifierEcosystemValBConformanceSuiteStateActive,
		OutputClassState:             VerifierEcosystemValBOutputClassStateActive,
		ObservedDiagnostics:          []string{VerifierEcosystemDiagnosticVerified},
		DerivedDiagnosticClass:       VerifierEcosystemDiagnosticVerified,
		StaleCoverageVisible:         true,
		RevokedCoverageVisible:       true,
		SupersededCoverageVisible:    true,
		UnsupportedCoverageVisible:   true,
		MalformedCoverageVisible:     true,
		Caveats:                      []string{"deterministic diagnostics and conformance remain bounded verifier behavior checks only"},
		ProjectionDisclaimer:         verifierEcosystemValDProjectionDisclaimer(),
	}
}

func VerifierEcosystemValDTrustKeyRotationGateModel() VerifierEcosystemValDTrustKeyRotationGate {
	return VerifierEcosystemValDTrustKeyRotationGate{
		CurrentState:             "verifier_ecosystem_vald_trust_key_rotation_gate_ready",
		TrustKeyRotationGateID:   "trust-key-rotation-gate-vald",
		TrustState:               VerifierEcosystemVal0TrustStateActive,
		TrustRootState:           VerifierEcosystemTrustRootTrusted,
		IssuerState:              verifierEcosystemValDIssuerStateBound,
		RevocationState:          VerifierEcosystemRevocationNotRevoked,
		KeyRotationState:         VerifierEcosystemKeyRotationCurrent,
		RolloverMetadataRef:      "rollover:trust-root-2026.04",
		TrustDistributionState:   VerifierEcosystemValCTrustDistributionStateActive,
		TrustDistributionMode:    VerifierEcosystemValCDistributionModePartnerScopedDir,
		OfflineDistributionScope: VerifierEcosystemScopePartnerSafe,
		EvidenceFreshnessState:   IntelligenceCalibrationFreshnessFresh,
		EvidenceRefs:             []string{"evidence:trust-key-rotation-gate-001"},
		Caveats:                  []string{"trust-root and key-rotation gate remains scoped, bounded, and non-global"},
		ProjectionDisclaimer:     verifierEcosystemValDProjectionDisclaimer(),
	}
}

func VerifierEcosystemValDNegativeDiagnosticsGateModel() VerifierEcosystemValDNegativeDiagnosticsGate {
	return VerifierEcosystemValDNegativeDiagnosticsGate{
		CurrentState:                "verifier_ecosystem_vald_negative_diagnostics_gate_ready",
		NegativeDiagnosticsGateID:   "negative-diagnostics-gate-vald",
		PublicOutputState:           VerifierEcosystemValCPublicOutputStateActive,
		PartnerOutputState:          VerifierEcosystemValCPartnerOutputStateActive,
		AuditorFlowState:            VerifierEcosystemValCAuditorFlowStateActive,
		PublicOverallResult:         VerifierEcosystemValAOverallResultVerified,
		PublicDiagnosticClass:       VerifierEcosystemDiagnosticVerified,
		PublicOutputClass:           VerifierEcosystemValBOutputClassVerified,
		PartnerOverallResult:        VerifierEcosystemValAOverallResultVerified,
		PartnerDiagnosticClass:      VerifierEcosystemDiagnosticVerified,
		PartnerOutputClass:          VerifierEcosystemValBOutputClassVerified,
		StaleArtifactVisible:        true,
		ExpiredArtifactVisible:      true,
		RevokedIssuerVisible:        true,
		SupersededProofVisible:      true,
		UnsupportedSchemaVisible:    true,
		UnsupportedProofTypeVisible: true,
		InsufficientTrustVisible:    true,
		RedactionBoundaryPreserved:  true,
		PublicPreservesNonVerified:  true,
		AuditorRepeatable:           true,
		AuditorEvidenceLinked:       true,
		Caveats:                     []string{"negative diagnostics remain visible as non-verified outcomes across bounded outputs"},
		ProjectionDisclaimer:        verifierEcosystemValDProjectionDisclaimer(),
	}
}

func VerifierEcosystemValDRedactionGateModel() VerifierEcosystemValDRedactionGate {
	return VerifierEcosystemValDRedactionGate{
		CurrentState:                "verifier_ecosystem_vald_redaction_gate_ready",
		RedactionGateID:             "redaction-gate-vald",
		AudienceSurfaceState:        VerifierEcosystemValCAudienceSurfaceStateActive,
		PublicOutputState:           VerifierEcosystemValCPublicOutputStateActive,
		PartnerOutputState:          VerifierEcosystemValCPartnerOutputStateActive,
		AuditorFlowState:            VerifierEcosystemValCAuditorFlowStateActive,
		OutputBoundaryState:         VerifierEcosystemVal0OutputBoundaryStateActive,
		RedactionPolicyRef:          "redaction-policy:public-safe",
		EvidenceVisibilityPolicy:    "public_caveated",
		TrustMaterialVisibility:     "summary_only",
		InternalDiagnosticSeparated: true,
		PartnerBroaderThanPublic:    true,
		Caveats:                     []string{"redaction gate preserves bounded audience separation and cannot convert invalid into verified"},
		ProjectionDisclaimer:        verifierEcosystemValDProjectionDisclaimer(),
	}
}

func VerifierEcosystemValDPublisherArtifactGateModel() VerifierEcosystemValDPublisherArtifactGate {
	return VerifierEcosystemValDPublisherArtifactGate{
		CurrentState:             "verifier_ecosystem_vald_publisher_artifact_gate_ready",
		PublisherArtifactGateID:  "publisher-artifact-gate-vald",
		PublisherProfileState:    VerifierEcosystemValCPublisherProfileStateActive,
		ArtifactRuleState:        VerifierEcosystemValCArtifactRuleStateActive,
		PublisherType:            VerifierEcosystemValCPublisherTypeVendor,
		RequiredSchemaPolicy:     "versioned_schema_required",
		RequiredSignaturePolicy:  "signature_ref_required",
		RequiredDigestPolicy:     "sha256_or_sha512_required",
		TrustRootPolicy:          "versioned_trust_root_required",
		SupportedArtifactTypes:   verifierEcosystemVal0SupportedProofTypes(),
		OutputBoundaryCompatible: true,
		ConformanceCaseRefs:      verifierEcosystemValBRequiredConformanceCaseIDs(),
		Caveats:                  []string{"publisher compatibility remains bounded guidance only and verifier-compatible artifacts are not automatically trusted"},
		ProjectionDisclaimer:     verifierEcosystemValDProjectionDisclaimer(),
	}
}

func VerifierEcosystemValDNoOverclaimGateModel() VerifierEcosystemValDNoOverclaimGate {
	return VerifierEcosystemValDNoOverclaimGate{
		CurrentState:         "verifier_ecosystem_vald_no_overclaim_gate_ready",
		NoOverclaimGateID:    "no-overclaim-gate-vald",
		OutputRefs:           verifierEcosystemValDRequiredNoOverclaimRefs(),
		Caveats:              []string{"Val D remains an advisory final verifier ecosystem gate and cannot certify, approve, or close Točka 7"},
		ProjectionDisclaimer: verifierEcosystemValDProjectionDisclaimer(),
	}
}

func verifierEcosystemValDRequiredEvidenceIDs() []string {
	return []string{
		"evidence:correctness-gate-001",
		"evidence:tooling-gate-001",
		"evidence:schema-compatibility-gate-001",
		"evidence:diagnostics-conformance-gate-001",
		"evidence:trust-key-rotation-gate-001",
		"evidence:negative-diagnostics-gate-001",
		"evidence:redaction-gate-001",
		"evidence:publisher-artifact-gate-001",
		"evidence:no-overclaim-gate-001",
		"evidence:point7-governance-004",
	}
}

func verifierEcosystemValDRequiredEvidenceScopes() []string {
	return []string{
		"correctness_gate",
		"tooling_gate",
		"schema_compatibility_gate",
		"diagnostics_conformance_gate",
		"trust_key_rotation_gate",
		"negative_diagnostics_gate",
		"redaction_gate",
		"publisher_artifact_gate",
		"no_overclaim_gate",
		"point7_governance",
	}
}

func VerifierEcosystemValDVerifierEvidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:correctness-gate-001", EvidenceType: "correctness_gate", Source: "verifier/vald/correctness-gate", Timestamp: "2026-04-27T20:30:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "correctness_gate", Caveats: []string{"correctness gate remains bounded to exact Val 0 through Val C state consistency only"}},
		{EvidenceID: "evidence:tooling-gate-001", EvidenceType: "tooling_gate", Source: "verifier/vald/tooling-gate", Timestamp: "2026-04-27T20:31:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "tooling_gate", Caveats: []string{"tooling gate remains deterministic and advisory only"}},
		{EvidenceID: "evidence:schema-compatibility-gate-001", EvidenceType: "schema_compatibility_gate", Source: "verifier/vald/schema-compatibility-gate", Timestamp: "2026-04-27T20:32:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "schema_compatibility_gate", Caveats: []string{"schema compatibility remains version-bound and fail-closed"}},
		{EvidenceID: "evidence:diagnostics-conformance-gate-001", EvidenceType: "diagnostics_conformance_gate", Source: "verifier/vald/diagnostics-conformance-gate", Timestamp: "2026-04-27T20:33:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "diagnostics_conformance_gate", Caveats: []string{"deterministic diagnostics and conformance remain bounded verifier behavior checks only"}},
		{EvidenceID: "evidence:trust-key-rotation-gate-001", EvidenceType: "trust_key_rotation_gate", Source: "verifier/vald/trust-key-rotation-gate", Timestamp: "2026-04-27T20:34:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "trust_key_rotation_gate", Caveats: []string{"trust-root and key rotation gate remains scoped and non-global"}},
		{EvidenceID: "evidence:negative-diagnostics-gate-001", EvidenceType: "negative_diagnostics_gate", Source: "verifier/vald/negative-diagnostics-gate", Timestamp: "2026-04-27T20:35:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "negative_diagnostics_gate", Caveats: []string{"negative diagnostics remain non-verified and visible where required by output boundary"}},
		{EvidenceID: "evidence:redaction-gate-001", EvidenceType: "redaction_gate", Source: "verifier/vald/redaction-gate", Timestamp: "2026-04-27T20:36:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "redaction_gate", Caveats: []string{"redaction gate preserves bounded audience separation and does not create canonical truth"}},
		{EvidenceID: "evidence:publisher-artifact-gate-001", EvidenceType: "publisher_artifact_gate", Source: "verifier/vald/publisher-artifact-gate", Timestamp: "2026-04-27T20:37:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "publisher_artifact_gate", Caveats: []string{"publisher compatibility remains bounded guidance only"}},
		{EvidenceID: "evidence:no-overclaim-gate-001", EvidenceType: "no_overclaim_gate", Source: "verifier/vald/no-overclaim-gate", Timestamp: "2026-04-27T20:38:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "no_overclaim_gate", Caveats: []string{"no-overclaim gate forbids certification, approval, universal authority, and point_7_pass claims"}},
		{EvidenceID: "evidence:point7-governance-004", EvidenceType: "state_governance", Source: "verifier/point7-governance", Timestamp: "2026-04-27T20:39:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point7_governance", Caveats: []string{"Val D is not final closure and keeps point_7_state not complete"}},
	}
}

func VerifierEcosystemValDProofEvidenceRefs() []string {
	return []string{
		"point6_integrated_closure",
		"point7_verifier_discipline_foundation",
		"point7_reference_verifier_tooling",
		"point7_compatibility_diagnostics_conformance",
		"point7_public_partner_auditor_publisher_ecosystem",
		"point7_final_verifier_ecosystem_gate",
		"evidence:correctness-gate-001",
		"evidence:tooling-gate-001",
		"evidence:schema-compatibility-gate-001",
		"evidence:diagnostics-conformance-gate-001",
		"evidence:trust-key-rotation-gate-001",
		"evidence:negative-diagnostics-gate-001",
		"evidence:redaction-gate-001",
		"evidence:publisher-artifact-gate-001",
		"evidence:no-overclaim-gate-001",
		"evidence:point7-governance-004",
	}
}

func verifierEcosystemValDPriorValStateSeverity(state string) (int, bool) {
	switch strings.TrimSpace(state) {
	case VerifierEcosystemVal0StateActive, VerifierEcosystemValAStateActive, VerifierEcosystemValBStateActive, VerifierEcosystemValCStateActive:
		return 0, true
	case VerifierEcosystemVal0StatePartial, VerifierEcosystemValAStatePartial, VerifierEcosystemValBStatePartial, VerifierEcosystemValCStatePartial:
		return 1, true
	case VerifierEcosystemVal0StateIncomplete, VerifierEcosystemValAStateIncomplete, VerifierEcosystemValBStateIncomplete, VerifierEcosystemValCStateIncomplete:
		return 2, true
	case VerifierEcosystemVal0StateUnknown, VerifierEcosystemValAStateUnknown, VerifierEcosystemValBStateUnknown, VerifierEcosystemValCStateUnknown:
		return 3, true
	case VerifierEcosystemVal0StateBlocked, VerifierEcosystemValAStateBlocked, VerifierEcosystemValBStateBlocked, VerifierEcosystemValCStateBlocked:
		return 4, true
	default:
		return 3, false
	}
}

func verifierEcosystemValDStateSeverity(state, active, partial, incomplete, blocked, unknown string) int {
	switch strings.TrimSpace(state) {
	case active:
		return 0
	case partial:
		return 1
	case incomplete:
		return 2
	case unknown:
		return 3
	case blocked:
		return 4
	default:
		return 3
	}
}

func verifierEcosystemValDHighestSeverity(values ...int) int {
	highest := 0
	for _, value := range values {
		if value > highest {
			highest = value
		}
	}
	return highest
}

func verifierEcosystemValDSeverityToState(highest int, active, partial, incomplete, blocked, unknown string) string {
	switch highest {
	case 4:
		return blocked
	case 3:
		return unknown
	case 2:
		return incomplete
	case 1:
		return partial
	default:
		return active
	}
}

func verifierEcosystemValDPoint6DependencyHealthy(snapshot VerifierEcosystemValDDependencySnapshot) bool {
	return strings.TrimSpace(snapshot.Point5State) == IntelligenceCalibrationPoint5StatePass &&
		strings.TrimSpace(snapshot.Point5DependencyState) == IntelligenceCalibrationValEStateActive &&
		strings.TrimSpace(snapshot.Point6State) == ReferenceArchitecturePoint6StatePass &&
		strings.TrimSpace(snapshot.Point6ClosureState) == ReferenceArchitectureValEStateActive &&
		strings.TrimSpace(snapshot.Point6ClosurePrerequisiteState) == ReferenceArchitectureValEPrerequisiteStateActive &&
		strings.TrimSpace(snapshot.Point6ClosureInvariantState) == ReferenceArchitectureValEInvariantStateActive &&
		strings.TrimSpace(snapshot.Point6ProofSurfaceState) == ReferenceArchitectureValEProofSurfaceStateActive &&
		strings.TrimSpace(snapshot.Point6PassRuleState) == ReferenceArchitectureValEPassRuleStateActive &&
		snapshot.Point6PassAllowed
}

func verifierEcosystemValDProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, evidenceRefs []string) bool {
	if !containsExactTrimmedStringSet(evidenceRefs, VerifierEcosystemValDProofEvidenceRefs()...) {
		return false
	}
	allFresh, stale, ok := verifierEcosystemVal0EvidenceValid(evidence)
	if !ok || !allFresh || stale {
		return false
	}
	evidenceIDs := make([]string, 0, len(evidence))
	evidenceScopes := make([]string, 0, len(evidence))
	for _, item := range evidence {
		evidenceIDs = append(evidenceIDs, item.EvidenceID)
		evidenceScopes = append(evidenceScopes, item.Scope)
	}
	return containsExactTrimmedStringSet(evidenceIDs, verifierEcosystemValDRequiredEvidenceIDs()...) &&
		containsExactTrimmedStringSet(evidenceScopes, verifierEcosystemValDRequiredEvidenceScopes()...)
}

func EvaluateVerifierEcosystemValDCorrectnessGateState(model VerifierEcosystemValDCorrectnessGate) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.CorrectnessGateID,
		model.Version,
		model.VerifierContractState,
		model.ReferenceEngineState,
		model.CompatibilityMatrixState,
		model.DiagnosticsState,
		model.ConformanceSuiteState,
		model.EcosystemSurfaceState,
		model.ProjectionDisclaimer,
		model.LifecycleState,
		model.CompatibilityState,
		model.CreatedAt,
		model.UpdatedAt,
	) || len(model.SourceValStates) == 0 || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 || len(model.Limitations) == 0 {
		return VerifierEcosystemValDCorrectnessGateStateIncomplete
	}
	if !containsTrimmedString([]string{"active"}, model.LifecycleState) ||
		!containsTrimmedString(verifierEcosystemVal0CompatibilityResults(), model.CompatibilityState) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValDCorrectnessGateStateUnknown
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(model.CreatedAt); !ok {
		return VerifierEcosystemValDCorrectnessGateStatePartial
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(model.UpdatedAt); !ok {
		return VerifierEcosystemValDCorrectnessGateStatePartial
	}
	if model.CertificationClaim || model.ApprovalClaim || len(model.BlockingReasons) > 0 || verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), strings.Join(model.Limitations, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValDCorrectnessGateStateBlocked
	}
	valHighest := 0
	for _, item := range model.SourceValStates {
		severity, ok := verifierEcosystemValDPriorValStateSeverity(item)
		if !ok {
			return VerifierEcosystemValDCorrectnessGateStateUnknown
		}
		if severity > valHighest {
			valHighest = severity
		}
	}
	highest := verifierEcosystemValDHighestSeverity(
		valHighest,
		verifierEcosystemValDStateSeverity(model.VerifierContractState, VerifierEcosystemVal0ContractStateActive, VerifierEcosystemVal0ContractStatePartial, VerifierEcosystemVal0ContractStateIncomplete, VerifierEcosystemVal0ContractStateBlocked, VerifierEcosystemVal0ContractStateUnknown),
		verifierEcosystemValDStateSeverity(model.ReferenceEngineState, VerifierEcosystemValAEngineStateActive, VerifierEcosystemValAEngineStatePartial, VerifierEcosystemValAEngineStateIncomplete, VerifierEcosystemValAEngineStateBlocked, VerifierEcosystemValAEngineStateUnknown),
		verifierEcosystemValDStateSeverity(model.CompatibilityMatrixState, VerifierEcosystemValBCompatibilityMatrixStateActive, VerifierEcosystemValBCompatibilityMatrixStatePartial, VerifierEcosystemValBCompatibilityMatrixStateIncomplete, VerifierEcosystemValBCompatibilityMatrixStateBlocked, VerifierEcosystemValBCompatibilityMatrixStateUnknown),
		verifierEcosystemValDStateSeverity(model.DiagnosticsState, VerifierEcosystemValBDiagnosticPrecedenceStateActive, VerifierEcosystemValBDiagnosticPrecedenceStatePartial, VerifierEcosystemValBDiagnosticPrecedenceStateIncomplete, VerifierEcosystemValBDiagnosticPrecedenceStateBlocked, VerifierEcosystemValBDiagnosticPrecedenceStateUnknown),
		verifierEcosystemValDStateSeverity(model.ConformanceSuiteState, VerifierEcosystemValBConformanceSuiteStateActive, VerifierEcosystemValBConformanceSuiteStatePartial, VerifierEcosystemValBConformanceSuiteStateIncomplete, VerifierEcosystemValBConformanceSuiteStateBlocked, VerifierEcosystemValBConformanceSuiteStateUnknown),
		verifierEcosystemValDStateSeverity(model.EcosystemSurfaceState, VerifierEcosystemValCStateActive, VerifierEcosystemValCStatePartial, VerifierEcosystemValCStateIncomplete, VerifierEcosystemValCStateBlocked, VerifierEcosystemValCStateUnknown),
	)
	return verifierEcosystemValDSeverityToState(highest, VerifierEcosystemValDCorrectnessGateStateActive, VerifierEcosystemValDCorrectnessGateStatePartial, VerifierEcosystemValDCorrectnessGateStateIncomplete, VerifierEcosystemValDCorrectnessGateStateBlocked, VerifierEcosystemValDCorrectnessGateStateUnknown)
}

func EvaluateVerifierEcosystemValDToolingGateState(model VerifierEcosystemValDToolingGate) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ToolingGateID,
		model.VerifierInputModelState,
		model.ReferenceVerifierEngineState,
		model.VerificationResultModelState,
		model.DiagnosticsMappingState,
		model.CommandSurfaceState,
		model.SDKEntrypointState,
		model.ProjectionDisclaimer,
	) || len(model.Caveats) == 0 {
		return VerifierEcosystemValDToolingGateStateIncomplete
	}
	if !verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValDToolingGateStateUnknown
	}
	if model.MutatesEvidence || model.ApprovesDeployment || model.SuppressesFailures || model.CertificationClaim || model.HiddenMainInstanceDependency || model.NetworkDependency || model.ClaimsRealCryptoWithoutPrimitive || verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValDToolingGateStateBlocked
	}
	if !model.DeterministicOutput {
		return VerifierEcosystemValDToolingGateStateBlocked
	}
	highest := verifierEcosystemValDHighestSeverity(
		verifierEcosystemValDStateSeverity(model.VerifierInputModelState, VerifierEcosystemValAInputStateActive, VerifierEcosystemValAInputStatePartial, VerifierEcosystemValAInputStateIncomplete, VerifierEcosystemValAInputStateBlocked, VerifierEcosystemValAInputStateUnknown),
		verifierEcosystemValDStateSeverity(model.ReferenceVerifierEngineState, VerifierEcosystemValAEngineStateActive, VerifierEcosystemValAEngineStatePartial, VerifierEcosystemValAEngineStateIncomplete, VerifierEcosystemValAEngineStateBlocked, VerifierEcosystemValAEngineStateUnknown),
		verifierEcosystemValDStateSeverity(model.VerificationResultModelState, VerifierEcosystemValAResultStateActive, VerifierEcosystemValAResultStatePartial, VerifierEcosystemValAResultStateIncomplete, VerifierEcosystemValAResultStateBlocked, VerifierEcosystemValAResultStateUnknown),
		verifierEcosystemValDStateSeverity(model.DiagnosticsMappingState, VerifierEcosystemValADiagnosticsMappingStateActive, VerifierEcosystemValADiagnosticsMappingStatePartial, VerifierEcosystemValADiagnosticsMappingStateIncomplete, VerifierEcosystemValADiagnosticsMappingStateBlocked, VerifierEcosystemValADiagnosticsMappingStateUnknown),
		verifierEcosystemValDStateSeverity(model.CommandSurfaceState, VerifierEcosystemValACommandContractStateActive, VerifierEcosystemValACommandContractStatePartial, VerifierEcosystemValACommandContractStateIncomplete, VerifierEcosystemValACommandContractStateBlocked, VerifierEcosystemValACommandContractStateUnknown),
		verifierEcosystemValDStateSeverity(model.SDKEntrypointState, VerifierEcosystemValASDKEntrypointStateActive, VerifierEcosystemValASDKEntrypointStatePartial, VerifierEcosystemValASDKEntrypointStateIncomplete, VerifierEcosystemValASDKEntrypointStateBlocked, VerifierEcosystemValASDKEntrypointStateUnknown),
	)
	return verifierEcosystemValDSeverityToState(highest, VerifierEcosystemValDToolingGateStateActive, VerifierEcosystemValDToolingGateStatePartial, VerifierEcosystemValDToolingGateStateIncomplete, VerifierEcosystemValDToolingGateStateBlocked, VerifierEcosystemValDToolingGateStateUnknown)
}

func EvaluateVerifierEcosystemValDSchemaCompatibilityGateState(model VerifierEcosystemValDSchemaCompatibilityGate) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.SchemaCompatibilityGateID,
		model.CompatibilityMatrixState,
		model.SchemaProofCompatibilityState,
		model.MixedVersionDiagnosticsState,
		model.CompatibilityState,
		model.ProjectionDisclaimer,
	) || len(model.SupportedSchemaVersions) == 0 || len(model.SupportedProofTypes) == 0 || len(model.SupportedVerifierVersions) == 0 || len(model.SupportedTrustRootVersions) == 0 || len(model.CompatibilityEntryKeys) == 0 || len(model.MixedVersionRuleCoverage) == 0 || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return VerifierEcosystemValDSchemaCompatibilityGateStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemVal0CompatibilityResults(), model.CompatibilityState) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValDSchemaCompatibilityGateStateUnknown
	}
	if !containsExactTrimmedStringSet(model.SupportedSchemaVersions, VerifierEcosystemValBCompatibilityMatrixModel().SupportedSchemaVersions...) ||
		!containsExactTrimmedStringSet(model.SupportedProofTypes, verifierEcosystemVal0SupportedProofTypes()...) ||
		!containsExactTrimmedStringSet(model.SupportedVerifierVersions, verifierEcosystemValBSupportedVerifierVersions()...) ||
		!containsExactTrimmedStringSet(model.SupportedTrustRootVersions, verifierEcosystemValBSupportedTrustRootVersions()...) ||
		!containsExactTrimmedStringSet(model.MixedVersionRuleCoverage, VerifierEcosystemValBCompatibilityMatrixModel().MixedVersionRules...) {
		return VerifierEcosystemValDSchemaCompatibilityGateStatePartial
	}
	if !containsExactTrimmedStringSet(model.CompatibilityEntryKeys, verifierEcosystemValBRequiredCompatibilityEntryKeys()...) {
		return VerifierEcosystemValDSchemaCompatibilityGateStatePartial
	}
	if verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValDSchemaCompatibilityGateStateBlocked
	}
	switch strings.TrimSpace(model.CompatibilityState) {
	case ReferenceArchitectureCompatibilityUnsupported:
		return VerifierEcosystemValDSchemaCompatibilityGateStateBlocked
	case ReferenceArchitectureCompatibilityUnknown:
		return VerifierEcosystemValDSchemaCompatibilityGateStateUnknown
	case ReferenceArchitectureCompatibilityCompatibleWithWarning:
		if len(model.Caveats) == 0 {
			return VerifierEcosystemValDSchemaCompatibilityGateStateBlocked
		}
	case ReferenceArchitectureCompatibilityDeprecated:
		if !model.DeprecatedMigrationVisible || len(model.Caveats) == 0 {
			return VerifierEcosystemValDSchemaCompatibilityGateStateBlocked
		}
	case ReferenceArchitectureCompatibilitySuperseded:
		if !model.SupersessionVisible || len(model.Caveats) == 0 {
			return VerifierEcosystemValDSchemaCompatibilityGateStateBlocked
		}
	}
	highest := verifierEcosystemValDHighestSeverity(
		verifierEcosystemValDStateSeverity(model.CompatibilityMatrixState, VerifierEcosystemValBCompatibilityMatrixStateActive, VerifierEcosystemValBCompatibilityMatrixStatePartial, VerifierEcosystemValBCompatibilityMatrixStateIncomplete, VerifierEcosystemValBCompatibilityMatrixStateBlocked, VerifierEcosystemValBCompatibilityMatrixStateUnknown),
		verifierEcosystemValDStateSeverity(model.SchemaProofCompatibilityState, VerifierEcosystemValBSchemaProofCompatibilityStateActive, VerifierEcosystemValBSchemaProofCompatibilityStatePartial, VerifierEcosystemValBSchemaProofCompatibilityStateIncomplete, VerifierEcosystemValBSchemaProofCompatibilityStateBlocked, VerifierEcosystemValBSchemaProofCompatibilityStateUnknown),
		verifierEcosystemValDStateSeverity(model.MixedVersionDiagnosticsState, VerifierEcosystemValBMixedVersionStateActive, VerifierEcosystemValBMixedVersionStatePartial, VerifierEcosystemValBMixedVersionStateIncomplete, VerifierEcosystemValBMixedVersionStateBlocked, VerifierEcosystemValBMixedVersionStateUnknown),
	)
	return verifierEcosystemValDSeverityToState(highest, VerifierEcosystemValDSchemaCompatibilityGateStateActive, VerifierEcosystemValDSchemaCompatibilityGateStatePartial, VerifierEcosystemValDSchemaCompatibilityGateStateIncomplete, VerifierEcosystemValDSchemaCompatibilityGateStateBlocked, VerifierEcosystemValDSchemaCompatibilityGateStateUnknown)
}

func EvaluateVerifierEcosystemValDDiagnosticsConformanceGateState(model VerifierEcosystemValDDiagnosticsConformanceGate) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DiagnosticsConformanceGateID,
		model.DiagnosticPrecedenceState,
		model.FixtureDescriptorState,
		model.ConformanceCaseState,
		model.ConformanceSuiteState,
		model.OutputClassState,
		model.DerivedDiagnosticClass,
		model.ProjectionDisclaimer,
	) || len(model.ObservedDiagnostics) == 0 || len(model.Caveats) == 0 {
		return VerifierEcosystemValDDiagnosticsConformanceGateStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), model.DerivedDiagnosticClass) ||
		!containsAllTrimmedStrings(verifierEcosystemVal0DiagnosticClasses(), model.ObservedDiagnostics...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValDDiagnosticsConformanceGateStateUnknown
	}
	if model.CertificationClaim || model.IntegrityRatingClaim || verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValDDiagnosticsConformanceGateStateBlocked
	}
	expected := DeriveVerifierEcosystemValBDiagnostic(model.ObservedDiagnostics, model.Caveats)
	if expected == VerifierEcosystemDiagnosticUnknown {
		return VerifierEcosystemValDDiagnosticsConformanceGateStateUnknown
	}
	if strings.TrimSpace(model.DerivedDiagnosticClass) != expected {
		return VerifierEcosystemValDDiagnosticsConformanceGateStatePartial
	}
	if !model.StaleCoverageVisible || !model.RevokedCoverageVisible || !model.SupersededCoverageVisible || !model.UnsupportedCoverageVisible || !model.MalformedCoverageVisible {
		return VerifierEcosystemValDDiagnosticsConformanceGateStatePartial
	}
	highest := verifierEcosystemValDHighestSeverity(
		verifierEcosystemValDStateSeverity(model.DiagnosticPrecedenceState, VerifierEcosystemValBDiagnosticPrecedenceStateActive, VerifierEcosystemValBDiagnosticPrecedenceStatePartial, VerifierEcosystemValBDiagnosticPrecedenceStateIncomplete, VerifierEcosystemValBDiagnosticPrecedenceStateBlocked, VerifierEcosystemValBDiagnosticPrecedenceStateUnknown),
		verifierEcosystemValDStateSeverity(model.FixtureDescriptorState, VerifierEcosystemValBFixtureDescriptorStateActive, VerifierEcosystemValBFixtureDescriptorStatePartial, VerifierEcosystemValBFixtureDescriptorStateIncomplete, VerifierEcosystemValBFixtureDescriptorStateBlocked, VerifierEcosystemValBFixtureDescriptorStateUnknown),
		verifierEcosystemValDStateSeverity(model.ConformanceCaseState, VerifierEcosystemValBConformanceCaseStateActive, VerifierEcosystemValBConformanceCaseStatePartial, VerifierEcosystemValBConformanceCaseStateIncomplete, VerifierEcosystemValBConformanceCaseStateBlocked, VerifierEcosystemValBConformanceCaseStateUnknown),
		verifierEcosystemValDStateSeverity(model.ConformanceSuiteState, VerifierEcosystemValBConformanceSuiteStateActive, VerifierEcosystemValBConformanceSuiteStatePartial, VerifierEcosystemValBConformanceSuiteStateIncomplete, VerifierEcosystemValBConformanceSuiteStateBlocked, VerifierEcosystemValBConformanceSuiteStateUnknown),
		verifierEcosystemValDStateSeverity(model.OutputClassState, VerifierEcosystemValBOutputClassStateActive, VerifierEcosystemValBOutputClassStatePartial, VerifierEcosystemValBOutputClassStateIncomplete, VerifierEcosystemValBOutputClassStateBlocked, VerifierEcosystemValBOutputClassStateUnknown),
	)
	return verifierEcosystemValDSeverityToState(highest, VerifierEcosystemValDDiagnosticsConformanceGateStateActive, VerifierEcosystemValDDiagnosticsConformanceGateStatePartial, VerifierEcosystemValDDiagnosticsConformanceGateStateIncomplete, VerifierEcosystemValDDiagnosticsConformanceGateStateBlocked, VerifierEcosystemValDDiagnosticsConformanceGateStateUnknown)
}

func EvaluateVerifierEcosystemValDTrustKeyRotationGateState(model VerifierEcosystemValDTrustKeyRotationGate) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.TrustKeyRotationGateID,
		model.TrustState,
		model.TrustRootState,
		model.IssuerState,
		model.RevocationState,
		model.KeyRotationState,
		model.TrustDistributionState,
		model.TrustDistributionMode,
		model.OfflineDistributionScope,
		model.EvidenceFreshnessState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return VerifierEcosystemValDTrustKeyRotationGateStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemVal0TrustRootStates(), model.TrustRootState) ||
		!containsTrimmedString(verifierEcosystemVal0RevocationStates(), model.RevocationState) ||
		!containsTrimmedString(verifierEcosystemValDIssuerStates(), model.IssuerState) ||
		!containsTrimmedString([]string{VerifierEcosystemKeyRotationCurrent, VerifierEcosystemKeyRotationRollover}, model.KeyRotationState) ||
		!containsTrimmedString(verifierEcosystemValCDistributionModes(), model.TrustDistributionMode) ||
		!containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), model.OfflineDistributionScope) ||
		!containsTrimmedString(verifierEcosystemValCFreshnessStates(), model.EvidenceFreshnessState) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValDTrustKeyRotationGateStateUnknown
	}
	if model.GlobalKeyDirectoryClaim || model.SensitivePublicKeyExposure || verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValDTrustKeyRotationGateStateBlocked
	}
	if strings.TrimSpace(model.TrustDistributionMode) == VerifierEcosystemValCDistributionModeUnknown {
		return VerifierEcosystemValDTrustKeyRotationGateStateUnknown
	}
	switch strings.TrimSpace(model.RevocationState) {
	case VerifierEcosystemRevocationNotRevoked:
	case VerifierEcosystemRevocationRevoked, VerifierEcosystemRevocationExpired, VerifierEcosystemRevocationUnsupported:
		return VerifierEcosystemValDTrustKeyRotationGateStateBlocked
	case VerifierEcosystemRevocationUnknown:
		return VerifierEcosystemValDTrustKeyRotationGateStateUnknown
	default:
		return VerifierEcosystemValDTrustKeyRotationGateStateUnknown
	}
	switch strings.TrimSpace(model.IssuerState) {
	case verifierEcosystemValDIssuerStateBound:
	case verifierEcosystemValDIssuerStateWarning:
	case verifierEcosystemValDIssuerStateRevoked:
		return VerifierEcosystemValDTrustKeyRotationGateStateBlocked
	default:
		return VerifierEcosystemValDTrustKeyRotationGateStateUnknown
	}
	trustResult := VerifierEcosystemValDTrustKeyRotationGateStateUnknown
	switch strings.TrimSpace(model.TrustRootState) {
	case VerifierEcosystemTrustRootTrusted:
		trustResult = VerifierEcosystemValDTrustKeyRotationGateStateActive
	case VerifierEcosystemTrustRootTrustedWithWarnings:
		trustResult = VerifierEcosystemValDTrustKeyRotationGateStatePartial
	case VerifierEcosystemTrustRootRotated:
		if strings.TrimSpace(model.RolloverMetadataRef) == "" {
			return VerifierEcosystemValDTrustKeyRotationGateStateBlocked
		}
		trustResult = VerifierEcosystemValDTrustKeyRotationGateStatePartial
	case VerifierEcosystemTrustRootRevoked, VerifierEcosystemTrustRootExpired, VerifierEcosystemTrustRootUnsupported:
		return VerifierEcosystemValDTrustKeyRotationGateStateBlocked
	case VerifierEcosystemTrustRootUnknown:
		return VerifierEcosystemValDTrustKeyRotationGateStateUnknown
	default:
		return VerifierEcosystemValDTrustKeyRotationGateStateUnknown
	}
	switch strings.TrimSpace(model.EvidenceFreshnessState) {
	case IntelligenceCalibrationFreshnessFresh:
	case IntelligenceCalibrationFreshnessStale, IntelligenceCalibrationFreshnessExpired, IntelligenceCalibrationFreshnessUnsupported:
		return VerifierEcosystemValDTrustKeyRotationGateStateBlocked
	case IntelligenceCalibrationFreshnessUnknown:
		return VerifierEcosystemValDTrustKeyRotationGateStateUnknown
	default:
		return VerifierEcosystemValDTrustKeyRotationGateStateUnknown
	}
	if strings.TrimSpace(model.KeyRotationState) == VerifierEcosystemKeyRotationRollover {
		if strings.TrimSpace(model.RolloverMetadataRef) == "" {
			return VerifierEcosystemValDTrustKeyRotationGateStateBlocked
		}
		if trustResult == VerifierEcosystemValDTrustKeyRotationGateStateActive || trustResult == VerifierEcosystemValDTrustKeyRotationGateStatePartial {
			trustResult = VerifierEcosystemValDTrustKeyRotationGateStatePartial
		} else {
			return trustResult
		}
	}
	if strings.TrimSpace(model.OfflineDistributionScope) == VerifierEcosystemScopePublicSafe {
		return VerifierEcosystemValDTrustKeyRotationGateStateBlocked
	}
	highest := verifierEcosystemValDHighestSeverity(
		verifierEcosystemValDStateSeverity(model.TrustState, VerifierEcosystemVal0TrustStateActive, VerifierEcosystemVal0TrustStatePartial, VerifierEcosystemVal0TrustStateIncomplete, VerifierEcosystemVal0TrustStateBlocked, VerifierEcosystemVal0TrustStateUnknown),
		verifierEcosystemValDStateSeverity(model.TrustDistributionState, VerifierEcosystemValCTrustDistributionStateActive, VerifierEcosystemValCTrustDistributionStatePartial, VerifierEcosystemValCTrustDistributionStateIncomplete, VerifierEcosystemValCTrustDistributionStateBlocked, VerifierEcosystemValCTrustDistributionStateUnknown),
		verifierEcosystemValDStateSeverity(trustResult, VerifierEcosystemValDTrustKeyRotationGateStateActive, VerifierEcosystemValDTrustKeyRotationGateStatePartial, VerifierEcosystemValDTrustKeyRotationGateStateIncomplete, VerifierEcosystemValDTrustKeyRotationGateStateBlocked, VerifierEcosystemValDTrustKeyRotationGateStateUnknown),
	)
	return verifierEcosystemValDSeverityToState(highest, VerifierEcosystemValDTrustKeyRotationGateStateActive, VerifierEcosystemValDTrustKeyRotationGateStatePartial, VerifierEcosystemValDTrustKeyRotationGateStateIncomplete, VerifierEcosystemValDTrustKeyRotationGateStateBlocked, VerifierEcosystemValDTrustKeyRotationGateStateUnknown)
}

func verifierEcosystemValDNonVerifiedDiagnostic(diagnostic string) bool {
	switch strings.TrimSpace(diagnostic) {
	case VerifierEcosystemDiagnosticInvalidSignature,
		VerifierEcosystemDiagnosticDigestMismatch,
		VerifierEcosystemDiagnosticUnsupportedSchema,
		VerifierEcosystemDiagnosticUnsupportedProofType,
		VerifierEcosystemDiagnosticStaleArtifact,
		VerifierEcosystemDiagnosticExpiredArtifact,
		VerifierEcosystemDiagnosticRevokedIssuer,
		VerifierEcosystemDiagnosticSupersededProof,
		VerifierEcosystemDiagnosticInsufficientTrustMaterial,
		VerifierEcosystemDiagnosticIncompleteArtifact,
		VerifierEcosystemDiagnosticScopeMismatch,
		VerifierEcosystemDiagnosticRedactionViolation:
		return true
	default:
		return false
	}
}

func EvaluateVerifierEcosystemValDNegativeDiagnosticsGateState(model VerifierEcosystemValDNegativeDiagnosticsGate) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.NegativeDiagnosticsGateID,
		model.PublicOutputState,
		model.PartnerOutputState,
		model.AuditorFlowState,
		model.PublicOverallResult,
		model.PublicDiagnosticClass,
		model.PublicOutputClass,
		model.PartnerOverallResult,
		model.PartnerDiagnosticClass,
		model.PartnerOutputClass,
		model.ProjectionDisclaimer,
	) || len(model.Caveats) == 0 {
		return VerifierEcosystemValDNegativeDiagnosticsGateStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemValAOverallResults(), model.PublicOverallResult) ||
		!containsTrimmedString(verifierEcosystemValAOverallResults(), model.PartnerOverallResult) ||
		!containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), model.PublicDiagnosticClass) ||
		!containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), model.PartnerDiagnosticClass) ||
		!containsTrimmedString(verifierEcosystemValBOutputClasses(), model.PublicOutputClass) ||
		!containsTrimmedString(verifierEcosystemValBOutputClasses(), model.PartnerOutputClass) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValDNegativeDiagnosticsGateStateUnknown
	}
	if verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) ||
		!model.RedactionBoundaryPreserved ||
		!model.PublicPreservesNonVerified ||
		model.PartnerInternalDiagnosticsExposed ||
		!model.AuditorRepeatable ||
		!model.AuditorEvidenceLinked {
		return VerifierEcosystemValDNegativeDiagnosticsGateStateBlocked
	}
	if (verifierEcosystemValDNonVerifiedDiagnostic(model.PublicDiagnosticClass) && strings.TrimSpace(model.PublicOutputClass) == VerifierEcosystemValBOutputClassVerified) ||
		(verifierEcosystemValDNonVerifiedDiagnostic(model.PartnerDiagnosticClass) && strings.TrimSpace(model.PartnerOutputClass) == VerifierEcosystemValBOutputClassVerified) {
		return VerifierEcosystemValDNegativeDiagnosticsGateStateBlocked
	}
	if !model.StaleArtifactVisible || !model.ExpiredArtifactVisible || !model.RevokedIssuerVisible || !model.SupersededProofVisible || !model.UnsupportedSchemaVisible || !model.UnsupportedProofTypeVisible || !model.InsufficientTrustVisible {
		return VerifierEcosystemValDNegativeDiagnosticsGateStatePartial
	}
	highest := verifierEcosystemValDHighestSeverity(
		verifierEcosystemValDStateSeverity(model.PublicOutputState, VerifierEcosystemValCPublicOutputStateActive, VerifierEcosystemValCPublicOutputStatePartial, VerifierEcosystemValCPublicOutputStateIncomplete, VerifierEcosystemValCPublicOutputStateBlocked, VerifierEcosystemValCPublicOutputStateUnknown),
		verifierEcosystemValDStateSeverity(model.PartnerOutputState, VerifierEcosystemValCPartnerOutputStateActive, VerifierEcosystemValCPartnerOutputStatePartial, VerifierEcosystemValCPartnerOutputStateIncomplete, VerifierEcosystemValCPartnerOutputStateBlocked, VerifierEcosystemValCPartnerOutputStateUnknown),
		verifierEcosystemValDStateSeverity(model.AuditorFlowState, VerifierEcosystemValCAuditorFlowStateActive, VerifierEcosystemValCAuditorFlowStatePartial, VerifierEcosystemValCAuditorFlowStateIncomplete, VerifierEcosystemValCAuditorFlowStateBlocked, VerifierEcosystemValCAuditorFlowStateUnknown),
	)
	return verifierEcosystemValDSeverityToState(highest, VerifierEcosystemValDNegativeDiagnosticsGateStateActive, VerifierEcosystemValDNegativeDiagnosticsGateStatePartial, VerifierEcosystemValDNegativeDiagnosticsGateStateIncomplete, VerifierEcosystemValDNegativeDiagnosticsGateStateBlocked, VerifierEcosystemValDNegativeDiagnosticsGateStateUnknown)
}

func EvaluateVerifierEcosystemValDRedactionGateState(model VerifierEcosystemValDRedactionGate) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.RedactionGateID,
		model.AudienceSurfaceState,
		model.PublicOutputState,
		model.PartnerOutputState,
		model.AuditorFlowState,
		model.OutputBoundaryState,
		model.RedactionPolicyRef,
		model.EvidenceVisibilityPolicy,
		model.TrustMaterialVisibility,
		model.ProjectionDisclaimer,
	) || len(model.Caveats) == 0 {
		return VerifierEcosystemValDRedactionGateStateIncomplete
	}
	if !verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValDRedactionGateStateUnknown
	}
	if !model.InternalDiagnosticSeparated || verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValDRedactionGateStateBlocked
	}
	if !model.PartnerBroaderThanPublic {
		return VerifierEcosystemValDRedactionGateStatePartial
	}
	highest := verifierEcosystemValDHighestSeverity(
		verifierEcosystemValDStateSeverity(model.AudienceSurfaceState, VerifierEcosystemValCAudienceSurfaceStateActive, VerifierEcosystemValCAudienceSurfaceStatePartial, VerifierEcosystemValCAudienceSurfaceStateIncomplete, VerifierEcosystemValCAudienceSurfaceStateBlocked, VerifierEcosystemValCAudienceSurfaceStateUnknown),
		verifierEcosystemValDStateSeverity(model.PublicOutputState, VerifierEcosystemValCPublicOutputStateActive, VerifierEcosystemValCPublicOutputStatePartial, VerifierEcosystemValCPublicOutputStateIncomplete, VerifierEcosystemValCPublicOutputStateBlocked, VerifierEcosystemValCPublicOutputStateUnknown),
		verifierEcosystemValDStateSeverity(model.PartnerOutputState, VerifierEcosystemValCPartnerOutputStateActive, VerifierEcosystemValCPartnerOutputStatePartial, VerifierEcosystemValCPartnerOutputStateIncomplete, VerifierEcosystemValCPartnerOutputStateBlocked, VerifierEcosystemValCPartnerOutputStateUnknown),
		verifierEcosystemValDStateSeverity(model.AuditorFlowState, VerifierEcosystemValCAuditorFlowStateActive, VerifierEcosystemValCAuditorFlowStatePartial, VerifierEcosystemValCAuditorFlowStateIncomplete, VerifierEcosystemValCAuditorFlowStateBlocked, VerifierEcosystemValCAuditorFlowStateUnknown),
		verifierEcosystemValDStateSeverity(model.OutputBoundaryState, VerifierEcosystemVal0OutputBoundaryStateActive, VerifierEcosystemVal0OutputBoundaryStatePartial, VerifierEcosystemVal0OutputBoundaryStateIncomplete, VerifierEcosystemVal0OutputBoundaryStateBlocked, VerifierEcosystemVal0OutputBoundaryStateUnknown),
	)
	return verifierEcosystemValDSeverityToState(highest, VerifierEcosystemValDRedactionGateStateActive, VerifierEcosystemValDRedactionGateStatePartial, VerifierEcosystemValDRedactionGateStateIncomplete, VerifierEcosystemValDRedactionGateStateBlocked, VerifierEcosystemValDRedactionGateStateUnknown)
}

func EvaluateVerifierEcosystemValDPublisherArtifactGateState(model VerifierEcosystemValDPublisherArtifactGate) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.PublisherArtifactGateID,
		model.PublisherProfileState,
		model.ArtifactRuleState,
		model.PublisherType,
		model.RequiredSchemaPolicy,
		model.RequiredSignaturePolicy,
		model.RequiredDigestPolicy,
		model.TrustRootPolicy,
		model.ProjectionDisclaimer,
	) || len(model.SupportedArtifactTypes) == 0 || len(model.ConformanceCaseRefs) == 0 || len(model.Caveats) == 0 {
		return VerifierEcosystemValDPublisherArtifactGateStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemValCPublisherTypes(), model.PublisherType) ||
		!containsAllTrimmedStrings(verifierEcosystemVal0SupportedProofTypes(), model.SupportedArtifactTypes...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValDPublisherArtifactGateStateUnknown
	}
	if strings.TrimSpace(model.PublisherType) == VerifierEcosystemValCPublisherTypeUnknown ||
		!model.OutputBoundaryCompatible ||
		model.AutomaticallyTrustedClaim ||
		verifierEcosystemValDClaimsBlocked(model.ObservedClaims) ||
		verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValDPublisherArtifactGateStateBlocked
	}
	if !containsExactTrimmedStringSet(model.ConformanceCaseRefs, verifierEcosystemValBRequiredConformanceCaseIDs()...) {
		return VerifierEcosystemValDPublisherArtifactGateStatePartial
	}
	highest := verifierEcosystemValDHighestSeverity(
		verifierEcosystemValDStateSeverity(model.PublisherProfileState, VerifierEcosystemValCPublisherProfileStateActive, VerifierEcosystemValCPublisherProfileStatePartial, VerifierEcosystemValCPublisherProfileStateIncomplete, VerifierEcosystemValCPublisherProfileStateBlocked, VerifierEcosystemValCPublisherProfileStateUnknown),
		verifierEcosystemValDStateSeverity(model.ArtifactRuleState, VerifierEcosystemValCArtifactRuleStateActive, VerifierEcosystemValCArtifactRuleStatePartial, VerifierEcosystemValCArtifactRuleStateIncomplete, VerifierEcosystemValCArtifactRuleStateBlocked, VerifierEcosystemValCArtifactRuleStateUnknown),
	)
	return verifierEcosystemValDSeverityToState(highest, VerifierEcosystemValDPublisherArtifactGateStateActive, VerifierEcosystemValDPublisherArtifactGateStatePartial, VerifierEcosystemValDPublisherArtifactGateStateIncomplete, VerifierEcosystemValDPublisherArtifactGateStateBlocked, VerifierEcosystemValDPublisherArtifactGateStateUnknown)
}

func EvaluateVerifierEcosystemValDNoOverclaimGateState(model VerifierEcosystemValDNoOverclaimGate) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.NoOverclaimGateID, model.ProjectionDisclaimer) || len(model.OutputRefs) == 0 || len(model.Caveats) == 0 {
		return VerifierEcosystemValDNoOverclaimGateStateIncomplete
	}
	if !verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValDNoOverclaimGateStateUnknown
	}
	if !containsExactTrimmedStringSet(model.OutputRefs, verifierEcosystemValDRequiredNoOverclaimRefs()...) {
		return VerifierEcosystemValDNoOverclaimGateStatePartial
	}
	if verifierEcosystemValDClaimsBlocked(model.ObservedClaims) || verifierEcosystemValDContainsDisallowedClaim(strings.Join(model.Caveats, " ")) {
		return VerifierEcosystemValDNoOverclaimGateStateBlocked
	}
	return VerifierEcosystemValDNoOverclaimGateStateActive
}

func EvaluateVerifierEcosystemValDState(
	dependency VerifierEcosystemValDDependencySnapshot,
	correctnessGateState, toolingGateState, schemaCompatibilityGateState, diagnosticsConformanceGateState, trustKeyRotationGateState, negativeDiagnosticsGateState, redactionGateState, publisherArtifactGateState, noOverclaimGateState string,
) string {
	if !verifierEcosystemValDPoint6DependencyHealthy(dependency) ||
		strings.TrimSpace(dependency.Val0CurrentState) != VerifierEcosystemVal0StateActive ||
		strings.TrimSpace(dependency.Val0State) != VerifierEcosystemVal0StateActive ||
		strings.TrimSpace(dependency.ValACurrentState) != VerifierEcosystemValAStateActive ||
		strings.TrimSpace(dependency.ValAState) != VerifierEcosystemValAStateActive ||
		strings.TrimSpace(dependency.ValBCurrentState) != VerifierEcosystemValBStateActive ||
		strings.TrimSpace(dependency.ValBState) != VerifierEcosystemValBStateActive ||
		strings.TrimSpace(dependency.ValCCurrentState) != VerifierEcosystemValCStateActive ||
		strings.TrimSpace(dependency.ValCState) != VerifierEcosystemValCStateActive ||
		strings.TrimSpace(dependency.Point7State) != VerifierEcosystemPoint7StateNotComplete {
		return VerifierEcosystemValDStateBlocked
	}
	highest := verifierEcosystemValDHighestSeverity(
		verifierEcosystemValDStateSeverity(correctnessGateState, VerifierEcosystemValDCorrectnessGateStateActive, VerifierEcosystemValDCorrectnessGateStatePartial, VerifierEcosystemValDCorrectnessGateStateIncomplete, VerifierEcosystemValDCorrectnessGateStateBlocked, VerifierEcosystemValDCorrectnessGateStateUnknown),
		verifierEcosystemValDStateSeverity(toolingGateState, VerifierEcosystemValDToolingGateStateActive, VerifierEcosystemValDToolingGateStatePartial, VerifierEcosystemValDToolingGateStateIncomplete, VerifierEcosystemValDToolingGateStateBlocked, VerifierEcosystemValDToolingGateStateUnknown),
		verifierEcosystemValDStateSeverity(schemaCompatibilityGateState, VerifierEcosystemValDSchemaCompatibilityGateStateActive, VerifierEcosystemValDSchemaCompatibilityGateStatePartial, VerifierEcosystemValDSchemaCompatibilityGateStateIncomplete, VerifierEcosystemValDSchemaCompatibilityGateStateBlocked, VerifierEcosystemValDSchemaCompatibilityGateStateUnknown),
		verifierEcosystemValDStateSeverity(diagnosticsConformanceGateState, VerifierEcosystemValDDiagnosticsConformanceGateStateActive, VerifierEcosystemValDDiagnosticsConformanceGateStatePartial, VerifierEcosystemValDDiagnosticsConformanceGateStateIncomplete, VerifierEcosystemValDDiagnosticsConformanceGateStateBlocked, VerifierEcosystemValDDiagnosticsConformanceGateStateUnknown),
		verifierEcosystemValDStateSeverity(trustKeyRotationGateState, VerifierEcosystemValDTrustKeyRotationGateStateActive, VerifierEcosystemValDTrustKeyRotationGateStatePartial, VerifierEcosystemValDTrustKeyRotationGateStateIncomplete, VerifierEcosystemValDTrustKeyRotationGateStateBlocked, VerifierEcosystemValDTrustKeyRotationGateStateUnknown),
		verifierEcosystemValDStateSeverity(negativeDiagnosticsGateState, VerifierEcosystemValDNegativeDiagnosticsGateStateActive, VerifierEcosystemValDNegativeDiagnosticsGateStatePartial, VerifierEcosystemValDNegativeDiagnosticsGateStateIncomplete, VerifierEcosystemValDNegativeDiagnosticsGateStateBlocked, VerifierEcosystemValDNegativeDiagnosticsGateStateUnknown),
		verifierEcosystemValDStateSeverity(redactionGateState, VerifierEcosystemValDRedactionGateStateActive, VerifierEcosystemValDRedactionGateStatePartial, VerifierEcosystemValDRedactionGateStateIncomplete, VerifierEcosystemValDRedactionGateStateBlocked, VerifierEcosystemValDRedactionGateStateUnknown),
		verifierEcosystemValDStateSeverity(publisherArtifactGateState, VerifierEcosystemValDPublisherArtifactGateStateActive, VerifierEcosystemValDPublisherArtifactGateStatePartial, VerifierEcosystemValDPublisherArtifactGateStateIncomplete, VerifierEcosystemValDPublisherArtifactGateStateBlocked, VerifierEcosystemValDPublisherArtifactGateStateUnknown),
		verifierEcosystemValDStateSeverity(noOverclaimGateState, VerifierEcosystemValDNoOverclaimGateStateActive, VerifierEcosystemValDNoOverclaimGateStatePartial, VerifierEcosystemValDNoOverclaimGateStateIncomplete, VerifierEcosystemValDNoOverclaimGateStateBlocked, VerifierEcosystemValDNoOverclaimGateStateUnknown),
	)
	return verifierEcosystemValDSeverityToState(highest, VerifierEcosystemValDStateActive, VerifierEcosystemValDStatePartial, VerifierEcosystemValDStateIncomplete, VerifierEcosystemValDStateBlocked, VerifierEcosystemValDStateUnknown)
}

func VerifierEcosystemValDProofSurfaceRefs() []string {
	return []string{
		"/v1/verifier-ecosystem/val0/proofs",
		"/v1/verifier-ecosystem/vala/proofs",
		"/v1/verifier-ecosystem/valb/proofs",
		"/v1/verifier-ecosystem/valc/proofs",
		"/v1/verifier-ecosystem/vald/correctness-gate",
		"/v1/verifier-ecosystem/vald/tooling-gate",
		"/v1/verifier-ecosystem/vald/schema-compatibility-gate",
		"/v1/verifier-ecosystem/vald/diagnostics-conformance-gate",
		"/v1/verifier-ecosystem/vald/trust-key-rotation-gate",
		"/v1/verifier-ecosystem/vald/negative-diagnostics-gate",
		"/v1/verifier-ecosystem/vald/redaction-gate",
		"/v1/verifier-ecosystem/vald/publisher-artifact-gate",
		"/v1/verifier-ecosystem/vald/no-overclaim-gate",
		"/v1/verifier-ecosystem/vald/proofs",
	}
}

func EvaluateVerifierEcosystemValDProofsState(
	currentState string,
	point7State string,
	val0CurrentState string,
	valACurrentState string,
	valBCurrentState string,
	valCCurrentState string,
	surfaceRefs, evidenceRefs, limitations, whyPoint7NotPass []string,
	projectionDisclaimer string,
) string {
	baseState := strings.TrimSpace(currentState)
	if strings.TrimSpace(val0CurrentState) != VerifierEcosystemVal0StateActive ||
		strings.TrimSpace(valACurrentState) != VerifierEcosystemValAStateActive ||
		strings.TrimSpace(valBCurrentState) != VerifierEcosystemValBStateActive ||
		strings.TrimSpace(valCCurrentState) != VerifierEcosystemValCStateActive ||
		!containsExactTrimmedStringSet(surfaceRefs, VerifierEcosystemValDProofSurfaceRefs()...) ||
		!verifierEcosystemValDProofEvidenceQualityValid(VerifierEcosystemValDVerifierEvidence(), evidenceRefs) ||
		len(limitations) == 0 ||
		len(whyPoint7NotPass) == 0 ||
		!verifierEcosystemVal0HasProjectionDisclaimer(projectionDisclaimer) {
		if baseState == VerifierEcosystemValDStateActive {
			return VerifierEcosystemValDStatePartial
		}
		return baseState
	}
	if baseState == VerifierEcosystemValDStateActive && strings.TrimSpace(point7State) != VerifierEcosystemPoint7StateNotComplete {
		return VerifierEcosystemValDStatePartial
	}
	return baseState
}

func EvaluateVerifierEcosystemValDPoint7State(valDState string) string {
	_ = valDState
	return VerifierEcosystemPoint7StateNotComplete
}
