package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	referenceArchitectureValCScenarioPackSchema = "point6.reference_architecture.valc.scenario_packs.v1"
	referenceArchitectureValCFailureTaxSchema   = "point6.reference_architecture.valc.failure_taxonomy.v1"
	referenceArchitectureValCDescriptorSchema   = "point6.reference_architecture.valc.scenario_descriptors.v1"
	referenceArchitectureValCDegradedSchema     = "point6.reference_architecture.valc.degraded_modes.v1"
	referenceArchitectureValCRecoverySchema     = "point6.reference_architecture.valc.recovery_expectations.v1"
	referenceArchitectureValCScalingSchema      = "point6.reference_architecture.valc.scaling_scenarios.v1"
	referenceArchitectureValCTrustSchema        = "point6.reference_architecture.valc.trust_path.v1"
	referenceArchitectureValCAuditSchema        = "point6.reference_architecture.valc.audit_path.v1"
	referenceArchitectureValCControlSchema      = "point6.reference_architecture.valc.control_plane_safety.v1"
	referenceArchitectureValCProofsSchema       = "point6.reference_architecture.valc.proofs.v1"
)

type referenceArchitectureValCFamilyStatus struct {
	Family                   string `json:"family"`
	ScenarioPackID           string `json:"scenario_pack_id"`
	ScenarioPackState        string `json:"scenario_pack_state"`
	ScenarioDescriptorState  string `json:"scenario_descriptor_state"`
	DegradedModeState        string `json:"degraded_mode_state"`
	RecoveryExpectationState string `json:"recovery_expectation_state"`
	ScalingScenarioState     string `json:"scaling_scenario_state"`
	TrustPathState           string `json:"trust_path_state"`
	AuditPathState           string `json:"audit_path_state"`
	ControlPlaneSafetyState  string `json:"control_plane_safety_state"`
	BlockingScenarioCount    int    `json:"blocking_scenario_count"`
	RecoveryExpectationCount int    `json:"recovery_expectation_count"`
	ScalingScenarioCount     int    `json:"scaling_scenario_count"`
}

type referenceArchitectureValCScenarioPackResponse struct {
	SchemaVersion string                                                          `json:"schema_version"`
	GeneratedAt   time.Time                                                       `json:"generated_at"`
	CurrentState  string                                                          `json:"current_state"`
	Model         operability.ReferenceArchitectureResilienceScenarioPackRegistry `json:"model"`
	FamilyStates  []referenceArchitectureValCFamilyStatus                         `json:"family_states,omitempty"`
	RouteRefs     []string                                                        `json:"route_refs,omitempty"`
	Limitations   []string                                                        `json:"limitations,omitempty"`
}

type referenceArchitectureValCCollectionResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	GeneratedAt   time.Time                               `json:"generated_at"`
	CurrentState  string                                  `json:"current_state"`
	FamilyStates  []referenceArchitectureValCFamilyStatus `json:"family_states,omitempty"`
	RouteRefs     []string                                `json:"route_refs,omitempty"`
	Limitations   []string                                `json:"limitations,omitempty"`
	Model         any                                     `json:"model"`
}

