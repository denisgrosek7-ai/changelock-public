package main

import (
	"context"
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	productionUsabilityVal0ConfigIntegritySchema = "point4.production_usability.val0.config_integrity.v1"
	productionUsabilityVal0ExplainabilitySchema  = "point4.production_usability.val0.explainability_contract.v1"
	productionUsabilityVal0StatusModelSchema     = "point4.production_usability.val0.status_model.v1"
	productionUsabilityVal0OperationSchema       = "point4.production_usability.val0.operation_contracts.v1"
	productionUsabilityVal0DecisionQualitySchema = "point4.production_usability.val0.decision_quality.v1"
	productionUsabilityVal0NotificationSchema    = "point4.production_usability.val0.notification_taxonomy.v1"
	productionUsabilityVal0PermissionSchema      = "point4.production_usability.val0.permission_redaction.v1"
	productionUsabilityVal0RecoverySchema        = "point4.production_usability.val0.recovery_contract.v1"
	productionUsabilityVal0ActionModesSchema     = "point4.production_usability.val0.action_modes.v1"
	productionUsabilityVal0ProofsSchema          = "point4.production_usability.val0.proofs.v1"
)

type productionUsabilityVal0ConfigIntegrityResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.ConfigIntegrityContract `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type productionUsabilityVal0ExplainabilityResponse struct {
	SchemaVersion string                                    `json:"schema_version"`
	GeneratedAt   time.Time                                 `json:"generated_at"`
	CurrentState  string                                    `json:"current_state"`
	Model         operability.ExplainabilityPayloadContract `json:"model"`
	RouteRefs     []string                                  `json:"route_refs,omitempty"`
	Limitations   []string                                  `json:"limitations,omitempty"`
}

type productionUsabilityVal0StatusModelResponse struct {
	SchemaVersion string                             `json:"schema_version"`
	GeneratedAt   time.Time                          `json:"generated_at"`
	CurrentState  string                             `json:"current_state"`
	Model         operability.OperationalStatusModel `json:"model"`
	RouteRefs     []string                           `json:"route_refs,omitempty"`
	Limitations   []string                           `json:"limitations,omitempty"`
}

type productionUsabilityVal0OperationResponse struct {
	SchemaVersion string                             `json:"schema_version"`
	GeneratedAt   time.Time                          `json:"generated_at"`
	CurrentState  string                             `json:"current_state"`
	Model         operability.OperationContractModel `json:"model"`
	RouteRefs     []string                           `json:"route_refs,omitempty"`
	Limitations   []string                           `json:"limitations,omitempty"`
}

type productionUsabilityVal0DecisionQualityResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.DecisionQualityContract `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type productionUsabilityVal0NotificationResponse struct {
	SchemaVersion string                                   `json:"schema_version"`
	GeneratedAt   time.Time                                `json:"generated_at"`
	CurrentState  string                                   `json:"current_state"`
	Model         operability.NotificationTaxonomyContract `json:"model"`
	RouteRefs     []string                                 `json:"route_refs,omitempty"`
	Limitations   []string                                 `json:"limitations,omitempty"`
}

type productionUsabilityVal0PermissionResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	GeneratedAt   time.Time                               `json:"generated_at"`
	CurrentState  string                                  `json:"current_state"`
	Model         operability.PermissionRedactionContract `json:"model"`
	RouteRefs     []string                                `json:"route_refs,omitempty"`
	Limitations   []string                                `json:"limitations,omitempty"`
}

type productionUsabilityVal0RecoveryResponse struct {
	SchemaVersion string                         `json:"schema_version"`
	GeneratedAt   time.Time                      `json:"generated_at"`
	CurrentState  string                         `json:"current_state"`
	Model         operability.RecoveryUXContract `json:"model"`
	RouteRefs     []string                       `json:"route_refs,omitempty"`
	Limitations   []string                       `json:"limitations,omitempty"`
}

type productionUsabilityVal0ActionModesResponse struct {
	SchemaVersion string                         `json:"schema_version"`
	GeneratedAt   time.Time                      `json:"generated_at"`
	CurrentState  string                         `json:"current_state"`
	Model         operability.ActionModeTaxonomy `json:"model"`
	RouteRefs     []string                       `json:"route_refs,omitempty"`
	Limitations   []string                       `json:"limitations,omitempty"`
}

