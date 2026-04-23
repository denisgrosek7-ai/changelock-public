package claims

import (
	"strings"

	"github.com/denisgrosek/changelock/internal/signing"
)

const (
	MeasuredPublicProofVal0ClaimRegistryStateActive     = "measured_public_proof_val0_claim_registry_active"
	MeasuredPublicProofVal0ClaimRegistryStatePartial    = "measured_public_proof_val0_claim_registry_partial"
	MeasuredPublicProofVal0ClaimRegistryStateIncomplete = "measured_public_proof_val0_claim_registry_incomplete"

	MeasuredPublicProofVal0RedactionTierStateActive     = "measured_public_proof_val0_redaction_tiers_active"
	MeasuredPublicProofVal0RedactionTierStatePartial    = "measured_public_proof_val0_redaction_tiers_partial"
	MeasuredPublicProofVal0RedactionTierStateIncomplete = "measured_public_proof_val0_redaction_tiers_incomplete"

	MeasuredPublicProofVal0SigningAuthorityStateActive     = "measured_public_proof_val0_signing_authority_active"
	MeasuredPublicProofVal0SigningAuthorityStatePartial    = "measured_public_proof_val0_signing_authority_partial"
	MeasuredPublicProofVal0SigningAuthorityStateIncomplete = "measured_public_proof_val0_signing_authority_incomplete"

	MeasuredPublicProofVal0CompatibilityStateActive     = "measured_public_proof_val0_compatibility_active"
	MeasuredPublicProofVal0CompatibilityStatePartial    = "measured_public_proof_val0_compatibility_partial"
	MeasuredPublicProofVal0CompatibilityStateIncomplete = "measured_public_proof_val0_compatibility_incomplete"

	MeasuredPublicProofVal0StateIncomplete  = "measured_public_proof_val0_incomplete"
	MeasuredPublicProofVal0StateSubstantial = "measured_public_proof_val0_substantially_ready"
	MeasuredPublicProofVal0StateActive      = "measured_public_proof_val0_active"

	PublicProofClaimClassPerformance     = "performance_claim"
	PublicProofClaimClassVerification    = "verification_claim"
	PublicProofClaimClassRuntimeBehavior = "runtime_behavior_claim"
	PublicProofClaimClassCompatibility   = "compatibility_claim"
	PublicProofClaimClassReplayability   = "replayability_claim"

	PublicProofStatusProven           = "proven"
	PublicProofStatusPartiallyProven  = "partially_proven"
	PublicProofStatusProofPending     = "proof_pending"
	PublicProofStatusProofFailed      = "proof_failed"
	PublicProofStatusClaimNotReissued = "claim_not_reissued"
	PublicProofStatusRestricted       = "restricted"
	PublicProofStatusSuperseded       = "superseded"
	PublicProofStatusWithdrawn        = "withdrawn"
	PublicProofStatusStale            = "stale"

	RedactionTierInternalFull  = "internal_full"
	RedactionTierPartnerScoped = "partner_scoped"
	RedactionTierPublicSafe    = "public_safe"
)

type PublicProofClaimClass struct {
	ClaimClass           string   `json:"claim_class"`
	Purpose              string   `json:"purpose"`
	RequiredFields       []string `json:"required_fields,omitempty"`
	SupportedScopes      []string `json:"supported_scopes,omitempty"`
	RequiredBoundaries   []string `json:"required_boundaries,omitempty"`
	DisallowedOverclaims []string `json:"disallowed_overclaims,omitempty"`
}

type PublicProofLifecycleState struct {
	Status              string `json:"status"`
	CurrentState        string `json:"current_state"`
	VerifierVisible     bool   `json:"verifier_visible"`
	Terminal            bool   `json:"terminal"`
	Description         string `json:"description"`
	PublicationBehavior string `json:"publication_behavior"`
}

type PublicProofClaimRegistryModel struct {
	CurrentState               string                      `json:"current_state"`
	RequiredClaimFields        []string                    `json:"required_claim_fields,omitempty"`
	ClaimClasses               []PublicProofClaimClass     `json:"claim_classes,omitempty"`
	LifecycleStates            []PublicProofLifecycleState `json:"lifecycle_states,omitempty"`
	MethodologyBoundaryRules   []string                    `json:"methodology_boundary_rules,omitempty"`
	FreshnessRules             []string                    `json:"freshness_rules,omitempty"`
	RevocationRules            []string                    `json:"revocation_rules,omitempty"`
	SupersessionRules          []string                    `json:"supersession_rules,omitempty"`
	RequiredProjectionPolicies []string                    `json:"required_projection_policies,omitempty"`
	Limitations                []string                    `json:"limitations,omitempty"`
}

