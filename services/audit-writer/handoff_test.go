package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandoffSealDownloadAndVerify(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)

	sealReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/handoff/seal?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"audience":"auditor_safe","include_forensics":true,"co_sign_mode":"cosign_required"}`),
	)
	sealReq.Header.Set("Authorization", "Bearer operator-demo-token")
	sealReq.Header.Set("Content-Type", "application/json")
	sealRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(sealRec, sealReq)
	if sealRec.Code != http.StatusOK {
		t.Fatalf("expected handoff seal 200, got %d: %s", sealRec.Code, sealRec.Body.String())
	}

	var sealResponse handoffSealResponse
	if err := json.NewDecoder(sealRec.Body).Decode(&sealResponse); err != nil {
		t.Fatalf("decode handoff seal response: %v", err)
	}
	if sealResponse.PackageID == "" || sealResponse.Bundle.ManifestHash == "" {
		t.Fatalf("expected package and manifest identity, got %#v", sealResponse)
	}
	if sealResponse.SchemaVersion != handoffSealResponseSchemaVersion {
		t.Fatalf("expected schema-versioned handoff seal response, got %#v", sealResponse)
	}
	if sealResponse.Manifest.PackageType != handoffPackageTypeIncidentPackage {
		t.Fatalf("expected incident package handoff, got %#v", sealResponse.Manifest)
	}
	if sealResponse.Manifest.RedactionProfile.Audience != incidentAudienceAuditorSafe {
		t.Fatalf("expected redaction metadata to be frozen before sealing, got %#v", sealResponse.Manifest.RedactionProfile)
	}
	if sealResponse.Bundle.SealStatus != handoffSealStatusPendingCosign {
		t.Fatalf("expected pending cosign status, got %#v", sealResponse.Bundle)
	}
	if len(sealResponse.Manifest.Artifacts) < 4 {
		t.Fatalf("expected sealed bundle artifacts, got %#v", sealResponse.Manifest.Artifacts)
	}
	if len(sealResponse.Manifest.ReadbackRefs) == 0 || len(sealResponse.Manifest.ForensicRefs) == 0 {
		t.Fatalf("expected 9b and 9f lineage in manifest, got %#v", sealResponse.Manifest)
	}

	secondSealReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/handoff/seal?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"audience":"auditor_safe","include_forensics":true,"co_sign_mode":"cosign_required"}`),
	)
	secondSealReq.Header.Set("Authorization", "Bearer operator-demo-token")
	secondSealReq.Header.Set("Content-Type", "application/json")
	secondSealRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(secondSealRec, secondSealReq)
	if secondSealRec.Code != http.StatusOK {
		t.Fatalf("expected second handoff seal 200, got %d: %s", secondSealRec.Code, secondSealRec.Body.String())
	}
	var secondSeal handoffSealResponse
	if err := json.NewDecoder(secondSealRec.Body).Decode(&secondSeal); err != nil {
		t.Fatalf("decode second handoff seal response: %v", err)
	}
	if secondSeal.Bundle.ManifestHash != sealResponse.Bundle.ManifestHash {
		t.Fatalf("expected deterministic manifest hash for the same package scope, got %q then %q", sealResponse.Bundle.ManifestHash, secondSeal.Bundle.ManifestHash)
	}

	manifestReq := httptest.NewRequest(http.MethodGet, "/v1/handoff/"+sealResponse.PackageID+"/manifest", nil)
	manifestReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	manifestRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(manifestRec, manifestReq)
	if manifestRec.Code != http.StatusOK {
		t.Fatalf("expected handoff manifest 200, got %d: %s", manifestRec.Code, manifestRec.Body.String())
	}

	verifyReq := httptest.NewRequest(http.MethodGet, "/v1/handoff/"+sealResponse.PackageID+"/verification", nil)
	verifyReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	verifyRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(verifyRec, verifyReq)
	if verifyRec.Code != http.StatusOK {
		t.Fatalf("expected handoff verification 200, got %d: %s", verifyRec.Code, verifyRec.Body.String())
	}
	var verification verificationResult
	if err := json.NewDecoder(verifyRec.Body).Decode(&verification); err != nil {
		t.Fatalf("decode handoff verification: %v", err)
	}
	if !verification.ManifestValid || !verification.ArtifactHashesValid || !verification.SignaturesValid || !verification.TimestampValid || !verification.TransparencyValid {
		t.Fatalf("expected fully valid verification result, got %#v", verification)
	}
	if verification.SchemaVersion != handoffVerificationSchemaVersion {
		t.Fatalf("expected schema-versioned handoff verification, got %#v", verification)
	}

	downloadReq := httptest.NewRequest(http.MethodGet, "/v1/handoff/"+sealResponse.PackageID+"/download", nil)
	downloadReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	downloadRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(downloadRec, downloadReq)
	if downloadRec.Code != http.StatusOK {
		t.Fatalf("expected handoff download 200, got %d: %s", downloadRec.Code, downloadRec.Body.String())
	}

	record, err := parseHandoffBundle(downloadRec.Body.Bytes())
	if err != nil {
		t.Fatalf("parse handoff bundle: %v", err)
	}
	if record.PackageID != sealResponse.PackageID || record.ManifestHash != sealResponse.Bundle.ManifestHash {
		t.Fatalf("expected bundle to retain stored handoff identity, got %#v", record)
	}

	verifyBundleReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/handoff/verify",
		bytes.NewBufferString(`{"bundle_base64":"`+base64.StdEncoding.EncodeToString(downloadRec.Body.Bytes())+`"}`),
	)
	verifyBundleReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	verifyBundleReq.Header.Set("Content-Type", "application/json")
	verifyBundleRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(verifyBundleRec, verifyBundleReq)
	if verifyBundleRec.Code != http.StatusOK {
		t.Fatalf("expected handoff bundle verify 200, got %d: %s", verifyBundleRec.Code, verifyBundleRec.Body.String())
	}
	var bundleVerification verificationResult
	if err := json.NewDecoder(verifyBundleRec.Body).Decode(&bundleVerification); err != nil {
		t.Fatalf("decode bundle verification: %v", err)
	}
	if bundleVerification.OverallStatus != handoffVerificationValid {
		t.Fatalf("expected valid offline bundle verification, got %#v", bundleVerification)
	}
	if bundleVerification.SchemaVersion != handoffVerificationSchemaVersion {
		t.Fatalf("expected schema-versioned bundle verification, got %#v", bundleVerification)
	}
}

