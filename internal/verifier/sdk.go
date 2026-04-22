package verifier

import (
	"strings"
	"time"
)

const (
	StateVerified            = "verified"
	StateInvalid             = "invalid"
	StateExpired             = "expired"
	StateUnsupported         = "unsupported_schema"
	StateIncomplete          = "incomplete_evidence"
	FreshnessStateFresh      = "fresh"
	FreshnessStateStale      = "stale"
	FreshnessStateExpired    = "expired"
	CompatibilitySupported   = "supported"
	CompatibilityUnsupported = "unsupported"
	DeprecationCurrent       = "current"
	DeprecationDeprecated    = "deprecated"
)

type Input struct {
	ArtifactID         string    `json:"artifact_id"`
	ArtifactType       string    `json:"artifact_type"`
	SchemaVersion      string    `json:"schema_version,omitempty"`
	VerifiedAt         time.Time `json:"verified_at,omitempty"`
	ValidUntil         time.Time `json:"valid_until,omitempty"`
	EvidenceRefs       []string  `json:"evidence_refs,omitempty"`
	ExportRefs         []string  `json:"export_refs,omitempty"`
	SupportedSchemas   []string  `json:"supported_schemas,omitempty"`
	IntegrityConfirmed bool      `json:"integrity_confirmed"`
	ChainContinuity    bool      `json:"chain_continuity"`
	MethodologyRef     string    `json:"methodology_ref,omitempty"`
	Deprecated         bool      `json:"deprecated"`
}

type Result struct {
	ArtifactID         string    `json:"artifact_id"`
	ArtifactType       string    `json:"artifact_type"`
	CurrentState       string    `json:"current_state"`
	FreshnessState     string    `json:"freshness_state"`
	CompatibilityState string    `json:"compatibility_state"`
	DeprecationState   string    `json:"deprecation_state"`
	VerifiedAt         time.Time `json:"verified_at,omitempty"`
	ValidUntil         time.Time `json:"valid_until,omitempty"`
	ExportRefs         []string  `json:"export_refs,omitempty"`
	ReasonCodes        []string  `json:"reason_codes,omitempty"`
	Limitations        []string  `json:"limitations,omitempty"`
}

func Evaluate(input Input, asOf time.Time) Result {
	if asOf.IsZero() {
		asOf = time.Now().UTC()
	}
	input.ArtifactID = strings.TrimSpace(input.ArtifactID)
	input.ArtifactType = strings.TrimSpace(input.ArtifactType)
	input.SchemaVersion = strings.TrimSpace(input.SchemaVersion)
	input.MethodologyRef = strings.TrimSpace(input.MethodologyRef)

	reasons := make([]string, 0, 4)
	state := StateVerified
	freshness := FreshnessStateFresh
	compatibilityState := CompatibilitySupported
	deprecationState := DeprecationCurrent

	if input.ArtifactType == "" || input.SchemaVersion == "" {
		state = StateUnsupported
		compatibilityState = CompatibilityUnsupported
		reasons = append(reasons, "artifact_type_or_schema_missing")
	}
	if len(input.SupportedSchemas) > 0 && !containsString(input.SupportedSchemas, input.SchemaVersion) {
		state = StateUnsupported
		compatibilityState = CompatibilityUnsupported
		reasons = append(reasons, "schema_version_not_supported")
	}
	if len(input.EvidenceRefs) == 0 {
		state = StateIncomplete
		reasons = append(reasons, "evidence_refs_missing")
	}
	if !input.IntegrityConfirmed || !input.ChainContinuity {
		state = StateInvalid
		reasons = append(reasons, "integrity_or_chain_continuity_failed")
	}
	if !input.ValidUntil.IsZero() {
		switch {
		case asOf.After(input.ValidUntil):
			state = StateExpired
			freshness = FreshnessStateExpired
			reasons = append(reasons, "artifact_freshness_expired")
		case asOf.After(input.ValidUntil.Add(-14 * 24 * time.Hour)):
			freshness = FreshnessStateStale
			reasons = append(reasons, "artifact_freshness_stale")
		}
	}
	if input.ArtifactType == "benchmark_pack" && input.MethodologyRef == "" {
		state = StateIncomplete
		reasons = append(reasons, "benchmark_methodology_missing")
	}
	if input.Deprecated {
		deprecationState = DeprecationDeprecated
		reasons = append(reasons, "schema_line_deprecated")
	}
	if state == StateVerified {
		reasons = append(reasons, "artifact_verifier_ready")
	}

	return Result{
		ArtifactID:         input.ArtifactID,
		ArtifactType:       input.ArtifactType,
		CurrentState:       state,
		FreshnessState:     freshness,
		CompatibilityState: compatibilityState,
		DeprecationState:   deprecationState,
		VerifiedAt:         input.VerifiedAt.UTC(),
		ValidUntil:         input.ValidUntil.UTC(),
		ExportRefs:         uniqueStrings(input.ExportRefs),
		ReasonCodes:        uniqueStrings(reasons),
		Limitations: []string{
			"Verifier result remains bounded to the supplied artifact, schema line, evidence refs, and freshness window.",
			"A verified artifact does not by itself prove global admissibility or customer-specific policy acceptance.",
		},
	}
}

func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := map[string]struct{}{}
	items := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		items = append(items, value)
	}
	if len(items) == 0 {
		return nil
	}
	return items
}

func containsString(values []string, target string) bool {
	target = strings.TrimSpace(target)
	for _, value := range values {
		if strings.TrimSpace(value) == target {
			return true
		}
	}
	return false
}
