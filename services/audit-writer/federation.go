package main

import (
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
	federationComponent = "federation-manager"

	federationLocalPeerID = "local-instance"

	federationPeerStatusActive      = "active"
	federationPeerStatusStale       = "stale"
	federationPeerStatusUnreachable = "unreachable"

	federationPolicyRoleLeader   = "leader"
	federationPolicyRoleFollower = "follower"
	federationPolicyRoleSupplier = "supplier"

	federationDecisionAccepted               = "accepted"
	federationDecisionAcceptedWithOverrides  = "accepted_with_local_overrides"
	federationDecisionRejectedUnverifiable   = "rejected_unverifiable"
	federationDecisionRejectedScopeMismatch  = "rejected_scope_mismatch"
	federationDecisionRejectedPolicyConflict = "rejected_policy_conflict"
	federationDecisionRejectedStale          = "rejected_stale"
	federationDecisionRejectedUntrustedPeer  = "rejected_untrusted_peer"

	federationSyncStatusLocalOnly           = "local_only"
	federationSyncStatusSynced              = "synced"
	federationSyncStatusSyncedWithOverrides = "synced_with_overrides"
	federationSyncStatusDiverged            = "diverged"
	federationSyncStatusStale               = "stale"

	federationProofStatusReady    = "ready"
	federationProofStatusAccepted = "accepted"
	federationProofStatusRejected = "rejected"
	federationProofTypeHandoff    = "sealed_handoff"
)

var (
	errFederationPeerNotFound = errors.New("federation peer not found")
	errFederationInvalidProof = errors.New("federated proof is invalid")
	errFederationRateLimited  = errors.New("federation peer is rate limited")
	errFederationCircuitOpen  = errors.New("federation peer circuit is open")
)

type federationScope struct {
	TenantID    string `json:"tenant_id,omitempty"`
	Environment string `json:"environment,omitempty"`
	Repo        string `json:"repo,omitempty"`
	TrustDomain string `json:"trust_domain,omitempty"`
	Audience    string `json:"audience,omitempty"`
}

type federatedIdentityBinding struct {
	BridgeID           string   `json:"bridge_id"`
	Provider           string   `json:"provider"`
	Issuer             string   `json:"issuer"`
	SubjectPattern     string   `json:"subject_pattern,omitempty"`
	NormalizedIdentity string   `json:"normalized_identity"`
	PrivateKeyImported bool     `json:"private_key_imported"`
	Limitations        []string `json:"limitations,omitempty"`
}

type federationPeerTrustState struct {
	IdentityVerified        bool     `json:"identity_verified"`
	TrustAnchorFingerprints []string `json:"trust_anchor_fingerprints"`
	ChannelMode             string   `json:"channel_mode"`
	FreshnessWindowMinutes  int      `json:"freshness_window_minutes"`
	Limitations             []string `json:"limitations,omitempty"`
}

type federationPeer struct {
	SchemaVersion     string                     `json:"schema_version"`
	PeerID            string                     `json:"peer_id"`
	Organization      string                     `json:"organization"`
	Region            string                     `json:"region,omitempty"`
	Cluster           string                     `json:"cluster,omitempty"`
	TrustDomain       string                     `json:"trust_domain,omitempty"`
	Endpoint          string                     `json:"endpoint,omitempty"`
	PublicKeys        []string                   `json:"public_keys"`
	Capabilities      []string                   `json:"capabilities,omitempty"`
	PolicyRole        string                     `json:"policy_role"`
	Status            string                     `json:"status"`
	LastSeen          time.Time                  `json:"last_seen"`
	AcceptedAudiences []string                   `json:"accepted_audiences,omitempty"`
	DisclosureMode    string                     `json:"disclosure_mode,omitempty"`
	IdentityBindings  []federatedIdentityBinding `json:"identity_bindings,omitempty"`
	MetadataHash      string                     `json:"metadata_hash"`
	MetadataSignature string                     `json:"metadata_signature"`
	TrustState        federationPeerTrustState   `json:"trust_state"`
	Limitations       []string                   `json:"limitations,omitempty"`
}

type federationPeerRequest struct {
	PeerID            string                     `json:"peer_id"`
	Organization      string                     `json:"organization"`
	Region            string                     `json:"region,omitempty"`
	Cluster           string                     `json:"cluster,omitempty"`
	TrustDomain       string                     `json:"trust_domain,omitempty"`
	Endpoint          string                     `json:"endpoint,omitempty"`
	PublicKeys        []string                   `json:"public_keys"`
	Capabilities      []string                   `json:"capabilities,omitempty"`
	PolicyRole        string                     `json:"policy_role,omitempty"`
	AcceptedAudiences []string                   `json:"accepted_audiences,omitempty"`
	DisclosureMode    string                     `json:"disclosure_mode,omitempty"`
	LastSeen          *time.Time                 `json:"last_seen,omitempty"`
	IdentityBindings  []federatedIdentityBinding `json:"identity_bindings,omitempty"`
}

type federationPeersResponse struct {
	SchemaVersion string           `json:"schema_version"`
	Peers         []federationPeer `json:"peers"`
}

type federatedProofRequest struct {
	RequestID      string          `json:"request_id"`
	RequestingPeer string          `json:"requesting_peer"`
	RespondingPeer string          `json:"responding_peer"`
	SubjectType    string          `json:"subject_type"`
	SubjectRef     string          `json:"subject_ref"`
	RequestedScope federationScope `json:"requested_scope"`
	RequestedAt    time.Time       `json:"requested_at"`
}

type federationProofRequestInput struct {
	PeerID         string          `json:"peer_id"`
	PackageID      string          `json:"package_id,omitempty"`
	SubjectType    string          `json:"subject_type,omitempty"`
	SubjectRef     string          `json:"subject_ref,omitempty"`
	RequestedScope federationScope `json:"requested_scope,omitempty"`
}

type federationProofFreshness struct {
	IssuedAt         time.Time `json:"issued_at"`
	ValidUntil       time.Time `json:"valid_until"`
	FreshnessMinutes int       `json:"freshness_minutes"`
	Stale            bool      `json:"stale"`
}

type federatedProofResponse struct {
	SchemaVersion     string                      `json:"schema_version"`
	RequestID         string                      `json:"request_id"`
	RespondingPeer    string                      `json:"responding_peer"`
	ProofType         string                      `json:"proof_type"`
	SealedManifestRef string                      `json:"sealed_manifest_ref"`
	ManifestHash      string                      `json:"manifest_hash"`
	SignatureRefs     []string                    `json:"signature_refs"`
	TimestampRef      string                      `json:"timestamp_ref"`
	TransparencyRef   string                      `json:"transparency_ref"`
	Scope             federationScope             `json:"scope"`
	RedactionProfile  string                      `json:"redaction_profile"`
	Freshness         federationProofFreshness    `json:"freshness"`
	ReadbackRefs      []sealedManifestReadbackRef `json:"readback_refs,omitempty"`
	ForensicRefs      []sealedManifestForensicRef `json:"forensic_refs,omitempty"`
	Status            string                      `json:"status"`
	Limitations       []string                    `json:"limitations,omitempty"`
}

