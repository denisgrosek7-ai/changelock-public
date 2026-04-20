package main

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	handoffComponent = "handoff-manager"

	handoffPackageTypeIncidentPackage = "incident_package"
	handoffManifestSchemaVersion      = "9g.manifest.v1"
	handoffBundleSchemaVersion        = "9g.bundle.v1"
	handoffRedactionProfileVersion    = "8q.v1"

	handoffSealStatusDraft          = "draft"
	handoffSealStatusSealedBySystem = "sealed_by_system"
	handoffSealStatusPendingCosign  = "pending_cosign"
	handoffSealStatusFullySealed    = "fully_sealed"
	handoffSealStatusFailed         = "failed"

	handoffCoSignSystemOnly         = "system_only"
	handoffCoSignOptional           = "optional_cosign"
	handoffCoSignRequired           = "cosign_required"
	handoffSignatureAlgorithm       = "ed25519"
	handoffTimestampTypeRFC3161Like = "changelock_detached_timestamp"
	handoffTransparencyStatusLogged = "logged"

	handoffSignerRoleSystem            = "system"
	handoffSignerRoleCISO              = "ciso"
	handoffSignerRoleAuditor           = "auditor"
	handoffSignerRoleDelegatedApprover = "delegated_approver"

	handoffVerificationValid   = "valid"
	handoffVerificationInvalid = "invalid"
	handoffVerificationPartial = "partial"
)

var (
	errHandoffNotFound           = errors.New("handoff package not found")
	errHandoffInvalidScope       = errors.New("handoff scope is invalid")
	errHandoffSigningDisabled    = errors.New("handoff signing seed is not configured")
	errHandoffAlreadyCosigned    = errors.New("handoff already has a cosign for that signer role")
	errHandoffCosignNotAllowed   = errors.New("handoff does not allow cosign")
	errHandoffManifestMismatch   = errors.New("handoff manifest hash mismatch")
	errHandoffInvalidBundle      = errors.New("handoff bundle is invalid")
	errHandoffVerificationFailed = errors.New("handoff verification failed")
)

type handoffSealRequest struct {
	Audience               string   `json:"audience,omitempty"`
	IncidentIDs            []string `json:"incident_ids,omitempty"`
	IncludeForensics       bool     `json:"include_forensics,omitempty"`
	IncludeRuntime         bool     `json:"include_runtime,omitempty"`
	IncludeRecommendations bool     `json:"include_recommendations,omitempty"`
	CoSignMode             string   `json:"co_sign_mode,omitempty"`
}

type handoffCosignRequest struct {
	SignerRole string `json:"signer_role,omitempty"`
}

type handoffVerifyRequest struct {
	BundleBase64 string `json:"bundle_base64,omitempty"`
	PackageID    string `json:"package_id,omitempty"`
}

type handoffRuntimeContext struct {
	AdvisoryOnly bool                      `json:"advisory_only"`
	Workloads    []runtimeWorkloadView     `json:"workloads"`
	Findings     []runtimeIntegrityFinding `json:"findings"`
	Limitations  []string                  `json:"limitations,omitempty"`
}

type sealedManifest struct {
	PackageID         string                      `json:"package_id"`
	PackageType       string                      `json:"package_type"`
	SchemaVersion     string                      `json:"schema_version"`
	CreatedAt         time.Time                   `json:"created_at"`
	GeneratorIdentity string                      `json:"generator_identity"`
	Scope             sealedManifestScope         `json:"scope"`
	RedactionProfile  sealedManifestRedaction     `json:"redaction_profile"`
	Artifacts         []sealedManifestArtifact    `json:"artifacts"`
	EvidenceRefs      []string                    `json:"evidence_refs"`
	ReadbackRefs      []sealedManifestReadbackRef `json:"readback_refs,omitempty"`
	ForensicRefs      []sealedManifestForensicRef `json:"forensic_refs,omitempty"`
	RootHash          string                      `json:"root_hash"`
	Limitations       []string                    `json:"limitations,omitempty"`
}

type sealedManifestScope struct {
	Audience         string   `json:"audience"`
	SelectionMode    string   `json:"selection_mode"`
	SelectionSummary string   `json:"selection_summary"`
	IncidentCount    int      `json:"incident_count"`
	IncidentRefs     []string `json:"incident_refs"`
	TenantID         string   `json:"tenant_id,omitempty"`
	Environment      string   `json:"environment,omitempty"`
	Repo             string   `json:"repo,omitempty"`
}

type sealedManifestRedaction struct {
	Audience       string   `json:"audience"`
	ProfileVersion string   `json:"profile_version"`
	Summary        []string `json:"summary"`
}

type sealedManifestArtifact struct {
	Path         string `json:"path"`
	MediaType    string `json:"media_type"`
	SHA256       string `json:"sha256"`
	Role         string `json:"role"`
	AdvisoryOnly bool   `json:"advisory_only,omitempty"`
}

type sealedManifestReadbackRef struct {
	ResourceType string `json:"resource_type"`
	ResourceID   string `json:"resource_id,omitempty"`
	EvidenceHash string `json:"evidence_hash"`
	ResourceURI  string `json:"resource_uri,omitempty"`
}

type sealedManifestForensicRef struct {
	ContextURI     string `json:"context_uri,omitempty"`
	ContextType    string `json:"context_type"`
	Timestamp      string `json:"timestamp"`
	AdvisoryOnly   bool   `json:"advisory_only"`
	Counterfactual bool   `json:"counterfactual,omitempty"`
}

type signatureRecord struct {
	SignatureID          string    `json:"signature_id"`
	SignerRole           string    `json:"signer_role"`
	SignerIdentity       string    `json:"signer_identity"`
	Algorithm            string    `json:"algorithm"`
	SignedObject         string    `json:"signed_object"`
	SignedHash           string    `json:"signed_hash"`
	SignatureValue       string    `json:"signature_value"`
	PublicKey            string    `json:"public_key"`
	CertificateChainRefs []string  `json:"certificate_chain_refs,omitempty"`
	SignedAt             time.Time `json:"signed_at"`
}

type timestampRecord struct {
	TimestampType  string    `json:"timestamp_type"`
	Authority      string    `json:"authority"`
	SubjectHash    string    `json:"subject_hash"`
	TokenRef       string    `json:"token_ref"`
	TimestampedAt  time.Time `json:"timestamped_at"`
	Algorithm      string    `json:"algorithm"`
	SignatureValue string    `json:"signature_value"`
	SignerIdentity string    `json:"signer_identity"`
	PublicKey      string    `json:"public_key"`
}

type transparencyInclusionProof struct {
	TreeSize  int      `json:"tree_size"`
	LeafIndex int      `json:"leaf_index"`
	LeafHash  string   `json:"leaf_hash"`
	RootHash  string   `json:"root_hash"`
	Path      []string `json:"path"`
}

type transparencyRecord struct {
	LogID          string                     `json:"log_id"`
	EntryID        string                     `json:"entry_id"`
	SubjectHash    string                     `json:"subject_hash"`
	InclusionProof transparencyInclusionProof `json:"inclusion_proof"`
	LoggedAt       time.Time                  `json:"logged_at"`
	Status         string                     `json:"status"`
	Algorithm      string                     `json:"algorithm"`
	SignatureValue string                     `json:"signature_value"`
	SignerIdentity string                     `json:"signer_identity"`
	PublicKey      string                     `json:"public_key"`
}

type sealedBundleMetadata struct {
	PackageID              string `json:"package_id"`
	BundlePath             string `json:"bundle_path"`
	ManifestHash           string `json:"manifest_hash"`
	SealStatus             string `json:"seal_status"`
	SignatureCount         int    `json:"signature_count"`
	TimestampStatus        string `json:"timestamp_status"`
	TransparencyStatus     string `json:"transparency_status"`
	VerificationURI        string `json:"verification_uri"`
	OfflineVerifierPresent bool   `json:"offline_verifier_present"`
}

type verificationResult struct {
	PackageID           string   `json:"package_id"`
	ManifestValid       bool     `json:"manifest_valid"`
	ArtifactHashesValid bool     `json:"artifact_hashes_valid"`
	SignaturesValid     bool     `json:"signatures_valid"`
	TimestampValid      bool     `json:"timestamp_valid"`
	TransparencyValid   bool     `json:"transparency_valid"`
	SignerIdentities    []string `json:"signer_identities"`
	RedactionProfile    string   `json:"redaction_profile"`
	OverallStatus       string   `json:"overall_status"`
	Limitations         []string `json:"limitations,omitempty"`
}

type handoffSessionRecord struct {
	SessionID      string    `json:"session_id"`
	PackageID      string    `json:"package_id"`
	PackageType    string    `json:"package_type"`
	ScopeSummary   string    `json:"scope_summary"`
	InitiatedBy    string    `json:"initiated_by"`
	InitiatedAt    time.Time `json:"initiated_at"`
	SignMode       string    `json:"sign_mode"`
	CoSignMode     string    `json:"co_sign_mode"`
	Status         string    `json:"status"`
	FinalBundleRef string    `json:"final_bundle_ref"`
	ManifestHash   string    `json:"manifest_hash"`
}

