package main

import (
	"net/http"
	"strings"
	"time"

	formalcore "github.com/denisgrosek/changelock/internal/formal"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	phase8PolicyProfilesSchema                 = "8a.formal_compliance_policy_profiles.v1"
	phase8RegulatoryMappingsSchema             = "8a.formal_compliance_regulatory_mappings.v1"
	phase8VerifierSurfacesSchema               = "8a.formal_compliance_verifier_surfaces.v1"
	phase8CertificationWorkflowSchema          = "8a.formal_compliance_certification_workflow.v1"
	phase8EvidenceAutomationSchema             = "8a.formal_compliance_evidence_automation.v1"
	phase8PolicyProfilesStateActive            = "phase8a_policy_profiles_active"
	phase8PolicyProfilesStatePartial           = "phase8a_policy_profiles_partial"
	phase8PolicyProfilesStateIncomplete        = "phase8a_policy_profiles_incomplete"
	phase8RegulatoryMappingsStateActive        = "phase8a_regulatory_mappings_active"
	phase8RegulatoryMappingsStatePartial       = "phase8a_regulatory_mappings_partial"
	phase8RegulatoryMappingsStateIncomplete    = "phase8a_regulatory_mappings_incomplete"
	phase8VerifierSurfacesStateActive          = "phase8a_verifier_surfaces_active"
	phase8VerifierSurfacesStatePartial         = "phase8a_verifier_surfaces_partial"
	phase8VerifierSurfacesStateIncomplete      = "phase8a_verifier_surfaces_incomplete"
	phase8CertificationWorkflowStateActive     = "phase8a_certification_workflow_active"
	phase8CertificationWorkflowStatePartial    = "phase8a_certification_workflow_partial"
	phase8CertificationWorkflowStateIncomplete = "phase8a_certification_workflow_incomplete"
	phase8EvidenceAutomationStateActive        = "phase8a_evidence_automation_active"
	phase8EvidenceAutomationStatePartial       = "phase8a_evidence_automation_partial"
	phase8EvidenceAutomationStateIncomplete    = "phase8a_evidence_automation_incomplete"
)

type phase8PolicyProfilesResponse struct {
	SchemaVersion        string                           `json:"schema_version"`
	GeneratedAt          time.Time                        `json:"generated_at"`
	CurrentState         string                           `json:"current_state"`
	ProfileState         string                           `json:"profile_state"`
	ManualInterpretation string                           `json:"manual_interpretation_state"`
	JurisdictionState    string                           `json:"jurisdiction_state"`
	Profiles             []formalcore.PolicyAsLawProfile  `json:"profiles,omitempty"`
	JurisdictionProfiles []formalcore.JurisdictionProfile `json:"jurisdiction_profiles,omitempty"`
	ActivationBoundaries []string                         `json:"activation_boundaries,omitempty"`
	RouteRefs            []string                         `json:"route_refs,omitempty"`
	Limitations          []string                         `json:"limitations,omitempty"`
}

type phase8RegulatoryMappingsResponse struct {
	SchemaVersion            string                         `json:"schema_version"`
	GeneratedAt              time.Time                      `json:"generated_at"`
	CurrentState             string                         `json:"current_state"`
	MappingState             string                         `json:"mapping_state"`
	ConflictHandlingState    string                         `json:"conflict_handling_state"`
	CompensatingControlState string                         `json:"compensating_control_state"`
	InheritedControlState    string                         `json:"inherited_control_state"`
	Mappings                 []formalcore.RegulatoryMapping `json:"mappings,omitempty"`
	RouteRefs                []string                       `json:"route_refs,omitempty"`
	Limitations              []string                       `json:"limitations,omitempty"`
}

