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

func TestReferenceArchitectureValAFamilyHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/reference-architecture/vala/family-registry?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValAFamilyRegistryResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode family registry: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValAFamilyRegistryStateActive || len(response.Model.Profiles) != 6 || len(response.FamilyStates) != 6 {
					t.Fatalf("unexpected family registry response %#v", response)
				}
			},
		},
		{
			path: "/v1/reference-architecture/vala/family-profiles?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValAFamilyProfilesResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode family profiles: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValAFamilyRegistryStateActive || len(response.Profiles) != 6 || len(response.FamilyStates) != 6 {
					t.Fatalf("unexpected family profiles response %#v", response)
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

func TestReferenceArchitectureValAProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/reference-architecture/vala/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response referenceArchitectureValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.ReferenceArchitectureValAStateActive {
		t.Fatalf("expected active proofs state, got %#v", response)
	}
	if response.Point5State != operability.IntelligenceCalibrationPoint5StatePass || response.Val0DependencyState != operability.ReferenceArchitectureVal0StateActive || response.Val0State != operability.ReferenceArchitectureVal0StateActive {
		t.Fatalf("expected active point5 and val0 dependencies, got %#v", response)
	}
	if response.ValAState != operability.ReferenceArchitectureValAStateActive || response.RegistryState != operability.ReferenceArchitectureValAFamilyRegistryStateActive {
		t.Fatalf("expected active Val A states, got %#v", response)
	}
	if response.Point6State != operability.ReferenceArchitecturePoint6StateNotComplete {
		t.Fatalf("expected point 6 to remain not complete, got %#v", response)
	}
	if len(response.SupportedFamilies) != 6 || len(response.FamilyStates) != 6 || len(response.WhyPoint6NotPass) == 0 || len(response.SurfaceRefs) < 4 || len(response.EvidenceRefs) < 8 || len(response.Limitations) == 0 {
		t.Fatalf("expected supported families, family states, refs, and limitations in proofs, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection-only disclaimer, got %#v", response)
	}
	for _, state := range response.FamilyStates {
		if state.CurrentState != operability.ReferenceArchitectureValAFamilyProfileStateActive {
			t.Fatalf("expected active per-family state, got %#v", response)
		}
	}
}
