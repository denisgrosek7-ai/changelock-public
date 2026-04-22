package formal

import "strings"

const (
	AudienceInternal        = "internal"
	AudienceAssessor        = "assessor"
	AudienceInsurer         = "insurer"
	AudienceRegulator       = "regulator"
	AudienceCertification   = "certification_body"
	AudienceGovernanceBoard = "governance_board"

	ClaimClassSupportingEvidenceOnly       = "supporting_evidence_only"
	ClaimClassMachineCheckableControl      = "machine_checkable_control_evidence"
	ClaimClassAssessorFacingSummary        = "assessor_facing_summary"
	ClaimClassInsurerFacingRiskInput       = "insurer_facing_risk_input"
	ClaimClassRegulatorSafeDisclosure      = "regulator_safe_disclosure"
	ClaimClassCertificationSupportArtifact = "certification_support_artifact"
	ClaimClassNotValidAsFormalClaim        = "not_valid_as_formal_claim"

	ProofClassTechnicalSupportOnly   = "technical_support_only"
	ProofClassFormalInternalReviewed = "formal_internal_reviewed"
	ProofClassAssessorReviewReady    = "assessor_review_ready"
	ProofClassExternalRelianceReady  = "external_reliance_ready"
	ProofClassNotSufficientExternal  = "not_sufficient_for_external_reliance"

	EntryGateStateReady      = "phase8_entry_gate_ready"
	EntryGateStateIncomplete = "phase8_entry_gate_incomplete"

	FoundationStateActive     = "phase8_contract_foundation_active"
	FoundationStateIncomplete = "phase8_contract_foundation_incomplete"

	FormalDisciplineStateActive     = "phase8_formal_discipline_active"
	FormalDisciplineStatePartial    = "phase8_formal_discipline_partial"
	FormalDisciplineStateIncomplete = "phase8_formal_discipline_incomplete"

	ComplianceCodificationStateActive     = "phase8_compliance_codification_active"
	ComplianceCodificationStatePartial    = "phase8_compliance_codification_partial"
	ComplianceCodificationStateIncomplete = "phase8_compliance_codification_incomplete"

	GovernedAutonomyStateActive     = "phase8_governed_autonomy_active"
	GovernedAutonomyStatePartial    = "phase8_governed_autonomy_partial"
	GovernedAutonomyStateIncomplete = "phase8_governed_autonomy_incomplete"

	Phase8StateIncomplete  = "phase8_formal_authority_core_incomplete"
	Phase8StateSubstantial = "phase8_formal_authority_core_substantially_ready"
	Phase8StateActive      = "phase8_formal_authority_core_active"
)

type EntryGate struct {
	CurrentState         string   `json:"current_state"`
	CarryOverLimitations []string `json:"carry_over_limitations,omitempty"`
	CarryOverDebt        []string `json:"carry_over_debt,omitempty"`
	ScopeBoundaries      []string `json:"scope_boundaries,omitempty"`
	ContractRefs         []string `json:"contract_refs,omitempty"`
	DeferredScope        []string `json:"deferred_scope,omitempty"`
}

type ClaimClass struct {
	SurfaceID              string   `json:"surface_id"`
	ClaimClass             string   `json:"claim_class"`
	AllowedUse             []string `json:"allowed_use,omitempty"`
	ProhibitedUse          []string `json:"prohibited_use,omitempty"`
	AudienceClass          string   `json:"audience_class"`
	StandardOfProof        string   `json:"standard_of_proof"`
	HumanReviewRequirement string   `json:"human_review_requirement"`
	MethodologyRefs        []string `json:"methodology_refs,omitempty"`
	InterpretationLimits   []string `json:"interpretation_limits,omitempty"`
	PublicationBoundaries  []string `json:"publication_boundaries,omitempty"`
}

type UsePermissionRule struct {
	SurfaceID          string   `json:"surface_id"`
	Audience           string   `json:"audience"`
	ClaimClass         string   `json:"claim_class"`
	AllowedUse         []string `json:"allowed_use,omitempty"`
	ForbiddenUse       []string `json:"forbidden_use,omitempty"`
	SharingScope       []string `json:"sharing_scope,omitempty"`
	RedactionProfile   []string `json:"redaction_profile,omitempty"`
	NeedsHumanApproval bool     `json:"needs_human_approval"`
	NeedsQuorum        bool     `json:"needs_quorum"`
	CanExport          bool     `json:"can_export"`
	CanCitePublicly    bool     `json:"can_cite_publicly"`
}

type StandardOfProofClass struct {
	SurfaceID                    string `json:"surface_id"`
	ProofClass                   string `json:"proof_class"`
	EvidenceSufficiencyFloor     string `json:"evidence_sufficiency_floor"`
	MinimumReviewLevel           string `json:"minimum_review_level"`
	HumanAttestationRequired     bool   `json:"human_attestation_required"`
	IndependentSecondPartyReview bool   `json:"independent_second_party_review"`
	ExternalAssessorConfirmation bool   `json:"external_assessor_confirmation"`
}

