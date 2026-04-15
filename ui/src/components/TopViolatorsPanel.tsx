import type { TopViolatorsResponse } from "../types";

type Props = {
  data: TopViolatorsResponse | null;
  loading: boolean;
};

export function TopViolatorsPanel({ data, loading }: Props) {
  if (loading) {
    return <section className="panel panel-empty">Loading top violators…</section>;
  }
  if (!data || data.items.length === 0) {
    return <section className="panel panel-empty">No deny-heavy repos, tenants, or environments found.</section>;
  }

  return (
    <section className="panel analytics-panel">
      <div className="table-toolbar">
        <span className="summary-label">Top Violators</span>
        <strong>{data.items.length}</strong>
      </div>
      <ul className="analytics-list">
        {data.items.map((item) => (
          <li key={item.key} className="analytics-list__item">
            <div>
              <strong>{item.key}</strong>
              <p>
                {item.top_reasons.length > 0
                  ? item.top_reasons.map((reason) => `${reason.reason} (${reason.count})`).join(" · ")
                  : "No deny reasons captured"}
              </p>
            </div>
            <span className="badge badge--deny">{item.deny_count}</span>
          </li>
        ))}
      </ul>
    </section>
  );
}