type PublicProofRedactionTier struct {
	TierID               string   `json:"tier_id"`
	CurrentState         string   `json:"current_state"`
	AllowedScopes        []string `json:"allowed_scopes,omitempty"`
	RemovedFields        []string `json:"removed_fields,omitempty"`
	AggregatedFields     []string `json:"aggregated_fields,omitempty"`
	NeverPublishedFields []string `json:"never_published_fields,omitempty"`
	PortalPolicy         string   `json:"portal_policy"`
	Limitations          []string `json:"limitations,omitempty"`
}

type PublicProofTrustRoot struct {
	TrustRootID           string   `json:"trust_root_id"`
	CurrentState          string   `json:"current_state"`
	SignerIdentity        string   `json:"signer_identity"`
	KeyVersion            string   `json:"key_version"`
	IssuanceWindow        string   `json:"issuance_window"`
	RotationPolicy        string   `json:"rotation_policy"`
	RevocationBehavior    string   `json:"revocation_behavior"`
	TimestampingPolicy    string   `json:"timestamping_policy"`
	SupportedClaimClasses []string `json:"supported_claim_classes,omitempty"`
	Limitations           []string `json:"limitations,omitempty"`
}

type PublicProofSigningAuthorityModel struct {
	CurrentState           string                     `json:"current_state"`
	Provider               signing.ProviderDescriptor `json:"provider"`
	TrustRoots             []PublicProofTrustRoot     `json:"trust_roots,omitempty"`
	KeyRotationPolicy      []string                   `json:"key_rotation_policy,omitempty"`
	RevokedSignerBehavior  []string                   `json:"revoked_signer_behavior,omitempty"`
	TimestampingDiscipline []string                   `json:"timestamping_discipline,omitempty"`
	RequiredArtifactFields []string                   `json:"required_artifact_fields,omitempty"`
	Limitations            []string                   `json:"limitations,omitempty"`
}

type PublicProofCompatibilityBaseline struct {
	CurrentState                 string   `json:"current_state"`
	SupportedArtifactSchemas     []string `json:"supported_artifact_schemas,omitempty"`
	SupportedVerifierSchemaLines []string `json:"supported_verifier_schema_lines,omitempty"`
	BackwardCompatibilityPolicy  []string `json:"backward_compatibility_policy,omitempty"`
	DeprecationPolicy            []string `json:"deprecation_policy,omitempty"`
	ReplayTolerancePolicy        []string `json:"replay_tolerance_policy,omitempty"`
	UnsupportedCases             []string `json:"unsupported_cases,omitempty"`
	FailureStates                []string `json:"failure_states,omitempty"`
	Limitations                  []string `json:"limitations,omitempty"`
}

