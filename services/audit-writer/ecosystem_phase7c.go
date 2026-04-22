package main

import (
	"net/http"
	"time"

	ecosystemcore "github.com/denisgrosek/changelock/internal/ecosystem"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	phase7MarketplaceReadinessSchema = "7c.ecosystem_marketplace_readiness.v1"
	phase7MSPIsolationSchema         = "7c.ecosystem_msp_isolation.v1"
	phase7PartnerExportSchema        = "7c.ecosystem_partner_export.v1"

	phase7MarketplaceReadinessStateActive     = "marketplace_readiness_pack_active"
	phase7MarketplaceReadinessStatePartial    = "marketplace_readiness_pack_partial"
	phase7MarketplaceReadinessStateIncomplete = "marketplace_readiness_pack_incomplete"
	phase7MSPIsolationStateActive             = "msp_isolation_pack_active"
	phase7MSPIsolationStatePartial            = "msp_isolation_pack_partial"
	phase7MSPIsolationStateIncomplete         = "msp_isolation_pack_incomplete"
	phase7PartnerExportStateActive            = "partner_export_pack_active"
	phase7PartnerExportStatePartial           = "partner_export_pack_partial"
	phase7PartnerExportStateIncomplete        = "partner_export_pack_incomplete"
)

type phase7MarketplaceReadinessCheck struct {
	CheckID       string   `json:"check_id"`
	CurrentState  string   `json:"current_state"`
	Category      string   `json:"category"`
	Summary       string   `json:"summary"`
	ReasonCodes   []string `json:"reason_codes,omitempty"`
	EvidenceRefs  []string `json:"evidence_refs,omitempty"`
	BoundaryRules []string `json:"boundary_rules,omitempty"`
	Limitations   []string `json:"limitations,omitempty"`
}

type phase7MarketplaceReadinessResponse struct {
	SchemaVersion         string                            `json:"schema_version"`
	GeneratedAt           time.Time                         `json:"generated_at"`
	CurrentState          string                            `json:"current_state"`
	ProfileDetection      string                            `json:"profile_detection"`
	ReadinessState        string                            `json:"readiness_state"`
	TrustPrerequisites    string                            `json:"trust_prerequisites"`
	UpgradeBoundaryState  string                            `json:"upgrade_boundary_state"`
	UnsupportedConditions []string                          `json:"unsupported_conditions,omitempty"`
	Checks                []phase7MarketplaceReadinessCheck `json:"checks,omitempty"`
	ExportBoundaryRefs    []string                          `json:"export_boundary_refs,omitempty"`
	CompatibilityRefs     []string                          `json:"compatibility_refs,omitempty"`
	FailSafeRefs          []string                          `json:"fail_safe_refs,omitempty"`
	RolloutRefs           []string                          `json:"rollout_refs,omitempty"`
	Limitations           []string                          `json:"limitations,omitempty"`
}

type phase7MSPIsolationControl struct {
	ControlID    string   `json:"control_id"`
	CurrentState string   `json:"current_state"`
	Scope        string   `json:"scope"`
	Summary      string   `json:"summary"`
	ReasonCodes  []string `json:"reason_codes,omitempty"`
	EvidenceRefs []string `json:"evidence_refs,omitempty"`
	Limitations  []string `json:"limitations,omitempty"`
}

type phase7MSPIsolationResponse struct {
	SchemaVersion        string                      `json:"schema_version"`
	GeneratedAt          time.Time                   `json:"generated_at"`
	CurrentState         string                      `json:"current_state"`
	TenantIsolation      string                      `json:"tenant_isolation"`
	AuditIsolation       string                      `json:"audit_isolation"`
	DelegatedManagement  string                      `json:"delegated_management"`
	AllowedDelegations   []string                    `json:"allowed_delegations,omitempty"`
	ForbiddenDelegations []string                    `json:"forbidden_delegations,omitempty"`
	SupportabilityScopes []string                    `json:"supportability_scopes,omitempty"`
	Controls             []phase7MSPIsolationControl `json:"controls,omitempty"`
	ExportBoundaryRefs   []string                    `json:"export_boundary_refs,omitempty"`
	AbuseControlRefs     []string                    `json:"abuse_control_refs,omitempty"`
	FailSafeRefs         []string                    `json:"fail_safe_refs,omitempty"`
	RolloutRefs          []string                    `json:"rollout_refs,omitempty"`
	Limitations          []string                    `json:"limitations,omitempty"`
}

type phase7PartnerExportVisibility struct {
	VisibilityClass string   `json:"visibility_class"`
	Fields          []string `json:"fields,omitempty"`
}

