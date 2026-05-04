package formal

import "strings"

const (
	Point12ValEStateActive         = "point12_vale_integrated_replayable_proof_pack_closure_active"
	Point12ValEStateBlocked        = "point12_vale_integrated_replayable_proof_pack_closure_blocked"
	Point12ValEStateReviewRequired = "point12_vale_integrated_replayable_proof_pack_closure_review_required"
	Point12ValEStateIncomplete     = "point12_vale_integrated_replayable_proof_pack_closure_incomplete"
	Point12ValEStateUnsupported    = "point12_vale_integrated_replayable_proof_pack_closure_unsupported"
	Point12ValEStateTampered       = "point12_vale_integrated_replayable_proof_pack_closure_tampered"
	Point12ValEStateFailed         = "point12_vale_integrated_replayable_proof_pack_closure_failed"
	Point12ValEStatePassConfirmed  = "point12_vale_integrated_replayable_proof_pack_closure_pass_confirmed"
)

const (
	point12ValEWaveID                        = "val_e"
	point12ValEScope                         = "integrated_replayable_proof_pack_closure_gate"
	point12ValEProjectionDisclaimerBaseline  = "projection_only not_canonical_truth point12_vale_integrated_replayable_proof_pack_closure_gate"
	point12ValEDependencySnapshotRefBaseline = "dependency_snapshot_point12_vale_vald_computed_001"
	point12ValEPoint12PassToken              = "point_12_pass"
	point12ValEReviewerResultPassConfirmed   = "PASS_CONFIRMED"
	point12ValEReviewerResultPass            = "PASS"
	point12ValEReviewerResultReviewRequired  = "REVIEW_REQUIRED"
)

type Point12ValEValDReviewContext struct {
	SnapshotFromComputedOutput   bool     `json:"snapshot_from_computed_output"`
	ValDPrematurePoint12PassSeen bool     `json:"vald_premature_point12_pass_seen"`
	OpenCLB0Findings             int      `json:"open_clb0_findings"`
	OpenCLB1Findings             int      `json:"open_clb1_findings"`
	OpenCLB2Findings             int      `json:"open_clb2_findings"`
	ReviewPrerequisites          []string `json:"review_prerequisites,omitempty"`
}

type Point12ValEDependencySnapshot struct {
	Val0SnapshotRef              string                `json:"val0_snapshot_ref"`
	ValASnapshotRef              string                `json:"vala_snapshot_ref"`
	ValBSnapshotRef              string                `json:"valb_snapshot_ref"`
	ValCSnapshotRef              string                `json:"valc_snapshot_ref"`
	ValDSnapshotRef              string                `json:"vald_snapshot_ref"`
	ValDCurrentState             string                `json:"vald_current_state"`
	ValDDependencyState          string                `json:"vald_dependency_state"`
	ValDBindingMatrixState       string                `json:"vald_binding_matrix_state"`
	ValDProofChainState          string                `json:"vald_proof_chain_state"`
	ValDQueryState               string                `json:"vald_query_state"`
	ValDExplanationState         string                `json:"vald_explanation_state"`
	ValDSupportProfileState      string                `json:"vald_support_profile_state"`
	ValDPortalCompatibilityState string                `json:"vald_portal_compatibility_state"`
	ValDPointID                  string                `json:"vald_point_id"`
	ValDWaveID                   string                `json:"vald_wave_id"`
	ProjectionDisclaimer         string                `json:"projection_disclaimer"`
	SnapshotRef                  string                `json:"snapshot_ref"`
	SnapshotFromComputedOutput   bool                  `json:"snapshot_from_computed_output"`
	ValDExternalAPIUsed          bool                  `json:"vald_external_api_used"`
	ValDPointPassEmitted         bool                  `json:"vald_point_pass_emitted"`
	ValDPrematurePoint12PassSeen bool                  `json:"vald_premature_point12_pass_seen"`
	OpenCLB0Findings             int                   `json:"open_clb0_findings"`
	OpenCLB1Findings             int                   `json:"open_clb1_findings"`
	OpenCLB2Findings             int                   `json:"open_clb2_findings"`
	ReviewPrerequisites          []string              `json:"review_prerequisites,omitempty"`
	Val0                         Point12Val0Foundation `json:"val0"`
	ValA                         Point12ValAFoundation `json:"vala"`
	ValB                         Point12ValBFoundation `json:"valb"`
	ValC                         Point12ValCFoundation `json:"valc"`
	ValD                         Point12ValDFoundation `json:"vald"`
}

type Point12ValEFinalReplayInvariants struct {
	ReviewID                                 string   `json:"review_id"`
	ReplayMode                               string   `json:"replay_mode"`
	ReplayResultTaxonomy                     string   `json:"replay_result_taxonomy"`
	OriginalDecisionState                    string   `json:"original_decision_state"`
	ReplayedDecisionState                    string   `json:"replayed_decision_state"`
	MatchOriginal                            bool     `json:"match_original"`
	OriginalContextUsesCurrentPolicySilently bool     `json:"original_context_uses_current_policy_silently"`
	CurrentPolicyContextExplicit             bool     `json:"current_policy_context_explicit"`
	ComparisonModeDriftExplanationPresent    bool     `json:"comparison_mode_drift_explanation_present"`
	SameDecisionIsTaxonomyOnly               bool     `json:"same_decision_is_taxonomy_only"`
	TamperDetected                           bool     `json:"tamper_detected"`
	UnsupportedVersion                       bool     `json:"unsupported_version"`
	InsufficientEvidence                     bool     `json:"insufficient_evidence"`
	BlockedReplay                            bool     `json:"blocked_replay"`
	DifferentDecision                        bool     `json:"different_decision"`
	PolicyMismatch                           bool     `json:"policy_mismatch"`
	EngineMismatch                           bool     `json:"engine_mismatch"`
	SchemaMismatch                           bool     `json:"schema_mismatch"`
	EvidenceMismatch                         bool     `json:"evidence_mismatch"`
	ClaimMismatch                            bool     `json:"claim_mismatch"`
	GovernanceMismatch                       bool     `json:"governance_mismatch"`
	RedactionLimitations                     bool     `json:"redaction_limitations"`
	RedactedDecisiveEvidence                 bool     `json:"redacted_decisive_evidence"`
	MissingDecisiveEvidence                  bool     `json:"missing_decisive_evidence"`
	MismatchExpectedActualPresent            bool     `json:"mismatch_expected_actual_present"`
	DecisionDriftReasonsPresent              bool     `json:"decision_drift_reasons_present"`
	ExternalAPIUsed                          bool     `json:"external_api_used"`
	PointPassEmittedOutsideValE              bool     `json:"point_pass_emitted_outside_vale"`
	ProofPackID                              string   `json:"proof_pack_id"`
	ManifestID                               string   `json:"manifest_id"`
	ReplayResultID                           string   `json:"replay_result_id"`
	ExportID                                 string   `json:"export_id"`
	OfflineBundleID                          string   `json:"offline_bundle_id"`
	ManifestPayloadHash                      string   `json:"manifest_payload_hash"`
	SignatureMetadataRef                     string   `json:"signature_metadata_ref"`
	EvidenceHashCheckResult                  string   `json:"evidence_hash_check_result"`
	ManifestIntegrityCheckResult             string   `json:"manifest_integrity_check_result"`
	SignatureMetadataCheckResult             string   `json:"signature_metadata_check_result"`
	CompatibilityCheckResult                 string   `json:"compatibility_check_result"`
	DecisionDriftExplanation                 string   `json:"decision_drift_explanation"`
	MismatchExplanations                     []string `json:"mismatch_explanations,omitempty"`
	CurrentState                             string   `json:"current_state"`
	Diagnostics                              []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer                     string   `json:"projection_disclaimer"`
}

type Point12ValEEvidenceQualityMap struct {
	EvidenceQualityMapID string   `json:"evidence_quality_map_id"`
	ProofPackID          string   `json:"proof_pack_id"`
	ManifestID           string   `json:"manifest_id"`
	TenantScope          string   `json:"tenant_scope"`
	ArtifactRef          string   `json:"artifact_ref"`
	EvidenceRefs         []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs     []string `json:"evidence_hash_refs,omitempty"`
	PolicyRef            string   `json:"policy_ref"`
	PolicyVersion        string   `json:"policy_version"`
	PolicyHash           string   `json:"policy_hash"`
	EngineVersion        string   `json:"engine_version"`
	EngineHash           string   `json:"engine_hash"`
	SchemaVersion        string   `json:"schema_version"`
	SchemaHash           string   `json:"schema_hash"`
	ClaimRefs            []string `json:"claim_refs,omitempty"`
	GovernanceEventRefs  []string `json:"governance_event_refs,omitempty"`
	EvidenceStates       []string `json:"evidence_states,omitempty"`
	PolicyState          string   `json:"policy_state"`
	EngineState          string   `json:"engine_state"`
	SchemaState          string   `json:"schema_state"`
	ClaimStates          []string `json:"claim_states,omitempty"`
	GovernanceStates     []string `json:"governance_states,omitempty"`
	RedactionState       string   `json:"redaction_state"`
	ReplayState          string   `json:"replay_state"`
	ExportState          string   `json:"export_state"`
	ProofChainState      string   `json:"proof_chain_state"`
	FreshnessWindow      string   `json:"freshness_window"`
	StaleRefs            []string `json:"stale_refs,omitempty"`
	RevokedRefs          []string `json:"revoked_refs,omitempty"`
	ExpiredRefs          []string `json:"expired_refs,omitempty"`
	SupersededRefs       []string `json:"superseded_refs,omitempty"`
	UnsupportedRefs      []string `json:"unsupported_refs,omitempty"`
	MissingRefs          []string `json:"missing_refs,omitempty"`
	MalformedRefs        []string `json:"malformed_refs,omitempty"`
	DuplicateRefs        []string `json:"duplicate_refs,omitempty"`
	UnrelatedRefs        []string `json:"unrelated_refs,omitempty"`
	CrossTenantRefs      []string `json:"cross_tenant_refs,omitempty"`
	TamperedRefs         []string `json:"tampered_refs,omitempty"`
	QualityState         string   `json:"quality_state"`
}

type Point12ValEBindingMutationClosureReview struct {
	ReviewID                        string   `json:"review_id"`
	ValAManifestBindingState        string   `json:"vala_manifest_binding_state"`
	ValBReplayBindingState          string   `json:"valb_replay_binding_state"`
	ValCExportOfflineRedactionState string   `json:"valc_export_offline_redaction_state"`
	ValDProofChainBindingState      string   `json:"vald_proof_chain_binding_state"`
	RequiredValDBindingFields       []string `json:"required_vald_binding_fields,omitempty"`
	AdvisoryOnlyAffectsAuthority    bool     `json:"advisory_only_affects_authority"`
	CurrentState                    string   `json:"current_state"`
	Diagnostics                     []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer            string   `json:"projection_disclaimer"`
}

type Point12ValEProjectionBoundaryResult struct {
	BoundaryResultID                         string   `json:"boundary_result_id"`
	ProofChainProjectionNotSourceOfTruth     bool     `json:"proof_chain_projection_not_source_of_truth"`
	ExportBoundedProjection                  bool     `json:"export_bounded_projection"`
	OfflineBoundedVerificationPackage        bool     `json:"offline_bounded_verification_package"`
	FinancialInsuranceAuditMetadataOnly      bool     `json:"financial_insurance_audit_metadata_only"`
	PortalCompatibilityModelOnly             bool     `json:"portal_compatibility_model_only"`
	AgentFindingAdvisoryOnly                 bool     `json:"agent_finding_advisory_only"`
	BuyerProductCustomerTextCreatesAuthority bool     `json:"buyer_product_customer_text_creates_authority"`
	AuditorNotesMutateOutcome                bool     `json:"auditor_notes_mutate_outcome"`
	MutatesCanonicalEvidenceSpine            bool     `json:"mutates_canonical_evidence_spine"`
	EmitsPoint12Pass                         bool     `json:"emits_point12_pass"`
	CurrentState                             string   `json:"current_state"`
	Diagnostics                              []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer                     string   `json:"projection_disclaimer"`
}

type Point12ValENoOverclaimReview struct {
	ReviewID                     string   `json:"review_id"`
	ObservedCustomerTexts        []string `json:"observed_customer_texts,omitempty"`
	ObservedExportTexts          []string `json:"observed_export_texts,omitempty"`
	ObservedSupportTexts         []string `json:"observed_support_texts,omitempty"`
	ObservedPortalTexts          []string `json:"observed_portal_texts,omitempty"`
	BlockedClaimLedger           []string `json:"blocked_claim_ledger,omitempty"`
	BlockedClaimLedgerClassified bool     `json:"blocked_claim_ledger_classified"`
	InternalDiagnosticTexts      []string `json:"internal_diagnostic_texts,omitempty"`
	GrepRefs                     []string `json:"grep_refs,omitempty"`
	CurrentState                 string   `json:"current_state"`
	Diagnostics                  []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
}

type Point12ValECleanRoomIPReview struct {
	ReviewID                           string   `json:"review_id"`
	ThirdPartyRefs                     []string `json:"third_party_refs,omitempty"`
	LicenseReviewRefs                  []string `json:"license_review_refs,omitempty"`
	IPReviewRefs                       []string `json:"ip_review_refs,omitempty"`
	AIReviewPackageRefs                []string `json:"ai_review_package_refs,omitempty"`
	CompetitorCopyDetected             bool     `json:"competitor_copy_detected"`
	ProprietaryWorkflowClaimDetected   bool     `json:"proprietary_workflow_claim_detected"`
	ReverseEngineeringClaimDetected    bool     `json:"reverse_engineering_claim_detected"`
	PatentClearanceClaimDetected       bool     `json:"patent_clearance_claim_detected"`
	LegalCertificationClaimDetected    bool     `json:"legal_certification_claim_detected"`
	UnreviewedCustomerFacingDependency bool     `json:"unreviewed_customer_facing_dependency"`
	CurrentState                       string   `json:"current_state"`
	Diagnostics                        []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer               string   `json:"projection_disclaimer"`
}

type Point12ValERetentionProvenanceReview struct {
	ReviewID                                      string   `json:"review_id"`
	ProofPackRetentionClassRef                    string   `json:"proof_pack_retention_class_ref"`
	ExportRetentionClassRef                       string   `json:"export_retention_class_ref"`
	OfflineBundleRetentionClassRef                string   `json:"offline_bundle_retention_class_ref"`
	RedactionManifestRetentionClassRef            string   `json:"redaction_manifest_retention_class_ref"`
	AuditRetentionClassRef                        string   `json:"audit_retention_class_ref"`
	RetentionOwnerRef                             string   `json:"retention_owner_ref"`
	DisposalPathRef                               string   `json:"disposal_path_ref"`
	TenantScope                                   string   `json:"tenant_scope"`
	PublicPrivateClassification                   string   `json:"public_private_classification"`
	ToolchainProvenanceRefs                       []string `json:"toolchain_provenance_refs,omitempty"`
	AgentLineageRefs                              []string `json:"agent_lineage_refs,omitempty"`
	AgentLineageAdvisoryOnly                      bool     `json:"agent_lineage_advisory_only"`
	SupportPilotArtifactPromotedWithoutGovernance bool     `json:"support_pilot_artifact_promoted_without_governance"`
	CurrentState                                  string   `json:"current_state"`
	Diagnostics                                   []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer                          string   `json:"projection_disclaimer"`
}

