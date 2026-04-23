package claims

import "testing"

func TestMeasuredPublicProofValEStateRequiresActiveValD(t *testing.T) {
	got := EvaluateMeasuredPublicProofValEState(
		MeasuredPublicProofValDStateSubstantial,
		MeasuredPublicProofValEReplayCorrectnessReviewStateActive,
		MeasuredPublicProofValESigningTrustReviewStateActive,
		MeasuredPublicProofValETransparencyReviewStateActive,
		MeasuredPublicProofValERedactionReviewStateActive,
		MeasuredPublicProofValECompatibilityReviewStateActive,
		MeasuredPublicProofValEIssuanceReviewStateActive,
		MeasuredPublicProofValEFailureStateReviewStateActive,
	)
	if got != MeasuredPublicProofValEStateIncomplete {
		t.Fatalf("expected incomplete vale state without active vald, got %q", got)
	}
}

func TestMeasuredPublicProofValEReplayCorrectnessReviewIsPartialWithoutToleranceBands(t *testing.T) {
	items := []PublicProofReplayCorrectnessReviewItem{{
		ClaimID:                     "point2_runtime_performance_claim",
		ArtifactID:                  "point2_runtime_performance_public_pack",
		CurrentState:                "replay_review_ready",
		ReviewOutcome:               "approved",
		ReplayState:                 "comparison_verified",
		ComparisonMode:              "bounded_local_vs_public_comparison",
		MethodologyRef:              "/v1/public/benchmarks/methodology",
		EvaluationRef:               "/v1/foundation/execution/benchmarks/evaluate",
		ToleranceDecision:           "within_declared_bands",
		SupportedEnvironmentClasses: []string{"runtime_hardened_enterprise_cluster"},
		UnsupportedReplayCases:      []string{"cross-provider replay outside declared environment class"},
		ReviewRefs:                  []string{"/v1/public/proof-expansion/valb/replay-verification"},
		EvidenceRefs:                []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
	}}
	if got := EvaluateMeasuredPublicProofValEReplayCorrectnessReviewState(items); got != MeasuredPublicProofValEReplayCorrectnessReviewStatePartial {
		t.Fatalf("expected partial replay correctness review without tolerance bands, got %q", got)
	}
}

func TestMeasuredPublicProofValESigningTrustReviewIsPartialWithoutTrustRootState(t *testing.T) {
	items := []PublicProofSigningTrustReviewItem{{
		ClaimID:                     "point2_runtime_performance_claim",
		ArtifactID:                  "point2_runtime_performance_public_pack",
		CurrentState:                "signing_trust_review_ready",
		ReviewOutcome:               "approved",
		VerificationState:           "verified",
		SigningPurposeState:         "purpose_enabled",
		HistoricalVerificationState: "historical_verification_ready",
		KeyRotationState:            "rotation_ready",
		RevocationState:             "revocation_ready",
		TimestampState:              "timestamp_bound",
		SignerMode:                  "software",
		TrustRootID:                 "public_proof_primary_root",
		KeyVersion:                  "v1",
		ReviewRefs:                  []string{"/v1/public/proof-expansion/val0/signing-authority"},
		EvidenceRefs:                []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
		FailureStates:               []string{"failed_verification", "trust_root_unavailable"},
	}}
	if got := EvaluateMeasuredPublicProofValESigningTrustReviewState(items); got != MeasuredPublicProofValESigningTrustReviewStatePartial {
		t.Fatalf("expected partial signing trust review without trust root state, got %q", got)
	}
}

func TestMeasuredPublicProofValEFailureStateReviewIsPartialWithoutRequiredFailureStates(t *testing.T) {
	items := []PublicProofFailureStateReviewItem{{
		ClaimID:                     "point2_verification_reference_claim",
		ArtifactID:                  "point2_verification_public_pack",
		CurrentState:                "failure_state_review_ready",
		ReviewOutcome:               "approved",
		FailureVisibilityState:      "verifier_visible_failures_modeled",
		RestrictionVisibilityState:  "restriction_visible",
		WithdrawalVisibilityState:   "withdrawal_visible_not_triggered",
		SupersessionVisibilityState: "supersession_visible_not_triggered",
		ReissueFailureState:         "claim_not_reissued_modeled",
		SupportedFailureStates:      []string{PublicProofStatusProofFailed, PublicProofStatusRestricted},
		ReviewRefs:                  []string{"/v1/public/proof-expansion/vald/claim-lifecycle"},
		EvidenceRefs:                []string{"/v1/public/proof-expansion/vald/correction-workflow"},
		FailureStates:               []string{"proof_generation_failed"},
	}}
	if got := EvaluateMeasuredPublicProofValEFailureStateReviewState(items); got != MeasuredPublicProofValEFailureStateReviewStatePartial {
		t.Fatalf("expected partial failure-state review without required failure states, got %q", got)
	}
}

