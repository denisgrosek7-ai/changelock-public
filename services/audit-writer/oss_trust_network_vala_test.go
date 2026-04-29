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

func TestOSSTrustNetworkValAHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/oss-trust-network/vala/status?tenant_id=acme&environment=prod",
		"/v1/oss-trust-network/vala/proofs?tenant_id=acme&environment=prod",
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

	for _, path := range []string{
		"/v1/oss-trust-network/vala/status?tenant_id=acme&environment=prod",
		"/v1/oss-trust-network/vala/proofs?tenant_id=acme&environment=prod",
	} {
		req := httptest.NewRequest(http.MethodPost, path, nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Fatalf("expected read-only Val A route for %s, got %d: %s", path, rec.Code, rec.Body.String())
		}
	}
}

func TestOSSTrustNetworkValAProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/oss-trust-network/vala/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if strings.Contains(rec.Body.String(), "point_9_pass") {
		t.Fatalf("expected proofs response to omit forbidden final pass field, got %s", rec.Body.String())
	}

	var response ossTrustNetworkValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.OSSTrustNetworkValAStateActive ||
		response.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.DependencyState != operability.OSSTrustNetworkValADependencyStateActive {
		t.Fatalf("expected active OSTN Val A proofs response with point 9 not complete, got %#v", response)
	}
	if response.Val0CurrentState != operability.OSSTrustNetworkVal0StateActive ||
		response.Val0Point9State != operability.OSSTrustNetworkPoint9StateNotComplete ||
		response.Val0DependencyState != operability.OSSTrustNetworkVal0DependencyStateActive ||
		response.Val0NoOverclaimState != operability.OSSTrustNetworkVal0NoOverclaimStateActive ||
		response.Val0Point8State != operability.DeveloperEcosystemPoint8StatePass ||
		!response.Val0Point8PassAllowed ||
		response.Val0Point8PassReason != operability.DeveloperEcosystemValEPoint8PassReasonAllowed ||
		response.Val0Point8ClosureState != operability.DeveloperEcosystemValEClosureStateActive {
		t.Fatalf("expected exact Val 0 dependency closure in Val A proofs response, got %#v", response)
	}
	if response.ReleaseTrustIntakeState != operability.OSSTrustNetworkValAReleaseTrustIntakeStateActive ||
		response.SigningSignalState != operability.OSSTrustNetworkValASigningSignalStateActive ||
		response.MaintainerAttestationState != operability.OSSTrustNetworkValAMaintainerAttestationStateActive ||
		response.ProvenanceMaterialState != operability.OSSTrustNetworkValAProvenanceMaterialStateActive ||
		response.RegistryDescriptorState != operability.OSSTrustNetworkValARegistryDescriptorStateActive ||
		response.RegistryMetadataState != operability.OSSTrustNetworkValARegistryMetadataStateActive ||
		response.TypoSquattingWarningState != operability.OSSTrustNetworkValATypoSquattingWarningStateActive ||
		response.DriftSignalState != operability.OSSTrustNetworkValADriftSignalStateActive ||
		response.NoOverclaimState != operability.OSSTrustNetworkValANoOverclaimStateActive {
		t.Fatalf("expected active Val A discipline states, got %#v", response)
	}
	if response.ReleaseTrustReviewState != operability.OSSTrustNetworkReviewStateReviewed ||
		response.ReleaseTrustFreshness != operability.IntelligenceCalibrationFreshnessFresh ||
		response.SigningState != operability.OSSTrustNetworkValASigningStateVerified ||
		response.MaintainerAttestation != operability.OSSTrustNetworkValAAttestationStateAttested ||
		response.ProvenanceState != operability.OSSTrustNetworkValAProvenanceStateVerified ||
		response.RegistryDescriptor != operability.OSSTrustNetworkValARegistryDescriptorGitHubReleases ||
		response.RegistryMetadataFreshness != operability.IntelligenceCalibrationFreshnessFresh ||
		response.TypoSquattingReviewState != operability.OSSTrustNetworkReviewStateCandidate ||
		response.DriftClass != operability.OSSTrustNetworkValADriftClassSigning ||
		response.DriftState != operability.OSSTrustNetworkValADriftStateCandidate {
		t.Fatalf("expected canonical Val A signal metadata, got %#v", response)
	}
	if len(response.SupportedSigningStates) != 7 ||
		len(response.SupportedAttestationStates) != 7 ||
		len(response.SupportedProvenanceStates) != 7 ||
		len(response.SupportedRegistryDescriptors) != 6 ||
		len(response.SupportedDriftClasses) != 5 ||
		len(response.SurfaceRefs) != len(operability.OSSTrustNetworkValAProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.OSSTrustNetworkValAProofEvidenceRefs()) ||
		len(response.WhyPoint9NotComplete) == 0 ||
		len(response.Limitations) == 0 ||
		len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact proof/evidence refs and summary fields, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "oss_trust_network_vala") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

func TestOSSTrustNetworkValADocsAndStatusStayBounded(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/oss-trust-network/vala/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if strings.Contains(rec.Body.String(), "point_9_pass") {
		t.Fatalf("expected status response to omit forbidden final pass field, got %s", rec.Body.String())
	}

	var response ossTrustNetworkValAStatusResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode status: %v", err)
	}
	if response.Model.CurrentState != operability.OSSTrustNetworkValAStateActive ||
		response.Model.Point9State != operability.OSSTrustNetworkPoint9StateNotComplete {
		t.Fatalf("expected active OSTN Val A status model with point 9 not complete, got %#v", response)
	}

	files := []string{
		"../../docs/oss-trust-network-vala-core.md",
		"../../docs/architecture/phase-index.md",
	}
	for _, path := range files {
		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		lower := strings.ToLower(string(content))
		for _, blocked := range []string{
			"changelock verified",
			"certified package",
			"regulator-approved",
			"production approved",
			"deployment approved",
			"audit passed",
			"point_9_pass",
		} {
			if strings.Contains(lower, blocked) {
				t.Fatalf("expected %s to avoid blocked wording %q", path, blocked)
			}
		}
	}
}
