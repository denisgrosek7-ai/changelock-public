package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	b2bSupplierOnboardingSchema  = "4b.supplier_onboarding.v1"
	b2bProofAcceptanceSchema     = "4b.sealed_proof_acceptance.v1"
	b2bProofAcceptanceEvalSchema = "4b.sealed_proof_acceptance_eval.v1"
	b2bDisclosureProfilesSchema  = "4b.disclosure_profiles.v1"
	b2bCustomerBundleSchema      = "4b.customer_trust_bundle.v1"
	b2bConsortiumReadinessSchema = "4b.consortium_readiness.v1"
)

type b2bSupplierOnboardingItem struct {
	PeerID                   string   `json:"peer_id"`
	Organization             string   `json:"organization"`
	TrustDomain              string   `json:"trust_domain,omitempty"`
	PolicyRole               string   `json:"policy_role"`
	Status                   string   `json:"status"`
	Capabilities             []string `json:"capabilities,omitempty"`
	AcceptedProofFormats     []string `json:"accepted_proof_formats,omitempty"`
	TrustAnchorFingerprints  []string `json:"trust_anchor_fingerprints,omitempty"`
	DisclosureBoundaries     []string `json:"disclosure_boundaries,omitempty"`
	OnboardingChecklist      []string `json:"onboarding_checklist,omitempty"`
	LocalAdmissibilityPolicy []string `json:"local_admissibility_policy,omitempty"`
	RevocationAndDistrust    []string `json:"revocation_and_distrust,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

type b2bSupplierOnboardingResponse struct {
	SchemaVersion        string                      `json:"schema_version"`
	AcceptedProofFormats []string                    `json:"accepted_proof_formats,omitempty"`
	Items                []b2bSupplierOnboardingItem `json:"items,omitempty"`
	LocalPolicyOverview  []string                    `json:"local_policy_overview,omitempty"`
	Limitations          []string                    `json:"limitations,omitempty"`
}

type b2bSealedProofAcceptanceResponse struct {
	SchemaVersion              string   `json:"schema_version"`
	AcceptedProofFormats       []string `json:"accepted_proof_formats,omitempty"`
	EvaluateEndpoint           string   `json:"evaluate_endpoint"`
	OfflineVerifySupported     bool     `json:"offline_verify_supported"`
	ProofFreshnessChecks       []string `json:"proof_freshness_checks,omitempty"`
	ProvenanceAndSignerChecks  []string `json:"provenance_and_signer_checks,omitempty"`
	AudienceSpecificDisclosure []string `json:"audience_specific_disclosure,omitempty"`
	LocalVerificationNarrative []string `json:"local_verification_narrative,omitempty"`
	RejectionSemantics         []string `json:"rejection_semantics,omitempty"`
	Limitations                []string `json:"limitations,omitempty"`
}

type b2bSealedProofAcceptanceEvaluation struct {
	SchemaVersion          string                   `json:"schema_version"`
	PeerID                 string                   `json:"peer_id"`
	ProofType              string                   `json:"proof_type"`
	OfflineVerifySupported bool                     `json:"offline_verify_supported"`
	LocalVerification      verificationResult       `json:"local_verification"`
	Freshness              federationProofFreshness `json:"freshness"`
	LocalDecision          federatedTrustDecision   `json:"local_decision"`
	AcceptanceNarrative    []string                 `json:"acceptance_narrative,omitempty"`
	Limitations            []string                 `json:"limitations,omitempty"`
}

type b2bDisclosureProfile struct {
	ProfileID                           string   `json:"profile_id"`
	DisplayName                         string   `json:"display_name"`
	DisclosureMode                      string   `json:"disclosure_mode"`
	AudienceClasses                     []string `json:"audience_classes,omitempty"`
	Includes                            []string `json:"includes,omitempty"`
	Excludes                            []string `json:"excludes,omitempty"`
	VerificationWithoutSourceDisclosure []string `json:"verification_without_source_disclosure,omitempty"`
	ExportVariants                      []string `json:"export_variants,omitempty"`
}

type b2bDisclosureProfilesResponse struct {
	SchemaVersion             string                 `json:"schema_version"`
	Profiles                  []b2bDisclosureProfile `json:"profiles,omitempty"`
	SelectiveDisclosurePolicy []string               `json:"selective_disclosure_policy,omitempty"`
	Limitations               []string               `json:"limitations,omitempty"`
}

type b2bCustomerTrustBundleResponse struct {
	SchemaVersion               string                    `json:"schema_version"`
	CurrentState                string                    `json:"current_state"`
	ProfileID                   string                    `json:"profile_id"`
	VerificationStatusSummary   string                    `json:"verification_status_summary"`
	BoundedTrustIndicators      []audit.TrustBadge        `json:"bounded_trust_indicators,omitempty"`
	CustomerSafeNarrative       []string                  `json:"customer_safe_narrative,omitempty"`
	PublicView                  *audit.PublishedTrustView `json:"public_view,omitempty"`
	MachineVerifiablePaths      []string                  `json:"machine_verifiable_paths,omitempty"`
	SealedProofVerificationPath string                    `json:"sealed_proof_verification_path,omitempty"`
	Limitations                 []string                  `json:"limitations,omitempty"`
}

type b2bConsortiumPeerSummary struct {
	TotalPeers        int `json:"total_peers"`
	ActivePeers       int `json:"active_peers"`
	StalePeers        int `json:"stale_peers"`
	SupplierPeers     int `json:"supplier_peers"`
	SharedAnchorCount int `json:"shared_anchor_count"`
	AcceptedArtifacts int `json:"accepted_artifacts"`
}

type b2bConsortiumReadinessResponse struct {
	SchemaVersion                string                   `json:"schema_version"`
	ReadinessState               string                   `json:"readiness_state"`
	PeerSummary                  b2bConsortiumPeerSummary `json:"peer_summary"`
	SharedTrustAnchorReadiness   []string                 `json:"shared_trust_anchor_readiness,omitempty"`
	ProofExchangeGovernanceHints []string                 `json:"proof_exchange_governance_hints,omitempty"`
	CrossOrgFreshnessSemantics   []string                 `json:"cross_org_freshness_semantics,omitempty"`
	SectorDisclosureProfiles     []string                 `json:"sector_disclosure_profiles,omitempty"`
	LocalOverrideAndDistrust     []string                 `json:"local_override_and_distrust,omitempty"`
	PolicyState                  policyFederationState    `json:"policy_state"`
	Limitations                  []string                 `json:"limitations,omitempty"`
}

func (s server) b2bSupplierOnboardingHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildB2BSupplierOnboarding(ctx)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) b2bSealedProofAcceptanceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, buildB2BSealedProofAcceptanceContract())
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var request federationProofVerifyRequest
		if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request.RequestedScope = enforcePrincipalFederationScope(principal, normalizeFederationScope(request.RequestedScope))
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		result, err := s.verifyFederatedProof(ctx, principal, request)
		if err != nil {
			writeFederationError(w, err)
			return
		}
		httpjson.Write(w, http.StatusOK, b2bSealedProofAcceptanceEvaluation{
			SchemaVersion:          b2bProofAcceptanceEvalSchema,
			PeerID:                 result.Decision.PeerID,
			ProofType:              result.Response.ProofType,
			OfflineVerifySupported: true,
			LocalVerification:      result.Verification,
			Freshness:              result.Response.Freshness,
			LocalDecision:          result.Decision,
			AcceptanceNarrative: []string{
				"Partner proof acceptance remains local-policy driven and requires local manifest, signature, timestamp, transparency, scope, freshness, and disclosure checks.",
				"Accepted remote proof does not import remote canonical evidence or bypass local override and distrust posture.",
			},
			Limitations: append([]string{
				"Acceptance evaluation records a bounded local trust decision; it does not onboard or authorize the partner globally.",
			}, uniqueStrings(append(cloneStrings(result.Verification.Limitations), cloneStrings(result.Decision.Limitations)...))...),
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) b2bDisclosureProfilesHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildB2BDisclosureProfiles())
}

func (s server) b2bCustomerBundleHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	scope, err := applyPrincipalTenantToTrustScopeRequest(principal, parseTrustScopeRequestFromQuery(r))
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildB2BCustomerBundle(ctx, scope, strings.TrimSpace(r.URL.Query().Get("package_id")))
	if err != nil {
		s.writeScorecardError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) b2bConsortiumReadinessHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildB2BConsortiumReadiness(ctx)
	if err != nil {
		writeFederationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) buildB2BSupplierOnboarding(ctx context.Context) (b2bSupplierOnboardingResponse, error) {
	peers, err := s.listFederationPeers(ctx)
	if err != nil {
		return b2bSupplierOnboardingResponse{}, err
	}
	items := make([]b2bSupplierOnboardingItem, 0, len(peers))
	for _, peer := range peers {
		item := b2bSupplierOnboardingItem{
			PeerID:                  peer.PeerID,
			Organization:            peer.Organization,
			TrustDomain:             peer.TrustDomain,
			PolicyRole:              peer.PolicyRole,
			Status:                  federationPeerDerivedStatus(peer),
			Capabilities:            cloneStrings(peer.Capabilities),
			AcceptedProofFormats:    acceptedProofFormatsForPeer(peer),
			TrustAnchorFingerprints: cloneStrings(peer.TrustState.TrustAnchorFingerprints),
			DisclosureBoundaries: []string{
				"accepted_audiences=" + strings.Join(cloneStrings(peer.AcceptedAudiences), ","),
				"disclosure_mode=" + firstNonEmpty(peer.DisclosureMode, "sealed_proof_only"),
			},
			OnboardingChecklist: []string{
				"register supplier identity and trust anchors",
				"declare accepted proof capabilities and redaction audiences",
				"verify local admissibility before proof reuse",
				"record distrust and revocation semantics before production acceptance",
			},
			LocalAdmissibilityPolicy: []string{
				"proof acceptance requires local manifest, signer, timestamp, transparency, scope, freshness, and disclosure checks",
				"remote proof validity never bypasses local policy overrides or distrust posture",
			},
			RevocationAndDistrust: []string{
				"stale or unreachable peers are degraded and can be rejected",
				"audience mismatch or policy divergence blocks reuse even if the bundle verifies cryptographically",
			},
			Limitations: cloneStrings(peer.Limitations),
		}
		if len(item.TrustAnchorFingerprints) == 0 {
			item.TrustAnchorFingerprints = federationFingerprints(peer.PublicKeys)
		}
		items = append(items, item)
	}
	return b2bSupplierOnboardingResponse{
		SchemaVersion:        b2bSupplierOnboardingSchema,
		AcceptedProofFormats: []string{federationProofTypeHandoff, "sealed_handoff_bundle"},
		Items:                items,
		LocalPolicyOverview: []string{
			"Supplier onboarding remains a local registration and admissibility process, not an implicit trust grant.",
			"Proof format acceptance and disclosure boundaries are scoped per peer and can be overridden or revoked locally.",
		},
		Limitations: []string{
			"Onboarding surface summarizes current partner registration posture; it does not prove remote operational health or remote evidence completeness.",
		},
	}, nil
}

func buildB2BSealedProofAcceptanceContract() b2bSealedProofAcceptanceResponse {
	return b2bSealedProofAcceptanceResponse{
		SchemaVersion:          b2bProofAcceptanceSchema,
		AcceptedProofFormats:   []string{federationProofTypeHandoff, "sealed_handoff_bundle"},
		EvaluateEndpoint:       "/v1/b2b/sealed-proof/acceptance",
		OfflineVerifySupported: true,
		ProofFreshnessChecks: []string{
			"proof freshness is evaluated against the locally accepted freshness window for the peer",
			"stale proofs can be rejected even if cryptographic verification succeeds",
		},
		ProvenanceAndSignerChecks: []string{
			"manifest hash, artifact hashes, signatures, detached timestamp, and transparency record must validate locally",
			"signer identities are recorded locally and never substituted by remote metadata alone",
		},
		AudienceSpecificDisclosure: []string{
			"accepted_audiences and peer disclosure mode gate remote proof reuse",
			"customer-safe or auditor-safe redaction still remains subject to local disclosure policy",
		},
		LocalVerificationNarrative: []string{
			"partner proof acceptance is a local verification narrative built on 9g handoff verification plus federation policy checks",
			"local policy overrides and distrust posture remain authoritative after cryptographic verification",
		},
		RejectionSemantics: []string{
			federationDecisionRejectedUnverifiable,
			federationDecisionRejectedScopeMismatch,
			federationDecisionRejectedPolicyConflict,
			federationDecisionRejectedStale,
			federationDecisionRejectedUntrustedPeer,
		},
		Limitations: []string{
			"Acceptance contract is bounded to sealed-handoff proof reuse and does not grant broad partner authorization outside the evaluated scope.",
		},
	}
}

func buildB2BDisclosureProfiles() b2bDisclosureProfilesResponse {
	return b2bDisclosureProfilesResponse{
		SchemaVersion: b2bDisclosureProfilesSchema,
		Profiles: []b2bDisclosureProfile{
			{
				ProfileID:       "sealed_proof_only",
				DisplayName:     "Sealed proof only",
				DisclosureMode:  "sealed_proof_only",
				AudienceClasses: []string{"partner", "supplier"},
				Includes:        []string{"manifest identity", "verification status", "freshness metadata", "sanitized readback refs"},
				Excludes:        []string{"raw audit database", "private keys", "internal-only operator notes"},
				VerificationWithoutSourceDisclosure: []string{
					"partner can verify manifest, signatures, timestamp, and transparency without access to internal source systems",
				},
				ExportVariants: []string{"/v1/handoff/{package_id}/download", "/v1/handoff/{package_id}/verification"},
			},
			{
				ProfileID:       incidentAudienceAuditorSafe,
				DisplayName:     "Auditor safe",
				DisclosureMode:  "selective_disclosure",
				AudienceClasses: []string{"auditor", "regulated_partner"},
				Includes:        []string{"audit-safe incident narrative", "verification lineage", "bounded evidence refs"},
				Excludes:        []string{"internal-only notes", "unredacted topology detail"},
				VerificationWithoutSourceDisclosure: []string{
					"auditor-safe redaction preserves verification lineage and bounded evidence narrative",
				},
				ExportVariants: []string{"/v1/incidents/{incident_id}/export?audience=auditor_safe", "/v1/trust/published"},
			},
			{
				ProfileID:       incidentAudienceCustomerSafe,
				DisplayName:     "Customer safe",
				DisclosureMode:  "customer_safe",
				AudienceClasses: []string{"customer"},
				Includes:        []string{"published trust indicators", "customer-safe verification narrative", "machine-verifiable verification path"},
				Excludes:        []string{"sensitive topology", "internal identifiers", "private remediation workflow"},
				VerificationWithoutSourceDisclosure: []string{
					"customer-safe view keeps public trust metrics and machine verification paths without exposing internal source data",
				},
				ExportVariants: []string{"/v1/trust/published", "/v1/audit/exports"},
			},
		},
		SelectiveDisclosurePolicy: []string{
			"selective disclosure is policy-controlled and profile-bound, not a best-effort redaction filter",
			"verification lineage must survive every export profile even when source-system detail is excluded",
		},
		Limitations: []string{
			"Disclosure profiles describe supported bounded export shapes; they do not imply universal per-field redaction for every artifact outside the current surfaces.",
		},
	}
}

func (s server) buildB2BCustomerBundle(ctx context.Context, scope trustScopeRequest, packageID string) (b2bCustomerTrustBundleResponse, error) {
	cfg, err := loadTrustAuditConfigFromEnv()
	if err != nil {
		return b2bCustomerTrustBundleResponse{}, err
	}
	input, err := s.collectTrustScorecardInput(ctx, scope, cfg)
	if err != nil {
		return b2bCustomerTrustBundleResponse{}, err
	}
	card := audit.ComputeTrustScorecard(input)
	badges := audit.BuildTrustBadges(card, input)
	publicView := audit.BuildPublishedTrustView(card, badges, audit.BuildStandardsMapping(card))
	publicBadges := make([]audit.TrustBadge, 0, len(badges))
	for _, badge := range badges {
		if badge.PublicPublishable {
			publicBadges = append(publicBadges, badge)
		}
	}
	sort.Slice(publicBadges, func(i, j int) bool { return publicBadges[i].ID < publicBadges[j].ID })
	state := "customer_safe_ready"
	if publicView == nil {
		state = "customer_safe_disabled"
	}
	verificationPath := ""
	if packageID != "" {
		if _, err := s.getStoredHandoffRecord(ctx, packageID); err == nil {
			verificationPath = "/v1/handoff/" + packageID + "/verification"
		}
	}
	return b2bCustomerTrustBundleResponse{
		SchemaVersion:             b2bCustomerBundleSchema,
		CurrentState:              state,
		ProfileID:                 incidentAudienceCustomerSafe,
		VerificationStatusSummary: customerBundleVerificationSummary(card, publicView != nil),
		BoundedTrustIndicators:    publicBadges,
		CustomerSafeNarrative: []string{
			"Customer trust bundle is a bounded verification view derived from measured trust posture and public-publishable indicators.",
			"It is not a certification claim and does not expose internal audit or remediation detail beyond the selected customer-safe profile.",
		},
		PublicView: publicView,
		MachineVerifiablePaths: uniqueStrings([]string{
			"/v1/trust/published",
			"/v1/audit/exports",
			verificationPath,
		}),
		SealedProofVerificationPath: verificationPath,
		Limitations: []string{
			"Customer bundle remains bounded to published trust and optional sealed-proof verification paths; it does not expose raw internal evidence or partner-private lineage.",
		},
	}, nil
}

func (s server) buildB2BConsortiumReadiness(ctx context.Context) (b2bConsortiumReadinessResponse, error) {
	view, err := s.buildFederationGlobalView(ctx)
	if err != nil {
		return b2bConsortiumReadinessResponse{}, err
	}
	summary := b2bConsortiumPeerSummary{
		TotalPeers:        len(view.Peers),
		StalePeers:        len(view.StalePeers),
		SharedAnchorCount: len(view.Anchors),
		AcceptedArtifacts: view.VerifiedArtifactsReused,
	}
	for _, peer := range view.Peers {
		if federationPeerDerivedStatus(peer) == federationPeerStatusActive {
			summary.ActivePeers++
		}
		if peer.PolicyRole == federationPolicyRoleSupplier {
			summary.SupplierPeers++
		}
	}
	readiness := "partner_local_only"
	if summary.TotalPeers > 0 && len(view.Anchors) > 0 {
		readiness = "consortium_ready_with_local_overrides"
	}
	if len(view.StalePeers) > 0 || view.PolicyState.SyncStatus == federationSyncStatusDiverged {
		readiness = "consortium_degraded"
	}
	return b2bConsortiumReadinessResponse{
		SchemaVersion:  b2bConsortiumReadinessSchema,
		ReadinessState: readiness,
		PeerSummary:    summary,
		SharedTrustAnchorReadiness: []string{
			"shared trust readiness is bounded to registered peers, local anchors, and proof-history reuse counts",
			"local system can exit to local-only mode when anchors stale, peers diverge, or distrust posture changes",
		},
		ProofExchangeGovernanceHints: []string{
			"proof exchange governance remains local-policy-first and preserves disclosure compatibility per peer",
			"cross-org proof reuse is gated by freshness, accepted audience, and local policy divergence state",
		},
		CrossOrgFreshnessSemantics: []string{
			"leader peers default to shorter freshness windows than supplier peers",
			"stale peers remain visible and can degrade consortium readiness without silently changing trust state",
		},
		SectorDisclosureProfiles: []string{"sealed_proof_only", incidentAudienceAuditorSafe, incidentAudienceCustomerSafe},
		LocalOverrideAndDistrust: []string{
			"local overrides remain active even after accepted remote proof decisions",
			"stale, unreachable, or policy-conflicting peers can be distrusted without affecting internal canonical truth",
		},
		PolicyState: view.PolicyState,
		Limitations: []string{
			"Consortium readiness is a bounded preparedness view; it does not imply a shared global log, globally mandatory trust root, or automatic cross-organization policy authority.",
		},
	}, nil
}

func acceptedProofFormatsForPeer(peer federationPeer) []string {
	formats := []string{}
	for _, capability := range peer.Capabilities {
		switch capability {
		case "sealed_handoff", "supplier_proof", "forensics_handoff":
			formats = append(formats, federationProofTypeHandoff)
		}
	}
	if len(formats) == 0 {
		formats = append(formats, federationProofTypeHandoff)
	}
	return uniqueStrings(formats)
}

func enforcePrincipalFederationScope(principal auth.Principal, scope federationScope) federationScope {
	if principal.TenantID != "" && !principal.GlobalScope {
		scope.TenantID = principal.TenantID
	}
	return scope
}

func customerBundleVerificationSummary(card audit.TrustScorecard, publicViewAvailable bool) string {
	switch {
	case !publicViewAvailable:
		return "Customer-safe publication is currently disabled, so only local readiness metadata is available."
	case card.OverallScore >= 85:
		return "Customer-safe trust bundle is available with strong measured posture and machine-verifiable publication paths."
	case card.OverallScore >= 60:
		return "Customer-safe trust bundle is available with partial trust indicators and bounded verification paths."
	default:
		return "Customer-safe trust bundle is available, but measured posture remains partial and should be read together with the included limitations."
	}
}
