package benchmark

import (
	"fmt"
	"strings"
	"time"
)

const (
	FoundationHarnessSchemaVersion    = "1.execution_benchmark_harness.v1"
	FoundationEvaluationSchemaVersion = "1.execution_benchmark_evaluation.v1"
)

type FoundationProfile struct {
	ProfileID       string   `json:"profile_id"`
	DisplayName     string   `json:"display_name"`
	Characteristics []string `json:"characteristics,omitempty"`
}

type FoundationFamily struct {
	FamilyID      string   `json:"family_id"`
	Component     string   `json:"component"`
	CommandHint   string   `json:"command_hint"`
	MetricClasses []string `json:"metric_classes,omitempty"`
	ProfileIDs    []string `json:"profile_ids,omitempty"`
	CurrentStatus string   `json:"current_status"`
}

type RegressionRule struct {
	MetricClass      string  `json:"metric_class"`
	Direction        string  `json:"direction"`
	ThresholdPercent float64 `json:"threshold_percent"`
}

type FoundationHarness struct {
	SchemaVersion string              `json:"schema_version"`
	Profiles      []FoundationProfile `json:"profiles,omitempty"`
	Families      []FoundationFamily  `json:"families,omitempty"`
	Rules         []RegressionRule    `json:"rules,omitempty"`
	Limitations   []string            `json:"limitations,omitempty"`
}

type Observation struct {
	FamilyID      string  `json:"family_id"`
	ProfileID     string  `json:"profile_id"`
	MetricClass   string  `json:"metric_class"`
	MetricName    string  `json:"metric_name"`
	Unit          string  `json:"unit"`
	BaselineValue float64 `json:"baseline_value"`
	ObservedValue float64 `json:"observed_value"`
}

type Override struct {
	Reason     string `json:"reason"`
	ApprovedBy string `json:"approved_by,omitempty"`
	TicketRef  string `json:"ticket_ref,omitempty"`
}

type EvaluationRequest struct {
	SchemaVersion string        `json:"schema_version"`
	ProfileID     string        `json:"profile_id"`
	ObservedAt    time.Time     `json:"observed_at,omitempty"`
	Observations  []Observation `json:"observations,omitempty"`
	Override      *Override     `json:"override,omitempty"`
}

type EvaluationResult struct {
	FamilyID      string  `json:"family_id"`
	MetricClass   string  `json:"metric_class"`
	MetricName    string  `json:"metric_name"`
	Status        string  `json:"status"`
	BaselineValue float64 `json:"baseline_value"`
	ObservedValue float64 `json:"observed_value"`
	DeltaPercent  float64 `json:"delta_percent"`
	ThresholdPct  float64 `json:"threshold_pct"`
	Summary       string  `json:"summary"`
}

type EvaluationResponse struct {
	SchemaVersion  string             `json:"schema_version"`
	CurrentState   string             `json:"current_state"`
	ProfileID      string             `json:"profile_id"`
	ObservedAt     time.Time          `json:"observed_at"`
	OverrideReason string             `json:"override_reason,omitempty"`
	Results        []EvaluationResult `json:"results,omitempty"`
	Limitations    []string           `json:"limitations,omitempty"`
}

func FoundationCatalog() FoundationHarness {
	profiles := []FoundationProfile{
		{ProfileID: "local_baseline", DisplayName: "Local Baseline", Characteristics: []string{"fast feedback", "developer workstation", "lower concurrency"}},
		{ProfileID: "production_like", DisplayName: "Production-like", Characteristics: []string{"representative control-plane mix", "evidence-heavy workload patterns", "realistic concurrency"}},
		{ProfileID: "stress", DisplayName: "Stress", Characteristics: []string{"burst-heavy workload", "retry pressure", "high fan-out dispatch"}},
	}
	families := []FoundationFamily{
		{FamilyID: "deploy_gate_admission", Component: "deploy-gate", CommandHint: "go test ./services/deploy-gate -run '^$' -bench BenchmarkAdmissionReview -benchmem", MetricClasses: []string{"user_facing_latency", "control_plane_latency", "memory_footprint"}, ProfileIDs: []string{"local_baseline", "production_like", "stress"}, CurrentStatus: "ready_for_regression_evaluation"},
		{FamilyID: "policy_evaluation", Component: "internal/policy", CommandHint: "go test ./internal/policy -run '^$' -bench 'BenchmarkEvaluate(Change|Artifact)' -benchmem", MetricClasses: []string{"control_plane_latency", "memory_footprint"}, ProfileIDs: []string{"local_baseline", "production_like"}, CurrentStatus: "ready_for_regression_evaluation"},
		{FamilyID: "runtime_compare", Component: "internal/runtime", CommandHint: "go test ./internal/runtime -run '^$' -bench BenchmarkCompare -benchmem", MetricClasses: []string{"control_plane_latency", "memory_footprint"}, ProfileIDs: []string{"local_baseline", "production_like"}, CurrentStatus: "ready_for_regression_evaluation"},
		{FamilyID: "audit_ingest", Component: "internal/audit", CommandHint: "go test ./internal/audit -run '^$' -bench BenchmarkMemoryStoreIngest -benchmem", MetricClasses: []string{"evidence_latency", "throughput", "memory_footprint"}, ProfileIDs: []string{"local_baseline", "production_like", "stress"}, CurrentStatus: "ready_for_regression_evaluation"},
		{FamilyID: "audit_writer_read_paths", Component: "services/audit-writer", CommandHint: "go test ./services/audit-writer -run '^$' -bench 'BenchmarkAuditWriter(TopologyBlastRadius|ForensicsState|RuntimeFindings)' -benchmem", MetricClasses: []string{"background_completion_latency", "evidence_latency", "memory_footprint"}, ProfileIDs: []string{"local_baseline", "production_like", "stress"}, CurrentStatus: "ready_for_regression_evaluation"},
		{FamilyID: "audit_writer_mutation_paths", Component: "services/audit-writer", CommandHint: "go test ./services/audit-writer -run '^$' -bench 'BenchmarkAuditWriter(HandoffSeal|HandoffVerify|FederationProofVerify|ValidationExecute)' -benchmem", MetricClasses: []string{"background_completion_latency", "evidence_latency", "memory_footprint"}, ProfileIDs: []string{"local_baseline", "production_like", "stress"}, CurrentStatus: "ready_for_regression_evaluation"},
	}
	rules := []RegressionRule{
		{MetricClass: "user_facing_latency", Direction: "max_increase_pct", ThresholdPercent: 15},
		{MetricClass: "control_plane_latency", Direction: "max_increase_pct", ThresholdPercent: 20},
		{MetricClass: "evidence_latency", Direction: "max_increase_pct", ThresholdPercent: 25},
		{MetricClass: "background_completion_latency", Direction: "max_increase_pct", ThresholdPercent: 30},
		{MetricClass: "memory_footprint", Direction: "max_increase_pct", ThresholdPercent: 25},
		{MetricClass: "throughput", Direction: "max_decrease_pct", ThresholdPercent: 15},
	}
	return FoundationHarness{
		SchemaVersion: FoundationHarnessSchemaVersion,
		Profiles:      profiles,
		Families:      families,
		Rules:         rules,
		Limitations: []string{
			"Regression evaluation is profile-aware but still depends on stable synthetic fixtures and reproducible runner conditions.",
			"Current rules are engineering gates for Phase 1 and are not public performance claims.",
		},
	}
}

