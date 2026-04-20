import type {
  AnalyticsAnomaliesResponse,
  AnalyticsDeltaResponse,
  AnalyticsMetricTrend,
  AnalyticsScorecardsResponse,
  AnalyticsSegmentsResponse,
  TrendsResponse,
} from "../types";

type Props = {
  trends: TrendsResponse | null;
  delta: AnalyticsDeltaResponse | null;
  anomalies: AnalyticsAnomaliesResponse | null;
  scorecards: AnalyticsScorecardsResponse | null;
  segments: AnalyticsSegmentsResponse | null;
  loading: boolean;
};

function formatMetricValue(metricKey: string, value: number) {
  if (metricKey === "policy_friction_rate") {
    return `${value.toFixed(1)}%`;
  }
  if (Math.abs(value) >= 10 || Number.isInteger(value)) {
    return value.toFixed(0);
  }
  return value.toFixed(1);
}

function directionTone(direction: string) {
  switch (direction) {
    case "improving":
      return "allow";
    case "worsening":
      return "deny";
    default:
      return "warning";
  }
}

function topMetricTrends(trends: TrendsResponse | null): AnalyticsMetricTrend[] {
  if (!trends?.metric_trends) {
    return [];
  }
  return [...trends.metric_trends].sort((left, right) => Math.abs(right.delta_value) - Math.abs(left.delta_value)).slice(0, 4);
}

export function AnalyticsInsightsPanel({ trends, delta, anomalies, scorecards, segments, loading }: Props) {
  if (loading) {
    return <section className="panel panel-empty analytics-panel analytics-panel--wide">Loading trend and delta analytics…</section>;
  }

  const metricTrends = topMetricTrends(trends);
  const scorecardCards = scorecards?.cards || [];
  const anomalyItems = anomalies?.items || [];
  const segmentItems = segments?.items || [];

  if (metricTrends.length === 0 && scorecardCards.length === 0 && anomalyItems.length === 0) {
    return <section className="panel panel-empty analytics-panel analytics-panel--wide">No trend and delta analytics are available for the current scope.</section>;
  }

  return (
    <section className="panel analytics-panel analytics-panel--wide">
      <div className="table-toolbar">
        <div>
          <span className="summary-label">Trend &amp; delta analytics</span>
          <strong>
            {trends?.comparison?.window || delta?.comparison.window || scorecards?.comparison.window || "28d"} vs{" "}
            {trends?.comparison?.compare_to || delta?.comparison.compare_to || scorecards?.comparison.compare_to || "previous_window"}
          </strong>
        </div>
        <span className="badge">{trends?.comparison?.group_by || delta?.comparison.group_by || "service"}</span>
      </div>

      {scorecardCards.length > 0 ? (
        <div className="analytics-scorecard-grid">
          {scorecardCards.map((card) => (
            <article className="analytics-scorecard-card" key={card.definition.key}>
              <div className="analytics-scorecard-card__header">
                <span className="summary-label">{card.definition.label}</span>
                <span className={`badge badge--${directionTone(card.direction)}`}>{card.status}</span>
              </div>
              <strong className="summary-value">{formatMetricValue(card.definition.key, card.current_value)}</strong>
              <p>{card.summary}</p>
              <small>
                Baseline {formatMetricValue(card.definition.key, card.baseline_value)} · Delta {formatMetricValue(card.definition.key, card.delta_value)}
              </small>
            </article>
          ))}
        </div>
      ) : null}

      <div className="analytics-insights-grid">
        <div>
          <h3>Top trend shifts</h3>
          {metricTrends.length > 0 ? (
            <ul className="analytics-list">
              {metricTrends.map((item) => (
                <li key={item.definition.key} className="analytics-list__item">
                  <div>
                    <strong>{item.definition.label}</strong>
                    <p>{item.summary}</p>
                  </div>
                  <span className={`badge badge--${directionTone(item.direction)}`}>
                    {item.direction} {formatMetricValue(item.definition.key, item.delta_value)}
                  </span>
                </li>
              ))}
            </ul>
          ) : (
            <div className="summary-list-empty">No metric trend shifts are available for this scope.</div>
          )}
        </div>

        <div>
          <h3>Explainable anomalies</h3>
          {anomalyItems.length > 0 ? (
            <ul className="analytics-list">
              {anomalyItems.slice(0, 5).map((item) => (
                <li key={`${item.type}-${item.segment}`} className="analytics-list__item">
                  <div>
                    <strong>{item.title}</strong>
                    <p>
                      {item.segment} · {item.deviation}
                    </p>
                    <small>{item.recommended_next_step}</small>
                  </div>
                  <span className={`badge badge--${item.severity === "high" ? "deny" : "warning"}`}>{item.severity}</span>
                </li>
              ))}
            </ul>
          ) : (
            <div className="summary-list-empty">No anomaly spikes crossed the explainable threshold in this window.</div>
          )}
        </div>
      </div>

      <div className="analytics-insights-grid">
        <div>
          <h3>Delta by segment</h3>
          {delta && delta.segments.length > 0 ? (
            <ul className="summary-list">
              {delta.segments.slice(0, 6).map((segment) => (
                <li key={segment.segment_key}>
                  <span>{segment.segment_label}</span>
                  <strong>
                    {formatMetricValue(delta.definition.key, segment.current_value)} / {formatMetricValue(delta.definition.key, segment.delta_value)}
                  </strong>
                </li>
              ))}
            </ul>
          ) : (
            <div className="summary-list-empty">No segment deltas are available for the selected metric.</div>
          )}
          {delta ? <p className="analytics-panel__helper">{delta.summary}</p> : null}
        </div>

        <div>
          <h3>Available segments</h3>
          {segmentItems.length > 0 ? (
            <div className="analytics-chip-groups">
              {segmentItems.map((item) => (
                <div key={item.group}>
                  <span className="summary-label">{item.group}</span>
                  <div className="chip-row">
                    {item.values.slice(0, 6).map((value) => (
                      <span className="chip" key={`${item.group}-${value}`}>
                        {value}
                      </span>
                    ))}
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className="summary-list-empty">No segment catalog is available for this scope.</div>
          )}
        </div>
      </div>

      {trends?.limitations && trends.limitations.length > 0 ? (
        <div className="analytics-panel__footnotes">
          {trends.limitations.slice(0, 3).map((item) => (
            <p key={item}>{item}</p>
          ))}
        </div>
      ) : null}
    </section>
  );
}