type federatedTrustDecision struct {
	SchemaVersion       string    `json:"schema_version"`
	Decision            string    `json:"decision"`
	DecisionID          string    `json:"decision_id"`
	SubjectRef          string    `json:"subject_ref"`
	PeerID              string    `json:"peer_id"`
	Reasons             []string  `json:"reasons"`
	LocalPolicyVersion  string    `json:"local_policy_version,omitempty"`
	RemotePolicyVersion string    `json:"remote_policy_version,omitempty"`
	ManifestHash        string    `json:"manifest_hash"`
	VerifiedAt          time.Time `json:"verified_at"`
	Limitations         []string  `json:"limitations,omitempty"`
}

type federationProofVerifyRequest struct {
	PeerID             string          `json:"peer_id"`
	BundleBase64       string          `json:"bundle_base64,omitempty"`
	PackageID          string          `json:"package_id,omitempty"`
	RequestedScope     federationScope `json:"requested_scope,omitempty"`
	LocalPolicyVersion string          `json:"local_policy_version,omitempty"`
}

type federationProofExchangeResult struct {
	SchemaVersion string                 `json:"schema_version"`
	Request       federatedProofRequest  `json:"request"`
	Response      federatedProofResponse `json:"response"`
}

type federationProofVerifyResult struct {
	SchemaVersion string                 `json:"schema_version"`
	Response      federatedProofResponse `json:"response"`
	Verification  verificationResult     `json:"verification"`
	Decision      federatedTrustDecision `json:"decision"`
}

type federationProofHistoryItem struct {
	RequestID    string                    `json:"request_id"`
	PeerID       string                    `json:"peer_id"`
	SubjectRef   string                    `json:"subject_ref"`
	ProofType    string                    `json:"proof_type"`
	ManifestHash string                    `json:"manifest_hash"`
	Status       string                    `json:"status"`
	Decision     string                    `json:"decision,omitempty"`
	VerifiedAt   *time.Time                `json:"verified_at,omitempty"`
	Freshness    *federationProofFreshness `json:"freshness,omitempty"`
	Reasons      []string                  `json:"reasons,omitempty"`
}

type federationProofHistoryResponse struct {
	SchemaVersion string                       `json:"schema_version"`
	Items         []federationProofHistoryItem `json:"items"`
}

type policyFederationState struct {
	SchemaVersion       string     `json:"schema_version"`
	LeaderPeer          string     `json:"leader_peer,omitempty"`
	GlobalPolicyRoot    string     `json:"global_policy_root,omitempty"`
	LocalPolicyRoot     string     `json:"local_policy_root,omitempty"`
	EffectivePolicyRoot string     `json:"effective_policy_root,omitempty"`
	SyncStatus          string     `json:"sync_status"`
	InheritedRules      []string   `json:"inherited_rules,omitempty"`
	LocalOverrides      []string   `json:"local_overrides,omitempty"`
	DivergenceReasons   []string   `json:"divergence_reasons,omitempty"`
	LastSyncAt          *time.Time `json:"last_sync_at,omitempty"`
	RemotePolicyVersion string     `json:"remote_policy_version,omitempty"`
}

type federationPolicySyncRequest struct {
	LeaderPeer          string   `json:"leader_peer"`
	GlobalPolicyRoot    string   `json:"global_policy_root"`
	LocalPolicyRoot     string   `json:"local_policy_root,omitempty"`
	InheritedRules      []string `json:"inherited_rules,omitempty"`
	LocalOverrides      []string `json:"local_overrides,omitempty"`
	RemotePolicyVersion string   `json:"remote_policy_version,omitempty"`
}

type federationAnchorRecord struct {
	PeerID             string    `json:"peer_id"`
	AuditRootHash      string    `json:"audit_root_hash"`
	PublishedAt        time.Time `json:"published_at"`
	VerificationStatus string    `json:"verification_status"`
	ProofRef           string    `json:"proof_ref,omitempty"`
	Limitations        []string  `json:"limitations,omitempty"`
}

type federationAnchorsResponse struct {
	SchemaVersion string                   `json:"schema_version"`
	Items         []federationAnchorRecord `json:"items"`
}

type federationGlobalView struct {
	SchemaVersion           string                       `json:"schema_version"`
	Peers                   []federationPeer             `json:"peers"`
	ProofHistory            []federationProofHistoryItem `json:"proof_history"`
	PolicyState             policyFederationState        `json:"policy_state"`
	Anchors                 []federationAnchorRecord     `json:"anchors"`
	TrustHealth             string                       `json:"trust_health"`
	StalePeers              []string                     `json:"stale_peers,omitempty"`
	PolicyDivergence        []string                     `json:"policy_divergence,omitempty"`
	VerifiedArtifactsReused int                          `json:"verified_artifacts_reused"`
	Limitations             []string                     `json:"limitations,omitempty"`
}

type federationPeerEvent struct {
	Peer federationPeer `json:"peer"`
}

type federationProofEvent struct {
	Request  federatedProofRequest  `json:"request,omitempty"`
	Response federatedProofResponse `json:"response,omitempty"`
	Decision federatedTrustDecision `json:"decision,omitempty"`
}

type federationPolicyEvent struct {
	State policyFederationState `json:"state"`
}

func (s server) federationPeersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		peers, err := s.listFederationPeers(ctx)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, federationPeersResponse{
			SchemaVersion: federationPeersSchemaVersion,
			Peers:         peers,
		})
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var request federationPeerRequest
		if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		peer, err := s.registerFederationPeer(ctx, principal, request)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, peer)
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) federationPeerByIDHandler(w http.ResponseWriter, r *http.Request) {
	if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
		return
	}
	peerID := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/federation/peers/"))
	if peerID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": errFederationPeerNotFound.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	peer, err := s.getFederationPeer(ctx, peerID)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, errFederationPeerNotFound) {
			status = http.StatusNotFound
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, peer)
}

func (s server) federationProofRequestHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request federationProofRequestInput
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	if err := s.enforceFederationResilience(ctx, strings.TrimSpace(request.PeerID)); err != nil {
		writeFederationError(w, err)
		return
	}
	result, err := s.createFederatedProofExchange(ctx, principal, request)
	if err != nil {
		writeFederationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, result)
}

func (s server) federationProofVerifyHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request federationProofVerifyRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	if err := s.enforceFederationResilience(ctx, strings.TrimSpace(request.PeerID)); err != nil {
		writeFederationError(w, err)
		return
	}
	result, err := s.verifyFederatedProof(ctx, principal, request)
	if err != nil {
		writeFederationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, result)
}

func (s server) federationProofHistoryHandler(w http.ResponseWriter, r *http.Request) {
	if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, err := s.listFederationProofHistory(ctx)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, federationProofHistoryResponse{
		SchemaVersion: federationProofHistorySchemaVersion,
		Items:         items,
	})
}

