package claims

import "testing"

func TestMeasuredPublicProofValBStateRequiresActiveValA(t *testing.T) {
	got := EvaluateMeasuredPublicProofValBState(
		MeasuredPublicProofValAStateSubstantial,
		MeasuredPublicProofValBTransparencyChainStateActive,
		MeasuredPublicProofValBVerifierCapabilityStateActive,
		MeasuredPublicProofValBSignatureVerificationStateActive,
		MeasuredPublicProofValBReplayVerificationStateActive,
	)
	if got != MeasuredPublicProofValBStateIncomplete {
		t.Fatalf("expected incomplete valb state without active vala, got %q", got)
	}
}

func TestMeasuredPublicProofValBSignatureVerificationIsPartialWithoutTrustRootState(t *testing.T) {
	items := []PublicProofSignatureVerificationItem{{
		ArtifactID:          "pack-a",
		CurrentState:        "verification_ready",
		VerificationState:   "verified",
		TrustRootState:      "",
		SchemaCompatibility: "supported",
		VerifierRef:         "/v1/public/verifier/sdk",
		TrustRootID:         "public_proof_primary_root",
		KeyVersion:          "v1",
		SignatureProvider:   "software",
		PayloadDigest:       "sha256:abc",
		EvidenceRefs:        []string{"/v1/public/proof-expansion/vala/downloadable-packs/pack-a"},
		FailureStates:       []string{"failed_verification"},
	}}
	if got := EvaluateMeasuredPublicProofValBSignatureVerificationState(items); got != MeasuredPublicProofValBSignatureVerificationStatePartial {
		t.Fatalf("expected partial signature verification without trust root state, got %q", got)
	}
}

func TestMeasuredPublicProofValBReplayVerificationIsPartialWithoutToleranceBands(t *testing.T) {
	items := []PublicProofReplayVerificationItem{{
		ArtifactID:                  "pack-a",
		CurrentState:                "replay_ready",
		ComparisonMode:              "bounded_tolerance_comparison",
		MethodologyRef:              "/v1/public/benchmarks/methodology",
		EvaluationState:             "replay_verified",
		EvaluationRef:               "/v1/foundation/execution/benchmarks/evaluate",
		SupportedEnvironmentClasses: []string{"runtime_hardened_enterprise_cluster"},
		ReplayCommandHints:          []string{"go test ./..."},
		UnsupportedReplayCases:      []string{"cross_environment_outside_tolerance"},
		EvidenceRefs:                []string{"/v1/public/proof-expansion/vala/downloadable-packs/pack-a"},
	}}
	if got := EvaluateMeasuredPublicProofValBReplayVerificationState(items); got != MeasuredPublicProofValBReplayVerificationStatePartial {
		t.Fatalf("expected partial replay verification without tolerance bands, got %q", got)
	}
}

func TestMeasuredPublicProofValBFoundationIsActive(t *testing.T) {
	transparency := PublicProofTransparencyChain{
		CurrentState:    "transparency_ready",
		ChainID:         "point2_valb_transparency_chain_v1",
		ParentAnchorRef: "/v1/public/transparency/anchor",
		Entries: []PublicProofTransparencyEntry{{
			ArtifactID:       "pack-a",
			CurrentState:     "anchored",
			ParentAnchorRef:  "/v1/public/transparency/anchor",
			AnchorID:         "point2_valb_transparency_chain_v1",
			EntryID:          "point2_valb_transparency_chain_v1/pack-a",
			EntryHash:        "sha256:abc",
			AnchoredAt:       "2026-04-23T10:00:00Z",
			TransparencyRefs: []string{"/v1/public/transparency/anchor", "/v1/public/proof-expansion/valb/transparency-chain"},
		}},
		IntegrityRules:   []string{"anchored artifacts remain digest-bound"},
		PublicationRules: []string{"anchored records remain public-proof projections"},
	}
	if got := EvaluateMeasuredPublicProofValBTransparencyChainState(transparency); got != MeasuredPublicProofValBTransparencyChainStateActive {
		t.Fatalf("expected active transparency chain state, got %q", got)
	}

	verifier := PublicProofVerifierCapability{
		CurrentState:         "verifier_ready",
		SDKRef:               "/v1/public/verifier/sdk",
		ReferencePackRef:     "/v1/public/verifier/reference-pack",
		OfflineGuideRef:      "/v1/public/verifier/offline-guide",
		SupportedSchemaLines: []string{"public.proof.sealed_artifact.v1"},
		ResultStates:         []string{"verified", "verified_with_limitations", "failed_verification"},
		TrustVerification:    []string{"trust root and signer metadata are checked"},
		ReplayVerification:   []string{"bounded replay uses explicit tolerance bands"},
		CommandHints:         []string{"go test ./..."},
	}
	if got := EvaluateMeasuredPublicProofValBVerifierCapabilityState(verifier); got != MeasuredPublicProofValBVerifierCapabilityStateActive {
		t.Fatalf("expected active verifier capability state, got %q", got)
	}

	signatureItems := []PublicProofSignatureVerificationItem{{
		ArtifactID:          "pack-a",
		CurrentState:        "verification_ready",
		VerificationState:   "verified",
		TrustRootState:      "trusted",
		SchemaCompatibility: "supported",
		VerifierRef:         "/v1/public/verifier/sdk",
		TrustRootID:         "public_proof_primary_root",
		KeyVersion:          "v1",
		SignatureProvider:   "software",
		PayloadDigest:       "sha256:abc",
		EvidenceRefs:        []string{"/v1/public/proof-expansion/vala/downloadable-packs/pack-a"},
		FailureStates:       []string{"failed_verification"},
	}}
	if got := EvaluateMeasuredPublicProofValBSignatureVerificationState(signatureItems); got != MeasuredPublicProofValBSignatureVerificationStateActive {
		t.Fatalf("expected active signature verification state, got %q", got)
	}

	replayItems := []PublicProofReplayVerificationItem{{
		ArtifactID:                  "pack-a",
		CurrentState:                "replay_ready",
		ComparisonMode:              "bounded_tolerance_comparison",
		MethodologyRef:              "/v1/public/benchmarks/methodology",
		EvaluationState:             "replay_verified",
		EvaluationRef:               "/v1/foundation/execution/benchmarks/evaluate",
		SupportedEnvironmentClasses: []string{"runtime_hardened_enterprise_cluster"},
		ToleranceBands:              []string{"capture_latency<=15%", "correlation_latency<=20%"},
		ReplayCommandHints:          []string{"go test ./..."},
		UnsupportedReplayCases:      []string{"cross_environment_outside_tolerance"},
		EvidenceRefs:                []string{"/v1/public/proof-expansion/vala/downloadable-packs/pack-a"},
	}}
	if got := EvaluateMeasuredPublicProofValBReplayVerificationState(replayItems); got != MeasuredPublicProofValBReplayVerificationStateActive {
		t.Fatalf("expected active replay verification state, got %q", got)
	}

	if got := EvaluateMeasuredPublicProofValBState(
		MeasuredPublicProofValAStateActive,
		MeasuredPublicProofValBTransparencyChainStateActive,
		MeasuredPublicProofValBVerifierCapabilityStateActive,
		MeasuredPublicProofValBSignatureVerificationStateActive,
		MeasuredPublicProofValBReplayVerificationStateActive,
	); got != MeasuredPublicProofValBStateActive {
		t.Fatalf("expected active valb state, got %q", got)
	}
}
