import type {
  RuntimeEnforcementDecision,
  RuntimeIntegrityFinding,
  RuntimeIntegrityState,
  RuntimeWorkloadView,
} from "../types";

type Props = {
  integrity: RuntimeIntegrityState[];
  workloads: RuntimeWorkloadView[];
  findings: RuntimeIntegrityFinding[];
  enforcement: RuntimeEnforcementDecision[];
  loading: boolean;
};

function countBy<T>(items: T[], predicate: (item: T) => boolean): number {
  return items.reduce((count, item) => count + (predicate(item) ? 1 : 0), 0);
}

export function RuntimeIntegrityPanel({ integrity, workloads, findings, enforcement, loading }: Props) {
  if (loading && integrity.length === 0 && workloads.length === 0 && findings.length === 0 && enforcement.length === 0) {
    return <section className="panel panel-empty">Loading runtime integrity mesh…</section>;
  }

  return (
    <section className="panel">
      <div className="panel-header">
        <div>
          <h3>Runtime Integrity Mesh</h3>
          <p>Backend-derived runtime integrity, drift, sandbox posture, and controlled enforcement history.</p>
        </div>
      </div>

      <section className="summary-grid">
        <article className="summary-card">
          <span className="summary-label">Verified</span>
          <strong className="summary-value">{countBy(integrity, (item) => item.identity_status === "verified")}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Critical Drift</span>
          <strong className="summary-value">{countBy(integrity, (item) => item.drift_level === "critical")}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Isolated Review</span>
          <strong className="summary-value">{countBy(integrity, (item) => item.current_sandbox_class === "isolated_review")}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Unverifiable</span>
          <strong className="summary-value">{countBy(integrity, (item) => item.sbom_verification.status === "unverifiable")}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Approval Gated</span>
          <strong className="summary-value">{countBy(enforcement, (item) => item.approval_mode === "approval_required")}</strong>
        </article>
      </section>

      {workloads.length === 0 ? (
        <div className="panel-empty">No runtime integrity workloads available in the current scope.</div>
      ) : (
        <div className="table-wrapper">
          <table className="data-table">
            <thead>
              <tr>
                <th>Workload</th>
                <th>Identity</th>
                <th>Integrity</th>
                <th>Drift</th>
                <th>Sandbox</th>
                <th>Posture</th>
              </tr>
            </thead>
            <tbody>
              {workloads.slice(0, 8).map((item) => (
                <tr key={item.subject_ref}>
                  <td>
                    <strong>{item.workload || item.subject_ref}</strong>
                    <div className="table-meta">
                      {item.namespace || "scope"} · {item.workload_kind || "workload"}
                    </div>
                  </td>
                  <td>{item.state.identity_status}</td>
                  <td>
                    <strong>{item.state.runtime_integrity_score}</strong>
                    <div className="table-meta">{item.state.sbom_verification.status}</div>
                  </td>
                  <td>{item.state.drift_level}</td>
                  <td>{item.sandbox_decision.assigned_sandbox_class}</td>
                  <td>{item.state.current_enforcement_posture}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {findings.length === 0 ? (
        <div className="panel-empty">No runtime integrity findings are active in the current scope.</div>
      ) : (
        <div className="table-wrapper">
          <table className="data-table">
            <thead>
              <tr>
                <th>Finding</th>
                <th>Severity</th>
                <th>Subject</th>
                <th>Action</th>
                <th>Confidence</th>
              </tr>
            </thead>
            <tbody>
              {findings.slice(0, 8).map((item) => (
                <tr key={item.finding_id}>
                  <td>
                    <strong>{item.finding_type}</strong>
                    <div className="table-meta">{item.summary}</div>
                  </td>
                  <td>{item.severity}</td>
                  <td>{item.subject_ref}</td>
                  <td>{item.recommended_action}</td>
                  <td>{item.confidence}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {enforcement.length === 0 ? (
        <div className="panel-empty">No runtime enforcement decisions recorded yet.</div>
      ) : (
        <div className="table-wrapper">
          <table className="data-table">
            <thead>
              <tr>
                <th>Action</th>
                <th>Subject</th>
                <th>Approval</th>
                <th>Executed</th>
                <th>Result</th>
              </tr>
            </thead>
            <tbody>
              {enforcement.slice(0, 6).map((item) => (
                <tr key={item.decision_id}>
                  <td>{item.action}</td>
                  <td>{item.subject_ref}</td>
                  <td>{item.approval_mode}</td>
                  <td>{item.executed ? "yes" : "no"}</td>
                  <td>{item.execution_result}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </section>
  );
}