type JurisdictionProfile struct {
	SurfaceID                  string   `json:"surface_id"`
	ProfileID                  string   `json:"profile_id"`
	Jurisdiction               string   `json:"jurisdiction"`
	EffectiveFrom              string   `json:"effective_from"`
	EffectiveUntil             string   `json:"effective_until,omitempty"`
	DeprecationState           string   `json:"deprecation_state"`
	ConflictsWith              []string `json:"conflicts_with,omitempty"`
	JurisdictionPriorityRule   string   `json:"jurisdiction_priority_rule"`
	FallbackInterpretationMode string   `json:"fallback_interpretation_mode"`
}

type ConflictResolutionRule struct {
	SurfaceID          string   `json:"surface_id"`
	ScenarioID         string   `json:"scenario_id"`
	CurrentState       string   `json:"current_state"`
	PriorityOrdering   []string `json:"priority_ordering,omitempty"`
	ResolutionBehavior []string `json:"resolution_behavior,omitempty"`
}

type EvidenceCustodyContract struct {
	SurfaceID               string   `json:"surface_id"`
	CustodyOwner            string   `json:"custody_owner"`
	RedactionPolicy         string   `json:"redaction_policy"`
	RetentionClass          string   `json:"retention_class"`
	LegalHoldMode           string   `json:"legal_hold_mode"`
	DisclosureAudience      string   `json:"disclosure_audience"`
	RequiredFields          []string `json:"required_fields,omitempty"`
	ReleaseApprovalRequired bool     `json:"release_approval_required"`
}

type ArtifactLifecycleWorkflow struct {
	SurfaceID               string   `json:"surface_id"`
	States                  []string `json:"states,omitempty"`
	FreezeOnChallenge       string   `json:"freeze_on_challenge"`
	WithdrawalBehavior      string   `json:"withdrawal_behavior"`
	ReactivationRequirement string   `json:"reactivation_requirement"`
}

type PolicyAsLawProfile struct {
	SurfaceID                    string   `json:"surface_id"`
	ProfileID                    string   `json:"profile_id"`
	CurrentState                 string   `json:"current_state"`
	MachineCheckableCoverageRate string   `json:"machine_checkable_coverage_rate"`
	ManualInterpretationSections []string `json:"manual_interpretation_sections,omitempty"`
	AmbiguityMarkers             []string `json:"ambiguity_markers,omitempty"`
	JurisdictionRefs             []string `json:"jurisdiction_refs,omitempty"`
}

type RegulatoryMapping struct {
	SurfaceID                    string   `json:"surface_id"`
	MappingID                    string   `json:"mapping_id"`
	CurrentState                 string   `json:"current_state"`
	ControlConflictMarkers       []string `json:"control_conflict_markers,omitempty"`
	CompensatingControlSemantics []string `json:"compensating_control_semantics,omitempty"`
	InheritedControlSemantics    []string `json:"inherited_control_semantics,omitempty"`
}

type CertificationEvidencePack struct {
	SurfaceID                 string   `json:"surface_id"`
	PackID                    string   `json:"pack_id"`
	CurrentState              string   `json:"current_state"`
	EvidenceSufficiencyClass  string   `json:"evidence_sufficiency_class"`
	EvidenceFreezeForSnapshot bool     `json:"evidence_freeze_for_snapshot"`
	AssessorIndependenceNotes []string `json:"assessor_independence_notes,omitempty"`
}

type VerifierSurface struct {
	SurfaceID              string   `json:"surface_id"`
	AudienceClass          string   `json:"audience_class"`
	DisclosureBasisClass   string   `json:"disclosure_basis_class"`
	ExportNarrowingProfile []string `json:"export_narrowing_profile,omitempty"`
	ConfidentialityWarning []string `json:"confidentiality_warning,omitempty"`
}

type ChallengeWorkflow struct {
	SurfaceID                string   `json:"surface_id"`
	ResponseSLA              string   `json:"response_sla"`
	MandatoryReviewerClasses []string `json:"mandatory_reviewer_classes,omitempty"`
	InterimValidityState     string   `json:"interim_validity_state"`
	States                   []string `json:"states,omitempty"`
}

type AuthorityControl struct {
	SurfaceID               string   `json:"surface_id"`
	NonDelegableActions     []string `json:"non_delegable_actions,omitempty"`
	SeparationOfDutiesRules []string `json:"separation_of_duties_rules,omitempty"`
	QuorumRules             []string `json:"quorum_rules,omitempty"`
	AlternateApproverPath   string   `json:"alternate_approver_path"`
	DeadlockResolutionPath  string   `json:"deadlock_resolution_path"`
	EmergencySuspensionPath string   `json:"emergency_suspension_path"`
}

type AIGuardrail struct {
	SurfaceID                       string   `json:"surface_id"`
	ProhibitedRecommendationClasses []string `json:"prohibited_recommendation_classes,omitempty"`
	EscalationToHumanThreshold      string   `json:"escalation_to_human_threshold"`
	ConfidenceFloor                 string   `json:"confidence_floor"`
	BoundaryRules                   []string `json:"boundary_rules,omitempty"`
}