type Point12ValEPassClosureManifest struct {
	CurrentState                string   `json:"current_state"`
	ClosureManifestID           string   `json:"closure_manifest_id"`
	PointID                     string   `json:"point_id"`
	WaveID                      string   `json:"wave_id"`
	Scope                       string   `json:"scope"`
	DependencyGateResult        string   `json:"dependency_gate_result"`
	Val0SnapshotRef             string   `json:"val0_snapshot_ref"`
	ValASnapshotRef             string   `json:"vala_snapshot_ref"`
	ValBSnapshotRef             string   `json:"valb_snapshot_ref"`
	ValCSnapshotRef             string   `json:"valc_snapshot_ref"`
	ValDSnapshotRef             string   `json:"vald_snapshot_ref"`
	ProofPackID                 string   `json:"proof_pack_id"`
	ManifestID                  string   `json:"manifest_id"`
	ReplayResultID              string   `json:"replay_result_id"`
	ExportID                    string   `json:"export_id"`
	OfflineBundleID             string   `json:"offline_bundle_id"`
	RedactionManifestID         string   `json:"redaction_manifest_id"`
	ProofChainID                string   `json:"proof_chain_id"`
	TenantScope                 string   `json:"tenant_scope"`
	ArtifactRef                 string   `json:"artifact_ref"`
	ArtifactHash                string   `json:"artifact_hash"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs            []string `json:"evidence_hash_refs,omitempty"`
	PolicyRef                   string   `json:"policy_ref"`
	PolicyVersion               string   `json:"policy_version"`
	PolicyHash                  string   `json:"policy_hash"`
	EngineVersion               string   `json:"engine_version"`
	EngineHash                  string   `json:"engine_hash"`
	SchemaVersion               string   `json:"schema_version"`
	SchemaHash                  string   `json:"schema_hash"`
	ClaimRefs                   []string `json:"claim_refs,omitempty"`
	GovernanceEventRefs         []string `json:"governance_event_refs,omitempty"`
	CompatibilityProfileRef     string   `json:"compatibility_profile_ref"`
	ManifestPayloadHash         string   `json:"manifest_payload_hash"`
	SignatureMetadataRef        string   `json:"signature_metadata_ref"`
	RetentionClassRef           string   `json:"retention_class_ref"`
	RetentionOwnerRef           string   `json:"retention_owner_ref"`
	DisposalPathRef             string   `json:"disposal_path_ref"`
	PublicPrivateClassification string   `json:"public_private_classification"`
	ToolchainProvenanceRefs     []string `json:"toolchain_provenance_refs,omitempty"`
	AgentLineageRefs            []string `json:"agent_lineage_refs,omitempty"`
	CommandsRun                 []string `json:"commands_run,omitempty"`
	TestsRun                    []string `json:"tests_run,omitempty"`
	NegativeFixturesRun         []string `json:"negative_fixtures_run,omitempty"`
	BindingMatrixResult         string   `json:"binding_matrix_result"`
	MutationTestResult          string   `json:"mutation_test_result"`
	ReplayInvariantResult       string   `json:"replay_invariant_result"`
	ExportOfflineBoundaryResult string   `json:"export_offline_boundary_result"`
	RedactionBoundaryResult     string   `json:"redaction_boundary_result"`
	ProofChainProjectionResult  string   `json:"proof_chain_projection_result"`
	EvidenceQualityMapResult    string   `json:"evidence_quality_map_result"`
	ProjectionBoundaryResult    string   `json:"projection_boundary_result"`
	NoOverclaimGrepResult       string   `json:"no_overclaim_grep_result"`
	CleanRoomIPResult           string   `json:"clean_room_ip_result"`
	RetentionDisposalResult     string   `json:"retention_disposal_result"`
	ReviewerResult              string   `json:"reviewer_result"`
	GeneratedAt                 string   `json:"generated_at"`
	CommitSHAIfAvailable        string   `json:"commit_sha_if_available"`
	Point12PassAllowed          bool     `json:"point12_pass_allowed"`
	Point12PassToken            string   `json:"point12_pass_token"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
}

type Point12ValEFoundation struct {
	CurrentState              string                                  `json:"current_state"`
	BlockingReasons           []string                                `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites       []string                                `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer      string                                  `json:"projection_disclaimer"`
	DependencyState           string                                  `json:"dependency_state"`
	ReplayInvariantState      string                                  `json:"replay_invariant_state"`
	EvidenceQualityState      string                                  `json:"evidence_quality_state"`
	BindingMutationState      string                                  `json:"binding_mutation_state"`
	ProjectionBoundaryState   string                                  `json:"projection_boundary_state"`
	NoOverclaimState          string                                  `json:"no_overclaim_state"`
	CleanRoomIPState          string                                  `json:"clean_room_ip_state"`
	RetentionProvenanceState  string                                  `json:"retention_provenance_state"`
	PassClosureManifestState  string                                  `json:"pass_closure_manifest_state"`
	Point12PassAllowed        bool                                    `json:"point12_pass_allowed"`
	Point12PassToken          string                                  `json:"point12_pass_token,omitempty"`
	Dependency                Point12ValEDependencySnapshot           `json:"dependency"`
	ReplayInvariants          Point12ValEFinalReplayInvariants        `json:"replay_invariants"`
	EvidenceQualityMap        Point12ValEEvidenceQualityMap           `json:"evidence_quality_map"`
	BindingMutationClosure    Point12ValEBindingMutationClosureReview `json:"binding_mutation_closure"`
	ProjectionBoundary        Point12ValEProjectionBoundaryResult     `json:"projection_boundary"`
	NoOverclaimReview         Point12ValENoOverclaimReview            `json:"no_overclaim_review"`
	CleanRoomIPReview         Point12ValECleanRoomIPReview            `json:"clean_room_ip_review"`
	RetentionProvenanceReview Point12ValERetentionProvenanceReview    `json:"retention_provenance_review"`
	PassClosureManifest       Point12ValEPassClosureManifest          `json:"pass_closure_manifest"`
}

func point12ValEStates() []string {
	return []string{
		Point12ValEStateActive,
		Point12ValEStateBlocked,
		Point12ValEStateReviewRequired,
		Point12ValEStateIncomplete,
		Point12ValEStateUnsupported,
		Point12ValEStateTampered,
		Point12ValEStateFailed,
		Point12ValEStatePassConfirmed,
	}
}

func point12ValEStateValid(value string) bool {
	return point11Val0ContainsTrimmed(point12ValEStates(), value)
}

func point12ValEReviewerResultValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValEReviewerResultPassConfirmed,
		point12ValEReviewerResultPass,
		point12ValEReviewerResultReviewRequired,
	}, value)
}

func point12ValEDependencySnapshotRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"dependency_snapshot_", "vald_snapshot_"})
}

func point12ValEClosureManifestRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"closure_manifest_", "manifest_"})
}

func point12ValEReviewRefValid(value string, prefixes ...string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, prefixes)
}

func point12ValECommandRunRefValid(value string) bool {
	return point12ValEReviewRefValid(value, "command_run_")
}

func point12ValETestRunRefValid(value string) bool {
	return point12ValEReviewRefValid(value, "test_run_")
}

func point12ValENegativeFixtureRunRefValid(value string) bool {
	return point12ValEReviewRefValid(value, "negative_fixture_")
}

func point12ValEGrepRunRefValid(value string) bool {
	return point12ValEReviewRefValid(value, "grep_run_")
}

func point12ValENonEmptyTextValid(value string) bool {
	return strings.TrimSpace(value) != ""
}

func point12ValEOptionalTextListValid(values []string) bool {
	for _, value := range values {
		if !point12ValENonEmptyTextValid(value) {
			return false
		}
	}
	return true
}

func point12ValEExactRefHashPairSetMatch(leftRefs, leftHashes, rightRefs, rightHashes []string) bool {
	if len(leftRefs) != len(leftHashes) || len(rightRefs) != len(rightHashes) || len(leftRefs) != len(rightRefs) {
		return false
	}
	left := map[string]string{}
	for idx, ref := range leftRefs {
		trimmedRef := strings.TrimSpace(ref)
		if trimmedRef == "" {
			return false
		}
		if _, exists := left[trimmedRef]; exists {
			return false
		}
		left[trimmedRef] = strings.TrimSpace(leftHashes[idx])
	}
	for idx, ref := range rightRefs {
		trimmedRef := strings.TrimSpace(ref)
		hash, exists := left[trimmedRef]
		if !exists || hash != strings.TrimSpace(rightHashes[idx]) {
			return false
		}
		delete(left, trimmedRef)
	}
	return len(left) == 0
}

func point12ValEForbiddenClaims() []string {
	return append(append([]string{}, point12Val0ForbiddenClaims()...),
		"legal protection guarantee",
		"rating boost",
		"premium discount",
	)
}

func point12ValEContainsForbiddenClaim(values ...string) bool {
	allowed := map[string]struct{}{}
	for _, value := range point12Val0AllowedClaims() {
		allowed[point11Val0NormalizeText(value)] = struct{}{}
	}
	for _, value := range values {
		normalized := point11Val0NormalizeText(value)
		if _, ok := allowed[normalized]; ok {
			continue
		}
		for _, forbidden := range point12ValEForbiddenClaims() {
			if strings.Contains(normalized, point11Val0NormalizeText(forbidden)) {
				return true
			}
		}
	}
	return false
}

func point12ValEInternalDiagnosticContextAllowed(value string) bool {
	normalized := point11Val0NormalizeText(value)
	return strings.Contains(normalized, "internal diagnostic") &&
		(strings.Contains(normalized, "removed") || strings.Contains(normalized, "disallowed") || strings.Contains(normalized, "blocked"))
}

func point12ValEDependencyReviewContextModel() Point12ValEValDReviewContext {
	return Point12ValEValDReviewContext{
		SnapshotFromComputedOutput: true,
	}
}

func SnapshotPoint12ValEDependencyFromComputed(
	val0 Point12Val0Foundation,
	valA Point12ValAFoundation,
	valB Point12ValBFoundation,
	valC Point12ValCFoundation,
	valD Point12ValDFoundation,
	review Point12ValEValDReviewContext,
) Point12ValEDependencySnapshot {
	reviewPrerequisites := append([]string{}, valD.ReviewPrerequisites...)
	reviewPrerequisites = append(reviewPrerequisites, review.ReviewPrerequisites...)
	return Point12ValEDependencySnapshot{
		Val0SnapshotRef:              valA.Dependency.SnapshotRef,
		ValASnapshotRef:              valB.Dependency.SnapshotRef,
		ValBSnapshotRef:              valC.Dependency.SnapshotRef,
		ValCSnapshotRef:              valD.Dependency.SnapshotRef,
		ValDSnapshotRef:              point12ValEDependencySnapshotRefBaseline,
		ValDCurrentState:             valD.CurrentState,
		ValDDependencyState:          valD.DependencyState,
		ValDBindingMatrixState:       valD.BindingMatrixState,
		ValDProofChainState:          valD.ProofChainState,
		ValDQueryState:               valD.QueryState,
		ValDExplanationState:         valD.ExplanationState,
		ValDSupportProfileState:      valD.SupportProfileState,
		ValDPortalCompatibilityState: valD.PortalCompatibilityState,
		ValDPointID:                  point12Val0PointID,
		ValDWaveID:                   point12ValDWaveID,
		ProjectionDisclaimer:         valD.ProjectionDisclaimer,
		SnapshotRef:                  point12ValEDependencySnapshotRefBaseline,
		SnapshotFromComputedOutput:   review.SnapshotFromComputedOutput,
		ValDExternalAPIUsed:          valD.Dependency.ValCExternalAPIUsed || valC.OfflineBundle.ExternalAPIUsed || valB.ReplayResult.ExternalAPIUsed,
		ValDPointPassEmitted:         review.ValDPrematurePoint12PassSeen,
		ValDPrematurePoint12PassSeen: review.ValDPrematurePoint12PassSeen,
		OpenCLB0Findings:             review.OpenCLB0Findings,
		OpenCLB1Findings:             review.OpenCLB1Findings,
		OpenCLB2Findings:             review.OpenCLB2Findings,
		ReviewPrerequisites:          reviewPrerequisites,
		Val0:                         val0,
		ValA:                         valA,
		ValB:                         valB,
		ValC:                         valC,
		ValD:                         valD,
	}
}

func point12ValEDependencySnapshotModel() Point12ValEDependencySnapshot {
	val0 := ComputePoint12Val0Foundation(Point12Val0FoundationModel())
	valA := ComputePoint12ValAFoundation(Point12ValAFoundationModel())
	valB := ComputePoint12ValBFoundation(Point12ValBFoundationModel())
	valC := ComputePoint12ValCFoundation(Point12ValCFoundationModel())
	valD := ComputePoint12ValDFoundation(Point12ValDFoundationModel())
	return SnapshotPoint12ValEDependencyFromComputed(val0, valA, valB, valC, valD, point12ValEDependencyReviewContextModel())
}

