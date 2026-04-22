package ecosystem

import "strings"

const (
	ScopeInternal = "internal"
	ScopeTenant   = "tenant"
	ScopePartner  = "partner"
	ScopePublic   = "public"
	ScopeVerifier = "verifier"

	SignalClassObservedFact      = "observed_fact"
	SignalClassDerivedRelevance  = "derived_relevance"
	SignalClassRecommendation    = "recommendation"
	SignalClassVerifierBacked    = "verifier_backed"
	SignalClassReviewRequirement = "review_required"

	StatusAdvisory       = "advisory"
	StatusVerifierBacked = "verifier_backed"
	StatusReviewRequired = "review_required"

	MutationNever      = "never"
	MutationReviewOnly = "review_only"

	EntryGateStateReady      = "phase7_entry_gate_ready"
	EntryGateStateIncomplete = "phase7_entry_gate_incomplete"

	FoundationStateActive     = "phase7_contract_foundation_active"
	FoundationStateIncomplete = "phase7_contract_foundation_incomplete"

	DeveloperPresenceStateActive     = "developer_presence_active"
	DeveloperPresenceStatePartial    = "developer_presence_partial"
	DeveloperPresenceStateIncomplete = "developer_presence_incomplete"

	OSSPresenceStateActive     = "oss_trust_presence_active"
	OSSPresenceStatePartial    = "oss_trust_presence_partial"
	OSSPresenceStateIncomplete = "oss_trust_presence_incomplete"

	DistributionPresenceStateActive     = "distribution_presence_active"
	DistributionPresenceStatePartial    = "distribution_presence_partial"
	DistributionPresenceStateIncomplete = "distribution_presence_incomplete"

	Phase7StateIncomplete  = "phase7_core_incomplete"
	Phase7StateSubstantial = "phase7_core_substantially_ready"
	Phase7StateActive      = "phase7_core_ecosystem_presence_active"
)

type EntryGate struct {
	CurrentState         string   `json:"current_state"`
	CanonicalWorkspace   string   `json:"canonical_workspace"`
	CarryOverLimitations []string `json:"carry_over_limitations,omitempty"`
	CarryOverDebt        []string `json:"carry_over_debt,omitempty"`
	ScopeBoundaries      []string `json:"scope_boundaries,omitempty"`
	ContractRefs         []string `json:"contract_refs,omitempty"`
}

type SignalContract struct {
	SignalID              string   `json:"signal_id"`
	SurfaceID             string   `json:"surface_id"`
	Source                string   `json:"source"`
	SignalClass           string   `json:"signal_class"`
	Scope                 string   `json:"scope"`
	Status                string   `json:"status"`
	FreshnessWindow       string   `json:"freshness_window"`
	Owner                 string   `json:"owner"`
	EvidenceRefs          []string `json:"evidence_refs,omitempty"`
	CompatibilityVersion  string   `json:"compatibility_version"`
	MutationCapability    string   `json:"mutation_capability"`
	SupersededSemantics   string   `json:"superseded_semantics"`
	RevokedSemantics      string   `json:"revoked_semantics"`
	DegradedState         string   `json:"degraded_state"`
	StaleState            string   `json:"stale_state"`
	ManualOverridePolicy  string   `json:"manual_override_policy,omitempty"`
	PublicationBoundaries []string `json:"publication_boundaries,omitempty"`
}

type AuthoritySurface struct {
	SurfaceID                string   `json:"surface_id"`
	Scope                    string   `json:"scope"`
	PresentsObservedFact     bool     `json:"presents_observed_fact"`
	PresentsDerivedRelevance bool     `json:"presents_derived_relevance"`
	PresentsRecommendation   bool     `json:"presents_recommendation"`
	PresentsVerifierBacked   bool     `json:"presents_verifier_backed_signal"`
	RequiresReview           bool     `json:"requires_review"`
	CanTriggerMutation       bool     `json:"can_trigger_mutation"`
	Publishers               []string `json:"publishers,omitempty"`
	Approvers                []string `json:"approvers,omitempty"`
	Superseders              []string `json:"superseders,omitempty"`
	Revokers                 []string `json:"revokers,omitempty"`
}

type FailSafeContract struct {
	SurfaceID            string   `json:"surface_id"`
	DegradedBehavior     string   `json:"degraded_behavior"`
	StaleBehavior        string   `json:"stale_behavior"`
	ManualOverridePolicy string   `json:"manual_override_policy"`
	DisablePath          string   `json:"disable_path"`
	KillSwitchOwner      string   `json:"kill_switch_owner"`
	RecoveryBehavior     string   `json:"recovery_behavior"`
	RecoveryVerification []string `json:"recovery_verification,omitempty"`
}

type PerformanceBudget struct {
	SurfaceID         string `json:"surface_id"`
	BudgetName        string `json:"budget_name"`
	TargetMS          int    `json:"target_ms"`
	BlockingPath      bool   `json:"blocking_path"`
	MeasurementSource string `json:"measurement_source"`
}

type SLOSpec struct {
	SurfaceID           string   `json:"surface_id"`
	AvailabilityTarget  string   `json:"availability_target"`
	FreshnessTarget     string   `json:"freshness_target"`
	ErrorBudgetPolicy   string   `json:"error_budget_policy"`
	AlertThresholds     []string `json:"alert_thresholds,omitempty"`
	CompletenessMonitor string   `json:"completeness_monitor"`
	NoiseBudgetPolicy   string   `json:"noise_budget_policy"`
}

type CompatibilityContract struct {
	SurfaceID                string   `json:"surface_id"`
	CurrentVersion           string   `json:"current_version"`
	BackwardCompatibility    []string `json:"backward_compatibility,omitempty"`
	DeprecationPolicy        []string `json:"deprecation_policy,omitempty"`
	MigrationGuidance        []string `json:"migration_guidance,omitempty"`
	BreakingChangeDisclosure []string `json:"breaking_change_disclosure,omitempty"`
}

type AbuseControl struct {
	SurfaceID        string   `json:"surface_id"`
	Threats          []string `json:"threats,omitempty"`
	Controls         []string `json:"controls,omitempty"`
	IncidentHandling []string `json:"incident_handling,omitempty"`
}

type RolloutContract struct {
	SurfaceID                string   `json:"surface_id"`
	Stages                   []string `json:"stages,omitempty"`
	CanaryCohort             string   `json:"canary_cohort"`
	RollbackConditions       []string `json:"rollback_conditions,omitempty"`
	DisableConditions        []string `json:"disable_conditions,omitempty"`
	ReleaseOwner             string   `json:"release_owner"`
	PostRollbackVerification []string `json:"post_rollback_verification,omitempty"`
}

type DataBoundary struct {
	SurfaceID          string   `json:"surface_id"`
	TenantConfidential []string `json:"tenant_confidential,omitempty"`
	InternalOnly       []string `json:"internal_only,omitempty"`
	PartnerVisible     []string `json:"partner_visible,omitempty"`
	PublicExportable   []string `json:"public_exportable,omitempty"`
	VerifierExportable []string `json:"verifier_exportable,omitempty"`
	AggregateOnly      []string `json:"aggregate_only,omitempty"`
	RedactedByDefault  bool     `json:"redacted_by_default"`
	RetentionPolicy    []string `json:"retention_policy,omitempty"`
	OffboardingPolicy  []string `json:"offboarding_policy,omitempty"`
}

type Coverage struct {
	SignalContracts        int `json:"signal_contracts"`
	AuthoritySurfaces      int `json:"authority_surfaces"`
	FailSafeContracts      int `json:"fail_safe_contracts"`
	PerformanceBudgets     int `json:"performance_budgets"`
	ObservabilitySLOs      int `json:"observability_slos"`
	CompatibilityContracts int `json:"compatibility_contracts"`
	AbuseControls          int `json:"abuse_controls"`
	RolloutContracts       int `json:"rollout_contracts"`
	DataBoundaries         int `json:"data_boundaries"`
}