type referenceArchitectureValCProofsResponse struct {
	SchemaVersion            string                                  `json:"schema_version"`
	GeneratedAt              time.Time                               `json:"generated_at"`
	CurrentState             string                                  `json:"current_state"`
	Point5DependencyState    string                                  `json:"point_5_dependency_state"`
	Point5State              string                                  `json:"point_5_state"`
	Val0DependencyState      string                                  `json:"val_0_dependency_state"`
	Val0State                string                                  `json:"val_0_state"`
	ValADependencyState      string                                  `json:"val_a_dependency_state"`
	ValAState                string                                  `json:"val_a_state"`
	ValBDependencyState      string                                  `json:"val_b_dependency_state"`
	ValBState                string                                  `json:"val_b_state"`
	ValCState                string                                  `json:"val_c_state"`
	Point6State              string                                  `json:"point_6_state"`
	ScenarioPackState        string                                  `json:"scenario_pack_state"`
	FailureTaxonomyState     string                                  `json:"failure_mode_taxonomy_state"`
	ScenarioDescriptorState  string                                  `json:"scenario_descriptor_state"`
	DegradedModeState        string                                  `json:"degraded_mode_state"`
	RecoveryExpectationState string                                  `json:"recovery_expectation_state"`
	ScalingScenarioState     string                                  `json:"scaling_scenario_state"`
	TrustPathState           string                                  `json:"trust_path_state"`
	AuditPathState           string                                  `json:"audit_path_state"`
	ControlPlaneSafetyState  string                                  `json:"control_plane_safety_state"`
	SupportedFamilies        []string                                `json:"supported_blueprint_families,omitempty"`
	FamilyStates             []referenceArchitectureValCFamilyStatus `json:"family_states,omitempty"`
	WhyPoint6NotPass         []string                                `json:"why_point_6_not_pass,omitempty"`
	SurfaceRefs              []string                                `json:"surface_refs,omitempty"`
	EvidenceRefs             []string                                `json:"evidence_refs,omitempty"`
	Limitations              []string                                `json:"limitations,omitempty"`
	ProjectionDisclaimer     string                                  `json:"projection_disclaimer"`
	IntegrationSummary       []string                                `json:"integration_summary,omitempty"`
}

func referenceArchitectureValCAllSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/proofs",
		"/v1/reference-architecture/valc/scenario-packs",
		"/v1/reference-architecture/valc/failure-taxonomy",
		"/v1/reference-architecture/valc/scenario-descriptors",
		"/v1/reference-architecture/valc/degraded-modes",
		"/v1/reference-architecture/valc/recovery-expectations",
		"/v1/reference-architecture/valc/scaling-scenarios",
		"/v1/reference-architecture/valc/trust-path",
		"/v1/reference-architecture/valc/audit-path",
		"/v1/reference-architecture/valc/control-plane-safety",
		"/v1/reference-architecture/valc/proofs",
	}
}

func referenceArchitectureValCProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_resilience_scaling_hardening"
}

func referenceArchitectureValCPoint5DependencyHealthy(state string) bool {
	return strings.TrimSpace(state) == operability.IntelligenceCalibrationValEStateActive
}

func referenceArchitectureValCProofCurrentState(
	valCState, point5DependencyState, point6State string,
	supportedFamilies, surfaceRefs, evidenceRefs, limitations []string,
	projectionDisclaimer string,
) string {
	baseState := strings.TrimSpace(valCState)
	if baseState == operability.ReferenceArchitectureValCStateActive &&
		!referenceArchitectureValCPoint5DependencyHealthy(point5DependencyState) {
		baseState = operability.ReferenceArchitectureValCStatePartial
	}
	return operability.EvaluateReferenceArchitectureValCProofsState(
		baseState,
		point6State,
		supportedFamilies,
		surfaceRefs,
		evidenceRefs,
		limitations,
		projectionDisclaimer,
	)
}

func referenceArchitectureValCEvidenceRefs(
	registry operability.ReferenceArchitectureResilienceScenarioPackRegistry,
	descriptors operability.ReferenceArchitectureScenarioDescriptorCollection,
	recovery operability.ReferenceArchitectureRecoveryExpectationCollection,
	trust operability.ReferenceArchitectureTrustPathCollection,
	audit operability.ReferenceArchitectureAuditPathCollection,
) []string {
	refs := []string{
		"point5_integrated_closure",
		"point6_val0_proofs",
		"point6_vala_proofs",
		"point6_valb_proofs",
		registry.RegistryID,
		descriptors.CollectionID,
		recovery.CollectionID,
		trust.CollectionID,
		audit.CollectionID,
	}
	for _, pack := range registry.ScenarioPacks {
		if pack.ScenarioPackID != "" {
			refs = append(refs, pack.ScenarioPackID)
		}
		for _, evidence := range pack.EvidenceRefs {
			if evidence.EvidenceID != "" {
				refs = append(refs, evidence.EvidenceID)
			}
		}
	}
	for _, check := range trust.Checks {
		if check.CheckID != "" {
			refs = append(refs, check.CheckID)
		}
	}
	for _, check := range audit.Checks {
		if check.CheckID != "" {
			refs = append(refs, check.CheckID)
		}
	}
	return refs
}

