import type { ReasonCount } from "../types";

type Props = {
  items: ReasonCount[];
  loading: boolean;
};

export function TopDenyReasons({ items, loading }: Props) {
  if (loading) {
    return <div className="summary-list-empty">Loading…</div>;
  }

  if (items.length === 0) {
    return <div className="summary-list-empty">No deny reasons yet.</div>;
  }

  return (
    <ul className="summary-list">
      {items.map((item) => (
        <li key={item.reason}>
          <span>{item.reason}</span>
          <strong>{item.count}</strong>
        </li>
      ))}
    </ul>
  );
}