func (s server) federationPolicyStateHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		state, err := s.currentFederationPolicyState(ctx)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, state)
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var request federationPolicySyncRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		state, err := s.syncFederationPolicy(ctx, principal, request)
		if err != nil {
			writeFederationError(w, err)
			return
		}
		httpjson.Write(w, http.StatusOK, state)
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) federationAnchorsHandler(w http.ResponseWriter, r *http.Request) {
	if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, err := s.listFederationAnchors(ctx)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, federationAnchorsResponse{
		SchemaVersion: federationAnchorsSchemaVersion,
		Items:         items,
	})
}

func (s server) federationGlobalViewHandler(w http.ResponseWriter, r *http.Request) {
	if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	view, err := s.buildFederationGlobalView(ctx)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, view)
}

func (s server) registerFederationPeer(ctx context.Context, principal auth.Principal, request federationPeerRequest) (federationPeer, error) {
	if strings.TrimSpace(request.PeerID) == "" {
		return federationPeer{}, audit.ErrInvalidEvent
	}
	if len(request.PublicKeys) == 0 {
		return federationPeer{}, audit.ErrInvalidEvent
	}
	seed, err := s.federationSeed()
	if err != nil {
		return federationPeer{}, err
	}
	lastSeen := time.Now().UTC()
	if request.LastSeen != nil && !request.LastSeen.IsZero() {
		lastSeen = request.LastSeen.UTC()
	}
	peer := federationPeer{
		SchemaVersion:     federationPeerSchemaVersion,
		PeerID:            strings.TrimSpace(request.PeerID),
		Organization:      strings.TrimSpace(request.Organization),
		Region:            strings.TrimSpace(request.Region),
		Cluster:           strings.TrimSpace(request.Cluster),
		TrustDomain:       strings.TrimSpace(request.TrustDomain),
		Endpoint:          strings.TrimSpace(request.Endpoint),
		PublicKeys:        sortedStrings(cloneStrings(request.PublicKeys)),
		Capabilities:      sortedStrings(cloneStrings(request.Capabilities)),
		PolicyRole:        firstNonEmpty(strings.TrimSpace(request.PolicyRole), federationPolicyRoleFollower),
		Status:            federationPeerStatusActive,
		LastSeen:          lastSeen,
		AcceptedAudiences: sortedStrings(cloneStrings(request.AcceptedAudiences)),
		DisclosureMode:    firstNonEmpty(strings.TrimSpace(request.DisclosureMode), "sealed_proof_only"),
		IdentityBindings:  normalizeFederatedIdentityBindings(request.IdentityBindings),
		Limitations: []string{
			"Federation peer trust is anchored in locally registered public keys and capabilities; private keys are never imported into the local instance.",
		},
	}
	peer.TrustState = federationPeerTrustState{
		IdentityVerified:        true,
		TrustAnchorFingerprints: federationFingerprints(peer.PublicKeys),
		ChannelMode:             "mutual_identity_metadata",
		FreshnessWindowMinutes:  federationFreshnessWindowMinutes(peer),
		Limitations: []string{
			"Peer acceptance verifies registered public keys, trust anchors, and declared capabilities before remote proof is considered.",
		},
	}
	peer.MetadataHash = federationPeerMetadataHash(peer)
	peer.MetadataSignature = federationSignValue(seed, "peer:"+peer.PeerID, peer.MetadataHash)

	payload, err := canonicalJSON(federationPeerEvent{Peer: peer})
	if err != nil {
		return federationPeer{}, err
	}
	_, err = s.store.Ingest(ctx, audit.Event{
		Component:   federationComponent,
		EventType:   audit.EventTypeFederationPeerRegistered,
		Actor:       principal.Subject,
		Decision:    audit.DecisionAllow,
		TenantID:    "",
		Environment: "",
		Reasons:     []string{"federation peer registered"},
		Federation:  payload,
	})
	if err != nil {
		return federationPeer{}, err
	}
	return peer, nil
}

func (s server) getFederationPeer(ctx context.Context, peerID string) (federationPeer, error) {
	peers, err := s.listFederationPeers(ctx)
	if err != nil {
		return federationPeer{}, err
	}
	for _, peer := range peers {
		if peer.PeerID == peerID {
			return peer, nil
		}
	}
	return federationPeer{}, errFederationPeerNotFound
}

func (s server) listFederationPeers(ctx context.Context) ([]federationPeer, error) {
	events, err := s.store.ListEvents(ctx, audit.EventFilter{Component: federationComponent, EventType: audit.EventTypeFederationPeerRegistered, Limit: 500})
	if err != nil {
		return nil, err
	}
	latest := map[string]federationPeer{}
	for _, event := range orderEventsAscending(events) {
		if event.EventType != audit.EventTypeFederationPeerRegistered || len(event.Federation) == 0 {
			continue
		}
		var payload federationPeerEvent
		if err := json.Unmarshal(event.Federation, &payload); err != nil {
			continue
		}
		peer := payload.Peer
		peer.SchemaVersion = federationPeerSchemaVersion
		peer.Status = federationPeerDerivedStatus(peer)
		latest[peer.PeerID] = peer
	}
	items := make([]federationPeer, 0, len(latest))
	for _, peer := range latest {
		items = append(items, peer)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].PeerID < items[j].PeerID })
	return items, nil
}

func (s server) createFederatedProofExchange(ctx context.Context, principal auth.Principal, request federationProofRequestInput) (federationProofExchangeResult, error) {
	peer, err := s.getFederationPeer(ctx, strings.TrimSpace(request.PeerID))
	if err != nil {
		return federationProofExchangeResult{}, err
	}
	record, err := s.resolveFederationHandoffRecord(ctx, request.PackageID, "")
	if err != nil {
		return federationProofExchangeResult{}, err
	}
	req := federatedProofRequest{
		RequestID:      recommendationID("federation-proof", peer.PeerID, record.PackageID),
		RequestingPeer: peer.PeerID,
		RespondingPeer: federationLocalPeerID,
		SubjectType:    firstNonEmpty(strings.TrimSpace(request.SubjectType), "package"),
		SubjectRef:     firstNonEmpty(strings.TrimSpace(request.SubjectRef), record.PackageID),
		RequestedScope: normalizeFederationScope(request.RequestedScope),
		RequestedAt:    time.Now().UTC(),
	}
	response := buildFederatedProofResponse(peer, record, req.RequestID)
	payload, err := canonicalJSON(federationProofEvent{Request: req, Response: response})
	if err != nil {
		return federationProofExchangeResult{}, err
	}
	_, err = s.store.Ingest(ctx, audit.Event{
		Component:   federationComponent,
		EventType:   audit.EventTypeFederationProofRequested,
		Actor:       principal.Subject,
		Decision:    audit.DecisionAllow,
		TenantID:    record.Manifest.Scope.TenantID,
		Environment: record.Manifest.Scope.Environment,
		Repo:        record.Manifest.Scope.Repo,
		Reasons:     []string{"federated proof prepared from sealed handoff"},
		Federation:  payload,
	})
	if err != nil {
		return federationProofExchangeResult{}, err
	}
	return federationProofExchangeResult{
		SchemaVersion: federationProofExchangeSchemaVersion,
		Request:       req,
		Response:      response,
	}, nil
}