type ModelRiskContract struct {
	SurfaceID             string   `json:"surface_id"`
	ModelID               string   `json:"model_id"`
	Owner                 string   `json:"owner"`
	IntendedUse           string   `json:"intended_use"`
	ForbiddenUse          []string `json:"forbidden_use,omitempty"`
	ReferenceBasis        []string `json:"reference_basis,omitempty"`
	DriftMonitoring       []string `json:"drift_monitoring,omitempty"`
	FailureModeNotes      []string `json:"failure_mode_notes,omitempty"`
	RollbackPath          string   `json:"rollback_path"`
	ChallengerModelReview string   `json:"challenger_model_review"`
}

type InstitutionalDependency struct {
	SurfaceID       string   `json:"surface_id"`
	DependencyID    string   `json:"dependency_id"`
	DependencyClass string   `json:"dependency_class"`
	Version         string   `json:"version"`
	Owner           string   `json:"owner"`
	ChangeTriggers  []string `json:"change_triggers,omitempty"`
}

type Coverage struct {
	ClaimClasses               int `json:"claim_classes"`
	UsePermissionRules         int `json:"use_permission_rules"`
	StandardOfProofClasses     int `json:"standard_of_proof_classes"`
	JurisdictionProfiles       int `json:"jurisdiction_profiles"`
	ConflictResolutionRules    int `json:"conflict_resolution_rules"`
	EvidenceCustodyContracts   int `json:"evidence_custody_contracts"`
	ArtifactLifecycleWorkflows int `json:"artifact_lifecycle_workflows"`
	PolicyAsLawProfiles        int `json:"policy_as_law_profiles"`
	RegulatoryMappings         int `json:"regulatory_mappings"`
	CertificationEvidencePacks int `json:"certification_evidence_packs"`
	VerifierSurfaces           int `json:"verifier_surfaces"`
	ChallengeWorkflows         int `json:"challenge_workflows"`
	AuthorityControls          int `json:"authority_controls"`
	AIGuardrails               int `json:"ai_guardrails"`
	ModelRiskContracts         int `json:"model_risk_contracts"`
	InstitutionalDependencies  int `json:"institutional_dependencies"`
}

var phase8CoreSurfacesByGroup = map[string][]string{
	"formal": {
		"formal.claim_taxonomy",
		"formal.use_permission_matrix",
		"formal.standard_of_proof",
		"formal.jurisdiction_model",
		"formal.conflict_resolution",
		"formal.evidence_custody",
		"formal.artifact_lifecycle",
	},
	"compliance": {
		"compliance.policy_as_law",
		"compliance.regulatory_mapping",
		"compliance.certification_pack",
		"compliance.verifier_surface",
	},
	"governance": {
		"governance.challenge_workflow",
		"governance.authority_controls",
		"governance.ai_guardrails",
		"governance.model_risk",
		"governance.dependency_registry",
	},
}

func EntryGateBaseline() EntryGate {
	return EntryGate{
		CurrentState: EntryGateStateReady,
		CarryOverLimitations: []string{
			"Phase 8 core reuses the existing execution, public-proof, and ecosystem evidence spine and does not create a separate formal truth store.",
			"Phase 8 core remains a bounded authority-support layer and does not claim legal, regulatory, insurer, or certification-body replacement authority.",
		},
		CarryOverDebt: []string{
			"Risk quantification, insurer export programs, incident attribution support, and actuarial benchmark discipline remain deferred to institutional expansion.",
			"Broader federated governance and advanced institutional disclosure programs remain outside the initial Phase 8 core pass.",
		},
		ScopeBoundaries: []string{
			"Formal outputs remain bounded by claim class, standard of proof, use permissions, and evidence custody discipline.",
			"AI-assisted and consensus-assisted governance remains advisory or review-routed and cannot perform non-delegable authority actions.",
		},
		ContractRefs: []string{
			"phase8.formal_claim_taxonomy",
			"phase8.use_permission_matrix",
			"phase8.standard_of_proof_model",
			"phase8.jurisdiction_and_conflict_model",
			"phase8.evidence_custody_and_release_workflow",
			"phase8.compliance_codification_baseline",
			"phase8.authority_governance_baseline",
		},
		DeferredScope: DeferredInstitutionalExpansion(),
	}
}

func DeferredInstitutionalExpansion() []string {
	return []string{
		"risk_quantification_baseline",
		"insurance_facing_evidence_exports",
		"incident_attribution_support",
		"actuarial_benchmark_discipline",
		"insurer_integration_program",
		"wider_federated_governance",
		"advanced_institutional_disclosure_programs",
	}
}