type phase8VerifierSurfacesResponse struct {
	SchemaVersion          string                       `json:"schema_version"`
	GeneratedAt            time.Time                    `json:"generated_at"`
	CurrentState           string                       `json:"current_state"`
	SurfaceState           string                       `json:"surface_state"`
	Surfaces               []formalcore.VerifierSurface `json:"surfaces,omitempty"`
	AllowedAudienceClasses []string                     `json:"allowed_audience_classes,omitempty"`
	ForbiddenAudiences     []string                     `json:"forbidden_audiences,omitempty"`
	RouteRefs              []string                     `json:"route_refs,omitempty"`
	Limitations            []string                     `json:"limitations,omitempty"`
}

type phase8CertificationWorkflowResponse struct {
	SchemaVersion           string                                 `json:"schema_version"`
	GeneratedAt             time.Time                              `json:"generated_at"`
	CurrentState            string                                 `json:"current_state"`
	WorkflowState           string                                 `json:"workflow_state"`
	SnapshotState           string                                 `json:"snapshot_state"`
	AssessorIndependence    string                                 `json:"assessor_independence_state"`
	IssueAgingState         string                                 `json:"issue_aging_state"`
	ReleaseBoundaryState    string                                 `json:"release_boundary_state"`
	EvidencePacks           []formalcore.CertificationEvidencePack `json:"evidence_packs,omitempty"`
	Lifecycle               formalcore.ArtifactLifecycleWorkflow   `json:"lifecycle"`
	ChallengeWorkflow       formalcore.ChallengeWorkflow           `json:"challenge_workflow"`
	RequiredApprovalActions []string                               `json:"required_approval_actions,omitempty"`
	RouteRefs               []string                               `json:"route_refs,omitempty"`
	Limitations             []string                               `json:"limitations,omitempty"`
}

type phase8EvidenceAutomationArtifact struct {
	ArtifactID    string   `json:"artifact_id"`
	CurrentState  string   `json:"current_state"`
	SourceSurface string   `json:"source_surface"`
	Summary       string   `json:"summary"`
	Completeness  string   `json:"completeness"`
	ReasonCodes   []string `json:"reason_codes,omitempty"`
	RouteRefs     []string `json:"route_refs,omitempty"`
	Limitations   []string `json:"limitations,omitempty"`
}

type phase8EvidenceAutomationExclusion struct {
	ItemID        string   `json:"item_id"`
	Reason        string   `json:"reason"`
	DeferredScope bool     `json:"deferred_scope"`
	ReasonCodes   []string `json:"reason_codes,omitempty"`
}

type phase8EvidenceAutomationResponse struct {
	SchemaVersion         string                              `json:"schema_version"`
	GeneratedAt           time.Time                           `json:"generated_at"`
	CurrentState          string                              `json:"current_state"`
	AutomationState       string                              `json:"automation_state"`
	PackCompletenessScore string                              `json:"pack_completeness_score"`
	IncludedArtifacts     []phase8EvidenceAutomationArtifact  `json:"included_artifacts,omitempty"`
	NotInPackReasons      []phase8EvidenceAutomationExclusion `json:"not_in_pack_reasons,omitempty"`
	ChallengeLinked       []string                            `json:"challenge_linked_artifacts,omitempty"`
	RouteRefs             []string                            `json:"route_refs,omitempty"`
	Limitations           []string                            `json:"limitations,omitempty"`
}

func (s server) phase8PolicyProfilesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8PolicyProfiles())
}

func (s server) phase8RegulatoryMappingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8RegulatoryMappings())
}

func (s server) phase8VerifierSurfacesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8VerifierSurfaces())
}

func (s server) phase8CertificationWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8CertificationWorkflow())
}

func (s server) phase8EvidenceAutomationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8EvidenceAutomation())
}

