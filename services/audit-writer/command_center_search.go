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

type commandCenterFocusTarget struct {
	Tab          string `json:"tab"`
	Kind         string `json:"kind"`
	Ref          string `json:"ref"`
	SecondaryRef string `json:"secondary_ref,omitempty"`
	ResourceURI  string `json:"resource_uri,omitempty"`
}

type commandCenterSearchResult struct {
	SchemaVersion     string                   `json:"schema_version"`
	ResultID          string                   `json:"result_id"`
	ResultType        string                   `json:"result_type"`
	Title             string                   `json:"title"`
	Summary           string                   `json:"summary"`
	Subtitle          string                   `json:"subtitle,omitempty"`
	SourceSubsystem   string                   `json:"source_subsystem"`
	Severity          string                   `json:"severity"`
	Target            commandCenterFocusTarget `json:"target"`
	IncidentRef       string                   `json:"incident_ref,omitempty"`
	RecommendationRef string                   `json:"recommendation_ref,omitempty"`
	EvidenceRefs      []string                 `json:"evidence_refs,omitempty"`
	PersonaHints      []string                 `json:"persona_hints,omitempty"`
	Limitations       []string                 `json:"limitations,omitempty"`
}

type commandCenterSearchResponse struct {
	SchemaVersion string                      `json:"schema_version"`
	Query         string                      `json:"query"`
	Results       []commandCenterSearchResult `json:"results"`
	Limitations   []string                    `json:"limitations,omitempty"`
}

type commandCenterSearchCandidate struct {
	score     int
	timestamp time.Time
	result    commandCenterSearchResult
}

