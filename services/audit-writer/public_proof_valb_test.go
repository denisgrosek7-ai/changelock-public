package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	claimscore "github.com/denisgrosek/changelock/internal/claims"
)

func TestPublicProofValBFoundationHandlers(t *testing.T) {
	handler := newHandlerWithRuntimesAndSigning(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "proof-valb-secret"),
	)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/valb/transparency-chain?as_of=2026-04-22T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valb transparency chain 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var transparency publicProofValBTransparencyChainResponse
	if err := json.NewDecoder(rec.Body).Decode(&transparency); err != nil {
		t.Fatalf("decode valb transparency chain: %v", err)
	}
	if transparency.CurrentState != claimscore.MeasuredPublicProofValBTransparencyChainStateActive {
		t.Fatalf("expected active valb transparency chain, got %#v", transparency)
	}
	if len(transparency.Model.Entries) != 2 {
		t.Fatalf("expected 2 transparency entries, got %#v", transparency.Model.Entries)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/valb/verifier-capability?as_of=2026-04-22T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valb verifier capability 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var verifier publicProofValBVerifierCapabilityResponse
	if err := json.NewDecoder(rec.Body).Decode(&verifier); err != nil {
		t.Fatalf("decode valb verifier capability: %v", err)
	}
	if verifier.CurrentState != claimscore.MeasuredPublicProofValBVerifierCapabilityStateActive {
		t.Fatalf("expected active valb verifier capability, got %#v", verifier)
	}
	if verifier.Model.ReferencePackRef == "" || !containsString(verifier.Model.SupportedSchemaLines, "public.proof.sealed_artifact.v1") {
		t.Fatalf("expected verifier capability to expose reference pack and sealed artifact schema line, got %#v", verifier.Model)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/valb/signature-verification?as_of=2026-04-22T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valb signature verification 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var signature publicProofValBSignatureVerificationResponse
	if err := json.NewDecoder(rec.Body).Decode(&signature); err != nil {
		t.Fatalf("decode valb signature verification: %v", err)
	}
	if signature.CurrentState != claimscore.MeasuredPublicProofValBSignatureVerificationStateActive {
		t.Fatalf("expected active valb signature verification, got %#v", signature)
	}
	if len(signature.Items) != 2 {
		t.Fatalf("expected 2 signature verification items, got %#v", signature.Items)
	}
	for _, item := range signature.Items {
		if item.VerificationState != "verified" || item.TrustRootState != "trusted" || item.SchemaCompatibility != "supported" {
			t.Fatalf("expected fully verified signature item, got %#v", item)
		}
	}
}

func TestPublicProofValBReplayAndProofsHandlers(t *testing.T) {
	handler := newHandlerWithRuntimesAndSigning(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "proof-valb-secret"),
	)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/valb/replay-verification?as_of=2026-04-22T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valb replay verification 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var replay publicProofValBReplayVerificationResponse
	if err := json.NewDecoder(rec.Body).Decode(&replay); err != nil {
		t.Fatalf("decode valb replay verification: %v", err)
	}
	if replay.CurrentState != claimscore.MeasuredPublicProofValBReplayVerificationStateActive {
		t.Fatalf("expected active valb replay verification, got %#v", replay)
	}
	if len(replay.Items) != 2 || len(replay.BenchmarkEvaluations) == 0 {
		t.Fatalf("expected 2 replay items and benchmark evaluations, got %#v", replay)
	}
	if !hasValBReplayItem(replay.Items, "point2_runtime_performance_public_pack", "comparison_verified") {
		t.Fatalf("expected runtime performance comparison_verified replay item, got %#v", replay.Items)
	}
	if !hasValBReplayItem(replay.Items, "point2_verification_public_pack", "replay_verified") {
		t.Fatalf("expected verification replay_verified replay item, got %#v", replay.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/valb/proofs?as_of=2026-04-22T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valb proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs publicProofValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode valb proofs: %v", err)
	}
	if proofs.CurrentState != claimscore.MeasuredPublicProofValBStateActive {
		t.Fatalf("expected active valb proofs, got %#v", proofs)
	}
	if proofs.ValAState != claimscore.MeasuredPublicProofValAStateActive {
		t.Fatalf("expected active vala dependency for valb proofs, got %#v", proofs)
	}
	if proofs.TransparencyChainState != claimscore.MeasuredPublicProofValBTransparencyChainStateActive ||
		proofs.VerifierCapabilityState != claimscore.MeasuredPublicProofValBVerifierCapabilityStateActive ||
		proofs.SignatureVerificationState != claimscore.MeasuredPublicProofValBSignatureVerificationStateActive ||
		proofs.ReplayVerificationState != claimscore.MeasuredPublicProofValBReplayVerificationStateActive {
		t.Fatalf("expected all valb substates active, got %#v", proofs)
	}
}

func TestPublicProofValBProofsStayInactiveWithoutValAArtifactSigningPurpose(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/valb/proofs?as_of=2026-04-22T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valb proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs publicProofValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode inactive valb proofs: %v", err)
	}
	if proofs.CurrentState == claimscore.MeasuredPublicProofValBStateActive {
		t.Fatalf("expected inactive valb proofs when vala signing purpose is unavailable, got %#v", proofs)
	}
	if proofs.ValAState == claimscore.MeasuredPublicProofValAStateActive {
		t.Fatalf("expected inactive vala dependency when signing purpose is unavailable, got %#v", proofs)
	}
	if proofs.SignatureVerificationState == claimscore.MeasuredPublicProofValBSignatureVerificationStateActive {
		t.Fatalf("expected inactive signature verification when signing purpose is unavailable, got %#v", proofs)
	}
}

func hasValBReplayItem(items []claimscore.PublicProofReplayVerificationItem, artifactID, evaluationState string) bool {
	for _, item := range items {
		if item.ArtifactID == artifactID && item.EvaluationState == evaluationState {
			return true
		}
	}
	return false
}
