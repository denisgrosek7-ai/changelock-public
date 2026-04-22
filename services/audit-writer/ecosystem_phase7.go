package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	ecosystemcore "github.com/denisgrosek/changelock/internal/ecosystem"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	phase7EntryGateSchema          = "7.ecosystem_entry_gate.v1"
	phase7ContractsSchema          = "7.ecosystem_contracts.v1"
	phase7DeveloperPresenceSchema  = "7.ecosystem_developer_presence.v1"
	phase7OSSNetworkSchema         = "7.ecosystem_oss_network.v1"
	phase7DistributionSchema       = "7.ecosystem_distribution_presence.v1"
	phase7ProofsSchema             = "7.ecosystem_phase7_proofs.v1"
	phase7CoverageScopeCorePass    = "core_pass"
	phase7DeferredExpandedScopeTag = "expanded_scope_deferred"
)

type phase7EntryGateResponse struct {
	SchemaVersion        string    `json:"schema_version"`
	GeneratedAt          time.Time `json:"generated_at"`
	CurrentState         string    `json:"current_state"`
	CanonicalWorkspace   string    `json:"canonical_workspace"`
	CarryOverLimitations []string  `json:"carry_over_limitations,omitempty"`
	CarryOverDebt        []string  `json:"carry_over_debt,omitempty"`
	ScopeBoundaries      []string  `json:"scope_boundaries,omitempty"`
	ContractRefs         []string  `json:"contract_refs,omitempty"`
	Limitations          []string  `json:"limitations,omitempty"`
}

type phase7ContractsResponse struct {
	SchemaVersion          string                                `json:"schema_version"`
	GeneratedAt            time.Time                             `json:"generated_at"`
	CurrentState           string                                `json:"current_state"`
	Coverage               ecosystemcore.Coverage                `json:"coverage"`
	SignalContracts        []ecosystemcore.SignalContract        `json:"signal_contracts,omitempty"`
	AuthoritySurfaces      []ecosystemcore.AuthoritySurface      `json:"authority_surfaces,omitempty"`
	FailSafeContracts      []ecosystemcore.FailSafeContract      `json:"fail_safe_contracts,omitempty"`
	PerformanceBudgets     []ecosystemcore.PerformanceBudget     `json:"performance_budgets,omitempty"`
	ObservabilitySLOs      []ecosystemcore.SLOSpec               `json:"observability_slos,omitempty"`
	CompatibilityContracts []ecosystemcore.CompatibilityContract `json:"compatibility_contracts,omitempty"`
	AbuseControls          []ecosystemcore.AbuseControl          `json:"abuse_controls,omitempty"`
	RolloutContracts       []ecosystemcore.RolloutContract       `json:"rollout_contracts,omitempty"`
	DataBoundaries         []ecosystemcore.DataBoundary          `json:"data_boundaries,omitempty"`
	Limitations            []string                              `json:"limitations,omitempty"`
}

type phase7DeveloperPresenceResponse struct {
	SchemaVersion        string    `json:"schema_version"`
	GeneratedAt          time.Time `json:"generated_at"`
	CurrentState         string    `json:"current_state"`
	IDEAdvisoryState     string    `json:"ide_advisory_state"`
	LocalValidationState string    `json:"local_validation_state"`
	InEditorVEXState     string    `json:"in_editor_vex_state"`
	PreCommitState       string    `json:"pre_commit_state"`
	AttentionBudgetState string    `json:"attention_budget_state"`
	OutputSemantics      []string  `json:"output_semantics,omitempty"`
	SignalRefs           []string  `json:"signal_refs,omitempty"`
	AuthorityRefs        []string  `json:"authority_refs,omitempty"`
	FailSafeRefs         []string  `json:"fail_safe_refs,omitempty"`
	PerformanceRefs      []string  `json:"performance_refs,omitempty"`
	CompatibilityRefs    []string  `json:"compatibility_refs,omitempty"`
	Limitations          []string  `json:"limitations,omitempty"`
}

type phase7PipelineSummary struct {
	CurrentState  string   `json:"current_state"`
	Stages        []string `json:"stages,omitempty"`
	BoundaryRules []string `json:"boundary_rules,omitempty"`
	ReviewStates  []string `json:"review_states,omitempty"`
}

type phase7RegistryConnector struct {
	RegistryID       string   `json:"registry_id"`
	CurrentState     string   `json:"current_state"`
	FreshnessState   string   `json:"freshness_state"`
	ReasonCodes      []string `json:"reason_codes,omitempty"`
	DegradedBehavior string   `json:"degraded_behavior"`
}