type surfaceCoverage struct {
	hasSignal        bool
	hasAuthority     bool
	hasFailSafe      bool
	hasPerformance   bool
	hasSLO           bool
	hasCompatibility bool
	hasAbuseControl  bool
	hasRollout       bool
	hasDataBoundary  bool
}

type coreSurfaceRequirement struct {
	requireSignal        bool
	requireAuthority     bool
	requireFailSafe      bool
	requirePerformance   bool
	requireSLO           bool
	requireCompatibility bool
	requireAbuseControl  bool
	requireRollout       bool
	requireDataBoundary  bool
}

var phase7CoreSurfaceRequirements = map[string]coreSurfaceRequirement{
	"developer.ide_plugin": {
		requireSignal:        true,
		requireAuthority:     true,
		requireFailSafe:      true,
		requirePerformance:   true,
		requireSLO:           true,
		requireCompatibility: true,
		requireRollout:       true,
	},
	"developer.local_validation": {
		requireSignal:        true,
		requireAuthority:     true,
		requireFailSafe:      true,
		requirePerformance:   true,
		requireSLO:           true,
		requireCompatibility: true,
	},
	"developer.pre_commit_hook": {
		requireSignal:        true,
		requireAuthority:     true,
		requireFailSafe:      true,
		requirePerformance:   true,
		requireSLO:           true,
		requireCompatibility: true,
		requireAbuseControl:  true,
		requireRollout:       true,
	},
	"oss.observation_pipeline": {
		requireSignal:        true,
		requireAuthority:     true,
		requireFailSafe:      true,
		requirePerformance:   true,
		requireSLO:           true,
		requireCompatibility: true,
		requireAbuseControl:  true,
		requireRollout:       true,
	},
	"oss.claim_pipeline": {
		requireSignal:        true,
		requireAuthority:     true,
		requireFailSafe:      true,
		requirePerformance:   true,
		requireSLO:           true,
		requireCompatibility: true,
	},
	"distribution.marketplace_deployment": {
		requireSignal:        true,
		requireAuthority:     true,
		requireFailSafe:      true,
		requirePerformance:   true,
		requireSLO:           true,
		requireCompatibility: true,
		requireAbuseControl:  true,
		requireRollout:       true,
		requireDataBoundary:  true,
	},
	"distribution.msp_operator": {
		requireSignal:        true,
		requireAuthority:     true,
		requireFailSafe:      true,
		requirePerformance:   true,
		requireSLO:           true,
		requireCompatibility: true,
		requireAbuseControl:  true,
		requireRollout:       true,
		requireDataBoundary:  true,
	},
	"distribution.partner_export": {
		requireSignal:        true,
		requireAuthority:     true,
		requireFailSafe:      true,
		requirePerformance:   true,
		requireSLO:           true,
		requireCompatibility: true,
		requireAbuseControl:  true,
		requireRollout:       true,
		requireDataBoundary:  true,
	},
}

var phase7CoreSurfacesByGroup = map[string][]string{
	"developer": {
		"developer.ide_plugin",
		"developer.local_validation",
		"developer.pre_commit_hook",
	},
	"oss": {
		"oss.observation_pipeline",
		"oss.claim_pipeline",
	},
	"distribution": {
		"distribution.marketplace_deployment",
		"distribution.msp_operator",
		"distribution.partner_export",
	},
}

func isDeferredExpandedSurface(surfaceID string) bool {
	switch strings.TrimSpace(surfaceID) {
	case "oss.remediation_pr":
		return true
	default:
		return false
	}
}

func EntryGateBaseline() EntryGate {
	return EntryGate{
		CurrentState:       EntryGateStateReady,
		CanonicalWorkspace: "",
		CarryOverLimitations: []string{
			"Phase 7 core slice reuses the Phase 1-6 audit and evidence spine and does not create a new canonical truth store.",
			"Phase 8 authority claims, certification posture, and regulator-facing authority remain out of scope.",
		},
		CarryOverDebt: []string{
			"Expanded automated PR orchestration remains deferred outside the initial Phase 7 core pass.",
			"Integrity-as-a-Service packaging remains deferred outside the initial Phase 7 core pass.",
		},
		ScopeBoundaries: []string{
			"IDE and local developer signals remain advisory or review-required and never become implicit production truth.",
			"Observation pipeline remains distinct from claim pipeline across OSS and partner-facing signals.",
		},
		ContractRefs: []string{
			"phase7.signal_contract_matrix",
			"phase7.authority_surface_matrix",
			"phase7.fail_safe_contracts",
			"phase7.performance_and_slo_baseline",
			"phase7.export_boundary_matrix",
		},
	}
}