type handoffBundleFile struct {
	Path         string `json:"path"`
	MediaType    string `json:"media_type"`
	Role         string `json:"role"`
	Content      string `json:"content"`
	AdvisoryOnly bool   `json:"advisory_only,omitempty"`
}

type handoffStoredRecord struct {
	PackageID       string               `json:"package_id"`
	PackageType     string               `json:"package_type"`
	Manifest        sealedManifest       `json:"manifest"`
	ManifestHash    string               `json:"manifest_hash"`
	Session         handoffSessionRecord `json:"session"`
	Bundle          sealedBundleMetadata `json:"bundle"`
	Files           []handoffBundleFile  `json:"files"`
	Signatures      []signatureRecord    `json:"signatures"`
	Timestamp       timestampRecord      `json:"timestamp"`
	Transparency    transparencyRecord   `json:"transparency"`
	Verification    verificationResult   `json:"verification"`
	DownloadURI     string               `json:"download_uri"`
	VerificationURI string               `json:"verification_uri"`
}

type handoffSealResponse struct {
	PackageID       string               `json:"package_id"`
	Manifest        sealedManifest       `json:"manifest"`
	Session         handoffSessionRecord `json:"session"`
	Bundle          sealedBundleMetadata `json:"bundle"`
	Verification    verificationResult   `json:"verification"`
	DownloadURI     string               `json:"download_uri"`
	VerificationURI string               `json:"verification_uri"`
}

type handoffReadbackLink struct {
	ResourceType string `json:"resource_type"`
	ResourceID   string `json:"resource_id,omitempty"`
	EvidenceHash string `json:"evidence_hash"`
	ResourceURI  string `json:"resource_uri,omitempty"`
}

func (s server) handoffSealHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request handoffSealRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	record, err := s.createSealedHandoff(ctx, principal, r, request)
	if err != nil {
		writeHandoffError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, handoffSealResponse{
		PackageID:       record.PackageID,
		Manifest:        record.Manifest,
		Session:         record.Session,
		Bundle:          record.Bundle,
		Verification:    record.Verification,
		DownloadURI:     record.DownloadURI,
		VerificationURI: record.VerificationURI,
	})
}

func (s server) handoffByIDHandler(w http.ResponseWriter, r *http.Request) {
	pathValue := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/handoff/"))
	if pathValue == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "handoff package not found"})
		return
	}
	parts := strings.Split(pathValue, "/")
	if len(parts) > 2 {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "handoff package not found"})
		return
	}
	packageID := strings.TrimSpace(parts[0])
	action := ""
	if len(parts) == 2 {
		action = strings.TrimSpace(parts[1])
	}
	if packageID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "handoff package not found"})
		return
	}
	switch {
	case action == "" && r.Method == http.MethodGet:
		s.getHandoffHandler(w, r, packageID)
	case action == "manifest" && r.Method == http.MethodGet:
		s.getHandoffManifestHandler(w, r, packageID)
	case action == "verification" && r.Method == http.MethodGet:
		s.getHandoffVerificationHandler(w, r, packageID)
	case action == "download" && r.Method == http.MethodGet:
		s.downloadHandoffHandler(w, r, packageID)
	case action == "cosign" && r.Method == http.MethodPost:
		s.cosignHandoffHandler(w, r, packageID)
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) handoffVerifyHandler(w http.ResponseWriter, r *http.Request) {
	if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
		return
	}
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request handoffVerifyRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	if strings.TrimSpace(request.PackageID) != "" {
		record, err := s.getStoredHandoffRecord(ctx, strings.TrimSpace(request.PackageID))
		if err != nil {
			writeHandoffError(w, err)
			return
		}
		httpjson.Write(w, http.StatusOK, s.verifyStoredHandoff(record))
		return
	}
	if strings.TrimSpace(request.BundleBase64) == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "bundle_base64 or package_id is required"})
		return
	}
	bundleBytes, err := base64.StdEncoding.DecodeString(strings.TrimSpace(request.BundleBase64))
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "bundle_base64 must be valid base64"})
		return
	}
	record, err := parseHandoffBundle(bundleBytes)
	if err != nil {
		writeHandoffError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, s.verifyStoredHandoff(record))
}

func (s server) getHandoffHandler(w http.ResponseWriter, r *http.Request, packageID string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	record, err := s.getStoredHandoffRecord(ctx, packageID)
	if err != nil {
		writeHandoffError(w, err)
		return
	}
	if err := ensurePrincipalHandoffScope(principal, record); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, handoffSealResponse{
		PackageID:       record.PackageID,
		Manifest:        record.Manifest,
		Session:         record.Session,
		Bundle:          record.Bundle,
		Verification:    s.verifyStoredHandoff(record),
		DownloadURI:     record.DownloadURI,
		VerificationURI: record.VerificationURI,
	})
}

func (s server) getHandoffManifestHandler(w http.ResponseWriter, r *http.Request, packageID string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	record, err := s.getStoredHandoffRecord(ctx, packageID)
	if err != nil {
		writeHandoffError(w, err)
		return
	}
	if err := ensurePrincipalHandoffScope(principal, record); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, record.Manifest)
}

func (s server) getHandoffVerificationHandler(w http.ResponseWriter, r *http.Request, packageID string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	record, err := s.getStoredHandoffRecord(ctx, packageID)
	if err != nil {
		writeHandoffError(w, err)
		return
	}
	if err := ensurePrincipalHandoffScope(principal, record); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, s.verifyStoredHandoff(record))
}

func (s server) downloadHandoffHandler(w http.ResponseWriter, r *http.Request, packageID string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	record, err := s.getStoredHandoffRecord(ctx, packageID)
	if err != nil {
		writeHandoffError(w, err)
		return
	}
	if err := ensurePrincipalHandoffScope(principal, record); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	bundleBytes, err := buildHandoffBundle(record)
	if err != nil {
		writeHandoffError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/vnd.changelock.safepkg+zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.safepkg\"", strings.ToLower(record.PackageID)))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bundleBytes)
}

func (s server) cosignHandoffHandler(w http.ResponseWriter, r *http.Request, packageID string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var request handoffCosignRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	record, err := s.getStoredHandoffRecord(ctx, packageID)
	if err != nil {
		writeHandoffError(w, err)
		return
	}
	if err := ensurePrincipalHandoffScope(principal, record); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	updated, err := s.cosignStoredHandoff(ctx, principal, record, request)
	if err != nil {
		writeHandoffError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, handoffSealResponse{
		PackageID:       updated.PackageID,
		Manifest:        updated.Manifest,
		Session:         updated.Session,
		Bundle:          updated.Bundle,
		Verification:    updated.Verification,
		DownloadURI:     updated.DownloadURI,
		VerificationURI: updated.VerificationURI,
	})
}