func referenceArchitectureValCCountBlockingScenarios(pack operability.ReferenceArchitectureScenarioDescriptorPack) int {
	count := 0
	for _, scenario := range pack.Scenarios {
		if scenario.BlocksMatched {
			count++
		}
	}
	return count
}

func buildReferenceArchitectureValCFamilyStatuses(
	registry operability.ReferenceArchitectureResilienceScenarioPackRegistry,
	descriptors operability.ReferenceArchitectureScenarioDescriptorCollection,
	degraded operability.ReferenceArchitectureDegradedModeCollection,
	recovery operability.ReferenceArchitectureRecoveryExpectationCollection,
	scaling operability.ReferenceArchitectureScalingScenarioCollection,
	trust operability.ReferenceArchitectureTrustPathCollection,
	audit operability.ReferenceArchitectureAuditPathCollection,
	control operability.ReferenceArchitectureControlPlaneSafetyCollection,
) []referenceArchitectureValCFamilyStatus {
	descriptorByFamily := map[string]operability.ReferenceArchitectureScenarioDescriptorPack{}
	for _, pack := range descriptors.Packs {
		descriptorByFamily[pack.BlueprintFamily] = pack
	}
	degradedByFamily := map[string]operability.ReferenceArchitectureDegradedModePack{}
	for _, pack := range degraded.Packs {
		degradedByFamily[pack.BlueprintFamily] = pack
	}
	recoveryByFamily := map[string]operability.ReferenceArchitectureRecoveryExpectationPack{}
	for _, pack := range recovery.Packs {
		recoveryByFamily[pack.BlueprintFamily] = pack
	}
	scalingByFamily := map[string]operability.ReferenceArchitectureScalingScenarioPack{}
	for _, pack := range scaling.Packs {
		scalingByFamily[pack.BlueprintFamily] = pack
	}
	trustByFamily := map[string]operability.ReferenceArchitectureTrustPathContinuityCheck{}
	for _, check := range trust.Checks {
		trustByFamily[check.BlueprintFamily] = check
	}
	auditByFamily := map[string]operability.ReferenceArchitectureAuditPathDegradationCheck{}
	for _, check := range audit.Checks {
		auditByFamily[check.BlueprintFamily] = check
	}
	controlByFamily := map[string]operability.ReferenceArchitectureControlPlaneSafetyCheck{}
	for _, check := range control.Checks {
		controlByFamily[check.BlueprintFamily] = check
	}

	statuses := make([]referenceArchitectureValCFamilyStatus, 0, len(registry.ScenarioPacks))
	for _, pack := range registry.ScenarioPacks {
		descriptorPack := descriptorByFamily[pack.BlueprintFamily]
		degradedPack := degradedByFamily[pack.BlueprintFamily]
		recoveryPack := recoveryByFamily[pack.BlueprintFamily]
		scalingPack := scalingByFamily[pack.BlueprintFamily]
		trustCheck := trustByFamily[pack.BlueprintFamily]
		auditCheck := auditByFamily[pack.BlueprintFamily]
		controlCheck := controlByFamily[pack.BlueprintFamily]

		statuses = append(statuses, referenceArchitectureValCFamilyStatus{
			Family:                   pack.BlueprintFamily,
			ScenarioPackID:           pack.ScenarioPackID,
			ScenarioPackState:        operability.EvaluateReferenceArchitectureValCScenarioPackState(pack),
			ScenarioDescriptorState:  operability.EvaluateReferenceArchitectureValCScenarioDescriptorPackState(descriptorPack),
			DegradedModeState:        operability.EvaluateReferenceArchitectureValCDegradedModePackState(degradedPack),
			RecoveryExpectationState: operability.EvaluateReferenceArchitectureValCRecoveryExpectationPackState(recoveryPack),
			ScalingScenarioState:     operability.EvaluateReferenceArchitectureValCScalingScenarioPackState(scalingPack),
			TrustPathState:           operability.EvaluateReferenceArchitectureValCTrustPathCheckState(trustCheck),
			AuditPathState:           operability.EvaluateReferenceArchitectureValCAuditPathCheckState(auditCheck),
			ControlPlaneSafetyState:  operability.EvaluateReferenceArchitectureValCControlPlaneCheckState(controlCheck),
			BlockingScenarioCount:    referenceArchitectureValCCountBlockingScenarios(descriptorPack),
			RecoveryExpectationCount: len(recoveryPack.Expectations),
			ScalingScenarioCount:     len(scalingPack.Scenarios),
		})
	}
	return statuses
}

