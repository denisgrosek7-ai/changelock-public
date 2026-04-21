package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	publicTrustBadgeProgramSchema  = "5d.public_conformance_badges.v1"
	publicVerifierProgramSchema    = "5d.public_verifier_program.v1"
	publicClaimsGovernanceSchema   = "5d.public_claims_governance.v1"
	publicTrustMarkLifecycleSchema = "5d.public_trust_mark_lifecycle.v1"
)

type publicTrustBadgeDefinition struct {
	BadgeID              string   `json:"badge_id"`
	DisplayName          string   `json:"display_name"`
	Meaning              []string `json:"meaning,omitempty"`
	BadgeCriteria        []string `json:"badge_criteria,omitempty"`
	EvidenceRequirements []string `json:"evidence_requirements,omitempty"`
	ValidityPeriod       string   `json:"validity_period"`
	RevocationRules      []string `json:"revocation_rules,omitempty"`
	VerificationMethod   []string `json:"verification_method,omitempty"`
	ConformanceProfiles  []string `json:"conformance_profiles,omitempty"`
	ProgramBoundaries    []string `json:"program_boundaries,omitempty"`
	Limitations          []string `json:"limitations,omitempty"`
}

type publicTrustBadgeProgramResponse struct {
	SchemaVersion     string                       `json:"schema_version"`
	ProgramID         string                       `json:"program_id"`
	BadgeDefinitions  []publicTrustBadgeDefinition `json:"badge_definitions,omitempty"`
	MeaningGuardrails []string                     `json:"meaning_guardrails,omitempty"`
	LookupRef         string                       `json:"lookup_ref,omitempty"`
	Limitations       []string                     `json:"limitations,omitempty"`
}

type publicVerifierProgramProfile struct {
	ProfileID              string   `json:"profile_id"`
	DisplayName            string   `json:"display_name"`
	OnboardingRequirements []string `json:"onboarding_requirements,omitempty"`
	ConformanceTargets     []string `json:"conformance_targets,omitempty"`
	VerificationGuidance   []string `json:"verification_guidance,omitempty"`
	CompatibilityGuidance  []string `json:"compatibility_guidance,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type publicVerifierProgramResponse struct {
	SchemaVersion        string                         `json:"schema_version"`
	ProgramID            string                         `json:"program_id"`
	OnboardingFlow       []string                       `json:"onboarding_flow,omitempty"`
	ConformanceTesting   []string                       `json:"conformance_testing,omitempty"`
	Profiles             []publicVerifierProgramProfile `json:"profiles,omitempty"`
	DisputeMismatchModel []string                       `json:"dispute_mismatch_model,omitempty"`
	VersionCompatibility []string                       `json:"version_compatibility,omitempty"`
	Limitations          []string                       `json:"limitations,omitempty"`
}

type publicClaimClass struct {
	ClaimClass            string   `json:"claim_class"`
	AllowedExamples       []string `json:"allowed_examples,omitempty"`
	RequiredEvidence      []string `json:"required_evidence,omitempty"`
	ApprovalPath          []string `json:"approval_path,omitempty"`
	DisallowedLanguage    []string `json:"disallowed_language,omitempty"`
	BenchmarkDiscipline   []string `json:"benchmark_discipline,omitempty"`
	CertificationBoundary []string `json:"certification_boundary,omitempty"`
	Limitations           []string `json:"limitations,omitempty"`
}

type publicClaimsGovernanceResponse struct {
	SchemaVersion         string             `json:"schema_version"`
	PolicyID              string             `json:"policy_id"`
	ClaimClasses          []publicClaimClass `json:"claim_classes,omitempty"`
	ReviewWorkflow        []string           `json:"review_workflow,omitempty"`
	RegionalCautions      []string           `json:"regional_cautions,omitempty"`
	PublicationBoundaries []string           `json:"publication_boundaries,omitempty"`
	Limitations           []string           `json:"limitations,omitempty"`
}

type publicTrustMarkHistory struct {
	State     string     `json:"state"`
	Summary   string     `json:"summary"`
	ChangedAt time.Time  `json:"changed_at"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

type publicTrustMarkStatus struct {
	MarkID           string                   `json:"mark_id"`
	BadgeID          string                   `json:"badge_id"`
	DisplayName      string                   `json:"display_name"`
	CurrentState     string                   `json:"current_state"`
	MeaningSummary   []string                 `json:"meaning_summary,omitempty"`
	IssuedAt         time.Time                `json:"issued_at"`
	ExpiresAt        *time.Time               `json:"expires_at,omitempty"`
	RevalidateBy     *time.Time               `json:"revalidate_by,omitempty"`
	VerificationURI  string                   `json:"verification_uri,omitempty"`
	RevocationReason string                   `json:"revocation_reason,omitempty"`
	HistoricalStatus []publicTrustMarkHistory `json:"historical_status,omitempty"`
	Limitations      []string                 `json:"limitations,omitempty"`
}

type publicTrustMarkLifecycleResponse struct {
	SchemaVersion   string                  `json:"schema_version"`
	ProgramID       string                  `json:"program_id"`
	LookupSemantics []string                `json:"lookup_semantics,omitempty"`
	IssuanceModel   []string                `json:"issuance_model,omitempty"`
	RevocationModel []string                `json:"revocation_model,omitempty"`
	Marks           []publicTrustMarkStatus `json:"marks,omitempty"`
	Limitations     []string                `json:"limitations,omitempty"`
}

func (s server) publicTrustBadgeProgramHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicTrustBadgeProgram())
}