func (s server) verifyFederatedProof(ctx context.Context, principal auth.Principal, request federationProofVerifyRequest) (federationProofVerifyResult, error) {
	peer, err := s.getFederationPeer(ctx, strings.TrimSpace(request.PeerID))
	if err != nil {
		return federationProofVerifyResult{}, err
	}
	record, err := s.resolveFederationHandoffRecord(ctx, request.PackageID, request.BundleBase64)
	if err != nil {
		return federationProofVerifyResult{}, err
	}
	verification := s.verifyStoredHandoff(record)
	response := buildFederatedProofResponse(peer, record, recommendationID("federation-verify", peer.PeerID, record.PackageID))
	policyState, err := s.currentFederationPolicyState(ctx)
	if err != nil {
		return federationProofVerifyResult{}, err
	}
	decision := s.makeFederatedTrustDecision(peer, record, response, verification, normalizeFederationScope(request.RequestedScope), policyState, strings.TrimSpace(request.LocalPolicyVersion))
	payload, err := canonicalJSON(federationProofEvent{
		Response: response,
		Decision: decision,
	})
	if err != nil {
		return federationProofVerifyResult{}, err
	}
	eventDecision := audit.DecisionAllow
	if !strings.HasPrefix(decision.Decision, "accepted") {
		eventDecision = audit.DecisionDeny
	}
	_, err = s.store.Ingest(ctx, audit.Event{
		Component:   federationComponent,
		EventType:   audit.EventTypeFederationProofVerified,
		Actor:       principal.Subject,
		Decision:    eventDecision,
		TenantID:    firstNonEmpty(request.RequestedScope.TenantID, record.Manifest.Scope.TenantID),
		Environment: firstNonEmpty(request.RequestedScope.Environment, record.Manifest.Scope.Environment),
		Repo:        firstNonEmpty(request.RequestedScope.Repo, record.Manifest.Scope.Repo),
		Reasons:     append([]string{"federated proof verified"}, decision.Reasons...),
		Federation:  payload,
	})
	if err != nil {
		return federationProofVerifyResult{}, err
	}
	return federationProofVerifyResult{
		SchemaVersion: federationProofVerifySchemaVersion,
		Response:      response,
		Verification:  verification,
		Decision:      decision,
	}, nil
}

func (s server) listFederationProofHistory(ctx context.Context) ([]federationProofHistoryItem, error) {
	events, err := s.store.ListEvents(ctx, audit.EventFilter{Component: federationComponent, Limit: 500})
	if err != nil {
		return nil, err
	}
	items := []federationProofHistoryItem{}
	for _, event := range orderEventsAscending(events) {
		if len(event.Federation) == 0 {
			continue
		}
		var payload federationProofEvent
		if err := json.Unmarshal(event.Federation, &payload); err != nil {
			continue
		}
		switch event.EventType {
		case audit.EventTypeFederationProofRequested:
			items = append(items, federationProofHistoryItem{
				RequestID:    payload.Request.RequestID,
				PeerID:       payload.Request.RequestingPeer,
				SubjectRef:   payload.Request.SubjectRef,
				ProofType:    payload.Response.ProofType,
				ManifestHash: payload.Response.ManifestHash,
				Status:       payload.Response.Status,
				Freshness:    timePointerFreshness(payload.Response.Freshness),
			})
		case audit.EventTypeFederationProofVerified:
			verifiedAt := payload.Decision.VerifiedAt
			items = append(items, federationProofHistoryItem{
				RequestID:    payload.Response.RequestID,
				PeerID:       payload.Decision.PeerID,
				SubjectRef:   payload.Decision.SubjectRef,
				ProofType:    payload.Response.ProofType,
				ManifestHash: payload.Response.ManifestHash,
				Status:       payload.Response.Status,
				Decision:     payload.Decision.Decision,
				VerifiedAt:   timePointer(verifiedAt),
				Freshness:    timePointerFreshness(payload.Response.Freshness),
				Reasons:      cloneStrings(payload.Decision.Reasons),
			})
		}
	}
	sort.Slice(items, func(i, j int) bool {
		left := federationHistorySortTime(items[i])
		right := federationHistorySortTime(items[j])
		if left.Equal(right) {
			if items[i].PeerID == items[j].PeerID {
				return items[i].RequestID < items[j].RequestID
			}
			return items[i].PeerID < items[j].PeerID
		}
		return left.After(right)
	})
	return items, nil
}

func (s server) currentFederationPolicyState(ctx context.Context) (policyFederationState, error) {
	events, err := s.store.ListEvents(ctx, audit.EventFilter{Component: federationComponent, EventType: audit.EventTypeFederationPolicySynced, Limit: 50})
	if err != nil {
		return policyFederationState{}, err
	}
	var latest policyFederationState
	found := false
	for _, event := range orderEventsAscending(events) {
		if len(event.Federation) == 0 {
			continue
		}
		var payload federationPolicyEvent
		if err := json.Unmarshal(event.Federation, &payload); err != nil {
			continue
		}
		payload.State.SyncStatus = deriveFederationSyncStatus(payload.State)
		latest = payload.State
		found = true
	}
	if found {
		return latest, nil
	}
	return policyFederationState{
		SchemaVersion: federationPolicyStateSchemaVersion,
		SyncStatus:    federationSyncStatusLocalOnly,
		DivergenceReasons: []string{
			"No federation leader is configured yet; local policy remains authoritative in the current scope.",
		},
	}, nil
}

func (s server) syncFederationPolicy(ctx context.Context, principal auth.Principal, request federationPolicySyncRequest) (policyFederationState, error) {
	if strings.TrimSpace(request.GlobalPolicyRoot) == "" {
		return policyFederationState{}, audit.ErrInvalidEvent
	}
	state := policyFederationState{
		SchemaVersion:       federationPolicyStateSchemaVersion,
		LeaderPeer:          strings.TrimSpace(request.LeaderPeer),
		GlobalPolicyRoot:    strings.TrimSpace(request.GlobalPolicyRoot),
		LocalPolicyRoot:     firstNonEmpty(strings.TrimSpace(request.LocalPolicyRoot), strings.TrimSpace(request.GlobalPolicyRoot)),
		InheritedRules:      sortedStrings(cloneStrings(request.InheritedRules)),
		LocalOverrides:      sortedStrings(cloneStrings(request.LocalOverrides)),
		LastSyncAt:          timePointer(time.Now().UTC()),
		RemotePolicyVersion: strings.TrimSpace(request.RemotePolicyVersion),
	}
	state.EffectivePolicyRoot = federationEffectivePolicyRoot(state)
	state.SyncStatus = deriveFederationSyncStatus(state)
	payload, err := canonicalJSON(federationPolicyEvent{State: state})
	if err != nil {
		return policyFederationState{}, err
	}
	eventDecision := audit.DecisionAllow
	if state.SyncStatus == federationSyncStatusDiverged {
		eventDecision = audit.DecisionDeny
	}
	_, err = s.store.Ingest(ctx, audit.Event{
		Component:  federationComponent,
		EventType:  audit.EventTypeFederationPolicySynced,
		Actor:      principal.Subject,
		Decision:   eventDecision,
		Reasons:    append([]string{"federation policy sync evaluated"}, state.DivergenceReasons...),
		Federation: payload,
	})
	if err != nil {
		return policyFederationState{}, err
	}
	return state, nil
}

