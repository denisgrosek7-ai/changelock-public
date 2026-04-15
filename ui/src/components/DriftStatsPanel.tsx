import type { DriftStatsResponse } from "../types";

type Props = {
  data: DriftStatsResponse | null;
  loading: boolean;
};

export function DriftStatsPanel({ data, loading }: Props) {
  if (loading) {
    return <section className="panel panel-empty">Loading drift statistics…</section>;
  }
  if (!data) {
    return <section className="panel panel-empty">No drift statistics available.</section>;
  }

  const driftClasses = Object.entries(data.counts_by_drift_class);

  return (
    <section className="panel analytics-panel">
      <div className="table-toolbar">
        <span className="summary-label">Runtime Drift Stats</span>
        <strong>{data.total_runtime_drift_denies}</strong>
      </div>
      <div className="analytics-metrics">
        <article className="summary-card">
          <span className="summary-label">Drift Denies</span>
          <strong className="summary-value">{data.total_runtime_drift_denies}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Approx. MTTR</span>
          <strong className="summary-value">
            {data.mean_time_to_resolve_seconds != null
              ? `${Math.round(data.mean_time_to_resolve_seconds / 60)} min`
              : "n/a"}
          </strong>
        </article>
      </div>
      <div className="analytics-split">
        <div>
          <h3>Counts by Drift Class</h3>
          {driftClasses.length > 0 ? (
            <ul className="summary-list">
              {driftClasses.map(([driftClass, count]) => (
                <li key={driftClass}>
                  <span>{driftClass}</span>
                  <strong>{count}</strong>
                </li>
              ))}
            </ul>
          ) : (
            <div className="summary-list-empty">No drift classes recorded.</div>
          )}
        </div>
        <div>
          <h3>Top Drifted Workloads</h3>
          {data.top_drifted_workloads.length > 0 ? (
            <ul className="summary-list">
              {data.top_drifted_workloads.map((workload) => (
                <li key={`${workload.namespace}-${workload.workload}`}>
                  <span>{workload.workload}</span>
                  <strong>{workload.count}</strong>
                </li>
              ))}
            </ul>
          ) : (
            <div className="summary-list-empty">No drifted workloads found.</div>
          )}
        </div>
      </div>
    </section>
  );
}
