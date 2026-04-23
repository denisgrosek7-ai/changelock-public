package claims

import (
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/signing"
)

func TestMeasuredPublicProofValAStateRequiresActiveVal0(t *testing.T) {
	got := EvaluateMeasuredPublicProofValAState(
		MeasuredPublicProofVal0StateSubstantial,
		MeasuredPublicProofValAArtifactSchemaStateActive,
		MeasuredPublicProofValASealingDisciplineStateActive,
		MeasuredPublicProofValAEnvironmentBindingStateActive,
		MeasuredPublicProofValADownloadablePackStateActive,
	)
	if got != MeasuredPublicProofValAStateIncomplete {
		t.Fatalf("expected incomplete vala state without active val0, got %q", got)
	}
}

func TestMeasuredPublicProofValASealingDisciplineRequiresPurposeEnablement(t *testing.T) {
	model := MeasuredPublicProofValASealingDiscipline(signing.ProviderDescriptor{
		ProviderMode:         signing.ModeSoftware,
		TrustBoundary:        signing.TrustBoundaryApplicationLocal,
		ActiveLifecycleState: signing.KeyStateActive,
		KeyClasses:           []string{signing.KeyClassSealing, signing.KeyClassVerificationRoot},
	}, false)
	if got := EvaluateMeasuredPublicProofValASealingDisciplineState(model); got != MeasuredPublicProofValASealingDisciplineStatePartial {
		t.Fatalf("expected partial sealing discipline without purpose enablement, got %q", got)
	}
}

func TestMeasuredPublicProofValADownloadablePackRequiresSignatureAndTimestampMetadata(t *testing.T) {
	pack := measuredPublicProofValAReadyPack()
	pack.SignatureEnvelope = nil
	if got := EvaluateMeasuredPublicProofValADownloadablePackState([]PublicSealedProofPack{pack}); got != MeasuredPublicProofValADownloadablePackStatePartial {
		t.Fatalf("expected partial pack state without signature envelope, got %q", got)
	}

	pack = measuredPublicProofValAReadyPack()
	pack.TimestampRef = ""
	if got := EvaluateMeasuredPublicProofValADownloadablePackState([]PublicSealedProofPack{pack}); got != MeasuredPublicProofValADownloadablePackStatePartial {
		t.Fatalf("expected partial pack state without timestamp ref, got %q", got)
	}
}

func TestMeasuredPublicProofValAFoundationIsActive(t *testing.T) {
	schema := MeasuredPublicProofValAArtifactSchema()
	if got := EvaluateMeasuredPublicProofValAArtifactSchemaState(schema); got != MeasuredPublicProofValAArtifactSchemaStateActive {
		t.Fatalf("expected active artifact schema state, got %q", got)
	}

	sealing := MeasuredPublicProofValASealingDiscipline(signing.ProviderDescriptor{
		ProviderMode:                   signing.ModeSoftware,
		TrustBoundary:                  signing.TrustBoundaryApplicationLocal,
		ActiveLifecycleState:           signing.KeyStateActive,
		SupportsHistoricalVerification: true,
		KeyClasses:                     []string{signing.KeyClassSealing, signing.KeyClassVerificationRoot},
	}, true)
	if got := EvaluateMeasuredPublicProofValASealingDisciplineState(sealing); got != MeasuredPublicProofValASealingDisciplineStateActive {
		t.Fatalf("expected active sealing discipline state, got %q", got)
	}

	binding := []PublicProofEnvironmentBindingItem{
		{
			ArtifactID:         "runtime_performance_public_pack",
			CurrentState:       "binding_ready",
			ClaimClass:         PublicProofClaimClassPerformance,
			RedactionTier:      RedactionTierPublicSafe,
			EnvironmentClass:   "runtime_hardened_enterprise_cluster",
			ExecutionProfile:   "production_like",
			WorkloadShape:      "runtime_substrate_depth_point1",
			BuildIdentity:      "changelock-2026.04.23",
			HarnessVersion:     "1.execution_benchmark_harness.v1",
			MethodologyRef:     "/v1/public/benchmarks/methodology",
			CompatibilityScope: "public.proof.sealed_artifact.v1",
			ProvenanceInputs:   []string{"/v1/public/phase6/proofs", "/v1/runtime/substrate-depth/complete"},
			ReplayBoundaries:   []string{"production_like only", "methodology-bound replay"},
			UnsupportedReplay:  []string{"cross-environment replay outside tolerance"},
		},
	}
	if got := EvaluateMeasuredPublicProofValAEnvironmentBindingState(binding); got != MeasuredPublicProofValAEnvironmentBindingStateActive {
		t.Fatalf("expected active environment binding state, got %q", got)
	}

	if got := EvaluateMeasuredPublicProofValADownloadablePackState([]PublicSealedProofPack{measuredPublicProofValAReadyPack()}); got != MeasuredPublicProofValADownloadablePackStateActive {
		t.Fatalf("expected active downloadable pack state, got %q", got)
	}

	if got := EvaluateMeasuredPublicProofValAState(
		MeasuredPublicProofVal0StateActive,
		schema.CurrentState,
		sealing.CurrentState,
		MeasuredPublicProofValAEnvironmentBindingStateActive,
		MeasuredPublicProofValADownloadablePackStateActive,
	); got != MeasuredPublicProofValAStateActive {
		t.Fatalf("expected active vala state, got %q", got)
	}
}

