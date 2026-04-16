import type { EventFilters, TabKey } from "../types";

type Props = {
  filters: EventFilters;
  tab: TabKey;
  enforcedTenantID?: string;
  onChange: (name: keyof EventFilters, value: string) => void;
  onRefresh: () => void;
  onReset: () => void;
};

export function Filters({ filters, tab, enforcedTenantID, onChange, onRefresh, onReset }: Props) {
  const isEventTab = tab === "overview" || tab === "events" || tab === "denies" || tab === "runtime";

  return (
    <section className="panel filters-panel">
      <div className="filters-grid">
        {tab === "events" ? (
          <label>
            <span>Decision</span>
            <select value={filters.decision} onChange={(event) => onChange("decision", event.target.value)}>
              <option value="">All</option>
              <option value="ALLOW">ALLOW</option>
              <option value="DENY">DENY</option>
              <option value="ERROR">ERROR</option>
            </select>
          </label>
        ) : null}

        {isEventTab ? (
          <label>
            <span>Component</span>
            <input
              value={filters.component}
              onChange={(event) => onChange("component", event.target.value)}
              placeholder="deploy-gate"
            />
          </label>
        ) : null}

        <label>
          <span>Repo</span>
          <input
            value={filters.repo}
            onChange={(event) => onChange("repo", event.target.value)}
            placeholder="my-org/acme-app"
          />
        </label>

        <label>
          <span>Environment</span>
          <input
            value={filters.environment}
            onChange={(event) => onChange("environment", event.target.value)}
            placeholder="prod"
          />
        </label>

        <label>
          <span>Tenant</span>
          <input
            value={enforcedTenantID || filters.tenant_id}
            onChange={(event) => onChange("tenant_id", event.target.value)}
            placeholder="acme"
            disabled={Boolean(enforcedTenantID)}
          />
        </label>

        <label>
          <span>Limit</span>
          <input value={filters.limit} onChange={(event) => onChange("limit", event.target.value)} placeholder="25" />
        </label>
      </div>

      <div className="filters-actions">
        <button className="button" onClick={onReset}>
          Reset
        </button>
        <button className="button button--primary" onClick={onRefresh}>
          Refresh
        </button>
      </div>
    </section>
  );
}
