import { useState } from "react";

import type { PolicyException } from "../types";

type Props = {
  pending: PolicyException[];
  canApprove: boolean;
  canRevoke: boolean;
  loading: boolean;
  onApprove: (exceptionID: string) => Promise<void>;
  onReject: (exceptionID: string, reason: string) => Promise<void>;
  onRevoke: (exceptionID: string) => Promise<void>;
};

export function PendingExceptionsPanel({ pending, canApprove, canRevoke, loading, onApprove, onReject, onRevoke }: Props) {
  const [busyID, setBusyID] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  async function runAction(action: () => Promise<void>, exceptionID: string) {
    try {
      setBusyID(exceptionID);
      setError(null);
      await action();
    } catch (actionError) {
      setError(actionError instanceof Error ? actionError.message : "Action failed.");
    } finally {
      setBusyID(null);
    }
  }

  return (
    <section className="panel analytics-panel">
      <div className="table-toolbar">
        <span className="summary-label">Pending Exception Queue</span>
        <strong>{loading ? "…" : pending.length}</strong>
      </div>
      {error ? <p className="panel-error">{error}</p> : null}
      {loading ? (
        <div className="panel-empty">Loading pending exceptions…</div>
      ) : pending.length === 0 ? (
        <div className="panel-empty">No pending exception requests.</div>
      ) : (
        <ul className="analytics-list">
          {pending.map((exception) => (
            <li key={exception.exception_id} className="exception-queue-item">
              <div>
                <div className="exception-queue-item__title">
                  <strong>{exception.exception_id}</strong>
                  <span className="chip">{exception.exception_type}</span>
                  <span className="chip">{exception.status}</span>
                </div>
                <p>{exception.reason}</p>
                <small>
                  {exception.requested_by || "unknown requester"} · {exception.ticket_id} · expires{" "}
                  {new Date(exception.expires_at).toLocaleString()}
                </small>
              </div>
              <div className="exception-queue-item__actions">
                {canApprove ? (
                  <>
                    <button
                      className="button button--primary"
                      disabled={busyID === exception.exception_id}
                      onClick={() => runAction(() => onApprove(exception.exception_id), exception.exception_id)}
                    >
                      Approve
                    </button>
                    <button
                      className="button"
                      disabled={busyID === exception.exception_id}
                      onClick={() => {
                        const reason = window.prompt("Rejection reason", "Insufficient scope or evidence");
                        if (!reason) {
                          return;
                        }
                        void runAction(() => onReject(exception.exception_id, reason), exception.exception_id);
                      }}
                    >
                      Reject
                    </button>
                  </>
                ) : null}
                {canRevoke ? (
                  <button
                    className="button"
                    disabled={busyID === exception.exception_id}
                    onClick={() => runAction(() => onRevoke(exception.exception_id), exception.exception_id)}
                  >
                    Revoke
                  </button>
                ) : null}
              </div>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
