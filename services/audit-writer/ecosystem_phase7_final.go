package main

import (
	"net/http"
	"time"

	ecosystemcore "github.com/denisgrosek/changelock/internal/ecosystem"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	phase7FinalSummarySchema = "7d.ecosystem_phase7_final_summary.v1"

	phase7FinalReviewStateActive     = "phase7_final_review_active"
	phase7FinalReviewStatePartial    = "phase7_final_review_partial"
	phase7FinalReviewStateIncomplete = "phase7_final_review_incomplete"

	phase7FinalizationStateReady       = "phase7_finalization_ready"
	phase7FinalizationStateSubstantial = "phase7_finalization_substantially_ready"
	phase7FinalizationStateIncomplete  = "phase7_finalization_incomplete"
)

type phase7FinalReviewCheck struct {
	CheckID      string   `json:"check_id"`
	CurrentState string   `json:"current_state"`
	Summary      string   `json:"summary"`
	ReasonCodes  []string `json:"reason_codes,omitempty"`
	RouteRefs    []string `json:"route_refs,omitempty"`
	DocRefs      []string `json:"doc_refs,omitempty"`
	Limitations  []string `json:"limitations,omitempty"`
}

type phase7FinalReviewSection struct {
	CurrentState string                   `json:"current_state"`
	RouteRefs    []string                 `json:"route_refs,omitempty"`
	DocRefs      []string                 `json:"doc_refs,omitempty"`
	Checks       []phase7FinalReviewCheck `json:"checks,omitempty"`
	Limitations  []string                 `json:"limitations,omitempty"`
}

type phase7FinalSummaryResponse struct {
	SchemaVersion              string                   `json:"schema_version"`
	GeneratedAt                time.Time                `json:"generated_at"`
	CurrentState               string                   `json:"current_state"`
	Phase7CoreState            string                   `json:"phase7_core_state"`
	DeveloperPresenceReview    phase7FinalReviewSection `json:"developer_presence_review"`
	OSSBoundaryReview          phase7FinalReviewSection `json:"oss_boundary_review"`
	DistributionBoundaryReview phase7FinalReviewSection `json:"distribution_boundary_review"`
	ContractAlignment          phase7FinalReviewSection `json:"contract_alignment"`
	DocsAndProofs              phase7FinalReviewSection `json:"docs_and_proofs"`
	DeferredScope              []string                 `json:"deferred_scope,omitempty"`
	Limitations                []string                 `json:"limitations,omitempty"`
}

func (s server) phase7FinalSummaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7FinalSummary())
}

func buildPhase7FinalSummary() phase7FinalSummaryResponse {
	coreProofs := buildPhase7Proofs()
	developerReview := buildPhase7DeveloperFinalReview()
	ossReview := buildPhase7OSSFinalReview()
	distributionReview := buildPhase7DistributionFinalReview()
	contractAlignment := buildPhase7ContractAlignmentReview()
	docsAndProofs := buildPhase7DocsAndProofsReview()
	return phase7FinalSummaryResponse{
		SchemaVersion:              phase7FinalSummarySchema,
		GeneratedAt:                publicSampleTime(),
		CurrentState:               phase7FinalizationState(coreProofs.CurrentState, developerReview, ossReview, distributionReview, contractAlignment, docsAndProofs),
		Phase7CoreState:            coreProofs.CurrentState,
		DeveloperPresenceReview:    developerReview,
		OSSBoundaryReview:          ossReview,
		DistributionBoundaryReview: distributionReview,
		ContractAlignment:          contractAlignment,
		DocsAndProofs:              docsAndProofs,
		DeferredScope: []string{
			"automated_pr_discipline",
			"broader_partner_api",
			"integrity_as_a_service_package",
			"additional_registry_provider_coverage",
		},
		Limitations: []string{
			"Phase 7 final summary is a bounded finalization pack over existing core and slice surfaces rather than a new authority layer.",
			"Formal external authority, certification posture, and regulator-facing trust remain reserved for a later authority phase.",
		},
	}
}

