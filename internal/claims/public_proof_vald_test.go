package claims

import "testing"

func TestMeasuredPublicProofValDStateRequiresActiveValC(t *testing.T) {
	got := EvaluateMeasuredPublicProofValDState(
		MeasuredPublicProofValCStateSubstantial,
		MeasuredPublicProofValDReleaseIssuanceStateActive,
		MeasuredPublicProofValDClaimLifecycleStateActive,
		MeasuredPublicProofValDPublicationDecisionStateActive,
		MeasuredPublicProofValDCorrectionWorkflowStateActive,
	)
	if got != MeasuredPublicProofValDStateIncomplete {
		t.Fatalf("expected incomplete vald state without active valc, got %q", got)
	}
}

func TestMeasuredPublicProofValDReleaseIssuanceIsPartialWithoutAuditRefs(t *testing.T) {
	items := []PublicProofReleaseIssuanceItem{{
		ClaimID:             "point2_runtime_performance_claim",
		ArtifactID:          "point2_runtime_performance_public_pack",
		CurrentState:        "issuance_gate_ready",
		ReleaseID:           "point2-2026-04-23-runtime",
		BuildIdentity:       "changelock-public-proof-2026.04.23",
		ReleaseChannel:      "public_proof_expansion",
		PriorReleaseRef:     "/v1/public/phase6/proofs?as_of=2026-04-23T10:00:00Z",
		ReissueDecision:     "reissued",
		PublicationDecision: "approved_public_safe_reissue",
		RequiredChecks:      []string{"signature_verified", "replay_verified"},
		SatisfiedChecks:     []string{"signature_verified", "replay_verified"},
		VerificationRefs:    []string{"/v1/public/proof-expansion/valb/signature-verification"},
		FailureStates:       []string{PublicProofStatusProofFailed, PublicProofStatusClaimNotReissued},
	}}
	if got := EvaluateMeasuredPublicProofValDReleaseIssuanceState(items); got != MeasuredPublicProofValDReleaseIssuanceStatePartial {
		t.Fatalf("expected partial release issuance without audit refs, got %q", got)
	}
}

func TestMeasuredPublicProofValDPublicationDecisionIsPartialWhenOverrideIsPermitted(t *testing.T) {
	items := []PublicProofPublicationDecisionItem{{
		ClaimID:           "point2_runtime_performance_claim",
		ArtifactID:        "point2_runtime_performance_public_pack",
		CurrentState:      "publication_decision_ready",
		PublicationStatus: "approved_public_safe_reissue",
		ApprovalBoundary:  "signature_replay_and_redaction_review_required",
		RedactionTier:     RedactionTierPublicSafe,
		PublicationScope:  ScopePublic,
		AutomationState:   "auto_issue_ready_no_override",
		OverridePermitted: true,
		DecisionAuditRefs: []string{"/v1/public/proof-expansion/valc/claim-lineage"},
		ProjectionRefs:    []string{"/v1/public/proof-expansion/valc/public-proof-portal"},
		FailureStates:     []string{PublicProofStatusProofFailed, PublicProofStatusWithdrawn},
	}}
	if got := EvaluateMeasuredPublicProofValDPublicationDecisionState(items); got != MeasuredPublicProofValDPublicationDecisionStatePartial {
		t.Fatalf("expected partial publication decision when override is permitted, got %q", got)
	}
}

