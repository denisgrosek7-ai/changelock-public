package main

import (
	"net/http"
	"strings"
	"time"

	formalcore "github.com/denisgrosek/changelock/internal/formal"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	phase8RiskQuantificationSchema           = "8c.formal_institutional_risk_quantification.v1"
	phase8InsuranceExportsSchema             = "8c.formal_institutional_insurance_exports.v1"
	phase8IncidentAttributionSchema          = "8c.formal_institutional_incident_attribution.v1"
	phase8ActuarialBenchmarksSchema          = "8c.formal_institutional_actuarial_benchmarks.v1"
	phase8RiskQuantificationStateActive      = "phase8c_risk_quantification_active"
	phase8RiskQuantificationStatePartial     = "phase8c_risk_quantification_partial"
	phase8RiskQuantificationStateIncomplete  = "phase8c_risk_quantification_incomplete"
	phase8InsuranceExportsStateActive        = "phase8c_insurance_exports_active"
	phase8InsuranceExportsStatePartial       = "phase8c_insurance_exports_partial"
	phase8InsuranceExportsStateIncomplete    = "phase8c_insurance_exports_incomplete"
	phase8IncidentAttributionStateActive     = "phase8c_incident_attribution_active"
	phase8IncidentAttributionStatePartial    = "phase8c_incident_attribution_partial"
	phase8IncidentAttributionStateIncomplete = "phase8c_incident_attribution_incomplete"
	phase8ActuarialBenchmarksStateActive     = "phase8c_actuarial_benchmarks_active"
	phase8ActuarialBenchmarksStatePartial    = "phase8c_actuarial_benchmarks_partial"
	phase8ActuarialBenchmarksStateIncomplete = "phase8c_actuarial_benchmarks_incomplete"
)

type phase8RiskQuantificationResponse struct {
	SchemaVersion            string    `json:"schema_version"`
	GeneratedAt              time.Time `json:"generated_at"`
	CurrentState             string    `json:"current_state"`
	QuantificationState      string    `json:"quantification_state"`
	RiskPostureBand          string    `json:"risk_posture_band"`
	ControlMaturityBand      string    `json:"control_maturity_band"`
	EvidenceCompletenessBand string    `json:"evidence_completeness_band"`
	ClaimSupportabilityBand  string    `json:"claim_supportability_band"`
	CalibrationClass         string    `json:"calibration_class"`
	PremiumModelBoundary     string    `json:"premium_model_boundary"`
	ForbiddenUseNotes        []string  `json:"forbidden_use_notes,omitempty"`
	RouteRefs                []string  `json:"route_refs,omitempty"`
	Limitations              []string  `json:"limitations,omitempty"`
}

type phase8AdverseDecisionSupport struct {
	ReasonCode         string   `json:"reason_code"`
	EvidenceBasis      []string `json:"evidence_basis,omitempty"`
	SufficiencyClass   string   `json:"sufficiency_class"`
	ChallengePath      string   `json:"challenge_path"`
	NextReviewPath     string   `json:"next_review_path"`
	DisclosureBoundary string   `json:"disclosure_boundary"`
}

type phase8InsuranceExportItem struct {
	ExportID                   string                       `json:"export_id"`
	CurrentState               string                       `json:"current_state"`
	Audience                   string                       `json:"audience"`
	ClaimClass                 string                       `json:"claim_class"`
	DisclosureScopeProfile     []string                     `json:"disclosure_scope_profile,omitempty"`
	ReleaseLifecycleState      string                       `json:"release_lifecycle_state"`
	WithdrawalState            string                       `json:"withdrawal_state"`
	ReleaseApprovalRequired    bool                         `json:"release_approval_required"`
	CanCitePublicly            bool                         `json:"can_cite_publicly"`
	AdverseDecisionExplanation phase8AdverseDecisionSupport `json:"adverse_decision_explanation"`
	RouteRefs                  []string                     `json:"route_refs,omitempty"`
	Limitations                []string                     `json:"limitations,omitempty"`
}

type phase8InsuranceExportsResponse struct {
	SchemaVersion         string                      `json:"schema_version"`
	GeneratedAt           time.Time                   `json:"generated_at"`
	CurrentState          string                      `json:"current_state"`
	ExportState           string                      `json:"export_state"`
	DisclosureScopeState  string                      `json:"disclosure_scope_state"`
	ReleaseLifecycleState string                      `json:"release_lifecycle_state"`
	Exports               []phase8InsuranceExportItem `json:"exports,omitempty"`
	RouteRefs             []string                    `json:"route_refs,omitempty"`
	Limitations           []string                    `json:"limitations,omitempty"`
}

