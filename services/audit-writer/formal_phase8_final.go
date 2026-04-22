package main

import (
	"net/http"
	"time"

	formalcore "github.com/denisgrosek/changelock/internal/formal"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	phase8FinalSummarySchema = "8d.formal_phase8_final_summary.v1"

	phase8FinalReviewStateActive     = "phase8_final_review_active"
	phase8FinalReviewStatePartial    = "phase8_final_review_partial"
	phase8FinalReviewStateIncomplete = "phase8_final_review_incomplete"

	phase8FinalizationStateReady       = "phase8_finalization_ready"
	phase8FinalizationStateSubstantial = "phase8_finalization_substantially_ready"
	phase8FinalizationStateIncomplete  = "phase8_finalization_incomplete"
)

type phase8FinalReviewCheck struct {
	CheckID      string   `json:"check_id"`
	CurrentState string   `json:"current_state"`
	Summary      string   `json:"summary"`
	ReasonCodes  []string `json:"reason_codes,omitempty"`
	RouteRefs    []string `json:"route_refs,omitempty"`
	DocRefs      []string `json:"doc_refs,omitempty"`
	Limitations  []string `json:"limitations,omitempty"`
}

type phase8FinalReviewSection struct {
	CurrentState string                   `json:"current_state"`
	RouteRefs    []string                 `json:"route_refs,omitempty"`
	DocRefs      []string                 `json:"doc_refs,omitempty"`
	Checks       []phase8FinalReviewCheck `json:"checks,omitempty"`
	Limitations  []string                 `json:"limitations,omitempty"`
}

type phase8FinalSummaryResponse struct {
	SchemaVersion                         string                   `json:"schema_version"`
	GeneratedAt                           time.Time                `json:"generated_at"`
	CurrentState                          string                   `json:"current_state"`
	Phase8CoreState                       string                   `json:"phase8_core_state"`
	LegalAndRegulatoryClaimReview         phase8FinalReviewSection `json:"legal_and_regulatory_claim_review"`
	UsePermissionAndStandardOfProofReview phase8FinalReviewSection `json:"use_permission_and_standard_of_proof_review"`
	CustodyRedactionAndLegalHoldReview    phase8FinalReviewSection `json:"custody_redaction_and_legal_hold_review"`
	ComplianceCodificationReview          phase8FinalReviewSection `json:"compliance_codification_review"`
	CertificationWorkflowReview           phase8FinalReviewSection `json:"certification_workflow_review"`
	ConsensusGovernanceReview             phase8FinalReviewSection `json:"consensus_governance_review"`
	AIGuardrailReview                     phase8FinalReviewSection `json:"ai_guardrail_review"`
	InsurerEvidenceReview                 phase8FinalReviewSection `json:"insurer_evidence_review"`
	FormalAuthorityBoundaryReview         phase8FinalReviewSection `json:"formal_authority_boundary_review"`
	DeferredInstitutionalScope            []string                 `json:"deferred_institutional_scope,omitempty"`
	Limitations                           []string                 `json:"limitations,omitempty"`
}

func (s server) phase8FinalSummaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8FinalSummary())
}

func buildPhase8FinalSummary() phase8FinalSummaryResponse {
	proofs := buildPhase8Proofs()
	legalAndRegulatory := buildPhase8LegalAndRegulatoryClaimReview()
	usePermissionAndProof := buildPhase8UsePermissionAndStandardOfProofReview()
	custodyReview := buildPhase8CustodyRedactionAndLegalHoldReview()
	complianceReview := buildPhase8ComplianceCodificationFinalReview()
	certificationReview := buildPhase8CertificationWorkflowFinalReview()
	consensusReview := buildPhase8ConsensusGovernanceReview()
	aiReview := buildPhase8AIGuardrailFinalReview()
	insurerEvidenceReview := buildPhase8InsurerEvidenceReview()
	boundaryReview := buildPhase8FormalAuthorityBoundaryReview()
	return phase8FinalSummaryResponse{
		SchemaVersion:                         phase8FinalSummarySchema,
		GeneratedAt:                           publicSampleTime(),
		CurrentState:                          phase8FinalizationState(proofs.CurrentState, legalAndRegulatory, usePermissionAndProof, custodyReview, complianceReview, certificationReview, consensusReview, aiReview, insurerEvidenceReview, boundaryReview),
		Phase8CoreState:                       proofs.CurrentState,
		LegalAndRegulatoryClaimReview:         legalAndRegulatory,
		UsePermissionAndStandardOfProofReview: usePermissionAndProof,
		CustodyRedactionAndLegalHoldReview:    custodyReview,
		ComplianceCodificationReview:          complianceReview,
		CertificationWorkflowReview:           certificationReview,
		ConsensusGovernanceReview:             consensusReview,
		AIGuardrailReview:                     aiReview,
		InsurerEvidenceReview:                 insurerEvidenceReview,
		FormalAuthorityBoundaryReview:         boundaryReview,
		DeferredInstitutionalScope:            phase8RemainingDeferredScope(),
		Limitations: []string{
			"Phase 8 final summary closes the bounded formal-authority package over core, 8A, 8B, and 8C without creating a new truth store or replacement authority layer.",
			"Insurer integration programs, wider federated governance, and advanced institutional disclosure programs remain outside the bounded Phase 8 finalization pack.",
		},
	}
}