func (s server) createSealedHandoff(ctx context.Context, principal auth.Principal, r *http.Request, request handoffSealRequest) (handoffStoredRecord, error) {
	filter, err := parseIncidentFilter(r)
	if err != nil {
		return handoffStoredRecord{}, err
	}
	incidents, err := s.listIncidents(ctx, filter)
	if err != nil {
		return handoffStoredRecord{}, err
	}
	incidents = sortHandoffIncidents(incidents)
	selectedIDs := uniqueStrings(append(parseIncidentIDList(r), request.IncidentIDs...))
	audience, err := parseIncidentExportAudience(firstNonEmpty(request.Audience, r.URL.Query().Get("audience")))
	if err != nil {
		return handoffStoredRecord{}, err
	}
	filtered, selectionMode := selectIncidentScope(incidents, selectedIDs)
	packagePayload := buildIncidentPackage(incidents, selectedIDs, filter, audience)
	manifestTime := deterministicManifestTimestamp(filtered, filter)
	packagePayload.GeneratedAt = manifestTime
	packagePayload.PackageIntel.GeneratedAt = manifestTime
	packagePayload.RedactionSummary = incidentExportRedactionSummary(audience)

	readbackRefs := s.collectPackageReadbackRefs(filtered, filter, audience)
	var forensicState *pointInTimeState
	var forensicRefs []sealedManifestForensicRef
	if request.IncludeForensics {
		state, err := s.buildPointInTimeState(ctx, forensicsFilter{
			Timestamp: manifestTime,
			event: audit.EventFilter{
				Decision:    filter.event.Decision,
				EventType:   filter.event.EventType,
				Component:   filter.event.Component,
				ClusterID:   filter.event.ClusterID,
				Repo:        filter.event.Repo,
				Environment: filter.event.Environment,
				TenantID:    filter.event.TenantID,
				Since:       filter.event.Since,
				Until:       filter.event.Until,
				Limit:       filter.event.Limit,
			},
		})
		if err != nil {
			return handoffStoredRecord{}, err
		}
		forensicState = &state
		forensicRefs = []sealedManifestForensicRef{{
			ContextURI:   fmt.Sprintf("/v1/forensics/state?tenant_id=%s&environment=%s&timestamp=%s", urlQueryEscape(filter.event.TenantID), urlQueryEscape(filter.event.Environment), manifestTime.Format(time.RFC3339)),
			ContextType:  state.Mode,
			Timestamp:    state.Timestamp.Format(time.RFC3339),
			AdvisoryOnly: true,
		}}
	}
	var runtimeContext *handoffRuntimeContext
	if request.IncludeRuntime {
		runtimeFilter := runtimeIntegrityFilter{
			ClusterID:   filter.event.ClusterID,
			TenantID:    filter.event.TenantID,
			Environment: filter.event.Environment,
			Repo:        filter.event.Repo,
			Limit:       6,
			event: audit.EventFilter{
				ClusterID:   filter.event.ClusterID,
				TenantID:    filter.event.TenantID,
				Environment: filter.event.Environment,
				Repo:        filter.event.Repo,
				Limit:       600,
			},
		}
		if len(filtered) == 1 {
			runtimeFilter.Workload = firstNonEmpty(filtered[0].ScopeRef, firstString(filtered[0].AffectedWorkloads))
		}
		workloads, workloadLimitations, err := s.buildRuntimeWorkloads(ctx, runtimeFilter)
		if err != nil {
			return handoffStoredRecord{}, err
		}
		findings, findingLimitations, err := s.buildRuntimeFindings(ctx, runtimeFilter)
		if err != nil {
			return handoffStoredRecord{}, err
		}
		runtimeContext = &handoffRuntimeContext{
			AdvisoryOnly: true,
			Workloads:    workloads,
			Findings:     findings,
			Limitations: uniqueStrings(append([]string{
				"Runtime context is exported as evidence-backed runtime state and advisory findings; it does not become canonical incident or report truth inside the sealed bundle.",
			}, append(workloadLimitations, findingLimitations...)...)),
		}
	}

	var recommendations []recommendation
	if request.IncludeRecommendations {
		recommendations, _ = s.listRecommendations(ctx, recommendationFilter{
			event: audit.EventFilter{
				ClusterID:   filter.event.ClusterID,
				Repo:        filter.event.Repo,
				Environment: filter.event.Environment,
				TenantID:    filter.event.TenantID,
				Since:       filter.event.Since,
				Until:       filter.event.Until,
				Limit:       filter.event.Limit,
			},
			SourceType:         "package",
			PackageIncidentIDs: selectedIDs,
			Limit:              10,
		})
	}

	artifactFiles, err := buildHandoffContentFiles(packagePayload, readbackRefs, forensicState, runtimeContext, recommendations)
	if err != nil {
		return handoffStoredRecord{}, err
	}
	manifest := buildSealedManifest(packagePayload, filter, selectionMode, artifactFiles, readbackRefs, forensicRefs, manifestTime)
	manifest.PackageID = handoffPackageID(manifest.RootHash)
	manifestBytes, err := canonicalJSON(manifest)
	if err != nil {
		return handoffStoredRecord{}, err
	}
	manifestHash := digestBytes(manifestBytes)
	manifestBytes, err = canonicalJSON(manifest)
	if err != nil {
		return handoffStoredRecord{}, err
	}
	manifestHash = digestBytes(manifestBytes)

	signSeed, err := s.handoffSeed()
	if err != nil {
		return handoffStoredRecord{}, err
	}
	primarySignature := signHandoffRecord(manifestHash, signSeed, handoffSignerRoleSystem)
	timestamp := buildTimestampRecord(manifestHash, signSeed, manifestTime)
	transparency := buildTransparencyRecord(manifestHash, signSeed, manifestTime)
	record := handoffStoredRecord{
		PackageID:    manifest.PackageID,
		PackageType:  handoffPackageTypeIncidentPackage,
		Manifest:     manifest,
		ManifestHash: manifestHash,
		Session: handoffSessionRecord{
			SessionID:      shortDigest("HOS-", manifestHash),
			PackageID:      manifest.PackageID,
			PackageType:    handoffPackageTypeIncidentPackage,
			ScopeSummary:   manifest.Scope.SelectionSummary,
			InitiatedBy:    incidentActor(principal),
			InitiatedAt:    time.Now().UTC(),
			SignMode:       "offline_public_key",
			CoSignMode:     normalizeCoSignMode(request.CoSignMode),
			Status:         handoffSealStatusSealedBySystem,
			FinalBundleRef: fmt.Sprintf("/v1/handoff/%s/download", manifest.PackageID),
			ManifestHash:   manifestHash,
		},
		Files:           artifactFiles,
		Signatures:      []signatureRecord{primarySignature},
		Timestamp:       timestamp,
		Transparency:    transparency,
		DownloadURI:     fmt.Sprintf("/v1/handoff/%s/download", manifest.PackageID),
		VerificationURI: fmt.Sprintf("/v1/handoff/%s/verification", manifest.PackageID),
	}
	if record.Session.CoSignMode == handoffCoSignRequired {
		record.Session.Status = handoffSealStatusPendingCosign
	} else if record.Session.CoSignMode == handoffCoSignSystemOnly {
		record.Session.Status = handoffSealStatusFullySealed
	}
	record.Bundle = sealedBundleMetadata{
		PackageID:              record.PackageID,
		BundlePath:             fmt.Sprintf("packages/%s.safepkg", strings.ToLower(record.PackageID)),
		ManifestHash:           manifestHash,
		SealStatus:             record.Session.Status,
		SignatureCount:         len(record.Signatures),
		TimestampStatus:        handoffVerificationValid,
		TransparencyStatus:     handoffTransparencyStatusLogged,
		VerificationURI:        record.VerificationURI,
		OfflineVerifierPresent: true,
	}
	record.Verification = s.verifyStoredHandoff(record)
	if err := s.persistHandoffRecord(ctx, principal, record, audit.EventTypeHandoffSealed); err != nil {
		return handoffStoredRecord{}, err
	}
	return record, nil
}

func (s server) cosignStoredHandoff(ctx context.Context, principal auth.Principal, record handoffStoredRecord, request handoffCosignRequest) (handoffStoredRecord, error) {
	if record.Session.CoSignMode == handoffCoSignSystemOnly {
		return handoffStoredRecord{}, errHandoffCosignNotAllowed
	}
	role := normalizeCoSignerRole(request.SignerRole)
	for _, signature := range record.Signatures {
		if signature.SignerRole == role {
			return handoffStoredRecord{}, errHandoffAlreadyCosigned
		}
	}
	signSeed, err := s.handoffSeed()
	if err != nil {
		return handoffStoredRecord{}, err
	}
	if digestBytesMust(canonicalJSONMust(record.Manifest)) != record.ManifestHash {
		return handoffStoredRecord{}, errHandoffManifestMismatch
	}
	record.Signatures = append(record.Signatures, signHandoffRecord(record.ManifestHash, signSeed, role))
	record.Session.Status = handoffSealStatusFullySealed
	record.Bundle.SealStatus = handoffSealStatusFullySealed
	record.Bundle.SignatureCount = len(record.Signatures)
	record.Verification = s.verifyStoredHandoff(record)
	if err := s.persistHandoffRecord(ctx, principal, record, audit.EventTypeHandoffCosigned); err != nil {
		return handoffStoredRecord{}, err
	}
	return record, nil
}

func (s server) persistHandoffRecord(ctx context.Context, principal auth.Principal, record handoffStoredRecord, eventType string) error {
	payload, err := json.Marshal(record)
	if err != nil {
		return err
	}
	_, err = s.store.Ingest(ctx, audit.Event{
		Component:   handoffComponent,
		EventType:   eventType,
		Decision:    audit.DecisionAllow,
		Actor:       incidentActor(principal),
		TenantID:    record.Manifest.Scope.TenantID,
		Environment: record.Manifest.Scope.Environment,
		Repo:        record.Manifest.Scope.Repo,
		Reasons: []string{
			fmt.Sprintf("handoff %s", strings.ReplaceAll(eventType, "handoff_", "")),
			record.PackageID,
		},
		Handoff: payload,
	})
	return err
}

func (s server) getStoredHandoffRecord(ctx context.Context, packageID string) (handoffStoredRecord, error) {
	events, err := s.store.ListEvents(ctx, audit.EventFilter{
		Component: handoffComponent,
		Limit:     2000,
	})
	if err != nil {
		return handoffStoredRecord{}, err
	}
	var latest *handoffStoredRecord
	var latestTime time.Time
	for _, event := range events {
		if event.EventType != audit.EventTypeHandoffSealed && event.EventType != audit.EventTypeHandoffCosigned {
			continue
		}
		if len(event.Handoff) == 0 || string(event.Handoff) == "null" {
			continue
		}
		var record handoffStoredRecord
		if err := json.Unmarshal(event.Handoff, &record); err != nil {
			continue
		}
		if record.PackageID != packageID {
			continue
		}
		if latest == nil || event.ReceivedAt.After(latestTime) {
			copyRecord := record
			latest = &copyRecord
			latestTime = event.ReceivedAt
		}
	}
	if latest == nil {
		return handoffStoredRecord{}, errHandoffNotFound
	}
	return *latest, nil
}

