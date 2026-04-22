package main

import (
	"net/http"
	"strings"
	"time"

	ecosystemcore "github.com/denisgrosek/changelock/internal/ecosystem"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	phase7DeveloperWorkbenchSchema = "7a.ecosystem_developer_workbench.v1"
	phase7DeveloperContextSchema   = "7a.ecosystem_developer_context_pack.v1"
	phase7DeveloperPreCommitSchema = "7a.ecosystem_developer_pre_commit.v1"

	phase7DeveloperWorkbenchStateActive  = "developer_workbench_active"
	phase7DeveloperWorkbenchStatePartial = "developer_workbench_partial"
	phase7DeveloperContextStateActive    = "developer_context_pack_active"
	phase7DeveloperPreCommitStateActive  = "developer_pre_commit_profile_active"
)

type phase7DeveloperCommand struct {
	CommandID      string   `json:"command_id"`
	CurrentState   string   `json:"current_state"`
	Command        string   `json:"command"`
	Purpose        string   `json:"purpose"`
	Scope          string   `json:"scope"`
	BlockingPath   bool     `json:"blocking_path"`
	ReviewRequired bool     `json:"review_required"`
	DisablePath    string   `json:"disable_path,omitempty"`
	LatencyBudget  int      `json:"latency_budget_ms,omitempty"`
	ReasonCodes    []string `json:"reason_codes,omitempty"`
	DocsRefs       []string `json:"docs_refs,omitempty"`
	EvidenceRefs   []string `json:"evidence_refs,omitempty"`
	Limitations    []string `json:"limitations,omitempty"`
}

type phase7DeveloperAttentionRule struct {
	RuleID       string   `json:"rule_id"`
	CurrentState string   `json:"current_state"`
	AppliesTo    string   `json:"applies_to"`
	Threshold    string   `json:"threshold"`
	Enforcement  string   `json:"enforcement"`
	ReasonCodes  []string `json:"reason_codes,omitempty"`
	Limitations  []string `json:"limitations,omitempty"`
}

type phase7DeveloperWorkbenchResponse struct {
	SchemaVersion        string                         `json:"schema_version"`
	GeneratedAt          time.Time                      `json:"generated_at"`
	CurrentState         string                         `json:"current_state"`
	IDEAdvisoryState     string                         `json:"ide_advisory_state"`
	LocalValidationState string                         `json:"local_validation_state"`
	InEditorContextState string                         `json:"in_editor_context_state"`
	PreCommitState       string                         `json:"pre_commit_state"`
	OutputSemantics      []string                       `json:"output_semantics,omitempty"`
	Commands             []phase7DeveloperCommand       `json:"commands,omitempty"`
	AttentionBudget      []phase7DeveloperAttentionRule `json:"attention_budget,omitempty"`
	RouteRefs            []string                       `json:"route_refs,omitempty"`
	Limitations          []string                       `json:"limitations,omitempty"`
}

type phase7DeveloperContextItem struct {
	ContextID        string   `json:"context_id"`
	CurrentState     string   `json:"current_state"`
	SignalClass      string   `json:"signal_class"`
	ReviewedState    string   `json:"reviewed_state"`
	Summary          string   `json:"summary"`
	RouteRefs        []string `json:"route_refs,omitempty"`
	EvidenceRefs     []string `json:"evidence_refs,omitempty"`
	ReasonCodes      []string `json:"reason_codes,omitempty"`
	UncertaintyNotes []string `json:"uncertainty_notes,omitempty"`
	Limitations      []string `json:"limitations,omitempty"`
}

type phase7DeveloperContextResponse struct {
	SchemaVersion string                       `json:"schema_version"`
	GeneratedAt   time.Time                    `json:"generated_at"`
	CurrentState  string                       `json:"current_state"`
	Items         []phase7DeveloperContextItem `json:"items,omitempty"`
	Limitations   []string                     `json:"limitations,omitempty"`
}

