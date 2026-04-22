package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	claimscore "github.com/denisgrosek/changelock/internal/claims"
	"github.com/denisgrosek/changelock/internal/httpjson"
	referencecore "github.com/denisgrosek/changelock/internal/reference"
	verifiercore "github.com/denisgrosek/changelock/internal/verifier"
)

const (
	phase6TransparencyAnchorSchema   = "6.market_transparency_anchor.v1"
	phase6ProofPortalSchema          = "6.market_public_proof_portal.v1"
	phase6TrustBadgeVerifySchema     = "6.market_trust_badge_verification.v1"
	phase6BenchmarkPacksSchema       = "6.market_benchmark_packs.v1"
	phase6ReferenceConformanceSchema = "6.market_reference_conformance.v1"
	phase6VerifierSDKSchema          = "6.market_verifier_sdk.v1"
	phase6ClaimsSummarySchema        = "6.market_claims_summary.v1"
	phase6AuditorWorkflowSchema      = "6.market_auditor_workflow.v1"
	phase6ProofsSchema               = "6.market_phase6_proofs.v1"

	phase6ProofStateIncomplete   = "phase6_incomplete"
	phase6ProofStateSubstantial  = "phase6_substantially_ready"
	phase6ProofStateMarketActive = "phase6_market_verifiability_active"

	phase6PublicScopePublic  = "public"
	phase6PublicScopePartner = "partner"
	phase6PublicScopeAuditor = "auditor"

	transparencyAnchorStateActive = "transparency_anchor_active"
	transparencyAnchorStateStale  = "transparency_anchor_stale"

	phase6PortalStateActive     = "public_proof_portal_active"
	phase6PortalStateStale      = "public_proof_portal_stale"
	phase6PortalStateIncomplete = "public_proof_portal_incomplete"

	phase6BenchmarkStateActive     = "benchmark_kit_active"
	phase6BenchmarkStateStale      = "benchmark_kit_stale"
	phase6BenchmarkStateIncomplete = "benchmark_kit_incomplete"

	phase6VerifierSDKStateActive     = "verifier_sdk_active"
	phase6VerifierSDKStateStale      = "verifier_sdk_stale"
	phase6VerifierSDKStateIncomplete = "verifier_sdk_incomplete"

	phase6ClaimsStateActive     = "claims_governance_active"
	phase6ClaimsStateStale      = "claims_governance_stale"
	phase6ClaimsStateIncomplete = "claims_governance_incomplete"

	phase6AuditorWorkflowStateActive     = "auditor_workflow_active"
	phase6AuditorWorkflowStateStale      = "auditor_workflow_stale"
	phase6AuditorWorkflowStateIncomplete = "auditor_workflow_incomplete"
)

type phase6TransparencyArtifact struct {
	ArtifactID string `json:"artifact_id"`
	URI        string `json:"uri"`
	Hash       string `json:"hash"`
}

type phase6TransparencyAnchorResponse struct {
	SchemaVersion     string                       `json:"schema_version"`
	AnchorID          string                       `json:"anchor_id"`
	CurrentState      string                       `json:"current_state"`
	FreshnessState    string                       `json:"freshness_state"`
	VerifiedAt        time.Time                    `json:"verified_at"`
	ValidUntil        time.Time                    `json:"valid_until"`
	RootHash          string                       `json:"root_hash"`
	Artifacts         []phase6TransparencyArtifact `json:"artifacts,omitempty"`
	VerificationHints []string                     `json:"verification_hints,omitempty"`
	Limitations       []string                     `json:"limitations,omitempty"`
}

type phase6ProofPortalItem struct {
	ProofID         string              `json:"proof_id"`
	DisplayName     string              `json:"display_name"`
	Category        string              `json:"category"`
	CurrentState    string              `json:"current_state"`
	FreshnessState  string              `json:"freshness_state"`
	VisibilityState string              `json:"visibility_state"`
	VerificationURI string              `json:"verification_uri"`
	Claim           claimscore.Decision `json:"claim"`
	Verification    verifiercore.Result `json:"verification"`
	TraceRefs       []string            `json:"trace_refs,omitempty"`
	Limitations     []string            `json:"limitations,omitempty"`
}

type phase6ProofPortalResponse struct {
	SchemaVersion    string                  `json:"schema_version"`
	Scope            string                  `json:"scope"`
	CurrentState     string                  `json:"current_state"`
	FreshnessState   string                  `json:"freshness_state"`
	AnchorRef        string                  `json:"anchor_ref,omitempty"`
	BadgeRef         string                  `json:"badge_ref,omitempty"`
	VerificationFlow []string                `json:"verification_flow,omitempty"`
	Items            []phase6ProofPortalItem `json:"items,omitempty"`
	Limitations      []string                `json:"limitations,omitempty"`
}

type phase6TrustBadgeVerificationResponse struct {
	SchemaVersion   string              `json:"schema_version"`
	Scope           string              `json:"scope"`
	BadgeID         string              `json:"badge_id"`
	CurrentState    string              `json:"current_state"`
	FreshnessState  string              `json:"freshness_state"`
	VerifiedAt      time.Time           `json:"verified_at"`
	ValidUntil      *time.Time          `json:"valid_until,omitempty"`
	VerificationURI string              `json:"verification_uri,omitempty"`
	AnchorRef       string              `json:"anchor_ref,omitempty"`
	PortalRef       string              `json:"portal_ref,omitempty"`
	TraceRefs       []string            `json:"trace_refs,omitempty"`
	Claim           claimscore.Decision `json:"claim"`
	Verification    verifiercore.Result `json:"verification"`
	Boundaries      []string            `json:"boundaries,omitempty"`
	Limitations     []string            `json:"limitations,omitempty"`
}