func buildPhase8PolicyProfiles() phase8PolicyProfilesResponse {
	return phase8PolicyProfilesResponse{
		SchemaVersion:        phase8PolicyProfilesSchema,
		GeneratedAt:          publicSampleTime(),
		CurrentState:         phase8ComplianceSliceState(phase8PolicyProfilesStateActive, phase8PolicyProfilesStatePartial, phase8PolicyProfilesStateIncomplete),
		ProfileState:         "policy_as_law_profile_activation_visible",
		ManualInterpretation: "manual_interpretation_sections_visible",
		JurisdictionState:    "jurisdiction_overlays_and_conflicts_visible",
		Profiles:             formalcore.PolicyAsLawProfiles(),
		JurisdictionProfiles: formalcore.JurisdictionProfiles(),
		ActivationBoundaries: []string{
			"policy_as_law_activation_requires_cross_function_review",
			"manual_interpretation_sections_remain_visible",
			"conflicting_jurisdiction_profiles_require_manual_resolution",
		},
		RouteRefs: []string{
			"/v1/formal/phase8/compliance-codification",
			"/v1/formal/phase8/compliance/regulatory-mappings",
			"/v1/formal/phase8/contracts",
		},
		Limitations: []string{
			"Policy-as-law profiles stay bounded to declared profiles, machine-checkable sections, and explicit manual interpretation gaps.",
			"Profile activation visibility does not claim universal machine-enforceable legal interpretation across all jurisdictions or organizational overlays.",
		},
	}
}

func buildPhase8RegulatoryMappings() phase8RegulatoryMappingsResponse {
	return phase8RegulatoryMappingsResponse{
		SchemaVersion:            phase8RegulatoryMappingsSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             phase8ComplianceSliceState(phase8RegulatoryMappingsStateActive, phase8RegulatoryMappingsStatePartial, phase8RegulatoryMappingsStateIncomplete),
		MappingState:             "regulatory_mapping_pack_active",
		ConflictHandlingState:    "control_conflict_handling_visible",
		CompensatingControlState: "compensating_control_semantics_visible",
		InheritedControlState:    "inherited_control_semantics_visible",
		Mappings:                 formalcore.RegulatoryMappings(),
		RouteRefs: []string{
			"/v1/formal/phase8/compliance/policy-profiles",
			"/v1/formal/phase8/compliance-codification",
			"/v1/formal/phase8/contracts",
		},
		Limitations: []string{
			"Regulatory mappings remain formalized control-alignment aids and do not become self-sufficient regulatory determinations.",
			"Conflict markers and compensating-control semantics stay visible rather than being flattened into optimistic compliance claims.",
		},
	}
}

func buildPhase8VerifierSurfaces() phase8VerifierSurfacesResponse {
	return phase8VerifierSurfacesResponse{
		SchemaVersion: phase8VerifierSurfacesSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  phase8ComplianceSliceState(phase8VerifierSurfacesStateActive, phase8VerifierSurfacesStatePartial, phase8VerifierSurfacesStateIncomplete),
		SurfaceState:  "regulator_and_certifier_safe_verifier_surfaces_active",
		Surfaces:      formalcore.VerifierSurfaces(),
		AllowedAudienceClasses: []string{
			formalcore.AudienceRegulator,
			formalcore.AudienceCertification,
		},
		ForbiddenAudiences: []string{
			"public",
			formalcore.AudienceInsurer,
			"unbounded_cross_jurisdiction_reuse",
		},
		RouteRefs: []string{
			"/v1/formal/phase8/compliance-codification",
			"/v1/formal/phase8/compliance/certification-workflow",
			"/v1/formal/phase8/proofs",
		},
		Limitations: []string{
			"Verifier surfaces remain audience-narrowed and cannot be widened into public certification, public marketing, or insurer export by default.",
			"Regulator-safe and certifier-safe routes remain disclosure supports rather than external authority issuance channels.",
		},
	}
}

