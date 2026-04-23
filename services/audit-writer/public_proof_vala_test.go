package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	claimscore "github.com/denisgrosek/changelock/internal/claims"
)

func TestPublicProofValAFoundationHandlers(t *testing.T) {
	handler := newHandlerWithRuntimesAndSigning(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "proof-vala-secret"),
	)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vala/sealed-artifact-schema", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vala artifact schema 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var schema publicProofValAArtifactSchemaResponse
	if err := json.NewDecoder(rec.Body).Decode(&schema); err != nil {
		t.Fatalf("decode vala artifact schema: %v", err)
	}
	if schema.CurrentState != claimscore.MeasuredPublicProofValAArtifactSchemaStateActive {
		t.Fatalf("expected active vala artifact schema, got %#v", schema)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vala/sealing-discipline", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vala sealing discipline 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var sealing publicProofValASealingDisciplineResponse
	if err := json.NewDecoder(rec.Body).Decode(&sealing); err != nil {
		t.Fatalf("decode vala sealing discipline: %v", err)
	}
	if sealing.CurrentState != claimscore.MeasuredPublicProofValASealingDisciplineStateActive {
		t.Fatalf("expected active vala sealing discipline, got %#v", sealing)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vala/environment-binding?as_of=2026-04-22T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vala environment binding 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var binding publicProofValAEnvironmentBindingResponse
	if err := json.NewDecoder(rec.Body).Decode(&binding); err != nil {
		t.Fatalf("decode vala environment binding: %v", err)
	}
	if binding.CurrentState != claimscore.MeasuredPublicProofValAEnvironmentBindingStateActive || len(binding.Items) != 2 {
		t.Fatalf("expected active vala environment binding with 2 items, got %#v", binding)
	}
}

func TestPublicProofValADownloadablePacksAndProofsHandlers(t *testing.T) {
	handler := newHandlerWithRuntimesAndSigning(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "proof-vala-secret"),
	)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vala/downloadable-packs?as_of=2026-04-22T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vala downloadable packs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var packs publicProofValADownloadablePacksResponse
	if err := json.NewDecoder(rec.Body).Decode(&packs); err != nil {
		t.Fatalf("decode vala downloadable packs: %v", err)
	}
	if packs.CurrentState != claimscore.MeasuredPublicProofValADownloadablePackStateActive {
		t.Fatalf("expected active vala downloadable packs, got %#v", packs)
	}
	if len(packs.Items) != 2 {
		t.Fatalf("expected 2 downloadable packs, got %#v", packs.Items)
	}
	if packs.Items[0].SignatureEnvelope == nil || packs.Items[0].SignatureEnvelope.Purpose != "public-proof-artifacts" {
		t.Fatalf("expected sealed pack signature envelope, got %#v", packs.Items[0])
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vala/downloadable-packs/point2_runtime_performance_public_pack?as_of=2026-04-22T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vala pack by id 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var pack publicProofValAPackResponse
	if err := json.NewDecoder(rec.Body).Decode(&pack); err != nil {
		t.Fatalf("decode vala pack by id: %v", err)
	}
	if pack.Pack.ArtifactID != "point2_runtime_performance_public_pack" || pack.CurrentState != "sealed_artifact_ready" {
		t.Fatalf("expected runtime performance pack, got %#v", pack)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vala/proofs?as_of=2026-04-22T10:00:00Z", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vala proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs publicProofValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode vala proofs: %v", err)
	}
	if proofs.CurrentState != claimscore.MeasuredPublicProofValAStateActive {
		t.Fatalf("expected active vala proofs, got %#v", proofs)
	}
	if proofs.Val0State != claimscore.MeasuredPublicProofVal0StateActive {
		t.Fatalf("expected active val0 dependency, got %#v", proofs)
	}
}

func TestPublicProofValAProofsStayInactiveWithoutPublicProofArtifactSigningPurpose(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vala/proofs?as_of=2026-04-22T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vala proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs publicProofValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode inactive vala proofs: %v", err)
	}
	if proofs.CurrentState == claimscore.MeasuredPublicProofValAStateActive {
		t.Fatalf("expected inactive vala proofs when signing purpose is unavailable, got %#v", proofs)
	}
	if proofs.SealingDisciplineState == claimscore.MeasuredPublicProofValASealingDisciplineStateActive {
		t.Fatalf("expected inactive sealing discipline when signing purpose is unavailable, got %#v", proofs)
	}
	if proofs.DownloadablePackState == claimscore.MeasuredPublicProofValADownloadablePackStateActive {
		t.Fatalf("expected inactive downloadable pack state when signing purpose is unavailable, got %#v", proofs)
	}
}
