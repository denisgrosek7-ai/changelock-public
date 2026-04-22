package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPhase7MarketplaceReadinessHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/distribution/marketplace-readiness", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected marketplace readiness 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase7MarketplaceReadinessResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode marketplace readiness: %v", err)
	}
	if response.CurrentState != phase7MarketplaceReadinessStateActive {
		t.Fatalf("expected active marketplace readiness pack, got %#v", response)
	}
	if response.ProfileDetection != "environment_profile_detection_active" || response.ReadinessState != "marketplace_readiness_gate_active" {
		t.Fatalf("expected active marketplace profile and readiness states, got %#v", response)
	}
	if len(response.Checks) < 4 {
		t.Fatalf("expected bounded readiness checks, got %#v", response.Checks)
	}
	if !containsString(response.UnsupportedConditions, "unsupported_marketplace_profile_requires_local_completion") {
		t.Fatalf("expected unsupported marketplace condition to stay visible, got %#v", response.UnsupportedConditions)
	}
	if len(response.FailSafeRefs) == 0 || len(response.RolloutRefs) == 0 {
		t.Fatalf("expected fail-safe and rollout refs, got %#v", response)
	}
}

func TestPhase7MSPIsolationHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/distribution/msp-isolation", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected msp isolation 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase7MSPIsolationResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode msp isolation: %v", err)
	}
	if response.CurrentState != phase7MSPIsolationStateActive {
		t.Fatalf("expected active msp isolation pack, got %#v", response)
	}
	if response.TenantIsolation != "strict_tenant_isolation_verified" || response.AuditIsolation != "per_tenant_audit_isolation_verified" {
		t.Fatalf("expected strict tenant and audit isolation, got %#v", response)
	}
	if !containsString(response.AllowedDelegations, "tenant_safe_read_only_exports") {
		t.Fatalf("expected tenant-safe read-only exports, got %#v", response.AllowedDelegations)
	}
	if !containsString(response.ForbiddenDelegations, "cross_tenant_mutation") {
		t.Fatalf("expected cross-tenant mutation to stay forbidden, got %#v", response.ForbiddenDelegations)
	}
	if len(response.ExportBoundaryRefs) == 0 || len(response.AbuseControlRefs) == 0 {
		t.Fatalf("expected export boundary and abuse refs, got %#v", response)
	}
}

func TestPhase7PartnerExportHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/distribution/partner-export", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected partner export 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase7PartnerExportResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode partner export: %v", err)
	}
	if response.CurrentState != phase7PartnerExportStateActive {
		t.Fatalf("expected active partner export pack, got %#v", response)
	}
	if response.Scope != "partner" || !response.RedactedByDefault {
		t.Fatalf("expected partner-scoped redacted-by-default export, got %#v", response)
	}
	if !containsString(response.ForbiddenOperations, "broader_partner_write_api") {
		t.Fatalf("expected broader partner write API to stay forbidden, got %#v", response.ForbiddenOperations)
	}
	if !containsString(response.ExpandedScopeDeferred, "integrity_as_a_service_package") {
		t.Fatalf("expected integrity-as-a-service to remain deferred, got %#v", response.ExpandedScopeDeferred)
	}
	if publicExport := visibilityFields(response.VisibilityClasses, "public_exportable"); len(publicExport) != 0 {
		t.Fatalf("expected no public exportable fields in bounded partner export, got %#v", publicExport)
	}
}

func visibilityFields(items []phase7PartnerExportVisibility, visibilityClass string) []string {
	for _, item := range items {
		if item.VisibilityClass == visibilityClass {
			return item.Fields
		}
	}
	return nil
}
