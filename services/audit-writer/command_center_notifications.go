package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

type commandCenterNotification struct {
	SchemaVersion    string                   `json:"schema_version"`
	NotificationID   string                   `json:"notification_id"`
	LifecyclePhase   string                   `json:"lifecycle_phase"`
	Severity         string                   `json:"severity"`
	CurrentState     string                   `json:"current_state"`
	Title            string                   `json:"title"`
	Summary          string                   `json:"summary"`
	OwnerHint        string                   `json:"owner_hint,omitempty"`
	NextAction       string                   `json:"next_action,omitempty"`
	SourceSubsystems []string                 `json:"source_subsystems,omitempty"`
	EvidenceRefs     []string                 `json:"evidence_refs,omitempty"`
	PersonaHints     []string                 `json:"persona_hints,omitempty"`
	Target           commandCenterFocusTarget `json:"target"`
}

type commandCenterNotificationsResponse struct {
	SchemaVersion string                      `json:"schema_version"`
	GeneratedAt   time.Time                   `json:"generated_at"`
	CountsByState map[string]int              `json:"counts_by_state"`
	Items         []commandCenterNotification `json:"items"`
	Limitations   []string                    `json:"limitations,omitempty"`
}

type commandCenterNotificationBucket struct {
	id               string
	lifecyclePhase   string
	severity         string
	currentState     string
	title            string
	summary          string
	ownerHint        string
	nextAction       string
	sourceSubsystems []string
	evidenceRefs     []string
	personaHints     []string
	target           commandCenterFocusTarget
	latestAt         time.Time
	entryCount       int
	subjectLabel     string
}

func (s server) commandCenterNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	lifecyclePhase := normalizeCommandCenterLifecyclePhase(r.URL.Query().Get("lifecycle_phase"))
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildCommandCenterNotifications(ctx, filter, lifecyclePhase)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) buildCommandCenterNotifications(ctx context.Context, filter audit.EventFilter, lifecyclePhase string) (commandCenterNotificationsResponse, error) {
	limit := filter.Limit
	if limit <= 0 {
		limit = 8
	}
	timeline, err := s.buildSecurityTimeline(ctx, filter, lifecyclePhase)
	if err != nil {
		return commandCenterNotificationsResponse{}, err
	}
	workflowOwners, err := s.phase4WorkflowOwnerHints(ctx, filter, maxInt(limit*8, 80))
	if err != nil {
		return commandCenterNotificationsResponse{}, err
	}

	buckets := map[string]*commandCenterNotificationBucket{}
	for _, entry := range timeline.Entries {
		if suppressCommandCenterNotificationEntry(entry) {
			continue
		}
		target, ok := commandCenterNotificationTarget(entry)
		if !ok {
			continue
		}
		key := fmt.Sprintf("%s|%s|%s", entry.LifecyclePhase, target.Kind, target.Ref)
		bucket, exists := buckets[key]
		if !exists {
			bucket = &commandCenterNotificationBucket{
				id:               key,
				lifecyclePhase:   entry.LifecyclePhase,
				severity:         entry.Severity,
				currentState:     commandCenterNotificationState(entry),
				title:            entry.Title,
				summary:          entry.Summary,
				ownerHint:        commandCenterNotificationOwnerHint(entry, workflowOwners),
				nextAction:       strings.TrimSpace(entry.NextAction),
				sourceSubsystems: []string{entry.SourceSubsystem},
				evidenceRefs:     append([]string{}, entry.EvidenceRefs...),
				personaHints:     append([]string{}, entry.PersonaHints...),
				target:           target,
				latestAt:         entry.Timestamp.UTC(),
				entryCount:       1,
				subjectLabel:     entry.SubjectLabel,
			}
			buckets[key] = bucket
			continue
		}
		bucket.entryCount++
		if commandCenterSeverityRank(entry.Severity) > commandCenterSeverityRank(bucket.severity) {
			bucket.severity = entry.Severity
			bucket.currentState = commandCenterNotificationState(entry)
			bucket.title = entry.Title
			bucket.summary = entry.Summary
			if entry.NextAction != "" {
				bucket.nextAction = entry.NextAction
			}
		}
		if entry.Timestamp.After(bucket.latestAt) {
			bucket.latestAt = entry.Timestamp.UTC()
		}
		if bucket.ownerHint == "" {
			bucket.ownerHint = commandCenterNotificationOwnerHint(entry, workflowOwners)
		}
		if bucket.nextAction == "" && entry.NextAction != "" {
			bucket.nextAction = entry.NextAction
		}
		bucket.sourceSubsystems = uniqueStrings(append(bucket.sourceSubsystems, entry.SourceSubsystem))
		bucket.evidenceRefs = uniqueStrings(append(bucket.evidenceRefs, entry.EvidenceRefs...))
		bucket.personaHints = uniqueStrings(append(bucket.personaHints, entry.PersonaHints...))
	}

	ordered := make([]*commandCenterNotificationBucket, 0, len(buckets))
	for _, bucket := range buckets {
		ordered = append(ordered, bucket)
	}
	sort.Slice(ordered, func(i, j int) bool {
		if commandCenterSeverityRank(ordered[i].severity) != commandCenterSeverityRank(ordered[j].severity) {
			return commandCenterSeverityRank(ordered[i].severity) > commandCenterSeverityRank(ordered[j].severity)
		}
		if !ordered[i].latestAt.Equal(ordered[j].latestAt) {
			return ordered[i].latestAt.After(ordered[j].latestAt)
		}
		return ordered[i].id < ordered[j].id
	})

	items := make([]commandCenterNotification, 0, len(ordered))
	for _, bucket := range ordered {
		title := bucket.title
		summary := bucket.summary
		if bucket.entryCount > 1 {
			title = fmt.Sprintf("%d %s signals for %s", bucket.entryCount, strings.ReplaceAll(bucket.lifecyclePhase, "_", " "), bucket.subjectLabel)
			summary = fmt.Sprintf("%s %d related signals were grouped into one actionable notification.", bucket.summary, bucket.entryCount-1)
		}
		items = append(items, commandCenterNotification{
			SchemaVersion:    commandNotificationSchemaVersion,
			NotificationID:   bucket.id,
			LifecyclePhase:   bucket.lifecyclePhase,
			Severity:         bucket.severity,
			CurrentState:     bucket.currentState,
			Title:            title,
			Summary:          summary,
			OwnerHint:        bucket.ownerHint,
			NextAction:       bucket.nextAction,
			SourceSubsystems: uniqueStrings(bucket.sourceSubsystems),
			EvidenceRefs:     limitStrings(bucket.evidenceRefs, 8),
			PersonaHints:     uniqueStrings(bucket.personaHints),
			Target:           bucket.target,
		})
	}
	if len(items) > limit {
		items = items[:limit]
	}

	countsByState := map[string]int{}
	for _, item := range items {
		countsByState[item.CurrentState]++
	}

	return commandCenterNotificationsResponse{
		SchemaVersion: commandNotificationsSchemaVersion,
		GeneratedAt:   timeline.GeneratedAt,
		CountsByState: countsByState,
		Items:         items,
		Limitations: []string{
			"Notifications are grouped projections over canonical timeline and workflow evidence; they do not create a separate alert truth layer.",
			"Low-signal resolved or allow-path events are intentionally suppressed so the command center stays actionable and state-aware.",
		},
	}, nil
}

