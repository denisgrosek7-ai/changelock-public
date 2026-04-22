package main

import (
	"net/http"
	"time"

	ecosystemcore "github.com/denisgrosek/changelock/internal/ecosystem"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	phase7OSSConnectorsSchema      = "7b.ecosystem_oss_connectors.v1"
	phase7OSSObservationsSchema    = "7b.ecosystem_oss_observation_feed.v1"
	phase7OSSReviewFlowSchema      = "7b.ecosystem_oss_review_flow.v1"
	phase7OSSReviewedSignalsSchema = "7b.ecosystem_oss_reviewed_signals.v1"

	phase7OSSConnectorsStateActive          = "oss_registry_connectors_active"
	phase7OSSConnectorsStatePartial         = "oss_registry_connectors_partial"
	phase7OSSConnectorsStateIncomplete      = "oss_registry_connectors_incomplete"
	phase7OSSObservationsStateActive        = "oss_observation_feed_active"
	phase7OSSObservationsStatePartial       = "oss_observation_feed_partial"
	phase7OSSObservationsStateIncomplete    = "oss_observation_feed_incomplete"
	phase7OSSReviewFlowStateActive          = "oss_review_flow_active"
	phase7OSSReviewFlowStatePartial         = "oss_review_flow_partial"
	phase7OSSReviewFlowStateIncomplete      = "oss_review_flow_incomplete"
	phase7OSSReviewedSignalsStateActive     = "oss_reviewed_signals_active"
	phase7OSSReviewedSignalsStatePartial    = "oss_reviewed_signals_partial"
	phase7OSSReviewedSignalsStateIncomplete = "oss_reviewed_signals_incomplete"
)

type phase7OSSConnectorDetail struct {
	RegistryID       string   `json:"registry_id"`
	CurrentState     string   `json:"current_state"`
	FreshnessState   string   `json:"freshness_state"`
	ProvenanceState  string   `json:"provenance_state"`
	ObservationState string   `json:"observation_state"`
	SourceTrustTier  string   `json:"source_trust_tier"`
	ReasonCodes      []string `json:"reason_codes,omitempty"`
	RouteRefs        []string `json:"route_refs,omitempty"`
	EvidenceRefs     []string `json:"evidence_refs,omitempty"`
	DegradedBehavior string   `json:"degraded_behavior"`
	Limitations      []string `json:"limitations,omitempty"`
}

type phase7OSSConnectorsResponse struct {
	SchemaVersion     string                     `json:"schema_version"`
	GeneratedAt       time.Time                  `json:"generated_at"`
	CurrentState      string                     `json:"current_state"`
	Connectors        []phase7OSSConnectorDetail `json:"connectors,omitempty"`
	CompatibilityRefs []string                   `json:"compatibility_refs,omitempty"`
	AbuseControlRefs  []string                   `json:"abuse_control_refs,omitempty"`
	RolloutRefs       []string                   `json:"rollout_refs,omitempty"`
	Limitations       []string                   `json:"limitations,omitempty"`
}

type phase7OSSObservationItem struct {
	ObservationID   string   `json:"observation_id"`
	RegistryID      string   `json:"registry_id"`
	CurrentState    string   `json:"current_state"`
	ReviewState     string   `json:"review_state"`
	SourceType      string   `json:"source_type"`
	ProvenanceState string   `json:"provenance_state"`
	FreshnessState  string   `json:"freshness_state"`
	Summary         string   `json:"summary"`
	ReasonCodes     []string `json:"reason_codes,omitempty"`
	RouteRefs       []string `json:"route_refs,omitempty"`
	EvidenceRefs    []string `json:"evidence_refs,omitempty"`
	Limitations     []string `json:"limitations,omitempty"`
}

type phase7OSSObservationsResponse struct {
	SchemaVersion            string                     `json:"schema_version"`
	GeneratedAt              time.Time                  `json:"generated_at"`
	CurrentState             string                     `json:"current_state"`
	ObservationPipelineState string                     `json:"observation_pipeline_state"`
	CandidateStates          []string                   `json:"candidate_states,omitempty"`
	CommunityInputState      string                     `json:"community_input_state"`
	Observations             []phase7OSSObservationItem `json:"observations,omitempty"`
	AbuseControlRefs         []string                   `json:"abuse_control_refs,omitempty"`
	Limitations              []string                   `json:"limitations,omitempty"`
}

