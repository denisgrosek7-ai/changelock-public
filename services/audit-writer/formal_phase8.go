package main

import (
	"net/http"
	"time"

	formalcore "github.com/denisgrosek/changelock/internal/formal"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	phase8EntryGateSchema              = "8.formal_entry_gate.v1"
	phase8ContractsSchema              = "8.formal_contracts.v1"
	phase8FormalDisciplineSchema       = "8.formal_discipline.v1"
	phase8ComplianceCodificationSchema = "8.formal_compliance_codification.v1"
	phase8GovernedAutonomySchema       = "8.formal_governed_autonomy.v1"
	phase8ProofsSchema                 = "8.formal_phase8_proofs.v1"
	phase8CoverageScopeCorePass        = "core_pass"
)

type phase8EntryGateResponse struct {
	SchemaVersion        string    `json:"schema_version"`
	GeneratedAt          time.Time `json:"generated_at"`
	CurrentState         string    `json:"current_state"`
	CarryOverLimitations []string  `json:"carry_over_limitations,omitempty"`
	CarryOverDebt        []string  `json:"carry_over_debt,omitempty"`
	ScopeBoundaries      []string  `json:"scope_boundaries,omitempty"`
	ContractRefs         []string  `json:"contract_refs,omitempty"`
	DeferredScope        []string  `json:"deferred_scope,omitempty"`
	Limitations          []string  `json:"limitations,omitempty"`
}

type phase8ContractsResponse struct {
	SchemaVersion              string                                 `json:"schema_version"`
	GeneratedAt                time.Time                              `json:"generated_at"`
	CurrentState               string                                 `json:"current_state"`
	Coverage                   formalcore.Coverage                    `json:"coverage"`
	ClaimClasses               []formalcore.ClaimClass                `json:"claim_classes,omitempty"`
	UsePermissionRules         []formalcore.UsePermissionRule         `json:"use_permission_rules,omitempty"`
	StandardOfProofClasses     []formalcore.StandardOfProofClass      `json:"standard_of_proof_classes,omitempty"`
	JurisdictionProfiles       []formalcore.JurisdictionProfile       `json:"jurisdiction_profiles,omitempty"`
	ConflictResolutionRules    []formalcore.ConflictResolutionRule    `json:"conflict_resolution_rules,omitempty"`
	EvidenceCustodyContracts   []formalcore.EvidenceCustodyContract   `json:"evidence_custody_contracts,omitempty"`
	ArtifactLifecycleWorkflows []formalcore.ArtifactLifecycleWorkflow `json:"artifact_lifecycle_workflows,omitempty"`
	PolicyAsLawProfiles        []formalcore.PolicyAsLawProfile        `json:"policy_as_law_profiles,omitempty"`
	RegulatoryMappings         []formalcore.RegulatoryMapping         `json:"regulatory_mappings,omitempty"`
	CertificationEvidencePacks []formalcore.CertificationEvidencePack `json:"certification_evidence_packs,omitempty"`
	VerifierSurfaces           []formalcore.VerifierSurface           `json:"verifier_surfaces,omitempty"`
	ChallengeWorkflows         []formalcore.ChallengeWorkflow         `json:"challenge_workflows,omitempty"`
	AuthorityControls          []formalcore.AuthorityControl          `json:"authority_controls,omitempty"`
	AIGuardrails               []formalcore.AIGuardrail               `json:"ai_guardrails,omitempty"`
	ModelRiskContracts         []formalcore.ModelRiskContract         `json:"model_risk_contracts,omitempty"`
	InstitutionalDependencies  []formalcore.InstitutionalDependency   `json:"institutional_dependencies,omitempty"`
	DeferredInstitutionalScope []string                               `json:"deferred_institutional_scope,omitempty"`
	Limitations                []string                               `json:"limitations,omitempty"`
}