func ensurePrincipalHandoffScope(principal auth.Principal, record handoffStoredRecord) error {
	tenantID := strings.TrimSpace(record.Manifest.Scope.TenantID)
	if tenantID == "" || principal.GlobalScope || strings.TrimSpace(principal.TenantID) == "" {
		return nil
	}
	if strings.TrimSpace(principal.TenantID) != tenantID {
		return auth.ErrInsufficientPermissions
	}
	return nil
}

func buildHandoffContentFiles(packagePayload incidentPackageResponse, readbackRefs []sealedManifestReadbackRef, forensicState *pointInTimeState, runtimeContext *handoffRuntimeContext, recommendations []recommendation) ([]handoffBundleFile, error) {
	packagePayload = normalizeIncidentPackageForHandoff(packagePayload)
	readbackRefs = normalizeSealedManifestReadbackRefs(readbackRefs)
	if forensicState != nil {
		normalized := normalizePointInTimeStateForHandoff(*forensicState)
		forensicState = &normalized
	}
	if runtimeContext != nil {
		normalized := *runtimeContext
		normalized.Workloads = append([]runtimeWorkloadView(nil), runtimeContext.Workloads...)
		normalized.Findings = append([]runtimeIntegrityFinding(nil), runtimeContext.Findings...)
		normalized.Limitations = append([]string(nil), runtimeContext.Limitations...)
		runtimeContext = &normalized
	}
	recommendations = normalizeRecommendationsForHandoff(recommendations)
	files := []handoffBundleFile{}
	htmlReport := renderHandoffHTML(packagePayload)
	files = append(files, handoffBundleFile{
		Path:      "report/report.html",
		MediaType: "text/html; charset=utf-8",
		Role:      "human_report",
		Content:   htmlReport,
	})
	packageBytes, err := canonicalJSON(packagePayload)
	if err != nil {
		return nil, err
	}
	files = append(files, handoffBundleFile{
		Path:      "evidence/package.json",
		MediaType: "application/json",
		Role:      "machine_package",
		Content:   string(packageBytes),
	})
	readbackBytes, err := canonicalJSON(struct {
		ReadbackRefs []sealedManifestReadbackRef `json:"readback_refs"`
	}{ReadbackRefs: readbackRefs})
	if err != nil {
		return nil, err
	}
	files = append(files, handoffBundleFile{
		Path:         "evidence/readback_refs.json",
		MediaType:    "application/json",
		Role:         "readback_lineage",
		Content:      string(readbackBytes),
		AdvisoryOnly: true,
	})
	if forensicState != nil {
		forensicsBytes, err := canonicalJSON(forensicState)
		if err != nil {
			return nil, err
		}
		files = append(files, handoffBundleFile{
			Path:         "evidence/forensics_context.json",
			MediaType:    "application/json",
			Role:         "forensic_context",
			Content:      string(forensicsBytes),
			AdvisoryOnly: true,
		})
	}
	if runtimeContext != nil {
		runtimeBytes, err := canonicalJSON(runtimeContext)
		if err != nil {
			return nil, err
		}
		files = append(files, handoffBundleFile{
			Path:         "evidence/runtime_context.json",
			MediaType:    "application/json",
			Role:         "runtime_context",
			Content:      string(runtimeBytes),
			AdvisoryOnly: true,
		})
	}
	if len(recommendations) > 0 {
		recommendationBytes, err := canonicalJSON(recommendationListResponse{Recommendations: recommendations})
		if err != nil {
			return nil, err
		}
		files = append(files, handoffBundleFile{
			Path:         "advisory/recommendations.json",
			MediaType:    "application/json",
			Role:         "recommendation_overlay",
			Content:      string(recommendationBytes),
			AdvisoryOnly: true,
		})
	}
	return sortHandoffFiles(files), nil
}

func buildSealedManifest(packagePayload incidentPackageResponse, filter incidentFilter, selectionMode string, files []handoffBundleFile, readbackRefs []sealedManifestReadbackRef, forensicRefs []sealedManifestForensicRef, manifestTime time.Time) sealedManifest {
	packagePayload = normalizeIncidentPackageForHandoff(packagePayload)
	readbackRefs = normalizeSealedManifestReadbackRefs(readbackRefs)
	forensicRefs = normalizeSealedManifestForensicRefs(forensicRefs)
	evidenceRefs := []string{}
	for _, item := range packagePayload.Incidents {
		evidenceRefs = append(evidenceRefs, item.IncidentID)
	}
	artifacts := make([]sealedManifestArtifact, 0, len(files))
	for _, file := range files {
		artifacts = append(artifacts, sealedManifestArtifact{
			Path:         file.Path,
			MediaType:    file.MediaType,
			SHA256:       digestString(file.Content),
			Role:         file.Role,
			AdvisoryOnly: file.AdvisoryOnly,
		})
	}
	sort.Slice(artifacts, func(i, j int) bool { return artifacts[i].Path < artifacts[j].Path })
	manifest := sealedManifest{
		PackageType:       handoffPackageTypeIncidentPackage,
		SchemaVersion:     handoffManifestSchemaVersion,
		CreatedAt:         manifestTime,
		GeneratorIdentity: handoffComponent,
		Scope: sealedManifestScope{
			Audience:         packagePayload.Audience,
			SelectionMode:    selectionMode,
			SelectionSummary: packagePayload.SelectionSummary,
			IncidentCount:    packagePayload.IncidentCount,
			IncidentRefs:     cloneStrings(packagePayload.IncidentRefs),
			TenantID:         filter.event.TenantID,
			Environment:      filter.event.Environment,
			Repo:             filter.event.Repo,
		},
		RedactionProfile: sealedManifestRedaction{
			Audience:       packagePayload.Audience,
			ProfileVersion: handoffRedactionProfileVersion,
			Summary:        sortedStrings(cloneStrings(packagePayload.RedactionSummary)),
		},
		Artifacts:    artifacts,
		EvidenceRefs: sortedStrings(limitStrings(uniqueStrings(evidenceRefs), 32)),
		ReadbackRefs: readbackRefs,
		ForensicRefs: forensicRefs,
		Limitations: sortedStrings([]string{
			"Sealed manifest covers the deterministic package scope and its content artifacts before signature, timestamp, and transparency metadata are attached.",
			"Redaction mode is fixed before sealing and is recorded directly in the manifest metadata.",
		}),
	}
	manifest.RootHash = digestBytesMust(canonicalJSONMust(manifestRootPreimage(manifest)))
	return manifest
}

func manifestRootPreimage(manifest sealedManifest) any {
	return struct {
		PackageType       string                      `json:"package_type"`
		SchemaVersion     string                      `json:"schema_version"`
		CreatedAt         time.Time                   `json:"created_at"`
		GeneratorIdentity string                      `json:"generator_identity"`
		Scope             sealedManifestScope         `json:"scope"`
		RedactionProfile  sealedManifestRedaction     `json:"redaction_profile"`
		Artifacts         []sealedManifestArtifact    `json:"artifacts"`
		EvidenceRefs      []string                    `json:"evidence_refs"`
		ReadbackRefs      []sealedManifestReadbackRef `json:"readback_refs,omitempty"`
		ForensicRefs      []sealedManifestForensicRef `json:"forensic_refs,omitempty"`
		Limitations       []string                    `json:"limitations,omitempty"`
	}{
		PackageType:       manifest.PackageType,
		SchemaVersion:     manifest.SchemaVersion,
		CreatedAt:         manifest.CreatedAt,
		GeneratorIdentity: manifest.GeneratorIdentity,
		Scope:             manifest.Scope,
		RedactionProfile:  manifest.RedactionProfile,
		Artifacts:         manifest.Artifacts,
		EvidenceRefs:      manifest.EvidenceRefs,
		ReadbackRefs:      manifest.ReadbackRefs,
		ForensicRefs:      manifest.ForensicRefs,
		Limitations:       manifest.Limitations,
	}
}

func (s server) collectPackageReadbackRefs(incidents []investigationIncident, filter incidentFilter, audience string) []sealedManifestReadbackRef {
	if len(incidents) == 0 {
		return nil
	}
	refs := []sealedManifestReadbackRef{}
	for _, incident := range incidents {
		defense := attachDefenseGapReadback(buildIncidentDefenseGapAssessment(incident, incidents), filter).Readback
		replay := attachPolicyReplayReadback(buildIncidentPolicyReplayAssessment(incident, incidents), filter).Readback
		for _, ref := range []advisoryReadbackRef{defense, replay} {
			if ref.ResourceType == "" || ref.EvidenceHash == "" {
				continue
			}
			item := sealedManifestReadbackRef{
				ResourceType: ref.ResourceType,
				ResourceID:   ref.ResourceID,
				EvidenceHash: ref.EvidenceHash,
				ResourceURI:  ref.ResourceURI,
			}
			if audience == incidentAudienceCustomerSafe {
				item.ResourceID = ""
				item.ResourceURI = ""
			}
			refs = append(refs, item)
		}
	}
	sort.Slice(refs, func(i, j int) bool {
		if refs[i].ResourceType == refs[j].ResourceType {
			return refs[i].EvidenceHash < refs[j].EvidenceHash
		}
		return refs[i].ResourceType < refs[j].ResourceType
	})
	dedup := []sealedManifestReadbackRef{}
	seen := map[string]struct{}{}
	for _, ref := range refs {
		key := strings.Join([]string{ref.ResourceType, ref.ResourceID, ref.EvidenceHash}, "|")
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		dedup = append(dedup, ref)
	}
	return dedup
}

