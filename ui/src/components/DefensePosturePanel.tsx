import type { DefensePostureState, HardeningExecutionRecord } from "../types";

type Props = {
  posture: DefensePostureState[];
  actions: HardeningExecutionRecord[];
  loading: boolean;
};

function countBy<T>(items: T[], predicate: (item: T) => boolean): number {
  return items.reduce((count, item) => count + (predicate(item) ? 1 : 0), 0);
}

export function DefensePosturePanel({ posture, actions, loading }: Props) {
  if (loading && posture.length === 0 && actions.length === 0) {
    return <section className="panel panel-empty">Loading runtime defense posture…</section>;
  }

  return (
    <section className="panel">
      <div className="panel-header">
        <div>
          <h3>Defense Posture</h3>
          <p>Policy-gated runtime hardening, forensic-first containment, rollback readiness, and trusted recovery history.</p>
        </div>
      </div>

      <section className="summary-grid">
        <article className="summary-card">
          <span className="summary-label">Soft Isolation</span>
          <strong className="summary-value">{countBy(posture, (item) => item.current_mode === "soft_isolation")}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Pending Approval</span>
          <strong className="summary-value">{countBy(posture, (item) => item.current_mode === "pending_approval")}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Rollback Ready</span>
          <strong className="summary-value">{countBy(posture, (item) => item.rollback_ready)}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Forensic Linked</span>
          <strong className="summary-value">{countBy(posture, (item) => item.forensic_status === "linked")}</strong>
        </article>
      </section>

      {posture.length === 0 ? (
        <div className="panel-empty">No active runtime hardening posture is recorded in the current scope.</div>
      ) : (
        <div className="table-wrapper">
          <table className="data-table">
            <thead>
              <tr>
                <th>Workload</th>
                <th>Mode</th>
                <th>Restrictions</th>
                <th>Forensics</th>
                <th>TTL</th>
                <th>Rollback</th>
              </tr>
            </thead>
            <tbody>
              {posture.slice(0, 8).map((item) => (
                <tr key={item.subject_ref}>
                  <td>
                    <strong>{item.subject_ref}</strong>
                    <div className="table-meta">{item.trigger_summary || "No active trigger summary."}</div>
                  </td>
                  <td>{item.current_mode}</td>
                  <td>{item.active_restrictions?.join(", ") || "none"}</td>
                  <td>{item.forensic_status || "not_requested"}</td>
                  <td>{item.expires_at ? new Date(item.expires_at).toLocaleString() : "n/a"}</td>
                  <td>{item.rollback_ready ? "ready" : "no"}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {actions.length === 0 ? (
        <div className="panel-empty">No hardening execution records have been recorded yet.</div>
      ) : (
        <div className="table-wrapper">
          <table className="data-table">
            <thead>
              <tr>
                <th>Execution</th>
                <th>Subject</th>
                <th>Actions</th>
                <th>Result</th>
                <th>Forensics</th>
              </tr>
            </thead>
            <tbody>
              {actions.slice(0, 6).map((item) => (
                <tr key={item.execution_id}>
                  <td>
                    <strong>{item.execution_id}</strong>
                    <div className="table-meta">{new Date(item.executed_at).toLocaleString()}</div>
                  </td>
                  <td>{item.subject_ref}</td>
                  <td>{item.actions_applied?.map((action) => action.action_type).join(", ") || "none"}</td>
                  <td>{item.execution_result}</td>
                  <td>{item.forensic_refs?.length || 0}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </section>
  );
}
