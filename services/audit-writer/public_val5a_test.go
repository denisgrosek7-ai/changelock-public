package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPublicHandoffSpecHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/specs/handoff", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public handoff spec 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicHandoffSpecResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public handoff spec: %v", err)
	}
	if response.SchemaVersion != publicHandoffSpecSchema || response.FormatID != federationProofTypeHandoff {
		t.Fatalf("expected public handoff schema and format, got %#v", response)
	}
	if response.Compatibility.StabilityStatus != "stable" || len(response.ManifestFields) < 4 {
		t.Fatalf("expected stable handoff compatibility and manifest fields, got %#v", response)
	}
	if len(response.OfflineVerificationSteps) < 5 || !containsString(response.ArchiveIntegrityFields, "verify/public_keys.json") {
		t.Fatalf("expected offline handoff verification steps and archive integrity fields, got %#v", response)
	}
}

func TestPublicProofVerificationSpecHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/specs/proof-verification", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public proof verification spec 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicProofVerificationSpecResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public proof verification spec: %v", err)
	}
	if response.SchemaVersion != publicProofVerificationSpecSchema {
		t.Fatalf("expected proof verification schema, got %#v", response)
	}
	if len(response.LocalVerificationSteps) < 3 || len(response.RejectionReasons) < 5 {
		t.Fatalf("expected local verification steps and rejection reasons, got %#v", response)
	}
	if !containsString(response.LocalPolicyOverride, "Remote proof validity never bypasses local policy overrides, local distrust, or local disclosure exclusions.") {
		t.Fatalf("expected local-policy-first semantics, got %#v", response.LocalPolicyOverride)
	}
}

func TestPublicValidationCertificateSpecHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/specs/validation-certificate", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public validation certificate spec 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicValidationCertificateSpecResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public validation certificate spec: %v", err)
	}
	if response.SchemaVersion != publicValidationCertificateSchema {
		t.Fatalf("expected validation certificate schema, got %#v", response)
	}
	if len(response.CertificateFields) < 5 || !containsString(response.PassFailSemantics, validationStatusFlaky+" means recent runs did not converge on one stable outcome and the scenario requires caution.") {
		t.Fatalf("expected public validation certificate semantics, got %#v", response)
	}
	if len(response.AuthoritativeVsAdvisory) == 0 || len(response.FailureStates) < 3 {
		t.Fatalf("expected authoritative/advisory semantics and failure states, got %#v", response)
	}
}

func TestPublicFederationExchangeSpecHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/specs/federation-proof-exchange", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public federation exchange spec 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicFederationExchangeSpecResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public federation exchange spec: %v", err)
	}
	if response.SchemaVersion != publicFederationExchangeSpecSchema || !response.NoGlobalAuthority {
		t.Fatalf("expected public federation exchange schema and bounded authority flag, got %#v", response)
	}
	if len(response.DisclosureProfiles) < 3 || len(response.DivergenceDistrustModel) == 0 {
		t.Fatalf("expected disclosure profiles and divergence/distrust model, got %#v", response)
	}
}

func TestPublicVerifierProfilesHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/verifier/profiles", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public verifier profiles 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicVerifierProfilesResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public verifier profiles: %v", err)
	}
	if response.SchemaVersion != publicVerifierProfilesSchema || len(response.Profiles) < 4 {
		t.Fatalf("expected public verifier profile contract, got %#v", response)
	}
	if findPublicVerifierProfile(t, response.Profiles, "partner_verifier").DisplayName == "" {
		t.Fatalf("expected partner verifier profile, got %#v", response.Profiles)
	}
}

func TestPublicOfflineGuideAndExplainabilityHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	guideReq := httptest.NewRequest(http.MethodGet, "/v1/public/verifier/offline-guide", nil)
	guideRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(guideRec, guideReq)
	if guideRec.Code != http.StatusOK {
		t.Fatalf("expected public offline guide 200, got %d: %s", guideRec.Code, guideRec.Body.String())
	}

	var guide publicOfflineGuideResponse
	if err := json.NewDecoder(guideRec.Body).Decode(&guide); err != nil {
		t.Fatalf("decode public offline guide: %v", err)
	}
	if guide.SchemaVersion != publicOfflineGuideSchema || len(guide.VerificationSteps) < 6 {
		t.Fatalf("expected public offline guide with verification steps, got %#v", guide)
	}
	if !containsString(guide.ConformanceTargets, "full_verifier") {
		t.Fatalf("expected full verifier conformance target, got %#v", guide.ConformanceTargets)
	}

	explainReq := httptest.NewRequest(http.MethodGet, "/v1/public/specs/explainability-boundaries", nil)
	explainRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(explainRec, explainReq)
	if explainRec.Code != http.StatusOK {
		t.Fatalf("expected public explainability boundaries 200, got %d: %s", explainRec.Code, explainRec.Body.String())
	}

	var explain publicExplainabilityResponse
	if err := json.NewDecoder(explainRec.Body).Decode(&explain); err != nil {
		t.Fatalf("decode public explainability boundaries: %v", err)
	}
	if explain.SchemaVersion != publicExplainabilitySchema {
		t.Fatalf("expected public explainability schema, got %#v", explain)
	}
	if len(explain.PubliclyProves) == 0 || len(explain.LocalPolicyOwned) == 0 || len(explain.NotPubliclyClaimed) == 0 {
		t.Fatalf("expected public explainability boundaries, got %#v", explain)
	}
}

func TestPublicSampleHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	handoffReq := httptest.NewRequest(http.MethodGet, "/v1/public/samples/handoff", nil)
	handoffRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(handoffRec, handoffReq)
	if handoffRec.Code != http.StatusOK {
		t.Fatalf("expected public handoff sample 200, got %d: %s", handoffRec.Code, handoffRec.Body.String())
	}

	var handoff publicHandoffSampleResponse
	if err := json.NewDecoder(handoffRec.Body).Decode(&handoff); err != nil {
		t.Fatalf("decode public handoff sample: %v", err)
	}
	if handoff.SchemaVersion != publicHandoffSampleSchema || handoff.Sample.PackageID == "" {
		t.Fatalf("expected public handoff sample schema and package, got %#v", handoff)
	}
	if handoff.ExpectedVerifierOutcome.OverallStatus != handoffVerificationValid {
		t.Fatalf("expected valid verifier outcome for handoff sample, got %#v", handoff.ExpectedVerifierOutcome)
	}

	proofReq := httptest.NewRequest(http.MethodGet, "/v1/public/samples/proof-verification", nil)
	proofRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(proofRec, proofReq)
	if proofRec.Code != http.StatusOK {
		t.Fatalf("expected public proof sample 200, got %d: %s", proofRec.Code, proofRec.Body.String())
	}

	var proof publicProofVerificationSampleResponse
	if err := json.NewDecoder(proofRec.Body).Decode(&proof); err != nil {
		t.Fatalf("decode public proof sample: %v", err)
	}
	if proof.SchemaVersion != publicProofSampleSchema {
		t.Fatalf("expected public proof sample schema, got %#v", proof)
	}
	if proof.AcceptedExample.Decision.Decision != federationDecisionAccepted || proof.RejectedExample.Decision.Decision != federationDecisionRejectedStale {
		t.Fatalf("expected accepted and stale-rejected proof samples, got %#v", proof)
	}

	validationReq := httptest.NewRequest(http.MethodGet, "/v1/public/samples/validation-certificate", nil)
	validationRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(validationRec, validationReq)
	if validationRec.Code != http.StatusOK {
		t.Fatalf("expected public validation sample 200, got %d: %s", validationRec.Code, validationRec.Body.String())
	}

	var validation publicValidationCertificateSampleResponse
	if err := json.NewDecoder(validationRec.Body).Decode(&validation); err != nil {
		t.Fatalf("decode public validation sample: %v", err)
	}
	if validation.SchemaVersion != publicValidationSampleSchema || validation.Sample.OverallStatus != validationStatusPass {
		t.Fatalf("expected pass validation sample, got %#v", validation)
	}
	if !validation.Sample.SealReady || len(validation.Sample.ScenarioResults) < 2 {
		t.Fatalf("expected seal-ready validation sample with scenarios, got %#v", validation.Sample)
	}

	federationReq := httptest.NewRequest(http.MethodGet, "/v1/public/samples/federation-proof-exchange", nil)
	federationRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(federationRec, federationReq)
	if federationRec.Code != http.StatusOK {
		t.Fatalf("expected public federation sample 200, got %d: %s", federationRec.Code, federationRec.Body.String())
	}

	var federation publicFederationExchangeSampleResponse
	if err := json.NewDecoder(federationRec.Body).Decode(&federation); err != nil {
		t.Fatalf("decode public federation sample: %v", err)
	}
	if federation.SchemaVersion != publicFederationSampleSchema {
		t.Fatalf("expected public federation sample schema, got %#v", federation)
	}
	if federation.ReadyExample.PolicyState.SyncStatus != federationSyncStatusSynced || federation.StaleExample.Decision.Decision != federationDecisionRejectedStale || federation.DivergedExample.PolicyState.SyncStatus != federationSyncStatusDiverged {
		t.Fatalf("expected ready, stale, and diverged federation sample cases, got %#v", federation)
	}
}

func TestPublicConformancePackHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/conformance-pack", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public conformance pack 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicConformancePackResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public conformance pack: %v", err)
	}
	if response.SchemaVersion != publicConformancePackSchema || len(response.SampleRefs) < 4 || len(response.Assertions) < 5 {
		t.Fatalf("expected public conformance pack with sample refs and assertions, got %#v", response)
	}
	if findPublicConformanceAssertion(t, response.Assertions, "partner-proof-reject-stale").ExpectedState != federationDecisionRejectedStale {
		t.Fatalf("expected stale rejection assertion in conformance pack, got %#v", response.Assertions)
	}
}

func TestPublicSchemaExportHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	indexReq := httptest.NewRequest(http.MethodGet, "/v1/public/schemas", nil)
	indexRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(indexRec, indexReq)
	if indexRec.Code != http.StatusOK {
		t.Fatalf("expected public schema index 200, got %d: %s", indexRec.Code, indexRec.Body.String())
	}

	var index publicSchemaIndexResponse
	if err := json.NewDecoder(indexRec.Body).Decode(&index); err != nil {
		t.Fatalf("decode public schema index: %v", err)
	}
	if index.SchemaVersion != publicSchemaIndexSchema || len(index.Schemas) < 4 {
		t.Fatalf("expected public schema index with schema entries, got %#v", index)
	}

	proofReq := httptest.NewRequest(http.MethodGet, "/v1/public/schemas/proof-verification", nil)
	proofRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(proofRec, proofReq)
	if proofRec.Code != http.StatusOK {
		t.Fatalf("expected public proof schema export 200, got %d: %s", proofRec.Code, proofRec.Body.String())
	}

	var proof publicSchemaExportResponse
	if err := json.NewDecoder(proofRec.Body).Decode(&proof); err != nil {
		t.Fatalf("decode public proof schema export: %v", err)
	}
	if proof.SchemaVersion != publicSchemaExportSchema || proof.ExportID != "proof-verification" {
		t.Fatalf("expected public schema export response, got %#v", proof)
	}
	if !equalStrings(proof.FailureStates, []string{
		federationDecisionRejectedUnverifiable,
		federationDecisionRejectedScopeMismatch,
		federationDecisionRejectedPolicyConflict,
		federationDecisionRejectedStale,
		federationDecisionRejectedUntrustedPeer,
	}) {
		t.Fatalf("expected stable proof rejection states, got %#v", proof.FailureStates)
	}
	if !containsString(proof.RequiredFields, "freshness") || len(proof.FieldDefinitions) < 5 {
		t.Fatalf("expected machine-readable proof schema export, got %#v", proof)
	}

	handoffReq := httptest.NewRequest(http.MethodGet, "/v1/public/schemas/handoff", nil)
	handoffRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(handoffRec, handoffReq)
	if handoffRec.Code != http.StatusOK {
		t.Fatalf("expected public handoff schema export 200, got %d: %s", handoffRec.Code, handoffRec.Body.String())
	}

	var handoff publicSchemaExportResponse
	if err := json.NewDecoder(handoffRec.Body).Decode(&handoff); err != nil {
		t.Fatalf("decode public handoff schema export: %v", err)
	}
	if !containsString(handoff.RequiredFields, "root_hash") || !containsString(handoff.ConformanceRefs, "/v1/public/verifier/reference-pack") {
		t.Fatalf("expected handoff schema export with conformance ref, got %#v", handoff)
	}
}

func TestPublicVerifierReferencePackHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/public/verifier/reference-pack", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected public verifier reference pack 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response publicVerifierReferencePackResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode public verifier reference pack: %v", err)
	}
	if response.SchemaVersion != publicVerifierReferencePackSchema || len(response.ReplayInputs) < 5 {
		t.Fatalf("expected reference verifier pack replay inputs, got %#v", response)
	}

	handoffInput := findPublicVerifierReplayInput(t, response.ReplayInputs, "handoff_bundle_valid")
	bundleBytes, err := base64.StdEncoding.DecodeString(handoffInput.Payload)
	if err != nil {
		t.Fatalf("decode handoff replay bundle: %v", err)
	}
	record, err := parseHandoffBundle(bundleBytes)
	if err != nil {
		t.Fatalf("parse handoff replay bundle: %v", err)
	}
	if verification := (server{}).verifyStoredHandoff(record); verification.OverallStatus != handoffVerificationValid {
		t.Fatalf("expected replayable valid handoff bundle, got %#v", verification)
	}

	staleInput := findPublicVerifierReplayInput(t, response.ReplayInputs, "proof_rejected_stale_json")
	if staleInput.ExpectedState != federationDecisionRejectedStale || !strings.Contains(staleInput.Payload, federationDecisionRejectedStale) {
		t.Fatalf("expected stale proof replay input to preserve rejection state, got %#v", staleInput)
	}
}

func findPublicVerifierProfile(t *testing.T, items []publicVerifierProfile, profileID string) publicVerifierProfile {
	t.Helper()
	for _, item := range items {
		if item.ProfileID == profileID {
			return item
		}
	}
	t.Fatalf("expected verifier profile %q, got %#v", profileID, items)
	return publicVerifierProfile{}
}

func findPublicConformanceAssertion(t *testing.T, items []publicConformanceAssertion, assertionID string) publicConformanceAssertion {
	t.Helper()
	for _, item := range items {
		if item.AssertionID == assertionID {
			return item
		}
	}
	t.Fatalf("expected conformance assertion %q, got %#v", assertionID, items)
	return publicConformanceAssertion{}
}

func findPublicVerifierReplayInput(t *testing.T, items []publicVerifierReplayInput, inputID string) publicVerifierReplayInput {
	t.Helper()
	for _, item := range items {
		if item.InputID == inputID {
			return item
		}
	}
	t.Fatalf("expected verifier replay input %q, got %#v", inputID, items)
	return publicVerifierReplayInput{}
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
