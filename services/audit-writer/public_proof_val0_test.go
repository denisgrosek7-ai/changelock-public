package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	claimscore "github.com/denisgrosek/changelock/internal/claims"
)

func TestPublicProofVal0FoundationHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/val0/claim-registry-model", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected val0 claim registry 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var registry publicProofVal0ClaimRegistryResponse
	if err := json.NewDecoder(rec.Body).Decode(&registry); err != nil {
		t.Fatalf("decode val0 claim registry: %v", err)
	}
	if registry.CurrentState != claimscore.MeasuredPublicProofVal0ClaimRegistryStateActive {
		t.Fatalf("expected active val0 claim registry, got %#v", registry)
	}
	if len(registry.Model.ClaimClasses) != 5 || len(registry.Model.LifecycleStates) < 8 {
		t.Fatalf("expected strict claim taxonomy and lifecycle states, got %#v", registry.Model)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/val0/redaction-tiers", nil)
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected val0 redaction tiers 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var redaction publicProofVal0RedactionTiersResponse
	if err := json.NewDecoder(rec.Body).Decode(&redaction); err != nil {
		t.Fatalf("decode val0 redaction tiers: %v", err)
	}
	if redaction.CurrentState != claimscore.MeasuredPublicProofVal0RedactionTierStateActive {
		t.Fatalf("expected active val0 redaction tiers, got %#v", redaction)
	}
	if len(redaction.Items) != 3 {
		t.Fatalf("expected 3 redaction tiers, got %#v", redaction.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/val0/signing-authority", nil)
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected val0 signing authority 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var signingAuthority publicProofVal0SigningAuthorityResponse
	if err := json.NewDecoder(rec.Body).Decode(&signingAuthority); err != nil {
		t.Fatalf("decode val0 signing authority: %v", err)
	}
	if signingAuthority.CurrentState != claimscore.MeasuredPublicProofVal0SigningAuthorityStateActive {
		t.Fatalf("expected active val0 signing authority, got %#v", signingAuthority)
	}
	if len(signingAuthority.Model.TrustRoots) == 0 || len(signingAuthority.Model.RequiredArtifactFields) == 0 {
		t.Fatalf("expected trust-root and artifact requirements, got %#v", signingAuthority.Model)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/val0/compatibility-baseline", nil)
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected val0 compatibility baseline 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var compatibility publicProofVal0CompatibilityResponse
	if err := json.NewDecoder(rec.Body).Decode(&compatibility); err != nil {
		t.Fatalf("decode val0 compatibility: %v", err)
	}
	if compatibility.CurrentState != claimscore.MeasuredPublicProofVal0CompatibilityStateActive {
		t.Fatalf("expected active val0 compatibility, got %#v", compatibility)
	}
	if len(compatibility.Model.ReplayTolerancePolicy) == 0 || len(compatibility.Model.FailureStates) == 0 {
		t.Fatalf("expected replay tolerance and failure states, got %#v", compatibility.Model)
	}
}

func TestPublicProofVal0ProofsHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/val0/proofs?as_of=2026-04-22T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected val0 proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicProofVal0ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode val0 proofs: %v", err)
	}
	if response.CurrentState != claimscore.MeasuredPublicProofVal0StateActive {
		t.Fatalf("expected active val0 proofs, got %#v", response)
	}
	if response.Phase6State != phase6ProofStateMarketActive {
		t.Fatalf("expected active phase6 dependency, got %#v", response)
	}
	if len(response.DeferredScope) == 0 || len(response.SurfaceRefs) < 5 || len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected deferred scope and summary refs, got %#v", response)
	}
}

func TestPublicProofVal0ProofsStayInactiveWhenPhase6IsNotActive(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/proof-expansion/val0/proofs?as_of=2026-09-30T10:00:00Z", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected val0 proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicProofVal0ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode stale val0 proofs: %v", err)
	}
	if response.CurrentState == claimscore.MeasuredPublicProofVal0StateActive {
		t.Fatalf("expected inactive val0 proofs when phase6 is stale, got %#v", response)
	}
	if response.Phase6State == phase6ProofStateMarketActive {
		t.Fatalf("expected stale phase6 dependency, got %#v", response)
	}
}