func MeasuredPublicProofVal0ClaimRegistryModel() PublicProofClaimRegistryModel {
	model := PublicProofClaimRegistryModel{
		RequiredClaimFields: []string{
			"claim_id",
			"claim_class",
			"owner",
			"scope",
			"methodology_ref",
			"environment_class",
			"artifact_refs",
			"freshness_class",
			"status",
			"projection_policy",
		},
		ClaimClasses: []PublicProofClaimClass{
			{
				ClaimClass:           PublicProofClaimClassPerformance,
				Purpose:              "Measured latency and benchmark-derived claims with explicit methodology and environment binding.",
				RequiredFields:       []string{"methodology_ref", "artifact_refs", "environment_class", "freshness_class"},
				SupportedScopes:      []string{ScopePublic, ScopePartner, ScopeAuditor, ScopeInternal},
				RequiredBoundaries:   []string{"methodology_bound", "environment_bound", "freshness_bound", "version_bound"},
				DisallowedOverclaims: []string{"universal_latency_truth", "cross_environment_parity_without_measurement"},
			},
			{
				ClaimClass:           PublicProofClaimClassVerification,
				Purpose:              "Verifier-oriented correctness or proof-verification claims tied to signed artifacts and trust roots.",
				RequiredFields:       []string{"methodology_ref", "artifact_refs", "trust_root", "key_version"},
				SupportedScopes:      []string{ScopePublic, ScopePartner, ScopeAuditor, ScopeInternal},
				RequiredBoundaries:   []string{"signing_bound", "trust_root_bound", "schema_bound", "freshness_bound"},
				DisallowedOverclaims: []string{"absolute_truth", "signature_replaces_methodology"},
			},
			{
				ClaimClass:           PublicProofClaimClassRuntimeBehavior,
				Purpose:              "Scoped runtime-behavior claims projected from bounded runtime evidence rather than raw tenant signal.",
				RequiredFields:       []string{"environment_class", "artifact_refs", "projection_policy", "limitations"},
				SupportedScopes:      []string{ScopePartner, ScopeAuditor, ScopeInternal},
				RequiredBoundaries:   []string{"projection_bound", "redaction_bound", "freshness_bound"},
				DisallowedOverclaims: []string{"global_runtime_score", "tenant_sensitive_raw_publication"},
			},
			{
				ClaimClass:           PublicProofClaimClassCompatibility,
				Purpose:              "Schema, verifier, and environment compatibility claims with visible unsupported and deprecated cases.",
				RequiredFields:       []string{"compatibility_policy", "schema_version", "artifact_refs", "environment_class"},
				SupportedScopes:      []string{ScopePublic, ScopePartner, ScopeAuditor, ScopeInternal},
				RequiredBoundaries:   []string{"schema_bound", "version_bound", "unsupported_cases_visible"},
				DisallowedOverclaims: []string{"universal_interoperability"},
			},
			{
				ClaimClass:           PublicProofClaimClassReplayability,
				Purpose:              "Replayability claims with explicit tolerance bands and unsupported replay cases.",
				RequiredFields:       []string{"replay_instructions", "tolerance_policy", "environment_class", "artifact_refs"},
				SupportedScopes:      []string{ScopePublic, ScopePartner, ScopeAuditor, ScopeInternal},
				RequiredBoundaries:   []string{"tolerance_bound", "environment_bound", "unsupported_replay_visible"},
				DisallowedOverclaims: []string{"bit_for_bit_replay_required_everywhere"},
			},
		},
		LifecycleStates: []PublicProofLifecycleState{
			{Status: PublicProofStatusProven, CurrentState: "claim_lifecycle_ready", VerifierVisible: true, Terminal: false, Description: "Signed, scoped, freshness-valid claim backed by active proof artifacts.", PublicationBehavior: "eligible_for_projection"},
			{Status: PublicProofStatusPartiallyProven, CurrentState: "claim_lifecycle_partial", VerifierVisible: true, Terminal: false, Description: "Bounded evidence exists but one or more publication or replay conditions remain partial.", PublicationBehavior: "restricted_projection"},
			{Status: PublicProofStatusProofPending, CurrentState: "claim_lifecycle_pending", VerifierVisible: true, Terminal: false, Description: "Proof reissue or publication gate is pending completion.", PublicationBehavior: "not_public"},
			{Status: PublicProofStatusProofFailed, CurrentState: "claim_lifecycle_failed", VerifierVisible: true, Terminal: false, Description: "Proof generation, signing, or verification failed.", PublicationBehavior: "not_public"},
			{Status: PublicProofStatusClaimNotReissued, CurrentState: "claim_lifecycle_reissue_required", VerifierVisible: true, Terminal: false, Description: "Previous claim exists but current release has not been reissued.", PublicationBehavior: "previous_claim_not_extended"},
			{Status: PublicProofStatusRestricted, CurrentState: "claim_lifecycle_restricted", VerifierVisible: true, Terminal: false, Description: "Claim is visible but publication or scope restrictions are active.", PublicationBehavior: "restricted_projection"},
			{Status: PublicProofStatusSuperseded, CurrentState: "claim_lifecycle_superseded", VerifierVisible: true, Terminal: true, Description: "Claim remains visible but is replaced by a newer approved claim.", PublicationBehavior: "supersession_chain_visible"},
			{Status: PublicProofStatusWithdrawn, CurrentState: "claim_lifecycle_withdrawn", VerifierVisible: true, Terminal: true, Description: "Claim or proof artifact is withdrawn and must remain verifier-visible.", PublicationBehavior: "withdrawn_notice_required"},
			{Status: PublicProofStatusStale, CurrentState: "claim_lifecycle_stale", VerifierVisible: true, Terminal: false, Description: "Freshness or re-verification window elapsed.", PublicationBehavior: "stale_notice_required"},
		},
		MethodologyBoundaryRules: []string{
			"methodology, measurement, interpretation, and limitations must remain explicitly separated",
			"signed artifacts do not replace measurement or methodology boundaries",
			"claims must stay environment-bound, methodology-bound, version-bound, and freshness-bound",
		},
		FreshnessRules: []string{
			"each claim requires issued_at and valid_through or freshness_class",
			"stale and withdrawn states must remain verifier-visible",
			"new releases do not inherit prior public claims without reissue or bounded restriction",
		},
		RevocationRules: []string{
			"withdrawn claims remain visible with withdrawn semantics",
			"revoked signer or trust-root state blocks new proof issuance",
			"incident correction requires restriction, supersession, or withdrawal instead of silent mutation",
		},
		SupersessionRules: []string{
			"superseded claims must keep superseded_by linkage visible",
			"supersession does not delete historical claim lineage",
			"claim_not_reissued must remain distinct from superseded",
		},
		RequiredProjectionPolicies: []string{
			RedactionTierInternalFull,
			RedactionTierPartnerScoped,
			RedactionTierPublicSafe,
		},
		Limitations: []string{
			"Val 0 defines the claim-registry discipline only and does not yet issue sealed proof artifacts.",
			"Public proof discipline does not become universal authority through taxonomy or status modeling alone.",
		},
	}
	model.CurrentState = EvaluateMeasuredPublicProofVal0ClaimRegistryState(model)
	return model
}