func SignalContracts() []SignalContract {
	return []SignalContract{
		{
			SignalID:             "developer.ide_trust_advisory",
			SurfaceID:            "developer.ide_plugin",
			Source:               "ide_plugin",
			SignalClass:          SignalClassDerivedRelevance,
			Scope:                ScopeTenant,
			Status:               StatusAdvisory,
			FreshnessWindow:      "15m",
			Owner:                "developer_experience",
			EvidenceRefs:         []string{"audit://phase3/vulnerability_relevance", "audit://phase5/command_center"},
			CompatibilityVersion: "7.0",
			MutationCapability:   MutationNever,
			SupersededSemantics:  "Superseded by fresher relevance or reviewed trust context for the same dependency subject.",
			RevokedSemantics:     "Revoked when plugin context loses subject identity or upstream evidence cannot be refreshed.",
			DegradedState:        "ide_signal_unavailable",
			StaleState:           "ide_context_stale",
			ManualOverridePolicy: "Developer may suppress repeated advisory prompts locally without mutating canonical trust.",
			PublicationBoundaries: []string{
				"tenant_visible",
				"never_public_truth",
			},
		},
		{
			SignalID:             "developer.local_validation_projection",
			SurfaceID:            "developer.local_validation",
			Source:               "local_cli",
			SignalClass:          SignalClassObservedFact,
			Scope:                ScopeTenant,
			Status:               StatusAdvisory,
			FreshnessWindow:      "run_scoped",
			Owner:                "developer_experience",
			EvidenceRefs:         []string{"audit://phase5/preview", "audit://phase5/check"},
			CompatibilityVersion: "7.0",
			MutationCapability:   MutationNever,
			SupersededSemantics:  "Superseded by a newer local validation run for the same declared subject and config set.",
			RevokedSemantics:     "Revoked when local inputs diverge from the inspected config or subject bundle.",
			DegradedState:        "local_validation_degraded",
			StaleState:           "local_validation_stale",
			ManualOverridePolicy: "Local validation can be skipped only with explicit developer acknowledgement and bounded output remains visible.",
			PublicationBoundaries: []string{
				"tenant_visible",
				"simulation_not_production_truth",
			},
		},
		{
			SignalID:             "developer.vex_relevance_context",
			SurfaceID:            "developer.ide_plugin",
			Source:               "phase3_intelligence_projection",
			SignalClass:          SignalClassDerivedRelevance,
			Scope:                ScopeTenant,
			Status:               StatusAdvisory,
			FreshnessWindow:      "30m",
			Owner:                "developer_experience",
			EvidenceRefs:         []string{"audit://phase3/grounded_query", "audit://phase3/supply_chain_pattern"},
			CompatibilityVersion: "7.0",
			MutationCapability:   MutationNever,
			SupersededSemantics:  "Superseded by a reviewed VEX candidate or fresher same-subject relevance evidence.",
			RevokedSemantics:     "Revoked when same-subject evidence is no longer bounded or confidence falls below review threshold.",
			DegradedState:        "vex_context_degraded",
			StaleState:           "vex_context_stale",
			ManualOverridePolicy: "Editors may hide stale context locally, but the stale marker remains visible when reopened.",
			PublicationBoundaries: []string{
				"tenant_visible",
				"reviewed_claims_only_public",
			},
		},
		{
			SignalID:             "developer.pre_commit_dependency_trust",
			SurfaceID:            "developer.pre_commit_hook",
			Source:               "local_hook",
			SignalClass:          SignalClassRecommendation,
			Scope:                ScopeTenant,
			Status:               StatusReviewRequired,
			FreshnessWindow:      "10m",
			Owner:                "developer_experience",
			EvidenceRefs:         []string{"audit://phase5/check", "audit://phase2/runtime_posture"},
			CompatibilityVersion: "7.0",
			MutationCapability:   MutationNever,
			SupersededSemantics:  "Superseded by the next pre-commit evaluation or by an explicit bounded override record.",
			RevokedSemantics:     "Revoked when hook context is stale, incomplete, or missing same-subject evidence.",
			DegradedState:        "pre_commit_check_degraded",
			StaleState:           "pre_commit_check_stale",
			ManualOverridePolicy: "Blocking behavior requires explicit developer override reason; override does not mutate production policy.",
			PublicationBoundaries: []string{
				"tenant_visible",
				"never_public_truth",
			},
		},
		{
			SignalID:             "oss.registry_provenance_observation",
			SurfaceID:            "oss.observation_pipeline",
			Source:               "registry_connector",
			SignalClass:          SignalClassObservedFact,
			Scope:                ScopeVerifier,
			Status:               StatusReviewRequired,
			FreshnessWindow:      "6h",
			Owner:                "oss_trust_ops",
			EvidenceRefs:         []string{"audit://phase6/transparency_anchor", "audit://phase3/supply_chain_pattern"},
			CompatibilityVersion: "7.0",
			MutationCapability:   MutationNever,
			SupersededSemantics:  "Superseded by the next normalized registry observation for the same package lineage.",
			RevokedSemantics:     "Revoked when provenance cannot be revalidated or replay/staleness checks fail.",
			DegradedState:        "registry_signal_unavailable",
			StaleState:           "registry_signal_stale",
			ManualOverridePolicy: "Observation remains candidate-only and cannot be promoted without review.",
			PublicationBoundaries: []string{
				"verifier_visible_candidate_only",
				"never_reviewed_claim_by_default",
			},
		},
		{
			SignalID:             "oss.reviewed_trust_claim",
			SurfaceID:            "oss.claim_pipeline",
			Source:               "reviewed_claim_pipeline",
			SignalClass:          SignalClassVerifierBacked,
			Scope:                ScopePublic,
			Status:               StatusVerifierBacked,
			FreshnessWindow:      "30d",
			Owner:                "oss_trust_ops",
			EvidenceRefs:         []string{"audit://phase6/proof_portal", "audit://phase6/claims_summary"},
			CompatibilityVersion: "7.0",
			MutationCapability:   MutationNever,
			SupersededSemantics:  "Superseded only by a newer reviewed claim or a formal supersession record.",
			RevokedSemantics:     "Revoked through explicit review outcome; revocation remains verifier-visible.",
			DegradedState:        "reviewed_claim_degraded",
			StaleState:           "reviewed_claim_stale",
			ManualOverridePolicy: "No local override; claim changes require review and publish discipline.",
			PublicationBoundaries: []string{
				"public_or_partner_visible_after_review",
				"verifier_exportable",
			},
		},
		{
			SignalID:             "oss.remediation_pr_recommendation",
			SurfaceID:            "oss.remediation_pr",
			Source:               "pr_automation_controller",
			SignalClass:          SignalClassRecommendation,
			Scope:                ScopeTenant,
			Status:               StatusReviewRequired,
			FreshnessWindow:      "24h",
			Owner:                "oss_trust_ops",
			EvidenceRefs:         []string{"audit://phase3/strategic_assessment", "audit://phase5/explain"},
			CompatibilityVersion: "7.0",
			MutationCapability:   MutationReviewOnly,
			SupersededSemantics:  "Superseded by newer evidence or a more precise remediation recommendation.",
			RevokedSemantics:     "Revoked when confidence, evidence completeness, or throttle rules fail.",
			DegradedState:        "remediation_pr_degraded",
			StaleState:           "remediation_pr_stale",
			ManualOverridePolicy: "PR creation remains disabled by default until review gates, throttle, and disable path permit exposure.",
			PublicationBoundaries: []string{
				"tenant_visible",
				"review_required_before_mutation",
			},
		},
		{
			SignalID:             "distribution.marketplace_deployment_readiness",
			SurfaceID:            "distribution.marketplace_deployment",
			Source:               "marketplace_readiness_projection",
			SignalClass:          SignalClassObservedFact,
			Scope:                ScopePartner,
			Status:               StatusReviewRequired,
			FreshnessWindow:      "1h",
			Owner:                "platform_operations",
			EvidenceRefs:         []string{"audit://phase5/readiness", "audit://phase5/support_bundle"},
			CompatibilityVersion: "7.0",
			MutationCapability:   MutationNever,
			SupersededSemantics:  "Superseded by the next readiness and deploy profile check for the same marketplace environment.",
			RevokedSemantics:     "Revoked when profile detection, trust prerequisites, or rollout verification no longer hold.",
			DegradedState:        "marketplace_deployment_degraded",
			StaleState:           "marketplace_readiness_stale",
			ManualOverridePolicy: "Deployment can proceed only with explicit degraded acknowledgement; no silent green path exists.",
			PublicationBoundaries: []string{
				"partner_visible",
				"aggregate_only_public",
			},
		},
		{
			SignalID:             "distribution.msp_tenant_isolation_posture",
			SurfaceID:            "distribution.msp_operator",
			Source:               "msp_isolation_projection",
			SignalClass:          SignalClassVerifierBacked,
			Scope:                ScopePartner,
			Status:               StatusVerifierBacked,
			FreshnessWindow:      "12h",
			Owner:                "platform_operations",
			EvidenceRefs:         []string{"audit://phase4/partner_trust", "audit://phase5/health_snapshot"},
			CompatibilityVersion: "7.0",
			MutationCapability:   MutationNever,
			SupersededSemantics:  "Superseded by the next bounded tenant-isolation verification run.",
			RevokedSemantics:     "Revoked immediately on cross-tenant boundary failure or audit-isolation regression.",
			DegradedState:        "tenant_isolation_degraded",
			StaleState:           "tenant_isolation_stale",
			ManualOverridePolicy: "No override for failing isolation posture; tenant boundary failure blocks partner-safe claims.",
			PublicationBoundaries: []string{
				"partner_visible",
				"verifier_exportable",
			},
		},
		{
			SignalID:             "distribution.partner_bounded_export",
			SurfaceID:            "distribution.partner_export",
			Source:               "partner_export_layer",
			SignalClass:          SignalClassVerifierBacked,
			Scope:                ScopePartner,
			Status:               StatusVerifierBacked,
			FreshnessWindow:      "24h",
			Owner:                "platform_operations",
			EvidenceRefs:         []string{"audit://phase6/verifier_sdk", "audit://phase5/support_bundle"},
			CompatibilityVersion: "7.0",
			MutationCapability:   MutationNever,
			SupersededSemantics:  "Superseded by a newer scoped export with the same tenant, audience, and contract version.",
			RevokedSemantics:     "Revoked on offboarding, scope reduction, or disclosure-boundary regression.",
			DegradedState:        "partner_export_degraded",
			StaleState:           "partner_export_stale",
			ManualOverridePolicy: "Export visibility is permissioned and revocable; no implicit cross-tenant merge is allowed.",
			PublicationBoundaries: []string{
				"partner_visible_only",
				"redacted_by_default",
			},
		},
	}
}