type phase7DeveloperPreCommitResponse struct {
	SchemaVersion     string                         `json:"schema_version"`
	GeneratedAt       time.Time                      `json:"generated_at"`
	CurrentState      string                         `json:"current_state"`
	HookState         string                         `json:"hook_state"`
	BlockingModel     string                         `json:"blocking_model"`
	OverridePolicy    string                         `json:"override_policy"`
	DisablePath       string                         `json:"disable_path"`
	LatencyBudget     int                            `json:"latency_budget_ms"`
	FailSafeBehaviors []string                       `json:"fail_safe_behaviors,omitempty"`
	Commands          []phase7DeveloperCommand       `json:"commands,omitempty"`
	AttentionBudget   []phase7DeveloperAttentionRule `json:"attention_budget,omitempty"`
	Limitations       []string                       `json:"limitations,omitempty"`
}

func (s server) phase7DeveloperWorkbenchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7DeveloperWorkbench())
}

func (s server) phase7DeveloperContextHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7DeveloperContextPack())
}

func (s server) phase7DeveloperPreCommitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7DeveloperPreCommitProfile())
}

func buildPhase7DeveloperWorkbench() phase7DeveloperWorkbenchResponse {
	commands := phase7DeveloperCommands()
	attention := phase7DeveloperAttentionBudget()
	currentState := phase7DeveloperWorkbenchStatePartial
	if ecosystemcore.EvaluateDeveloperPresenceState() == ecosystemcore.DeveloperPresenceStateActive {
		currentState = phase7DeveloperWorkbenchStateActive
	}
	return phase7DeveloperWorkbenchResponse{
		SchemaVersion:        phase7DeveloperWorkbenchSchema,
		GeneratedAt:          publicSampleTime(),
		CurrentState:         currentState,
		IDEAdvisoryState:     "ide_advisory_delivery_active",
		LocalValidationState: "local_validation_workbench_active",
		InEditorContextState: phase7DeveloperContextStateActive,
		PreCommitState:       phase7DeveloperPreCommitStateActive,
		OutputSemantics: []string{
			ecosystemcore.SignalClassObservedFact,
			ecosystemcore.SignalClassDerivedRelevance,
			ecosystemcore.SignalClassRecommendation,
			"uncertainty",
		},
		Commands:        commands,
		AttentionBudget: attention,
		RouteRefs: []string{
			"/v1/ecosystem/phase7/developer/context",
			"/v1/ecosystem/phase7/developer/pre-commit",
			"/v1/intelligence/vulnerability-relevance",
			"/v1/vex/status",
		},
		Limitations: []string{
			"Developer workbench remains a bounded developer-presence projection and does not become a new local policy authority.",
			"IDE and local validation helpers stay advisory or review-required and do not claim production equivalence or live runtime truth.",
		},
	}
}