type phase7OSSTrustSignal struct {
	SignalID     string   `json:"signal_id"`
	CurrentState string   `json:"current_state"`
	ReviewState  string   `json:"review_state"`
	Semantics    []string `json:"semantics,omitempty"`
	EvidenceRefs []string `json:"evidence_refs,omitempty"`
}

type phase7OSSNetworkResponse struct {
	SchemaVersion       string                    `json:"schema_version"`
	GeneratedAt         time.Time                 `json:"generated_at"`
	CurrentState        string                    `json:"current_state"`
	ObservationPipeline phase7PipelineSummary     `json:"observation_pipeline"`
	ClaimPipeline       phase7PipelineSummary     `json:"claim_pipeline"`
	Connectors          []phase7RegistryConnector `json:"connectors,omitempty"`
	ReviewedSignals     []phase7OSSTrustSignal    `json:"reviewed_signals,omitempty"`
	AutomatedPRState    string                    `json:"automated_pr_state"`
	AbuseControlRefs    []string                  `json:"abuse_control_refs,omitempty"`
	RolloutRefs         []string                  `json:"rollout_refs,omitempty"`
	Limitations         []string                  `json:"limitations,omitempty"`
}

type phase7MarketplaceDeploymentSummary struct {
	CurrentState       string   `json:"current_state"`
	ProfileDetection   string   `json:"profile_detection"`
	ReadinessState     string   `json:"readiness_state"`
	DegradedBehavior   string   `json:"degraded_behavior"`
	RollbackBehavior   string   `json:"rollback_behavior"`
	ExportBoundaryRefs []string `json:"export_boundary_refs,omitempty"`
}

type phase7MSPIsolationSummary struct {
	CurrentState       string   `json:"current_state"`
	TenantIsolation    string   `json:"tenant_isolation"`
	AuditIsolation     string   `json:"audit_isolation"`
	AutomationScope    []string `json:"automation_scope,omitempty"`
	ExportBoundaryRefs []string `json:"export_boundary_refs,omitempty"`
}

type phase7PartnerSurfaceSummary struct {
	CurrentState      string   `json:"current_state"`
	Scope             string   `json:"scope"`
	AllowedOperations []string `json:"allowed_operations,omitempty"`
	ForbiddenOps      []string `json:"forbidden_operations,omitempty"`
	CredentialModel   []string `json:"credential_model,omitempty"`
}

type phase7DistributionResponse struct {
	SchemaVersion         string                             `json:"schema_version"`
	GeneratedAt           time.Time                          `json:"generated_at"`
	CurrentState          string                             `json:"current_state"`
	MarketplaceDeployment phase7MarketplaceDeploymentSummary `json:"marketplace_deployment"`
	MSPIsolation          phase7MSPIsolationSummary          `json:"msp_isolation"`
	PartnerSurface        phase7PartnerSurfaceSummary        `json:"partner_surface"`
	CompatibilityRefs     []string                           `json:"compatibility_refs,omitempty"`
	FailSafeRefs          []string                           `json:"fail_safe_refs,omitempty"`
	ExpandedScopeDeferred []string                           `json:"expanded_scope_deferred,omitempty"`
	Limitations           []string                           `json:"limitations,omitempty"`
}

type phase7ProofSection struct {
	CurrentState string   `json:"current_state"`
	KeyRefs      []string `json:"key_refs,omitempty"`
}

type phase7FoundationProofSection struct {
	CurrentState string                 `json:"current_state"`
	Coverage     ecosystemcore.Coverage `json:"coverage"`
	KeyRefs      []string               `json:"key_refs,omitempty"`
}

type phase7ProofsResponse struct {
	SchemaVersion         string                       `json:"schema_version"`
	GeneratedAt           time.Time                    `json:"generated_at"`
	CoverageScope         string                       `json:"coverage_scope"`
	CurrentState          string                       `json:"current_state"`
	EntryGate             phase7ProofSection           `json:"entry_gate"`
	ContractFoundation    phase7FoundationProofSection `json:"contract_foundation"`
	DeveloperPresence     phase7ProofSection           `json:"developer_presence"`
	OSSNetwork            phase7ProofSection           `json:"oss_network"`
	DistributionPresence  phase7ProofSection           `json:"distribution_presence"`
	ExpandedScopeDeferred []string                     `json:"expanded_scope_deferred,omitempty"`
	Limitations           []string                     `json:"limitations,omitempty"`
}

func (s server) phase7EntryGateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7EntryGate())
}

func (s server) phase7ContractsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7Contracts())
}

func (s server) phase7DeveloperPresenceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7DeveloperPresence())
}