type phase8FormalDisciplineResponse struct {
	SchemaVersion              string                                 `json:"schema_version"`
	GeneratedAt                time.Time                              `json:"generated_at"`
	CurrentState               string                                 `json:"current_state"`
	ClaimTaxonomyState         string                                 `json:"claim_taxonomy_state"`
	UsePermissionState         string                                 `json:"use_permission_state"`
	StandardOfProofState       string                                 `json:"standard_of_proof_state"`
	JurisdictionModelState     string                                 `json:"jurisdiction_model_state"`
	ConflictResolutionState    string                                 `json:"conflict_resolution_state"`
	EvidenceCustodyState       string                                 `json:"evidence_custody_state"`
	ArtifactLifecycleState     string                                 `json:"artifact_lifecycle_state"`
	ClaimClasses               []formalcore.ClaimClass                `json:"claim_classes,omitempty"`
	UsePermissionRules         []formalcore.UsePermissionRule         `json:"use_permission_rules,omitempty"`
	StandardOfProofClasses     []formalcore.StandardOfProofClass      `json:"standard_of_proof_classes,omitempty"`
	JurisdictionProfiles       []formalcore.JurisdictionProfile       `json:"jurisdiction_profiles,omitempty"`
	ConflictResolutionRules    []formalcore.ConflictResolutionRule    `json:"conflict_resolution_rules,omitempty"`
	EvidenceCustodyContracts   []formalcore.EvidenceCustodyContract   `json:"evidence_custody_contracts,omitempty"`
	ArtifactLifecycleWorkflows []formalcore.ArtifactLifecycleWorkflow `json:"artifact_lifecycle_workflows,omitempty"`
	Limitations                []string                               `json:"limitations,omitempty"`
}

type phase8ComplianceCodificationResponse struct {
	SchemaVersion              string                                 `json:"schema_version"`
	GeneratedAt                time.Time                              `json:"generated_at"`
	CurrentState               string                                 `json:"current_state"`
	PolicyAsLawState           string                                 `json:"policy_as_law_state"`
	RegulatoryMappingState     string                                 `json:"regulatory_mapping_state"`
	CertificationPackState     string                                 `json:"certification_pack_state"`
	VerifierSurfaceState       string                                 `json:"verifier_surface_state"`
	PolicyAsLawProfiles        []formalcore.PolicyAsLawProfile        `json:"policy_as_law_profiles,omitempty"`
	RegulatoryMappings         []formalcore.RegulatoryMapping         `json:"regulatory_mappings,omitempty"`
	CertificationEvidencePacks []formalcore.CertificationEvidencePack `json:"certification_evidence_packs,omitempty"`
	VerifierSurfaces           []formalcore.VerifierSurface           `json:"verifier_surfaces,omitempty"`
	Limitations                []string                               `json:"limitations,omitempty"`
}

type phase8GovernedAutonomyResponse struct {
	SchemaVersion              string                               `json:"schema_version"`
	GeneratedAt                time.Time                            `json:"generated_at"`
	CurrentState               string                               `json:"current_state"`
	ChallengeWorkflowState     string                               `json:"challenge_workflow_state"`
	AuthorityControlState      string                               `json:"authority_control_state"`
	AIGuardrailState           string                               `json:"ai_guardrail_state"`
	ModelRiskState             string                               `json:"model_risk_state"`
	DependencyRegistryState    string                               `json:"dependency_registry_state"`
	ChallengeWorkflows         []formalcore.ChallengeWorkflow       `json:"challenge_workflows,omitempty"`
	AuthorityControls          []formalcore.AuthorityControl        `json:"authority_controls,omitempty"`
	AIGuardrails               []formalcore.AIGuardrail             `json:"ai_guardrails,omitempty"`
	ModelRiskContracts         []formalcore.ModelRiskContract       `json:"model_risk_contracts,omitempty"`
	InstitutionalDependencies  []formalcore.InstitutionalDependency `json:"institutional_dependencies,omitempty"`
	DeferredInstitutionalScope []string                             `json:"deferred_institutional_scope,omitempty"`
	Limitations                []string                             `json:"limitations,omitempty"`
}

type phase8ProofSection struct {
	CurrentState string   `json:"current_state"`
	KeyRefs      []string `json:"key_refs,omitempty"`
}

type phase8FoundationProofSection struct {
	CurrentState string              `json:"current_state"`
	Coverage     formalcore.Coverage `json:"coverage"`
	KeyRefs      []string            `json:"key_refs,omitempty"`
}

type phase8ProofsResponse struct {
	SchemaVersion              string                       `json:"schema_version"`
	GeneratedAt                time.Time                    `json:"generated_at"`
	CoverageScope              string                       `json:"coverage_scope"`
	CurrentState               string                       `json:"current_state"`
	EntryGate                  phase8ProofSection           `json:"entry_gate"`
	ContractFoundation         phase8FoundationProofSection `json:"contract_foundation"`
	FormalDiscipline           phase8ProofSection           `json:"formal_discipline"`
	ComplianceCodification     phase8ProofSection           `json:"compliance_codification"`
	GovernedAutonomy           phase8ProofSection           `json:"governed_autonomy"`
	DeferredInstitutionalScope []string                     `json:"deferred_institutional_scope,omitempty"`
	Limitations                []string                     `json:"limitations,omitempty"`
}