func AuthoritySurfaceMatrix() []AuthoritySurface {
	return []AuthoritySurface{
		{
			SurfaceID:                "developer.ide_plugin",
			Scope:                    ScopeTenant,
			PresentsDerivedRelevance: true,
			PresentsRecommendation:   true,
			RequiresReview:           false,
			CanTriggerMutation:       false,
			Publishers:               []string{"developer_experience"},
			Approvers:                []string{"developer"},
			Superseders:              []string{"developer_experience", "security_reviewer"},
			Revokers:                 []string{"developer_experience", "security_reviewer"},
		},
		{
			SurfaceID:              "developer.local_validation",
			Scope:                  ScopeTenant,
			PresentsObservedFact:   true,
			PresentsRecommendation: true,
			RequiresReview:         false,
			CanTriggerMutation:     false,
			Publishers:             []string{"developer_experience"},
			Approvers:              []string{"developer"},
			Superseders:            []string{"developer"},
			Revokers:               []string{"developer"},
		},
		{
			SurfaceID:              "developer.pre_commit_hook",
			Scope:                  ScopeTenant,
			PresentsObservedFact:   true,
			PresentsRecommendation: true,
			RequiresReview:         true,
			CanTriggerMutation:     false,
			Publishers:             []string{"developer_experience"},
			Approvers:              []string{"developer", "security_reviewer"},
			Superseders:            []string{"developer", "security_reviewer"},
			Revokers:               []string{"developer_experience"},
		},
		{
			SurfaceID:            "oss.observation_pipeline",
			Scope:                ScopeVerifier,
			PresentsObservedFact: true,
			RequiresReview:       true,
			CanTriggerMutation:   false,
			Publishers:           []string{"registry_connector"},
			Approvers:            []string{"oss_reviewer"},
			Superseders:          []string{"oss_reviewer"},
			Revokers:             []string{"oss_reviewer"},
		},
		{
			SurfaceID:              "oss.claim_pipeline",
			Scope:                  ScopePublic,
			PresentsVerifierBacked: true,
			RequiresReview:         true,
			CanTriggerMutation:     false,
			Publishers:             []string{"oss_reviewer"},
			Approvers:              []string{"oss_reviewer", "security_reviewer"},
			Superseders:            []string{"oss_reviewer"},
			Revokers:               []string{"oss_reviewer", "security_reviewer"},
		},
		{
			SurfaceID:              "oss.remediation_pr",
			Scope:                  ScopeTenant,
			PresentsRecommendation: true,
			RequiresReview:         true,
			CanTriggerMutation:     true,
			Publishers:             []string{"oss_pr_controller"},
			Approvers:              []string{"repo_owner", "security_reviewer"},
			Superseders:            []string{"repo_owner", "security_reviewer"},
			Revokers:               []string{"repo_owner", "security_reviewer"},
		},
		{
			SurfaceID:              "distribution.marketplace_deployment",
			Scope:                  ScopePartner,
			PresentsObservedFact:   true,
			PresentsRecommendation: true,
			RequiresReview:         true,
			CanTriggerMutation:     false,
			Publishers:             []string{"platform_operations"},
			Approvers:              []string{"tenant_operator", "platform_operations"},
			Superseders:            []string{"platform_operations"},
			Revokers:               []string{"platform_operations"},
		},
		{
			SurfaceID:              "distribution.msp_operator",
			Scope:                  ScopePartner,
			PresentsObservedFact:   true,
			PresentsVerifierBacked: true,
			RequiresReview:         true,
			CanTriggerMutation:     false,
			Publishers:             []string{"platform_operations"},
			Approvers:              []string{"msp_operator", "security_reviewer"},
			Superseders:            []string{"platform_operations"},
			Revokers:               []string{"platform_operations", "security_reviewer"},
		},
		{
			SurfaceID:              "distribution.partner_export",
			Scope:                  ScopePartner,
			PresentsVerifierBacked: true,
			RequiresReview:         true,
			CanTriggerMutation:     false,
			Publishers:             []string{"platform_operations"},
			Approvers:              []string{"tenant_operator", "partner_operator"},
			Superseders:            []string{"platform_operations"},
			Revokers:               []string{"platform_operations", "tenant_operator"},
		},
	}
}

func FailSafeContracts() []FailSafeContract {
	return []FailSafeContract{
		{
			SurfaceID:            "developer.ide_plugin",
			DegradedBehavior:     "Show signal unavailable or degraded; never paint stale dependency context green.",
			StaleBehavior:        "Keep advisory visible with stale marker and suppress blocking semantics.",
			ManualOverridePolicy: "Developer may dismiss repeated stale hints locally.",
			DisablePath:          "settings.ide_plugin.enabled=false",
			KillSwitchOwner:      "developer_experience",
			RecoveryBehavior:     "Refresh same-subject context and verify evidence refs before clearing degraded state.",
			RecoveryVerification: []string{"Confirm dependency subject identity.", "Confirm evidence refs refresh successfully."},
		},
		{
			SurfaceID:            "developer.local_validation",
			DegradedBehavior:     "Return bounded unknowns instead of pretending production equivalence.",
			StaleBehavior:        "Mark local projection stale when config or subject inputs drift.",
			ManualOverridePolicy: "Developer may rerun or skip locally with explicit acknowledgement.",
			DisablePath:          "changelock-cli --skip-local-validation",
			KillSwitchOwner:      "developer_experience",
			RecoveryBehavior:     "Regenerate validation context and re-evaluate same declared subject.",
			RecoveryVerification: []string{"Confirm config hash matches inspected state.", "Confirm runtime-only unknowns remain visible."},
		},
		{
			SurfaceID:            "developer.pre_commit_hook",
			DegradedBehavior:     "Downgrade to warning with reason code when local evidence is unavailable.",
			StaleBehavior:        "Treat stale hook data as review-required and keep override reason visible.",
			ManualOverridePolicy: "Override requires bounded reason string and local audit note.",
			DisablePath:          "CHGLOCK_PRECOMMIT_DISABLE=1",
			KillSwitchOwner:      "developer_experience",
			RecoveryBehavior:     "Re-run trust check with fresh context before re-enabling blocking behavior.",
			RecoveryVerification: []string{"Confirm hook latency stays within budget.", "Confirm override path remains auditable."},
		},
		{
			SurfaceID:            "oss.observation_pipeline",
			DegradedBehavior:     "Keep candidate observations unpublished and visible as incomplete ingestion only.",
			StaleBehavior:        "Mark connector state stale and suppress reviewed delivery for affected package lines.",
			ManualOverridePolicy: "No manual promotion from degraded candidate to reviewed claim.",
			DisablePath:          "oss.registry_connectors.enabled=false",
			KillSwitchOwner:      "oss_trust_ops",
			RecoveryBehavior:     "Replay intake through provenance, freshness, and abuse filters before publishing new candidates.",
			RecoveryVerification: []string{"Confirm replay and staleness checks pass.", "Confirm candidates remain candidate-only."},
		},
		{
			SurfaceID:            "oss.claim_pipeline",
			DegradedBehavior:     "Freeze reviewed publication and retain last explicit reviewed/superseded/revoked state.",
			StaleBehavior:        "Surface stale claim marker instead of claiming readiness.",
			ManualOverridePolicy: "Reviewers may revoke but not silently widen claim scope.",
			DisablePath:          "oss.claim_pipeline.publish=false",
			KillSwitchOwner:      "oss_trust_ops",
			RecoveryBehavior:     "Resume publication only after evidence completeness and review queue health are restored.",
			RecoveryVerification: []string{"Confirm candidate/reviewed separation still holds.", "Confirm revocation semantics remain visible."},
		},
		{
			SurfaceID:            "distribution.marketplace_deployment",
			DegradedBehavior:     "Return not-ready or degraded instead of a silent ready state.",
			StaleBehavior:        "Keep marketplace posture stale until readiness checks rerun successfully.",
			ManualOverridePolicy: "Tenant operator must acknowledge degraded deploy posture before continuing.",
			DisablePath:          "marketplace.readiness.enabled=false",
			KillSwitchOwner:      "platform_operations",
			RecoveryBehavior:     "Re-run profile detection, trust prerequisites, and readiness checks before clearing not-ready state.",
			RecoveryVerification: []string{"Confirm deploy profile matches installed topology.", "Confirm rollback path remains available."},
		},
		{
			SurfaceID:            "distribution.msp_operator",
			DegradedBehavior:     "Downgrade to failing isolation posture and stop tenant-safe automation.",
			StaleBehavior:        "Mark isolation verification stale and suppress partner-ready claims.",
			ManualOverridePolicy: "No override may bypass cross-tenant isolation failure.",
			DisablePath:          "msp.operator.enabled=false",
			KillSwitchOwner:      "platform_operations",
			RecoveryBehavior:     "Run tenant-boundary verification and audit isolation checks before re-enabling operations.",
			RecoveryVerification: []string{"Confirm per-tenant audit isolation.", "Confirm no cross-tenant export paths remain."},
		},
		{
			SurfaceID:            "distribution.partner_export",
			DegradedBehavior:     "Return redacted incomplete export rather than widening disclosure.",
			StaleBehavior:        "Flag export stale and suppress verifier-visible promotion.",
			ManualOverridePolicy: "Scope narrowing is allowed; scope widening requires new review.",
			DisablePath:          "partner.export.enabled=false",
			KillSwitchOwner:      "platform_operations",
			RecoveryBehavior:     "Re-issue scoped credentials and regenerate export under current boundary matrix.",
			RecoveryVerification: []string{"Confirm export remains partner-only.", "Confirm redacted-by-default fields stay redacted."},
		},
	}
}

