package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	claimscore "github.com/denisgrosek/changelock/internal/claims"
)

func TestPublicProofValDFoundationHandlers(t *testing.T) {
	handler := newHandlerWithRuntimesAndSigning(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "proof-vald-secret"),
	)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vald/release-issuance-gate?as_of=2026-04-23T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vald release issuance 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var issuance publicProofValDReleaseIssuanceResponse
	if err := json.NewDecoder(rec.Body).Decode(&issuance); err != nil {
		t.Fatalf("decode vald release issuance: %v", err)
	}
	if issuance.CurrentState != claimscore.MeasuredPublicProofValDReleaseIssuanceStateActive {
		t.Fatalf("expected active release issuance, got %#v", issuance)
	}
	if len(issuance.Items) != 2 {
		t.Fatalf("expected 2 issuance items, got %#v", issuance.Items)
	}
	if !hasValDReleaseIssuanceItem(issuance.Items, "point2_verification_public_pack", "restricted_reissue_ready") {
		t.Fatalf("expected restricted verification reissue item, got %#v", issuance.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vald/claim-lifecycle?as_of=2026-04-23T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vald claim lifecycle 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var lifecycle publicProofValDClaimLifecycleResponse
	if err := json.NewDecoder(rec.Body).Decode(&lifecycle); err != nil {
		t.Fatalf("decode vald claim lifecycle: %v", err)
	}
	if lifecycle.CurrentState != claimscore.MeasuredPublicProofValDClaimLifecycleStateActive {
		t.Fatalf("expected active claim lifecycle, got %#v", lifecycle)
	}
	if !hasValDClaimLifecycleItem(lifecycle.Items, "point2_verification_public_pack", claimscore.PublicProofStatusRestricted, "restricted_to_partner_scope") {
		t.Fatalf("expected restricted partner lifecycle item, got %#v", lifecycle.Items)
	}
}

func TestPublicProofValDPublicationAndProofsHandlers(t *testing.T) {
	handler := newHandlerWithRuntimesAndSigning(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "proof-vald-secret"),
	)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vald/publication-decisions?as_of=2026-04-23T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vald publication decisions 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var publication publicProofValDPublicationDecisionResponse
	if err := json.NewDecoder(rec.Body).Decode(&publication); err != nil {
		t.Fatalf("decode vald publication decisions: %v", err)
	}
	if publication.CurrentState != claimscore.MeasuredPublicProofValDPublicationDecisionStateActive {
		t.Fatalf("expected active publication decisions, got %#v", publication)
	}
	if !hasValDPublicationItem(publication.Items, "point2_runtime_performance_public_pack", "approved_public_safe_reissue") {
		t.Fatalf("expected approved runtime publication item, got %#v", publication.Items)
	}
	if !hasValDPublicationItem(publication.Items, "point2_verification_public_pack", "restricted_partner_scoped_reissue") {
		t.Fatalf("expected restricted verification publication item, got %#v", publication.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vald/correction-workflow?as_of=2026-04-23T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vald correction workflow 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var correction publicProofValDCorrectionWorkflowResponse
	if err := json.NewDecoder(rec.Body).Decode(&correction); err != nil {
		t.Fatalf("decode vald correction workflow: %v", err)
	}
	if correction.CurrentState != claimscore.MeasuredPublicProofValDCorrectionWorkflowStateActive {
		t.Fatalf("expected active correction workflow, got %#v", correction)
	}
	if !hasValDCorrectionItem(correction.Items, "point2_runtime_performance_public_pack", "/v1/public/proof-expansion/valc/public-proof-portal") {
		t.Fatalf("expected runtime correction workflow item, got %#v", correction.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vald/proofs?as_of=2026-04-23T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vald proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs publicProofValDProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode vald proofs: %v", err)
	}
	if proofs.CurrentState != claimscore.MeasuredPublicProofValDStateActive {
		t.Fatalf("expected active vald proofs, got %#v", proofs)
	}
	if proofs.ValCState != claimscore.MeasuredPublicProofValCStateActive ||
		proofs.ReleaseIssuanceState != claimscore.MeasuredPublicProofValDReleaseIssuanceStateActive ||
		proofs.ClaimLifecycleState != claimscore.MeasuredPublicProofValDClaimLifecycleStateActive ||
		proofs.PublicationDecisionState != claimscore.MeasuredPublicProofValDPublicationDecisionStateActive ||
		proofs.CorrectionWorkflowState != claimscore.MeasuredPublicProofValDCorrectionWorkflowStateActive {
		t.Fatalf("expected all vald substates active, got %#v", proofs)
	}
}

func TestPublicProofValDProofsStayInactiveWithoutValAArtifactSigningPurpose(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vald/proofs?as_of=2026-04-23T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vald proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs publicProofValDProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode inactive vald proofs: %v", err)
	}
	if proofs.CurrentState == claimscore.MeasuredPublicProofValDStateActive {
		t.Fatalf("expected inactive vald proofs without vala signing purpose, got %#v", proofs)
	}
	if proofs.ValCState == claimscore.MeasuredPublicProofValCStateActive {
		t.Fatalf("expected inactive valc dependency, got %#v", proofs)
	}
	if proofs.ReleaseIssuanceState == claimscore.MeasuredPublicProofValDReleaseIssuanceStateActive {
		t.Fatalf("expected inactive release issuance when valc dependency is unavailable, got %#v", proofs)
	}
}

func hasValDReleaseIssuanceItem(items []claimscore.PublicProofReleaseIssuanceItem, artifactID, reissueDecision string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.ReissueDecision == reissueDecision {
			return true
		}
	}
	return false
}

func hasValDClaimLifecycleItem(items []claimscore.PublicProofClaimLifecycleItem, artifactID, claimStatus, restrictionState string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.ClaimStatus == claimStatus && item.RestrictionState == restrictionState {
			return true
		}
	}
	return false
}

func hasValDPublicationItem(items []claimscore.PublicProofPublicationDecisionItem, artifactID, publicationStatus string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.PublicationStatus == publicationStatus {
			return true
		}
	}
	return false
}

func hasValDCorrectionItem(items []claimscore.PublicProofCorrectionWorkflowItem, artifactID, correctionNoticeRef string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.CorrectionNoticeRef == correctionNoticeRef {
			return true
		}
	}
	return false
}
