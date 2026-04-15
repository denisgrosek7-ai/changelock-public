import { useState } from "react";

import type { StoredEvent } from "../types";

type Props = {
  event: StoredEvent | null;
};

function asRecord(value: unknown): Record<string, unknown> | null {
  if (!value || typeof value !== "object" || Array.isArray(value)) {
    return null;
  }
  return value as Record<string, unknown>;
}

function readString(value: unknown): string | null {
  return typeof value === "string" && value.trim() !== "" ? value : null;
}

function formatVulnerabilitySummary(value: unknown) {
  const summary = asRecord(value);
  if (!summary) {
    return null;
  }

  const parts = ["critical", "high", "medium", "low", "unknown"]
    .map((key) => {
      const raw = summary[key];
      return typeof raw === "number" && raw > 0 ? `${key}:${raw}` : null;
    })
    .filter(Boolean);

  const total = typeof summary.total === "number" ? `total:${summary.total}` : null;
  return [...parts, total].filter(Boolean).join(" · ");
}

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
  const artifactEvidence = asRecord(asRecord(event?.evidence)?.artifact);
  const sbomArtifactRef = readString(artifactEvidence?.sbom_artifact_ref);
  const sbomDigestRef = readString(artifactEvidence?.sbom_digest_ref);
  const sbomFormat = readString(artifactEvidence?.sbom_format);
  const sbomHash = readString(artifactEvidence?.sbom_hash);
  const vulnerabilityStatus = readString(artifactEvidence?.vulnerability_scan_status);
  const vulnerabilityTool = readString(artifactEvidence?.vulnerability_scan_tool);
  const vulnerabilityThreshold = readString(artifactEvidence?.vulnerability_scan_severity_threshold);
  const vulnerabilityReportRef = readString(artifactEvidence?.vulnerability_report_ref);
  const vulnerabilitySummary = formatVulnerabilitySummary(artifactEvidence?.vulnerability_summary);

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
          <dt>Policy Bundle</dt>
          <dd>{event.policy_bundle_id || "-"}</dd>
        </div>
        <div>
          <dt>Bundle Hash</dt>
          <dd>
            <code>{event.policy_bundle_hash || "-"}</code>
          </dd>
        </div>
        <div>
          <dt>Decision Hash</dt>
          <dd>
            <code>{event.decision_hash || "-"}</code>
          </dd>
        </div>
        <div>
          <dt>Exception</dt>
          <dd>{event.is_exception ? "yes" : "no"}</dd>
        </div>
        <div>
          <dt>Exception ID</dt>
          <dd>{event.exception_id || "-"}</dd>
        </div>
        <div>
          <dt>Exception Type</dt>
          <dd>{event.exception_type || "-"}</dd>
        </div>
        <div>
          <dt>Exception Status</dt>
          <dd>{event.exception_status || "-"}</dd>
        </div>
        <div>
          <dt>Exception Ticket</dt>
          <dd>{event.exception_ticket_id || "-"}</dd>
        </div>
        <div>
          <dt>Requested By</dt>
          <dd>{event.exception_requested_by || "-"}</dd>
        </div>
        <div>
          <dt>Requested At</dt>
          <dd>{formatTimestamp(event.exception_requested_at)}</dd>
        </div>
        <div>
          <dt>Approved By</dt>
          <dd>{event.exception_approved_by || "-"}</dd>
        </div>
        <div>
          <dt>Approved At</dt>
          <dd>{formatTimestamp(event.exception_approved_at)}</dd>
        </div>
        <div>
          <dt>Rejected By</dt>
          <dd>{event.exception_rejected_by || "-"}</dd>
        </div>
        <div>
          <dt>Rejected At</dt>
          <dd>{formatTimestamp(event.exception_rejected_at)}</dd>
        </div>
        <div>
          <dt>Rejection Reason</dt>
          <dd>{event.exception_rejection_reason || "-"}</dd>
        </div>
        <div>
          <dt>Exception Expires</dt>
          <dd>{formatTimestamp(event.exception_expires_at)}</dd>
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
        <div>
          <dt>CVE ID</dt>
          <dd>{event.cve_id || "-"}</dd>
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

      {event.exception_id || event.exception_reason ? (
        <section className="details-section">
          <h3>Exception Metadata</h3>
          <dl className="details-grid details-grid--compact">
            <div>
              <dt>Reason</dt>
              <dd>{event.exception_reason || "-"}</dd>
            </div>
            <div>
              <dt>Ticket</dt>
              <dd>{event.exception_ticket_id || "-"}</dd>
            </div>
            <div>
              <dt>Approved By</dt>
              <dd>{event.exception_approved_by || "-"}</dd>
            </div>
            <div>
              <dt>Expires</dt>
              <dd>{formatTimestamp(event.exception_expires_at)}</dd>
            </div>
          </dl>
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
        <h3>Supply Chain Evidence</h3>
        {sbomArtifactRef || sbomDigestRef || vulnerabilityStatus ? (
          <dl className="details-grid details-grid--compact">
            <div>
              <dt>SBOM</dt>
              <dd>{sbomArtifactRef || "-"}</dd>
            </div>
            <div>
              <dt>SBOM Format</dt>
              <dd>{sbomFormat || "-"}</dd>
            </div>
            <div>
              <dt>SBOM Digest Ref</dt>
              <dd className="details-break">{sbomDigestRef || "-"}</dd>
            </div>
            <div>
              <dt>SBOM Hash</dt>
              <dd>
                <code>{sbomHash || "-"}</code>
              </dd>
            </div>
            <div>
              <dt>Scan Status</dt>
              <dd>{vulnerabilityStatus || "-"}</dd>
            </div>
            <div>
              <dt>Scan Tool</dt>
              <dd>{vulnerabilityTool || "-"}</dd>
            </div>
            <div>
              <dt>Severity Threshold</dt>
              <dd>{vulnerabilityThreshold || "-"}</dd>
            </div>
            <div>
              <dt>Report Ref</dt>
              <dd>{vulnerabilityReportRef || "-"}</dd>
            </div>
            <div>
              <dt>Vulnerability Summary</dt>
              <dd>{vulnerabilitySummary || "-"}</dd>
            </div>
          </dl>
        ) : (
          <div className="details-empty">No SBOM or vulnerability evidence recorded.</div>
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