func (s server) publicVerifierProgramHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicVerifierProgram())
}

func (s server) publicClaimsGovernanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicClaimsGovernance())
}

func (s server) publicTrustMarkLifecycleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicTrustMarkLifecycle())
}

func (s server) publicTrustMarkByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	markID := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/public/trust-program/marks/"))
	if markID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "trust mark not found"})
		return
	}
	mark, ok := publicTrustMarkByID(markID)
	if !ok {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "trust mark not found"})
		return
	}
	httpjson.Write(w, http.StatusOK, mark)
}

func buildPublicTrustBadgeProgram() publicTrustBadgeProgramResponse {
	return publicTrustBadgeProgramResponse{
		SchemaVersion: publicTrustBadgeProgramSchema,
		ProgramID:     "public_trust_badge_program_v1",
		BadgeDefinitions: []publicTrustBadgeDefinition{
			{
				BadgeID:     "verification_ready",
				DisplayName: "Verification Ready",
				Meaning: []string{
					"Indicates that the published artifact or surface can be independently checked through public specs, public samples, and verifier guidance.",
					"It does not mean the underlying deployment is universally secure or certified.",
				},
				BadgeCriteria: []string{
					"public schema export exists",
					"public sample or replay input exists",
					"verifier expectations and failure states are documented",
				},
				EvidenceRequirements: []string{
					"/v1/public/specs/handoff",
					"/v1/public/schemas",
					"/v1/public/verifier/reference-pack",
				},
				ValidityPeriod:      "90d",
				RevocationRules:     []string{"schema breaking change without updated verifier guidance", "public verifier pack no longer replays successfully", "sample or schema removed without replacement"},
				VerificationMethod:  []string{"replay the reference pack", "verify sample-to-schema compatibility", "confirm conformance assertions still pass"},
				ConformanceProfiles: []string{"minimal_verifier", "full_verifier", "auditor", "partner_verifier"},
				ProgramBoundaries:   []string{"This mark is about verifier-facing artifact quality, not environment-wide trust or certification."},
			},
			{
				BadgeID:     "public_architecture_documented",
				DisplayName: "Public Architecture Documented",
				Meaning: []string{
					"Indicates that the deployment profile, maturity expectations, and decision guidance are publicly documented in bounded form.",
					"It does not imply one-click deployment correctness or universal best-fit architecture.",
				},
				BadgeCriteria: []string{
					"reference architecture exists",
					"maturity level exists",
					"decision guide and decision matrix linkage exists",
				},
				EvidenceRequirements: []string{
					"/v1/public/reference-architectures",
					"/v1/public/maturity-map",
					"/v1/public/decision-guides/matrix",
				},
				ValidityPeriod:     "180d",
				RevocationRules:    []string{"reference architecture assumptions removed", "decision guides no longer map to active contracts", "maturity criteria drift beyond the published surfaces"},
				VerificationMethod: []string{"validate references and maturity linkage", "confirm documented limitations remain visible"},
				ProgramBoundaries:  []string{"This mark documents bounded architecture clarity only; it is not a compliance or sector approval mark."},
			},
			{
				BadgeID:     "public_benchmark_disciplined",
				DisplayName: "Public Benchmark Disciplined",
				Meaning: []string{
					"Indicates that benchmark publication follows explicit methodology, limitation labeling, and non-claim discipline.",
					"It does not imply that all benchmark entries are already measured or marketing-ready.",
				},
				BadgeCriteria: []string{
					"public benchmark methodology exists",
					"benchmark set publication status is explicit",
					"analytics publication rules and case-study discipline are documented",
				},
				EvidenceRequirements: []string{
					"/v1/public/benchmarks/methodology",
					"/v1/public/benchmarks/set",
					"/v1/public/analytics/publication-discipline",
					"/v1/public/case-studies",
				},
				ValidityPeriod:     "90d",
				RevocationRules:    []string{"benchmark publication starts over-claiming beyond measured status", "publication statuses disappear", "case-study reproducibility or limitations are removed"},
				VerificationMethod: []string{"check benchmark status labels", "check not-claimed language", "confirm case-study limitations are still published"},
				ProgramBoundaries:  []string{"This mark confirms publication discipline, not best-in-class performance."},
			},
		},
		MeaningGuardrails: []string{
			"Badge meaning must stay public, evidence-backed, time-bounded, and revocable.",
			"Badges must not be described as certifications, legal attestations, or universal customer guarantees without a separate formal program.",
		},
		LookupRef: "/v1/public/trust-program/marks",
		Limitations: []string{
			"This badge program is a bounded conformance and publication program. It is not a formal certification body, legal authority, or blanket product endorsement program.",
		},
	}
}