func point12ValEDependencyStateAndReasons(model Point12ValEDependencySnapshot) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12ValEDependencySnapshotRefValid(model.SnapshotRef) ||
		!model.SnapshotFromComputedOutput ||
		!point12ValADependencySnapshotRefValid(model.Val0SnapshotRef) ||
		!point12ValBDependencySnapshotRefValid(model.ValASnapshotRef) ||
		!point12ValCDependencySnapshotRefValid(model.ValBSnapshotRef) ||
		!point12ValDDependencySnapshotRefValid(model.ValCSnapshotRef) ||
		!point12ValEDependencySnapshotRefValid(model.ValDSnapshotRef) ||
		strings.TrimSpace(model.ValDPointID) != point12Val0PointID ||
		strings.TrimSpace(model.ValDWaveID) != point12ValDWaveID ||
		model.ValDExternalAPIUsed ||
		model.ValDPointPassEmitted ||
		model.ValDPrematurePoint12PassSeen ||
		model.OpenCLB0Findings > 0 ||
		model.OpenCLB1Findings > 0 ||
		model.OpenCLB2Findings > 0 {
		blockedReasons = append(blockedReasons, "dependency_identity_or_preflight_invalid")
	}
	if strings.TrimSpace(model.ValD.PortalCompatibility.PortalContractID) == "" ||
		!model.ValD.PortalCompatibility.ReadOnly ||
		!model.ValD.PortalCompatibility.NotesAnnotationOnly ||
		model.ValD.PortalCompatibility.EvidenceMutationAllowed ||
		model.ValD.PortalCompatibility.DecisionMutationAllowed ||
		model.ValD.PortalCompatibility.CertificationAllowed ||
		model.ValD.PortalCompatibility.PointPassAllowed {
		blockedReasons = append(blockedReasons, "dependency_portal_contract_invalid")
	}
	if !model.ValD.SupportProfile.AdvisoryOnly ||
		!model.ValD.SupportProfile.RequiredCustomerReview ||
		!model.ValD.SupportProfile.NoPremiumGuarantee ||
		!model.ValD.SupportProfile.NoRatingClaim ||
		!model.ValD.SupportProfile.NoComplianceGuarantee ||
		!model.ValD.SupportProfile.NoFinancialGuarantee ||
		!model.ValD.SupportProfile.NoLegalProtectionGuarantee {
		blockedReasons = append(blockedReasons, "dependency_support_profile_boundary_invalid")
	}
	if point12Val0ContainsPrematurePassToken(
		model.ValD.ProofChain.ProofChainID,
		model.ValD.Query.RequestedExplanation,
		model.ValD.Explanation.CustomerVisibleStatement,
		model.ValD.SupportProfile.SupportStatement,
		model.ValD.PortalCompatibility.RequiredProjectionDisclaimer,
		model.ValA.Manifest.ProofPackID,
		model.ValB.ReplayRequest.ProofPackID,
		model.ValC.ExportBundle.ExportID,
	) {
		blockedReasons = append(blockedReasons, "dependency_contains_point12_pass_input")
	}
	if strings.TrimSpace(model.ValDCurrentState) == Point12ValDStateBlocked ||
		strings.TrimSpace(model.ValDDependencyState) == Point12ValDDependencyStateBlocked ||
		strings.TrimSpace(model.ValDBindingMatrixState) == Point12ValDBindingMatrixStateBlocked ||
		strings.TrimSpace(model.ValDProofChainState) == Point12ValDProofChainStateBlocked ||
		strings.TrimSpace(model.ValDExplanationState) == Point12ValDExplanationStateBlocked ||
		strings.TrimSpace(model.ValDSupportProfileState) == Point12ValDSupportProfileStateBlocked ||
		strings.TrimSpace(model.ValDPortalCompatibilityState) == Point12ValDPortalCompatibilityStateBlocked {
		blockedReasons = append(blockedReasons, "dependency_vald_blocked")
	}
	if strings.TrimSpace(model.ValDCurrentState) == Point12ValDStateReviewRequired ||
		strings.TrimSpace(model.ValDDependencyState) == Point12ValDDependencyStateReviewRequired ||
		strings.TrimSpace(model.ValDBindingMatrixState) == Point12ValDBindingMatrixStateReviewRequired ||
		strings.TrimSpace(model.ValDProofChainState) == Point12ValDProofChainStateReviewRequired ||
		strings.TrimSpace(model.ValDQueryState) == Point12ValDQueryStateReviewRequired ||
		strings.TrimSpace(model.ValDExplanationState) == Point12ValDExplanationStateReviewRequired ||
		strings.TrimSpace(model.ValDSupportProfileState) == Point12ValDSupportProfileStateReviewRequired ||
		len(model.ReviewPrerequisites) > 0 {
		reviewReasons = append(reviewReasons, "dependency_vald_review_required")
	}
	if strings.TrimSpace(model.ValDCurrentState) != Point12ValDStateActive ||
		strings.TrimSpace(model.ValDDependencyState) != Point12ValDDependencyStateActive ||
		strings.TrimSpace(model.ValDBindingMatrixState) != Point12ValDBindingMatrixStateActive ||
		strings.TrimSpace(model.ValDProofChainState) != Point12ValDProofChainStateActive ||
		strings.TrimSpace(model.ValDQueryState) != Point12ValDQueryStateActive ||
		strings.TrimSpace(model.ValDExplanationState) != Point12ValDExplanationStateActive ||
		strings.TrimSpace(model.ValDSupportProfileState) != Point12ValDSupportProfileStateActive ||
		strings.TrimSpace(model.ValDPortalCompatibilityState) != Point12ValDPortalCompatibilityStateActive {
		if len(blockedReasons) == 0 && len(reviewReasons) == 0 {
			blockedReasons = append(blockedReasons, "dependency_vald_not_active")
		}
	}
	if len(blockedReasons) > 0 {
		return Point12ValEStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValEStateReviewRequired, reviewReasons
	}
	return Point12ValEStateActive, nil
}

func EvaluatePoint12ValEDependencyState(model Point12ValEDependencySnapshot) string {
	state, _ := point12ValEDependencyStateAndReasons(model)
	return state
}

func point12ValEFinalReplayInvariantsModel(dependency Point12ValEDependencySnapshot) Point12ValEFinalReplayInvariants {
	result := dependency.ValB.ReplayResult
	request := dependency.ValB.ReplayRequest
	semantics := point12ValEComputeDependencyReplaySemantics(result)
	return Point12ValEFinalReplayInvariants{
		ReviewID:                              "replay_invariant_point12_vale_001",
		ReplayMode:                            semantics.ReplayMode,
		ReplayResultTaxonomy:                  semantics.ReplayResultTaxonomy,
		OriginalDecisionState:                 result.OriginalDecisionState,
		ReplayedDecisionState:                 result.ReplayedDecisionState,
		MatchOriginal:                         result.MatchOriginal,
		CurrentPolicyContextExplicit:          strings.TrimSpace(semantics.ReplayMode) != point12Val0ReplayModeCurrentPolicyContext || (strings.TrimSpace(request.CurrentPolicyRef) != "" && strings.TrimSpace(request.CurrentPolicyHash) != "" && strings.TrimSpace(request.CurrentPolicyVersion) != ""),
		ComparisonModeDriftExplanationPresent: strings.TrimSpace(semantics.ReplayMode) != point12Val0ReplayModeComparisonMode || semantics.DecisionDriftReasonsPresent,
		SameDecisionIsTaxonomyOnly:            true,
		TamperDetected:                        semantics.TamperDetected,
		UnsupportedVersion:                    semantics.UnsupportedVersion,
		InsufficientEvidence:                  semantics.InsufficientEvidence,
		BlockedReplay:                         semantics.BlockedReplay,
		DifferentDecision:                     semantics.DifferentDecision,
		PolicyMismatch:                        semantics.PolicyMismatch,
		EngineMismatch:                        semantics.EngineMismatch,
		SchemaMismatch:                        semantics.SchemaMismatch,
		EvidenceMismatch:                      semantics.EvidenceMismatch,
		ClaimMismatch:                         semantics.ClaimMismatch,
		GovernanceMismatch:                    semantics.GovernanceMismatch,
		RedactionLimitations:                  semantics.RedactionLimitations,
		RedactedDecisiveEvidence:              dependency.ValC.RedactionImpactVerdict.DecisiveEvidenceRemoved,
		MissingDecisiveEvidence:               semantics.InsufficientEvidence,
		MismatchExpectedActualPresent:         semantics.MismatchExpectedActualPresent,
		DecisionDriftReasonsPresent:           strings.TrimSpace(semantics.ReplayMode) != point12Val0ReplayModeComparisonMode || semantics.DecisionDriftReasonsPresent,
		ExternalAPIUsed:                       semantics.ExternalAPIUsed || dependency.ValC.OfflineBundle.ExternalAPIUsed,
		PointPassEmittedOutsideValE:           semantics.PointPassEmitted,
		ProofPackID:                           result.ProofPackID,
		ManifestID:                            result.ManifestID,
		ReplayResultID:                        result.ReplayResultID,
		ExportID:                              dependency.ValC.ExportBundle.ExportID,
		OfflineBundleID:                       dependency.ValC.OfflineBundle.OfflineBundleID,
		ManifestPayloadHash:                   request.ManifestPayloadHash,
		SignatureMetadataRef:                  dependency.ValA.Manifest.SignatureMetadataRef,
		EvidenceHashCheckResult:               result.EvidenceHashCheckResult,
		ManifestIntegrityCheckResult:          result.ManifestIntegrityCheckResult,
		SignatureMetadataCheckResult:          result.SignatureMetadataCheckResult,
		CompatibilityCheckResult:              result.CompatibilityCheckResult,
		DecisionDriftExplanation:              result.DecisionDriftExplanation,
		MismatchExplanations:                  append([]string{}, result.MismatchExplanations...),
		CurrentState:                          Point12ValEStateActive,
		Diagnostics:                           []string{"final_replay_invariants_bound_to_vala_and_valc"},
		ProjectionDisclaimer:                  point12ValEProjectionDisclaimerBaseline,
	}
}

func point12ValEReplayMismatchesHaveExpectedActual(result Point12ValBReplayResult) bool {
	if len(result.Mismatches) == 0 {
		return true
	}
	for _, mismatch := range result.Mismatches {
		if strings.TrimSpace(mismatch.ExpectedRef) == "" && strings.TrimSpace(mismatch.ExpectedHash) == "" && strings.TrimSpace(mismatch.ExpectedVersion) == "" {
			return false
		}
		if strings.TrimSpace(mismatch.ActualRef) == "" && strings.TrimSpace(mismatch.ActualHash) == "" && strings.TrimSpace(mismatch.ActualVersion) == "" {
			return false
		}
	}
	return true
}

type point12ValEDependencyReplaySemantics struct {
	ReplayMode                    string
	ReplayResultTaxonomy          string
	OriginalDecisionState         string
	ReplayedDecisionState         string
	MatchOriginal                 bool
	TamperDetected                bool
	UnsupportedVersion            bool
	InsufficientEvidence          bool
	BlockedReplay                 bool
	DifferentDecision             bool
	PolicyMismatch                bool
	EngineMismatch                bool
	SchemaMismatch                bool
	EvidenceMismatch              bool
	ClaimMismatch                 bool
	GovernanceMismatch            bool
	RedactionLimitations          bool
	MismatchExpectedActualPresent bool
	DecisionDriftReasonsPresent   bool
	DecisionDriftExplanation      string
	MismatchExplanations          []string
	EvidenceHashCheckResult       string
	ManifestIntegrityCheckResult  string
	SignatureMetadataCheckResult  string
	CompatibilityCheckResult      string
	ExternalAPIUsed               bool
	PointPassEmitted              bool
	ProjectionDisclaimer          string
}

func point12ValEComputeDependencyReplaySemantics(result Point12ValBReplayResult) point12ValEDependencyReplaySemantics {
	taxonomy := strings.TrimSpace(result.ReplayResultTaxonomy)
	return point12ValEDependencyReplaySemantics{
		ReplayMode:                    result.ReplayMode,
		ReplayResultTaxonomy:          taxonomy,
		OriginalDecisionState:         result.OriginalDecisionState,
		ReplayedDecisionState:         result.ReplayedDecisionState,
		MatchOriginal:                 result.MatchOriginal,
		TamperDetected:                result.TamperDetected || taxonomy == Point12Val0ReplayResultTamperDetected || result.ManifestIntegrityCheckResult == point12ValBCheckResultTampered || result.SignatureMetadataCheckResult == point12ValBCheckResultTampered || result.EvidenceHashCheckResult == point12ValBCheckResultTampered || point12ValBHasMismatchType(result.Mismatches, point12ValBMismatchTypeTamperDetected),
		UnsupportedVersion:            result.UnsupportedVersion || taxonomy == Point12Val0ReplayResultUnsupportedVersion || result.ManifestIntegrityCheckResult == point12ValBCheckResultUnsupported || result.CompatibilityCheckResult == point12ValBCheckResultUnsupported || point12ValBHasMismatchType(result.Mismatches, point12ValBMismatchTypeUnsupportedVersion),
		InsufficientEvidence:          result.InsufficientEvidence || taxonomy == Point12Val0ReplayResultInsufficientEvidence || result.EvidenceHashCheckResult == point12ValBCheckResultMissing || point12ValBHasMismatchType(result.Mismatches, point12ValBMismatchTypeMissingEvidence),
		BlockedReplay:                 strings.TrimSpace(result.ReplayState) == Point12ValBReplayResultStateBlocked || taxonomy == Point12Val0ReplayResultBlockedReplay || point12ValBHasReplayBlockingMismatch(result.Mismatches),
		DifferentDecision:             taxonomy == Point12Val0ReplayResultDifferentDecision,
		PolicyMismatch:                taxonomy == Point12Val0ReplayResultPolicyMismatch || point12ValBHasMismatchType(result.Mismatches, point12ValBMismatchTypePolicyMismatch),
		EngineMismatch:                taxonomy == Point12Val0ReplayResultEngineMismatch || point12ValBHasMismatchType(result.Mismatches, point12ValBMismatchTypeEngineMismatch),
		SchemaMismatch:                taxonomy == Point12Val0ReplayResultSchemaMismatch || point12ValBHasMismatchType(result.Mismatches, point12ValBMismatchTypeSchemaMismatch),
		EvidenceMismatch:              taxonomy == Point12Val0ReplayResultEvidenceMismatch || point12ValBHasMismatchType(result.Mismatches, point12ValBMismatchTypeEvidenceMismatch),
		ClaimMismatch:                 taxonomy == Point12Val0ReplayResultClaimMismatch || point12ValBHasMismatchType(result.Mismatches, point12ValBMismatchTypeClaimMismatch),
		GovernanceMismatch:            taxonomy == Point12Val0ReplayResultGovernanceMismatch || point12ValBHasMismatchType(result.Mismatches, point12ValBMismatchTypeGovernanceMismatch),
		RedactionLimitations:          result.RedactionLimitations || taxonomy == Point12Val0ReplayResultRedactedLimitations || point12ValBHasMismatchType(result.Mismatches, point12ValBMismatchTypeRedactionMismatch),
		MismatchExpectedActualPresent: point12ValEReplayMismatchesHaveExpectedActual(result),
		DecisionDriftReasonsPresent:   strings.TrimSpace(result.ReplayMode) != point12Val0ReplayModeComparisonMode || strings.TrimSpace(result.DecisionDriftExplanation) != "",
		DecisionDriftExplanation:      result.DecisionDriftExplanation,
		MismatchExplanations:          append([]string{}, result.MismatchExplanations...),
		EvidenceHashCheckResult:       result.EvidenceHashCheckResult,
		ManifestIntegrityCheckResult:  result.ManifestIntegrityCheckResult,
		SignatureMetadataCheckResult:  result.SignatureMetadataCheckResult,
		CompatibilityCheckResult:      result.CompatibilityCheckResult,
		ExternalAPIUsed:               result.ExternalAPIUsed,
		PointPassEmitted:              result.PointPassEmitted,
		ProjectionDisclaimer:          result.ProjectionDisclaimer,
	}
}

