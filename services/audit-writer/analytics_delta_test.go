package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
)

func TestAnalyticsDeltaAndAuxiliaryEndpoints(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	now := time.Now().UTC().Truncate(time.Minute)
	mustIngest := func(event audit.Event) {
		t.Helper()
		if _, err := store.Ingest(context.Background(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}

	// Baseline window.
	mustIngest(analyticsTestEvent(now.Add(-45*24*time.Hour), "prod", audit.DecisionAllow, "signer-a", "sha256:base-1", 3, 1, "", nil))
	mustIngest(analyticsTestEvent(now.Add(-44*24*time.Hour), "prod", audit.DecisionAllow, "signer-a", "sha256:base-2", 2, 1, "", nil))
	mustIngest(analyticsTestEvent(now.Add(-43*24*time.Hour), "stage", audit.DecisionAllow, "signer-a", "sha256:base-2", 2, 1, "", nil))

	// Current window.
	mustIngest(analyticsTestEvent(now.Add(-10*24*time.Hour), "prod", audit.DecisionDeny, "signer-b", "sha256:curr-1", 2, 2, "EX-1", []string{"workflow mismatch"}))
	mustIngest(analyticsTestEvent(now.Add(-9*24*time.Hour), "prod", audit.DecisionDeny, "signer-c", "sha256:curr-2", 1, 2, "EX-2", []string{"workflow mismatch"}))
	mustIngest(analyticsTestEvent(now.Add(-8*24*time.Hour), "stage", audit.DecisionAllow, "signer-c", "sha256:curr-2", 1, 2, "", nil))
	mustIngest(analyticsTestEvent(now.Add(-7*24*time.Hour), "prod", audit.DecisionAllow, "signer-c", "sha256:curr-3", 0, 1, "", nil))
	mustIngest(analyticsTestEvent(now.Add(-6*24*time.Hour), "prod", audit.DecisionAllow, "signer-c", "sha256:curr-4", 0, 0, "", nil))

	handler := newHandler(store, "memory")

	run := func(path string, target any) {
		t.Helper()
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200 for %s, got %d: %s", path, rec.Code, rec.Body.String())
		}
		if err := json.NewDecoder(rec.Body).Decode(target); err != nil {
			t.Fatalf("decode %s: %v", path, err)
		}
	}

	var trends audit.TrendsResponse
	run("/v1/analytics/trends?window=28d&group_by=service&tenant_id=acme", &trends)
	if len(trends.MetricTrends) != len(analyticsMetricDefinitions()) {
		t.Fatalf("expected metric trends in trends response, got %#v", trends)
	}
	if trends.SchemaVersion != audit.TrendsSchemaVersion {
		t.Fatalf("expected schema-versioned trends response, got %#v", trends)
	}
	if trends.Comparison == nil || trends.Comparison.GroupBy != "service" {
		t.Fatalf("expected comparison context on trends response, got %#v", trends.Comparison)
	}

	var delta audit.AnalyticsDeltaResponse
	run("/v1/analytics/delta?window=28d&metric=policy_friction_rate&group_by=service&tenant_id=acme", &delta)
	if delta.Definition.Key != analyticsMetricPolicyFrictionRate {
		t.Fatalf("unexpected delta definition %#v", delta.Definition)
	}
	if len(delta.Segments) == 0 {
		t.Fatalf("expected segment deltas, got %#v", delta)
	}
	if delta.SchemaVersion != audit.AnalyticsDeltaSchemaVersion {
		t.Fatalf("expected schema-versioned analytics delta, got %#v", delta)
	}

	var topologyDelta audit.AnalyticsDeltaResponse
	run("/v1/analytics/delta?window=28d&metric=blast_radius_trend&group_by=service&tenant_id=acme", &topologyDelta)
	if topologyDelta.Definition.Key != analyticsMetricBlastRadiusTrend {
		t.Fatalf("unexpected topology delta definition %#v", topologyDelta.Definition)
	}
	if len(topologyDelta.Segments) == 0 || !strings.Contains(strings.ToLower(topologyDelta.Summary), "blast radius") {
		t.Fatalf("expected first-class topology metric output, got %#v", topologyDelta)
	}

	var anomalies audit.AnalyticsAnomaliesResponse
	run("/v1/analytics/anomalies?window=28d&group_by=service&tenant_id=acme", &anomalies)
	if len(anomalies.Items) == 0 {
		t.Fatalf("expected explainable anomalies, got %#v", anomalies)
	}
	if anomalies.SchemaVersion != audit.AnalyticsAnomaliesSchemaVersion {
		t.Fatalf("expected schema-versioned analytics anomalies, got %#v", anomalies)
	}

	var scorecards audit.AnalyticsScorecardsResponse
	run("/v1/analytics/scorecards?window=28d&tenant_id=acme", &scorecards)
	if len(scorecards.Cards) != len(analyticsMetricDefinitions()) {
		t.Fatalf("expected scorecards for all analytics metrics, got %#v", scorecards)
	}
	if scorecards.SchemaVersion != audit.AnalyticsScorecardsSchemaVersion {
		t.Fatalf("expected schema-versioned analytics scorecards, got %#v", scorecards)
	}

	var segments audit.AnalyticsSegmentsResponse
	run("/v1/analytics/segments?window=28d&tenant_id=acme", &segments)
	if len(segments.Items) != 3 {
		t.Fatalf("expected team/service/environment segment catalog, got %#v", segments)
	}
	if segments.SchemaVersion != audit.AnalyticsSegmentsSchemaVersion {
		t.Fatalf("expected schema-versioned analytics segments, got %#v", segments)
	}
}

func analyticsTestEvent(timestamp time.Time, environment string, decision string, signer string, digest string, critical int, high int, exceptionID string, reasons []string) audit.Event {
	return audit.Event{
		Timestamp:      timestamp,
		Component:      "deploy-gate",
		EventType:      audit.EventTypeDeployGateDecision,
		Decision:       decision,
		TenantID:       "acme",
		Repo:           "acme/payments-api",
		Environment:    environment,
		Namespace:      environment + "-acme",
		Workload:       "payments-api",
		Digest:         digest,
		Reasons:        reasons,
		IsException:    exceptionID != "",
		ExceptionID:    exceptionID,
		PolicyBundleID: "bundle-9c",
		VerifierSummary: &audit.VerifierSummary{
			SignatureValid:   signer != "",
			AttestationValid: true,
		},
		Evidence: &audit.Evidence{
			Artifact: &audit.ArtifactEvidence{
				SignerIdentity: signer,
				SBOMHash:       "sbom-" + strings.TrimPrefix(digest, "sha256:"),
				VulnerabilitySummary: &audit.VulnerabilitySummary{
					Critical: critical,
					High:     high,
					Total:    critical + high,
				},
			},
		},
	}
}
