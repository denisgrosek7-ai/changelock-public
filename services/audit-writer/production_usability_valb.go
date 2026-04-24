package main

import (
	"context"
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	productionUsabilityValBUIDataSchema          = "point4.production_usability.valb.ui_data_resilience.v1"
	productionUsabilityValBWindowingSchema       = "point4.production_usability.valb.windowing.v1"
	productionUsabilityValBResultSemanticsSchema = "point4.production_usability.valb.result_semantics.v1"
	productionUsabilityValBCommandCenterSchema   = "point4.production_usability.valb.command_center_tasks.v1"
	productionUsabilityValBNoiseBudgetSchema     = "point4.production_usability.valb.noise_budget.v1"
	productionUsabilityValBAPIProtectionSchema   = "point4.production_usability.valb.api_protection.v1"
	productionUsabilityValBCLIResilienceSchema   = "point4.production_usability.valb.cli_resilience.v1"
	productionUsabilityValBScaleEnvelopeSchema   = "point4.production_usability.valb.scale_envelope.v1"
	productionUsabilityValBActionModesSchema     = "point4.production_usability.valb.action_mode_enforcement.v1"
	productionUsabilityValBProofsSchema          = "point4.production_usability.valb.proofs.v1"
)

type productionUsabilityValBUIDataResponse struct {
	SchemaVersion string                            `json:"schema_version"`
	GeneratedAt   time.Time                         `json:"generated_at"`
	CurrentState  string                            `json:"current_state"`
	Model         operability.UIDataResilienceModel `json:"model"`
	RouteRefs     []string                          `json:"route_refs,omitempty"`
	Limitations   []string                          `json:"limitations,omitempty"`
}

type productionUsabilityValBWindowingResponse struct {
	SchemaVersion string                                `json:"schema_version"`
	GeneratedAt   time.Time                             `json:"generated_at"`
	CurrentState  string                                `json:"current_state"`
	Model         operability.VirtualDataWindowContract `json:"model"`
	RouteRefs     []string                              `json:"route_refs,omitempty"`
	Limitations   []string                              `json:"limitations,omitempty"`
}

type productionUsabilityValBResultSemanticsResponse struct {
	SchemaVersion string                           `json:"schema_version"`
	GeneratedAt   time.Time                        `json:"generated_at"`
	CurrentState  string                           `json:"current_state"`
	Model         operability.ResultSemanticsModel `json:"model"`
	RouteRefs     []string                         `json:"route_refs,omitempty"`
	Limitations   []string                         `json:"limitations,omitempty"`
}

type productionUsabilityValBCommandCenterResponse struct {
	SchemaVersion string                             `json:"schema_version"`
	GeneratedAt   time.Time                          `json:"generated_at"`
	CurrentState  string                             `json:"current_state"`
	Model         operability.CommandCenterTaskModel `json:"model"`
	RouteRefs     []string                           `json:"route_refs,omitempty"`
	Limitations   []string                           `json:"limitations,omitempty"`
}

type productionUsabilityValBNoiseBudgetResponse struct {
	SchemaVersion string                       `json:"schema_version"`
	GeneratedAt   time.Time                    `json:"generated_at"`
	CurrentState  string                       `json:"current_state"`
	Model         operability.NoiseBudgetModel `json:"model"`
	RouteRefs     []string                     `json:"route_refs,omitempty"`
	Limitations   []string                     `json:"limitations,omitempty"`
}

type productionUsabilityValBAPIProtectionResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.APIProtectionDiscipline `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type productionUsabilityValBCLIResilienceResponse struct {
	SchemaVersion string                           `json:"schema_version"`
	GeneratedAt   time.Time                        `json:"generated_at"`
	CurrentState  string                           `json:"current_state"`
	Model         operability.CLIResilienceSurface `json:"model"`
	RouteRefs     []string                         `json:"route_refs,omitempty"`
	Limitations   []string                         `json:"limitations,omitempty"`
}

type productionUsabilityValBScaleEnvelopeResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.ProductionScaleEnvelope `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type productionUsabilityValBActionModeResponse struct {
	SchemaVersion string                                 `json:"schema_version"`
	GeneratedAt   time.Time                              `json:"generated_at"`
	CurrentState  string                                 `json:"current_state"`
	Model         operability.ActionModeEnforcementModel `json:"model"`
	RouteRefs     []string                               `json:"route_refs,omitempty"`
	Limitations   []string                               `json:"limitations,omitempty"`
}

