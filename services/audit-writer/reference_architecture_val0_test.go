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

func TestReferenceArchitectureVal0FoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/reference-architecture/val0/blueprint-discipline?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureVal0BlueprintResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode blueprint discipline: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureVal0BlueprintDisciplineStateActive || response.Model.BlueprintID == "" {
					t.Fatalf("unexpected blueprint discipline response %#v", response)
				}
			},
		},
		{
			path: "/v1/reference-architecture/val0/environment-fit?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureVal0BlueprintResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode environment fit: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureVal0EnvironmentFitStateActive || response.Model.TargetEnvironment.TrustAnchorMode == "" {
					t.Fatalf("unexpected environment fit response %#v", response)
				}
			},
		},
		{
			path: "/v1/reference-architecture/val0/conformance-evidence?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureVal0BlueprintResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode conformance evidence: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureVal0EvidenceDisciplineStateActive || len(response.Model.EvidenceRefs) < 5 {
					t.Fatalf("unexpected conformance evidence response %#v", response)
				}
			},
		},
		{
			path: "/v1/reference-architecture/val0/compatibility-baseline?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureVal0BlueprintResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode compatibility baseline: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureVal0CompatibilityBaselineStateActive || response.Model.CompatibilityState != operability.ReferenceArchitectureCompatibilityCompatible {
					t.Fatalf("unexpected compatibility response %#v", response)
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

func TestReferenceArchitectureVal0ProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/reference-architecture/val0/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response referenceArchitectureVal0ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.ReferenceArchitectureVal0StateActive {
		t.Fatalf("expected active proofs state, got %#v", response)
	}
	if response.Point5DependencyState != operability.IntelligenceCalibrationValEStateActive || response.Point5State != operability.IntelligenceCalibrationPoint5StatePass {
		t.Fatalf("expected active point 5 dependency closure, got %#v", response)
	}
	if response.Val0State != operability.ReferenceArchitectureVal0StateActive {
		t.Fatalf("expected active Val 0 state, got %#v", response)
	}
	if response.Point6State != operability.ReferenceArchitecturePoint6StateNotComplete {
		t.Fatalf("expected point 6 to remain not complete in Val 0, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection-only disclaimer, got %#v", response)
	}
	if len(response.WhyPoint6NotPass) == 0 || len(response.SurfaceRefs) < 5 || len(response.EvidenceRefs) < 6 || len(response.Limitations) == 0 {
		t.Fatalf("expected why_not_pass, refs, and limitations in proofs, got %#v", response)
	}
	if containsSubstringValue(response.IntegrationSummary, "certified architecture") {
		t.Fatalf("did not expect certification language in integration summary, got %#v", response)
	}
}

func containsSubstringValue(values []string, expected string) bool {
	for _, value := range values {
		if strings.Contains(value, expected) {
			return true
		}
	}
	return false
}