func (s server) referenceArchitectureValCScenarioPacksHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValCScenarioPacks())
}

func (s server) referenceArchitectureValCFailureTaxonomyHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValCFailureTaxonomy())
}

func (s server) referenceArchitectureValCScenarioDescriptorsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValCScenarioDescriptors())
}

func (s server) referenceArchitectureValCDegradedModesHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValCDegradedModes())
}

func (s server) referenceArchitectureValCRecoveryExpectationsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValCRecoveryExpectations())
}

func (s server) referenceArchitectureValCScalingScenariosHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValCScalingScenarios())
}

func (s server) referenceArchitectureValCTrustPathHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValCTrustPath())
}

func (s server) referenceArchitectureValCAuditPathHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValCAuditPath())
}

func (s server) referenceArchitectureValCControlPlaneSafetyHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValCControlPlaneSafety())
}

func (s server) referenceArchitectureValCProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValCProofs())
}

func buildReferenceArchitectureValCScenarioPacks() referenceArchitectureValCScenarioPackResponse {
	registry := operability.ReferenceArchitectureValCScenarioPackRegistry()
	descriptors := operability.ReferenceArchitectureValCScenarioDescriptorCollection()
	degraded := operability.ReferenceArchitectureValCDegradedModeCollection()
	recovery := operability.ReferenceArchitectureValCRecoveryExpectationCollection()
	scaling := operability.ReferenceArchitectureValCScalingScenarioCollection()
	trust := operability.ReferenceArchitectureValCTrustPathCollection()
	audit := operability.ReferenceArchitectureValCAuditPathCollection()
	control := operability.ReferenceArchitectureValCControlPlaneSafetyCollection()
	return referenceArchitectureValCScenarioPackResponse{
		SchemaVersion: referenceArchitectureValCScenarioPackSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValCScenarioPackRegistryState(registry),
		Model:         registry,
		FamilyStates:  buildReferenceArchitectureValCFamilyStatuses(registry, descriptors, degraded, recovery, scaling, trust, audit, control),
		RouteRefs: []string{
			"/v1/reference-architecture/valb/proofs",
			"/v1/reference-architecture/valc/proofs",
		},
		Limitations: []string{
			"Val C scenario packs define bounded resilience and scaling contracts only and do not execute chaos or recovery.",
			"Scenario packs remain advisory projections and do not approve deployment or mutate canonical truth.",
		},
	}
}

func buildReferenceArchitectureValCFailureTaxonomy() referenceArchitectureValCCollectionResponse {
	taxonomy := operability.ReferenceArchitectureValCFailureModeTaxonomy()
	registry := operability.ReferenceArchitectureValCScenarioPackRegistry()
	descriptors := operability.ReferenceArchitectureValCScenarioDescriptorCollection()
	degraded := operability.ReferenceArchitectureValCDegradedModeCollection()
	recovery := operability.ReferenceArchitectureValCRecoveryExpectationCollection()
	scaling := operability.ReferenceArchitectureValCScalingScenarioCollection()
	trust := operability.ReferenceArchitectureValCTrustPathCollection()
	audit := operability.ReferenceArchitectureValCAuditPathCollection()
	control := operability.ReferenceArchitectureValCControlPlaneSafetyCollection()
	return referenceArchitectureValCCollectionResponse{
		SchemaVersion: referenceArchitectureValCFailureTaxSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValCFailureModeTaxonomyState(taxonomy),
		Model:         taxonomy,
		FamilyStates:  buildReferenceArchitectureValCFamilyStatuses(registry, descriptors, degraded, recovery, scaling, trust, audit, control),
		RouteRefs: []string{
			"/v1/reference-architecture/valc/scenario-packs",
			"/v1/reference-architecture/valc/proofs",
		},
		Limitations: []string{
			"Failure-mode taxonomy is bounded and fail-closed; unknown categories cannot become active.",
			"Taxonomy output remains descriptive and does not authorize remediation or deployment.",
		},
	}
}

