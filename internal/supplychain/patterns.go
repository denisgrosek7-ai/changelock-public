package supplychain

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
	"time"

	federationint "github.com/denisgrosek/changelock/internal/federation"
)

const (
	PatternVerdictSchemaVersion       = "3.supply_chain_pattern_verdict.v1"
	PatternVerdictModelVersion        = "phase3-supply-chain-baseline-v1"
	PatternStateStableTrusted         = "stable_trusted"
	PatternStateTrustDriftObserved    = "trust_drift_observed"
	PatternStateSuspiciousPublication = "suspicious_publication_pattern"
	PatternStateTyposquat             = "suspected_typo_squat"
	PatternStateCrossClusterConcern   = "cross_cluster_concern_active"
	PatternStateUnderReview           = "under_review"
)

type ExplanationPayload struct {
	Observed []string `json:"observed,omitempty"`
	Derived  []string `json:"derived,omitempty"`
}

type BaselineContext struct {
	WorkloadClass               string   `json:"workload_class,omitempty"`
	LanguageEcosystem           string   `json:"language_ecosystem,omitempty"`
	ExpectedBehaviors           []string `json:"expected_behaviors,omitempty"`
	ExpectedPublishCadenceHours int      `json:"expected_publish_cadence_hours,omitempty"`
}

type Input struct {
	SubjectRef           string                             `json:"subject_ref"`
	PackageName          string                             `json:"package_name"`
	PackageVersion       string                             `json:"package_version,omitempty"`
	Publisher            string                             `json:"publisher,omitempty"`
	PreviousPublisher    string                             `json:"previous_publisher,omitempty"`
	TyposquatCandidateOf string                             `json:"typosquat_candidate_of,omitempty"`
	PublishDeltaHours    int                                `json:"publish_delta_hours,omitempty"`
	RuntimeBehaviors     []string                           `json:"runtime_behaviors,omitempty"`
	ProvenanceConsistent bool                               `json:"provenance_consistent"`
	SigningConsistent    bool                               `json:"signing_consistent"`
	Baseline             BaselineContext                    `json:"baseline"`
	FederatedSignals     []federationint.IntelligenceSignal `json:"federated_signals,omitempty"`
	EvidenceRefs         []string                           `json:"evidence_refs,omitempty"`
}

type PatternVerdict struct {
	SchemaVersion       string                           `json:"schema_version"`
	ModelVersion        string                           `json:"model_version"`
	PatternID           string                           `json:"pattern_id"`
	SubjectRef          string                           `json:"subject_ref"`
	PackageName         string                           `json:"package_name"`
	PackageVersion      string                           `json:"package_version,omitempty"`
	CurrentState        string                           `json:"current_state"`
	TrustScore          int                              `json:"trust_score"`
	SourceConfidence    int                              `json:"source_confidence"`
	ReasonCodes         []string                         `json:"reason_codes,omitempty"`
	Explanation         ExplanationPayload               `json:"explanation"`
	FederatedAssessment federationint.WeightedAssessment `json:"federated_assessment"`
	EvidenceRefs        []string                         `json:"evidence_refs,omitempty"`
	AdvisoryOnly        bool                             `json:"advisory_only"`
	ObservedAt          time.Time                        `json:"observed_at"`
	Limitations         []string                         `json:"limitations,omitempty"`
}