func (s server) listFederationAnchors(ctx context.Context) ([]federationAnchorRecord, error) {
	events, err := s.store.ListEvents(ctx, audit.EventFilter{Limit: 5000})
	if err != nil {
		return nil, err
	}
	peers, err := s.listFederationPeers(ctx)
	if err != nil {
		return nil, err
	}
	history, err := s.listFederationProofHistory(ctx)
	if err != nil {
		return nil, err
	}
	localAnchor := buildLocalFederationAnchor(events)
	items := []federationAnchorRecord{localAnchor}
	for _, peer := range peers {
		item := federationAnchorRecord{
			PeerID:             peer.PeerID,
			AuditRootHash:      firstNonEmpty(latestPeerManifestHash(history, peer.PeerID), peer.MetadataHash),
			PublishedAt:        peer.LastSeen.UTC(),
			VerificationStatus: "verified",
			ProofRef:           firstNonEmpty(latestPeerManifestHash(history, peer.PeerID), peer.MetadataHash),
			Limitations: []string{
				"Remote anchor status is derived from the latest verified proof or peer metadata hash; raw remote audit history is not replicated locally.",
			},
		}
		if federationPeerDerivedStatus(peer) == federationPeerStatusStale {
			item.VerificationStatus = "stale"
			item.Limitations = append(item.Limitations, "Peer is stale; remote anchor freshness should be revalidated before reuse.")
		}
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].PeerID < items[j].PeerID })
	return items, nil
}

func (s server) buildFederationGlobalView(ctx context.Context) (federationGlobalView, error) {
	peers, err := s.listFederationPeers(ctx)
	if err != nil {
		return federationGlobalView{}, err
	}
	history, err := s.listFederationProofHistory(ctx)
	if err != nil {
		return federationGlobalView{}, err
	}
	state, err := s.currentFederationPolicyState(ctx)
	if err != nil {
		return federationGlobalView{}, err
	}
	anchors, err := s.listFederationAnchors(ctx)
	if err != nil {
		return federationGlobalView{}, err
	}
	stalePeers := []string{}
	verifiedArtifacts := 0
	for _, peer := range peers {
		if federationPeerDerivedStatus(peer) == federationPeerStatusStale {
			stalePeers = append(stalePeers, peer.PeerID)
		}
	}
	for _, item := range history {
		if strings.HasPrefix(item.Decision, "accepted") {
			verifiedArtifacts++
		}
	}
	trustHealth := "healthy"
	limitations := []string{
		"Federation view reflects proof exchange, local trust decisions, and policy sync state; it does not centralize remote canonical audit or evidence truth.",
	}
	if len(stalePeers) > 0 || state.SyncStatus == federationSyncStatusDiverged {
		trustHealth = "degraded"
	}
	return federationGlobalView{
		SchemaVersion:           federationGlobalViewSchemaVersion,
		Peers:                   peers,
		ProofHistory:            history,
		PolicyState:             state,
		Anchors:                 anchors,
		TrustHealth:             trustHealth,
		StalePeers:              stalePeers,
		PolicyDivergence:        cloneStrings(state.DivergenceReasons),
		VerifiedArtifactsReused: verifiedArtifacts,
		Limitations:             limitations,
	}, nil
}

func (s server) resolveFederationHandoffRecord(ctx context.Context, packageID string, bundleBase64 string) (handoffStoredRecord, error) {
	if strings.TrimSpace(packageID) != "" {
		return s.getStoredHandoffRecord(ctx, strings.TrimSpace(packageID))
	}
	if strings.TrimSpace(bundleBase64) == "" {
		return handoffStoredRecord{}, errHandoffInvalidBundle
	}
	bundleBytes, err := base64.StdEncoding.DecodeString(strings.TrimSpace(bundleBase64))
	if err != nil {
		return handoffStoredRecord{}, errHandoffInvalidBundle
	}
	return parseHandoffBundle(bundleBytes)
}

func buildFederatedProofResponse(peer federationPeer, record handoffStoredRecord, requestID string) federatedProofResponse {
	freshnessMinutes := federationFreshnessWindowMinutes(peer)
	issuedAt := record.Timestamp.TimestampedAt.UTC()
	if issuedAt.IsZero() {
		issuedAt = record.Manifest.CreatedAt.UTC()
	}
	validUntil := issuedAt.Add(time.Duration(freshnessMinutes) * time.Minute)
	return federatedProofResponse{
		SchemaVersion:     federationProofResponseSchemaVersion,
		RequestID:         requestID,
		RespondingPeer:    federationLocalPeerID,
		ProofType:         federationProofTypeHandoff,
		SealedManifestRef: fmt.Sprintf("/v1/handoff/%s/manifest", record.PackageID),
		ManifestHash:      record.ManifestHash,
		SignatureRefs: []string{
			fmt.Sprintf("/v1/handoff/%s/manifest", record.PackageID),
			fmt.Sprintf("/v1/handoff/%s/verification", record.PackageID),
		},
		TimestampRef:     fmt.Sprintf("/v1/handoff/%s/verification", record.PackageID),
		TransparencyRef:  fmt.Sprintf("/v1/handoff/%s/verification", record.PackageID),
		Scope:            federationScope{TenantID: record.Manifest.Scope.TenantID, Environment: record.Manifest.Scope.Environment, Repo: record.Manifest.Scope.Repo, Audience: record.Manifest.RedactionProfile.Audience, TrustDomain: peer.TrustDomain},
		RedactionProfile: record.Manifest.RedactionProfile.ProfileVersion + ":" + record.Manifest.RedactionProfile.Audience,
		Freshness: federationProofFreshness{
			IssuedAt:         issuedAt,
			ValidUntil:       validUntil,
			FreshnessMinutes: freshnessMinutes,
			Stale:            time.Now().UTC().After(validUntil),
		},
		ReadbackRefs: cloneSealedManifestReadbackRefs(record.Manifest.ReadbackRefs),
		ForensicRefs: cloneSealedManifestForensicRefs(record.Manifest.ForensicRefs),
		Status:       federationProofStatusReady,
		Limitations: []string{
			"Federated proof response reuses the sealed 9g handoff as the trust artifact and does not replicate the remote raw audit store.",
			"Remote proof remains subject to local trust policy, freshness, scope, and disclosure compatibility checks before acceptance.",
		},
	}
}