func point12ValEReplaySemanticsMismatchReasons(model Point12ValEFinalReplayInvariants, semantics point12ValEDependencyReplaySemantics) []string {
	reasons := []string{}
	for _, mismatch := range []struct {
		field string
		got   bool
		want  bool
	}{
		{field: "tamper_detected", got: model.TamperDetected, want: semantics.TamperDetected},
		{field: "unsupported_version", got: model.UnsupportedVersion, want: semantics.UnsupportedVersion},
		{field: "insufficient_evidence", got: model.InsufficientEvidence, want: semantics.InsufficientEvidence},
		{field: "blocked_replay", got: model.BlockedReplay, want: semantics.BlockedReplay},
		{field: "different_decision", got: model.DifferentDecision, want: semantics.DifferentDecision},
		{field: "policy_mismatch", got: model.PolicyMismatch, want: semantics.PolicyMismatch},
		{field: "engine_mismatch", got: model.EngineMismatch, want: semantics.EngineMismatch},
		{field: "schema_mismatch", got: model.SchemaMismatch, want: semantics.SchemaMismatch},
		{field: "evidence_mismatch", got: model.EvidenceMismatch, want: semantics.EvidenceMismatch},
		{field: "claim_mismatch", got: model.ClaimMismatch, want: semantics.ClaimMismatch},
		{field: "governance_mismatch", got: model.GovernanceMismatch, want: semantics.GovernanceMismatch},
		{field: "redaction_limitations", got: model.RedactionLimitations, want: semantics.RedactionLimitations},
		{field: "mismatch_expected_actual_present", got: model.MismatchExpectedActualPresent, want: semantics.MismatchExpectedActualPresent},
		{field: "decision_drift_reasons_present", got: model.DecisionDriftReasonsPresent, want: semantics.DecisionDriftReasonsPresent},
		{field: "external_api_used", got: model.ExternalAPIUsed, want: semantics.ExternalAPIUsed},
		{field: "point_pass_emitted", got: model.PointPassEmittedOutsideValE, want: semantics.PointPassEmitted},
	} {
		if mismatch.got != mismatch.want {
			reasons = append(reasons, "replay_invariant_dependency_semantics_mismatch:"+mismatch.field)
		}
	}
	if strings.TrimSpace(model.ReplayMode) != strings.TrimSpace(semantics.ReplayMode) {
		reasons = append(reasons, "replay_invariant_dependency_semantics_mismatch:replay_mode")
	}
	if strings.TrimSpace(model.ReplayResultTaxonomy) != strings.TrimSpace(semantics.ReplayResultTaxonomy) {
		reasons = append(reasons, "replay_invariant_dependency_semantics_mismatch:replay_result_taxonomy")
	}
	if strings.TrimSpace(model.OriginalDecisionState) != strings.TrimSpace(semantics.OriginalDecisionState) {
		reasons = append(reasons, "replay_invariant_dependency_semantics_mismatch:original_decision_state")
	}
	if strings.TrimSpace(model.ReplayedDecisionState) != strings.TrimSpace(semantics.ReplayedDecisionState) {
		reasons = append(reasons, "replay_invariant_dependency_semantics_mismatch:replayed_decision_state")
	}
	if model.MatchOriginal != semantics.MatchOriginal {
		reasons = append(reasons, "replay_invariant_dependency_semantics_mismatch:match_original")
	}
	if strings.TrimSpace(model.DecisionDriftExplanation) != strings.TrimSpace(semantics.DecisionDriftExplanation) {
		reasons = append(reasons, "replay_invariant_dependency_semantics_mismatch:decision_drift_explanation")
	}
	if !point12Val0ExactStringSetMatch(model.MismatchExplanations, semantics.MismatchExplanations) {
		reasons = append(reasons, "replay_invariant_dependency_semantics_mismatch:mismatch_explanations")
	}
	if strings.TrimSpace(model.EvidenceHashCheckResult) != strings.TrimSpace(semantics.EvidenceHashCheckResult) {
		reasons = append(reasons, "replay_invariant_dependency_semantics_mismatch:evidence_hash_check_result")
	}
	if strings.TrimSpace(model.ManifestIntegrityCheckResult) != strings.TrimSpace(semantics.ManifestIntegrityCheckResult) {
		reasons = append(reasons, "replay_invariant_dependency_semantics_mismatch:manifest_integrity_check_result")
	}
	if strings.TrimSpace(model.SignatureMetadataCheckResult) != strings.TrimSpace(semantics.SignatureMetadataCheckResult) {
		reasons = append(reasons, "replay_invariant_dependency_semantics_mismatch:signature_metadata_check_result")
	}
	if strings.TrimSpace(model.CompatibilityCheckResult) != strings.TrimSpace(semantics.CompatibilityCheckResult) {
		reasons = append(reasons, "replay_invariant_dependency_semantics_mismatch:compatibility_check_result")
	}
	return reasons
}

func point12ValEFinalReplayInvariantStateAndReasons(model Point12ValEFinalReplayInvariants, dependency Point12ValEDependencySnapshot) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	dependencyReplaySemantics := point12ValEComputeDependencyReplaySemantics(dependency.ValB.ReplayResult)
	if !point12ValEReviewRefValid(model.ReviewID, "replay_invariant_") ||
		!point12Val0ReplayModeValid(model.ReplayMode) ||
		!point11Val0ContainsTrimmed(point12Val0ReplayResults(), model.ReplayResultTaxonomy) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point12ValBReplayResultRefValid(model.ReplayResultID) ||
		!point12ValCExportRefValid(model.ExportID) ||
		!point12ValCOfflineBundleRefValid(model.OfflineBundleID) ||
		!point12Val0HashValid(model.ManifestPayloadHash) ||
		!point12ValASignatureMetadataRefValid(model.SignatureMetadataRef) ||
		!point12ValBDecisionStateValueValid(model.OriginalDecisionState) ||
		!point12ValBDecisionStateValueValid(model.ReplayedDecisionState) ||
		!point12ValBCheckResultValid(model.EvidenceHashCheckResult) ||
		!point12ValBCheckResultValid(model.ManifestIntegrityCheckResult) ||
		!point12ValBCheckResultValid(model.SignatureMetadataCheckResult) ||
		!point12ValBCheckResultValid(model.CompatibilityCheckResult) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12ValEStateValid(model.CurrentState) {
		blockedReasons = append(blockedReasons, "replay_invariant_identity_or_metadata_invalid")
	}
	if !point11Val0ValidProjectionDisclaimer(dependencyReplaySemantics.ProjectionDisclaimer) {
		blockedReasons = append(blockedReasons, "replay_invariant_dependency_projection_disclaimer_invalid")
	}
	if model.ExternalAPIUsed || model.PointPassEmittedOutsideValE {
		blockedReasons = append(blockedReasons, "replay_invariant_external_api_or_point_pass_detected")
	}
	blockedReasons = append(blockedReasons, point12ValEReplaySemanticsMismatchReasons(model, dependencyReplaySemantics)...)
	if model.OriginalContextUsesCurrentPolicySilently {
		blockedReasons = append(blockedReasons, "replay_invariant_original_context_silent_current_policy")
	}
	if strings.TrimSpace(model.ReplayMode) == point12Val0ReplayModeCurrentPolicyContext && !model.CurrentPolicyContextExplicit {
		blockedReasons = append(blockedReasons, "replay_invariant_current_policy_context_not_explicit")
	}
	if strings.TrimSpace(model.ReplayMode) == point12Val0ReplayModeComparisonMode && !model.ComparisonModeDriftExplanationPresent {
		blockedReasons = append(blockedReasons, "replay_invariant_comparison_mode_drift_explanation_missing")
	}
	if strings.TrimSpace(model.ReplayResultTaxonomy) == Point12Val0ReplayResultSameDecision && !model.SameDecisionIsTaxonomyOnly {
		blockedReasons = append(blockedReasons, "replay_invariant_same_decision_overclaimed")
	}
	if strings.TrimSpace(model.ProofPackID) != strings.TrimSpace(dependency.ValA.Manifest.ProofPackID) ||
		strings.TrimSpace(model.ManifestID) != strings.TrimSpace(dependency.ValA.Manifest.ManifestID) ||
		strings.TrimSpace(model.ReplayResultID) != strings.TrimSpace(dependency.ValB.ReplayResult.ReplayResultID) ||
		strings.TrimSpace(model.ExportID) != strings.TrimSpace(dependency.ValC.ExportBundle.ExportID) ||
		strings.TrimSpace(model.OfflineBundleID) != strings.TrimSpace(dependency.ValC.OfflineBundle.OfflineBundleID) ||
		strings.TrimSpace(model.ManifestPayloadHash) != strings.TrimSpace(dependency.ValA.Manifest.ManifestPayloadHash) ||
		strings.TrimSpace(model.SignatureMetadataRef) != strings.TrimSpace(dependency.ValA.Manifest.SignatureMetadataRef) {
		blockedReasons = append(blockedReasons, "replay_invariant_binding_mismatch")
	}
	if dependencyReplaySemantics.BlockedReplay {
		blockedReasons = append(blockedReasons, "replay_invariant_dependency_blocked_replay")
	}
	if dependencyReplaySemantics.InsufficientEvidence || model.RedactedDecisiveEvidence || model.MissingDecisiveEvidence {
		blockedReasons = append(blockedReasons, "replay_invariant_insufficient_or_redacted_decisive_evidence")
	}
	if dependencyReplaySemantics.DifferentDecision {
		blockedReasons = append(blockedReasons, "replay_invariant_dependency_different_decision")
	}
	if dependencyReplaySemantics.PolicyMismatch {
		blockedReasons = append(blockedReasons, "replay_invariant_dependency_policy_mismatch")
	}
	if dependencyReplaySemantics.EngineMismatch {
		blockedReasons = append(blockedReasons, "replay_invariant_dependency_engine_mismatch")
	}
	if dependencyReplaySemantics.SchemaMismatch {
		blockedReasons = append(blockedReasons, "replay_invariant_dependency_schema_mismatch")
	}
	if dependencyReplaySemantics.EvidenceMismatch {
		blockedReasons = append(blockedReasons, "replay_invariant_dependency_evidence_mismatch")
	}
	if dependencyReplaySemantics.ClaimMismatch {
		blockedReasons = append(blockedReasons, "replay_invariant_dependency_claim_mismatch")
	}
	if dependencyReplaySemantics.GovernanceMismatch {
		blockedReasons = append(blockedReasons, "replay_invariant_dependency_governance_mismatch")
	}
	if dependencyReplaySemantics.RedactionLimitations {
		blockedReasons = append(blockedReasons, "replay_invariant_dependency_redaction_limitations")
	}
	if dependencyReplaySemantics.ExternalAPIUsed {
		blockedReasons = append(blockedReasons, "replay_invariant_dependency_external_api_used")
	}
	if dependencyReplaySemantics.PointPassEmitted {
		blockedReasons = append(blockedReasons, "replay_invariant_dependency_point_pass_emitted")
	}
	if len(dependency.ValB.ReplayResult.Mismatches) > 0 && !model.MismatchExpectedActualPresent {
		blockedReasons = append(blockedReasons, "replay_invariant_expected_actual_missing")
	}
	if strings.TrimSpace(dependency.ValB.ReplayResult.ReplayResultTaxonomy) == Point12Val0ReplayResultDifferentDecision && !model.DecisionDriftReasonsPresent {
		blockedReasons = append(blockedReasons, "replay_invariant_decision_drift_reason_missing")
	}
	if len(blockedReasons) > 0 {
		return Point12ValEStateBlocked, blockedReasons
	}
	if dependencyReplaySemantics.TamperDetected {
		return Point12ValEStateTampered, []string{"replay_invariant_tamper_detected"}
	}
	if dependencyReplaySemantics.UnsupportedVersion {
		return Point12ValEStateUnsupported, []string{"replay_invariant_unsupported_version"}
	}
	if len(reviewReasons) > 0 {
		return Point12ValEStateReviewRequired, reviewReasons
	}
	return Point12ValEStateActive, nil
}

func point12ValEEvidenceQualityMapModel(dependency Point12ValEDependencySnapshot) Point12ValEEvidenceQualityMap {
	proofChain := dependency.ValD.ProofChain
	evidenceStates := make([]string, 0, len(proofChain.EvidenceRefs))
	for range proofChain.EvidenceRefs {
		evidenceStates = append(evidenceStates, "active")
	}
	claimStates := make([]string, 0, len(proofChain.ClaimRefs))
	for range proofChain.ClaimRefs {
		claimStates = append(claimStates, "active")
	}
	governanceStates := make([]string, 0, len(proofChain.GovernanceEventRefs))
	for range proofChain.GovernanceEventRefs {
		governanceStates = append(governanceStates, "active")
	}
	return Point12ValEEvidenceQualityMap{
		EvidenceQualityMapID: "quality_map_point12_vale_001",
		ProofPackID:          proofChain.ProofPackID,
		ManifestID:           proofChain.ManifestID,
		TenantScope:          proofChain.TenantScope,
		ArtifactRef:          proofChain.ArtifactRef,
		EvidenceRefs:         append([]string{}, proofChain.EvidenceRefs...),
		EvidenceHashRefs:     append([]string{}, proofChain.EvidenceHashRefs...),
		PolicyRef:            proofChain.PolicyRef,
		PolicyVersion:        proofChain.PolicyVersion,
		PolicyHash:           proofChain.PolicyHash,
		EngineVersion:        proofChain.EngineVersion,
		EngineHash:           proofChain.EngineHash,
		SchemaVersion:        proofChain.SchemaVersion,
		SchemaHash:           proofChain.SchemaHash,
		ClaimRefs:            append([]string{}, proofChain.ClaimRefs...),
		GovernanceEventRefs:  append([]string{}, proofChain.GovernanceEventRefs...),
		EvidenceStates:       evidenceStates,
		PolicyState:          "active",
		EngineState:          "active",
		SchemaState:          "active",
		ClaimStates:          claimStates,
		GovernanceStates:     governanceStates,
		RedactionState:       dependency.ValC.RedactionImpactState,
		ReplayState:          dependency.ValB.ReplayResult.ReplayResultTaxonomy,
		ExportState:          dependency.ValC.ExportState,
		ProofChainState:      dependency.ValD.ProofChainState,
		FreshnessWindow:      dependency.ValA.Manifest.FreshnessWindow,
		QualityState:         Point12ValEStateActive,
	}
}