func buildPhase8CertificationWorkflow() phase8CertificationWorkflowResponse {
	return phase8CertificationWorkflowResponse{
		SchemaVersion:           phase8CertificationWorkflowSchema,
		GeneratedAt:             publicSampleTime(),
		CurrentState:            phase8ComplianceSliceState(phase8CertificationWorkflowStateActive, phase8CertificationWorkflowStatePartial, phase8CertificationWorkflowStateIncomplete),
		WorkflowState:           "certification_body_workflow_baseline_active",
		SnapshotState:           "assessment_snapshot_and_evidence_freeze_visible",
		AssessorIndependence:    "assessor_independence_notes_visible",
		IssueAgingState:         "issue_aging_visibility_active",
		ReleaseBoundaryState:    "support_artifact_not_certification_issuance",
		EvidencePacks:           formalcore.CertificationEvidencePacks(),
		Lifecycle:               phase8CertificationLifecycle(),
		ChallengeWorkflow:       phase8FormalChallengeWorkflow(),
		RequiredApprovalActions: phase8CertificationApprovalActions(),
		RouteRefs: []string{
			"/v1/formal/phase8/compliance/verifier-surfaces",
			"/v1/formal/phase8/compliance/evidence-automation",
			"/v1/formal/phase8/proofs",
		},
		Limitations: []string{
			"Certification workflow remains assessor-facing support only and never becomes self-issued certification or hidden release authority.",
			"Challenged or withdrawn artifacts remain traceable and review-routed rather than silently returning to active external use.",
		},
	}
}

func buildPhase8EvidenceAutomation() phase8EvidenceAutomationResponse {
	return phase8EvidenceAutomationResponse{
		SchemaVersion:         phase8EvidenceAutomationSchema,
		GeneratedAt:           publicSampleTime(),
		CurrentState:          phase8ComplianceSliceState(phase8EvidenceAutomationStateActive, phase8EvidenceAutomationStatePartial, phase8EvidenceAutomationStateIncomplete),
		AutomationState:       "compliance_evidence_automation_baseline_active",
		PackCompletenessScore: "bounded_machine_checkable_plus_manual_interpretation_visible",
		IncludedArtifacts:     phase8EvidenceAutomationArtifacts(),
		NotInPackReasons:      phase8EvidenceAutomationExclusions(),
		ChallengeLinked: []string{
			"certification_support_pack_baseline",
			"formal_control_mapping_baseline",
		},
		RouteRefs: []string{
			"/v1/formal/phase8/compliance/policy-profiles",
			"/v1/formal/phase8/compliance/regulatory-mappings",
			"/v1/formal/phase8/compliance/certification-workflow",
		},
		Limitations: []string{
			"Compliance evidence automation remains bounded to support-pack assembly, completeness visibility, and review-linked artifacts.",
			"Automation does not turn evidence assembly into legal conclusions, direct certification issuance, or insurer-facing institutional exports.",
		},
	}
}

func phase8ComplianceSliceState(active, partial, incomplete string) string {
	switch formalcore.EvaluateComplianceCodificationState() {
	case formalcore.ComplianceCodificationStateActive:
		return active
	case formalcore.ComplianceCodificationStatePartial:
		return partial
	default:
		return incomplete
	}
}

func phase8CertificationLifecycle() formalcore.ArtifactLifecycleWorkflow {
	workflows := formalcore.ArtifactLifecycleWorkflows()
	if len(workflows) == 0 {
		return formalcore.ArtifactLifecycleWorkflow{}
	}
	return workflows[0]
}

func phase8FormalChallengeWorkflow() formalcore.ChallengeWorkflow {
	workflows := formalcore.ChallengeWorkflows()
	if len(workflows) == 0 {
		return formalcore.ChallengeWorkflow{}
	}
	return workflows[0]
}

func phase8CertificationApprovalActions() []string {
	controls := formalcore.AuthorityControls()
	if len(controls) == 0 {
		return nil
	}
	out := []string{}
	for _, action := range controls[0].NonDelegableActions {
		if strings.Contains(action, "certification") || strings.Contains(action, "external_disclosure") {
			out = append(out, action)
		}
	}
	return out
}