func (s server) makeFederatedTrustDecision(peer federationPeer, record handoffStoredRecord, response federatedProofResponse, verification verificationResult, requestedScope federationScope, policyState policyFederationState, localPolicyVersion string) federatedTrustDecision {
	reasons := []string{}
	decision := federationDecisionAccepted
	if federationPeerDerivedStatus(peer) != federationPeerStatusActive {
		decision = federationDecisionRejectedStale
		reasons = append(reasons, "peer is stale or unreachable")
	}
	if !verification.ManifestValid || !verification.ArtifactHashesValid || !verification.SignaturesValid || !verification.TimestampValid || !verification.TransparencyValid {
		decision = federationDecisionRejectedUnverifiable
		reasons = append(reasons, "sealed handoff verification did not fully validate manifest, signature, timestamp, and transparency proof")
	}
	if requestedScope.TenantID != "" && record.Manifest.Scope.TenantID != "" && requestedScope.TenantID != record.Manifest.Scope.TenantID {
		decision = federationDecisionRejectedScopeMismatch
		reasons = append(reasons, "tenant scope does not match the sealed manifest scope")
	}
	if requestedScope.Environment != "" && record.Manifest.Scope.Environment != "" && requestedScope.Environment != record.Manifest.Scope.Environment {
		decision = federationDecisionRejectedScopeMismatch
		reasons = append(reasons, "environment scope does not match the sealed manifest scope")
	}
	if len(peer.AcceptedAudiences) > 0 && !containsString(peer.AcceptedAudiences, record.Manifest.RedactionProfile.Audience) {
		decision = federationDecisionRejectedPolicyConflict
		reasons = append(reasons, "redaction audience is not allowed by the local disclosure policy for this peer")
	}
	if response.Freshness.Stale {
		decision = federationDecisionRejectedStale
		reasons = append(reasons, "proof is outside the locally accepted freshness window")
	}
	if policyState.SyncStatus == federationSyncStatusDiverged {
		decision = federationDecisionRejectedPolicyConflict
		reasons = append(reasons, "local policy federation state is divergent and blocks remote proof reuse until overrides are reconciled")
	}
	if decision == federationDecisionAccepted && len(policyState.LocalOverrides) > 0 {
		decision = federationDecisionAcceptedWithOverrides
		reasons = append(reasons, "local policy overrides remain in force even though the remote proof is valid")
	}
	if len(reasons) == 0 {
		reasons = append(reasons, "remote proof validated against local trust anchors, scope, freshness, and disclosure policy")
	}
	return federatedTrustDecision{
		SchemaVersion:       federationTrustDecisionSchemaVersion,
		DecisionID:          recommendationID("federation-decision", peer.PeerID, record.ManifestHash),
		SubjectRef:          firstNonEmpty(requestedScope.Repo, record.PackageID),
		PeerID:              peer.PeerID,
		Decision:            decision,
		Reasons:             uniqueStrings(reasons),
		LocalPolicyVersion:  firstNonEmpty(localPolicyVersion, policyState.EffectivePolicyRoot, policyState.LocalPolicyRoot),
		RemotePolicyVersion: firstNonEmpty(policyState.RemotePolicyVersion, response.RedactionProfile),
		ManifestHash:        response.ManifestHash,
		VerifiedAt:          time.Now().UTC(),
		Limitations: []string{
			"Remote proof validity never bypasses local policy overrides or disclosure rules.",
			"Federated trust decision remains a local backend-native acceptance step and does not import remote canonical truth.",
		},
	}
}

func buildLocalFederationAnchor(events []audit.StoredEvent) federationAnchorRecord {
	type canonicalAnchorItem struct {
		Timestamp    time.Time `json:"timestamp"`
		ReceivedAt   time.Time `json:"received_at"`
		DecisionHash string    `json:"decision_hash,omitempty"`
		Component    string    `json:"component"`
		EventType    string    `json:"event_type"`
		RequestID    string    `json:"request_id"`
	}
	items := make([]canonicalAnchorItem, 0, len(events))
	publishedAt := time.Time{}
	for _, event := range orderEventsAscending(events) {
		if ts := eventTimestamp(event).UTC(); ts.After(publishedAt) {
			publishedAt = ts
		}
		if received := event.ReceivedAt.UTC(); received.After(publishedAt) {
			publishedAt = received
		}
		items = append(items, canonicalAnchorItem{
			Timestamp:    eventTimestamp(event).UTC(),
			ReceivedAt:   event.ReceivedAt.UTC(),
			DecisionHash: strings.TrimSpace(event.DecisionHash),
			Component:    strings.TrimSpace(event.Component),
			EventType:    strings.TrimSpace(event.EventType),
			RequestID:    strings.TrimSpace(event.RequestID),
		})
	}
	payload, _ := canonicalJSON(items)
	sum := sha256.Sum256(payload)
	if publishedAt.IsZero() {
		publishedAt = time.Unix(0, 0).UTC()
	}
	return federationAnchorRecord{
		PeerID:             federationLocalPeerID,
		AuditRootHash:      "sha256:" + hex.EncodeToString(sum[:]),
		PublishedAt:        publishedAt,
		VerificationStatus: "valid",
		ProofRef:           "/v1/federation/anchors",
		Limitations: []string{
			"Local anchor is a signed summary over canonical audit event lineage, not a replicated remote truth database.",
		},
	}
}

func federationPeerDerivedStatus(peer federationPeer) string {
	lastSeen := peer.LastSeen.UTC()
	if lastSeen.IsZero() {
		return federationPeerStatusStale
	}
	if time.Since(lastSeen) > 6*time.Hour {
		return federationPeerStatusStale
	}
	return federationPeerStatusActive
}

func federationPeerMetadataHash(peer federationPeer) string {
	payload, _ := canonicalJSON(struct {
		PeerID            string                     `json:"peer_id"`
		Organization      string                     `json:"organization"`
		Region            string                     `json:"region,omitempty"`
		Cluster           string                     `json:"cluster,omitempty"`
		TrustDomain       string                     `json:"trust_domain,omitempty"`
		Endpoint          string                     `json:"endpoint,omitempty"`
		PublicKeys        []string                   `json:"public_keys"`
		Capabilities      []string                   `json:"capabilities,omitempty"`
		PolicyRole        string                     `json:"policy_role"`
		AcceptedAudiences []string                   `json:"accepted_audiences,omitempty"`
		DisclosureMode    string                     `json:"disclosure_mode,omitempty"`
		IdentityBindings  []federatedIdentityBinding `json:"identity_bindings,omitempty"`
	}{
		PeerID:            peer.PeerID,
		Organization:      peer.Organization,
		Region:            peer.Region,
		Cluster:           peer.Cluster,
		TrustDomain:       peer.TrustDomain,
		Endpoint:          peer.Endpoint,
		PublicKeys:        peer.PublicKeys,
		Capabilities:      peer.Capabilities,
		PolicyRole:        peer.PolicyRole,
		AcceptedAudiences: peer.AcceptedAudiences,
		DisclosureMode:    peer.DisclosureMode,
		IdentityBindings:  peer.IdentityBindings,
	})
	sum := sha256.Sum256(payload)
	return "sha256:" + hex.EncodeToString(sum[:])
}