func MeasuredPublicProofVal0RedactionTiers() []PublicProofRedactionTier {
	return []PublicProofRedactionTier{
		{
			TierID:               RedactionTierInternalFull,
			CurrentState:         "redaction_tier_ready",
			AllowedScopes:        []string{ScopeInternal},
			RemovedFields:        []string{},
			AggregatedFields:     []string{},
			NeverPublishedFields: []string{},
			PortalPolicy:         "internal viewers may inspect full bounded artifact and raw evidence lineage within existing access controls",
			Limitations:          []string{"Internal full does not authorize cross-tenant disclosure or public projection."},
		},
		{
			TierID:               RedactionTierPartnerScoped,
			CurrentState:         "redaction_tier_ready",
			AllowedScopes:        []string{ScopePartner, ScopeAuditor, ScopeInternal},
			RemovedFields:        []string{"tenant_sensitive_raw_events", "customer_identifiers", "raw_runtime_signal"},
			AggregatedFields:     []string{"environment_descriptors", "measurement_bands", "artifact_lineage_summary"},
			NeverPublishedFields: []string{"tenant_sensitive_raw_events", "private_anchor_credentials"},
			PortalPolicy:         "partner-safe views expose bounded evidence and replay posture without tenant-sensitive raw signal",
			Limitations:          []string{"Partner-scoped projection still requires explicit approval and bounded disclosure posture."},
		},
		{
			TierID:               RedactionTierPublicSafe,
			CurrentState:         "redaction_tier_ready",
			AllowedScopes:        []string{ScopePublic, ScopePartner, ScopeAuditor, ScopeInternal},
			RemovedFields:        []string{"tenant_sensitive_raw_events", "partner_private_refs", "customer_identifiers", "raw_runtime_signal", "non_publishable_anchor_material"},
			AggregatedFields:     []string{"metric_summaries", "environment_class", "freshness_state", "compatibility_scope", "supersession_chain"},
			NeverPublishedFields: []string{"tenant_sensitive_raw_events", "customer_identifiers", "non_publishable_anchor_material"},
			PortalPolicy:         "public-safe views remain methodology-bound projections over the evidence spine and never become a new truth store",
			Limitations:          []string{"Public-safe projection may aggregate or remove fields even when partner-scoped views remain more detailed."},
		},
	}
}

