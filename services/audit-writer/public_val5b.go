package main

import (
	"net/http"

	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	publicReferenceArchitecturesSchema = "5b.public_reference_architectures.v1"
	publicMaturityMapSchema            = "5b.public_zero_trust_maturity_map.v1"
	publicDecisionGuidesSchema         = "5b.public_architecture_decision_guides.v1"
	publicSectorProfilesSchema         = "5b.public_sector_profiles.v1"
	publicDecisionMatrixSchema         = "5b.public_deployment_decision_matrix.v1"
)

type publicReferenceArchitecture struct {
	ArchitectureID      string   `json:"architecture_id"`
	DisplayName         string   `json:"display_name"`
	SectorProfile       string   `json:"sector_profile"`
	DeploymentProfile   string   `json:"deployment_profile"`
	TrustBoundaries     []string `json:"trust_boundaries,omitempty"`
	CriticalServices    []string `json:"critical_services,omitempty"`
	TopologySummary     []string `json:"topology_summary,omitempty"`
	HandoffModel        []string `json:"handoff_model,omitempty"`
	ValidationModel     []string `json:"validation_model,omitempty"`
	RuntimeModel        []string `json:"runtime_model,omitempty"`
	GovernanceOverlays  []string `json:"governance_overlays,omitempty"`
	IntegrationPatterns []string `json:"integration_patterns,omitempty"`
	Assumptions         []string `json:"assumptions,omitempty"`
	KnownLimitations    []string `json:"known_limitations,omitempty"`
}

type publicReferenceArchitecturesResponse struct {
	SchemaVersion     string                        `json:"schema_version"`
	Architectures     []publicReferenceArchitecture `json:"architectures,omitempty"`
	DecisionGuideRefs []string                      `json:"decision_guide_refs,omitempty"`
	Limitations       []string                      `json:"limitations,omitempty"`
}

