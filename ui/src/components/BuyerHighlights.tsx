import type { Summary } from "../types";

type Props = {
  summary: Summary | null;
  loading: boolean;
};

function metric(summary: Summary | null, key: string) {
  if (!summary) {
    return 0;
  }
  return summary.counts_by_event_type[key] ?? 0;
}

export function BuyerHighlights({ summary, loading }: Props) {
  const cards = [
    {
      label: "Blocked Before Deploy",
      value: summary?.total_deny ?? 0,
      tone: "deny",
      description: "Risky changes and untrusted artifacts were stopped before they reached runtime.",
    },
    {
      label: "Verified Decisions",
      value: summary?.total_allow ?? 0,
      tone: "allow",
      description: "Allowed operations remained explainable and traceable through the audit store.",
    },
    {
      label: "Monitored Signals",
      value:
        metric(summary, "artifact_verification_result") +
        metric(summary, "deploy_gate_decision") +
        metric(summary, "runtime_drift_result"),
      tone: "neutral",
      description: "Verification, admission, and runtime drift are visible in one operator surface.",
    },
  ] as const;

  return (
    <section className="highlights-grid">
      {cards.map((card) => (
        <article className={`panel highlight-card highlight-card--${card.tone}`} key={card.label}>
          <span className="summary-label">{card.label}</span>
          <strong className="highlight-value">{loading ? "…" : card.value}</strong>
          <p>{card.description}</p>
        </article>
      ))}
    </section>
  );
}