func ClaimClasses() []ClaimClass {
	return []ClaimClass{
		{
			SurfaceID:              "formal.claim_taxonomy",
			ClaimClass:             ClaimClassSupportingEvidenceOnly,
			AllowedUse:             []string{"technical support", "internal evidence review"},
			ProhibitedUse:          []string{"public certification claim", "regulatory approval claim"},
			AudienceClass:          AudienceInternal,
			StandardOfProof:        ProofClassTechnicalSupportOnly,
			HumanReviewRequirement: "required",
			MethodologyRefs:        []string{"docs/formal-phase8-plan.md#6-formal-claim-taxonomy"},
			InterpretationLimits:   []string{"Evidence support only; not valid as a formal external reliance statement."},
			PublicationBoundaries:  []string{"internal_only"},
		},
		{
			SurfaceID:              "formal.claim_taxonomy",
			ClaimClass:             ClaimClassMachineCheckableControl,
			AllowedUse:             []string{"machine-checkable control evidence", "assessor preparation"},
			ProhibitedUse:          []string{"blanket legal compliance claim"},
			AudienceClass:          AudienceAssessor,
			StandardOfProof:        ProofClassFormalInternalReviewed,
			HumanReviewRequirement: "required",
			MethodologyRefs:        []string{"docs/formal-phase8-plan.md#23-legislative-and-compliance-codification"},
			InterpretationLimits:   []string{"Covers control evidence only within the declared profile and version."},
			PublicationBoundaries:  []string{"assessor_or_internal_only"},
		},
		{
			SurfaceID:              "formal.claim_taxonomy",
			ClaimClass:             ClaimClassInsurerFacingRiskInput,
			AllowedUse:             []string{"insurer review support", "risk committee input"},
			ProhibitedUse:          []string{"automatic pricing promise", "public safety claim"},
			AudienceClass:          AudienceInsurer,
			StandardOfProof:        ProofClassAssessorReviewReady,
			HumanReviewRequirement: "required_with_quorum",
			MethodologyRefs:        []string{"docs/formal-phase8-plan.md#21-cyber-insurance-and-actuarial-integration"},
			InterpretationLimits:   []string{"Input to insurer process only; not a final economic or legal determination."},
			PublicationBoundaries:  []string{"insurer_scoped_export_only"},
		},
		{
			SurfaceID:              "formal.claim_taxonomy",
			ClaimClass:             ClaimClassRegulatorSafeDisclosure,
			AllowedUse:             []string{"bounded regulator-safe disclosure", "formal disclosure pack support"},
			ProhibitedUse:          []string{"unbounded public marketing use"},
			AudienceClass:          AudienceRegulator,
			StandardOfProof:        ProofClassExternalRelianceReady,
			HumanReviewRequirement: "required_with_quorum",
			MethodologyRefs:        []string{"docs/formal-phase8-plan.md#23-legislative-and-compliance-codification"},
			InterpretationLimits:   []string{"Disclosure remains bounded by audience, redaction profile, and legal-hold discipline."},
			PublicationBoundaries:  []string{"regulator_scoped_release_only"},
		},
		{
			SurfaceID:              "formal.claim_taxonomy",
			ClaimClass:             ClaimClassCertificationSupportArtifact,
			AllowedUse:             []string{"certification support artifact", "assessment snapshot"},
			ProhibitedUse:          []string{"self-issued certification statement"},
			AudienceClass:          AudienceCertification,
			StandardOfProof:        ProofClassAssessorReviewReady,
			HumanReviewRequirement: "required_with_quorum",
			MethodologyRefs:        []string{"docs/formal-phase8-plan.md#12-formal-artifact-release-and-withdrawal-workflow"},
			InterpretationLimits:   []string{"Supports certification workflow; does not itself create a certification outcome."},
			PublicationBoundaries:  []string{"certification_body_scoped_release_only"},
		},
		{
			SurfaceID:              "formal.claim_taxonomy",
			ClaimClass:             ClaimClassNotValidAsFormalClaim,
			AllowedUse:             []string{"internal explanatory context"},
			ProhibitedUse:          []string{"formal external reliance", "policy enforcement by authority implication"},
			AudienceClass:          AudienceInternal,
			StandardOfProof:        ProofClassNotSufficientExternal,
			HumanReviewRequirement: "required",
			MethodologyRefs:        []string{"docs/formal-phase8-plan.md#6-formal-claim-taxonomy"},
			InterpretationLimits:   []string{"This output must remain explicitly non-formal."},
			PublicationBoundaries:  []string{"internal_only"},
		},
	}
}

func UsePermissionRules() []UsePermissionRule {
	return []UsePermissionRule{
		{
			SurfaceID:          "formal.use_permission_matrix",
			Audience:           AudienceAssessor,
			ClaimClass:         ClaimClassMachineCheckableControl,
			AllowedUse:         []string{"assessment preparation", "control verification review"},
			ForbiddenUse:       []string{"public citation as certification"},
			SharingScope:       []string{"assessor_scoped"},
			RedactionProfile:   []string{"assessment_minimum_necessary"},
			NeedsHumanApproval: true,
			NeedsQuorum:        false,
			CanExport:          true,
			CanCitePublicly:    false,
		},
		{
			SurfaceID:          "formal.use_permission_matrix",
			Audience:           AudienceInsurer,
			ClaimClass:         ClaimClassInsurerFacingRiskInput,
			AllowedUse:         []string{"underwriting review support", "renewal posture review"},
			ForbiddenUse:       []string{"automatic pricing engine input without human sign-off", "public benchmark reuse"},
			SharingScope:       []string{"insurer_scoped", "tenant_approved_only"},
			RedactionProfile:   []string{"insurer_minimum_necessary"},
			NeedsHumanApproval: true,
			NeedsQuorum:        true,
			CanExport:          true,
			CanCitePublicly:    false,
		},
		{
			SurfaceID:          "formal.use_permission_matrix",
			Audience:           AudienceRegulator,
			ClaimClass:         ClaimClassRegulatorSafeDisclosure,
			AllowedUse:         []string{"formal disclosure support", "evidence-backed control demonstration"},
			ForbiddenUse:       []string{"public marketing reuse", "implicit cross-jurisdiction claim reuse"},
			SharingScope:       []string{"regulator_scoped", "release_approved_only"},
			RedactionProfile:   []string{"regulator_safe_disclosure"},
			NeedsHumanApproval: true,
			NeedsQuorum:        true,
			CanExport:          true,
			CanCitePublicly:    false,
		},
		{
			SurfaceID:          "formal.use_permission_matrix",
			Audience:           AudienceCertification,
			ClaimClass:         ClaimClassCertificationSupportArtifact,
			AllowedUse:         []string{"assessment snapshot", "issue-aging review"},
			ForbiddenUse:       []string{"self-certification statement"},
			SharingScope:       []string{"certification_body_scoped"},
			RedactionProfile:   []string{"certification_minimum_necessary"},
			NeedsHumanApproval: true,
			NeedsQuorum:        true,
			CanExport:          true,
			CanCitePublicly:    false,
		},
	}
}

