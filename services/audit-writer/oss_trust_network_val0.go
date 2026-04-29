package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	ossTrustNetworkVal0StatusSchema = "point9.oss_trust_network.val0.status.v1"
	ossTrustNetworkVal0ProofsSchema = "point9.oss_trust_network.val0.proofs.v1"
)

type ossTrustNetworkVal0StatusResponse struct {
	SchemaVersion string                                    `json:"schema_version"`
	GeneratedAt   time.Time                                 `json:"generated_at"`
	CurrentState  string                                    `json:"current_state"`
	Model         operability.OSSTrustNetworkVal0Foundation `json:"model"`
	RouteRefs     []string                                  `json:"route_refs,omitempty"`
	Limitations   []string                                  `json:"limitations,omitempty"`
}

type ossTrustNetworkVal0ProofsResponse struct {
	SchemaVersion              string    `json:"schema_version"`
	GeneratedAt                time.Time `json:"generated_at"`
	CurrentState               string    `json:"current_state"`
	Point9State                string    `json:"point_9_state"`
	DependencyState            string    `json:"dependency_state"`
	Point8CurrentState         string    `json:"point_8_current_state"`
	Point8State                string    `json:"point_8_state"`
	Point8PassAllowed          bool      `json:"point_8_pass_allowed"`
	Point8PassReason           string    `json:"point_8_pass_reason"`
	Point8ClosureState         string    `json:"point_8_closure_state"`
	Point8NoOverclaimState     string    `json:"point_8_no_overclaim_state"`
	SignalContractState        string    `json:"signal_contract_state"`
	TrustMarkingState          string    `json:"trust_marking_state"`
	MaintainerIdentityState    string    `json:"maintainer_identity_state"`
	RegistryFreshnessState     string    `json:"registry_freshness_state"`
	SharedVEXState             string    `json:"shared_vex_state"`
	PropagationState           string    `json:"propagation_state"`
	LocalApplicabilityState    string    `json:"local_applicability_state"`
	NoOverclaimState           string    `json:"no_overclaim_state"`
	SignalReviewState          string    `json:"signal_review_state"`
	SharedVEXReviewState       string    `json:"shared_vex_review_state"`
	RegistryFreshness          string    `json:"registry_signal_freshness"`
	AllowedTrustMarkingClasses []string  `json:"allowed_trust_marking_classes,omitempty"`
	SupportedReviewStates      []string  `json:"supported_review_states,omitempty"`
	SurfaceRefs                []string  `json:"surface_refs,omitempty"`
	EvidenceRefs               []string  `json:"evidence_refs,omitempty"`
	BlockingReasons            []string  `json:"blocking_reasons,omitempty"`
	WhyPoint9NotComplete       []string  `json:"why_point_9_not_complete,omitempty"`
	Limitations                []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer       string    `json:"projection_disclaimer"`
	IntegrationSummary         []string  `json:"integration_summary,omitempty"`
}

func ossTrustNetworkVal0AllSurfaceRefs() []string {
	return operability.OSSTrustNetworkVal0ProofSurfaceRefs()
}

func buildOSSTrustNetworkVal0FoundationModel() operability.OSSTrustNetworkVal0Foundation {
	model := operability.OSSTrustNetworkVal0FoundationModel()
	return operability.ComputeOSSTrustNetworkVal0Foundation(model)
}

func (s server) ossTrustNetworkVal0StatusHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildOSSTrustNetworkVal0Status())
}

func (s server) ossTrustNetworkVal0ProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildOSSTrustNetworkVal0Proofs())
}

func buildOSSTrustNetworkVal0Status() ossTrustNetworkVal0StatusResponse {
	model := buildOSSTrustNetworkVal0FoundationModel()
	limitations := []string{
		"Val 0 defines only the OSS trust discipline foundation and does not implement registry connectors, signing integrations, provenance verification engines, typo-squatting detection engines, shared reviewed intelligence workflow, dashboards, or later-wave closure.",
		"OSS trust network signals remain advisory or projection-only and cannot certify packages, approve production or deployment, create enterprise policy authority, or mutate canonical evidence.",
	}
	return ossTrustNetworkVal0StatusResponse{
		SchemaVersion: ossTrustNetworkVal0StatusSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     ossTrustNetworkVal0AllSurfaceRefs(),
		Limitations:   limitations,
	}
}

func buildOSSTrustNetworkVal0Proofs() ossTrustNetworkVal0ProofsResponse {
	model := buildOSSTrustNetworkVal0FoundationModel()
	limitations := []string{
		"Val 0 keeps Točka 9 incomplete and reserves any final pass semantics for later integrated closure waves only.",
		"Registry metadata, maintainer attestations, shared VEX, and network propagation remain bounded evidence-linked signals rather than canonical truth or enterprise approval authority.",
		"Točka 10 and every later OSTN wave remain out of scope here.",
	}
	currentState := operability.EvaluateOSSTrustNetworkVal0ProofsState(model, limitations)
	return ossTrustNetworkVal0ProofsResponse{
		SchemaVersion:              ossTrustNetworkVal0ProofsSchema,
		GeneratedAt:                publicSampleTime(),
		CurrentState:               currentState,
		Point9State:                model.Point9State,
		DependencyState:            model.DependencyState,
		Point8CurrentState:         model.Dependency.CurrentState,
		Point8State:                model.Dependency.Point8State,
		Point8PassAllowed:          model.Dependency.Point8PassAllowed,
		Point8PassReason:           model.Dependency.Point8PassReason,
		Point8ClosureState:         model.Dependency.ClosureState,
		Point8NoOverclaimState:     model.Dependency.NoOverclaimState,
		SignalContractState:        model.SignalContractState,
		TrustMarkingState:          model.TrustMarkingState,
		MaintainerIdentityState:    model.MaintainerIdentityState,
		RegistryFreshnessState:     model.RegistryFreshnessState,
		SharedVEXState:             model.SharedVEXState,
		PropagationState:           model.PropagationState,
		LocalApplicabilityState:    model.LocalApplicabilityState,
		NoOverclaimState:           model.NoOverclaimState,
		SignalReviewState:          model.SignalContract.ReviewState,
		SharedVEXReviewState:       model.SharedVEXTriage.ReviewState,
		RegistryFreshness:          model.RegistryFreshness.FreshnessState,
		AllowedTrustMarkingClasses: model.TrustMarking.AllowedTrustMarkingClasses,
		SupportedReviewStates:      model.SignalContract.SupportedReviewStates,
		SurfaceRefs:                model.ProofSurfaceRefs,
		EvidenceRefs:               model.EvidenceRefs,
		BlockingReasons:            model.BlockingReasons,
		WhyPoint9NotComplete:       model.WhyPoint9NotComplete,
		Limitations:                limitations,
		ProjectionDisclaimer:       model.ProjectionDisclaimer,
		IntegrationSummary: []string{
			"Val 0 establishes bounded OSS signal contract, trust marking, maintainer identity, registry freshness, shared VEX, propagation, local applicability, and no-overclaim discipline only.",
			"Network, registry, community, and shared intelligence signals remain advisory inputs and do not become canonical truth, broad certification, or enterprise override authority here.",
		},
	}
}