func PerformanceBudgets() []PerformanceBudget {
	return []PerformanceBudget{
		{SurfaceID: "developer.ide_plugin", BudgetName: "ide_advisory_fetch_latency", TargetMS: 350, BlockingPath: false, MeasurementSource: "plugin telemetry"},
		{SurfaceID: "developer.local_validation", BudgetName: "local_validation_time", TargetMS: 1200, BlockingPath: false, MeasurementSource: "local cli timings"},
		{SurfaceID: "developer.pre_commit_hook", BudgetName: "pre_commit_max_blocking", TargetMS: 1500, BlockingPath: true, MeasurementSource: "hook runtime telemetry"},
		{SurfaceID: "oss.observation_pipeline", BudgetName: "registry_connector_refresh", TargetMS: 60000, BlockingPath: false, MeasurementSource: "connector ingest metrics"},
		{SurfaceID: "oss.claim_pipeline", BudgetName: "reviewed_claim_delivery", TargetMS: 2500, BlockingPath: false, MeasurementSource: "claim pipeline telemetry"},
		{SurfaceID: "distribution.marketplace_deployment", BudgetName: "marketplace_readiness_evaluation", TargetMS: 4000, BlockingPath: true, MeasurementSource: "deployment readiness telemetry"},
		{SurfaceID: "distribution.msp_operator", BudgetName: "msp_isolation_verification", TargetMS: 5000, BlockingPath: true, MeasurementSource: "tenant isolation telemetry"},
		{SurfaceID: "distribution.partner_export", BudgetName: "partner_export_generation", TargetMS: 1500, BlockingPath: false, MeasurementSource: "export generation telemetry"},
	}
}

func ObservabilitySLOs() []SLOSpec {
	return []SLOSpec{
		{
			SurfaceID:           "developer.ide_plugin",
			AvailabilityTarget:  "99.5%",
			FreshnessTarget:     "dependency context refreshed within 15m",
			ErrorBudgetPolicy:   "Suppress noisy advice before violating developer attention budget.",
			AlertThresholds:     []string{"stale signal exposure > 2%", "median advisory latency > 350ms"},
			CompletenessMonitor: "ide advisory evidence completeness",
			NoiseBudgetPolicy:   "no more than one repeated advisory per dependency per session without state change",
		},
		{
			SurfaceID:           "developer.pre_commit_hook",
			AvailabilityTarget:  "99.0%",
			FreshnessTarget:     "pre-commit evidence newer than 10m",
			ErrorBudgetPolicy:   "Fail open to warning when evidence is incomplete, never to silent pass.",
			AlertThresholds:     []string{"blocking time > 1500ms", "bypass rate > 10%"},
			CompletenessMonitor: "hook evidence completeness",
			NoiseBudgetPolicy:   "warn before block when confidence or freshness is below hard threshold",
		},
		{
			SurfaceID:           "developer.local_validation",
			AvailabilityTarget:  "99.5%",
			FreshnessTarget:     "local validation inputs remain tied to the current config and subject bundle",
			ErrorBudgetPolicy:   "Bounded unknowns remain visible instead of silent success when local context is incomplete.",
			AlertThresholds:     []string{"validation runtime > 1200ms", "stale local context exposure > 2%"},
			CompletenessMonitor: "local validation context completeness",
			NoiseBudgetPolicy:   "avoid repeated identical local warnings without state change",
		},
		{
			SurfaceID:           "oss.observation_pipeline",
			AvailabilityTarget:  "99.0%",
			FreshnessTarget:     "registry observations refreshed within 6h",
			ErrorBudgetPolicy:   "Candidate backlog may grow, but reviewed claims freeze before stale publication.",
			AlertThresholds:     []string{"stale connector exposure > 1%", "candidate backlog age > 24h"},
			CompletenessMonitor: "registry observation completeness",
			NoiseBudgetPolicy:   "suppress repeated identical observations pending review",
		},
		{
			SurfaceID:           "oss.claim_pipeline",
			AvailabilityTarget:  "99.5%",
			FreshnessTarget:     "reviewed claims published or marked stale within 30d window",
			ErrorBudgetPolicy:   "Freeze publication before letting stale reviewed claims appear active.",
			AlertThresholds:     []string{"reviewed claim stale exposure > 0", "claim publish latency > 2500ms"},
			CompletenessMonitor: "reviewed claim publication completeness",
			NoiseBudgetPolicy:   "group repeated reviewed-state updates for the same package lineage",
		},
		{
			SurfaceID:           "distribution.marketplace_deployment",
			AvailabilityTarget:  "99.5%",
			FreshnessTarget:     "readiness recomputed within 1h of install or upgrade",
			ErrorBudgetPolicy:   "Show degraded or not-ready before claiming ready.",
			AlertThresholds:     []string{"not-ready false negative > 0", "readiness latency > 4s"},
			CompletenessMonitor: "deployment readiness check completeness",
			NoiseBudgetPolicy:   "emit deploy warnings only when posture changes or blocker appears",
		},
		{
			SurfaceID:           "distribution.msp_operator",
			AvailabilityTarget:  "99.9%",
			FreshnessTarget:     "tenant isolation proof refreshed within 12h",
			ErrorBudgetPolicy:   "Cross-tenant failures exhaust budget immediately and trigger kill-switch review.",
			AlertThresholds:     []string{"tenant boundary probe > 0", "audit isolation drift > 0"},
			CompletenessMonitor: "tenant boundary verification completeness",
			NoiseBudgetPolicy:   "aggregate repeated isolation success signals; surface failures immediately",
		},
		{
			SurfaceID:           "distribution.partner_export",
			AvailabilityTarget:  "99.5%",
			FreshnessTarget:     "partner export contracts refreshed within 24h",
			ErrorBudgetPolicy:   "Suppress export widening and prefer incomplete redacted output on contract mismatch.",
			AlertThresholds:     []string{"partner export boundary mismatch > 0", "export generation latency > 1500ms"},
			CompletenessMonitor: "partner export boundary completeness",
			NoiseBudgetPolicy:   "aggregate export success signals and surface only contract or redaction changes",
		},
	}
}

