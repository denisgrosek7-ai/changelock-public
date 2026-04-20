import type { TopologyDeltaResponse, TopologyGraphResponse, TopologyHeatmapResponse } from "../types";

type Props = {
  graph: TopologyGraphResponse | null;
  heatmap: TopologyHeatmapResponse | null;
  delta: TopologyDeltaResponse | null;
  loading: boolean;
};

function renderEmpty(message: string) {
  return <div className="summary-list-empty">{message}</div>;
}

export function TopologyInsightsPanel({ graph, heatmap, delta, loading }: Props) {
  if (loading) {
    return <section className="panel analytics-panel analytics-panel--wide">Loading topology intelligence…</section>;
  }

  return (
    <>
      <section className="panel analytics-panel analytics-panel--wide">
        <div className="incident-report-section__header">
          <span className="summary-label">Service-graph blast radius</span>
          <strong>Declared, observed, and effective connectivity</strong>
        </div>
        {graph ? (
          <>
            <div className="summary-grid">
              <article className="summary-card">
                <span className="summary-label">Declared graph</span>
                <strong className="summary-value">{graph.summary.declared_edges}</strong>
                <p>{graph.summary.declared_nodes} nodes with policy/config-derived reachability.</p>
              </article>
              <article className="summary-card">
                <span className="summary-label">Observed graph</span>
                <strong className="summary-value">{graph.summary.observed_edges}</strong>
                <p>{graph.summary.observed_nodes} nodes with telemetry-derived paths.</p>
              </article>
              <article className="summary-card">
                <span className="summary-label">Effective graph</span>
                <strong className="summary-value">{graph.summary.effective_edges}</strong>
                <p>{graph.summary.effective_nodes} nodes in the security-oriented graph.</p>
              </article>
              <article className="summary-card">
                <span className="summary-label">Critical exposure</span>
                <strong className="summary-value">{graph.summary.high_blast_radius}</strong>
                <p>{graph.summary.public_entry_nodes} public entry nodes · {graph.summary.critical_nodes} critical nodes.</p>
              </article>
            </div>

            <div className="incident-evidence-grid">
              <div>
                <span className="summary-label">Top effective paths</span>
                {graph.effective_graph.edges.length > 0 ? (
                  <ul className="summary-list summary-list--compact">
                    {graph.effective_graph.edges.slice(0, 6).map((edge) => (
                      <li key={`${edge.source}-${edge.target}-${edge.edge_type}`}>
                        <span>{edge.source} → {edge.target} · {edge.edge_type} · {edge.connectivity_class}</span>
                      </li>
                    ))}
                  </ul>
                ) : renderEmpty("No effective connectivity edges matched the current scope.")}
              </div>
              <div>
                <span className="summary-label">Limitations</span>
                {graph.limitations && graph.limitations.length > 0
                  ? (
                    <ul className="summary-list summary-list--compact">
                      {graph.limitations.map((item) => (
                        <li key={item}>
                          <span>{item}</span>
                        </li>
                      ))}
                    </ul>
                  )
                  : renderEmpty("No explicit topology limitations recorded for this scope.")}
              </div>
            </div>
          </>
        ) : renderEmpty("Topology graph is not available for the current scope.")}
      </section>

      <section className="panel analytics-panel analytics-panel--wide">
        <div className="incident-report-section__header">
          <span className="summary-label">Heatmap</span>
          <strong>Services ranked by node risk and blast radius</strong>
        </div>
        {heatmap && heatmap.items.length > 0 ? (
          <div className="incident-package-table">
            <div className="incident-package-table__row incident-package-table__row--header">
              <span>Service</span>
              <span>Node risk</span>
              <span>Blast radius</span>
              <span>Critical reach</span>
              <span>Exposure</span>
            </div>
            {heatmap.items.slice(0, 10).map((item) => (
              <div className="incident-package-table__row" key={item.node_id}>
                <span>
                  <strong>{item.service}</strong>
                  <small>{item.environment || item.cluster || "unknown scope"}</small>
                </span>
                <span>{item.node_risk_score}</span>
                <span>{item.blast_radius_score}</span>
                <span>{item.critical_reach_count}</span>
                <span>{item.public_exposure ? "public" : "internal"} · {item.propagation_class}</span>
              </div>
            ))}
          </div>
        ) : renderEmpty("No topology heatmap items matched the current scope.")}
      </section>

      <section className="panel analytics-panel analytics-panel--wide">
        <div className="incident-report-section__header">
          <span className="summary-label">Topology delta</span>
          <strong>Reachability drift against the previous window</strong>
        </div>
        {delta && delta.items.length > 0 ? (
          <div className="incident-impact-list">
            {delta.items.slice(0, 6).map((item) => (
              <article className="incident-impact-card incident-defense-gap" key={`${item.node_id}-${item.service}`}>
                <div className="incident-impact-card__header">
                  <strong>{item.service}</strong>
                  <span className={`chip chip--${item.delta > 0 ? "deny" : item.delta < 0 ? "allow" : "muted"}`}>
                    delta {item.delta > 0 ? "+" : ""}{item.delta}
                  </span>
                </div>
                <p>
                  current {item.current_blast_radius_score} · baseline {item.baseline_blast_radius_score} · edge additions {item.edge_additions}
                </p>
                {item.drift_signals && item.drift_signals.length > 0 ? (
                  <div className="chip-row">
                    {item.drift_signals.map((signal) => (
                      <span className="chip chip--muted" key={`${item.node_id}-${signal}`}>{signal}</span>
                    ))}
                  </div>
                ) : null}
              </article>
            ))}
          </div>
        ) : renderEmpty("No topology drift items matched the current scope.")}
      </section>
    </>
  );
}
