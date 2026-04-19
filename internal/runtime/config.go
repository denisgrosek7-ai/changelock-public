package runtime

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

type SelfHealingConfig struct {
	Mode                      RemediationMode
	MaxAttempts               int
	Window                    time.Duration
	AllowedKinds              map[string]struct{}
	RequireSignedDesiredState bool
	CriticalOnly              bool
	FailMode                  RemediationMode
	VerifyDesiredStateOnLoop  bool
	ReconcileInterval         time.Duration
	ProtectedNamespaces       map[string]struct{}
	ProtectedWorkloads        map[string]struct{}
	QuarantineNetworkPolicy   bool
	RuntimeVEXQuarantine      bool
	RuntimeVEXSeverity        string
	RuntimeVEXRequireNet      bool
}

func LoadSelfHealingConfig() (SelfHealingConfig, error) {
	mode := RemediationMode(strings.ToLower(strings.TrimSpace(envOrDefault("CHANGELOCK_SELF_HEALING_MODE", string(RemediationModeDisabled)))))
	switch mode {
	case RemediationModeDisabled, RemediationModeAlertOnly, RemediationModeQuarantine, RemediationModePatchApprovedState, RemediationModeRestartApprovedState:
	default:
		return SelfHealingConfig{}, errors.New("unsupported CHANGELOCK_SELF_HEALING_MODE: " + string(mode))
	}

	maxAttempts := 3
	if raw := strings.TrimSpace(os.Getenv("CHANGELOCK_SELF_HEALING_MAX_ATTEMPTS")); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil || value <= 0 {
			return SelfHealingConfig{}, errors.New("CHANGELOCK_SELF_HEALING_MAX_ATTEMPTS must be a positive integer")
		}
		maxAttempts = value
	}

	window := 15 * time.Minute
	if raw := strings.TrimSpace(os.Getenv("CHANGELOCK_SELF_HEALING_WINDOW")); raw != "" {
		value, err := time.ParseDuration(raw)
		if err != nil || value <= 0 {
			return SelfHealingConfig{}, errors.New("CHANGELOCK_SELF_HEALING_WINDOW must be a positive duration")
		}
		window = value
	}

	allowedKinds, err := parseAllowedKinds(os.Getenv("CHANGELOCK_SELF_HEALING_ALLOWED_KINDS"))
	if err != nil {
		return SelfHealingConfig{}, err
	}

	failMode, err := parseFailMode(os.Getenv("CHANGELOCK_CLOSED_LOOP_FAIL_MODE"))
	if err != nil {
		return SelfHealingConfig{}, err
	}
	reconcileInterval, err := parseReconcileInterval(os.Getenv("CHANGELOCK_CLOSED_LOOP_RECONCILE_INTERVAL"))
	if err != nil {
		return SelfHealingConfig{}, err
	}
	protectedNamespaces := parseCSVSet(os.Getenv("CHANGELOCK_CLOSED_LOOP_PROTECTED_NAMESPACES"), []string{"changelock", "changelock-system"})
	protectedWorkloads := parseCSVSet(os.Getenv("CHANGELOCK_CLOSED_LOOP_PROTECTED_WORKLOADS"), nil)
	runtimeVEXSeverity, err := parseRuntimeVEXSeverity(os.Getenv("CHANGELOCK_RUNTIME_VEX_QUARANTINE_SEVERITY"))
	if err != nil {
		return SelfHealingConfig{}, err
	}
	requireSignedDesiredState := parseBoolEnv("CHANGELOCK_CLOSED_LOOP_REQUIRE_SIGNED_DESIRED_STATE")
	if !requireSignedDesiredState {
		requireSignedDesiredState = parseBoolEnv("CHANGELOCK_SELF_HEALING_REQUIRE_SIGNED_DESIRED_STATE")
	}

	return SelfHealingConfig{
		Mode:                      mode,
		MaxAttempts:               maxAttempts,
		Window:                    window,
		AllowedKinds:              allowedKinds,
		RequireSignedDesiredState: requireSignedDesiredState,
		CriticalOnly:              parseBoolEnv("CHANGELOCK_SELF_HEALING_CRITICAL_ONLY"),
		FailMode:                  failMode,
		VerifyDesiredStateOnLoop:  parseBoolEnv("CHANGELOCK_CLOSED_LOOP_VERIFY_DESIRED_STATE_ON_RECONCILE"),
		ReconcileInterval:         reconcileInterval,
		ProtectedNamespaces:       protectedNamespaces,
		ProtectedWorkloads:        protectedWorkloads,
		QuarantineNetworkPolicy:   parseBoolEnv("CHANGELOCK_RUNTIME_QUARANTINE_NETWORK_POLICY_ENABLED"),
		RuntimeVEXQuarantine:      parseBoolEnv("CHANGELOCK_RUNTIME_VEX_QUARANTINE_ENABLED"),
		RuntimeVEXSeverity:        runtimeVEXSeverity,
		RuntimeVEXRequireNet:      !strings.EqualFold(strings.TrimSpace(os.Getenv("CHANGELOCK_RUNTIME_VEX_QUARANTINE_REQUIRE_NET_ACTIONABLE")), "false"),
	}, nil
}