func buildReferenceArchitectureValCScenarioDescriptors() referenceArchitectureValCCollectionResponse {
	registry := operability.ReferenceArchitectureValCScenarioPackRegistry()
	descriptors := operability.ReferenceArchitectureValCScenarioDescriptorCollection()
	degraded := operability.ReferenceArchitectureValCDegradedModeCollection()
	recovery := operability.ReferenceArchitectureValCRecoveryExpectationCollection()
	scaling := operability.ReferenceArchitectureValCScalingScenarioCollection()
	trust := operability.ReferenceArchitectureValCTrustPathCollection()
	audit := operability.ReferenceArchitectureValCAuditPathCollection()
	control := operability.ReferenceArchitectureValCControlPlaneSafetyCollection()
	return referenceArchitectureValCCollectionResponse{
		SchemaVersion: referenceArchitectureValCDescriptorSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValCScenarioDescriptorCollectionState(descriptors),
		Model:         descriptors,
		FamilyStates:  buildReferenceArchitectureValCFamilyStatuses(registry, descriptors, degraded, recovery, scaling, trust, audit, control),
		RouteRefs: []string{
			"/v1/reference-architecture/valc/scenario-packs",
			"/v1/reference-architecture/valc/proofs",
		},
		Limitations: []string{
			"Scenario descriptors define deterministic expected outcomes only and do not execute fault injection.",
			"Blocking scenarios remain explicit and advisory_only does not bypass fail-closed semantics.",
		},
	}
}

func buildReferenceArchitectureValCDegradedModes() referenceArchitectureValCCollectionResponse {
	registry := operability.ReferenceArchitectureValCScenarioPackRegistry()
	descriptors := operability.ReferenceArchitectureValCScenarioDescriptorCollection()
	degraded := operability.ReferenceArchitectureValCDegradedModeCollection()
	recovery := operability.ReferenceArchitectureValCRecoveryExpectationCollection()
	scaling := operability.ReferenceArchitectureValCScalingScenarioCollection()
	trust := operability.ReferenceArchitectureValCTrustPathCollection()
	audit := operability.ReferenceArchitectureValCAuditPathCollection()
	control := operability.ReferenceArchitectureValCControlPlaneSafetyCollection()
	return referenceArchitectureValCCollectionResponse{
		SchemaVersion: referenceArchitectureValCDegradedSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValCDegradedModeCollectionState(degraded),
		Model:         degraded,
		FamilyStates:  buildReferenceArchitectureValCFamilyStatuses(registry, descriptors, degraded, recovery, scaling, trust, audit, control),
		RouteRefs: []string{
			"/v1/reference-architecture/valc/scenario-descriptors",
			"/v1/reference-architecture/valc/proofs",
		},
		Limitations: []string{
			"Bounded degraded modes preserve blocked operations, operator action, and evidence retention requirements.",
			"Degraded output cannot silently become matched or suppress canonical failure state.",
		},
	}
}