type publicMaturityLevel struct {
	LevelID                string   `json:"level_id"`
	DisplayName            string   `json:"display_name"`
	RequiredCapabilities   []string `json:"required_capabilities,omitempty"`
	EvidenceExpectations   []string `json:"evidence_expectations,omitempty"`
	GovernanceExpectations []string `json:"governance_expectations,omitempty"`
	RuntimeExpectations    []string `json:"runtime_expectations,omitempty"`
	ValidationExpectations []string `json:"validation_expectations,omitempty"`
	B2BTrustExpectations   []string `json:"b2b_trust_expectations,omitempty"`
	MeasurableCriteria     []string `json:"measurable_criteria,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type publicMaturityMapResponse struct {
	SchemaVersion string                `json:"schema_version"`
	Levels        []publicMaturityLevel `json:"levels,omitempty"`
	Limitations   []string              `json:"limitations,omitempty"`
}

type publicDecisionGuide struct {
	GuideID                   string   `json:"guide_id"`
	Question                  string   `json:"question"`
	UseWhen                   []string `json:"use_when,omitempty"`
	AvoidWhen                 []string `json:"avoid_when,omitempty"`
	RecommendationSummary     string   `json:"recommendation_summary"`
	Tradeoffs                 []string `json:"tradeoffs,omitempty"`
	RelatedContracts          []string `json:"related_contracts,omitempty"`
	ReferenceArchitectureRefs []string `json:"reference_architecture_refs,omitempty"`
	Limitations               []string `json:"limitations,omitempty"`
}

type publicDecisionGuidesResponse struct {
	SchemaVersion string                `json:"schema_version"`
	Guides        []publicDecisionGuide `json:"guides,omitempty"`
	Limitations   []string              `json:"limitations,omitempty"`
}

type publicSectorProfile struct {
	SectorProfileID             string   `json:"sector_profile_id"`
	DisplayName                 string   `json:"display_name"`
	RecommendedArchitectureRefs []string `json:"recommended_architecture_refs,omitempty"`
	TrustPriorities             []string `json:"trust_priorities,omitempty"`
	RequiredContracts           []string `json:"required_contracts,omitempty"`
	GovernanceExpectations      []string `json:"governance_expectations,omitempty"`
	DeploymentAssumptions       []string `json:"deployment_assumptions,omitempty"`
	KnownLimitations            []string `json:"known_limitations,omitempty"`
}

type publicSectorProfilesResponse struct {
	SchemaVersion    string                `json:"schema_version"`
	Profiles         []publicSectorProfile `json:"profiles,omitempty"`
	MaturityMapRef   string                `json:"maturity_map_ref,omitempty"`
	DecisionGuideRef string                `json:"decision_guide_ref,omitempty"`
	Limitations      []string              `json:"limitations,omitempty"`
}

type publicDeploymentDecisionRow struct {
	DecisionID                string   `json:"decision_id"`
	DecisionPoint             string   `json:"decision_point"`
	PreferredOption           string   `json:"preferred_option"`
	Options                   []string `json:"options,omitempty"`
	ChooseWhen                []string `json:"choose_when,omitempty"`
	AvoidWhen                 []string `json:"avoid_when,omitempty"`
	RequiredContracts         []string `json:"required_contracts,omitempty"`
	ReferenceArchitectureRefs []string `json:"reference_architecture_refs,omitempty"`
	MaturityLevelRefs         []string `json:"maturity_level_refs,omitempty"`
	Limitations               []string `json:"limitations,omitempty"`
}

type publicDeploymentDecisionMatrixResponse struct {
	SchemaVersion string                        `json:"schema_version"`
	Rows          []publicDeploymentDecisionRow `json:"rows,omitempty"`
	Limitations   []string                      `json:"limitations,omitempty"`
}

func (s server) publicReferenceArchitecturesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicReferenceArchitectures())
}

func (s server) publicMaturityMapHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicMaturityMap())
}

func (s server) publicDecisionGuidesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicDecisionGuides())
}

func (s server) publicSectorProfilesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicSectorProfiles())
}

func (s server) publicDeploymentDecisionMatrixHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicDeploymentDecisionMatrix())
}

func buildPublicReferenceArchitectures() publicReferenceArchitecturesResponse {
	return publicReferenceArchitecturesResponse{
		SchemaVersion: publicReferenceArchitecturesSchema,
		Architectures: []publicReferenceArchitecture{
			{
				ArchitectureID:      "regulated-saas",
				DisplayName:         "Regulated SaaS reference architecture",
				SectorProfile:       "regulated_saas",
				DeploymentProfile:   "managed multi-tenant cloud",
				TrustBoundaries:     []string{"customer-facing ingress boundary", "build/deploy trust boundary", "runtime evidence boundary", "public publication boundary"},
				CriticalServices:    []string{"deploy-gate", "audit-writer", "runtime-agent", "attestation-verifier"},
				TopologySummary:     []string{"central control plane with bounded publication surfaces", "tenant-scoped runtime and validation evidence"},
				HandoffModel:        []string{"sealed handoff for incidents and audit exports", "customer-safe trust bundle for bounded publication"},
				ValidationModel:     []string{"strict regression and release validation before promotion", "seal-ready validation certificate for gated export"},
				RuntimeModel:        []string{"runtime rule packs and bounded response policy", "attestation-aware runtime posture linkage"},
				GovernanceOverlays:  []string{"enterprise identity fabric", "approval/autonomy split", "public claims governance before publication"},
				IntegrationPatterns: []string{"ITSM draft-before-write", "bounded SIEM enrichment", "customer-safe trust publication"},
				Assumptions:         []string{"connected control plane available", "operator-owned governance workflow", "customer-safe publication allowed by policy"},
				KnownLimitations:    []string{"Does not imply formal sector certification or legal attestation by itself."},
			},
			{
				ArchitectureID:      "fintech-multi-region",
				DisplayName:         "Fintech multi-region reference architecture",
				SectorProfile:       "fintech_multi_region",
				DeploymentProfile:   "active-active regional trust cells",
				TrustBoundaries:     []string{"regional runtime boundary", "cross-region federation boundary", "payment-service clearance boundary"},
				CriticalServices:    []string{"audit-writer", "runtime-agent", "federation-manager", "attestation-verifier"},
				TopologySummary:     []string{"regional isolation with bounded proof exchange", "local override remains authoritative in every region"},
				HandoffModel:        []string{"sealed handoff for regulated incident exchange", "cross-region proof acceptance with local freshness checks"},
				ValidationModel:     []string{"compatibility validation before region-wide policy shifts", "strict release validation per region"},
				RuntimeModel:        []string{"bounded runtime hardening for payments and settlement workloads", "proof-based regional trust posture rollups"},
				GovernanceOverlays:  []string{"clearance and approval mapping for privileged payment workflows", "regional override and distrust model"},
				IntegrationPatterns: []string{"bounded partner proof exchange", "connector health visibility for enterprise integrations"},
				Assumptions:         []string{"regional operators exist", "trust anchors are regionally managed", "cross-region proofs remain policy-gated"},
				KnownLimitations:    []string{"Does not promise automatic cross-region convergence or shared global authority."},
			},
			{
				ArchitectureID:      "air-gapped-regulated",
				DisplayName:         "Air-gapped regulated environment reference architecture",
				SectorProfile:       "air_gapped_regulated",
				DeploymentProfile:   "disconnected or delayed-sync execution",
				TrustBoundaries:     []string{"offline verification boundary", "local trust anchor boundary", "delayed sync boundary"},
				CriticalServices:    []string{"audit-writer", "runtime-agent", "validation-harness"},
				TopologySummary:     []string{"local-only evidence and verifier surfaces", "delayed-sync semantics explicitly visible"},
				HandoffModel:        []string{"sealed handoff as primary transfer format", "offline verifier guide and reference pack for auditors"},
				ValidationModel:     []string{"local validation certificates retained with short, bounded export path", "compatibility validation for offline policy changes"},
				RuntimeModel:        []string{"bounded runtime posture without dependency on live external control plane", "degraded mode remains explicit"},
				GovernanceOverlays:  []string{"manual approval and local operator accountability", "claims discipline for disconnected deployments"},
				IntegrationPatterns: []string{"offline ITSM evidence handoff", "no dependency on live remote proof verification"},
				Assumptions:         []string{"local trust anchors managed on site", "operators can perform offline verification", "sync is delayed and visible"},
				KnownLimitations:    []string{"No claim of instant federation, live remote revocation, or hidden convergence."},
			},
			{
				ArchitectureID:      "supplier-federation",
				DisplayName:         "Supplier federation reference architecture",
				SectorProfile:       "supplier_federation",
				DeploymentProfile:   "bounded partner proof exchange",
				TrustBoundaries:     []string{"supplier onboarding boundary", "sealed proof acceptance boundary", "disclosure-minimized export boundary"},
				CriticalServices:    []string{"audit-writer", "federation-manager", "attestation-verifier"},
				TopologySummary:     []string{"local policy over remote proofs", "supplier/customer-safe publication surfaces"},
				HandoffModel:        []string{"sealed handoff for partner proof exchange", "customer-safe bundle and auditor-safe proof profiles"},
				ValidationModel:     []string{"validation evidence shared selectively", "seal-ready certificate used as bounded trust signal"},
				RuntimeModel:        []string{"runtime state remains locally owned even when partner proofs are accepted", "stale/diverged partner state remains visible"},
				GovernanceOverlays:  []string{"revocation/distrust model for partners", "public claims constrained by disclosure policy"},
				IntegrationPatterns: []string{"supplier onboarding contract", "consortium/shared trust readiness hints"},
				Assumptions:         []string{"partners can supply sealed proof bundles", "local organization owns admissibility policy", "disclosure profiles are enforced"},
				KnownLimitations:    []string{"No implicit trust in remote peers and no requirement to share internal databases or private keys."},
			},
			{
				ArchitectureID:      "runtime-hardened-enterprise-cluster",
				DisplayName:         "Runtime-hardened enterprise cluster reference architecture",
				SectorProfile:       "runtime_hardened_enterprise_cluster",
				DeploymentProfile:   "enterprise cluster with strong runtime and enterprise integration",
				TrustBoundaries:     []string{"admission/runtime boundary", "incident collaboration boundary", "trust-hub governance boundary"},
				CriticalServices:    []string{"deploy-gate", "runtime-agent", "audit-writer", "policy-engine", "attestation-verifier"},
				TopologySummary:     []string{"runtime moat with explainable bounded enforcement", "integration-safe enterprise collaboration surfaces"},
				HandoffModel:        []string{"handoff for incident response and forensic preservation", "public sample/export path only where policy allows"},
				ValidationModel:     []string{"release, runtime, and compatibility validation gates", "policy-tuned approval flows"},
				RuntimeModel:        []string{"rule-pack registry, response policy, posture linkage, execution coverage expansion", "ambient/confidential readiness remains bounded"},
				GovernanceOverlays:  []string{"identity fabric integration", "trust-hub boundaries", "bounded clearance model"},
				IntegrationPatterns: []string{"ITSM lifecycle contracts", "SIEM/SOAR bounded sync", "incident collaboration flow"},
				Assumptions:         []string{"enterprise identity and workflow systems exist", "operators accept evidence-native governance model"},
				KnownLimitations:    []string{"Reference architecture remains a bounded deployment blueprint, not a claim of universal best fit."},
			},
		},
		DecisionGuideRefs: []string{
			"/v1/public/decision-guides",
			"/v1/public/maturity-map",
		},
		Limitations: []string{
			"Reference architectures are bounded deployment blueprints with explicit assumptions and limitations; they are not certifications or one-click deployment recipes.",
		},
	}
}

func buildPublicMaturityMap() publicMaturityMapResponse {
	return publicMaturityMapResponse{
		SchemaVersion: publicMaturityMapSchema,
		Levels: []publicMaturityLevel{
			{
				LevelID:                "level_1_foundation",
				DisplayName:            "Level 1 Foundation",
				RequiredCapabilities:   []string{"deploy-gate trust decisions", "audit evidence collection", "baseline incident and handoff surfaces"},
				EvidenceExpectations:   []string{"bounded audit lineage exists", "handoff and validation outputs are available in local surfaces"},
				GovernanceExpectations: []string{"basic approval roles exist", "policy ownership is identifiable"},
				RuntimeExpectations:    []string{"runtime findings are visible even if enforcement is still conservative"},
				ValidationExpectations: []string{"bounded validation runs exist"},
				B2BTrustExpectations:   []string{"none required"},
				MeasurableCriteria:     []string{"core trust control plane is active", "evidence and incident surfaces are queryable"},
			},
			{
				LevelID:                "level_2_hardened_operations",
				DisplayName:            "Level 2 Hardened Operations",
				RequiredCapabilities:   []string{"sealed handoff", "federation hardening", "strict validation harness", "runtime response policy"},
				EvidenceExpectations:   []string{"handoff verification is local and replayable", "validation outputs are seal-ready"},
				GovernanceExpectations: []string{"forensic-first and approval/autonomy split are explicit"},
				RuntimeExpectations:    []string{"runtime rule packs and explainability contracts are active"},
				ValidationExpectations: []string{"strict validation certificate semantics are stable"},
				B2BTrustExpectations:   []string{"optional bounded partner proof acceptance"},
				MeasurableCriteria:     []string{"critical runtime and validation contracts are explainable and test-backed"},
			},
			{
				LevelID:                "level_3_execution_and_enterprise",
				DisplayName:            "Level 3 Execution and Enterprise",
				RequiredCapabilities:   []string{"execution coverage expansion", "identity fabric integration", "ITSM lifecycle contracts", "incident collaboration"},
				EvidenceExpectations:   []string{"hybrid/VM/ephemeral evidence is retained in bounded form"},
				GovernanceExpectations: []string{"enterprise integration safety model is documented", "connector health is visible"},
				RuntimeExpectations:    []string{"attestation-aware posture linkage and runtime boundaries are explicit"},
				ValidationExpectations: []string{"validation outputs connect to enterprise workflows"},
				B2BTrustExpectations:   []string{"optional customer-safe or partner-safe export views"},
				MeasurableCriteria:     []string{"enterprise workflow embedding is in place", "execution substrates beyond default Kubernetes model are covered"},
			},
			{
				LevelID:                "level_4_trust_exchange",
				DisplayName:            "Level 4 Trust Exchange",
				RequiredCapabilities:   []string{"supplier onboarding", "sealed proof acceptance", "disclosure-minimized exchange", "customer trust bundles"},
				EvidenceExpectations:   []string{"publicly bounded proof exchange and customer-safe trust narratives exist"},
				GovernanceExpectations: []string{"partner distrust/revocation path is explicit"},
				RuntimeExpectations:    []string{"local runtime and partner trust remain clearly separated"},
				ValidationExpectations: []string{"validation outputs are interpretable in partner-facing exchange"},
				B2BTrustExpectations:   []string{"partner proof exchange is local-policy-first", "consortium readiness is documented"},
				MeasurableCriteria:     []string{"B2B exchange can happen without sharing private keys or internal data stores"},
			},
			{
				LevelID:                "level_5_public_verifiability",
				DisplayName:            "Level 5 Public Verifiability",
				RequiredCapabilities:   []string{"public verification specs", "public schema exports", "reference verifier pack", "reference architectures", "benchmark and claims discipline"},
				EvidenceExpectations:   []string{"third parties can verify bounded public artifacts without trusting internal implementation details"},
				GovernanceExpectations: []string{"public claims governance and publication discipline exist"},
				RuntimeExpectations:    []string{"runtime claims exposed publicly stay bounded by measured or published semantics"},
				ValidationExpectations: []string{"validation certificate semantics remain public and stable"},
				B2BTrustExpectations:   []string{"public trust exchange semantics are bounded and verifiable"},
				MeasurableCriteria:     []string{"public verifier inputs exist", "public maturity criteria are measurable and not tied only to feature purchase"},
				Limitations:            []string{"This level reflects public verifiability discipline, not an assertion of formal certification."},
			},
		},
		Limitations: []string{
			"Maturity levels are bounded adoption markers tied to capability and evidence expectations; they are not marketing tiers or legal attestations.",
		},
	}
}

func buildPublicDecisionGuides() publicDecisionGuidesResponse {
	return publicDecisionGuidesResponse{
		SchemaVersion: publicDecisionGuidesSchema,
		Guides: []publicDecisionGuide{
			{
				GuideID:                   "sealed_only_vs_broader_exchange",
				Question:                  "When should an organization use sealed-only exchange instead of broader evidence exchange?",
				UseWhen:                   []string{"partner only needs verifier inputs and bounded trust narrative", "disclosure minimization is a hard requirement", "local policy must stay fully authoritative"},
				AvoidWhen:                 []string{"deep bilateral troubleshooting requires richer evidence", "both sides already operate under a stronger shared disclosure agreement"},
				RecommendationSummary:     "Default to sealed-only exchange and add broader evidence sharing only when governance and disclosure assumptions are explicit.",
				Tradeoffs:                 []string{"sealed-only improves minimization but may slow deep investigations", "broader exchange improves operator context but increases disclosure surface"},
				RelatedContracts:          []string{"/v1/public/specs/handoff", "/v1/public/specs/proof-verification", "/v1/b2b/disclosure-profiles"},
				ReferenceArchitectureRefs: []string{"supplier-federation", "air-gapped-regulated"},
			},
			{
				GuideID:                   "require_validation_gates",
				Question:                  "When should validation gates be required before promotion or publication?",
				UseWhen:                   []string{"runtime or policy changes affect critical workloads", "regulated change control is required", "public or partner-facing trust claims depend on the result"},
				AvoidWhen:                 []string{"change scope is explicitly advisory-only and outside promotion path"},
				RecommendationSummary:     "Require validation gates for promoted, partner-facing, or governance-relevant changes; keep advisory projections clearly marked.",
				Tradeoffs:                 []string{"more validation increases confidence and evidence quality", "more validation adds operational time and compute overhead"},
				RelatedContracts:          []string{"/v1/public/specs/validation-certificate", "/v1/runtime/response-policy", "/v1/public/maturity-map"},
				ReferenceArchitectureRefs: []string{"regulated-saas", "fintech-multi-region", "runtime-hardened-enterprise-cluster"},
			},
			{
				GuideID:                   "stronger_runtime_hardening",
				Question:                  "When should stronger runtime hardening or lower-latency enforcement be enabled?",
				UseWhen:                   []string{"workload handles sensitive trust-critical functions", "blast radius of compromise is high", "operator team accepts bounded forensic-first response model"},
				AvoidWhen:                 []string{"runtime profile is poorly understood", "false-positive tolerance is extremely low and not yet measured"},
				RecommendationSummary:     "Enable stronger runtime hardening only after response policy, forensic-first semantics, and runtime boundaries are explicit and tested.",
				Tradeoffs:                 []string{"higher protection may increase operational tuning effort", "conservative posture reduces disruption but may leave more manual review"},
				RelatedContracts:          []string{"/v1/runtime/response-policy", "/v1/runtime/boundaries", "/v1/runtime/posture-linkage"},
				ReferenceArchitectureRefs: []string{"runtime-hardened-enterprise-cluster", "fintech-multi-region"},
			},
			{
				GuideID:                   "hybrid_or_air_gapped",
				Question:                  "When should hybrid or air-gapped execution models be adopted?",
				UseWhen:                   []string{"connectivity is constrained or policy-limited", "local trust anchors are required", "delayed sync is acceptable and visible"},
				AvoidWhen:                 []string{"workflow depends on immediate remote coordination or instant cross-org convergence"},
				RecommendationSummary:     "Adopt hybrid or air-gapped execution only if offline verification, delayed sync semantics, and degraded-mode evidence are explicit.",
				Tradeoffs:                 []string{"offline posture reduces dependency on live control planes", "offline posture increases operational burden for sync, distribution, and review"},
				RelatedContracts:          []string{"/v1/execution/coverage/matrix", "/v1/public/verifier/offline-guide", "/v1/public/reference-architectures"},
				ReferenceArchitectureRefs: []string{"air-gapped-regulated"},
			},
			{
				GuideID:                   "b2b_proof_exchange_adoption",
				Question:                  "When should B2B proof exchange be enabled for partners or suppliers?",
				UseWhen:                   []string{"supplier/customer trust needs repeatable verification", "local admissibility policy and disclosure profiles are defined", "sealed proof acceptance and revocation semantics are in place"},
				AvoidWhen:                 []string{"partner onboarding is ad hoc", "distrust/revocation path is undefined"},
				RecommendationSummary:     "Enable B2B proof exchange only after supplier onboarding, disclosure profiles, and local-policy-first acceptance are explicit.",
				Tradeoffs:                 []string{"partner proof exchange improves portability and verification", "it introduces governance overhead around trust anchors, freshness, and revocation"},
				RelatedContracts:          []string{"/v1/b2b/suppliers/onboarding", "/v1/b2b/sealed-proof/acceptance", "/v1/public/specs/federation-proof-exchange"},
				ReferenceArchitectureRefs: []string{"supplier-federation", "fintech-multi-region"},
			},
		},
		Limitations: []string{
			"Decision guides are bounded operational guides tied to ChangeLock contracts; they do not replace organization-specific architecture review or legal/compliance assessment.",
		},
	}
}

func buildPublicSectorProfiles() publicSectorProfilesResponse {
	return publicSectorProfilesResponse{
		SchemaVersion:    publicSectorProfilesSchema,
		MaturityMapRef:   "/v1/public/maturity-map",
		DecisionGuideRef: "/v1/public/decision-guides",
		Profiles: []publicSectorProfile{
			{
				SectorProfileID:             "regulated_saas",
				DisplayName:                 "Regulated SaaS",
				RecommendedArchitectureRefs: []string{"regulated-saas", "runtime-hardened-enterprise-cluster"},
				TrustPriorities:             []string{"customer-safe publication controls", "release and runtime evidence continuity", "approval lineage for public claims"},
				RequiredContracts:           []string{"/v1/public/specs/handoff", "/v1/public/specs/validation-certificate", "/v1/integrations/itsm-lifecycle", "/v1/runtime/response-policy"},
				GovernanceExpectations:      []string{"public claims governance exists", "approval/autonomy split remains explicit", "identity-driven approval lineage is retained"},
				DeploymentAssumptions:       []string{"connected control plane is acceptable", "tenant-scoped evidence retention exists", "customer-safe export policy is enforced"},
				KnownLimitations:            []string{"This profile does not imply sector certification or legal attestation by itself."},
			},
			{
				SectorProfileID:             "fintech_multi_region",
				DisplayName:                 "Fintech multi-region",
				RecommendedArchitectureRefs: []string{"fintech-multi-region", "runtime-hardened-enterprise-cluster"},
				TrustPriorities:             []string{"regional isolation with local overrides", "payment-service runtime hardening", "bounded cross-region proof portability"},
				RequiredContracts:           []string{"/v1/public/specs/federation-proof-exchange", "/v1/runtime/posture-linkage", "/v1/trust-hub/clearance", "/v1/b2b/sealed-proof/acceptance"},
				GovernanceExpectations:      []string{"regional distrust path exists", "clearance issuance remains explainable and time-bounded", "connector health stays visible"},
				DeploymentAssumptions:       []string{"operators are regionally distributed", "regional trust anchors are managed locally", "policy remains locally enforceable in every region"},
				KnownLimitations:            []string{"This profile does not promise automatic regional convergence or shared global trust authority."},
			},
			{
				SectorProfileID:             "air_gapped_regulated",
				DisplayName:                 "Air-gapped regulated environment",
				RecommendedArchitectureRefs: []string{"air-gapped-regulated"},
				TrustPriorities:             []string{"offline verification", "local trust anchor ownership", "delayed-sync visibility"},
				RequiredContracts:           []string{"/v1/execution/coverage/matrix", "/v1/public/verifier/offline-guide", "/v1/public/verifier/reference-pack"},
				GovernanceExpectations:      []string{"manual approval remains possible", "local operator accountability is explicit", "publication claims stay conservative"},
				DeploymentAssumptions:       []string{"operators can move sealed bundles offline", "sync is delayed and visible", "external connectivity is intermittent or prohibited"},
				KnownLimitations:            []string{"No claim of live remote revocation, hidden convergence, or permanently current external posture."},
			},
			{
				SectorProfileID:             "supplier_federation",
				DisplayName:                 "Supplier federation",
				RecommendedArchitectureRefs: []string{"supplier-federation"},
				TrustPriorities:             []string{"local admissibility over remote proof", "disclosure minimization", "partner distrust and revocation"},
				RequiredContracts:           []string{"/v1/b2b/suppliers/onboarding", "/v1/b2b/sealed-proof/acceptance", "/v1/b2b/disclosure-profiles", "/v1/public/specs/proof-verification"},
				GovernanceExpectations:      []string{"partner trust-anchor registration is audited", "customer-safe and auditor-safe export profiles are separated", "rejection reasons remain explainable"},
				DeploymentAssumptions:       []string{"partners can deliver sealed proof bundles", "local policy owns final admissibility", "proof freshness is checked locally"},
				KnownLimitations:            []string{"This profile does not remove the need for bilateral governance or disclosure agreements."},
			},
			{
				SectorProfileID:             "runtime_hardened_enterprise_cluster",
				DisplayName:                 "Runtime-hardened enterprise cluster",
				RecommendedArchitectureRefs: []string{"runtime-hardened-enterprise-cluster", "regulated-saas"},
				TrustPriorities:             []string{"runtime explainability", "forensic-first response", "enterprise workflow embedding"},
				RequiredContracts:           []string{"/v1/runtime/rule-packs", "/v1/runtime/response-policy", "/v1/runtime/boundaries", "/v1/integrations/safety"},
				GovernanceExpectations:      []string{"enterprise integration safety remains explicit", "trust-hub boundaries are visible", "runtime response remains bounded rather than destructive by default"},
				DeploymentAssumptions:       []string{"enterprise identity and workflow systems exist", "operators accept evidence-native governance flow", "runtime-agent deployment is feasible"},
				KnownLimitations:            []string{"This profile is a bounded architecture blueprint, not a universal best-fit deployment recipe."},
			},
		},
		Limitations: []string{
			"Sector profiles map priorities and contracts to deployment profiles; they do not replace organization-specific threat modeling, legal review, or certification programs.",
		},
	}
}

func buildPublicDeploymentDecisionMatrix() publicDeploymentDecisionMatrixResponse {
	return publicDeploymentDecisionMatrixResponse{
		SchemaVersion: publicDecisionMatrixSchema,
		Rows: []publicDeploymentDecisionRow{
			{
				DecisionID:                "exchange_scope",
				DecisionPoint:             "Choose sealed-only exchange or broader evidence exchange.",
				PreferredOption:           "sealed_only",
				Options:                   []string{"sealed_only", "bounded_richer_exchange"},
				ChooseWhen:                []string{"disclosure minimization is mandatory", "partner only needs verifier inputs", "local policy must stay clearly authoritative"},
				AvoidWhen:                 []string{"deep bilateral investigation requires richer context and governance already permits it"},
				RequiredContracts:         []string{"/v1/public/specs/handoff", "/v1/public/specs/proof-verification", "/v1/b2b/disclosure-profiles"},
				ReferenceArchitectureRefs: []string{"supplier-federation", "air-gapped-regulated"},
				MaturityLevelRefs:         []string{"level_4_trust_exchange", "level_5_public_verifiability"},
			},
			{
				DecisionID:                "validation_gate_strength",
				DecisionPoint:             "Choose advisory validation or required validation gate before promotion/publication.",
				PreferredOption:           "required_gate_for_promoted_changes",
				Options:                   []string{"advisory_only", "required_gate_for_promoted_changes"},
				ChooseWhen:                []string{"change affects critical workloads", "public or partner-facing claims depend on the result", "regulated change control is expected"},
				AvoidWhen:                 []string{"change stays outside promotion path and is explicitly advisory-only"},
				RequiredContracts:         []string{"/v1/public/specs/validation-certificate", "/v1/validation/executions", "/v1/runtime/response-policy"},
				ReferenceArchitectureRefs: []string{"regulated-saas", "fintech-multi-region", "runtime-hardened-enterprise-cluster"},
				MaturityLevelRefs:         []string{"level_2_hardened_operations", "level_3_execution_and_enterprise"},
			},
			{
				DecisionID:                "runtime_hardening_profile",
				DecisionPoint:             "Choose conservative or stronger runtime hardening profile.",
				PreferredOption:           "bounded_stronger_hardening_after_tuning",
				Options:                   []string{"conservative_detect_first", "bounded_stronger_hardening_after_tuning"},
				ChooseWhen:                []string{"blast radius is high", "runtime response policy is explicit", "false-positive behavior has been measured"},
				AvoidWhen:                 []string{"runtime profile is poorly understood", "forensic-first path is not ready"},
				RequiredContracts:         []string{"/v1/runtime/response-policy", "/v1/runtime/boundaries", "/v1/runtime/posture-linkage"},
				ReferenceArchitectureRefs: []string{"runtime-hardened-enterprise-cluster", "fintech-multi-region"},
				MaturityLevelRefs:         []string{"level_2_hardened_operations", "level_3_execution_and_enterprise"},
			},
			{
				DecisionID:                "execution_connectivity_model",
				DecisionPoint:             "Choose connected, hybrid, or air-gapped execution posture.",
				PreferredOption:           "hybrid_or_air_gapped_only_with_offline_verify",
				Options:                   []string{"connected", "hybrid", "air_gapped"},
				ChooseWhen:                []string{"connectivity is constrained by policy or environment", "local trust anchors are required", "delayed sync is acceptable and visible"},
				AvoidWhen:                 []string{"workflow depends on immediate remote coordination or hidden convergence"},
				RequiredContracts:         []string{"/v1/execution/coverage/matrix", "/v1/public/verifier/offline-guide", "/v1/public/reference-architectures"},
				ReferenceArchitectureRefs: []string{"air-gapped-regulated", "regulated-saas"},
				MaturityLevelRefs:         []string{"level_3_execution_and_enterprise", "level_5_public_verifiability"},
			},
			{
				DecisionID:                "partner_trust_exchange",
				DecisionPoint:             "Choose whether to enable B2B proof exchange and customer-safe publication.",
				PreferredOption:           "enable_after_onboarding_and_disclosure_policy",
				Options:                   []string{"local_only", "partner_proof_exchange", "partner_and_customer_safe_publication"},
				ChooseWhen:                []string{"supplier onboarding is explicit", "disclosure profiles are enforced", "rejection and revocation semantics are defined"},
				AvoidWhen:                 []string{"partner trust anchors are unmanaged", "customer-safe narrative would over-disclose or over-claim"},
				RequiredContracts:         []string{"/v1/b2b/suppliers/onboarding", "/v1/b2b/customer-bundles", "/v1/public/specs/federation-proof-exchange"},
				ReferenceArchitectureRefs: []string{"supplier-federation", "regulated-saas"},
				MaturityLevelRefs:         []string{"level_4_trust_exchange", "level_5_public_verifiability"},
			},
		},
		Limitations: []string{
			"Decision matrix rows are bounded guidance tied to current ChangeLock contracts; they do not replace organization-specific architecture review, incident history, or legal/compliance obligations.",
		},
	}
}
