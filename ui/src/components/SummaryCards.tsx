import { TopDenyReasons } from "./TopDenyReasons";
import type { Summary } from "../types";

type Props = {
  summary: Summary | null;
  loading: boolean;
};

const cards = [
  { key: "total_events", label: "Total Events", tone: "neutral" },
  { key: "total_allow", label: "ALLOW", tone: "allow" },
  { key: "total_deny", label: "DENY", tone: "deny" },
  { key: "total_error", label: "ERROR", tone: "error" },
  { key: "recent_runtime_drift_deny", label: "Runtime Drift (24h)", tone: "drift" },
] as const;

export function SummaryCards({ summary, loading }: Props) {
  const eventTypeEntries = summary ? Object.entries(summary.counts_by_event_type) : [];

  return (
    <section className="summary-grid">
      {cards.map((card) => (
        <article className={`summary-card summary-card--${card.tone}`} key={card.key}>
          <span className="summary-label">{card.label}</span>
          <strong className="summary-value">
            {loading ? "…" : summary ? summary[card.key] : "-"}
          </strong>
        </article>
      ))}

      <article className="summary-card summary-card--wide">
        <span className="summary-label">Top Deny Reasons</span>
        <TopDenyReasons items={summary?.top_deny_reasons || []} loading={loading} />
      </article>

      <article className="summary-card summary-card--wide">
        <span className="summary-label">Signal Mix</span>
        {loading ? (
          <div className="summary-list-empty">Loading…</div>
        ) : eventTypeEntries.length > 0 ? (
          <ul className="summary-list">
            {eventTypeEntries.map(([eventType, count]) => (
              <li key={eventType}>
                <span>{eventType}</span>
                <strong>{count}</strong>
              </li>
            ))}
          </ul>
        ) : (
          <div className="summary-list-empty">No event types recorded yet.</div>
        )}
      </article>
    </section>
  );
}