func (s server) commandCenterSearchHandler(w http.ResponseWriter, r *http.Request) {
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
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildCommandCenterSearch(ctx, filter, query)
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

func (s server) buildCommandCenterSearch(ctx context.Context, filter audit.EventFilter, query string) (commandCenterSearchResponse, error) {
	trimmedQuery := strings.TrimSpace(query)
	limit := filter.Limit
	if limit <= 0 {
		limit = 8
	}
	if trimmedQuery == "" {
		return commandCenterSearchResponse{
			SchemaVersion: commandSearchResponseSchemaVersion,
			Query:         "",
			Results:       []commandCenterSearchResult{},
			Limitations: []string{
				"Command-center search stays semantic and evidence-backed; empty queries intentionally do not fan out into a global unbounded object listing.",
			},
		}, nil
	}

	contextFilter := securityTimelineContextFilter(filter, limit)
	candidates := []commandCenterSearchCandidate{}

	incidents, err := s.listIncidents(ctx, incidentFilter{event: contextFilter})
	if err != nil {
		return commandCenterSearchResponse{}, err
	}
	for _, incident := range incidents {
		score := commandCenterMatchScore(trimmedQuery,
			incident.ID,
			incident.Title,
			incident.Summary,
			incident.ScopeRef,
			incident.Repository,
			strings.Join(incident.AffectedWorkloads, " "),
			strings.Join(incident.ReasonCodes, " "),
		)
		if score == 0 {
			continue
		}
		timestamp := time.Unix(0, 0).UTC()
		if incident.LastActivityAt != nil && !incident.LastActivityAt.IsZero() {
			timestamp = incident.LastActivityAt.UTC()
		} else if incident.UpdatedAt != nil && !incident.UpdatedAt.IsZero() {
			timestamp = incident.UpdatedAt.UTC()
		}
		candidates = append(candidates, commandCenterSearchCandidate{
			score:     score,
			timestamp: timestamp,
			result: commandCenterSearchResult{
				SchemaVersion:   commandSearchResultSchemaVersion,
				ResultID:        "incident:" + incident.ID,
				ResultType:      "incident",
				Title:           incident.Title,
				Summary:         firstNonEmpty(strings.TrimSpace(incident.CaseSummary), strings.TrimSpace(incident.Summary), strings.TrimSpace(incident.LikelyCause)),
				Subtitle:        firstNonEmpty(strings.TrimSpace(incident.ScopeRef), strings.TrimSpace(incident.Repository), strings.TrimSpace(incident.Environment)),
				SourceSubsystem: "incident",
				Severity:        incident.Severity,
				Target: commandCenterFocusTarget{
					Tab:  "events",
					Kind: "incident",
					Ref:  incident.ID,
				},
				IncidentRef:  incident.ID,
				EvidenceRefs: limitStrings(incident.EvidenceRefs, 8),
				PersonaHints: securityTimelinePersonaHints("incident", incident.Severity),
			},
		})
	}

	recommendations, err := s.listRecommendations(ctx, recommendationFilter{
		event: contextFilter,
		Limit: maxInt(limit*8, 80),
	})
	if err != nil {
		return commandCenterSearchResponse{}, err
	}
	for _, item := range recommendations {
		score := commandCenterMatchScore(trimmedQuery,
			item.RecommendationID,
			item.Title,
			item.SubjectRef,
			item.Service,
			item.Repo,
			item.RecommendedAction,
			item.Rationale,
		)
		if score == 0 {
			continue
		}
		candidates = append(candidates, commandCenterSearchCandidate{
			score:     score,
			timestamp: item.CreatedAt.UTC(),
			result: commandCenterSearchResult{
				SchemaVersion:   commandSearchResultSchemaVersion,
				ResultID:        "recommendation:" + item.RecommendationID,
				ResultType:      "recommendation",
				Title:           item.Title,
				Summary:         item.RecommendedAction,
				Subtitle:        firstNonEmpty(item.SubjectRef, item.Service, item.Repo),
				SourceSubsystem: "recommendation",
				Severity:        commandCenterRecommendationSeverity(item),
				Target: commandCenterFocusTarget{
					Tab:          "overview",
					Kind:         "recommendation",
					Ref:          item.RecommendationID,
					SecondaryRef: firstNonEmpty(item.RelatedIncidentRefs...),
				},
				IncidentRef:       firstNonEmpty(item.RelatedIncidentRefs...),
				RecommendationRef: item.RecommendationID,
				EvidenceRefs:      limitStrings(item.EvidenceRefs, 8),
				PersonaHints:      securityTimelinePersonaHints("recommendation", commandCenterRecommendationSeverity(item)),
			},
		})
	}

	runtimeFilter := runtimeIntegrityFilter{
		event:       contextFilter,
		ClusterID:   contextFilter.ClusterID,
		TenantID:    contextFilter.TenantID,
		Environment: contextFilter.Environment,
		Repo:        contextFilter.Repo,
		Limit:       maxInt(limit*8, 80),
	}
	runtimeFindings, _, err := s.buildRuntimeFindings(ctx, runtimeFilter)
	if err != nil {
		return commandCenterSearchResponse{}, err
	}
	for _, finding := range runtimeFindings {
		score := commandCenterMatchScore(trimmedQuery,
			finding.FindingID,
			finding.FindingType,
			finding.SubjectRef,
			finding.Summary,
			finding.RecommendedAction,
		)
		if score == 0 {
			continue
		}
		timestamp := time.Unix(0, 0).UTC()
		candidates = append(candidates, commandCenterSearchCandidate{
			score:     score,
			timestamp: timestamp,
			result: commandCenterSearchResult{
				SchemaVersion:   commandSearchResultSchemaVersion,
				ResultID:        "runtime_finding:" + finding.FindingID,
				ResultType:      "runtime_finding",
				Title:           finding.FindingType,
				Summary:         finding.Summary,
				Subtitle:        finding.SubjectRef,
				SourceSubsystem: "runtime",
				Severity:        finding.Severity,
				Target: commandCenterFocusTarget{
					Tab:          "runtime",
					Kind:         "runtime_finding",
					Ref:          finding.FindingID,
					SecondaryRef: finding.SubjectRef,
				},
				EvidenceRefs: limitStrings(finding.EvidenceRefs, 8),
				PersonaHints: securityTimelinePersonaHints("runtime", finding.Severity),
			},
		})
	}
	runtimeStates, _, err := s.buildRuntimeIntegrityStates(ctx, runtimeFilter)
	if err != nil {
		return commandCenterSearchResponse{}, err
	}
	for _, state := range runtimeStates {
		score := commandCenterMatchScore(trimmedQuery, state.SubjectRef, strings.Join(state.ActiveFindings, " "))
		if score == 0 {
			continue
		}
		candidates = append(candidates, commandCenterSearchCandidate{
			score:     score,
			timestamp: state.LastVerifiedAt.UTC(),
			result: commandCenterSearchResult{
				SchemaVersion:   commandSearchResultSchemaVersion,
				ResultID:        "runtime_subject:" + state.SubjectRef,
				ResultType:      "runtime_subject",
				Title:           state.SubjectRef,
				Summary:         fmt.Sprintf("Runtime integrity score %d with drift %s and posture %s.", state.RuntimeIntegrityScore, state.DriftLevel, state.CurrentEnforcementPosture),
				Subtitle:        state.CurrentSandboxClass,
				SourceSubsystem: "runtime",
				Severity:        state.DriftLevel,
				Target: commandCenterFocusTarget{
					Tab:  "runtime",
					Kind: "runtime_subject",
					Ref:  state.SubjectRef,
				},
				EvidenceRefs: limitStrings(state.EvidenceRefs, 8),
				PersonaHints: securityTimelinePersonaHints("runtime", state.DriftLevel),
			},
		})
	}

	validationFilter := validationHarnessFilter{
		event:       contextFilter,
		ClusterID:   contextFilter.ClusterID,
		TenantID:    contextFilter.TenantID,
		Environment: contextFilter.Environment,
		Repo:        contextFilter.Repo,
		Limit:       maxInt(limit*4, 24),
	}
	validationRuns, _, err := s.listStrictValidationRuns(ctx, validationFilter)
	if err != nil {
		return commandCenterSearchResponse{}, err
	}
	for _, run := range validationRuns {
		score := commandCenterMatchScore(trimmedQuery, run.RunID, run.Scope, run.Mode, run.Certificate.OverallStatus)
		if score == 0 {
			continue
		}
		candidates = append(candidates, commandCenterSearchCandidate{
			score:     score,
			timestamp: run.Certificate.IssuedAt.UTC(),
			result: commandCenterSearchResult{
				SchemaVersion:   commandSearchResultSchemaVersion,
				ResultID:        "validation_run:" + run.RunID,
				ResultType:      "validation_run",
				Title:           run.Scope,
				Summary:         fmt.Sprintf("Validation run %s finished with certificate %s across %d scenario verdicts.", run.Mode, run.Certificate.OverallStatus, len(run.Verdicts)),
				Subtitle:        run.RunID,
				SourceSubsystem: "validation",
				Severity:        commandCenterValidationSeverity(run.Certificate.OverallStatus),
				Target: commandCenterFocusTarget{
					Tab:  "validation",
					Kind: "validation_run",
					Ref:  run.RunID,
				},
				EvidenceRefs: limitStrings(run.Certificate.EvidenceRefs, 8),
				PersonaHints: securityTimelinePersonaHints("validation", commandCenterValidationSeverity(run.Certificate.OverallStatus)),
			},
		})
	}

	federationView, err := s.buildFederationGlobalView(ctx)
	if err != nil {
		return commandCenterSearchResponse{}, err
	}
	for _, peer := range federationView.Peers {
		score := commandCenterMatchScore(trimmedQuery, peer.PeerID, peer.Organization, peer.Region, peer.TrustDomain, peer.PolicyRole)
		if score == 0 {
			continue
		}
		candidates = append(candidates, commandCenterSearchCandidate{
			score:     score,
			timestamp: peer.LastSeen.UTC(),
			result: commandCenterSearchResult{
				SchemaVersion:   commandSearchResultSchemaVersion,
				ResultID:        "federation_peer:" + peer.PeerID,
				ResultType:      "federation_peer",
				Title:           peer.Organization,
				Summary:         fmt.Sprintf("Peer %s is %s with %s policy role and %s disclosure mode.", peer.PeerID, peer.Status, peer.PolicyRole, peer.DisclosureMode),
				Subtitle:        firstNonEmpty(peer.Region, peer.TrustDomain, peer.Cluster),
				SourceSubsystem: "federation",
				Severity:        commandCenterFederationPeerSeverity(peer, federationView.StalePeers),
				Target: commandCenterFocusTarget{
					Tab:  "federation",
					Kind: "federation_peer",
					Ref:  peer.PeerID,
				},
				EvidenceRefs: limitStrings(peer.TrustState.TrustAnchorFingerprints, 4),
				PersonaHints: securityTimelinePersonaHints("federation", commandCenterFederationPeerSeverity(peer, federationView.StalePeers)),
			},
		})
	}

	handoffEvents, err := s.store.ListEvents(ctx, audit.EventFilter{
		ClusterID:   contextFilter.ClusterID,
		TenantID:    contextFilter.TenantID,
		Environment: contextFilter.Environment,
		Repo:        contextFilter.Repo,
		Component:   handoffComponent,
		Limit:       maxInt(limit*10, 120),
	})
	if err != nil {
		return commandCenterSearchResponse{}, err
	}
	for _, event := range handoffEvents {
		record := parseSecurityTimelineHandoff(event.Handoff)
		if record == nil {
			continue
		}
		score := commandCenterMatchScore(trimmedQuery, record.PackageID, record.ManifestHash, record.Bundle.ManifestHash, record.Manifest.Scope.SelectionSummary)
		if score == 0 {
			continue
		}
		candidates = append(candidates, commandCenterSearchCandidate{
			score:     score,
			timestamp: eventTimestamp(event).UTC(),
			result: commandCenterSearchResult{
				SchemaVersion:   commandSearchResultSchemaVersion,
				ResultID:        "handoff_package:" + record.PackageID,
				ResultType:      "handoff_package",
				Title:           record.PackageID,
				Summary:         fmt.Sprintf("Sealed handoff %s with %s verification and manifest %s.", record.PackageType, record.Verification.OverallStatus, record.ManifestHash),
				Subtitle:        record.Manifest.Scope.SelectionSummary,
				SourceSubsystem: "handoff",
				Severity:        commandCenterHandoffSeverity(record.Verification.OverallStatus),
				Target: commandCenterFocusTarget{
					Tab:         "overview",
					Kind:        "handoff_package",
					Ref:         record.PackageID,
					ResourceURI: fmt.Sprintf("/v1/handoff/%s", record.PackageID),
				},
				EvidenceRefs: limitStrings(record.Manifest.EvidenceRefs, 8),
				PersonaHints: securityTimelinePersonaHints("handoff", commandCenterHandoffSeverity(record.Verification.OverallStatus)),
				Limitations: []string{
					"Sealed handoff search results point at package metadata and verification state; downstream disclosure remains bounded by package audience and redaction profile.",
				},
			},
		})
	}

	intelligenceEvents, err := s.store.ListEvents(ctx, audit.EventFilter{
		ClusterID:   contextFilter.ClusterID,
		TenantID:    contextFilter.TenantID,
		Environment: contextFilter.Environment,
		Repo:        contextFilter.Repo,
		Component:   phase3IntelligenceComponent,
		Limit:       maxInt(limit*10, 120),
	})
	if err != nil {
		return commandCenterSearchResponse{}, err
	}
	for _, event := range intelligenceEvents {
		payload := parsePhase3IntelligencePayload(event.Intelligence)
		switch {
		case payload.Vulnerability != nil:
			item := payload.Vulnerability
			score := commandCenterMatchScore(trimmedQuery, item.VerdictID, item.SubjectRef, item.VulnerabilityID, item.PackageName, item.ImageDigest, item.PURL)
			if score == 0 {
				continue
			}
			candidates = append(candidates, commandCenterSearchCandidate{
				score:     score,
				timestamp: item.EvaluatedAt.UTC(),
				result: commandCenterSearchResult{
					SchemaVersion:   commandSearchResultSchemaVersion,
					ResultID:        "vulnerability_relevance:" + item.VerdictID,
					ResultType:      "vulnerability_relevance",
					Title:           firstNonEmpty(item.VulnerabilityID, item.PackageName, item.VerdictID),
					Summary:         firstNonEmpty(firstString(item.Explanation.Derived), firstString(item.Explanation.Recommended), "Bounded vulnerability relevance verdict is available."),
					Subtitle:        firstNonEmpty(item.SubjectRef, item.PackageName, item.ImageDigest),
					SourceSubsystem: "intelligence",
					Severity:        commandCenterIntelligenceSeverity(item.CurrentState),
					Target: commandCenterFocusTarget{
						Tab:          "vulnerabilities",
						Kind:         "vulnerability_relevance",
						Ref:          item.VerdictID,
						SecondaryRef: item.VulnerabilityID,
						ResourceURI:  fmt.Sprintf("/v1/intelligence/vulnerability-relevance?subject_ref=%s&vulnerability_id=%s", item.SubjectRef, item.VulnerabilityID),
					},
					EvidenceRefs: limitStrings(append([]string{}, item.EvidenceRefs...), 8),
					PersonaHints: []string{"developer", "security_engineer", "executive"},
				},
			})
		case payload.SupplyChain != nil:
			item := payload.SupplyChain
			score := commandCenterMatchScore(trimmedQuery, item.PatternID, item.SubjectRef, item.PackageName, item.PackageVersion, strings.Join(item.ReasonCodes, " "))
			if score == 0 {
				continue
			}
			candidates = append(candidates, commandCenterSearchCandidate{
				score:     score,
				timestamp: item.ObservedAt.UTC(),
				result: commandCenterSearchResult{
					SchemaVersion:   commandSearchResultSchemaVersion,
					ResultID:        "supply_chain_pattern:" + item.PatternID,
					ResultType:      "supply_chain_pattern",
					Title:           firstNonEmpty(item.PackageName, item.PatternID),
					Summary:         firstNonEmpty(firstString(item.Explanation.Derived), "Bounded supply-chain pattern verdict is available."),
					Subtitle:        firstNonEmpty(item.SubjectRef, item.PackageVersion),
					SourceSubsystem: "intelligence",
					Severity:        commandCenterSupplyChainSeverity(item.CurrentState),
					Target: commandCenterFocusTarget{
						Tab:          "guidance",
						Kind:         "supply_chain_pattern",
						Ref:          item.PatternID,
						SecondaryRef: item.PackageName,
						ResourceURI:  fmt.Sprintf("/v1/intelligence/supply-chain/patterns?subject_ref=%s&package_name=%s", item.SubjectRef, item.PackageName),
					},
					EvidenceRefs: limitStrings(item.EvidenceRefs, 8),
					PersonaHints: []string{"developer", "security_engineer", "auditor"},
				},
			})
		case payload.Strategic != nil:
			item := payload.Strategic
			score := commandCenterMatchScore(trimmedQuery, item.AssessmentID, item.SubjectRef, strings.Join(item.RecommendedActions, " "), item.Recommendation.Action)
			if score == 0 {
				continue
			}
			candidates = append(candidates, commandCenterSearchCandidate{
				score:     score,
				timestamp: item.EvaluatedAt.UTC(),
				result: commandCenterSearchResult{
					SchemaVersion:   commandSearchResultSchemaVersion,
					ResultID:        "strategic_assessment:" + item.AssessmentID,
					ResultType:      "strategic_assessment",
					Title:           "Strategic advisory",
					Summary:         firstNonEmpty(item.Recommendation.Summary, firstString(item.RecommendedActions), "Bounded strategic assessment is available."),
					Subtitle:        firstNonEmpty(item.SubjectRef, item.Recommendation.Action),
					SourceSubsystem: "intelligence",
					Severity:        strings.ToLower(item.Recommendation.PriorityBand),
					Target: commandCenterFocusTarget{
						Tab:          "guidance",
						Kind:         "strategic_assessment",
						Ref:          item.AssessmentID,
						SecondaryRef: item.SubjectRef,
						ResourceURI:  fmt.Sprintf("/v1/intelligence/strategic/query?subject_ref=%s", item.SubjectRef),
					},
					EvidenceRefs: limitStrings(item.EvidenceRefs, 8),
					PersonaHints: []string{"security_engineer", "platform_operator", "executive"},
				},
			})
		case payload.Query != nil:
			item := payload.Query
			score := commandCenterMatchScore(trimmedQuery, item.QueryID, item.Query, item.Scope.SubjectRef, item.Scope.VulnerabilityID, item.Scope.PackageName)
			if score == 0 {
				continue
			}
			candidates = append(candidates, commandCenterSearchCandidate{
				score:     score,
				timestamp: item.AnsweredAt.UTC(),
				result: commandCenterSearchResult{
					SchemaVersion:   commandSearchResultSchemaVersion,
					ResultID:        "grounded_query:" + item.QueryID,
					ResultType:      "grounded_query",
					Title:           firstNonEmpty(item.Query, item.QueryID),
					Summary:         firstNonEmpty(firstString(item.RecommendedActions), firstString(item.InferredConclusions), "Grounded security query response is available."),
					Subtitle:        firstNonEmpty(item.Scope.SubjectRef, item.Scope.VulnerabilityID, item.Scope.PackageName),
					SourceSubsystem: "intelligence",
					Severity:        "low",
					Target: commandCenterFocusTarget{
						Tab:          "guidance",
						Kind:         "grounded_query",
						Ref:          item.QueryID,
						SecondaryRef: item.Scope.SubjectRef,
						ResourceURI:  fmt.Sprintf("/v1/intelligence/strategic/query?subject_ref=%s&vulnerability_id=%s&package_name=%s", item.Scope.SubjectRef, item.Scope.VulnerabilityID, item.Scope.PackageName),
					},
					EvidenceRefs: limitStrings(item.EvidenceRefs, 8),
					PersonaHints: []string{"developer", "security_engineer", "auditor"},
				},
			})
		}
	}

	enterpriseEvents, err := s.store.ListEvents(ctx, audit.EventFilter{
		TenantID:    contextFilter.TenantID,
		Environment: contextFilter.Environment,
		Repo:        contextFilter.Repo,
		Component:   phase4EnterpriseComponent,
		Limit:       maxInt(limit*10, 120),
	})
	if err != nil {
		return commandCenterSearchResponse{}, err
	}
	for _, event := range enterpriseEvents {
		payload := parsePhase4EnterprisePayload(event.Enterprise)
		switch {
		case payload.Workflow != nil:
			item := payload.Workflow
			score := commandCenterMatchScore(trimmedQuery, item.WorkflowID, item.SubjectRef, item.CurrentState, item.CanonicalState, item.Routing.PrimaryOwner, strings.Join(item.ReasonCodes, " "))
			if score == 0 {
				continue
			}
			candidates = append(candidates, commandCenterSearchCandidate{
				score:     score,
				timestamp: item.ObservedAt.UTC(),
				result: commandCenterSearchResult{
					SchemaVersion:   commandSearchResultSchemaVersion,
					ResultID:        "workflow_record:" + item.WorkflowID,
					ResultType:      "workflow_record",
					Title:           firstNonEmpty(item.WorkflowID, item.SubjectRef),
					Summary:         firstNonEmpty(firstString(item.ReasonCodes), "Enterprise workflow lifecycle record is available."),
					Subtitle:        firstNonEmpty(item.Routing.PrimaryOwner, item.SubjectRef, item.CurrentState),
					SourceSubsystem: "workflow",
					Severity:        commandCenterWorkflowSeverity(item.CurrentState, item.CanonicalState),
					Target: commandCenterFocusTarget{
						Tab:          "exceptions",
						Kind:         "workflow_record",
						Ref:          item.WorkflowID,
						SecondaryRef: item.SubjectRef,
						ResourceURI:  fmt.Sprintf("/v1/enterprise/workflow/lifecycle?subject_ref=%s&workflow_id=%s", item.SubjectRef, item.WorkflowID),
					},
					EvidenceRefs: limitStrings(item.EvidenceRefs, 8),
					PersonaHints: []string{"platform_operator", "security_engineer", "executive"},
				},
			})
		case payload.Reconciliation != nil:
			item := payload.Reconciliation
			score := commandCenterMatchScore(trimmedQuery, item.WorkflowID, item.SubjectRef, item.ConnectorRef, item.CurrentState, item.ConnectorSystem, strings.Join(item.ReasonCodes, " "))
			if score == 0 {
				continue
			}
			candidates = append(candidates, commandCenterSearchCandidate{
				score:     score,
				timestamp: item.ObservedAt.UTC(),
				result: commandCenterSearchResult{
					SchemaVersion:   commandSearchResultSchemaVersion,
					ResultID:        "workflow_reconciliation:" + item.WorkflowID,
					ResultType:      "workflow_reconciliation",
					Title:           firstNonEmpty(item.WorkflowID, item.ConnectorRef),
					Summary:         firstNonEmpty(firstString(item.ReasonCodes), "Connector reconciliation record is available."),
					Subtitle:        firstNonEmpty(item.ConnectorSystem, item.ConnectorRef, item.SubjectRef),
					SourceSubsystem: "workflow",
					Severity:        commandCenterConnectorSeverity(item.CurrentState),
					Target: commandCenterFocusTarget{
						Tab:          "exceptions",
						Kind:         "workflow_reconciliation",
						Ref:          item.WorkflowID,
						SecondaryRef: item.ConnectorRef,
						ResourceURI:  fmt.Sprintf("/v1/enterprise/workflow/connectors/reconcile?subject_ref=%s&workflow_id=%s", item.SubjectRef, item.WorkflowID),
					},
					EvidenceRefs: limitStrings(item.EvidenceRefs, 8),
					PersonaHints: []string{"platform_operator", "security_engineer", "auditor"},
				},
			})
		case payload.PartnerIntake != nil:
			item := payload.PartnerIntake
			score := commandCenterMatchScore(trimmedQuery, item.PartnerID, item.Organization, item.HandoffRef, item.TrustDomain, strings.Join(item.ReasonCodes, " "))
			if score == 0 {
				continue
			}
			candidates = append(candidates, commandCenterSearchCandidate{
				score:     score,
				timestamp: item.ObservedAt.UTC(),
				result: commandCenterSearchResult{
					SchemaVersion:   commandSearchResultSchemaVersion,
					ResultID:        "partner_trust:" + item.PartnerID,
					ResultType:      "partner_trust",
					Title:           firstNonEmpty(item.Organization, item.PartnerID),
					Summary:         item.Dashboard.TrustSummary,
					Subtitle:        firstNonEmpty(item.HandoffRef, item.TrustDomain, item.PartnerID),
					SourceSubsystem: "partner",
					Severity:        commandCenterPartnerSeverity(item.CurrentState),
					Target: commandCenterFocusTarget{
						Tab:          "federation",
						Kind:         "partner_trust",
						Ref:          item.PartnerID,
						SecondaryRef: item.HandoffRef,
						ResourceURI:  fmt.Sprintf("/v1/enterprise/partner-trust/dashboard?partner_id=%s", item.PartnerID),
					},
					EvidenceRefs: limitStrings(item.EvidenceRefs, 8),
					PersonaHints: []string{"security_engineer", "auditor", "executive"},
				},
			})
		case payload.Compliance != nil:
			item := payload.Compliance
			score := commandCenterMatchScore(trimmedQuery, item.ControlID, item.ControlFamily, item.SubjectRef, item.CoverageState, strings.Join(item.ReasonCodes, " "))
			if score == 0 {
				continue
			}
			candidates = append(candidates, commandCenterSearchCandidate{
				score:     score,
				timestamp: item.ObservedAt.UTC(),
				result: commandCenterSearchResult{
					SchemaVersion:   commandSearchResultSchemaVersion,
					ResultID:        "compliance_mapping:" + item.ControlID,
					ResultType:      "compliance_mapping",
					Title:           firstNonEmpty(item.ControlID, item.ControlFamily),
					Summary:         firstNonEmpty(firstString(item.ReasonCodes), "Compliance mapping record is available."),
					Subtitle:        firstNonEmpty(item.SubjectRef, item.ControlFamily, item.CoverageState),
					SourceSubsystem: "governance",
					Severity:        commandCenterGovernanceSeverity(item.CoverageState),
					Target: commandCenterFocusTarget{
						Tab:          "scorecard",
						Kind:         "compliance_mapping",
						Ref:          item.ControlID,
						SecondaryRef: item.SubjectRef,
						ResourceURI:  fmt.Sprintf("/v1/enterprise/governance/compliance-mapping?subject_ref=%s", item.SubjectRef),
					},
					EvidenceRefs: limitStrings(append(append([]string{}, item.EvidenceRefs...), item.TechnicalEventRefs...), 8),
					PersonaHints: []string{"security_engineer", "auditor", "executive"},
				},
			})
		case payload.PolicyDrift != nil:
			item := payload.PolicyDrift
			score := commandCenterMatchScore(trimmedQuery, item.SubjectRef, item.ExceptionID, item.Actor, item.CurrentMode, item.PreviousMode, strings.Join(item.ImpactSummary, " "))
			if score == 0 {
				continue
			}
			candidates = append(candidates, commandCenterSearchCandidate{
				score:     score,
				timestamp: item.ObservedAt.UTC(),
				result: commandCenterSearchResult{
					SchemaVersion:   commandSearchResultSchemaVersion,
					ResultID:        "policy_drift:" + firstNonEmpty(item.ExceptionID, item.SubjectRef, item.Actor),
					ResultType:      "policy_drift",
					Title:           firstNonEmpty(item.ExceptionID, item.SubjectRef, "policy drift"),
					Summary:         firstNonEmpty(firstString(item.ImpactSummary), item.ChangeReason, "Policy drift record is available."),
					Subtitle:        firstNonEmpty(item.Actor, item.SubjectRef, item.CurrentMode),
					SourceSubsystem: "governance",
					Severity:        commandCenterDriftSeverity(item.CurrentState),
					Target: commandCenterFocusTarget{
						Tab:          "analytics",
						Kind:         "policy_drift",
						Ref:          firstNonEmpty(item.ExceptionID, item.SubjectRef, item.Actor),
						SecondaryRef: item.SubjectRef,
						ResourceURI:  fmt.Sprintf("/v1/enterprise/governance/policy-drift?subject_ref=%s", item.SubjectRef),
					},
					EvidenceRefs: limitStrings(item.EvidenceRefs, 8),
					PersonaHints: []string{"security_engineer", "auditor", "executive"},
				},
			})
		case payload.Executive != nil:
			item := payload.Executive
			score := commandCenterMatchScore(trimmedQuery, item.ScopeRef, strings.Join(item.Highlights, " "), strings.Join(item.EvidenceTraceRefs, " "))
			if score == 0 {
				continue
			}
			candidates = append(candidates, commandCenterSearchCandidate{
				score:     score,
				timestamp: item.GeneratedAt.UTC(),
				result: commandCenterSearchResult{
					SchemaVersion:   commandSearchResultSchemaVersion,
					ResultID:        "executive_report:" + firstNonEmpty(item.ScopeRef, "executive_report"),
					ResultType:      "executive_report",
					Title:           firstNonEmpty(item.ScopeRef, "Executive report"),
					Summary:         firstNonEmpty(firstString(item.Highlights), "Executive governance summary is available."),
					Subtitle:        item.CurrentState,
					SourceSubsystem: "governance",
					Severity:        commandCenterExecutiveSeverity(item.CurrentState),
					Target: commandCenterFocusTarget{
						Tab:         "analytics",
						Kind:        "executive_report",
						Ref:         firstNonEmpty(item.ScopeRef, "executive_report"),
						ResourceURI: fmt.Sprintf("/v1/enterprise/governance/executive-report?scope_ref=%s", item.ScopeRef),
					},
					EvidenceRefs: limitStrings(item.EvidenceTraceRefs, 8),
					PersonaHints: []string{"executive", "security_engineer", "auditor"},
				},
			})
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].score != candidates[j].score {
			return candidates[i].score > candidates[j].score
		}
		if commandCenterSeverityRank(candidates[i].result.Severity) != commandCenterSeverityRank(candidates[j].result.Severity) {
			return commandCenterSeverityRank(candidates[i].result.Severity) > commandCenterSeverityRank(candidates[j].result.Severity)
		}
		if !candidates[i].timestamp.Equal(candidates[j].timestamp) {
			return candidates[i].timestamp.After(candidates[j].timestamp)
		}
		return candidates[i].result.ResultID < candidates[j].result.ResultID
	})

	results := make([]commandCenterSearchResult, 0, minInt(len(candidates), limit))
	seen := map[string]struct{}{}
	for _, candidate := range candidates {
		if _, ok := seen[candidate.result.ResultID]; ok {
			continue
		}
		seen[candidate.result.ResultID] = struct{}{}
		results = append(results, candidate.result)
		if len(results) >= limit {
			break
		}
	}

	return commandCenterSearchResponse{
		SchemaVersion: commandSearchResponseSchemaVersion,
		Query:         trimmedQuery,
		Results:       results,
		Limitations: []string{
			"Command-center search is semantic and bounded to evidence-backed incidents, runtime, validation, intelligence, workflow, partner, governance, federation, and sealed handoff objects; it is not a generic full-text index.",
			"Deep links route to existing operator surfaces or exact package metadata targets without introducing a new command-center truth store.",
		},
	}, nil
}

