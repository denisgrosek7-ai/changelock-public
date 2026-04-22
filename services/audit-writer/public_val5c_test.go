package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPublicBenchmarkMethodologyHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/benchmarks/methodology", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public benchmark methodology 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicBenchmarkMethodologyResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public benchmark methodology: %v", err)
	}
	if response.SchemaVersion != publicBenchmarkMethodologySchema || response.MethodologyID == "" {
		t.Fatalf("expected benchmark methodology payload, got %#v", response)
	}
	if !containsString(response.NotMeasured, "No public benchmark in this methodology claims universal security gain percentages, universal runtime prevention rates, or blanket latency guarantees across all substrates.") {
		t.Fatalf("expected bounded not-measured guardrail, got %#v", response.NotMeasured)
	}
}

func TestPublicBenchmarkSetHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/benchmarks/set", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public benchmark set 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicBenchmarkSetResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public benchmark set: %v", err)
	}
	if response.SchemaVersion != publicBenchmarkSetSchema || len(response.Benchmarks) < 8 {
		t.Fatalf("expected public benchmark set catalog, got %#v", response)
	}
	item := findPublicBenchmarkDefinition(t, response.Benchmarks, "runtime_overhead")
	if item.PublicationStatus != "starting_points_only_not_public_claim" {
		t.Fatalf("expected bounded runtime overhead status, got %#v", item)
	}
}

func TestPublicAnalyticsPublicationDisciplineHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/analytics/publication-discipline", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected analytics publication discipline 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicAnalyticsPublicationDisciplineResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public analytics publication discipline: %v", err)
	}
	if response.SchemaVersion != publicAnalyticsPublicationDisciplineSchema || response.PublicationMode != "aggregated_and_anonymized_only" {
		t.Fatalf("expected analytics publication discipline payload, got %#v", response)
	}
	if len(response.DoNotPublishConditions) == 0 {
		t.Fatalf("expected do-not-publish conditions, got %#v", response)
	}
}

func TestPublicCaseStudyPacksHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/case-studies", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public case study packs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicCaseStudyPacksResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public case study packs: %v", err)
	}
	if response.SchemaVersion != publicCaseStudyPacksSchema || len(response.Packs) < 3 {
		t.Fatalf("expected public case study packs, got %#v", response)
	}
	pack := findPublicCaseStudyPack(t, response.Packs, "offline_handoff_verification")
	if len(pack.EvidenceRefs) == 0 || !containsString(pack.EvidenceRefs, "/v1/public/verifier/reference-pack") {
		t.Fatalf("expected replayable offline handoff pack, got %#v", pack)
	}
}

func findPublicBenchmarkDefinition(t *testing.T, items []publicBenchmarkDefinition, benchmarkID string) publicBenchmarkDefinition {
	t.Helper()
	for _, item := range items {
		if item.BenchmarkID == benchmarkID {
			return item
		}
	}
	t.Fatalf("expected benchmark %q, got %#v", benchmarkID, items)
	return publicBenchmarkDefinition{}
}

func findPublicCaseStudyPack(t *testing.T, items []publicCaseStudyPack, packID string) publicCaseStudyPack {
	t.Helper()
	for _, item := range items {
		if item.PackID == packID {
			return item
		}
	}
	t.Fatalf("expected case study pack %q, got %#v", packID, items)
	return publicCaseStudyPack{}
}
