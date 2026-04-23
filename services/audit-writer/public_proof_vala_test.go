package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
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
	if !strings.Contains(pack.Pack.DownloadRef, "as_of="+url.QueryEscape("2026-04-22T10:00:00Z")) {
		t.Fatalf("expected pack download ref to preserve as_of, got %#v", pack.Pack)
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

func TestPublicProofValADownloadRefPreservesAsOfAndArtifactDeterminism(t *testing.T) {
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
	listed, ok := findValAPack(packs.Items, "point2_runtime_performance_public_pack")
	if !ok {
		t.Fatalf("expected runtime performance pack in list, got %#v", packs.Items)
	}
	if !strings.Contains(listed.DownloadRef, "as_of="+url.QueryEscape("2026-04-22T10:00:00Z")) {
		t.Fatalf("expected download ref to preserve as_of, got %#v", listed)
	}

	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, listed.DownloadRef, nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vala pack by download ref 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var pack publicProofValAPackResponse
	if err := json.NewDecoder(rec.Body).Decode(&pack); err != nil {
		t.Fatalf("decode vala pack by download ref: %v", err)
	}
	if pack.Pack.PayloadDigest != listed.PayloadDigest {
		t.Fatalf("expected identical payload digest from list and download, got list=%q download=%q", listed.PayloadDigest, pack.Pack.PayloadDigest)
	}
	if !reflect.DeepEqual(pack.Pack.MetricSummaries, listed.MetricSummaries) {
		t.Fatalf("expected identical metric summaries from list and download, got list=%#v download=%#v", listed.MetricSummaries, pack.Pack.MetricSummaries)
	}
}

func TestPublicProofValAPackByIDRequiresAsOf(t *testing.T) {
	handler := newHandlerWithRuntimesAndSigning(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "proof-vala-secret"),
	)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vala/downloadable-packs/point2_runtime_performance_public_pack", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected vala pack by id without as_of to fail, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestPublicProofValADifferentAsOfYieldDifferentRefsAndDigests(t *testing.T) {
	handler := newHandlerWithRuntimesAndSigning(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "proof-vala-secret"),
	)

	a := mustValAPacksAt(t, handler, "2026-04-22T10:00:00Z")
	b := mustValAPacksAt(t, handler, "2026-09-30T10:00:00Z")

	packA, ok := findValAPack(a.Items, "point2_runtime_performance_public_pack")
	if !ok {
		t.Fatalf("expected runtime performance pack in first list, got %#v", a.Items)
	}
	packB, ok := findValAPack(b.Items, "point2_runtime_performance_public_pack")
	if !ok {
		t.Fatalf("expected runtime performance pack in second list, got %#v", b.Items)
	}
	if packA.DownloadRef == packB.DownloadRef {
		t.Fatalf("expected as_of-specific download refs to differ, got %#v and %#v", packA.DownloadRef, packB.DownloadRef)
	}
	if packA.PayloadDigest == packB.PayloadDigest {
		t.Fatalf("expected as_of-dependent payload digest to differ, got %#v and %#v", packA.PayloadDigest, packB.PayloadDigest)
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

func mustValAPacksAt(t *testing.T, handler http.Handler, asOf string) publicProofValADownloadablePacksResponse {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/vala/downloadable-packs?as_of="+url.QueryEscape(asOf), nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vala downloadable packs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicProofValADownloadablePacksResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode vala downloadable packs: %v", err)
	}
	return response
}

func findValAPack(items []claimscore.PublicSealedProofPack, artifactID string) (claimscore.PublicSealedProofPack, bool) {
	for _, item := range items {
		if item.ArtifactID == artifactID {
			return item, true
		}
	}
	return claimscore.PublicSealedProofPack{}, false
}
