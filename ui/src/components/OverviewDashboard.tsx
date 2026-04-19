import { buildOverviewModel } from "../overview";
import type { ExecutiveDefenseReport, SystemicWeaknessResponse } from "../incidents";
import type {
  AuditHealth,
  DriftStatsResponse,
  ExceptionReport,
  Summary,
  SyncStatus,
  TopViolatorsResponse,
  TrendsResponse,
} from "../types";

type Props = {
  health: AuditHealth | null;
  summary: Summary | null;
  trends: TrendsResponse | null;
  topViolators: TopViolatorsResponse | null;
  driftStats: DriftStatsResponse | null;
  exceptionReport: ExceptionReport | null;
  systemicWeaknesses: SystemicWeaknessResponse | null;
  executiveReport: ExecutiveDefenseReport | null;
  syncStatus: SyncStatus | null;
  loading: boolean;
  onSelectTrustMetric?: (metricKey: string, label: string) => void;
};

function severityClass(value: string) {
  return value === "critical" || value === "high" ? "deny" : value === "medium" ? "warning" : "muted";
}

function priorityClass(value: string) {
  return value === "now" ? "deny" : value === "next" ? "warning" : "muted";
}

function trustStatusClass(value: string) {
  return value === "verified" ? "allow" : value === "partial" ? "warning" : value === "gap" ? "deny" : "muted";
}

function confidenceClass(value: string) {
  return value === "high" ? "allow" : value === "medium" ? "warning" : "muted";
}

function trendClass(value: string) {
  return value === "improving" ? "allow" : value === "worsening" ? "deny" : "warning";
}

function shieldBandClass(value: string) {
  return value === "strong" ? "allow" : value === "watch" ? "warning" : "deny";
}