func buildPublicVerifierProgram() publicVerifierProgramResponse {
	return publicVerifierProgramResponse{
		SchemaVersion: publicVerifierProgramSchema,
		ProgramID:     "public_verifier_program_v1",
		OnboardingFlow: []string{
			"Choose a verifier profile and supported schema_version line.",
			"Replay the public reference pack and conformance pack.",
			"Document which public specs, failure states, and limitation classes the verifier actually supports.",
			"Retest whenever public schema or conformance guidance changes.",
		},
		ConformanceTesting: []string{
			"Use /v1/public/conformance-pack for bounded assertion targets.",
			"Use /v1/public/verifier/reference-pack for replay inputs.",
			"Treat stale, rejected, and divergence examples as first-class expected outcomes rather than generic errors.",
		},
		Profiles: []publicVerifierProgramProfile{
			{
				ProfileID:              "minimal_verifier",
				DisplayName:            "Minimal verifier onboarding",
				OnboardingRequirements: []string{"verify sealed bundle integrity", "render failure states distinctly", "preserve limitation labels"},
				ConformanceTargets:     []string{"/v1/public/conformance-pack", "/v1/public/samples/handoff"},
				VerificationGuidance:   []string{"must not silently trust remote peers", "must preserve exact accept/reject semantics"},
				CompatibilityGuidance:  []string{"compatible within one stable schema line", "requires retest on new schema_version"},
			},
			{
				ProfileID:              "full_verifier",
				DisplayName:            "Full verifier onboarding",
				OnboardingRequirements: []string{"support public handoff, proof, validation, and federation formats", "preserve conformance semantics across all public samples"},
				ConformanceTargets:     []string{"/v1/public/conformance-pack", "/v1/public/schemas"},
				VerificationGuidance:   []string{"must preserve divergence, stale, and advisory states", "must not upgrade advisory semantics into stronger trust claims"},
				CompatibilityGuidance:  []string{"must publish supported schema versions explicitly"},
			},
			{
				ProfileID:              "auditor",
				DisplayName:            "Auditor verifier onboarding",
				OnboardingRequirements: []string{"support offline verification", "support historical limitation interpretation", "support bounded verification narrative export"},
				ConformanceTargets:     []string{"/v1/public/verifier/offline-guide", "/v1/public/verifier/reference-pack"},
				VerificationGuidance:   []string{"must preserve offline failure semantics", "must not treat archived proof as perpetual live trust"},
				CompatibilityGuidance:  []string{"auditor profile remains scoped to archived and public artifacts"},
			},
			{
				ProfileID:              "partner_verifier",
				DisplayName:            "Partner verifier onboarding",
				OnboardingRequirements: []string{"support local admissibility", "support rejection and distrust semantics", "support disclosure-minimized proof handling"},
				ConformanceTargets:     []string{"/v1/public/samples/proof-verification", "/v1/public/samples/federation-proof-exchange"},
				VerificationGuidance:   []string{"must distinguish cryptographic validity from local admissibility", "must keep local override authoritative"},
				CompatibilityGuidance:  []string{"partner support must declare accepted disclosure profiles and version lines"},
			},
		},
		DisputeMismatchModel: []string{
			"Version mismatch, schema mismatch, and semantic mismatch must be reported separately.",
			"When two verifiers disagree, the disagreement must identify whether the difference is cryptographic, freshness-related, schema-related, or local-policy-related.",
			"Public verifier program does not centralize dispute resolution into a global authority.",
		},
		VersionCompatibility: []string{
			"Support is declared per public schema_version and profile.",
			"Breaking semantic changes require new public schema versions and renewed conformance testing.",
			"Older stable verifier lines remain valid only within the compatibility guidance of their published schema line.",
		},
		Limitations: []string{
			"This verifier program defines onboarding and conformance discipline for third-party verifiers. It is not an accreditation body or legal attestation program.",
		},
	}
}