func commandCenterMatchScore(query string, values ...string) int {
	needle := strings.ToLower(strings.TrimSpace(query))
	if needle == "" {
		return 0
	}
	tokens := strings.Fields(needle)
	best := 0
	for _, raw := range values {
		value := strings.ToLower(strings.TrimSpace(raw))
		if value == "" {
			continue
		}
		switch {
		case value == needle:
			best = maxInt(best, 140)
		case strings.HasPrefix(value, needle):
			best = maxInt(best, 110)
		case strings.Contains(value, needle):
			best = maxInt(best, 85)
		}
		if len(tokens) > 1 {
			matched := 0
			for _, token := range tokens {
				if strings.Contains(value, token) {
					matched++
				}
			}
			if matched == len(tokens) {
				best = maxInt(best, 70+matched*5)
			}
		}
	}
	return best
}

func commandCenterRecommendationSeverity(item recommendation) string {
	switch recommendationPriorityBand(item.PriorityBand) {
	case "NOW":
		return "high"
	case "TODAY":
		return "medium"
	default:
		return "low"
	}
}

func commandCenterValidationSeverity(status string) string {
	switch strings.ToLower(strings.TrimSpace(status)) {
	case validationStatusFail, validationStatusFlaky:
		return "high"
	case validationStatusPartial:
		return "medium"
	default:
		return "low"
	}
}

