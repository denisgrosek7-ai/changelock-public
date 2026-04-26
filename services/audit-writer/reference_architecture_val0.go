package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	referenceArchitectureVal0BlueprintSchema     = "point6.reference_architecture.val0.blueprint_discipline.v1"
	referenceArchitectureVal0EnvironmentSchema   = "point6.reference_architecture.val0.environment_fit.v1"
	referenceArchitectureVal0EvidenceSchema      = "point6.reference_architecture.val0.conformance_evidence.v1"
	referenceArchitectureVal0CompatibilitySchema = "point6.reference_architecture.val0.compatibility_baseline.v1"
	referenceArchitectureVal0ProofsSchema        = "point6.reference_architecture.val0.proofs.v1"
)

type referenceArchitectureVal0BlueprintResponse struct {
	SchemaVersion string                                             `json:"schema_version"`
	GeneratedAt   time.Time                                          `json:"generated_at"`
	CurrentState  string                                             `json:"current_state"`
	Model         operability.ReferenceArchitectureBlueprintContract `json:"model"`
	RouteRefs     []string                                           `json:"route_refs,omitempty"`
	Limitations   []string                                           `json:"limitations,omitempty"`
}

type referenceArchitectureVal0ProofsResponse struct {
	SchemaVersion              string    `json:"schema_version"`
	GeneratedAt                time.Time `json:"generated_at"`
	CurrentState               string    `json:"current_state"`
	Point5DependencyState      string    `json:"point_5_dependency_state"`
	Point5State                string    `json:"point_5_state"`
	Val0State                  string    `json:"val_0_state"`
	Point6State                string    `json:"point_6_state"`
	BlueprintDisciplineState   string    `json:"blueprint_discipline_state"`
	TaxonomyState              string    `json:"blueprint_family_taxonomy_state"`
	EnvironmentFitState        string    `json:"environment_fit_state"`
	EvidenceDisciplineState    string    `json:"conformance_evidence_discipline_state"`
	CompatibilityBaselineState string    `json:"compatibility_deprecation_baseline_state"`
	ConformanceState           string    `json:"reference_conformance_state"`
	SupportedFamilies          []string  `json:"supported_blueprint_families,omitempty"`
	SupportedConformanceStates []string  `json:"supported_conformance_states,omitempty"`
	SupportedCompatibility     []string  `json:"supported_compatibility_states,omitempty"`
	SupportedLifecycle         []string  `json:"supported_lifecycle_states,omitempty"`
	WhyPoint6NotPass           []string  `json:"why_point_6_not_pass,omitempty"`
	SurfaceRefs                []string  `json:"surface_refs,omitempty"`
	EvidenceRefs               []string  `json:"evidence_refs,omitempty"`
	Limitations                []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer       string    `json:"projection_disclaimer"`
	IntegrationSummary         []string  `json:"integration_summary,omitempty"`
}

func referenceArchitectureVal0AllSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/blueprint-discipline",
		"/v1/reference-architecture/val0/environment-fit",
		"/v1/reference-architecture/val0/conformance-evidence",
		"/v1/reference-architecture/val0/compatibility-baseline",
		"/v1/reference-architecture/val0/proofs",
	}
}

func referenceArchitectureVal0EvidenceRefs(model operability.ReferenceArchitectureBlueprintContract) []string {
	refs := []string{"point5_integrated_closure", "evidence_spine"}
	for _, evidence := range model.EvidenceRefs {
		if evidence.EvidenceID == "" {
			continue
		}
		refs = append(refs, evidence.EvidenceID)
	}
	return refs
}

func referenceArchitectureVal0ProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_reference_architecture_foundation"
}

func (s server) referenceArchitectureVal0BlueprintHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureVal0BlueprintDiscipline())
}

func (s server) referenceArchitectureVal0EnvironmentFitHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureVal0EnvironmentFit())
}

func (s server) referenceArchitectureVal0EvidenceHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureVal0EvidenceDiscipline())
}

func (s server) referenceArchitectureVal0CompatibilityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureVal0CompatibilityBaseline())
}

func (s server) referenceArchitectureVal0ProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureVal0Proofs())
}

func buildReferenceArchitectureVal0BlueprintDiscipline() referenceArchitectureVal0BlueprintResponse {
	model := operability.ReferenceArchitectureVal0BlueprintContract()
	return referenceArchitectureVal0BlueprintResponse{
		SchemaVersion: referenceArchitectureVal0BlueprintSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureVal0BlueprintDisciplineState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/reference-architecture/val0/environment-fit",
			"/v1/reference-architecture/val0/proofs",
		},
		Limitations: []string{
			"Val 0 defines only blueprint contract discipline and bounded reference semantics.",
			"Reference architecture remains a projection over canonical evidence and is not certification or mutation authority.",
		},
	}
}

func buildReferenceArchitectureVal0EnvironmentFit() referenceArchitectureVal0BlueprintResponse {
	model := operability.ReferenceArchitectureVal0BlueprintContract()
	return referenceArchitectureVal0BlueprintResponse{
		SchemaVersion: referenceArchitectureVal0EnvironmentSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureVal0EnvironmentFitState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/reference-architecture/val0/conformance-evidence",
			"/v1/reference-architecture/val0/proofs",
		},
		Limitations: []string{
			"Environment fit remains bounded to declared topology, trust anchor, audit path, connectivity, residency, control, and access modes.",
			"Missing capabilities or support boundary mismatches degrade or block matched reference alignment.",
		},
	}
}