func StandardOfProofClasses() []StandardOfProofClass {
	return []StandardOfProofClass{
		{
			SurfaceID:                    "formal.standard_of_proof",
			ProofClass:                   ProofClassTechnicalSupportOnly,
			EvidenceSufficiencyFloor:     "partial_evidence_allowed",
			MinimumReviewLevel:           "single_internal_review",
			HumanAttestationRequired:     true,
			IndependentSecondPartyReview: false,
			ExternalAssessorConfirmation: false,
		},
		{
			SurfaceID:                    "formal.standard_of_proof",
			ProofClass:                   ProofClassFormalInternalReviewed,
			EvidenceSufficiencyFloor:     "internally_verified",
			MinimumReviewLevel:           "formal_internal_review",
			HumanAttestationRequired:     true,
			IndependentSecondPartyReview: false,
			ExternalAssessorConfirmation: false,
		},
		{
			SurfaceID:                    "formal.standard_of_proof",
			ProofClass:                   ProofClassAssessorReviewReady,
			EvidenceSufficiencyFloor:     "review_complete",
			MinimumReviewLevel:           "two_party_internal_review",
			HumanAttestationRequired:     true,
			IndependentSecondPartyReview: true,
			ExternalAssessorConfirmation: false,
		},
		{
			SurfaceID:                    "formal.standard_of_proof",
			ProofClass:                   ProofClassExternalRelianceReady,
			EvidenceSufficiencyFloor:     "externally_verified_or_equivalent",
			MinimumReviewLevel:           "quorum_review",
			HumanAttestationRequired:     true,
			IndependentSecondPartyReview: true,
			ExternalAssessorConfirmation: true,
		},
		{
			SurfaceID:                    "formal.standard_of_proof",
			ProofClass:                   ProofClassNotSufficientExternal,
			EvidenceSufficiencyFloor:     "insufficient_for_external_reliance",
			MinimumReviewLevel:           "single_internal_review",
			HumanAttestationRequired:     true,
			IndependentSecondPartyReview: false,
			ExternalAssessorConfirmation: false,
		},
	}
}

func JurisdictionProfiles() []JurisdictionProfile {
	return []JurisdictionProfile{
		{
			SurfaceID:                  "formal.jurisdiction_model",
			ProfileID:                  "eu_dora_control_profile",
			Jurisdiction:               "eu",
			EffectiveFrom:              "2025-01-17",
			DeprecationState:           "current",
			ConflictsWith:              []string{"us_sectoral_control_overlay"},
			JurisdictionPriorityRule:   "explicit_regional_profile_overrides_global_baseline",
			FallbackInterpretationMode: "manual_resolution_required",
		},
		{
			SurfaceID:                  "formal.jurisdiction_model",
			ProfileID:                  "us_sectoral_control_overlay",
			Jurisdiction:               "us",
			EffectiveFrom:              "2025-01-01",
			DeprecationState:           "current",
			ConflictsWith:              []string{"eu_dora_control_profile"},
			JurisdictionPriorityRule:   "organization_specific_stricter_overlay_wins_when_declared",
			FallbackInterpretationMode: "manual_resolution_required",
		},
	}
}

func ConflictResolutionRules() []ConflictResolutionRule {
	return []ConflictResolutionRule{
		{
			SurfaceID:        "formal.conflict_resolution",
			ScenarioID:       "multi_jurisdiction_control_conflict",
			CurrentState:     "manual_resolution_required",
			PriorityOrdering: []string{"organization_specific_stricter_overlay", "regional_specific_profile", "global_baseline"},
			ResolutionBehavior: []string{
				"mark multiple_profiles_conflicting",
				"freeze release pending manual resolution",
			},
		},
		{
			SurfaceID:        "formal.conflict_resolution",
			ScenarioID:       "effective_date_overlap_conflict",
			CurrentState:     "multiple_profiles_conflicting",
			PriorityOrdering: []string{"current_effective_profile", "latest_non_deprecated_profile"},
			ResolutionBehavior: []string{
				"emit manual_resolution_required when overlap cannot be ordered safely",
			},
		},
	}
}

