import { useState } from "react";

import type { StoredEvent } from "../types";

type Props = {
  event: StoredEvent | null;
};

function formatTimestamp(timestamp?: string) {
  if (!timestamp) {
    return "-";
  }
  return new Date(timestamp).toLocaleString();
}

function formatRelativeTime(timestamp?: string) {
  if (!timestamp) {
    return "";
  }

  const diff = new Date(timestamp).getTime() - Date.now();
  const formatter = new Intl.RelativeTimeFormat(undefined, { numeric: "auto" });
  const minutes = Math.round(diff / 60000);

  if (Math.abs(minutes) < 60) {
    return formatter.format(minutes, "minute");
  }

  const hours = Math.round(minutes / 60);
  if (Math.abs(hours) < 24) {
    return formatter.format(hours, "hour");
  }

  const days = Math.round(hours / 24);
  return formatter.format(days, "day");
}

export function EventDetails({ event }: Props) {
  const [copied, setCopied] = useState(false);
  const evidence = event?.evidence ? JSON.stringify(event.evidence, null, 2) : null;
  const rawEvent = event?.raw_event ? JSON.stringify(event.raw_event, null, 2) : null;

  async function copyRequestID() {
    if (!event?.request_id) {
      return;
    }
    try {
      await navigator.clipboard.writeText(event.request_id);
      setCopied(true);
      window.setTimeout(() => setCopied(false), 1500);
    } catch {
      setCopied(false);
    }
  }

  if (!event) {
    return (
      <aside className="panel details-panel details-panel--empty">
        Select an event to inspect reasons, verifier output, and evidence.
      </aside>
    );
  }

  return (
    <aside className="panel details-panel">
      <div className="details-header">
        <div>
          <h2>Event Details</h2>
          <p>{event.component} · {event.event_type}</p>
        </div>
        <div className="details-header__actions">
          {event.request_id ? (
            <button className="button button--ghost" onClick={copyRequestID}>
              {copied ? "Copied" : "Copy Request ID"}
            </button>
          ) : null}
          <span className={`badge badge--${event.decision.toLowerCase()}`}>{event.decision}</span>
        </div>
      </div>

      <dl className="details-grid">
        <div>
          <dt>Request ID</dt>
          <dd>
            <code>{event.request_id || "-"}</code>
          </dd>
        </div>
        <div>
          <dt>Timestamp</dt>
          <dd>
            {formatTimestamp(event.timestamp || event.received_at)}
            <span className="details-subtext">{formatRelativeTime(event.timestamp || event.received_at)}</span>
          </dd>
        </div>
        <div>
          <dt>Repo</dt>
          <dd>{event.repo || "-"}</dd>
        </div>
        <div>
          <dt>Environment</dt>
          <dd>{event.environment || "-"}</dd>
        </div>
        <div>
          <dt>Tenant</dt>
          <dd>{event.tenant_id || "-"}</dd>
        </div>
        <div>
          <dt>Actor</dt>
          <dd>{event.actor || "-"}</dd>
        </div>
        <div>
          <dt>Policy Version</dt>
          <dd>{event.policy_version || "-"}</dd>
        </div>
        <div>
          <dt>Drift Result</dt>
          <dd>{event.drift_result || "-"}</dd>
        </div>
        <div>
          <dt>Image</dt>
          <dd className="details-break">{event.image || "-"}</dd>
        </div>
        <div>
          <dt>Digest</dt>
          <dd>
            <code>{event.digest || "-"}</code>
          </dd>
        </div>
      </dl>

      {event.drift_classes && event.drift_classes.length > 0 ? (
        <section className="details-section">
          <h3>Drift Classes</h3>
          <div className="chip-row">
            {event.drift_classes.map((value) => (
              <span className="chip" key={value}>
                {value}
              </span>
            ))}
          </div>
        </section>
      ) : null}

      <section className="details-section">
        <h3>Reasons</h3>
        {event.reasons && event.reasons.length > 0 ? (
          <ul className="details-list">
            {event.reasons.map((reason) => (
              <li key={reason}>{reason}</li>
            ))}
          </ul>
        ) : (
          <div className="details-empty">No reasons recorded.</div>
        )}
      </section>

      <section className="details-section">
        <h3>Verifier Summary</h3>
        {event.verifier_summary ? (
          <dl className="details-grid details-grid--compact">
            <div>
              <dt>Signature</dt>
              <dd>{event.verifier_summary.signature_valid ? "valid" : "not valid"}</dd>
            </div>
            <div>
              <dt>Attestation</dt>
              <dd>{event.verifier_summary.attestation_valid ? "valid" : "not valid"}</dd>
            </div>
          </dl>
        ) : (
          <div className="details-empty">No verifier summary recorded.</div>
        )}
      </section>

      <section className="details-section">
        <h3>Evidence</h3>
        {evidence ? <pre className="json-block">{evidence}</pre> : <div className="details-empty">No evidence recorded.</div>}
      </section>

      <section className="details-section">
        <h3>Raw Event</h3>
        {rawEvent ? <pre className="json-block">{rawEvent}</pre> : <div className="details-empty">No raw event payload recorded.</div>}
      </section>
    </aside>
  );
}
