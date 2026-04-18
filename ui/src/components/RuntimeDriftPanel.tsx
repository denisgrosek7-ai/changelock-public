import type { RuntimeActiveState, RuntimeClosedLoopStatus, RuntimeDriftFinding, RuntimeDriftStatus } from "../types";

type Props = {
  findings: RuntimeDriftFinding[];
  status: RuntimeDriftStatus | null;
  activeStates: RuntimeActiveState[];
  closedLoopStatus: RuntimeClosedLoopStatus | null;
  loading: boolean;
};

export function RuntimeDriftPanel({ findings, status, activeStates, closedLoopStatus, loading }: Props) {
  if (loading && !status && !closedLoopStatus && findings.length === 0 && activeStates.length === 0) {
    return <section className="panel panel-empty">Loading runtime drift status…</section>;
  }

  return (
    <section className="panel">
      <div className="panel-header">
        <div>
          <h3>Runtime Self-Healing</h3>
          <p>Active drift, remediation outcomes, and quarantine state.</p>
        </div>
      </div>

      <section className="summary-grid">
        <article className="summary-card">
          <span className="summary-label">In Sync</span>
          <strong className="summary-value">{closedLoopStatus?.in_sync ?? 0}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Active</span>
          <strong className="summary-value">{status?.active_findings ?? 0}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Quarantined</span>
          <strong className="summary-value">{status?.quarantined ?? 0}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Failed</span>
          <strong className="summary-value">{status?.failed ?? 0}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Remediated</span>
          <strong className="summary-value">{status?.remediated ?? 0}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Protected</span>
          <strong className="summary-value">{closedLoopStatus?.protected_blocked ?? 0}</strong>
        </article>
      </section>

      {activeStates.length === 0 ? (
        <div className="panel-empty">No active closed-loop runtime state recorded.</div>
      ) : (
        <div className="table-wrapper">
          <table className="data-table">
            <thead>
              <tr>
                <th>Target</th>
                <th>Reconcile</th>
                <th>Mode</th>
                <th>Trust</th>
                <th>Containment</th>
              </tr>
            </thead>
            <tbody>
              {activeStates.map((item) => (
                <tr key={item.id}>
                  <td>
                    <strong>{item.workload}</strong>
                    <div className="table-meta">
                      {item.namespace} · {item.workload_kind}
                    </div>
                  </td>
                  <td>{item.reconciliation_status}</td>
                  <td>{item.remediation_mode || "-"}</td>
                  <td>{item.desired_state_verification_state || "-"}</td>
                  <td>
                    {item.quarantine_type || item.protected_reason || item.last_error || "-"}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {findings.length === 0 ? (
        <div className="panel-empty">No active runtime drift findings.</div>
      ) : (
        <div className="table-wrapper">
          <table className="data-table">
            <thead>
              <tr>
                <th>Workload</th>
                <th>Severity</th>
                <th>Status</th>
                <th>Mode</th>
                <th>Attempts</th>
                <th>Verification</th>
              </tr>
            </thead>
            <tbody>
              {findings.map((finding) => (
                <tr key={finding.id}>
                  <td>
                    <strong>{finding.workload}</strong>
                    <div className="table-meta">
                      {finding.namespace} · {finding.workload_kind}
                    </div>
                  </td>
                  <td>{finding.drift_severity || "-"}</td>
                  <td>{finding.status}</td>
                  <td>{finding.remediation_mode || "-"}</td>
                  <td>{finding.remediation_attempt ?? 0}</td>
                  <td>{finding.desired_state_verification_state || "-"}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </section>
  );
}
