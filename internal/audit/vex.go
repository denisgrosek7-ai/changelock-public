package audit

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	internalvex "github.com/denisgrosek/changelock/internal/vex"
)

const (
	EventTypeVEXStatementRecorded = "vex_statement_recorded"
	EventTypeVEXStatementRevoked  = "vex_statement_revoked"
)

type VEXMatch struct {
	ID              int64      `json:"id"`
	SourceFormat    string     `json:"source_format"`
	SourceRef       string     `json:"source_ref,omitempty"`
	VulnerabilityID string     `json:"vulnerability_id"`
	Status          string     `json:"status"`
	Justification   string     `json:"justification,omitempty"`
	ActionStatement string     `json:"action_statement,omitempty"`
	ImpactStatement string     `json:"impact_statement,omitempty"`
	FixedVersion    string     `json:"fixed_version,omitempty"`
	CreatedBy       string     `json:"created_by,omitempty"`
	UpdatedBy       string     `json:"updated_by,omitempty"`
	ExpiresAt       *time.Time `json:"expires_at,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type VulnerabilityNetResponse struct {
	RawCount                int                    `json:"raw_count"`
	ResolvedByVEXCount      int                    `json:"resolved_by_vex_count"`
	ActionableCount         int                    `json:"actionable_count"`
	UnderInvestigationCount int                    `json:"under_investigation_count"`
	SeverityThreshold       string                 `json:"severity_threshold,omitempty"`
	ThresholdBreached       bool                   `json:"threshold_breached"`
	Findings                []VulnerabilityFinding `json:"findings"`
	AppliedFilters          map[string]string      `json:"applied_filters"`
}

func cloneVEXMatch(match *VEXMatch) *VEXMatch {
	if match == nil {
		return nil
	}
	copy := *match
	if match.ExpiresAt != nil {
		expiresAt := match.ExpiresAt.UTC()
		copy.ExpiresAt = &expiresAt
	}
	return &copy
}

func cloneVulnerabilityFindingWithVEX(finding VulnerabilityFinding) VulnerabilityFinding {
	finding = cloneVulnerabilityFinding(finding)
	finding.VEX = cloneVEXMatch(finding.VEX)
	return finding
}

func statementToVEXMatch(statement internalvex.Statement) *VEXMatch {
	match := &VEXMatch{
		ID:              statement.ID,
		SourceFormat:    statement.SourceFormat,
		SourceRef:       statement.SourceRef,
		VulnerabilityID: statement.VulnerabilityID,
		Status:          statement.Status,
		Justification:   statement.Justification,
		ActionStatement: statement.ActionStatement,
		ImpactStatement: statement.ImpactStatement,
		FixedVersion:    statement.FixedVersion,
		CreatedBy:       statement.CreatedBy,
		UpdatedBy:       statement.UpdatedBy,
		CreatedAt:       statement.CreatedAt,
		UpdatedAt:       statement.UpdatedAt,
	}
	if statement.ExpiresAt != nil {
		expiresAt := statement.ExpiresAt.UTC()
		match.ExpiresAt = &expiresAt
	}
	return match
}

func statementToLegacyDecision(statement internalvex.Statement) *VulnerabilityDecision {
	decision := &VulnerabilityDecision{
		ID:            statement.ID,
		ImageDigest:   statement.Scope.ImageDigest,
		CVEID:         statement.VulnerabilityID,
		Decision:      legacyDecisionValue(statement),
		Justification: firstNonEmpty(statement.Justification, statement.ImpactStatement),
		DecidedBy:     firstNonEmpty(statement.UpdatedBy, statement.CreatedBy),
		ExpiresAt:     statement.ExpiresAt,
		Active:        statement.Active && statement.RevokedAt == nil,
		Metadata:      cloneLegacyMetadata(statement.Metadata),
		CreatedAt:     statement.CreatedAt,
		UpdatedAt:     statement.UpdatedAt,
	}
	return decision
}

func legacyDecisionValue(statement internalvex.Statement) string {
	metadata := legacyDecisionMetadata(statement.Metadata)
	if value := strings.TrimSpace(metadata["legacy_decision"]); value != "" {
		return strings.ToUpper(value)
	}
	switch strings.TrimSpace(statement.Status) {
	case internalvex.StatusNotAffected:
		return VulnerabilityDecisionNotAffected
	case internalvex.StatusUnderInvestigation:
		return VulnerabilityDecisionUnderInvestigation
	case internalvex.StatusFixed:
		return VulnerabilityDecisionFixRequired
	default:
		return VulnerabilityDecisionAcceptedRisk
	}
}

func legacyDecisionMetadata(raw json.RawMessage) map[string]string {
	if len(raw) == 0 {
		return nil
	}
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil
	}
	results := map[string]string{}
	for key, value := range payload {
		if text, ok := value.(string); ok {
			results[key] = text
		}
	}
	return results
}

func cloneLegacyMetadata(raw json.RawMessage) json.RawMessage {
	if len(raw) == 0 {
		return nil
	}
	return append(json.RawMessage(nil), raw...)
}

func vexScopeFromFilter(filter VulnerabilityActiveFilter) internalvex.EvaluationScope {
	return internalvex.EvaluationScope{
		TenantID:    strings.TrimSpace(filter.TenantID),
		Environment: strings.TrimSpace(filter.Environment),
	}
}

func findingRefFromFinding(finding VulnerabilityFinding) internalvex.FindingRef {
	return internalvex.FindingRef{
		VulnerabilityID: finding.CVEID,
		ImageDigest:     finding.ImageDigest,
		PackageName:     finding.PackageName,
		PURL:            finding.PURL,
	}
}

func matchStatementForFinding(statements []internalvex.Statement, finding VulnerabilityFinding, scope internalvex.EvaluationScope, now time.Time) *internalvex.Statement {
	var matched *internalvex.Statement
	findingRef := findingRefFromFinding(finding)
	for _, statement := range statements {
		if !internalvex.Matches(statement, findingRef, scope, now) {
			continue
		}
		if matched == nil || statement.UpdatedAt.After(matched.UpdatedAt) {
			copy := statement
			matched = &copy
		}
	}
	return matched
}

func statementsByDigestAndVulnerability(statements []internalvex.Statement) map[string][]internalvex.Statement {
	grouped := map[string][]internalvex.Statement{}
	for _, statement := range statements {
		key := strings.Join([]string{strings.TrimSpace(statement.Scope.ImageDigest), strings.TrimSpace(statement.VulnerabilityID)}, "|")
		grouped[key] = append(grouped[key], statement)
	}
	return grouped
}

func resolveMatchingVEX(ctx context.Context, store Store, filter VulnerabilityActiveFilter, findings []VulnerabilityFinding) (map[string]*internalvex.Statement, error) {
	digests := map[string]struct{}{}
	for _, finding := range findings {
		digests[finding.ImageDigest] = struct{}{}
	}
	matches := map[string]*internalvex.Statement{}
	if len(digests) == 0 {
		return matches, nil
	}
	now := time.Now().UTC()
	scope := vexScopeFromFilter(filter)
	for digest := range digests {
		statements, err := store.ListVEXStatements(ctx, internalvex.Filter{
			ImageDigest: digest,
			Active:      boolPointer(true),
			Limit:       500,
		})
		if err != nil {
			return nil, err
		}
		for _, finding := range findings {
			if finding.ImageDigest != digest {
				continue
			}
			key := strings.Join([]string{finding.ImageDigest, finding.CVEID, firstNonEmpty(finding.PURL, finding.PackageName)}, "|")
			if _, ok := matches[key]; ok {
				continue
			}
			match := matchStatementForFinding(statements, finding, scope, now)
			if match != nil {
				matches[key] = match
			}
		}
	}
	return matches, nil
}

func legacyDecisionCreateRequestToVEX(request VulnerabilityDecisionCreateRequest) (internalvex.CreateRequest, error) {
	request, err := NormalizeVulnerabilityDecisionCreateRequest(request, time.Now)
	if err != nil {
		return internalvex.CreateRequest{}, err
	}
	createRequest := internalvex.CreateRequest{
		SourceFormat:    internalvex.SourceFormatAPI,
		VulnerabilityID: request.CVEID,
		Scope: internalvex.Scope{
			ImageDigest: request.ImageDigest,
		},
		Justification: request.Justification,
		ExpiresAt:     request.ExpiresAt,
	}
	metadataMap := map[string]string{
		"legacy_decision": strings.ToUpper(strings.TrimSpace(request.Decision)),
	}
	switch strings.ToUpper(strings.TrimSpace(request.Decision)) {
	case VulnerabilityDecisionNotAffected:
		createRequest.Status = internalvex.StatusNotAffected
	case VulnerabilityDecisionUnderInvestigation:
		createRequest.Status = internalvex.StatusUnderInvestigation
	case VulnerabilityDecisionFixRequired:
		createRequest.Status = internalvex.StatusAffected
		createRequest.ActionStatement = "fix_required"
	case VulnerabilityDecisionAcceptedRisk:
		createRequest.Status = internalvex.StatusAffected
		createRequest.ActionStatement = "accepted_risk"
	default:
		return internalvex.CreateRequest{}, fmt.Errorf("%w: unsupported vulnerability decision %q", ErrInvalidException, request.Decision)
	}
	metadata, _ := json.Marshal(metadataMap)
	createRequest.Metadata = metadata
	return internalvex.NormalizeCreateRequest(createRequest, time.Now)
}