func federationFingerprints(keys []string) []string {
	items := make([]string, 0, len(keys))
	for _, key := range keys {
		sum := sha256.Sum256([]byte(strings.TrimSpace(key)))
		items = append(items, "sha256:"+hex.EncodeToString(sum[:8]))
	}
	sort.Strings(items)
	return items
}

func federationFreshnessWindowMinutes(peer federationPeer) int {
	if peer.TrustState.FreshnessWindowMinutes > 0 {
		return peer.TrustState.FreshnessWindowMinutes
	}
	switch peer.PolicyRole {
	case federationPolicyRoleLeader:
		return 180
	case federationPolicyRoleSupplier:
		return 720
	default:
		return 360
	}
}

func deriveFederationSyncStatus(state policyFederationState) string {
	state.DivergenceReasons = federationPolicyDivergenceReasons(state)
	if state.GlobalPolicyRoot == "" {
		return federationSyncStatusLocalOnly
	}
	if len(state.DivergenceReasons) > 0 {
		if len(state.LocalOverrides) > 0 {
			return federationSyncStatusSyncedWithOverrides
		}
		return federationSyncStatusDiverged
	}
	return federationSyncStatusSynced
}

func federationPolicyDivergenceReasons(state policyFederationState) []string {
	reasons := []string{}
	if len(state.LocalOverrides) > 0 {
		reasons = append(reasons, "local compliance overrides remain active and are preserved over the inherited global policy root")
	}
	if state.LeaderPeer == "" {
		reasons = append(reasons, "no federation leader is configured")
	}
	if state.LocalPolicyRoot != "" && state.GlobalPolicyRoot != "" && state.LocalPolicyRoot != state.GlobalPolicyRoot {
		reasons = append(reasons, "local policy root differs from the inherited global root")
	}
	return uniqueStrings(reasons)
}

func federationEffectivePolicyRoot(state policyFederationState) string {
	payload, _ := canonicalJSON(struct {
		Global    string   `json:"global"`
		Local     string   `json:"local"`
		Overrides []string `json:"overrides,omitempty"`
	}{
		Global:    state.GlobalPolicyRoot,
		Local:     state.LocalPolicyRoot,
		Overrides: sortedStrings(cloneStrings(state.LocalOverrides)),
	})
	sum := sha256.Sum256(payload)
	return "sha256:" + hex.EncodeToString(sum[:])
}

func latestPeerManifestHash(items []federationProofHistoryItem, peerID string) string {
	for _, item := range items {
		if item.PeerID == peerID && item.ManifestHash != "" {
			return item.ManifestHash
		}
	}
	return ""
}

func normalizeFederationScope(scope federationScope) federationScope {
	scope.TenantID = strings.TrimSpace(scope.TenantID)
	scope.Environment = strings.TrimSpace(scope.Environment)
	scope.Repo = strings.TrimSpace(scope.Repo)
	scope.TrustDomain = strings.TrimSpace(scope.TrustDomain)
	scope.Audience = strings.TrimSpace(scope.Audience)
	return scope
}

func normalizeFederatedIdentityBindings(bindings []federatedIdentityBinding) []federatedIdentityBinding {
	items := append([]federatedIdentityBinding(nil), bindings...)
	for index := range items {
		items[index].BridgeID = firstNonEmpty(strings.TrimSpace(items[index].BridgeID), recommendationID("federation-identity", items[index].Provider, items[index].NormalizedIdentity))
		items[index].PrivateKeyImported = false
		items[index].Limitations = uniqueStrings(append(cloneStrings(items[index].Limitations),
			"Federated identity bridge normalizes issuer and subject claims without importing private signing keys across trust domains.",
		))
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Provider == items[j].Provider {
			return items[i].NormalizedIdentity < items[j].NormalizedIdentity
		}
		return items[i].Provider < items[j].Provider
	})
	return items
}

func cloneSealedManifestReadbackRefs(values []sealedManifestReadbackRef) []sealedManifestReadbackRef {
	items := append([]sealedManifestReadbackRef(nil), values...)
	return normalizeSealedManifestReadbackRefs(items)
}

func cloneSealedManifestForensicRefs(values []sealedManifestForensicRef) []sealedManifestForensicRef {
	items := append([]sealedManifestForensicRef(nil), values...)
	return normalizeSealedManifestForensicRefs(items)
}

func federationSeedEnv() string {
	return strings.TrimSpace(os.Getenv("CHANGELOCK_FEDERATION_SIGNING_SEED"))
}

func (s server) federationSeed() (string, error) {
	if seed := federationSeedEnv(); seed != "" {
		return seed, nil
	}
	return s.handoffSeed()
}

func federationKeypair(seedSource, role string) (ed25519.PublicKey, ed25519.PrivateKey) {
	sum := sha256.Sum256([]byte("changelock-federation|" + role + "|" + seedSource))
	privateKey := ed25519.NewKeyFromSeed(sum[:])
	publicKey := privateKey.Public().(ed25519.PublicKey)
	return publicKey, privateKey
}

func federationSignValue(seedSource, role, value string) string {
	publicKey, privateKey := federationKeypair(seedSource, role)
	signature := ed25519.Sign(privateKey, []byte(value))
	return base64.StdEncoding.EncodeToString(append(publicKey, signature...))
}

func writeFederationError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, errFederationPeerNotFound):
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	case errors.Is(err, errHandoffInvalidBundle), errors.Is(err, errHandoffNotFound), errors.Is(err, errFederationInvalidProof):
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	case errors.Is(err, errFederationRateLimited):
		httpjson.Write(w, http.StatusTooManyRequests, map[string]string{"error": err.Error()})
	case errors.Is(err, errFederationCircuitOpen):
		httpjson.Write(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
	default:
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
}

func timePointerFreshness(value federationProofFreshness) *federationProofFreshness {
	copy := value
	return &copy
}

func federationHistorySortTime(item federationProofHistoryItem) time.Time {
	if item.VerifiedAt != nil && !item.VerifiedAt.IsZero() {
		return item.VerifiedAt.UTC()
	}
	if item.Freshness != nil {
		if !item.Freshness.IssuedAt.IsZero() {
			return item.Freshness.IssuedAt.UTC()
		}
		if !item.Freshness.ValidUntil.IsZero() {
			return item.Freshness.ValidUntil.UTC()
		}
	}
	return time.Unix(0, 0).UTC()
}

