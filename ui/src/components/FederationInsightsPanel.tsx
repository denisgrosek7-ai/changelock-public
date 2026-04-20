import type { FederationGlobalView } from "../types";

type Props = {
  view: FederationGlobalView | null;
  loading: boolean;
  focusPeerID?: string | null;
};

function renderEmpty(message: string) {
  return <div className="summary-list-empty">{message}</div>;
}

function formatTimestamp(value?: string) {
  if (!value) {
    return "n/a";
  }
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) {
    return value;
  }
  return date.toLocaleString();
}

export function FederationInsightsPanel({ view, loading, focusPeerID }: Props) {
  if (loading) {
    return <section className="panel analytics-panel analytics-panel--wide">Loading federation trust state…</section>;
  }

  const visiblePeers = view ? [...view.peers].sort((left, right) => {
    const leftFocused = focusPeerID && left.peer_id === focusPeerID;
    const rightFocused = focusPeerID && right.peer_id === focusPeerID;
    if (leftFocused !== rightFocused) {
      return leftFocused ? -1 : 1;
    }
    return left.peer_id.localeCompare(right.peer_id);
  }) : [];
  const visibleProofHistory = view ? [...view.proof_history].sort((left, right) => {
    const leftFocused = focusPeerID && left.peer_id === focusPeerID;
    const rightFocused = focusPeerID && right.peer_id === focusPeerID;
    if (leftFocused !== rightFocused) {
      return leftFocused ? -1 : 1;
    }
    return left.request_id.localeCompare(right.request_id);
  }) : [];

  return (
    <>
      <section className="panel analytics-panel analytics-panel--wide">
        <div className="incident-report-section__header">
          <span className="summary-label">Federation trust</span>
          <strong>Peer health, reused proofs, and local trust outcomes</strong>
        </div>
        {view ? (
          <div className="summary-grid">
            <article className="summary-card">
              <span className="summary-label">Trust health</span>
              <strong className="summary-value">{view.trust_health}</strong>
              <p>{view.peers.length} peer(s) · {view.stale_peers?.length || 0} stale.</p>
            </article>
            <article className="summary-card">
              <span className="summary-label">Reused artifacts</span>
              <strong className="summary-value">{view.verified_artifacts_reused}</strong>
              <p>Remote sealed handoffs verified through the local trust policy.</p>
            </article>
            <article className="summary-card">
              <span className="summary-label">Policy sync</span>
              <strong className="summary-value">{view.policy_state.sync_status}</strong>
              <p>{view.policy_state.local_overrides?.length || 0} local override(s) preserved.</p>
            </article>
            <article className="summary-card">
              <span className="summary-label">Anchors</span>
              <strong className="summary-value">{view.anchors.length}</strong>
              <p>{view.policy_divergence?.length || 0} divergence signal(s) recorded.</p>
            </article>
          </div>
        ) : renderEmpty("Federation state is not available for the current scope.")}
      </section>

      <section className="panel analytics-panel analytics-panel--wide">
        <div className="incident-report-section__header">
          <span className="summary-label">Peers</span>
          <strong>Verified trust domains and disclosure boundaries</strong>
        </div>
        {view && view.peers.length > 0 ? (
          <div className="incident-package-table">
            <div className="incident-package-table__row incident-package-table__row--header">
              <span>Peer</span>
              <span>Status</span>
              <span>Role</span>
              <span>Capabilities</span>
              <span>Trust channel</span>
            </div>
            {visiblePeers.map((peer) => (
              <div className={`incident-package-table__row ${focusPeerID === peer.peer_id ? "is-selected" : ""}`} key={peer.peer_id}>
                <span>
                  <strong>{peer.organization}</strong>
                  <small>{peer.peer_id} · {peer.region || peer.trust_domain || "unknown domain"}</small>
                </span>
                <span>{peer.status}</span>
                <span>{peer.policy_role}</span>
                <span>{(peer.capabilities || []).slice(0, 3).join(", ") || "none declared"}</span>
                <span>{peer.trust_state.channel_mode} · last seen {formatTimestamp(peer.last_seen)}</span>
              </div>
            ))}
          </div>
        ) : renderEmpty("No federation peers are currently registered.")}
      </section>

      <section className="panel analytics-panel analytics-panel--wide">
        <div className="incident-report-section__header">
          <span className="summary-label">Proof exchange</span>
          <strong>Recent remote proof decisions</strong>
        </div>
        {view && view.proof_history.length > 0 ? (
          <div className="incident-impact-list">
            {visibleProofHistory.slice(0, 6).map((item) => (
              <article className={`incident-impact-card incident-defense-gap ${focusPeerID === item.peer_id ? "is-selected" : ""}`} key={item.request_id}>
                <div className="incident-impact-card__header">
                  <strong>{item.peer_id}</strong>
                  <span className={`chip chip--${item.decision?.startsWith("accepted") ? "allow" : item.decision ? "deny" : "muted"}`}>
                    {item.decision || item.status}
                  </span>
                </div>
                <p>{item.subject_ref} · {item.proof_type} · {item.manifest_hash.slice(0, 18)}…</p>
                <small>
                  {item.freshness ? `valid until ${formatTimestamp(item.freshness.valid_until)}` : "no freshness window"}
                  {item.verified_at ? ` · verified ${formatTimestamp(item.verified_at)}` : ""}
                </small>
              </article>
            ))}
          </div>
        ) : renderEmpty("No federated proof exchange history has been recorded yet.")}
      </section>

      <section className="panel analytics-panel analytics-panel--wide">
        <div className="incident-report-section__header">
          <span className="summary-label">Policy and anchors</span>
          <strong>Leader/follower sync and regional root integrity</strong>
        </div>
        {view ? (
          <div className="incident-evidence-grid">
            <div>
              <span className="summary-label">Policy state</span>
              <ul className="summary-list summary-list--compact">
                <li><span>Leader: {view.policy_state.leader_peer || "none"}</span></li>
                <li><span>Global root: {view.policy_state.global_policy_root || "n/a"}</span></li>
                <li><span>Effective root: {view.policy_state.effective_policy_root || "n/a"}</span></li>
                {(view.policy_state.divergence_reasons || []).slice(0, 4).map((reason) => (
                  <li key={reason}><span>Divergence: {reason}</span></li>
                ))}
              </ul>
            </div>
            <div>
              <span className="summary-label">Anchor status</span>
              {view.anchors.length > 0 ? (
                <ul className="summary-list summary-list--compact">
                  {view.anchors.slice(0, 5).map((anchor) => (
                    <li key={`${anchor.peer_id}-${anchor.audit_root_hash}`}>
                      <span>{anchor.peer_id} · {anchor.verification_status} · published {formatTimestamp(anchor.published_at)}</span>
                    </li>
                  ))}
                </ul>
              ) : renderEmpty("No federation anchor records are available yet.")}
            </div>
          </div>
        ) : renderEmpty("No federation policy or anchor state is available.")}
      </section>
    </>
  );
}