func EvidenceCustodyContracts() []EvidenceCustodyContract {
	return []EvidenceCustodyContract{
		{
			SurfaceID:          "formal.evidence_custody",
			CustodyOwner:       "formal_authority_ops",
			RedactionPolicy:    "institutional_minimum_necessary",
			RetentionClass:     "formal_release_retention",
			LegalHoldMode:      "supported",
			DisclosureAudience: "approved_external_audience_only",
			RequiredFields: []string{
				"export_id",
				"signing_status",
				"timestamp",
				"custody_owner",
				"redaction_policy",
				"retention_class",
				"legal_hold_mode",
				"disclosure_audience",
				"integrity_reference",
				"supersession_or_revocation_record",
				"release_approval_record",
			},
			ReleaseApprovalRequired: true,
		},
	}
}

func ArtifactLifecycleWorkflows() []ArtifactLifecycleWorkflow {
	return []ArtifactLifecycleWorkflow{
		{
			SurfaceID: "formal.artifact_lifecycle",
			States: []string{
				"draft",
				"internal_review",
				"approved_for_release",
				"released",
				"challenged",
				"superseded",
				"withdrawn",
				"revoked",
			},
			FreezeOnChallenge:       "challenged_artifacts_freeze_external_reuse_until_review",
			WithdrawalBehavior:      "withdrawn_artifacts_remain_traceable_and_not_silently_removed",
			ReactivationRequirement: "formal_re_review_required",
		},
	}
}

func PolicyAsLawProfiles() []PolicyAsLawProfile {
	return []PolicyAsLawProfile{
		{
			SurfaceID:                    "compliance.policy_as_law",
			ProfileID:                    "phase8_core_formal_control_profile",
			CurrentState:                 "policy_as_law_profile_active",
			MachineCheckableCoverageRate: "78_percent_machine_checkable",
			ManualInterpretationSections: []string{"human_materiality_judgment", "organization_specific_legal_overlay"},
			AmbiguityMarkers:             []string{"manual_interpretation_required_sections_visible"},
			JurisdictionRefs:             []string{"eu_dora_control_profile", "us_sectoral_control_overlay"},
		},
	}
}

func RegulatoryMappings() []RegulatoryMapping {
	return []RegulatoryMapping{
		{
			SurfaceID:                    "compliance.regulatory_mapping",
			MappingID:                    "formal_control_mapping_baseline",
			CurrentState:                 "regulatory_mapping_active",
			ControlConflictMarkers:       []string{"control_conflict_marker_supported"},
			CompensatingControlSemantics: []string{"compensating_control_supported"},
			InheritedControlSemantics:    []string{"inherited_control_supported"},
		},
	}
}

func CertificationEvidencePacks() []CertificationEvidencePack {
	return []CertificationEvidencePack{
		{
			SurfaceID:                 "compliance.certification_pack",
			PackID:                    "certification_support_pack_baseline",
			CurrentState:              "certification_pack_active",
			EvidenceSufficiencyClass:  "assessor_review_ready",
			EvidenceFreezeForSnapshot: true,
			AssessorIndependenceNotes: []string{"assessment snapshot is bounded and remains external-assessor dependent for any final certification outcome"},
		},
	}
}

func VerifierSurfaces() []VerifierSurface {
	return []VerifierSurface{
		{
			SurfaceID:              "compliance.verifier_surface",
			AudienceClass:          AudienceRegulator,
			DisclosureBasisClass:   ClaimClassRegulatorSafeDisclosure,
			ExportNarrowingProfile: []string{"regulator_safe_disclosure", "confidentiality_warning_required"},
			ConfidentialityWarning: []string{"disclosure remains audience-bounded and cannot be widened into public publication by default"},
		},
		{
			SurfaceID:              "compliance.verifier_surface",
			AudienceClass:          AudienceCertification,
			DisclosureBasisClass:   ClaimClassCertificationSupportArtifact,
			ExportNarrowingProfile: []string{"assessment_snapshot_only"},
			ConfidentialityWarning: []string{"support artifact is not equivalent to certification issuance"},
		},
	}
}

func ChallengeWorkflows() []ChallengeWorkflow {
	return []ChallengeWorkflow{
		{
			SurfaceID:                "governance.challenge_workflow",
			ResponseSLA:              "5_business_days",
			MandatoryReviewerClasses: []string{"formal_authority_reviewer", "legal_or_compliance_reviewer"},
			InterimValidityState:     "challenged_pending_review",
			States: []string{
				"challenged_pending_review",
				"rebuttal_submitted",
				"override_recorded",
				"superseded_after_review",
				"challenge_rejected",
				"withdrawn_pending_reissue",
			},
		},
	}
}

func AuthorityControls() []AuthorityControl {
	return []AuthorityControl{
		{
			SurfaceID: "governance.authority_controls",
			NonDelegableActions: []string{
				"regulator_facing_disclosure_approval",
				"insurer_facing_export_approval",
				"certification_support_release_approval",
				"release_from_legal_hold",
				"reactivate_challenged_artifact",
				"external_disclosure_audience_approval",
				"formal_profile_methodology_change",
			},
			SeparationOfDutiesRules: []string{
				"proposer_cannot_be_final_approver",
				"redaction_policy_change_requires_independent_reviewer",
				"policy_as_law_activation_requires_cross_function_review",
			},
			QuorumRules: []string{
				"regulator_disclosure_requires_two_person_rule",
				"insurer_risk_input_export_requires_cross_function_quorum",
				"certification_support_release_requires_assessor_ready_review",
			},
			AlternateApproverPath:   "named_alternate_authority_reviewer",
			DeadlockResolutionPath:  "formal_authority_board_escalation",
			EmergencySuspensionPath: "formal_authority_emergency_suspend_release",
		},
	}
}