func commandCenterFederationPeerSeverity(peer federationPeer, stalePeers []string) string {
	if containsString(stalePeers, peer.PeerID) {
		return "high"
	}
	if strings.Contains(strings.ToLower(strings.TrimSpace(peer.Status)), "stale") || strings.Contains(strings.ToLower(strings.TrimSpace(peer.Status)), "degraded") {
		return "medium"
	}
	return "low"
}

func commandCenterHandoffSeverity(status string) string {
	switch strings.ToLower(strings.TrimSpace(status)) {
	case handoffVerificationInvalid:
		return "high"
	case handoffVerificationPartial:
		return "medium"
	default:
		return "low"
	}
}

func commandCenterIntelligenceSeverity(state string) string {
	switch strings.ToLower(strings.TrimSpace(state)) {
	case "active_priority_candidate", "reachable_and_externally_exposed", "cross_cluster_concern_active", "suspected_typo_squat":
		return "high"
	case "reachable_but_low_exploitability", "trust_drift_observed", "suspicious_publication_pattern", "under_review", "strategic_advisory_ready":
		return "medium"
	default:
		return "low"
	}
}

func commandCenterSupplyChainSeverity(state string) string {
	switch strings.ToLower(strings.TrimSpace(state)) {
	case "cross_cluster_concern_active", "suspected_typo_squat":
		return "high"
	case "trust_drift_observed", "suspicious_publication_pattern", "under_review":
		return "medium"
	default:
		return "low"
	}
}