func buildPublicClaimsGovernance() publicClaimsGovernanceResponse {
	return publicClaimsGovernanceResponse{
		SchemaVersion: publicClaimsGovernanceSchema,
		PolicyID:      "public_claims_governance_v1",
		ClaimClasses: []publicClaimClass{
			{
				ClaimClass:         "verification_claim",
				AllowedExamples:    []string{"publicly documented and third-party verifiable", "bounded verifier-ready format", "offline verification supported for the published bundle shape"},
				RequiredEvidence:   []string{"/v1/public/specs/handoff", "/v1/public/schemas", "/v1/public/verifier/reference-pack"},
				ApprovalPath:       []string{"engineering confirms current public artifacts replay", "documentation confirms limitation language remains intact"},
				DisallowedLanguage: []string{"universally trusted", "guaranteed interoperable everywhere", "global proof authority"},
				Limitations:        []string{"Verification claims must stay tied to the exact public formats and samples currently published."},
			},
			{
				ClaimClass:          "benchmark_claim",
				AllowedExamples:     []string{"methodology published", "benchmark entry measured in the declared substrate", "starting points only where measurement is not yet public"},
				RequiredEvidence:    []string{"/v1/public/benchmarks/methodology", "/v1/public/benchmarks/set"},
				ApprovalPath:        []string{"measurement owner reviews publication status", "claims reviewer checks non-claim language and substrate context"},
				DisallowedLanguage:  []string{"always low latency", "under one percent overhead everywhere", "400% better"},
				BenchmarkDiscipline: []string{"claim must not exceed benchmark publication status", "methodology-only entries cannot be marketed as measured wins"},
			},
			{
				ClaimClass:            "badge_or_mark_claim",
				AllowedExamples:       []string{"verification ready under the bounded public trust program", "public benchmark disciplined under the published methodology"},
				RequiredEvidence:      []string{"/v1/public/trust-program/badges", "/v1/public/trust-program/marks"},
				ApprovalPath:          []string{"program owner confirms issuance status", "expiration and revocation state are checked before publication"},
				DisallowedLanguage:    []string{"certified secure", "industry standard by default", "official certification"},
				CertificationBoundary: []string{"trust marks are bounded conformance or publication signals only", "formal certification language remains prohibited without a separate legal program"},
			},
			{
				ClaimClass:            "sector_or_compliance_adjacent_claim",
				AllowedExamples:       []string{"regulated deployment profile documented", "FIPS-readiness mapping exists where applicable"},
				RequiredEvidence:      []string{"/v1/public/reference-architectures", "/v1/execution/compliance-readiness"},
				ApprovalPath:          []string{"engineering confirms evidence-backed readiness wording", "legal or commercial review confirms region-specific wording where needed"},
				DisallowedLanguage:    []string{"certified compliant", "formally approved in every region", "guaranteed regulator accepted"},
				CertificationBoundary: []string{"readiness language must not be upgraded into certification language", "sector profiles remain bounded architecture guidance"},
			},
		},
		ReviewWorkflow: []string{
			"Map the claim to one claim class before publication.",
			"Check the required evidence and current publication status.",
			"Reject wording that exceeds the evidence class or current trust mark state.",
			"Re-review claims after mark expiry, revocation, benchmark methodology changes, or schema-version changes.",
		},
		RegionalCautions: []string{
			"Sector, certification, or regulatory wording may require additional regional legal review outside this repository.",
			"Public claims governance is technical discipline plus publication control; it is not a substitute for jurisdiction-specific legal advice.",
		},
		PublicationBoundaries: []string{
			"Public claims must preserve limitation and freshness context where the underlying surface requires it.",
			"Marketing wording must not outrun evidence, conformance status, or benchmark publication status.",
		},
		Limitations: []string{
			"This policy governs public claim language for the current repository surfaces. It does not by itself create a formal certification or legal approval framework.",
		},
	}
}

