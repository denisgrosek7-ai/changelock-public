import type { ExecutiveDefenseReport, Recommendation } from "../incidents";
import type {
  CommandCenterFocusTarget,
  CommandCenterNotification,
  CommandCenterNotificationsResponse,
  CommandCenterPersona,
  CommandCenterSearchResponse,
  SecurityTimelineEntry,
  SecurityTimelineResponse,
  TabKey,
} from "../types";

type Props = {
  persona: CommandCenterPersona;
  timeline: SecurityTimelineResponse | null;
  notifications: CommandCenterNotificationsResponse | null;
  recommendations: Recommendation[];
  executiveReport: ExecutiveDefenseReport | null;
  loading: boolean;
  searchQuery: string;
  lifecyclePhase: string;
  searchResults: CommandCenterSearchResponse | null;
  searchLoading: boolean;
  focusTarget: CommandCenterFocusTarget | null;
  onPersonaChange: (persona: CommandCenterPersona) => void;
  onOpenTab: (tab: TabKey) => void;
  onOpenTarget: (target: CommandCenterFocusTarget) => void;
  onLifecyclePhaseChange: (value: string) => void;
  onSearchQueryChange: (value: string) => void;
  onSearchSubmit: () => void;
};

type PersonaConfig = {
  label: string;
  summary: string;
  tabs: Array<{ tab: TabKey; title: string; description: string; sources: string[] }>;
  recommendationSources: string[];
};

const personaConfig: Record<CommandCenterPersona, PersonaConfig> = {
  developer: {
    label: "Developer",
    summary: "Focus on deny reasons, repo-linked trust drift, and the next verification-safe fix path.",
    tabs: [
      { tab: "events", title: "Deploy and incident flow", description: "See why the trust gate blocked or reopened a path.", sources: ["deploy", "incident", "recommendation"] },
      { tab: "validation", title: "Validation and replay", description: "Check whether the proposed fix closes the control gap.", sources: ["validation"] },
      { tab: "runtime", title: "Runtime drift", description: "Inspect workload drift that now needs a code or image fix.", sources: ["runtime", "hardening"] },
    ],
    recommendationSources: ["incident", "runtime_signal", "validation_signal"],
  },
  security_engineer: {
    label: "Security Engineer",
    summary: "Track active trust failures, bounded remediation paths, and whether validation or federation state changes widen risk.",
    tabs: [
      { tab: "events", title: "Incidents and approvals", description: "Own the current triage and approval-bound response path.", sources: ["incident", "recommendation", "deploy"] },
      { tab: "federation", title: "Federation trust", description: "Review peer proof reuse, freshness, and divergence handling.", sources: ["federation"] },
      { tab: "validation", title: "Validation pressure", description: "Confirm that regressions or chaos rehearsal stay bounded and explainable.", sources: ["validation"] },
    ],
    recommendationSources: ["incident", "validation_signal", "federation_signal", "hardening_signal", "runtime_signal", "forensic_signal"],
  },
  platform_operator: {
    label: "Platform / Operator",
    summary: "Watch runtime containment, blast radius, and whether recovery is safe before traffic or workload posture changes.",
    tabs: [
      { tab: "runtime", title: "Runtime posture", description: "Containment, hardening, and recovery live here.", sources: ["runtime", "hardening"] },
      { tab: "topology", title: "Blast radius", description: "See the current topology-aware impact before stronger isolation.", sources: ["incident", "runtime", "hardening"] },
      { tab: "forensics", title: "Forensics", description: "Preserve and reconstruct volatile context before destructive action.", sources: ["runtime", "hardening", "incident"] },
    ],
    recommendationSources: ["runtime_signal", "hardening_signal", "forensic_signal", "incident"],
  },
  auditor: {
    label: "Auditor",
    summary: "Follow evidence lineage, sealed packages, validation certificates, and remote proof decisions without losing scope or limitations.",
    tabs: [
      { tab: "events", title: "Incident evidence", description: "Open the investigation workbench and linked package outputs.", sources: ["incident", "handoff", "recommendation"] },
      { tab: "validation", title: "Validation certificates", description: "Review controlled validation runs and scope-bound confidence output.", sources: ["validation"] },
      { tab: "federation", title: "Proof portability", description: "Inspect peer trust, local acceptance, and disclosure-bounded exchange.", sources: ["handoff", "federation"] },
    ],
    recommendationSources: ["validation_signal", "federation_signal", "package"],
  },
  executive: {
    label: "Executive",
    summary: "Stay on material risk, trend direction, and the bounded operator action path that changes business exposure.",
    tabs: [
      { tab: "analytics", title: "Trust health trends", description: "Open measured governance and operating pressure indicators.", sources: ["incident", "recommendation"] },
      { tab: "events", title: "Material incidents", description: "Inspect the operator-owned incidents that drive current posture.", sources: ["incident", "hardening", "runtime"] },
      { tab: "scorecard", title: "Measured trust score", description: "Review the current trust scorecard and audit mapping.", sources: ["handoff", "validation", "federation"] },
    ],
    recommendationSources: ["incident", "hardening_signal", "validation_signal", "federation_signal"],
  },
};