func MeasuredPublicProofVal0SigningAuthority(provider signing.ProviderDescriptor) PublicProofSigningAuthorityModel {
	model := PublicProofSigningAuthorityModel{
		Provider: provider,
		TrustRoots: []PublicProofTrustRoot{
			{
				TrustRootID:           "public_proof_primary_root",
				CurrentState:          "trust_root_policy_ready",
				SignerIdentity:        firstNonEmpty(provider.ProviderMode, "disabled_signer"),
				KeyVersion:            "v1",
				IssuanceWindow:        "2026-01-01T00:00:00Z/2026-12-31T23:59:59Z",
				RotationPolicy:        "new key versions require overlap, verify-only retirement, and signer-identity continuity review",
				RevocationBehavior:    "revoked signer blocks new public proof issuance and forces verifier-visible restriction or withdrawal",
				TimestampingPolicy:    "public or partner proof issuance requires artifact timestamping before activation",
				SupportedClaimClasses: []string{PublicProofClaimClassPerformance, PublicProofClaimClassVerification, PublicProofClaimClassRuntimeBehavior, PublicProofClaimClassCompatibility, PublicProofClaimClassReplayability},
				Limitations:           []string{"Trust-root discipline is modeled here; multi-root issuance and transparency anchoring arrive in later Point 2 waves."},
			},
		},
		KeyRotationPolicy: []string{
			"key rotation requires trust-root and key-version visibility",
			"retired keys may remain verify-only until supersession windows close",
			"revoked keys remain verifier-visible and block new artifact issuance",
		},
		RevokedSignerBehavior: []string{
			"revoked signer blocks new signed artifact publication",
			"claims signed by revoked material transition to restricted, superseded, or withdrawn rather than silently disappearing",
		},
		TimestampingDiscipline: []string{
			"issuance requires timestamp linkage before public or partner activation",
			"timestamp discipline does not replace methodology or scope boundaries",
		},
		RequiredArtifactFields: []string{
			"artifact_id",
			"artifact_schema_version",
			"claim_linkage",
			"signer_metadata",
			"evidence_refs",
			"redaction_class",
			"timestamp",
		},
		Limitations: uniqueStrings(append([]string{
			"Val 0 defines signing authority, trust-root, rotation, and revocation discipline, but not yet sealed proof issuance.",
		}, provider.Limitations...)),
	}
	model.CurrentState = EvaluateMeasuredPublicProofVal0SigningAuthorityState(model)
	return model
}

func MeasuredPublicProofVal0CompatibilityBaseline() PublicProofCompatibilityBaseline {
	model := PublicProofCompatibilityBaseline{
		SupportedArtifactSchemas: []string{
			"public.proof.claim_registry.v1",
			"public.proof.status_lifecycle.v1",
			"public.proof.redaction_projection.v1",
			"public.proof.signing_authority.v1",
			"public.proof.compatibility_baseline.v1",
		},
		SupportedVerifierSchemaLines: []string{
			"/v1/public/verifier/sdk",
			"/v1/public/specs/proof-verification",
			"/v1/public/phase6/proofs",
		},
		BackwardCompatibilityPolicy: []string{
			"minor additive fields remain backward-compatible within a supported schema line",
			"breaking field removals require a new schema line and visible deprecation notice",
			"portal projections must show verifier incompatibility instead of silently coercing schema lines",
		},
		DeprecationPolicy: []string{
			"deprecated schema lines remain verifier-visible before removal",
			"withdrawn or superseded claims must surface compatibility and deprecation state together where relevant",
		},
		ReplayTolerancePolicy: []string{
			"replay remains environment-compatible and methodology-compatible, not bit-for-bit identical",
			"tolerance bands must be visible for replayability claims",
			"unsupported replay cases must remain explicit",
		},
		UnsupportedCases: []string{
			"cross-schema replay without compatibility declaration",
			"cross-environment comparison outside declared tolerance bands",
			"public-safe replay over tenant-sensitive raw evidence",
		},
		FailureStates: []string{
			"proof_generation_failed",
			"signature_failed",
			"anchoring_unavailable",
			"replay_unavailable",
			"freshness_expired",
			"claim_restricted",
			"claim_withdrawn",
			"claim_superseded",
		},
		Limitations: []string{
			"Compatibility discipline is defined here before later waves add full verifier SDK and sealed artifact issuance.",
		},
	}
	model.CurrentState = EvaluateMeasuredPublicProofVal0CompatibilityState(model)
	return model
}

func EvaluateMeasuredPublicProofVal0ClaimRegistryState(model PublicProofClaimRegistryModel) string {
	if len(model.RequiredClaimFields) == 0 || len(model.ClaimClasses) == 0 || len(model.LifecycleStates) == 0 {
		return MeasuredPublicProofVal0ClaimRegistryStateIncomplete
	}
	if len(model.MethodologyBoundaryRules) == 0 || len(model.FreshnessRules) == 0 || len(model.RevocationRules) == 0 || len(model.SupersessionRules) == 0 || len(model.RequiredProjectionPolicies) == 0 {
		return MeasuredPublicProofVal0ClaimRegistryStatePartial
	}
	return MeasuredPublicProofVal0ClaimRegistryStateActive
}

