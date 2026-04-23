package claims

import "testing"

func TestMeasuredPublicProofValCStateRequiresActiveValB(t *testing.T) {
	got := EvaluateMeasuredPublicProofValCState(
		MeasuredPublicProofValBStateSubstantial,
		MeasuredPublicProofValCPublicPortalStateActive,
		MeasuredPublicProofValCPartnerPortalStateActive,
		MeasuredPublicProofValCClaimLineageStateActive,
		MeasuredPublicProofValCDownloadProjectionStateActive,
	)
	if got != MeasuredPublicProofValCStateIncomplete {
		t.Fatalf("expected incomplete valc state without active valb, got %q", got)
	}
}

func TestMeasuredPublicProofValCPublicPortalIsPartialWithoutVerificationRef(t *testing.T) {
	items := []PublicProofPortalProjectionItem{{
		ClaimID:         "point2_runtime_performance_claim",
		ArtifactID:      "point2_runtime_performance_public_pack",
		CurrentState:    "portal_projection_ready",
		ClaimClass:      PublicProofClaimClassPerformance,
		Scope:           ScopePublic,
		VisibilityState: VisibilityPublicSafe,
		FreshnessState:  FreshnessFresh,
		MethodologyRef:  "/v1/public/benchmarks/methodology",
		DownloadRef:     "/v1/public/proof-expansion/vala/downloadable-packs/point2_runtime_performance_public_pack?as_of=2026-04-22T10:00:00Z",
		ReplayRef:       "/v1/public/proof-expansion/valb/replay-verification",
		LineageRef:      "/v1/public/proof-expansion/valc/claim-lineage",
		EvidenceRefs:    []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
	}}
	if got := EvaluateMeasuredPublicProofValCPublicPortalState(items); got != MeasuredPublicProofValCPublicPortalStatePartial {
		t.Fatalf("expected partial public portal without verification ref, got %q", got)
	}
}

