import type { StoredEvent } from "../types";

type Props = {
  events: StoredEvent[];
  selectedEventID: number | null;
  onSelect: (event: StoredEvent) => void;
  loading: boolean;
  error: string | null;
};

function formatTimestamp(timestamp?: string, fallback?: string) {
  const value = timestamp || fallback;
  if (!value) {
    return "-";
  }

  const date = new Date(value);
  return `${date.toLocaleDateString()} ${date.toLocaleTimeString()}`;
}

function displayValue(...values: Array<string | undefined>) {
  for (const value of values) {
    if (value) {
      return value;
    }
  }
  return "-";
}

export function EventsTable({ events, selectedEventID, onSelect, loading, error }: Props) {
  if (loading) {
    return <section className="panel panel-empty">Loading events from the audit store…</section>;
  }

  if (error) {
    return <section className="panel panel-empty panel-error">Unable to load events. {error}</section>;
  }

  if (events.length === 0) {
    return <section className="panel panel-empty">No events matched the current view.</section>;
  }

  return (
    <section className="panel table-panel">
      <div className="table-toolbar">
        <span className="summary-label">Recent Events</span>
        <strong>{events.length}</strong>
      </div>
      <div className="table-scroll">
        <table className="events-table">
          <thead>
            <tr>
              <th>Timestamp</th>
              <th>Decision</th>
              <th>Component</th>
              <th>Event Type</th>
              <th>Repo</th>
              <th>Environment</th>
              <th>Tenant</th>
              <th>Image / Digest</th>
            </tr>
          </thead>
          <tbody>
            {events.map((event) => (
              <tr
                key={event.id}
                className={selectedEventID === event.id ? "is-selected" : undefined}
                onClick={() => onSelect(event)}
              >
                <td>{formatTimestamp(event.timestamp, event.received_at)}</td>
                <td>
                  <span className={`badge badge--${event.decision.toLowerCase()}`}>{event.decision}</span>
                </td>
                <td title={event.component}>{event.component}</td>
                <td>
                  <div className="event-meta-primary" title={event.event_type}>
                    {event.event_type}
                  </div>
                  {event.drift_result ? <span className="chip chip--tight">{event.drift_result}</span> : null}
                </td>
                <td>
                  <span className="truncate" title={displayValue(event.repo, event.namespace)}>
                    {displayValue(event.repo, event.namespace)}
                  </span>
                </td>
                <td>{displayValue(event.environment)}</td>
                <td>{displayValue(event.tenant_id)}</td>
                <td className="events-table__artifact">
                  <div className="truncate" title={displayValue(event.image)}>
                    {displayValue(event.image)}
                  </div>
                  <code className="truncate" title={displayValue(event.digest)}>
                    {displayValue(event.digest)}
                  </code>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </section>
  );
}
