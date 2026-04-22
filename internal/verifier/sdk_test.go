package verifier

import (
	"testing"
	"time"
)

func TestEvaluateReturnsExpiredForStaleArtifact(t *testing.T) {
	result := Evaluate(Input{
		ArtifactID:         "sample-proof",
		ArtifactType:       "proof_bundle",
		SchemaVersion:      "6.sample.v1",
		VerifiedAt:         time.Date(2026, time.April, 21, 10, 0, 0, 0, time.UTC),
		ValidUntil:         time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC),
		EvidenceRefs:       []string{"/v1/public/proof-portal"},
		IntegrityConfirmed: true,
		ChainContinuity:    true,
	}, time.Date(2026, time.April, 23, 10, 0, 0, 0, time.UTC))

	if result.CurrentState != StateExpired {
		t.Fatalf("expected expired result, got %#v", result)
	}
}

func TestEvaluateReturnsIncompleteWithoutEvidence(t *testing.T) {
	result := Evaluate(Input{
		ArtifactID:         "sample-proof",
		ArtifactType:       "proof_bundle",
		SchemaVersion:      "6.sample.v1",
		IntegrityConfirmed: true,
		ChainContinuity:    true,
	}, time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC))

	if result.CurrentState != StateIncomplete {
		t.Fatalf("expected incomplete result, got %#v", result)
	}
}

func TestEvaluateReturnsUnsupportedForUnknownSchemaLine(t *testing.T) {
	result := Evaluate(Input{
		ArtifactID:         "sample-proof",
		ArtifactType:       "proof_bundle",
		SchemaVersion:      "6.sample.v9",
		SupportedSchemas:   []string{"6.sample.v1", "6.sample.v2"},
		EvidenceRefs:       []string{"/v1/public/proof-portal"},
		IntegrityConfirmed: true,
		ChainContinuity:    true,
	}, time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC))

	if result.CurrentState != StateUnsupported {
		t.Fatalf("expected unsupported result, got %#v", result)
	}
	if result.CompatibilityState != CompatibilityUnsupported {
		t.Fatalf("expected unsupported compatibility state, got %#v", result)
	}
}

func TestEvaluateKeepsExpiredIncompleteAndUnsupportedDistinctAcrossArtifacts(t *testing.T) {
	cases := []struct {
		name   string
		input  Input
		expect string
	}{
		{
			name: "expired_benchmark",
			input: Input{
				ArtifactID:         "bench",
				ArtifactType:       "benchmark_pack",
				SchemaVersion:      "6.benchmark.v1",
				SupportedSchemas:   []string{"6.benchmark.v1"},
				VerifiedAt:         time.Date(2026, time.April, 21, 10, 0, 0, 0, time.UTC),
				ValidUntil:         time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC),
				EvidenceRefs:       []string{"/v1/public/benchmarks/packs"},
				MethodologyRef:     "/v1/public/benchmarks/methodology",
				IntegrityConfirmed: true,
				ChainContinuity:    true,
			},
			expect: StateExpired,
		},
		{
			name: "incomplete_benchmark",
			input: Input{
				ArtifactID:         "bench",
				ArtifactType:       "benchmark_pack",
				SchemaVersion:      "6.benchmark.v1",
				SupportedSchemas:   []string{"6.benchmark.v1"},
				EvidenceRefs:       []string{"/v1/public/benchmarks/packs"},
				IntegrityConfirmed: true,
				ChainContinuity:    true,
			},
			expect: StateIncomplete,
		},
		{
			name: "unsupported_handoff",
			input: Input{
				ArtifactID:         "handoff",
				ArtifactType:       "handoff_bundle",
				SchemaVersion:      "6.handoff.v9",
				SupportedSchemas:   []string{"6.handoff.v1"},
				EvidenceRefs:       []string{"/v1/public/specs/handoff"},
				IntegrityConfirmed: true,
				ChainContinuity:    true,
			},
			expect: StateUnsupported,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := Evaluate(tc.input, time.Date(2026, time.April, 23, 10, 0, 0, 0, time.UTC))
			if result.CurrentState != tc.expect {
				t.Fatalf("expected %s, got %#v", tc.expect, result)
			}
		})
	}
}
