package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStrictValidationExecutionSurfaces(t *testing.T) {
	fixture := forensicsTestFixture(t)

	scenariosReq := httptest.NewRequest(http.MethodGet, "/v1/validation/scenarios", nil)
	scenariosReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	scenariosRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(scenariosRec, scenariosReq)
	if scenariosRec.Code != http.StatusOK {
		t.Fatalf("expected strict validation scenarios 200, got %d: %s", scenariosRec.Code, scenariosRec.Body.String())
	}

	var scenarioResponse validationScenarioListResponse
	if err := json.NewDecoder(scenariosRec.Body).Decode(&scenarioResponse); err != nil {
		t.Fatalf("decode strict validation scenarios: %v", err)
	}
	if len(scenarioResponse.Scenarios) < 7 {
		t.Fatalf("expected strict validation scenario registry, got %#v", scenarioResponse)
	}
	if scenarioResponse.Scenarios[0].Version == "" || len(scenarioResponse.Scenarios[0].CleanupPlan) == 0 {
		t.Fatalf("expected versioned scenario metadata with cleanup, got %#v", scenarioResponse.Scenarios[0])
	}

	executeReq := httptest.NewRequest(http.MethodPost, "/v1/validation/execute?tenant_id=acme&environment=prod", bytes.NewBufferString(`{"scenario_ids":["safe_release_positive","unsigned_image_block"]}`))
	executeReq.Header.Set("Authorization", "Bearer operator-demo-token")
	executeReq.Header.Set("Content-Type", "application/json")
	executeRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(executeRec, executeReq)
	if executeRec.Code != http.StatusOK {
		t.Fatalf("expected strict validation execute 200, got %d: %s", executeRec.Code, executeRec.Body.String())
	}

	var run validationExecutionRun
	if err := json.NewDecoder(executeRec.Body).Decode(&run); err != nil {
		t.Fatalf("decode strict validation run: %v", err)
	}
	if run.RunID == "" || len(run.Executions) != 2 || len(run.Verdicts) != 2 {
		t.Fatalf("expected strict validation run with executions and verdicts, got %#v", run)
	}
	if run.Certificate.CertificateID == "" || !run.Certificate.SealReady {
		t.Fatalf("expected seal-ready validation certificate, got %#v", run.Certificate)
	}
	if !strings.Contains(run.Executions[0].Namespace, "validation") || len(run.Executions[0].CleanupPlan) == 0 || len(run.Executions[0].RollbackPlan) == 0 {
		t.Fatalf("expected isolated execution metadata with cleanup/rollback, got %#v", run.Executions[0])
	}
	if run.Verdicts[0].VerdictID == "" || run.Verdicts[0].ExpectedOutcome.Verdict == "" || run.Verdicts[0].ObservedOutcome.Summary == "" {
		t.Fatalf("expected strict verdict contract with expected/observed split, got %#v", run.Verdicts[0])
	}

	executionsReq := httptest.NewRequest(http.MethodGet, "/v1/validation/executions?tenant_id=acme&environment=prod", nil)
	executionsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	executionsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(executionsRec, executionsReq)
	if executionsRec.Code != http.StatusOK {
		t.Fatalf("expected strict validation executions 200, got %d: %s", executionsRec.Code, executionsRec.Body.String())
	}

	var executionList validationExecutionListResponse
	if err := json.NewDecoder(executionsRec.Body).Decode(&executionList); err != nil {
		t.Fatalf("decode strict validation execution list: %v", err)
	}
	if len(executionList.Executions) == 0 {
		t.Fatalf("expected strict validation execution list, got %#v", executionList)
	}

	executionReq := httptest.NewRequest(http.MethodGet, "/v1/validation/executions/"+run.Executions[0].ExecutionID+"?tenant_id=acme&environment=prod", nil)
	executionReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	executionRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(executionRec, executionReq)
	if executionRec.Code != http.StatusOK {
		t.Fatalf("expected strict validation execution by id 200, got %d: %s", executionRec.Code, executionRec.Body.String())
	}

	verdictReq := httptest.NewRequest(http.MethodGet, "/v1/validation/verdicts/"+run.Verdicts[0].VerdictID+"?tenant_id=acme&environment=prod", nil)
	verdictReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	verdictRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(verdictRec, verdictReq)
	if verdictRec.Code != http.StatusOK {
		t.Fatalf("expected strict validation verdict by id 200, got %d: %s", verdictRec.Code, verdictRec.Body.String())
	}

	var verdict validationVerdict
	if err := json.NewDecoder(verdictRec.Body).Decode(&verdict); err != nil {
		t.Fatalf("decode strict validation verdict: %v", err)
	}
	if verdict.ExecutionID == "" || verdict.ExpectedOutcome.LatencyThresholdMS == 0 {
		t.Fatalf("expected strict validation verdict payload, got %#v", verdict)
	}

	certificateReq := httptest.NewRequest(http.MethodGet, "/v1/validation/certificates/"+run.Certificate.CertificateID+"?tenant_id=acme&environment=prod", nil)
	certificateReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	certificateRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(certificateRec, certificateReq)
	if certificateRec.Code != http.StatusOK {
		t.Fatalf("expected strict validation certificate by id 200, got %d: %s", certificateRec.Code, certificateRec.Body.String())
	}

	var certificate validationCertificate
	if err := json.NewDecoder(certificateRec.Body).Decode(&certificate); err != nil {
		t.Fatalf("decode strict validation certificate: %v", err)
	}
	if certificate.Scope == "" || len(certificate.ScenarioResults) == 0 {
		t.Fatalf("expected strict validation certificate payload, got %#v", certificate)
	}
}

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