func TestMeasuredPublicProofValDClaimLifecycleIsPartialWithoutRequiredLifecycleStates(t *testing.T) {
	items := []PublicProofClaimLifecycleItem{{
		ClaimID:                  "point2_verification_reference_claim",
		ArtifactID:               "point2_verification_public_pack",
		CurrentState:             "claim_lifecycle_governed",
		ClaimStatus:              PublicProofStatusRestricted,
		ReissueState:             "reissued_for_current_release",
		FreshnessState:           FreshnessFresh,
		PublicationScope:         ScopePartner,
		RestrictionState:         "restricted_to_partner_scope",
		WithdrawalState:          "not_withdrawn",
		SupersessionState:        "not_superseded",
		SupportedLifecycleStates: []string{PublicProofStatusProven, PublicProofStatusRestricted},
		VerifierNoticeRefs:       []string{"/v1/public/proof-expansion/valb/proofs"},
		PortalRefs:               []string{"/v1/public/proof-expansion/valc/partner-proof-portal"},
		EvidenceRefs:             []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
	}}
	if got := EvaluateMeasuredPublicProofValDClaimLifecycleState(items); got != MeasuredPublicProofValDClaimLifecycleStatePartial {
		t.Fatalf("expected partial claim lifecycle without required lifecycle states, got %q", got)
	}
}