func CompatibilityContracts() []CompatibilityContract {
	return []CompatibilityContract{
		{
			SurfaceID:                "developer.ide_plugin",
			CurrentVersion:           "7.0",
			BackwardCompatibility:    []string{"support previous advisory schema for one minor line", "unknown fields remain visible but non-authoritative"},
			DeprecationPolicy:        []string{"mark schema line deprecated before removal", "expose plugin upgrade notice before breaking change"},
			MigrationGuidance:        []string{"reinstall plugin against matching CLI version", "rerun local config inspection after upgrade"},
			BreakingChangeDisclosure: []string{"publish plugin compatibility note", "document noisy or blocking behavior changes before release"},
		},
		{
			SurfaceID:                "developer.local_validation",
			CurrentVersion:           "7.0",
			BackwardCompatibility:    []string{"accept previous preview reason codes for one minor line"},
			DeprecationPolicy:        []string{"retain bounded deprecated output mode during transition"},
			MigrationGuidance:        []string{"rerun local validation after CLI upgrade"},
			BreakingChangeDisclosure: []string{"document scope or uncertainty semantics before breaking output changes"},
		},
		{
			SurfaceID:                "developer.pre_commit_hook",
			CurrentVersion:           "7.0",
			BackwardCompatibility:    []string{"respect prior disable env var for one minor line"},
			DeprecationPolicy:        []string{"announce blocking threshold changes before removal"},
			MigrationGuidance:        []string{"refresh hook install after plugin upgrade"},
			BreakingChangeDisclosure: []string{"document new blocking conditions before they become default"},
		},
		{
			SurfaceID:                "oss.observation_pipeline",
			CurrentVersion:           "7.0",
			BackwardCompatibility:    []string{"keep candidate schema readable during one minor deprecation window"},
			DeprecationPolicy:        []string{"publish connector schema deprecations before removal"},
			MigrationGuidance:        []string{"replay staged observations against the new schema line"},
			BreakingChangeDisclosure: []string{"announce provenance and freshness rule changes before deployment"},
		},
		{
			SurfaceID:                "oss.claim_pipeline",
			CurrentVersion:           "7.0",
			BackwardCompatibility:    []string{"keep reviewed claim schema readable during one minor deprecation window"},
			DeprecationPolicy:        []string{"publish reviewed-claim schema deprecations before removal"},
			MigrationGuidance:        []string{"re-verify reviewed publication before switching schema line"},
			BreakingChangeDisclosure: []string{"announce reviewed claim state or scope changes before release"},
		},
		{
			SurfaceID:                "distribution.marketplace_deployment",
			CurrentVersion:           "7.0",
			BackwardCompatibility:    []string{"preserve readiness result fields during one minor transition"},
			DeprecationPolicy:        []string{"mark deprecated deployment profiles before removal"},
			MigrationGuidance:        []string{"rerun readiness and rollback checks after upgrade"},
			BreakingChangeDisclosure: []string{"document topology or profile assumptions before breaking support boundaries"},
		},
		{
			SurfaceID:                "distribution.msp_operator",
			CurrentVersion:           "7.0",
			BackwardCompatibility:    []string{"preserve per-tenant audit boundary fields through one minor line"},
			DeprecationPolicy:        []string{"announce export boundary changes before enforcement"},
			MigrationGuidance:        []string{"reissue tenant credentials during compatibility transition"},
			BreakingChangeDisclosure: []string{"disclose cross-tenant safety rule changes before release"},
		},
		{
			SurfaceID:                "distribution.partner_export",
			CurrentVersion:           "7.0",
			BackwardCompatibility:    []string{"keep verifier export schema readable during deprecation window"},
			DeprecationPolicy:        []string{"publish export version sunset dates"},
			MigrationGuidance:        []string{"rotate scoped credentials before old export line sunset"},
			BreakingChangeDisclosure: []string{"document export field removals and redaction changes ahead of rollout"},
		},
	}
}

func AbuseControls() []AbuseControl {
	return []AbuseControl{
		{
			SurfaceID:        "oss.observation_pipeline",
			Threats:          []string{"malicious_registry_signal_injection", "connector_poisoning", "replayed_external_signal_input"},
			Controls:         []string{"source trust tiering", "freshness and replay detection", "candidate-only staging before review"},
			IncidentHandling: []string{"freeze reviewed publication", "open abuse incident", "replay recent connector inputs"},
		},
		{
			SurfaceID:        "oss.remediation_pr",
			Threats:          []string{"pr_flood", "noisy_remediation", "weak_evidence_mutation"},
			Controls:         []string{"pr throttling", "confidence threshold", "disable path and kill-switch"},
			IncidentHandling: []string{"suspend PR creation", "review queue drain", "rerun evidence completeness checks"},
		},
		{
			SurfaceID:        "distribution.msp_operator",
			Threats:          []string{"tenant_boundary_probing", "cross_tenant_leakage", "over_broad_partner_automation"},
			Controls:         []string{"strict tenant isolation", "scoped credentials", "per-tenant audit isolation"},
			IncidentHandling: []string{"disable tenant-safe automation", "rotate credentials", "verify per-tenant audit logs"},
		},
		{
			SurfaceID:        "distribution.partner_export",
			Threats:          []string{"scope_widening_export", "stale_partner_signal_replay"},
			Controls:         []string{"redacted by default", "partner scope minimization", "export revocation support"},
			IncidentHandling: []string{"revoke export", "reissue scoped token", "rerun export boundary verification"},
		},
		{
			SurfaceID:        "distribution.marketplace_deployment",
			Threats:          []string{"false_green_readiness", "unsupported_topology_claim", "stale_marketplace_profile"},
			Controls:         []string{"profile detection", "readiness gate", "rollback and disable conditions"},
			IncidentHandling: []string{"mark deployment not-ready", "disable profile rollout", "rerun readiness verification"},
		},
		{
			SurfaceID:        "developer.pre_commit_hook",
			Threats:          []string{"misleading_local_confidence_signal", "bypass_without_reason"},
			Controls:         []string{"bounded uncertainty output", "explicit override reason", "hook telemetry"},
			IncidentHandling: []string{"disable blocking mode", "review repeated bypasses", "restore bounded warning output"},
		},
	}
}

func RolloutContracts() []RolloutContract {
	return []RolloutContract{
		{
			SurfaceID:                "developer.ide_plugin",
			Stages:                   []string{"internal", "canary", "tenant_opt_in", "default"},
			CanaryCohort:             "internal_repos",
			RollbackConditions:       []string{"advisory latency budget exceeded", "noise budget violated"},
			DisableConditions:        []string{"stale context spike", "plugin compatibility failure"},
			ReleaseOwner:             "developer_experience",
			PostRollbackVerification: []string{"confirm plugin disabled cleanly", "confirm local CLI still provides bounded fallback"},
		},
		{
			SurfaceID:                "developer.pre_commit_hook",
			Stages:                   []string{"warn_only", "opt_in_blocking", "default_blocking"},
			CanaryCohort:             "security_guarded_repos",
			RollbackConditions:       []string{"false positive complaints exceed threshold", "hook latency exceeds blocking budget"},
			DisableConditions:        []string{"evidence completeness drops below floor"},
			ReleaseOwner:             "developer_experience",
			PostRollbackVerification: []string{"confirm hook reverts to warning mode", "confirm override reasons remain logged"},
		},
		{
			SurfaceID:                "oss.observation_pipeline",
			Stages:                   []string{"shadow_ingest", "candidate_delivery", "reviewed_publication"},
			CanaryCohort:             "npm_and_pypi_subset",
			RollbackConditions:       []string{"candidate backlog age exceeds threshold", "connector poisoning suspected"},
			DisableConditions:        []string{"replay detection failure", "source trust tier regression"},
			ReleaseOwner:             "oss_trust_ops",
			PostRollbackVerification: []string{"confirm reviewed publication frozen", "confirm candidate-only mode remains explicit"},
		},
		{
			SurfaceID:                "distribution.marketplace_deployment",
			Stages:                   []string{"internal_profile", "partner_canary", "supported_marketplace_profile"},
			CanaryCohort:             "enterprise_default_profile",
			RollbackConditions:       []string{"readiness false green detected", "degraded deploy hidden from operator"},
			DisableConditions:        []string{"trust prerequisite regression", "rollback path unavailable"},
			ReleaseOwner:             "platform_operations",
			PostRollbackVerification: []string{"confirm deployment marked not-ready", "confirm previous profile guidance restored"},
		},
		{
			SurfaceID:                "distribution.msp_operator",
			Stages:                   []string{"internal_tenant_fleet", "msp_canary", "supported_msp_profile"},
			CanaryCohort:             "strict_isolation_tenants",
			RollbackConditions:       []string{"tenant isolation regression detected", "partner-safe automation exceeds tenant scope"},
			DisableConditions:        []string{"cross-tenant boundary probe detected", "per-tenant audit isolation incomplete"},
			ReleaseOwner:             "platform_operations",
			PostRollbackVerification: []string{"confirm tenant-safe automation disabled cleanly", "confirm tenant-boundary verification reruns successfully"},
		},
		{
			SurfaceID:                "distribution.partner_export",
			Stages:                   []string{"internal_export", "partner_canary", "supported_partner_profile"},
			CanaryCohort:             "scoped_verifier_exports",
			RollbackConditions:       []string{"export boundary mismatch", "partner scope escalation bug"},
			DisableConditions:        []string{"credential scoping failure", "redaction regression"},
			ReleaseOwner:             "platform_operations",
			PostRollbackVerification: []string{"confirm revoked export cannot be reused", "confirm redaction rules remain enforced"},
		},
	}
}