func EvaluateMeasuredPublicProofVal0RedactionTierState(items []PublicProofRedactionTier) string {
	if len(items) == 0 {
		return MeasuredPublicProofVal0RedactionTierStateIncomplete
	}
	required := map[string]struct{}{
		RedactionTierInternalFull:  {},
		RedactionTierPartnerScoped: {},
		RedactionTierPublicSafe:    {},
	}
	for _, item := range items {
		if strings.TrimSpace(item.TierID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.PortalPolicy) == "" {
			return MeasuredPublicProofVal0RedactionTierStatePartial
		}
		if len(item.AllowedScopes) == 0 || len(item.NeverPublishedFields) == 0 && item.TierID != RedactionTierInternalFull {
			return MeasuredPublicProofVal0RedactionTierStatePartial
		}
		delete(required, strings.TrimSpace(item.TierID))
	}
	if len(required) != 0 {
		return MeasuredPublicProofVal0RedactionTierStatePartial
	}
	return MeasuredPublicProofVal0RedactionTierStateActive
}

func EvaluateMeasuredPublicProofVal0SigningAuthorityState(model PublicProofSigningAuthorityModel) string {
	if strings.TrimSpace(model.Provider.ProviderMode) == "" || strings.TrimSpace(model.Provider.TrustBoundary) == "" {
		return MeasuredPublicProofVal0SigningAuthorityStateIncomplete
	}
	if len(model.TrustRoots) == 0 || len(model.KeyRotationPolicy) == 0 || len(model.RevokedSignerBehavior) == 0 || len(model.TimestampingDiscipline) == 0 || len(model.RequiredArtifactFields) == 0 {
		return MeasuredPublicProofVal0SigningAuthorityStatePartial
	}
	for _, root := range model.TrustRoots {
		if strings.TrimSpace(root.TrustRootID) == "" || strings.TrimSpace(root.CurrentState) == "" || strings.TrimSpace(root.KeyVersion) == "" || strings.TrimSpace(root.RotationPolicy) == "" || strings.TrimSpace(root.RevocationBehavior) == "" || strings.TrimSpace(root.TimestampingPolicy) == "" {
			return MeasuredPublicProofVal0SigningAuthorityStatePartial
		}
	}
	return MeasuredPublicProofVal0SigningAuthorityStateActive
}

func EvaluateMeasuredPublicProofVal0CompatibilityState(model PublicProofCompatibilityBaseline) string {
	if len(model.SupportedArtifactSchemas) == 0 || len(model.SupportedVerifierSchemaLines) == 0 {
		return MeasuredPublicProofVal0CompatibilityStateIncomplete
	}
	if len(model.BackwardCompatibilityPolicy) == 0 || len(model.DeprecationPolicy) == 0 || len(model.ReplayTolerancePolicy) == 0 || len(model.UnsupportedCases) == 0 || len(model.FailureStates) == 0 {
		return MeasuredPublicProofVal0CompatibilityStatePartial
	}
	return MeasuredPublicProofVal0CompatibilityStateActive
}

func EvaluateMeasuredPublicProofVal0State(claimRegistryState, redactionTierState, signingAuthorityState, compatibilityState string) string {
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(claimRegistryState),
		strings.TrimSpace(redactionTierState),
		strings.TrimSpace(signingAuthorityState),
		strings.TrimSpace(compatibilityState),
	} {
		switch state {
		case MeasuredPublicProofVal0ClaimRegistryStateActive,
			MeasuredPublicProofVal0RedactionTierStateActive,
			MeasuredPublicProofVal0SigningAuthorityStateActive,
			MeasuredPublicProofVal0CompatibilityStateActive:
		case MeasuredPublicProofVal0ClaimRegistryStatePartial,
			MeasuredPublicProofVal0RedactionTierStatePartial,
			MeasuredPublicProofVal0SigningAuthorityStatePartial,
			MeasuredPublicProofVal0CompatibilityStatePartial:
			hasPartial = true
		default:
			return MeasuredPublicProofVal0StateIncomplete
		}
	}
	if hasPartial {
		return MeasuredPublicProofVal0StateSubstantial
	}
	return MeasuredPublicProofVal0StateActive
}
