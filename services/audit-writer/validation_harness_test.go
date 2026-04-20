package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidationHarnessScenariosScoreRunAndWhatIf(t *testing.T) {
	fixture := forensicsTestFixture(t)

	scenariosReq := httptest.NewRequest(http.MethodGet, "/v1/validation/harness/scenarios", nil)
	scenariosReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	scenariosRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(scenariosRec, scenariosReq)
	if scenariosRec.Code != http.StatusOK {
		t.Fatalf("expected validation scenarios 200, got %d: %s", scenariosRec.Code, scenariosRec.Body.String())
	}

	var scenarios validationHarnessScenarioListResponse
	if err := json.NewDecoder(scenariosRec.Body).Decode(&scenarios); err != nil {
		t.Fatalf("decode validation scenarios: %v", err)
	}
	if len(scenarios.Scenarios) < 5 || scenarios.Scenarios[0].ScenarioID == "" {
		t.Fatalf("expected validation scenario catalog, got %#v", scenarios)
	}

	scoreReq := httptest.NewRequest(http.MethodGet, "/v1/validation/harness/score?tenant_id=acme&environment=prod&limit=10", nil)
	scoreReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	scoreRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(scoreRec, scoreReq)
	if scoreRec.Code != http.StatusOK {
		t.Fatalf("expected validation score 200, got %d: %s", scoreRec.Code, scoreRec.Body.String())
	}

	var score validationHarnessScoreResponse
	if err := json.NewDecoder(scoreRec.Body).Decode(&score); err != nil {
		t.Fatalf("decode validation score: %v", err)
	}
	if len(score.Results) == 0 || score.ConfidenceLevel == "" {
		t.Fatalf("expected validation score results, got %#v", score)
	}

	runReq := httptest.NewRequest(http.MethodPost, "/v1/validation/harness/runs?tenant_id=acme&environment=prod", bytes.NewBufferString(`{"mode":"policy_dry_run"}`))
	runReq.Header.Set("Authorization", "Bearer operator-demo-token")
	runReq.Header.Set("Content-Type", "application/json")
	runRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(runRec, runReq)
	if runRec.Code != http.StatusOK {
		t.Fatalf("expected validation run 200, got %d: %s", runRec.Code, runRec.Body.String())
	}

	var run validationHarnessRun
	if err := json.NewDecoder(runRec.Body).Decode(&run); err != nil {
		t.Fatalf("decode validation run: %v", err)
	}
	if run.RunID == "" || run.CertificateStatus == "" || len(run.Results) == 0 {
		t.Fatalf("expected persisted validation run, got %#v", run)
	}

	runsReq := httptest.NewRequest(http.MethodGet, "/v1/validation/harness/runs?tenant_id=acme&environment=prod&limit=5", nil)
	runsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	runsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(runsRec, runsReq)
	if runsRec.Code != http.StatusOK {
		t.Fatalf("expected validation runs list 200, got %d: %s", runsRec.Code, runsRec.Body.String())
	}

	var runs validationHarnessRunsResponse
	if err := json.NewDecoder(runsRec.Body).Decode(&runs); err != nil {
		t.Fatalf("decode validation runs: %v", err)
	}
	if len(runs.Runs) == 0 || runs.Runs[0].RunID == "" {
		t.Fatalf("expected validation run history, got %#v", runs)
	}

	whatIfReq := httptest.NewRequest(http.MethodPost, "/v1/validation/harness/what-if?tenant_id=acme&environment=prod", bytes.NewBufferString(`{"rekor_unavailable":true,"inject_critical_vulnerability":true,"tighten_runtime_restrictions":true}`))
	whatIfReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	whatIfReq.Header.Set("Content-Type", "application/json")
	whatIfRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(whatIfRec, whatIfReq)
	if whatIfRec.Code != http.StatusOK {
		t.Fatalf("expected validation what-if 200, got %d: %s", whatIfRec.Code, whatIfRec.Body.String())
	}

	var whatIf validationHarnessWhatIfResponse
	if err := json.NewDecoder(whatIfRec.Body).Decode(&whatIf); err != nil {
		t.Fatalf("decode validation what-if: %v", err)
	}
	if len(whatIf.ChangeSet) == 0 || len(whatIf.Results) == 0 {
		t.Fatalf("expected projected validation what-if output, got %#v", whatIf)
	}
}

func TestValidationRecommendationsAndHandoffExport(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)

	runReq := httptest.NewRequest(http.MethodPost, "/v1/validation/harness/runs?tenant_id=acme&environment=prod", bytes.NewBufferString(`{"mode":"policy_dry_run"}`))
	runReq.Header.Set("Authorization", "Bearer operator-demo-token")
	runReq.Header.Set("Content-Type", "application/json")
	runRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(runRec, runReq)
	if runRec.Code != http.StatusOK {
		t.Fatalf("expected validation run 200, got %d: %s", runRec.Code, runRec.Body.String())
	}

	recommendationReq := httptest.NewRequest(http.MethodGet, "/v1/recommendations?tenant_id=acme&environment=prod&source_type=validation_signal", nil)
	recommendationReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	recommendationRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(recommendationRec, recommendationReq)
	if recommendationRec.Code != http.StatusOK {
		t.Fatalf("expected validation recommendations 200, got %d: %s", recommendationRec.Code, recommendationRec.Body.String())
	}

	var recommendations recommendationListResponse
	if err := json.NewDecoder(recommendationRec.Body).Decode(&recommendations); err != nil {
		t.Fatalf("decode validation recommendations: %v", err)
	}
	if len(recommendations.Recommendations) == 0 {
		t.Fatalf("expected validation recommendation candidate, got %#v", recommendations)
	}
	if recommendations.Recommendations[0].SourceType != "validation_signal" {
		t.Fatalf("expected validation recommendation source, got %#v", recommendations.Recommendations[0])
	}

	sealReq := httptest.NewRequest(http.MethodPost, "/v1/handoff/seal?tenant_id=acme&environment=prod", bytes.NewBufferString(`{"audience":"internal","include_validation":true}`))
	sealReq.Header.Set("Authorization", "Bearer operator-demo-token")
	sealReq.Header.Set("Content-Type", "application/json")
	sealRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(sealRec, sealReq)
	if sealRec.Code != http.StatusOK {
		t.Fatalf("expected validation-inclusive handoff seal 200, got %d: %s", sealRec.Code, sealRec.Body.String())
	}

	var sealed handoffSealResponse
	if err := json.NewDecoder(sealRec.Body).Decode(&sealed); err != nil {
		t.Fatalf("decode validation-inclusive handoff: %v", err)
	}
	if sealed.PackageID == "" || len(sealed.Manifest.Artifacts) == 0 {
		t.Fatalf("expected sealed handoff payload, got %#v", sealed)
	}
	foundValidationArtifact := false
	for _, artifact := range sealed.Manifest.Artifacts {
		if artifact.Path == "evidence/validation_harness.json" {
			foundValidationArtifact = true
			break
		}
	}
	if !foundValidationArtifact {
		t.Fatalf("expected validation harness artifact in sealed handoff manifest, got %#v", sealed.Manifest.Artifacts)
	}
}
