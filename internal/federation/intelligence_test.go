package federation

import "testing"

func TestWeightSignalsCrossClusterConcern(t *testing.T) {
	assessment := WeightSignals([]IntelligenceSignal{
		{SourcePeerID: "local", LocalSource: true, SuspicionLevel: SignalSuspicionHigh, SourceConfidence: 90, ReasonCodes: []string{"maintainer_drift"}},
		{SourcePeerID: "peer-a", LocalSource: false, SuspicionLevel: SignalSuspicionHigh, SourceConfidence: 90, ReasonCodes: []string{"cross_cluster_typo_squat"}},
	}, nil)
	if assessment.CurrentState != WeightedAssessmentCrossClusterRisk {
		t.Fatalf("expected cross-cluster concern, got %#v", assessment)
	}
	if !assessment.BoundedPropagation || !assessment.LocalOverrideVisible {
		t.Fatalf("expected bounded propagation and local override visibility, got %#v", assessment)
	}
}

func TestWeightSignalsNoSignals(t *testing.T) {
	assessment := WeightSignals(nil, nil)
	if assessment.CurrentState != WeightedAssessmentLocalOnly {
		t.Fatalf("expected local only without signals, got %#v", assessment)
	}
}
