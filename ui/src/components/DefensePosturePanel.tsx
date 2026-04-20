import type { DefensePostureState, HardeningExecutionRecord } from "../types";

type Props = {
  posture: DefensePostureState[];
  actions: HardeningExecutionRecord[];
  loading: boolean;
  focusSubjectRef?: string | null;
  focusExecutionID?: string | null;
};

function countBy<T>(items: T[], predicate: (item: T) => boolean): number {
  return items.reduce((count, item) => count + (predicate(item) ? 1 : 0), 0);
}

export function DefensePosturePanel({ posture, actions, loading, focusSubjectRef, focusExecutionID }: Props) {
  if (loading && posture.length === 0 && actions.length === 0) {
    return <section className="panel panel-empty">Loading runtime defense posture…</section>;
  }

  const visiblePosture = [...posture].sort((left, right) => {
    const leftFocused = focusSubjectRef && left.subject_ref === focusSubjectRef;
    const rightFocused = focusSubjectRef && right.subject_ref === focusSubjectRef;
    if (leftFocused !== rightFocused) {
      return leftFocused ? -1 : 1;
    }
    return left.subject_ref.localeCompare(right.subject_ref);
  });
  const visibleActions = [...actions].sort((left, right) => {
    const leftFocused = (focusExecutionID && left.execution_id === focusExecutionID) || (focusSubjectRef && left.subject_ref === focusSubjectRef);
    const rightFocused = (focusExecutionID && right.execution_id === focusExecutionID) || (focusSubjectRef && right.subject_ref === focusSubjectRef);
    if (leftFocused !== rightFocused) {
      return leftFocused ? -1 : 1;
    }
    return left.execution_id.localeCompare(right.execution_id);
  });

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
              {visiblePosture.slice(0, 8).map((item) => (
                <tr key={item.subject_ref} className={focusSubjectRef === item.subject_ref ? "is-selected" : undefined}>
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
              {visibleActions.slice(0, 6).map((item) => (
                <tr key={item.execution_id} className={(focusExecutionID === item.execution_id || focusSubjectRef === item.subject_ref) ? "is-selected" : undefined}>
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