func (s server) phase7OSSNetworkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7OSSNetwork())
}

func (s server) phase7DistributionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7Distribution())
}

func (s server) phase7ProofsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7Proofs())
}

func buildPhase7EntryGate() phase7EntryGateResponse {
	entryGate := ecosystemcore.EntryGateBaseline()
	return phase7EntryGateResponse{
		SchemaVersion:        phase7EntryGateSchema,
		GeneratedAt:          publicSampleTime(),
		CurrentState:         entryGate.CurrentState,
		CanonicalWorkspace:   phase7CanonicalWorkspace(),
		CarryOverLimitations: entryGate.CarryOverLimitations,
		CarryOverDebt:        entryGate.CarryOverDebt,
		ScopeBoundaries:      entryGate.ScopeBoundaries,
		ContractRefs:         entryGate.ContractRefs,
		Limitations: []string{
			"Phase 7 entry gate is a core-pass governance boundary and does not claim full expanded ecosystem completion.",
		},
	}
}

func phase7CanonicalWorkspace() string {
	if value := strings.TrimSpace(os.Getenv("CHANGELOCK_CANONICAL_WORKSPACE")); value != "" {
		return value
	}
	if cwd, err := os.Getwd(); err == nil {
		return cwd
	}
	return ""
}

func buildPhase7Contracts() phase7ContractsResponse {
	coverage := ecosystemcore.ContractsCoverage()
	return phase7ContractsResponse{
		SchemaVersion:          phase7ContractsSchema,
		GeneratedAt:            publicSampleTime(),
		CurrentState:           ecosystemcore.EvaluateFoundationState(coverage),
		Coverage:               coverage,
		SignalContracts:        ecosystemcore.SignalContracts(),
		AuthoritySurfaces:      ecosystemcore.AuthoritySurfaceMatrix(),
		FailSafeContracts:      ecosystemcore.FailSafeContracts(),
		PerformanceBudgets:     ecosystemcore.PerformanceBudgets(),
		ObservabilitySLOs:      ecosystemcore.ObservabilitySLOs(),
		CompatibilityContracts: ecosystemcore.CompatibilityContracts(),
		AbuseControls:          ecosystemcore.AbuseControls(),
		RolloutContracts:       ecosystemcore.RolloutContracts(),
		DataBoundaries:         ecosystemcore.DataBoundaries(),
		Limitations: []string{
			"Contracts surface defines core-pass ecosystem rules and does not widen authority beyond bounded reviewer, tenant, or partner scopes.",
			"Expanded automated PR and Integrity-as-a-Service delivery remain intentionally deferred outside the core pass.",
		},
	}
}

func buildPhase7DeveloperPresence() phase7DeveloperPresenceResponse {
	signals := ecosystemcore.SignalContractsForGroup("developer")
	authorities := ecosystemcore.AuthoritySurfacesForGroup("developer")
	failsafes := ecosystemcore.FailSafeContractsForGroup("developer")
	budgets := ecosystemcore.PerformanceBudgetsForGroup("developer")
	compat := ecosystemcore.CompatibilityContractsForGroup("developer")
	return phase7DeveloperPresenceResponse{
		SchemaVersion:        phase7DeveloperPresenceSchema,
		GeneratedAt:          publicSampleTime(),
		CurrentState:         ecosystemcore.EvaluateDeveloperPresenceState(),
		IDEAdvisoryState:     "ide_advisory_active",
		LocalValidationState: "local_validation_active",
		InEditorVEXState:     "in_editor_vex_context_active",
		PreCommitState:       "pre_commit_trust_check_active",
		AttentionBudgetState: "attention_budget_active",
		OutputSemantics: []string{
			ecosystemcore.SignalClassObservedFact,
			ecosystemcore.SignalClassDerivedRelevance,
			ecosystemcore.SignalClassRecommendation,
			"uncertainty",
		},
		SignalRefs:        phase7SignalIDs(signals),
		AuthorityRefs:     phase7AuthorityIDs(authorities),
		FailSafeRefs:      phase7FailSafeIDs(failsafes),
		PerformanceRefs:   phase7BudgetIDs(budgets),
		CompatibilityRefs: phase7CompatibilityIDs(compat),
		Limitations: []string{
			"IDE and local validation surfaces remain advisory or review-required helpers and never become canonical production truth.",
			"Local sandbox output remains explicit about runtime-only unknowns and does not claim production equivalence.",
		},
	}
}

