package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	ossTrustNetworkValEClosureSchema = "point9.oss_trust_network.vale.closure.v1"
	ossTrustNetworkValEProofsSchema  = "point9.oss_trust_network.vale.proofs.v1"
)

type ossTrustNetworkValEClosureResponse struct {
	SchemaVersion string                                           `json:"schema_version"`
	GeneratedAt   time.Time                                        `json:"generated_at"`
	CurrentState  string                                           `json:"current_state"`
	Model         operability.OSSTrustNetworkValEIntegratedClosure `json:"model"`
	RouteRefs     []string                                         `json:"route_refs,omitempty"`
	Limitations   []string                                         `json:"limitations,omitempty"`
}

type ossTrustNetworkValEProofsResponse struct {
	SchemaVersion                           string    `json:"schema_version"`
	GeneratedAt                             time.Time `json:"generated_at"`
	CurrentState                            string    `json:"current_state"`
	Point9State                             string    `json:"point_9_state"`
	Point9PassAllowed                       bool      `json:"point_9_pass_allowed"`
	Point9PassReason                        string    `json:"point_9_pass_reason"`
	ClosureState                            string    `json:"closure_state"`
	DependencyState                         string    `json:"dependency_state"`
	Val0SourceState                         string    `json:"val0_source_state"`
	ValASourceState                         string    `json:"vala_source_state"`
	ValBSourceState                         string    `json:"valb_source_state"`
	ValCSourceState                         string    `json:"valc_source_state"`
	ValDSourceState                         string    `json:"vald_source_state"`
	IntegratedClosureState                  string    `json:"integrated_closure_state"`
	CanonicalBoundaryState                  string    `json:"canonical_boundary_state"`
	EvidenceQualityState                    string    `json:"evidence_quality_state"`
	NoOverclaimState                        string    `json:"no_overclaim_state"`
	FinalPassRuleState                      string    `json:"final_pass_rule_state"`
	ValDCurrentState                        string    `json:"vald_current_state"`
	ValDPoint9State                         string    `json:"vald_point_9_state"`
	ValDDependencyState                     string    `json:"vald_dependency_state"`
	ValDSignalCorrectnessState              string    `json:"vald_signal_correctness_state"`
	ValDReleaseFoundationState              string    `json:"vald_release_foundation_state"`
	ValDReviewedIntelligenceState           string    `json:"vald_reviewed_intelligence_state"`
	ValDPropagationSafetyState              string    `json:"vald_propagation_safety_state"`
	ValDRemediationPRSafetyState            string    `json:"vald_remediation_pr_safety_state"`
	ValDEcosystemVisibilityConsistencyState string    `json:"vald_ecosystem_visibility_consistency_state"`
	ValDEvidenceQualityState                string    `json:"vald_evidence_quality_state"`
	ValDNoOverclaimState                    string    `json:"vald_no_overclaim_state"`
	SurfaceRefs                             []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                            []string  `json:"evidence_refs,omitempty"`
	BlockingReasons                         []string  `json:"blocking_reasons,omitempty"`
	WhyPoint9Pass                           []string  `json:"why_point_9_pass,omitempty"`
	Limitations                             []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer                    string    `json:"projection_disclaimer"`
	IntegratedClosureSummary                []string  `json:"integrated_closure_summary,omitempty"`
}

func ossTrustNetworkValEAllSurfaceRefs() []string {
	return operability.OSSTrustNetworkValEProofSurfaceRefs()
}

func buildOSSTrustNetworkValEModel() operability.OSSTrustNetworkValEIntegratedClosure {
	model := operability.OSSTrustNetworkValEIntegratedClosureModel()
	return operability.ComputeOSSTrustNetworkValEClosure(model)
}

func (s server) ossTrustNetworkValEClosureHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildOSSTrustNetworkValEClosure())
}

func (s server) ossTrustNetworkValEProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildOSSTrustNetworkValEProofs())
}

func buildOSSTrustNetworkValEClosure() ossTrustNetworkValEClosureResponse {
	model := buildOSSTrustNetworkValEModel()
	limitations := []string{
		"Val E is the integrated closure for Točka 9 only and does not implement Točka 10 or any live registry connector behavior.",
		"Integrated closure remains bounded, evidence-linked, and non-authoritative outside the canonical execution, audit, and evidence spine.",
		"It does not authorize deployment or production use and does not create regulator-facing approvals, legal or IP clearance, or promotional badge semantics.",
	}
	return ossTrustNetworkValEClosureResponse{
		SchemaVersion: ossTrustNetworkValEClosureSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     ossTrustNetworkValEAllSurfaceRefs(),
		Limitations:   limitations,
	}
}

func buildOSSTrustNetworkValEProofs() ossTrustNetworkValEProofsResponse {
	model := buildOSSTrustNetworkValEModel()
	limitations := []string{
		"Only Val E may emit point_9_pass and Točka 9 becomes complete only when the Val E final pass rule is active.",
		"Val 0 through Val D remain prerequisites and cannot close Točka 9 on their own.",
		"Integrated closure remains bounded and does not authorize deployment, production use, regulator-facing approvals, legal or IP clearance, promotional badge semantics, or universal authority claims.",
	}
	return ossTrustNetworkValEProofsResponse{
		SchemaVersion:                           ossTrustNetworkValEProofsSchema,
		GeneratedAt:                             publicSampleTime(),
		CurrentState:                            model.CurrentState,
		Point9State:                             model.Point9State,
		Point9PassAllowed:                       model.Point9PassAllowed,
		Point9PassReason:                        model.Point9PassReason,
		ClosureState:                            model.ClosureState,
		DependencyState:                         model.DependencyState,
		Val0SourceState:                         model.Val0SourceState,
		ValASourceState:                         model.ValASourceState,
		ValBSourceState:                         model.ValBSourceState,
		ValCSourceState:                         model.ValCSourceState,
		ValDSourceState:                         model.ValDSourceState,
		IntegratedClosureState:                  model.IntegratedClosureState,
		CanonicalBoundaryState:                  model.CanonicalBoundaryState,
		EvidenceQualityState:                    model.EvidenceQualityState,
		NoOverclaimState:                        model.NoOverclaimState,
		FinalPassRuleState:                      model.FinalPassRuleState,
		ValDCurrentState:                        model.ValDSource.CurrentState,
		ValDPoint9State:                         model.ValDSource.Point9State,
		ValDDependencyState:                     model.ValDSource.DependencyState,
		ValDSignalCorrectnessState:              model.ValDSource.SignalCorrectnessState,
		ValDReleaseFoundationState:              model.ValDSource.ReleaseFoundationState,
		ValDReviewedIntelligenceState:           model.ValDSource.ReviewedIntelligenceState,
		ValDPropagationSafetyState:              model.ValDSource.PropagationSafetyState,
		ValDRemediationPRSafetyState:            model.ValDSource.RemediationPRSafetyState,
		ValDEcosystemVisibilityConsistencyState: model.ValDSource.EcosystemVisibilityConsistencyState,
		ValDEvidenceQualityState:                model.ValDSource.EvidenceQualityState,
		ValDNoOverclaimState:                    model.ValDSource.NoOverclaimState,
		SurfaceRefs:                             model.ProofSurfaceRefs,
		EvidenceRefs:                            model.EvidenceRefs,
		BlockingReasons:                         model.BlockingReasons,
		WhyPoint9Pass:                           model.WhyPoint9Pass,
		Limitations:                             limitations,
		ProjectionDisclaimer:                    model.ProjectionDisclaimer,
		IntegratedClosureSummary:                model.IntegratedClosureSummary,
	}
}
