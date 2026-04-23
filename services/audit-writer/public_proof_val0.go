package main

import (
	"net/http"
	"time"

	claimscore "github.com/denisgrosek/changelock/internal/claims"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/signing"
)

const (
	publicProofVal0ClaimRegistrySchema    = "point2.measured_public_proof.val0.claim_registry.v1"
	publicProofVal0RedactionTiersSchema   = "point2.measured_public_proof.val0.redaction_tiers.v1"
	publicProofVal0SigningAuthoritySchema = "point2.measured_public_proof.val0.signing_authority.v1"
	publicProofVal0CompatibilitySchema    = "point2.measured_public_proof.val0.compatibility.v1"
	publicProofVal0ProofsSchema           = "point2.measured_public_proof.val0.proofs.v1"
)

type publicProofVal0ClaimRegistryResponse struct {
	SchemaVersion string                                   `json:"schema_version"`
	GeneratedAt   time.Time                                `json:"generated_at"`
	CurrentState  string                                   `json:"current_state"`
	Model         claimscore.PublicProofClaimRegistryModel `json:"model"`
	RouteRefs     []string                                 `json:"route_refs,omitempty"`
	Limitations   []string                                 `json:"limitations,omitempty"`
}

type publicProofVal0RedactionTiersResponse struct {
	SchemaVersion string                                `json:"schema_version"`
	GeneratedAt   time.Time                             `json:"generated_at"`
	CurrentState  string                                `json:"current_state"`
	Items         []claimscore.PublicProofRedactionTier `json:"items,omitempty"`
	RouteRefs     []string                              `json:"route_refs,omitempty"`
	Limitations   []string                              `json:"limitations,omitempty"`
}

type publicProofVal0SigningAuthorityResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         claimscore.PublicProofSigningAuthorityModel `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type publicProofVal0CompatibilityResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         claimscore.PublicProofCompatibilityBaseline `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type publicProofVal0ProofsResponse struct {
	SchemaVersion         string    `json:"schema_version"`
	GeneratedAt           time.Time `json:"generated_at"`
	CurrentState          string    `json:"current_state"`
	Phase6State           string    `json:"phase6_state"`
	ClaimRegistryState    string    `json:"claim_registry_state"`
	RedactionTierState    string    `json:"redaction_tier_state"`
	SigningAuthorityState string    `json:"signing_authority_state"`
	CompatibilityState    string    `json:"compatibility_state"`
	SurfaceRefs           []string  `json:"surface_refs,omitempty"`
	EvidenceRefs          []string  `json:"evidence_refs,omitempty"`
	DeferredScope         []string  `json:"deferred_scope,omitempty"`
	Limitations           []string  `json:"limitations,omitempty"`
	IntegrationSummary    []string  `json:"integration_summary,omitempty"`
}

func (s server) publicProofVal0ClaimRegistryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicProofVal0ClaimRegistry())
}

func (s server) publicProofVal0RedactionTiersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicProofVal0RedactionTiers())
}

func (s server) publicProofVal0SigningAuthorityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, s.buildPublicProofVal0SigningAuthority())
}

func (s server) publicProofVal0CompatibilityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicProofVal0Compatibility())
}

func (s server) publicProofVal0ProofsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofVal0Proofs(asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildPublicProofVal0ClaimRegistry() publicProofVal0ClaimRegistryResponse {
	model := claimscore.MeasuredPublicProofVal0ClaimRegistryModel()
	return publicProofVal0ClaimRegistryResponse{
		SchemaVersion: publicProofVal0ClaimRegistrySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/public/trust-program/claims-governance",
			"/v1/public/claims/summary",
			"/v1/public/proof-expansion/val0/proofs",
		},
		Limitations: []string{
			"Val 0 claim registry is a discipline foundation and not yet a live signed claim registry with issuance workflows.",
		},
	}
}

func buildPublicProofVal0RedactionTiers() publicProofVal0RedactionTiersResponse {
	items := claimscore.MeasuredPublicProofVal0RedactionTiers()
	return publicProofVal0RedactionTiersResponse{
		SchemaVersion: publicProofVal0RedactionTiersSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofVal0RedactionTierState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-portal",
			"/v1/public/auditor/workflows",
			"/v1/public/proof-expansion/val0/proofs",
		},
		Limitations: []string{
			"Val 0 redaction tiers define projection policy only; later Point 2 waves will bind them to sealed proof artifact issuance.",
		},
	}
}

