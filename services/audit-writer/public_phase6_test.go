package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	referencecore "github.com/denisgrosek/changelock/internal/reference"
)

func TestPublicTransparencyAnchorAndProofPortalHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	anchorReq := httptest.NewRequest(http.MethodGet, "/v1/public/transparency/anchor?as_of=2026-04-22T10:00:00Z", nil)
	anchorRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(anchorRec, anchorReq)
	if anchorRec.Code != http.StatusOK {
		t.Fatalf("expected transparency anchor 200, got %d: %s", anchorRec.Code, anchorRec.Body.String())
	}

	var anchor phase6TransparencyAnchorResponse
	if err := json.NewDecoder(anchorRec.Body).Decode(&anchor); err != nil {
		t.Fatalf("decode transparency anchor: %v", err)
	}
	if anchor.SchemaVersion != phase6TransparencyAnchorSchema || anchor.CurrentState != transparencyAnchorStateActive {
		t.Fatalf("expected active transparency anchor, got %#v", anchor)
	}
	if len(anchor.Artifacts) < 5 || anchor.RootHash == "" {
		t.Fatalf("expected hashed public artifacts in anchor, got %#v", anchor)
	}

	portalReq := httptest.NewRequest(http.MethodGet, "/v1/public/proof-portal?scope=public&as_of=2026-04-22T10:00:00Z", nil)
	portalRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(portalRec, portalReq)
	if portalRec.Code != http.StatusOK {
		t.Fatalf("expected proof portal 200, got %d: %s", portalRec.Code, portalRec.Body.String())
	}

	var portal phase6ProofPortalResponse
	if err := json.NewDecoder(portalRec.Body).Decode(&portal); err != nil {
		t.Fatalf("decode proof portal: %v", err)
	}
	if portal.SchemaVersion != phase6ProofPortalSchema || portal.CurrentState != phase6PortalStateActive {
		t.Fatalf("expected active proof portal, got %#v", portal)
	}
	if portal.FreshnessState != "fresh" || portal.AnchorRef == "" || len(portal.Items) < 5 {
		t.Fatalf("expected fresh proof portal with multiple items, got %#v", portal)
	}
}

func TestPublicTrustBadgeVerificationAndConformanceHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	badgeReq := httptest.NewRequest(http.MethodGet, "/v1/public/trust-program/badges/verify?badge_id=verification_ready&scope=public&as_of=2026-04-22T10:00:00Z", nil)
	badgeRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(badgeRec, badgeReq)
	if badgeRec.Code != http.StatusOK {
		t.Fatalf("expected trust badge verification 200, got %d: %s", badgeRec.Code, badgeRec.Body.String())
	}

	var badge phase6TrustBadgeVerificationResponse
	if err := json.NewDecoder(badgeRec.Body).Decode(&badge); err != nil {
		t.Fatalf("decode trust badge verification: %v", err)
	}
	if badge.CurrentState != "active" || badge.Verification.CurrentState != "verified" {
		t.Fatalf("expected active verified badge, got %#v", badge)
	}

	conformanceReq := httptest.NewRequest(http.MethodGet, "/v1/public/reference/conformance?architecture_id=runtime-hardened-enterprise-cluster&as_of=2026-04-22T10:00:00Z", nil)
	conformanceRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(conformanceRec, conformanceReq)
	if conformanceRec.Code != http.StatusOK {
		t.Fatalf("expected reference conformance 200, got %d: %s", conformanceRec.Code, conformanceRec.Body.String())
	}

	var conformance phase6ReferenceConformanceResponse
	if err := json.NewDecoder(conformanceRec.Body).Decode(&conformance); err != nil {
		t.Fatalf("decode reference conformance: %v", err)
	}
	if conformance.SchemaVersion != phase6ReferenceConformanceSchema || conformance.CurrentState == referencecore.StateIncomplete {
		t.Fatalf("expected bounded conformance response, got %#v", conformance)
	}
	if !containsString(conformance.UnsupportedChecks, "formal_certification") {
		t.Fatalf("expected unsupported certification boundary to remain visible, got %#v", conformance)
	}
	if !containsComparisonState(conformance.ComparisonItems, "formal_certification", "unsupported") {
		t.Fatalf("expected unsupported comparison item to remain visible, got %#v", conformance.ComparisonItems)
	}
}

