package main

import (
	"context"
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	productionUsabilityValAConfigFactorySchema    = "point4.production_usability.vala.config_factory.v1"
	productionUsabilityValABootstrapSchema        = "point4.production_usability.vala.bootstrap_validation.v1"
	productionUsabilityValAPolicySchema           = "point4.production_usability.vala.policy_schema.v1"
	productionUsabilityValAEffectiveConfigSchema  = "point4.production_usability.vala.effective_config.v1"
	productionUsabilityValARejectionsSchema       = "point4.production_usability.vala.rejections.v1"
	productionUsabilityValADryRunSchema           = "point4.production_usability.vala.policy_dry_run.v1"
	productionUsabilityValAExplainSchema          = "point4.production_usability.vala.explain_outputs.v1"
	productionUsabilityValARecoveryGuidanceSchema = "point4.production_usability.vala.recovery_guidance.v1"
	productionUsabilityValAFirstRunSchema         = "point4.production_usability.vala.first_run_bootstrap.v1"
	productionUsabilityValAUpgradePreviewSchema   = "point4.production_usability.vala.upgrade_impact_preview.v1"
	productionUsabilityValAProofsSchema           = "point4.production_usability.vala.proofs.v1"
)

type productionUsabilityValAConfigFactoryResponse struct {
	SchemaVersion string                                `json:"schema_version"`
	GeneratedAt   time.Time                             `json:"generated_at"`
	CurrentState  string                                `json:"current_state"`
	Model         operability.SchemaStrictConfigFactory `json:"model"`
	RouteRefs     []string                              `json:"route_refs,omitempty"`
	Limitations   []string                              `json:"limitations,omitempty"`
}

type productionUsabilityValABootstrapResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.BootstrapValidationCore `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type productionUsabilityValAPolicySchemaResponse struct {
	SchemaVersion string                             `json:"schema_version"`
	GeneratedAt   time.Time                          `json:"generated_at"`
	CurrentState  string                             `json:"current_state"`
	Model         operability.PolicySchemaDiscipline `json:"model"`
	RouteRefs     []string                           `json:"route_refs,omitempty"`
	Limitations   []string                           `json:"limitations,omitempty"`
}

type productionUsabilityValAEffectiveConfigResponse struct {
	SchemaVersion string                                `json:"schema_version"`
	GeneratedAt   time.Time                             `json:"generated_at"`
	CurrentState  string                                `json:"current_state"`
	Model         operability.EffectiveConfigInspection `json:"model"`
	RouteRefs     []string                              `json:"route_refs,omitempty"`
	Limitations   []string                              `json:"limitations,omitempty"`
}

type productionUsabilityValARejectionResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	GeneratedAt   time.Time                               `json:"generated_at"`
	CurrentState  string                                  `json:"current_state"`
	Model         operability.HumanReadableRejectionLayer `json:"model"`
	RouteRefs     []string                                `json:"route_refs,omitempty"`
	Limitations   []string                                `json:"limitations,omitempty"`
}

type productionUsabilityValADryRunResponse struct {
	SchemaVersion string                            `json:"schema_version"`
	GeneratedAt   time.Time                         `json:"generated_at"`
	CurrentState  string                            `json:"current_state"`
	Model         operability.PolicyDryRunAuditFlow `json:"model"`
	RouteRefs     []string                          `json:"route_refs,omitempty"`
	Limitations   []string                          `json:"limitations,omitempty"`
}

type productionUsabilityValAExplainResponse struct {
	SchemaVersion string                                    `json:"schema_version"`
	GeneratedAt   time.Time                                 `json:"generated_at"`
	CurrentState  string                                    `json:"current_state"`
	Model         operability.PermissionAwareExplainOutputs `json:"model"`
	RouteRefs     []string                                  `json:"route_refs,omitempty"`
	Limitations   []string                                  `json:"limitations,omitempty"`
}

type productionUsabilityValARecoveryGuidanceResponse struct {
	SchemaVersion string                           `json:"schema_version"`
	GeneratedAt   time.Time                        `json:"generated_at"`
	CurrentState  string                           `json:"current_state"`
	Model         operability.RecoveryGuidanceCore `json:"model"`
	RouteRefs     []string                         `json:"route_refs,omitempty"`
	Limitations   []string                         `json:"limitations,omitempty"`
}

type productionUsabilityValAFirstRunResponse struct {
	SchemaVersion string                            `json:"schema_version"`
	GeneratedAt   time.Time                         `json:"generated_at"`
	CurrentState  string                            `json:"current_state"`
	Model         operability.FirstRunSafeBootstrap `json:"model"`
	RouteRefs     []string                          `json:"route_refs,omitempty"`
	Limitations   []string                          `json:"limitations,omitempty"`
}

type productionUsabilityValAUpgradePreviewResponse struct {
	SchemaVersion string                           `json:"schema_version"`
	GeneratedAt   time.Time                        `json:"generated_at"`
	CurrentState  string                           `json:"current_state"`
	Model         operability.UpgradeImpactPreview `json:"model"`
	RouteRefs     []string                         `json:"route_refs,omitempty"`
	Limitations   []string                         `json:"limitations,omitempty"`
}

type productionUsabilityValAProofsResponse struct {
	SchemaVersion             string    `json:"schema_version"`
	GeneratedAt               time.Time `json:"generated_at"`
	CurrentState              string    `json:"current_state"`
	Val0DependencyState       string    `json:"val_0_dependency_state"`
	Val0FoundationState       string    `json:"val_0_foundation_state"`
	ValAState                 string    `json:"val_a_state"`
	Point4State               string    `json:"point_4_state"`
	ConfigFactoryState        string    `json:"config_factory_state"`
	BootstrapValidationState  string    `json:"bootstrap_validation_state"`
	PolicySchemaState         string    `json:"policy_schema_state"`
	EffectiveConfigState      string    `json:"effective_config_state"`
	RejectionLayerState       string    `json:"rejection_layer_state"`
	DryRunState               string    `json:"dry_run_state"`
	ExplainState              string    `json:"explain_state"`
	RecoveryGuidanceState     string    `json:"recovery_guidance_state"`
	FirstRunState             string    `json:"first_run_state"`
	UpgradeImpactPreviewState string    `json:"upgrade_impact_preview_state"`
	WhyPoint4NotPass          []string  `json:"why_point_4_not_pass,omitempty"`
	SurfaceRefs               []string  `json:"surface_refs,omitempty"`
	EvidenceRefs              []string  `json:"evidence_refs,omitempty"`
	Limitations               []string  `json:"limitations,omitempty"`
	IntegrationSummary        []string  `json:"integration_summary,omitempty"`
}

func (s server) productionUsabilityValAConfigFactoryHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValAConfigFactory())
}

func (s server) productionUsabilityValABootstrapValidationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValABootstrapValidation())
}

func (s server) productionUsabilityValAPolicySchemaHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValAPolicySchema())
}

func (s server) productionUsabilityValAEffectiveConfigHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValAEffectiveConfig())
}

func (s server) productionUsabilityValARejectionsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValARejectionLayer())
}

func (s server) productionUsabilityValADryRunHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValADryRun())
}

func (s server) productionUsabilityValAExplainHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValAExplain())
}

func (s server) productionUsabilityValARecoveryGuidanceHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValARecoveryGuidance())
}

func (s server) productionUsabilityValAFirstRunHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValAFirstRun())
}

func (s server) productionUsabilityValAUpgradePreviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValAUpgradePreview())
}

func (s server) productionUsabilityValAProofsHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildProductionUsabilityValAProofs(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildProductionUsabilityValAConfigFactory() productionUsabilityValAConfigFactoryResponse {
	model := operability.ProductionUsabilityValAConfigFactory()
	return productionUsabilityValAConfigFactoryResponse{
		SchemaVersion: productionUsabilityValAConfigFactorySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValAConfigFactoryState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vala/effective-config",
			"/v1/production/usability-operability-recovery/vala/bootstrap-validation",
			"/v1/production/usability-operability-recovery/vala/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValABootstrapValidation() productionUsabilityValABootstrapResponse {
	model := operability.ProductionUsabilityValABootstrapValidation()
	return productionUsabilityValABootstrapResponse{
		SchemaVersion: productionUsabilityValABootstrapSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValABootstrapValidationState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vala/rejections",
			"/v1/production/usability-operability-recovery/vala/recovery-guidance",
			"/v1/production/usability-operability-recovery/vala/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValAPolicySchema() productionUsabilityValAPolicySchemaResponse {
	model := operability.ProductionUsabilityValAPolicySchemaDiscipline()
	return productionUsabilityValAPolicySchemaResponse{
		SchemaVersion: productionUsabilityValAPolicySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValAPolicySchemaState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vala/policy-dry-run",
			"/v1/production/usability-operability-recovery/vala/effective-config",
			"/v1/production/usability-operability-recovery/vala/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValAEffectiveConfig() productionUsabilityValAEffectiveConfigResponse {
	model := operability.ProductionUsabilityValAEffectiveConfigInspection()
	return productionUsabilityValAEffectiveConfigResponse{
		SchemaVersion: productionUsabilityValAEffectiveConfigSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValAEffectiveConfigState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vala/config-factory",
			"/v1/production/usability-operability-recovery/vala/policy-schema",
			"/v1/production/usability-operability-recovery/vala/proofs",
		},
		Limitations: model.LimitationNotes,
	}
}

func buildProductionUsabilityValARejectionLayer() productionUsabilityValARejectionResponse {
	model := operability.ProductionUsabilityValAHumanReadableRejectionLayer()
	return productionUsabilityValARejectionResponse{
		SchemaVersion: productionUsabilityValARejectionsSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValARejectionLayerState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vala/bootstrap-validation",
			"/v1/production/usability-operability-recovery/vala/explain",
			"/v1/production/usability-operability-recovery/vala/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValADryRun() productionUsabilityValADryRunResponse {
	model := operability.ProductionUsabilityValAPolicyDryRunAuditFlow()
	return productionUsabilityValADryRunResponse{
		SchemaVersion: productionUsabilityValADryRunSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValADryRunState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vala/rejections",
			"/v1/production/usability-operability-recovery/vala/recovery-guidance",
			"/v1/production/usability-operability-recovery/vala/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValAExplain() productionUsabilityValAExplainResponse {
	model := operability.ProductionUsabilityValAPermissionAwareExplainOutputs()
	return productionUsabilityValAExplainResponse{
		SchemaVersion: productionUsabilityValAExplainSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValAExplainState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vala/rejections",
			"/v1/production/usability-operability-recovery/vala/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValARecoveryGuidance() productionUsabilityValARecoveryGuidanceResponse {
	model := operability.ProductionUsabilityValARecoveryGuidance()
	return productionUsabilityValARecoveryGuidanceResponse{
		SchemaVersion: productionUsabilityValARecoveryGuidanceSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValARecoveryGuidanceState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vala/bootstrap-validation",
			"/v1/production/usability-operability-recovery/vala/policy-dry-run",
			"/v1/production/usability-operability-recovery/vala/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValAFirstRun() productionUsabilityValAFirstRunResponse {
	model := operability.ProductionUsabilityValAFirstRunSafeBootstrap()
	return productionUsabilityValAFirstRunResponse{
		SchemaVersion: productionUsabilityValAFirstRunSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValAFirstRunState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vala/bootstrap-validation",
			"/v1/production/usability-operability-recovery/vala/effective-config",
			"/v1/production/usability-operability-recovery/vala/proofs",
		},
		Limitations: []string{
			"First-run bootstrap remains a non-mutating safe validation path and does not claim full production completion.",
		},
	}
}

func buildProductionUsabilityValAUpgradePreview() productionUsabilityValAUpgradePreviewResponse {
	model := operability.ProductionUsabilityValAUpgradeImpactPreview()
	return productionUsabilityValAUpgradePreviewResponse{
		SchemaVersion: productionUsabilityValAUpgradePreviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValAUpgradePreviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vala/config-factory",
			"/v1/production/usability-operability-recovery/vala/policy-schema",
			"/v1/production/usability-operability-recovery/vala/proofs",
		},
		Limitations: []string{
			"Upgrade impact preview is bounded to config and policy schema perspective and does not perform live upgrade or rollback orchestration.",
		},
	}
}

func buildProductionUsabilityValAProofsCurrentState(val0State string, configFactory operability.SchemaStrictConfigFactory, bootstrap operability.BootstrapValidationCore, policySchema operability.PolicySchemaDiscipline, effectiveConfig operability.EffectiveConfigInspection, rejectionLayer operability.HumanReadableRejectionLayer, dryRun operability.PolicyDryRunAuditFlow, explain operability.PermissionAwareExplainOutputs, recovery operability.RecoveryGuidanceCore, firstRun operability.FirstRunSafeBootstrap, upgradePreview operability.UpgradeImpactPreview, surfaceRefs, evidenceRefs, limitations, whyPoint4NotPass []string) string {
	return operability.EvaluateProductionUsabilityValAProofsState(
		val0State,
		operability.EvaluateProductionUsabilityValAConfigFactoryState(configFactory),
		operability.EvaluateProductionUsabilityValABootstrapValidationState(bootstrap),
		operability.EvaluateProductionUsabilityValAPolicySchemaState(policySchema),
		operability.EvaluateProductionUsabilityValAEffectiveConfigState(effectiveConfig),
		operability.EvaluateProductionUsabilityValARejectionLayerState(rejectionLayer),
		operability.EvaluateProductionUsabilityValADryRunState(dryRun),
		operability.EvaluateProductionUsabilityValAExplainState(explain),
		operability.EvaluateProductionUsabilityValARecoveryGuidanceState(recovery),
		operability.EvaluateProductionUsabilityValAFirstRunState(firstRun),
		operability.EvaluateProductionUsabilityValAUpgradePreviewState(upgradePreview),
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint4NotPass,
	)
}

func (s server) buildProductionUsabilityValAProofs(ctx context.Context, filter phase4EnterpriseFilter) (productionUsabilityValAProofsResponse, error) {
	val0, err := s.buildProductionUsabilityVal0Proofs(ctx, filter)
	if err != nil {
		return productionUsabilityValAProofsResponse{}, err
	}

	configFactory := operability.ProductionUsabilityValAConfigFactory()
	bootstrap := operability.ProductionUsabilityValABootstrapValidation()
	policySchema := operability.ProductionUsabilityValAPolicySchemaDiscipline()
	effectiveConfig := operability.ProductionUsabilityValAEffectiveConfigInspection()
	rejectionLayer := operability.ProductionUsabilityValAHumanReadableRejectionLayer()
	dryRun := operability.ProductionUsabilityValAPolicyDryRunAuditFlow()
	explain := operability.ProductionUsabilityValAPermissionAwareExplainOutputs()
	recovery := operability.ProductionUsabilityValARecoveryGuidance()
	firstRun := operability.ProductionUsabilityValAFirstRunSafeBootstrap()
	upgradePreview := operability.ProductionUsabilityValAUpgradeImpactPreview()

	whyPoint4NotPass := []string{
		"Točka 4 remains not complete because later waves still need resilience hardening, supportability flows, lifecycle usability operations, and final usability gating.",
		"Val A proves Config & Explainability Core readiness only; it does not prove broader support bundle, upgrade lifecycle, or integrated closure behavior.",
	}
	surfaceRefs := []string{
		"/v1/production/usability-operability-recovery/val0/proofs",
		"/v1/production/usability-operability-recovery/vala/config-factory",
		"/v1/production/usability-operability-recovery/vala/bootstrap-validation",
		"/v1/production/usability-operability-recovery/vala/policy-schema",
		"/v1/production/usability-operability-recovery/vala/effective-config",
		"/v1/production/usability-operability-recovery/vala/rejections",
		"/v1/production/usability-operability-recovery/vala/policy-dry-run",
		"/v1/production/usability-operability-recovery/vala/explain",
		"/v1/production/usability-operability-recovery/vala/recovery-guidance",
		"/v1/production/usability-operability-recovery/vala/first-run-bootstrap",
		"/v1/production/usability-operability-recovery/vala/upgrade-impact-preview",
		"/v1/production/usability-operability-recovery/vala/proofs",
	}
	evidenceRefs := []string{
		"/v1/production/usability-operability-recovery/val0/proofs",
		"config_factory_core",
		"bootstrap_validation_core",
		"policy_schema_core",
		"effective_config_inspection",
		"human_readable_rejection_layer",
		"policy_dry_run_audit_flow",
		"permission_aware_explain_outputs",
		"recovery_guidance_core",
		"first_run_safe_bootstrap",
		"upgrade_impact_preview",
	}
	limitations := []string{
		"Val A proves config and explainability core readiness only and does not claim full production usability completion.",
		"Effective config, dry-run, explain, and upgrade preview outputs remain projection-only and never replace canonical truth.",
	}
	currentState := buildProductionUsabilityValAProofsCurrentState(
		val0.Val0State,
		configFactory,
		bootstrap,
		policySchema,
		effectiveConfig,
		rejectionLayer,
		dryRun,
		explain,
		recovery,
		firstRun,
		upgradePreview,
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint4NotPass,
	)
	valAState := operability.EvaluateProductionUsabilityValAState(
		val0.Val0State,
		operability.EvaluateProductionUsabilityValAConfigFactoryState(configFactory),
		operability.EvaluateProductionUsabilityValABootstrapValidationState(bootstrap),
		operability.EvaluateProductionUsabilityValAPolicySchemaState(policySchema),
		operability.EvaluateProductionUsabilityValAEffectiveConfigState(effectiveConfig),
		operability.EvaluateProductionUsabilityValARejectionLayerState(rejectionLayer),
		operability.EvaluateProductionUsabilityValADryRunState(dryRun),
		operability.EvaluateProductionUsabilityValAExplainState(explain),
		operability.EvaluateProductionUsabilityValARecoveryGuidanceState(recovery),
		operability.EvaluateProductionUsabilityValAFirstRunState(firstRun),
		operability.EvaluateProductionUsabilityValAUpgradePreviewState(upgradePreview),
	)

	return productionUsabilityValAProofsResponse{
		SchemaVersion:             productionUsabilityValAProofsSchema,
		GeneratedAt:               publicSampleTime(),
		CurrentState:              currentState,
		Val0DependencyState:       val0.CurrentState,
		Val0FoundationState:       val0.Val0State,
		ValAState:                 valAState,
		Point4State:               operability.ProductionUsabilityPoint4StateNotComplete,
		ConfigFactoryState:        operability.EvaluateProductionUsabilityValAConfigFactoryState(configFactory),
		BootstrapValidationState:  operability.EvaluateProductionUsabilityValABootstrapValidationState(bootstrap),
		PolicySchemaState:         operability.EvaluateProductionUsabilityValAPolicySchemaState(policySchema),
		EffectiveConfigState:      operability.EvaluateProductionUsabilityValAEffectiveConfigState(effectiveConfig),
		RejectionLayerState:       operability.EvaluateProductionUsabilityValARejectionLayerState(rejectionLayer),
		DryRunState:               operability.EvaluateProductionUsabilityValADryRunState(dryRun),
		ExplainState:              operability.EvaluateProductionUsabilityValAExplainState(explain),
		RecoveryGuidanceState:     operability.EvaluateProductionUsabilityValARecoveryGuidanceState(recovery),
		FirstRunState:             operability.EvaluateProductionUsabilityValAFirstRunState(firstRun),
		UpgradeImpactPreviewState: operability.EvaluateProductionUsabilityValAUpgradePreviewState(upgradePreview),
		WhyPoint4NotPass:          whyPoint4NotPass,
		SurfaceRefs:               surfaceRefs,
		EvidenceRefs:              evidenceRefs,
		Limitations:               limitations,
		IntegrationSummary: []string{
			"Val A turns the Val 0 contracts into schema-strict config, policy inspection, explainability, dry-run, recovery, first-run, and upgrade-preview production surfaces.",
			"Val A remains fail-closed on active Val 0 foundation and keeps all inspection and preview outputs as projections only.",
			"Point 4 still remains not complete because later waves must add resilience, supportability operations, and the final usability gate.",
		},
	}, nil
}