type phase7OSSReviewFlowStep struct {
	StepID            string   `json:"step_id"`
	CurrentState      string   `json:"current_state"`
	FromState         string   `json:"from_state"`
	ToStates          []string `json:"to_states,omitempty"`
	Requirements      []string `json:"requirements,omitempty"`
	PublicationScopes []string `json:"publication_scopes,omitempty"`
	BoundaryRules     []string `json:"boundary_rules,omitempty"`
}

type phase7OSSReviewFlowResponse struct {
	SchemaVersion            string                    `json:"schema_version"`
	GeneratedAt              time.Time                 `json:"generated_at"`
	CurrentState             string                    `json:"current_state"`
	ObservationPipelineState string                    `json:"observation_pipeline_state"`
	ClaimPipelineState       string                    `json:"claim_pipeline_state"`
	ReviewStates             []string                  `json:"review_states,omitempty"`
	Steps                    []phase7OSSReviewFlowStep `json:"steps,omitempty"`
	PublicationRules         []string                  `json:"publication_rules,omitempty"`
	RolloutRefs              []string                  `json:"rollout_refs,omitempty"`
	ExpandedScopeDeferred    []string                  `json:"expanded_scope_deferred,omitempty"`
	Limitations              []string                  `json:"limitations,omitempty"`
}

type phase7OSSReviewedSignalRecord struct {
	RecordID         string   `json:"record_id"`
	SignalClass      string   `json:"signal_class"`
	CurrentState     string   `json:"current_state"`
	ReviewState      string   `json:"review_state"`
	PublicationScope string   `json:"publication_scope"`
	DeliveryState    string   `json:"delivery_state"`
	Semantics        []string `json:"semantics,omitempty"`
	RouteRefs        []string `json:"route_refs,omitempty"`
	EvidenceRefs     []string `json:"evidence_refs,omitempty"`
	ReasonCodes      []string `json:"reason_codes,omitempty"`
	Limitations      []string `json:"limitations,omitempty"`
}

type phase7OSSReviewedSignalsResponse struct {
	SchemaVersion string                          `json:"schema_version"`
	GeneratedAt   time.Time                       `json:"generated_at"`
	CurrentState  string                          `json:"current_state"`
	Records       []phase7OSSReviewedSignalRecord `json:"records,omitempty"`
	RouteRefs     []string                        `json:"route_refs,omitempty"`
	Limitations   []string                        `json:"limitations,omitempty"`
}

func (s server) phase7OSSConnectorsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7OSSConnectors())
}

func (s server) phase7OSSObservationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7OSSObservations())
}

func (s server) phase7OSSReviewFlowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7OSSReviewFlow())
}

func (s server) phase7OSSReviewedSignalsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase7OSSReviewedSignals())
}

func buildPhase7OSSConnectors() phase7OSSConnectorsResponse {
	return phase7OSSConnectorsResponse{
		SchemaVersion:     phase7OSSConnectorsSchema,
		GeneratedAt:       publicSampleTime(),
		CurrentState:      phase7OSSProjectionState(phase7OSSConnectorsStateActive, phase7OSSConnectorsStatePartial, phase7OSSConnectorsStateIncomplete),
		Connectors:        phase7OSSConnectorDetails(),
		CompatibilityRefs: phase7CompatibilityIDs(ecosystemcore.CompatibilityContractsForGroup("oss")),
		AbuseControlRefs:  phase7AbuseIDs(ecosystemcore.AbuseControlsForGroup("oss")),
		RolloutRefs:       phase7RolloutIDs(ecosystemcore.RolloutContractsForGroup("oss")),
		Limitations: []string{
			"Registry connectors stay bounded to provenance and freshness-aware observation intake and do not create reviewed claims on their own.",
			"Connector posture remains verifier-friendly and reason-coded rather than a universal trust score.",
		},
	}
}

func buildPhase7OSSObservations() phase7OSSObservationsResponse {
	return phase7OSSObservationsResponse{
		SchemaVersion:            phase7OSSObservationsSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             phase7OSSProjectionState(phase7OSSObservationsStateActive, phase7OSSObservationsStatePartial, phase7OSSObservationsStateIncomplete),
		ObservationPipelineState: "observation_pipeline_active",
		CandidateStates:          []string{"candidate", "stale_candidate", "blocked_candidate"},
		CommunityInputState:      "community_candidate_only_review_required",
		Observations:             phase7OSSObservationItems(),
		AbuseControlRefs:         phase7AbuseIDs(ecosystemcore.AbuseControlsForGroup("oss")),
		Limitations: []string{
			"Community and registry input remain candidate observations until explicit review binds evidence and publishes a bounded claim.",
			"Observation feed never auto-promotes to public or verifier-ready reviewed publication.",
		},
	}
}