func point12ValEEvidenceQualityStateAndReasons(model Point12ValEEvidenceQualityMap, dependency Point12ValEDependencySnapshot) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point12ValEReviewRefValid(model.EvidenceQualityMapID, "quality_map_") ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point12Val0ArtifactRefValid(model.ArtifactRef) ||
		!point12Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point12Val0StringListValid(model.EvidenceHashRefs, point12Val0EvidenceHashRefValid) ||
		!point12Val0PolicyRefValid(model.PolicyRef) ||
		!point12Val0VersionIdentityValid(model.PolicyVersion) ||
		!point12Val0HashValid(model.PolicyHash) ||
		!point12Val0VersionIdentityValid(model.EngineVersion) ||
		!point12Val0HashValid(model.EngineHash) ||
		!point12Val0VersionIdentityValid(model.SchemaVersion) ||
		!point12Val0HashValid(model.SchemaHash) ||
		!point12Val0VersionIdentityValid(model.FreshnessWindow) ||
		!point12Val0OptionalStringListValid(model.ClaimRefs, point12Val0ClaimRefValid) ||
		!point12Val0OptionalStringListValid(model.GovernanceEventRefs, point12Val0GovernanceEventRefValid) ||
		!point12Val0OptionalStringListValid(model.StaleRefs, point12Val0EvidenceRefValid) ||
		!point12Val0OptionalStringListValid(model.RevokedRefs, point12Val0EvidenceRefValid) ||
		!point12Val0OptionalStringListValid(model.ExpiredRefs, point12Val0EvidenceRefValid) ||
		!point12Val0OptionalStringListValid(model.SupersededRefs, point12Val0EvidenceRefValid) ||
		!point12Val0OptionalStringListValid(model.UnsupportedRefs, point12Val0EvidenceRefValid) ||
		!point12Val0OptionalStringListValid(model.MissingRefs, point12Val0EvidenceRefValid) ||
		!point12Val0OptionalStringListValid(model.MalformedRefs, point11Val0IdentityValueValid) ||
		!point12Val0OptionalStringListValid(model.DuplicateRefs, point12Val0EvidenceRefValid) ||
		!point12Val0OptionalStringListValid(model.UnrelatedRefs, point12Val0EvidenceRefValid) ||
		!point12Val0OptionalStringListValid(model.CrossTenantRefs, point12Val0EvidenceRefValid) ||
		!point12Val0OptionalStringListValid(model.TamperedRefs, point12Val0EvidenceRefValid) ||
		!point12ValEStateValid(model.QualityState) {
		blockedReasons = append(blockedReasons, "evidence_quality_map_identity_or_metadata_invalid")
	}
	if !point12ValEExactRefHashPairSetMatch(model.EvidenceRefs, model.EvidenceHashRefs, dependency.ValD.ProofChain.EvidenceRefs, dependency.ValD.ProofChain.EvidenceHashRefs) ||
		!point12ValEExactRefHashPairSetMatch(model.EvidenceRefs, model.EvidenceHashRefs, dependency.ValA.Manifest.EvidenceRefs, dependency.ValA.Manifest.EvidenceHashRefs) ||
		strings.TrimSpace(model.ArtifactRef) != strings.TrimSpace(dependency.ValD.ProofChain.ArtifactRef) ||
		strings.TrimSpace(model.PolicyHash) != strings.TrimSpace(dependency.ValA.Manifest.PolicyHash) ||
		strings.TrimSpace(model.EngineHash) != strings.TrimSpace(dependency.ValA.Manifest.EngineHash) ||
		strings.TrimSpace(model.SchemaHash) != strings.TrimSpace(dependency.ValA.Manifest.SchemaHash) ||
		!point12Val0ExactStringSetMatch(model.ClaimRefs, dependency.ValD.ProofChain.ClaimRefs) ||
		!point12Val0ExactStringSetMatch(model.GovernanceEventRefs, dependency.ValD.ProofChain.GovernanceEventRefs) {
		blockedReasons = append(blockedReasons, "evidence_quality_map_binding_mismatch")
	}
	if len(model.TamperedRefs) > 0 || strings.TrimSpace(model.ReplayState) == Point12Val0ReplayResultTamperDetected {
		return Point12ValEStateTampered, append(blockedReasons, "evidence_quality_tampered")
	}
	if len(model.UnsupportedRefs) > 0 || strings.TrimSpace(model.SchemaState) == Point12ValEStateUnsupported || strings.TrimSpace(model.ReplayState) == Point12Val0ReplayResultUnsupportedVersion {
		return Point12ValEStateUnsupported, append(blockedReasons, "evidence_quality_unsupported")
	}
	if len(model.MissingRefs) > 0 {
		return Point12ValEStateIncomplete, append(blockedReasons, "evidence_quality_missing_decisive_refs")
	}
	if len(model.RevokedRefs) > 0 || len(model.ExpiredRefs) > 0 || len(model.DuplicateRefs) > 0 || len(model.UnrelatedRefs) > 0 || len(model.CrossTenantRefs) > 0 || len(model.MalformedRefs) > 0 {
		blockedReasons = append(blockedReasons, "evidence_quality_decisive_ref_invalid")
	}
	if len(model.StaleRefs) > 0 || len(model.SupersededRefs) > 0 {
		reviewReasons = append(reviewReasons, "evidence_quality_stale_or_superseded")
	}
	if strings.TrimSpace(model.RedactionState) == Point12ValCRedactionImpactInsufficient ||
		strings.TrimSpace(model.RedactionState) == Point12ValCRedactionImpactBlockedReplay ||
		strings.TrimSpace(model.ReplayState) == Point12Val0ReplayResultInsufficientEvidence {
		blockedReasons = append(blockedReasons, "evidence_quality_replay_or_redaction_insufficient")
	}
	if strings.TrimSpace(model.ExportState) != Point12ValCExportStateReady || strings.TrimSpace(model.ProofChainState) != Point12ValDProofChainStateActive {
		blockedReasons = append(blockedReasons, "evidence_quality_projection_context_not_active")
	}
	if len(blockedReasons) > 0 {
		return Point12ValEStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValEStateReviewRequired, reviewReasons
	}
	return Point12ValEStateActive, nil
}

func point12ValERequiredValDBindingFields() []string {
	return []string{
		"export_id",
		"redaction_manifest_id",
		"tenant_scope",
		"artifact_hash",
		"evidence_hash_refs",
		"policy_hash",
		"engine_hash",
		"schema_hash",
		"manifest_payload_hash",
		"proof_pack_id",
		"manifest_id",
		"replay_result_id",
		"projection_disclaimer",
		"source_to_evidence_from_ref",
	}
}

func point12ValEBindingMutationClosureModel(dependency Point12ValEDependencySnapshot) Point12ValEBindingMutationClosureReview {
	return Point12ValEBindingMutationClosureReview{
		ReviewID:                        "binding_mutation_review_point12_vale_001",
		ValAManifestBindingState:        Point12ValEStateActive,
		ValBReplayBindingState:          Point12ValEStateActive,
		ValCExportOfflineRedactionState: Point12ValEStateActive,
		ValDProofChainBindingState:      Point12ValEStateActive,
		RequiredValDBindingFields:       point12ValERequiredValDBindingFields(),
		CurrentState:                    Point12ValEStateActive,
		Diagnostics:                     []string{"binding_matrix_and_mutation_closure_active"},
		ProjectionDisclaimer:            dependency.ProjectionDisclaimer,
	}
}

func point12ValEBindingFieldExists(fields []Point12ValDBindingMatrixField, fieldName string) bool {
	for _, field := range fields {
		if strings.TrimSpace(field.FieldName) != strings.TrimSpace(fieldName) {
			continue
		}
		switch strings.TrimSpace(field.BindingClass) {
		case point12ValDBindingClassExactRequired:
			if field.ValidationRequired && field.MutationTestRequired && strings.TrimSpace(field.UpstreamSource) != "" {
				return true
			}
		case point12ValDBindingClassIntentionallyNotBound:
			if strings.TrimSpace(field.Reason) != "" {
				return true
			}
		}
	}
	return false
}

func point12ValEIntentionallyNotBoundReasonsPresent(fields []Point12ValDBindingMatrixField) bool {
	for _, field := range fields {
		if strings.TrimSpace(field.BindingClass) == point12ValDBindingClassIntentionallyNotBound && strings.TrimSpace(field.Reason) == "" {
			return false
		}
	}
	return true
}

func point12ValEBindingMutationStateAndReasons(model Point12ValEBindingMutationClosureReview, dependency Point12ValEDependencySnapshot) (string, []string) {
	blockedReasons := []string{}
	resolvedValAManifestState := EvaluatePoint12ValAManifestIntegrityState(dependency.ValA.Manifest, dependency.ValA.Dependency)
	resolvedValDBindingMatrixState := EvaluatePoint12ValDBindingMatrixState(dependency.ValD.BindingMatrix)
	resolvedValDProofChainState := EvaluatePoint12ValDProofChainProjectionState(dependency.ValD.ProofChain, dependency.ValD.Dependency)
	resolvedValDQueryState := EvaluatePoint12ValDProofChainQueryState(dependency.ValD.Query, dependency.ValD.ProofChain, dependency.ValD.Dependency)
	resolvedValDExplanationState := EvaluatePoint12ValDExplanationState(dependency.ValD.Explanation, dependency.ValD.Query, dependency.ValD.ProofChain, dependency.ValD.Dependency)
	resolvedValDSupportProfileState := EvaluatePoint12ValDSupportProfileState(dependency.ValD.SupportProfile, dependency.ValD.ProofChain)
	resolvedValDPortalCompatibilityState := EvaluatePoint12ValDPortalCompatibilityState(dependency.ValD.PortalCompatibility, dependency.ValD.ProofChain)
	if !point12ValEReviewRefValid(model.ReviewID, "binding_mutation_review_") ||
		!point12Val0OptionalStringListValid(model.RequiredValDBindingFields, point11Val0IdentityValueValid) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12ValEStateValid(model.CurrentState) ||
		!point12ValEStateValid(model.ValAManifestBindingState) ||
		!point12ValEStateValid(model.ValBReplayBindingState) ||
		!point12ValEStateValid(model.ValCExportOfflineRedactionState) ||
		!point12ValEStateValid(model.ValDProofChainBindingState) {
		blockedReasons = append(blockedReasons, "binding_mutation_identity_or_metadata_invalid")
	}
	if dependency.ValA.ManifestIntegrityState != Point12ValAManifestIntegrityStateActive ||
		resolvedValAManifestState != Point12ValAManifestIntegrityStateActive ||
		dependency.ValA.Manifest.SchemaHash != dependency.ValA.Dependency.Val0Manifest.SchemaHash ||
		dependency.ValA.Manifest.ManifestPayloadHash != point12ValAComputedManifestPayloadHash(dependency.ValA.Manifest) ||
		dependency.ValA.Manifest.SignatureBoundManifestPayloadHash != dependency.ValA.Manifest.ManifestPayloadHash {
		blockedReasons = append(blockedReasons, "binding_mutation_vala_schema_or_payload_drift")
	}
	if strings.TrimSpace(dependency.ValB.ReplayRequest.ProofPackID) != strings.TrimSpace(dependency.ValA.Manifest.ProofPackID) ||
		strings.TrimSpace(dependency.ValB.ReplayRequest.ManifestID) != strings.TrimSpace(dependency.ValA.Manifest.ManifestID) ||
		strings.TrimSpace(dependency.ValB.ReplayRequest.ArtifactRef) != strings.TrimSpace(dependency.ValA.Manifest.ArtifactRef) ||
		strings.TrimSpace(dependency.ValB.ReplayRequest.ArtifactHash) != strings.TrimSpace(dependency.ValA.Manifest.ArtifactHash) ||
		strings.TrimSpace(dependency.ValB.ReplayRequest.PolicyHash) != strings.TrimSpace(dependency.ValA.Manifest.PolicyHash) ||
		strings.TrimSpace(dependency.ValB.ReplayRequest.EngineHash) != strings.TrimSpace(dependency.ValA.Manifest.EngineHash) ||
		strings.TrimSpace(dependency.ValB.ReplayRequest.SchemaHash) != strings.TrimSpace(dependency.ValA.Manifest.SchemaHash) ||
		strings.TrimSpace(dependency.ValB.ReplayRequest.CompatibilityProfileRef) != strings.TrimSpace(dependency.ValA.Manifest.CompatibilityProfileRef) ||
		strings.TrimSpace(dependency.ValB.ReplayRequest.ManifestPayloadHash) != strings.TrimSpace(dependency.ValA.Manifest.ManifestPayloadHash) {
		blockedReasons = append(blockedReasons, "binding_mutation_valb_request_manifest_binding_invalid")
	}
	if strings.TrimSpace(dependency.ValB.ReplayResult.ReplayRequestID) != strings.TrimSpace(dependency.ValB.ReplayRequest.ReplayRequestID) ||
		strings.TrimSpace(dependency.ValB.ReplayResult.ProofPackID) != strings.TrimSpace(dependency.ValB.ReplayRequest.ProofPackID) ||
		strings.TrimSpace(dependency.ValB.ReplayResult.ManifestID) != strings.TrimSpace(dependency.ValB.ReplayRequest.ManifestID) ||
		strings.TrimSpace(dependency.ValB.ReplayResult.ReplayMode) != strings.TrimSpace(dependency.ValB.ReplayRequest.ReplayMode) ||
		dependency.ValB.ReplayResult.PointPassEmitted {
		blockedReasons = append(blockedReasons, "binding_mutation_valb_result_request_binding_invalid")
	}
	if strings.TrimSpace(dependency.ValC.ExportBundle.ProofPackID) != strings.TrimSpace(dependency.ValB.ReplayRequest.ProofPackID) ||
		strings.TrimSpace(dependency.ValC.ExportBundle.ManifestID) != strings.TrimSpace(dependency.ValB.ReplayRequest.ManifestID) ||
		strings.TrimSpace(dependency.ValC.ExportBundle.ReplayResultID) != strings.TrimSpace(dependency.ValB.ReplayResult.ReplayResultID) ||
		strings.TrimSpace(dependency.ValC.ExportBundle.ArtifactRef) != strings.TrimSpace(dependency.ValB.ReplayRequest.ArtifactRef) ||
		strings.TrimSpace(dependency.ValC.ExportBundle.ArtifactHash) != strings.TrimSpace(dependency.ValB.ReplayRequest.ArtifactHash) ||
		strings.TrimSpace(dependency.ValC.ExportBundle.ManifestPayloadHash) != strings.TrimSpace(dependency.ValB.ReplayRequest.ManifestPayloadHash) ||
		strings.TrimSpace(dependency.ValC.ExportBundle.SignatureMetadataRef) != strings.TrimSpace(dependency.ValA.Manifest.SignatureMetadataRef) ||
		strings.TrimSpace(dependency.ValC.ExportBundle.CompatibilityProfileRef) != strings.TrimSpace(dependency.ValB.ReplayRequest.CompatibilityProfileRef) ||
		strings.TrimSpace(dependency.ValC.ExportBundle.RedactionManifestRef) != strings.TrimSpace(dependency.ValB.ReplayRequest.RedactionManifestRef) ||
		strings.TrimSpace(dependency.ValC.ExportBundle.OfflineBundleRef) != strings.TrimSpace(dependency.ValC.OfflineBundle.OfflineBundleID) {
		blockedReasons = append(blockedReasons, "binding_mutation_valc_export_binding_invalid")
	}
	if strings.TrimSpace(dependency.ValC.OfflineBundle.ProofPackID) != strings.TrimSpace(dependency.ValB.ReplayRequest.ProofPackID) ||
		strings.TrimSpace(dependency.ValC.OfflineBundle.ManifestID) != strings.TrimSpace(dependency.ValB.ReplayRequest.ManifestID) ||
		strings.TrimSpace(dependency.ValC.OfflineBundle.ReplayRequestID) != strings.TrimSpace(dependency.ValB.ReplayRequest.ReplayRequestID) ||
		strings.TrimSpace(dependency.ValC.OfflineBundle.ReplayResultID) != strings.TrimSpace(dependency.ValB.ReplayResult.ReplayResultID) ||
		strings.TrimSpace(dependency.ValC.OfflineBundle.ArtifactRef) != strings.TrimSpace(dependency.ValB.ReplayRequest.ArtifactRef) ||
		strings.TrimSpace(dependency.ValC.OfflineBundle.ArtifactHash) != strings.TrimSpace(dependency.ValB.ReplayRequest.ArtifactHash) ||
		strings.TrimSpace(dependency.ValC.OfflineBundle.ManifestPayloadHash) != strings.TrimSpace(dependency.ValB.ReplayRequest.ManifestPayloadHash) ||
		strings.TrimSpace(dependency.ValC.OfflineBundle.SignatureMetadataRef) != strings.TrimSpace(dependency.ValA.Manifest.SignatureMetadataRef) ||
		strings.TrimSpace(dependency.ValC.OfflineBundle.CompatibilityProfileRef) != strings.TrimSpace(dependency.ValB.ReplayRequest.CompatibilityProfileRef) ||
		strings.TrimSpace(dependency.ValC.OfflineBundle.RedactionManifestRef) != strings.TrimSpace(dependency.ValB.ReplayRequest.RedactionManifestRef) {
		blockedReasons = append(blockedReasons, "binding_mutation_valc_offline_binding_invalid")
	}
	if dependency.ValC.RedactionManifest.RedactionManifestID != dependency.ValB.ReplayRequest.RedactionManifestRef ||
		dependency.ValC.RedactionImpactVerdict.RedactionManifestID != dependency.ValC.RedactionManifest.RedactionManifestID ||
		dependency.ValC.RedactionManifest.RedactionManifestID != dependency.ValD.ProofChain.RedactionManifestID {
		blockedReasons = append(blockedReasons, "binding_mutation_valc_redaction_identity_invalid")
	}
	if strings.TrimSpace(dependency.ValC.PublicPrivateBoundary.ExportID) != strings.TrimSpace(dependency.ValC.ExportBundle.ExportID) ||
		dependency.ValC.PublicPrivateBoundary.OfflineBundleID != dependency.ValC.OfflineBundle.OfflineBundleID {
		blockedReasons = append(blockedReasons, "binding_mutation_valc_boundary_binding_invalid")
	}
	if dependency.ValD.BindingMatrixState != Point12ValDBindingMatrixStateActive ||
		resolvedValDBindingMatrixState != Point12ValDBindingMatrixStateActive ||
		dependency.ValD.ProofChainState != Point12ValDProofChainStateActive ||
		resolvedValDProofChainState != Point12ValDProofChainStateActive ||
		dependency.ValD.QueryState != Point12ValDQueryStateActive ||
		resolvedValDQueryState != Point12ValDQueryStateActive ||
		dependency.ValD.ExplanationState != Point12ValDExplanationStateActive ||
		resolvedValDExplanationState != Point12ValDExplanationStateActive ||
		dependency.ValD.SupportProfileState != Point12ValDSupportProfileStateActive ||
		resolvedValDSupportProfileState != Point12ValDSupportProfileStateActive ||
		dependency.ValD.PortalCompatibilityState != Point12ValDPortalCompatibilityStateActive ||
		resolvedValDPortalCompatibilityState != Point12ValDPortalCompatibilityStateActive ||
		dependency.ValD.ProofChain.ProjectionHash != point12ValDComputedProjectionHash(dependency.ValD.ProofChain) {
		blockedReasons = append(blockedReasons, "binding_mutation_vald_binding_invalid")
	}
	for _, fieldName := range model.RequiredValDBindingFields {
		if !point12ValEBindingFieldExists(dependency.ValD.BindingMatrix.BoundFields, fieldName) {
			blockedReasons = append(blockedReasons, "binding_mutation_required_vald_field_missing:"+fieldName)
			break
		}
	}
	if !point12ValEIntentionallyNotBoundReasonsPresent(dependency.ValD.BindingMatrix.BoundFields) {
		blockedReasons = append(blockedReasons, "binding_mutation_intentionally_not_bound_reason_missing")
	}
	if model.AdvisoryOnlyAffectsAuthority {
		blockedReasons = append(blockedReasons, "binding_mutation_advisory_only_affects_authority")
	}
	if len(blockedReasons) > 0 {
		return Point12ValEStateBlocked, blockedReasons
	}
	return Point12ValEStateActive, nil
}