func normalizeCoSignMode(value string) string {
	switch strings.TrimSpace(strings.ToLower(value)) {
	case "", handoffCoSignSystemOnly:
		return handoffCoSignSystemOnly
	case handoffCoSignOptional:
		return handoffCoSignOptional
	case handoffCoSignRequired:
		return handoffCoSignRequired
	default:
		return handoffCoSignSystemOnly
	}
}

func normalizeCoSignerRole(value string) string {
	switch strings.TrimSpace(strings.ToLower(value)) {
	case handoffSignerRoleCISO:
		return handoffSignerRoleCISO
	case handoffSignerRoleAuditor:
		return handoffSignerRoleAuditor
	case handoffSignerRoleDelegatedApprover:
		return handoffSignerRoleDelegatedApprover
	default:
		return handoffSignerRoleCISO
	}
}

func deterministicManifestTimestamp(incidents []investigationIncident, filter incidentFilter) time.Time {
	values := []time.Time{}
	for _, incident := range incidents {
		for _, pointer := range []*time.Time{incident.UpdatedAt, incident.ResolvedAt, incident.OpenedAt, incident.LastActivityAt, incident.LastSeenAt} {
			if pointer != nil && !pointer.IsZero() {
				values = append(values, pointer.UTC())
			}
		}
	}
	if len(values) == 0 {
		if filter.event.Until != nil {
			return filter.event.Until.UTC().Truncate(time.Second)
		}
		if filter.event.Since != nil {
			return filter.event.Since.UTC().Truncate(time.Second)
		}
		return time.Unix(0, 0).UTC()
	}
	sort.Slice(values, func(i, j int) bool { return values[i].Before(values[j]) })
	return values[len(values)-1].UTC().Truncate(time.Second)
}

func renderHandoffHTML(payload incidentPackageResponse) string {
	builder := &strings.Builder{}
	builder.WriteString("<!doctype html><html><head><meta charset=\"utf-8\"><title>ChangeLock sealed handoff</title>")
	builder.WriteString("<style>body{font-family:ui-sans-serif,system-ui,sans-serif;margin:32px;color:#15202b}h1,h2{margin:0 0 12px}section{margin:0 0 24px}table{width:100%;border-collapse:collapse}th,td{padding:8px 10px;border-bottom:1px solid #d8dee4;text-align:left}code{font-size:12px;background:#f6f8fa;padding:2px 4px;border-radius:4px}.chips span{display:inline-block;margin:0 8px 8px 0;padding:4px 8px;border-radius:999px;background:#eef2f7;font-size:12px}</style>")
	builder.WriteString("</head><body>")
	builder.WriteString("<section><h1>ChangeLock sealed handoff</h1>")
	builder.WriteString(fmt.Sprintf("<p>%s</p>", escapeHTML(payload.PackageSummary)))
	builder.WriteString("<div class=\"chips\">")
	builder.WriteString(fmt.Sprintf("<span>Audience: %s</span>", escapeHTML(payload.Audience)))
	builder.WriteString(fmt.Sprintf("<span>Incidents: %d</span>", payload.IncidentCount))
	builder.WriteString(fmt.Sprintf("<span>Selection: %s</span>", escapeHTML(payload.SelectionMode)))
	builder.WriteString("</div></section>")
	builder.WriteString("<section><h2>Package intelligence</h2>")
	builder.WriteString(fmt.Sprintf("<p>%s</p>", escapeHTML(payload.PackageIntel.RecommendedActions.WhyThisMattersNow)))
	builder.WriteString("</section>")
	builder.WriteString("<section><h2>Included incidents</h2><table><thead><tr><th>ID</th><th>Title</th><th>State</th><th>Severity</th><th>Scope</th></tr></thead><tbody>")
	for _, item := range payload.Incidents {
		builder.WriteString("<tr>")
		builder.WriteString(fmt.Sprintf("<td><code>%s</code></td>", escapeHTML(item.IncidentID)))
		builder.WriteString(fmt.Sprintf("<td>%s</td>", escapeHTML(item.Title)))
		builder.WriteString(fmt.Sprintf("<td>%s</td>", escapeHTML(item.State)))
		builder.WriteString(fmt.Sprintf("<td>%s</td>", escapeHTML(item.Severity)))
		builder.WriteString(fmt.Sprintf("<td>%s</td>", escapeHTML(item.ScopeLabel)))
		builder.WriteString("</tr>")
	}
	builder.WriteString("</tbody></table></section>")
	if len(payload.RedactionSummary) > 0 {
		builder.WriteString("<section><h2>Redaction profile</h2><ul>")
		for _, item := range payload.RedactionSummary {
			builder.WriteString(fmt.Sprintf("<li>%s</li>", escapeHTML(item)))
		}
		builder.WriteString("</ul></section>")
	}
	builder.WriteString("</body></html>")
	return builder.String()
}

func signHandoffRecord(manifestHash, seedSource, signerRole string) signatureRecord {
	publicKey, privateKey, signerIdentity := handoffKeypair(seedSource, signerRole)
	signature := ed25519.Sign(privateKey, []byte(manifestHash))
	signedAt := time.Now().UTC()
	return signatureRecord{
		SignatureID:    shortDigest("SIG-", manifestHash+":"+signerRole),
		SignerRole:     signerRole,
		SignerIdentity: signerIdentity,
		Algorithm:      handoffSignatureAlgorithm,
		SignedObject:   "sealed_manifest",
		SignedHash:     manifestHash,
		SignatureValue: base64.StdEncoding.EncodeToString(signature),
		PublicKey:      base64.StdEncoding.EncodeToString(publicKey),
		CertificateChainRefs: []string{
			fmt.Sprintf("verify/public_keys.json#%s", signerRole),
		},
		SignedAt: signedAt,
	}
}

func buildTimestampRecord(manifestHash, seedSource string, timestamp time.Time) timestampRecord {
	publicKey, privateKey, signerIdentity := handoffKeypair(seedSource, "timestamp")
	payload := fmt.Sprintf("%s|%s|%s", handoffTimestampTypeRFC3161Like, manifestHash, timestamp.UTC().Format(time.RFC3339))
	signature := ed25519.Sign(privateKey, []byte(payload))
	return timestampRecord{
		TimestampType:  handoffTimestampTypeRFC3161Like,
		Authority:      "changelock-timestamp-authority",
		SubjectHash:    manifestHash,
		TokenRef:       "timestamp/manifest.timestamp.json",
		TimestampedAt:  timestamp.UTC(),
		Algorithm:      handoffSignatureAlgorithm,
		SignatureValue: base64.StdEncoding.EncodeToString(signature),
		SignerIdentity: signerIdentity,
		PublicKey:      base64.StdEncoding.EncodeToString(publicKey),
	}
}

func buildTransparencyRecord(manifestHash, seedSource string, timestamp time.Time) transparencyRecord {
	publicKey, privateKey, signerIdentity := handoffKeypair(seedSource, "transparency")
	entryID := shortDigest("LOG-", manifestHash)
	proof := transparencyInclusionProof{
		TreeSize:  1,
		LeafIndex: 0,
		LeafHash:  digestString(manifestHash),
		RootHash:  digestString("embedded-log|" + manifestHash),
		Path:      []string{},
	}
	payload := fmt.Sprintf("%s|%s|%s|%s", "embedded-handoff-log", entryID, manifestHash, proof.RootHash)
	signature := ed25519.Sign(privateKey, []byte(payload))
	return transparencyRecord{
		LogID:          "embedded-handoff-log",
		EntryID:        entryID,
		SubjectHash:    manifestHash,
		InclusionProof: proof,
		LoggedAt:       timestamp.UTC(),
		Status:         handoffTransparencyStatusLogged,
		Algorithm:      handoffSignatureAlgorithm,
		SignatureValue: base64.StdEncoding.EncodeToString(signature),
		SignerIdentity: signerIdentity,
		PublicKey:      base64.StdEncoding.EncodeToString(publicKey),
	}
}