func EvaluatePattern(input Input, now func() time.Time) PatternVerdict {
	if now == nil {
		now = time.Now
	}
	input.SubjectRef = strings.TrimSpace(input.SubjectRef)
	input.PackageName = strings.TrimSpace(input.PackageName)
	input.PackageVersion = strings.TrimSpace(input.PackageVersion)
	input.Publisher = strings.TrimSpace(input.Publisher)
	input.PreviousPublisher = strings.TrimSpace(input.PreviousPublisher)
	input.TyposquatCandidateOf = strings.TrimSpace(input.TyposquatCandidateOf)
	input.RuntimeBehaviors = uniqueStrings(input.RuntimeBehaviors)
	input.EvidenceRefs = uniqueStrings(input.EvidenceRefs)
	input.Baseline.WorkloadClass = strings.TrimSpace(input.Baseline.WorkloadClass)
	input.Baseline.LanguageEcosystem = strings.TrimSpace(input.Baseline.LanguageEcosystem)
	input.Baseline.ExpectedBehaviors = uniqueStrings(input.Baseline.ExpectedBehaviors)

	federated := federationint.WeightSignals(input.FederatedSignals, now)
	score := 78
	reasons := []string{}
	observed := []string{}
	derived := []string{}
	state := PatternStateStableTrusted

	if input.TyposquatCandidateOf != "" {
		score -= 45
		state = PatternStateTyposquat
		reasons = append(reasons, "typosquat_similarity_detected")
		observed = append(observed, "package name closely resembles a different trusted package")
	}
	if input.PreviousPublisher != "" && input.Publisher != "" && input.PreviousPublisher != input.Publisher {
		score -= 20
		reasons = append(reasons, "maintainer_drift_observed")
		observed = append(observed, "publisher identity changed from the previous observed state")
		if state == PatternStateStableTrusted {
			state = PatternStateTrustDriftObserved
		}
	}
	if input.PublishDeltaHours > 0 && input.Baseline.ExpectedPublishCadenceHours > 0 && input.PublishDeltaHours < input.Baseline.ExpectedPublishCadenceHours/4 {
		score -= 10
		reasons = append(reasons, "sudden_publish_cadence_anomaly")
		if state == PatternStateStableTrusted {
			state = PatternStateSuspiciousPublication
		}
	}
	if !input.ProvenanceConsistent || !input.SigningConsistent {
		score -= 18
		reasons = append(reasons, "publication_trust_inconsistent")
		if state == PatternStateStableTrusted {
			state = PatternStateSuspiciousPublication
		}
	}
	if hasUnexpectedBehavior(input.RuntimeBehaviors, input.Baseline.ExpectedBehaviors) {
		score -= 15
		reasons = append(reasons, "behavioral_baseline_deviation")
		derived = append(derived, "runtime behavior deviates from the bounded workload baseline")
		if state == PatternStateStableTrusted {
			state = PatternStateUnderReview
		}
	}
	if federated.CurrentState == federationint.WeightedAssessmentCrossClusterRisk {
		score -= 20
		reasons = append(reasons, "cross_cluster_concern_active")
		state = PatternStateCrossClusterConcern
	}

	score = clamp(score, 0, 100)
	sourceConfidence := clamp(55+len(reasons)*8, 0, 100)
	if len(reasons) == 0 {
		reasons = []string{"supply_chain_pattern_stable"}
		observed = append(observed, "no bounded trust drift or anomaly signal crossed the current threshold")
		derived = append(derived, "package remains stable in the current bounded baseline window")
	}
	return PatternVerdict{
		SchemaVersion:       PatternVerdictSchemaVersion,
		ModelVersion:        PatternVerdictModelVersion,
		PatternID:           stableID(input.SubjectRef, input.PackageName, input.PackageVersion),
		SubjectRef:          input.SubjectRef,
		PackageName:         input.PackageName,
		PackageVersion:      input.PackageVersion,
		CurrentState:        state,
		TrustScore:          score,
		SourceConfidence:    sourceConfidence,
		ReasonCodes:         uniqueStrings(reasons),
		Explanation:         ExplanationPayload{Observed: uniqueStrings(observed), Derived: uniqueStrings(derived)},
		FederatedAssessment: federated,
		EvidenceRefs:        input.EvidenceRefs,
		AdvisoryOnly:        true,
		ObservedAt:          now().UTC(),
		Limitations: []string{
			"Anomaly and trust-drift signals are bounded heuristics with reason codes; they are not standalone proof of maliciousness.",
			"Federated signals are weighted inputs and do not automatically produce deny decisions without local policy confirmation.",
		},
	}
}

func hasUnexpectedBehavior(observed, expected []string) bool {
	if len(observed) == 0 || len(expected) == 0 {
		return false
	}
	allowed := map[string]struct{}{}
	for _, item := range expected {
		allowed[strings.TrimSpace(strings.ToLower(item))] = struct{}{}
	}
	for _, item := range observed {
		if _, ok := allowed[strings.TrimSpace(strings.ToLower(item))]; !ok {
			return true
		}
	}
	return false
}

func stableID(values ...string) string {
	h := sha1.New()
	for _, value := range values {
		h.Write([]byte(strings.TrimSpace(value)))
		h.Write([]byte{0})
	}
	return "pat-" + hex.EncodeToString(h.Sum(nil))[:16]
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
