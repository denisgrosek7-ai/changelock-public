package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

func TestReferenceArchitectureValBFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/reference-architecture/valb/pack-registry?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValBPackRegistryResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode pack registry: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValBPackStateActive || len(response.Model.Packs) != 6 || len(response.FamilyStates) != 6 {
					t.Fatalf("unexpected pack registry response %#v", response)
				}
			},
		},
		{
			path: "/v1/reference-architecture/valb/artifact-manifests?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValBCollectionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode artifact manifests: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValBManifestStateActive || len(response.FamilyStates) != 6 {
					t.Fatalf("unexpected artifact manifest response %#v", response)
				}
			},
		},
		{
			path: "/v1/reference-architecture/valb/readiness-checks?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValBCollectionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode readiness checks: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValBReadinessStateActive || len(response.FamilyStates) != 6 {
					t.Fatalf("unexpected readiness response %#v", response)
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

func TestReferenceArchitectureValBProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/reference-architecture/valb/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response referenceArchitectureValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.ReferenceArchitectureValBStateActive {
		t.Fatalf("expected active proofs state, got %#v", response)
	}
	if response.Point5State != operability.IntelligenceCalibrationPoint5StatePass ||
		response.Val0DependencyState != operability.ReferenceArchitectureVal0StateActive ||
		response.Val0State != operability.ReferenceArchitectureVal0StateActive ||
		response.ValADependencyState != operability.ReferenceArchitectureValAStateActive ||
		response.ValAState != operability.ReferenceArchitectureValAStateActive {
		t.Fatalf("expected active dependency states, got %#v", response)
	}
	if response.ValBState != operability.ReferenceArchitectureValBStateActive ||
		response.PackRegistryState != operability.ReferenceArchitectureValBPackStateActive ||
		response.ArtifactManifestState != operability.ReferenceArchitectureValBManifestStateActive ||
		response.BundleState != operability.ReferenceArchitectureValBBundleStateActive ||
		response.ReadinessState != operability.ReferenceArchitectureValBReadinessStateActive ||
		response.ValidationHookState != operability.ReferenceArchitectureValBHookStateActive ||
		response.ConformanceKitState != operability.ReferenceArchitectureValBConformanceKitStateActive ||
		response.DeviationState != operability.ReferenceArchitectureValBDeviationStateActive {
		t.Fatalf("expected active Val B component states, got %#v", response)
	}
	if response.Point6State != operability.ReferenceArchitecturePoint6StateNotComplete {
		t.Fatalf("expected point 6 to remain not complete, got %#v", response)
	}
	if len(response.SupportedFamilies) != 6 || len(response.FamilyStates) != 6 || len(response.WhyPoint6NotPass) == 0 || len(response.SurfaceRefs) < 10 || len(response.EvidenceRefs) < 10 || len(response.Limitations) == 0 {
		t.Fatalf("expected supported families, refs, and limitations in proofs, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection-only disclaimer, got %#v", response)
	}
}