func (s server) buildPublicProofVal0SigningAuthority() publicProofVal0SigningAuthorityResponse {
	model := claimscore.MeasuredPublicProofVal0SigningAuthority(publicProofVal0ProviderDescriptor(s.signing))
	return publicProofVal0SigningAuthorityResponse{
		SchemaVersion: publicProofVal0SigningAuthoritySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/public/transparency/anchor",
			"/v1/public/trust-program/claims-governance",
			"/v1/public/proof-expansion/val0/proofs",
		},
		Limitations: []string{
			"Val 0 signing authority defines trust-root and rotation discipline before later Point 2 waves add sealed artifact issuance and anchoring.",
		},
	}
}

func buildPublicProofVal0Compatibility() publicProofVal0CompatibilityResponse {
	model := claimscore.MeasuredPublicProofVal0CompatibilityBaseline()
	return publicProofVal0CompatibilityResponse{
		SchemaVersion: publicProofVal0CompatibilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/public/verifier/sdk",
			"/v1/public/specs/proof-verification",
			"/v1/public/proof-expansion/val0/proofs",
		},
		Limitations: []string{
			"Val 0 compatibility defines schema, deprecation, and replay-tolerance policy before later Point 2 waves add full verifier and replay flows.",
		},
	}
}

func (s server) buildPublicProofVal0Proofs(asOf time.Time) (publicProofVal0ProofsResponse, error) {
	phase6, err := buildPhase6Proofs(asOf)
	if err != nil {
		return publicProofVal0ProofsResponse{}, err
	}
	claimRegistry := claimscore.MeasuredPublicProofVal0ClaimRegistryModel()
	redaction := claimscore.MeasuredPublicProofVal0RedactionTiers()
	signingAuthority := claimscore.MeasuredPublicProofVal0SigningAuthority(publicProofVal0ProviderDescriptor(s.signing))
	compatibility := claimscore.MeasuredPublicProofVal0CompatibilityBaseline()

	currentState := claimscore.MeasuredPublicProofVal0StateIncomplete
	if phase6.CurrentState == phase6ProofStateMarketActive {
		currentState = claimscore.EvaluateMeasuredPublicProofVal0State(
			claimRegistry.CurrentState,
			claimscore.EvaluateMeasuredPublicProofVal0RedactionTierState(redaction),
			signingAuthority.CurrentState,
			compatibility.CurrentState,
		)
	}

	return publicProofVal0ProofsResponse{
		SchemaVersion:         publicProofVal0ProofsSchema,
		GeneratedAt:           publicSampleTime(),
		CurrentState:          currentState,
		Phase6State:           phase6.CurrentState,
		ClaimRegistryState:    claimRegistry.CurrentState,
		RedactionTierState:    claimscore.EvaluateMeasuredPublicProofVal0RedactionTierState(redaction),
		SigningAuthorityState: signingAuthority.CurrentState,
		CompatibilityState:    compatibility.CurrentState,
		SurfaceRefs: []string{
			"/v1/public/phase6/proofs",
			"/v1/public/proof-expansion/val0/claim-registry-model",
			"/v1/public/proof-expansion/val0/redaction-tiers",
			"/v1/public/proof-expansion/val0/signing-authority",
			"/v1/public/proof-expansion/val0/compatibility-baseline",
			"/v1/public/proof-expansion/val0/proofs",
		},
		EvidenceRefs: []string{
			"/v1/public/claims/summary",
			"/v1/public/trust-program/claims-governance",
			"/v1/public/verifier/sdk",
			"/v1/public/transparency/anchor",
			"/v1/public/phase6/proofs",
		},
		DeferredScope: []string{
			"point2_vala_sealed_proof_artifacts",
			"point2_valb_transparency_and_verification",
			"point2_valc_public_and_partner_proof_portal",
			"point2_vald_automated_issuance_and_revocation_gate",
			"point2_vale_final_proof_gate",
		},
		Limitations: []string{
			"Val 0 remains a discipline foundation over claim taxonomy, redaction, signing authority, and compatibility; it does not yet issue sealed proof artifacts.",
			"Point 2 cannot advance from Val 0 without the existing Phase 6 public-proof baseline staying verifier-friendly and freshness-bounded.",
		},
		IntegrationSummary: []string{
			"Claim taxonomy, lifecycle state, and claim-registry requirements are now explicit.",
			"Public, partner, and internal projection tiers are now explicit and fail-closed.",
			"Signing authority, trust-root, rotation, and revoked-signer behavior are now explicit.",
			"Schema compatibility, replay tolerance, deprecation, and failure-state discipline are now explicit.",
		},
	}, nil
}

func publicProofVal0ProviderDescriptor(runtime *signingRuntime) signing.ProviderDescriptor {
	if runtime == nil || runtime.runtime == nil {
		return (*signing.Runtime)(nil).DescribeProvider()
	}
	return runtime.runtime.DescribeProvider()
}
