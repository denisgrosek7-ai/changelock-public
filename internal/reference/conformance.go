package reference

import "strings"

const (
	StateActive     = "bounded_conformance_active"
	StatePartial    = "bounded_conformance_partial"
	StateIncomplete = "bounded_conformance_incomplete"
)

type Check struct {
	CheckID     string `json:"check_id"`
	Required    bool   `json:"required"`
	State       string `json:"state"`
	Summary     string `json:"summary"`
	EvidenceRef string `json:"evidence_ref,omitempty"`
}

type Input struct {
	ArchitectureID string  `json:"architecture_id"`
	Checks         []Check `json:"checks,omitempty"`
}

type Result struct {
	ArchitectureID       string   `json:"architecture_id"`
	CurrentState         string   `json:"current_state"`
	RequiredChecksPassed []string `json:"required_checks_passed,omitempty"`
	OptionalChecksPassed []string `json:"optional_checks_passed,omitempty"`
	UnsupportedChecks    []string `json:"unsupported_checks,omitempty"`
	DegradedChecks       []string `json:"degraded_checks,omitempty"`
	Deviations           []string `json:"deviations,omitempty"`
	Limitations          []string `json:"limitations,omitempty"`
}

func Evaluate(input Input) Result {
	result := Result{
		ArchitectureID: input.ArchitectureID,
		CurrentState:   StateActive,
	}

	requiredFailed := false
	optionalDegraded := false
	for _, check := range input.Checks {
		switch strings.TrimSpace(check.State) {
		case "active", "ready", "verified":
			if check.Required {
				result.RequiredChecksPassed = append(result.RequiredChecksPassed, check.CheckID)
			} else {
				result.OptionalChecksPassed = append(result.OptionalChecksPassed, check.CheckID)
			}
		case "unsupported":
			result.UnsupportedChecks = append(result.UnsupportedChecks, check.CheckID)
			if check.Required {
				requiredFailed = true
			}
		case "degraded", "warning", "partial":
			result.DegradedChecks = append(result.DegradedChecks, check.CheckID)
			if check.Required {
				requiredFailed = true
			} else {
				optionalDegraded = true
			}
		default:
			result.Deviations = append(result.Deviations, check.CheckID)
			if check.Required {
				requiredFailed = true
			} else {
				optionalDegraded = true
			}
		}
	}

	switch {
	case requiredFailed:
		result.CurrentState = StateIncomplete
	case optionalDegraded:
		result.CurrentState = StatePartial
	default:
		result.CurrentState = StateActive
	}
	result.Limitations = []string{
		"Conformance remains bounded to the declared reference architecture and the published public-proof surfaces.",
		"Unsupported or degraded checks remain visible and are not collapsed into a score-only result.",
	}
	return result
}