func measuredPublicProofValAReadyPack() PublicSealedProofPack {
	now := time.Date(2026, 4, 23, 10, 0, 0, 0, time.UTC)
	return PublicSealedProofPack{
		ArtifactID:            "runtime_performance_public_pack",
		ArtifactSchemaVersion: "public.proof.sealed_artifact.v1",
		ArtifactType:          PublicProofArtifactTypeBenchmarkPack,
		CurrentState:          "sealed_artifact_ready",
		ClaimID:               "point2_runtime_performance_claim",
		ClaimClass:            PublicProofClaimClassPerformance,
		RedactionTier:         RedactionTierPublicSafe,
		EnvironmentClass:      "runtime_hardened_enterprise_cluster",
		ExecutionProfile:      "production_like",
		WorkloadShape:         "runtime_substrate_depth_point1",
		BuildIdentity:         "changelock-2026.04.23",
		HarnessVersion:        "1.execution_benchmark_harness.v1",
		MethodologyRef:        "/v1/public/benchmarks/methodology",
		IssuedAt:              now,
		ValidThrough:          now.Add(30 * 24 * time.Hour),
		MeasurementSource:     "runtime_substrate_vale_latency.standard_node.v1",
		EvidenceRefs:          []string{"/v1/public/phase6/proofs", "/v1/runtime/substrate-depth/complete"},
		DownloadRef:           "/v1/public/proof-expansion/vala/downloadable-packs/runtime_performance_public_pack",
		PayloadDigest:         "sha256:ready-pack",
		SignatureEnvelope: &signing.Envelope{
			Provider:      signing.ModeSoftware,
			KeyID:         "test-signing-key",
			Algorithm:     signing.AlgorithmHMACSHA256,
			Purpose:       signing.PurposePublicProofArtifact,
			PayloadDigest: "sha256:ready-pack",
			Signature:     "signature",
			SignedAt:      now,
		},
		TrustRootID:  "public_proof_primary_root",
		KeyVersion:   "v1",
		TimestampRef: "/v1/public/proof-expansion/vala/downloadable-packs/runtime_performance_public_pack#timestamp",
		PackagingFiles: []PublicSealedProofPackFile{
			{Path: "manifest.json", MediaType: "application/json", Role: "manifest", SHA256: "sha256:manifest"},
			{Path: "payload/measurement_summary.json", MediaType: "application/json", Role: "measurement_summary", SHA256: "sha256:payload"},
			{Path: "environment/binding.json", MediaType: "application/json", Role: "environment_binding", SHA256: "sha256:binding"},
			{Path: "signature/envelope.json", MediaType: "application/json", Role: "signature_envelope", SHA256: "sha256:signature"},
			{Path: "timestamp/receipt.json", MediaType: "application/json", Role: "timestamp_receipt", SHA256: "sha256:timestamp"},
		},
		MetricSummaries: []string{
			"capture_p99_micros=340",
			"correlation_p99_micros=580",
			"false_positive_rate_pct=0.94",
		},
	}
}
