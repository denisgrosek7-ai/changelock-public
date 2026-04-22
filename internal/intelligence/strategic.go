package intelligence

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
	"time"
)

const (
	StrategicAssessmentSchemaVersion = "3.intelligence_strategic_assessment.v1"
	QueryResponseSchemaVersion       = "3.intelligence_query_response.v1"
	ModelVersion                     = "phase3-strategic-support-v1"
)

type Recommendation struct {
	RecommendationID   string   `json:"recommendation_id"`
	Title              string   `json:"title"`
	Summary            string   `json:"summary"`
	Action             string   `json:"action"`
	PriorityBand       string   `json:"priority_band"`
	EffortBand         string   `json:"effort_band"`
	RiskReductionScore int      `json:"risk_reduction_score"`
	ConfidenceScore    int      `json:"confidence_score"`
	EvidenceRefs       []string `json:"evidence_refs,omitempty"`
	Explanation        []string `json:"explanation,omitempty"`
	AdvisoryOnly       bool     `json:"advisory_only"`
}

type StrategicAssessmentInput struct {
	SubjectRef          string   `json:"subject_ref"`
	CandidateAction     string   `json:"candidate_action,omitempty"`
	RelevanceScore      int      `json:"relevance_score"`
	PatternTrustScore   int      `json:"pattern_trust_score"`
	BlastRadiusScore    int      `json:"blast_radius_score"`
	DelayDays           int      `json:"delay_days,omitempty"`
	EffortBand          string   `json:"effort_band,omitempty"`
	ObservedFacts       []string `json:"observed_facts,omitempty"`
	InferredConclusions []string `json:"inferred_conclusions,omitempty"`
	EvidenceRefs        []string `json:"evidence_refs,omitempty"`
}

type StrategicAssessment struct {
	SchemaVersion       string         `json:"schema_version"`
	ModelVersion        string         `json:"model_version"`
	AssessmentID        string         `json:"assessment_id"`
	SubjectRef          string         `json:"subject_ref"`
	ObservedFacts       []string       `json:"observed_facts,omitempty"`
	InferredConclusions []string       `json:"inferred_conclusions,omitempty"`
	RecommendedActions  []string       `json:"recommended_actions,omitempty"`
	RiskDelta           int            `json:"risk_delta"`
	BlastRadiusDelta    int            `json:"blast_radius_delta"`
	ResponseCost        string         `json:"response_cost"`
	ConfidenceScore     int            `json:"confidence_score"`
	Uncertainty         []string       `json:"uncertainty,omitempty"`
	Recommendation      Recommendation `json:"recommendation"`
	EvidenceRefs        []string       `json:"evidence_refs,omitempty"`
	CurrentState        string         `json:"current_state"`
	AdvisoryOnly        bool           `json:"advisory_only"`
	EvaluatedAt         time.Time      `json:"evaluated_at"`
	Limitations         []string       `json:"limitations,omitempty"`
}

type QueryScope struct {
	SubjectRef      string `json:"subject_ref,omitempty"`
	VulnerabilityID string `json:"vulnerability_id,omitempty"`
	PackageName     string `json:"package_name,omitempty"`
	TenantID        string `json:"tenant_id,omitempty"`
	Environment     string `json:"environment,omitempty"`
	Repo            string `json:"repo,omitempty"`
}

type QueryResponse struct {
	SchemaVersion       string     `json:"schema_version"`
	ModelVersion        string     `json:"model_version"`
	QueryID             string     `json:"query_id"`
	Query               string     `json:"query"`
	Scope               QueryScope `json:"scope"`
	ObservedFacts       []string   `json:"observed_facts,omitempty"`
	InferredConclusions []string   `json:"inferred_conclusions,omitempty"`
	RecommendedActions  []string   `json:"recommended_actions,omitempty"`
	Uncertainty         []string   `json:"uncertainty,omitempty"`
	EvidenceRefs        []string   `json:"evidence_refs,omitempty"`
	CurrentState        string     `json:"current_state"`
	RetrievalMode       string     `json:"retrieval_mode"`
	AdvisoryOnly        bool       `json:"advisory_only"`
	AnsweredAt          time.Time  `json:"answered_at"`
	Limitations         []string   `json:"limitations,omitempty"`
}

