package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPublicTrustBadgeProgramHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/trust-program/badges", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public trust badge program 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicTrustBadgeProgramResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public trust badge program: %v", err)
	}
	if response.SchemaVersion != publicTrustBadgeProgramSchema || len(response.BadgeDefinitions) < 3 {
		t.Fatalf("expected trust badge program catalog, got %#v", response)
	}
	badge := findPublicTrustBadgeDefinition(t, response.BadgeDefinitions, "verification_ready")
	if len(badge.EvidenceRequirements) == 0 || badge.ValidityPeriod == "" {
		t.Fatalf("expected evidence and validity semantics on badge definition, got %#v", badge)
	}
}

func TestPublicVerifierProgramHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/trust-program/verifier-program", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public verifier program 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicVerifierProgramResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public verifier program: %v", err)
	}
	if response.SchemaVersion != publicVerifierProgramSchema || len(response.Profiles) < 4 {
		t.Fatalf("expected verifier program profiles, got %#v", response)
	}
	profile := findPublicVerifierProgramProfile(t, response.Profiles, "partner_verifier")
	if len(profile.ConformanceTargets) == 0 || len(response.DisputeMismatchModel) == 0 {
		t.Fatalf("expected partner verifier guidance and dispute model, got profile=%#v response=%#v", profile, response)
	}
}

func TestPublicClaimsGovernanceHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/trust-program/claims-governance", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public claims governance 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicClaimsGovernanceResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public claims governance: %v", err)
	}
	if response.SchemaVersion != publicClaimsGovernanceSchema || len(response.ClaimClasses) < 4 {
		t.Fatalf("expected public claims governance classes, got %#v", response)
	}
	claimClass := findPublicClaimClass(t, response.ClaimClasses, "benchmark_claim")
	if len(claimClass.BenchmarkDiscipline) == 0 || len(claimClass.DisallowedLanguage) == 0 {
		t.Fatalf("expected benchmark claim discipline, got %#v", claimClass)
	}
}

func TestPublicTrustMarkLifecycleHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/trust-program/marks", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public trust mark lifecycle 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicTrustMarkLifecycleResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public trust mark lifecycle: %v", err)
	}
	if response.SchemaVersion != publicTrustMarkLifecycleSchema || len(response.Marks) < 2 {
		t.Fatalf("expected public trust mark lifecycle catalog, got %#v", response)
	}
	revoked := findPublicTrustMarkStatus(t, response.Marks, "mark-benchmark-discipline-sample")
	if revoked.CurrentState != "revoked" || revoked.RevocationReason == "" {
		t.Fatalf("expected revoked sample trust mark, got %#v", revoked)
	}

	markReq := httptest.NewRequest(http.MethodGet, "/v1/public/trust-program/marks/mark-verification-ready-sample", nil)
	markRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(markRec, markReq)
	if markRec.Code != http.StatusOK {
		t.Fatalf("expected public trust mark lookup 200, got %d: %s", markRec.Code, markRec.Body.String())
	}

	var mark publicTrustMarkStatus
	if err := json.NewDecoder(markRec.Body).Decode(&mark); err != nil {
		t.Fatalf("decode public trust mark lookup: %v", err)
	}
	if mark.MarkID != "mark-verification-ready-sample" || len(mark.HistoricalStatus) == 0 {
		t.Fatalf("expected concrete trust mark lookup with history, got %#v", mark)
	}
}

func findPublicTrustBadgeDefinition(t *testing.T, items []publicTrustBadgeDefinition, badgeID string) publicTrustBadgeDefinition {
	t.Helper()
	for _, item := range items {
		if item.BadgeID == badgeID {
			return item
		}
	}
	t.Fatalf("expected public trust badge definition %q, got %#v", badgeID, items)
	return publicTrustBadgeDefinition{}
}

func findPublicVerifierProgramProfile(t *testing.T, items []publicVerifierProgramProfile, profileID string) publicVerifierProgramProfile {
	t.Helper()
	for _, item := range items {
		if item.ProfileID == profileID {
			return item
		}
	}
	t.Fatalf("expected public verifier program profile %q, got %#v", profileID, items)
	return publicVerifierProgramProfile{}
}

func findPublicClaimClass(t *testing.T, items []publicClaimClass, claimClass string) publicClaimClass {
	t.Helper()
	for _, item := range items {
		if item.ClaimClass == claimClass {
			return item
		}
	}
	t.Fatalf("expected public claim class %q, got %#v", claimClass, items)
	return publicClaimClass{}
}

func findPublicTrustMarkStatus(t *testing.T, items []publicTrustMarkStatus, markID string) publicTrustMarkStatus {
	t.Helper()
	for _, item := range items {
		if item.MarkID == markID {
			return item
		}
	}
	t.Fatalf("expected public trust mark %q, got %#v", markID, items)
	return publicTrustMarkStatus{}
}