func point12ValEProjectionBoundaryModel(dependency Point12ValEDependencySnapshot) Point12ValEProjectionBoundaryResult {
	agentAdvisoryOnly := true
	for _, edge := range dependency.ValD.ProofChain.LineageEdges {
		if strings.TrimSpace(edge.EdgeType) == point12ValDLineageEdgeTypeAgentFindingAdvisory && (!edge.AdvisoryOnly || edge.ClaimsCertification || edge.ClaimsSourceOfTruth || edge.EmitsPrematurePass) {
			agentAdvisoryOnly = false
			break
		}
	}
	return Point12ValEProjectionBoundaryResult{
		BoundaryResultID:                     "projection_boundary_point12_vale_001",
		ProofChainProjectionNotSourceOfTruth: dependency.ValD.ProofChain.AdvisoryOnly,
		ExportBoundedProjection:              dependency.ValC.ExportBundle.AdvisoryOnly,
		OfflineBoundedVerificationPackage:    dependency.ValC.OfflineBundle.NoExternalAPIRequired && !dependency.ValC.OfflineBundle.ExternalAPIUsed,
		FinancialInsuranceAuditMetadataOnly:  dependency.ValD.SupportProfile.AdvisoryOnly && dependency.ValD.SupportProfile.NoFinancialGuarantee && dependency.ValD.SupportProfile.NoComplianceGuarantee,
		PortalCompatibilityModelOnly:         dependency.ValD.PortalCompatibility.ReadOnly && dependency.ValD.PortalCompatibility.NotesAnnotationOnly,
		AgentFindingAdvisoryOnly:             agentAdvisoryOnly,
		CurrentState:                         Point12ValEStateActive,
		Diagnostics:                          []string{"projection_boundaries_remain_advisory_only"},
		ProjectionDisclaimer:                 dependency.ProjectionDisclaimer,
	}
}

func point12ValEProjectionBoundaryStateAndReasons(model Point12ValEProjectionBoundaryResult) (string, []string) {
	reasons := []string{}
	if !point12ValEReviewRefValid(model.BoundaryResultID, "projection_boundary_") ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12ValEStateValid(model.CurrentState) {
		reasons = append(reasons, "projection_boundary_identity_or_metadata_invalid")
	}
	if !model.ProofChainProjectionNotSourceOfTruth ||
		!model.ExportBoundedProjection ||
		!model.OfflineBoundedVerificationPackage ||
		!model.FinancialInsuranceAuditMetadataOnly ||
		!model.PortalCompatibilityModelOnly ||
		!model.AgentFindingAdvisoryOnly ||
		model.BuyerProductCustomerTextCreatesAuthority ||
		model.AuditorNotesMutateOutcome ||
		model.MutatesCanonicalEvidenceSpine ||
		model.EmitsPoint12Pass {
		reasons = append(reasons, "projection_boundary_guardrail_failed")
	}
	if len(reasons) > 0 {
		return Point12ValEStateBlocked, reasons
	}
	return Point12ValEStateActive, nil
}

func point12ValENoOverclaimReviewModel(dependency Point12ValEDependencySnapshot) Point12ValENoOverclaimReview {
	return Point12ValENoOverclaimReview{
		ReviewID:                     "no_overclaim_review_point12_vale_001",
		ObservedCustomerTexts:        []string{dependency.ValD.Explanation.CustomerVisibleStatement},
		ObservedExportTexts:          []string{dependency.ValC.ExportBundle.CustomerVisibleSummary, dependency.ValC.OfflineBundle.CustomerVisibleExplanation},
		ObservedSupportTexts:         []string{dependency.ValD.SupportProfile.SupportStatement},
		ObservedPortalTexts:          []string{"read_only_projection", "notes_annotation_only"},
		BlockedClaimLedger:           append(append([]string{}, dependency.ValD.SupportProfile.BlockedWordingRefs...), dependency.ValC.RedactionManifest.DisallowedClaimsAfterRedaction...),
		BlockedClaimLedgerClassified: true,
		InternalDiagnosticTexts:      []string{dependency.ValD.Explanation.InternalDiagnosticSummary, dependency.ValD.SupportProfile.InternalDiagnosticSummary, dependency.ValC.RedactionManifest.RedactionSummary},
		GrepRefs:                     []string{"grep_run_point12_vale_no_overclaim_001"},
		CurrentState:                 Point12ValEStateActive,
		Diagnostics:                  []string{"forbidden_wording_blocked_outside_internal_and_ledger_context"},
		ProjectionDisclaimer:         dependency.ProjectionDisclaimer,
	}
}

func point12ValENoOverclaimStateAndReasons(model Point12ValENoOverclaimReview) (string, []string) {
	reasons := []string{}
	if !point12ValEReviewRefValid(model.ReviewID, "no_overclaim_review_") ||
		!point12ValEOptionalTextListValid(model.ObservedCustomerTexts) ||
		!point12ValEOptionalTextListValid(model.ObservedExportTexts) ||
		!point12ValEOptionalTextListValid(model.ObservedSupportTexts) ||
		!point12ValEOptionalTextListValid(model.ObservedPortalTexts) ||
		!point12ValEOptionalTextListValid(model.BlockedClaimLedger) ||
		!point12ValEOptionalTextListValid(model.InternalDiagnosticTexts) ||
		!point12Val0StringListValid(model.GrepRefs, point12ValEGrepRunRefValid) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12ValEStateValid(model.CurrentState) {
		reasons = append(reasons, "no_overclaim_identity_or_metadata_invalid")
	}
	if point12ValEContainsForbiddenClaim(model.ObservedCustomerTexts...) ||
		point12ValEContainsForbiddenClaim(model.ObservedExportTexts...) ||
		point12ValEContainsForbiddenClaim(model.ObservedSupportTexts...) ||
		point12ValEContainsForbiddenClaim(model.ObservedPortalTexts...) {
		reasons = append(reasons, "no_overclaim_customer_or_export_overclaim_detected")
	}
	if point12ValEContainsForbiddenClaim(model.BlockedClaimLedger...) && !model.BlockedClaimLedgerClassified {
		reasons = append(reasons, "no_overclaim_blocked_ledger_unclassified")
	}
	for _, value := range model.InternalDiagnosticTexts {
		if point12ValEContainsForbiddenClaim(value) && !point12ValEInternalDiagnosticContextAllowed(value) {
			reasons = append(reasons, "no_overclaim_internal_diagnostic_unclassified")
			break
		}
	}
	if len(reasons) > 0 {
		return Point12ValEStateBlocked, reasons
	}
	return Point12ValEStateActive, nil
}

func point12ValECleanRoomIPReviewModel(dependency Point12ValEDependencySnapshot) Point12ValECleanRoomIPReview {
	return Point12ValECleanRoomIPReview{
		ReviewID:             "clean_room_review_point12_vale_001",
		ThirdPartyRefs:       []string{"third_party_point12_vale_001"},
		LicenseReviewRefs:    []string{"license_review_point12_vale_001"},
		IPReviewRefs:         []string{"ip_review_point12_vale_001"},
		AIReviewPackageRefs:  []string{"ip_review_point12_vale_codex_review_package_001"},
		CurrentState:         Point12ValEStateActive,
		Diagnostics:          []string{"clean_room_ip_guardrail_active_not_legal_opinion"},
		ProjectionDisclaimer: dependency.ProjectionDisclaimer,
	}
}

func point12ValECleanRoomIPStateAndReasons(model Point12ValECleanRoomIPReview) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point12ValEReviewRefValid(model.ReviewID, "clean_room_review_") ||
		!point12Val0StringListValid(model.ThirdPartyRefs, func(value string) bool { return point12ValEReviewRefValid(value, "third_party_") }) ||
		!point12Val0StringListValid(model.LicenseReviewRefs, func(value string) bool { return point12ValEReviewRefValid(value, "license_review_") }) ||
		!point12Val0StringListValid(model.IPReviewRefs, func(value string) bool { return point12ValEReviewRefValid(value, "ip_review_") }) ||
		!point12Val0StringListValid(model.AIReviewPackageRefs, func(value string) bool { return point12ValEReviewRefValid(value, "ip_review_", "review_package_") }) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12ValEStateValid(model.CurrentState) {
		blockedReasons = append(blockedReasons, "clean_room_ip_identity_or_metadata_invalid")
	}
	if model.CompetitorCopyDetected ||
		model.ProprietaryWorkflowClaimDetected ||
		model.ReverseEngineeringClaimDetected ||
		model.PatentClearanceClaimDetected ||
		model.LegalCertificationClaimDetected {
		blockedReasons = append(blockedReasons, "clean_room_ip_policy_violation")
	}
	if model.UnreviewedCustomerFacingDependency {
		reviewReasons = append(reviewReasons, "clean_room_ip_unreviewed_customer_facing_dependency")
	}
	if len(blockedReasons) > 0 {
		return Point12ValEStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValEStateReviewRequired, reviewReasons
	}
	return Point12ValEStateActive, nil
}