func (s server) phase8EntryGateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8EntryGate())
}

func (s server) phase8ContractsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8Contracts())
}

func (s server) phase8FormalDisciplineHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8FormalDiscipline())
}

func (s server) phase8ComplianceCodificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8ComplianceCodification())
}

func (s server) phase8GovernedAutonomyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8GovernedAutonomy())
}

func (s server) phase8ProofsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8Proofs())
}

func buildPhase8EntryGate() phase8EntryGateResponse {
	entry := formalcore.EntryGateBaseline()
	return phase8EntryGateResponse{
		SchemaVersion:        phase8EntryGateSchema,
		GeneratedAt:          publicSampleTime(),
		CurrentState:         entry.CurrentState,
		CarryOverLimitations: entry.CarryOverLimitations,
		CarryOverDebt:        entry.CarryOverDebt,
		ScopeBoundaries:      entry.ScopeBoundaries,
		ContractRefs:         entry.ContractRefs,
		DeferredScope:        entry.DeferredScope,
		Limitations: []string{
			"Phase 8 entry gate is a bounded authority-support readiness boundary and does not itself create external institutional authority.",
		},
	}
}

func buildPhase8Contracts() phase8ContractsResponse {
	coverage := formalcore.ContractsCoverage()
	return phase8ContractsResponse{
		SchemaVersion:              phase8ContractsSchema,
		GeneratedAt:                publicSampleTime(),
		CurrentState:               formalcore.EvaluateFoundationState(coverage),
		Coverage:                   coverage,
		ClaimClasses:               formalcore.ClaimClasses(),
		UsePermissionRules:         formalcore.UsePermissionRules(),
		StandardOfProofClasses:     formalcore.StandardOfProofClasses(),
		JurisdictionProfiles:       formalcore.JurisdictionProfiles(),
		ConflictResolutionRules:    formalcore.ConflictResolutionRules(),
		EvidenceCustodyContracts:   formalcore.EvidenceCustodyContracts(),
		ArtifactLifecycleWorkflows: formalcore.ArtifactLifecycleWorkflows(),
		PolicyAsLawProfiles:        formalcore.PolicyAsLawProfiles(),
		RegulatoryMappings:         formalcore.RegulatoryMappings(),
		CertificationEvidencePacks: formalcore.CertificationEvidencePacks(),
		VerifierSurfaces:           formalcore.VerifierSurfaces(),
		ChallengeWorkflows:         formalcore.ChallengeWorkflows(),
		AuthorityControls:          formalcore.AuthorityControls(),
		AIGuardrails:               formalcore.AIGuardrails(),
		ModelRiskContracts:         formalcore.ModelRiskContracts(),
		InstitutionalDependencies:  formalcore.InstitutionalDependencies(),
		DeferredInstitutionalScope: formalcore.DeferredInstitutionalExpansion(),
		Limitations: []string{
			"Phase 8 contracts are bounded governance and formal-use contracts and do not create direct legal or certification authority.",
		},
	}
}

func buildPhase8FormalDiscipline() phase8FormalDisciplineResponse {
	return phase8FormalDisciplineResponse{
		SchemaVersion:              phase8FormalDisciplineSchema,
		GeneratedAt:                publicSampleTime(),
		CurrentState:               formalcore.EvaluateFormalDisciplineState(),
		ClaimTaxonomyState:         "formal_claim_taxonomy_active",
		UsePermissionState:         "formal_use_permission_matrix_active",
		StandardOfProofState:       "formal_standard_of_proof_active",
		JurisdictionModelState:     "formal_jurisdiction_model_active",
		ConflictResolutionState:    "formal_conflict_resolution_active",
		EvidenceCustodyState:       "formal_evidence_custody_active",
		ArtifactLifecycleState:     "formal_artifact_lifecycle_active",
		ClaimClasses:               formalcore.ClaimClasses(),
		UsePermissionRules:         formalcore.UsePermissionRules(),
		StandardOfProofClasses:     formalcore.StandardOfProofClasses(),
		JurisdictionProfiles:       formalcore.JurisdictionProfiles(),
		ConflictResolutionRules:    formalcore.ConflictResolutionRules(),
		EvidenceCustodyContracts:   formalcore.EvidenceCustodyContracts(),
		ArtifactLifecycleWorkflows: formalcore.ArtifactLifecycleWorkflows(),
		Limitations: []string{
			"Formal discipline remains claim- and permission-bounded and does not convert support artifacts into self-issued external authority.",
		},
	}
}