func DataBoundaries() []DataBoundary {
	return []DataBoundary{
		{
			SurfaceID:          "developer.ide_plugin",
			TenantConfidential: []string{"dependency subject identity", "workspace-local config hints"},
			InternalOnly:       []string{"plugin telemetry raw events"},
			AggregateOnly:      []string{"advisory adoption metrics"},
			RedactedByDefault:  true,
			RetentionPolicy:    []string{"session-local hints expire with workspace session", "telemetry aggregates retain no raw subject identifiers"},
			OffboardingPolicy:  []string{"clear local caches", "remove workspace-scoped plugin tokens"},
		},
		{
			SurfaceID:          "oss.claim_pipeline",
			PartnerVisible:     []string{"reviewed bounded trust claims"},
			PublicExportable:   []string{"reviewed claims with bounded semantics"},
			VerifierExportable: []string{"claim refs", "verification refs", "freshness refs"},
			AggregateOnly:      []string{"candidate review ratios"},
			RedactedByDefault:  true,
			RetentionPolicy:    []string{"candidate observations retain bounded review history", "public claims retain superseded or revoked markers"},
			OffboardingPolicy:  []string{"revoke unpublished candidate tokens", "retain published reviewed claims with explicit revocation semantics"},
		},
		{
			SurfaceID:          "distribution.marketplace_deployment",
			TenantConfidential: []string{"deployment topology specifics", "tenant config posture"},
			PartnerVisible:     []string{"bounded readiness verdict", "unsupported conditions"},
			AggregateOnly:      []string{"marketplace readiness completion rate"},
			RedactedByDefault:  true,
			RetentionPolicy:    []string{"retain readiness outputs for bounded support window", "aggregate install metrics only outside tenant scope"},
			OffboardingPolicy:  []string{"revoke marketplace tokens", "purge tenant-specific deployment posture exports"},
		},
		{
			SurfaceID:          "distribution.msp_operator",
			TenantConfidential: []string{"per-tenant audit records", "tenant proof exports"},
			InternalOnly:       []string{"operator control-plane notes"},
			PartnerVisible:     []string{"tenant-isolation verdicts", "scoped supportability status"},
			AggregateOnly:      []string{"fleet-wide readiness rates"},
			RedactedByDefault:  true,
			RetentionPolicy:    []string{"retain per-tenant audit isolation evidence per tenant policy", "aggregate fleet metrics without tenant identifiers"},
			OffboardingPolicy:  []string{"revoke tenant-scoped operator access", "confirm tenant export cleanup"},
		},
		{
			SurfaceID:          "distribution.partner_export",
			TenantConfidential: []string{"raw support bundle contents", "internal-only tenant notes"},
			PartnerVisible:     []string{"scoped verifier-facing export fields"},
			VerifierExportable: []string{"bounded proof bundles", "export contract version"},
			PublicExportable:   []string{},
			AggregateOnly:      []string{"export success rate"},
			RedactedByDefault:  true,
			RetentionPolicy:    []string{"bounded export retention per tenant contract", "revoked exports remain listed but unusable"},
			OffboardingPolicy:  []string{"revoke scoped credentials", "invalidate old export refs", "confirm partner cache purge instructions"},
		},
	}
}

func ContractsCoverage() Coverage {
	return Coverage{
		SignalContracts:        countCoreSignalContracts(),
		AuthoritySurfaces:      countCoreAuthoritySurfaces(),
		FailSafeContracts:      countCoreFailSafeContracts(),
		PerformanceBudgets:     countCorePerformanceBudgets(),
		ObservabilitySLOs:      countCoreObservabilitySLOs(),
		CompatibilityContracts: countCoreCompatibilityContracts(),
		AbuseControls:          countCoreAbuseControls(),
		RolloutContracts:       countCoreRolloutContracts(),
		DataBoundaries:         countCoreDataBoundaries(),
	}
}

func EvaluateFoundationState(coverage Coverage) string {
	if coverage.SignalContracts == 0 || coverage.AuthoritySurfaces == 0 || coverage.FailSafeContracts == 0 || coverage.PerformanceBudgets == 0 || coverage.ObservabilitySLOs == 0 || coverage.CompatibilityContracts == 0 || coverage.AbuseControls == 0 || coverage.RolloutContracts == 0 || coverage.DataBoundaries == 0 {
		return FoundationStateIncomplete
	}
	if evaluateFoundationStateForCoverageMap(buildSurfaceCoverageMap()) != FoundationStateActive {
		return FoundationStateIncomplete
	}
	return FoundationStateActive
}

func EvaluateDeveloperPresenceState() string {
	return evaluateGroupStateForCoverageMap("developer", buildSurfaceCoverageMap())
}

func EvaluateOSSPresenceState() string {
	return evaluateGroupStateForCoverageMap("oss", buildSurfaceCoverageMap())
}

func EvaluateDistributionPresenceState() string {
	return evaluateGroupStateForCoverageMap("distribution", buildSurfaceCoverageMap())
}

func EvaluatePhase7State(entryGateState, foundationState, developerState, ossState, distributionState string) string {
	if strings.TrimSpace(entryGateState) != EntryGateStateReady || strings.TrimSpace(foundationState) != FoundationStateActive {
		return Phase7StateIncomplete
	}
	developerState = strings.TrimSpace(developerState)
	ossState = strings.TrimSpace(ossState)
	distributionState = strings.TrimSpace(distributionState)
	if !isDeveloperPhase7State(developerState) || !isOSSPhase7State(ossState) || !isDistributionPhase7State(distributionState) {
		return Phase7StateIncomplete
	}
	if developerState == DeveloperPresenceStateIncomplete || ossState == OSSPresenceStateIncomplete || distributionState == DistributionPresenceStateIncomplete {
		return Phase7StateIncomplete
	}
	if developerState == DeveloperPresenceStatePartial || ossState == OSSPresenceStatePartial || distributionState == DistributionPresenceStatePartial {
		return Phase7StateSubstantial
	}
	return Phase7StateActive
}

func isDeveloperPhase7State(state string) bool {
	switch strings.TrimSpace(state) {
	case DeveloperPresenceStateActive, DeveloperPresenceStatePartial, DeveloperPresenceStateIncomplete:
		return true
	default:
		return false
	}
}

func isOSSPhase7State(state string) bool {
	switch strings.TrimSpace(state) {
	case OSSPresenceStateActive, OSSPresenceStatePartial, OSSPresenceStateIncomplete:
		return true
	default:
		return false
	}
}

func isDistributionPhase7State(state string) bool {
	switch strings.TrimSpace(state) {
	case DistributionPresenceStateActive, DistributionPresenceStatePartial, DistributionPresenceStateIncomplete:
		return true
	default:
		return false
	}
}

func SignalContractsForGroup(group string) []SignalContract {
	group = strings.TrimSpace(group)
	if group == "" {
		return nil
	}
	out := []SignalContract{}
	for _, item := range SignalContracts() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		if strings.HasPrefix(item.SignalID, group+".") {
			out = append(out, item)
		}
	}
	return out
}

func AuthoritySurfacesForGroup(group string) []AuthoritySurface {
	group = strings.TrimSpace(group)
	if group == "" {
		return nil
	}
	out := []AuthoritySurface{}
	for _, item := range AuthoritySurfaceMatrix() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		if strings.HasPrefix(item.SurfaceID, group+".") {
			out = append(out, item)
		}
	}
	return out
}

func FailSafeContractsForGroup(group string) []FailSafeContract {
	group = strings.TrimSpace(group)
	if group == "" {
		return nil
	}
	out := []FailSafeContract{}
	for _, item := range FailSafeContracts() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		if strings.HasPrefix(item.SurfaceID, group+".") {
			out = append(out, item)
		}
	}
	return out
}