func buildHandoffBundle(record handoffStoredRecord) ([]byte, error) {
	buffer := &bytes.Buffer{}
	zipWriter := zip.NewWriter(buffer)
	files := append([]handoffBundleFile{}, record.Files...)
	manifestBytes, err := canonicalJSON(record.Manifest)
	if err != nil {
		return nil, err
	}
	signatureBytes, err := canonicalJSON(record.Signatures)
	if err != nil {
		return nil, err
	}
	timestampBytes, err := canonicalJSON(record.Timestamp)
	if err != nil {
		return nil, err
	}
	transparencyBytes, err := canonicalJSON(record.Transparency)
	if err != nil {
		return nil, err
	}
	publicKeyBytes, err := canonicalJSON(buildPublicKeyBundle(record))
	if err != nil {
		return nil, err
	}
	files = append(files,
		handoffBundleFile{Path: "manifest.json", MediaType: "application/json", Role: "sealed_manifest", Content: string(manifestBytes)},
		handoffBundleFile{Path: "signatures/signatures.json", MediaType: "application/json", Role: "signature_set", Content: string(signatureBytes)},
		handoffBundleFile{Path: "timestamp/manifest.timestamp.json", MediaType: "application/json", Role: "timestamp_record", Content: string(timestampBytes)},
		handoffBundleFile{Path: "transparency/manifest.transparency.json", MediaType: "application/json", Role: "transparency_record", Content: string(transparencyBytes)},
		handoffBundleFile{Path: "verify/public_keys.json", MediaType: "application/json", Role: "offline_verifier_keys", Content: string(publicKeyBytes)},
		handoffBundleFile{Path: "verify/verification_instructions.txt", MediaType: "text/plain; charset=utf-8", Role: "offline_verifier", Content: buildVerificationInstructions(record)},
		handoffBundleFile{Path: "verify/manifest.sha256", MediaType: "text/plain; charset=utf-8", Role: "offline_verifier", Content: record.ManifestHash + "\n"},
	)
	files = sortHandoffFiles(files)
	for _, file := range files {
		header := &zip.FileHeader{Name: file.Path, Method: zip.Deflate}
		header.Modified = time.Unix(0, 0).UTC()
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return nil, err
		}
		if _, err := writer.Write([]byte(file.Content)); err != nil {
			return nil, err
		}
	}
	if err := zipWriter.Close(); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func parseHandoffBundle(bundle []byte) (handoffStoredRecord, error) {
	reader, err := zip.NewReader(bytes.NewReader(bundle), int64(len(bundle)))
	if err != nil {
		return handoffStoredRecord{}, errHandoffInvalidBundle
	}
	files := map[string]string{}
	for _, item := range reader.File {
		handle, err := item.Open()
		if err != nil {
			return handoffStoredRecord{}, errHandoffInvalidBundle
		}
		content, err := io.ReadAll(handle)
		handle.Close()
		if err != nil {
			return handoffStoredRecord{}, errHandoffInvalidBundle
		}
		files[item.Name] = string(content)
	}
	var manifest sealedManifest
	if err := json.Unmarshal([]byte(files["manifest.json"]), &manifest); err != nil {
		return handoffStoredRecord{}, errHandoffInvalidBundle
	}
	var signatures []signatureRecord
	if err := json.Unmarshal([]byte(files["signatures/signatures.json"]), &signatures); err != nil {
		return handoffStoredRecord{}, errHandoffInvalidBundle
	}
	var timestamp timestampRecord
	if err := json.Unmarshal([]byte(files["timestamp/manifest.timestamp.json"]), &timestamp); err != nil {
		return handoffStoredRecord{}, errHandoffInvalidBundle
	}
	var transparency transparencyRecord
	if err := json.Unmarshal([]byte(files["transparency/manifest.transparency.json"]), &transparency); err != nil {
		return handoffStoredRecord{}, errHandoffInvalidBundle
	}
	contentFiles := []handoffBundleFile{}
	for _, artifact := range manifest.Artifacts {
		content, ok := files[artifact.Path]
		if !ok {
			return handoffStoredRecord{}, errHandoffInvalidBundle
		}
		contentFiles = append(contentFiles, handoffBundleFile{
			Path:         artifact.Path,
			MediaType:    artifact.MediaType,
			Role:         artifact.Role,
			Content:      content,
			AdvisoryOnly: artifact.AdvisoryOnly,
		})
	}
	manifestBytes, err := canonicalJSON(manifest)
	if err != nil {
		return handoffStoredRecord{}, errHandoffInvalidBundle
	}
	record := handoffStoredRecord{
		PackageID:    manifest.PackageID,
		PackageType:  manifest.PackageType,
		Manifest:     manifest,
		ManifestHash: digestBytes(manifestBytes),
		Files:        sortHandoffFiles(contentFiles),
		Signatures:   signatures,
		Timestamp:    timestamp,
		Transparency: transparency,
	}
	record.Bundle = sealedBundleMetadata{
		PackageID:              record.PackageID,
		BundlePath:             fmt.Sprintf("packages/%s.safepkg", strings.ToLower(record.PackageID)),
		ManifestHash:           record.ManifestHash,
		SealStatus:             handoffSealStatusFullySealed,
		SignatureCount:         len(record.Signatures),
		TimestampStatus:        handoffVerificationValid,
		TransparencyStatus:     record.Transparency.Status,
		VerificationURI:        fmt.Sprintf("/v1/handoff/%s/verification", record.PackageID),
		OfflineVerifierPresent: true,
	}
	record.DownloadURI = fmt.Sprintf("/v1/handoff/%s/download", record.PackageID)
	record.VerificationURI = record.Bundle.VerificationURI
	record.Verification = verificationResult{}
	return record, nil
}

func (s server) verifyStoredHandoff(record handoffStoredRecord) verificationResult {
	result := verificationResult{
		PackageID:        record.PackageID,
		SignerIdentities: []string{},
		RedactionProfile: record.Manifest.RedactionProfile.Audience,
	}
	manifestHash := digestBytesMust(canonicalJSONMust(record.Manifest))
	rootHash := digestBytesMust(canonicalJSONMust(manifestRootPreimage(record.Manifest)))
	result.ManifestValid = manifestHash == record.ManifestHash && rootHash == record.Manifest.RootHash
	artifactValid := true
	for _, artifact := range record.Manifest.Artifacts {
		content := ""
		found := false
		for _, file := range record.Files {
			if file.Path == artifact.Path {
				content = file.Content
				found = true
				break
			}
		}
		if !found || digestString(content) != artifact.SHA256 {
			artifactValid = false
			break
		}
	}
	result.ArtifactHashesValid = artifactValid
	signaturesValid := len(record.Signatures) > 0
	for _, signature := range record.Signatures {
		pub, err := decodeHandoffPublicKey(signature.PublicKey)
		if err != nil {
			signaturesValid = false
			break
		}
		signatureBytes, err := base64.StdEncoding.DecodeString(signature.SignatureValue)
		if err != nil || !ed25519.Verify(pub, []byte(signature.SignedHash), signatureBytes) {
			signaturesValid = false
			break
		}
		result.SignerIdentities = append(result.SignerIdentities, signature.SignerIdentity)
	}
	result.SignaturesValid = signaturesValid
	result.TimestampValid = verifyTimestampRecord(record.Timestamp, record.ManifestHash)
	result.TransparencyValid = verifyTransparencyRecord(record.Transparency, record.ManifestHash)
	switch {
	case result.ManifestValid && result.ArtifactHashesValid && result.SignaturesValid && result.TimestampValid && result.TransparencyValid:
		result.OverallStatus = handoffVerificationValid
	case result.ManifestValid && result.ArtifactHashesValid:
		result.OverallStatus = handoffVerificationPartial
	default:
		result.OverallStatus = handoffVerificationInvalid
	}
	if !result.TransparencyValid {
		result.Limitations = append(result.Limitations, "Transparency verification did not validate against the embedded log proof.")
	}
	if !result.TimestampValid {
		result.Limitations = append(result.Limitations, "Timestamp verification did not validate against the detached timestamp token.")
	}
	if !result.SignaturesValid {
		result.Limitations = append(result.Limitations, "One or more manifest signatures did not verify against the included signer public keys.")
	}
	if !result.ManifestValid {
		result.Limitations = append(result.Limitations, "Manifest hash or root hash no longer matches the canonical sealed manifest structure.")
	}
	if !result.ArtifactHashesValid {
		result.Limitations = append(result.Limitations, "At least one sealed artifact no longer matches the hash recorded in the manifest.")
	}
	return result
}

func verifyTimestampRecord(record timestampRecord, manifestHash string) bool {
	publicKey, err := decodeHandoffPublicKey(record.PublicKey)
	if err != nil || record.SubjectHash != manifestHash {
		return false
	}
	payload := fmt.Sprintf("%s|%s|%s", record.TimestampType, record.SubjectHash, record.TimestampedAt.UTC().Format(time.RFC3339))
	signature, err := base64.StdEncoding.DecodeString(record.SignatureValue)
	return err == nil && ed25519.Verify(publicKey, []byte(payload), signature)
}

func verifyTransparencyRecord(record transparencyRecord, manifestHash string) bool {
	publicKey, err := decodeHandoffPublicKey(record.PublicKey)
	if err != nil || record.SubjectHash != manifestHash {
		return false
	}
	if record.InclusionProof.RootHash != digestString("embedded-log|"+manifestHash) {
		return false
	}
	payload := fmt.Sprintf("%s|%s|%s|%s", record.LogID, record.EntryID, record.SubjectHash, record.InclusionProof.RootHash)
	signature, err := base64.StdEncoding.DecodeString(record.SignatureValue)
	return err == nil && ed25519.Verify(publicKey, []byte(payload), signature)
}

func buildPublicKeyBundle(record handoffStoredRecord) map[string]string {
	keys := map[string]string{}
	for _, signature := range record.Signatures {
		if signature.PublicKey != "" {
			keys[signature.SignerRole] = signature.PublicKey
		}
	}
	if record.Timestamp.PublicKey != "" {
		keys["timestamp"] = record.Timestamp.PublicKey
	}
	if record.Transparency.PublicKey != "" {
		keys["transparency"] = record.Transparency.PublicKey
	}
	return keys
}