func buildPhase8LegalAndRegulatoryClaimReview() phase8FinalReviewSection {
	formal := buildPhase8FormalDiscipline()
	verifier := buildPhase8VerifierSurfaces()
	regulatorClaim := phase8ClaimClassByID(formalcore.ClaimClassRegulatorSafeDisclosure)
	regulatorRule := phase8UsePermissionRuleByAudience(formalcore.AudienceRegulator)
	active := formal.CurrentState == formalcore.FormalDisciplineStateActive &&
		verifier.CurrentState == phase8VerifierSurfacesStateActive &&
		regulatorClaim.AudienceClass == formalcore.AudienceRegulator &&
		regulatorClaim.StandardOfProof == formalcore.ProofClassExternalRelianceReady &&
		regulatorRule.NeedsHumanApproval &&
		regulatorRule.NeedsQuorum &&
		containsString(verifier.AllowedAudienceClasses, formalcore.AudienceRegulator) &&
		containsString(verifier.ForbiddenAudiences, "public") &&
		containsString(verifier.ForbiddenAudiences, formalcore.AudienceInsurer)
	return phase8FinalReviewSection{
		CurrentState: phase8SectionState(active, formal.CurrentState != formalcore.FormalDisciplineStateIncomplete && verifier.CurrentState != phase8VerifierSurfacesStateIncomplete),
		RouteRefs: []string{
			"/v1/formal/phase8/formal-discipline",
			"/v1/formal/phase8/compliance/verifier-surfaces",
			"/v1/formal/phase8/contracts",
		},
		DocRefs: []string{
			"docs/formal-phase8-core.md",
			"docs/formal-phase8-vala.md",
			"docs/formal-phase8-final.md",
		},
		Checks: []phase8FinalReviewCheck{
			{
				CheckID:      "regulator_safe_claim_class",
				CurrentState: phase8ReviewCheckState(regulatorClaim.AudienceClass == formalcore.AudienceRegulator && regulatorClaim.StandardOfProof == formalcore.ProofClassExternalRelianceReady),
				Summary:      "Regulator-safe disclosure keeps an explicit claim class, bounded audience, and external-reliance-ready proof threshold.",
				ReasonCodes:  []string{"regulator_safe_disclosure", "external_reliance_ready"},
				RouteRefs:    []string{"/v1/formal/phase8/formal-discipline"},
				DocRefs:      []string{"docs/formal-phase8-core.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "regulator_use_permission_boundary",
				CurrentState: phase8ReviewCheckState(regulatorRule.NeedsHumanApproval && regulatorRule.NeedsQuorum && !regulatorRule.CanCitePublicly),
				Summary:      "Regulator-facing disclosure remains release-approved, quorum-bound, and non-public by default.",
				ReasonCodes:  []string{"use_permission_matrix_enforced", "no_public_reuse"},
				RouteRefs:    []string{"/v1/formal/phase8/contracts"},
				DocRefs:      []string{"docs/formal-phase8-core.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "verifier_surface_audience_narrowing",
				CurrentState: phase8ReviewCheckState(containsString(verifier.AllowedAudienceClasses, formalcore.AudienceRegulator) && containsString(verifier.ForbiddenAudiences, "public") && containsString(verifier.ForbiddenAudiences, formalcore.AudienceInsurer)),
				Summary:      "Regulator-safe verifier surfaces remain audience-narrowed and cannot widen into public or insurer-facing disclosure paths.",
				ReasonCodes:  []string{"audience_narrowed_verifier_surface", "insurer_and_public_blocked"},
				RouteRefs:    []string{"/v1/formal/phase8/compliance/verifier-surfaces"},
				DocRefs:      []string{"docs/formal-phase8-vala.md", "docs/formal-phase8-final.md"},
			},
		},
		Limitations: []string{
			"Legal and regulatory claim review confirms bounded disclosure discipline; it does not claim direct legal or regulatory authority.",
		},
	}
}

func buildPhase8UsePermissionAndStandardOfProofReview() phase8FinalReviewSection {
	formal := buildPhase8FormalDiscipline()
	insurerRule := phase8UsePermissionRuleByAudience(formalcore.AudienceInsurer)
	regulatorRule := phase8UsePermissionRuleByAudience(formalcore.AudienceRegulator)
	externalReliance := phase8ProofClassByID(formalcore.ProofClassExternalRelianceReady)
	notSufficientExternal := phase8ProofClassByID(formalcore.ProofClassNotSufficientExternal)
	nonFormalClaim := phase8ClaimClassByID(formalcore.ClaimClassNotValidAsFormalClaim)
	active := formal.CurrentState == formalcore.FormalDisciplineStateActive &&
		insurerRule.NeedsHumanApproval &&
		insurerRule.NeedsQuorum &&
		regulatorRule.NeedsHumanApproval &&
		regulatorRule.NeedsQuorum &&
		externalReliance.ExternalAssessorConfirmation &&
		externalReliance.IndependentSecondPartyReview &&
		nonFormalClaim.StandardOfProof == formalcore.ProofClassNotSufficientExternal &&
		notSufficientExternal.EvidenceSufficiencyFloor != ""
	return phase8FinalReviewSection{
		CurrentState: phase8SectionState(active, formal.CurrentState != formalcore.FormalDisciplineStateIncomplete),
		RouteRefs: []string{
			"/v1/formal/phase8/formal-discipline",
			"/v1/formal/phase8/contracts",
		},
		DocRefs: []string{
			"docs/formal-phase8-core.md",
			"docs/formal-phase8-final.md",
		},
		Checks: []phase8FinalReviewCheck{
			{
				CheckID:      "external_use_permissions_require_review",
				CurrentState: phase8ReviewCheckState(insurerRule.NeedsHumanApproval && insurerRule.NeedsQuorum && regulatorRule.NeedsHumanApproval && regulatorRule.NeedsQuorum),
				Summary:      "External-facing use-permission rules remain human-reviewed and quorum-bound for insurer and regulator audiences.",
				ReasonCodes:  []string{"human_approval_required", "quorum_required"},
				RouteRefs:    []string{"/v1/formal/phase8/contracts"},
				DocRefs:      []string{"docs/formal-phase8-core.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "proof_classes_preserve_reliance_levels",
				CurrentState: phase8ReviewCheckState(externalReliance.ExternalAssessorConfirmation && externalReliance.IndependentSecondPartyReview && notSufficientExternal.EvidenceSufficiencyFloor != ""),
				Summary:      "Standard-of-proof classes remain stratified between external reliance and explicitly insufficient external use.",
				ReasonCodes:  []string{"standard_of_proof_visible", "external_reliance_not_equivalent_to_non_formal"},
				RouteRefs:    []string{"/v1/formal/phase8/formal-discipline"},
				DocRefs:      []string{"docs/formal-phase8-core.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "non_formal_outputs_stay_non_external",
				CurrentState: phase8ReviewCheckState(nonFormalClaim.StandardOfProof == formalcore.ProofClassNotSufficientExternal),
				Summary:      "Outputs marked not valid as formal claims remain explicitly outside external reliance paths.",
				ReasonCodes:  []string{"not_valid_as_formal_claim", "no_external_reliance"},
				RouteRefs:    []string{"/v1/formal/phase8/formal-discipline"},
				DocRefs:      []string{"docs/formal-phase8-core.md", "docs/formal-phase8-final.md"},
			},
		},
		Limitations: []string{
			"Use-permission and proof review confirms bounded external reliance discipline and does not widen claim classes by implication.",
		},
	}
}

func buildPhase8CustodyRedactionAndLegalHoldReview() phase8FinalReviewSection {
	formal := buildPhase8FormalDiscipline()
	insurance := buildPhase8InsuranceExports()
	custody := phase8EvidenceCustodyContract()
	lifecycle := phase8CertificationLifecycle()
	exportBounded := len(insurance.Exports) > 0 && insurance.Exports[0].ReleaseApprovalRequired && !insurance.Exports[0].CanCitePublicly
	active := formal.CurrentState == formalcore.FormalDisciplineStateActive &&
		insurance.CurrentState == phase8InsuranceExportsStateActive &&
		custody.ReleaseApprovalRequired &&
		containsString(custody.RequiredFields, "redaction_policy") &&
		containsString(custody.RequiredFields, "legal_hold_mode") &&
		containsString(custody.RequiredFields, "release_approval_record") &&
		containsString(custody.RequiredFields, "integrity_reference") &&
		exportBounded &&
		lifecycle.ReactivationRequirement == "formal_re_review_required"
	return phase8FinalReviewSection{
		CurrentState: phase8SectionState(active, formal.CurrentState != formalcore.FormalDisciplineStateIncomplete && insurance.CurrentState != phase8InsuranceExportsStateIncomplete),
		RouteRefs: []string{
			"/v1/formal/phase8/formal-discipline",
			"/v1/formal/phase8/institutional/insurance-exports",
			"/v1/formal/phase8/contracts",
		},
		DocRefs: []string{
			"docs/formal-phase8-core.md",
			"docs/formal-phase8-valc.md",
			"docs/formal-phase8-final.md",
		},
		Checks: []phase8FinalReviewCheck{
			{
				CheckID:      "custody_required_fields_visible",
				CurrentState: phase8ReviewCheckState(custody.ReleaseApprovalRequired && containsString(custody.RequiredFields, "redaction_policy") && containsString(custody.RequiredFields, "legal_hold_mode") && containsString(custody.RequiredFields, "release_approval_record")),
				Summary:      "Formal custody keeps release approval, redaction, legal-hold, and integrity references explicit.",
				ReasonCodes:  []string{"custody_owner_visible", "release_approval_record_required", "legal_hold_supported"},
				RouteRefs:    []string{"/v1/formal/phase8/contracts"},
				DocRefs:      []string{"docs/formal-phase8-core.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "insurer_exports_remain_non_public",
				CurrentState: phase8ReviewCheckState(exportBounded),
				Summary:      "Insurer-facing exports remain release-approved and non-public under explicit disclosure boundaries.",
				ReasonCodes:  []string{"insurer_scoped_release_only", "public_citation_blocked"},
				RouteRefs:    []string{"/v1/formal/phase8/institutional/insurance-exports"},
				DocRefs:      []string{"docs/formal-phase8-valc.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "challenged_artifacts_require_rereview",
				CurrentState: phase8ReviewCheckState(lifecycle.ReactivationRequirement == "formal_re_review_required"),
				Summary:      "Challenged or withdrawn formal artifacts cannot silently return to active release without re-review.",
				ReasonCodes:  []string{"formal_re_review_required", "challenge_freeze_preserved"},
				RouteRefs:    []string{"/v1/formal/phase8/compliance/certification-workflow"},
				DocRefs:      []string{"docs/formal-phase8-vala.md", "docs/formal-phase8-final.md"},
			},
		},
		Limitations: []string{
			"Custody review confirms release discipline and traceability; it does not imply unrestricted external reuse of formal artifacts.",
		},
	}
}

func buildPhase8ComplianceCodificationFinalReview() phase8FinalReviewSection {
	compliance := buildPhase8ComplianceCodification()
	policy := buildPhase8PolicyProfiles()
	mappings := buildPhase8RegulatoryMappings()
	automation := buildPhase8EvidenceAutomation()
	active := compliance.CurrentState == formalcore.ComplianceCodificationStateActive &&
		policy.CurrentState == phase8PolicyProfilesStateActive &&
		mappings.CurrentState == phase8RegulatoryMappingsStateActive &&
		automation.CurrentState == phase8EvidenceAutomationStateActive &&
		containsString(policy.ActivationBoundaries, "manual_interpretation_sections_remain_visible") &&
		containsString(policy.ActivationBoundaries, "conflicting_jurisdiction_profiles_require_manual_resolution") &&
		hasEvidenceAutomationExclusion(automation.NotInPackReasons, "direct_certification_issuance", false) &&
		hasEvidenceAutomationExclusion(automation.NotInPackReasons, "legal_verdict_output", false)
	return phase8FinalReviewSection{
		CurrentState: phase8SectionState(active, compliance.CurrentState != formalcore.ComplianceCodificationStateIncomplete),
		RouteRefs: []string{
			"/v1/formal/phase8/compliance-codification",
			"/v1/formal/phase8/compliance/policy-profiles",
			"/v1/formal/phase8/compliance/regulatory-mappings",
			"/v1/formal/phase8/compliance/evidence-automation",
		},
		DocRefs: []string{
			"docs/formal-phase8-core.md",
			"docs/formal-phase8-vala.md",
			"docs/formal-phase8-final.md",
		},
		Checks: []phase8FinalReviewCheck{
			{
				CheckID:      "policy_profiles_preserve_manual_resolution",
				CurrentState: phase8ReviewCheckState(containsString(policy.ActivationBoundaries, "manual_interpretation_sections_remain_visible") && containsString(policy.ActivationBoundaries, "conflicting_jurisdiction_profiles_require_manual_resolution")),
				Summary:      "Policy-as-law profiles keep manual interpretation and jurisdiction conflict resolution visible.",
				ReasonCodes:  []string{"manual_interpretation_visible", "manual_resolution_required"},
				RouteRefs:    []string{"/v1/formal/phase8/compliance/policy-profiles"},
				DocRefs:      []string{"docs/formal-phase8-vala.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "regulatory_mappings_keep_conflicts_visible",
				CurrentState: phase8ReviewCheckState(mappings.ConflictHandlingState == "control_conflict_handling_visible" && mappings.CompensatingControlState == "compensating_control_semantics_visible" && mappings.InheritedControlState == "inherited_control_semantics_visible"),
				Summary:      "Regulatory mappings retain conflict, compensating-control, and inherited-control semantics rather than flattening them away.",
				ReasonCodes:  []string{"control_conflict_visible", "compensating_controls_visible", "inherited_controls_visible"},
				RouteRefs:    []string{"/v1/formal/phase8/compliance/regulatory-mappings"},
				DocRefs:      []string{"docs/formal-phase8-vala.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "evidence_automation_excludes_forbidden_outputs",
				CurrentState: phase8ReviewCheckState(hasEvidenceAutomationExclusion(automation.NotInPackReasons, "direct_certification_issuance", false) && hasEvidenceAutomationExclusion(automation.NotInPackReasons, "legal_verdict_output", false)),
				Summary:      "Compliance evidence automation stays bounded and excludes direct certification issuance and legal verdict outputs.",
				ReasonCodes:  []string{"support_artifact_only", "no_legal_verdict_output"},
				RouteRefs:    []string{"/v1/formal/phase8/compliance/evidence-automation"},
				DocRefs:      []string{"docs/formal-phase8-vala.md", "docs/formal-phase8-final.md"},
			},
		},
		Limitations: []string{
			"Compliance codification review confirms bounded machine-checkable support and does not elevate codification into direct legal or certification authority.",
		},
	}
}

func buildPhase8CertificationWorkflowFinalReview() phase8FinalReviewSection {
	workflow := buildPhase8CertificationWorkflow()
	hasFrozenSnapshot := len(workflow.EvidencePacks) > 0 && workflow.EvidencePacks[0].EvidenceFreezeForSnapshot
	active := workflow.CurrentState == phase8CertificationWorkflowStateActive &&
		workflow.ReleaseBoundaryState == "support_artifact_not_certification_issuance" &&
		workflow.SnapshotState == "assessment_snapshot_and_evidence_freeze_visible" &&
		hasFrozenSnapshot
	return phase8FinalReviewSection{
		CurrentState: phase8SectionState(active, workflow.CurrentState != phase8CertificationWorkflowStateIncomplete),
		RouteRefs: []string{
			"/v1/formal/phase8/compliance/certification-workflow",
			"/v1/formal/phase8/compliance/verifier-surfaces",
			"/v1/formal/phase8/proofs",
		},
		DocRefs: []string{
			"docs/formal-phase8-vala.md",
			"docs/formal-phase8-final.md",
		},
		Checks: []phase8FinalReviewCheck{
			{
				CheckID:      "certification_support_artifact_only",
				CurrentState: phase8ReviewCheckState(workflow.ReleaseBoundaryState == "support_artifact_not_certification_issuance"),
				Summary:      "Certification workflow remains support-artifact only and never becomes self-issued certification.",
				ReasonCodes:  []string{"support_artifact_not_certification_issuance", "no_self_issued_authority"},
				RouteRefs:    []string{"/v1/formal/phase8/compliance/certification-workflow"},
				DocRefs:      []string{"docs/formal-phase8-vala.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "assessment_snapshot_freeze_visible",
				CurrentState: phase8ReviewCheckState(len(workflow.EvidencePacks) > 0 && workflow.EvidencePacks[0].EvidenceFreezeForSnapshot),
				Summary:      "Assessment snapshots preserve evidence freeze and assessor-ready review boundaries.",
				ReasonCodes:  []string{"assessment_snapshot_frozen", "assessor_review_ready"},
				RouteRefs:    []string{"/v1/formal/phase8/compliance/certification-workflow"},
				DocRefs:      []string{"docs/formal-phase8-vala.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "challenge_and_withdrawal_traceability",
				CurrentState: phase8ReviewCheckState(workflow.Lifecycle.ReactivationRequirement == "formal_re_review_required" && workflow.ChallengeWorkflow.InterimValidityState == "challenged_pending_review"),
				Summary:      "Challenged certification support artifacts remain traceable and re-review-gated before any return to active use.",
				ReasonCodes:  []string{"challenge_workflow_active", "formal_re_review_required"},
				RouteRefs:    []string{"/v1/formal/phase8/compliance/certification-workflow"},
				DocRefs:      []string{"docs/formal-phase8-vala.md", "docs/formal-phase8-final.md"},
			},
		},
		Limitations: []string{
			"Certification workflow review confirms assessor-facing support discipline and does not imply certification issuance authority.",
		},
	}
}

func buildPhase8ConsensusGovernanceReview() phase8FinalReviewSection {
	consensus := buildPhase8ConsensusReview()
	suggestions := buildPhase8PolicySuggestions()
	routing := buildPhase8AuthorityRouting()
	active := consensus.CurrentState == phase8ConsensusReviewStateActive &&
		suggestions.CurrentState == phase8PolicySuggestionsStateActive &&
		routing.CurrentState == phase8AuthorityRoutingStateActive &&
		consensus.ConsensusState == "consensus_assisted_review_visible" &&
		consensus.QuorumState == "quorum_threshold_and_abstain_visible" &&
		suggestions.ApprovalBoundary == "advisory_until_formally_approved" &&
		containsForbiddenAction(suggestions.Suggestions, "automatic_profile_activation") &&
		containsForbiddenAction(suggestions.Suggestions, "direct_external_release") &&
		routing.AlternateApproverPath != "" &&
		routing.DeadlockResolutionPath != "" &&
		routing.EmergencySuspensionPath != ""
	return phase8FinalReviewSection{
		CurrentState: phase8SectionState(active, consensus.CurrentState != phase8ConsensusReviewStateIncomplete && suggestions.CurrentState != phase8PolicySuggestionsStateIncomplete && routing.CurrentState != phase8AuthorityRoutingStateIncomplete),
		RouteRefs: []string{
			"/v1/formal/phase8/governance/consensus-review",
			"/v1/formal/phase8/governance/policy-suggestions",
			"/v1/formal/phase8/governance/authority-routing",
		},
		DocRefs: []string{
			"docs/formal-phase8-valb.md",
			"docs/formal-phase8-final.md",
		},
		Checks: []phase8FinalReviewCheck{
			{
				CheckID:      "consensus_review_remains_review_support",
				CurrentState: phase8ReviewCheckState(consensus.ConsensusState == "consensus_assisted_review_visible" && consensus.QuorumState == "quorum_threshold_and_abstain_visible" && consensus.MinorityState == "weighted_disagreement_and_minority_report_visible"),
				Summary:      "Consensus support remains review-support only with quorum, abstain, disagreement, and minority-report visibility.",
				ReasonCodes:  []string{"consensus_assist_only", "minority_report_preserved"},
				RouteRefs:    []string{"/v1/formal/phase8/governance/consensus-review"},
				DocRefs:      []string{"docs/formal-phase8-valb.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "policy_suggestions_cannot_auto_activate",
				CurrentState: phase8ReviewCheckState(suggestions.ApprovalBoundary == "advisory_until_formally_approved" && containsForbiddenAction(suggestions.Suggestions, "automatic_profile_activation") && containsForbiddenAction(suggestions.Suggestions, "direct_external_release")),
				Summary:      "Policy suggestions remain advisory until formally approved and cannot auto-activate profiles or trigger direct release.",
				ReasonCodes:  []string{"advisory_until_formally_approved", "automatic_activation_forbidden"},
				RouteRefs:    []string{"/v1/formal/phase8/governance/policy-suggestions"},
				DocRefs:      []string{"docs/formal-phase8-valb.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "authority_routing_has_explicit_paths",
				CurrentState: phase8ReviewCheckState(routing.AlternateApproverPath != "" && routing.DeadlockResolutionPath != "" && routing.EmergencySuspensionPath != ""),
				Summary:      "Authority routing keeps explicit alternate approver, deadlock, and emergency suspension paths visible.",
				ReasonCodes:  []string{"alternate_approver_visible", "deadlock_resolution_visible", "emergency_suspension_visible"},
				RouteRefs:    []string{"/v1/formal/phase8/governance/authority-routing"},
				DocRefs:      []string{"docs/formal-phase8-valb.md", "docs/formal-phase8-final.md"},
			},
		},
		Limitations: []string{
			"Consensus governance review confirms bounded routing and advisory autonomy rather than a self-authorizing governance body.",
		},
	}
}

func buildPhase8AIGuardrailFinalReview() phase8FinalReviewSection {
	guardrails := buildPhase8AIGuardrails()
	modelRisk := buildPhase8ModelRisk()
	active := guardrails.CurrentState == phase8AIGuardrailsStateActive &&
		modelRisk.CurrentState == phase8ModelRiskStateActive &&
		hasForbiddenRecommendationClass(guardrails.Guardrails, "legal_verdict") &&
		hasForbiddenRecommendationClass(guardrails.Guardrails, "non_delegable_action_execution") &&
		hasDependencyClass(modelRisk.Dependencies, "regulatory_framework") &&
		modelRisk.ReviewState == "challenger_review_and_rollback_path_visible"
	return phase8FinalReviewSection{
		CurrentState: phase8SectionState(active, guardrails.CurrentState != phase8AIGuardrailsStateIncomplete && modelRisk.CurrentState != phase8ModelRiskStateIncomplete),
		RouteRefs: []string{
			"/v1/formal/phase8/governance/ai-guardrails",
			"/v1/formal/phase8/governance/model-risk",
			"/v1/formal/phase8/proofs",
		},
		DocRefs: []string{
			"docs/formal-phase8-valb.md",
			"docs/formal-phase8-final.md",
		},
		Checks: []phase8FinalReviewCheck{
			{
				CheckID:      "ai_cannot_issue_legal_verdicts",
				CurrentState: phase8ReviewCheckState(hasForbiddenRecommendationClass(guardrails.Guardrails, "legal_verdict")),
				Summary:      "AI guardrails explicitly prohibit legal verdict outputs in the formal-authority slice.",
				ReasonCodes:  []string{"legal_verdict_forbidden", "bounded_ai_support"},
				RouteRefs:    []string{"/v1/formal/phase8/governance/ai-guardrails"},
				DocRefs:      []string{"docs/formal-phase8-valb.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "ai_cannot_execute_non_delegable_actions",
				CurrentState: phase8ReviewCheckState(hasForbiddenRecommendationClass(guardrails.Guardrails, "non_delegable_action_execution")),
				Summary:      "AI and consensus layers remain blocked from executing non-delegable authority actions.",
				ReasonCodes:  []string{"non_delegable_execution_forbidden", "human_routing_required"},
				RouteRefs:    []string{"/v1/formal/phase8/governance/ai-guardrails", "/v1/formal/phase8/governance/authority-routing"},
				DocRefs:      []string{"docs/formal-phase8-valb.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "model_risk_and_dependencies_visible",
				CurrentState: phase8ReviewCheckState(modelRisk.ReviewState == "challenger_review_and_rollback_path_visible" && hasDependencyClass(modelRisk.Dependencies, "regulatory_framework")),
				Summary:      "Model risk governance keeps rollback, challenger review, and dependency change triggers visible.",
				ReasonCodes:  []string{"challenger_review_visible", "dependency_registry_visible"},
				RouteRefs:    []string{"/v1/formal/phase8/governance/model-risk"},
				DocRefs:      []string{"docs/formal-phase8-valb.md", "docs/formal-phase8-final.md"},
			},
		},
		Limitations: []string{
			"AI guardrail review confirms bounded governance support and does not turn models into a tribunal or hidden override authority.",
		},
	}
}

func buildPhase8InsurerEvidenceReview() phase8FinalReviewSection {
	risk := buildPhase8RiskQuantification()
	exports := buildPhase8InsuranceExports()
	incident := buildPhase8IncidentAttribution()
	benchmarks := buildPhase8ActuarialBenchmarks()
	exportBounded := len(exports.Exports) > 0 &&
		exports.Exports[0].Audience == formalcore.AudienceInsurer &&
		exports.Exports[0].ReleaseApprovalRequired &&
		!exports.Exports[0].CanCitePublicly
	aggregateBenchmarks := len(benchmarks.Items) > 0 &&
		benchmarks.Items[0].MinimumCohortSize >= 50 &&
		benchmarks.Items[0].AggregationScope == "aggregate_only_cross_tenant_safe_band" &&
		containsString(benchmarks.Items[0].ForbiddenUse, "tenant_level_pricing") &&
		containsString(benchmarks.Items[0].ForbiddenUse, "raw_subject_disclosure")
	incidentBounded := len(incident.Items) > 0 &&
		incident.NonLegalConclusionState == "non_legal_conclusion_marker_active" &&
		incident.Items[0].NonLegalConclusionMarker == "support_only_not_legal_conclusion"
	active := risk.CurrentState == phase8RiskQuantificationStateActive &&
		exports.CurrentState == phase8InsuranceExportsStateActive &&
		incident.CurrentState == phase8IncidentAttributionStateActive &&
		benchmarks.CurrentState == phase8ActuarialBenchmarksStateActive &&
		risk.PremiumModelBoundary == "never_automatic_pricing_promise" &&
		exportBounded &&
		incidentBounded &&
		aggregateBenchmarks
	return phase8FinalReviewSection{
		CurrentState: phase8SectionState(active, risk.CurrentState != phase8RiskQuantificationStateIncomplete && exports.CurrentState != phase8InsuranceExportsStateIncomplete && incident.CurrentState != phase8IncidentAttributionStateIncomplete && benchmarks.CurrentState != phase8ActuarialBenchmarksStateIncomplete),
		RouteRefs: []string{
			"/v1/formal/phase8/institutional/risk-quantification",
			"/v1/formal/phase8/institutional/insurance-exports",
			"/v1/formal/phase8/institutional/incident-attribution",
			"/v1/formal/phase8/institutional/actuarial-benchmarks",
		},
		DocRefs: []string{
			"docs/formal-phase8-valc.md",
			"docs/formal-phase8-final.md",
		},
		Checks: []phase8FinalReviewCheck{
			{
				CheckID:      "risk_quantification_not_pricing_promise",
				CurrentState: phase8ReviewCheckState(risk.PremiumModelBoundary == "never_automatic_pricing_promise"),
				Summary:      "Risk quantification remains integration-ready input and never becomes an automatic pricing promise.",
				ReasonCodes:  []string{"integration_ready_not_pricing_promise", "automatic_pricing_forbidden"},
				RouteRefs:    []string{"/v1/formal/phase8/institutional/risk-quantification"},
				DocRefs:      []string{"docs/formal-phase8-valc.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "insurer_exports_scoped_and_non_public",
				CurrentState: phase8ReviewCheckState(exportBounded),
				Summary:      "Insurer-facing exports remain insurer-scoped, release-approved, and non-public.",
				ReasonCodes:  []string{"insurer_scoped_export_only", "release_approval_required", "public_citation_blocked"},
				RouteRefs:    []string{"/v1/formal/phase8/institutional/insurance-exports"},
				DocRefs:      []string{"docs/formal-phase8-valc.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "incident_attribution_remains_non_legal",
				CurrentState: phase8ReviewCheckState(incidentBounded),
				Summary:      "Incident attribution support keeps ambiguity visible and remains explicitly non-legal.",
				ReasonCodes:  []string{"support_only_not_legal_conclusion", "ambiguity_visible"},
				RouteRefs:    []string{"/v1/formal/phase8/institutional/incident-attribution"},
				DocRefs:      []string{"docs/formal-phase8-valc.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "actuarial_benchmarks_stay_aggregate_only",
				CurrentState: phase8ReviewCheckState(aggregateBenchmarks),
				Summary:      "Actuarial benchmarks remain aggregate-only, privacy-guarded, and blocked from tenant-level pricing or raw disclosure.",
				ReasonCodes:  []string{"aggregate_only_cross_tenant_safe_band", "reidentification_guarded"},
				RouteRefs:    []string{"/v1/formal/phase8/institutional/actuarial-benchmarks"},
				DocRefs:      []string{"docs/formal-phase8-valc.md", "docs/formal-phase8-final.md"},
			},
		},
		Limitations: []string{
			"Insurer evidence review confirms institutional usefulness without widening into public authority, pricing automation, or legal adjudication.",
		},
	}
}

func buildPhase8FormalAuthorityBoundaryReview() phase8FinalReviewSection {
	entry := buildPhase8EntryGate()
	proofs := buildPhase8Proofs()
	remainingDeferred := phase8RemainingDeferredScope()
	active := entry.CurrentState == formalcore.EntryGateStateReady &&
		proofs.CurrentState == formalcore.Phase8StateActive &&
		proofs.CoverageScope == phase8CoverageScopeCorePass &&
		hasEntryGateLimitation(entry.CarryOverLimitations, "does not create a separate formal truth store") &&
		hasEntryGateLimitation(entry.CarryOverLimitations, "does not claim legal, regulatory, insurer, or certification-body replacement authority") &&
		containsString(remainingDeferred, "insurer_integration_program") &&
		containsString(remainingDeferred, "wider_federated_governance") &&
		containsString(remainingDeferred, "advanced_institutional_disclosure_programs")
	return phase8FinalReviewSection{
		CurrentState: phase8SectionState(active, proofs.CurrentState != formalcore.Phase8StateIncomplete),
		RouteRefs: []string{
			"/v1/formal/phase8/entry-gate",
			"/v1/formal/phase8/proofs",
			"/v1/formal/phase8/final-summary",
		},
		DocRefs: []string{
			"docs/formal-phase8-core.md",
			"docs/formal-phase8-final.md",
		},
		Checks: []phase8FinalReviewCheck{
			{
				CheckID:      "no_shadow_truth_store",
				CurrentState: phase8ReviewCheckState(hasEntryGateLimitation(entry.CarryOverLimitations, "does not create a separate formal truth store")),
				Summary:      "Phase 8 finalization stays on the existing evidence spine and does not introduce a second formal truth store.",
				ReasonCodes:  []string{"single_evidence_spine", "no_shadow_truth_store"},
				RouteRefs:    []string{"/v1/formal/phase8/entry-gate"},
				DocRefs:      []string{"docs/formal-phase8-core.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "no_replacement_authority_claim",
				CurrentState: phase8ReviewCheckState(hasEntryGateLimitation(entry.CarryOverLimitations, "does not claim legal, regulatory, insurer, or certification-body replacement authority")),
				Summary:      "Formal-authority outputs remain bounded support and do not claim replacement authority over insurers, regulators, or certifiers.",
				ReasonCodes:  []string{"bounded_formal_authority", "no_replacement_authority_claim"},
				RouteRefs:    []string{"/v1/formal/phase8/entry-gate", "/v1/formal/phase8/proofs"},
				DocRefs:      []string{"docs/formal-phase8-core.md", "docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "remaining_deferred_programs_visible",
				CurrentState: phase8ReviewCheckState(containsString(remainingDeferred, "insurer_integration_program") && containsString(remainingDeferred, "wider_federated_governance") && containsString(remainingDeferred, "advanced_institutional_disclosure_programs")),
				Summary:      "Programs beyond the bounded Phase 8 package remain explicitly deferred rather than hidden behind optimistic completion language.",
				ReasonCodes:  []string{"remaining_deferred_scope_visible", "no_silent_expansion"},
				RouteRefs:    []string{"/v1/formal/phase8/final-summary"},
				DocRefs:      []string{"docs/formal-phase8-final.md"},
			},
			{
				CheckID:      "core_pass_proofs_remain_fail_closed",
				CurrentState: phase8ReviewCheckState(proofs.CurrentState == formalcore.Phase8StateActive && proofs.CoverageScope == phase8CoverageScopeCorePass),
				Summary:      "Core proofs remain fail-closed and keep formal, compliance, and governance activation aligned under core-pass semantics.",
				ReasonCodes:  []string{"fail_closed_core_proofs", "core_pass_scope"},
				RouteRefs:    []string{"/v1/formal/phase8/proofs"},
				DocRefs:      []string{"docs/formal-phase8-core.md", "docs/formal-phase8-final.md"},
			},
		},
		Limitations: []string{
			"Formal-authority boundary review confirms bounded completion of Phase 8 and does not claim authority or institutional programs beyond this pack.",
		},
	}
}

func phase8FinalizationState(phase8State string, sections ...phase8FinalReviewSection) string {
	if phase8State != formalcore.Phase8StateActive {
		return phase8FinalizationStateIncomplete
	}
	hasPartial := false
	for _, section := range sections {
		switch section.CurrentState {
		case phase8FinalReviewStateActive:
		case phase8FinalReviewStatePartial:
			hasPartial = true
		default:
			return phase8FinalizationStateIncomplete
		}
	}
	if hasPartial {
		return phase8FinalizationStateSubstantial
	}
	return phase8FinalizationStateReady
}

func phase8SectionState(active bool, hasCoverage bool) string {
	switch {
	case active:
		return phase8FinalReviewStateActive
	case hasCoverage:
		return phase8FinalReviewStatePartial
	default:
		return phase8FinalReviewStateIncomplete
	}
}

func phase8ReviewCheckState(active bool) string {
	if active {
		return phase8FinalReviewStateActive
	}
	return phase8FinalReviewStateIncomplete
}

func phase8RemainingDeferredScope() []string {
	return []string{
		"insurer_integration_program",
		"wider_federated_governance",
		"advanced_institutional_disclosure_programs",
	}
}

func phase8ClaimClassByID(target string) formalcore.ClaimClass {
	for _, claim := range formalcore.ClaimClasses() {
		if claim.ClaimClass == target {
			return claim
		}
	}
	return formalcore.ClaimClass{}
}

func phase8UsePermissionRuleByAudience(target string) formalcore.UsePermissionRule {
	for _, rule := range formalcore.UsePermissionRules() {
		if rule.Audience == target {
			return rule
		}
	}
	return formalcore.UsePermissionRule{}
}

func phase8ProofClassByID(target string) formalcore.StandardOfProofClass {
	for _, proofClass := range formalcore.StandardOfProofClasses() {
		if proofClass.ProofClass == target {
			return proofClass
		}
	}
	return formalcore.StandardOfProofClass{}
}

func hasEvidenceAutomationExclusion(items []phase8EvidenceAutomationExclusion, itemID string, deferred bool) bool {
	for _, item := range items {
		if item.ItemID == itemID && item.DeferredScope == deferred {
			return true
		}
	}
	return false
}

func containsForbiddenAction(items []phase8PolicySuggestion, target string) bool {
	for _, item := range items {
		if containsString(item.ForbiddenActions, target) {
			return true
		}
	}
	return false
}

func hasEntryGateLimitation(items []string, target string) bool {
	return containsSubstring(items, target)
}
