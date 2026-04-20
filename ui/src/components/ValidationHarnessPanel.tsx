import { useState } from "react";

import type {
  ValidationHarnessRun,
  ValidationHarnessScenario,
  ValidationHarnessScore,
  ValidationHarnessWhatIfResponse,
} from "../types";

type Props = {
  scenarios: ValidationHarnessScenario[];
  score: ValidationHarnessScore | null;
  runs: ValidationHarnessRun[];
  whatIf: ValidationHarnessWhatIfResponse | null;
  loading: boolean;
  onRunHarness?: () => Promise<void>;
  onRunWhatIf?: () => Promise<void>;
};

export function ValidationHarnessPanel({ scenarios, score, runs, whatIf, loading, onRunHarness, onRunWhatIf }: Props) {
  const [actionError, setActionError] = useState<string | null>(null);
  const [actionLoading, setActionLoading] = useState<"run" | "what-if" | null>(null);

  async function runAction(kind: "run" | "what-if", action?: () => Promise<void>) {
    if (!action) {
      return;
    }
    setActionError(null);
    setActionLoading(kind);
    try {
      await action();
    } catch (error) {
      setActionError(error instanceof Error ? error.message : "Validation harness action failed.");
    } finally {
      setActionLoading(null);
    }
  }

  if (loading && !score && scenarios.length === 0 && runs.length === 0) {
    return <section className="panel panel-empty">Loading controlled validation harness…</section>;
  }

  return (
    <section className="panel">
      <div className="panel-header">
        <div>
          <h3>Controlled Validation Harness</h3>
          <p>Dry-run security regression, bounded chaos rehearsal, and what-if confidence checks over current policy and runtime state.</p>
        </div>
        <div className="toolbar-inline">
          <button className="button" disabled={actionLoading !== null || !onRunHarness} onClick={() => runAction("run", onRunHarness)}>
            {actionLoading === "run" ? "Running…" : "Run Harness"}
          </button>
          <button className="button button-secondary" disabled={actionLoading !== null || !onRunWhatIf} onClick={() => runAction("what-if", onRunWhatIf)}>
            {actionLoading === "what-if" ? "Projecting…" : "Run What-If"}
          </button>
        </div>
      </div>

      {actionError ? <div className="status-banner"><p>{actionError}</p></div> : null}

      {score ? (
        <section className="summary-grid">
          <article className="summary-card">
            <span className="summary-label">Confidence</span>
            <strong className="summary-value">{score.confidence_level}</strong>
          </article>
          <article className="summary-card">
            <span className="summary-label">Overall</span>
            <strong className="summary-value">{score.overall_status}</strong>
          </article>
          <article className="summary-card">
            <span className="summary-label">Passed</span>
            <strong className="summary-value">{score.passed_scenarios}</strong>
          </article>
          <article className="summary-card">
            <span className="summary-label">Partial</span>
            <strong className="summary-value">{score.partial_scenarios}</strong>
          </article>
          <article className="summary-card">
            <span className="summary-label">Failed</span>
            <strong className="summary-value">{score.failed_scenarios}</strong>
          </article>
          <article className="summary-card">
            <span className="summary-label">Avg Response</span>
            <strong className="summary-value">{score.average_response_ms}ms</strong>
          </article>
        </section>
      ) : null}

      {score?.critical_gaps?.length ? (
        <section className="panel-subsection">
          <span className="summary-label">Critical Gaps</span>
          <ul className="summary-list">
            {score.critical_gaps.slice(0, 4).map((item) => (
              <li key={item}>{item}</li>
            ))}
          </ul>
        </section>
      ) : null}

      {scenarios.length > 0 ? (
        <div className="table-wrapper">
          <table className="data-table">
            <thead>
              <tr>
                <th>Scenario</th>
                <th>Category</th>
                <th>Mode</th>
                <th>Expected Outcome</th>
              </tr>
            </thead>
            <tbody>
              {scenarios.map((scenario) => (
                <tr key={scenario.scenario_id}>
                  <td>
                    <strong>{scenario.title}</strong>
                    <div className="table-meta">{scenario.description}</div>
                  </td>
                  <td>{scenario.category}</td>
                  <td>{scenario.validation_mode}</td>
                  <td>{scenario.expected_outcome}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      ) : (
        <div className="panel-empty">No validation harness scenarios are available in the current scope.</div>
      )}

      {score?.results?.length ? (
        <div className="table-wrapper">
          <table className="data-table">
            <thead>
              <tr>
                <th>Scenario Result</th>
                <th>Status</th>
                <th>Response</th>
                <th>Evidence</th>
              </tr>
            </thead>
            <tbody>
              {score.results.map((result) => (
                <tr key={result.scenario_id}>
                  <td>
                    <strong>{result.scenario_id}</strong>
                    <div className="table-meta">{result.summary}</div>
                  </td>
                  <td>{result.status}</td>
                  <td>{result.response_time_ms}ms</td>
                  <td>{result.evidence_refs?.length || 0}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      ) : null}

      {runs.length > 0 ? (
        <div className="table-wrapper">
          <table className="data-table">
            <thead>
              <tr>
                <th>Run</th>
                <th>Status</th>
                <th>Certificate</th>
                <th>Completed</th>
              </tr>
            </thead>
            <tbody>
              {runs.slice(0, 5).map((run) => (
                <tr key={run.run_id}>
                  <td>
                    <strong>{run.scope_summary}</strong>
                    <div className="table-meta">{run.mode}</div>
                  </td>
                  <td>{run.overall_status}</td>
                  <td>{run.certificate_status}</td>
                  <td>{new Date(run.completed_at).toLocaleString()}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      ) : (
        <div className="panel-empty">No persisted validation harness runs recorded yet.</div>
      )}

      {whatIf ? (
        <section className="panel-subsection">
          <div className="panel-header">
            <div>
              <h4>What-If Projection</h4>
              <p>{whatIf.change_set.join(" · ")}</p>
            </div>
          </div>
          <section className="summary-grid">
            <article className="summary-card">
              <span className="summary-label">Projected Status</span>
              <strong className="summary-value">{whatIf.overall_status}</strong>
            </article>
            <article className="summary-card">
              <span className="summary-label">Pass</span>
              <strong className="summary-value">{whatIf.projected_pass}</strong>
            </article>
            <article className="summary-card">
              <span className="summary-label">Partial</span>
              <strong className="summary-value">{whatIf.projected_partial}</strong>
            </article>
            <article className="summary-card">
              <span className="summary-label">Fail</span>
              <strong className="summary-value">{whatIf.projected_fail}</strong>
            </article>
          </section>
          {whatIf.compatibility_risks?.length ? (
            <ul className="summary-list">
              {whatIf.compatibility_risks.map((item) => (
                <li key={item}>{item}</li>
              ))}
            </ul>
          ) : null}
        </section>
      ) : null}
    </section>
  );
}