func buildPhase7DeveloperContextPack() phase7DeveloperContextResponse {
	signals := ecosystemcore.SignalContractsForGroup("developer")
	items := []phase7DeveloperContextItem{
		{
			ContextID:     "dependency_trust_advisory",
			CurrentState:  "dependency_trust_advisory_active",
			SignalClass:   ecosystemcore.SignalClassDerivedRelevance,
			ReviewedState: "advisory_only",
			Summary:       "IDE trust feedback stays same-subject and bounded to dependency trust posture, drift, and safer-alternative hints.",
			RouteRefs:     []string{"/v1/ecosystem/phase7/developer/workbench"},
			EvidenceRefs:  signalEvidenceRefs(signals, "developer.ide_trust_advisory"),
			ReasonCodes:   []string{"developer_same_subject_advisory", "never_canonical_truth"},
			UncertaintyNotes: []string{
				"IDE context can become stale between local edits and upstream evidence refresh.",
				"Advisory state does not imply server-side approval or live runtime verification.",
			},
			Limitations: []string{
				"Derived relevance does not replace deterministic config or deploy-time enforcement.",
			},
		},
		{
			ContextID:     "local_validation_observed_fact",
			CurrentState:  "local_validation_context_active",
			SignalClass:   ecosystemcore.SignalClassObservedFact,
			ReviewedState: "deterministic_local_only",
			Summary:       "Local validation exposes deterministic config and subject checks while keeping runtime-only unknowns explicit.",
			RouteRefs:     []string{"/v1/ecosystem/phase7/developer/workbench"},
			EvidenceRefs:  signalEvidenceRefs(signals, "developer.local_validation_projection"),
			ReasonCodes:   []string{"simulation_not_production_truth", "deterministic_local_scope"},
			UncertaintyNotes: []string{
				"Local validation does not know live attestation freshness, tenant workflow state, or cluster-only substrate changes.",
			},
			Limitations: []string{
				"Observed-fact output remains local-run scoped rather than authoritative production state.",
			},
		},
		{
			ContextID:     "vex_relevance_context",
			CurrentState:  "vex_relevance_context_active",
			SignalClass:   ecosystemcore.SignalClassDerivedRelevance,
			ReviewedState: "candidate_or_reviewed_boundary_visible",
			Summary:       "In-editor relevance and VEX context separate bounded candidate context from reviewed publication and keep evidence refs visible.",
			RouteRefs:     []string{"/v1/intelligence/vulnerability-relevance", "/v1/vex/status"},
			EvidenceRefs:  signalEvidenceRefs(signals, "developer.vex_relevance_context"),
			ReasonCodes:   []string{"vex_context_bounded", "review_boundary_visible"},
			UncertaintyNotes: []string{
				"Candidate VEX context is not a published VEX statement.",
				"Relevance may change when digest, exposure, or reviewed VEX status changes.",
			},
			Limitations: []string{
				"Context pack does not mutate VEX, exceptions, or policy; it only exposes bounded explanation paths.",
			},
		},
		{
			ContextID:     "pre_commit_recommendation_boundary",
			CurrentState:  phase7DeveloperPreCommitStateActive,
			SignalClass:   ecosystemcore.SignalClassRecommendation,
			ReviewedState: ecosystemcore.StatusReviewRequired,
			Summary:       "Pre-commit output keeps bounded blocking and explicit override reasons, and fails open to warning when evidence is incomplete.",
			RouteRefs:     []string{"/v1/ecosystem/phase7/developer/pre-commit"},
			EvidenceRefs:  signalEvidenceRefs(signals, "developer.pre_commit_dependency_trust"),
			ReasonCodes:   []string{"bounded_blocking_model", "override_reason_required"},
			UncertaintyNotes: []string{
				"Blocking posture can downgrade to warning when freshness or evidence completeness falls below the configured floor.",
			},
			Limitations: []string{
				"Pre-commit recommendation is not a hidden mutation or approval engine.",
			},
		},
	}
	return phase7DeveloperContextResponse{
		SchemaVersion: phase7DeveloperContextSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  phase7DeveloperContextStateActive,
		Items:         items,
		Limitations: []string{
			"Developer context pack stays advisory and same-subject; it does not widen into global trust or policy authority.",
		},
	}
}