func buildReferenceArchitectureVal0EvidenceDiscipline() referenceArchitectureVal0BlueprintResponse {
	model := operability.ReferenceArchitectureVal0BlueprintContract()
	return referenceArchitectureVal0BlueprintResponse{
		SchemaVersion: referenceArchitectureVal0EvidenceSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureVal0EvidenceDisciplineState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/reference-architecture/val0/compatibility-baseline",
			"/v1/reference-architecture/val0/proofs",
		},
		Limitations: []string{
			"Conformance claims remain bounded by explicit evidence scope, RFC3339-valid timestamps, freshness, and caveats.",
			"Stale or malformed evidence cannot support matched reference alignment.",
		},
	}
}

func buildReferenceArchitectureVal0CompatibilityBaseline() referenceArchitectureVal0BlueprintResponse {
	model := operability.ReferenceArchitectureVal0BlueprintContract()
	return referenceArchitectureVal0BlueprintResponse{
		SchemaVersion: referenceArchitectureVal0CompatibilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureVal0CompatibilityBaselineState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/reference-architecture/val0/blueprint-discipline",
			"/v1/reference-architecture/val0/proofs",
		},
		Limitations: []string{
			"Deprecated or superseded blueprints remain visible and cannot silently pass as clean matched references.",
			"Compatibility warnings do not imply universal deployment safety or certification.",
		},
	}
}

func buildReferenceArchitectureVal0Proofs() referenceArchitectureVal0ProofsResponse {
	model := operability.ReferenceArchitectureVal0BlueprintContract()
	point5 := buildIntelligenceCalibrationValEProofs()

	blueprintDisciplineState := operability.EvaluateReferenceArchitectureVal0BlueprintDisciplineState(model)
	taxonomyState := operability.EvaluateReferenceArchitectureVal0TaxonomyState(model)
	environmentFitState := operability.EvaluateReferenceArchitectureVal0EnvironmentFitState(model)
	evidenceDisciplineState := operability.EvaluateReferenceArchitectureVal0EvidenceDisciplineState(model)
	compatibilityBaselineState := operability.EvaluateReferenceArchitectureVal0CompatibilityBaselineState(model)
	conformanceState := operability.EvaluateReferenceArchitectureVal0ReferenceConformanceState(model)

	val0State := operability.EvaluateReferenceArchitectureVal0State(
		point5.Point5State,
		blueprintDisciplineState,
		taxonomyState,
		environmentFitState,
		evidenceDisciplineState,
		compatibilityBaselineState,
		conformanceState,
	)

	surfaceRefs := referenceArchitectureVal0AllSurfaceRefs()
	evidenceRefs := referenceArchitectureVal0EvidenceRefs(model)
	limitations := []string{
		"Val 0 defines blueprint discipline, taxonomy, environment fit, conformance evidence, and compatibility baseline only.",
		"Točka 6 remains not complete because blueprint families, blueprint-as-code, resilience hardening, final reference gate, and integrated closure remain for Val A through Val E.",
		"Reference architecture remains a bounded projection over canonical evidence and must not become certification, policy authority, or mutation authority.",
	}
	whyPoint6NotPass := []string{
		"Val 0 does not ship real blueprint families beyond taxonomy and contract placeholders.",
		"Val A through Val E remain required before point_6_pass can exist.",
		"Točka 6 final PASS is reserved for Val E integrated closure and is impossible in Val 0.",
	}

	point6State := operability.ReferenceArchitecturePoint6StateNotComplete
	currentState := operability.EvaluateReferenceArchitectureVal0ProofsState(
		point5.Point5State,
		val0State,
		blueprintDisciplineState,
		taxonomyState,
		environmentFitState,
		evidenceDisciplineState,
		compatibilityBaselineState,
		conformanceState,
		point6State,
		model.SupportedFamilies,
		model.SupportedConformanceStates,
		model.SupportedCompatibilityStates,
		surfaceRefs,
		evidenceRefs,
		limitations,
		referenceArchitectureVal0ProjectionDisclaimer(),
	)

	return referenceArchitectureVal0ProofsResponse{
		SchemaVersion:              referenceArchitectureVal0ProofsSchema,
		GeneratedAt:                publicSampleTime(),
		CurrentState:               currentState,
		Point5DependencyState:      point5.CurrentState,
		Point5State:                point5.Point5State,
		Val0State:                  val0State,
		Point6State:                point6State,
		BlueprintDisciplineState:   blueprintDisciplineState,
		TaxonomyState:              taxonomyState,
		EnvironmentFitState:        environmentFitState,
		EvidenceDisciplineState:    evidenceDisciplineState,
		CompatibilityBaselineState: compatibilityBaselineState,
		ConformanceState:           conformanceState,
		SupportedFamilies:          model.SupportedFamilies,
		SupportedConformanceStates: model.SupportedConformanceStates,
		SupportedCompatibility:     model.SupportedCompatibilityStates,
		SupportedLifecycle:         model.SupportedLifecycleStates,
		WhyPoint6NotPass:           whyPoint6NotPass,
		SurfaceRefs:                surfaceRefs,
		EvidenceRefs:               evidenceRefs,
		Limitations:                limitations,
		ProjectionDisclaimer:       referenceArchitectureVal0ProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val 0 establishes bounded validated reference blueprint discipline with explicit family taxonomy, environment fit rules, conformance evidence rules, and compatibility baseline handling.",
			"Matched reference alignment requires fresh scoped evidence, supported enums, compatible lifecycle, and full required capability coverage.",
			"Točka 6 remains not complete until later vals add blueprint families, blueprint-as-code, resilience hardening, final reference gate, and integrated closure.",
		},
	}
}
