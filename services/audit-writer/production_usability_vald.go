package main

import (
	"context"
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	productionUsabilityValDConfigReviewSchema          = "point4.production_usability.vald.config_review.v1"
	productionUsabilityValDExplainabilityReviewSchema  = "point4.production_usability.vald.explainability_review.v1"
	productionUsabilityValDDryRunReviewSchema          = "point4.production_usability.vald.dry_run_review.v1"
	productionUsabilityValDRedactionReviewSchema       = "point4.production_usability.vald.redaction_review.v1"
	productionUsabilityValDDegradedReviewSchema        = "point4.production_usability.vald.degraded_state_review.v1"
	productionUsabilityValDUIWindowingReviewSchema     = "point4.production_usability.vald.ui_windowing_review.v1"
	productionUsabilityValDCommandNoiseReviewSchema    = "point4.production_usability.vald.command_noise_review.v1"
	productionUsabilityValDAPIProtectionReviewSchema   = "point4.production_usability.vald.api_protection_review.v1"
	productionUsabilityValDCLIResilienceReviewSchema   = "point4.production_usability.vald.cli_resilience_review.v1"
	productionUsabilityValDSupportabilityReviewSchema  = "point4.production_usability.vald.supportability_review.v1"
	productionUsabilityValDRecoveryReviewSchema        = "point4.production_usability.vald.recovery_review.v1"
	productionUsabilityValDUpgradeRollbackReviewSchema = "point4.production_usability.vald.upgrade_rollback_review.v1"
	productionUsabilityValDScaleEnvelopeReviewSchema   = "point4.production_usability.vald.scale_envelope_review.v1"
	productionUsabilityValDGovernanceReviewSchema      = "point4.production_usability.vald.governance_boundary_review.v1"
	productionUsabilityValDRegressionGateSchema        = "point4.production_usability.vald.regression_gate.v1"
	productionUsabilityValDProofsSchema                = "point4.production_usability.vald.proofs.v1"
)

type productionUsabilityValDConfigReviewResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.ConfigCorrectnessReview `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type productionUsabilityValDExplainabilityReviewResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	GeneratedAt   time.Time                               `json:"generated_at"`
	CurrentState  string                                  `json:"current_state"`
	Model         operability.ExplainabilityClarityReview `json:"model"`
	RouteRefs     []string                                `json:"route_refs,omitempty"`
	Limitations   []string                                `json:"limitations,omitempty"`
}

type productionUsabilityValDDryRunReviewResponse struct {
	SchemaVersion string                                   `json:"schema_version"`
	GeneratedAt   time.Time                                `json:"generated_at"`
	CurrentState  string                                   `json:"current_state"`
	Model         operability.DryRunAuditCorrectnessReview `json:"model"`
	RouteRefs     []string                                 `json:"route_refs,omitempty"`
	Limitations   []string                                 `json:"limitations,omitempty"`
}

type productionUsabilityValDRedactionReviewResponse struct {
	SchemaVersion string                                `json:"schema_version"`
	GeneratedAt   time.Time                             `json:"generated_at"`
	CurrentState  string                                `json:"current_state"`
	Model         operability.PermissionRedactionReview `json:"model"`
	RouteRefs     []string                              `json:"route_refs,omitempty"`
	Limitations   []string                              `json:"limitations,omitempty"`
}

type productionUsabilityValDDegradedReviewResponse struct {
	SchemaVersion string                             `json:"schema_version"`
	GeneratedAt   time.Time                          `json:"generated_at"`
	CurrentState  string                             `json:"current_state"`
	Model         operability.DegradedBehaviorReview `json:"model"`
	RouteRefs     []string                           `json:"route_refs,omitempty"`
	Limitations   []string                           `json:"limitations,omitempty"`
}

type productionUsabilityValDUIWindowingReviewResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.UIWindowingResultReview `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type productionUsabilityValDCommandNoiseReviewResponse struct {
	SchemaVersion string                         `json:"schema_version"`
	GeneratedAt   time.Time                      `json:"generated_at"`
	CurrentState  string                         `json:"current_state"`
	Model         operability.CommandNoiseReview `json:"model"`
	RouteRefs     []string                       `json:"route_refs,omitempty"`
	Limitations   []string                       `json:"limitations,omitempty"`
}