func TestPublicVerifierSDKClaimsSummaryAndPhase6ProofsHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	sdkReq := httptest.NewRequest(http.MethodGet, "/v1/public/verifier/sdk?as_of=2026-04-22T10:00:00Z", nil)
	sdkRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(sdkRec, sdkReq)
	if sdkRec.Code != http.StatusOK {
		t.Fatalf("expected verifier sdk 200, got %d: %s", sdkRec.Code, sdkRec.Body.String())
	}

	var sdk phase6VerifierSDKResponse
	if err := json.NewDecoder(sdkRec.Body).Decode(&sdk); err != nil {
		t.Fatalf("decode verifier sdk: %v", err)
	}
	if !sdk.NoVendorBackend || sdk.CurrentState != phase6VerifierSDKStateActive {
		t.Fatalf("expected vendor-independent verifier sdk baseline, got %#v", sdk)
	}
	if len(sdk.ArtifactExports) < 5 || len(sdk.SupportedSchemaLines) == 0 {
		t.Fatalf("expected completed verifier sdk metadata, got %#v", sdk)
	}

	claimsReq := httptest.NewRequest(http.MethodGet, "/v1/public/claims/summary?scope=public&as_of=2026-04-22T10:00:00Z", nil)
	claimsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(claimsRec, claimsReq)
	if claimsRec.Code != http.StatusOK {
		t.Fatalf("expected claims summary 200, got %d: %s", claimsRec.Code, claimsRec.Body.String())
	}

	var claims phase6ClaimsSummaryResponse
	if err := json.NewDecoder(claimsRec.Body).Decode(&claims); err != nil {
		t.Fatalf("decode claims summary: %v", err)
	}
	if claims.CurrentState != phase6ClaimsStateActive || len(claims.Items) < 4 {
		t.Fatalf("expected active claims summary, got %#v", claims)
	}
	if claimPresent(claims.Items, "partner_exchange_claim") || claimPresent(claims.Items, "auditor_ready_claim") {
		t.Fatalf("expected public claims summary to stay scope-bounded, got %#v", claims.Items)
	}

	proofsReq := httptest.NewRequest(http.MethodGet, "/v1/public/phase6/proofs?as_of=2026-04-22T10:00:00Z", nil)
	proofsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(proofsRec, proofsReq)
	if proofsRec.Code != http.StatusOK {
		t.Fatalf("expected phase6 proofs 200, got %d: %s", proofsRec.Code, proofsRec.Body.String())
	}

	var proofs phase6ProofsResponse
	if err := json.NewDecoder(proofsRec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode phase6 proofs: %v", err)
	}
	if proofs.CurrentState != phase6ProofStateMarketActive {
		t.Fatalf("expected active phase6 proofs, got %#v", proofs)
	}
}

func TestPublicPhase6ProofsRemainInactiveWhenAnchorIsStale(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/phase6/proofs?as_of=2026-09-30T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase6 proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs phase6ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode stale phase6 proofs: %v", err)
	}
	if proofs.CurrentState == phase6ProofStateMarketActive {
		t.Fatalf("expected stale anchor to block active phase6 state, got %#v", proofs)
	}
	if proofs.TransparencyAnchor.CurrentState != transparencyAnchorStateStale {
		t.Fatalf("expected stale transparency anchor, got %#v", proofs.TransparencyAnchor)
	}
}

func TestPhase6TrustBadgeBecomesStaleWhenUnderlyingProofIsStale(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/trust-program/badges/verify?badge_id=verification_ready&scope=public&as_of=2026-09-30T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected trust badge verification 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var badge phase6TrustBadgeVerificationResponse
	if err := json.NewDecoder(rec.Body).Decode(&badge); err != nil {
		t.Fatalf("decode stale trust badge: %v", err)
	}
	if badge.CurrentState == "active" || badge.FreshnessState == "fresh" {
		t.Fatalf("expected stale underlying proof to stale the badge, got %#v", badge)
	}
}