function sourceCount(timeline: SecurityTimelineResponse | null, sources: string[]) {
  if (!timeline) {
    return 0;
  }
  return sources.reduce((total, source) => total + (timeline.counts_by_source[source] || 0), 0);
}

function severityClass(value: string) {
  return value === "critical" || value === "high" ? "deny" : value === "medium" ? "warning" : "muted";
}

const lifecycleLabels: Record<string, string> = {
  build_verify: "Build / Verify",
  runtime: "Runtime",
  validation: "Validation",
  intelligence: "Intelligence",
  workflow: "Workflow",
  partner: "Partner",
  governance: "Governance",
};

function recommendationStatusClass(value: Recommendation["status"]) {
  if (value === "verified_successful") {
    return "allow";
  }
  if (value === "regressed" || value === "executed_no_effect" || value === "rejected") {
    return "deny";
  }
  return "warning";
}

function formatTimestamp(value: string) {
  return new Date(value).toLocaleString();
}

function timelineMatchesPersona(entry: SecurityTimelineEntry, persona: CommandCenterPersona) {
  return entry.persona_hints?.includes(persona) ?? false;
}

function notificationMatchesPersona(entry: CommandCenterNotification, persona: CommandCenterPersona) {
  return entry.persona_hints?.includes(persona) ?? false;
}

function topTimelineForPersona(timeline: SecurityTimelineResponse | null, persona: CommandCenterPersona) {
  if (!timeline) {
    return [];
  }
  const direct = timeline.entries.filter((entry) => timelineMatchesPersona(entry, persona));
  if (direct.length > 0) {
    return direct.slice(0, 8);
  }
  if (persona === "executive") {
    return timeline.entries.filter((entry) => entry.severity === "critical" || entry.severity === "high").slice(0, 8);
  }
  return timeline.entries.slice(0, 8);
}

function topNotificationsForPersona(notifications: CommandCenterNotificationsResponse | null, persona: CommandCenterPersona) {
  if (!notifications) {
    return [];
  }
  const direct = notifications.items.filter((entry) => notificationMatchesPersona(entry, persona));
  return (direct.length > 0 ? direct : notifications.items).slice(0, 6);
}

function topRecommendationsForPersona(recommendations: Recommendation[], persona: CommandCenterPersona) {
  const config = personaConfig[persona];
  const filtered = recommendations.filter((item) => config.recommendationSources.includes(item.sourceType));
  return (filtered.length > 0 ? filtered : recommendations).slice(0, 4);
}

function recommendationMatchesFocus(item: Recommendation, focusTarget: CommandCenterFocusTarget | null) {
  return focusTarget?.kind === "recommendation" && focusTarget.ref === item.recommendationID;
}

function searchResultMatchesFocus(result: NonNullable<CommandCenterSearchResponse["results"]>[number], focusTarget: CommandCenterFocusTarget | null) {
  return (
    focusTarget?.tab === result.target.tab &&
    focusTarget.kind === result.target.kind &&
    focusTarget.ref === result.target.ref
  );
}