func buildPhase7OSSNetwork() phase7OSSNetworkResponse {
	signals := ecosystemcore.SignalContractsForGroup("oss")
	abuseControls := ecosystemcore.AbuseControlsForGroup("oss")
	rollouts := ecosystemcore.RolloutContractsForGroup("oss")
	return phase7OSSNetworkResponse{
		SchemaVersion: phase7OSSNetworkSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  ecosystemcore.EvaluateOSSPresenceState(),
		ObservationPipeline: phase7PipelineSummary{
			CurrentState: "observation_pipeline_active",
			Stages:       []string{"intake", "normalization", "provenance_tagging", "freshness_marking", "abuse_filtering", "candidate_staging"},
			BoundaryRules: []string{
				"Observation pipeline never promotes reviewed trust claims on its own.",
				"Connector degradation leaves observations candidate-only and explicitly stale or unavailable.",
			},
			ReviewStates: []string{"candidate", "stale_candidate", "blocked_candidate"},
		},
		ClaimPipeline: phase7PipelineSummary{
			CurrentState: "claim_pipeline_active",
			Stages:       []string{"review", "evidence_binding", "publication", "supersession", "revocation"},
			BoundaryRules: []string{
				"Reviewed publication requires explicit review and evidence binding.",
				"Reviewed signals remain bounded, verifier-friendly, and revocable.",
			},
			ReviewStates: []string{"reviewed", "rejected", "superseded", "revoked"},
		},
		Connectors: []phase7RegistryConnector{
			{RegistryID: "npm", CurrentState: "connector_active", FreshnessState: "fresh", ReasonCodes: []string{"provenance_observation_ready"}, DegradedBehavior: "show candidate-only stale state if refresh or provenance verification fails"},
			{RegistryID: "pypi", CurrentState: "connector_active", FreshnessState: "fresh", ReasonCodes: []string{"provenance_observation_ready"}, DegradedBehavior: "show candidate-only stale state if refresh or provenance verification fails"},
			{RegistryID: "maven", CurrentState: "connector_active", FreshnessState: "fresh", ReasonCodes: []string{"provenance_observation_ready"}, DegradedBehavior: "show candidate-only stale state if refresh or provenance verification fails"},
		},
		ReviewedSignals: []phase7OSSTrustSignal{
			{
				SignalID:     "oss.reviewed_trust_claim",
				CurrentState: "reviewed_signal_active",
				ReviewState:  "reviewed",
				Semantics: []string{
					"bounded_trust_signal",
					"verifier_friendly",
					"not_a_marketing_quality_score",
				},
				EvidenceRefs: signalEvidenceRefs(signals, "oss.reviewed_trust_claim"),
			},
		},
		AutomatedPRState: phase7DeferredExpandedScopeTag,
		AbuseControlRefs: phase7AbuseIDs(abuseControls),
		RolloutRefs:      phase7RolloutIDs(rollouts),
		Limitations: []string{
			"Automated PR creation remains deferred outside the Phase 7 core pass.",
			"Community-assisted candidate intake remains bounded to observation staging until formal review discipline is widened.",
		},
	}
}

func buildPhase7Distribution() phase7DistributionResponse {
	failsafes := ecosystemcore.FailSafeContractsForGroup("distribution")
	compat := ecosystemcore.CompatibilityContractsForGroup("distribution")
	boundaries := ecosystemcore.DataBoundariesForGroup("distribution")
	return phase7DistributionResponse{
		SchemaVersion: phase7DistributionSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  ecosystemcore.EvaluateDistributionPresenceState(),
		MarketplaceDeployment: phase7MarketplaceDeploymentSummary{
			CurrentState:       "marketplace_deployment_active",
			ProfileDetection:   "environment_profile_detection_active",
			ReadinessState:     "deployment_readiness_active",
			DegradedBehavior:   "show not-ready or degraded instead of a silent ready state",
			RollbackBehavior:   "preserve rollback and disable path before declaring install ready",
			ExportBoundaryRefs: phase7BoundaryIDs(boundaries, "distribution.marketplace_deployment"),
		},
		MSPIsolation: phase7MSPIsolationSummary{
			CurrentState:       "msp_isolation_active",
			TenantIsolation:    "strict_tenant_isolation_verified",
			AuditIsolation:     "per_tenant_audit_isolation_verified",
			AutomationScope:    []string{"tenant_safe_read_only_exports", "scoped_supportability", "no_cross_tenant_mutation"},
			ExportBoundaryRefs: phase7BoundaryIDs(boundaries, "distribution.msp_operator"),
		},
		PartnerSurface: phase7PartnerSurfaceSummary{
			CurrentState:      "partner_surface_active",
			Scope:             ecosystemcore.ScopePartner,
			AllowedOperations: []string{"scoped_read", "verifier_friendly_export", "tenant_safe_hook"},
			ForbiddenOps:      []string{"cross_tenant_read", "implicit_orchestration", "broad_mutation_authority"},
			CredentialModel:   []string{"scoped_credentials", "lifecycle_safe_onboarding", "revocable_exports"},
		},
		CompatibilityRefs:     phase7CompatibilityIDs(compat),
		FailSafeRefs:          phase7FailSafeIDs(failsafes),
		ExpandedScopeDeferred: []string{"integrity_as_a_service_package", "broader_partner_write_api"},
		Limitations: []string{
			"Marketplace presence remains readiness-bound and does not imply click-and-forget production completion.",
			"Partner surface stays read/export scoped in the core pass and does not introduce broad mutation authority.",
		},
	}
}

