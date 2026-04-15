import { useState, type FormEvent } from "react";

import type { ExceptionRequestInput } from "../types";

type Props = {
  enabled: boolean;
  submitting: boolean;
  onSubmit: (input: ExceptionRequestInput) => Promise<void>;
};

const initialState: ExceptionRequestInput = {
  exception_id: "",
  exception_type: "BREAK_GLASS",
  tenant_id: "",
  environment: "",
  namespace: "",
  repo: "",
  image_digest: "",
  cve_id: "",
  reason: "",
  ticket_id: "",
  ttl_hours: 2,
};

export function ExceptionRequestForm({ enabled, submitting, onSubmit }: Props) {
  const [form, setForm] = useState<ExceptionRequestInput>(initialState);
  const [error, setError] = useState<string | null>(null);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    if (!enabled) {
      return;
    }
    try {
      setError(null);
      await onSubmit({
        ...form,
        ttl_hours: form.ttl_hours ? Number(form.ttl_hours) : undefined,
      });
      setForm(initialState);
    } catch (submitError) {
      setError(submitError instanceof Error ? submitError.message : "Request failed.");
    }
  }

  return (
    <section className="panel exception-form-panel">
      <div className="table-toolbar">
        <span className="summary-label">Request Exception</span>
        <strong>{enabled ? "Enabled" : "Read only"}</strong>
      </div>
      <form className="exception-form" onSubmit={handleSubmit}>
        <label>
          <span>Exception ID</span>
          <input
            value={form.exception_id}
            onChange={(event) => setForm((current) => ({ ...current, exception_id: event.target.value }))}
            placeholder="EX-2026-700"
            disabled={!enabled || submitting}
          />
        </label>
        <label>
          <span>Type</span>
          <select
            value={form.exception_type}
            onChange={(event) => setForm((current) => ({ ...current, exception_type: event.target.value }))}
            disabled={!enabled || submitting}
          >
            <option value="BREAK_GLASS">BREAK_GLASS</option>
            <option value="DIGEST_BYPASS">DIGEST_BYPASS</option>
            <option value="CVE_WHITELIST">CVE_WHITELIST</option>
          </select>
        </label>
        <label>
          <span>Tenant</span>
          <input value={form.tenant_id || ""} onChange={(event) => setForm((current) => ({ ...current, tenant_id: event.target.value }))} disabled={!enabled || submitting} />
        </label>
        <label>
          <span>Environment</span>
          <input value={form.environment || ""} onChange={(event) => setForm((current) => ({ ...current, environment: event.target.value }))} disabled={!enabled || submitting} />
        </label>
        <label>
          <span>Namespace</span>
          <input value={form.namespace || ""} onChange={(event) => setForm((current) => ({ ...current, namespace: event.target.value }))} disabled={!enabled || submitting} />
        </label>
        <label>
          <span>Repo</span>
          <input value={form.repo || ""} onChange={(event) => setForm((current) => ({ ...current, repo: event.target.value }))} disabled={!enabled || submitting} />
        </label>
        <label>
          <span>Digest</span>
          <input value={form.image_digest || ""} onChange={(event) => setForm((current) => ({ ...current, image_digest: event.target.value }))} disabled={!enabled || submitting} />
        </label>
        <label>
          <span>CVE</span>
          <input value={form.cve_id || ""} onChange={(event) => setForm((current) => ({ ...current, cve_id: event.target.value }))} disabled={!enabled || submitting} />
        </label>
        <label>
          <span>Ticket</span>
          <input value={form.ticket_id} onChange={(event) => setForm((current) => ({ ...current, ticket_id: event.target.value }))} disabled={!enabled || submitting} />
        </label>
        <label>
          <span>TTL Hours</span>
          <input
            type="number"
            min="1"
            value={form.ttl_hours || 2}
            onChange={(event) => setForm((current) => ({ ...current, ttl_hours: Number(event.target.value) }))}
            disabled={!enabled || submitting}
          />
        </label>
        <label className="exception-form__wide">
          <span>Reason</span>
          <textarea
            value={form.reason}
            onChange={(event) => setForm((current) => ({ ...current, reason: event.target.value }))}
            rows={3}
            disabled={!enabled || submitting}
          />
        </label>
        {error ? <p className="panel-error">{error}</p> : null}
        <div className="filters-actions">
          <button className="button button--primary" type="submit" disabled={!enabled || submitting}>
            {submitting ? "Submitting…" : "Request Exception"}
          </button>
        </div>
      </form>
    </section>
  );
}