func TestPhase6ClaimsAndPortalKeepScopesSeparated(t *testing.T) {
	fixture := forensicsTestFixture(t)

	publicClaims := mustPhase6ClaimsSummary(t, fixture.handler, phase6PublicScopePublic)
	if claimPresent(publicClaims.Items, "partner_exchange_claim") || claimPresent(publicClaims.Items, "auditor_ready_claim") {
		t.Fatalf("expected public scope to omit partner/auditor claims, got %#v", publicClaims.Items)
	}

	partnerClaims := mustPhase6ClaimsSummary(t, fixture.handler, phase6PublicScopePartner)
	if !claimPresent(partnerClaims.Items, "partner_exchange_claim") || claimPresent(partnerClaims.Items, "auditor_ready_claim") {
		t.Fatalf("expected partner scope to expose only partner claim, got %#v", partnerClaims.Items)
	}

	auditorClaims := mustPhase6ClaimsSummary(t, fixture.handler, phase6PublicScopeAuditor)
	if !claimPresent(auditorClaims.Items, "auditor_ready_claim") {
		t.Fatalf("expected auditor scope to expose auditor claim, got %#v", auditorClaims.Items)
	}

	portalReq := httptest.NewRequest(http.MethodGet, "/v1/public/proof-portal?scope=public&as_of=2026-04-22T10:00:00Z", nil)
	portalRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(portalRec, portalReq)
	var publicPortal phase6ProofPortalResponse
	if err := json.NewDecoder(portalRec.Body).Decode(&publicPortal); err != nil {
		t.Fatalf("decode public portal: %v", err)
	}
	for _, item := range publicPortal.Items {
		if item.Category == "auditor" || item.Category == "partner" {
			t.Fatalf("expected public portal to omit scoped-only items, got %#v", publicPortal.Items)
		}
	}
}

func TestPhase6FinalProofStateRequiresAllCoreSurfaces(t *testing.T) {
	if got := phase6FinalProofState(transparencyAnchorStateActive, phase6PortalStateActive, "active", phase6BenchmarkStateIncomplete, referencecore.StateActive, phase6VerifierSDKStateActive, phase6ClaimsStateActive, phase6AuditorWorkflowStateActive); got != phase6ProofStateIncomplete {
		t.Fatalf("expected incomplete without benchmarks, got %q", got)
	}
	if got := phase6FinalProofState(transparencyAnchorStateActive, phase6PortalStateActive, "active", phase6BenchmarkStateActive, referencecore.StatePartial, phase6VerifierSDKStateActive, phase6ClaimsStateActive, phase6AuditorWorkflowStateActive); got != phase6ProofStateSubstantial {
		t.Fatalf("expected substantial with partial conformance, got %q", got)
	}
	if got := phase6FinalProofState(transparencyAnchorStateActive, phase6PortalStateActive, "active", phase6BenchmarkStateActive, referencecore.StateActive, phase6VerifierSDKStateActive, phase6ClaimsStateActive, phase6AuditorWorkflowStateActive); got != phase6ProofStateMarketActive {
		t.Fatalf("expected active with all required surfaces, got %q", got)
	}
}

func mustPhase6ClaimsSummary(t *testing.T, handler http.Handler, scope string) phase6ClaimsSummaryResponse {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, "/v1/public/claims/summary?scope="+scope+"&as_of=2026-04-22T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected claims summary 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var response phase6ClaimsSummaryResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode claims summary: %v", err)
	}
	return response
}

func claimPresent(items []phase6ClaimSummaryItem, claimID string) bool {
	for _, item := range items {
		if item.ClaimID == claimID {
			return true
		}
	}
	return false
}

func containsComparisonState(items []phase6ReferenceComparisonItem, checkID, state string) bool {
	for _, item := range items {
		if item.CheckID == checkID && item.CurrentState == state {
			return true
		}
	}
	return false
}