func buildPhase7Proofs() phase7ProofsResponse {
	entry := buildPhase7EntryGate()
	contracts := buildPhase7Contracts()
	developer := buildPhase7DeveloperPresence()
	oss := buildPhase7OSSNetwork()
	distribution := buildPhase7Distribution()
	return phase7ProofsResponse{
		SchemaVersion: phase7ProofsSchema,
		GeneratedAt:   publicSampleTime(),
		CoverageScope: phase7CoverageScopeCorePass,
		CurrentState: ecosystemcore.EvaluatePhase7State(
			entry.CurrentState,
			contracts.CurrentState,
			developer.CurrentState,
			oss.CurrentState,
			distribution.CurrentState,
		),
		EntryGate: phase7ProofSection{
			CurrentState: entry.CurrentState,
			KeyRefs:      entry.ContractRefs,
		},
		ContractFoundation: phase7FoundationProofSection{
			CurrentState: contracts.CurrentState,
			Coverage:     contracts.Coverage,
			KeyRefs: []string{
				"phase7.signal_contract_matrix",
				"phase7.authority_surface_matrix",
				"phase7.fail_safe_contracts",
				"phase7.compatibility_deprecation_matrix",
				"phase7.export_boundary_matrix",
			},
		},
		DeveloperPresence: phase7ProofSection{
			CurrentState: developer.CurrentState,
			KeyRefs:      developer.SignalRefs,
		},
		OSSNetwork: phase7ProofSection{
			CurrentState: oss.CurrentState,
			KeyRefs: []string{
				"oss.observation_pipeline",
				"oss.claim_pipeline",
				"oss.reviewed_trust_claim",
			},
		},
		DistributionPresence: phase7ProofSection{
			CurrentState: distribution.CurrentState,
			KeyRefs: []string{
				"distribution.marketplace_deployment_readiness",
				"distribution.msp_tenant_isolation_posture",
				"distribution.partner_bounded_export",
			},
		},
		ExpandedScopeDeferred: []string{
			"automated_pr_discipline",
			"broader_partner_api",
			"integrity_as_a_service_package",
			"additional_registry_provider_coverage",
		},
		Limitations: []string{
			"Phase 7 proofs are core-pass scoped and do not claim full expanded ecosystem completion.",
			"Formal external authority, certification posture, and regulator-grade trust remain reserved for a later authority phase.",
		},
	}
}

func phase7SignalIDs(items []ecosystemcore.SignalContract) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, item.SignalID)
	}
	return out
}

func phase7AuthorityIDs(items []ecosystemcore.AuthoritySurface) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, item.SurfaceID)
	}
	return out
}

func phase7FailSafeIDs(items []ecosystemcore.FailSafeContract) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, item.SurfaceID)
	}
	return out
}

func phase7BudgetIDs(items []ecosystemcore.PerformanceBudget) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, item.SurfaceID+":"+item.BudgetName)
	}
	return out
}

func phase7CompatibilityIDs(items []ecosystemcore.CompatibilityContract) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, item.SurfaceID)
	}
	return out
}

func phase7AbuseIDs(items []ecosystemcore.AbuseControl) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, item.SurfaceID)
	}
	return out
}

func phase7RolloutIDs(items []ecosystemcore.RolloutContract) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, item.SurfaceID)
	}
	return out
}

func phase7BoundaryIDs(items []ecosystemcore.DataBoundary, surfaceID string) []string {
	out := []string{}
	for _, item := range items {
		if item.SurfaceID == surfaceID {
			out = append(out, item.SurfaceID)
		}
	}
	return out
}

func signalEvidenceRefs(items []ecosystemcore.SignalContract, signalID string) []string {
	for _, item := range items {
		if item.SignalID == signalID {
			return item.EvidenceRefs
		}
	}
	return nil
}