func Assess(input StrategicAssessmentInput, now func() time.Time) StrategicAssessment {
	if now == nil {
		now = time.Now
	}
	input.SubjectRef = strings.TrimSpace(input.SubjectRef)
	input.CandidateAction = firstNonEmpty(strings.TrimSpace(input.CandidateAction), "prioritize_patch_and_validate")
	input.EffortBand = firstNonEmpty(strings.TrimSpace(input.EffortBand), "moderate")
	input.ObservedFacts = uniqueStrings(input.ObservedFacts)
	input.InferredConclusions = uniqueStrings(input.InferredConclusions)
	input.EvidenceRefs = uniqueStrings(input.EvidenceRefs)

	riskDelta := clamp(input.RelevanceScore+(100-input.PatternTrustScore)/2+input.BlastRadiusScore/2+input.DelayDays, 0, 100)
	blastDelta := clamp(input.BlastRadiusScore-input.DelayDays/3, 0, 100)
	confidence := clamp(55+len(input.ObservedFacts)*5+len(input.EvidenceRefs)*3, 0, 100)
	priority := "medium"
	switch {
	case riskDelta >= 80:
		priority = "critical"
	case riskDelta >= 65:
		priority = "high"
	case riskDelta <= 35:
		priority = "low"
	}
	recommended := []string{
		"apply the bounded candidate action only after reviewing the linked evidence and uncertainty notes",
		input.CandidateAction,
	}
	if priority == "critical" || priority == "high" {
		recommended = append(recommended, "validate remediation timing against current blast-radius reduction potential")
	}
	responseCost := map[string]string{
		"low":      "bounded_low_cost",
		"moderate": "bounded_moderate_cost",
		"high":     "bounded_high_cost",
	}[strings.ToLower(input.EffortBand)]
	if responseCost == "" {
		responseCost = "bounded_moderate_cost"
	}
	recommendation := Recommendation{
		RecommendationID:   stableID(input.SubjectRef, input.CandidateAction, priority),
		Title:              "Phase 3 advisory recommendation",
		Summary:            "Bounded intelligence-derived recommendation built from current execution, vulnerability, and supply-chain evidence.",
		Action:             input.CandidateAction,
		PriorityBand:       priority,
		EffortBand:         strings.ToLower(input.EffortBand),
		RiskReductionScore: clamp(riskDelta-blastDelta/4, 0, 100),
		ConfidenceScore:    confidence,
		EvidenceRefs:       input.EvidenceRefs,
		Explanation: []string{
			"Recommendation ranking is advisory-only and grounded in the currently linked evidence set.",
			"Observed facts remain separate from inferred strategic conclusions.",
		},
		AdvisoryOnly: true,
	}
	return StrategicAssessment{
		SchemaVersion:       StrategicAssessmentSchemaVersion,
		ModelVersion:        ModelVersion,
		AssessmentID:        stableID(input.SubjectRef, input.CandidateAction, "assessment"),
		SubjectRef:          input.SubjectRef,
		ObservedFacts:       input.ObservedFacts,
		InferredConclusions: append([]string{}, input.InferredConclusions...),
		RecommendedActions:  uniqueStrings(recommended),
		RiskDelta:           riskDelta,
		BlastRadiusDelta:    blastDelta,
		ResponseCost:        responseCost,
		ConfidenceScore:     confidence,
		Uncertainty: []string{
			"Cost and blast-radius values are bounded operational estimates, not financial forecasts.",
			"Simulation does not mutate policy, runtime, or VEX state.",
		},
		Recommendation: recommendation,
		EvidenceRefs:   input.EvidenceRefs,
		CurrentState:   "strategic_advisory_ready",
		AdvisoryOnly:   true,
		EvaluatedAt:    now().UTC(),
		Limitations: []string{
			"Strategic support is recommendation-only and depends on the freshness of the linked evidence set.",
		},
	}
}

func BuildGroundedQuery(query string, scope QueryScope, observed, inferred, recommended, uncertainty, evidenceRefs []string, now func() time.Time) QueryResponse {
	if now == nil {
		now = time.Now
	}
	scope.SubjectRef = strings.TrimSpace(scope.SubjectRef)
	scope.VulnerabilityID = strings.ToUpper(strings.TrimSpace(scope.VulnerabilityID))
	scope.PackageName = strings.TrimSpace(scope.PackageName)
	scope.TenantID = strings.TrimSpace(scope.TenantID)
	scope.Environment = strings.TrimSpace(scope.Environment)
	scope.Repo = strings.TrimSpace(scope.Repo)
	return QueryResponse{
		SchemaVersion:       QueryResponseSchemaVersion,
		ModelVersion:        ModelVersion,
		QueryID:             stableID(query, scope.SubjectRef, scope.VulnerabilityID, scope.PackageName, scope.TenantID, scope.Environment, scope.Repo, "query"),
		Query:               strings.TrimSpace(query),
		Scope:               scope,
		ObservedFacts:       uniqueStrings(observed),
		InferredConclusions: uniqueStrings(inferred),
		RecommendedActions:  uniqueStrings(recommended),
		Uncertainty:         uniqueStrings(uncertainty),
		EvidenceRefs:        uniqueStrings(evidenceRefs),
		CurrentState:        "grounded_advisory_response",
		RetrievalMode:       "bounded_evidence_retrieval",
		AdvisoryOnly:        true,
		AnsweredAt:          now().UTC(),
		Limitations: []string{
			"Natural-language output is retrieval-grounded and advisory-only; it does not approve or mutate trust state.",
		},
	}
}

func stableID(values ...string) string {
	h := sha1.New()
	for _, value := range values {
		h.Write([]byte(strings.TrimSpace(value)))
		h.Write([]byte{0})
	}
	return "intel-" + hex.EncodeToString(h.Sum(nil))[:16]
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
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
