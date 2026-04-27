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

func TestReferenceArchitectureValDFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/reference-architecture/vald/operational-visibility?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValDVisibilityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode operational visibility: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValDVisibilityStateActive || len(response.Model.Reports) != 6 || len(response.FamilyStates) != 6 {
					t.Fatalf("unexpected operational visibility response %#v", response)
				}
			},
		},
		{
			path: "/v1/reference-architecture/vald/compatibility-gate?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValDCollectionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode compatibility gate: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValDCompatibilityGateStateActive || len(response.FamilyStates) != 6 {
					t.Fatalf("unexpected compatibility gate response %#v", response)
				}
			},
		},
		{
			path: "/v1/reference-architecture/vald/final-gate?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValDCollectionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode final gate: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValDFinalGateStateActive || len(response.FamilyStates) != 6 {
					t.Fatalf("unexpected final gate response %#v", response)
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

func TestReferenceArchitectureValDProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/reference-architecture/vald/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response referenceArchitectureValDProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.ReferenceArchitectureValDStateActive {
		t.Fatalf("expected active proofs state, got %#v", response)
	}
	if response.Point5State != operability.IntelligenceCalibrationPoint5StatePass ||
		response.Point5DependencyState != operability.IntelligenceCalibrationValEStateActive ||
		response.Val0DependencyState != operability.ReferenceArchitectureVal0StateActive ||
		response.Val0State != operability.ReferenceArchitectureVal0StateActive ||
		response.ValADependencyState != operability.ReferenceArchitectureValAStateActive ||
		response.ValAState != operability.ReferenceArchitectureValAStateActive ||
		response.ValBDependencyState != operability.ReferenceArchitectureValBStateActive ||
		response.ValBState != operability.ReferenceArchitectureValBStateActive ||
		response.ValCDependencyState != operability.ReferenceArchitectureValCStateActive ||
		response.ValCState != operability.ReferenceArchitectureValCStateActive {
		t.Fatalf("expected active dependency states, got %#v", response)
	}
	if response.ValDState != operability.ReferenceArchitectureValDStateActive ||
		response.OperationalVisibilityState != operability.ReferenceArchitectureValDVisibilityStateActive ||
		response.AlignmentSummaryState != operability.ReferenceArchitectureValDAlignmentStateActive ||
		response.DeviationAlertState != operability.ReferenceArchitectureValDAlertStateActive ||
		response.SupportBoundaryState != operability.ReferenceArchitectureValDSupportBoundaryStateActive ||
		response.MigrationUpgradeState != operability.ReferenceArchitectureValDMigrationStateActive ||
		response.TopologyGateState != operability.ReferenceArchitectureValDTopologyGateStateActive ||
		response.SecurityBoundaryGateState != operability.ReferenceArchitectureValDSecurityGateStateActive ||
		response.OperabilityGateState != operability.ReferenceArchitectureValDOperabilityGateStateActive ||
		response.CompatibilityGateState != operability.ReferenceArchitectureValDCompatibilityGateStateActive ||
		response.FinalGateState != operability.ReferenceArchitectureValDFinalGateStateActive {
		t.Fatalf("expected active Val D component states, got %#v", response)
	}
	if response.Point6State != operability.ReferenceArchitecturePoint6StateNotComplete {
		t.Fatalf("expected point 6 to remain not complete, got %#v", response)
	}
	if len(response.SupportedFamilies) != 6 || len(response.FamilyStates) != 6 || len(response.WhyPoint6NotPass) == 0 || len(response.SurfaceRefs) < 15 || len(response.EvidenceRefs) < 14 || len(response.Limitations) == 0 {
		t.Fatalf("expected supported families, refs, and limitations in proofs, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection-only disclaimer, got %#v", response)
	}
}