func buildReferenceArchitectureValCRecoveryExpectations() referenceArchitectureValCCollectionResponse {
	registry := operability.ReferenceArchitectureValCScenarioPackRegistry()
	descriptors := operability.ReferenceArchitectureValCScenarioDescriptorCollection()
	degraded := operability.ReferenceArchitectureValCDegradedModeCollection()
	recovery := operability.ReferenceArchitectureValCRecoveryExpectationCollection()
	scaling := operability.ReferenceArchitectureValCScalingScenarioCollection()
	trust := operability.ReferenceArchitectureValCTrustPathCollection()
	audit := operability.ReferenceArchitectureValCAuditPathCollection()
	control := operability.ReferenceArchitectureValCControlPlaneSafetyCollection()
	return referenceArchitectureValCCollectionResponse{
		SchemaVersion: referenceArchitectureValCRecoverySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValCRecoveryExpectationCollectionState(recovery),
		Model:         recovery,
		FamilyStates:  buildReferenceArchitectureValCFamilyStatuses(registry, descriptors, degraded, recovery, scaling, trust, audit, control),
		RouteRefs: []string{
			"/v1/reference-architecture/valc/degraded-modes",
			"/v1/reference-architecture/valc/proofs",
		},
		Limitations: []string{
			"Recovery expectations are contracts and evidence requirements only; Val C does not execute recovery.",
			"Missing recovery path, stale evidence, or unsupported conditions fail closed.",
		},
	}
}

func buildReferenceArchitectureValCScalingScenarios() referenceArchitectureValCCollectionResponse {
	registry := operability.ReferenceArchitectureValCScenarioPackRegistry()
	descriptors := operability.ReferenceArchitectureValCScenarioDescriptorCollection()
	degraded := operability.ReferenceArchitectureValCDegradedModeCollection()
	recovery := operability.ReferenceArchitectureValCRecoveryExpectationCollection()
	scaling := operability.ReferenceArchitectureValCScalingScenarioCollection()
	trust := operability.ReferenceArchitectureValCTrustPathCollection()
	audit := operability.ReferenceArchitectureValCAuditPathCollection()
	control := operability.ReferenceArchitectureValCControlPlaneSafetyCollection()
	return referenceArchitectureValCCollectionResponse{
		SchemaVersion: referenceArchitectureValCScalingSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValCScalingScenarioCollectionState(scaling),
		Model:         scaling,
		FamilyStates:  buildReferenceArchitectureValCFamilyStatuses(registry, descriptors, degraded, recovery, scaling, trust, audit, control),
		RouteRefs: []string{
			"/v1/reference-architecture/valc/scenario-packs",
			"/v1/reference-architecture/valc/proofs",
		},
		Limitations: []string{
			"Scaling descriptors are bounded assumptions and thresholds only, not real load-test execution or performance guarantees.",
			"Unknown categories, malformed timestamps, or stale evidence fail closed.",
		},
	}
}

func buildReferenceArchitectureValCTrustPath() referenceArchitectureValCCollectionResponse {
	registry := operability.ReferenceArchitectureValCScenarioPackRegistry()
	descriptors := operability.ReferenceArchitectureValCScenarioDescriptorCollection()
	degraded := operability.ReferenceArchitectureValCDegradedModeCollection()
	recovery := operability.ReferenceArchitectureValCRecoveryExpectationCollection()
	scaling := operability.ReferenceArchitectureValCScalingScenarioCollection()
	trust := operability.ReferenceArchitectureValCTrustPathCollection()
	audit := operability.ReferenceArchitectureValCAuditPathCollection()
	control := operability.ReferenceArchitectureValCControlPlaneSafetyCollection()
	return referenceArchitectureValCCollectionResponse{
		SchemaVersion: referenceArchitectureValCTrustSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValCTrustPathCollectionState(trust),
		Model:         trust,
		FamilyStates:  buildReferenceArchitectureValCFamilyStatuses(registry, descriptors, degraded, recovery, scaling, trust, audit, control),
		RouteRefs: []string{
			"/v1/reference-architecture/valc/scenario-packs",
			"/v1/reference-architecture/valc/proofs",
		},
		Limitations: []string{
			"Trust-path continuity remains evidence-linked and advisory; stale or unsupported trust paths cannot be active.",
			"High-assurance and sovereign profiles preserve stricter local trust semantics where required.",
		},
	}
}

