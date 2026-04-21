package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPublicReferenceArchitecturesHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/reference-architectures", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public reference architectures 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicReferenceArchitecturesResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public reference architectures: %v", err)
	}
	if response.SchemaVersion != publicReferenceArchitecturesSchema || len(response.Architectures) < 5 {
		t.Fatalf("expected reference architecture catalog, got %#v", response)
	}
	if findPublicReferenceArchitecture(t, response.Architectures, "air-gapped-regulated").DisplayName == "" {
		t.Fatalf("expected air-gapped regulated architecture, got %#v", response.Architectures)
	}
	if !containsString(response.DecisionGuideRefs, "/v1/public/decision-guides") {
		t.Fatalf("expected decision guide ref, got %#v", response.DecisionGuideRefs)
	}
}

func TestPublicMaturityMapHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/maturity-map", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public maturity map 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicMaturityMapResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public maturity map: %v", err)
	}
	if response.SchemaVersion != publicMaturityMapSchema || len(response.Levels) < 5 {
		t.Fatalf("expected maturity map levels, got %#v", response)
	}
	level := findPublicMaturityLevel(t, response.Levels, "level_5_public_verifiability")
	if len(level.RequiredCapabilities) == 0 || !containsString(level.RequiredCapabilities, "public verification specs") {
		t.Fatalf("expected public verifiability maturity level, got %#v", level)
	}
}

func TestPublicDecisionGuidesHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/decision-guides", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public decision guides 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicDecisionGuidesResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public decision guides: %v", err)
	}
	if response.SchemaVersion != publicDecisionGuidesSchema || len(response.Guides) < 5 {
		t.Fatalf("expected public decision guide catalog, got %#v", response)
	}
	guide := findPublicDecisionGuide(t, response.Guides, "b2b_proof_exchange_adoption")
	if len(guide.RelatedContracts) == 0 || len(guide.ReferenceArchitectureRefs) == 0 {
		t.Fatalf("expected bounded B2B decision guide linkage, got %#v", guide)
	}
}

func TestPublicSectorProfilesHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/reference-architectures/sector-profiles", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public sector profiles 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicSectorProfilesResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public sector profiles: %v", err)
	}
	if response.SchemaVersion != publicSectorProfilesSchema || len(response.Profiles) < 5 {
		t.Fatalf("expected sector profiles catalog, got %#v", response)
	}
	profile := findPublicSectorProfile(t, response.Profiles, "air_gapped_regulated")
	if len(profile.RequiredContracts) == 0 || !containsString(profile.RequiredContracts, "/v1/public/verifier/offline-guide") {
		t.Fatalf("expected offline verifier linkage in air-gapped profile, got %#v", profile)
	}
	if response.MaturityMapRef != "/v1/public/maturity-map" {
		t.Fatalf("expected maturity map ref, got %#v", response)
	}
}

func TestPublicDeploymentDecisionMatrixHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/decision-guides/matrix", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public decision matrix 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicDeploymentDecisionMatrixResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public decision matrix: %v", err)
	}
	if response.SchemaVersion != publicDecisionMatrixSchema || len(response.Rows) < 5 {
		t.Fatalf("expected deployment decision matrix, got %#v", response)
	}
	row := findPublicDeploymentDecisionRow(t, response.Rows, "partner_trust_exchange")
	if len(row.RequiredContracts) == 0 || !containsString(row.RequiredContracts, "/v1/b2b/suppliers/onboarding") {
		t.Fatalf("expected partner exchange contract linkage, got %#v", row)
	}
}

func findPublicReferenceArchitecture(t *testing.T, items []publicReferenceArchitecture, architectureID string) publicReferenceArchitecture {
	t.Helper()
	for _, item := range items {
		if item.ArchitectureID == architectureID {
			return item
		}
	}
	t.Fatalf("expected reference architecture %q, got %#v", architectureID, items)
	return publicReferenceArchitecture{}
}

func findPublicMaturityLevel(t *testing.T, items []publicMaturityLevel, levelID string) publicMaturityLevel {
	t.Helper()
	for _, item := range items {
		if item.LevelID == levelID {
			return item
		}
	}
	t.Fatalf("expected maturity level %q, got %#v", levelID, items)
	return publicMaturityLevel{}
}

func findPublicDecisionGuide(t *testing.T, items []publicDecisionGuide, guideID string) publicDecisionGuide {
	t.Helper()
	for _, item := range items {
		if item.GuideID == guideID {
			return item
		}
	}
	t.Fatalf("expected decision guide %q, got %#v", guideID, items)
	return publicDecisionGuide{}
}

func findPublicSectorProfile(t *testing.T, items []publicSectorProfile, profileID string) publicSectorProfile {
	t.Helper()
	for _, item := range items {
		if item.SectorProfileID == profileID {
			return item
		}
	}
	t.Fatalf("expected sector profile %q, got %#v", profileID, items)
	return publicSectorProfile{}
}

func findPublicDeploymentDecisionRow(t *testing.T, items []publicDeploymentDecisionRow, decisionID string) publicDeploymentDecisionRow {
	t.Helper()
	for _, item := range items {
		if item.DecisionID == decisionID {
			return item
		}
	}
	t.Fatalf("expected deployment decision row %q, got %#v", decisionID, items)
	return publicDeploymentDecisionRow{}
}