func parseAllowedKinds(raw string) (map[string]struct{}, error) {
	normalized := strings.TrimSpace(raw)
	if normalized == "" {
		normalized = "Deployment,DaemonSet,StatefulSet"
	}

	allowed := make(map[string]struct{})
	for _, item := range strings.Split(normalized, ",") {
		kind := normalizeWorkloadKind(item)
		if kind == "" {
			continue
		}
		switch kind {
		case "Deployment", "DaemonSet", "StatefulSet":
		default:
			return nil, errors.New("unsupported CHANGELOCK_SELF_HEALING_ALLOWED_KINDS entry: " + kind)
		}
		allowed[kind] = struct{}{}
	}
	if len(allowed) == 0 {
		return nil, errors.New("CHANGELOCK_SELF_HEALING_ALLOWED_KINDS must contain at least one supported workload kind")
	}
	return allowed, nil
}

func parseFailMode(raw string) (RemediationMode, error) {
	value := strings.ToLower(strings.TrimSpace(raw))
	if value == "" {
		return RemediationModeQuarantine, nil
	}
	switch RemediationMode(value) {
	case RemediationModeQuarantine, RemediationModeAlertOnly:
		return RemediationMode(value), nil
	default:
		return "", errors.New("CHANGELOCK_CLOSED_LOOP_FAIL_MODE must be quarantine or alert-only")
	}
}

func parseReconcileInterval(raw string) (time.Duration, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return 2 * time.Minute, nil
	}
	value, err := time.ParseDuration(trimmed)
	if err != nil || value <= 0 {
		return 0, errors.New("CHANGELOCK_CLOSED_LOOP_RECONCILE_INTERVAL must be a positive duration")
	}
	return value, nil
}

func parseCSVSet(raw string, defaults []string) map[string]struct{} {
	values := defaults
	if strings.TrimSpace(raw) != "" {
		values = strings.Split(raw, ",")
	}
	result := map[string]struct{}{}
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed != "" {
			result[trimmed] = struct{}{}
		}
	}
	return result
}

func parseRuntimeVEXSeverity(raw string) (string, error) {
	value := strings.ToLower(strings.TrimSpace(raw))
	if value == "" {
		return "critical", nil
	}
	switch value {
	case "critical", "high", "medium", "low", "unknown":
		return strings.ToUpper(value), nil
	default:
		return "", errors.New("CHANGELOCK_RUNTIME_VEX_QUARANTINE_SEVERITY must be critical, high, medium, low, or unknown")
	}
}

func parseBoolEnv(key string) bool {
	value := strings.ToLower(strings.TrimSpace(os.Getenv(key)))
	return value == "1" || value == "true" || value == "yes" || value == "on"
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); strings.TrimSpace(value) != "" {
		return value
	}
	return fallback
}