func buildReferenceArchitectureValCAuditPath() referenceArchitectureValCCollectionResponse {
	registry := operability.ReferenceArchitectureValCScenarioPackRegistry()
	descriptors := operability.ReferenceArchitectureValCScenarioDescriptorCollection()
	degraded := operability.ReferenceArchitectureValCDegradedModeCollection()
	recovery := operability.ReferenceArchitectureValCRecoveryExpectationCollection()
	scaling := operability.ReferenceArchitectureValCScalingScenarioCollection()
	trust := operability.ReferenceArchitectureValCTrustPathCollection()
	audit := operability.ReferenceArchitectureValCAuditPathCollection()
	control := operability.ReferenceArchitectureValCControlPlaneSafetyCollection()
	return referenceArchitectureValCCollectionResponse{
		SchemaVersion: referenceArchitectureValCAuditSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValCAuditPathCollectionState(audit),
		Model:         audit,
		FamilyStates:  buildReferenceArchitectureValCFamilyStatuses(registry, descriptors, degraded, recovery, scaling, trust, audit, control),
		RouteRefs: []string{
			"/v1/reference-architecture/valc/recovery-expectations",
			"/v1/reference-architecture/valc/proofs",
		},
		Limitations: []string{
			"Audit-path degradation remains bounded and cannot suppress canonical failure state.",
			"Audit writer unavailability or missing evidence custody path cannot silently pass.",
		},
	}
}

func buildReferenceArchitectureValCControlPlaneSafety() referenceArchitectureValCCollectionResponse {
	registry := operability.ReferenceArchitectureValCScenarioPackRegistry()
	descriptors := operability.ReferenceArchitectureValCScenarioDescriptorCollection()
	degraded := operability.ReferenceArchitectureValCDegradedModeCollection()
	recovery := operability.ReferenceArchitectureValCRecoveryExpectationCollection()
	scaling := operability.ReferenceArchitectureValCScalingScenarioCollection()
	trust := operability.ReferenceArchitectureValCTrustPathCollection()
	audit := operability.ReferenceArchitectureValCAuditPathCollection()
	control := operability.ReferenceArchitectureValCControlPlaneSafetyCollection()
	return referenceArchitectureValCCollectionResponse{
		SchemaVersion: referenceArchitectureValCControlSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValCControlPlaneCollectionState(control),
		Model:         control,
		FamilyStates:  buildReferenceArchitectureValCFamilyStatuses(registry, descriptors, degraded, recovery, scaling, trust, audit, control),
		RouteRefs: []string{
			"/v1/reference-architecture/valc/scaling-scenarios",
			"/v1/reference-architecture/valc/proofs",
		},
		Limitations: []string{
			"Control-plane safety checks define bounded overload, timeout, and evidence expectations only.",
			"No automatic approval, mutation, or deployment authorization is introduced in Val C.",
		},
	}
}

