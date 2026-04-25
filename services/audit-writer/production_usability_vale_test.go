package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

func TestProductionUsabilityValEFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedEnterprisePhase4AuthorityArtifacts(t, handler)

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/production/usability-operability-recovery/vale/dependency-closure?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValEDependencyClosureResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode dependency closure: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValEDependencyClosureStateActive || response.Model.DependencyStatus != operability.ProductionUsabilityDependencyPass {
					t.Fatalf("unexpected dependency closure response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vale/coherence-review?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValECoherenceReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode coherence review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValECoherenceReviewStateActive || len(response.Model.CarriedForwardLimitations) == 0 {
					t.Fatalf("unexpected coherence review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vale/pass-rule?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValEPassRuleResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode pass rule: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValEPassRuleStateActive || response.Model.Point4State != operability.ProductionUsabilityPoint4StatePass {
					t.Fatalf("unexpected pass rule response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vale/canonical-truth-boundary?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValECanonicalBoundaryReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode canonical boundary review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValECanonicalBoundaryReviewStateActive || response.Model.ProjectionClaimsCanonicalTruth {
					t.Fatalf("unexpected canonical boundary review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vale/redaction-export-review?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValERedactionExportReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode redaction/export review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValERedactionExportReviewStateActive || len(response.Model.CheckedScopes) != 5 {
					t.Fatalf("unexpected redaction/export review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vale/supportability-recovery-review?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValESupportabilityRecoveryResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode supportability/recovery review: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValESupportabilityRecoveryReviewStateActive || response.Model.SupportabilityOverridesFailedProof {
					t.Fatalf("unexpected supportability/recovery review response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vale/regression-closure?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValERegressionClosureResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode regression closure: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValERegressionClosureStateActive || len(response.Model.CoveredCategories) == 0 {
					t.Fatalf("unexpected regression closure response %#v", response)
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

func TestProductionUsabilityValEProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedEnterprisePhase4AuthorityArtifacts(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/production/usability-operability-recovery/vale/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val E proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response productionUsabilityValEProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val E proofs: %v", err)
	}
	if response.CurrentState != operability.ProductionUsabilityValEStateActive {
		t.Fatalf("expected active Val E proofs state, got %#v", response)
	}
	if response.Val0FoundationState != operability.ProductionUsabilityVal0StateActive ||
		response.ValACoreState != operability.ProductionUsabilityValAStateActive ||
		response.ValBResilienceState != operability.ProductionUsabilityValBStateActive ||
		response.ValCSupportabilityState != operability.ProductionUsabilityValCStateActive ||
		response.ValDFinalGateState != operability.ProductionUsabilityValDStateActive {
		t.Fatalf("expected active Val 0/A/B/C/D dependencies, got %#v", response)
	}
	if response.ValEState != operability.ProductionUsabilityValEStateActive {
		t.Fatalf("expected active Val E state, got %#v", response)
	}
	if response.Point4State != operability.ProductionUsabilityPoint4StatePass || !response.PassCriteriaMet {
		t.Fatalf("expected point 4 pass only through active Val E, got %#v", response)
	}
	if len(response.Limitations) == 0 || len(response.SurfaceRefs) < 13 || len(response.EvidenceRefs) < 12 {
		t.Fatalf("expected limitations, surface refs, and evidence refs, got %#v", response)
	}
	if response.ProjectionDisclaimer != "projection_only not_canonical_truth integrated_closure_summary" {
		t.Fatalf("expected projection-only disclaimer, got %#v", response)
	}
}

func TestProductionUsabilityValEProofsStayNotCompleteWithoutDependencies(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/production/usability-operability-recovery/vale/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val E proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response productionUsabilityValEProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val E proofs: %v", err)
	}
	if response.Point4State != operability.ProductionUsabilityPoint4StateNotComplete {
		t.Fatalf("expected point 4 to remain not complete without dependencies, got %#v", response)
	}
	if response.ValEState == operability.ProductionUsabilityValEStateActive {
		t.Fatalf("expected non-active Val E state without dependencies, got %#v", response)
	}
}