func TestHandoffCosignKeepsManifestHashBoundToSamePackage(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)

	sealReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/handoff/seal?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"audience":"internal","include_forensics":true,"co_sign_mode":"cosign_required"}`),
	)
	sealReq.Header.Set("Authorization", "Bearer operator-demo-token")
	sealReq.Header.Set("Content-Type", "application/json")
	sealRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(sealRec, sealReq)
	if sealRec.Code != http.StatusOK {
		t.Fatalf("expected handoff seal 200, got %d: %s", sealRec.Code, sealRec.Body.String())
	}
	var sealed handoffSealResponse
	if err := json.NewDecoder(sealRec.Body).Decode(&sealed); err != nil {
		t.Fatalf("decode sealed handoff: %v", err)
	}
	if sealed.Bundle.SealStatus != handoffSealStatusPendingCosign {
		t.Fatalf("expected pending cosign handoff, got %#v", sealed.Bundle)
	}

	cosignReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/handoff/"+sealed.PackageID+"/cosign",
		bytes.NewBufferString(`{"signer_role":"auditor"}`),
	)
	cosignReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	cosignReq.Header.Set("Content-Type", "application/json")
	cosignRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(cosignRec, cosignReq)
	if cosignRec.Code != http.StatusOK {
		t.Fatalf("expected handoff cosign 200, got %d: %s", cosignRec.Code, cosignRec.Body.String())
	}

	var cosigned handoffSealResponse
	if err := json.NewDecoder(cosignRec.Body).Decode(&cosigned); err != nil {
		t.Fatalf("decode cosigned handoff: %v", err)
	}
	if cosigned.Bundle.SealStatus != handoffSealStatusFullySealed || cosigned.Bundle.SignatureCount != 2 {
		t.Fatalf("expected fully sealed handoff with two signatures, got %#v", cosigned.Bundle)
	}

	recordReq := httptest.NewRequest(http.MethodGet, "/v1/handoff/"+sealed.PackageID, nil)
	recordReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	recordRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(recordRec, recordReq)
	if recordRec.Code != http.StatusOK {
		t.Fatalf("expected handoff record 200, got %d: %s", recordRec.Code, recordRec.Body.String())
	}
	var stored handoffSealResponse
	if err := json.NewDecoder(recordRec.Body).Decode(&stored); err != nil {
		t.Fatalf("decode stored handoff: %v", err)
	}
	if stored.Bundle.ManifestHash != sealed.Bundle.ManifestHash {
		t.Fatalf("expected cosign to keep the same manifest hash, got %q then %q", sealed.Bundle.ManifestHash, stored.Bundle.ManifestHash)
	}
	if stored.Session.InitiatedAt.After(time.Now().UTC().Add(5 * time.Second)) {
		t.Fatalf("expected stable handoff session timestamps, got %#v", stored.Session)
	}

	verifyReq := httptest.NewRequest(http.MethodGet, "/v1/handoff/"+sealed.PackageID+"/verification", nil)
	verifyReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	verifyRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(verifyRec, verifyReq)
	if verifyRec.Code != http.StatusOK {
		t.Fatalf("expected handoff verification 200, got %d: %s", verifyRec.Code, verifyRec.Body.String())
	}
	var verification verificationResult
	if err := json.NewDecoder(verifyRec.Body).Decode(&verification); err != nil {
		t.Fatalf("decode handoff verification: %v", err)
	}
	if verification.OverallStatus != handoffVerificationValid || len(verification.SignerIdentities) != 2 {
		t.Fatalf("expected valid cosigned handoff verification, got %#v", verification)
	}
}

func TestHandoffVerifyReportsTimestampTransparencyAndArtifactFailures(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)

	sealReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/handoff/seal?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"audience":"internal"}`),
	)
	sealReq.Header.Set("Authorization", "Bearer operator-demo-token")
	sealReq.Header.Set("Content-Type", "application/json")
	sealRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(sealRec, sealReq)
	if sealRec.Code != http.StatusOK {
		t.Fatalf("expected handoff seal 200, got %d: %s", sealRec.Code, sealRec.Body.String())
	}
	var sealed handoffSealResponse
	if err := json.NewDecoder(sealRec.Body).Decode(&sealed); err != nil {
		t.Fatalf("decode sealed handoff: %v", err)
	}

	downloadReq := httptest.NewRequest(http.MethodGet, "/v1/handoff/"+sealed.PackageID+"/download", nil)
	downloadReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	downloadRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(downloadRec, downloadReq)
	if downloadRec.Code != http.StatusOK {
		t.Fatalf("expected handoff download 200, got %d: %s", downloadRec.Code, downloadRec.Body.String())
	}

	record, err := parseHandoffBundle(downloadRec.Body.Bytes())
	if err != nil {
		t.Fatalf("parse handoff bundle: %v", err)
	}
	record.Timestamp.SubjectHash = "sha256:tampered-timestamp"
	record.Transparency.InclusionProof.RootHash = "sha256:tampered-transparency"
	if len(record.Files) == 0 {
		t.Fatalf("expected sealed bundle files, got %#v", record)
	}
	record.Files[0].Content += "\ncorrupted"

	tamperedBundle, err := buildHandoffBundle(record)
	if err != nil {
		t.Fatalf("build tampered handoff bundle: %v", err)
	}

	verifyReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/handoff/verify",
		bytes.NewBufferString(`{"bundle_base64":"`+base64.StdEncoding.EncodeToString(tamperedBundle)+`"}`),
	)
	verifyReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	verifyReq.Header.Set("Content-Type", "application/json")
	verifyRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(verifyRec, verifyReq)
	if verifyRec.Code != http.StatusOK {
		t.Fatalf("expected degraded handoff verify 200, got %d: %s", verifyRec.Code, verifyRec.Body.String())
	}

	var verification verificationResult
	if err := json.NewDecoder(verifyRec.Body).Decode(&verification); err != nil {
		t.Fatalf("decode degraded handoff verification: %v", err)
	}
	if verification.OverallStatus != handoffVerificationInvalid {
		t.Fatalf("expected invalid verification after timestamp/transparency/artifact corruption, got %#v", verification)
	}
	if verification.TimestampValid || verification.TransparencyValid || verification.ArtifactHashesValid {
		t.Fatalf("expected invalid timestamp, transparency, and artifact hashes, got %#v", verification)
	}
	limitations := strings.ToLower(strings.Join(verification.Limitations, " "))
	if !strings.Contains(limitations, "timestamp") || !strings.Contains(limitations, "transparency") || !strings.Contains(limitations, "artifact") {
		t.Fatalf("expected explicit degraded verification limitations, got %#v", verification.Limitations)
	}
}