type phase8IncidentAttributionItem struct {
	RecordID                 string   `json:"record_id"`
	CurrentState             string   `json:"current_state"`
	AttributionBasisClass    string   `json:"attribution_basis_class"`
	EvidenceSufficiencyClass string   `json:"evidence_sufficiency_class"`
	UnresolvedAmbiguityNote  string   `json:"unresolved_ambiguity_note"`
	NonLegalConclusionMarker string   `json:"non_legal_conclusion_marker"`
	ReasonCodes              []string `json:"reason_codes,omitempty"`
	RouteRefs                []string `json:"route_refs,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

type phase8IncidentAttributionResponse struct {
	SchemaVersion           string                          `json:"schema_version"`
	GeneratedAt             time.Time                       `json:"generated_at"`
	CurrentState            string                          `json:"current_state"`
	AttributionState        string                          `json:"attribution_state"`
	AmbiguityHandlingState  string                          `json:"ambiguity_handling_state"`
	NonLegalConclusionState string                          `json:"non_legal_conclusion_state"`
	Items                   []phase8IncidentAttributionItem `json:"items,omitempty"`
	RouteRefs               []string                        `json:"route_refs,omitempty"`
	Limitations             []string                        `json:"limitations,omitempty"`
}

type phase8ActuarialBenchmarkItem struct {
	BenchmarkID                  string   `json:"benchmark_id"`
	CurrentState                 string   `json:"current_state"`
	MinimumCohortSize            int      `json:"minimum_cohort_size"`
	AggregationScope             string   `json:"aggregation_scope"`
	ReidentificationThreshold    string   `json:"reidentification_risk_threshold"`
	PublicationWithdrawalTrigger string   `json:"publication_withdrawal_trigger"`
	ForbiddenUse                 []string `json:"forbidden_use,omitempty"`
	RouteRefs                    []string `json:"route_refs,omitempty"`
	Limitations                  []string `json:"limitations,omitempty"`
}

type phase8ActuarialBenchmarksResponse struct {
	SchemaVersion            string                         `json:"schema_version"`
	GeneratedAt              time.Time                      `json:"generated_at"`
	CurrentState             string                         `json:"current_state"`
	BenchmarkState           string                         `json:"benchmark_state"`
	PrivacyBoundaryState     string                         `json:"privacy_boundary_state"`
	PublicationBoundaryState string                         `json:"publication_boundary_state"`
	Items                    []phase8ActuarialBenchmarkItem `json:"items,omitempty"`
	RouteRefs                []string                       `json:"route_refs,omitempty"`
	Limitations              []string                       `json:"limitations,omitempty"`
}

func (s server) phase8RiskQuantificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8RiskQuantification())
}

func (s server) phase8InsuranceExportsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8InsuranceExports())
}

func (s server) phase8IncidentAttributionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8IncidentAttribution())
}

func (s server) phase8ActuarialBenchmarksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8ActuarialBenchmarks())
}

func buildPhase8RiskQuantification() phase8RiskQuantificationResponse {
	return phase8RiskQuantificationResponse{
		SchemaVersion:            phase8RiskQuantificationSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             phase8InstitutionalSliceState(phase8RiskQuantificationStateActive, phase8RiskQuantificationStatePartial, phase8RiskQuantificationStateIncomplete),
		QuantificationState:      "bounded_risk_quantification_baseline_active",
		RiskPostureBand:          "moderate_controlled_exposure_band",
		ControlMaturityBand:      "evidence_backed_maturity_band_visible",
		EvidenceCompletenessBand: "review_complete_but_scope_bounded",
		ClaimSupportabilityBand:  "institutional_support_input_only",
		CalibrationClass:         "integration_ready_not_pricing_promise",
		PremiumModelBoundary:     "never_automatic_pricing_promise",
		ForbiddenUseNotes: []string{
			"Do not use as an automatic pricing engine output.",
			"Do not reuse as a public trust ranking or universal organizational safety score.",
		},
		RouteRefs: []string{
			"/v1/formal/phase8/contracts",
			"/v1/formal/phase8/institutional/insurance-exports",
			"/v1/formal/phase8/proofs",
		},
		Limitations: []string{
			"Risk quantification remains a bounded institutional input and does not become a legal verdict, final underwriting decision, or universal score.",
		},
	}
}

func buildPhase8InsuranceExports() phase8InsuranceExportsResponse {
	return phase8InsuranceExportsResponse{
		SchemaVersion:         phase8InsuranceExportsSchema,
		GeneratedAt:           publicSampleTime(),
		CurrentState:          phase8InstitutionalSliceState(phase8InsuranceExportsStateActive, phase8InsuranceExportsStatePartial, phase8InsuranceExportsStateIncomplete),
		ExportState:           "insurer_facing_evidence_exports_active",
		DisclosureScopeState:  "insurer_scoped_export_only",
		ReleaseLifecycleState: "release_and_withdrawal_lifecycle_visible",
		Exports:               phase8InsuranceExportItems(),
		RouteRefs: []string{
			"/v1/formal/phase8/institutional/risk-quantification",
			"/v1/formal/phase8/governance/authority-routing",
			"/v1/formal/phase8/proofs",
		},
		Limitations: []string{
			"Insurance-facing exports remain scoped evidence contracts and do not create automatic pricing, public disclosure, or regulator-facing release authority.",
		},
	}
}

func buildPhase8IncidentAttribution() phase8IncidentAttributionResponse {
	return phase8IncidentAttributionResponse{
		SchemaVersion:           phase8IncidentAttributionSchema,
		GeneratedAt:             publicSampleTime(),
		CurrentState:            phase8InstitutionalSliceState(phase8IncidentAttributionStateActive, phase8IncidentAttributionStatePartial, phase8IncidentAttributionStateIncomplete),
		AttributionState:        "incident_attribution_support_active",
		AmbiguityHandlingState:  "unresolved_ambiguity_visible",
		NonLegalConclusionState: "non_legal_conclusion_marker_active",
		Items:                   phase8IncidentAttributionItems(),
		RouteRefs: []string{
			"/v1/formal/phase8/institutional/insurance-exports",
			"/v1/formal/phase8/governance/model-risk",
			"/v1/formal/phase8/proofs",
		},
		Limitations: []string{
			"Incident attribution support remains evidence-backed and non-legal; unresolved ambiguity stays visible rather than being flattened into definitive blame or legal conclusions.",
		},
	}
}

func buildPhase8ActuarialBenchmarks() phase8ActuarialBenchmarksResponse {
	return phase8ActuarialBenchmarksResponse{
		SchemaVersion:            phase8ActuarialBenchmarksSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             phase8InstitutionalSliceState(phase8ActuarialBenchmarksStateActive, phase8ActuarialBenchmarksStatePartial, phase8ActuarialBenchmarksStateIncomplete),
		BenchmarkState:           "actuarial_benchmark_discipline_active",
		PrivacyBoundaryState:     "aggregate_only_and_reidentification_guarded",
		PublicationBoundaryState: "withdrawal_trigger_visible",
		Items:                    phase8ActuarialBenchmarkItems(),
		RouteRefs: []string{
			"/v1/formal/phase8/institutional/risk-quantification",
			"/v1/formal/phase8/institutional/insurance-exports",
			"/v1/formal/phase8/proofs",
		},
		Limitations: []string{
			"Actuarial benchmarks remain aggregated and privacy-guarded; they do not become tenant-level pricing outputs, public rankings, or raw cross-organization disclosure streams.",
		},
	}
}

func phase8InstitutionalSliceState(active, partial, incomplete string) string {
	switch buildPhase8Proofs().CurrentState {
	case formalcore.Phase8StateActive:
		return active
	case formalcore.Phase8StateSubstantial:
		return partial
	default:
		return incomplete
	}
}

func phase8InsuranceExportItems() []phase8InsuranceExportItem {
	rule := phase8InsurerUsePermissionRule()
	claim := phase8InsurerClaimClass()
	custody := phase8EvidenceCustodyContract()
	lifecycle := phase8CertificationLifecycle()
	return []phase8InsuranceExportItem{
		{
			ExportID:                "insurer_risk_input_export_baseline",
			CurrentState:            "release_ready_after_formal_approval",
			Audience:                rule.Audience,
			ClaimClass:              claim.ClaimClass,
			DisclosureScopeProfile:  append([]string{}, rule.SharingScope...),
			ReleaseLifecycleState:   strings.Join([]string{lifecycle.States[1], lifecycle.States[2], lifecycle.States[3]}, " -> "),
			WithdrawalState:         lifecycle.WithdrawalBehavior,
			ReleaseApprovalRequired: custody.ReleaseApprovalRequired,
			CanCitePublicly:         rule.CanCitePublicly,
			AdverseDecisionExplanation: phase8AdverseDecisionSupport{
				ReasonCode:         "bounded_institutional_risk_input",
				EvidenceBasis:      []string{"/v1/formal/phase8/institutional/risk-quantification", "/v1/formal/phase8/governance/authority-routing"},
				SufficiencyClass:   claim.StandardOfProof,
				ChallengePath:      "/v1/formal/phase8/governance/authority-routing",
				NextReviewPath:     "/v1/formal/phase8/governance/model-risk",
				DisclosureBoundary: "insurer_scoped_release_only",
			},
			RouteRefs: []string{
				"/v1/formal/phase8/institutional/risk-quantification",
				"/v1/formal/phase8/governance/authority-routing",
			},
			Limitations: []string{
				"Export remains insurer-scoped and cannot be publicly cited or widened into regulator or certification release by default.",
			},
		},
	}
}

func phase8IncidentAttributionItems() []phase8IncidentAttributionItem {
	return []phase8IncidentAttributionItem{
		{
			RecordID:                 "incident_custody_correlation_baseline",
			CurrentState:             "support_ready",
			AttributionBasisClass:    "custody_timeline_and_control_correlation",
			EvidenceSufficiencyClass: "assessor_review_ready",
			UnresolvedAmbiguityNote:  "Shared-control overlap remains visible where evidence cannot safely collapse to a single actor or legal cause.",
			NonLegalConclusionMarker: "support_only_not_legal_conclusion",
			ReasonCodes:              []string{"custody_chain_visible", "shared_control_ambiguity_preserved"},
			RouteRefs: []string{
				"/v1/formal/phase8/contracts",
				"/v1/formal/phase8/institutional/insurance-exports",
			},
			Limitations: []string{
				"Attribution support cannot replace external legal or claims adjudication.",
			},
		},
		{
			RecordID:                 "incident_configuration_and_exposure_linkage",
			CurrentState:             "support_ready",
			AttributionBasisClass:    "configuration_exposure_and_review_linkage",
			EvidenceSufficiencyClass: "formal_internal_reviewed",
			UnresolvedAmbiguityNote:  "Indirect dependency and operator timing ambiguity remains visible when root-cause certainty is insufficient.",
			NonLegalConclusionMarker: "support_only_not_legal_conclusion",
			ReasonCodes:              []string{"configuration_linkage_visible", "uncertainty_not_hidden"},
			RouteRefs: []string{
				"/v1/formal/phase8/institutional/risk-quantification",
				"/v1/formal/phase8/governance/model-risk",
			},
			Limitations: []string{
				"Support record remains bounded and cannot be used as a definitive liability conclusion.",
			},
		},
	}
}

func phase8ActuarialBenchmarkItems() []phase8ActuarialBenchmarkItem {
	return []phase8ActuarialBenchmarkItem{
		{
			BenchmarkID:                  "bounded_control_maturity_cohort",
			CurrentState:                 "aggregate_benchmark_ready",
			MinimumCohortSize:            50,
			AggregationScope:             "aggregate_only_cross_tenant_safe_band",
			ReidentificationThreshold:    "strict_low_reidentification_risk_only",
			PublicationWithdrawalTrigger: "withdraw_if_cohort_drops_below_floor_or_reidentification_risk_increases",
			ForbiddenUse: []string{
				"tenant_level_pricing",
				"public_ranked_league_table",
				"raw_subject_disclosure",
			},
			RouteRefs: []string{
				"/v1/formal/phase8/institutional/risk-quantification",
				"/v1/formal/phase8/institutional/insurance-exports",
			},
			Limitations: []string{
				"Benchmark publication remains aggregate-only and cannot be inverted into tenant-identifiable or public marketing claims.",
			},
		},
	}
}

func phase8InsurerUsePermissionRule() formalcore.UsePermissionRule {
	for _, rule := range formalcore.UsePermissionRules() {
		if rule.Audience == formalcore.AudienceInsurer {
			return rule
		}
	}
	return formalcore.UsePermissionRule{}
}

func phase8InsurerClaimClass() formalcore.ClaimClass {
	for _, claim := range formalcore.ClaimClasses() {
		if claim.ClaimClass == formalcore.ClaimClassInsurerFacingRiskInput {
			return claim
		}
	}
	return formalcore.ClaimClass{}
}

func phase8EvidenceCustodyContract() formalcore.EvidenceCustodyContract {
	contracts := formalcore.EvidenceCustodyContracts()
	if len(contracts) == 0 {
		return formalcore.EvidenceCustodyContract{}
	}
	return contracts[0]
}
