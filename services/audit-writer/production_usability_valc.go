package main

import (
	"context"
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	productionUsabilityValCReadinessSchema         = "point4.production_usability.valc.readiness.v1"
	productionUsabilityValCGuidedReadinessSchema   = "point4.production_usability.valc.guided_readiness.v1"
	productionUsabilityValCSupportBundleSchema     = "point4.production_usability.valc.support_bundle.v1"
	productionUsabilityValCDiagnosticsSchema       = "point4.production_usability.valc.diagnostics.v1"
	productionUsabilityValCHealthSnapshotSchema    = "point4.production_usability.valc.health_snapshot.v1"
	productionUsabilityValCRecoveryPlaybooksSchema = "point4.production_usability.valc.recovery_playbooks.v1"
	productionUsabilityValCUpgradeAdvisorySchema   = "point4.production_usability.valc.upgrade_rollback_advisory.v1"
	productionUsabilityValCPermissionFlowsSchema   = "point4.production_usability.valc.permission_support_flows.v1"
	productionUsabilityValCExportSafetySchema      = "point4.production_usability.valc.redaction_export_safety.v1"
	productionUsabilityValCProofsSchema            = "point4.production_usability.valc.proofs.v1"
)

type productionUsabilityValCReadinessResponse struct {
	SchemaVersion string                          `json:"schema_version"`
	GeneratedAt   time.Time                       `json:"generated_at"`
	CurrentState  string                          `json:"current_state"`
	Model         operability.ReadinessCheckModel `json:"model"`
	RouteRefs     []string                        `json:"route_refs,omitempty"`
	Limitations   []string                        `json:"limitations,omitempty"`
}

type productionUsabilityValCGuidedReadinessResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.GuidedReadinessBaseline `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type productionUsabilityValCSupportBundleResponse struct {
	SchemaVersion string                               `json:"schema_version"`
	GeneratedAt   time.Time                            `json:"generated_at"`
	CurrentState  string                               `json:"current_state"`
	Model         operability.SupportBundleQualityGate `json:"model"`
	RouteRefs     []string                             `json:"route_refs,omitempty"`
}

type productionUsabilityValCDiagnosticsResponse struct {
	SchemaVersion string                                `json:"schema_version"`
	GeneratedAt   time.Time                             `json:"generated_at"`
	CurrentState  string                                `json:"current_state"`
	Model         operability.DiagnosticsHardeningModel `json:"model"`
	RouteRefs     []string                              `json:"route_refs,omitempty"`
	Limitations   []string                              `json:"limitations,omitempty"`
}

type productionUsabilityValCHealthSnapshotResponse struct {
	SchemaVersion string                          `json:"schema_version"`
	GeneratedAt   time.Time                       `json:"generated_at"`
	CurrentState  string                          `json:"current_state"`
	Model         operability.HealthSnapshotModel `json:"model"`
	RouteRefs     []string                        `json:"route_refs,omitempty"`
	Limitations   []string                        `json:"limitations,omitempty"`
}

type productionUsabilityValCRecoveryPlaybookResponse struct {
	SchemaVersion string                            `json:"schema_version"`
	GeneratedAt   time.Time                         `json:"generated_at"`
	CurrentState  string                            `json:"current_state"`
	Model         operability.RecoveryPlaybookModel `json:"model"`
	RouteRefs     []string                          `json:"route_refs,omitempty"`
	Limitations   []string                          `json:"limitations,omitempty"`
}

type productionUsabilityValCUpgradeAdvisoryResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.UpgradeRollbackAdvisory `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type productionUsabilityValCPermissionSupportResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         operability.PermissionAwareSupportFlowModel `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type productionUsabilityValCExportSafetyResponse struct {
	SchemaVersion string                               `json:"schema_version"`
	GeneratedAt   time.Time                            `json:"generated_at"`
	CurrentState  string                               `json:"current_state"`
	Model         operability.RedactionSafeExportModel `json:"model"`
	RouteRefs     []string                             `json:"route_refs,omitempty"`
}

type productionUsabilityValCProofsResponse struct {
	SchemaVersion                string    `json:"schema_version"`
	GeneratedAt                  time.Time `json:"generated_at"`
	CurrentState                 string    `json:"current_state"`
	Val0DependencyState          string    `json:"val_0_dependency_state"`
	Val0FoundationState          string    `json:"val_0_foundation_state"`
	ValADependencyState          string    `json:"val_a_dependency_state"`
	ValACoreState                string    `json:"val_a_core_state"`
	ValBDependencyState          string    `json:"val_b_dependency_state"`
	ValBResilienceState          string    `json:"val_b_resilience_state"`
	ValCState                    string    `json:"val_c_state"`
	Point4State                  string    `json:"point_4_state"`
	ReadinessState               string    `json:"readiness_state"`
	GuidedReadinessState         string    `json:"guided_readiness_state"`
	SupportBundleState           string    `json:"support_bundle_state"`
	DiagnosticsState             string    `json:"diagnostics_state"`
	HealthSnapshotState          string    `json:"health_snapshot_state"`
	RecoveryPlaybookState        string    `json:"recovery_playbook_state"`
	UpgradeRollbackAdvisoryState string    `json:"upgrade_rollback_advisory_state"`
	PermissionSupportFlowState   string    `json:"permission_support_flow_state"`
	RedactionExportSafetyState   string    `json:"redaction_export_safety_state"`
	WhyPoint4NotPass             []string  `json:"why_point_4_not_pass,omitempty"`
	SurfaceRefs                  []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                 []string  `json:"evidence_refs,omitempty"`
	Limitations                  []string  `json:"limitations,omitempty"`
	IntegrationSummary           []string  `json:"integration_summary,omitempty"`
}

func (s server) productionUsabilityValCReadinessHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValCReadiness())
}

func (s server) productionUsabilityValCGuidedReadinessHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValCGuidedReadiness())
}

func (s server) productionUsabilityValCSupportBundleHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValCSupportBundle())
}

func (s server) productionUsabilityValCDiagnosticsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValCDiagnostics())
}

func (s server) productionUsabilityValCHealthSnapshotHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValCHealthSnapshot())
}

func (s server) productionUsabilityValCRecoveryPlaybooksHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValCRecoveryPlaybooks())
}

func (s server) productionUsabilityValCUpgradeAdvisoryHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValCUpgradeAdvisory())
}

func (s server) productionUsabilityValCPermissionSupportHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValCPermissionSupport())
}

func (s server) productionUsabilityValCExportSafetyHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValCExportSafety())
}

func (s server) productionUsabilityValCProofsHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parsePhase4EnterpriseFilter(req)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildProductionUsabilityValCProofs(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildProductionUsabilityValCReadiness() productionUsabilityValCReadinessResponse {
	model := operability.ProductionUsabilityValCReadinessChecks()
	return productionUsabilityValCReadinessResponse{
		SchemaVersion: productionUsabilityValCReadinessSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValCReadinessState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valc/guided-readiness",
			"/v1/production/usability-operability-recovery/valc/health-snapshot",
			"/v1/production/usability-operability-recovery/valc/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValCGuidedReadiness() productionUsabilityValCGuidedReadinessResponse {
	model := operability.ProductionUsabilityValCGuidedReadiness()
	return productionUsabilityValCGuidedReadinessResponse{
		SchemaVersion: productionUsabilityValCGuidedReadinessSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValCGuidedReadinessState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valc/readiness",
			"/v1/production/usability-operability-recovery/valc/upgrade-rollback-advisory",
			"/v1/production/usability-operability-recovery/valc/proofs",
		},
		Limitations: []string{model.LimitationMessage},
	}
}

func buildProductionUsabilityValCSupportBundle() productionUsabilityValCSupportBundleResponse {
	model := operability.ProductionUsabilityValCSupportBundleQualityGate()
	return productionUsabilityValCSupportBundleResponse{
		SchemaVersion: productionUsabilityValCSupportBundleSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValCSupportBundleState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valc/diagnostics",
			"/v1/production/usability-operability-recovery/valc/permission-support-flows",
			"/v1/production/usability-operability-recovery/valc/proofs",
		},
	}
}

func buildProductionUsabilityValCDiagnostics() productionUsabilityValCDiagnosticsResponse {
	model := operability.ProductionUsabilityValCDiagnosticsHardening()
	return productionUsabilityValCDiagnosticsResponse{
		SchemaVersion: productionUsabilityValCDiagnosticsSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValCDiagnosticsState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valc/support-bundle",
			"/v1/production/usability-operability-recovery/valc/health-snapshot",
			"/v1/production/usability-operability-recovery/valc/proofs",
		},
		Limitations: []string{model.LimitationMessage},
	}
}

func buildProductionUsabilityValCHealthSnapshot() productionUsabilityValCHealthSnapshotResponse {
	model := operability.ProductionUsabilityValCHealthSnapshot()
	return productionUsabilityValCHealthSnapshotResponse{
		SchemaVersion: productionUsabilityValCHealthSnapshotSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValCHealthSnapshotState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valc/readiness",
			"/v1/production/usability-operability-recovery/valc/diagnostics",
			"/v1/production/usability-operability-recovery/valc/proofs",
		},
		Limitations: []string{model.LimitationMessage},
	}
}

func buildProductionUsabilityValCRecoveryPlaybooks() productionUsabilityValCRecoveryPlaybookResponse {
	model := operability.ProductionUsabilityValCRecoveryPlaybooks()
	return productionUsabilityValCRecoveryPlaybookResponse{
		SchemaVersion: productionUsabilityValCRecoveryPlaybooksSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValCRecoveryPlaybookState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valc/support-bundle",
			"/v1/production/usability-operability-recovery/valc/upgrade-rollback-advisory",
			"/v1/production/usability-operability-recovery/valc/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValCUpgradeAdvisory() productionUsabilityValCUpgradeAdvisoryResponse {
	model := operability.ProductionUsabilityValCUpgradeRollbackAdvisory()
	return productionUsabilityValCUpgradeAdvisoryResponse{
		SchemaVersion: productionUsabilityValCUpgradeAdvisorySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValCUpgradeAdvisoryState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valc/guided-readiness",
			"/v1/production/usability-operability-recovery/valc/recovery-playbooks",
			"/v1/production/usability-operability-recovery/valc/proofs",
		},
		Limitations: []string{model.LimitationDisclaimer},
	}
}

func buildProductionUsabilityValCPermissionSupport() productionUsabilityValCPermissionSupportResponse {
	model := operability.ProductionUsabilityValCPermissionSupportFlows()
	return productionUsabilityValCPermissionSupportResponse{
		SchemaVersion: productionUsabilityValCPermissionFlowsSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValCPermissionSupportState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valc/support-bundle",
			"/v1/production/usability-operability-recovery/valc/redaction-export-safety",
			"/v1/production/usability-operability-recovery/valc/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValCExportSafety() productionUsabilityValCExportSafetyResponse {
	model := operability.ProductionUsabilityValCRedactionSafeExport()
	return productionUsabilityValCExportSafetyResponse{
		SchemaVersion: productionUsabilityValCExportSafetySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValCExportSafetyState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valc/support-bundle",
			"/v1/production/usability-operability-recovery/valc/permission-support-flows",
			"/v1/production/usability-operability-recovery/valc/proofs",
		},
	}
}

func buildProductionUsabilityValCProofsCurrentState(val0State, valAState, valBState string, readiness operability.ReadinessCheckModel, guided operability.GuidedReadinessBaseline, supportBundle operability.SupportBundleQualityGate, diagnostics operability.DiagnosticsHardeningModel, health operability.HealthSnapshotModel, recovery operability.RecoveryPlaybookModel, advisory operability.UpgradeRollbackAdvisory, permissionFlows operability.PermissionAwareSupportFlowModel, exportSafety operability.RedactionSafeExportModel, surfaceRefs, evidenceRefs, limitations, whyPoint4NotPass []string) string {
	return operability.EvaluateProductionUsabilityValCProofsState(
		val0State,
		valAState,
		valBState,
		operability.EvaluateProductionUsabilityValCReadinessState(readiness),
		operability.EvaluateProductionUsabilityValCGuidedReadinessState(guided),
		operability.EvaluateProductionUsabilityValCSupportBundleState(supportBundle),
		operability.EvaluateProductionUsabilityValCDiagnosticsState(diagnostics),
		operability.EvaluateProductionUsabilityValCHealthSnapshotState(health),
		operability.EvaluateProductionUsabilityValCRecoveryPlaybookState(recovery),
		operability.EvaluateProductionUsabilityValCUpgradeAdvisoryState(advisory),
		operability.EvaluateProductionUsabilityValCPermissionSupportState(permissionFlows),
		operability.EvaluateProductionUsabilityValCExportSafetyState(exportSafety),
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint4NotPass,
	)
}

func (s server) buildProductionUsabilityValCProofs(ctx context.Context, filter phase4EnterpriseFilter) (productionUsabilityValCProofsResponse, error) {
	val0, err := s.buildProductionUsabilityVal0Proofs(ctx, filter)
	if err != nil {
		return productionUsabilityValCProofsResponse{}, err
	}
	valA, err := s.buildProductionUsabilityValAProofs(ctx, filter)
	if err != nil {
		return productionUsabilityValCProofsResponse{}, err
	}
	valB, err := s.buildProductionUsabilityValBProofs(ctx, filter)
	if err != nil {
		return productionUsabilityValCProofsResponse{}, err
	}

	readiness := operability.ProductionUsabilityValCReadinessChecks()
	guided := operability.ProductionUsabilityValCGuidedReadiness()
	supportBundle := operability.ProductionUsabilityValCSupportBundleQualityGate()
	diagnostics := operability.ProductionUsabilityValCDiagnosticsHardening()
	health := operability.ProductionUsabilityValCHealthSnapshot()
	recovery := operability.ProductionUsabilityValCRecoveryPlaybooks()
	advisory := operability.ProductionUsabilityValCUpgradeRollbackAdvisory()
	permissionFlows := operability.ProductionUsabilityValCPermissionSupportFlows()
	exportSafety := operability.ProductionUsabilityValCRedactionSafeExport()

	whyPoint4NotPass := []string{
		"Točka 4 remains not complete because later waves still need final usability gate review and integrated closure.",
		"Val C proves supportability and lifecycle operations readiness only; it does not prove final usability completion or integrated closure.",
	}
	surfaceRefs := []string{
		"/v1/production/usability-operability-recovery/val0/proofs",
		"/v1/production/usability-operability-recovery/vala/proofs",
		"/v1/production/usability-operability-recovery/valb/proofs",
		"/v1/production/usability-operability-recovery/valc/readiness",
		"/v1/production/usability-operability-recovery/valc/guided-readiness",
		"/v1/production/usability-operability-recovery/valc/support-bundle",
		"/v1/production/usability-operability-recovery/valc/diagnostics",
		"/v1/production/usability-operability-recovery/valc/health-snapshot",
		"/v1/production/usability-operability-recovery/valc/recovery-playbooks",
		"/v1/production/usability-operability-recovery/valc/upgrade-rollback-advisory",
		"/v1/production/usability-operability-recovery/valc/permission-support-flows",
		"/v1/production/usability-operability-recovery/valc/redaction-export-safety",
		"/v1/production/usability-operability-recovery/valc/proofs",
	}
	evidenceRefs := []string{
		"val0_proofs",
		"vala_proofs",
		"valb_proofs",
		"readiness_projection",
		"guided_readiness_projection",
		"support_bundle_manifest",
		"diagnostics_projection",
		"health_snapshot_projection",
		"recovery_playbook_contract",
		"upgrade_rollback_advisory_contract",
		"permission_support_flow_contract",
		"redaction_export_safety_contract",
	}
	limitations := []string{
		"Val C proves bounded supportability and lifecycle-operation projections only and does not claim final usability completion.",
		"Support, diagnostics, readiness, health, advisory, and export surfaces remain projection-only and do not replace canonical truth or governed workflow state.",
	}
	currentState := buildProductionUsabilityValCProofsCurrentState(
		val0.Val0State,
		valA.ValAState,
		valB.ValBState,
		readiness,
		guided,
		supportBundle,
		diagnostics,
		health,
		recovery,
		advisory,
		permissionFlows,
		exportSafety,
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint4NotPass,
	)
	valCState := operability.EvaluateProductionUsabilityValCState(
		val0.Val0State,
		valA.ValAState,
		valB.ValBState,
		operability.EvaluateProductionUsabilityValCReadinessState(readiness),
		operability.EvaluateProductionUsabilityValCGuidedReadinessState(guided),
		operability.EvaluateProductionUsabilityValCSupportBundleState(supportBundle),
		operability.EvaluateProductionUsabilityValCDiagnosticsState(diagnostics),
		operability.EvaluateProductionUsabilityValCHealthSnapshotState(health),
		operability.EvaluateProductionUsabilityValCRecoveryPlaybookState(recovery),
		operability.EvaluateProductionUsabilityValCUpgradeAdvisoryState(advisory),
		operability.EvaluateProductionUsabilityValCPermissionSupportState(permissionFlows),
		operability.EvaluateProductionUsabilityValCExportSafetyState(exportSafety),
	)

	return productionUsabilityValCProofsResponse{
		SchemaVersion:                productionUsabilityValCProofsSchema,
		GeneratedAt:                  publicSampleTime(),
		CurrentState:                 currentState,
		Val0DependencyState:          val0.CurrentState,
		Val0FoundationState:          val0.Val0State,
		ValADependencyState:          valA.CurrentState,
		ValACoreState:                valA.ValAState,
		ValBDependencyState:          valB.CurrentState,
		ValBResilienceState:          valB.ValBState,
		ValCState:                    valCState,
		Point4State:                  operability.ProductionUsabilityPoint4StateNotComplete,
		ReadinessState:               operability.EvaluateProductionUsabilityValCReadinessState(readiness),
		GuidedReadinessState:         operability.EvaluateProductionUsabilityValCGuidedReadinessState(guided),
		SupportBundleState:           operability.EvaluateProductionUsabilityValCSupportBundleState(supportBundle),
		DiagnosticsState:             operability.EvaluateProductionUsabilityValCDiagnosticsState(diagnostics),
		HealthSnapshotState:          operability.EvaluateProductionUsabilityValCHealthSnapshotState(health),
		RecoveryPlaybookState:        operability.EvaluateProductionUsabilityValCRecoveryPlaybookState(recovery),
		UpgradeRollbackAdvisoryState: operability.EvaluateProductionUsabilityValCUpgradeAdvisoryState(advisory),
		PermissionSupportFlowState:   operability.EvaluateProductionUsabilityValCPermissionSupportState(permissionFlows),
		RedactionExportSafetyState:   operability.EvaluateProductionUsabilityValCExportSafetyState(exportSafety),
		WhyPoint4NotPass:             whyPoint4NotPass,
		SurfaceRefs:                  surfaceRefs,
		EvidenceRefs:                 evidenceRefs,
		Limitations:                  limitations,
		IntegrationSummary: []string{
			"Val C adds bounded readiness, support bundle, diagnostics, health snapshot, recovery playbook, advisory, permission-aware support, and export safety contracts over the active Val 0, Val A, and Val B spine.",
			"Val C keeps supportability outputs projection-only, redaction-safe, and fail-closed without introducing new lifecycle mutation authority.",
			"Point 4 remains not complete because final usability gate review and integrated closure remain later work.",
		},
	}, nil
}
