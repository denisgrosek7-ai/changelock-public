package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

func TestOSSTrustNetworkVal0Handlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/oss-trust-network/val0/status?tenant_id=acme&environment=prod",
		"/v1/oss-trust-network/val0/proofs?tenant_id=acme&environment=prod",
	}

	for _, path := range paths {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200 for %s, got %d: %s", path, rec.Code, rec.Body.String())
		}
	}

	req := httptest.NewRequest(http.MethodPost, "/v1/oss-trust-network/val0/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status route to remain read-only, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestOSSTrustNetworkVal0ProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/oss-trust-network/val0/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if strings.Contains(rec.Body.String(), "point_9_pass") {
		t.Fatalf("expected proofs response to omit forbidden final pass field, got %s", rec.Body.String())
	}

	var response ossTrustNetworkVal0ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.OSSTrustNetworkVal0StateActive ||
		response.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.DependencyState != operability.OSSTrustNetworkVal0DependencyStateActive {
		t.Fatalf("expected active OSTN proofs response with point 9 not complete, got %#v", response)
	}
	if response.Point8CurrentState != operability.DeveloperEcosystemValEStatePass ||
		response.Point8State != operability.DeveloperEcosystemPoint8StatePass ||
		!response.Point8PassAllowed ||
		response.Point8ClosureState != operability.DeveloperEcosystemValEClosureStateActive ||
		response.Point8NoOverclaimState != operability.DeveloperEcosystemValENoOverclaimStateActive {
		t.Fatalf("expected accepted Točka 8 / Val E dependency closure, got %#v", response)
	}
	if response.SignalContractState != operability.OSSTrustNetworkVal0SignalContractStateActive ||
		response.TrustMarkingState != operability.OSSTrustNetworkVal0TrustMarkingStateActive ||
		response.MaintainerIdentityState != operability.OSSTrustNetworkVal0MaintainerIdentityStateActive ||
		response.RegistryFreshnessState != operability.OSSTrustNetworkVal0RegistryFreshnessStateActive ||
		response.SharedVEXState != operability.OSSTrustNetworkVal0SharedVEXStateActive ||
		response.PropagationState != operability.OSSTrustNetworkVal0PropagationStateActive ||
		response.LocalApplicabilityState != operability.OSSTrustNetworkVal0LocalApplicabilityStateActive ||
		response.NoOverclaimState != operability.OSSTrustNetworkVal0NoOverclaimStateActive {
		t.Fatalf("expected active OSTN discipline states, got %#v", response)
	}
	if response.SignalReviewState != operability.OSSTrustNetworkReviewStateReviewed ||
		response.SharedVEXReviewState != operability.OSSTrustNetworkReviewStateReviewed ||
		response.RegistryFreshness != operability.IntelligenceCalibrationFreshnessFresh {
		t.Fatalf("expected canonical review and freshness metadata, got %#v", response)
	}
	if len(response.AllowedTrustMarkingClasses) != len(operability.OSSTrustNetworkVal0TrustMarkingModel().AllowedTrustMarkingClasses) ||
		len(response.SupportedReviewStates) != len(operability.OSSTrustNetworkVal0SignalContractModel().SupportedReviewStates) ||
		len(response.SurfaceRefs) != len(operability.OSSTrustNetworkVal0ProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.OSSTrustNetworkVal0ProofEvidenceRefs()) ||
		len(response.WhyPoint9NotComplete) == 0 ||
		len(response.Limitations) == 0 ||
		len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact proof/evidence refs and summary fields, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "oss_trust_network_val0") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

func TestOSSTrustNetworkVal0DocsAndStatusStayBounded(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/oss-trust-network/val0/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if strings.Contains(rec.Body.String(), "point_9_pass") {
		t.Fatalf("expected status response to omit forbidden final pass field, got %s", rec.Body.String())
	}

	var response ossTrustNetworkVal0StatusResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode status: %v", err)
	}
	if response.Model.CurrentState != operability.OSSTrustNetworkVal0StateActive ||
		response.Model.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete {
		t.Fatalf("expected active OSTN status model with point 9 not complete, got %#v", response)
	}

	files := []string{
		"../../docs/oss-trust-network-val0-core.md",
		"../../README.md",
		"../../public-docs/README.md",
	}
	for _, path := range files {
		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		lower := strings.ToLower(string(content))
		for _, blocked := range []string{
			"de-facto standard",
			"immune system for open source",
			"certified",
			"regulator-approved",
			"production approved",
			"audit passed",
			"point_9_pass",
		} {
			if strings.Contains(lower, blocked) {
				t.Fatalf("expected %s to avoid blocked wording %q", path, blocked)
			}
		}
	}

	readme, err := os.ReadFile("../../README.md")
	if err != nil {
		t.Fatalf("read repo readme: %v", err)
	}
	if !strings.Contains(string(readme), "Phase 9: Open-Source Trust Network Expansion") {
		t.Fatalf("expected repo readme Phase 9 section to reflect OSTN direction")
	}

	publicReadme, err := os.ReadFile("../../public-docs/README.md")
	if err != nil {
		t.Fatalf("read public readme: %v", err)
	}
	if !strings.Contains(string(publicReadme), "Phase 9: Open-Source Trust Network Expansion") {
		t.Fatalf("expected public readme Phase 9 section to reflect OSTN direction")
	}
}
