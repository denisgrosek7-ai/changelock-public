package federation

import (
	"strings"
	"time"
)

const (
	IntelligenceSignalSchemaVersion    = "3.federation_intelligence_signal.v1"
	WeightedAssessmentSchemaVersion    = "3.federation_weighted_assessment.v1"
	SignalSuspicionLow                 = "low"
	SignalSuspicionMedium              = "medium"
	SignalSuspicionHigh                = "high"
	WeightedAssessmentLocalOnly        = "local_only"
	WeightedAssessmentBoundedObserved  = "bounded_federated_signal"
	WeightedAssessmentCrossClusterRisk = "cross_cluster_concern_active"
	WeightedAssessmentUnderReview      = "under_review"
)

type IntelligenceSignal struct {
	SchemaVersion    string    `json:"schema_version,omitempty"`
	SourcePeerID     string    `json:"source_peer_id"`
	SignalType       string    `json:"signal_type"`
	SuspicionLevel   string    `json:"suspicion_level"`
	SourceConfidence int       `json:"source_confidence"`
	LocalSource      bool      `json:"local_source"`
	ObservedAt       time.Time `json:"observed_at"`
	ReasonCodes      []string  `json:"reason_codes,omitempty"`
}

type WeightedAssessment struct {
	SchemaVersion        string         `json:"schema_version"`
	CurrentState         string         `json:"current_state"`
	WeightedScore        int            `json:"weighted_score"`
	BoundedPropagation   bool           `json:"bounded_propagation"`
	LocalOverrideVisible bool           `json:"local_override_visible"`
	ReasonCodes          []string       `json:"reason_codes,omitempty"`
	SourceWeights        map[string]int `json:"source_weights,omitempty"`
	ObservedAt           time.Time      `json:"observed_at"`
}

func NormalizeSignal(signal IntelligenceSignal, now func() time.Time) IntelligenceSignal {
	if now == nil {
		now = time.Now
	}
	if strings.TrimSpace(signal.SchemaVersion) == "" {
		signal.SchemaVersion = IntelligenceSignalSchemaVersion
	}
	signal.SourcePeerID = strings.TrimSpace(signal.SourcePeerID)
	signal.SignalType = strings.TrimSpace(signal.SignalType)
	signal.SuspicionLevel = normalizeSuspicionLevel(signal.SuspicionLevel)
	if signal.SourceConfidence <= 0 {
		signal.SourceConfidence = 50
	}
	if signal.SourceConfidence > 100 {
		signal.SourceConfidence = 100
	}
	if signal.ObservedAt.IsZero() {
		signal.ObservedAt = now().UTC()
	}
	signal.ReasonCodes = uniqueStrings(signal.ReasonCodes)
	return signal
}

func WeightSignals(signals []IntelligenceSignal, now func() time.Time) WeightedAssessment {
	if now == nil {
		now = time.Now
	}
	assessment := WeightedAssessment{
		SchemaVersion:        WeightedAssessmentSchemaVersion,
		CurrentState:         WeightedAssessmentLocalOnly,
		BoundedPropagation:   true,
		LocalOverrideVisible: true,
		SourceWeights:        map[string]int{},
		ObservedAt:           now().UTC(),
	}
	if len(signals) == 0 {
		assessment.ReasonCodes = []string{"no_federated_signals"}
		return assessment
	}
	localHigh := false
	remoteHigh := 0
	score := 0
	reasons := []string{}
	for _, raw := range signals {
		signal := NormalizeSignal(raw, now)
		base := suspicionWeight(signal.SuspicionLevel)
		weight := (base * signal.SourceConfidence) / 100
		if !signal.LocalSource {
			weight = weight / 2
		}
		if weight <= 0 {
			weight = 1
		}
		sourceID := signal.SourcePeerID
		if sourceID == "" {
			sourceID = "unknown"
		}
		assessment.SourceWeights[sourceID] = weight
		score += weight
		reasons = append(reasons, signal.ReasonCodes...)
		if signal.LocalSource && signal.SuspicionLevel == SignalSuspicionHigh {
			localHigh = true
		}
		if !signal.LocalSource && signal.SuspicionLevel == SignalSuspicionHigh {
			remoteHigh++
		}
	}
	assessment.WeightedScore = clamp(score, 0, 100)
	switch {
	case localHigh || remoteHigh >= 2 || assessment.WeightedScore >= 55:
		assessment.CurrentState = WeightedAssessmentCrossClusterRisk
	case assessment.WeightedScore >= 30:
		assessment.CurrentState = WeightedAssessmentBoundedObserved
	default:
		assessment.CurrentState = WeightedAssessmentUnderReview
	}
	assessment.ReasonCodes = uniqueStrings(reasons)
	if len(assessment.ReasonCodes) == 0 {
		assessment.ReasonCodes = []string{"weighted_federated_signal_observed"}
	}
	return assessment
}

func suspicionWeight(level string) int {
	switch normalizeSuspicionLevel(level) {
	case SignalSuspicionHigh:
		return 60
	case SignalSuspicionMedium:
		return 35
	default:
		return 15
	}
}

func normalizeSuspicionLevel(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case SignalSuspicionHigh:
		return SignalSuspicionHigh
	case SignalSuspicionMedium:
		return SignalSuspicionMedium
	default:
		return SignalSuspicionLow
	}
}

func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := map[string]struct{}{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func clamp(value, lower, upper int) int {
	if value < lower {
		return lower
	}
	if value > upper {
		return upper
	}
	return value
}
