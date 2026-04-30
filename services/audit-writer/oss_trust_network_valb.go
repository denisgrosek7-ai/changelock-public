package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	ossTrustNetworkValBStatusSchema = "point9.oss_trust_network.valb.status.v1"
	ossTrustNetworkValBProofsSchema = "point9.oss_trust_network.valb.proofs.v1"
)

type ossTrustNetworkValBStatusResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.OSSTrustNetworkValBCore `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type ossTrustNetworkValBProofsResponse struct {
	SchemaVersion                     string    `json:"schema_version"`
	GeneratedAt                       time.Time `json:"generated_at"`
	CurrentState                      string    `json:"current_state"`
	Point9State                       string    `json:"point_9_state"`
	DependencyState                   string    `json:"dependency_state"`
	ValACurrentState                  string    `json:"vala_current_state"`
	ValAPoint9State                   string    `json:"vala_point_9_state"`
	ValADependencyState               string    `json:"vala_dependency_state"`
	ValAReleaseTrustIntakeState       string    `json:"vala_release_trust_intake_state"`
	ValASigningSignalState            string    `json:"vala_signing_signal_state"`
	ValAMaintainerState               string    `json:"vala_maintainer_attestation_state"`
	ValAProvenanceState               string    `json:"vala_provenance_material_state"`
	ValARegistryDescriptorState       string    `json:"vala_registry_descriptor_state"`
	ValARegistryMetadataState         string    `json:"vala_registry_metadata_state"`
	ValATypoWarningState              string    `json:"vala_typo_squatting_warning_state"`
	ValADriftSignalState              string    `json:"vala_drift_signal_state"`
	ValANoOverclaimState              string    `json:"vala_no_overclaim_state"`
	CandidateSignalIntakeState        string    `json:"candidate_signal_intake_state"`
	ReviewWorkflowState               string    `json:"review_workflow_state"`
	SharedVEXTriageState              string    `json:"shared_vex_triage_state"`
	SourceWeightingState              string    `json:"source_weighting_state"`
	LocalApplicabilityState           string    `json:"local_applicability_state"`
	PropagationExchangeState          string    `json:"propagation_exchange_state"`
	SupersessionRevocationState       string    `json:"supersession_revocation_state"`
	ReviewerAuditabilityState         string    `json:"reviewer_auditability_state"`
	NoOverclaimState                  string    `json:"no_overclaim_state"`
	CandidateIntakeState              string    `json:"candidate_intake_state"`
	CandidateSourceClass              string    `json:"candidate_source_class"`
	CandidateFreshness                string    `json:"candidate_freshness"`
	ReviewState                       string    `json:"review_state"`
	ReviewerDecisionState             string    `json:"reviewer_decision_state"`
	SharedVEXState                    string    `json:"shared_vex_state"`
	SourceClass                       string    `json:"source_class"`
	SourceWeightClass                 string    `json:"source_weight_class"`
	LocalApplicabilityStatus          string    `json:"local_applicability_status"`
	PropagationState                  string    `json:"propagation_state"`
	LifecycleState                    string    `json:"lifecycle_state"`
	ReviewerRoleClass                 string    `json:"reviewer_role_class"`
	SupportedCandidateSources         []string  `json:"supported_candidate_sources,omitempty"`
	SupportedCandidateIntakeStates    []string  `json:"supported_candidate_intake_states,omitempty"`
	SupportedReviewStates             []string  `json:"supported_review_states,omitempty"`
	SupportedReviewerDecisionStates   []string  `json:"supported_reviewer_decision_states,omitempty"`
	SupportedSharedVEXStates          []string  `json:"supported_shared_vex_states,omitempty"`
	SupportedSourceWeightClasses      []string  `json:"supported_source_weight_classes,omitempty"`
	SupportedLocalApplicabilityStates []string  `json:"supported_local_applicability_states,omitempty"`
	SupportedPropagationStates        []string  `json:"supported_propagation_states,omitempty"`
	SurfaceRefs                       []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                      []string  `json:"evidence_refs,omitempty"`
	BlockingReasons                   []string  `json:"blocking_reasons,omitempty"`
	WhyPoint9NotComplete              []string  `json:"why_point_9_not_complete,omitempty"`
	Limitations                       []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer              string    `json:"projection_disclaimer"`
	IntegrationSummary                []string  `json:"integration_summary,omitempty"`
}

func ossTrustNetworkValBAllSurfaceRefs() []string {
	return operability.OSSTrustNetworkValBProofSurfaceRefs()
}

func ossTrustNetworkValBCandidateSources() []string {
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

func ossTrustNetworkValBCandidateIntakeStates() []string {
	return []string{
		operability.OSSTrustNetworkValBCandidateIntakeStateReceived,
		operability.OSSTrustNetworkValBCandidateIntakeStateNormalized,
		operability.OSSTrustNetworkValBCandidateIntakeStateRejectedAtIntake,
		operability.OSSTrustNetworkValBCandidateIntakeStateUnsupported,
		operability.OSSTrustNetworkValBCandidateIntakeStateStale,
		operability.OSSTrustNetworkValBCandidateIntakeStateMalformed,
		operability.OSSTrustNetworkValBCandidateIntakeStateUnknown,
	}
}

func ossTrustNetworkValBReviewStates() []string {
	return []string{
		operability.OSSTrustNetworkValBReviewStateCandidate,
		operability.OSSTrustNetworkValBReviewStateInReview,
		operability.OSSTrustNetworkValBReviewStateReviewed,
		operability.OSSTrustNetworkValBReviewStateRejected,
		operability.OSSTrustNetworkValBReviewStateSuperseded,
		operability.OSSTrustNetworkValBReviewStateRevoked,
	}
}

func ossTrustNetworkValBReviewerDecisionStates() []string {
	return []string{
		operability.OSSTrustNetworkValBReviewerDecisionStateNone,
		operability.OSSTrustNetworkValBReviewerDecisionStateAccepted,
		operability.OSSTrustNetworkValBReviewerDecisionStateRejected,
		operability.OSSTrustNetworkValBReviewerDecisionStateSuperseded,
		operability.OSSTrustNetworkValBReviewerDecisionStateRevoked,
		operability.OSSTrustNetworkValBReviewerDecisionStateNeedsMoreEvidence,
	}
}

func ossTrustNetworkValBSharedVEXStates() []string {
	return []string{
		operability.OSSTrustNetworkValBSharedVEXStateCandidate,
		operability.OSSTrustNetworkValBSharedVEXStateReviewed,
		operability.OSSTrustNetworkValBSharedVEXStateRejected,
		operability.OSSTrustNetworkValBSharedVEXStateSuperseded,
		operability.OSSTrustNetworkValBSharedVEXStateRevoked,
		operability.OSSTrustNetworkValBSharedVEXStateUnsupported,
		operability.OSSTrustNetworkValBSharedVEXStateUnknown,
	}
}

func ossTrustNetworkValBSourceWeightClasses() []string {
	return []string{
		operability.OSSTrustNetworkValBSourceWeightClassLow,
		operability.OSSTrustNetworkValBSourceWeightClassMedium,
		operability.OSSTrustNetworkValBSourceWeightClassHigh,
		operability.OSSTrustNetworkValBSourceWeightClassBounded,
	}
}

func ossTrustNetworkValBLocalApplicabilityStates() []string {
	return []string{
		operability.OSSTrustNetworkValBLocalApplicabilityStatusApplicable,
		operability.OSSTrustNetworkValBLocalApplicabilityStatusNotApplicable,
		operability.OSSTrustNetworkValBLocalApplicabilityStatusUnknown,
		operability.OSSTrustNetworkValBLocalApplicabilityStatusNeedsLocalReview,
		operability.OSSTrustNetworkValBLocalApplicabilityStatusUnsupported,
	}
}

func ossTrustNetworkValBPropagationStates() []string {
	return []string{
		operability.OSSTrustNetworkValBPropagationStateNotShared,
		operability.OSSTrustNetworkValBPropagationStateCandidateExchange,
		operability.OSSTrustNetworkValBPropagationStateReviewedExchange,
		operability.OSSTrustNetworkValBPropagationStateRejected,
		operability.OSSTrustNetworkValBPropagationStateRevoked,
		operability.OSSTrustNetworkValBPropagationStateSuperseded,
		operability.OSSTrustNetworkValBPropagationStateUnsupported,
		operability.OSSTrustNetworkValBPropagationStateUnknown,
	}
}

func buildOSSTrustNetworkValBModel() operability.OSSTrustNetworkValBCore {
	model := operability.OSSTrustNetworkValBCoreModel()
	return operability.ComputeOSSTrustNetworkValBCore(model)
}

func (s server) ossTrustNetworkValBStatusHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildOSSTrustNetworkValBStatus())
}

func (s server) ossTrustNetworkValBProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildOSSTrustNetworkValBProofs())
}

func buildOSSTrustNetworkValBStatus() ossTrustNetworkValBStatusResponse {
	model := buildOSSTrustNetworkValBModel()
	limitations := []string{
		"Val B defines bounded shared reviewed intelligence only and does not implement dashboards, remediation workflows, final closure, or Točka 10.",
		"Shared intelligence remains advisory and bounded by local applicability, source weighting, review workflow, and explicit no-overclaim discipline.",
	}
	return ossTrustNetworkValBStatusResponse{
		SchemaVersion: ossTrustNetworkValBStatusSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     ossTrustNetworkValBAllSurfaceRefs(),
		Limitations:   limitations,
	}
}

func buildOSSTrustNetworkValBProofs() ossTrustNetworkValBProofsResponse {
	model := buildOSSTrustNetworkValBModel()
	limitations := []string{
		"Val B keeps Točka 9 incomplete and reserves any final pass semantics for later integrated closure waves only.",
		"Shared reviewed intelligence remains bounded advisory exchange and does not become canonical truth, certification, or approval authority.",
		"Val C through Val E and Točka 10 remain out of scope here.",
	}
	currentState := operability.EvaluateOSSTrustNetworkValBProofsState(model, limitations)
	return ossTrustNetworkValBProofsResponse{
		SchemaVersion:                     ossTrustNetworkValBProofsSchema,
		GeneratedAt:                       publicSampleTime(),
		CurrentState:                      currentState,
		Point9State:                       model.Point9State,
		DependencyState:                   model.DependencyState,
		ValACurrentState:                  model.Dependency.ValACurrentState,
		ValAPoint9State:                   model.Dependency.ValAPoint9State,
		ValADependencyState:               model.Dependency.ValADependencyState,
		ValAReleaseTrustIntakeState:       model.Dependency.ValAReleaseTrustIntakeState,
		ValASigningSignalState:            model.Dependency.ValASigningSignalState,
		ValAMaintainerState:               model.Dependency.ValAMaintainerState,
		ValAProvenanceState:               model.Dependency.ValAProvenanceState,
		ValARegistryDescriptorState:       model.Dependency.ValARegistryDescriptorState,
		ValARegistryMetadataState:         model.Dependency.ValARegistryMetadataState,
		ValATypoWarningState:              model.Dependency.ValATypoWarningState,
		ValADriftSignalState:              model.Dependency.ValADriftSignalState,
		ValANoOverclaimState:              model.Dependency.ValANoOverclaimState,
		CandidateSignalIntakeState:        model.CandidateSignalIntakeState,
		ReviewWorkflowState:               model.ReviewWorkflowState,
		SharedVEXTriageState:              model.SharedVEXTriageState,
		SourceWeightingState:              model.SourceWeightingState,
		LocalApplicabilityState:           model.LocalApplicabilityState,
		PropagationExchangeState:          model.PropagationExchangeState,
		SupersessionRevocationState:       model.SupersessionRevocationState,
		ReviewerAuditabilityState:         model.ReviewerAuditabilityState,
		NoOverclaimState:                  model.NoOverclaimState,
		CandidateIntakeState:              model.CandidateSignalIntake.IntakeState,
		CandidateSourceClass:              model.CandidateSignalIntake.CandidateSourceClass,
		CandidateFreshness:                model.CandidateSignalIntake.FreshnessState,
		ReviewState:                       model.ReviewWorkflow.ReviewState,
		ReviewerDecisionState:             model.ReviewWorkflow.ReviewerDecisionState,
		SharedVEXState:                    model.SharedVEXTriage.ReviewState,
		SourceClass:                       model.SourceWeighting.SourceClass,
		SourceWeightClass:                 model.SourceWeighting.SourceWeightClass,
		LocalApplicabilityStatus:          model.LocalApplicability.ApplicabilityState,
		PropagationState:                  model.PropagationExchange.PropagationState,
		LifecycleState:                    model.SupersessionRevocation.LifecycleState,
		ReviewerRoleClass:                 model.ReviewerAuditability.ReviewerRoleClass,
		SupportedCandidateSources:         ossTrustNetworkValBCandidateSources(),
		SupportedCandidateIntakeStates:    ossTrustNetworkValBCandidateIntakeStates(),
		SupportedReviewStates:             ossTrustNetworkValBReviewStates(),
		SupportedReviewerDecisionStates:   ossTrustNetworkValBReviewerDecisionStates(),
		SupportedSharedVEXStates:          ossTrustNetworkValBSharedVEXStates(),
		SupportedSourceWeightClasses:      ossTrustNetworkValBSourceWeightClasses(),
		SupportedLocalApplicabilityStates: ossTrustNetworkValBLocalApplicabilityStates(),
		SupportedPropagationStates:        ossTrustNetworkValBPropagationStates(),
		SurfaceRefs:                       model.ProofSurfaceRefs,
		EvidenceRefs:                      model.EvidenceRefs,
		BlockingReasons:                   model.BlockingReasons,
		WhyPoint9NotComplete:              model.WhyPoint9NotComplete,
		Limitations:                       limitations,
		ProjectionDisclaimer:              model.ProjectionDisclaimer,
		IntegrationSummary: []string{
			"Val B adds bounded candidate intake, reviewed shared intelligence workflow, shared VEX, source weighting, local applicability, propagation, lifecycle, reviewer auditability, and no-overclaim disciplines on top of exact Val A dependency.",
			"Shared network signals remain advisory reviewed intelligence only and do not become canonical truth, enterprise override, badge marketing, or final Point 9 closure here.",
		},
	}
}