func TestMeasuredPublicProofValEFoundationIsActive(t *testing.T) {
	replayItems := []PublicProofReplayCorrectnessReviewItem{{
		ClaimID:                     "point2_runtime_performance_claim",
		ArtifactID:                  "point2_runtime_performance_public_pack",
		CurrentState:                "replay_review_ready",
		ReviewOutcome:               "approved",
		ReplayState:                 "comparison_verified",
		ComparisonMode:              "bounded_local_vs_public_comparison",
		MethodologyRef:              "/v1/public/benchmarks/methodology",
		EvaluationRef:               "/v1/foundation/execution/benchmarks/evaluate",
		ToleranceDecision:           "within_declared_bands",
		SupportedEnvironmentClasses: []string{"runtime_hardened_enterprise_cluster"},
		ToleranceBands:              []string{"capture_p99<=400us", "correlation_p99<=700us"},
		UnsupportedReplayCases:      []string{"cross-provider replay outside declared environment class"},
		ReviewRefs:                  []string{"/v1/public/proof-expansion/valb/replay-verification"},
		EvidenceRefs:                []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
	}}
	if got := EvaluateMeasuredPublicProofValEReplayCorrectnessReviewState(replayItems); got != MeasuredPublicProofValEReplayCorrectnessReviewStateActive {
		t.Fatalf("expected active replay correctness review state, got %q", got)
	}

	signingItems := []PublicProofSigningTrustReviewItem{{
		ClaimID:                     "point2_runtime_performance_claim",
		ArtifactID:                  "point2_runtime_performance_public_pack",
		CurrentState:                "signing_trust_review_ready",
		ReviewOutcome:               "approved",
		VerificationState:           "verified",
		TrustRootState:              "trusted",
		SigningPurposeState:         "purpose_enabled",
		HistoricalVerificationState: "historical_verification_ready",
		KeyRotationState:            "rotation_ready",
		RevocationState:             "revocation_ready",
		TimestampState:              "timestamp_bound",
		SignerMode:                  "software",
		TrustRootID:                 "public_proof_primary_root",
		KeyVersion:                  "v1",
		ReviewRefs:                  []string{"/v1/public/proof-expansion/val0/signing-authority"},
		EvidenceRefs:                []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
		FailureStates:               []string{"failed_verification", "trust_root_unavailable"},
	}}
	if got := EvaluateMeasuredPublicProofValESigningTrustReviewState(signingItems); got != MeasuredPublicProofValESigningTrustReviewStateActive {
		t.Fatalf("expected active signing trust review state, got %q", got)
	}

	transparencyItems := []PublicProofTransparencyReviewItem{{
		ClaimID:                "point2_runtime_performance_claim",
		ArtifactID:             "point2_runtime_performance_public_pack",
		CurrentState:           "transparency_review_ready",
		ReviewOutcome:          "approved",
		TransparencyState:      "anchored_projection_ready",
		EntryHashState:         "digest_bound",
		AnchorState:            "anchor_active",
		SupersessionVisibility: "visible_not_superseded",
		AnchorRef:              "/v1/public/transparency/anchor",
		EntryID:                "point2_valb_transparency_chain_v1/point2_runtime_performance_public_pack",
		EntryHash:              "sha256:abc",
		LineageRef:             "/v1/public/proof-expansion/valc/claim-lineage",
		ReviewRefs:             []string{"/v1/public/proof-expansion/valb/transparency-chain"},
		EvidenceRefs:           []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
		FailureStates:          []string{"anchoring_unavailable", PublicProofStatusWithdrawn},
	}}
	if got := EvaluateMeasuredPublicProofValETransparencyReviewState(transparencyItems); got != MeasuredPublicProofValETransparencyReviewStateActive {
		t.Fatalf("expected active transparency review state, got %q", got)
	}

	redactionItems := []PublicProofRedactionReviewItem{{
		ClaimID:               "point2_runtime_performance_claim",
		ArtifactID:            "point2_runtime_performance_public_pack",
		CurrentState:          "redaction_review_ready",
		ReviewOutcome:         "approved",
		RedactionTier:         RedactionTierPublicSafe,
		PublicationScope:      ScopePublic,
		PortalProjectionState: "portal_projection_ready",
		RedactionDecision:     "tier_scope_match",
		ProjectionDiscipline:  "projection_only_enforced",
		RemovedFields:         []string{"tenant_sensitive_raw_events"},
		NeverPublishedFields:  []string{"tenant_sensitive_raw_events"},
		ReviewRefs:            []string{"/v1/public/proof-expansion/val0/redaction-tiers"},
		EvidenceRefs:          []string{"/v1/public/proof-expansion/valc/public-proof-portal"},
		FailureStates:         []string{"claim_restricted", "claim_withdrawn"},
	}}
	if got := EvaluateMeasuredPublicProofValERedactionReviewState(redactionItems); got != MeasuredPublicProofValERedactionReviewStateActive {
		t.Fatalf("expected active redaction review state, got %q", got)
	}

	compatibilityItems := []PublicProofCompatibilityReviewItem{{
		ClaimID:                  "point2_runtime_performance_claim",
		ArtifactID:               "point2_runtime_performance_public_pack",
		CurrentState:             "compatibility_review_ready",
		ReviewOutcome:            "approved",
		SchemaCompatibility:      "supported",
		VerifierCompatibility:    "supported",
		DeprecationState:         "not_deprecated",
		ReplayCompatibility:      "supported",
		MethodologyCompatibility: "bounded_supported",
		SupportedSchemaLines:     []string{"public.proof.sealed_artifact.v1", "/v1/public/verifier/sdk"},
		UnsupportedCases:         []string{"cross-schema replay without compatibility declaration"},
		ReviewRefs:               []string{"/v1/public/proof-expansion/val0/compatibility-baseline"},
		EvidenceRefs:             []string{"/v1/public/proof-expansion/valb/replay-verification"},
		FailureStates:            []string{"proof_generation_failed", "freshness_expired"},
	}}
	if got := EvaluateMeasuredPublicProofValECompatibilityReviewState(compatibilityItems); got != MeasuredPublicProofValECompatibilityReviewStateActive {
		t.Fatalf("expected active compatibility review state, got %q", got)
	}

	issuanceItems := []PublicProofIssuanceReviewItem{{
		ClaimID:                 "point2_verification_reference_claim",
		ArtifactID:              "point2_verification_public_pack",
		CurrentState:            "issuance_review_ready",
		ReviewOutcome:           "approved",
		ReleaseIssuanceState:    "issuance_gate_ready",
		ClaimLifecycleStatus:    PublicProofStatusRestricted,
		PublicationDecision:     "restricted_partner_scoped_reissue",
		CorrectionWorkflowState: "correction_workflow_ready",
		OverrideState:           "override_not_permitted",
		ReviewRefs:              []string{"/v1/public/proof-expansion/vald/release-issuance-gate"},
		EvidenceRefs:            []string{"/v1/public/proof-expansion/vald/claim-lifecycle"},
		FailureStates:           []string{"claim_not_reissued", "claim_restricted"},
	}}
	if got := EvaluateMeasuredPublicProofValEIssuanceReviewState(issuanceItems); got != MeasuredPublicProofValEIssuanceReviewStateActive {
		t.Fatalf("expected active issuance review state, got %q", got)
	}

	failureStateItems := []PublicProofFailureStateReviewItem{{
		ClaimID:                     "point2_runtime_performance_claim",
		ArtifactID:                  "point2_runtime_performance_public_pack",
		CurrentState:                "failure_state_review_ready",
		ReviewOutcome:               "approved",
		FailureVisibilityState:      "verifier_visible_failures_modeled",
		RestrictionVisibilityState:  "restriction_visible",
		WithdrawalVisibilityState:   "withdrawal_visible_not_triggered",
		SupersessionVisibilityState: "supersession_visible_not_triggered",
		ReissueFailureState:         "claim_not_reissued_modeled",
		SupportedFailureStates: []string{
			PublicProofStatusProofFailed,
			PublicProofStatusClaimNotReissued,
			PublicProofStatusRestricted,
			PublicProofStatusSuperseded,
			PublicProofStatusWithdrawn,
			PublicProofStatusStale,
		},
		ReviewRefs:   []string{"/v1/public/proof-expansion/vald/claim-lifecycle"},
		EvidenceRefs: []string{"/v1/public/proof-expansion/vald/correction-workflow"},
		FailureStates: []string{
			"proof_generation_failed",
			"signature_failed",
			"claim_restricted",
		},
	}}
	if got := EvaluateMeasuredPublicProofValEFailureStateReviewState(failureStateItems); got != MeasuredPublicProofValEFailureStateReviewStateActive {
		t.Fatalf("expected active failure-state review state, got %q", got)
	}

	if got := EvaluateMeasuredPublicProofValEState(
		MeasuredPublicProofValDStateActive,
		MeasuredPublicProofValEReplayCorrectnessReviewStateActive,
		MeasuredPublicProofValESigningTrustReviewStateActive,
		MeasuredPublicProofValETransparencyReviewStateActive,
		MeasuredPublicProofValERedactionReviewStateActive,
		MeasuredPublicProofValECompatibilityReviewStateActive,
		MeasuredPublicProofValEIssuanceReviewStateActive,
		MeasuredPublicProofValEFailureStateReviewStateActive,
	); got != MeasuredPublicProofValEStateActive {
		t.Fatalf("expected active vale state, got %q", got)
	}
}