func TestMeasuredPublicProofValDFoundationIsActive(t *testing.T) {
	releaseItems := []PublicProofReleaseIssuanceItem{
		{
			ClaimID:             "point2_runtime_performance_claim",
			ArtifactID:          "point2_runtime_performance_public_pack",
			CurrentState:        "issuance_gate_ready",
			ReleaseID:           "point2-2026-04-23-runtime",
			BuildIdentity:       "changelock-public-proof-2026.04.23",
			ReleaseChannel:      "public_proof_expansion",
			PriorReleaseRef:     "/v1/public/phase6/proofs?as_of=2026-04-23T10:00:00Z",
			ReissueDecision:     "reissued",
			PublicationDecision: "approved_public_safe_reissue",
			RequiredChecks:      []string{"signature_verified", "replay_verified", "transparency_verified"},
			SatisfiedChecks:     []string{"signature_verified", "replay_verified", "transparency_verified"},
			VerificationRefs:    []string{"/v1/public/proof-expansion/valb/signature-verification", "/v1/public/proof-expansion/valb/replay-verification"},
			AuditRefs:           []string{"/v1/public/proof-expansion/valc/claim-lineage", "/v1/public/proof-expansion/valb/transparency-chain"},
			FailureStates:       []string{PublicProofStatusProofFailed, PublicProofStatusClaimNotReissued, PublicProofStatusWithdrawn},
		},
		{
			ClaimID:             "point2_verification_reference_claim",
			ArtifactID:          "point2_verification_public_pack",
			CurrentState:        "issuance_gate_ready",
			ReleaseID:           "point2-2026-04-23-verification",
			BuildIdentity:       "changelock-public-proof-2026.04.23",
			ReleaseChannel:      "public_proof_expansion",
			PriorReleaseRef:     "/v1/public/phase6/proofs?as_of=2026-04-23T10:00:00Z",
			ReissueDecision:     "restricted_reissue_ready",
			PublicationDecision: "restricted_partner_scoped_reissue",
			RequiredChecks:      []string{"signature_verified", "replay_verified", "partner_redaction_review"},
			SatisfiedChecks:     []string{"signature_verified", "replay_verified", "partner_redaction_review"},
			VerificationRefs:    []string{"/v1/public/proof-expansion/valb/signature-verification", "/v1/public/proof-expansion/valb/replay-verification"},
			AuditRefs:           []string{"/v1/public/proof-expansion/valc/claim-lineage", "/v1/public/proof-expansion/valb/transparency-chain"},
			FailureStates:       []string{PublicProofStatusRestricted, PublicProofStatusProofFailed, PublicProofStatusWithdrawn},
		},
	}
	if got := EvaluateMeasuredPublicProofValDReleaseIssuanceState(releaseItems); got != MeasuredPublicProofValDReleaseIssuanceStateActive {
		t.Fatalf("expected active release issuance state, got %q", got)
	}

	lifecycleItems := []PublicProofClaimLifecycleItem{
		{
			ClaimID:                  "point2_runtime_performance_claim",
			ArtifactID:               "point2_runtime_performance_public_pack",
			CurrentState:             "claim_lifecycle_governed",
			ClaimStatus:              PublicProofStatusProven,
			ReissueState:             "reissued_for_current_release",
			FreshnessState:           FreshnessFresh,
			PublicationScope:         ScopePublic,
			RestrictionState:         "unrestricted_public_safe",
			WithdrawalState:          "not_withdrawn",
			SupersessionState:        "not_superseded",
			SupportedLifecycleStates: []string{PublicProofStatusProven, PublicProofStatusRestricted, PublicProofStatusSuperseded, PublicProofStatusWithdrawn, PublicProofStatusClaimNotReissued},
			VerifierNoticeRefs:       []string{"/v1/public/proof-expansion/valb/proofs"},
			PortalRefs:               []string{"/v1/public/proof-expansion/valc/public-proof-portal", "/v1/public/proof-expansion/valc/claim-lineage"},
			EvidenceRefs:             []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
		},
		{
			ClaimID:                  "point2_verification_reference_claim",
			ArtifactID:               "point2_verification_public_pack",
			CurrentState:             "claim_lifecycle_governed",
			ClaimStatus:              PublicProofStatusRestricted,
			ReissueState:             "reissued_for_current_release",
			FreshnessState:           FreshnessFresh,
			PublicationScope:         ScopePartner,
			RestrictionState:         "restricted_to_partner_scope",
			WithdrawalState:          "not_withdrawn",
			SupersessionState:        "not_superseded",
			SupportedLifecycleStates: []string{PublicProofStatusProven, PublicProofStatusRestricted, PublicProofStatusSuperseded, PublicProofStatusWithdrawn, PublicProofStatusClaimNotReissued},
			VerifierNoticeRefs:       []string{"/v1/public/proof-expansion/valb/proofs"},
			PortalRefs:               []string{"/v1/public/proof-expansion/valc/partner-proof-portal", "/v1/public/proof-expansion/valc/claim-lineage"},
			EvidenceRefs:             []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
		},
	}
	if got := EvaluateMeasuredPublicProofValDClaimLifecycleState(lifecycleItems); got != MeasuredPublicProofValDClaimLifecycleStateActive {
		t.Fatalf("expected active claim lifecycle state, got %q", got)
	}

	publicationItems := []PublicProofPublicationDecisionItem{
		{
			ClaimID:           "point2_runtime_performance_claim",
			ArtifactID:        "point2_runtime_performance_public_pack",
			CurrentState:      "publication_decision_ready",
			PublicationStatus: "approved_public_safe_reissue",
			ApprovalBoundary:  "signature_replay_and_redaction_review_required",
			RedactionTier:     RedactionTierPublicSafe,
			PublicationScope:  ScopePublic,
			AutomationState:   "auto_issue_ready_no_override",
			DecisionAuditRefs: []string{"/v1/public/proof-expansion/valc/claim-lineage", "/v1/public/proof-expansion/valb/proofs"},
			ProjectionRefs:    []string{"/v1/public/proof-expansion/valc/public-proof-portal", "/v1/public/proof-expansion/valc/download-projections"},
			FailureStates:     []string{PublicProofStatusProofFailed, PublicProofStatusClaimNotReissued, PublicProofStatusWithdrawn},
		},
		{
			ClaimID:           "point2_verification_reference_claim",
			ArtifactID:        "point2_verification_public_pack",
			CurrentState:      "publication_decision_ready",
			PublicationStatus: "restricted_partner_scoped_reissue",
			ApprovalBoundary:  "signature_replay_and_redaction_review_required",
			RedactionTier:     RedactionTierPartnerScoped,
			PublicationScope:  ScopePartner,
			AutomationState:   "auto_issue_restricted_no_override",
			DecisionAuditRefs: []string{"/v1/public/proof-expansion/valc/claim-lineage", "/v1/public/proof-expansion/valb/proofs"},
			ProjectionRefs:    []string{"/v1/public/proof-expansion/valc/partner-proof-portal", "/v1/public/proof-expansion/valc/download-projections"},
			FailureStates:     []string{PublicProofStatusRestricted, PublicProofStatusProofFailed, PublicProofStatusWithdrawn},
		},
	}
	if got := EvaluateMeasuredPublicProofValDPublicationDecisionState(publicationItems); got != MeasuredPublicProofValDPublicationDecisionStateActive {
		t.Fatalf("expected active publication decision state, got %q", got)
	}

	correctionItems := []PublicProofCorrectionWorkflowItem{
		{
			ClaimID:               "point2_runtime_performance_claim",
			ArtifactID:            "point2_runtime_performance_public_pack",
			CurrentState:          "correction_workflow_ready",
			TriggerClass:          "release_regression_or_incident",
			TriggerState:          "monitoring_ready",
			RestrictionActionRef:  "/v1/public/proof-expansion/vald/publication-decisions",
			WithdrawalActionRef:   "/v1/public/proof-expansion/vald/claim-lifecycle",
			SupersessionActionRef: "/v1/public/proof-expansion/valc/claim-lineage",
			CorrectionNoticeRef:   "/v1/public/proof-expansion/valc/public-proof-portal",
			AuditRefs:             []string{"/v1/public/proof-expansion/valb/transparency-chain", "/v1/public/proof-expansion/valc/claim-lineage"},
			VerifierRefs:          []string{"/v1/public/proof-expansion/valb/signature-verification", "/v1/public/proof-expansion/valb/replay-verification"},
			PortalNoticeRefs:      []string{"/v1/public/proof-expansion/valc/public-proof-portal"},
			FailureStates:         []string{PublicProofStatusProofFailed, PublicProofStatusRestricted, PublicProofStatusWithdrawn, PublicProofStatusSuperseded},
		},
		{
			ClaimID:               "point2_verification_reference_claim",
			ArtifactID:            "point2_verification_public_pack",
			CurrentState:          "correction_workflow_ready",
			TriggerClass:          "release_regression_or_incident",
			TriggerState:          "monitoring_ready",
			RestrictionActionRef:  "/v1/public/proof-expansion/vald/publication-decisions",
			WithdrawalActionRef:   "/v1/public/proof-expansion/vald/claim-lifecycle",
			SupersessionActionRef: "/v1/public/proof-expansion/valc/claim-lineage",
			CorrectionNoticeRef:   "/v1/public/proof-expansion/valc/partner-proof-portal",
			AuditRefs:             []string{"/v1/public/proof-expansion/valb/transparency-chain", "/v1/public/proof-expansion/valc/claim-lineage"},
			VerifierRefs:          []string{"/v1/public/proof-expansion/valb/signature-verification", "/v1/public/proof-expansion/valb/replay-verification"},
			PortalNoticeRefs:      []string{"/v1/public/proof-expansion/valc/partner-proof-portal"},
			FailureStates:         []string{PublicProofStatusProofFailed, PublicProofStatusRestricted, PublicProofStatusWithdrawn, PublicProofStatusSuperseded},
		},
	}
	if got := EvaluateMeasuredPublicProofValDCorrectionWorkflowState(correctionItems); got != MeasuredPublicProofValDCorrectionWorkflowStateActive {
		t.Fatalf("expected active correction workflow state, got %q", got)
	}

	if got := EvaluateMeasuredPublicProofValDState(
		MeasuredPublicProofValCStateActive,
		MeasuredPublicProofValDReleaseIssuanceStateActive,
		MeasuredPublicProofValDClaimLifecycleStateActive,
		MeasuredPublicProofValDPublicationDecisionStateActive,
		MeasuredPublicProofValDCorrectionWorkflowStateActive,
	); got != MeasuredPublicProofValDStateActive {
		t.Fatalf("expected active vald state, got %q", got)
	}
}
