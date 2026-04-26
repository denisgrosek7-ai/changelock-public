package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	referenceArchitectureValAFamilyRegistrySchema = "point6.reference_architecture.vala.family_registry.v1"
	referenceArchitectureValAFamilyProfilesSchema = "point6.reference_architecture.vala.family_profiles.v1"
	referenceArchitectureValAProofsSchema         = "point6.reference_architecture.vala.proofs.v1"
)

type referenceArchitectureValAFamilyStatus struct {
	Family                      string `json:"family"`
	BlueprintID                 string `json:"blueprint_id"`
	Title                       string `json:"title"`
	CurrentState                string `json:"current_state"`
	LifecycleState              string `json:"lifecycle_state"`
	CompatibilityState          string `json:"compatibility_state"`
	RequiredCapabilitiesCount   int    `json:"required_capabilities_count"`
	OptionalCapabilitiesCount   int    `json:"optional_capabilities_count"`
	HasDegradedConditions       bool   `json:"has_degraded_conditions"`
	HasUnsupportedConditions    bool   `json:"has_unsupported_conditions"`
	HasEvidenceRequirements     bool   `json:"has_evidence_requirements"`
	SupportBoundaryPresent      bool   `json:"support_boundary_present"`
	ProjectionDisclaimerPresent bool   `json:"projection_disclaimer_present"`
}

type referenceArchitectureValAFamilyRegistryResponse struct {
	SchemaVersion string                                                   `json:"schema_version"`
	GeneratedAt   time.Time                                                `json:"generated_at"`
	CurrentState  string                                                   `json:"current_state"`
	Model         operability.ReferenceArchitectureBlueprintFamilyRegistry `json:"model"`
	FamilyStates  []referenceArchitectureValAFamilyStatus                  `json:"family_states,omitempty"`
	RouteRefs     []string                                                 `json:"route_refs,omitempty"`
	Limitations   []string                                                 `json:"limitations,omitempty"`
}

type referenceArchitectureValAFamilyProfilesResponse struct {
	SchemaVersion string                                                    `json:"schema_version"`
	GeneratedAt   time.Time                                                 `json:"generated_at"`
	CurrentState  string                                                    `json:"current_state"`
	Profiles      []operability.ReferenceArchitectureBlueprintFamilyProfile `json:"profiles,omitempty"`
	FamilyStates  []referenceArchitectureValAFamilyStatus                   `json:"family_states,omitempty"`
	RouteRefs     []string                                                  `json:"route_refs,omitempty"`
	Limitations   []string                                                  `json:"limitations,omitempty"`
}

type referenceArchitectureValAProofsResponse struct {
	SchemaVersion         string                                  `json:"schema_version"`
	GeneratedAt           time.Time                               `json:"generated_at"`
	CurrentState          string                                  `json:"current_state"`
	Point5DependencyState string                                  `json:"point_5_dependency_state"`
	Point5State           string                                  `json:"point_5_state"`
	Val0DependencyState   string                                  `json:"val_0_dependency_state"`
	Val0State             string                                  `json:"val_0_state"`
	ValAState             string                                  `json:"val_a_state"`
	Point6State           string                                  `json:"point_6_state"`
	RegistryState         string                                  `json:"family_registry_state"`
	SupportedFamilies     []string                                `json:"supported_blueprint_families,omitempty"`
	FamilyStates          []referenceArchitectureValAFamilyStatus `json:"family_states,omitempty"`
	WhyPoint6NotPass      []string                                `json:"why_point_6_not_pass,omitempty"`
	SurfaceRefs           []string                                `json:"surface_refs,omitempty"`
	EvidenceRefs          []string                                `json:"evidence_refs,omitempty"`
	Limitations           []string                                `json:"limitations,omitempty"`
	ProjectionDisclaimer  string                                  `json:"projection_disclaimer"`
	IntegrationSummary    []string                                `json:"integration_summary,omitempty"`
}

func referenceArchitectureValAAllSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/family-registry",
		"/v1/reference-architecture/vala/family-profiles",
		"/v1/reference-architecture/vala/proofs",
	}
}

func referenceArchitectureValAEvidenceRefs(registry operability.ReferenceArchitectureBlueprintFamilyRegistry) []string {
	refs := []string{"point5_integrated_closure", "point6_val0_proofs", registry.RegistryID}
	for _, profile := range registry.Profiles {
		if profile.BlueprintID != "" {
			refs = append(refs, profile.BlueprintID)
		}
	}
	return refs
}

func referenceArchitectureValAProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_reference_architecture_family_profiles"
}

func buildReferenceArchitectureValAFamilyStatuses(point5State, val0CurrentState, val0State, point6State string, profiles []operability.ReferenceArchitectureBlueprintFamilyProfile) []referenceArchitectureValAFamilyStatus {
	statuses := make([]referenceArchitectureValAFamilyStatus, 0, len(profiles))
	for _, profile := range profiles {
		statuses = append(statuses, referenceArchitectureValAFamilyStatus{
			Family:                      profile.Family,
			BlueprintID:                 profile.BlueprintID,
			Title:                       profile.Title,
			CurrentState:                operability.EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile),
			LifecycleState:              profile.LifecycleState,
			CompatibilityState:          profile.CompatibilityState,
			RequiredCapabilitiesCount:   len(profile.RequiredCapabilities),
			OptionalCapabilitiesCount:   len(profile.OptionalCapabilities),
			HasDegradedConditions:       len(profile.DegradedConditions) > 0,
			HasUnsupportedConditions:    len(profile.UnsupportedConditions) > 0,
			HasEvidenceRequirements:     len(profile.RequiredEvidenceTypes) > 0,
			SupportBoundaryPresent:      profile.SupportBoundaryRef != "",
			ProjectionDisclaimerPresent: profile.ProjectionDisclaimer != "",
		})
	}
	return statuses
}

func (s server) referenceArchitectureValAFamilyRegistryHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValAFamilyRegistry())
}

func (s server) referenceArchitectureValAFamilyProfilesHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValAFamilyProfiles())
}

func (s server) referenceArchitectureValAProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValAProofs())
}

func buildReferenceArchitectureValAFamilyRegistry() referenceArchitectureValAFamilyRegistryResponse {
	val0 := buildReferenceArchitectureVal0Proofs()
	registry := operability.ReferenceArchitectureValAFamilyRegistry()
	familyStates := buildReferenceArchitectureValAFamilyStatuses(val0.Point5State, val0.CurrentState, val0.Val0State, val0.Point6State, registry.Profiles)
	currentState := operability.EvaluateReferenceArchitectureValAFamilyRegistryState(val0.Point5State, val0.CurrentState, val0.Val0State, val0.Point6State, registry)
	return referenceArchitectureValAFamilyRegistryResponse{
		SchemaVersion: referenceArchitectureValAFamilyRegistrySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  currentState,
		Model:         registry,
		FamilyStates:  familyStates,
		RouteRefs: []string{
			"/v1/reference-architecture/val0/proofs",
			"/v1/reference-architecture/vala/family-profiles",
			"/v1/reference-architecture/vala/proofs",
		},
		Limitations: []string{
			"Val A defines bounded core blueprint family profiles only and does not ship deployment recipes or blueprint-as-code packs.",
			"Blueprint family registry remains advisory and does not create certification, approval authority, or canonical truth.",
		},
	}
}

