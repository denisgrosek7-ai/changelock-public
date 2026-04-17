package audit

import (
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/signingidentity"
	internalvex "github.com/denisgrosek/changelock/internal/vex"
)

func TestComputeTrustScorecardVerifiedSignalsProducesHighGrade(t *testing.T) {
	now := time.Date(2026, 4, 17, 12, 0, 0, 0, time.UTC)
	report := ExceptionReport{}
	input := TrustScorecardInput{
		ScopeType:    "tenant",
		ScopeRef:     "acme",
		TenantID:     "acme",
		CalculatedAt: now,
		ArtifactVerificationEvents: []StoredEvent{
			artifactVerificationEvent("acme", "repo/app", "verified", true, true, true),
			artifactVerificationEvent("acme", "repo/app", "verified", true, true, true),
		},
		PolicyDecisionEvents: []StoredEvent{
			{Event: Event{EventType: EventTypePolicyDecision, PolicyBundleID: "bundle-a", PolicyBundleHash: "sha256:abc"}},
			{Event: Event{EventType: EventTypeDeployGateDecision, PolicyVersion: "2026.04.17"}},
		},
		VulnerabilityNet: VulnerabilityNetResponse{
			RawCount:           2,
			ResolvedByVEXCount: 2,
			ActionableCount:    0,
		},
		VEXStatus: internalvex.StatusSummary{ActiveCount: 2},
		SigningIdentityStatus: signingidentity.StatusSummary{
			EnforcementMode:    signingidentity.EnforcementEnforce,
			ObservedIdentities: 2,
			Authorized:         2,
			EnabledPolicies:    2,
			TotalPolicies:      2,
		},
		RuntimeStatus: RuntimeClosedLoopStatus{
			TotalTargets: 4,
			InSync:       4,
		},
		RuntimeActiveStates: []RuntimeActiveStateView{
			{ID: "one", ReconciliationStatus: "in_sync"},
			{ID: "two", ReconciliationStatus: "in_sync"},
			{ID: "three", ReconciliationStatus: "in_sync"},
			{ID: "four", ReconciliationStatus: "in_sync"},
		},
		ExceptionReport: report,
	}

	scorecard := ComputeTrustScorecard(input)
	if scorecard.OverallGrade != TrustGradeA {
		t.Fatalf("expected grade A, got %s (%d)", scorecard.OverallGrade, scorecard.OverallScore)
	}
	if metric := findMetric(scorecard.Metrics, ScorecardMetricArtifactIntegrity); metric.Status != TrustMetricStatusVerified {
		t.Fatalf("expected verified artifact metric, got %#v", metric)
	}
}

func TestComputeTrustScorecardMissingDataStaysUnknown(t *testing.T) {
	scorecard := ComputeTrustScorecard(TrustScorecardInput{
		ScopeType:    "tenant",
		ScopeRef:     "acme",
		TenantID:     "acme",
		CalculatedAt: time.Date(2026, 4, 17, 12, 0, 0, 0, time.UTC),
	})

	if scorecard.OverallGrade == TrustGradeA || scorecard.OverallGrade == TrustGradeB {
		t.Fatalf("missing data should not inflate scorecard grade, got %s", scorecard.OverallGrade)
	}
	if metric := findMetric(scorecard.Metrics, ScorecardMetricArtifactIntegrity); metric.Status != TrustMetricStatusUnknown {
		t.Fatalf("expected unknown artifact metric, got %#v", metric)
	}
}

func TestBuildPublishedTrustViewSanitizesInternalOnlyDetails(t *testing.T) {
	scorecard := ComputeTrustScorecard(TrustScorecardInput{
		ScopeType:       "tenant",
		ScopeRef:        "acme",
		TenantID:        "acme",
		CalculatedAt:    time.Date(2026, 4, 17, 12, 0, 0, 0, time.UTC),
		PublicationMode: TrustPublicationPreview,
		ArtifactVerificationEvents: []StoredEvent{
			artifactVerificationEvent("acme", "repo/app", "verified", true, true, true),
		},
	})
	badges := BuildTrustBadges(scorecard, TrustScorecardInput{VEXStatus: internalvex.StatusSummary{}})
	view := BuildPublishedTrustView(scorecard, badges, BuildStandardsMapping(scorecard))
	if view == nil {
		t.Fatal("expected published trust view")
	}
	for _, metric := range view.Metrics {
		if len(metric.EvidenceRefs) != 0 {
			t.Fatalf("expected sanitized metrics without evidence refs, got %#v", metric)
		}
	}
	for _, badge := range view.Badges {
		if !badge.PublicPublishable {
			t.Fatalf("expected only public badges, got %#v", badge)
		}
	}
}

func TestBuildHardeningReviewFindings(t *testing.T) {
	now := time.Date(2026, 4, 17, 12, 0, 0, 0, time.UTC)
	report := ExceptionReport{
		Active: []PolicyException{
			{
				ExceptionID: "EX-1",
				Status:      ExceptionStatusApproved,
				CreatedAt:   now.Add(-30 * 24 * time.Hour),
				Active:      true,
			},
		},
	}
	input := TrustScorecardInput{
		ScopeType:       "tenant",
		ScopeRef:        "acme",
		TenantID:        "acme",
		CalculatedAt:    now,
		ExceptionReport: report,
		VulnerabilityNet: VulnerabilityNetResponse{
			RawCount:          2,
			ActionableCount:   2,
			ThresholdBreached: true,
		},
		SigningIdentityStatus: signingidentity.StatusSummary{
			ObservedIdentities: 1,
		},
		RuntimeStatus: RuntimeClosedLoopStatus{
			TotalTargets: 1,
			Failed:       1,
		},
		PolicyDecisionEvents: []StoredEvent{{Event: Event{EventType: EventTypePolicyDecision}}},
		StaleExceptionDays:   14,
	}

	card := ComputeTrustScorecard(input)
	findings := BuildHardeningReview(input, card)
	if len(findings) < 3 {
		t.Fatalf("expected multiple hardening findings, got %#v", findings)
	}
	foundStale := false
	foundVuln := false
	for _, finding := range findings {
		if finding.ReasonCode == HardeningFindingStaleExceptions {
			foundStale = true
		}
		if finding.ReasonCode == HardeningFindingVulnerabilityDebt {
			foundVuln = true
		}
	}
	if !foundStale || !foundVuln {
		t.Fatalf("expected stale exception and vulnerability debt findings, got %#v", findings)
	}
}

func artifactVerificationEvent(tenantID, repo, transparencyState string, signatureValid bool, attestationValid bool, hasSBOM bool) StoredEvent {
	artifact := &ArtifactEvidence{
		SignerIdentity: "https://github.com/example/repo/.github/workflows/release.yml@refs/heads/main",
		Repository:     repo,
		Digest:         "sha256:abc",
	}
	if hasSBOM {
		artifact.SBOMArtifactRef = "oci://ghcr.io/example/repo:sbom"
	}
	return StoredEvent{
		Event: Event{
			EventType: auditEventTypeArtifactVerificationResult(),
			TenantID:  tenantID,
			Repo:      repo,
			Evidence: &Evidence{
				Artifact:          artifact,
				VerificationState: transparencyState,
			},
			VerifierSummary: &VerifierSummary{
				SignatureValid:   signatureValid,
				AttestationValid: attestationValid,
			},
		},
	}
}

func auditEventTypeArtifactVerificationResult() string {
	return EventTypeArtifactVerificationResult
}