func commandCenterWorkflowSeverity(currentState, canonicalState string) string {
	switch firstNonEmpty(strings.ToLower(strings.TrimSpace(canonicalState)), strings.ToLower(strings.TrimSpace(currentState))) {
	case "rejected", "under_validation":
		return "high"
	case "exception_active", "reopened", "assigned":
		return "medium"
	default:
		return "low"
	}
}

func commandCenterConnectorSeverity(state string) string {
	switch strings.ToLower(strings.TrimSpace(state)) {
	case "external_closure_pending_validation", "reopened_for_validation":
		return "high"
	case "connector_degraded_core_preserved", "awaiting_external_reconciliation":
		return "medium"
	default:
		return "low"
	}
}

func commandCenterPartnerSeverity(state string) string {
	switch strings.ToLower(strings.TrimSpace(state)) {
	case "rejected", "expired":
		return "high"
	case "under_review", "superseded", "received":
		return "medium"
	default:
		return "low"
	}
}

func commandCenterGovernanceSeverity(coverageState string) string {
	switch strings.ToLower(strings.TrimSpace(coverageState)) {
	case "missing":
		return "high"
	case "partial", "inferred":
		return "medium"
	default:
		return "low"
	}
}

func commandCenterDriftSeverity(state string) string {
	switch strings.ToLower(strings.TrimSpace(state)) {
	case "policy_softened":
		return "high"
	case "under_review", "policy_strengthened":
		return "medium"
	default:
		return "low"
	}
}

func commandCenterExecutiveSeverity(state string) string {
	if strings.EqualFold(strings.TrimSpace(state), "executive_governance_attention_required") {
		return "high"
	}
	return "medium"
}

func commandCenterSeverityRank(severity string) int {
	switch strings.ToLower(strings.TrimSpace(severity)) {
	case "critical":
		return 4
	case "high":
		return 3
	case "medium", "warning", "watch":
		return 2
	default:
		return 1
	}
}
