package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	claimscore "github.com/denisgrosek/changelock/internal/claims"
)

func TestPublicProofValCPortalHandlers(t *testing.T) {
	handler := newHandlerWithRuntimesAndSigning(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "proof-valc-secret"),
	)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/valc/public-proof-portal?as_of=2026-04-22T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valc public portal 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var publicPortal publicProofValCPortalResponse
	if err := json.NewDecoder(rec.Body).Decode(&publicPortal); err != nil {
		t.Fatalf("decode valc public portal: %v", err)
	}
	if publicPortal.CurrentState != claimscore.MeasuredPublicProofValCPublicPortalStateActive {
		t.Fatalf("expected active valc public portal, got %#v", publicPortal)
	}
	if len(publicPortal.Items) != 1 {
		t.Fatalf("expected 1 public portal item, got %#v", publicPortal.Items)
	}
	if publicPortal.Items[0].ArtifactID != "point2_runtime_performance_public_pack" || publicPortal.Items[0].VisibilityState != claimscore.VisibilityPublicSafe {
		t.Fatalf("expected public-safe runtime performance portal item, got %#v", publicPortal.Items[0])
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/valc/partner-proof-portal?as_of=2026-04-22T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valc partner portal 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var partnerPortal publicProofValCPortalResponse
	if err := json.NewDecoder(rec.Body).Decode(&partnerPortal); err != nil {
		t.Fatalf("decode valc partner portal: %v", err)
	}
	if partnerPortal.CurrentState != claimscore.MeasuredPublicProofValCPartnerPortalStateActive {
		t.Fatalf("expected active valc partner portal, got %#v", partnerPortal)
	}
	if len(partnerPortal.Items) != 2 {
		t.Fatalf("expected 2 partner portal items, got %#v", partnerPortal.Items)
	}
	if !hasValCPortalItem(partnerPortal.Items, "point2_verification_public_pack", claimscore.VisibilityPartnerSafe) {
		t.Fatalf("expected partner-scoped verification portal item, got %#v", partnerPortal.Items)
	}
}

func TestPublicProofValCLineageAndProofsHandlers(t *testing.T) {
	handler := newHandlerWithRuntimesAndSigning(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "proof-valc-secret"),
	)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/valc/claim-lineage?as_of=2026-04-22T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valc claim lineage 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var lineage publicProofValCClaimLineageResponse
	if err := json.NewDecoder(rec.Body).Decode(&lineage); err != nil {
		t.Fatalf("decode valc claim lineage: %v", err)
	}
	if lineage.CurrentState != claimscore.MeasuredPublicProofValCClaimLineageStateActive {
		t.Fatalf("expected active valc claim lineage, got %#v", lineage)
	}
	if len(lineage.Items) != 2 {
		t.Fatalf("expected 2 lineage items, got %#v", lineage.Items)
	}
	if !hasValCLineageItem(lineage.Items, "point2_runtime_performance_public_pack", "not_superseded") {
		t.Fatalf("expected runtime pack lineage item, got %#v", lineage.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/valc/download-projections?as_of=2026-04-22T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valc download projections 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var downloads publicProofValCDownloadProjectionsResponse
	if err := json.NewDecoder(rec.Body).Decode(&downloads); err != nil {
		t.Fatalf("decode valc download projections: %v", err)
	}
	if downloads.CurrentState != claimscore.MeasuredPublicProofValCDownloadProjectionStateActive {
		t.Fatalf("expected active valc download projections, got %#v", downloads)
	}
	if len(downloads.Items) != 2 {
		t.Fatalf("expected 2 download projections, got %#v", downloads.Items)
	}
	if !hasValCDownloadProjection(downloads.Items, "point2_verification_public_pack", claimscore.ScopePartner, "reference_replay_available") {
		t.Fatalf("expected partner verification download projection, got %#v", downloads.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/valc/proofs?as_of=2026-04-22T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valc proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs publicProofValCProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode valc proofs: %v", err)
	}
	if proofs.CurrentState != claimscore.MeasuredPublicProofValCStateActive {
		t.Fatalf("expected active valc proofs, got %#v", proofs)
	}
	if proofs.ValBState != claimscore.MeasuredPublicProofValBStateActive ||
		proofs.PublicPortalState != claimscore.MeasuredPublicProofValCPublicPortalStateActive ||
		proofs.PartnerPortalState != claimscore.MeasuredPublicProofValCPartnerPortalStateActive ||
		proofs.ClaimLineageState != claimscore.MeasuredPublicProofValCClaimLineageStateActive ||
		proofs.DownloadProjectionState != claimscore.MeasuredPublicProofValCDownloadProjectionStateActive {
		t.Fatalf("expected all valc substates active, got %#v", proofs)
	}
}

func TestPublicProofValCProofsStayInactiveWithoutValAArtifactSigningPurpose(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/valc/proofs?as_of=2026-04-22T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valc proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs publicProofValCProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode inactive valc proofs: %v", err)
	}
	if proofs.CurrentState == claimscore.MeasuredPublicProofValCStateActive {
		t.Fatalf("expected inactive valc proofs without vala signing purpose, got %#v", proofs)
	}
	if proofs.ValBState == claimscore.MeasuredPublicProofValBStateActive {
		t.Fatalf("expected inactive valb dependency, got %#v", proofs)
	}
	if proofs.PublicPortalState == claimscore.MeasuredPublicProofValCPublicPortalStateActive {
		t.Fatalf("expected inactive public portal when valb dependency is unavailable, got %#v", proofs)
	}
}

func hasValCPortalItem(items []claimscore.PublicProofPortalProjectionItem, artifactID, visibility string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.VisibilityState == visibility {
			return true
		}
	}
	return false
}

func hasValCLineageItem(items []claimscore.PublicProofClaimLineageItem, artifactID, supersessionState string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.SupersessionState == supersessionState {
			return true
		}
	}
	return false
}

func hasValCDownloadProjection(items []claimscore.PublicProofDownloadProjectionItem, artifactID, scope, replayAvailability string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.PublicationScope == scope && item.ReplayAvailability == replayAvailability {
			return true
		}
	}
	return false
}