func AIGuardrails() []AIGuardrail {
	return []AIGuardrail{
		{
			SurfaceID:                       "governance.ai_guardrails",
			ProhibitedRecommendationClasses: []string{"legal_verdict", "formal_authority_override", "non_delegable_action_execution"},
			EscalationToHumanThreshold:      "always_for_formal_release_or_override",
			ConfidenceFloor:                 "high_confidence_plus_human_review",
			BoundaryRules: []string{
				"AI suggestions remain advisory until approved through formal routing.",
				"Consensus and AI outputs cannot silently widen disclosure audience or claim class.",
			},
		},
	}
}

func ModelRiskContracts() []ModelRiskContract {
	return []ModelRiskContract{
		{
			SurfaceID:    "governance.model_risk",
			ModelID:      "phase8_formal_review_assist_v1",
			Owner:        "formal_authority_ops",
			IntendedUse:  "bounded review assistance for formal claim and routing preparation",
			ForbiddenUse: []string{"legal conclusion", "automatic release decision", "automatic adverse institutional decision"},
			ReferenceBasis: []string{
				"phase6_claim_governance",
				"phase7_authority_surface_matrix",
				"phase8_formal_claim_taxonomy",
			},
			DriftMonitoring:       []string{"review disagreement rate", "challenge reversal rate", "release rollback trigger"},
			FailureModeNotes:      []string{"overstated authority tone", "insufficient explanation quality", "jurisdiction conflict oversimplification"},
			RollbackPath:          "disable_formal_review_assist",
			ChallengerModelReview: "required_for_material_model_change",
		},
	}
}

func InstitutionalDependencies() []InstitutionalDependency {
	return []InstitutionalDependency{
		{
			SurfaceID:       "governance.dependency_registry",
			DependencyID:    "eu_dora_profile_reference",
			DependencyClass: "regulatory_framework",
			Version:         "2025-01",
			Owner:           "formal_authority_ops",
			ChangeTriggers:  []string{"effective_date_change", "regulator_interpretation_update"},
		},
		{
			SurfaceID:       "governance.dependency_registry",
			DependencyID:    "certification_scheme_baseline",
			DependencyClass: "certification_scheme",
			Version:         "1.0",
			Owner:           "formal_authority_ops",
			ChangeTriggers:  []string{"scheme_update", "assessor_guidance_change"},
		},
	}
}

func ContractsCoverage() Coverage {
	return Coverage{
		ClaimClasses:               len(ClaimClasses()),
		UsePermissionRules:         len(UsePermissionRules()),
		StandardOfProofClasses:     len(StandardOfProofClasses()),
		JurisdictionProfiles:       len(JurisdictionProfiles()),
		ConflictResolutionRules:    len(ConflictResolutionRules()),
		EvidenceCustodyContracts:   len(EvidenceCustodyContracts()),
		ArtifactLifecycleWorkflows: len(ArtifactLifecycleWorkflows()),
		PolicyAsLawProfiles:        len(PolicyAsLawProfiles()),
		RegulatoryMappings:         len(RegulatoryMappings()),
		CertificationEvidencePacks: len(CertificationEvidencePacks()),
		VerifierSurfaces:           len(VerifierSurfaces()),
		ChallengeWorkflows:         len(ChallengeWorkflows()),
		AuthorityControls:          len(AuthorityControls()),
		AIGuardrails:               len(AIGuardrails()),
		ModelRiskContracts:         len(ModelRiskContracts()),
		InstitutionalDependencies:  len(InstitutionalDependencies()),
	}
}

func EvaluateFoundationState(coverage Coverage) string {
	if coverage.ClaimClasses == 0 ||
		coverage.UsePermissionRules == 0 ||
		coverage.StandardOfProofClasses == 0 ||
		coverage.JurisdictionProfiles == 0 ||
		coverage.ConflictResolutionRules == 0 ||
		coverage.EvidenceCustodyContracts == 0 ||
		coverage.ArtifactLifecycleWorkflows == 0 ||
		coverage.PolicyAsLawProfiles == 0 ||
		coverage.RegulatoryMappings == 0 ||
		coverage.CertificationEvidencePacks == 0 ||
		coverage.VerifierSurfaces == 0 ||
		coverage.ChallengeWorkflows == 0 ||
		coverage.AuthorityControls == 0 ||
		coverage.AIGuardrails == 0 ||
		coverage.ModelRiskContracts == 0 ||
		coverage.InstitutionalDependencies == 0 {
		return FoundationStateIncomplete
	}
	if evaluateFoundationStateForPresence(buildCoreSurfacePresence()) != FoundationStateActive {
		return FoundationStateIncomplete
	}
	return FoundationStateActive
}