func buildPhase7DeveloperPreCommitProfile() phase7DeveloperPreCommitResponse {
	failSafe := phase7DeveloperFailSafe("developer.pre_commit_hook")
	budget := phase7DeveloperBudget("developer.pre_commit_hook", "pre_commit_max_blocking")
	commands := []phase7DeveloperCommand{
		{
			CommandID:      "pre_commit_check",
			CurrentState:   "developer_command_ready",
			Command:        "changelock-cli check --config <path> --output json",
			Purpose:        "Run deterministic local config and workflow checks before commit.",
			Scope:          "local",
			BlockingPath:   true,
			ReviewRequired: true,
			DisablePath:    "CHGLOCK_PRECOMMIT_DISABLE=1",
			LatencyBudget:  budget.TargetMS,
			ReasonCodes:    []string{"bounded_blocking_model", "explicit_override_reason"},
			DocsRefs:       []string{"docs/developer-preflight-cli.md", "docs/ecosystem-phase7-core.md"},
			EvidenceRefs:   signalEvidenceRefs(ecosystemcore.SignalContractsForGroup("developer"), "developer.pre_commit_dependency_trust"),
			Limitations: []string{
				"Blocking path is bounded to local evidence and cannot silently claim production readiness.",
			},
		},
		{
			CommandID:      "pre_commit_preview",
			CurrentState:   "developer_command_ready",
			Command:        "changelock-cli preview --config <path> --output json",
			Purpose:        "Show bounded startup, sync, and trusted-execution preview before commit or push.",
			Scope:          "local",
			BlockingPath:   false,
			ReviewRequired: true,
			LatencyBudget:  budget.TargetMS,
			ReasonCodes:    []string{"preview_bounded_local_only", "warn_before_block"},
			DocsRefs:       []string{"docs/developer-preflight-cli.md"},
			Limitations: []string{
				"Preview keeps live runtime and tenant-only unknowns explicit instead of flattening them into local certainty.",
			},
		},
	}
	return phase7DeveloperPreCommitResponse{
		SchemaVersion: phase7DeveloperPreCommitSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  phase7DeveloperPreCommitStateActive,
		HookState:     "pre_commit_hook_active",
		BlockingModel: "warn_before_block_with_fail_open_on_incomplete_evidence",
		OverridePolicy: firstNonEmpty(
			strings.TrimSpace(failSafe.ManualOverridePolicy),
			"Override requires explicit developer reason and remains locally auditable.",
		),
		DisablePath: firstNonEmpty(
			strings.TrimSpace(failSafe.DisablePath),
			"CHGLOCK_PRECOMMIT_DISABLE=1",
		),
		LatencyBudget: budget.TargetMS,
		FailSafeBehaviors: []string{
			firstNonEmpty(strings.TrimSpace(failSafe.DegradedBehavior), "Downgrade to warning rather than silently passing."),
			firstNonEmpty(strings.TrimSpace(failSafe.StaleBehavior), "Treat stale hook context as review-required."),
			"Pre-commit stays non-mutating and cannot auto-open PRs or apply policy changes.",
		},
		Commands:        commands,
		AttentionBudget: phase7DeveloperAttentionBudget(),
		Limitations: []string{
			"Pre-commit discipline is intentionally bounded and does not become a hidden repo mutation path.",
			"Disable path and override reason remain explicit to keep hook behavior explainable and auditable.",
		},
	}
}