func buildVerificationInstructions(record handoffStoredRecord) string {
	return strings.Join([]string{
		"ChangeLock sealed handoff verification instructions",
		fmt.Sprintf("Package ID: %s", record.PackageID),
		fmt.Sprintf("Manifest hash: %s", record.ManifestHash),
		"",
		"Offline verification steps:",
		"1. Recompute sha256 for each file listed in manifest.json and compare them with manifest.artifacts[].sha256.",
		"2. Recompute the manifest hash from manifest.json and compare it with verify/manifest.sha256.",
		"3. Verify each signature in signatures/signatures.json against the manifest hash using the matching public key in verify/public_keys.json.",
		"4. Verify timestamp/manifest.timestamp.json against the manifest hash and the timestamp authority public key.",
		"5. Verify transparency/manifest.transparency.json against the manifest hash and the transparency public key.",
		"",
		"Offline verification is authoritative. The ChangeLock verification endpoint is only a convenience surface.",
	}, "\n")
}

func (s server) handoffSeed() (string, error) {
	if value := strings.TrimSpace(os.Getenv("CHANGELOCK_HANDOFF_SIGNING_SEED")); value != "" {
		return value, nil
	}
	if s.signing != nil && s.signing.runtime != nil && strings.TrimSpace(s.signing.runtime.Config.SoftwareSecret) != "" {
		return strings.TrimSpace(s.signing.runtime.Config.SoftwareSecret), nil
	}
	return "", errHandoffSigningDisabled
}

func decodeHandoffPublicKey(value string) (ed25519.PublicKey, error) {
	decoded, err := base64.StdEncoding.DecodeString(strings.TrimSpace(value))
	if err != nil {
		return nil, err
	}
	return ed25519.PublicKey(decoded), nil
}

func handoffKeypair(seedSource, role string) (ed25519.PublicKey, ed25519.PrivateKey, string) {
	sum := sha256.Sum256([]byte("changelock-handoff|" + role + "|" + seedSource))
	privateKey := ed25519.NewKeyFromSeed(sum[:])
	publicKey := privateKey.Public().(ed25519.PublicKey)
	signerIdentity := fmt.Sprintf("handoff:%s:%s", role, hex.EncodeToString(publicKey[:4]))
	return publicKey, privateKey, signerIdentity
}

func handoffPackageID(manifestHash string) string {
	return "HOF-" + strings.ToUpper(strings.TrimPrefix(shortDigest("", manifestHash), ""))
}

func sortHandoffFiles(files []handoffBundleFile) []handoffBundleFile {
	items := append([]handoffBundleFile(nil), files...)
	sort.Slice(items, func(i, j int) bool { return items[i].Path < items[j].Path })
	return items
}

func sortHandoffIncidents(incidents []investigationIncident) []investigationIncident {
	items := append([]investigationIncident(nil), incidents...)
	sort.Slice(items, func(i, j int) bool {
		if items[i].ID == items[j].ID {
			return firstNonEmpty(items[i].ScopeRef, items[i].Title) < firstNonEmpty(items[j].ScopeRef, items[j].Title)
		}
		return items[i].ID < items[j].ID
	})
	return items
}

func normalizeIncidentPackageForHandoff(payload incidentPackageResponse) incidentPackageResponse {
	payload.RedactionSummary = sortedStrings(cloneStrings(payload.RedactionSummary))
	payload.IncidentRefs = sortedStrings(cloneStrings(payload.IncidentRefs))
	payload.Limitations = sortedStrings(cloneStrings(payload.Limitations))
	payload.Incidents = append([]incidentPackageItem(nil), payload.Incidents...)
	sort.Slice(payload.Incidents, func(i, j int) bool {
		if payload.Incidents[i].IncidentID == payload.Incidents[j].IncidentID {
			return payload.Incidents[i].Title < payload.Incidents[j].Title
		}
		return payload.Incidents[i].IncidentID < payload.Incidents[j].IncidentID
	})
	payload.PackageIntel = normalizePackageIntelligenceForHandoff(payload.PackageIntel)
	return payload
}

func normalizePackageIntelligenceForHandoff(value packageIntelligence) packageIntelligence {
	value.DefenseGapSummary.TopGapTypes = sortedStrings(cloneStrings(value.DefenseGapSummary.TopGapTypes))
	value.DefenseGapSummary.TopFindings = normalizeDefenseGapFindingsForHandoff(value.DefenseGapSummary.TopFindings)
	value.DefenseGapSummary.Limitations = sortedStrings(cloneStrings(value.DefenseGapSummary.Limitations))
	value.PolicyReplaySummary.BlastRadius.TopScopes = sortedStrings(cloneStrings(value.PolicyReplaySummary.BlastRadius.TopScopes))
	value.PolicyReplaySummary.TopCoverageGaps = normalizeCoverageGapFindingsForHandoff(value.PolicyReplaySummary.TopCoverageGaps)
	value.PolicyReplaySummary.Limitations = sortedStrings(cloneStrings(value.PolicyReplaySummary.Limitations))
	value.SystemicWeakness.TopPatterns = normalizePackageSystemicPatternsForHandoff(value.SystemicWeakness.TopPatterns)
	value.SystemicWeakness.Limitations = sortedStrings(cloneStrings(value.SystemicWeakness.Limitations))
	value.RecommendedActions.ImmediateContainment = sortedStrings(cloneStrings(value.RecommendedActions.ImmediateContainment))
	value.RecommendedActions.NearTermHardening = sortedStrings(cloneStrings(value.RecommendedActions.NearTermHardening))
	value.RecommendedActions.GovernanceFix = sortedStrings(cloneStrings(value.RecommendedActions.GovernanceFix))
	return value
}

func normalizeDefenseGapFindingsForHandoff(values []defenseGapFinding) []defenseGapFinding {
	items := append([]defenseGapFinding(nil), values...)
	for index := range items {
		items[index].EvidenceRefs = sortedStrings(cloneStrings(items[index].EvidenceRefs))
		items[index].RelatedIncidentRefs = sortedStrings(cloneStrings(items[index].RelatedIncidentRefs))
		items[index].RecommendedActions.Containment = sortedStrings(cloneStrings(items[index].RecommendedActions.Containment))
		items[index].RecommendedActions.Hardening = sortedStrings(cloneStrings(items[index].RecommendedActions.Hardening))
		items[index].RecommendedActions.GovernanceFix = sortedStrings(cloneStrings(items[index].RecommendedActions.GovernanceFix))
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].GapType == items[j].GapType {
			return items[i].Title < items[j].Title
		}
		return items[i].GapType < items[j].GapType
	})
	return items
}

func normalizeCoverageGapFindingsForHandoff(values []coverageGapFinding) []coverageGapFinding {
	items := append([]coverageGapFinding(nil), values...)
	for index := range items {
		items[index].EvidenceRefs = sortedStrings(cloneStrings(items[index].EvidenceRefs))
		items[index].RelatedIncidentRefs = sortedStrings(cloneStrings(items[index].RelatedIncidentRefs))
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].GapType == items[j].GapType {
			return items[i].Title < items[j].Title
		}
		return items[i].GapType < items[j].GapType
	})
	return items
}

func normalizePackageSystemicPatternsForHandoff(values []packageSystemicPattern) []packageSystemicPattern {
	items := append([]packageSystemicPattern(nil), values...)
	for index := range items {
		items[index].RelatedIncidentRefs = sortedStrings(cloneStrings(items[index].RelatedIncidentRefs))
		items[index].EvidenceRefs = sortedStrings(cloneStrings(items[index].EvidenceRefs))
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].PatternKey == items[j].PatternKey {
			return items[i].Title < items[j].Title
		}
		return items[i].PatternKey < items[j].PatternKey
	})
	return items
}

