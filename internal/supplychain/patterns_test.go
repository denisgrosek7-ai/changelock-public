package supplychain

import (
	"testing"

	federationint "github.com/denisgrosek/changelock/internal/federation"
)

func TestEvaluatePatternTyposquatAndFederation(t *testing.T) {
	verdict := EvaluatePattern(Input{
		SubjectRef:           "cluster-a|acme-prod|Deployment|api",
		PackageName:          "reqeusts",
		PackageVersion:       "1.2.3",
		TyposquatCandidateOf: "requests",
		Publisher:            "attacker",
		PreviousPublisher:    "trusted",
		ProvenanceConsistent: false,
		SigningConsistent:    false,
		FederatedSignals: []federationint.IntelligenceSignal{
			{SourcePeerID: "peer-a", SuspicionLevel: federationint.SignalSuspicionHigh, SourceConfidence: 95},
			{SourcePeerID: "peer-b", SuspicionLevel: federationint.SignalSuspicionHigh, SourceConfidence: 90},
		},
	}, nil)
	if verdict.CurrentState != PatternStateCrossClusterConcern {
		t.Fatalf("expected cross-cluster concern, got %#v", verdict)
	}
	if verdict.TrustScore >= 60 {
		t.Fatalf("expected degraded trust score, got %#v", verdict)
	}
}

func TestEvaluatePatternStableTrusted(t *testing.T) {
	verdict := EvaluatePattern(Input{
		SubjectRef:           "cluster-a|acme-prod|Deployment|api",
		PackageName:          "requests",
		PackageVersion:       "2.32.0",
		Publisher:            "trusted",
		PreviousPublisher:    "trusted",
		ProvenanceConsistent: true,
		SigningConsistent:    true,
		Baseline: BaselineContext{
			ExpectedBehaviors: []string{"http_client"},
		},
		RuntimeBehaviors: []string{"http_client"},
	}, nil)
	if verdict.CurrentState != PatternStateStableTrusted {
		t.Fatalf("expected stable trusted, got %#v", verdict)
	}
}
