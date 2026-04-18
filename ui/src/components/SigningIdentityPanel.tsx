import { useEffect, useState } from "react";

import {
  getSigningIdentityFindings,
  getSigningIdentityObservations,
  getSigningIdentityPolicies,
  getSigningIdentityStatus,
} from "../api";
import type {
  SigningIdentityFinding,
  SigningIdentityObservation,
  SigningIdentityPolicy,
  SigningIdentityStatus,
} from "../types";

type Props = {
  tenantID?: string;
};

function formatTimestamp(value?: string) {
  if (!value) {
    return "-";
  }
  return new Date(value).toLocaleString();
}

export function SigningIdentityPanel({ tenantID }: Props) {
  const [status, setStatus] = useState<SigningIdentityStatus | null>(null);
  const [policies, setPolicies] = useState<SigningIdentityPolicy[]>([]);
  const [observations, setObservations] = useState<SigningIdentityObservation[]>([]);
  const [findings, setFindings] = useState<SigningIdentityFinding[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;

    async function load() {
      setLoading(true);
      setError(null);
      try {
        const params = tenantID ? { tenant_id: tenantID, limit: "50" } : { limit: "50" };
        const [statusResult, policiesResult, observationsResult, findingsResult] = await Promise.all([
          getSigningIdentityStatus(params),
          getSigningIdentityPolicies(),
          getSigningIdentityObservations(params),
          getSigningIdentityFindings(params),
        ]);
        if (cancelled) {
          return;
        }
        setStatus(statusResult);
        setPolicies(policiesResult.policies);
        setObservations(observationsResult.items);
        setFindings(findingsResult.items);
      } catch (loadError) {
        if (!cancelled) {
          setError(loadError instanceof Error ? loadError.message : "Unable to load signing identities.");
          setStatus(null);
          setPolicies([]);
          setObservations([]);
          setFindings([]);
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

  return (
    <>
      <section className="summary-grid">
        <article className="summary-card">
          <span className="summary-label">Enforcement</span>
          <strong className="summary-value">{status?.enforcement_mode || "disabled"}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Observed Signers</span>
          <strong className="summary-value">{status?.observed_identities || 0}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Authorized</span>
          <strong className="summary-value">{status?.authorized || 0}</strong>
        </article>
        <article className="summary-card">
          <span className="summary-label">Unauthorized / Unknown</span>
          <strong className="summary-value">{(status?.unauthorized || 0) + (status?.unknown || 0)}</strong>
        </article>
      </section>

      {error ? (
        <section className="panel status-banner">
          <div>
            <strong>Signing identity monitoring is unavailable.</strong>
            <p>{error}</p>
          </div>
        </section>
      ) : null}

      <section className="analytics-grid">
        <article className="panel">
          <header className="panel-header">
            <div>
              <h3>Observed Identities</h3>
              <p>Recently observed signing identities, authorization status, and transparency state.</p>
            </div>
          </header>
          <div className="table-shell">
            <table>
              <thead>
                <tr>
                  <th>Signer</th>
                  <th>Repository</th>
                  <th>Workflow</th>
                  <th>Status</th>
                  <th>Last Seen</th>
                </tr>
              </thead>
              <tbody>
                {observations.map((item) => (
                  <tr key={item.id}>
                    <td>{item.signer_identity || "-"}</td>
                    <td>{item.repository || "-"}</td>
                    <td>{item.workflow || "-"}</td>
                    <td>{item.authorized}{item.reason_code ? ` · ${item.reason_code}` : ""}</td>
                    <td>{formatTimestamp(item.last_seen_at)}</td>
                  </tr>
                ))}
                {!loading && observations.length === 0 ? (
                  <tr>
                    <td colSpan={5}>No signing identity observations recorded.</td>
                  </tr>
                ) : null}
              </tbody>
            </table>
          </div>
        </article>

        <article className="panel">
          <header className="panel-header">
            <div>
              <h3>Findings</h3>
              <p>Unauthorized, distrusted, or transparency-related signer findings plus workflow drift advisories.</p>
            </div>
          </header>
          <div className="table-shell">
            <table>
              <thead>
                <tr>
                  <th>Type</th>
                  <th>Severity</th>
                  <th>Workflow</th>
                  <th>Reason</th>
                </tr>
              </thead>
              <tbody>
                {findings.map((item) => (
                  <tr key={item.id}>
                    <td>{item.type}</td>
                    <td>{item.severity}</td>
                    <td>{item.workflow || "-"}</td>
                    <td>{item.reason}</td>
                  </tr>
                ))}
                {!loading && findings.length === 0 ? (
                  <tr>
                    <td colSpan={4}>No signer findings detected.</td>
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
            <h3>Authorized Policies</h3>
            <p>Explicit GitHub OIDC signer policies currently recorded in the control plane.</p>
          </div>
        </header>
        <div className="table-shell">
          <table>
            <thead>
              <tr>
                <th>Name</th>
                <th>Repository</th>
                <th>Workflow</th>
                <th>Ref</th>
                <th>State</th>
              </tr>
            </thead>
            <tbody>
              {policies.map((policy) => (
                <tr key={policy.id}>
                  <td>{policy.name || policy.id}</td>
                  <td>{policy.repository || "-"}</td>
                  <td>{policy.workflow || "-"}</td>
                  <td>{policy.ref || "-"}</td>
                  <td>{policy.enabled ? "enabled" : "disabled"}{policy.distrusted_after ? ` · distrusted ${formatTimestamp(policy.distrusted_after)}` : ""}</td>
                </tr>
              ))}
              {!loading && policies.length === 0 ? (
                <tr>
                  <td colSpan={5}>No signing identity policies recorded.</td>
                </tr>
              ) : null}
            </tbody>
          </table>
        </div>
      </section>
    </>
  );
}
