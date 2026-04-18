import { useEffect, useState } from "react";

import { getAIGuidance, getAIInsights } from "../api";
import type { AIInsightsResponse, GuidanceItem, GuidanceResponse } from "../types";

type Props = {
  tenantID?: string;
};

function docsLabel(ref?: string) {
  if (!ref) {
    return "";
  }
  return ref.replace(/^docs\//, "").replace(/\.md$/, "");
}

export function AIInsightsPanel({ tenantID }: Props) {
  const [guidance, setGuidance] = useState<GuidanceResponse | null>(null);
  const [insights, setInsights] = useState<AIInsightsResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;

    async function load() {
      setLoading(true);
      setError(null);
      try {
        const params = tenantID ? { tenant_id: tenantID, limit: "12" } : { limit: "12" };
        const [guidanceResult, insightsResult] = await Promise.all([
          getAIGuidance(params),
          getAIInsights(params),
        ]);
        if (cancelled) {
          return;
        }
        setGuidance(guidanceResult);
        setInsights(insightsResult);
      } catch (loadError) {
        if (!cancelled) {
          setError(loadError instanceof Error ? loadError.message : "Unable to load contextual guidance.");
          setGuidance(null);
          setInsights(null);
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

  const items = guidance?.items || [];
  const vexDrafts = items.filter((item) => item.vex_draft);
  const breakGlass = items.filter((item) => item.break_glass_guidance);

  return (
    <>
      <section className="summary-grid">
        <article className="summary-card">
          <span className="summary-label">Guidance Mode</span>
          <strong className="summary-value">{guidance?.summary.guidance_mode || "disabled"}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Top Priority Items</span>
          <strong className="summary-value">{insights?.top_items.length || 0}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Deterministic Only</span>
          <strong className="summary-value">{guidance?.summary.deterministic_only ? "yes" : "no"}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">VEX Draft Candidates</span>
          <strong className="summary-value">{vexDrafts.length}</strong>
        </article>
      </section>

      {error ? (
        <section className="panel status-banner">
          <div>
            <strong>AI guidance is unavailable.</strong>
            <p>{error}</p>
          </div>
        </section>
      ) : null}

      <section className="analytics-grid">
        <article className="panel">
          <header className="panel-header">
            <div>
              <h3>Prioritized Guidance</h3>
              <p>Advisory grouping, contextual priority, and bounded remediation guidance derived from deterministic ChangeLock findings.</p>
            </div>
          </header>
          <div className="table-shell">
            <table>
              <thead>
                <tr>
                  <th>Priority</th>
                  <th>Category</th>
                  <th>Confidence</th>
                  <th>Recommendation</th>
                </tr>
              </thead>
              <tbody>
                {items.map((item) => (
                  <tr key={item.id}>
                    <td>{item.priority}</td>
                    <td>{item.category}</td>
                    <td>{item.confidence}</td>
                    <td>{item.recommendation_summary || item.explanation || "-"}</td>
                  </tr>
                ))}
                {!loading && items.length === 0 ? (
                  <tr>
                    <td colSpan={4}>No contextual guidance items are available for the current scope.</td>
                  </tr>
                ) : null}
              </tbody>
            </table>
          </div>
        </article>

        <article className="panel">
          <header className="panel-header">
            <div>
              <h3>Insight Limitations</h3>
              <p>Guidance remains advisory. Deterministic trust decisions, evidence, and enforcement still stay authoritative.</p>
            </div>
          </header>
          <ul className="panel-list">
            {(guidance?.summary.limitations || []).map((item) => (
              <li key={item}>{item}</li>
            ))}
            {!loading && !(guidance?.summary.limitations || []).length ? <li>No additional limitations were reported.</li> : null}
          </ul>
        </article>
      </section>

      <section className="analytics-grid">
        <article className="panel">
          <header className="panel-header">
            <div>
              <h3>VEX Draft Candidates</h3>
              <p>Review-only draft candidates. These are not authoritative VEX records and are never auto-published.</p>
            </div>
          </header>
          <div className="table-shell">
            <table>
              <thead>
                <tr>
                  <th>Candidate</th>
                  <th>Confidence</th>
                  <th>Justification</th>
                </tr>
              </thead>
              <tbody>
                {vexDrafts.map((item) => (
                  <tr key={item.id}>
                    <td>{item.vex_draft?.candidate_status}</td>
                    <td>{item.vex_draft?.confidence}</td>
                    <td>{item.vex_draft?.justification}</td>
                  </tr>
                ))}
                {!loading && vexDrafts.length === 0 ? (
                  <tr>
                    <td colSpan={3}>No VEX draft candidates are active for the current scope.</td>
                  </tr>
                ) : null}
              </tbody>
            </table>
          </div>
        </article>

        <article className="panel">
          <header className="panel-header">
            <div>
              <h3>Break-Glass Guidance</h3>
              <p>Scope explanation and cleanup reminders for active break-glass posture. Guidance does not approve or extend exceptions.</p>
            </div>
          </header>
          <div className="table-shell">
            <table>
              <thead>
                <tr>
                  <th>Scope</th>
                  <th>Narrower Alternative</th>
                  <th>Confidence</th>
                </tr>
              </thead>
              <tbody>
                {breakGlass.map((item) => (
                  <tr key={item.id}>
                    <td>{item.break_glass_guidance?.scope_explanation}</td>
                    <td>{item.break_glass_guidance?.narrower_alternative || "-"}</td>
                    <td>{item.break_glass_guidance?.confidence}</td>
                  </tr>
                ))}
                {!loading && breakGlass.length === 0 ? (
                  <tr>
                    <td colSpan={3}>No break-glass guidance items are active for the current scope.</td>
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
            <h3>Detailed Guidance</h3>
            <p>Each item references deterministic reason codes, evidence scope, and bounded remediation steps.</p>
          </div>
        </header>
        <div className="table-shell">
          <table>
            <thead>
              <tr>
                <th>Grouping</th>
                <th>Why</th>
                <th>Safer Alternative</th>
                <th>Docs</th>
              </tr>
            </thead>
            <tbody>
              {items.map((item: GuidanceItem) => (
                <tr key={`${item.id}-detail`}>
                  <td>{item.grouping.label}</td>
                  <td>{item.explanation || item.impact_summary || "-"}</td>
                  <td>{item.safer_alternative || "-"}</td>
                  <td>{(item.docs_refs || []).slice(0, 2).map(docsLabel).filter(Boolean).join(", ") || "-"}</td>
                </tr>
              ))}
              {!loading && items.length === 0 ? (
                <tr>
                  <td colSpan={4}>No detailed contextual guidance is available.</td>
                </tr>
              ) : null}
            </tbody>
          </table>
        </div>
      </section>
    </>
  );
}