type phase7PartnerExportResponse struct {
	SchemaVersion         string                          `json:"schema_version"`
	GeneratedAt           time.Time                       `json:"generated_at"`
	CurrentState          string                          `json:"current_state"`
	Scope                 string                          `json:"scope"`
	VisibilityState       string                          `json:"visibility_state"`
	AllowedOperations     []string                        `json:"allowed_operations,omitempty"`
	ForbiddenOperations   []string                        `json:"forbidden_operations,omitempty"`
	CredentialModel       []string                        `json:"credential_model,omitempty"`
	VisibilityClasses     []phase7PartnerExportVisibility `json:"visibility_classes,omitempty"`
	RedactedByDefault     bool                            `json:"redacted_by_default"`
	RetentionPolicy       []string                        `json:"retention_policy,omitempty"`
	OffboardingPolicy     []string                        `json:"offboarding_policy,omitempty"`
	CompatibilityRefs     []string                        `json:"compatibility_refs,omitempty"`
	FailSafeRefs          []string                        `json:"fail_safe_refs,omitempty"`
	ExpandedScopeDeferred []string                        `json:"expanded_scope_deferred,omitempty"`
	Limitations           []string                        `json:"limitations,omitempty"`
}

func (s server) phase7MarketplaceReadinessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7MarketplaceReadiness())
}

func (s server) phase7MSPIsolationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7MSPIsolation())
}

func (s server) phase7PartnerExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7PartnerExport())
}

func buildPhase7MarketplaceReadiness() phase7MarketplaceReadinessResponse {
	signals := ecosystemcore.SignalContractsForGroup("distribution")
	failSafe := phase7DistributionFailSafe("distribution.marketplace_deployment")
	boundary := phase7DistributionBoundary("distribution.marketplace_deployment")
	evidenceRefs := signalEvidenceRefs(signals, "distribution.marketplace_deployment_readiness")
	return phase7MarketplaceReadinessResponse{
		SchemaVersion:        phase7MarketplaceReadinessSchema,
		GeneratedAt:          publicSampleTime(),
		CurrentState:         phase7DistributionProjectionState(phase7MarketplaceReadinessStateActive, phase7MarketplaceReadinessStatePartial, phase7MarketplaceReadinessStateIncomplete),
		ProfileDetection:     "environment_profile_detection_active",
		ReadinessState:       "marketplace_readiness_gate_active",
		TrustPrerequisites:   "trust_and_config_prerequisites_visible",
		UpgradeBoundaryState: "upgrade_boundary_visibility_active",
		UnsupportedConditions: []string{
			"unsupported_marketplace_profile_requires_local_completion",
			"tenant_specific_secret_bootstrap_remains_local_completion",
		},
		Checks: []phase7MarketplaceReadinessCheck{
			{
				CheckID:      "profile_detection",
				CurrentState: "marketplace_check_active",
				Category:     ecosystemcore.SignalClassObservedFact,
				Summary:      "Environment profile detection stays explicit before any readiness verdict is shown.",
				ReasonCodes:  []string{"profile_detection_required", "unsupported_profile_not_silent_green"},
				EvidenceRefs: evidenceRefs,
				BoundaryRules: []string{
					"Unsupported marketplace topology remains not-ready until local completion is explicit.",
				},
			},
			{
				CheckID:      "trust_prerequisites",
				CurrentState: "marketplace_check_active",
				Category:     ecosystemcore.SignalClassReviewRequirement,
				Summary:      "Readiness checks keep trust and configuration prerequisites visible instead of burying them behind one-click setup.",
				ReasonCodes:  []string{"trust_prerequisites_visible", "config_prerequisites_visible"},
				EvidenceRefs: evidenceRefs,
				BoundaryRules: []string{
					"Missing prerequisites degrade or block readiness rather than falling through as ready.",
				},
			},
			{
				CheckID:      "post_deploy_readiness",
				CurrentState: "marketplace_check_active",
				Category:     ecosystemcore.SignalClassObservedFact,
				Summary:      "Post-deploy readiness remains evidence-bound and recomputed after install or upgrade.",
				ReasonCodes:  []string{"post_deploy_readiness", "bounded_support_window"},
				EvidenceRefs: evidenceRefs,
				BoundaryRules: []string{
					"Readiness cannot remain green when stale or degraded checks are present.",
				},
			},
			{
				CheckID:      "rollback_and_disable_path",
				CurrentState: "marketplace_check_active",
				Category:     ecosystemcore.SignalClassRecommendation,
				Summary:      "Rollback and disable paths stay visible before a deployment profile is presented as ready.",
				ReasonCodes:  []string{"rollback_path_required", "disable_path_visible"},
				EvidenceRefs: evidenceRefs,
				BoundaryRules: []string{
					"Absence of rollback readiness prevents a full-ready posture.",
				},
			},
		},
		ExportBoundaryRefs: phase7BoundaryIDs(ecosystemcore.DataBoundariesForGroup("distribution"), "distribution.marketplace_deployment"),
		CompatibilityRefs:  phase7DistributionCompatibilityRefs("distribution.marketplace_deployment"),
		FailSafeRefs:       phase7DistributionFailSafeRefs("distribution.marketplace_deployment"),
		RolloutRefs:        phase7DistributionRolloutRefs("distribution.marketplace_deployment"),
		Limitations: []string{
			firstNonEmpty(failSafe.DegradedBehavior, "Return not-ready or degraded instead of a silent ready state."),
			"Marketplace readiness stays evidence-bound and does not imply click-and-forget production completion.",
			"Only aggregate marketplace posture can become public; tenant-specific readiness detail remains bounded by export boundaries.",
			firstNonEmpty(firstString(boundary.PartnerVisible), "bounded readiness verdict remains partner-visible"),
		},
	}
}