export function CommandCenterPanel({
  persona,
  timeline,
  notifications,
  recommendations,
  executiveReport,
  loading,
  searchQuery,
  lifecyclePhase,
  searchResults,
  searchLoading,
  focusTarget,
  onPersonaChange,
  onOpenTab,
  onOpenTarget,
  onLifecyclePhaseChange,
  onSearchQueryChange,
  onSearchSubmit,
}: Props) {
  const config = personaConfig[persona];
  const visibleTimeline = topTimelineForPersona(timeline, persona);
  const visibleNotifications = topNotificationsForPersona(notifications, persona);
  const visibleRecommendations = topRecommendationsForPersona(recommendations, persona).sort((left, right) => {
    const leftFocused = recommendationMatchesFocus(left, focusTarget);
    const rightFocused = recommendationMatchesFocus(right, focusTarget);
    if (leftFocused !== rightFocused) {
      return leftFocused ? -1 : 1;
    }
    return left.recommendationID.localeCompare(right.recommendationID);
  });

  return (
    <section className="command-center">
      <article className="panel command-center__hero">
        <div>
          <span className="summary-label">Command center</span>
          <h2>{config.label} view</h2>
          <p>{config.summary}</p>
          {persona === "executive" && executiveReport ? (
            <p className="command-center__hero-note">
              {executiveReport.executiveSummary.whatMattersNow} · shield health {executiveReport.shieldHealth.score}/100
            </p>
          ) : null}
        </div>
        <div className="persona-switch" role="tablist" aria-label="Command center personas">
          {Object.entries(personaConfig).map(([key, value]) => (
            <button
              key={key}
              type="button"
              className={`persona-switch__button ${persona === key ? "is-active" : ""}`}
              onClick={() => onPersonaChange(key as CommandCenterPersona)}
            >
              {value.label}
            </button>
          ))}
        </div>
      </article>

      <article className="panel command-center__search">
        <div className="command-center__search-header">
          <div>
            <span className="summary-label">Contextual command bar</span>
            <strong>Semantic lookup across canonical subject refs, CVEs, packages, workflow ids, partner refs, and evidence-linked trust objects</strong>
          </div>
        </div>
        <div className="command-center__search-form">
          <input
            type="search"
            value={searchQuery}
            placeholder="Search subject ref, CVE, package, workflow id, incident id, partner handoff ref, or evidence-linked object"
            onChange={(event) => onSearchQueryChange(event.target.value)}
            onKeyDown={(event) => {
              if (event.key === "Enter") {
                event.preventDefault();
                onSearchSubmit();
              }
            }}
          />
          <button type="button" className="button" disabled={searchLoading} onClick={onSearchSubmit}>
            {searchLoading ? "Searching…" : "Search"}
          </button>
        </div>
        {searchResults ? (
          searchResults.results.length > 0 ? (
            <div className="command-center__search-results">
              {searchResults.results.map((result) => (
                <article
                  className={`command-center__search-result ${searchResultMatchesFocus(result, focusTarget) ? "is-selected" : ""}`}
                  key={result.result_id}
                >
                  <div>
                    <div className="overview-list-item__title">
                      <strong>{result.title}</strong>
                      <span className={`chip chip--${severityClass(result.severity)}`}>{result.severity}</span>
                      <span className="chip chip--muted">{result.result_type.replace(/_/g, " ")}</span>
                    </div>
                    <p>{result.summary}</p>
                    {result.subtitle ? <small>{result.subtitle}</small> : null}
                    {result.target.resource_uri ? <small>resource {result.target.resource_uri}</small> : null}
                  </div>
                  <div className="command-center__timeline-footer">
                    {result.evidence_refs?.length ? (
                      <span className="chip chip--muted">{result.evidence_refs.length} evidence refs</span>
                    ) : null}
                    <button type="button" className="button button-secondary" onClick={() => onOpenTarget(result.target)}>
                      Open
                    </button>
                  </div>
                </article>
              ))}
            </div>
          ) : (
            <p className="panel-empty">No semantic command-center results matched the current scope.</p>
          )
        ) : (
          <p className="panel-empty">Search stays bounded and returns exact trust objects rather than generic full-text fragments.</p>
        )}
      </article>

      <section className="command-center__lanes">
        {config.tabs.map((lane) => (
          <article className="panel command-center__lane" key={`${persona}-${lane.tab}`}>
            <div className="command-center__lane-header">
              <div>
                <span className="summary-label">{lane.title}</span>
                <strong>{sourceCount(timeline, lane.sources)} timeline signals</strong>
              </div>
              <button type="button" className="button button-secondary" onClick={() => onOpenTab(lane.tab)}>
                Open {lane.tab}
              </button>
            </div>
            <p>{lane.description}</p>
          </article>
        ))}
      </section>

      <section className="command-center__grid">
        <article className="panel analytics-panel">
          <div className="table-toolbar">
            <span className="summary-label">Grouped notifications</span>
            <strong>{visibleNotifications.length}</strong>
          </div>
          <div className="persona-switch" role="tablist" aria-label="Command center lifecycle phases">
            <button
              type="button"
              className={`persona-switch__button ${lifecyclePhase === "" ? "is-active" : ""}`}
              onClick={() => onLifecyclePhaseChange("")}
            >
              All phases
            </button>
            {Object.entries(lifecycleLabels).map(([key, label]) => (
              <button
                key={key}
                type="button"
                className={`persona-switch__button ${lifecyclePhase === key ? "is-active" : ""}`}
                onClick={() => onLifecyclePhaseChange(key)}
              >
                {label} {(timeline?.counts_by_lifecycle[key] || 0) > 0 ? `(${timeline?.counts_by_lifecycle[key] || 0})` : ""}
              </button>
            ))}
          </div>
          {loading && visibleNotifications.length === 0 ? (
            <p className="panel-empty">Grouping stateful notifications…</p>
          ) : visibleNotifications.length === 0 ? (
            <p className="panel-empty">No grouped notifications matched the current scope and lifecycle filter.</p>
          ) : (
            <ul className="analytics-list">
              {visibleNotifications.map((item) => (
                <li className="analytics-list__item overview-list-item" key={item.notification_id}>
                  <div>
                    <div className="overview-list-item__title">
                      <strong>{item.title}</strong>
                      <span className={`chip chip--${severityClass(item.severity)}`}>{item.severity}</span>
                      <span className="chip chip--muted">{item.current_state.replace(/_/g, " ")}</span>
                    </div>
                    <p>{item.summary}</p>
                    <small>
                      {lifecycleLabels[item.lifecycle_phase] || item.lifecycle_phase}
                      {item.owner_hint ? ` · owner ${item.owner_hint}` : ""}
                    </small>
                    {item.next_action ? <p className="incident-inline-copy"><strong>Next:</strong> {item.next_action}</p> : null}
                  </div>
                  <div className="command-center__timeline-footer">
                    {item.evidence_refs?.length ? <span className="chip chip--muted">{item.evidence_refs.length} evidence refs</span> : null}
                    <button type="button" className="button button-secondary" onClick={() => onOpenTarget(item.target)}>
                      Open
                    </button>
                  </div>
                </li>
              ))}
            </ul>
          )}
        </article>

        <article className="panel analytics-panel">
          <div className="table-toolbar">
            <span className="summary-label">Next best actions</span>
            <strong>{visibleRecommendations.length}</strong>
          </div>
          {loading && visibleRecommendations.length === 0 ? (
            <p className="panel-empty">Loading recommendation-linked action flow…</p>
          ) : visibleRecommendations.length === 0 ? (
            <p className="panel-empty">No recommendation-backed next action is currently in scope.</p>
          ) : (
            <ul className="analytics-list">
              {visibleRecommendations.map((item) => (
                <li className="analytics-list__item overview-list-item" key={item.recommendationID}>
                  <div>
                    <div className="overview-list-item__title">
                      <strong>{item.title}</strong>
                      <span className={`chip chip--${recommendationStatusClass(item.status)}`}>{item.status.replace(/_/g, " ")}</span>
                    </div>
                    <p>{item.recommendedAction}</p>
                    <small>{item.rationale}</small>
                  </div>
                  <div className="overview-list-item__aside">
                    {item.approvalMode === "approval_required" ? "Approval required before execution." : "Auto-safe launcher path available."}
                  </div>
                </li>
              ))}
            </ul>
          )}
        </article>

        <article className="panel analytics-panel">
          <div className="table-toolbar">
            <span className="summary-label">Unified security timeline</span>
            <strong>{visibleTimeline.length}</strong>
          </div>
          {loading && visibleTimeline.length === 0 ? (
            <p className="panel-empty">Aggregating current trust timeline…</p>
          ) : visibleTimeline.length === 0 ? (
            <p className="panel-empty">No evidence-backed timeline entries matched the current scope.</p>
          ) : (
            <ol className="incident-timeline command-center__timeline">
              {visibleTimeline.map((entry) => (
                <li className="incident-timeline__item" key={entry.entry_id}>
                  <div className="incident-timeline__meta">
                    <strong>{entry.title}</strong>
                    <span className={`chip chip--${severityClass(entry.severity)}`}>{entry.severity}</span>
                    <span className="chip chip--muted">{lifecycleLabels[entry.lifecycle_phase] || entry.lifecycle_phase}</span>
                    <span className="chip chip--muted">{entry.source_subsystem.replace(/_/g, " ")}</span>
                    <small>{formatTimestamp(entry.timestamp)}</small>
                  </div>
                  <p>{entry.summary}</p>
                  <small>
                    Subject {entry.subject_label}
                    {entry.recommendation_ref ? ` · recommendation ${entry.recommendation_ref}` : ""}
                    {entry.incident_ref ? ` · incident ${entry.incident_ref}` : ""}
                  </small>
                  {entry.next_action ? <p className="incident-inline-copy"><strong>Next:</strong> {entry.next_action}</p> : null}
                  <div className="command-center__timeline-footer">
                    {entry.evidence_refs && entry.evidence_refs.length > 0 ? (
                      <span className="chip chip--muted">{entry.evidence_refs.length} evidence refs</span>
                    ) : (
                      <span className="chip chip--muted">Evidence linked through source subsystem</span>
                    )}
                    {entry.drilldown_tab ? (
                      <button
                        type="button"
                        className="button button-secondary"
                        onClick={() =>
                          entry.drilldown_target_kind && entry.drilldown_target_ref
                            ? onOpenTarget({
                                tab: entry.drilldown_tab!,
                                kind: entry.drilldown_target_kind,
                                ref: entry.drilldown_target_ref,
                                secondary_ref: entry.drilldown_target_secondary_ref,
                                resource_uri: entry.resource_uri,
                              })
                            : onOpenTab(entry.drilldown_tab!)
                        }
                      >
                        {entry.drilldown_label || `Open ${entry.drilldown_tab}`}
                      </button>
                    ) : null}
                  </div>
                </li>
              ))}
            </ol>
          )}
        </article>
      </section>
    </section>
  );
}
