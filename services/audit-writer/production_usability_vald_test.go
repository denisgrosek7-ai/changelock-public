package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

func TestProductionUsabilityValDFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/production/usability-operability-recovery/vald/config-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDConfigReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode config review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDConfigReviewStateActive || response.Model.ReviewState != operability.ProductionUsabilityFinalGatePass {
					t.Fatalf("unexpected config review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/explainability-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDExplainabilityReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode explainability review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDExplainabilityReviewStateActive || !response.Model.PolicyRefsPresent {
					t.Fatalf("unexpected explainability review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/dry-run-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDDryRunReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode dry-run review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDDryRunReviewStateActive || response.Model.PreviewSuccessImpliesActivate {
					t.Fatalf("unexpected dry-run review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/redaction-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDRedactionReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode redaction review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDRedactionReviewStateActive || response.Model.PartnerOrPublicExposeFullEvidence {
					t.Fatalf("unexpected redaction review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/degraded-state-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDDegradedReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode degraded review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDDegradedBehaviorReviewStateActive || response.Model.StaleReportedAsFresh {
					t.Fatalf("unexpected degraded review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/ui-windowing-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDUIWindowingReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode ui/windowing review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDUIWindowingReviewStateActive || response.Model.LimitExceedsMaxWindow {
					t.Fatalf("unexpected ui/windowing review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/command-noise-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDCommandNoiseReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode command/noise review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDCommandNoiseReviewStateActive || response.Model.UngovernedTaskMutation {
					t.Fatalf("unexpected command/noise review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/api-protection-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDAPIProtectionReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode api protection review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDAPIProtectionReviewStateActive || response.Model.PriorityBypassesGovernance {
					t.Fatalf("unexpected api protection review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/cli-resilience-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDCLIResilienceReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode cli review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDCLIResilienceReviewStateActive || response.Model.MissingRequiredKey {
					t.Fatalf("unexpected cli review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/supportability-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDSupportabilityReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode supportability review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDSupportabilityReviewStateActive || response.Model.SupportBundleManifestMissing {
					t.Fatalf("unexpected supportability review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/recovery-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDRecoveryReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode recovery review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDRecoveryReviewStateActive || response.Model.PolicyBypassRecommended {
					t.Fatalf("unexpected recovery review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/upgrade-rollback-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDUpgradeRollbackReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode upgrade review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDUpgradeRollbackReviewStateActive || response.Model.AdvisoryMutatesState {
					t.Fatalf("unexpected upgrade review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/scale-envelope-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDScaleEnvelopeReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode scale review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDScaleEnvelopeReviewStateActive || response.Model.MarketedAsGuarantee {
					t.Fatalf("unexpected scale review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/governance-boundary-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDGovernanceReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode governance review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDGovernanceBoundaryReviewStateActive || response.Model.ProjectionClaimsCanonicalTruth {
					t.Fatalf("unexpected governance review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vald/regression-gate?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValDRegressionGateResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode regression gate: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValDRegressionGateStateActive || response.Model.MissingCriticalFixtureCoverage {
					t.Fatalf("unexpected regression gate response %#v", response)
				}
			},
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, tc.path, nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200 for %s, got %d: %s", tc.path, rec.Code, rec.Body.String())
		}
		tc.decode(t, rec)
	}
}

func TestProductionUsabilityValDProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedEnterprisePhase4AuthorityArtifacts(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/production/usability-operability-recovery/vald/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val D proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response productionUsabilityValDProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val D proofs: %v", err)
	}
	if response.CurrentState != operability.ProductionUsabilityValDStateActive {
		t.Fatalf("expected active Val D proofs state, got %#v", response)
	}
	if response.Val0FoundationState != operability.ProductionUsabilityVal0StateActive ||
		response.ValACoreState != operability.ProductionUsabilityValAStateActive ||
		response.ValBResilienceState != operability.ProductionUsabilityValBStateActive ||
		response.ValCSupportabilityState != operability.ProductionUsabilityValCStateActive {
		t.Fatalf("expected active Val 0/A/B/C dependencies, got %#v", response)
	}
	if response.ValDState != operability.ProductionUsabilityValDStateActive {
		t.Fatalf("expected active Val D final gate state, got %#v", response)
	}
	if response.Point4State != operability.ProductionUsabilityPoint4StateNotComplete {
		t.Fatalf("expected point 4 to remain not complete, got %#v", response)
	}
	if len(response.WhyPoint4NotPass) == 0 || len(response.SurfaceRefs) < 20 || len(response.EvidenceRefs) < 16 {
		t.Fatalf("expected why_not_pass, surface refs, and evidence refs, got %#v", response)
	}
}

func TestProductionUsabilityValDProofsStayInactiveWithoutDependencies(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/production/usability-operability-recovery/vald/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val D proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response productionUsabilityValDProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode inactive Val D proofs: %v", err)
	}
	if response.CurrentState == operability.ProductionUsabilityValDStateActive {
		t.Fatalf("expected inactive Val D proofs without dependencies, got %#v", response)
	}
	if response.Val0DependencyState == "" || response.ValADependencyState == "" || response.ValBDependencyState == "" || response.ValCDependencyState == "" {
		t.Fatalf("expected explicit dependency states, got %#v", response)
	}
}