type productionUsabilityVal0ProofsResponse struct {
	SchemaVersion            string    `json:"schema_version"`
	GeneratedAt              time.Time `json:"generated_at"`
	CurrentState             string    `json:"current_state"`
	Point3DependencyState    string    `json:"point_3_dependency_state"`
	Val0State                string    `json:"val_0_state"`
	Point4State              string    `json:"point_4_state"`
	ConfigIntegrityState     string    `json:"config_integrity_state"`
	ExplainabilityState      string    `json:"explainability_state"`
	StatusModelState         string    `json:"status_model_state"`
	OperationContractState   string    `json:"operation_contract_state"`
	DecisionQualityState     string    `json:"decision_quality_state"`
	NotificationState        string    `json:"notification_state"`
	PermissionRedactionState string    `json:"permission_redaction_state"`
	RecoveryState            string    `json:"recovery_state"`
	ActionModeState          string    `json:"action_mode_state"`
	WhyPoint4NotPass         []string  `json:"why_point_4_not_pass,omitempty"`
	SurfaceRefs              []string  `json:"surface_refs,omitempty"`
	EvidenceRefs             []string  `json:"evidence_refs,omitempty"`
	Limitations              []string  `json:"limitations,omitempty"`
	IntegrationSummary       []string  `json:"integration_summary,omitempty"`
}

func (s server) productionUsabilityVal0ConfigIntegrityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityVal0ConfigIntegrity())
}

func (s server) productionUsabilityVal0ExplainabilityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityVal0Explainability())
}

func (s server) productionUsabilityVal0StatusModelHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityVal0StatusModel())
}

func (s server) productionUsabilityVal0OperationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityVal0OperationContracts())
}

func (s server) productionUsabilityVal0DecisionQualityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityVal0DecisionQuality())
}

func (s server) productionUsabilityVal0NotificationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityVal0NotificationTaxonomy())
}

func (s server) productionUsabilityVal0PermissionHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityVal0PermissionRedaction())
}

func (s server) productionUsabilityVal0RecoveryHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityVal0RecoveryContract())
}

func (s server) productionUsabilityVal0ActionModesHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityVal0ActionModes())
}