func buildPhase7MSPIsolation() phase7MSPIsolationResponse {
	signals := ecosystemcore.SignalContractsForGroup("distribution")
	evidenceRefs := signalEvidenceRefs(signals, "distribution.msp_tenant_isolation_posture")
	return phase7MSPIsolationResponse{
		SchemaVersion:       phase7MSPIsolationSchema,
		GeneratedAt:         publicSampleTime(),
		CurrentState:        phase7DistributionProjectionState(phase7MSPIsolationStateActive, phase7MSPIsolationStatePartial, phase7MSPIsolationStateIncomplete),
		TenantIsolation:     "strict_tenant_isolation_verified",
		AuditIsolation:      "per_tenant_audit_isolation_verified",
		DelegatedManagement: "bounded_delegated_management_active",
		AllowedDelegations: []string{
			"tenant_safe_read_only_exports",
			"scoped_supportability",
			"tenant_scoped_operator_actions",
		},
		ForbiddenDelegations: []string{
			"cross_tenant_mutation",
			"shared_audit_stream",
			"broad_partner_automation",
		},
		SupportabilityScopes: []string{
			"tenant_scoped_support_bundle",
			"tenant_scoped_health_snapshot",
			"fleet_aggregate_only_outside_tenant_scope",
		},
		Controls: []phase7MSPIsolationControl{
			{
				ControlID:    "tenant_boundary_enforcement",
				CurrentState: "msp_control_active",
				Scope:        "tenant_boundary",
				Summary:      "Tenant boundary checks remain explicit and block partner-safe claims on any cross-tenant regression.",
				ReasonCodes:  []string{"strict_tenant_isolation", "cross_tenant_failure_blocks_claims"},
				EvidenceRefs: evidenceRefs,
			},
			{
				ControlID:    "audit_isolation",
				CurrentState: "msp_control_active",
				Scope:        "audit",
				Summary:      "Per-tenant audit isolation remains part of the bounded MSP operating contract.",
				ReasonCodes:  []string{"per_tenant_audit_isolation", "audit_boundary_enforced"},
				EvidenceRefs: evidenceRefs,
			},
			{
				ControlID:    "tenant_safe_automation",
				CurrentState: "msp_control_active",
				Scope:        "automation",
				Summary:      "Automation remains tenant-safe, read/export scoped, and explicitly non-mutating across tenant boundaries.",
				ReasonCodes:  []string{"tenant_safe_automation", "no_cross_tenant_mutation"},
				EvidenceRefs: evidenceRefs,
				Limitations: []string{
					"Automation scope remains bounded and does not widen into cross-tenant orchestration.",
				},
			},
		},
		ExportBoundaryRefs: phase7BoundaryIDs(ecosystemcore.DataBoundariesForGroup("distribution"), "distribution.msp_operator"),
		AbuseControlRefs:   phase7DistributionAbuseRefs("distribution.msp_operator"),
		FailSafeRefs:       phase7DistributionFailSafeRefs("distribution.msp_operator"),
		RolloutRefs:        phase7DistributionRolloutRefs("distribution.msp_operator"),
		Limitations: []string{
			"No MSP override may bypass cross-tenant isolation failure.",
			"Delegated management remains bounded to tenant-safe operations and scoped supportability.",
		},
	}
}