func EvaluateFormalDisciplineState() string {
	return evaluateGroupStateForPresence("formal", buildCoreSurfacePresence())
}

func EvaluateComplianceCodificationState() string {
	return evaluateGroupStateForPresence("compliance", buildCoreSurfacePresence())
}

func EvaluateGovernedAutonomyState() string {
	return evaluateGroupStateForPresence("governance", buildCoreSurfacePresence())
}

func EvaluatePhase8State(entryGateState, foundationState, formalState, complianceState, governanceState string) string {
	if strings.TrimSpace(entryGateState) != EntryGateStateReady || strings.TrimSpace(foundationState) != FoundationStateActive {
		return Phase8StateIncomplete
	}
	if !isFormalDisciplineState(formalState) || !isComplianceCodificationState(complianceState) || !isGovernedAutonomyState(governanceState) {
		return Phase8StateIncomplete
	}
	if formalState == FormalDisciplineStateIncomplete || complianceState == ComplianceCodificationStateIncomplete || governanceState == GovernedAutonomyStateIncomplete {
		return Phase8StateIncomplete
	}
	if formalState == FormalDisciplineStatePartial || complianceState == ComplianceCodificationStatePartial || governanceState == GovernedAutonomyStatePartial {
		return Phase8StateSubstantial
	}
	return Phase8StateActive
}

func buildCoreSurfacePresence() map[string]bool {
	presence := map[string]bool{}
	for _, item := range ClaimClasses() {
		presence[item.SurfaceID] = true
	}
	for _, item := range UsePermissionRules() {
		presence[item.SurfaceID] = true
	}
	for _, item := range StandardOfProofClasses() {
		presence[item.SurfaceID] = true
	}
	for _, item := range JurisdictionProfiles() {
		presence[item.SurfaceID] = true
	}
	for _, item := range ConflictResolutionRules() {
		presence[item.SurfaceID] = true
	}
	for _, item := range EvidenceCustodyContracts() {
		presence[item.SurfaceID] = true
	}
	for _, item := range ArtifactLifecycleWorkflows() {
		presence[item.SurfaceID] = true
	}
	for _, item := range PolicyAsLawProfiles() {
		presence[item.SurfaceID] = true
	}
	for _, item := range RegulatoryMappings() {
		presence[item.SurfaceID] = true
	}
	for _, item := range CertificationEvidencePacks() {
		presence[item.SurfaceID] = true
	}
	for _, item := range VerifierSurfaces() {
		presence[item.SurfaceID] = true
	}
	for _, item := range ChallengeWorkflows() {
		presence[item.SurfaceID] = true
	}
	for _, item := range AuthorityControls() {
		presence[item.SurfaceID] = true
	}
	for _, item := range AIGuardrails() {
		presence[item.SurfaceID] = true
	}
	for _, item := range ModelRiskContracts() {
		presence[item.SurfaceID] = true
	}
	for _, item := range InstitutionalDependencies() {
		presence[item.SurfaceID] = true
	}
	return presence
}

func evaluateFoundationStateForPresence(presence map[string]bool) string {
	for _, surfaces := range phase8CoreSurfacesByGroup {
		for _, surfaceID := range surfaces {
			if !presence[surfaceID] {
				return FoundationStateIncomplete
			}
		}
	}
	return FoundationStateActive
}

func evaluateGroupStateForPresence(group string, presence map[string]bool) string {
	surfaces := phase8CoreSurfacesByGroup[strings.TrimSpace(group)]
	if len(surfaces) == 0 {
		return Phase8StateIncomplete
	}
	anyPresent := false
	allPresent := true
	for _, surfaceID := range surfaces {
		if presence[surfaceID] {
			anyPresent = true
			continue
		}
		allPresent = false
	}
	switch strings.TrimSpace(group) {
	case "formal":
		if allPresent {
			return FormalDisciplineStateActive
		}
		if anyPresent {
			return FormalDisciplineStatePartial
		}
		return FormalDisciplineStateIncomplete
	case "compliance":
		if allPresent {
			return ComplianceCodificationStateActive
		}
		if anyPresent {
			return ComplianceCodificationStatePartial
		}
		return ComplianceCodificationStateIncomplete
	case "governance":
		if allPresent {
			return GovernedAutonomyStateActive
		}
		if anyPresent {
			return GovernedAutonomyStatePartial
		}
		return GovernedAutonomyStateIncomplete
	default:
		return Phase8StateIncomplete
	}
}

func isFormalDisciplineState(value string) bool {
	switch strings.TrimSpace(value) {
	case FormalDisciplineStateActive, FormalDisciplineStatePartial, FormalDisciplineStateIncomplete:
		return true
	default:
		return false
	}
}

func isComplianceCodificationState(value string) bool {
	switch strings.TrimSpace(value) {
	case ComplianceCodificationStateActive, ComplianceCodificationStatePartial, ComplianceCodificationStateIncomplete:
		return true
	default:
		return false
	}
}

func isGovernedAutonomyState(value string) bool {
	switch strings.TrimSpace(value) {
	case GovernedAutonomyStateActive, GovernedAutonomyStatePartial, GovernedAutonomyStateIncomplete:
		return true
	default:
		return false
	}
}