func buildPhase8ComplianceCodification() phase8ComplianceCodificationResponse {
	return phase8ComplianceCodificationResponse{
		SchemaVersion:              phase8ComplianceCodificationSchema,
		GeneratedAt:                publicSampleTime(),
		CurrentState:               formalcore.EvaluateComplianceCodificationState(),
		PolicyAsLawState:           "policy_as_law_profile_active",
		RegulatoryMappingState:     "regulatory_mapping_active",
		CertificationPackState:     "certification_support_pack_active",
		VerifierSurfaceState:       "regulator_and_certifier_safe_surface_active",
		PolicyAsLawProfiles:        formalcore.PolicyAsLawProfiles(),
		RegulatoryMappings:         formalcore.RegulatoryMappings(),
		CertificationEvidencePacks: formalcore.CertificationEvidencePacks(),
		VerifierSurfaces:           formalcore.VerifierSurfaces(),
		Limitations: []string{
			"Compliance codification remains machine-checkable where possible and explicitly marks manual interpretation and unresolved ambiguity where needed.",
			"Certification support artifacts remain bounded support artifacts and do not become self-issued certifications.",
		},
	}
}

func buildPhase8GovernedAutonomy() phase8GovernedAutonomyResponse {
	return phase8GovernedAutonomyResponse{
		SchemaVersion:              phase8GovernedAutonomySchema,
		GeneratedAt:                publicSampleTime(),
		CurrentState:               formalcore.EvaluateGovernedAutonomyState(),
		ChallengeWorkflowState:     "formal_challenge_workflow_active",
		AuthorityControlState:      "non_delegable_authority_controls_active",
		AIGuardrailState:           "formal_ai_guardrails_active",
		ModelRiskState:             "formal_model_risk_management_active",
		DependencyRegistryState:    "institutional_dependency_registry_active",
		ChallengeWorkflows:         formalcore.ChallengeWorkflows(),
		AuthorityControls:          formalcore.AuthorityControls(),
		AIGuardrails:               formalcore.AIGuardrails(),
		ModelRiskContracts:         formalcore.ModelRiskContracts(),
		InstitutionalDependencies:  formalcore.InstitutionalDependencies(),
		DeferredInstitutionalScope: formalcore.DeferredInstitutionalExpansion(),
		Limitations: []string{
			"Governed autonomy remains advisory or review-routed and cannot execute non-delegable authority actions.",
			"Institutional expansion such as risk quantification and insurer integration remains explicitly deferred from the core pass.",
		},
	}
}

func buildPhase8Proofs() phase8ProofsResponse {
	entry := buildPhase8EntryGate()
	contracts := buildPhase8Contracts()
	formal := buildPhase8FormalDiscipline()
	compliance := buildPhase8ComplianceCodification()
	governance := buildPhase8GovernedAutonomy()
	return phase8ProofsResponse{
		SchemaVersion: phase8ProofsSchema,
		GeneratedAt:   publicSampleTime(),
		CoverageScope: phase8CoverageScopeCorePass,
		CurrentState: formalcore.EvaluatePhase8State(
			entry.CurrentState,
			contracts.CurrentState,
			formal.CurrentState,
			compliance.CurrentState,
			governance.CurrentState,
		),
		EntryGate: phase8ProofSection{
			CurrentState: entry.CurrentState,
			KeyRefs:      []string{"/v1/formal/phase8/entry-gate"},
		},
		ContractFoundation: phase8FoundationProofSection{
			CurrentState: contracts.CurrentState,
			Coverage:     contracts.Coverage,
			KeyRefs:      []string{"/v1/formal/phase8/contracts"},
		},
		FormalDiscipline: phase8ProofSection{
			CurrentState: formal.CurrentState,
			KeyRefs:      []string{"/v1/formal/phase8/formal-discipline"},
		},
		ComplianceCodification: phase8ProofSection{
			CurrentState: compliance.CurrentState,
			KeyRefs:      []string{"/v1/formal/phase8/compliance-codification"},
		},
		GovernedAutonomy: phase8ProofSection{
			CurrentState: governance.CurrentState,
			KeyRefs:      []string{"/v1/formal/phase8/governed-autonomy"},
		},
		DeferredInstitutionalScope: formalcore.DeferredInstitutionalExpansion(),
		Limitations: []string{
			"Phase 8 proofs remain a bounded formal-authority core summary and do not claim external legal or certification authority.",
			"Insurance-facing expansion, actuarial benchmarks, and broader federated governance remain outside the initial core-pass proofs gate.",
		},
	}
}