func point12ValERetentionProvenanceReviewModel(dependency Point12ValEDependencySnapshot) Point12ValERetentionProvenanceReview {
	agentAdvisoryOnly := true
	for _, lineage := range dependency.Val0.ProvenanceProfile.AgentLineages {
		if !lineage.LineageInputOnly || lineage.ClaimsCertification || lineage.ClaimsSourceOfTruth || lineage.EmitsPrematurePass {
			agentAdvisoryOnly = false
			break
		}
	}
	for _, edge := range dependency.ValD.ProofChain.LineageEdges {
		if strings.TrimSpace(edge.EdgeType) == point12ValDLineageEdgeTypeAgentFindingAdvisory && (!edge.AdvisoryOnly || edge.ClaimsCertification || edge.ClaimsSourceOfTruth || edge.EmitsPrematurePass) {
			agentAdvisoryOnly = false
			break
		}
	}
	return Point12ValERetentionProvenanceReview{
		ReviewID:                           "retention_review_point12_vale_001",
		ProofPackRetentionClassRef:         dependency.ValA.Manifest.RetentionClassRef,
		ExportRetentionClassRef:            dependency.ValC.ExportBundle.RetentionClassRef,
		OfflineBundleRetentionClassRef:     dependency.ValC.OfflineBundle.RetentionClassRef,
		RedactionManifestRetentionClassRef: dependency.ValC.RedactionManifest.RetentionClassRef,
		AuditRetentionClassRef:             dependency.ValC.ExportBundle.RetentionClassRef,
		RetentionOwnerRef:                  dependency.ValC.ExportBundle.RetentionOwnerRef,
		DisposalPathRef:                    dependency.ValC.ExportBundle.DisposalPathRef,
		TenantScope:                        dependency.ValD.ProofChain.TenantScope,
		PublicPrivateClassification:        dependency.ValC.ExportBundle.PublicPrivateClassification,
		ToolchainProvenanceRefs:            append([]string{}, dependency.ValA.Manifest.ToolchainProvenanceRefs...),
		AgentLineageRefs:                   append([]string{}, dependency.ValA.Manifest.AgentLineageRefs...),
		AgentLineageAdvisoryOnly:           agentAdvisoryOnly,
		CurrentState:                       Point12ValEStateActive,
		Diagnostics:                        []string{"retention_and_provenance_requirements_complete"},
		ProjectionDisclaimer:               dependency.ProjectionDisclaimer,
	}
}

func point12ValERetentionProvenanceStateAndReasons(model Point12ValERetentionProvenanceReview, dependency Point12ValEDependencySnapshot) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point12ValEReviewRefValid(model.ReviewID, "retention_review_") ||
		!point12Val0RetentionClassRefValid(model.ProofPackRetentionClassRef) ||
		!point12Val0RetentionClassRefValid(model.ExportRetentionClassRef) ||
		!point12Val0RetentionClassRefValid(model.OfflineBundleRetentionClassRef) ||
		!point12Val0RetentionClassRefValid(model.RedactionManifestRetentionClassRef) ||
		!point12Val0RetentionClassRefValid(model.AuditRetentionClassRef) ||
		!point12ValCRetentionOwnerRefValid(model.RetentionOwnerRef) ||
		!point12ValCDisposalPathRefValid(model.DisposalPathRef) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point12ValCPublicPrivateClassificationValid(model.PublicPrivateClassification) ||
		!point12Val0OptionalStringListValid(model.ToolchainProvenanceRefs, point12Val0ToolchainProvenanceRefValid) ||
		!point12Val0OptionalStringListValid(model.AgentLineageRefs, point12Val0AgentLineageRefValid) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12ValEStateValid(model.CurrentState) {
		blockedReasons = append(blockedReasons, "retention_provenance_identity_or_metadata_invalid")
	}
	if strings.TrimSpace(model.ProofPackRetentionClassRef) != strings.TrimSpace(dependency.ValA.Manifest.RetentionClassRef) ||
		strings.TrimSpace(model.ExportRetentionClassRef) != strings.TrimSpace(dependency.ValC.ExportBundle.RetentionClassRef) ||
		strings.TrimSpace(model.OfflineBundleRetentionClassRef) != strings.TrimSpace(dependency.ValC.OfflineBundle.RetentionClassRef) ||
		strings.TrimSpace(model.RedactionManifestRetentionClassRef) != strings.TrimSpace(dependency.ValC.RedactionManifest.RetentionClassRef) ||
		strings.TrimSpace(model.RetentionOwnerRef) != strings.TrimSpace(dependency.ValC.ExportBundle.RetentionOwnerRef) ||
		strings.TrimSpace(model.DisposalPathRef) != strings.TrimSpace(dependency.ValC.ExportBundle.DisposalPathRef) {
		blockedReasons = append(blockedReasons, "retention_provenance_binding_mismatch")
	}
	if dependency.Val0.ProvenanceProfile.DecisiveToolchainProvenanceRequired && len(model.ToolchainProvenanceRefs) == 0 {
		reviewReasons = append(reviewReasons, "retention_provenance_toolchain_missing")
	}
	if !model.AgentLineageAdvisoryOnly || model.SupportPilotArtifactPromotedWithoutGovernance {
		blockedReasons = append(blockedReasons, "retention_provenance_agent_or_support_boundary_invalid")
	}
	if len(blockedReasons) > 0 {
		return Point12ValEStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValEStateReviewRequired, reviewReasons
	}
	return Point12ValEStateActive, nil
}

func point12ValEPassClosureManifestModel(dependency Point12ValEDependencySnapshot) Point12ValEPassClosureManifest {
	return Point12ValEPassClosureManifest{
		CurrentState:                Point12ValEStateActive,
		ClosureManifestID:           "closure_manifest_point12_vale_001",
		PointID:                     point12Val0PointID,
		WaveID:                      point12ValEWaveID,
		Scope:                       point12ValEScope,
		DependencyGateResult:        Point12ValEStateActive,
		Val0SnapshotRef:             dependency.Val0SnapshotRef,
		ValASnapshotRef:             dependency.ValASnapshotRef,
		ValBSnapshotRef:             dependency.ValBSnapshotRef,
		ValCSnapshotRef:             dependency.ValCSnapshotRef,
		ValDSnapshotRef:             dependency.ValDSnapshotRef,
		ProofPackID:                 dependency.ValD.ProofChain.ProofPackID,
		ManifestID:                  dependency.ValD.ProofChain.ManifestID,
		ReplayResultID:              dependency.ValD.ProofChain.ReplayResultID,
		ExportID:                    dependency.ValD.ProofChain.ExportID,
		OfflineBundleID:             dependency.ValD.ProofChain.OfflineBundleID,
		RedactionManifestID:         dependency.ValD.ProofChain.RedactionManifestID,
		ProofChainID:                dependency.ValD.ProofChain.ProofChainID,
		TenantScope:                 dependency.ValD.ProofChain.TenantScope,
		ArtifactRef:                 dependency.ValD.ProofChain.ArtifactRef,
		ArtifactHash:                dependency.ValD.ProofChain.ArtifactHash,
		EvidenceRefs:                append([]string{}, dependency.ValD.ProofChain.EvidenceRefs...),
		EvidenceHashRefs:            append([]string{}, dependency.ValD.ProofChain.EvidenceHashRefs...),
		PolicyRef:                   dependency.ValD.ProofChain.PolicyRef,
		PolicyVersion:               dependency.ValD.ProofChain.PolicyVersion,
		PolicyHash:                  dependency.ValD.ProofChain.PolicyHash,
		EngineVersion:               dependency.ValD.ProofChain.EngineVersion,
		EngineHash:                  dependency.ValD.ProofChain.EngineHash,
		SchemaVersion:               dependency.ValD.ProofChain.SchemaVersion,
		SchemaHash:                  dependency.ValD.ProofChain.SchemaHash,
		ClaimRefs:                   append([]string{}, dependency.ValD.ProofChain.ClaimRefs...),
		GovernanceEventRefs:         append([]string{}, dependency.ValD.ProofChain.GovernanceEventRefs...),
		CompatibilityProfileRef:     dependency.ValD.ProofChain.CompatibilityProfileRef,
		ManifestPayloadHash:         dependency.ValA.Manifest.ManifestPayloadHash,
		SignatureMetadataRef:        dependency.ValA.Manifest.SignatureMetadataRef,
		RetentionClassRef:           dependency.ValC.ExportBundle.RetentionClassRef,
		RetentionOwnerRef:           dependency.ValC.ExportBundle.RetentionOwnerRef,
		DisposalPathRef:             dependency.ValC.ExportBundle.DisposalPathRef,
		PublicPrivateClassification: dependency.ValC.ExportBundle.PublicPrivateClassification,
		ToolchainProvenanceRefs:     append([]string{}, dependency.ValA.Manifest.ToolchainProvenanceRefs...),
		AgentLineageRefs:            append([]string{}, dependency.ValA.Manifest.AgentLineageRefs...),
		CommandsRun: []string{
			"command_run_point12_vale_gofmt_001",
			"command_run_point12_vale_go_test_formal_001",
			"command_run_point12_vale_go_test_all_001",
		},
		TestsRun: []string{
			"test_run_point12_vale_internal_formal_001",
			"test_run_point12_vale_point11_regressions_001",
			"test_run_point12_vale_go_test_all_001",
		},
		NegativeFixturesRun: []string{
			"negative_fixture_point12_vale_dependency_gate_001",
			"negative_fixture_point12_vale_no_overclaim_001",
			"negative_fixture_point12_vale_binding_mutation_001",
		},
		BindingMatrixResult:         Point12ValEStateActive,
		MutationTestResult:          Point12ValEStateActive,
		ReplayInvariantResult:       Point12ValEStateActive,
		ExportOfflineBoundaryResult: Point12ValEStateActive,
		RedactionBoundaryResult:     Point12ValEStateActive,
		ProofChainProjectionResult:  Point12ValEStateActive,
		EvidenceQualityMapResult:    Point12ValEStateActive,
		ProjectionBoundaryResult:    Point12ValEStateActive,
		NoOverclaimGrepResult:       Point12ValEStateActive,
		CleanRoomIPResult:           Point12ValEStateActive,
		RetentionDisposalResult:     Point12ValEStateActive,
		ReviewerResult:              point12ValEReviewerResultPassConfirmed,
		GeneratedAt:                 "2026-05-04T11:30:00Z",
		Point12PassAllowed:          true,
		Point12PassToken:            point12ValEPoint12PassToken,
		ProjectionDisclaimer:        dependency.ProjectionDisclaimer,
	}
}

func point12ValEPassCandidate(foundation Point12ValEFoundation) bool {
	return foundation.DependencyState == Point12ValEStateActive &&
		foundation.ReplayInvariantState == Point12ValEStateActive &&
		foundation.EvidenceQualityState == Point12ValEStateActive &&
		foundation.BindingMutationState == Point12ValEStateActive &&
		foundation.ProjectionBoundaryState == Point12ValEStateActive &&
		foundation.NoOverclaimState == Point12ValEStateActive &&
		foundation.CleanRoomIPState == Point12ValEStateActive &&
		foundation.RetentionProvenanceState == Point12ValEStateActive
}