export function OverviewDashboard({
  health,
  summary,
  trends,
  topViolators,
  driftStats,
  exceptionReport,
  systemicWeaknesses,
  executiveReport,
  syncStatus,
  loading,
  onSelectTrustMetric,
}: Props) {
  if (loading && !summary) {
    return <section className="panel panel-empty">Building the current operator posture…</section>;
  }

  const model = buildOverviewModel({
    health,
    summary,
    trends,
    topViolators,
    driftStats,
    exceptionReport,
    syncStatus,
  });

  return (
    <>
      <section className={`panel posture-hero posture-hero--${model.posture.level}`}>
        <div className="posture-hero__main">
          <span className="summary-label">Current posture</span>
          <h2>{model.posture.title}</h2>
          <p>{model.posture.summary}</p>
          <div className="chip-row">
            <span className={`chip chip--${model.posture.level === "stable" ? "allow" : model.posture.level === "degraded" ? "warning" : "deny"}`}>
              {model.posture.level === "stable" ? "Stable" : model.posture.level === "degraded" ? "Degraded" : "At Risk"}
            </span>
            <span className="chip chip--muted">{model.posture.changed}</span>
          </div>
        </div>

        <div className="posture-hero__support">
          <section>
            <span className="summary-label">Why this status</span>
            <ul className="summary-list summary-list--compact">
              {model.posture.reasons.map((reason) => (
                <li key={reason}>
                  <span>{reason}</span>
                </li>
              ))}
            </ul>
          </section>

          <section>
            <span className="summary-label">Affected scope</span>
            <p className="posture-hero__scope">{model.posture.scope}</p>
          </section>
        </div>
      </section>

      <section className="overview-priority-strip">
        <article className="panel analytics-panel">
          <div className="table-toolbar">
            <span className="summary-label">Trust scorecard</span>
            <strong>{model.trust.grade}</strong>
          </div>
          <div className="overview-scorecard">
            <div className="overview-scorecard__hero">
              <strong className="overview-scorecard__grade">{model.trust.grade}</strong>
              <div>
                <p className="overview-scorecard__score">{model.trust.score}/100 measured trust score</p>
                <p>{model.trust.summary}</p>
              </div>
            </div>

            <div className="overview-scorecard__metrics">
              {model.trust.metrics.map((metric) => (
                <button
                  type="button"
                  className={`overview-scorecard__metric ${onSelectTrustMetric ? "overview-scorecard__metric--clickable" : ""}`}
                  key={metric.id}
                  onClick={() => onSelectTrustMetric?.(metric.id, metric.label)}
                >
                  <div className="overview-scorecard__metric-header">
                    <strong>{metric.label}</strong>
                    <span className={`chip chip--${trustStatusClass(metric.status)}`}>{metric.status}</span>
                  </div>
                  <p>{metric.score}/100 · weight {metric.weight}</p>
                  <small>{metric.detail}</small>
                  <small><strong>Likely defense gap:</strong> {metric.defenseGap}</small>
                  <small><strong>Common root weakness:</strong> {metric.rootWeakness}</small>
                  <small><strong>Suggested next control move:</strong> {metric.nextMove}</small>
                </button>
              ))}
            </div>

            <div className="overview-scorecard__footer">
              <p>{model.trust.derivedFrom}</p>
              {model.trust.draggers.length > 0 ? (
                <div className="chip-row">
                  {model.trust.draggers.map((item) => (
                    <span className="chip chip--warning" key={item}>{item}</span>
                  ))}
                </div>
              ) : null}
            </div>
          </div>
        </article>

        <article className="panel analytics-panel">
          <div className="table-toolbar">
            <span className="summary-label">Contextual guidance</span>
            <strong>{model.guidance.length}</strong>
          </div>
          <ul className="analytics-list">
            {model.guidance.map((item) => (
              <li key={item.id} className="analytics-list__item overview-list-item">
                <div>
                  <div className="overview-list-item__title">
                    <strong>{item.title}</strong>
                    <span className={`chip chip--${confidenceClass(item.confidence)}`}>{item.confidence}</span>
                  </div>
                  <p>{item.summary}</p>
                  <small>{item.source}</small>
                </div>
                <div className="overview-list-item__aside">{item.nextStep}</div>
              </li>
            ))}
          </ul>
        </article>
      </section>

      <section className="overview-workbench">
        <article className="panel analytics-panel">
          <div className="table-toolbar">
            <span className="summary-label">Executive defense reporting</span>
            <strong>{executiveReport?.shieldHealth.score ?? 0}</strong>
          </div>
          {executiveReport ? (
            <div className="overview-executive">
              <div className="overview-executive__header">
                <div>
                  <strong>{executiveReport.boardPackage.headline}</strong>
                  <p>{executiveReport.boardPackage.narrative}</p>
                  <small>{executiveReport.executiveSummary.whatMattersNow}</small>
                </div>
                <div className="chip-row">
                  <span className={`chip chip--${shieldBandClass(executiveReport.shieldHealth.band)}`}>{executiveReport.shieldHealth.band.replace("_", " ")}</span>
                  <span className="chip chip--muted">{executiveReport.audience}</span>
                </div>
              </div>

              <div className="overview-executive__summary">
                <div>
                  <span className="summary-label">Top risks</span>
                  <ul className="summary-list summary-list--compact">
                    {executiveReport.executiveSummary.topRisks.map((risk) => (
                      <li key={risk}><span>{risk}</span></li>
                    ))}
                  </ul>
                </div>
                <div>
                  <span className="summary-label">Top improvements</span>
                  <ul className="summary-list summary-list--compact">
                    {executiveReport.executiveSummary.topImprovements.map((item) => (
                      <li key={item}><span>{item}</span></li>
                    ))}
                  </ul>
                </div>
              </div>

              <div className="overview-executive__shield">
                <div>
                  <span className="summary-label">Shield health score</span>
                  <p className="overview-executive__score">{executiveReport.shieldHealth.score}/100</p>
                  <p>{executiveReport.shieldHealth.summary}</p>
                </div>
                <div className="overview-executive__components">
                  {executiveReport.shieldHealth.components.map((component) => (
                    <article className="summary-card summary-card--compact" key={component.key}>
                      <span className="summary-label">{component.label}</span>
                      <strong className="summary-value">{component.score}</strong>
                      <p>{component.summary}</p>
                    </article>
                  ))}
                </div>
              </div>

              <div className="overview-executive__trends">
                {executiveReport.riskReductionTrends.map((trend) => (
                  <article className="summary-card summary-card--compact" key={trend.key}>
                    <div className="overview-list-item__title">
                      <strong>{trend.label}</strong>
                      <span className={`chip chip--${trendClass(trend.direction)}`}>{trend.direction}</span>
                    </div>
                    <p>{trend.value}</p>
                    <small>{trend.summary}</small>
                  </article>
                ))}
              </div>

              {executiveReport.strategicGaps.length > 0 ? (
                <div>
                  <span className="summary-label">Strategic gap analysis</span>
                  <ul className="analytics-list">
                    {executiveReport.strategicGaps.slice(0, 3).map((gap) => (
                      <li key={gap.id} className="analytics-list__item overview-list-item">
                        <div>
                          <div className="overview-list-item__title">
                            <strong>{gap.title}</strong>
                            <span className={`chip chip--${confidenceClass(gap.confidence)}`}>{gap.confidence}</span>
                          </div>
                          <p>{gap.summary}</p>
                        </div>
                        <div className="overview-list-item__aside">{gap.investmentTarget}</div>
                      </li>
                    ))}
                  </ul>
                </div>
              ) : null}

              <div className="overview-executive__footer">
                <div>
                  <span className="summary-label">Board package</span>
                  <p>{executiveReport.boardPackage.packageSummary}</p>
                </div>
                <div>
                  <span className="summary-label">Trend change</span>
                  <p>{executiveReport.executiveSummary.trendChange}</p>
                </div>
              </div>
            </div>
          ) : (
            <div className="panel-empty">No executive defense report is currently loaded for the selected posture scope.</div>
          )}
        </article>

        <article className="panel analytics-panel">
          <div className="table-toolbar">
            <span className="summary-label">Systemic weaknesses</span>
            <strong>{systemicWeaknesses?.weaknesses.length || 0}</strong>
          </div>
          {systemicWeaknesses && systemicWeaknesses.weaknesses.length > 0 ? (
            <ul className="analytics-list">
              {systemicWeaknesses.weaknesses.slice(0, 4).map((weakness) => (
                <li key={weakness.patternKey} className="analytics-list__item overview-list-item">
                  <div>
                    <div className="overview-list-item__title">
                      <strong>{weakness.title}</strong>
                      <span className={`chip chip--${severityClass(weakness.priority)}`}>{weakness.priority}</span>
                    </div>
                    <p>{weakness.summary}</p>
                    <small>{weakness.rootCauseHypothesis}</small>
                  </div>
                  <div className="overview-list-item__aside">{weakness.executiveRecommendation}</div>
                </li>
              ))}
            </ul>
          ) : (
            <div className="panel-empty">No repeated systemic weakness pattern stands out in the currently loaded posture scope.</div>
          )}
        </article>

        <article className="panel analytics-panel">
          <div className="table-toolbar">
            <span className="summary-label">Top active incidents</span>
            <strong>{model.incidents.length}</strong>
          </div>
          {model.incidents.length > 0 ? (
            <ul className="analytics-list">
              {model.incidents.map((incident) => (
                <li key={incident.id} className="analytics-list__item overview-list-item">
                  <div>
                    <div className="overview-list-item__title">
                      <strong>{incident.title}</strong>
                      <span className={`chip chip--${severityClass(incident.severity)}`}>{incident.severity}</span>
                    </div>
                    <p>{incident.summary}</p>
                    <small>{incident.scope} · {incident.evidence}</small>
                  </div>
                  <div className="overview-list-item__aside">{incident.action}</div>
                </li>
              ))}
            </ul>
          ) : (
            <div className="panel-empty">No concentrated incident pattern is visible in the loaded scope.</div>
          )}
        </article>

        <article className="panel analytics-panel">
          <div className="table-toolbar">
            <span className="summary-label">Next operator actions</span>
            <strong>{model.actions.length}</strong>
          </div>
          <ul className="analytics-list">
            {model.actions.map((action) => (
              <li key={action.id} className="analytics-list__item overview-list-item">
                <div>
                  <div className="overview-list-item__title">
                    <strong>{action.title}</strong>
                    <span className={`chip chip--${priorityClass(action.priority)}`}>{action.priority}</span>
                  </div>
                  <p>{action.detail}</p>
                  <small>{action.source}</small>
                </div>
              </li>
            ))}
          </ul>
        </article>

        <article className="panel analytics-panel">
          <div className="table-toolbar">
            <span className="summary-label">Current drivers</span>
            <strong>{model.denyDrivers.length + model.blastRadius.length}</strong>
          </div>

          <div className="overview-split">
            <div>
              <h3>Top deny reasons</h3>
              {model.denyDrivers.length > 0 ? (
                <ul className="summary-list">
                  {model.denyDrivers.map((item) => (
                    <li key={item.label}>
                      <span>{item.label}</span>
                      <strong>{item.count}</strong>
                    </li>
                  ))}
                </ul>
              ) : (
                <div className="summary-list-empty">No deny reasons recorded yet.</div>
              )}
            </div>

            <div>
              <h3>Blast radius</h3>
              {model.blastRadius.length > 0 ? (
                <ul className="summary-list">
                  {model.blastRadius.map((item) => (
                    <li key={`${item.label}-${item.detail}`}>
                      <span>
                        {item.label}
                        <small className="details-subtext">{item.detail}</small>
                      </span>
                      <strong>{item.count}</strong>
                    </li>
                  ))}
                </ul>
              ) : (
                <div className="summary-list-empty">No concentrated scope stands out yet.</div>
              )}
            </div>
          </div>
        </article>
      </section>

      <section className="overview-metrics">
        {model.metrics.map((metric) => (
          <article className={`panel overview-metric overview-metric--${metric.tone}`} key={metric.id}>
            <span className="summary-label">{metric.label}</span>
            <strong className="summary-value">{metric.value}</strong>
            <p className="overview-metric__trend">{metric.trend}</p>
            <p>{metric.context}</p>
          </article>
        ))}
      </section>
    </>
  );
}
