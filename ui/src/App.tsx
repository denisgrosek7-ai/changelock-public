import { useEffect, useState } from "react";

import { APIError, apiBaseURL, apiTokenConfigured, getAuthStatus, getEvents, getHealth, getSummary } from "./api";
import { BuyerHighlights } from "./components/BuyerHighlights";
import { EventDetails } from "./components/EventDetails";
import { EventsTable } from "./components/EventsTable";
import { Filters } from "./components/Filters";
import { HealthBadge } from "./components/HealthBadge";
import { SummaryCards } from "./components/SummaryCards";
import type { AuditHealth, AuthStatus, EventFilters, StoredEvent, Summary, TabKey } from "./types";

const initialFilters: EventFilters = {
  decision: "",
  component: "",
  repo: "",
  environment: "",
  tenant_id: "",
  limit: "25",
};

const tabs: Array<{ key: TabKey; label: string; description: string }> = [
  { key: "overview", label: "Overview", description: "Summary plus latest audit events." },
  { key: "events", label: "Events", description: "All audit events with operator filters." },
  { key: "denies", label: "Denies", description: "Rejected operations and why they were blocked." },
  { key: "runtime", label: "Runtime Drift", description: "Runtime scans and drift findings." },
];

export default function App() {
  const [activeTab, setActiveTab] = useState<TabKey>("overview");
  const [filters, setFilters] = useState<EventFilters>(initialFilters);
  const [summary, setSummary] = useState<Summary | null>(null);
  const [events, setEvents] = useState<StoredEvent[]>([]);
  const [selectedEvent, setSelectedEvent] = useState<StoredEvent | null>(null);
  const [health, setHealth] = useState<AuditHealth | null>(null);
  const [authStatus, setAuthStatus] = useState<AuthStatus | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [refreshIndex, setRefreshIndex] = useState(0);
  const [lastLoadedAt, setLastLoadedAt] = useState<string>("");

  useEffect(() => {
    let ignore = false;

    async function load() {
      setLoading(true);
      setError(null);

      try {
        const [healthResult, authResult] = await Promise.allSettled([getHealth(), getAuthStatus()]);

        if (ignore) {
          return;
        }

        const problems: string[] = [];

        if (healthResult.status === "fulfilled") {
          setHealth(healthResult.value);
        } else {
          setHealth({
            status: "error",
            error: healthResult.reason instanceof Error ? healthResult.reason.message : "Health unavailable.",
          });
          problems.push("Backend health is unavailable.");
        }

        if (authResult.status === "fulfilled") {
          setAuthStatus(authResult.value);
        } else {
          setSummary(null);
          setEvents([]);
          setSelectedEvent(null);

          if (authResult.reason instanceof APIError) {
            if (authResult.reason.status === 401) {
              problems.push("Dashboard API requires a valid bearer token.");
            } else if (authResult.reason.status === 403) {
              problems.push("Configured token is not allowed to access dashboard reports.");
            } else {
              problems.push(authResult.reason.message);
            }
          } else {
            problems.push(authResult.reason instanceof Error ? authResult.reason.message : "Auth status unavailable.");
          }
        }

        if (authResult.status === "fulfilled") {
          const [summaryResult, eventsResult] = await Promise.allSettled([
            getSummary({ environment: filters.environment, tenant_id: filters.tenant_id }),
            getEvents(activeTab, filters),
          ]);

          if (ignore) {
            return;
          }

          if (summaryResult.status === "fulfilled") {
            setSummary(summaryResult.value);
          } else {
            problems.push(summaryResult.reason instanceof Error ? summaryResult.reason.message : "Summary unavailable.");
          }

          if (eventsResult.status === "fulfilled") {
            setEvents(eventsResult.value.events);
            setSelectedEvent((current) =>
              eventsResult.value.events.find((event) => event.id === current?.id) || eventsResult.value.events[0] || null,
            );
          } else {
            problems.push(eventsResult.reason instanceof Error ? eventsResult.reason.message : "Events unavailable.");
          }
        }

        if (problems.length === 0) {
          setLastLoadedAt(new Date().toLocaleTimeString());
        } else {
          setError(problems.join(" "));
        }
      } catch (loadError) {
        if (ignore) {
          return;
        }
        setError(loadError instanceof Error ? loadError.message : "Failed to load dashboard data.");
      } finally {
        if (!ignore) {
          setLoading(false);
        }
      }
    }

    void load();
    return () => {
      ignore = true;
    };
  }, [activeTab, filters.component, filters.decision, filters.environment, filters.limit, filters.repo, filters.tenant_id, refreshIndex]);

  const activeTabMeta = tabs.find((tab) => tab.key === activeTab) || tabs[0];

  return (
    <main className="app-shell">
      <header className="hero">
        <div>
          <p className="eyebrow">ChangeLock Security Dashboard</p>
          <h1>Audit Visibility for Trust, Denies, and Runtime Drift</h1>
          <p className="hero-copy">
            This dashboard renders the live PostgreSQL-backed reports API from the ChangeLock control plane and
            keeps deployment, verification, and runtime findings in one review surface.
          </p>
        </div>

        <div className="hero-meta">
          <div className="panel health-panel">
            <span className="summary-label">Audit Backend</span>
            <strong>{health?.backend || "unknown"}</strong>
            <HealthBadge health={health} />
          </div>
          <div className="panel health-panel">
            <span className="summary-label">API Base</span>
            <code>{apiBaseURL()}</code>
          </div>
          <div className="panel health-panel">
            <span className="summary-label">Access</span>
            <strong>{authStatus?.role || authStatus?.auth_mode || "unknown"}</strong>
            <small>
              {authStatus?.subject
                ? `${authStatus.subject}${authStatus.token_id ? ` (${authStatus.token_id})` : ""}`
                : apiTokenConfigured()
                  ? "Token configured"
                  : "No token configured"}
            </small>
          </div>
        </div>
      </header>

      {error ? (
        <section className="panel status-banner">
          <div>
            <strong>Dashboard data is partially unavailable.</strong>
            <p>{error}</p>
          </div>
          <HealthBadge health={health} />
        </section>
      ) : null}

      <BuyerHighlights summary={summary} loading={loading} />

      <SummaryCards summary={summary} loading={loading} />

      <section className="tabs">
        {tabs.map((tab) => (
          <button
            key={tab.key}
            className={`tab ${tab.key === activeTab ? "is-active" : ""}`}
            onClick={() => setActiveTab(tab.key)}
          >
            <span>{tab.label}</span>
          </button>
        ))}
      </section>

      <section className="tab-header panel">
        <div>
          <h2>{activeTabMeta.label}</h2>
          <p>{activeTabMeta.description}</p>
          {lastLoadedAt ? <p className="tab-header__meta">Last refresh at {lastLoadedAt}</p> : null}
        </div>
        <button className="button" onClick={() => setRefreshIndex((value) => value + 1)}>
          Refresh Data
        </button>
      </section>

      <Filters
        filters={filters}
        tab={activeTab}
        onChange={(name, value) => setFilters((current) => ({ ...current, [name]: value }))}
        onRefresh={() => setRefreshIndex((value) => value + 1)}
        onReset={() => setFilters(initialFilters)}
      />

      <section className="content-grid">
        <EventsTable
          events={events}
          selectedEventID={selectedEvent?.id || null}
          onSelect={setSelectedEvent}
          loading={loading}
          error={error}
        />
        <EventDetails event={selectedEvent} />
      </section>
    </main>
  );
}