type productionUsabilityValBProofsResponse struct {
	SchemaVersion              string    `json:"schema_version"`
	GeneratedAt                time.Time `json:"generated_at"`
	CurrentState               string    `json:"current_state"`
	Val0DependencyState        string    `json:"val_0_dependency_state"`
	Val0FoundationState        string    `json:"val_0_foundation_state"`
	ValADependencyState        string    `json:"val_a_dependency_state"`
	ValACoreState              string    `json:"val_a_core_state"`
	ValBState                  string    `json:"val_b_state"`
	Point4State                string    `json:"point_4_state"`
	UIDataResilienceState      string    `json:"ui_data_resilience_state"`
	WindowingState             string    `json:"windowing_state"`
	ResultSemanticsState       string    `json:"result_semantics_state"`
	CommandCenterState         string    `json:"command_center_state"`
	NoiseBudgetState           string    `json:"noise_budget_state"`
	APIProtectionState         string    `json:"api_protection_state"`
	CLIResilienceState         string    `json:"cli_resilience_state"`
	ScaleEnvelopeState         string    `json:"scale_envelope_state"`
	ActionModeEnforcementState string    `json:"action_mode_enforcement_state"`
	WhyPoint4NotPass           []string  `json:"why_point_4_not_pass,omitempty"`
	SurfaceRefs                []string  `json:"surface_refs,omitempty"`
	EvidenceRefs               []string  `json:"evidence_refs,omitempty"`
	Limitations                []string  `json:"limitations,omitempty"`
	IntegrationSummary         []string  `json:"integration_summary,omitempty"`
}

func (s server) productionUsabilityValBUIDataHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValBUIData())
}

func (s server) productionUsabilityValBWindowingHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValBWindowing())
}

func (s server) productionUsabilityValBResultSemanticsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValBResultSemantics())
}

func (s server) productionUsabilityValBCommandCenterHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValBCommandCenter())
}

func (s server) productionUsabilityValBNoiseBudgetHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValBNoiseBudget())
}

func (s server) productionUsabilityValBAPIProtectionHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValBAPIProtection())
}

func (s server) productionUsabilityValBCLIResilienceHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValBCLIResilience())
}

func (s server) productionUsabilityValBScaleEnvelopeHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValBScaleEnvelope())
}

func (s server) productionUsabilityValBActionModeHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValBActionModes())
}

func (s server) productionUsabilityValBProofsHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildProductionUsabilityValBProofs(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildProductionUsabilityValBUIData() productionUsabilityValBUIDataResponse {
	model := operability.ProductionUsabilityValBUIDataResilience()
	return productionUsabilityValBUIDataResponse{
		SchemaVersion: productionUsabilityValBUIDataSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValBUIDataResilienceState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valb/result-semantics",
			"/v1/production/usability-operability-recovery/valb/windowing",
			"/v1/production/usability-operability-recovery/valb/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValBWindowing() productionUsabilityValBWindowingResponse {
	model := operability.ProductionUsabilityValBWindowing()
	return productionUsabilityValBWindowingResponse{
		SchemaVersion: productionUsabilityValBWindowingSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValBWindowingState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valb/ui-data-resilience",
			"/v1/production/usability-operability-recovery/valb/result-semantics",
			"/v1/production/usability-operability-recovery/valb/proofs",
		},
		Limitations: []string{model.LimitationMessage},
	}
}

func buildProductionUsabilityValBResultSemantics() productionUsabilityValBResultSemanticsResponse {
	model := operability.ProductionUsabilityValBResultSemantics()
	return productionUsabilityValBResultSemanticsResponse{
		SchemaVersion: productionUsabilityValBResultSemanticsSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValBResultSemanticsState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valb/ui-data-resilience",
			"/v1/production/usability-operability-recovery/valb/windowing",
			"/v1/production/usability-operability-recovery/valb/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValBCommandCenter() productionUsabilityValBCommandCenterResponse {
	model := operability.ProductionUsabilityValBCommandCenterTasks()
	return productionUsabilityValBCommandCenterResponse{
		SchemaVersion: productionUsabilityValBCommandCenterSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValBCommandCenterState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valb/noise-budget",
			"/v1/production/usability-operability-recovery/valb/ui-data-resilience",
			"/v1/production/usability-operability-recovery/valb/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValBNoiseBudget() productionUsabilityValBNoiseBudgetResponse {
	model := operability.ProductionUsabilityValBNoiseBudget()
	return productionUsabilityValBNoiseBudgetResponse{
		SchemaVersion: productionUsabilityValBNoiseBudgetSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValBNoiseBudgetState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valb/command-center-tasks",
			"/v1/production/usability-operability-recovery/valb/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValBAPIProtection() productionUsabilityValBAPIProtectionResponse {
	model := operability.ProductionUsabilityValBAPIProtection()
	return productionUsabilityValBAPIProtectionResponse{
		SchemaVersion: productionUsabilityValBAPIProtectionSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValBAPIProtectionState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valb/cli-resilience",
			"/v1/production/usability-operability-recovery/valb/action-mode-enforcement",
			"/v1/production/usability-operability-recovery/valb/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValBCLIResilience() productionUsabilityValBCLIResilienceResponse {
	model := operability.ProductionUsabilityValBCLIResilience()
	return productionUsabilityValBCLIResilienceResponse{
		SchemaVersion: productionUsabilityValBCLIResilienceSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValBCLIResilienceState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valb/api-protection",
			"/v1/production/usability-operability-recovery/valb/action-mode-enforcement",
			"/v1/production/usability-operability-recovery/valb/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValBScaleEnvelope() productionUsabilityValBScaleEnvelopeResponse {
	model := operability.ProductionUsabilityValBScaleEnvelope()
	return productionUsabilityValBScaleEnvelopeResponse{
		SchemaVersion: productionUsabilityValBScaleEnvelopeSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValBScaleEnvelopeState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valb/windowing",
			"/v1/production/usability-operability-recovery/valb/api-protection",
			"/v1/production/usability-operability-recovery/valb/proofs",
		},
		Limitations: []string{model.LimitationDisclaimer},
	}
}

func buildProductionUsabilityValBActionModes() productionUsabilityValBActionModeResponse {
	model := operability.ProductionUsabilityValBActionModeEnforcement()
	return productionUsabilityValBActionModeResponse{
		SchemaVersion: productionUsabilityValBActionModesSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValBActionModeEnforcementState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/valb/api-protection",
			"/v1/production/usability-operability-recovery/valb/cli-resilience",
			"/v1/production/usability-operability-recovery/valb/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValBProofsCurrentState(val0State, valAState string, uiData operability.UIDataResilienceModel, windowing operability.VirtualDataWindowContract, resultSemantics operability.ResultSemanticsModel, commandCenter operability.CommandCenterTaskModel, noiseBudget operability.NoiseBudgetModel, apiProtection operability.APIProtectionDiscipline, cliResilience operability.CLIResilienceSurface, scaleEnvelope operability.ProductionScaleEnvelope, actionModes operability.ActionModeEnforcementModel, surfaceRefs, evidenceRefs, limitations, whyPoint4NotPass []string) string {
	return operability.EvaluateProductionUsabilityValBProofsState(
		val0State,
		valAState,
		operability.EvaluateProductionUsabilityValBUIDataResilienceState(uiData),
		operability.EvaluateProductionUsabilityValBWindowingState(windowing),
		operability.EvaluateProductionUsabilityValBResultSemanticsState(resultSemantics),
		operability.EvaluateProductionUsabilityValBCommandCenterState(commandCenter),
		operability.EvaluateProductionUsabilityValBNoiseBudgetState(noiseBudget),
		operability.EvaluateProductionUsabilityValBAPIProtectionState(apiProtection),
		operability.EvaluateProductionUsabilityValBCLIResilienceState(cliResilience),
		operability.EvaluateProductionUsabilityValBScaleEnvelopeState(scaleEnvelope),
		operability.EvaluateProductionUsabilityValBActionModeEnforcementState(actionModes),
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint4NotPass,
	)
}

func (s server) buildProductionUsabilityValBProofs(ctx context.Context, filter phase4EnterpriseFilter) (productionUsabilityValBProofsResponse, error) {
	val0, err := s.buildProductionUsabilityVal0Proofs(ctx, filter)
	if err != nil {
		return productionUsabilityValBProofsResponse{}, err
	}
	valA, err := s.buildProductionUsabilityValAProofs(ctx, filter)
	if err != nil {
		return productionUsabilityValBProofsResponse{}, err
	}

	uiData := operability.ProductionUsabilityValBUIDataResilience()
	windowing := operability.ProductionUsabilityValBWindowing()
	resultSemantics := operability.ProductionUsabilityValBResultSemantics()
	commandCenter := operability.ProductionUsabilityValBCommandCenterTasks()
	noiseBudget := operability.ProductionUsabilityValBNoiseBudget()
	apiProtection := operability.ProductionUsabilityValBAPIProtection()
	cliResilience := operability.ProductionUsabilityValBCLIResilience()
	scaleEnvelope := operability.ProductionUsabilityValBScaleEnvelope()
	actionModes := operability.ProductionUsabilityValBActionModeEnforcement()

	whyPoint4NotPass := []string{
		"Točka 4 remains not complete because later waves still need support bundle and diagnostics lifecycle work, final usability gate review, and integrated closure.",
		"Val B proves UI/API/CLI resilience readiness only; it does not prove support operations, install or upgrade lifecycle execution, or final usability closure.",
	}
	surfaceRefs := []string{
		"/v1/production/usability-operability-recovery/val0/proofs",
		"/v1/production/usability-operability-recovery/vala/proofs",
		"/v1/production/usability-operability-recovery/valb/ui-data-resilience",
		"/v1/production/usability-operability-recovery/valb/windowing",
		"/v1/production/usability-operability-recovery/valb/result-semantics",
		"/v1/production/usability-operability-recovery/valb/command-center-tasks",
		"/v1/production/usability-operability-recovery/valb/noise-budget",
		"/v1/production/usability-operability-recovery/valb/api-protection",
		"/v1/production/usability-operability-recovery/valb/cli-resilience",
		"/v1/production/usability-operability-recovery/valb/scale-envelope",
		"/v1/production/usability-operability-recovery/valb/action-mode-enforcement",
		"/v1/production/usability-operability-recovery/valb/proofs",
	}
	evidenceRefs := []string{
		"val0_proofs",
		"vala_proofs",
		"ui_data_resilience_projection_contract",
		"windowing_contract",
		"result_semantics_contract",
		"command_center_task_contract",
		"noise_budget_contract",
		"api_protection_contract",
		"cli_resilience_contract",
		"scale_envelope_contract",
		"action_mode_enforcement_contract",
	}
	limitations := []string{
		"Val B proves bounded resilience contracts only and does not claim final usability or support lifecycle completion.",
		"UI, CLI, API, and windowing outputs remain projection-only and never replace canonical truth or governed mutation authority.",
	}
	currentState := buildProductionUsabilityValBProofsCurrentState(
		val0.Val0State,
		valA.ValAState,
		uiData,
		windowing,
		resultSemantics,
		commandCenter,
		noiseBudget,
		apiProtection,
		cliResilience,
		scaleEnvelope,
		actionModes,
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint4NotPass,
	)
	valBState := operability.EvaluateProductionUsabilityValBState(
		val0.Val0State,
		valA.ValAState,
		operability.EvaluateProductionUsabilityValBUIDataResilienceState(uiData),
		operability.EvaluateProductionUsabilityValBWindowingState(windowing),
		operability.EvaluateProductionUsabilityValBResultSemanticsState(resultSemantics),
		operability.EvaluateProductionUsabilityValBCommandCenterState(commandCenter),
		operability.EvaluateProductionUsabilityValBNoiseBudgetState(noiseBudget),
		operability.EvaluateProductionUsabilityValBAPIProtectionState(apiProtection),
		operability.EvaluateProductionUsabilityValBCLIResilienceState(cliResilience),
		operability.EvaluateProductionUsabilityValBScaleEnvelopeState(scaleEnvelope),
		operability.EvaluateProductionUsabilityValBActionModeEnforcementState(actionModes),
	)

	return productionUsabilityValBProofsResponse{
		SchemaVersion:              productionUsabilityValBProofsSchema,
		GeneratedAt:                publicSampleTime(),
		CurrentState:               currentState,
		Val0DependencyState:        val0.CurrentState,
		Val0FoundationState:        val0.Val0State,
		ValADependencyState:        valA.CurrentState,
		ValACoreState:              valA.ValAState,
		ValBState:                  valBState,
		Point4State:                operability.ProductionUsabilityPoint4StateNotComplete,
		UIDataResilienceState:      operability.EvaluateProductionUsabilityValBUIDataResilienceState(uiData),
		WindowingState:             operability.EvaluateProductionUsabilityValBWindowingState(windowing),
		ResultSemanticsState:       operability.EvaluateProductionUsabilityValBResultSemanticsState(resultSemantics),
		CommandCenterState:         operability.EvaluateProductionUsabilityValBCommandCenterState(commandCenter),
		NoiseBudgetState:           operability.EvaluateProductionUsabilityValBNoiseBudgetState(noiseBudget),
		APIProtectionState:         operability.EvaluateProductionUsabilityValBAPIProtectionState(apiProtection),
		CLIResilienceState:         operability.EvaluateProductionUsabilityValBCLIResilienceState(cliResilience),
		ScaleEnvelopeState:         operability.EvaluateProductionUsabilityValBScaleEnvelopeState(scaleEnvelope),
		ActionModeEnforcementState: operability.EvaluateProductionUsabilityValBActionModeEnforcementState(actionModes),
		WhyPoint4NotPass:           whyPoint4NotPass,
		SurfaceRefs:                surfaceRefs,
		EvidenceRefs:               evidenceRefs,
		Limitations:                limitations,
		IntegrationSummary: []string{
			"Val B adds bounded UI, API, and CLI resilience contracts over the active Val 0 and Val A production usability spine.",
			"Val B keeps projection health, windowing, command-center tasks, API protection, CLI retry behavior, scale envelope, and action-mode enforcement explicit and fail-closed.",
			"Point 4 remains not complete because support lifecycle flows, final usability gate review, and integrated closure remain later work.",
		},
	}, nil
}