func TestStrictValidationRegressionChaosAndCompatibilityRuns(t *testing.T) {
	fixture := forensicsTestFixture(t)

	regressionReq := httptest.NewRequest(http.MethodPost, "/v1/validation/regression/run?tenant_id=acme&environment=prod", nil)
	regressionReq.Header.Set("Authorization", "Bearer operator-demo-token")
	regressionRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(regressionRec, regressionReq)
	if regressionRec.Code != http.StatusOK {
		t.Fatalf("expected validation regression run 200, got %d: %s", regressionRec.Code, regressionRec.Body.String())
	}

	var regression validationExecutionRun
	if err := json.NewDecoder(regressionRec.Body).Decode(&regression); err != nil {
		t.Fatalf("decode validation regression run: %v", err)
	}
	if regression.Mode != validationModeRegression || len(regression.Executions) < 3 {
		t.Fatalf("expected regression suite execution bundle, got %#v", regression)
	}

	chaosReq := httptest.NewRequest(http.MethodPost, "/v1/validation/chaos/run?tenant_id=acme&environment=prod", nil)
	chaosReq.Header.Set("Authorization", "Bearer operator-demo-token")
	chaosRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(chaosRec, chaosReq)
	if chaosRec.Code != http.StatusOK {
		t.Fatalf("expected validation chaos run 200, got %d: %s", chaosRec.Code, chaosRec.Body.String())
	}

	var chaos validationExecutionRun
	if err := json.NewDecoder(chaosRec.Body).Decode(&chaos); err != nil {
		t.Fatalf("decode validation chaos run: %v", err)
	}
	if chaos.Mode != validationModeControlledChaos || len(chaos.Executions) == 0 {
		t.Fatalf("expected chaos execution bundle, got %#v", chaos)
	}
	if chaos.Executions[0].ApprovalMode != recommendationApprovalHumanReview {
		t.Fatalf("expected approval-gated chaos execution, got %#v", chaos.Executions[0])
	}

	compatReq := httptest.NewRequest(http.MethodPost, "/v1/validation/compatibility/run?tenant_id=acme&environment=prod", bytes.NewBufferString(`{"kubernetes_version":"1.33","rekor_unavailable":true,"tighten_runtime_restrictions":true}`))
	compatReq.Header.Set("Authorization", "Bearer operator-demo-token")
	compatReq.Header.Set("Content-Type", "application/json")
	compatRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(compatRec, compatReq)
	if compatRec.Code != http.StatusOK {
		t.Fatalf("expected validation compatibility run 200, got %d: %s", compatRec.Code, compatRec.Body.String())
	}

	var compatibility validationExecutionRun
	if err := json.NewDecoder(compatRec.Body).Decode(&compatibility); err != nil {
		t.Fatalf("decode validation compatibility run: %v", err)
	}
	if compatibility.Mode != validationModeCompatibility || !compatibility.SimulationDerived {
		t.Fatalf("expected simulation-derived compatibility run, got %#v", compatibility)
	}
	if len(compatibility.ChangeSet) == 0 || len(compatibility.CompatibilityRisks) == 0 {
		t.Fatalf("expected compatibility change set and risks, got %#v", compatibility)
	}
}
