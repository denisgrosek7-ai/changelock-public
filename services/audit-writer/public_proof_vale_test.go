package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	claimscore "github.com/denisgrosek/changelock/internal/claims"
)

func TestPublicProofValEFoundationHandlers(t *testing.T) {
	handler := newHandlerWithRuntimesAndSigning(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "proof-vale-secret"),
	)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vale/replay-correctness-review?as_of=2026-04-23T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale replay correctness 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var replay publicProofValEReplayCorrectnessReviewResponse
	if err := json.NewDecoder(rec.Body).Decode(&replay); err != nil {
		t.Fatalf("decode vale replay correctness: %v", err)
	}
	if replay.CurrentState != claimscore.MeasuredPublicProofValEReplayCorrectnessReviewStateActive {
		t.Fatalf("expected active replay correctness review, got %#v", replay)
	}
	if len(replay.Items) != 2 {
		t.Fatalf("expected 2 replay review items, got %#v", replay.Items)
	}
	if !hasValEReplayItem(replay.Items, "point2_runtime_performance_public_pack", "within_declared_bands") {
		t.Fatalf("expected runtime replay correctness item, got %#v", replay.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vale/signing-trust-review?as_of=2026-04-23T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale signing trust 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var signingTrust publicProofValESigningTrustReviewResponse
	if err := json.NewDecoder(rec.Body).Decode(&signingTrust); err != nil {
		t.Fatalf("decode vale signing trust: %v", err)
	}
	if signingTrust.CurrentState != claimscore.MeasuredPublicProofValESigningTrustReviewStateActive {
		t.Fatalf("expected active signing trust review, got %#v", signingTrust)
	}
	if !hasValESigningTrustItem(signingTrust.Items, "point2_verification_public_pack", "trusted", "purpose_enabled") {
		t.Fatalf("expected verification signing trust item, got %#v", signingTrust.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vale/transparency-review?as_of=2026-04-23T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale transparency review 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var transparency publicProofValETransparencyReviewResponse
	if err := json.NewDecoder(rec.Body).Decode(&transparency); err != nil {
		t.Fatalf("decode vale transparency review: %v", err)
	}
	if transparency.CurrentState != claimscore.MeasuredPublicProofValETransparencyReviewStateActive {
		t.Fatalf("expected active transparency review, got %#v", transparency)
	}
	if !hasValETransparencyItem(transparency.Items, "point2_runtime_performance_public_pack", "visible_not_superseded") {
		t.Fatalf("expected runtime transparency review item, got %#v", transparency.Items)
	}
}

func TestPublicProofValERedactionIssuanceFailureAndProofsHandlers(t *testing.T) {
	handler := newHandlerWithRuntimesAndSigning(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "proof-vale-secret"),
	)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vale/redaction-review?as_of=2026-04-23T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale redaction review 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var redaction publicProofValERedactionReviewResponse
	if err := json.NewDecoder(rec.Body).Decode(&redaction); err != nil {
		t.Fatalf("decode vale redaction review: %v", err)
	}
	if redaction.CurrentState != claimscore.MeasuredPublicProofValERedactionReviewStateActive {
		t.Fatalf("expected active redaction review, got %#v", redaction)
	}
	if !hasValERedactionItem(redaction.Items, "point2_verification_public_pack", claimscore.RedactionTierPartnerScoped, claimscore.ScopePartner) {
		t.Fatalf("expected partner redaction review item, got %#v", redaction.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vale/compatibility-review?as_of=2026-04-23T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale compatibility review 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var compatibility publicProofValECompatibilityReviewResponse
	if err := json.NewDecoder(rec.Body).Decode(&compatibility); err != nil {
		t.Fatalf("decode vale compatibility review: %v", err)
	}
	if compatibility.CurrentState != claimscore.MeasuredPublicProofValECompatibilityReviewStateActive {
		t.Fatalf("expected active compatibility review, got %#v", compatibility)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vale/issuance-review?as_of=2026-04-23T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale issuance review 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var issuance publicProofValEIssuanceReviewResponse
	if err := json.NewDecoder(rec.Body).Decode(&issuance); err != nil {
		t.Fatalf("decode vale issuance review: %v", err)
	}
	if issuance.CurrentState != claimscore.MeasuredPublicProofValEIssuanceReviewStateActive {
		t.Fatalf("expected active issuance review, got %#v", issuance)
	}
	if !hasValEIssuanceItem(issuance.Items, "point2_verification_public_pack", claimscore.PublicProofStatusRestricted, "restricted_partner_scoped_reissue") {
		t.Fatalf("expected restricted verification issuance review item, got %#v", issuance.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vale/failure-state-review?as_of=2026-04-23T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale failure-state review 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var failureStates publicProofValEFailureStateReviewResponse
	if err := json.NewDecoder(rec.Body).Decode(&failureStates); err != nil {
		t.Fatalf("decode vale failure-state review: %v", err)
	}
	if failureStates.CurrentState != claimscore.MeasuredPublicProofValEFailureStateReviewStateActive {
		t.Fatalf("expected active failure-state review, got %#v", failureStates)
	}
	if !hasValEFailureStateItem(failureStates.Items, "point2_runtime_performance_public_pack", "claim_not_reissued_modeled") {
		t.Fatalf("expected runtime failure-state review item, got %#v", failureStates.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vale/proofs?as_of=2026-04-23T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs publicProofValEProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode vale proofs: %v", err)
	}
	if proofs.CurrentState != claimscore.MeasuredPublicProofValEStateActive {
		t.Fatalf("expected active vale proofs, got %#v", proofs)
	}
	if proofs.ValDState != claimscore.MeasuredPublicProofValDStateActive ||
		proofs.ReplayCorrectnessReviewState != claimscore.MeasuredPublicProofValEReplayCorrectnessReviewStateActive ||
		proofs.SigningTrustReviewState != claimscore.MeasuredPublicProofValESigningTrustReviewStateActive ||
		proofs.TransparencyReviewState != claimscore.MeasuredPublicProofValETransparencyReviewStateActive ||
		proofs.RedactionReviewState != claimscore.MeasuredPublicProofValERedactionReviewStateActive ||
		proofs.CompatibilityReviewState != claimscore.MeasuredPublicProofValECompatibilityReviewStateActive ||
		proofs.IssuanceReviewState != claimscore.MeasuredPublicProofValEIssuanceReviewStateActive ||
		proofs.FailureStateReviewState != claimscore.MeasuredPublicProofValEFailureStateReviewStateActive {
		t.Fatalf("expected all vale substates active, got %#v", proofs)
	}
	if len(proofs.DeferredScope) != 0 {
		t.Fatalf("expected no deferred scope in final proof gate, got %#v", proofs.DeferredScope)
	}
}

func TestPublicProofValEProofsStayInactiveWithoutValAArtifactSigningPurpose(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vale/proofs?as_of=2026-04-23T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs publicProofValEProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode inactive vale proofs: %v", err)
	}
	if proofs.CurrentState == claimscore.MeasuredPublicProofValEStateActive {
		t.Fatalf("expected inactive vale proofs without vala signing purpose, got %#v", proofs)
	}
	if proofs.ValDState == claimscore.MeasuredPublicProofValDStateActive {
		t.Fatalf("expected inactive vald dependency, got %#v", proofs)
	}
}

func hasValEReplayItem(items []claimscore.PublicProofReplayCorrectnessReviewItem, artifactID, toleranceDecision string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.ToleranceDecision == toleranceDecision {
			return true
		}
	}
	return false
}

func hasValESigningTrustItem(items []claimscore.PublicProofSigningTrustReviewItem, artifactID, trustRootState, signingPurposeState string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.TrustRootState == trustRootState && item.SigningPurposeState == signingPurposeState {
			return true
		}
	}
	return false
}

func hasValETransparencyItem(items []claimscore.PublicProofTransparencyReviewItem, artifactID, supersessionVisibility string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.SupersessionVisibility == supersessionVisibility {
			return true
		}
	}
	return false
}

func hasValERedactionItem(items []claimscore.PublicProofRedactionReviewItem, artifactID, redactionTier, publicationScope string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.RedactionTier == redactionTier && item.PublicationScope == publicationScope {
			return true
		}
	}
	return false
}

func hasValEIssuanceItem(items []claimscore.PublicProofIssuanceReviewItem, artifactID, lifecycleStatus, publicationDecision string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.ClaimLifecycleStatus == lifecycleStatus && item.PublicationDecision == publicationDecision {
			return true
		}
	}
	return false
}

func hasValEFailureStateItem(items []claimscore.PublicProofFailureStateReviewItem, artifactID, reissueFailureState string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.ReissueFailureState == reissueFailureState {
			return true
		}
	}
	return false
}