func (s server) buildFederationRecommendations(ctx context.Context, incidents []investigationIncident, filter recommendationFilter) ([]recommendation, error) {
	view, err := s.buildFederationGlobalView(ctx)
	if err != nil {
		return nil, err
	}
	recommendations := []recommendation{}
	if len(view.StalePeers) > 0 {
		peerID := view.StalePeers[0]
		stalePeerSeenAt := federationPeerLastSeen(view.Peers, peerID)
		template := recommendationTemplateCatalog[6]
		recommendations = append(recommendations, recommendation{
			RecommendationID:   recommendationID("federation", peerID, template.TemplateID),
			SourceType:         "federation_signal",
			SourceRef:          federationRecommendationSourceRef("peer", peerID),
			SubjectType:        "peer",
			SubjectRef:         peerID,
			Team:               firstIncidentTenant(incidents),
			Service:            peerID,
			Environment:        filter.event.Environment,
			RecommendationType: template.RecommendationType,
			Title:              fmt.Sprintf("Revalidate stale federation peer %s", peerID),
			Description:        "Peer freshness has fallen outside the currently accepted proof-reuse window.",
			RecommendedAction:  "Refresh peer metadata, verify trust anchors, and rerun proof exchange before accepting any more remote artifacts.",
			Rationale:          "A stale peer undermines freshness guarantees for remote proof reuse and policy inheritance.",
			EvidenceRefs:       []string{peerID},
			PriorityBand:       "TODAY",
			ImpactScore:        74,
			EffortScore:        recommendationEffortScore(template.TemplateID),
			ConfidenceScore:    82,
			ApprovalMode:       template.ApprovalMode,
			Status:             recommendationStatusShown,
			CreatedAt:          recommendationCreatedAt(stalePeerSeenAt),
			ExpiresAt:          recommendationExpiry(stalePeerSeenAt),
			VerificationPlan: []string{
				"Confirm the peer moves back to active status and no longer appears in the stale federation peer list.",
				"Re-run a bounded proof verification and confirm local trust decision returns accepted or accepted_with_local_overrides.",
			},
			ActionTemplate: template,
			Outcome:        recommendationOutcome{Status: recommendationStatusShown},
			AdvisoryOnly:   true,
			Limitations: []string{
				"Federation recommendation is an overlay signal and does not mutate canonical audit, evidence, or remote peer truth.",
			},
		})
	}
	if len(view.PolicyDivergence) > 0 {
		template := recommendationTemplateCatalog[0]
		policyTimestamp := view.PolicyState.LastSyncAt
		recommendations = append(recommendations, recommendation{
			RecommendationID:   recommendationID("federation-policy", view.PolicyState.LeaderPeer, template.TemplateID),
			SourceType:         "federation_signal",
			SourceRef:          federationRecommendationSourceRef("policy", firstNonEmpty(view.PolicyState.LeaderPeer, "local")),
			SubjectType:        "policy_federation",
			SubjectRef:         firstNonEmpty(view.PolicyState.LeaderPeer, "local"),
			Team:               firstIncidentTenant(incidents),
			Service:            firstNonEmpty(view.PolicyState.LeaderPeer, "local"),
			Environment:        filter.event.Environment,
			RecommendationType: template.RecommendationType,
			Title:              "Reconcile federated policy divergence",
			Description:        firstString(view.PolicyDivergence),
			RecommendedAction:  "Review inherited rules against local overrides and clear the divergence before widening remote proof reuse.",
			Rationale:          strings.Join(view.PolicyDivergence, " "),
			EvidenceRefs:       append([]string{}, view.PolicyState.LocalOverrides...),
			PriorityBand:       "NOW",
			ImpactScore:        78,
			EffortScore:        recommendationEffortScore(template.TemplateID),
			ConfidenceScore:    84,
			ApprovalMode:       template.ApprovalMode,
			Status:             recommendationStatusShown,
			CreatedAt:          recommendationCreatedAt(policyTimestamp),
			ExpiresAt:          recommendationExpiry(policyTimestamp),
			VerificationPlan: []string{
				"Confirm the federation policy state returns to synced or synced_with_local_overrides without unresolved divergence reasons.",
				"Verify that remote proof acceptance decisions no longer fail because of policy conflict in the same scope.",
			},
			ActionTemplate: template,
			Outcome:        recommendationOutcome{Status: recommendationStatusShown},
			AdvisoryOnly:   true,
			Limitations: []string{
				"Policy federation recommendation reflects local override preservation and remains an advisory workflow overlay.",
			},
		})
	}
	for _, item := range view.ProofHistory {
		if !strings.HasPrefix(item.Decision, "rejected") && item.Decision != federationDecisionRejectedStale {
			continue
		}
		template := recommendationTemplateCatalog[1]
		recommendations = append(recommendations, recommendation{
			RecommendationID:   recommendationID("federation-proof", item.PeerID, template.TemplateID),
			SourceType:         "federation_signal",
			SourceRef:          federationRecommendationSourceRef("proof", item.PeerID),
			SubjectType:        "peer",
			SubjectRef:         item.PeerID,
			Team:               firstIncidentTenant(incidents),
			Service:            item.PeerID,
			Environment:        filter.event.Environment,
			RecommendationType: template.RecommendationType,
			Title:              fmt.Sprintf("Investigate rejected proof from %s", item.PeerID),
			Description:        firstString(item.Reasons),
			RecommendedAction:  "Re-verify sealed handoff lineage, freshness, disclosure policy, and scope mapping before retrying remote proof reuse.",
			Rationale:          strings.Join(item.Reasons, " "),
			EvidenceRefs:       compactStrings(item.ManifestHash, item.PeerID),
			PriorityBand:       "TODAY",
			ImpactScore:        70,
			EffortScore:        recommendationEffortScore(template.TemplateID),
			ConfidenceScore:    80,
			ApprovalMode:       template.ApprovalMode,
			Status:             recommendationStatusShown,
			CreatedAt:          recommendationCreatedAt(item.VerifiedAt, nil, nil),
			ExpiresAt:          recommendationExpiry(item.VerifiedAt),
			VerificationPlan: []string{
				"Retry the same federated proof verification and confirm the trust decision no longer returns rejected or stale.",
				"Confirm the peer, scope, and disclosure policy align with the sealed manifest before reusing the remote artifact.",
			},
			ActionTemplate: template,
			Outcome:        recommendationOutcome{Status: recommendationStatusShown},
			AdvisoryOnly:   true,
			Limitations: []string{
				"Federated proof recommendation remains advisory and does not import remote raw evidence into local truth.",
			},
		})
		break
	}
	return recommendations, nil
}

func federationPeerLastSeen(peers []federationPeer, peerID string) *time.Time {
	for _, peer := range peers {
		if peer.PeerID != peerID || peer.LastSeen.IsZero() {
			continue
		}
		timestamp := peer.LastSeen.UTC()
		return &timestamp
	}
	return nil
}

func federationRecommendationSourceRef(kind, value string) string {
	return strings.TrimSpace(kind) + ":" + strings.TrimSpace(value)
}

func parseFederationRecommendationSourceRef(value string) (string, string) {
	parts := strings.SplitN(strings.TrimSpace(value), ":", 2)
	if len(parts) != 2 {
		return "proof", strings.TrimSpace(value)
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}