func (s server) phase4WorkflowOwnerHints(ctx context.Context, filter audit.EventFilter, limit int) (map[string]string, error) {
	items, err := s.listPhase4WorkflowArtifacts(ctx, phase4EnterpriseFilter{
		ClusterID:   filter.ClusterID,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Limit:       limit,
	})
	if err != nil {
		return nil, err
	}
	hints := map[string]string{}
	for _, item := range items {
		owner := strings.TrimSpace(item.Routing.PrimaryOwner)
		if owner == "" {
			continue
		}
		if workflowID := strings.TrimSpace(item.WorkflowID); workflowID != "" {
			hints[workflowID] = owner
		}
		if subjectRef := strings.TrimSpace(item.SubjectRef); subjectRef != "" {
			hints[subjectRef] = owner
		}
	}
	return hints, nil
}

func suppressCommandCenterNotificationEntry(entry securityTimelineEntry) bool {
	return commandCenterSeverityRank(entry.Severity) <= 1 && (entry.Outcome == "allow" || entry.Outcome == "signal") && entry.NextAction == ""
}

func commandCenterNotificationTarget(entry securityTimelineEntry) (commandCenterFocusTarget, bool) {
	if entry.DrilldownTargetKind != "" && entry.DrilldownTargetRef != "" {
		return commandCenterFocusTarget{
			Tab:          entry.DrilldownTab,
			Kind:         entry.DrilldownTargetKind,
			Ref:          entry.DrilldownTargetRef,
			SecondaryRef: entry.DrilldownTargetSecondaryRef,
			ResourceURI:  entry.ResourceURI,
		}, true
	}

	switch entry.LifecyclePhase {
	case "runtime":
		if entry.SubjectRef != "" {
			return commandCenterFocusTarget{Tab: "runtime", Kind: "runtime_subject", Ref: entry.SubjectRef}, true
		}
	case "validation":
		if entry.SubjectRef != "" {
			return commandCenterFocusTarget{Tab: "validation", Kind: "validation_scenario", Ref: entry.SubjectRef}, true
		}
	case "workflow":
		if entry.SubjectRef != "" {
			return commandCenterFocusTarget{Tab: "exceptions", Kind: "workflow_record", Ref: entry.SubjectRef}, true
		}
	case "partner":
		if entry.SubjectRef != "" {
			return commandCenterFocusTarget{Tab: "federation", Kind: "partner_trust", Ref: entry.SubjectRef}, true
		}
	case "governance":
		if entry.SubjectRef != "" {
			return commandCenterFocusTarget{Tab: "analytics", Kind: "executive_report", Ref: entry.SubjectRef}, true
		}
	case "intelligence":
		if entry.SubjectRef != "" {
			return commandCenterFocusTarget{Tab: "guidance", Kind: "grounded_query", Ref: entry.SubjectRef}, true
		}
	}
	return commandCenterFocusTarget{}, false
}

func commandCenterNotificationOwnerHint(entry securityTimelineEntry, workflowOwners map[string]string) string {
	for _, key := range []string{entry.DrilldownTargetRef, entry.DrilldownTargetSecondaryRef, entry.SubjectRef} {
		if owner := strings.TrimSpace(workflowOwners[key]); owner != "" {
			return owner
		}
	}
	return ""
}

func commandCenterNotificationState(entry securityTimelineEntry) string {
	switch {
	case entry.Outcome == "error":
		return "degraded"
	case entry.Outcome == "deny" || commandCenterSeverityRank(entry.Severity) >= 3:
		return "action_required"
	case entry.NextAction != "":
		return "review_required"
	default:
		return "watch"
	}
}