func point12ValEPassClosureManifestStateAndReasons(model Point12ValEPassClosureManifest, foundation Point12ValEFoundation, expectedPassAllowed bool) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12ValEClosureManifestRefValid(model.ClosureManifestID) ||
		strings.TrimSpace(model.PointID) != point12Val0PointID ||
		strings.TrimSpace(model.WaveID) != point12ValEWaveID ||
		strings.TrimSpace(model.Scope) != point12ValEScope ||
		!point12ValEStateValid(model.CurrentState) {
		reasons = append(reasons, "pass_closure_manifest_identity_or_metadata_invalid")
	}
	if strings.TrimSpace(model.DependencyGateResult) != Point12ValEStateActive ||
		strings.TrimSpace(model.BindingMatrixResult) != Point12ValEStateActive ||
		strings.TrimSpace(model.MutationTestResult) != Point12ValEStateActive ||
		strings.TrimSpace(model.ReplayInvariantResult) != Point12ValEStateActive ||
		strings.TrimSpace(model.ExportOfflineBoundaryResult) != Point12ValEStateActive ||
		strings.TrimSpace(model.RedactionBoundaryResult) != Point12ValEStateActive ||
		strings.TrimSpace(model.ProofChainProjectionResult) != Point12ValEStateActive ||
		strings.TrimSpace(model.EvidenceQualityMapResult) != Point12ValEStateActive ||
		strings.TrimSpace(model.ProjectionBoundaryResult) != Point12ValEStateActive ||
		strings.TrimSpace(model.NoOverclaimGrepResult) != Point12ValEStateActive ||
		strings.TrimSpace(model.CleanRoomIPResult) != Point12ValEStateActive ||
		strings.TrimSpace(model.RetentionDisposalResult) != Point12ValEStateActive {
		reasons = append(reasons, "pass_closure_manifest_result_state_invalid")
	}
	if !point12ValADependencySnapshotRefValid(model.Val0SnapshotRef) ||
		!point12ValBDependencySnapshotRefValid(model.ValASnapshotRef) ||
		!point12ValCDependencySnapshotRefValid(model.ValBSnapshotRef) ||
		!point12ValDDependencySnapshotRefValid(model.ValCSnapshotRef) ||
		!point12ValEDependencySnapshotRefValid(model.ValDSnapshotRef) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point12ValBReplayResultRefValid(model.ReplayResultID) ||
		!point12ValCExportRefValid(model.ExportID) ||
		!point12ValCOfflineBundleRefValid(model.OfflineBundleID) ||
		!point12Val0RedactionManifestRefValid(model.RedactionManifestID) ||
		!point12ValDProofChainRefValid(model.ProofChainID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point12Val0ArtifactRefValid(model.ArtifactRef) ||
		!point12Val0HashValid(model.ArtifactHash) ||
		!point12Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point12Val0StringListValid(model.EvidenceHashRefs, point12Val0EvidenceHashRefValid) ||
		!point12Val0PolicyRefValid(model.PolicyRef) ||
		!point12Val0VersionIdentityValid(model.PolicyVersion) ||
		!point12Val0HashValid(model.PolicyHash) ||
		!point12Val0VersionIdentityValid(model.EngineVersion) ||
		!point12Val0HashValid(model.EngineHash) ||
		!point12Val0VersionIdentityValid(model.SchemaVersion) ||
		!point12Val0HashValid(model.SchemaHash) ||
		!point12Val0OptionalStringListValid(model.ClaimRefs, point12Val0ClaimRefValid) ||
		!point12Val0OptionalStringListValid(model.GovernanceEventRefs, point12Val0GovernanceEventRefValid) ||
		!point12Val0CompatibilityProfileRefValid(model.CompatibilityProfileRef) ||
		!point12Val0HashValid(model.ManifestPayloadHash) ||
		!point12ValASignatureMetadataRefValid(model.SignatureMetadataRef) ||
		!point12Val0RetentionClassRefValid(model.RetentionClassRef) ||
		!point12ValCRetentionOwnerRefValid(model.RetentionOwnerRef) ||
		!point12ValCDisposalPathRefValid(model.DisposalPathRef) ||
		!point12ValCPublicPrivateClassificationValid(model.PublicPrivateClassification) ||
		!point12Val0StringListValid(model.ToolchainProvenanceRefs, point12Val0ToolchainProvenanceRefValid) ||
		!point12Val0StringListValid(model.AgentLineageRefs, point12Val0AgentLineageRefValid) ||
		!point12Val0StringListValid(model.CommandsRun, point12ValECommandRunRefValid) ||
		!point12Val0StringListValid(model.TestsRun, point12ValETestRunRefValid) ||
		!point12Val0StringListValid(model.NegativeFixturesRun, point12ValENegativeFixtureRunRefValid) ||
		!point11Val0ValidTimestamp(model.GeneratedAt) {
		reasons = append(reasons, "pass_closure_manifest_required_fields_invalid")
	}
	if strings.TrimSpace(model.Val0SnapshotRef) != strings.TrimSpace(foundation.Dependency.Val0SnapshotRef) ||
		strings.TrimSpace(model.ValASnapshotRef) != strings.TrimSpace(foundation.Dependency.ValASnapshotRef) ||
		strings.TrimSpace(model.ValBSnapshotRef) != strings.TrimSpace(foundation.Dependency.ValBSnapshotRef) ||
		strings.TrimSpace(model.ValCSnapshotRef) != strings.TrimSpace(foundation.Dependency.ValCSnapshotRef) ||
		strings.TrimSpace(model.ValDSnapshotRef) != strings.TrimSpace(foundation.Dependency.ValDSnapshotRef) ||
		strings.TrimSpace(model.ProofPackID) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.ProofPackID) ||
		strings.TrimSpace(model.ManifestID) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.ManifestID) ||
		strings.TrimSpace(model.ReplayResultID) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.ReplayResultID) ||
		strings.TrimSpace(model.ExportID) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.ExportID) ||
		strings.TrimSpace(model.OfflineBundleID) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.OfflineBundleID) ||
		strings.TrimSpace(model.RedactionManifestID) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.RedactionManifestID) ||
		strings.TrimSpace(model.ProofChainID) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.ProofChainID) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.TenantScope) ||
		strings.TrimSpace(model.ArtifactRef) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.ArtifactRef) ||
		strings.TrimSpace(model.ArtifactHash) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.ArtifactHash) ||
		!point12ValEExactRefHashPairSetMatch(model.EvidenceRefs, model.EvidenceHashRefs, foundation.Dependency.ValD.ProofChain.EvidenceRefs, foundation.Dependency.ValD.ProofChain.EvidenceHashRefs) ||
		strings.TrimSpace(model.PolicyRef) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.PolicyRef) ||
		strings.TrimSpace(model.PolicyVersion) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.PolicyVersion) ||
		strings.TrimSpace(model.PolicyHash) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.PolicyHash) ||
		strings.TrimSpace(model.EngineVersion) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.EngineVersion) ||
		strings.TrimSpace(model.EngineHash) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.EngineHash) ||
		strings.TrimSpace(model.SchemaVersion) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.SchemaVersion) ||
		strings.TrimSpace(model.SchemaHash) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.SchemaHash) ||
		!point12Val0ExactStringSetMatch(model.ClaimRefs, foundation.Dependency.ValD.ProofChain.ClaimRefs) ||
		!point12Val0ExactStringSetMatch(model.GovernanceEventRefs, foundation.Dependency.ValD.ProofChain.GovernanceEventRefs) ||
		strings.TrimSpace(model.CompatibilityProfileRef) != strings.TrimSpace(foundation.Dependency.ValD.ProofChain.CompatibilityProfileRef) ||
		strings.TrimSpace(model.ManifestPayloadHash) != strings.TrimSpace(foundation.Dependency.ValA.Manifest.ManifestPayloadHash) ||
		strings.TrimSpace(model.SignatureMetadataRef) != strings.TrimSpace(foundation.Dependency.ValA.Manifest.SignatureMetadataRef) ||
		strings.TrimSpace(model.RetentionClassRef) != strings.TrimSpace(foundation.Dependency.ValC.ExportBundle.RetentionClassRef) ||
		strings.TrimSpace(model.RetentionOwnerRef) != strings.TrimSpace(foundation.Dependency.ValC.ExportBundle.RetentionOwnerRef) ||
		strings.TrimSpace(model.DisposalPathRef) != strings.TrimSpace(foundation.Dependency.ValC.ExportBundle.DisposalPathRef) ||
		strings.TrimSpace(model.PublicPrivateClassification) != strings.TrimSpace(foundation.Dependency.ValC.ExportBundle.PublicPrivateClassification) ||
		!point12Val0ExactStringSetMatch(model.ToolchainProvenanceRefs, foundation.Dependency.ValA.Manifest.ToolchainProvenanceRefs) ||
		!point12Val0ExactStringSetMatch(model.AgentLineageRefs, foundation.Dependency.ValA.Manifest.AgentLineageRefs) {
		reasons = append(reasons, "pass_closure_manifest_binding_mismatch")
	}
	if strings.TrimSpace(model.CommitSHAIfAvailable) != "" {
		reasons = append(reasons, "pass_closure_manifest_commit_sha_not_allowed_before_commit")
	}
	if !point12ValEReviewerResultValid(model.ReviewerResult) {
		reasons = append(reasons, "pass_closure_manifest_reviewer_result_invalid")
	}
	switch strings.TrimSpace(model.ReviewerResult) {
	case point12ValEReviewerResultPassConfirmed:
		if !expectedPassAllowed || !model.Point12PassAllowed || strings.TrimSpace(model.Point12PassToken) != point12ValEPoint12PassToken {
			reasons = append(reasons, "pass_closure_manifest_pass_confirmed_not_fully_authorized")
		}
	case point12ValEReviewerResultPass:
		if model.Point12PassAllowed || strings.TrimSpace(model.Point12PassToken) != "" {
			reasons = append(reasons, "pass_closure_manifest_point12_pass_emitted_before_final_confirmation")
		}
	case point12ValEReviewerResultReviewRequired:
		if model.Point12PassAllowed || strings.TrimSpace(model.Point12PassToken) != "" {
			reasons = append(reasons, "pass_closure_manifest_review_required_emits_point12_pass")
		}
		return Point12ValEStateReviewRequired, append(reasons, "pass_closure_manifest_reviewer_requires_review")
	}
	if strings.TrimSpace(model.Point12PassToken) != "" && strings.TrimSpace(model.Point12PassToken) != point12ValEPoint12PassToken {
		reasons = append(reasons, "pass_closure_manifest_token_invalid")
	}
	if !expectedPassAllowed && strings.TrimSpace(model.Point12PassToken) != "" {
		reasons = append(reasons, "pass_closure_manifest_token_present_before_final_happy_path")
	}
	if !point12ValEPassCandidate(foundation) {
		reasons = append(reasons, "pass_closure_manifest_foundation_gates_not_active")
	}
	if len(reasons) > 0 {
		return Point12ValEStateBlocked, reasons
	}
	return Point12ValEStateActive, nil
}

func point12ValEFoundationModelFromUpstream(
	val0 Point12Val0Foundation,
	valA Point12ValAFoundation,
	valB Point12ValBFoundation,
	valC Point12ValCFoundation,
	valD Point12ValDFoundation,
) Point12ValEFoundation {
	dependency := SnapshotPoint12ValEDependencyFromComputed(val0, valA, valB, valC, valD, point12ValEDependencyReviewContextModel())
	return Point12ValEFoundation{
		CurrentState:              Point12ValEStatePassConfirmed,
		ProjectionDisclaimer:      point12ValEProjectionDisclaimerBaseline,
		DependencyState:           Point12ValEStateActive,
		ReplayInvariantState:      Point12ValEStateActive,
		EvidenceQualityState:      Point12ValEStateActive,
		BindingMutationState:      Point12ValEStateActive,
		ProjectionBoundaryState:   Point12ValEStateActive,
		NoOverclaimState:          Point12ValEStateActive,
		CleanRoomIPState:          Point12ValEStateActive,
		RetentionProvenanceState:  Point12ValEStateActive,
		PassClosureManifestState:  Point12ValEStateActive,
		Point12PassAllowed:        true,
		Point12PassToken:          point12ValEPoint12PassToken,
		Dependency:                dependency,
		ReplayInvariants:          point12ValEFinalReplayInvariantsModel(dependency),
		EvidenceQualityMap:        point12ValEEvidenceQualityMapModel(dependency),
		BindingMutationClosure:    point12ValEBindingMutationClosureModel(dependency),
		ProjectionBoundary:        point12ValEProjectionBoundaryModel(dependency),
		NoOverclaimReview:         point12ValENoOverclaimReviewModel(dependency),
		CleanRoomIPReview:         point12ValECleanRoomIPReviewModel(dependency),
		RetentionProvenanceReview: point12ValERetentionProvenanceReviewModel(dependency),
		PassClosureManifest:       point12ValEPassClosureManifestModel(dependency),
	}
}

func Point12ValEFoundationModel() Point12ValEFoundation {
	val0 := ComputePoint12Val0Foundation(Point12Val0FoundationModel())
	valA := ComputePoint12ValAFoundation(Point12ValAFoundationModel())
	valB := ComputePoint12ValBFoundation(Point12ValBFoundationModel())
	valC := ComputePoint12ValCFoundation(Point12ValCFoundationModel())
	valD := ComputePoint12ValDFoundation(Point12ValDFoundationModel())
	return point12ValEFoundationModelFromUpstream(val0, valA, valB, valC, valD)
}

func point12ValECurrentState(model Point12ValEFoundation) string {
	states := []string{
		model.DependencyState,
		model.ReplayInvariantState,
		model.EvidenceQualityState,
		model.BindingMutationState,
		model.ProjectionBoundaryState,
		model.NoOverclaimState,
		model.CleanRoomIPState,
		model.RetentionProvenanceState,
		model.PassClosureManifestState,
	}
	for _, state := range states {
		if state == Point12ValEStateTampered {
			return Point12ValEStateTampered
		}
	}
	for _, state := range states {
		if state == Point12ValEStateFailed {
			return Point12ValEStateFailed
		}
	}
	for _, state := range states {
		if state == Point12ValEStateBlocked {
			return Point12ValEStateBlocked
		}
	}
	for _, state := range states {
		if state == Point12ValEStateUnsupported {
			return Point12ValEStateUnsupported
		}
	}
	for _, state := range states {
		if state == Point12ValEStateIncomplete {
			return Point12ValEStateIncomplete
		}
	}
	for _, state := range states {
		if state == Point12ValEStateReviewRequired {
			return Point12ValEStateReviewRequired
		}
	}
	if model.Point12PassAllowed && strings.TrimSpace(model.Point12PassToken) == point12ValEPoint12PassToken && strings.TrimSpace(model.PassClosureManifest.ReviewerResult) == point12ValEReviewerResultPassConfirmed {
		return Point12ValEStatePassConfirmed
	}
	return Point12ValEStateActive
}

func point12ValEBlockingReasons(model Point12ValEFoundation) []string {
	reasons := []string{}
	componentStates := map[string]string{
		"dependency":            model.DependencyState,
		"replay_invariant":      model.ReplayInvariantState,
		"evidence_quality":      model.EvidenceQualityState,
		"binding_mutation":      model.BindingMutationState,
		"projection_boundary":   model.ProjectionBoundaryState,
		"no_overclaim":          model.NoOverclaimState,
		"clean_room_ip":         model.CleanRoomIPState,
		"retention_provenance":  model.RetentionProvenanceState,
		"pass_closure_manifest": model.PassClosureManifestState,
	}
	for name, state := range componentStates {
		switch state {
		case Point12ValEStateBlocked, Point12ValEStateTampered, Point12ValEStateUnsupported, Point12ValEStateFailed, Point12ValEStateIncomplete:
			reasons = append(reasons, name+":"+state)
		}
	}
	if model.Point12PassAllowed && strings.TrimSpace(model.Point12PassToken) != point12ValEPoint12PassToken {
		reasons = append(reasons, "point12_pass_token_invalid")
	}
	return reasons
}

func ComputePoint12ValEFoundation(model Point12ValEFoundation) Point12ValEFoundation {
	dependencyState, dependencyReasons := point12ValEDependencyStateAndReasons(model.Dependency)
	model.DependencyState = dependencyState
	replayState, replayReasons := point12ValEFinalReplayInvariantStateAndReasons(model.ReplayInvariants, model.Dependency)
	model.ReplayInvariantState = replayState
	model.ReplayInvariants.CurrentState = replayState
	qualityState, qualityReasons := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency)
	model.EvidenceQualityState = qualityState
	model.EvidenceQualityMap.QualityState = qualityState
	bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
	model.BindingMutationState = bindingState
	model.BindingMutationClosure.CurrentState = bindingState
	projectionState, _ := point12ValEProjectionBoundaryStateAndReasons(model.ProjectionBoundary)
	model.ProjectionBoundaryState = projectionState
	model.ProjectionBoundary.CurrentState = projectionState
	noOverclaimState, noOverclaimReasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
	model.NoOverclaimState = noOverclaimState
	model.NoOverclaimReview.CurrentState = noOverclaimState
	cleanRoomState, cleanRoomReasons := point12ValECleanRoomIPStateAndReasons(model.CleanRoomIPReview)
	model.CleanRoomIPState = cleanRoomState
	model.CleanRoomIPReview.CurrentState = cleanRoomState
	retentionState, retentionReasons := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency)
	model.RetentionProvenanceState = retentionState
	model.RetentionProvenanceReview.CurrentState = retentionState

	passCandidate := point12ValEPassCandidate(model)
	manifestState, manifestReasons := point12ValEPassClosureManifestStateAndReasons(model.PassClosureManifest, model, passCandidate)
	model.PassClosureManifestState = manifestState
	model.PassClosureManifest.CurrentState = manifestState
	model.Point12PassAllowed = passCandidate &&
		manifestState == Point12ValEStateActive &&
		strings.TrimSpace(model.PassClosureManifest.ReviewerResult) == point12ValEReviewerResultPassConfirmed &&
		model.PassClosureManifest.Point12PassAllowed &&
		strings.TrimSpace(model.PassClosureManifest.Point12PassToken) == point12ValEPoint12PassToken
	if model.Point12PassAllowed {
		model.Point12PassToken = point12ValEPoint12PassToken
	} else {
		model.Point12PassToken = ""
	}
	model.CurrentState = point12ValECurrentState(model)
	model.BlockingReasons = point12ValEBlockingReasons(model)
	model.ReviewPrerequisites = nil
	if model.DependencyState == Point12ValEStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, dependencyReasons...)
	}
	if model.ReplayInvariantState == Point12ValEStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, replayReasons...)
	}
	if model.EvidenceQualityState == Point12ValEStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, qualityReasons...)
	}
	if model.BindingMutationState == Point12ValEStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, bindingReasons...)
	}
	if model.NoOverclaimState == Point12ValEStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, noOverclaimReasons...)
	}
	if model.CleanRoomIPState == Point12ValEStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, cleanRoomReasons...)
	}
	if model.RetentionProvenanceState == Point12ValEStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, retentionReasons...)
	}
	if model.PassClosureManifestState == Point12ValEStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, manifestReasons...)
	}
	return model
}
