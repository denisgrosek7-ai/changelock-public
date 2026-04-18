import { useEffect, useState } from "react";

import { createAuditReport } from "../api";
import type { AuditReport } from "../types";

type Props = {
  tenantID?: string;
};

function formatTimestamp(value?: string) {
  if (!value) {
    return "-";
  }
  return new Date(value).toLocaleString();
}

export function TrustScorecardPanel({ tenantID }: Props) {
  const [report, setReport] = useState<AuditReport | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;

    async function load() {
      setLoading(true);
      setError(null);
      try {
        const next = await createAuditReport({
          tenant_id: tenantID,
          include_public_view: true,
        });
        if (!cancelled) {
          setReport(next);
        }
      } catch (loadError) {
        if (!cancelled) {
          setError(loadError instanceof Error ? loadError.message : "Unable to load trust scorecard.");
          setReport(null);
        }
      } finally {
        if (!cancelled) {
          setLoading(false);
        }
      }
    }

    void load();
    return () => {
      cancelled = true;
    };
  }, [tenantID]);

  if (loading && !report) {
    return <section className="panel panel-empty">Loading trust scorecard…</section>;
  }

  const scorecard = report?.scorecard;

  return (
    <>
      <section className="summary-grid">
        <article className="summary-card">
          <span className="summary-label">Overall Grade</span>
          <strong className="summary-value">{scorecard?.overall_grade || "-"}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Overall Score</span>
          <strong className="summary-value">{scorecard?.overall_score ?? 0}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Signing Coverage</span>
          <strong className="summary-value">{scorecard?.signing_coverage ?? 0}%</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Transparency Coverage</span>
          <strong className="summary-value">{scorecard?.transparency_coverage ?? 0}%</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Net Actionable Vulns</span>
          <strong className="summary-value">{scorecard?.actionable_vulnerability_count ?? 0}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Stale Exceptions</span>
          <strong className="summary-value">{scorecard?.stale_exception_count ?? 0}</strong>
        </article>
      </section>

      {error ? (
        <section className="panel status-banner">
          <div>
            <strong>Trust scorecard is unavailable.</strong>
            <p>{error}</p>
          </div>
        </section>
      ) : null}

      <section className="analytics-grid">
        <article className="panel">
          <header className="panel-header">
            <div>
              <h3>Metric Breakdown</h3>
              <p>Explainable weighted metrics derived from measured policy, evidence, VEX, signer, runtime, and exception signals.</p>
            </div>
          </header>
          <div className="table-shell">
            <table>
              <thead>
                <tr>
                  <th>Metric</th>
                  <th>Status</th>
                  <th>Score</th>
                  <th>Reason</th>
                </tr>
              </thead>
              <tbody>
                {scorecard?.metrics.map((metric) => (
                  <tr key={metric.id}>
                    <td>{metric.name}</td>
                    <td>{metric.status}</td>
                    <td>{metric.score}/{metric.weight}</td>
                    <td>{metric.reason_detail || metric.reason_code}</td>
                  </tr>
                ))}
                {!loading && !scorecard?.metrics.length ? (
                  <tr>
                    <td colSpan={4}>No scorecard metrics are available.</td>
                  </tr>
                ) : null}
              </tbody>
            </table>
          </div>
        </article>

        <article className="panel">
          <header className="panel-header">
            <div>
              <h3>Trust Badges</h3>
              <p>Sanitized trust claims derived from the same internal scorecard.</p>
            </div>
          </header>
          <div className="table-shell">
            <table>
              <thead>
                <tr>
                  <th>Badge</th>
                  <th>State</th>
                  <th>Public</th>
                  <th>Summary</th>
                </tr>
              </thead>
              <tbody>
                {report?.badges.map((badge) => (
                  <tr key={badge.id}>
                    <td>{badge.label}</td>
                    <td>{badge.state}</td>
                    <td>{badge.public_publishable ? "yes" : "internal-only"}</td>
                    <td>{badge.summary}</td>
                  </tr>
                ))}
                {!loading && !report?.badges.length ? (
                  <tr>
                    <td colSpan={4}>No trust badges are available.</td>
                  </tr>
                ) : null}
              </tbody>
            </table>
          </div>
        </article>
      </section>

      <section className="analytics-grid">
        <article className="panel">
          <header className="panel-header">
            <div>
              <h3>Hardening Review</h3>
              <p>Advisory-only continuous review findings for stale exceptions, signer gaps, vulnerability debt, and runtime containment.</p>
            </div>
          </header>
          <div className="table-shell">
            <table>
              <thead>
                <tr>
                  <th>Category</th>
                  <th>Severity</th>
                  <th>Status</th>
                  <th>Reason</th>
                </tr>
              </thead>
              <tbody>
                {report?.findings.map((finding) => (
                  <tr key={finding.id}>
                    <td>{finding.category}</td>
                    <td>{finding.severity}</td>
                    <td>{finding.status}</td>
                    <td>{finding.reason_detail || finding.reason_code}</td>
                  </tr>
                ))}
                {!loading && !report?.findings.length ? (
                  <tr>
                    <td colSpan={4}>No active hardening review findings.</td>
                  </tr>
                ) : null}
              </tbody>
            </table>
          </div>
        </article>

        <article className="panel">
          <header className="panel-header">
            <div>
              <h3>Standards Mapping</h3>
              <p>Readiness and evidence mapping only. These mappings do not claim formal certification.</p>
            </div>
          </header>
          <div className="table-shell">
            <table>
              <thead>
                <tr>
                  <th>Standard</th>
                  <th>Control</th>
                  <th>Status</th>
                  <th>Summary</th>
                </tr>
              </thead>
              <tbody>
                {report?.standards_mapping.map((mapping) => (
                  <tr key={`${mapping.standard}-${mapping.control}`}>
                    <td>{mapping.standard}</td>
                    <td>{mapping.control}</td>
                    <td>{mapping.status}</td>
                    <td>{mapping.summary}</td>
                  </tr>
                ))}
                {!loading && !report?.standards_mapping.length ? (
                  <tr>
                    <td colSpan={4}>No standards mappings are available.</td>
                  </tr>
                ) : null}
              </tbody>
            </table>
          </div>
        </article>
      </section>

      <section className="panel">
        <header className="panel-header">
          <div>
            <h3>Audit Export Readiness</h3>
            <p>Read-only audit reports and deterministic evidence export bundles are available through the audit API.</p>
          </div>
        </header>
        <div className="summary-grid">
          <article className="summary-card">
            <span className="summary-label">Report ID</span>
            <strong className="summary-value" style={{ fontSize: "1.05rem" }}>{report?.id || "-"}</strong>
          </article>
          <article className="summary-card">
            <span className="summary-label">Generated</span>
            <strong className="summary-value" style={{ fontSize: "1.05rem" }}>{formatTimestamp(report?.generated_at)}</strong>
          </article>
          <article className="summary-card">
            <span className="summary-label">Scope</span>
            <strong className="summary-value" style={{ fontSize: "1.05rem" }}>{report?.scope_ref || "-"}</strong>
          </article>
          <article className="summary-card">
            <span className="summary-label">Publication</span>
            <strong className="summary-value" style={{ fontSize: "1.05rem" }}>{scorecard?.publication_mode || "disabled"}</strong>
          </article>
          <article className="summary-card">
            <span className="summary-label">Public Preview</span>
            <strong className="summary-value" style={{ fontSize: "1.05rem" }}>{report?.public_view ? "available" : "disabled"}</strong>
          </article>
          <article className="summary-card">
            <span className="summary-label">Export Path</span>
            <strong className="summary-value" style={{ fontSize: "1.05rem" }}>/v1/audit/exports</strong>
          </article>
        </div>

        {report?.limitations?.length ? (
          <div className="panel-empty">
            <strong>Scope notes</strong>
            <ul className="summary-list">
              {report.limitations.map((item) => (
                <li key={item}>{item}</li>
              ))}
            </ul>
          </div>
        ) : null}
      </section>
    </>
  );
}