func TestMeasuredPublicProofValCFoundationIsActive(t *testing.T) {
	publicItems := []PublicProofPortalProjectionItem{{
		ClaimID:         "point2_runtime_performance_claim",
		ArtifactID:      "point2_runtime_performance_public_pack",
		CurrentState:    "portal_projection_ready",
		ClaimClass:      PublicProofClaimClassPerformance,
		Scope:           ScopePublic,
		VisibilityState: VisibilityPublicSafe,
		FreshnessState:  FreshnessFresh,
		MethodologyRef:  "/v1/public/benchmarks/methodology",
		DownloadRef:     "/v1/public/proof-expansion/vala/downloadable-packs/point2_runtime_performance_public_pack?as_of=2026-04-22T10:00:00Z",
		VerificationRef: "/v1/public/proof-expansion/valb/signature-verification",
		ReplayRef:       "/v1/public/proof-expansion/valb/replay-verification",
		LineageRef:      "/v1/public/proof-expansion/valc/claim-lineage",
		EvidenceRefs:    []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
	}}
	if got := EvaluateMeasuredPublicProofValCPublicPortalState(publicItems); got != MeasuredPublicProofValCPublicPortalStateActive {
		t.Fatalf("expected active public portal state, got %q", got)
	}

	partnerItems := []PublicProofPortalProjectionItem{
		{
			ClaimID:         "point2_runtime_performance_claim",
			ArtifactID:      "point2_runtime_performance_public_pack",
			CurrentState:    "portal_projection_ready",
			ClaimClass:      PublicProofClaimClassPerformance,
			Scope:           ScopePartner,
			VisibilityState: VisibilityPartnerSafe,
			FreshnessState:  FreshnessFresh,
			MethodologyRef:  "/v1/public/benchmarks/methodology",
			DownloadRef:     "/v1/public/proof-expansion/vala/downloadable-packs/point2_runtime_performance_public_pack?as_of=2026-04-22T10:00:00Z",
			VerificationRef: "/v1/public/proof-expansion/valb/signature-verification",
			ReplayRef:       "/v1/public/proof-expansion/valb/replay-verification",
			LineageRef:      "/v1/public/proof-expansion/valc/claim-lineage",
			EvidenceRefs:    []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
		},
		{
			ClaimID:         "point2_verification_reference_claim",
			ArtifactID:      "point2_verification_public_pack",
			CurrentState:    "portal_projection_ready",
			ClaimClass:      PublicProofClaimClassVerification,
			Scope:           ScopePartner,
			VisibilityState: VisibilityPartnerSafe,
			FreshnessState:  FreshnessFresh,
			MethodologyRef:  "/v1/public/benchmarks/methodology",
			DownloadRef:     "/v1/public/proof-expansion/vala/downloadable-packs/point2_verification_public_pack?as_of=2026-04-22T10:00:00Z",
			VerificationRef: "/v1/public/proof-expansion/valb/signature-verification",
			ReplayRef:       "/v1/public/proof-expansion/valb/replay-verification",
			LineageRef:      "/v1/public/proof-expansion/valc/claim-lineage",
			EvidenceRefs:    []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
		},
	}
	if got := EvaluateMeasuredPublicProofValCPartnerPortalState(partnerItems); got != MeasuredPublicProofValCPartnerPortalStateActive {
		t.Fatalf("expected active partner portal state, got %q", got)
	}

	lineage := []PublicProofClaimLineageItem{{
		ClaimID:           "point2_runtime_performance_claim",
		ArtifactID:        "point2_runtime_performance_public_pack",
		CurrentState:      "lineage_ready",
		FreshnessState:    FreshnessFresh,
		PublicationScope:  ScopePublic,
		VisibilityState:   VisibilityPublicSafe,
		SupersessionState: "not_superseded",
		ArtifactRefs:      []string{"/v1/public/proof-expansion/vala/downloadable-packs/point2_runtime_performance_public_pack?as_of=2026-04-22T10:00:00Z"},
		TransparencyRefs:  []string{"/v1/public/proof-expansion/valb/transparency-chain"},
		VerifierRefs:      []string{"/v1/public/proof-expansion/valb/signature-verification", "/v1/public/proof-expansion/valb/replay-verification"},
		MethodologyRefs:   []string{"/v1/public/benchmarks/methodology"},
		EvidenceRefs:      []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
	}}
	if got := EvaluateMeasuredPublicProofValCClaimLineageState(lineage); got != MeasuredPublicProofValCClaimLineageStateActive {
		t.Fatalf("expected active claim lineage state, got %q", got)
	}

	downloads := []PublicProofDownloadProjectionItem{{
		ArtifactID:         "point2_runtime_performance_public_pack",
		ClaimID:            "point2_runtime_performance_claim",
		CurrentState:       "download_projection_ready",
		RedactionTier:      RedactionTierPublicSafe,
		PublicationScope:   ScopePublic,
		VisibilityState:    VisibilityPublicSafe,
		DownloadRef:        "/v1/public/proof-expansion/vala/downloadable-packs/point2_runtime_performance_public_pack?as_of=2026-04-22T10:00:00Z",
		TimestampRef:       "/v1/public/proof-expansion/vala/downloadable-packs/point2_runtime_performance_public_pack?as_of=2026-04-22T10:00:00Z#timestamp",
		PayloadDigest:      "sha256:abc",
		ReplayAvailability: "bounded_replay_available",
		AllowedScopes:      []string{ScopePublic, ScopePartner, ScopeAuditor, ScopeInternal},
		EvidenceRefs:       []string{"/v1/public/proof-expansion/vala/downloadable-packs"},
	}}
	if got := EvaluateMeasuredPublicProofValCDownloadProjectionState(downloads); got != MeasuredPublicProofValCDownloadProjectionStateActive {
		t.Fatalf("expected active download projection state, got %q", got)
	}

	if got := EvaluateMeasuredPublicProofValCState(
		MeasuredPublicProofValBStateActive,
		MeasuredPublicProofValCPublicPortalStateActive,
		MeasuredPublicProofValCPartnerPortalStateActive,
		MeasuredPublicProofValCClaimLineageStateActive,
		MeasuredPublicProofValCDownloadProjectionStateActive,
	); got != MeasuredPublicProofValCStateActive {
		t.Fatalf("expected active valc state, got %q", got)
	}
}
