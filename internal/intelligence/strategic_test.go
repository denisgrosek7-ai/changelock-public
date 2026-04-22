package intelligence

import "testing"

func TestAssessGeneratesAdvisoryRecommendation(t *testing.T) {
	assessment := Assess(StrategicAssessmentInput{
		SubjectRef:          "cluster-a|acme-prod|Deployment|api",
		CandidateAction:     "patch_and_validate",
		RelevanceScore:      88,
		PatternTrustScore:   30,
		BlastRadiusScore:    70,
		DelayDays:           7,
		EffortBand:          "high",
		ObservedFacts:       []string{"CVE is reachable", "package trust drift observed"},
		InferredConclusions: []string{"risk remains concentrated on an internet-exposed API"},
		EvidenceRefs:        []string{"report:/v1/intelligence/vulnerability-relevance"},
	}, nil)
	if assessment.CurrentState != "strategic_advisory_ready" {
		t.Fatalf("expected strategic advisory ready, got %#v", assessment)
	}
	if !assessment.Recommendation.AdvisoryOnly || assessment.Recommendation.PriorityBand == "" {
		t.Fatalf("expected advisory recommendation with priority, got %#v", assessment.Recommendation)
	}
}

func TestBuildGroundedQuery(t *testing.T) {
	response := BuildGroundedQuery(
		"what is highest priority?",
		QueryScope{
			SubjectRef:      "cluster-a|acme-prod|Deployment|api",
			VulnerabilityID: "CVE-2026-0001",
			PackageName:     "openssl",
			TenantID:        "acme",
			Environment:     "prod",
			Repo:            "github.com/acme/api",
		},
		[]string{"CVE-2026-0001 is observed reachable"},
		[]string{"this is the current highest-risk path"},
		[]string{"patch_and_validate"},
		[]string{"bounded estimate only"},
		[]string{"report:/v1/intelligence/phase3/proofs"},
		nil,
	)
	if !response.AdvisoryOnly || response.CurrentState != "grounded_advisory_response" {
		t.Fatalf("expected grounded advisory response, got %#v", response)
	}
	if response.Scope.SubjectRef != "cluster-a|acme-prod|Deployment|api" || response.Scope.Repo != "github.com/acme/api" {
		t.Fatalf("expected grounded query scope metadata, got %#v", response.Scope)
	}
}