func buildPublicTrustMarkLifecycle() publicTrustMarkLifecycleResponse {
	return publicTrustMarkLifecycleResponse{
		SchemaVersion: publicTrustMarkLifecycleSchema,
		ProgramID:     "public_trust_mark_lifecycle_v1",
		LookupSemantics: []string{
			"Public lookup reports current mark state, expiry, revalidation target, and historical state changes.",
			"Lookup is bounded to program-issued public marks and does not expose internal tenant or unpublished trust data.",
		},
		IssuanceModel: []string{
			"Marks are issued only when their evidence requirements and program criteria are satisfied.",
			"Every mark has a bounded lifetime and a revalidation deadline.",
			"Mark meaning is defined by the public badge program, not by ad hoc sales or support language.",
		},
		RevocationModel: []string{
			"Marks can be revoked when evidence falls out of date, required public surfaces disappear, or published semantics drift beyond the mark criteria.",
			"Revocation remains historically visible rather than silently deleting prior issuance state.",
		},
		Marks: publicTrustMarkCatalog(),
		Limitations: []string{
			"Trust marks are bounded public program statuses. They are not permanent labels, customer-specific guarantees, or formal compliance certificates.",
		},
	}
}

func publicTrustMarkCatalog() []publicTrustMarkStatus {
	issuedAt := publicSampleTime()
	revalidateBy := issuedAt.Add(60 * 24 * time.Hour)
	expiresAt := issuedAt.Add(90 * 24 * time.Hour)
	revokedAt := issuedAt.Add(30 * 24 * time.Hour)
	return []publicTrustMarkStatus{
		{
			MarkID:          "mark-verification-ready-sample",
			BadgeID:         "verification_ready",
			DisplayName:     "Verification Ready Sample Mark",
			CurrentState:    "active",
			MeaningSummary:  []string{"Public verifier artifacts exist and replay inputs remain available.", "Mark remains bounded to verifier-facing surfaces."},
			IssuedAt:        issuedAt,
			ExpiresAt:       &expiresAt,
			RevalidateBy:    &revalidateBy,
			VerificationURI: "/v1/public/trust-program/marks/mark-verification-ready-sample",
			HistoricalStatus: []publicTrustMarkHistory{
				{State: "issued", Summary: "Issued after public schema, sample, and verifier pack surfaces became available.", ChangedAt: issuedAt, ExpiresAt: &expiresAt},
				{State: "revalidated", Summary: "Revalidated against current reference-pack replay and conformance assertions.", ChangedAt: issuedAt.Add(14 * 24 * time.Hour), ExpiresAt: &expiresAt},
			},
			Limitations: []string{"Mark applies only to the public verification program surfaces and not to all private deployment states."},
		},
		{
			MarkID:           "mark-benchmark-discipline-sample",
			BadgeID:          "public_benchmark_disciplined",
			DisplayName:      "Benchmark Discipline Sample Mark",
			CurrentState:     "revoked",
			MeaningSummary:   []string{"Demonstrates revocation semantics when publication discipline changes or evidence expires."},
			IssuedAt:         issuedAt,
			ExpiresAt:        &expiresAt,
			RevalidateBy:     &revalidateBy,
			VerificationURI:  "/v1/public/trust-program/marks/mark-benchmark-discipline-sample",
			RevocationReason: "example revocation: benchmark methodology changed and the older mark was not revalidated in time",
			HistoricalStatus: []publicTrustMarkHistory{
				{State: "issued", Summary: "Issued when benchmark methodology and publication set were first published.", ChangedAt: issuedAt, ExpiresAt: &expiresAt},
				{State: "revoked", Summary: "Revoked after the sample mark missed the declared revalidation window.", ChangedAt: revokedAt},
			},
			Limitations: []string{"This revoked sample mark exists to make lifecycle and historical-status semantics public and testable."},
		},
	}
}

func publicTrustMarkByID(markID string) (publicTrustMarkStatus, bool) {
	for _, item := range publicTrustMarkCatalog() {
		if item.MarkID == markID {
			return item, true
		}
	}
	return publicTrustMarkStatus{}, false
}
