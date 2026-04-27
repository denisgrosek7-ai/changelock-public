package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	referenceArchitectureValEClosureSchema = "point6.reference_architecture.vale.closure.v1"
	referenceArchitectureValEProofsSchema  = "point6.reference_architecture.vale.proofs.v1"
)

type referenceArchitectureValEClosureResponse struct {
	SchemaVersion string                                             `json:"schema_version"`
	GeneratedAt   time.Time                                          `json:"generated_at"`
	CurrentState  string                                             `json:"current_state"`
	Model         operability.ReferenceArchitectureIntegratedClosure `json:"model"`
	RouteRefs     []string                                           `json:"route_refs,omitempty"`
	Limitations   []string                                           `json:"limitations,omitempty"`
}

type referenceArchitectureValEProofsResponse struct {
	SchemaVersion            string                                              `json:"schema_version"`
	GeneratedAt              time.Time                                           `json:"generated_at"`
	CurrentState             string                                              `json:"current_state"`
	Point5DependencyState    string                                              `json:"point_5_dependency_state"`
	Point5State              string                                              `json:"point_5_state"`
	Val0DependencyState      string                                              `json:"val_0_dependency_state"`
	Val0State                string                                              `json:"val_0_state"`
	ValADependencyState      string                                              `json:"val_a_dependency_state"`
	ValAState                string                                              `json:"val_a_state"`
	ValBDependencyState      string                                              `json:"val_b_dependency_state"`
	ValBState                string                                              `json:"val_b_state"`
	ValCDependencyState      string                                              `json:"val_c_dependency_state"`
	ValCState                string                                              `json:"val_c_state"`
	ValDDependencyState      string                                              `json:"val_d_dependency_state"`
	ValDState                string                                              `json:"val_d_state"`
	ValDFinalGateState       string                                              `json:"val_d_final_gate_state"`
	ClosurePrerequisiteState string                                              `json:"closure_prerequisite_state"`
	ClosureInvariantState    string                                              `json:"closure_invariant_state"`
	ProofSurfaceState        string                                              `json:"proof_surface_state"`
	PassRuleState            string                                              `json:"pass_rule_state"`
	ValEState                string                                              `json:"val_e_state"`
	Point6State              string                                              `json:"point_6_state"`
	Point6PassAllowed        bool                                                `json:"point_6_pass_allowed"`
	Point6PassReason         string                                              `json:"point_6_pass_reason"`
	ClosureInvariants        []operability.ReferenceArchitectureClosureInvariant `json:"closure_invariants,omitempty"`
	BlockingReasons          []string                                            `json:"blocking_reasons,omitempty"`
	Caveats                  []string                                            `json:"caveats,omitempty"`
	Limitations              []string                                            `json:"limitations,omitempty"`
	SurfaceRefs              []string                                            `json:"surface_refs,omitempty"`
	EvidenceRefs             []string                                            `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer     string                                              `json:"projection_disclaimer"`
	IntegrationSummary       []string                                            `json:"integration_summary,omitempty"`
}

func referenceArchitectureValEAllSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/proofs",
		"/v1/reference-architecture/valc/proofs",
		"/v1/reference-architecture/vald/proofs",
		"/v1/reference-architecture/vald/final-gate",
		"/v1/reference-architecture/vale/closure",
		"/v1/reference-architecture/vale/proofs",
	}
}

func referenceArchitectureValEProjectionDisclaimer() string {
	return "projection_only not_canonical_truth integrated_reference_architecture_closure evidence_linked_closure"
}

func referenceArchitectureValEUniqueRefs(groups ...[]string) []string {
	seen := map[string]struct{}{}
	refs := []string{}
	for _, group := range groups {
		for _, ref := range group {
			trimmed := strings.TrimSpace(ref)
			if trimmed == "" {
				continue
			}
			if _, ok := seen[trimmed]; ok {
				continue
			}
			seen[trimmed] = struct{}{}
			refs = append(refs, trimmed)
		}
	}
	return refs
}

func referenceArchitectureValEEvidenceFresh(states ...string) bool {
	supportedActiveStates := []string{
		operability.ReferenceArchitectureVal0StateActive,
		operability.ReferenceArchitectureValAStateActive,
		operability.ReferenceArchitectureValBStateActive,
		operability.ReferenceArchitectureValCStateActive,
		operability.ReferenceArchitectureValDStateActive,
		operability.ReferenceArchitectureValDFinalGateStateActive,
	}
	for _, state := range states {
		matched := false
		for _, supported := range supportedActiveStates {
			if strings.TrimSpace(state) == strings.TrimSpace(supported) {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

func (s server) referenceArchitectureValEClosureHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValEClosure())
}

func (s server) referenceArchitectureValEProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValEProofs())
}

func buildReferenceArchitectureValEClosureModel() operability.ReferenceArchitectureIntegratedClosure {
	val0 := buildReferenceArchitectureVal0Proofs()
	valA := buildReferenceArchitectureValAProofs()
	valB := buildReferenceArchitectureValBProofs()
	valC := buildReferenceArchitectureValCProofs()
	valD := buildReferenceArchitectureValDProofs()

	evidenceRefs := referenceArchitectureValEUniqueRefs(
		val0.EvidenceRefs,
		valA.EvidenceRefs,
		valB.EvidenceRefs,
		valC.EvidenceRefs,
		valD.EvidenceRefs,
	)
	model := operability.ReferenceArchitectureIntegratedClosure{
		ClosureID:        "reference-architecture-point-6-closure",
		Version:          "2026.04",
		Point:            "point_6",
		ClosureVal:       "val_e",
		Point6PassReason: "point_6_pass through Val E only after evidence-linked fail-closed closure across Val 0 through Val D.",
		SourceValStates: operability.ReferenceArchitectureValESourceValStates{
			Val0State: val0.Val0State,
			ValAState: valA.ValAState,
			ValBState: valB.ValBState,
			ValCState: valC.ValCState,
			ValDState: valD.ValDState,
		},
		SourceCurrentStates: operability.ReferenceArchitectureValESourceCurrentStates{
			Val0CurrentState: val0.CurrentState,
			ValACurrentState: valA.CurrentState,
			ValBCurrentState: valB.CurrentState,
			ValCCurrentState: valC.CurrentState,
			ValDCurrentState: valD.CurrentState,
		},
		DependencyStates: operability.ReferenceArchitectureValEDependencyStates{
			Point5State:         val0.Point5State,
			Point5Dependency:    val0.Point5DependencyState,
			Val0DependencyState: val0.CurrentState,
			ValADependencyState: valA.CurrentState,
			ValBDependencyState: valB.CurrentState,
			ValCDependencyState: valC.CurrentState,
			ValDFinalGateState:  valD.FinalGateState,
			PreClosurePoint6:    val0.Point6State,
		},
		Val0: operability.ReferenceArchitectureVal0ProofSnapshot{
			CurrentState:               val0.CurrentState,
			Point5DependencyState:      val0.Point5DependencyState,
			Point5State:                val0.Point5State,
			Val0State:                  val0.Val0State,
			Point6State:                val0.Point6State,
			BlueprintDisciplineState:   val0.BlueprintDisciplineState,
			TaxonomyState:              val0.TaxonomyState,
			EnvironmentFitState:        val0.EnvironmentFitState,
			EvidenceDisciplineState:    val0.EvidenceDisciplineState,
			CompatibilityBaselineState: val0.CompatibilityBaselineState,
			ConformanceState:           val0.ConformanceState,
			SupportedFamilies:          val0.SupportedFamilies,
			SupportedConformanceStates: val0.SupportedConformanceStates,
			SupportedCompatibility:     val0.SupportedCompatibility,
			SupportedLifecycle:         val0.SupportedLifecycle,
			SurfaceRefs:                val0.SurfaceRefs,
			EvidenceRefs:               val0.EvidenceRefs,
			ProjectionDisclaimer:       val0.ProjectionDisclaimer,
		},
		ValA: operability.ReferenceArchitectureValAProofSnapshot{
			CurrentState:          valA.CurrentState,
			Point5DependencyState: valA.Point5DependencyState,
			Point5State:           valA.Point5State,
			Val0DependencyState:   valA.Val0DependencyState,
			Val0State:             valA.Val0State,
			ValAState:             valA.ValAState,
			Point6State:           valA.Point6State,
			RegistryState:         valA.RegistryState,
			SupportedFamilies:     valA.SupportedFamilies,
			SurfaceRefs:           valA.SurfaceRefs,
			EvidenceRefs:          valA.EvidenceRefs,
			ProjectionDisclaimer:  valA.ProjectionDisclaimer,
		},
		ValB: operability.ReferenceArchitectureValBProofSnapshot{
			CurrentState:          valB.CurrentState,
			Point5DependencyState: valB.Point5DependencyState,
			Point5State:           valB.Point5State,
			Val0DependencyState:   valB.Val0DependencyState,
			Val0State:             valB.Val0State,
			ValADependencyState:   valB.ValADependencyState,
			ValAState:             valB.ValAState,
			ValBState:             valB.ValBState,
			Point6State:           valB.Point6State,
			PackRegistryState:     valB.PackRegistryState,
			ArtifactManifestState: valB.ArtifactManifestState,
			BundleState:           valB.BundleState,
			ReadinessState:        valB.ReadinessState,
			ValidationHookState:   valB.ValidationHookState,
			ConformanceKitState:   valB.ConformanceKitState,
			DeviationState:        valB.DeviationState,
			SupportedFamilies:     valB.SupportedFamilies,
			SurfaceRefs:           valB.SurfaceRefs,
			EvidenceRefs:          valB.EvidenceRefs,
			ProjectionDisclaimer:  valB.ProjectionDisclaimer,
		},
		ValC: operability.ReferenceArchitectureValCProofSnapshot{
			CurrentState:             valC.CurrentState,
			Point5DependencyState:    valC.Point5DependencyState,
			Point5State:              valC.Point5State,
			Val0DependencyState:      valC.Val0DependencyState,
			Val0State:                valC.Val0State,
			ValADependencyState:      valC.ValADependencyState,
			ValAState:                valC.ValAState,
			ValBDependencyState:      valC.ValBDependencyState,
			ValBState:                valC.ValBState,
			ValCState:                valC.ValCState,
			Point6State:              valC.Point6State,
			ScenarioPackState:        valC.ScenarioPackState,
			FailureTaxonomyState:     valC.FailureTaxonomyState,
			ScenarioDescriptorState:  valC.ScenarioDescriptorState,
			DegradedModeState:        valC.DegradedModeState,
			RecoveryExpectationState: valC.RecoveryExpectationState,
			ScalingScenarioState:     valC.ScalingScenarioState,
			TrustPathState:           valC.TrustPathState,
			AuditPathState:           valC.AuditPathState,
			ControlPlaneSafetyState:  valC.ControlPlaneSafetyState,
			SupportedFamilies:        valC.SupportedFamilies,
			SurfaceRefs:              valC.SurfaceRefs,
			EvidenceRefs:             valC.EvidenceRefs,
			ProjectionDisclaimer:     valC.ProjectionDisclaimer,
		},
		ValD: operability.ReferenceArchitectureValDProofSnapshot{
			CurrentState:               valD.CurrentState,
			Point5DependencyState:      valD.Point5DependencyState,
			Point5State:                valD.Point5State,
			Val0DependencyState:        valD.Val0DependencyState,
			Val0State:                  valD.Val0State,
			ValADependencyState:        valD.ValADependencyState,
			ValAState:                  valD.ValAState,
			ValBDependencyState:        valD.ValBDependencyState,
			ValBState:                  valD.ValBState,
			ValCDependencyState:        valD.ValCDependencyState,
			ValCState:                  valD.ValCState,
			ValDState:                  valD.ValDState,
			Point6State:                valD.Point6State,
			OperationalVisibilityState: valD.OperationalVisibilityState,
			AlignmentSummaryState:      valD.AlignmentSummaryState,
			DeviationAlertState:        valD.DeviationAlertState,
			SupportBoundaryState:       valD.SupportBoundaryState,
			MigrationUpgradeState:      valD.MigrationUpgradeState,
			TopologyGateState:          valD.TopologyGateState,
			SecurityBoundaryGateState:  valD.SecurityBoundaryGateState,
			OperabilityGateState:       valD.OperabilityGateState,
			CompatibilityGateState:     valD.CompatibilityGateState,
			FinalGateState:             valD.FinalGateState,
			SupportedFamilies:          valD.SupportedFamilies,
			SurfaceRefs:                valD.SurfaceRefs,
			EvidenceRefs:               valD.EvidenceRefs,
			ProjectionDisclaimer:       valD.ProjectionDisclaimer,
		},
		ProofSurfaceRefs: referenceArchitectureValEAllSurfaceRefs(),
		EvidenceRefs:     evidenceRefs,
		Caveats: []string{
			"Integrated reference architecture closure remains a measured advisory projection over the canonical execution, audit, and evidence spine.",
		},
		Limitations: []string{
			"Val E closes Točka 6 only and does not start Točka 7 or create deployment approval authority.",
			"Integrated closure remains fail-closed and blocks on stale, partial, degraded, unsupported, drifted, or unknown source states.",
		},
		ProjectionDisclaimer: referenceArchitectureValEProjectionDisclaimer(),
		CreatedAt:            publicSampleTime().Format(time.RFC3339),
		UpdatedAt:            publicSampleTime().Format(time.RFC3339),
		EvidenceFresh: referenceArchitectureValEEvidenceFresh(
			val0.CurrentState,
			valA.CurrentState,
			valB.CurrentState,
			valC.CurrentState,
			valD.CurrentState,
			valD.FinalGateState,
		),
		StaleEvidenceDetected:         false,
		RedactionKeepsFailuresVisible: true,
	}
	model = operability.ComputeReferenceArchitectureValEClosure(model)
	if model.Point6PassAllowed {
		model.Point6PassReason = "point_6_pass through Val E only after actual Val 0 through Val D proof states, exact proof surfaces, and closure invariants all remain active and fresh."
	} else {
		model.Point6PassReason = "point_6_pass remains blocked until all Val 0 through Val D prerequisites, exact proof surfaces, and closure invariants are active and fresh."
	}
	return operability.ComputeReferenceArchitectureValEClosure(model)
}

func buildReferenceArchitectureValEClosure() referenceArchitectureValEClosureResponse {
	model := buildReferenceArchitectureValEClosureModel()
	return referenceArchitectureValEClosureResponse{
		SchemaVersion: referenceArchitectureValEClosureSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     referenceArchitectureValEAllSurfaceRefs(),
		Limitations:   model.Limitations,
	}
}

func buildReferenceArchitectureValEProofs() referenceArchitectureValEProofsResponse {
	model := buildReferenceArchitectureValEClosureModel()
	return referenceArchitectureValEProofsResponse{
		SchemaVersion:            referenceArchitectureValEProofsSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             model.CurrentState,
		Point5DependencyState:    model.DependencyStates.Point5Dependency,
		Point5State:              model.DependencyStates.Point5State,
		Val0DependencyState:      model.DependencyStates.Val0DependencyState,
		Val0State:                model.SourceValStates.Val0State,
		ValADependencyState:      model.DependencyStates.ValADependencyState,
		ValAState:                model.SourceValStates.ValAState,
		ValBDependencyState:      model.DependencyStates.ValBDependencyState,
		ValBState:                model.SourceValStates.ValBState,
		ValCDependencyState:      model.DependencyStates.ValCDependencyState,
		ValCState:                model.SourceValStates.ValCState,
		ValDDependencyState:      model.SourceCurrentStates.ValDCurrentState,
		ValDState:                model.SourceValStates.ValDState,
		ValDFinalGateState:       model.DependencyStates.ValDFinalGateState,
		ClosurePrerequisiteState: model.ClosurePrerequisiteState,
		ClosureInvariantState:    model.ClosureInvariantState,
		ProofSurfaceState:        model.ProofSurfaceState,
		PassRuleState:            model.PassRuleState,
		ValEState:                model.CurrentState,
		Point6State:              model.Point6State,
		Point6PassAllowed:        model.Point6PassAllowed,
		Point6PassReason:         model.Point6PassReason,
		ClosureInvariants:        model.ClosureInvariants,
		BlockingReasons:          model.BlockingReasons,
		Caveats:                  model.Caveats,
		Limitations:              model.Limitations,
		SurfaceRefs:              model.ProofSurfaceRefs,
		EvidenceRefs:             model.EvidenceRefs,
		ProjectionDisclaimer:     model.ProjectionDisclaimer,
		IntegrationSummary: []string{
			"Val E integrates Val 0 through Val D into an evidence-linked integrated reference architecture closure and is the only layer that can return point_6_pass.",
			"Integrated closure remains advisory and projection-only, does not approve deployment, and does not create a new canonical truth layer.",
		},
	}
}