func buildPhase7OSSReviewFlow() phase7OSSReviewFlowResponse {
	return phase7OSSReviewFlowResponse{
		SchemaVersion:            phase7OSSReviewFlowSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             phase7OSSProjectionState(phase7OSSReviewFlowStateActive, phase7OSSReviewFlowStatePartial, phase7OSSReviewFlowStateIncomplete),
		ObservationPipelineState: "observation_pipeline_active",
		ClaimPipelineState:       "claim_pipeline_active",
		ReviewStates:             []string{"candidate", "reviewed", "rejected", "superseded", "revoked"},
		Steps: []phase7OSSReviewFlowStep{
			{
				StepID:       "candidate_intake",
				CurrentState: "review_step_active",
				FromState:    "candidate",
				ToStates:     []string{"under_review", "blocked_candidate"},
				Requirements: []string{"provenance_tagging", "freshness_marking", "abuse_filtering"},
				BoundaryRules: []string{
					"Candidate observations remain unpublished and review-required.",
					"Blocked candidates remain visible as blocked rather than silently dropped into reviewed delivery.",
				},
			},
			{
				StepID:            "review_decision",
				CurrentState:      "review_step_active",
				FromState:         "under_review",
				ToStates:          []string{"reviewed", "rejected"},
				Requirements:      []string{"bounded_semantics", "evidence_binding", "explicit_reviewer_decision"},
				PublicationScopes: []string{"review_workspace_only"},
				BoundaryRules: []string{
					"Community input cannot bypass reviewer state discipline.",
					"Review must stay explainable and evidence-traceable.",
				},
			},
			{
				StepID:            "publication_and_lifecycle",
				CurrentState:      "review_step_active",
				FromState:         "reviewed",
				ToStates:          []string{"superseded", "revoked"},
				Requirements:      []string{"bounded_publication", "subscriber_delivery", "revocation_visibility"},
				PublicationScopes: []string{"public_bounded", "partner_bounded", "verifier_exportable"},
				BoundaryRules: []string{
					"Superseded and revoked reviewed claims remain visible with explicit lifecycle state.",
					"Automated PR and mutation authority stay outside this bounded review flow.",
				},
			},
		},
		PublicationRules: []string{
			"Reviewed publication requires explicit review and evidence binding.",
			"Candidate observations never become reviewed claims by connector freshness alone.",
			"Superseded and revoked reviewed claims remain verifier-visible rather than silently disappearing.",
		},
		RolloutRefs:           phase7RolloutIDs(ecosystemcore.RolloutContractsForGroup("oss")),
		ExpandedScopeDeferred: []string{"automated_pr_discipline", "community_mutation_without_review", "additional_registry_provider_coverage"},
		Limitations: []string{
			"Review flow remains bounded to reviewed signal publication and does not open a hidden remediation or orchestration path.",
			"Subscriber delivery stays limited to reviewed bounded signals and not raw candidate intake.",
		},
	}
}

func buildPhase7OSSReviewedSignals() phase7OSSReviewedSignalsResponse {
	return phase7OSSReviewedSignalsResponse{
		SchemaVersion: phase7OSSReviewedSignalsSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  phase7OSSProjectionState(phase7OSSReviewedSignalsStateActive, phase7OSSReviewedSignalsStatePartial, phase7OSSReviewedSignalsStateIncomplete),
		Records:       phase7OSSReviewedSignalRecords(),
		RouteRefs: []string{
			"/v1/public/transparency/anchor",
			"/v1/public/proof-portal",
			"/v1/public/claims/summary",
		},
		Limitations: []string{
			"Reviewed signals remain bounded trust publications and not global quality or safety scores.",
			"Rejected candidate observations do not appear in reviewed-signal delivery because they never become published claims.",
		},
	}
}

func phase7OSSProjectionState(activeState, partialState, incompleteState string) string {
	switch ecosystemcore.EvaluateOSSPresenceState() {
	case ecosystemcore.OSSPresenceStateActive:
		return activeState
	case ecosystemcore.OSSPresenceStatePartial:
		return partialState
	default:
		return incompleteState
	}
}