func phase7DeveloperCommands() []phase7DeveloperCommand {
	signals := ecosystemcore.SignalContractsForGroup("developer")
	return []phase7DeveloperCommand{
		{
			CommandID:      "local_check",
			CurrentState:   "developer_command_ready",
			Command:        "changelock-cli check --config <path> --output json",
			Purpose:        "Deterministic local validation before commit, push, or PR.",
			Scope:          "local",
			BlockingPath:   false,
			ReviewRequired: false,
			LatencyBudget:  phase7DeveloperBudget("developer.local_validation", "local_validation_time").TargetMS,
			ReasonCodes:    []string{"deterministic_local_scope"},
			DocsRefs:       []string{"docs/developer-preflight-cli.md"},
			EvidenceRefs:   signalEvidenceRefs(signals, "developer.local_validation_projection"),
		},
		{
			CommandID:      "local_preview",
			CurrentState:   "developer_command_ready",
			Command:        "changelock-cli preview --config <path> --output json",
			Purpose:        "Bounded preview of sync, startup, and trusted-execution drift before push.",
			Scope:          "local",
			BlockingPath:   false,
			ReviewRequired: true,
			LatencyBudget:  phase7DeveloperBudget("developer.local_validation", "local_validation_time").TargetMS,
			ReasonCodes:    []string{"preview_bounded_local_only"},
			DocsRefs:       []string{"docs/developer-preflight-cli.md"},
			EvidenceRefs:   signalEvidenceRefs(signals, "developer.local_validation_projection"),
		},
		{
			CommandID:      "local_inspect",
			CurrentState:   "developer_command_ready",
			Command:        "changelock-cli inspect --config <path> --output json",
			Purpose:        "Show normalized and effective local config and runtime-self-healing posture.",
			Scope:          "local",
			BlockingPath:   false,
			ReviewRequired: false,
			ReasonCodes:    []string{"effective_state_visible"},
			DocsRefs:       []string{"docs/developer-preflight-cli.md", "docs/production-phase5-core.md"},
			EvidenceRefs:   signalEvidenceRefs(signals, "developer.local_validation_projection"),
		},
		{
			CommandID:      "local_explain",
			CurrentState:   "developer_command_ready",
			Command:        "changelock-cli explain --config <path> --topic sync --output json",
			Purpose:        "Explain local effective state, sync conflicts, and deterministic limits in text form.",
			Scope:          "local",
			BlockingPath:   false,
			ReviewRequired: false,
			ReasonCodes:    []string{"explainable_local_state"},
			DocsRefs:       []string{"docs/developer-preflight-cli.md"},
			EvidenceRefs:   signalEvidenceRefs(signals, "developer.local_validation_projection"),
		},
		{
			CommandID:      "post_run_guidance",
			CurrentState:   "developer_command_ready",
			Command:        "changelock-cli guidance --input ./artifacts/preflight.json --format markdown",
			Purpose:        "Render bounded relevance and VEX-aware guidance without mutating policy, VEX, or runtime state.",
			Scope:          "local_or_ci",
			BlockingPath:   false,
			ReviewRequired: true,
			ReasonCodes:    []string{"guidance_advisory_only", "vex_context_bounded"},
			DocsRefs:       []string{"docs/developer-preflight-cli.md", "docs/vex-exploitability-ops.md"},
			EvidenceRefs:   signalEvidenceRefs(signals, "developer.vex_relevance_context"),
			Limitations: []string{
				"Guidance output is advisory-only and does not auto-publish VEX or remediation decisions.",
			},
		},
	}
}

func phase7DeveloperAttentionBudget() []phase7DeveloperAttentionRule {
	return []phase7DeveloperAttentionRule{
		{
			RuleID:       "repeat_advisory_suppression",
			CurrentState: "attention_budget_active",
			AppliesTo:    "developer.ide_plugin",
			Threshold:    "no more than one repeated advisory per dependency per session without state change",
			Enforcement:  "suppress repeat advisory until evidence or subject state changes",
			ReasonCodes:  []string{"noise_budget_enforced", "repeat_suppression_active"},
			Limitations: []string{
				"Suppression reduces noise but does not hide changed or degraded state.",
			},
		},
		{
			RuleID:       "warn_before_block",
			CurrentState: "attention_budget_active",
			AppliesTo:    "developer.pre_commit_hook",
			Threshold:    "warn when confidence or freshness is below hard threshold before enabling blocking",
			Enforcement:  "degrade to warning when evidence completeness is insufficient",
			ReasonCodes:  []string{"warn_before_block", "fail_open_on_incomplete_evidence"},
		},
		{
			RuleID:       "stale_context_quiet_default",
			CurrentState: "attention_budget_active",
			AppliesTo:    "developer.vex_relevance_context",
			Threshold:    "default quiet behavior when same-subject context is stale and unchanged",
			Enforcement:  "show stale marker and context boundary instead of repeated interruption",
			ReasonCodes:  []string{"stale_context_marked", "quiet_default_when_weak"},
		},
	}
}

func phase7DeveloperFailSafe(surfaceID string) ecosystemcore.FailSafeContract {
	for _, item := range ecosystemcore.FailSafeContractsForGroup("developer") {
		if item.SurfaceID == surfaceID {
			return item
		}
	}
	return ecosystemcore.FailSafeContract{}
}

func phase7DeveloperBudget(surfaceID, budgetName string) ecosystemcore.PerformanceBudget {
	for _, item := range ecosystemcore.PerformanceBudgetsForGroup("developer") {
		if item.SurfaceID == surfaceID && item.BudgetName == budgetName {
			return item
		}
	}
	return ecosystemcore.PerformanceBudget{}
}