type phase6BenchmarkPack struct {
	PackID                 string   `json:"pack_id"`
	BenchmarkID            string   `json:"benchmark_id"`
	DisplayName            string   `json:"display_name"`
	CurrentState           string   `json:"current_state"`
	FreshnessState         string   `json:"freshness_state"`
	PublicationStatus      string   `json:"publication_status"`
	MethodologyRef         string   `json:"methodology_ref"`
	MethodologyDigest      string   `json:"methodology_digest"`
	SignatureRef           string   `json:"signature_ref"`
	VerificationURI        string   `json:"verification_uri"`
	ReplayCommand          string   `json:"replay_command"`
	ReferenceArchitectures []string `json:"reference_architectures,omitempty"`
	EnvironmentAssumptions []string `json:"environment_assumptions,omitempty"`
	UnsupportedConditions  []string `json:"unsupported_conditions,omitempty"`
	EvidenceRefs           []string `json:"evidence_refs,omitempty"`
	NotClaimed             []string `json:"not_claimed,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type phase6BenchmarkPacksResponse struct {
	SchemaVersion             string                `json:"schema_version"`
	CurrentState              string                `json:"current_state"`
	FreshnessState            string                `json:"freshness_state"`
	VerifiedAt                time.Time             `json:"verified_at"`
	ValidUntil                time.Time             `json:"valid_until"`
	MethodologyRef            string                `json:"methodology_ref"`
	MethodologyDigest         string                `json:"methodology_digest"`
	CustomerRunKitRefs        []string              `json:"customer_run_kit_refs,omitempty"`
	PublicationDisciplineRefs []string              `json:"publication_discipline_refs,omitempty"`
	Packs                     []phase6BenchmarkPack `json:"packs,omitempty"`
	Limitations               []string              `json:"limitations,omitempty"`
}

type phase6ReferenceComparisonItem struct {
	CheckID      string `json:"check_id"`
	CurrentState string `json:"current_state"`
	Summary      string `json:"summary"`
	EvidenceRef  string `json:"evidence_ref,omitempty"`
}

type phase6ReferenceConformanceResponse struct {
	SchemaVersion     string                          `json:"schema_version"`
	ArchitectureID    string                          `json:"architecture_id"`
	DisplayName       string                          `json:"display_name"`
	SectorProfile     string                          `json:"sector_profile"`
	DeploymentProfile string                          `json:"deployment_profile"`
	CurrentState      string                          `json:"current_state"`
	RequiredPassed    []string                        `json:"required_passed,omitempty"`
	OptionalPassed    []string                        `json:"optional_passed,omitempty"`
	UnsupportedChecks []string                        `json:"unsupported_checks,omitempty"`
	DegradedChecks    []string                        `json:"degraded_checks,omitempty"`
	Deviations        []string                        `json:"deviations,omitempty"`
	Assumptions       []string                        `json:"assumptions,omitempty"`
	MethodologyRef    string                          `json:"methodology_ref,omitempty"`
	BenchmarkRefs     []string                        `json:"benchmark_refs,omitempty"`
	ComparisonItems   []phase6ReferenceComparisonItem `json:"comparison_items,omitempty"`
	ComparisonRefs    []string                        `json:"comparison_refs,omitempty"`
	Limitations       []string                        `json:"limitations,omitempty"`
}

type phase6VerifierExport struct {
	ArtifactType string   `json:"artifact_type"`
	SpecRef      string   `json:"spec_ref"`
	SampleRef    string   `json:"sample_ref,omitempty"`
	SchemaRef    string   `json:"schema_ref,omitempty"`
	Profiles     []string `json:"profiles,omitempty"`
}

type phase6VerifierSDKResponse struct {
	SchemaVersion          string                 `json:"schema_version"`
	SDKVersion             string                 `json:"sdk_version"`
	CurrentState           string                 `json:"current_state"`
	FreshnessState         string                 `json:"freshness_state"`
	VerifiedAt             time.Time              `json:"verified_at"`
	ValidUntil             time.Time              `json:"valid_until"`
	SupportedArtifacts     []string               `json:"supported_artifacts,omitempty"`
	ResultStates           []string               `json:"result_states,omitempty"`
	SupportedProfiles      []string               `json:"supported_profiles,omitempty"`
	SupportedSchemaLines   []string               `json:"supported_schema_lines,omitempty"`
	LanguageBindings       []string               `json:"language_bindings,omitempty"`
	CompatibilityRules     []string               `json:"compatibility_rules,omitempty"`
	DeprecationRules       []string               `json:"deprecation_rules,omitempty"`
	ResultConsistencyRules []string               `json:"result_consistency_rules,omitempty"`
	VerificationCommands   []string               `json:"verification_commands,omitempty"`
	ArtifactExports        []phase6VerifierExport `json:"artifact_exports,omitempty"`
	ReferencePackRef       string                 `json:"reference_pack_ref,omitempty"`
	AuditorRefs            []string               `json:"auditor_refs,omitempty"`
	NoVendorBackend        bool                   `json:"no_vendor_backend"`
	Limitations            []string               `json:"limitations,omitempty"`
}

type phase6ClaimSummaryItem struct {
	ClaimID               string   `json:"claim_id"`
	ClaimClass            string   `json:"claim_class"`
	PublicationClass      string   `json:"publication_class"`
	CurrentState          string   `json:"current_state"`
	FreshnessState        string   `json:"freshness_state"`
	VisibilityState       string   `json:"visibility_state"`
	AllowedScopes         []string `json:"allowed_scopes,omitempty"`
	TraceRefs             []string `json:"trace_refs,omitempty"`
	RequiredPreconditions []string `json:"required_preconditions,omitempty"`
	ReasonCodes           []string `json:"reason_codes,omitempty"`
}

type phase6ClaimsSummaryResponse struct {
	SchemaVersion         string                   `json:"schema_version"`
	Scope                 string                   `json:"scope"`
	CurrentState          string                   `json:"current_state"`
	FreshnessState        string                   `json:"freshness_state"`
	Items                 []phase6ClaimSummaryItem `json:"items,omitempty"`
	PublicationBoundaries []string                 `json:"publication_boundaries,omitempty"`
	Limitations           []string                 `json:"limitations,omitempty"`
}

type phase6AuditorWorkflowStep struct {
	Step          int      `json:"step"`
	Action        string   `json:"action"`
	ExpectedRefs  []string `json:"expected_refs,omitempty"`
	FailureStates []string `json:"failure_states,omitempty"`
}

type phase6AuditorWorkflowResponse struct {
	SchemaVersion   string                      `json:"schema_version"`
	CurrentState    string                      `json:"current_state"`
	FreshnessState  string                      `json:"freshness_state"`
	VerifiedAt      time.Time                   `json:"verified_at"`
	ValidUntil      time.Time                   `json:"valid_until"`
	WorkflowID      string                      `json:"workflow_id"`
	PermittedScopes []string                    `json:"permitted_scopes,omitempty"`
	AccessModel     []string                    `json:"access_model,omitempty"`
	VerifierRefs    []string                    `json:"verifier_refs,omitempty"`
	Steps           []phase6AuditorWorkflowStep `json:"steps,omitempty"`
	ExportRefs      []string                    `json:"export_refs,omitempty"`
	Limitations     []string                    `json:"limitations,omitempty"`
}

type phase6ProofsResponse struct {
	SchemaVersion        string                               `json:"schema_version"`
	CurrentState         string                               `json:"current_state"`
	FreshnessState       string                               `json:"freshness_state"`
	TransparencyAnchor   phase6TransparencyAnchorResponse     `json:"transparency_anchor"`
	ProofPortal          phase6ProofPortalResponse            `json:"proof_portal"`
	TrustBadge           phase6TrustBadgeVerificationResponse `json:"trust_badge"`
	BenchmarkPacks       phase6BenchmarkPacksResponse         `json:"benchmark_packs"`
	ReferenceConformance phase6ReferenceConformanceResponse   `json:"reference_conformance"`
	VerifierSDK          phase6VerifierSDKResponse            `json:"verifier_sdk"`
	ClaimsSummary        phase6ClaimsSummaryResponse          `json:"claims_summary"`
	AuditorWorkflow      phase6AuditorWorkflowResponse        `json:"auditor_workflow"`
	Limitations          []string                             `json:"limitations,omitempty"`
}

func (s server) publicTransparencyAnchorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := buildPhase6TransparencyAnchor(asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofPortalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	scope, err := parsePhase6Scope(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := buildPhase6ProofPortal(scope, asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicTrustBadgeVerificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	scope, err := parsePhase6Scope(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	badgeID := firstNonEmpty(strings.TrimSpace(r.URL.Query().Get("badge_id")), "verification_ready")
	response, ok, err := buildPhase6TrustBadgeVerification(badgeID, scope, asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if !ok {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "badge not found"})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicBenchmarkPacksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase6BenchmarkPacks(asOf))
}

func (s server) publicReferenceConformanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	architectureID := firstNonEmpty(strings.TrimSpace(r.URL.Query().Get("architecture_id")), "runtime-hardened-enterprise-cluster")
	response, ok, err := buildPhase6ReferenceConformance(architectureID, asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if !ok {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "reference architecture not found"})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicVerifierSDKHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := buildPhase6VerifierSDK(asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicClaimsSummaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	scope, err := parsePhase6Scope(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := buildPhase6ClaimsSummary(scope, asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicAuditorWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := buildPhase6AuditorWorkflow(asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicPhase6ProofsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := buildPhase6Proofs(asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildPhase6TransparencyAnchor(asOf time.Time) (phase6TransparencyAnchorResponse, error) {
	artifacts, err := phase6TransparencyArtifacts()
	if err != nil {
		return phase6TransparencyAnchorResponse{}, err
	}
	lines := make([]string, 0, len(artifacts))
	for _, item := range artifacts {
		lines = append(lines, item.ArtifactID+"|"+item.Hash)
	}
	sort.Strings(lines)
	sum := sha256.Sum256([]byte(strings.Join(lines, "\n")))
	verifiedAt := publicSampleTime()
	validUntil := verifiedAt.Add(90 * 24 * time.Hour)
	freshness := phase6FreshnessState(validUntil, asOf)
	currentState := transparencyAnchorStateActive
	if freshness != verifiercore.FreshnessStateFresh {
		currentState = transparencyAnchorStateStale
	}
	return phase6TransparencyAnchorResponse{
		SchemaVersion:  phase6TransparencyAnchorSchema,
		AnchorID:       "phase6_public_transparency_anchor_v1",
		CurrentState:   currentState,
		FreshnessState: freshness,
		VerifiedAt:     verifiedAt,
		ValidUntil:     validUntil,
		RootHash:       "sha256:" + hex.EncodeToString(sum[:]),
		Artifacts:      artifacts,
		VerificationHints: []string{
			"Fetch each listed public artifact, recompute its hash locally, and then recompute the deterministic root hash in lexical artifact order.",
			"Treat the anchor as a bounded integrity checkpoint over verifier-facing artifacts only.",
		},
		Limitations: []string{
			"Transparency anchor intentionally excludes internal tenant, workflow, incident, and unpublished partner state.",
		},
	}, nil
}

func buildPhase6ProofPortal(scope string, asOf time.Time) (phase6ProofPortalResponse, error) {
	anchor, err := buildPhase6TransparencyAnchor(asOf)
	if err != nil {
		return phase6ProofPortalResponse{}, err
	}
	sdk, err := buildPhase6VerifierSDK(asOf)
	if err != nil {
		return phase6ProofPortalResponse{}, err
	}
	benchmarkPacks := buildPhase6BenchmarkPacks(asOf)
	conformance, _, err := buildPhase6ReferenceConformance("runtime-hardened-enterprise-cluster", asOf)
	if err != nil {
		return phase6ProofPortalResponse{}, err
	}
	badge, _, err := buildPhase6TrustBadgeVerification("verification_ready", scope, asOf)
	if err != nil {
		return phase6ProofPortalResponse{}, err
	}
	items := []phase6ProofPortalItem{
		phase6PortalItemFromArtifact(scope, claimscore.Input{
			ClaimID:                  "transparency_anchor_claim",
			ClaimClass:               "verification_claim",
			PublicationClass:         "public_proof_claim",
			Scope:                    scope,
			VerifiedAt:               anchor.VerifiedAt,
			ValidUntil:               anchor.ValidUntil,
			EvidenceRefs:             []string{"/v1/public/transparency/anchor"},
			ProofRefs:                []string{"/v1/public/transparency/anchor"},
			VerifierRefs:             []string{"/v1/public/verifier/sdk"},
			SupportsIndependentCheck: true,
		}, verifiercore.Input{
			ArtifactID:         anchor.AnchorID,
			ArtifactType:       "transparency_anchor",
			SchemaVersion:      anchor.SchemaVersion,
			VerifiedAt:         anchor.VerifiedAt,
			ValidUntil:         anchor.ValidUntil,
			EvidenceRefs:       []string{"/v1/public/transparency/anchor"},
			ExportRefs:         []string{"/v1/public/transparency/anchor"},
			SupportedSchemas:   []string{phase6TransparencyAnchorSchema},
			IntegrityConfirmed: anchor.CurrentState == transparencyAnchorStateActive,
			ChainContinuity:    true,
		}, asOf, "Transparency anchor", "transparency", "/v1/public/transparency/anchor"),
		phase6PortalItemFromArtifact(scope, claimscore.Input{
			ClaimID:                  "verifier_reference_pack_claim",
			ClaimClass:               "verification_claim",
			PublicationClass:         "public_proof_claim",
			Scope:                    scope,
			VerifiedAt:               sdk.VerifiedAt,
			ValidUntil:               sdk.ValidUntil,
			EvidenceRefs:             []string{"/v1/public/verifier/reference-pack"},
			ProofRefs:                []string{"/v1/public/verifier/reference-pack"},
			VerifierRefs:             []string{"/v1/public/verifier/sdk"},
			SupportsIndependentCheck: true,
		}, verifiercore.Input{
			ArtifactID:         "phase6_verifier_reference_pack",
			ArtifactType:       "reference_pack",
			SchemaVersion:      phase6VerifierSDKSchema,
			VerifiedAt:         sdk.VerifiedAt,
			ValidUntil:         sdk.ValidUntil,
			EvidenceRefs:       []string{"/v1/public/verifier/reference-pack", "/v1/public/verifier/sdk"},
			ExportRefs:         []string{"/v1/public/verifier/reference-pack"},
			SupportedSchemas:   []string{phase6VerifierSDKSchema},
			IntegrityConfirmed: sdk.CurrentState == phase6VerifierSDKStateActive,
			ChainContinuity:    true,
		}, asOf, "Verifier reference pack", "verification", "/v1/public/verifier/sdk"),
		phase6PortalItemFromArtifact(scope, claimscore.Input{
			ClaimID:                  "benchmark_pack_claim",
			ClaimClass:               "benchmark_claim",
			PublicationClass:         "benchmark_claim",
			Scope:                    scope,
			VerifiedAt:               benchmarkPacks.VerifiedAt,
			ValidUntil:               benchmarkPacks.ValidUntil,
			EvidenceRefs:             []string{"/v1/public/benchmarks/packs"},
			ProofRefs:                []string{"/v1/public/benchmarks/packs"},
			MethodologyRefs:          []string{benchmarkPacks.MethodologyRef},
			VerifierRefs:             []string{"/v1/public/verifier/sdk"},
			MethodologyRef:           benchmarkPacks.MethodologyRef,
			SupportsIndependentCheck: true,
		}, verifiercore.Input{
			ArtifactID:         "phase6_public_benchmark_packs",
			ArtifactType:       "benchmark_pack",
			SchemaVersion:      phase6BenchmarkPacksSchema,
			VerifiedAt:         benchmarkPacks.VerifiedAt,
			ValidUntil:         benchmarkPacks.ValidUntil,
			EvidenceRefs:       []string{"/v1/public/benchmarks/packs", benchmarkPacks.MethodologyRef},
			ExportRefs:         benchmarkPacks.CustomerRunKitRefs,
			SupportedSchemas:   []string{phase6BenchmarkPacksSchema},
			MethodologyRef:     benchmarkPacks.MethodologyRef,
			IntegrityConfirmed: benchmarkPacks.CurrentState == phase6BenchmarkStateActive,
			ChainContinuity:    true,
		}, asOf, "Benchmark packs", "benchmark", "/v1/public/benchmarks/packs"),
		phase6PortalItemFromArtifact(scope, claimscore.Input{
			ClaimID:                  "reference_conformance_claim",
			ClaimClass:               "conformance_claim",
			PublicationClass:         "conformance_claim",
			Scope:                    scope,
			VerifiedAt:               anchor.VerifiedAt,
			ValidUntil:               benchmarkPacks.ValidUntil,
			EvidenceRefs:             []string{"/v1/public/reference/conformance"},
			ProofRefs:                []string{"/v1/public/reference/conformance"},
			MethodologyRefs:          []string{benchmarkPacks.MethodologyRef},
			VerifierRefs:             []string{"/v1/public/verifier/sdk"},
			MethodologyRef:           benchmarkPacks.MethodologyRef,
			SupportsIndependentCheck: true,
		}, verifiercore.Input{
			ArtifactID:         conformance.ArchitectureID,
			ArtifactType:       "conformance_comparison",
			SchemaVersion:      phase6ReferenceConformanceSchema,
			VerifiedAt:         anchor.VerifiedAt,
			ValidUntil:         benchmarkPacks.ValidUntil,
			EvidenceRefs:       []string{"/v1/public/reference/conformance"},
			ExportRefs:         conformance.ComparisonRefs,
			SupportedSchemas:   []string{phase6ReferenceConformanceSchema},
			MethodologyRef:     benchmarkPacks.MethodologyRef,
			IntegrityConfirmed: conformance.CurrentState == referencecore.StateActive,
			ChainContinuity:    true,
		}, asOf, "Reference conformance", "conformance", "/v1/public/reference/conformance"),
		{
			ProofID:         "verification_ready_badge",
			DisplayName:     "Verification-ready trust badge",
			Category:        "trust_badge",
			CurrentState:    badge.CurrentState,
			FreshnessState:  badge.FreshnessState,
			VisibilityState: badge.Claim.VisibilityState,
			VerificationURI: badge.VerificationURI,
			Claim:           badge.Claim,
			Verification:    badge.Verification,
			TraceRefs:       badge.TraceRefs,
			Limitations:     badge.Limitations,
		},
	}

	if scope == phase6PublicScopePartner || scope == phase6PublicScopeAuditor {
		items = append(items, phase6PortalItemFromArtifact(scope, claimscore.Input{
			ClaimID:                  "partner_exchange_claim",
			ClaimClass:               "verification_claim",
			PublicationClass:         "partner_exchange_claim",
			Scope:                    scope,
			VerifiedAt:               anchor.VerifiedAt,
			ValidUntil:               anchor.ValidUntil,
			EvidenceRefs:             []string{"/v1/public/specs/federation-proof-exchange"},
			ProofRefs:                []string{"/v1/public/specs/federation-proof-exchange"},
			VerifierRefs:             []string{"/v1/public/verifier/sdk"},
			SupportsIndependentCheck: true,
			PartnerVisibleOnly:       true,
		}, verifiercore.Input{
			ArtifactID:         "partner_exchange_profile",
			ArtifactType:       "proof_bundle",
			SchemaVersion:      publicFederationExchangeSpecSchema,
			VerifiedAt:         anchor.VerifiedAt,
			ValidUntil:         anchor.ValidUntil,
			EvidenceRefs:       []string{"/v1/public/specs/federation-proof-exchange", "/v1/public/verifier/sdk"},
			ExportRefs:         []string{"/v1/public/specs/federation-proof-exchange"},
			SupportedSchemas:   []string{publicFederationExchangeSpecSchema},
			IntegrityConfirmed: anchor.CurrentState == transparencyAnchorStateActive,
			ChainContinuity:    true,
		}, asOf, "Partner proof exchange", "partner", "/v1/public/specs/federation-proof-exchange"))
	}
	if scope == phase6PublicScopeAuditor {
		auditorWorkflow, err := buildPhase6AuditorWorkflow(asOf)
		if err != nil {
			return phase6ProofPortalResponse{}, err
		}
		items = append(items, phase6PortalItemFromArtifact(scope, claimscore.Input{
			ClaimID:                  "auditor_workflow_claim",
			ClaimClass:               "auditor_ready_claim",
			PublicationClass:         "auditor_ready_claim",
			Scope:                    scope,
			VerifiedAt:               auditorWorkflow.VerifiedAt,
			ValidUntil:               auditorWorkflow.ValidUntil,
			EvidenceRefs:             []string{"/v1/public/auditor/workflows"},
			ProofRefs:                []string{"/v1/public/auditor/workflows"},
			VerifierRefs:             auditorWorkflow.VerifierRefs,
			SupportsIndependentCheck: true,
			AuditorVisibleOnly:       true,
		}, verifiercore.Input{
			ArtifactID:         auditorWorkflow.WorkflowID,
			ArtifactType:       "proof_bundle",
			SchemaVersion:      phase6AuditorWorkflowSchema,
			VerifiedAt:         auditorWorkflow.VerifiedAt,
			ValidUntil:         auditorWorkflow.ValidUntil,
			EvidenceRefs:       []string{"/v1/public/auditor/workflows"},
			ExportRefs:         auditorWorkflow.ExportRefs,
			SupportedSchemas:   []string{phase6AuditorWorkflowSchema},
			IntegrityConfirmed: auditorWorkflow.CurrentState == phase6AuditorWorkflowStateActive,
			ChainContinuity:    true,
		}, asOf, "Auditor workflow", "auditor", "/v1/public/auditor/workflows"))
	}

	visible := make([]phase6ProofPortalItem, 0, len(items))
	aggregateFreshness := []string{anchor.FreshnessState}
	hasBlocked := false
	for _, item := range items {
		if !claimscore.AllowsScope(item.Claim, scope) {
			continue
		}
		visible = append(visible, item)
		aggregateFreshness = append(aggregateFreshness, item.FreshnessState)
		if item.Claim.CurrentState == claimscore.StateBlocked || item.Verification.CurrentState == verifiercore.StateInvalid || item.Verification.CurrentState == verifiercore.StateUnsupported || item.Verification.CurrentState == verifiercore.StateIncomplete {
			hasBlocked = true
		}
	}
	freshness := phase6AggregateFreshness(aggregateFreshness...)
	currentState := phase6PortalStateActive
	switch {
	case len(visible) == 0 || hasBlocked:
		currentState = phase6PortalStateIncomplete
	case freshness != verifiercore.FreshnessStateFresh:
		currentState = phase6PortalStateStale
	}
	if badge.CurrentState != "active" && currentState == phase6PortalStateActive {
		currentState = phase6PortalStateStale
	}

	return phase6ProofPortalResponse{
		SchemaVersion:  phase6ProofPortalSchema,
		Scope:          scope,
		CurrentState:   currentState,
		FreshnessState: freshness,
		AnchorRef:      "/v1/public/transparency/anchor",
		BadgeRef:       "/v1/public/trust-program/badges/verify?badge_id=verification_ready",
		VerificationFlow: []string{
			"Resolve the transparency anchor and linked verifier-facing artifacts for the chosen scope.",
			"Replay the verifier pack or schema-compatible artifacts locally and preserve exact verifier result states.",
			"Interpret only the bounded claim language, freshness, and limitations attached to each item.",
		},
		Items: visible,
		Limitations: []string{
			"Proof portal remains a bounded verifier-facing projection and does not expose customer-sensitive runtime, incident, workflow, or unpublished partner data.",
		},
	}, nil
}

func buildPhase6TrustBadgeVerification(badgeID, scope string, asOf time.Time) (phase6TrustBadgeVerificationResponse, bool, error) {
	definition, ok := phase6TrustBadgeDefinitionByID(badgeID)
	if !ok {
		return phase6TrustBadgeVerificationResponse{}, false, nil
	}
	mark, markFound := phase6TrustMarkByBadgeID(badgeID)
	if !markFound {
		return phase6TrustBadgeVerificationResponse{}, false, nil
	}

	anchor, err := buildPhase6TransparencyAnchor(asOf)
	if err != nil {
		return phase6TrustBadgeVerificationResponse{}, false, err
	}
	benchmarkPacks := buildPhase6BenchmarkPacks(asOf)
	sdk, err := buildPhase6VerifierSDK(asOf)
	if err != nil {
		return phase6TrustBadgeVerificationResponse{}, false, err
	}
	architectureSet := buildPublicReferenceArchitectures()
	verifiedAt := publicSampleTime()
	validUntil := phase6MinTime(mark.ExpiresAt, &anchor.ValidUntil, &sdk.ValidUntil, &benchmarkPacks.ValidUntil)
	var proofRefs []string
	var verifierRefs []string
	var methodologyRefs []string
	integrityConfirmed := mark.CurrentState == "active"
	switch badgeID {
	case "verification_ready":
		proofRefs = []string{"/v1/public/transparency/anchor", "/v1/public/verifier/reference-pack"}
		verifierRefs = []string{"/v1/public/verifier/sdk"}
		integrityConfirmed = integrityConfirmed && anchor.CurrentState == transparencyAnchorStateActive && sdk.CurrentState == phase6VerifierSDKStateActive
	case "public_benchmark_disciplined":
		proofRefs = []string{"/v1/public/benchmarks/packs"}
		verifierRefs = []string{"/v1/public/verifier/sdk"}
		methodologyRefs = []string{benchmarkPacks.MethodologyRef}
		integrityConfirmed = integrityConfirmed && benchmarkPacks.CurrentState == phase6BenchmarkStateActive
	case "public_architecture_documented":
		proofRefs = []string{"/v1/public/reference-architectures", "/v1/public/reference/conformance"}
		verifierRefs = []string{"/v1/public/verifier/sdk"}
		integrityConfirmed = integrityConfirmed && len(architectureSet.Architectures) > 0
	default:
		proofRefs = definition.EvidenceRequirements
		verifierRefs = []string{"/v1/public/verifier/sdk"}
	}

	claim := claimscore.Evaluate(claimscore.Input{
		ClaimID:                  "badge_" + badgeID,
		ClaimClass:               "trust_badge_claim",
		PublicationClass:         "trust_badge_claim",
		Scope:                    scope,
		VerifiedAt:               verifiedAt,
		ValidUntil:               validUntil,
		EvidenceRefs:             definition.EvidenceRequirements,
		ProofRefs:                proofRefs,
		MethodologyRefs:          methodologyRefs,
		VerifierRefs:             verifierRefs,
		SupportsIndependentCheck: true,
		BadgeState:               firstNonEmpty(mark.CurrentState, "inactive"),
	}, asOf)
	verification := verifiercore.Evaluate(verifiercore.Input{
		ArtifactID:         badgeID,
		ArtifactType:       "trust_badge",
		SchemaVersion:      publicTrustBadgeProgramSchema,
		VerifiedAt:         verifiedAt,
		ValidUntil:         validUntil,
		EvidenceRefs:       definition.EvidenceRequirements,
		ExportRefs:         uniquePhase6Strings(append(proofRefs, verifierRefs...)),
		SupportedSchemas:   []string{publicTrustBadgeProgramSchema},
		IntegrityConfirmed: integrityConfirmed,
		ChainContinuity:    markFound,
		Deprecated:         false,
	}, asOf)

	currentState := mark.CurrentState
	switch {
	case verification.CurrentState == verifiercore.StateExpired || verification.FreshnessState != verifiercore.FreshnessStateFresh || claim.CurrentState == claimscore.StateStale || anchor.CurrentState != transparencyAnchorStateActive:
		currentState = "stale"
	case claim.CurrentState == claimscore.StateBlocked || verification.CurrentState != verifiercore.StateVerified:
		currentState = "bounded_review_required"
	case currentState == "":
		currentState = "badge_unavailable"
	}

	return phase6TrustBadgeVerificationResponse{
		SchemaVersion:   phase6TrustBadgeVerifySchema,
		Scope:           scope,
		BadgeID:         badgeID,
		CurrentState:    currentState,
		FreshnessState:  phase6AggregateFreshness(anchor.FreshnessState, verification.FreshnessState, claim.FreshnessState),
		VerifiedAt:      verifiedAt,
		ValidUntil:      phase6TimePointer(validUntil),
		VerificationURI: firstNonEmpty(mark.VerificationURI, "/v1/public/trust-program/marks"),
		AnchorRef:       "/v1/public/transparency/anchor",
		PortalRef:       "/v1/public/proof-portal",
		TraceRefs:       uniquePhase6Strings(append(append(proofRefs, verifierRefs...), definition.EvidenceRequirements...)),
		Claim:           claim,
		Verification:    verification,
		Boundaries: []string{
			"Trust badge remains a bounded freshness and proof-availability signal rather than an absolute trust or security guarantee.",
			"When underlying proof freshness or verifier readiness expires, the badge must stop rendering as active.",
		},
		Limitations: uniquePhase6Strings(append(definition.Limitations, mark.Limitations...)),
	}, true, nil
}

func buildPhase6BenchmarkPacks(asOf time.Time) phase6BenchmarkPacksResponse {
	methodology := buildPublicBenchmarkMethodology()
	set := buildPublicBenchmarkSet()
	commandHints := phase6BenchmarkCommandHints()
	verifiedAt := publicSampleTime()
	validUntil := verifiedAt.Add(90 * 24 * time.Hour)
	freshness := phase6FreshnessState(validUntil, asOf)
	methodologyDigest := phase6DigestPayload(methodology)
	packs := make([]phase6BenchmarkPack, 0, len(set.Benchmarks))
	hasReady := false
	for _, benchmark := range set.Benchmarks {
		state := "methodology_only"
		switch benchmark.PublicationStatus {
		case "sample_and_reference_pack_ready", "sample_and_spec_ready", "spec_and_contract_ready":
			state = "benchmark_pack_ready"
			hasReady = true
		case "starting_points_only_not_public_claim":
			state = "starting_points_only"
		}
		if freshness != verifiercore.FreshnessStateFresh && state == "benchmark_pack_ready" {
			state = "benchmark_pack_stale"
		}
		payload := canonicalJSONMust(struct {
			Benchmark   publicBenchmarkDefinition          `json:"benchmark"`
			Methodology publicBenchmarkMethodologyResponse `json:"methodology"`
		}{Benchmark: benchmark, Methodology: methodology})
		sum := sha256.Sum256(payload)
		packs = append(packs, phase6BenchmarkPack{
			PackID:                 "phase6_" + benchmark.BenchmarkID,
			BenchmarkID:            benchmark.BenchmarkID,
			DisplayName:            benchmark.DisplayName,
			CurrentState:           state,
			FreshnessState:         freshness,
			PublicationStatus:      benchmark.PublicationStatus,
			MethodologyRef:         set.MethodologyRef,
			MethodologyDigest:      methodologyDigest,
			SignatureRef:           "sha256:" + hex.EncodeToString(sum[:]),
			VerificationURI:        "/v1/public/verifier/sdk",
			ReplayCommand:          commandHints[benchmark.BenchmarkID],
			ReferenceArchitectures: phase6BenchmarkArchitectureRefs(benchmark.BenchmarkID),
			EnvironmentAssumptions: benchmark.EnvironmentAssumptions,
			UnsupportedConditions:  benchmark.NotClaimed,
			EvidenceRefs:           benchmark.EvidenceRefs,
			NotClaimed:             benchmark.NotClaimed,
			Limitations:            benchmark.Limitations,
		})
	}
	currentState := phase6BenchmarkStateIncomplete
	switch {
	case hasReady && freshness == verifiercore.FreshnessStateFresh:
		currentState = phase6BenchmarkStateActive
	case hasReady:
		currentState = phase6BenchmarkStateStale
	}
	return phase6BenchmarkPacksResponse{
		SchemaVersion:      phase6BenchmarkPacksSchema,
		CurrentState:       currentState,
		FreshnessState:     freshness,
		VerifiedAt:         verifiedAt,
		ValidUntil:         validUntil,
		MethodologyRef:     set.MethodologyRef,
		MethodologyDigest:  methodologyDigest,
		CustomerRunKitRefs: []string{"/v1/public/benchmarks/methodology", "/v1/public/benchmarks/packs", "/v1/public/verifier/reference-pack"},
		PublicationDisciplineRefs: []string{
			"/v1/public/benchmarks/methodology",
			"/v1/public/analytics/publication-discipline",
			"/v1/public/case-studies",
		},
		Packs: packs,
		Limitations: []string{
			"Benchmark packs remain reproducible publication baselines with explicit methodology and non-claim discipline.",
		},
	}
}

func buildPhase6ReferenceConformance(architectureID string, asOf time.Time) (phase6ReferenceConformanceResponse, bool, error) {
	architectures := buildPublicReferenceArchitectures()
	architecture, ok := phase6ReferenceArchitectureByID(architectures.Architectures, architectureID)
	if !ok {
		return phase6ReferenceConformanceResponse{}, false, nil
	}
	anchor, err := buildPhase6TransparencyAnchor(asOf)
	if err != nil {
		return phase6ReferenceConformanceResponse{}, false, err
	}
	sdk, err := buildPhase6VerifierSDK(asOf)
	if err != nil {
		return phase6ReferenceConformanceResponse{}, false, err
	}
	benchmarkPacks := buildPhase6BenchmarkPacks(asOf)
	methodology := buildPublicBenchmarkMethodology()
	checks := []referencecore.Check{
		{CheckID: "transparency_anchor", Required: true, State: phase6NormalizeState(anchor.CurrentState), Summary: "public transparency anchor available", EvidenceRef: "/v1/public/transparency/anchor"},
		{CheckID: "verifier_sdk", Required: true, State: phase6NormalizeState(sdk.CurrentState), Summary: "independent verifier SDK baseline available", EvidenceRef: "/v1/public/verifier/sdk"},
		{CheckID: "benchmark_packs", Required: true, State: phase6NormalizeState(benchmarkPacks.CurrentState), Summary: "public benchmark packs and methodology available", EvidenceRef: "/v1/public/benchmarks/packs"},
		{CheckID: "public_conformance_pack", Required: true, State: "active", Summary: "public conformance pack available", EvidenceRef: "/v1/public/conformance-pack"},
		{CheckID: "handoff_public_spec", Required: true, State: "active", Summary: "public handoff spec available", EvidenceRef: "/v1/public/specs/handoff"},
		{CheckID: "proof_verification_public_spec", Required: true, State: "active", Summary: "public proof verification spec available", EvidenceRef: "/v1/public/specs/proof-verification"},
		{CheckID: "validation_public_spec", Required: false, State: "active", Summary: "public validation certificate spec available", EvidenceRef: "/v1/public/specs/validation-certificate"},
		{CheckID: "federation_public_spec", Required: false, State: "active", Summary: "public federation proof exchange spec available", EvidenceRef: "/v1/public/specs/federation-proof-exchange"},
		{CheckID: "formal_certification", Required: false, State: "unsupported", Summary: "formal certification remains explicitly out of scope"},
	}
	evaluation := referencecore.Evaluate(referencecore.Input{
		ArchitectureID: architecture.ArchitectureID,
		Checks:         checks,
	})
	comparison := make([]phase6ReferenceComparisonItem, 0, len(checks))
	for _, check := range checks {
		comparison = append(comparison, phase6ReferenceComparisonItem{
			CheckID:      check.CheckID,
			CurrentState: phase6ComparisonState(check.State),
			Summary:      check.Summary,
			EvidenceRef:  check.EvidenceRef,
		})
	}
	return phase6ReferenceConformanceResponse{
		SchemaVersion:     phase6ReferenceConformanceSchema,
		ArchitectureID:    architecture.ArchitectureID,
		DisplayName:       architecture.DisplayName,
		SectorProfile:     architecture.SectorProfile,
		DeploymentProfile: architecture.DeploymentProfile,
		CurrentState:      evaluation.CurrentState,
		RequiredPassed:    evaluation.RequiredChecksPassed,
		OptionalPassed:    evaluation.OptionalChecksPassed,
		UnsupportedChecks: evaluation.UnsupportedChecks,
		DegradedChecks:    evaluation.DegradedChecks,
		Deviations:        evaluation.Deviations,
		Assumptions:       architecture.Assumptions,
		MethodologyRef:    "/v1/public/benchmarks/methodology",
		BenchmarkRefs:     []string{"/v1/public/benchmarks/packs", "/v1/public/benchmarks/methodology"},
		ComparisonItems:   comparison,
		ComparisonRefs: []string{
			"/v1/public/reference-architectures",
			"/v1/public/benchmarks/packs",
			"/v1/public/transparency/anchor",
			"/v1/public/verifier/sdk",
		},
		Limitations: uniquePhase6Strings(append(append(architecture.KnownLimitations, evaluation.Limitations...), methodology.Limitations...)),
	}, true, nil
}

func buildPhase6VerifierSDK(asOf time.Time) (phase6VerifierSDKResponse, error) {
	referencePack, err := buildPublicVerifierReferencePack()
	if err != nil {
		return phase6VerifierSDKResponse{}, err
	}
	index := buildPublicSchemaIndex()
	program := buildPublicVerifierProgram()
	verifiedAt := publicSampleTime()
	validUntil := verifiedAt.Add(180 * 24 * time.Hour)
	freshness := phase6FreshnessState(validUntil, asOf)
	currentState := phase6VerifierSDKStateActive
	if len(referencePack.ReplayInputs) == 0 || len(index.Schemas) == 0 {
		currentState = phase6VerifierSDKStateIncomplete
	} else if freshness != verifiercore.FreshnessStateFresh {
		currentState = phase6VerifierSDKStateStale
	}
	supportedSchemaLines := make([]string, 0, len(index.Schemas))
	for _, item := range index.Schemas {
		supportedSchemaLines = append(supportedSchemaLines, item.SchemaID)
	}
	exports := []phase6VerifierExport{
		{ArtifactType: "handoff_bundle", SpecRef: "/v1/public/specs/handoff", SampleRef: "/v1/public/samples/handoff", SchemaRef: "/v1/public/schemas/handoff", Profiles: []string{"minimal_verifier", "full_verifier", "auditor"}},
		{ArtifactType: "proof_verification_case", SpecRef: "/v1/public/specs/proof-verification", SampleRef: "/v1/public/samples/proof-verification", SchemaRef: "/v1/public/schemas/proof-verification", Profiles: []string{"partner_verifier", "full_verifier"}},
		{ArtifactType: "validation_certificate", SpecRef: "/v1/public/specs/validation-certificate", SampleRef: "/v1/public/samples/validation-certificate", SchemaRef: "/v1/public/schemas/validation-certificate", Profiles: []string{"full_verifier", "auditor"}},
		{ArtifactType: "federation_exchange_case", SpecRef: "/v1/public/specs/federation-proof-exchange", SampleRef: "/v1/public/samples/federation-proof-exchange", SchemaRef: "/v1/public/schemas/federation-proof-exchange", Profiles: []string{"partner_verifier", "full_verifier"}},
		{ArtifactType: "benchmark_pack", SpecRef: "/v1/public/benchmarks/methodology", SampleRef: "/v1/public/benchmarks/packs", SchemaRef: "/v1/public/benchmarks/packs", Profiles: []string{"full_verifier", "auditor"}},
		{ArtifactType: "transparency_anchor", SpecRef: "/v1/public/transparency/anchor", SampleRef: "/v1/public/transparency/anchor", SchemaRef: "/v1/public/transparency/anchor", Profiles: []string{"minimal_verifier", "full_verifier", "auditor", "partner_verifier"}},
	}
	return phase6VerifierSDKResponse{
		SchemaVersion:        phase6VerifierSDKSchema,
		SDKVersion:           "phase6.verifier_sdk.v2",
		CurrentState:         currentState,
		FreshnessState:       freshness,
		VerifiedAt:           verifiedAt,
		ValidUntil:           validUntil,
		SupportedArtifacts:   []string{"handoff_bundle", "proof_verification_case", "validation_certificate", "federation_exchange_case", "benchmark_pack", "transparency_anchor"},
		ResultStates:         []string{verifiercore.StateVerified, verifiercore.StateInvalid, verifiercore.StateExpired, verifiercore.StateUnsupported, verifiercore.StateIncomplete},
		SupportedProfiles:    []string{"minimal_verifier", "full_verifier", "auditor", "partner_verifier"},
		SupportedSchemaLines: uniquePhase6Strings(supportedSchemaLines),
		LanguageBindings:     []string{"go_reference", "json_schema_http"},
		CompatibilityRules:   append([]string{}, program.VersionCompatibility...),
		DeprecationRules: []string{
			"Breaking semantic changes require a new public schema version and refreshed conformance testing.",
			"Deprecated schema lines remain verifier-visible until their announced replacement line is published and the compatibility window expires.",
		},
		ResultConsistencyRules: []string{
			"The same artifact must map to the same bounded result classes: verified, invalid, expired, unsupported_schema, or incomplete_evidence.",
			"Freshness, schema compatibility, and local admissibility remain distinct and must not be collapsed into a single optimistic verdict.",
		},
		VerificationCommands: []string{
			"Download /v1/public/verifier/reference-pack and replay each input locally.",
			"Use /v1/public/schemas and the public spec exports to validate schema compatibility before trusting replay results.",
			"Preserve exact verifier states and limitations instead of upgrading them into stronger trust claims.",
		},
		ArtifactExports:  exports,
		ReferencePackRef: "/v1/public/verifier/reference-pack",
		AuditorRefs:      []string{"/v1/public/verifier/reference-pack", "/v1/public/verifier/offline-guide", "/v1/public/auditor/workflows"},
		NoVendorBackend:  true,
		Limitations: []string{
			"Verifier SDK completion remains bounded to published public artifacts and does not require a vendor-only verification backend.",
			"SDK output remains a verifier result and does not replace local admissibility, legal review, or customer-specific policy acceptance.",
		},
	}, nil
}

func buildPhase6ClaimsSummary(scope string, asOf time.Time) (phase6ClaimsSummaryResponse, error) {
	catalog, err := buildPhase6ClaimsCatalog(scope, asOf)
	if err != nil {
		return phase6ClaimsSummaryResponse{}, err
	}
	items := make([]phase6ClaimSummaryItem, 0, len(catalog))
	blocked := false
	freshnessStates := make([]string, 0, len(catalog))
	for _, decision := range catalog {
		if !claimscore.AllowsScope(decision, scope) {
			continue
		}
		items = append(items, phase6ClaimSummaryItem{
			ClaimID:               decision.ClaimID,
			ClaimClass:            decision.ClaimClass,
			PublicationClass:      decision.PublicationClass,
			CurrentState:          decision.CurrentState,
			FreshnessState:        decision.FreshnessState,
			VisibilityState:       decision.VisibilityState,
			AllowedScopes:         decision.AllowedScopes,
			TraceRefs:             decision.TraceRefs,
			RequiredPreconditions: decision.RequiredPreconditions,
			ReasonCodes:           decision.ReasonCodes,
		})
		freshnessStates = append(freshnessStates, decision.FreshnessState)
		if decision.CurrentState == claimscore.StateBlocked {
			blocked = true
		}
	}
	freshness := phase6AggregateFreshness(freshnessStates...)
	currentState := phase6ClaimsStateActive
	switch {
	case blocked:
		currentState = phase6ClaimsStateIncomplete
	case freshness != verifiercore.FreshnessStateFresh:
		currentState = phase6ClaimsStateStale
	}
	return phase6ClaimsSummaryResponse{
		SchemaVersion:  phase6ClaimsSummarySchema,
		Scope:          scope,
		CurrentState:   currentState,
		FreshnessState: freshness,
		Items:          items,
		PublicationBoundaries: []string{
			"Public, partner, and auditor scopes remain distinct and publication is limited by the allowed scope set on each claim.",
			"Stale or incomplete claims remain non-ready and cannot silently stay publication-ready.",
		},
		Limitations: []string{
			"Claims summary remains a bounded publication-control surface and not a substitute for the full legal or commercial review path.",
		},
	}, nil
}

func buildPhase6AuditorWorkflow(asOf time.Time) (phase6AuditorWorkflowResponse, error) {
	sdk, err := buildPhase6VerifierSDK(asOf)
	if err != nil {
		return phase6AuditorWorkflowResponse{}, err
	}
	anchor, err := buildPhase6TransparencyAnchor(asOf)
	if err != nil {
		return phase6AuditorWorkflowResponse{}, err
	}
	verifiedAt := publicSampleTime()
	validUntil := phase6MinTime(&sdk.ValidUntil, &anchor.ValidUntil)
	freshness := phase6AggregateFreshness(sdk.FreshnessState, anchor.FreshnessState)
	currentState := phase6AuditorWorkflowStateActive
	switch {
	case sdk.CurrentState == phase6VerifierSDKStateIncomplete:
		currentState = phase6AuditorWorkflowStateIncomplete
	case freshness != verifiercore.FreshnessStateFresh:
		currentState = phase6AuditorWorkflowStateStale
	}
	return phase6AuditorWorkflowResponse{
		SchemaVersion:   phase6AuditorWorkflowSchema,
		CurrentState:    currentState,
		FreshnessState:  freshness,
		VerifiedAt:      verifiedAt,
		ValidUntil:      validUntil,
		WorkflowID:      "phase6_public_auditor_verification_v1",
		PermittedScopes: []string{phase6PublicScopeAuditor},
		AccessModel: []string{
			"Auditor workflow is bounded to published public or auditor-safe references.",
			"Auditor exports remain permissioned and do not escalate into internal tenant evidence.",
		},
		VerifierRefs: []string{
			"/v1/public/verifier/sdk",
			"/v1/public/verifier/reference-pack",
			"/v1/public/verifier/offline-guide",
		},
		Steps: []phase6AuditorWorkflowStep{
			{Step: 1, Action: "Fetch the transparency anchor, verifier SDK guidance, and reference pack.", ExpectedRefs: []string{"/v1/public/transparency/anchor", "/v1/public/verifier/sdk", "/v1/public/verifier/reference-pack"}},
			{Step: 2, Action: "Replay the public verifier inputs locally and preserve exact result states.", ExpectedRefs: []string{"/v1/public/conformance-pack", "/v1/public/schemas"}, FailureStates: []string{verifiercore.StateInvalid, verifiercore.StateExpired, verifiercore.StateUnsupported, verifiercore.StateIncomplete}},
			{Step: 3, Action: "Record bounded findings with freshness, scope, and limitation context.", ExpectedRefs: []string{"/v1/public/claims/summary", "/v1/public/reference/conformance"}},
		},
		ExportRefs: []string{
			"/v1/public/verifier/reference-pack",
			"/v1/public/conformance-pack",
			"/v1/public/schemas",
			"/v1/public/claims/summary",
			"/v1/public/reference/conformance",
		},
		Limitations: []string{
			"Auditor workflow remains bounded to public or auditor-safe artifacts and does not expose internal tenant evidence, unpublished partner state, or customer-sensitive runtime detail.",
		},
	}, nil
}

func buildPhase6Proofs(asOf time.Time) (phase6ProofsResponse, error) {
	anchor, err := buildPhase6TransparencyAnchor(asOf)
	if err != nil {
		return phase6ProofsResponse{}, err
	}
	portal, err := buildPhase6ProofPortal(phase6PublicScopePublic, asOf)
	if err != nil {
		return phase6ProofsResponse{}, err
	}
	badge, _, err := buildPhase6TrustBadgeVerification("verification_ready", phase6PublicScopePublic, asOf)
	if err != nil {
		return phase6ProofsResponse{}, err
	}
	benchmarkPacks := buildPhase6BenchmarkPacks(asOf)
	conformance, _, err := buildPhase6ReferenceConformance("runtime-hardened-enterprise-cluster", asOf)
	if err != nil {
		return phase6ProofsResponse{}, err
	}
	sdk, err := buildPhase6VerifierSDK(asOf)
	if err != nil {
		return phase6ProofsResponse{}, err
	}
	claimsSummary, err := buildPhase6ClaimsSummary(phase6PublicScopePublic, asOf)
	if err != nil {
		return phase6ProofsResponse{}, err
	}
	auditorWorkflow, err := buildPhase6AuditorWorkflow(asOf)
	if err != nil {
		return phase6ProofsResponse{}, err
	}
	freshness := phase6AggregateFreshness(anchor.FreshnessState, portal.FreshnessState, badge.FreshnessState, benchmarkPacks.FreshnessState, sdk.FreshnessState, claimsSummary.FreshnessState, auditorWorkflow.FreshnessState)
	currentState := phase6FinalProofState(anchor.CurrentState, portal.CurrentState, badge.CurrentState, benchmarkPacks.CurrentState, conformance.CurrentState, sdk.CurrentState, claimsSummary.CurrentState, auditorWorkflow.CurrentState)
	return phase6ProofsResponse{
		SchemaVersion:        phase6ProofsSchema,
		CurrentState:         currentState,
		FreshnessState:       freshness,
		TransparencyAnchor:   anchor,
		ProofPortal:          portal,
		TrustBadge:           badge,
		BenchmarkPacks:       benchmarkPacks,
		ReferenceConformance: conformance,
		VerifierSDK:          sdk,
		ClaimsSummary:        claimsSummary,
		AuditorWorkflow:      auditorWorkflow,
		Limitations: []string{
			"Phase 6 proofs gate remains a bounded public-verifiability summary and does not claim certification, regulator approval, or customer-specific runtime disclosure.",
		},
	}, nil
}

func buildPhase6ClaimsCatalog(scope string, asOf time.Time) ([]claimscore.Decision, error) {
	anchor, err := buildPhase6TransparencyAnchor(asOf)
	if err != nil {
		return nil, err
	}
	benchmarkPacks := buildPhase6BenchmarkPacks(asOf)
	sdk, err := buildPhase6VerifierSDK(asOf)
	if err != nil {
		return nil, err
	}
	conformance, _, err := buildPhase6ReferenceConformance("runtime-hardened-enterprise-cluster", asOf)
	if err != nil {
		return nil, err
	}
	badge, _, err := buildPhase6TrustBadgeVerification("verification_ready", scope, asOf)
	if err != nil {
		return nil, err
	}
	decisions := []claimscore.Decision{
		claimscore.Evaluate(claimscore.Input{
			ClaimID:                  "verification_ready_claim",
			ClaimClass:               "verification_claim",
			PublicationClass:         "verification_claim",
			Scope:                    scope,
			VerifiedAt:               anchor.VerifiedAt,
			ValidUntil:               anchor.ValidUntil,
			EvidenceRefs:             []string{"/v1/public/transparency/anchor"},
			ProofRefs:                []string{"/v1/public/verifier/reference-pack"},
			VerifierRefs:             []string{"/v1/public/verifier/sdk"},
			SupportsIndependentCheck: true,
		}, asOf),
		claimscore.Evaluate(claimscore.Input{
			ClaimID:                  "benchmark_methodology_claim",
			ClaimClass:               "benchmark_claim",
			PublicationClass:         "benchmark_claim",
			Scope:                    scope,
			VerifiedAt:               benchmarkPacks.VerifiedAt,
			ValidUntil:               benchmarkPacks.ValidUntil,
			EvidenceRefs:             []string{"/v1/public/benchmarks/packs"},
			MethodologyRefs:          []string{benchmarkPacks.MethodologyRef},
			ProofRefs:                []string{"/v1/public/benchmarks/packs"},
			VerifierRefs:             []string{"/v1/public/verifier/sdk"},
			MethodologyRef:           benchmarkPacks.MethodologyRef,
			SupportsIndependentCheck: true,
		}, asOf),
		claimscore.Evaluate(claimscore.Input{
			ClaimID:                  "conformance_claim",
			ClaimClass:               "conformance_claim",
			PublicationClass:         "conformance_claim",
			Scope:                    scope,
			VerifiedAt:               anchor.VerifiedAt,
			ValidUntil:               benchmarkPacks.ValidUntil,
			EvidenceRefs:             []string{"/v1/public/reference/conformance"},
			ProofRefs:                []string{"/v1/public/reference/conformance"},
			MethodologyRefs:          []string{benchmarkPacks.MethodologyRef},
			VerifierRefs:             []string{"/v1/public/verifier/sdk"},
			MethodologyRef:           benchmarkPacks.MethodologyRef,
			SupportsIndependentCheck: conformance.CurrentState != referencecore.StateIncomplete && sdk.CurrentState == phase6VerifierSDKStateActive,
		}, asOf),
		badge.Claim,
		claimscore.Evaluate(claimscore.Input{
			ClaimID:                  "partner_exchange_claim",
			ClaimClass:               "verification_claim",
			PublicationClass:         "partner_exchange_claim",
			Scope:                    scope,
			VerifiedAt:               anchor.VerifiedAt,
			ValidUntil:               anchor.ValidUntil,
			EvidenceRefs:             []string{"/v1/public/specs/federation-proof-exchange"},
			ProofRefs:                []string{"/v1/public/specs/federation-proof-exchange"},
			VerifierRefs:             []string{"/v1/public/verifier/sdk"},
			SupportsIndependentCheck: true,
			PartnerVisibleOnly:       true,
		}, asOf),
		claimscore.Evaluate(claimscore.Input{
			ClaimID:                  "auditor_ready_claim",
			ClaimClass:               "auditor_ready_claim",
			PublicationClass:         "auditor_ready_claim",
			Scope:                    scope,
			VerifiedAt:               sdk.VerifiedAt,
			ValidUntil:               sdk.ValidUntil,
			EvidenceRefs:             []string{"/v1/public/auditor/workflows"},
			ProofRefs:                []string{"/v1/public/auditor/workflows"},
			VerifierRefs:             []string{"/v1/public/verifier/reference-pack", "/v1/public/verifier/sdk"},
			SupportsIndependentCheck: true,
			AuditorVisibleOnly:       true,
		}, asOf),
	}
	return decisions, nil
}

func phase6TransparencyArtifacts() ([]phase6TransparencyArtifact, error) {
	referencePack, err := buildPublicVerifierReferencePack()
	if err != nil {
		return nil, err
	}
	items := []struct {
		id      string
		uri     string
		payload any
	}{
		{id: "schemas", uri: "/v1/public/schemas", payload: buildPublicSchemaIndex()},
		{id: "conformance_pack", uri: "/v1/public/conformance-pack", payload: buildPublicConformancePack()},
		{id: "verifier_reference_pack", uri: "/v1/public/verifier/reference-pack", payload: referencePack},
		{id: "benchmark_methodology", uri: "/v1/public/benchmarks/methodology", payload: buildPublicBenchmarkMethodology()},
		{id: "trust_badge_program", uri: "/v1/public/trust-program/badges", payload: buildPublicTrustBadgeProgram()},
		{id: "claims_governance", uri: "/v1/public/trust-program/claims-governance", payload: buildPublicClaimsGovernance()},
	}
	artifacts := make([]phase6TransparencyArtifact, 0, len(items))
	for _, item := range items {
		payload, err := canonicalJSON(item.payload)
		if err != nil {
			return nil, err
		}
		sum := sha256.Sum256(payload)
		artifacts = append(artifacts, phase6TransparencyArtifact{
			ArtifactID: item.id,
			URI:        item.uri,
			Hash:       "sha256:" + hex.EncodeToString(sum[:]),
		})
	}
	sort.Slice(artifacts, func(i, j int) bool { return artifacts[i].ArtifactID < artifacts[j].ArtifactID })
	return artifacts, nil
}

func phase6PortalItemFromArtifact(scope string, claimInput claimscore.Input, verificationInput verifiercore.Input, asOf time.Time, displayName, category, verificationURI string) phase6ProofPortalItem {
	claim := claimscore.Evaluate(claimInput, asOf)
	verification := verifiercore.Evaluate(verificationInput, asOf)
	currentState := verification.CurrentState
	switch {
	case claim.CurrentState == claimscore.StateBlocked:
		currentState = "claim_blocked"
	case claim.CurrentState == claimscore.StateStale || verification.CurrentState == verifiercore.StateExpired || verification.FreshnessState != verifiercore.FreshnessStateFresh:
		currentState = "stale"
	case verification.CurrentState == verifiercore.StateVerified && scope == phase6PublicScopeAuditor && category == "auditor":
		currentState = "verified"
	}
	return phase6ProofPortalItem{
		ProofID:         claimInput.ClaimID,
		DisplayName:     displayName,
		Category:        category,
		CurrentState:    currentState,
		FreshnessState:  phase6AggregateFreshness(claim.FreshnessState, verification.FreshnessState),
		VisibilityState: claim.VisibilityState,
		VerificationURI: verificationURI,
		Claim:           claim,
		Verification:    verification,
		TraceRefs:       uniquePhase6Strings(append(claim.TraceRefs, verification.ExportRefs...)),
		Limitations:     uniquePhase6Strings(append(claim.Limitations, verification.Limitations...)),
	}
}

func phase6TrustBadgeDefinitionByID(badgeID string) (publicTrustBadgeDefinition, bool) {
	for _, item := range buildPublicTrustBadgeProgram().BadgeDefinitions {
		if item.BadgeID == badgeID {
			return item, true
		}
	}
	return publicTrustBadgeDefinition{}, false
}

func phase6TrustMarkByBadgeID(badgeID string) (publicTrustMarkStatus, bool) {
	for _, item := range publicTrustMarkCatalog() {
		if item.BadgeID == badgeID {
			return item, true
		}
	}
	return publicTrustMarkStatus{}, false
}

func phase6ReferenceArchitectureByID(items []publicReferenceArchitecture, architectureID string) (publicReferenceArchitecture, bool) {
	for _, item := range items {
		if item.ArchitectureID == architectureID {
			return item, true
		}
	}
	return publicReferenceArchitecture{}, false
}

func phase6NormalizeState(value string) string {
	switch strings.TrimSpace(value) {
	case transparencyAnchorStateActive, phase6BenchmarkStateActive, phase6VerifierSDKStateActive, phase6ClaimsStateActive, phase6AuditorWorkflowStateActive, phase6PortalStateActive:
		return "active"
	case transparencyAnchorStateStale, phase6BenchmarkStateStale, phase6VerifierSDKStateStale, phase6ClaimsStateStale, phase6AuditorWorkflowStateStale, phase6PortalStateStale:
		return "degraded"
	case referencecore.StateActive:
		return "active"
	case referencecore.StatePartial:
		return "degraded"
	case referencecore.StateIncomplete:
		return "partial"
	default:
		return "unsupported"
	}
}

func phase6ComparisonState(value string) string {
	switch strings.TrimSpace(value) {
	case "active", "ready", "verified":
		return "matched"
	case "degraded", "warning", "partial":
		return "degraded"
	case "unsupported":
		return "unsupported"
	default:
		return "deviated"
	}
}

func phase6BenchmarkCommandHints() map[string]string {
	items := map[string]string{}
	for _, family := range phase6FoundationFamilies() {
		items[family.BenchmarkID] = family.CommandHint
	}
	return items
}

type benchmarkFamily struct {
	BenchmarkID string
	CommandHint string
}

func phase6FoundationFamilies() []benchmarkFamily {
	return []benchmarkFamily{
		{BenchmarkID: "deploy_gate_latency", CommandHint: "go test ./services/deploy-gate -run '^$' -bench BenchmarkAdmissionReview -benchmem"},
		{BenchmarkID: "audit_ingest_throughput", CommandHint: "go test ./internal/audit -run '^$' -bench BenchmarkMemoryStoreIngest -benchmem"},
		{BenchmarkID: "handoff_seal_and_verify", CommandHint: "go test ./services/audit-writer -run '^$' -bench BenchmarkAuditWriterHandoffSeal -benchmem"},
		{BenchmarkID: "federation_proof_verification", CommandHint: "go test ./services/audit-writer -run '^$' -bench BenchmarkAuditWriterFederationProofVerify -benchmem"},
		{BenchmarkID: "validation_execution", CommandHint: "go test ./services/audit-writer -run '^$' -bench BenchmarkAuditWriterValidationExecute -benchmem"},
		{BenchmarkID: "runtime_overhead", CommandHint: "go test ./internal/runtime -run '^$' -bench BenchmarkCompare -benchmem"},
		{BenchmarkID: "runtime_response_latency", CommandHint: "go test ./services/audit-writer -run '^$' -bench BenchmarkAuditWriterRuntimeResponsePolicy -benchmem"},
		{BenchmarkID: "degraded_mode_behavior", CommandHint: "go test ./services/audit-writer -run '^$' -bench BenchmarkAuditWriterPhase6DegradedMode -benchmem"},
	}
}

func phase6BenchmarkArchitectureRefs(benchmarkID string) []string {
	switch benchmarkID {
	case "handoff_seal_and_verify":
		return []string{"supplier-federation", "air-gapped-regulated"}
	case "federation_proof_verification":
		return []string{"supplier-federation", "fintech-multi-region"}
	case "runtime_overhead", "runtime_response_latency":
		return []string{"runtime-hardened-enterprise-cluster", "regulated-saas"}
	default:
		return []string{"runtime-hardened-enterprise-cluster"}
	}
}

func parsePhase6Scope(r *http.Request) (string, error) {
	scope := firstNonEmpty(strings.TrimSpace(r.URL.Query().Get("scope")), phase6PublicScopePublic)
	switch scope {
	case phase6PublicScopePublic, phase6PublicScopePartner, phase6PublicScopeAuditor:
		return scope, nil
	default:
		return "", fmt.Errorf("unsupported scope %q", scope)
	}
}

func parsePhase6AsOf(r *http.Request) (time.Time, error) {
	raw := strings.TrimSpace(r.URL.Query().Get("as_of"))
	if raw == "" {
		return time.Now().UTC(), nil
	}
	ts, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid as_of timestamp")
	}
	return ts.UTC(), nil
}

func phase6FreshnessState(validUntil, asOf time.Time) string {
	if validUntil.IsZero() {
		return verifiercore.FreshnessStateFresh
	}
	switch {
	case asOf.After(validUntil):
		return verifiercore.FreshnessStateExpired
	case asOf.After(validUntil.Add(-14 * 24 * time.Hour)):
		return verifiercore.FreshnessStateStale
	default:
		return verifiercore.FreshnessStateFresh
	}
}

func phase6AggregateFreshness(values ...string) string {
	hasStale := false
	for _, value := range values {
		switch strings.TrimSpace(value) {
		case verifiercore.FreshnessStateExpired:
			return verifiercore.FreshnessStateExpired
		case verifiercore.FreshnessStateStale:
			hasStale = true
		}
	}
	if hasStale {
		return verifiercore.FreshnessStateStale
	}
	return verifiercore.FreshnessStateFresh
}

func phase6FinalProofState(anchorState, portalState, badgeState, benchmarkState, conformanceState, sdkState, claimsState, auditorState string) string {
	switch {
	case anchorState != transparencyAnchorStateActive ||
		portalState == phase6PortalStateIncomplete ||
		badgeState == "badge_unavailable" || badgeState == "bounded_review_required" || badgeState == "revoked" ||
		benchmarkState == phase6BenchmarkStateIncomplete ||
		conformanceState == referencecore.StateIncomplete ||
		sdkState == phase6VerifierSDKStateIncomplete ||
		claimsState == phase6ClaimsStateIncomplete ||
		auditorState == phase6AuditorWorkflowStateIncomplete:
		return phase6ProofStateIncomplete
	case portalState != phase6PortalStateActive ||
		badgeState != "active" ||
		benchmarkState != phase6BenchmarkStateActive ||
		conformanceState != referencecore.StateActive ||
		sdkState != phase6VerifierSDKStateActive ||
		claimsState != phase6ClaimsStateActive ||
		auditorState != phase6AuditorWorkflowStateActive:
		return phase6ProofStateSubstantial
	default:
		return phase6ProofStateMarketActive
	}
}

func phase6DigestPayload(value any) string {
	sum := sha256.Sum256(canonicalJSONMust(value))
	return "sha256:" + hex.EncodeToString(sum[:])
}

func phase6MinTime(times ...*time.Time) time.Time {
	var earliest time.Time
	for _, ts := range times {
		if ts == nil || ts.IsZero() {
			continue
		}
		if earliest.IsZero() || ts.Before(earliest) {
			earliest = ts.UTC()
		}
	}
	return earliest
}

func phase6TimePointer(value time.Time) *time.Time {
	if value.IsZero() {
		return nil
	}
	copy := value.UTC()
	return &copy
}

func uniquePhase6Strings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := map[string]struct{}{}
	items := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		items = append(items, value)
	}
	if len(items) == 0 {
		return nil
	}
	return items
}
