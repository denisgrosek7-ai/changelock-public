package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	ossTrustNetworkValDStatusSchema = "point9.oss_trust_network.vald.status.v1"
	ossTrustNetworkValDProofsSchema = "point9.oss_trust_network.vald.proofs.v1"
)

type ossTrustNetworkValDStatusResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.OSSTrustNetworkValDCore `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type ossTrustNetworkValDProofsResponse struct {
	SchemaVersion                       string    `json:"schema_version"`
	GeneratedAt                         time.Time `json:"generated_at"`
	CurrentState                        string    `json:"current_state"`
	Point9State                         string    `json:"point_9_state"`
	DependencyState                     string    `json:"dependency_state"`
	ValCCurrentState                    string    `json:"valc_current_state"`
	ValCPoint9State                     string    `json:"valc_point_9_state"`
	ValCDependencyState                 string    `json:"valc_dependency_state"`
	ValCTrustVisibilityState            string    `json:"valc_trust_visibility_state"`
	ValCPackageTrustStatusState         string    `json:"valc_package_trust_status_state"`
	ValCExportBoundaryState             string    `json:"valc_export_boundary_state"`
	ValCRemediationSuggestionState      string    `json:"valc_remediation_suggestion_state"`
	ValCPRProposalState                 string    `json:"valc_pr_proposal_state"`
	ValCLocalOverrideState              string    `json:"valc_local_override_state"`
	ValCRemediationSafetyState          string    `json:"valc_remediation_safety_state"`
	ValCEcosystemConsistencyState       string    `json:"valc_ecosystem_consistency_state"`
	ValCNoOverclaimState                string    `json:"valc_no_overclaim_state"`
	SignalCorrectnessState              string    `json:"signal_correctness_state"`
	ReleaseFoundationState              string    `json:"release_foundation_state"`
	ReviewedIntelligenceState           string    `json:"reviewed_intelligence_state"`
	PropagationSafetyState              string    `json:"propagation_safety_state"`
	RemediationPRSafetyState            string    `json:"remediation_pr_safety_state"`
	EcosystemVisibilityConsistencyState string    `json:"ecosystem_visibility_consistency_state"`
	EvidenceQualityState                string    `json:"evidence_quality_state"`
	NoOverclaimState                    string    `json:"no_overclaim_state"`
	SignalLifecycleState                string    `json:"signal_lifecycle_state"`
	ReviewState                         string    `json:"review_state"`
	ReviewerDecisionState               string    `json:"reviewer_decision_state"`
	PropagationState                    string    `json:"propagation_state"`
	PackageStatusClass                  string    `json:"package_status_class"`
	ExportClass                         string    `json:"export_class"`
	SuggestionClass                     string    `json:"suggestion_class"`
	ProposalState                       string    `json:"proposal_state"`
	SupportedSignalLifecycleStates      []string  `json:"supported_signal_lifecycle_states,omitempty"`
	SupportedSourceClasses              []string  `json:"supported_source_classes,omitempty"`
	SupportedSourceWeightClasses        []string  `json:"supported_source_weight_classes,omitempty"`
	SupportedExportClasses              []string  `json:"supported_export_classes,omitempty"`
	SupportedSuggestionClasses          []string  `json:"supported_suggestion_classes,omitempty"`
	SupportedProposalStates             []string  `json:"supported_proposal_states,omitempty"`
	SurfaceRefs                         []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                        []string  `json:"evidence_refs,omitempty"`
	BlockingReasons                     []string  `json:"blocking_reasons,omitempty"`
	WhyPoint9NotComplete                []string  `json:"why_point_9_not_complete,omitempty"`
	Limitations                         []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer                string    `json:"projection_disclaimer"`
	FinalReadinessSummary               []string  `json:"final_readiness_summary,omitempty"`
}

func ossTrustNetworkValDAllSurfaceRefs() []string {
	return operability.OSSTrustNetworkValDProofSurfaceRefs()
}

func ossTrustNetworkValDSignalLifecycleStates() []string {
	return []string{
		operability.OSSTrustNetworkValDSignalLifecycleCandidate,
		operability.OSSTrustNetworkValDSignalLifecycleReviewed,
		operability.OSSTrustNetworkValDSignalLifecycleRejected,
		operability.OSSTrustNetworkValDSignalLifecycleSuperseded,
		operability.OSSTrustNetworkValDSignalLifecycleRevoked,
		operability.OSSTrustNetworkValDSignalLifecycleUnsupported,
		operability.OSSTrustNetworkValDSignalLifecycleUnknown,
	}
}

func ossTrustNetworkValDSourceClasses() []string {
	return []string{
		operability.OSSTrustNetworkValBCandidateSourceClassMaintainer,
		operability.OSSTrustNetworkValBCandidateSourceClassRegistry,
		operability.OSSTrustNetworkValBCandidateSourceClassCommunity,
		operability.OSSTrustNetworkValBCandidateSourceClassEnterpriseObservation,
		operability.OSSTrustNetworkValBCandidateSourceClassVendor,
		operability.OSSTrustNetworkValBCandidateSourceClassVerifier,
		operability.OSSTrustNetworkValBCandidateSourceClassAutomatedHeuristic,
	}
}

func ossTrustNetworkValDSourceWeightClasses() []string {
	return []string{
		operability.OSSTrustNetworkValBSourceWeightClassLow,
		operability.OSSTrustNetworkValBSourceWeightClassMedium,
		operability.OSSTrustNetworkValBSourceWeightClassHigh,
		operability.OSSTrustNetworkValBSourceWeightClassBounded,
	}
}

func ossTrustNetworkValDExportClasses() []string {
	return []string{
		operability.OSSTrustNetworkValCExportClassInternalOperatorView,
		operability.OSSTrustNetworkValCExportClassEnterpriseCustomerView,
		operability.OSSTrustNetworkValCExportClassMaintainerView,
		operability.OSSTrustNetworkValCExportClassPartnerView,
		operability.OSSTrustNetworkValCExportClassPublicSummaryView,
		operability.OSSTrustNetworkValCExportClassUnsupportedView,
	}
}

func ossTrustNetworkValDSuggestionClasses() []string {
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

func ossTrustNetworkValDProposalStates() []string {
	return []string{
		operability.OSSTrustNetworkValCProposalStateProposalReady,
		operability.OSSTrustNetworkValCProposalStateNeedsReview,
		operability.OSSTrustNetworkValCProposalStateUnsupported,
		operability.OSSTrustNetworkValCProposalStateBlocked,
		operability.OSSTrustNetworkValCProposalStateUnknown,
	}
}

func buildOSSTrustNetworkValDModel() operability.OSSTrustNetworkValDCore {
	model := operability.OSSTrustNetworkValDCoreModel()
	return operability.ComputeOSSTrustNetworkValDCore(model)
}

func (s server) ossTrustNetworkValDStatusHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildOSSTrustNetworkValDStatus())
}

func (s server) ossTrustNetworkValDProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildOSSTrustNetworkValDProofs())
}

func buildOSSTrustNetworkValDStatus() ossTrustNetworkValDStatusResponse {
	model := buildOSSTrustNetworkValDModel()
	limitations := []string{
		"Val D defines the final OSTN readiness gate only and does not implement integrated closure, Val E, or Točka 10.",
		"Readiness, visibility, remediation, and propagation outputs remain bounded advisory projections and do not become canonical truth, approval authority, or hidden mutation paths.",
	}
	return ossTrustNetworkValDStatusResponse{
		SchemaVersion: ossTrustNetworkValDStatusSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     ossTrustNetworkValDAllSurfaceRefs(),
		Limitations:   limitations,
	}
}

func buildOSSTrustNetworkValDProofs() ossTrustNetworkValDProofsResponse {
	model := buildOSSTrustNetworkValDModel()
	limitations := []string{
		"Val D keeps Točka 9 incomplete and requires Val E for integrated closure and any final pass semantics.",
		"Val D reports bounded final readiness only and does not create canonical truth, approval authority, promotional badge surfaces, or hidden mutation paths.",
		"Val E and Točka 10 remain out of scope here.",
	}
	currentState := operability.EvaluateOSSTrustNetworkValDProofsState(model, limitations)
	return ossTrustNetworkValDProofsResponse{
		SchemaVersion:                       ossTrustNetworkValDProofsSchema,
		GeneratedAt:                         publicSampleTime(),
		CurrentState:                        currentState,
		Point9State:                         model.Point9State,
		DependencyState:                     model.DependencyState,
		ValCCurrentState:                    model.Dependency.ValCCurrentState,
		ValCPoint9State:                     model.Dependency.ValCPoint9State,
		ValCDependencyState:                 model.Dependency.ValCDependencyState,
		ValCTrustVisibilityState:            model.Dependency.ValCTrustVisibilityState,
		ValCPackageTrustStatusState:         model.Dependency.ValCPackageTrustStatusState,
		ValCExportBoundaryState:             model.Dependency.ValCExportBoundaryState,
		ValCRemediationSuggestionState:      model.Dependency.ValCRemediationSuggestionState,
		ValCPRProposalState:                 model.Dependency.ValCPRProposalState,
		ValCLocalOverrideState:              model.Dependency.ValCLocalOverrideState,
		ValCRemediationSafetyState:          model.Dependency.ValCRemediationSafetyState,
		ValCEcosystemConsistencyState:       model.Dependency.ValCEcosystemConsistencyState,
		ValCNoOverclaimState:                model.Dependency.ValCNoOverclaimState,
		SignalCorrectnessState:              model.SignalCorrectnessState,
		ReleaseFoundationState:              model.ReleaseFoundationState,
		ReviewedIntelligenceState:           model.ReviewedIntelligenceState,
		PropagationSafetyState:              model.PropagationSafetyState,
		RemediationPRSafetyState:            model.RemediationPRSafetyState,
		EcosystemVisibilityConsistencyState: model.EcosystemVisibilityConsistencyState,
		EvidenceQualityState:                model.EvidenceQualityState,
		NoOverclaimState:                    model.NoOverclaimState,
		SignalLifecycleState:                model.SignalCorrectness.SignalLifecycleState,
		ReviewState:                         model.SignalCorrectness.ReviewState,
		ReviewerDecisionState:               model.SignalCorrectness.ReviewerDecisionState,
		PropagationState:                    model.PropagationSafety.PropagationState,
		PackageStatusClass:                  model.EcosystemVisibilityConsistency.PackageStatusClass,
		ExportClass:                         model.EcosystemVisibilityConsistency.ExportClass,
		SuggestionClass:                     model.RemediationPRSafety.SuggestionClass,
		ProposalState:                       model.RemediationPRSafety.ProposalState,
		SupportedSignalLifecycleStates:      ossTrustNetworkValDSignalLifecycleStates(),
		SupportedSourceClasses:              ossTrustNetworkValDSourceClasses(),
		SupportedSourceWeightClasses:        ossTrustNetworkValDSourceWeightClasses(),
		SupportedExportClasses:              ossTrustNetworkValDExportClasses(),
		SupportedSuggestionClasses:          ossTrustNetworkValDSuggestionClasses(),
		SupportedProposalStates:             ossTrustNetworkValDProposalStates(),
		SurfaceRefs:                         model.ProofSurfaceRefs,
		EvidenceRefs:                        model.EvidenceRefs,
		BlockingReasons:                     model.BlockingReasons,
		WhyPoint9NotComplete:                model.WhyPoint9NotComplete,
		Limitations:                         limitations,
		ProjectionDisclaimer:                model.ProjectionDisclaimer,
		FinalReadinessSummary:               model.FinalReadinessSummary,
	}
}