func buildPhase7DeveloperFinalReview() phase7FinalReviewSection {
	presence := buildPhase7DeveloperPresence()
	workbench := buildPhase7DeveloperWorkbench()
	context := buildPhase7DeveloperContextPack()
	preCommit := buildPhase7DeveloperPreCommitProfile()
	active := presence.CurrentState == ecosystemcore.DeveloperPresenceStateActive &&
		workbench.CurrentState == phase7DeveloperWorkbenchStateActive &&
		context.CurrentState == phase7DeveloperContextStateActive &&
		preCommit.CurrentState == phase7DeveloperPreCommitStateActive &&
		containsString(workbench.OutputSemantics, "uncertainty") &&
		containsString(preCommit.FailSafeBehaviors, "Pre-commit stays non-mutating and cannot auto-open PRs or apply policy changes.")
	return phase7FinalReviewSection{
		CurrentState: phase7SectionState(active, presence.CurrentState != ecosystemcore.DeveloperPresenceStateIncomplete),
		RouteRefs: []string{
			"/v1/ecosystem/phase7/developer-presence",
			"/v1/ecosystem/phase7/developer/workbench",
			"/v1/ecosystem/phase7/developer/context",
			"/v1/ecosystem/phase7/developer/pre-commit",
		},
		DocRefs: []string{
			"docs/ecosystem-phase7-core.md",
			"docs/ecosystem-phase7-vala.md",
		},
		Checks: []phase7FinalReviewCheck{
			{
				CheckID:      "developer_workbench_boundary",
				CurrentState: phase7ReviewCheckState(workbench.CurrentState == phase7DeveloperWorkbenchStateActive),
				Summary:      "Developer workbench remains bounded to local advisory and command-pack projections.",
				ReasonCodes:  []string{"bounded_workbench_projection", "no_local_policy_authority"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/developer/workbench"},
				DocRefs:      []string{"docs/ecosystem-phase7-vala.md"},
			},
			{
				CheckID:      "developer_context_boundary",
				CurrentState: phase7ReviewCheckState(context.CurrentState == phase7DeveloperContextStateActive && containsString(workbench.OutputSemantics, "uncertainty")),
				Summary:      "Developer context keeps observed fact, derived relevance, recommendation, and uncertainty explicitly separated.",
				ReasonCodes:  []string{"fact_relevance_recommendation_uncertainty_split", "bounded_vex_context"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/developer/context"},
				DocRefs:      []string{"docs/ecosystem-phase7-vala.md"},
			},
			{
				CheckID:      "pre_commit_boundary",
				CurrentState: phase7ReviewCheckState(preCommit.CurrentState == phase7DeveloperPreCommitStateActive && preCommit.DisablePath != "" && containsString(preCommit.FailSafeBehaviors, "Pre-commit stays non-mutating and cannot auto-open PRs or apply policy changes.")),
				Summary:      "Pre-commit remains bounded, explainable, and non-mutating with explicit override and disable behavior.",
				ReasonCodes:  []string{"bounded_pre_commit", "non_mutating_fail_safe", "override_and_disable_visible"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/developer/pre-commit"},
				DocRefs:      []string{"docs/ecosystem-phase7-vala.md"},
			},
		},
		Limitations: []string{
			"Developer presence remains advisory and never becomes canonical production truth or mutation authority.",
		},
	}
}

func buildPhase7OSSFinalReview() phase7FinalReviewSection {
	network := buildPhase7OSSNetwork()
	connectors := buildPhase7OSSConnectors()
	observations := buildPhase7OSSObservations()
	reviewFlow := buildPhase7OSSReviewFlow()
	reviewed := buildPhase7OSSReviewedSignals()
	hasReviewedLifecycle := false
	for _, item := range reviewed.Records {
		if item.ReviewState == "reviewed" || item.ReviewState == "superseded" || item.ReviewState == "revoked" {
			hasReviewedLifecycle = true
			break
		}
	}
	active := network.CurrentState == ecosystemcore.OSSPresenceStateActive &&
		connectors.CurrentState == phase7OSSConnectorsStateActive &&
		observations.CurrentState == phase7OSSObservationsStateActive &&
		reviewFlow.CurrentState == phase7OSSReviewFlowStateActive &&
		reviewed.CurrentState == phase7OSSReviewedSignalsStateActive &&
		network.AutomatedPRState == phase7DeferredExpandedScopeTag &&
		observations.CommunityInputState == "community_candidate_only_review_required" &&
		hasReviewedLifecycle
	return phase7FinalReviewSection{
		CurrentState: phase7SectionState(active, network.CurrentState != ecosystemcore.OSSPresenceStateIncomplete),
		RouteRefs: []string{
			"/v1/ecosystem/phase7/oss-network",
			"/v1/ecosystem/phase7/oss/connectors",
			"/v1/ecosystem/phase7/oss/observations",
			"/v1/ecosystem/phase7/oss/review-flow",
			"/v1/ecosystem/phase7/oss/reviewed-signals",
		},
		DocRefs: []string{
			"docs/ecosystem-phase7-core.md",
			"docs/ecosystem-phase7-valb.md",
		},
		Checks: []phase7FinalReviewCheck{
			{
				CheckID:      "observation_claim_boundary",
				CurrentState: phase7ReviewCheckState(observations.CommunityInputState == "community_candidate_only_review_required" && !containsString(observations.CandidateStates, "reviewed")),
				Summary:      "Observation intake remains distinct from reviewed claim publication.",
				ReasonCodes:  []string{"candidate_only_until_review", "observation_not_claim"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/oss/observations", "/v1/ecosystem/phase7/oss/review-flow"},
				DocRefs:      []string{"docs/ecosystem-phase7-valb.md"},
			},
			{
				CheckID:      "review_lifecycle_visibility",
				CurrentState: phase7ReviewCheckState(hasReviewedLifecycle && containsString(reviewFlow.ReviewStates, "revoked")),
				Summary:      "Reviewed OSS publication keeps reviewed, superseded, and revoked lifecycle states visible.",
				ReasonCodes:  []string{"reviewed_superseded_revoked_visible", "bounded_reviewed_publication"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/oss/reviewed-signals", "/v1/public/proof-portal"},
				DocRefs:      []string{"docs/ecosystem-phase7-valb.md"},
			},
			{
				CheckID:      "automation_deferred_boundary",
				CurrentState: phase7ReviewCheckState(network.AutomatedPRState == phase7DeferredExpandedScopeTag && containsString(reviewFlow.ExpandedScopeDeferred, "automated_pr_discipline")),
				Summary:      "Automated remediation PR discipline remains deferred and outside bounded review publication.",
				ReasonCodes:  []string{"automated_pr_deferred", "no_hidden_mutation_path"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/oss-network", "/v1/ecosystem/phase7/oss/review-flow"},
				DocRefs:      []string{"docs/ecosystem-phase7-valb.md"},
			},
		},
		Limitations: []string{
			"OSS trust remains bounded to reviewed publication over candidate intake and does not become crowd-sourced canonical truth.",
		},
	}
}

func buildPhase7DistributionFinalReview() phase7FinalReviewSection {
	distribution := buildPhase7Distribution()
	marketplace := buildPhase7MarketplaceReadiness()
	msp := buildPhase7MSPIsolation()
	partner := buildPhase7PartnerExport()
	active := distribution.CurrentState == ecosystemcore.DistributionPresenceStateActive &&
		marketplace.CurrentState == phase7MarketplaceReadinessStateActive &&
		msp.CurrentState == phase7MSPIsolationStateActive &&
		partner.CurrentState == phase7PartnerExportStateActive &&
		containsString(msp.ForbiddenDelegations, "cross_tenant_mutation") &&
		containsString(partner.ForbiddenOperations, "broader_partner_write_api") &&
		len(visibilityFields(partner.VisibilityClasses, "public_exportable")) == 0
	return phase7FinalReviewSection{
		CurrentState: phase7SectionState(active, distribution.CurrentState != ecosystemcore.DistributionPresenceStateIncomplete),
		RouteRefs: []string{
			"/v1/ecosystem/phase7/distribution",
			"/v1/ecosystem/phase7/distribution/marketplace-readiness",
			"/v1/ecosystem/phase7/distribution/msp-isolation",
			"/v1/ecosystem/phase7/distribution/partner-export",
		},
		DocRefs: []string{
			"docs/ecosystem-phase7-core.md",
			"docs/ecosystem-phase7-valc.md",
		},
		Checks: []phase7FinalReviewCheck{
			{
				CheckID:      "marketplace_readiness_boundary",
				CurrentState: phase7ReviewCheckState(len(marketplace.UnsupportedConditions) > 0 && marketplace.CurrentState == phase7MarketplaceReadinessStateActive),
				Summary:      "Marketplace readiness remains bounded, unsupported conditions remain visible, and readiness never implies completion-optimism.",
				ReasonCodes:  []string{"unsupported_conditions_visible", "no_click_and_forget_completion"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/distribution/marketplace-readiness"},
				DocRefs:      []string{"docs/ecosystem-phase7-valc.md"},
			},
			{
				CheckID:      "msp_isolation_boundary",
				CurrentState: phase7ReviewCheckState(containsString(msp.ForbiddenDelegations, "cross_tenant_mutation") && containsString(msp.ForbiddenDelegations, "shared_audit_stream")),
				Summary:      "MSP isolation remains tenant- and audit-bounded and forbids cross-tenant mutation paths.",
				ReasonCodes:  []string{"strict_tenant_isolation_verified", "per_tenant_audit_isolation_verified"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/distribution/msp-isolation"},
				DocRefs:      []string{"docs/ecosystem-phase7-valc.md"},
			},
			{
				CheckID:      "partner_export_boundary",
				CurrentState: phase7ReviewCheckState(containsString(partner.ForbiddenOperations, "broader_partner_write_api") && len(visibilityFields(partner.VisibilityClasses, "public_exportable")) == 0),
				Summary:      "Partner export remains partner-scoped, redacted-by-default, and non-orchestrating.",
				ReasonCodes:  []string{"broader_partner_write_api_deferred", "public_exportable_empty", "redacted_by_default"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/distribution/partner-export"},
				DocRefs:      []string{"docs/ecosystem-phase7-valc.md"},
			},
		},
		Limitations: []string{
			"Distribution surfaces remain readiness- and boundary-bound and do not claim broad partner orchestration or full service packaging closure.",
		},
	}
}

func buildPhase7ContractAlignmentReview() phase7FinalReviewSection {
	contracts := buildPhase7Contracts()
	active := contracts.CurrentState == ecosystemcore.FoundationStateActive &&
		contracts.Coverage.SignalContracts >= 8 &&
		contracts.Coverage.AuthoritySurfaces >= 8 &&
		contracts.Coverage.DataBoundaries >= 4
	return phase7FinalReviewSection{
		CurrentState: phase7SectionState(active, contracts.CurrentState != ecosystemcore.FoundationStateIncomplete),
		RouteRefs: []string{
			"/v1/ecosystem/phase7/entry-gate",
			"/v1/ecosystem/phase7/contracts",
		},
		DocRefs: []string{
			"docs/ecosystem-phase7-core.md",
		},
		Checks: []phase7FinalReviewCheck{
			{
				CheckID:      "entry_gate_and_contract_foundation",
				CurrentState: phase7ReviewCheckState(contracts.CurrentState == ecosystemcore.FoundationStateActive),
				Summary:      "Entry gate and contract foundation remain active and fail-closed across required core surfaces.",
				ReasonCodes:  []string{"entry_gate_ready", "contract_foundation_active", "fail_closed_core_coverage"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/entry-gate", "/v1/ecosystem/phase7/contracts"},
				DocRefs:      []string{"docs/ecosystem-phase7-core.md"},
			},
			{
				CheckID:      "core_coverage_alignment",
				CurrentState: phase7ReviewCheckState(contracts.Coverage.SignalContracts >= 8 && contracts.Coverage.AuthoritySurfaces >= 8 && contracts.Coverage.DataBoundaries >= 4),
				Summary:      "Core-pass coverage remains aligned with bounded developer, OSS, and distribution surfaces.",
				ReasonCodes:  []string{"core_pass_coverage", "deferred_surfaces_excluded"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/contracts"},
				DocRefs:      []string{"docs/ecosystem-phase7-core.md"},
			},
		},
		Limitations: []string{
			"Contract alignment review stays scoped to core-pass coverage and bounded finalization semantics.",
		},
	}
}

func buildPhase7DocsAndProofsReview() phase7FinalReviewSection {
	proofs := buildPhase7Proofs()
	docRefs := []string{
		"docs/ecosystem-phase7-core.md",
		"docs/ecosystem-phase7-vala.md",
		"docs/ecosystem-phase7-valb.md",
		"docs/ecosystem-phase7-valc.md",
		"docs/ecosystem-phase7-final.md",
	}
	active := proofs.CurrentState == ecosystemcore.Phase7StateActive &&
		proofs.CoverageScope == phase7CoverageScopeCorePass &&
		containsString(proofs.ExpandedScopeDeferred, "automated_pr_discipline") &&
		containsString(proofs.ExpandedScopeDeferred, "integrity_as_a_service_package")
	return phase7FinalReviewSection{
		CurrentState: phase7SectionState(active, proofs.CurrentState != ecosystemcore.Phase7StateIncomplete),
		RouteRefs: []string{
			"/v1/ecosystem/phase7/proofs",
			"/v1/ecosystem/phase7/final-summary",
		},
		DocRefs: docRefs,
		Checks: []phase7FinalReviewCheck{
			{
				CheckID:      "proofs_gate_alignment",
				CurrentState: phase7ReviewCheckState(proofs.CurrentState == ecosystemcore.Phase7StateActive && proofs.CoverageScope == phase7CoverageScopeCorePass),
				Summary:      "Phase 7 proofs remain active and core-pass scoped across the assembled ecosystem package.",
				ReasonCodes:  []string{"phase7_core_ecosystem_presence_active", "core_pass_scoped_proofs"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/proofs"},
				DocRefs:      docRefs,
			},
			{
				CheckID:      "deferred_scope_visibility",
				CurrentState: phase7ReviewCheckState(containsString(proofs.ExpandedScopeDeferred, "automated_pr_discipline") && containsString(proofs.ExpandedScopeDeferred, "integrity_as_a_service_package")),
				Summary:      "Deferred expanded-scope boundaries remain visible in final proofs and summary output.",
				ReasonCodes:  []string{"deferred_scope_visible", "no_phase8_authority_drift"},
				RouteRefs:    []string{"/v1/ecosystem/phase7/proofs", "/v1/ecosystem/phase7/final-summary"},
				DocRefs:      docRefs,
			},
		},
		Limitations: []string{
			"Docs and proofs review confirms bounded alignment; it does not widen Phase 7 into a later authority phase.",
		},
	}
}

func phase7FinalizationState(phase7State string, sections ...phase7FinalReviewSection) string {
	if phase7State != ecosystemcore.Phase7StateActive {
		return phase7FinalizationStateIncomplete
	}
	hasPartial := false
	for _, section := range sections {
		switch section.CurrentState {
		case phase7FinalReviewStateActive:
		case phase7FinalReviewStatePartial:
			hasPartial = true
		default:
			return phase7FinalizationStateIncomplete
		}
	}
	if hasPartial {
		return phase7FinalizationStateSubstantial
	}
	return phase7FinalizationStateReady
}

func phase7SectionState(active bool, hasCoverage bool) string {
	switch {
	case active:
		return phase7FinalReviewStateActive
	case hasCoverage:
		return phase7FinalReviewStatePartial
	default:
		return phase7FinalReviewStateIncomplete
	}
}

func phase7ReviewCheckState(active bool) string {
	if active {
		return phase7FinalReviewStateActive
	}
	return phase7FinalReviewStateIncomplete
}