func phase7OSSConnectorDetails() []phase7OSSConnectorDetail {
	signals := ecosystemcore.SignalContractsForGroup("oss")
	observationFailSafe := phase7OSSFailSafe("oss.observation_pipeline")
	baseEvidenceRefs := signalEvidenceRefs(signals, "oss.registry_provenance_observation")
	degradedBehavior := firstNonEmpty(
		observationFailSafe.DegradedBehavior,
		"Keep candidate observations unpublished and explicit when connector freshness or provenance verification degrades.",
	)
	return []phase7OSSConnectorDetail{
		{
			RegistryID:       "npm",
			CurrentState:     "connector_active",
			FreshnessState:   "fresh",
			ProvenanceState:  "provenance_observation_ready",
			ObservationState: "candidate_only_until_review",
			SourceTrustTier:  "bounded_registry_connector",
			ReasonCodes:      []string{"provenance_observation_ready", "freshness_window_enforced", "candidate_only_until_review"},
			RouteRefs:        []string{"/v1/ecosystem/phase7/oss/observations", "/v1/intelligence/supply-chain/patterns"},
			EvidenceRefs:     baseEvidenceRefs,
			DegradedBehavior: degradedBehavior,
			Limitations: []string{
				"Registry observation is not itself a reviewed claim or public trust mark.",
			},
		},
		{
			RegistryID:       "pypi",
			CurrentState:     "connector_active",
			FreshnessState:   "fresh",
			ProvenanceState:  "provenance_observation_ready",
			ObservationState: "candidate_only_until_review",
			SourceTrustTier:  "bounded_registry_connector",
			ReasonCodes:      []string{"provenance_observation_ready", "freshness_window_enforced", "candidate_only_until_review"},
			RouteRefs:        []string{"/v1/ecosystem/phase7/oss/observations", "/v1/intelligence/supply-chain/patterns"},
			EvidenceRefs:     baseEvidenceRefs,
			DegradedBehavior: degradedBehavior,
			Limitations: []string{
				"Connector freshness does not widen a candidate observation into a published trust claim.",
			},
		},
		{
			RegistryID:       "maven",
			CurrentState:     "connector_active",
			FreshnessState:   "fresh",
			ProvenanceState:  "provenance_observation_ready",
			ObservationState: "candidate_only_until_review",
			SourceTrustTier:  "bounded_registry_connector",
			ReasonCodes:      []string{"provenance_observation_ready", "freshness_window_enforced", "candidate_only_until_review"},
			RouteRefs:        []string{"/v1/ecosystem/phase7/oss/observations", "/v1/intelligence/supply-chain/patterns"},
			EvidenceRefs:     baseEvidenceRefs,
			DegradedBehavior: degradedBehavior,
			Limitations: []string{
				"Connector output remains observation-only until review binds bounded semantics and evidence.",
			},
		},
	}
}

func phase7OSSObservationItems() []phase7OSSObservationItem {
	signals := ecosystemcore.SignalContractsForGroup("oss")
	baseEvidenceRefs := signalEvidenceRefs(signals, "oss.registry_provenance_observation")
	return []phase7OSSObservationItem{
		{
			ObservationID:   "npm:leftpad:signing_continuity_shift",
			RegistryID:      "npm",
			CurrentState:    "candidate_observation_active",
			ReviewState:     "candidate",
			SourceType:      "registry_connector",
			ProvenanceState: "provenance_shift_observed",
			FreshnessState:  "fresh",
			Summary:         "Registry intake observed a signing-continuity change for the same package lineage and kept it in candidate staging pending review.",
			ReasonCodes:     []string{"candidate_only_until_review", "signing_continuity_shift", "reason_coded_registry_intelligence"},
			RouteRefs:       []string{"/v1/ecosystem/phase7/oss/review-flow", "/v1/intelligence/supply-chain/patterns"},
			EvidenceRefs:    baseEvidenceRefs,
			Limitations: []string{
				"Candidate observation is not a public reviewed trust statement.",
			},
		},
		{
			ObservationID:   "pypi:requests:provenance_gap",
			RegistryID:      "pypi",
			CurrentState:    "candidate_observation_stale",
			ReviewState:     "stale_candidate",
			SourceType:      "registry_connector",
			ProvenanceState: "provenance_incomplete",
			FreshnessState:  "stale",
			Summary:         "The connector retained a provenance gap observation as stale candidate intake rather than presenting it as reviewed signal delivery.",
			ReasonCodes:     []string{"connector_stale", "candidate_only_until_review", "stale_signal_not_promoted"},
			RouteRefs:       []string{"/v1/ecosystem/phase7/oss/connectors", "/v1/ecosystem/phase7/oss/review-flow"},
			EvidenceRefs:    baseEvidenceRefs,
			Limitations: []string{
				"Stale candidate intake stays visible as stale and never collapses into a silent green state.",
			},
		},
		{
			ObservationID:   "maven:log4j:community_submission_blocked",
			RegistryID:      "maven",
			CurrentState:    "candidate_observation_blocked",
			ReviewState:     "blocked_candidate",
			SourceType:      "community_submission",
			ProvenanceState: "unverified_submission",
			FreshnessState:  "fresh",
			Summary:         "Community-assisted intake remained blocked in candidate staging because provenance and abuse checks did not support reviewer promotion.",
			ReasonCodes:     []string{"community_input_needs_review", "abuse_filter_blocked", "never_auto_publish"},
			RouteRefs:       []string{"/v1/ecosystem/phase7/oss/review-flow"},
			EvidenceRefs:    baseEvidenceRefs,
			Limitations: []string{
				"Community input never becomes reviewed publication without explicit reviewer action and evidence binding.",
			},
		},
	}
}