func buildPhase7PartnerExport() phase7PartnerExportResponse {
	boundary := phase7DistributionBoundary("distribution.partner_export")
	failSafe := phase7DistributionFailSafe("distribution.partner_export")
	return phase7PartnerExportResponse{
		SchemaVersion:   phase7PartnerExportSchema,
		GeneratedAt:     publicSampleTime(),
		CurrentState:    phase7DistributionProjectionState(phase7PartnerExportStateActive, phase7PartnerExportStatePartial, phase7PartnerExportStateIncomplete),
		Scope:           ecosystemcore.ScopePartner,
		VisibilityState: "partner_bounded_export_visible",
		AllowedOperations: []string{
			"scoped_read",
			"verifier_friendly_export",
			"tenant_safe_hook",
		},
		ForbiddenOperations: []string{
			"cross_tenant_read",
			"implicit_orchestration",
			"broad_mutation_authority",
			"broader_partner_write_api",
		},
		CredentialModel: []string{
			"scoped_credentials",
			"lifecycle_safe_onboarding",
			"revocable_exports",
		},
		VisibilityClasses: []phase7PartnerExportVisibility{
			{VisibilityClass: "partner_visible", Fields: boundary.PartnerVisible},
			{VisibilityClass: "verifier_exportable", Fields: boundary.VerifierExportable},
			{VisibilityClass: "public_exportable", Fields: boundary.PublicExportable},
			{VisibilityClass: "aggregate_only", Fields: boundary.AggregateOnly},
		},
		RedactedByDefault:     boundary.RedactedByDefault,
		RetentionPolicy:       boundary.RetentionPolicy,
		OffboardingPolicy:     boundary.OffboardingPolicy,
		CompatibilityRefs:     phase7DistributionCompatibilityRefs("distribution.partner_export"),
		FailSafeRefs:          phase7DistributionFailSafeRefs("distribution.partner_export"),
		ExpandedScopeDeferred: []string{"broader_partner_write_api", "integrity_as_a_service_package"},
		Limitations: []string{
			firstNonEmpty(failSafe.DegradedBehavior, "Return redacted incomplete export rather than widening disclosure."),
			"Partner export stays scoped and revocable and never becomes a broad write-capable orchestration layer.",
			"Public exportable fields remain empty for partner export in the bounded core pass.",
		},
	}
}

func phase7DistributionProjectionState(activeState, partialState, incompleteState string) string {
	switch ecosystemcore.EvaluateDistributionPresenceState() {
	case ecosystemcore.DistributionPresenceStateActive:
		return activeState
	case ecosystemcore.DistributionPresenceStatePartial:
		return partialState
	default:
		return incompleteState
	}
}

func phase7DistributionBoundary(surfaceID string) ecosystemcore.DataBoundary {
	for _, item := range ecosystemcore.DataBoundariesForGroup("distribution") {
		if item.SurfaceID == surfaceID {
			return item
		}
	}
	return ecosystemcore.DataBoundary{}
}

func phase7DistributionFailSafe(surfaceID string) ecosystemcore.FailSafeContract {
	for _, item := range ecosystemcore.FailSafeContractsForGroup("distribution") {
		if item.SurfaceID == surfaceID {
			return item
		}
	}
	return ecosystemcore.FailSafeContract{}
}

func phase7DistributionCompatibilityRefs(surfaceID string) []string {
	out := []string{}
	for _, item := range ecosystemcore.CompatibilityContractsForGroup("distribution") {
		if item.SurfaceID == surfaceID {
			out = append(out, item.SurfaceID)
		}
	}
	return out
}

func phase7DistributionFailSafeRefs(surfaceID string) []string {
	out := []string{}
	for _, item := range ecosystemcore.FailSafeContractsForGroup("distribution") {
		if item.SurfaceID == surfaceID {
			out = append(out, item.SurfaceID)
		}
	}
	return out
}

func phase7DistributionRolloutRefs(surfaceID string) []string {
	out := []string{}
	for _, item := range ecosystemcore.RolloutContractsForGroup("distribution") {
		if item.SurfaceID == surfaceID {
			out = append(out, item.SurfaceID)
		}
	}
	return out
}

func phase7DistributionAbuseRefs(surfaceID string) []string {
	out := []string{}
	for _, item := range ecosystemcore.AbuseControlsForGroup("distribution") {
		if item.SurfaceID == surfaceID {
			out = append(out, item.SurfaceID)
		}
	}
	return out
}