func phase8EvidenceAutomationArtifacts() []phase8EvidenceAutomationArtifact {
	return []phase8EvidenceAutomationArtifact{
		{
			ArtifactID:    "policy_as_law_profile_snapshot",
			CurrentState:  "automation_ready",
			SourceSurface: "compliance.policy_as_law",
			Summary:       "Policy profile snapshot keeps machine-checkable coverage, manual interpretation sections, and jurisdiction refs visible.",
			Completeness:  "machine_checkable_plus_manual_sections_visible",
			ReasonCodes:   []string{"policy_profile_snapshot", "manual_interpretation_not_hidden"},
			RouteRefs:     []string{"/v1/formal/phase8/compliance/policy-profiles"},
			Limitations: []string{
				"Profile snapshot is bounded to declared profile scope and does not claim universal enforceability.",
			},
		},
		{
			ArtifactID:    "regulatory_control_mapping_pack",
			CurrentState:  "automation_ready",
			SourceSurface: "compliance.regulatory_mapping",
			Summary:       "Mapping pack retains control conflicts, compensating controls, and inherited-control semantics in machine-readable form.",
			Completeness:  "control_conflict_and_compensating_control_visible",
			ReasonCodes:   []string{"formalized_mapping", "conflict_marker_preserved"},
			RouteRefs:     []string{"/v1/formal/phase8/compliance/regulatory-mappings"},
			Limitations: []string{
				"Mappings remain bounded support packs and are not standalone regulatory conclusions.",
			},
		},
		{
			ArtifactID:    "certification_assessment_snapshot",
			CurrentState:  "automation_ready",
			SourceSurface: "compliance.certification_pack",
			Summary:       "Certification snapshot preserves evidence freeze, assessor independence notes, and challenge-linked release boundaries.",
			Completeness:  "assessor_review_ready_snapshot",
			ReasonCodes:   []string{"assessment_snapshot_frozen", "challenge_linked_release_path"},
			RouteRefs:     []string{"/v1/formal/phase8/compliance/certification-workflow"},
			Limitations: []string{
				"Assessment snapshot supports certification workflow but never becomes certification issuance.",
			},
		},
		{
			ArtifactID:    "regulator_safe_disclosure_bundle",
			CurrentState:  "automation_ready",
			SourceSurface: "compliance.verifier_surface",
			Summary:       "Disclosure bundle keeps regulator-safe and certifier-safe narrowing profiles visible for bounded external review.",
			Completeness:  "audience_narrowed_export_only",
			ReasonCodes:   []string{"disclosure_basis_visible", "public_widening_blocked"},
			RouteRefs:     []string{"/v1/formal/phase8/compliance/verifier-surfaces"},
			Limitations: []string{
				"Bundle remains audience-bounded and cannot be reused as public proof or insurer-facing export by default.",
			},
		},
	}
}

func phase8EvidenceAutomationExclusions() []phase8EvidenceAutomationExclusion {
	return []phase8EvidenceAutomationExclusion{
		{
			ItemID:        "insurance_facing_evidence_exports",
			Reason:        "Insurer-facing exports remain deferred institutional expansion rather than part of bounded Phase 8A compliance automation.",
			DeferredScope: true,
			ReasonCodes:   []string{"institutional_expansion_deferred", "no_insurer_export_in_8a"},
		},
		{
			ItemID:        "actuarial_benchmark_discipline",
			Reason:        "Actuarial benchmark discipline remains outside Phase 8A and cannot be implied by compliance evidence pack assembly.",
			DeferredScope: true,
			ReasonCodes:   []string{"deferred_benchmarking", "not_compliance_codification_core"},
		},
		{
			ItemID:        "direct_certification_issuance",
			Reason:        "Certification issuance remains external-assessor dependent and is not created by automated support-pack assembly.",
			DeferredScope: false,
			ReasonCodes:   []string{"support_artifact_not_certification", "no_self_issued_authority"},
		},
		{
			ItemID:        "legal_verdict_output",
			Reason:        "Compliance automation remains evidence-traceable support and does not issue direct legal conclusions.",
			DeferredScope: false,
			ReasonCodes:   []string{"no_legal_verdict", "bounded_formal_support_only"},
		},
	}
}