func normalizePointInTimeStateForHandoff(state pointInTimeState) pointInTimeState {
	state.PolicyContext.ActiveRules = sortedStrings(cloneStrings(state.PolicyContext.ActiveRules))
	state.PolicyContext.RuleVersions = sortedStrings(cloneStrings(state.PolicyContext.RuleVersions))
	state.InventoryContext.RunningSubjects = sortedStrings(cloneStrings(state.InventoryContext.RunningSubjects))
	state.InventoryContext.ArtifactDigests = sortedStrings(cloneStrings(state.InventoryContext.ArtifactDigests))
	state.InventoryContext.SBOMRefs = sortedStrings(cloneStrings(state.InventoryContext.SBOMRefs))
	state.VulnerabilityContext.KnownFindings = normalizeHistoricalVulnerabilityFindingsForHandoff(state.VulnerabilityContext.KnownFindings)
	state.VulnerabilityContext.UnknownLaterDisclosedRefs = sortedStrings(cloneStrings(state.VulnerabilityContext.UnknownLaterDisclosedRefs))
	state.VulnerabilityContext.VEXState = normalizeHistoricalVEXStatesForHandoff(state.VulnerabilityContext.VEXState)
	state.IdentityContext.Signers = sortedStrings(cloneStrings(state.IdentityContext.Signers))
	state.IdentityContext.TrustRoots = sortedStrings(cloneStrings(state.IdentityContext.TrustRoots))
	state.IdentityContext.IdentityDriftFlags = sortedStrings(cloneStrings(state.IdentityContext.IdentityDriftFlags))
	state.ExceptionContext.ActiveExceptions = sortedStrings(cloneStrings(state.ExceptionContext.ActiveExceptions))
	state.IncidentContext.RelevantIncidents = normalizeForensicsIncidentSummariesForHandoff(state.IncidentContext.RelevantIncidents)
	if state.TopologyContext != nil {
		topologyContext := *state.TopologyContext
		topologyContext.TopRiskPaths = sortedStrings(cloneStrings(topologyContext.TopRiskPaths))
		topologyContext.Heatmap = normalizeTopologyNodesForHandoff(topologyContext.Heatmap)
		topologyContext.Limitations = sortedStrings(cloneStrings(topologyContext.Limitations))
		state.TopologyContext = &topologyContext
	}
	state.EvidenceRefs = sortedStrings(cloneStrings(state.EvidenceRefs))
	state.ReadbackRefs = normalizeAdvisoryReadbackRefsForHandoff(state.ReadbackRefs)
	state.Limitations = sortedStrings(cloneStrings(state.Limitations))
	return state
}

func normalizeHistoricalVulnerabilityFindingsForHandoff(values []historicalVulnerabilityFinding) []historicalVulnerabilityFinding {
	items := append([]historicalVulnerabilityFinding(nil), values...)
	for index := range items {
		items[index].EvidenceRefs = sortedStrings(cloneStrings(items[index].EvidenceRefs))
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].CVEID == items[j].CVEID {
			return items[i].ImageDigest < items[j].ImageDigest
		}
		return items[i].CVEID < items[j].CVEID
	})
	return items
}

func normalizeHistoricalVEXStatesForHandoff(values []historicalVEXState) []historicalVEXState {
	items := append([]historicalVEXState(nil), values...)
	sort.Slice(items, func(i, j int) bool {
		if items[i].CreatedAt.Equal(items[j].CreatedAt) {
			return items[i].StatementID < items[j].StatementID
		}
		return items[i].CreatedAt.Before(items[j].CreatedAt)
	})
	return items
}

func normalizeForensicsIncidentSummariesForHandoff(values []forensicsIncidentSummary) []forensicsIncidentSummary {
	items := append([]forensicsIncidentSummary(nil), values...)
	sort.Slice(items, func(i, j int) bool {
		if items[i].IncidentID == items[j].IncidentID {
			return items[i].ScopeRef < items[j].ScopeRef
		}
		return items[i].IncidentID < items[j].IncidentID
	})
	return items
}

func normalizeRecommendationsForHandoff(values []recommendation) []recommendation {
	items := append([]recommendation(nil), values...)
	for index := range items {
		items[index].EvidenceRefs = sortedStrings(cloneStrings(items[index].EvidenceRefs))
		items[index].ReadbackRefs = normalizeAdvisoryReadbackRefsForHandoff(items[index].ReadbackRefs)
		items[index].RelatedIncidentRefs = sortedStrings(cloneStrings(items[index].RelatedIncidentRefs))
		items[index].VerificationPlan = sortedStrings(cloneStrings(items[index].VerificationPlan))
		items[index].ActionTemplate.RequiredInputs = sortedStrings(cloneStrings(items[index].ActionTemplate.RequiredInputs))
		items[index].ActionTemplate.AllowedAudiences = sortedStrings(cloneStrings(items[index].ActionTemplate.AllowedAudiences))
		items[index].Comments = normalizeRecommendationCommentsForHandoff(items[index].Comments)
		items[index].History = normalizeRecommendationHistoryForHandoff(items[index].History)
		items[index].Limitations = sortedStrings(cloneStrings(items[index].Limitations))
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].RecommendationID == items[j].RecommendationID {
			return items[i].SourceRef < items[j].SourceRef
		}
		return items[i].RecommendationID < items[j].RecommendationID
	})
	return items
}

func normalizeRecommendationCommentsForHandoff(values []recommendationComment) []recommendationComment {
	items := append([]recommendationComment(nil), values...)
	sort.Slice(items, func(i, j int) bool {
		left := time.Time{}
		right := time.Time{}
		if items[i].Timestamp != nil {
			left = items[i].Timestamp.UTC()
		}
		if items[j].Timestamp != nil {
			right = items[j].Timestamp.UTC()
		}
		if left.Equal(right) {
			return items[i].ID < items[j].ID
		}
		return left.Before(right)
	})
	return items
}

func normalizeRecommendationHistoryForHandoff(values []recommendationHistoryEntry) []recommendationHistoryEntry {
	items := append([]recommendationHistoryEntry(nil), values...)
	sort.Slice(items, func(i, j int) bool {
		left := time.Time{}
		right := time.Time{}
		if items[i].Timestamp != nil {
			left = items[i].Timestamp.UTC()
		}
		if items[j].Timestamp != nil {
			right = items[j].Timestamp.UTC()
		}
		if left.Equal(right) {
			return items[i].ID < items[j].ID
		}
		return left.Before(right)
	})
	return items
}

func normalizeAdvisoryReadbackRefsForHandoff(values []advisoryReadbackRef) []advisoryReadbackRef {
	items := append([]advisoryReadbackRef(nil), values...)
	sort.Slice(items, func(i, j int) bool {
		if items[i].ResourceType == items[j].ResourceType {
			if items[i].ResourceID == items[j].ResourceID {
				return items[i].EvidenceHash < items[j].EvidenceHash
			}
			return items[i].ResourceID < items[j].ResourceID
		}
		return items[i].ResourceType < items[j].ResourceType
	})
	return items
}

func normalizeSealedManifestReadbackRefs(values []sealedManifestReadbackRef) []sealedManifestReadbackRef {
	items := append([]sealedManifestReadbackRef(nil), values...)
	sort.Slice(items, func(i, j int) bool {
		if items[i].ResourceType == items[j].ResourceType {
			if items[i].ResourceID == items[j].ResourceID {
				return items[i].EvidenceHash < items[j].EvidenceHash
			}
			return items[i].ResourceID < items[j].ResourceID
		}
		return items[i].ResourceType < items[j].ResourceType
	})
	return items
}

func normalizeSealedManifestForensicRefs(values []sealedManifestForensicRef) []sealedManifestForensicRef {
	items := append([]sealedManifestForensicRef(nil), values...)
	sort.Slice(items, func(i, j int) bool {
		if items[i].Timestamp == items[j].Timestamp {
			return items[i].ContextURI < items[j].ContextURI
		}
		return items[i].Timestamp < items[j].Timestamp
	})
	return items
}

func normalizeTopologyNodesForHandoff(values []topologyNode) []topologyNode {
	items := append([]topologyNode(nil), values...)
	for index := range items {
		items[index].EvidenceRefs = sortedStrings(cloneStrings(items[index].EvidenceRefs))
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].NodeID == items[j].NodeID {
			if items[i].Service == items[j].Service {
				return items[i].ArtifactDigest < items[j].ArtifactDigest
			}
			return items[i].Service < items[j].Service
		}
		return items[i].NodeID < items[j].NodeID
	})
	return items
}

func sortedStrings(values []string) []string {
	items := append([]string(nil), values...)
	sort.Strings(items)
	return items
}

func canonicalJSON(value any) ([]byte, error) {
	return json.Marshal(value)
}

func canonicalJSONMust(value any) []byte {
	payload, err := canonicalJSON(value)
	if err != nil {
		panic(err)
	}
	return payload
}

func digestString(value string) string {
	return digestBytes([]byte(value))
}

func digestBytes(payload []byte) string {
	sum := sha256.Sum256(payload)
	return "sha256:" + hex.EncodeToString(sum[:])
}

func digestBytesMust(payload []byte) string {
	return digestBytes(payload)
}

func shortDigest(prefix, value string) string {
	sum := sha256.Sum256([]byte(value))
	return prefix + strings.ToUpper(hex.EncodeToString(sum[:]))[:12]
}

func escapeHTML(value string) string {
	replacer := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "\"", "&quot;")
	return replacer.Replace(value)
}

func urlQueryEscape(value string) string {
	replacer := strings.NewReplacer(" ", "%20", ":", "%3A", "/", "%2F", "@", "%40")
	return replacer.Replace(value)
}

func writeHandoffError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, errHandoffNotFound):
		status = http.StatusNotFound
	case errors.Is(err, errHandoffSigningDisabled), errors.Is(err, errHandoffInvalidScope), errors.Is(err, errHandoffInvalidBundle), errors.Is(err, errHandoffManifestMismatch):
		status = http.StatusBadRequest
	case errors.Is(err, errHandoffAlreadyCosigned), errors.Is(err, errHandoffCosignNotAllowed):
		status = http.StatusConflict
	case errors.Is(err, context.DeadlineExceeded):
		status = http.StatusGatewayTimeout
	}
	httpjson.Write(w, status, map[string]string{"error": err.Error()})
}
