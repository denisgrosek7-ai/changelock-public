package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	ossTrustNetworkValCStatusSchema = "point9.oss_trust_network.valc.status.v1"
	ossTrustNetworkValCProofsSchema = "point9.oss_trust_network.valc.proofs.v1"
)

type ossTrustNetworkValCStatusResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.OSSTrustNetworkValCCore `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type ossTrustNetworkValCProofsResponse struct {
	SchemaVersion                   string    `json:"schema_version"`
	GeneratedAt                     time.Time `json:"generated_at"`
	CurrentState                    string    `json:"current_state"`
	Point9State                     string    `json:"point_9_state"`
	DependencyState                 string    `json:"dependency_state"`
	ValBCurrentState                string    `json:"valb_current_state"`
	ValBPoint9State                 string    `json:"valb_point_9_state"`
	ValBDependencyState             string    `json:"valb_dependency_state"`
	ValBCandidateSignalIntakeState  string    `json:"valb_candidate_signal_intake_state"`
	ValBReviewWorkflowState         string    `json:"valb_review_workflow_state"`
	ValBSharedVEXTriageState        string    `json:"valb_shared_vex_triage_state"`
	ValBSourceWeightingState        string    `json:"valb_source_weighting_state"`
	ValBLocalApplicabilityState     string    `json:"valb_local_applicability_state"`
	ValBPropagationExchangeState    string    `json:"valb_propagation_exchange_state"`
	ValBSupersessionRevocationState string    `json:"valb_supersession_revocation_state"`
	ValBReviewerAuditabilityState   string    `json:"valb_reviewer_auditability_state"`
	ValBNoOverclaimState            string    `json:"valb_no_overclaim_state"`
	TrustVisibilityState            string    `json:"trust_visibility_state"`
	PackageTrustStatusState         string    `json:"package_trust_status_state"`
	ExportBoundaryState             string    `json:"export_boundary_state"`
	RemediationSuggestionState      string    `json:"remediation_suggestion_state"`
	PRProposalState                 string    `json:"pr_proposal_state"`
	LocalOverrideState              string    `json:"local_override_state"`
	RemediationSafetyState          string    `json:"remediation_safety_state"`
	EcosystemConsistencyState       string    `json:"ecosystem_consistency_state"`
	NoOverclaimState                string    `json:"no_overclaim_state"`
	VisibilityState                 string    `json:"visibility_state"`
	PackageStatusClass              string    `json:"package_status_class"`
	ExportClass                     string    `json:"export_class"`
	SuggestionClass                 string    `json:"suggestion_class"`
	SuggestionConfidenceClass       string    `json:"suggestion_confidence_class"`
	ProposalState                   string    `json:"proposal_state"`
	LocalOverrideVisibilityState    string    `json:"local_override_visibility_state"`
	RemediationRiskClass            string    `json:"remediation_risk_class"`
	SupportedVisibilityStates       []string  `json:"supported_visibility_states,omitempty"`
	SupportedPackageStatusClasses   []string  `json:"supported_package_status_classes,omitempty"`
	SupportedExportClasses          []string  `json:"supported_export_classes,omitempty"`
	SupportedSuggestionClasses      []string  `json:"supported_suggestion_classes,omitempty"`
	SupportedProposalStates         []string  `json:"supported_proposal_states,omitempty"`
	SupportedLocalOverrideStates    []string  `json:"supported_local_override_states,omitempty"`
	SupportedRiskClasses            []string  `json:"supported_risk_classes,omitempty"`
	SurfaceRefs                     []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                    []string  `json:"evidence_refs,omitempty"`
	BlockingReasons                 []string  `json:"blocking_reasons,omitempty"`
	WhyPoint9NotComplete            []string  `json:"why_point_9_not_complete,omitempty"`
	Limitations                     []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer            string    `json:"projection_disclaimer"`
	IntegrationSummary              []string  `json:"integration_summary,omitempty"`
}

func ossTrustNetworkValCAllSurfaceRefs() []string {
	return operability.OSSTrustNetworkValCProofSurfaceRefs()
}

func ossTrustNetworkValCVisibilityStates() []string {
	return []string{
		operability.OSSTrustNetworkValCVisibilityVisible,
		operability.OSSTrustNetworkValCVisibilityLimited,
		operability.OSSTrustNetworkValCVisibilityHidden,
		operability.OSSTrustNetworkValCVisibilityUnsupported,
		operability.OSSTrustNetworkValCVisibilityStale,
		operability.OSSTrustNetworkValCVisibilityUnknown,
	}
}

func ossTrustNetworkValCPackageStatusClasses() []string {
	return []string{
		operability.OSSTrustNetworkValCPackageStatusReviewedSignalAvailable,
		operability.OSSTrustNetworkValCPackageStatusCandidateSignalAvailable,
		operability.OSSTrustNetworkValCPackageStatusLocalReviewNeeded,
		operability.OSSTrustNetworkValCPackageStatusSupersededSignal,
		operability.OSSTrustNetworkValCPackageStatusRevokedSignal,
		operability.OSSTrustNetworkValCPackageStatusUnsupportedSignal,
		operability.OSSTrustNetworkValCPackageStatusUnknownSignal,
	}
}

func ossTrustNetworkValCExportClasses() []string {
	return []string{
		operability.OSSTrustNetworkValCExportClassInternalOperatorView,
		operability.OSSTrustNetworkValCExportClassEnterpriseCustomerView,
		operability.OSSTrustNetworkValCExportClassMaintainerView,
		operability.OSSTrustNetworkValCExportClassPartnerView,
		operability.OSSTrustNetworkValCExportClassPublicSummaryView,
		operability.OSSTrustNetworkValCExportClassUnsupportedView,
	}
}

func ossTrustNetworkValCSuggestionClasses() []string {
	return []string{
		operability.OSSTrustNetworkValCSuggestionClassVersionUpgrade,
		operability.OSSTrustNetworkValCSuggestionClassPinOrHold,
		operability.OSSTrustNetworkValCSuggestionClassReplaceDependency,
		operability.OSSTrustNetworkValCSuggestionClassMaintainerContact,
		operability.OSSTrustNetworkValCSuggestionClassReviewRequired,
		operability.OSSTrustNetworkValCSuggestionClassNoAction,
		operability.OSSTrustNetworkValCSuggestionClassUnsupported,
	}
}

func ossTrustNetworkValCProposalStates() []string {
	return []string{
		operability.OSSTrustNetworkValCProposalStateProposalReady,
		operability.OSSTrustNetworkValCProposalStateNeedsReview,
		operability.OSSTrustNetworkValCProposalStateUnsupported,
		operability.OSSTrustNetworkValCProposalStateBlocked,
		operability.OSSTrustNetworkValCProposalStateUnknown,
	}
}

func ossTrustNetworkValCLocalOverrideStates() []string {
	return []string{
		operability.OSSTrustNetworkValCOverrideStateNoOverride,
		operability.OSSTrustNetworkValCOverrideStateOverridePresent,
		operability.OSSTrustNetworkValCOverrideStateOverrideRequiresReview,
		operability.OSSTrustNetworkValCOverrideStateOverrideRejected,
		operability.OSSTrustNetworkValCOverrideStateUnsupported,
		operability.OSSTrustNetworkValCOverrideStateUnknown,
	}
}

func ossTrustNetworkValCRiskClasses() []string {
	return []string{
		operability.OSSTrustNetworkValCRiskClassLow,
		operability.OSSTrustNetworkValCRiskClassMedium,
		operability.OSSTrustNetworkValCRiskClassHigh,
	}
}

func buildOSSTrustNetworkValCModel() operability.OSSTrustNetworkValCCore {
	model := operability.OSSTrustNetworkValCCoreModel()
	return operability.ComputeOSSTrustNetworkValCCore(model)
}

func (s server) ossTrustNetworkValCStatusHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildOSSTrustNetworkValCStatus())
}

func (s server) ossTrustNetworkValCProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildOSSTrustNetworkValCProofs())
}

func buildOSSTrustNetworkValCStatus() ossTrustNetworkValCStatusResponse {
	model := buildOSSTrustNetworkValCModel()
	limitations := []string{
		"Val C defines bounded remediation and ecosystem visibility only and does not implement final OSTN gates, integrated closure, or Točka 10.",
		"Visibility, remediation suggestions, proposal descriptors, and local overrides remain advisory and bounded by explicit no-hidden-mutation and no-overclaim discipline.",
	}
	return ossTrustNetworkValCStatusResponse{
		SchemaVersion: ossTrustNetworkValCStatusSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     ossTrustNetworkValCAllSurfaceRefs(),
		Limitations:   limitations,
	}
}

func buildOSSTrustNetworkValCProofs() ossTrustNetworkValCProofsResponse {
	model := buildOSSTrustNetworkValCModel()
	limitations := []string{
		"Val C keeps Točka 9 incomplete and reserves any final pass semantics for later integrated closure waves only.",
		"Visibility, remediation, proposal, and override surfaces remain bounded advisory projections and do not become canonical truth, approval authority, or hidden mutation paths.",
		"Val D through Val E and Točka 10 remain out of scope here.",
	}
	currentState := operability.EvaluateOSSTrustNetworkValCProofsState(model, limitations)
	return ossTrustNetworkValCProofsResponse{
		SchemaVersion:                   ossTrustNetworkValCProofsSchema,
		GeneratedAt:                     publicSampleTime(),
		CurrentState:                    currentState,
		Point9State:                     model.Point9State,
		DependencyState:                 model.DependencyState,
		ValBCurrentState:                model.Dependency.ValBCurrentState,
		ValBPoint9State:                 model.Dependency.ValBPoint9State,
		ValBDependencyState:             model.Dependency.ValBDependencyState,
		ValBCandidateSignalIntakeState:  model.Dependency.ValBCandidateSignalIntakeState,
		ValBReviewWorkflowState:         model.Dependency.ValBReviewWorkflowState,
		ValBSharedVEXTriageState:        model.Dependency.ValBSharedVEXTriageState,
		ValBSourceWeightingState:        model.Dependency.ValBSourceWeightingState,
		ValBLocalApplicabilityState:     model.Dependency.ValBLocalApplicabilityState,
		ValBPropagationExchangeState:    model.Dependency.ValBPropagationExchangeState,
		ValBSupersessionRevocationState: model.Dependency.ValBSupersessionRevocationState,
		ValBReviewerAuditabilityState:   model.Dependency.ValBReviewerAuditabilityState,
		ValBNoOverclaimState:            model.Dependency.ValBNoOverclaimState,
		TrustVisibilityState:            model.TrustVisibilityState,
		PackageTrustStatusState:         model.PackageTrustStatusState,
		ExportBoundaryState:             model.ExportBoundaryState,
		RemediationSuggestionState:      model.RemediationSuggestionState,
		PRProposalState:                 model.PRProposalState,
		LocalOverrideState:              model.LocalOverrideState,
		RemediationSafetyState:          model.RemediationSafetyState,
		EcosystemConsistencyState:       model.EcosystemConsistencyState,
		NoOverclaimState:                model.NoOverclaimState,
		VisibilityState:                 model.TrustVisibility.VisibilityState,
		PackageStatusClass:              model.PackageTrustStatus.StatusClass,
		ExportClass:                     model.ExportBoundary.ExportClass,
		SuggestionClass:                 model.RemediationSuggestion.SuggestionClass,
		SuggestionConfidenceClass:       model.RemediationSuggestion.ConfidenceClass,
		ProposalState:                   model.PRProposal.ProposalState,
		LocalOverrideVisibilityState:    model.LocalOverride.OverrideState,
		RemediationRiskClass:            model.RemediationSafety.RiskClass,
		SupportedVisibilityStates:       ossTrustNetworkValCVisibilityStates(),
		SupportedPackageStatusClasses:   ossTrustNetworkValCPackageStatusClasses(),
		SupportedExportClasses:          ossTrustNetworkValCExportClasses(),
		SupportedSuggestionClasses:      ossTrustNetworkValCSuggestionClasses(),
		SupportedProposalStates:         ossTrustNetworkValCProposalStates(),
		SupportedLocalOverrideStates:    ossTrustNetworkValCLocalOverrideStates(),
		SupportedRiskClasses:            ossTrustNetworkValCRiskClasses(),
		SurfaceRefs:                     model.ProofSurfaceRefs,
		EvidenceRefs:                    model.EvidenceRefs,
		BlockingReasons:                 model.BlockingReasons,
		WhyPoint9NotComplete:            model.WhyPoint9NotComplete,
		Limitations:                     limitations,
		ProjectionDisclaimer:            model.ProjectionDisclaimer,
		IntegrationSummary: []string{
			"Val C adds bounded ecosystem visibility, package trust summaries, advisory remediation suggestions, reviewer-required PR proposal descriptors, local override visibility, and consistency discipline on top of exact Val B reviewed intelligence.",
			"Val C keeps all visibility and remediation surfaces evidence-linked, local-context-bounded, no-hidden-mutation, and not canonical truth.",
		},
	}
}