func EvaluateFoundationRegression(request EvaluationRequest) EvaluationResponse {
	observedAt := request.ObservedAt.UTC()
	if request.ObservedAt.IsZero() {
		observedAt = time.Now().UTC()
	}
	catalog := FoundationCatalog()
	ruleByClass := map[string]RegressionRule{}
	for _, rule := range catalog.Rules {
		ruleByClass[rule.MetricClass] = rule
	}

	results := make([]EvaluationResult, 0, len(request.Observations))
	failed := false
	for _, item := range request.Observations {
		rule, ok := ruleByClass[strings.TrimSpace(item.MetricClass)]
		if !ok || item.BaselineValue <= 0 {
			results = append(results, EvaluationResult{
				FamilyID:      item.FamilyID,
				MetricClass:   item.MetricClass,
				MetricName:    item.MetricName,
				Status:        "insufficient_baseline",
				BaselineValue: item.BaselineValue,
				ObservedValue: item.ObservedValue,
				Summary:       "No supported regression rule or usable baseline was supplied for this observation.",
			})
			failed = true
			continue
		}

		deltaPct := ((item.ObservedValue - item.BaselineValue) / item.BaselineValue) * 100
		status := "pass"
		switch rule.Direction {
		case "max_increase_pct":
			if deltaPct > rule.ThresholdPercent {
				status = "regression"
			}
		case "max_decrease_pct":
			if deltaPct < -rule.ThresholdPercent {
				status = "regression"
			}
		default:
			status = "unsupported_rule"
		}
		if status != "pass" {
			failed = true
		}
		results = append(results, EvaluationResult{
			FamilyID:      item.FamilyID,
			MetricClass:   item.MetricClass,
			MetricName:    item.MetricName,
			Status:        status,
			BaselineValue: item.BaselineValue,
			ObservedValue: item.ObservedValue,
			DeltaPercent:  deltaPct,
			ThresholdPct:  rule.ThresholdPercent,
			Summary:       evaluationSummary(item, rule, deltaPct, status),
		})
	}

	currentState := "passed"
	overrideReason := ""
	if failed {
		currentState = "failed"
		if request.Override != nil && strings.TrimSpace(request.Override.Reason) != "" {
			currentState = "passed_with_override"
			overrideReason = strings.TrimSpace(request.Override.Reason)
		}
	}

	return EvaluationResponse{
		SchemaVersion:  FoundationEvaluationSchemaVersion,
		CurrentState:   currentState,
		ProfileID:      strings.TrimSpace(request.ProfileID),
		ObservedAt:     observedAt,
		OverrideReason: overrideReason,
		Results:        results,
		Limitations: []string{
			"Benchmark gate evaluation compares observed metrics to the supplied baseline under the selected metric-class rules; it does not normalize hardware or runner drift automatically.",
		},
	}
}

func evaluationSummary(item Observation, rule RegressionRule, deltaPct float64, status string) string {
	switch status {
	case "pass":
		return fmt.Sprintf("%s stayed within the %.1f%% %s budget for %s.", item.MetricName, rule.ThresholdPercent, strings.ReplaceAll(rule.Direction, "_", " "), item.MetricClass)
	case "regression":
		return fmt.Sprintf("%s moved by %.2f%% against baseline %.4f%s and exceeded the %.1f%% %s budget.", item.MetricName, deltaPct, item.BaselineValue, item.Unit, rule.ThresholdPercent, strings.ReplaceAll(rule.Direction, "_", " "))
	default:
		return fmt.Sprintf("%s could not be evaluated because the supplied observation did not match a supported regression rule.", item.MetricName)
	}
}