func buildReferenceArchitectureValCProofs() referenceArchitectureValCProofsResponse {
	val0 := buildReferenceArchitectureVal0Proofs()
	valA := buildReferenceArchitectureValAProofs()
	valB := buildReferenceArchitectureValBProofs()
	registry := operability.ReferenceArchitectureValCScenarioPackRegistry()
	taxonomy := operability.ReferenceArchitectureValCFailureModeTaxonomy()
	descriptors := operability.ReferenceArchitectureValCScenarioDescriptorCollection()
	degraded := operability.ReferenceArchitectureValCDegradedModeCollection()
	recovery := operability.ReferenceArchitectureValCRecoveryExpectationCollection()
	scaling := operability.ReferenceArchitectureValCScalingScenarioCollection()
	trust := operability.ReferenceArchitectureValCTrustPathCollection()
	audit := operability.ReferenceArchitectureValCAuditPathCollection()
	control := operability.ReferenceArchitectureValCControlPlaneSafetyCollection()

	scenarioPackState := operability.EvaluateReferenceArchitectureValCScenarioPackRegistryState(registry)
	failureTaxonomyState := operability.EvaluateReferenceArchitectureValCFailureModeTaxonomyState(taxonomy)
	scenarioDescriptorState := operability.EvaluateReferenceArchitectureValCScenarioDescriptorCollectionState(descriptors)
	degradedModeState := operability.EvaluateReferenceArchitectureValCDegradedModeCollectionState(degraded)
	recoveryExpectationState := operability.EvaluateReferenceArchitectureValCRecoveryExpectationCollectionState(recovery)
	scalingScenarioState := operability.EvaluateReferenceArchitectureValCScalingScenarioCollectionState(scaling)
	trustPathState := operability.EvaluateReferenceArchitectureValCTrustPathCollectionState(trust)
	auditPathState := operability.EvaluateReferenceArchitectureValCAuditPathCollectionState(audit)
	controlPlaneState := operability.EvaluateReferenceArchitectureValCControlPlaneCollectionState(control)

	valCState := operability.EvaluateReferenceArchitectureValCState(
		val0.Point5State,
		val0.CurrentState,
		val0.Val0State,
		valA.CurrentState,
		valA.ValAState,
		valB.CurrentState,
		valB.ValBState,
		val0.Point6State,
		scenarioPackState,
		failureTaxonomyState,
		scenarioDescriptorState,
		degradedModeState,
		recoveryExpectationState,
		scalingScenarioState,
		trustPathState,
		auditPathState,
		controlPlaneState,
	)

	surfaceRefs := referenceArchitectureValCAllSurfaceRefs()
	evidenceRefs := referenceArchitectureValCEvidenceRefs(registry, descriptors, recovery, trust, audit)
	limitations := []string{
		"Val C defines bounded resilience and scaling contracts only and does not execute chaos, load, or recovery workflows.",
		"Recovery expectations are evidence-linked contracts, not proof of executed disaster recovery.",
		"Scaling descriptors are bounded capacity assumptions and thresholds, not performance guarantees.",
		"Točka 6 remains not_complete until final integrated closure in Val E.",
	}
	effectiveValCState := valCState
	if effectiveValCState == operability.ReferenceArchitectureValCStateActive &&
		!referenceArchitectureValCPoint5DependencyHealthy(val0.Point5DependencyState) {
		effectiveValCState = operability.ReferenceArchitectureValCStatePartial
	}
	currentState := referenceArchitectureValCProofCurrentState(
		effectiveValCState,
		val0.Point5DependencyState,
		val0.Point6State,
		registry.SupportedFamilies,
		surfaceRefs,
		evidenceRefs,
		limitations,
		referenceArchitectureValCProjectionDisclaimer(),
	)

	return referenceArchitectureValCProofsResponse{
		SchemaVersion:            referenceArchitectureValCProofsSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             currentState,
		Point5DependencyState:    val0.Point5DependencyState,
		Point5State:              val0.Point5State,
		Val0DependencyState:      val0.CurrentState,
		Val0State:                val0.Val0State,
		ValADependencyState:      valA.CurrentState,
		ValAState:                valA.ValAState,
		ValBDependencyState:      valB.CurrentState,
		ValBState:                valB.ValBState,
		ValCState:                effectiveValCState,
		Point6State:              operability.ReferenceArchitecturePoint6StateNotComplete,
		ScenarioPackState:        scenarioPackState,
		FailureTaxonomyState:     failureTaxonomyState,
		ScenarioDescriptorState:  scenarioDescriptorState,
		DegradedModeState:        degradedModeState,
		RecoveryExpectationState: recoveryExpectationState,
		ScalingScenarioState:     scalingScenarioState,
		TrustPathState:           trustPathState,
		AuditPathState:           auditPathState,
		ControlPlaneSafetyState:  controlPlaneState,
		SupportedFamilies:        registry.SupportedFamilies,
		FamilyStates:             buildReferenceArchitectureValCFamilyStatuses(registry, descriptors, degraded, recovery, scaling, trust, audit, control),
		WhyPoint6NotPass: []string{
			"Val C adds bounded resilience and scaling hardening only and does not implement operational visibility or integrated closure.",
			"Točka 6 final PASS remains reserved for Val E integrated closure.",
		},
		SurfaceRefs:          surfaceRefs,
		EvidenceRefs:         evidenceRefs,
		Limitations:          limitations,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val C binds Val B delivery packs to bounded resilience scenario packs, degraded-mode rules, recovery expectations, and scaling descriptors.",
			"Trust-path, audit-path, and control-plane safety checks remain evidence-linked advisory projections and do not approve deployment.",
		},
	}
}