type productionUsabilityValDAPIProtectionReviewResponse struct {
	SchemaVersion string                          `json:"schema_version"`
	GeneratedAt   time.Time                       `json:"generated_at"`
	CurrentState  string                          `json:"current_state"`
	Model         operability.APIProtectionReview `json:"model"`
	RouteRefs     []string                        `json:"route_refs,omitempty"`
	Limitations   []string                        `json:"limitations,omitempty"`
}

type productionUsabilityValDCLIResilienceReviewResponse struct {
	SchemaVersion string                          `json:"schema_version"`
	GeneratedAt   time.Time                       `json:"generated_at"`
	CurrentState  string                          `json:"current_state"`
	Model         operability.CLIResilienceReview `json:"model"`
	RouteRefs     []string                        `json:"route_refs,omitempty"`
	Limitations   []string                        `json:"limitations,omitempty"`
}

type productionUsabilityValDSupportabilityReviewResponse struct {
	SchemaVersion string                           `json:"schema_version"`
	GeneratedAt   time.Time                        `json:"generated_at"`
	CurrentState  string                           `json:"current_state"`
	Model         operability.SupportabilityReview `json:"model"`
	RouteRefs     []string                         `json:"route_refs,omitempty"`
	Limitations   []string                         `json:"limitations,omitempty"`
}

type productionUsabilityValDRecoveryReviewResponse struct {
	SchemaVersion string                       `json:"schema_version"`
	GeneratedAt   time.Time                    `json:"generated_at"`
	CurrentState  string                       `json:"current_state"`
	Model         operability.RecoveryUXReview `json:"model"`
	RouteRefs     []string                     `json:"route_refs,omitempty"`
	Limitations   []string                     `json:"limitations,omitempty"`
}

type productionUsabilityValDUpgradeRollbackReviewResponse struct {
	SchemaVersion string                            `json:"schema_version"`
	GeneratedAt   time.Time                         `json:"generated_at"`
	CurrentState  string                            `json:"current_state"`
	Model         operability.UpgradeRollbackReview `json:"model"`
	RouteRefs     []string                          `json:"route_refs,omitempty"`
	Limitations   []string                          `json:"limitations,omitempty"`
}

type productionUsabilityValDScaleEnvelopeReviewResponse struct {
	SchemaVersion string                          `json:"schema_version"`
	GeneratedAt   time.Time                       `json:"generated_at"`
	CurrentState  string                          `json:"current_state"`
	Model         operability.ScaleEnvelopeReview `json:"model"`
	RouteRefs     []string                        `json:"route_refs,omitempty"`
	Limitations   []string                        `json:"limitations,omitempty"`
}

type productionUsabilityValDGovernanceReviewResponse struct {
	SchemaVersion string                               `json:"schema_version"`
	GeneratedAt   time.Time                            `json:"generated_at"`
	CurrentState  string                               `json:"current_state"`
	Model         operability.GovernanceBoundaryReview `json:"model"`
	RouteRefs     []string                             `json:"route_refs,omitempty"`
	Limitations   []string                             `json:"limitations,omitempty"`
}

type productionUsabilityValDRegressionGateResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.UsabilityRegressionGate `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type productionUsabilityValDProofsResponse struct {
	SchemaVersion                 string    `json:"schema_version"`
	GeneratedAt                   time.Time `json:"generated_at"`
	CurrentState                  string    `json:"current_state"`
	Val0DependencyState           string    `json:"val_0_dependency_state"`
	Val0FoundationState           string    `json:"val_0_foundation_state"`
	ValADependencyState           string    `json:"val_a_dependency_state"`
	ValACoreState                 string    `json:"val_a_core_state"`
	ValBDependencyState           string    `json:"val_b_dependency_state"`
	ValBResilienceState           string    `json:"val_b_resilience_state"`
	ValCDependencyState           string    `json:"val_c_dependency_state"`
	ValCSupportabilityState       string    `json:"val_c_supportability_state"`
	ValDState                     string    `json:"val_d_state"`
	Point4State                   string    `json:"point_4_state"`
	ConfigReviewState             string    `json:"config_review_state"`
	ExplainabilityReviewState     string    `json:"explainability_review_state"`
	DryRunReviewState             string    `json:"dry_run_review_state"`
	RedactionReviewState          string    `json:"redaction_review_state"`
	DegradedBehaviorReviewState   string    `json:"degraded_behavior_review_state"`
	UIWindowingReviewState        string    `json:"ui_windowing_review_state"`
	CommandNoiseReviewState       string    `json:"command_noise_review_state"`
	APIProtectionReviewState      string    `json:"api_protection_review_state"`
	CLIResilienceReviewState      string    `json:"cli_resilience_review_state"`
	SupportabilityReviewState     string    `json:"supportability_review_state"`
	RecoveryReviewState           string    `json:"recovery_review_state"`
	UpgradeRollbackReviewState    string    `json:"upgrade_rollback_review_state"`
	ScaleEnvelopeReviewState      string    `json:"scale_envelope_review_state"`
	GovernanceBoundaryReviewState string    `json:"governance_boundary_review_state"`
	RegressionGateState           string    `json:"regression_gate_state"`
	WhyPoint4NotPass              []string  `json:"why_point_4_not_pass,omitempty"`
	SurfaceRefs                   []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                  []string  `json:"evidence_refs,omitempty"`
	Limitations                   []string  `json:"limitations,omitempty"`
	IntegrationSummary            []string  `json:"integration_summary,omitempty"`
}

func (s server) productionUsabilityValDConfigReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDConfigReview())
}

func (s server) productionUsabilityValDExplainabilityReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDExplainabilityReview())
}

func (s server) productionUsabilityValDDryRunReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDDryRunReview())
}

func (s server) productionUsabilityValDRedactionReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDRedactionReview())
}

func (s server) productionUsabilityValDDegradedReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDDegradedReview())
}

func (s server) productionUsabilityValDUIWindowingReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDUIWindowingReview())
}

func (s server) productionUsabilityValDCommandNoiseReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDCommandNoiseReview())
}

func (s server) productionUsabilityValDAPIProtectionReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDAPIProtectionReview())
}

func (s server) productionUsabilityValDCLIResilienceReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDCLIResilienceReview())
}

func (s server) productionUsabilityValDSupportabilityReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDSupportabilityReview())
}

func (s server) productionUsabilityValDRecoveryReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDRecoveryReview())
}

func (s server) productionUsabilityValDUpgradeRollbackReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDUpgradeRollbackReview())
}

func (s server) productionUsabilityValDScaleEnvelopeReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDScaleEnvelopeReview())
}

func (s server) productionUsabilityValDGovernanceReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDGovernanceReview())
}

func (s server) productionUsabilityValDRegressionGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildProductionUsabilityValDRegressionGate())
}

func (s server) productionUsabilityValDProofsHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildProductionUsabilityValDProofs(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildProductionUsabilityValDConfigReview() productionUsabilityValDConfigReviewResponse {
	model := operability.ProductionUsabilityValDConfigCorrectnessReview()
	return productionUsabilityValDConfigReviewResponse{
		SchemaVersion: productionUsabilityValDConfigReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDConfigReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/explainability-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDExplainabilityReview() productionUsabilityValDExplainabilityReviewResponse {
	model := operability.ProductionUsabilityValDExplainabilityClarityReview()
	return productionUsabilityValDExplainabilityReviewResponse{
		SchemaVersion: productionUsabilityValDExplainabilityReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDExplainabilityReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/redaction-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDDryRunReview() productionUsabilityValDDryRunReviewResponse {
	model := operability.ProductionUsabilityValDDryRunAuditReview()
	return productionUsabilityValDDryRunReviewResponse{
		SchemaVersion: productionUsabilityValDDryRunReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDDryRunReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/explainability-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDRedactionReview() productionUsabilityValDRedactionReviewResponse {
	model := operability.ProductionUsabilityValDPermissionRedactionReview()
	return productionUsabilityValDRedactionReviewResponse{
		SchemaVersion: productionUsabilityValDRedactionReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDRedactionReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/supportability-review",
			"/v1/production/usability-operability-recovery/vald/governance-boundary-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDDegradedReview() productionUsabilityValDDegradedReviewResponse {
	model := operability.ProductionUsabilityValDDegradedBehaviorReview()
	return productionUsabilityValDDegradedReviewResponse{
		SchemaVersion: productionUsabilityValDDegradedReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDDegradedBehaviorReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/ui-windowing-review",
			"/v1/production/usability-operability-recovery/vald/supportability-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDUIWindowingReview() productionUsabilityValDUIWindowingReviewResponse {
	model := operability.ProductionUsabilityValDUIWindowingResultReview()
	return productionUsabilityValDUIWindowingReviewResponse{
		SchemaVersion: productionUsabilityValDUIWindowingReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDUIWindowingReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/degraded-state-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDCommandNoiseReview() productionUsabilityValDCommandNoiseReviewResponse {
	model := operability.ProductionUsabilityValDCommandNoiseReview()
	return productionUsabilityValDCommandNoiseReviewResponse{
		SchemaVersion: productionUsabilityValDCommandNoiseReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDCommandNoiseReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/api-protection-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDAPIProtectionReview() productionUsabilityValDAPIProtectionReviewResponse {
	model := operability.ProductionUsabilityValDAPIProtectionReview()
	return productionUsabilityValDAPIProtectionReviewResponse{
		SchemaVersion: productionUsabilityValDAPIProtectionReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDAPIProtectionReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/cli-resilience-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDCLIResilienceReview() productionUsabilityValDCLIResilienceReviewResponse {
	model := operability.ProductionUsabilityValDCLIResilienceReview()
	return productionUsabilityValDCLIResilienceReviewResponse{
		SchemaVersion: productionUsabilityValDCLIResilienceReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDCLIResilienceReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/recovery-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDSupportabilityReview() productionUsabilityValDSupportabilityReviewResponse {
	model := operability.ProductionUsabilityValDSupportabilityReview()
	return productionUsabilityValDSupportabilityReviewResponse{
		SchemaVersion: productionUsabilityValDSupportabilityReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDSupportabilityReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/recovery-review",
			"/v1/production/usability-operability-recovery/vald/governance-boundary-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDRecoveryReview() productionUsabilityValDRecoveryReviewResponse {
	model := operability.ProductionUsabilityValDRecoveryUXReview()
	return productionUsabilityValDRecoveryReviewResponse{
		SchemaVersion: productionUsabilityValDRecoveryReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDRecoveryReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/upgrade-rollback-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDUpgradeRollbackReview() productionUsabilityValDUpgradeRollbackReviewResponse {
	model := operability.ProductionUsabilityValDUpgradeRollbackReview()
	return productionUsabilityValDUpgradeRollbackReviewResponse{
		SchemaVersion: productionUsabilityValDUpgradeRollbackReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDUpgradeRollbackReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/supportability-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDScaleEnvelopeReview() productionUsabilityValDScaleEnvelopeReviewResponse {
	model := operability.ProductionUsabilityValDScaleEnvelopeReview()
	return productionUsabilityValDScaleEnvelopeReviewResponse{
		SchemaVersion: productionUsabilityValDScaleEnvelopeReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDScaleEnvelopeReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/governance-boundary-review",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDGovernanceReview() productionUsabilityValDGovernanceReviewResponse {
	model := operability.ProductionUsabilityValDGovernanceBoundaryReview()
	return productionUsabilityValDGovernanceReviewResponse{
		SchemaVersion: productionUsabilityValDGovernanceReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDGovernanceBoundaryReviewState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/regression-gate",
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDRegressionGate() productionUsabilityValDRegressionGateResponse {
	model := operability.ProductionUsabilityValDUsabilityRegressionGate()
	return productionUsabilityValDRegressionGateResponse{
		SchemaVersion: productionUsabilityValDRegressionGateSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateProductionUsabilityValDRegressionGateState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/production/usability-operability-recovery/vald/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildProductionUsabilityValDProofsCurrentState(val0State, valAState, valBState, valCState string, configReview operability.ConfigCorrectnessReview, explainReview operability.ExplainabilityClarityReview, dryRunReview operability.DryRunAuditCorrectnessReview, redactionReview operability.PermissionRedactionReview, degradedReview operability.DegradedBehaviorReview, uiWindowingReview operability.UIWindowingResultReview, commandNoiseReview operability.CommandNoiseReview, apiReview operability.APIProtectionReview, cliReview operability.CLIResilienceReview, supportabilityReview operability.SupportabilityReview, recoveryReview operability.RecoveryUXReview, upgradeReview operability.UpgradeRollbackReview, scaleReview operability.ScaleEnvelopeReview, governanceReview operability.GovernanceBoundaryReview, regressionGate operability.UsabilityRegressionGate, surfaceRefs, evidenceRefs, limitations, whyPoint4NotPass []string) string {
	return operability.EvaluateProductionUsabilityValDProofsState(
		val0State,
		valAState,
		valBState,
		valCState,
		operability.EvaluateProductionUsabilityValDConfigReviewState(configReview),
		operability.EvaluateProductionUsabilityValDExplainabilityReviewState(explainReview),
		operability.EvaluateProductionUsabilityValDDryRunReviewState(dryRunReview),
		operability.EvaluateProductionUsabilityValDRedactionReviewState(redactionReview),
		operability.EvaluateProductionUsabilityValDDegradedBehaviorReviewState(degradedReview),
		operability.EvaluateProductionUsabilityValDUIWindowingReviewState(uiWindowingReview),
		operability.EvaluateProductionUsabilityValDCommandNoiseReviewState(commandNoiseReview),
		operability.EvaluateProductionUsabilityValDAPIProtectionReviewState(apiReview),
		operability.EvaluateProductionUsabilityValDCLIResilienceReviewState(cliReview),
		operability.EvaluateProductionUsabilityValDSupportabilityReviewState(supportabilityReview),
		operability.EvaluateProductionUsabilityValDRecoveryReviewState(recoveryReview),
		operability.EvaluateProductionUsabilityValDUpgradeRollbackReviewState(upgradeReview),
		operability.EvaluateProductionUsabilityValDScaleEnvelopeReviewState(scaleReview),
		operability.EvaluateProductionUsabilityValDGovernanceBoundaryReviewState(governanceReview),
		operability.EvaluateProductionUsabilityValDRegressionGateState(regressionGate),
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint4NotPass,
	)
}

func (s server) buildProductionUsabilityValDProofs(ctx context.Context, filter phase4EnterpriseFilter) (productionUsabilityValDProofsResponse, error) {
	val0, err := s.buildProductionUsabilityVal0Proofs(ctx, filter)
	if err != nil {
		return productionUsabilityValDProofsResponse{}, err
	}
	valA, err := s.buildProductionUsabilityValAProofs(ctx, filter)
	if err != nil {
		return productionUsabilityValDProofsResponse{}, err
	}
	valB, err := s.buildProductionUsabilityValBProofs(ctx, filter)
	if err != nil {
		return productionUsabilityValDProofsResponse{}, err
	}
	valC, err := s.buildProductionUsabilityValCProofs(ctx, filter)
	if err != nil {
		return productionUsabilityValDProofsResponse{}, err
	}

	configReview := operability.ProductionUsabilityValDConfigCorrectnessReview()
	explainReview := operability.ProductionUsabilityValDExplainabilityClarityReview()
	dryRunReview := operability.ProductionUsabilityValDDryRunAuditReview()
	redactionReview := operability.ProductionUsabilityValDPermissionRedactionReview()
	degradedReview := operability.ProductionUsabilityValDDegradedBehaviorReview()
	uiWindowingReview := operability.ProductionUsabilityValDUIWindowingResultReview()
	commandNoiseReview := operability.ProductionUsabilityValDCommandNoiseReview()
	apiReview := operability.ProductionUsabilityValDAPIProtectionReview()
	cliReview := operability.ProductionUsabilityValDCLIResilienceReview()
	supportabilityReview := operability.ProductionUsabilityValDSupportabilityReview()
	recoveryReview := operability.ProductionUsabilityValDRecoveryUXReview()
	upgradeReview := operability.ProductionUsabilityValDUpgradeRollbackReview()
	scaleReview := operability.ProductionUsabilityValDScaleEnvelopeReview()
	governanceReview := operability.ProductionUsabilityValDGovernanceBoundaryReview()
	regressionGate := operability.ProductionUsabilityValDUsabilityRegressionGate()

	whyPoint4NotPass := []string{
		"Točka 4 remains not complete because integrated closure is still deferred to Val E.",
		"Val D proves final usability gate readiness only; it does not prove integrated closure or final point completion.",
	}
	surfaceRefs := []string{
		"/v1/production/usability-operability-recovery/val0/proofs",
		"/v1/production/usability-operability-recovery/vala/proofs",
		"/v1/production/usability-operability-recovery/valb/proofs",
		"/v1/production/usability-operability-recovery/valc/proofs",
		"/v1/production/usability-operability-recovery/vald/config-review",
		"/v1/production/usability-operability-recovery/vald/explainability-review",
		"/v1/production/usability-operability-recovery/vald/dry-run-review",
		"/v1/production/usability-operability-recovery/vald/redaction-review",
		"/v1/production/usability-operability-recovery/vald/degraded-state-review",
		"/v1/production/usability-operability-recovery/vald/ui-windowing-review",
		"/v1/production/usability-operability-recovery/vald/command-noise-review",
		"/v1/production/usability-operability-recovery/vald/api-protection-review",
		"/v1/production/usability-operability-recovery/vald/cli-resilience-review",
		"/v1/production/usability-operability-recovery/vald/supportability-review",
		"/v1/production/usability-operability-recovery/vald/recovery-review",
		"/v1/production/usability-operability-recovery/vald/upgrade-rollback-review",
		"/v1/production/usability-operability-recovery/vald/scale-envelope-review",
		"/v1/production/usability-operability-recovery/vald/governance-boundary-review",
		"/v1/production/usability-operability-recovery/vald/regression-gate",
		"/v1/production/usability-operability-recovery/vald/proofs",
	}
	evidenceRefs := []string{
		"val0_proofs",
		"vala_proofs",
		"valb_proofs",
		"valc_proofs",
		"config_correctness_review_contract",
		"explainability_review_contract",
		"dry_run_review_contract",
		"redaction_review_contract",
		"degraded_behavior_review_contract",
		"ui_windowing_review_contract",
		"command_noise_review_contract",
		"api_protection_review_contract",
		"cli_resilience_review_contract",
		"supportability_review_contract",
		"recovery_review_contract",
		"upgrade_rollback_review_contract",
		"scale_review_contract",
		"governance_boundary_review_contract",
		"regression_gate_contract",
	}
	limitations := []string{
		"Val D proves final usability gate readiness only and does not claim integrated closure or final point completion.",
		"All gate, readiness, support, explainability, and advisory outputs remain projection-only and do not replace canonical truth.",
	}
	currentState := buildProductionUsabilityValDProofsCurrentState(
		val0.Val0State,
		valA.ValAState,
		valB.ValBState,
		valC.ValCState,
		configReview,
		explainReview,
		dryRunReview,
		redactionReview,
		degradedReview,
		uiWindowingReview,
		commandNoiseReview,
		apiReview,
		cliReview,
		supportabilityReview,
		recoveryReview,
		upgradeReview,
		scaleReview,
		governanceReview,
		regressionGate,
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint4NotPass,
	)
	valDState := operability.EvaluateProductionUsabilityValDState(
		val0.Val0State,
		valA.ValAState,
		valB.ValBState,
		valC.ValCState,
		operability.EvaluateProductionUsabilityValDConfigReviewState(configReview),
		operability.EvaluateProductionUsabilityValDExplainabilityReviewState(explainReview),
		operability.EvaluateProductionUsabilityValDDryRunReviewState(dryRunReview),
		operability.EvaluateProductionUsabilityValDRedactionReviewState(redactionReview),
		operability.EvaluateProductionUsabilityValDDegradedBehaviorReviewState(degradedReview),
		operability.EvaluateProductionUsabilityValDUIWindowingReviewState(uiWindowingReview),
		operability.EvaluateProductionUsabilityValDCommandNoiseReviewState(commandNoiseReview),
		operability.EvaluateProductionUsabilityValDAPIProtectionReviewState(apiReview),
		operability.EvaluateProductionUsabilityValDCLIResilienceReviewState(cliReview),
		operability.EvaluateProductionUsabilityValDSupportabilityReviewState(supportabilityReview),
		operability.EvaluateProductionUsabilityValDRecoveryReviewState(recoveryReview),
		operability.EvaluateProductionUsabilityValDUpgradeRollbackReviewState(upgradeReview),
		operability.EvaluateProductionUsabilityValDScaleEnvelopeReviewState(scaleReview),
		operability.EvaluateProductionUsabilityValDGovernanceBoundaryReviewState(governanceReview),
		operability.EvaluateProductionUsabilityValDRegressionGateState(regressionGate),
	)

	return productionUsabilityValDProofsResponse{
		SchemaVersion:                 productionUsabilityValDProofsSchema,
		GeneratedAt:                   publicSampleTime(),
		CurrentState:                  currentState,
		Val0DependencyState:           val0.CurrentState,
		Val0FoundationState:           val0.Val0State,
		ValADependencyState:           valA.CurrentState,
		ValACoreState:                 valA.ValAState,
		ValBDependencyState:           valB.CurrentState,
		ValBResilienceState:           valB.ValBState,
		ValCDependencyState:           valC.CurrentState,
		ValCSupportabilityState:       valC.ValCState,
		ValDState:                     valDState,
		Point4State:                   operability.ProductionUsabilityPoint4StateNotComplete,
		ConfigReviewState:             operability.EvaluateProductionUsabilityValDConfigReviewState(configReview),
		ExplainabilityReviewState:     operability.EvaluateProductionUsabilityValDExplainabilityReviewState(explainReview),
		DryRunReviewState:             operability.EvaluateProductionUsabilityValDDryRunReviewState(dryRunReview),
		RedactionReviewState:          operability.EvaluateProductionUsabilityValDRedactionReviewState(redactionReview),
		DegradedBehaviorReviewState:   operability.EvaluateProductionUsabilityValDDegradedBehaviorReviewState(degradedReview),
		UIWindowingReviewState:        operability.EvaluateProductionUsabilityValDUIWindowingReviewState(uiWindowingReview),
		CommandNoiseReviewState:       operability.EvaluateProductionUsabilityValDCommandNoiseReviewState(commandNoiseReview),
		APIProtectionReviewState:      operability.EvaluateProductionUsabilityValDAPIProtectionReviewState(apiReview),
		CLIResilienceReviewState:      operability.EvaluateProductionUsabilityValDCLIResilienceReviewState(cliReview),
		SupportabilityReviewState:     operability.EvaluateProductionUsabilityValDSupportabilityReviewState(supportabilityReview),
		RecoveryReviewState:           operability.EvaluateProductionUsabilityValDRecoveryReviewState(recoveryReview),
		UpgradeRollbackReviewState:    operability.EvaluateProductionUsabilityValDUpgradeRollbackReviewState(upgradeReview),
		ScaleEnvelopeReviewState:      operability.EvaluateProductionUsabilityValDScaleEnvelopeReviewState(scaleReview),
		GovernanceBoundaryReviewState: operability.EvaluateProductionUsabilityValDGovernanceBoundaryReviewState(governanceReview),
		RegressionGateState:           operability.EvaluateProductionUsabilityValDRegressionGateState(regressionGate),
		WhyPoint4NotPass:              whyPoint4NotPass,
		SurfaceRefs:                   surfaceRefs,
		EvidenceRefs:                  evidenceRefs,
		Limitations:                   limitations,
		IntegrationSummary: []string{
			"Val D reviews final usability correctness posture across the active Val 0, Val A, Val B, and Val C layers.",
			"Val D keeps config, explainability, resilience, supportability, and governance checks projection-only and fail-closed.",
			"Point 4 remains not complete because integrated closure is intentionally deferred to Val E.",
		},
	}, nil
}