func phase7OSSReviewedSignalRecords() []phase7OSSReviewedSignalRecord {
	signals := ecosystemcore.SignalContractsForGroup("oss")
	claimEvidenceRefs := signalEvidenceRefs(signals, "oss.reviewed_trust_claim")
	return []phase7OSSReviewedSignalRecord{
		{
			RecordID:         "npm:leftpad:reviewed_claim:2026-04",
			SignalClass:      "oss.reviewed_trust_claim",
			CurrentState:     "reviewed_signal_active",
			ReviewState:      "reviewed",
			PublicationScope: "public_bounded",
			DeliveryState:    "subscriber_delivery_active",
			Semantics: []string{
				"bounded_trust_signal",
				"verifier_friendly",
				"not_a_marketing_quality_score",
			},
			RouteRefs:    []string{"/v1/public/proof-portal", "/v1/public/claims/summary", "/v1/public/transparency/anchor"},
			EvidenceRefs: claimEvidenceRefs,
			ReasonCodes:  []string{"reviewed_claim_published", "evidence_bound", "freshness_disciplined"},
			Limitations: []string{
				"Reviewed publication remains bounded to explicit claim semantics rather than universal trust authority.",
			},
		},
		{
			RecordID:         "pypi:urllib3:reviewed_claim:2026-03",
			SignalClass:      "oss.reviewed_trust_claim",
			CurrentState:     "superseded_signal_visible",
			ReviewState:      "superseded",
			PublicationScope: "verifier_exportable",
			DeliveryState:    "superseded_record_retained_visible",
			Semantics: []string{
				"bounded_trust_signal",
				"lifecycle_state_visible",
				"superseded_not_silently_hidden",
			},
			RouteRefs:    []string{"/v1/public/proof-portal", "/v1/public/claims/summary"},
			EvidenceRefs: claimEvidenceRefs,
			ReasonCodes:  []string{"superseded_claim_visible", "replacement_claim_available"},
			Limitations: []string{
				"Superseded reviewed claims remain visible for verifier traceability and are not presented as current trust posture.",
			},
		},
		{
			RecordID:         "maven:log4j:reviewed_claim:2026-02",
			SignalClass:      "oss.reviewed_trust_claim",
			CurrentState:     "revoked_signal_visible",
			ReviewState:      "revoked",
			PublicationScope: "partner_bounded",
			DeliveryState:    "revoked_record_retained_visible",
			Semantics: []string{
				"bounded_trust_signal",
				"revocation_visible",
				"reviewer_action_retained_visible",
			},
			RouteRefs:    []string{"/v1/public/proof-portal", "/v1/public/transparency/anchor"},
			EvidenceRefs: claimEvidenceRefs,
			ReasonCodes:  []string{"revoked_claim_visible", "revocation_never_collapses_into_generic_success"},
			Limitations: []string{
				"Revoked reviewed signals remain explicit lifecycle records and do not imply current approval.",
			},
		},
	}
}

func phase7OSSFailSafe(surfaceID string) ecosystemcore.FailSafeContract {
	for _, item := range ecosystemcore.FailSafeContractsForGroup("oss") {
		if item.SurfaceID == surfaceID {
			return item
		}
	}
	return ecosystemcore.FailSafeContract{}
}
