import type {
  ForensicReplayResponse,
  ForensicTimelineResponse,
  PointInTimeState,
  TimeDeltaResult,
  VEXFlashbackResponse,
} from "../types";

type Props = {
  state: PointInTimeState | null;
  delta: TimeDeltaResult | null;
  timeline: ForensicTimelineResponse | null;
  flashback: VEXFlashbackResponse | null;
  replay: ForensicReplayResponse | null;
  loading: boolean;
  timestamp: string;
  onTimestampChange: (value: string) => void;
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

export function ForensicsInsightsPanel({ state, delta, timeline, flashback, replay, loading, timestamp, onTimestampChange }: Props) {
  if (loading) {
    return <section className="panel analytics-panel analytics-panel--wide">Loading forensic reconstruction…</section>;
  }

  return (
    <>
      <section className="panel analytics-panel analytics-panel--wide">
        <div className="incident-report-section__header">
          <div>
            <span className="summary-label">Time-travel forensics</span>
            <strong>Historical reconstruction, delta, and counterfactual replay</strong>
          </div>
          <label className="filters-inline-field">
            <span className="summary-label">Historical timestamp</span>
            <input type="datetime-local" value={timestamp} onChange={(event) => onTimestampChange(event.target.value)} />
          </label>
        </div>
        {state ? (
          <div className="summary-grid">
            <article className="summary-card">
              <span className="summary-label">Historical state</span>
              <strong className="summary-value">{formatTimestamp(state.timestamp)}</strong>
              <p>{state.subject_summary || "current-scope"} · {state.environment || "unknown environment"}</p>
            </article>
            <article className="summary-card">
              <span className="summary-label">Policy bundle</span>
              <strong className="summary-value">{state.policy_context.policy_bundle_hash || "n/a"}</strong>
              <p>{state.policy_context.rule_versions.length} rule version(s) · {state.policy_context.active_rules.length} active rule hint(s)</p>
            </article>
            <article className="summary-card">
              <span className="summary-label">Known findings at T</span>
              <strong className="summary-value">{state.vulnerability_context.known_findings.length}</strong>
              <p>{state.vulnerability_context.unknown_later_disclosed_refs.length} later disclosure ref(s) stay separate from historical truth.</p>
            </article>
            <article className="summary-card">
              <span className="summary-label">Topology at T</span>
              <strong className="summary-value">{state.topology_context?.blast_radius_score ?? 0}</strong>
              <p>{state.topology_context?.critical_reach_count ?? 0} critical reach · {state.topology_context?.primary_service || "no mapped service"}</p>
            </article>
          </div>
        ) : renderEmpty("Forensic point-in-time state is not available for the current scope.")}
      </section>

      <section className="panel analytics-panel analytics-panel--wide">
        <div className="incident-report-section__header">
          <span className="summary-label">Time delta</span>
          <strong>What changed between T1 and T2</strong>
        </div>
        {delta ? (
          <div className="incident-evidence-grid">
            <div>
              <span className="summary-label">Comparison window</span>
              <p>{formatTimestamp(delta.comparison.t1)} → {formatTimestamp(delta.comparison.t2)}</p>
              <ul className="summary-list summary-list--compact">
                {(delta.policy_delta.modified || []).slice(0, 4).map((item) => (
                  <li key={`policy-${item}`}><span>Policy drift: {item}</span></li>
                ))}
                {(delta.identity_delta.modified || []).slice(0, 4).map((item) => (
                  <li key={`identity-${item}`}><span>Identity drift: {item}</span></li>
                ))}
                {(delta.vulnerability_delta.added || []).slice(0, 4).map((item) => (
                  <li key={`vuln-${item}`}><span>Later disclosure: {item}</span></li>
                ))}
              </ul>
            </div>
            <div>
              <span className="summary-label">Topology delta</span>
              {delta.topology_delta && delta.topology_delta.length > 0 ? (
                <ul className="summary-list summary-list--compact">
                  {delta.topology_delta.slice(0, 4).map((item) => (
                    <li key={`${item.node_id}-${item.service}`}>
                      <span>{item.service} · delta {item.delta > 0 ? "+" : ""}{item.delta} · critical reach {item.critical_reach_delta > 0 ? "+" : ""}{item.critical_reach_delta}</span>
                    </li>
                  ))}
                </ul>
              ) : renderEmpty("No topology delta items matched the current forensic comparison.")}
            </div>
          </div>
        ) : renderEmpty("Forensic delta is not available for the current scope.")}
      </section>

      <section className="panel analytics-panel analytics-panel--wide">
        <div className="incident-report-section__header">
          <span className="summary-label">Replay</span>
          <strong>Historical verdict vs modern control stack</strong>
        </div>
        {replay ? (
          <div className="summary-grid">
            <article className="summary-card">
              <span className="summary-label">Historical verdict</span>
              <strong className="summary-value">{replay.historical_verdict}</strong>
              <p>{formatTimestamp(replay.historical_timestamp)}</p>
            </article>
            <article className="summary-card">
              <span className="summary-label">Replay verdict</span>
              <strong className="summary-value">{replay.replay_verdict}</strong>
              <p>{replay.replay_mode} · {replay.counterfactual ? "counterfactual" : "historical"}</p>
            </article>
            <article className="summary-card">
              <span className="summary-label">Verdict delta</span>
              <strong className="summary-value">{replay.verdict_delta}</strong>
              <p>{(replay.explanations || [])[0] || "No replay explanation recorded."}</p>
            </article>
          </div>
        ) : renderEmpty("Replay result is not available for the current scope.")}
      </section>

      <section className="panel analytics-panel analytics-panel--wide">
        <div className="incident-report-section__header">
          <span className="summary-label">VEX flashback</span>
          <strong>Historical vuln known-state and active VEX basis</strong>
        </div>
        {flashback ? (
          <div className="incident-evidence-grid">
            <div>
              <span className="summary-label">Historical basis</span>
              <p>{flashback.historical_decision_basis}</p>
              <ul className="summary-list summary-list--compact">
                {flashback.historical_vulnerability_state.slice(0, 4).map((item) => (
                  <li key={`${item.cve_id}-${item.image_digest || "scope"}`}>
                    <span>{item.cve_id} · {item.known_at_t ? "known at T" : "later"} · {item.severity || "unknown severity"}</span>
                  </li>
                ))}
              </ul>
            </div>
            <div>
              <span className="summary-label">Active VEX / later disclosures</span>
              <ul className="summary-list summary-list--compact">
                {flashback.vex_flashback.slice(0, 3).map((item) => (
                  <li key={`vex-${item.statement_id}`}>
                    <span>{item.vulnerability_id} · {item.status}</span>
                  </li>
                ))}
                {(flashback.disclosed_after_t_refs || []).slice(0, 3).map((item) => (
                  <li key={`later-${item}`}>
                    <span>Disclosed after T: {item}</span>
                  </li>
                ))}
              </ul>
            </div>
          </div>
        ) : renderEmpty("VEX flashback is not available for the current scope.")}
      </section>

      <section className="panel analytics-panel analytics-panel--wide">
        <div className="incident-report-section__header">
          <span className="summary-label">Timeline</span>
          <strong>Evidence-backed markers across the selected window</strong>
        </div>
        {timeline && timeline.markers.length > 0 ? (
          <ul className="summary-list">
            {timeline.markers.slice(0, 8).map((marker) => (
              <li key={marker.marker_id}>
                <span>{formatTimestamp(marker.timestamp)} · {marker.marker_type} · {marker.title}</span>
              </li>
            ))}
          </ul>
        ) : renderEmpty("No forensic markers matched the selected window.")}
      </section>
    </>
  );
}