func PerformanceBudgetsForGroup(group string) []PerformanceBudget {
	group = strings.TrimSpace(group)
	if group == "" {
		return nil
	}
	out := []PerformanceBudget{}
	for _, item := range PerformanceBudgets() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		if strings.HasPrefix(item.SurfaceID, group+".") {
			out = append(out, item)
		}
	}
	return out
}

func CompatibilityContractsForGroup(group string) []CompatibilityContract {
	group = strings.TrimSpace(group)
	if group == "" {
		return nil
	}
	out := []CompatibilityContract{}
	for _, item := range CompatibilityContracts() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		if strings.HasPrefix(item.SurfaceID, group+".") {
			out = append(out, item)
		}
	}
	return out
}

func AbuseControlsForGroup(group string) []AbuseControl {
	group = strings.TrimSpace(group)
	if group == "" {
		return nil
	}
	out := []AbuseControl{}
	for _, item := range AbuseControls() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		if strings.HasPrefix(item.SurfaceID, group+".") {
			out = append(out, item)
		}
	}
	return out
}

func RolloutContractsForGroup(group string) []RolloutContract {
	group = strings.TrimSpace(group)
	if group == "" {
		return nil
	}
	out := []RolloutContract{}
	for _, item := range RolloutContracts() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		if strings.HasPrefix(item.SurfaceID, group+".") {
			out = append(out, item)
		}
	}
	return out
}

func DataBoundariesForGroup(group string) []DataBoundary {
	group = strings.TrimSpace(group)
	if group == "" {
		return nil
	}
	out := []DataBoundary{}
	for _, item := range DataBoundaries() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		if strings.HasPrefix(item.SurfaceID, group+".") {
			out = append(out, item)
		}
	}
	return out
}

func countCoreSignalContracts() int {
	return countCoreItems(SignalContracts(), func(item SignalContract) string { return item.SurfaceID })
}

func countCoreAuthoritySurfaces() int {
	return countCoreItems(AuthoritySurfaceMatrix(), func(item AuthoritySurface) string { return item.SurfaceID })
}

func countCoreFailSafeContracts() int {
	return countCoreItems(FailSafeContracts(), func(item FailSafeContract) string { return item.SurfaceID })
}

func countCorePerformanceBudgets() int {
	return countCoreItems(PerformanceBudgets(), func(item PerformanceBudget) string { return item.SurfaceID })
}

func countCoreObservabilitySLOs() int {
	return countCoreItems(ObservabilitySLOs(), func(item SLOSpec) string { return item.SurfaceID })
}

func countCoreCompatibilityContracts() int {
	return countCoreItems(CompatibilityContracts(), func(item CompatibilityContract) string { return item.SurfaceID })
}

func countCoreAbuseControls() int {
	return countCoreItems(AbuseControls(), func(item AbuseControl) string { return item.SurfaceID })
}

func countCoreRolloutContracts() int {
	return countCoreItems(RolloutContracts(), func(item RolloutContract) string { return item.SurfaceID })
}

func countCoreDataBoundaries() int {
	return countCoreItems(DataBoundaries(), func(item DataBoundary) string { return item.SurfaceID })
}

func countCoreItems[T any](items []T, surfaceID func(T) string) int {
	count := 0
	for _, item := range items {
		if isDeferredExpandedSurface(surfaceID(item)) {
			continue
		}
		count++
	}
	return count
}

func buildSurfaceCoverageMap() map[string]surfaceCoverage {
	coverage := map[string]surfaceCoverage{}
	for _, item := range SignalContracts() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		current := coverage[item.SurfaceID]
		current.hasSignal = true
		coverage[item.SurfaceID] = current
	}
	for _, item := range AuthoritySurfaceMatrix() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		current := coverage[item.SurfaceID]
		current.hasAuthority = true
		coverage[item.SurfaceID] = current
	}
	for _, item := range FailSafeContracts() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		current := coverage[item.SurfaceID]
		current.hasFailSafe = true
		coverage[item.SurfaceID] = current
	}
	for _, item := range PerformanceBudgets() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		current := coverage[item.SurfaceID]
		current.hasPerformance = true
		coverage[item.SurfaceID] = current
	}
	for _, item := range ObservabilitySLOs() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		current := coverage[item.SurfaceID]
		current.hasSLO = true
		coverage[item.SurfaceID] = current
	}
	for _, item := range CompatibilityContracts() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		current := coverage[item.SurfaceID]
		current.hasCompatibility = true
		coverage[item.SurfaceID] = current
	}
	for _, item := range AbuseControls() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		current := coverage[item.SurfaceID]
		current.hasAbuseControl = true
		coverage[item.SurfaceID] = current
	}
	for _, item := range RolloutContracts() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		current := coverage[item.SurfaceID]
		current.hasRollout = true
		coverage[item.SurfaceID] = current
	}
	for _, item := range DataBoundaries() {
		if isDeferredExpandedSurface(item.SurfaceID) {
			continue
		}
		current := coverage[item.SurfaceID]
		current.hasDataBoundary = true
		coverage[item.SurfaceID] = current
	}
	return coverage
}

func evaluateFoundationStateForCoverageMap(coverage map[string]surfaceCoverage) string {
	for surfaceID, requirement := range phase7CoreSurfaceRequirements {
		if !surfaceSatisfiesRequirement(coverage[surfaceID], requirement) {
			return FoundationStateIncomplete
		}
	}
	return FoundationStateActive
}

func evaluateGroupStateForCoverageMap(group string, coverage map[string]surfaceCoverage) string {
	coreSurfaces := phase7CoreSurfacesByGroup[strings.TrimSpace(group)]
	if len(coreSurfaces) == 0 {
		return Phase7StateIncomplete
	}
	anyCovered := false
	allCovered := true
	for _, surfaceID := range coreSurfaces {
		surfaceCoverage := coverage[surfaceID]
		if surfaceHasAnyCoverage(surfaceCoverage) {
			anyCovered = true
		}
		if !surfaceSatisfiesRequirement(surfaceCoverage, phase7CoreSurfaceRequirements[surfaceID]) {
			allCovered = false
		}
	}
	if allCovered {
		switch group {
		case "developer":
			return DeveloperPresenceStateActive
		case "oss":
			return OSSPresenceStateActive
		case "distribution":
			return DistributionPresenceStateActive
		default:
			return Phase7StateIncomplete
		}
	}
	if anyCovered {
		switch group {
		case "developer":
			return DeveloperPresenceStatePartial
		case "oss":
			return OSSPresenceStatePartial
		case "distribution":
			return DistributionPresenceStatePartial
		default:
			return Phase7StateIncomplete
		}
	}
	switch group {
	case "developer":
		return DeveloperPresenceStateIncomplete
	case "oss":
		return OSSPresenceStateIncomplete
	case "distribution":
		return DistributionPresenceStateIncomplete
	default:
		return Phase7StateIncomplete
	}
}

func surfaceSatisfiesRequirement(coverage surfaceCoverage, requirement coreSurfaceRequirement) bool {
	if requirement.requireSignal && !coverage.hasSignal {
		return false
	}
	if requirement.requireAuthority && !coverage.hasAuthority {
		return false
	}
	if requirement.requireFailSafe && !coverage.hasFailSafe {
		return false
	}
	if requirement.requirePerformance && !coverage.hasPerformance {
		return false
	}
	if requirement.requireSLO && !coverage.hasSLO {
		return false
	}
	if requirement.requireCompatibility && !coverage.hasCompatibility {
		return false
	}
	if requirement.requireAbuseControl && !coverage.hasAbuseControl {
		return false
	}
	if requirement.requireRollout && !coverage.hasRollout {
		return false
	}
	if requirement.requireDataBoundary && !coverage.hasDataBoundary {
		return false
	}
	return true
}

func surfaceHasAnyCoverage(coverage surfaceCoverage) bool {
	return coverage.hasSignal || coverage.hasAuthority || coverage.hasFailSafe || coverage.hasPerformance || coverage.hasSLO || coverage.hasCompatibility || coverage.hasAbuseControl || coverage.hasRollout || coverage.hasDataBoundary
}