func buildReferenceArchitectureValAFamilyProfiles() referenceArchitectureValAFamilyProfilesResponse {
	val0 := buildReferenceArchitectureVal0Proofs()
	registry := operability.ReferenceArchitectureValAFamilyRegistry()
	currentState := operability.EvaluateReferenceArchitectureValAFamilyRegistryState(val0.Point5State, val0.CurrentState, val0.Val0State, val0.Point6State, registry)
	return referenceArchitectureValAFamilyProfilesResponse{
		SchemaVersion: referenceArchitectureValAFamilyProfilesSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  currentState,
		Profiles:      registry.Profiles,
		FamilyStates:  buildReferenceArchitectureValAFamilyStatuses(val0.Point5State, val0.CurrentState, val0.Val0State, val0.Point6State, registry.Profiles),
		RouteRefs: []string{
			"/v1/reference-architecture/val0/proofs",
			"/v1/reference-architecture/vala/family-registry",
			"/v1/reference-architecture/vala/proofs",
		},
		Limitations: []string{
			"Family profiles are validated reference blueprint profiles with bounded assumptions, capability requirements, evidence requirements, and support boundaries.",
			"Val A remains advisory and does not execute Val B blueprint-as-code, Val C resilience validation, Val D final gate, or Val E integrated closure.",
		},
	}
}

func buildReferenceArchitectureValAProofs() referenceArchitectureValAProofsResponse {
	val0 := buildReferenceArchitectureVal0Proofs()
	registry := operability.ReferenceArchitectureValAFamilyRegistry()
	registryState := operability.EvaluateReferenceArchitectureValAFamilyRegistryState(val0.Point5State, val0.CurrentState, val0.Val0State, val0.Point6State, registry)
	valAState := operability.EvaluateReferenceArchitectureValAState(val0.Point5State, val0.CurrentState, val0.Val0State, val0.Point6State, registryState)
	surfaceRefs := referenceArchitectureValAAllSurfaceRefs()
	evidenceRefs := referenceArchitectureValAEvidenceRefs(registry)
	limitations := []string{
		"Val A implements only core blueprint family profiles and keeps Točka 6 not complete.",
		"Later Val B through Val E remain required for blueprint-as-code, resilience hardening, final reference gate, and integrated closure.",
		"Reference architecture family profiles remain bounded advisory projections over canonical evidence.",
	}
	whyPoint6NotPass := []string{
		"Val A introduces core family profiles only and does not implement blueprint-as-code delivery packs or final closure.",
		"Točka 6 final PASS remains reserved for Val E integrated closure.",
	}
	currentState := operability.EvaluateReferenceArchitectureValAProofsState(
		valAState,
		operability.ReferenceArchitecturePoint6StateNotComplete,
		registry.SupportedFamilies,
		surfaceRefs,
		evidenceRefs,
		limitations,
		referenceArchitectureValAProjectionDisclaimer(),
	)
	return referenceArchitectureValAProofsResponse{
		SchemaVersion:         referenceArchitectureValAProofsSchema,
		GeneratedAt:           publicSampleTime(),
		CurrentState:          currentState,
		Point5DependencyState: val0.Point5DependencyState,
		Point5State:           val0.Point5State,
		Val0DependencyState:   val0.CurrentState,
		Val0State:             val0.Val0State,
		ValAState:             valAState,
		Point6State:           operability.ReferenceArchitecturePoint6StateNotComplete,
		RegistryState:         registryState,
		SupportedFamilies:     registry.SupportedFamilies,
		FamilyStates:          buildReferenceArchitectureValAFamilyStatuses(val0.Point5State, val0.CurrentState, val0.Val0State, val0.Point6State, registry.Profiles),
		WhyPoint6NotPass:      whyPoint6NotPass,
		SurfaceRefs:           surfaceRefs,
		EvidenceRefs:          evidenceRefs,
		Limitations:           limitations,
		ProjectionDisclaimer:  referenceArchitectureValAProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val A turns the Val 0 taxonomy into six bounded validated reference blueprint family profiles.",
			"Every family profile remains anchored to Val 0 contract and conformance discipline and cannot claim certification or point_6_pass.",
			"Točka 6 remains not complete until later vals add blueprint-as-code, resilience hardening, final reference gate, and integrated closure.",
		},
	}
}
