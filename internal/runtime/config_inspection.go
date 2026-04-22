package runtime

import (
	"os"
	"sort"
	"strings"
)

const SelfHealingInspectionSchemaVersion = "5.runtime_self_healing_config_inspection.v1"

type ConfigIssue struct {
	Severity string `json:"severity"`
	Code     string `json:"code"`
	Field    string `json:"field,omitempty"`
	Message  string `json:"message"`
}

type ConfigDefault struct {
	Field  string `json:"field"`
	Value  string `json:"value"`
	Source string `json:"source"`
}

type SelfHealingInspection struct {
	SchemaVersion   string            `json:"schema_version"`
	CurrentState    string            `json:"current_state"`
	EffectiveConfig SelfHealingConfig `json:"effective_config"`
	DeclaredValues  map[string]string `json:"declared_values,omitempty"`
	DefaultsApplied []ConfigDefault   `json:"defaults_applied,omitempty"`
	Issues          []ConfigIssue     `json:"issues,omitempty"`
	Limitations     []string          `json:"limitations,omitempty"`
}

func InspectSelfHealingConfig(getenv func(string) string) SelfHealingInspection {
	if getenv == nil {
		getenv = os.Getenv
	}
	report := SelfHealingInspection{
		SchemaVersion:  SelfHealingInspectionSchemaVersion,
		DeclaredValues: map[string]string{},
		Limitations: []string{
			"Runtime self-healing config inspection reflects local environment variables and effective defaults only; it does not claim current cluster-side rollout state.",
			"Compatibility warnings remain bounded to the declared env configuration and do not simulate live remediation execution.",
		},
	}

	recordDeclared := func(key string) {
		if value := strings.TrimSpace(getenv(key)); value != "" {
			report.DeclaredValues[key] = value
		}
	}
	for _, key := range []string{
		"CHANGELOCK_SELF_HEALING_MODE",
		"CHANGELOCK_SELF_HEALING_MAX_ATTEMPTS",
		"CHANGELOCK_SELF_HEALING_WINDOW",
		"CHANGELOCK_SELF_HEALING_ALLOWED_KINDS",
		"CHANGELOCK_CLOSED_LOOP_FAIL_MODE",
		"CHANGELOCK_CLOSED_LOOP_RECONCILE_INTERVAL",
		"CHANGELOCK_CLOSED_LOOP_PROTECTED_NAMESPACES",
		"CHANGELOCK_CLOSED_LOOP_PROTECTED_WORKLOADS",
		"CHANGELOCK_RUNTIME_VEX_QUARANTINE_SEVERITY",
		"CHANGELOCK_CLOSED_LOOP_REQUIRE_SIGNED_DESIRED_STATE",
		"CHANGELOCK_SELF_HEALING_REQUIRE_SIGNED_DESIRED_STATE",
		"CHANGELOCK_SELF_HEALING_CRITICAL_ONLY",
		"CHANGELOCK_CLOSED_LOOP_VERIFY_DESIRED_STATE_ON_RECONCILE",
		"CHANGELOCK_RUNTIME_QUARANTINE_NETWORK_POLICY_ENABLED",
		"CHANGELOCK_RUNTIME_VEX_QUARANTINE_ENABLED",
		"CHANGELOCK_RUNTIME_VEX_QUARANTINE_REQUIRE_NET_ACTIONABLE",
	} {
		recordDeclared(key)
	}

	cfg, err := LoadSelfHealingConfigFromEnv(getenv)
	if err != nil {
		report.CurrentState = "invalid"
		report.Issues = []ConfigIssue{{
			Severity: "error",
			Code:     "runtime_self_healing_invalid",
			Field:    "env",
			Message:  err.Error(),
		}}
		return report
	}
	report.EffectiveConfig = cfg
	report.DefaultsApplied = runtimeSelfHealingDefaults(getenv, cfg)

	if cfg.Mode != RemediationModeDisabled && !cfg.RequireSignedDesiredState {
		report.Issues = append(report.Issues, ConfigIssue{
			Severity: "warning",
			Code:     "signed_desired_state_relaxed",
			Field:    "CHANGELOCK_CLOSED_LOOP_REQUIRE_SIGNED_DESIRED_STATE",
			Message:  "closed-loop remediation is enabled without signed desired-state enforcement",
		})
	}
	if len(report.Issues) > 0 {
		report.CurrentState = "valid_with_warnings"
	} else if len(report.DefaultsApplied) > 0 {
		report.CurrentState = "valid_with_defaults"
	} else {
		report.CurrentState = "valid"
	}
	return report
}

func runtimeSelfHealingDefaults(getenv func(string) string, cfg SelfHealingConfig) []ConfigDefault {
	defaults := []ConfigDefault{}
	addDefault := func(envKey, field, value string) {
		if strings.TrimSpace(getenv(envKey)) == "" {
			defaults = append(defaults, ConfigDefault{
				Field:  field,
				Value:  value,
				Source: "default",
			})
		}
	}
	addDefault("CHANGELOCK_SELF_HEALING_MODE", "mode", string(cfg.Mode))
	addDefault("CHANGELOCK_SELF_HEALING_MAX_ATTEMPTS", "max_attempts", "3")
	addDefault("CHANGELOCK_SELF_HEALING_WINDOW", "window", "15m0s")
	addDefault("CHANGELOCK_SELF_HEALING_ALLOWED_KINDS", "allowed_kinds", "Deployment,DaemonSet,StatefulSet")
	addDefault("CHANGELOCK_CLOSED_LOOP_FAIL_MODE", "fail_mode", string(cfg.FailMode))
	addDefault("CHANGELOCK_CLOSED_LOOP_RECONCILE_INTERVAL", "reconcile_interval", cfg.ReconcileInterval.String())
	addDefault("CHANGELOCK_CLOSED_LOOP_PROTECTED_NAMESPACES", "protected_namespaces", "changelock,changelock-system")
	addDefault("CHANGELOCK_RUNTIME_VEX_QUARANTINE_SEVERITY", "runtime_vex_severity", cfg.RuntimeVEXSeverity)
	addDefault("CHANGELOCK_RUNTIME_VEX_QUARANTINE_REQUIRE_NET_ACTIONABLE", "runtime_vex_require_net", boolString(cfg.RuntimeVEXRequireNet))
	sort.Slice(defaults, func(i, j int) bool {
		return defaults[i].Field < defaults[j].Field
	})
	return defaults
}

func boolString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}