func (s server) productionUsabilityVal0ProofsHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildProductionUsabilityVal0Proofs(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildProductionUsabilityVal0ConfigIntegrity() productionUsabilityVal0ConfigIntegrityResponse {
	model := operability.ProductionUsabilityVal0ConfigIntegrity()
	return productionUsabilityVal0ConfigIntegrityResponse{
		SchemaVersion: productionUsabilityVal0ConfigIntegritySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityVal0ConfigIntegrityState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityVal0Explainability() productionUsabilityVal0ExplainabilityResponse {
	model := operability.ProductionUsabilityVal0ExplainabilityContract()
	return productionUsabilityVal0ExplainabilityResponse{
		SchemaVersion: productionUsabilityVal0ExplainabilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityVal0ExplainabilityState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/val0/permission-redaction",
			"/v1/production/usability-operability-recovery/val0/recovery-contract",
			"/v1/production/usability-operability-recovery/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityVal0StatusModel() productionUsabilityVal0StatusModelResponse {
	model := operability.ProductionUsabilityVal0OperationalStatusModel()
	return productionUsabilityVal0StatusModelResponse{
		SchemaVersion: productionUsabilityVal0StatusModelSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityVal0StatusModelState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityVal0OperationContracts() productionUsabilityVal0OperationResponse {
	model := operability.ProductionUsabilityVal0OperationContractModel()
	return productionUsabilityVal0OperationResponse{
		SchemaVersion: productionUsabilityVal0OperationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityVal0OperationContractState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/val0/action-modes",
			"/v1/production/usability-operability-recovery/val0/recovery-contract",
			"/v1/production/usability-operability-recovery/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityVal0DecisionQuality() productionUsabilityVal0DecisionQualityResponse {
	model := operability.ProductionUsabilityVal0DecisionQualityContract()
	return productionUsabilityVal0DecisionQualityResponse{
		SchemaVersion: productionUsabilityVal0DecisionQualitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityVal0DecisionQualityState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityVal0NotificationTaxonomy() productionUsabilityVal0NotificationResponse {
	model := operability.ProductionUsabilityVal0NotificationTaxonomyContract()
	return productionUsabilityVal0NotificationResponse{
		SchemaVersion: productionUsabilityVal0NotificationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityVal0NotificationState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityVal0PermissionRedaction() productionUsabilityVal0PermissionResponse {
	model := operability.ProductionUsabilityVal0PermissionRedactionContract()
	return productionUsabilityVal0PermissionResponse{
		SchemaVersion: productionUsabilityVal0PermissionSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityVal0PermissionRedactionState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/val0/explainability-contract",
			"/v1/production/usability-operability-recovery/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityVal0RecoveryContract() productionUsabilityVal0RecoveryResponse {
	model := operability.ProductionUsabilityVal0RecoveryUXContract()
	return productionUsabilityVal0RecoveryResponse{
		SchemaVersion: productionUsabilityVal0RecoverySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityVal0RecoveryState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/val0/operation-contracts",
			"/v1/production/usability-operability-recovery/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityVal0ActionModes() productionUsabilityVal0ActionModesResponse {
	model := operability.ProductionUsabilityVal0ActionModeTaxonomy()
	return productionUsabilityVal0ActionModesResponse{
		SchemaVersion: productionUsabilityVal0ActionModesSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityVal0ActionModeState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/val0/operation-contracts",
			"/v1/production/usability-operability-recovery/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityVal0ProofsCurrentState(point3DependencyState string, configIntegrity operability.ConfigIntegrityContract, explainability operability.ExplainabilityPayloadContract, statusModel operability.OperationalStatusModel, operationContracts operability.OperationContractModel, decisionQuality operability.DecisionQualityContract, notification operability.NotificationTaxonomyContract, permission operability.PermissionRedactionContract, recovery operability.RecoveryUXContract, actionModes operability.ActionModeTaxonomy, surfaceRefs, evidenceRefs, limitations, whyPoint4NotPass []string) string {
	return operability.EvaluateProductionUsabilityVal0ProofsState(
		point3DependencyState,
		operability.EvaluateProductionUsabilityVal0ConfigIntegrityState(configIntegrity),
		operability.EvaluateProductionUsabilityVal0ExplainabilityState(explainability),
		operability.EvaluateProductionUsabilityVal0StatusModelState(statusModel),
		operability.EvaluateProductionUsabilityVal0OperationContractState(operationContracts),
		operability.EvaluateProductionUsabilityVal0DecisionQualityState(decisionQuality),
		operability.EvaluateProductionUsabilityVal0NotificationState(notification),
		operability.EvaluateProductionUsabilityVal0PermissionRedactionState(permission),
		operability.EvaluateProductionUsabilityVal0RecoveryState(recovery),
		operability.EvaluateProductionUsabilityVal0ActionModeState(actionModes),
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint4NotPass,
	)
}

func (s server) buildProductionUsabilityVal0Proofs(ctx context.Context, filter phase4EnterpriseFilter) (productionUsabilityVal0ProofsResponse, error) {
	point3, err := s.buildEnterpriseWorkflowAuthorityValDProofs(ctx, filter)
	if err != nil {
		return productionUsabilityVal0ProofsResponse{}, err
	}

	configIntegrity := operability.ProductionUsabilityVal0ConfigIntegrity()
	explainability := operability.ProductionUsabilityVal0ExplainabilityContract()
	statusModel := operability.ProductionUsabilityVal0OperationalStatusModel()
	operationContracts := operability.ProductionUsabilityVal0OperationContractModel()
	decisionQuality := operability.ProductionUsabilityVal0DecisionQualityContract()
	notification := operability.ProductionUsabilityVal0NotificationTaxonomyContract()
	permission := operability.ProductionUsabilityVal0PermissionRedactionContract()
	recovery := operability.ProductionUsabilityVal0RecoveryUXContract()
	actionModes := operability.ProductionUsabilityVal0ActionModeTaxonomy()

	whyPoint4NotPass := []string{
		"Točka 4 full PASS remains closed because later waves still need stricter config execution, UI and API resilience, support bundle and upgrade flows, and a final usability gate.",
		"Val 0 establishes only the semantic and fail-closed contract foundation for production usability, operability, explainability, recovery, and supportability.",
	}
	surfaceRefs := []string{
		"/v1/enterprise/workflow-authority/vald/proofs",
		"/v1/production/usability-operability-recovery/val0/config-integrity",
		"/v1/production/usability-operability-recovery/val0/explainability-contract",
		"/v1/production/usability-operability-recovery/val0/status-model",
		"/v1/production/usability-operability-recovery/val0/operation-contracts",
		"/v1/production/usability-operability-recovery/val0/decision-quality",
		"/v1/production/usability-operability-recovery/val0/notification-taxonomy",
		"/v1/production/usability-operability-recovery/val0/permission-redaction",
		"/v1/production/usability-operability-recovery/val0/recovery-contract",
		"/v1/production/usability-operability-recovery/val0/action-modes",
		"/v1/production/usability-operability-recovery/val0/proofs",
	}
	evidenceRefs := []string{
		"/v1/enterprise/phase4/proofs",
		"/v1/enterprise/workflow-authority/vald/proofs",
		"config_integrity_contract",
		"explainability_contract",
		"operational_status_model",
		"operation_retry_contract",
		"permission_redaction_contract",
	}
	limitations := []string{
		"Val 0 only proves discipline foundation and does not claim Točka 4 production usability completion.",
		"All usability, diagnostics, cache, notification, and support outputs remain projections and never replace canonical evidence truth.",
	}
	currentState := buildProductionUsabilityVal0ProofsCurrentState(
		point3.CurrentState,
		configIntegrity,
		explainability,
		statusModel,
		operationContracts,
		decisionQuality,
		notification,
		permission,
		recovery,
		actionModes,
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint4NotPass,
	)
	val0State := operability.EvaluateProductionUsabilityVal0State(
		point3.CurrentState,
		operability.EvaluateProductionUsabilityVal0ConfigIntegrityState(configIntegrity),
		operability.EvaluateProductionUsabilityVal0ExplainabilityState(explainability),
		operability.EvaluateProductionUsabilityVal0StatusModelState(statusModel),
		operability.EvaluateProductionUsabilityVal0OperationContractState(operationContracts),
		operability.EvaluateProductionUsabilityVal0DecisionQualityState(decisionQuality),
		operability.EvaluateProductionUsabilityVal0NotificationState(notification),
		operability.EvaluateProductionUsabilityVal0PermissionRedactionState(permission),
		operability.EvaluateProductionUsabilityVal0RecoveryState(recovery),
		operability.EvaluateProductionUsabilityVal0ActionModeState(actionModes),
	)

	return productionUsabilityVal0ProofsResponse{
		SchemaVersion:            productionUsabilityVal0ProofsSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             currentState,
		Point3DependencyState:    point3.CurrentState,
		Val0State:                val0State,
		Point4State:              operability.ProductionUsabilityPoint4StateNotComplete,
		ConfigIntegrityState:     operability.EvaluateProductionUsabilityVal0ConfigIntegrityState(configIntegrity),
		ExplainabilityState:      operability.EvaluateProductionUsabilityVal0ExplainabilityState(explainability),
		StatusModelState:         operability.EvaluateProductionUsabilityVal0StatusModelState(statusModel),
		OperationContractState:   operability.EvaluateProductionUsabilityVal0OperationContractState(operationContracts),
		DecisionQualityState:     operability.EvaluateProductionUsabilityVal0DecisionQualityState(decisionQuality),
		NotificationState:        operability.EvaluateProductionUsabilityVal0NotificationState(notification),
		PermissionRedactionState: operability.EvaluateProductionUsabilityVal0PermissionRedactionState(permission),
		RecoveryState:            operability.EvaluateProductionUsabilityVal0RecoveryState(recovery),
		ActionModeState:          operability.EvaluateProductionUsabilityVal0ActionModeState(actionModes),
		WhyPoint4NotPass:         whyPoint4NotPass,
		SurfaceRefs:              surfaceRefs,
		EvidenceRefs:             evidenceRefs,
		Limitations:              limitations,
		IntegrationSummary: []string{
			"Val 0 defines fail-closed production usability, operability, explainability, retry, notification, permission, recovery, and safe automation contracts over the existing canonical evidence spine.",
			"Val 0 keeps effective config view, UI, CLI, API, diagnostics, support, cache, and notification surfaces as projections only rather than a new source of truth.",
			"Point 4 remains not complete because later waves still need execution-grade config, resilience, support bundle, upgrade, and final usability gate behavior.",
		},
	}, nil
}
