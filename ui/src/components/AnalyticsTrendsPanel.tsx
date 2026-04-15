import type { TrendsResponse } from "../types";

type Props = {
  trends: TrendsResponse | null;
  loading: boolean;
};

function maxBucketValue(trends: TrendsResponse | null) {
  if (!trends || trends.buckets.length === 0) {
    return 1;
  }
  return Math.max(
    ...trends.buckets.map((bucket) => Math.max(bucket.allow_count, bucket.deny_count, bucket.error_count, 1)),
  );
}

function formatBucketLabel(timestamp: string) {
  const date = new Date(timestamp);
  return `${date.getMonth() + 1}/${date.getDate()}`;
}

export function AnalyticsTrendsPanel({ trends, loading }: Props) {
  if (loading) {
    return <section className="panel panel-empty">Loading analytics trends…</section>;
  }
  if (!trends || trends.buckets.length === 0) {
    return <section className="panel panel-empty">No trend data available for the current filters.</section>;
  }

  const maxValue = maxBucketValue(trends);

  return (
    <section className="panel analytics-panel">
      <div className="table-toolbar">
        <span className="summary-label">Decision Trends</span>
        <strong>{trends.buckets.length} buckets</strong>
      </div>
      <div className="trend-totals">
        <span className="chip chip--allow">ALLOW {trends.totals.allow ?? 0}</span>
        <span className="chip chip--deny">DENY {trends.totals.deny ?? 0}</span>
        <span className="chip chip--error">ERROR {trends.totals.error ?? 0}</span>
      </div>
      <div className="trend-bars">
        {trends.buckets.map((bucket) => (
          <div className="trend-bar-group" key={bucket.timestamp}>
            <div className="trend-bar-group__stack">
              <span
                className="trend-bar trend-bar--allow"
                style={{ height: `${Math.max((bucket.allow_count / maxValue) * 100, 4)}%` }}
                title={`ALLOW ${bucket.allow_count}`}
              />
              <span
                className="trend-bar trend-bar--deny"
                style={{ height: `${Math.max((bucket.deny_count / maxValue) * 100, 4)}%` }}
                title={`DENY ${bucket.deny_count}`}
              />
              <span
                className="trend-bar trend-bar--error"
                style={{ height: `${Math.max((bucket.error_count / maxValue) * 100, 4)}%` }}
                title={`ERROR ${bucket.error_count}`}
              />
            </div>
            <span className="trend-bar-group__label">{formatBucketLabel(bucket.timestamp)}</span>
          </div>
        ))}
      </div>
    </section>
  );
}
