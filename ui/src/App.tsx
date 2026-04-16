import { useEffect, useState } from "react";

import {
  APIError,
  apiBaseURL,
  apiTokenConfigured,
  approveException,
  getAuthStatus,
  getDriftStats,
  getEvents,
  getExceptionReport,
  getExceptions,
  getHealth,
  getSummary,
  getTopViolators,
  getTrends,
  rejectException,
  requestException,
  revokeException,
} from "./api";
import { AnalyticsTrendsPanel } from "./components/AnalyticsTrendsPanel";
import { BuyerHighlights } from "./components/BuyerHighlights";
import { DriftStatsPanel } from "./components/DriftStatsPanel";
import { EventDetails } from "./components/EventDetails";
import { EventsTable } from "./components/EventsTable";
import { ExceptionRequestForm } from "./components/ExceptionRequestForm";
import { Filters } from "./components/Filters";
import { HealthBadge } from "./components/HealthBadge";
import { PendingExceptionsPanel } from "./components/PendingExceptionsPanel";
import { SBOMInventoryPanel } from "./components/SBOMInventoryPanel";
import { SummaryCards } from "./components/SummaryCards";
import { TopViolatorsPanel } from "./components/TopViolatorsPanel";
import { VulnerabilityOpsPanel } from "./components/VulnerabilityOpsPanel";
import type {
  AuditHealth,
  AuthStatus,
  DriftStatsResponse,
  EventFilters,
  ExceptionReport,
  ExceptionRequestInput,
  PolicyException,
  StoredEvent,
  Summary,
  TabKey,
  TopViolatorsResponse,
  TrendsResponse,
} from "./types";

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
  { key: "analytics", label: "Analytics", description: "Trends, violators, and drift statistics." },
  { key: "exceptions", label: "Exceptions", description: "Approval queue, status counts, and recent exception use." },
  { key: "inventory", label: "SBOM Inventory", description: "Search stored SBOM components by digest, package, or PURL." },
  { key: "vulnerabilities", label: "Vulnerability Ops", description: "Active findings, blast radius, timelines, and VEX-lite decisions." },
];

function isHumanRole(role?: string) {
  return role === "viewer" || role === "operator" || role === "security_admin";
}

export default function App() {
  const [activeTab, setActiveTab] = useState<TabKey>("overview");
  const [filters, setFilters] = useState<EventFilters>(initialFilters);
  const [summary, setSummary] = useState<Summary | null>(null);
  const [events, setEvents] = useState<StoredEvent[]>([]);
  const [selectedEvent, setSelectedEvent] = useState<StoredEvent | null>(null);
  const [health, setHealth] = useState<AuditHealth | null>(null);
  const [authStatus, setAuthStatus] = useState<AuthStatus | null>(null);
  const [trends, setTrends] = useState<TrendsResponse | null>(null);
  const [topViolators, setTopViolators] = useState<TopViolatorsResponse | null>(null);
  const [driftStats, setDriftStats] = useState<DriftStatsResponse | null>(null);
  const [exceptionReport, setExceptionReport] = useState<ExceptionReport | null>(null);
  const [pendingExceptions, setPendingExceptions] = useState<PolicyException[]>([]);
  const [loading, setLoading] = useState(true);
  const [requestSubmitting, setRequestSubmitting] = useState(false);
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
          setTrends(null);
          setTopViolators(null);
          setDriftStats(null);
          setExceptionReport(null);
          setPendingExceptions([]);

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
          const baseFilters = {
            tenant_id: filters.tenant_id,
            environment: filters.environment,
            repo: filters.repo,
          };

          const promises: Array<Promise<void>> = [
            getSummary({ environment: filters.environment, tenant_id: filters.tenant_id }).then(setSummary),
          ];

          if (activeTab === "analytics") {
            promises.push(
              getTrends({
                window_days: "30",
                granularity: "day",
                ...baseFilters,
              }).then(setTrends),
            );
            promises.push(
              getTopViolators({
                window_days: "30",
                limit: "5",
                dimension: "repo",
                ...baseFilters,
              }).then(setTopViolators),
            );
            promises.push(
              getDriftStats({
                window_days: "30",
                ...baseFilters,
              }).then(setDriftStats),
            );
            setEvents([]);
            setSelectedEvent(null);
          } else if (activeTab === "exceptions") {
            promises.push(getExceptionReport(baseFilters).then((report) => {
              setExceptionReport(report);
              setEvents(report.recent_used);
              setSelectedEvent((current) => report.recent_used.find((event) => event.id === current?.id) || report.recent_used[0] || null);
            }));
            promises.push(
              getExceptions({
                ...baseFilters,
                status: "PENDING",
                limit: filters.limit,
              }).then((response) => setPendingExceptions(response.exceptions)),
            );
          } else if (activeTab === "inventory" || activeTab === "vulnerabilities") {
            setEvents([]);
            setSelectedEvent(null);
            setTrends(null);
            setTopViolators(null);
            setDriftStats(null);
            setExceptionReport(null);
            setPendingExceptions([]);
          } else {
            promises.push(
              getEvents(activeTab, filters).then((response) => {
                setEvents(response.events);
                setSelectedEvent((current) => response.events.find((event) => event.id === current?.id) || response.events[0] || null);
              }),
            );
            setTrends(null);
            setTopViolators(null);
            setDriftStats(null);
            setExceptionReport(null);
            setPendingExceptions([]);
          }

          const settled = await Promise.allSettled(promises);
          if (ignore) {
            return;
          }
          for (const result of settled) {
            if (result.status === "rejected") {
              problems.push(result.reason instanceof Error ? result.reason.message : "Dashboard data unavailable.");
            }
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
  const role = authStatus?.role;
  const canRequest = role === "operator" || role === "security_admin";
  const canApprove = role === "security_admin";

  async function handleRequestException(input: ExceptionRequestInput) {
    setRequestSubmitting(true);
    try {
      await requestException(input);
      setRefreshIndex((value) => value + 1);
    } finally {
      setRequestSubmitting(false);
    }
  }

  async function handleApprove(exceptionID: string) {
    await approveException(exceptionID);
    setRefreshIndex((value) => value + 1);
  }

  async function handleReject(exceptionID: string, reason: string) {
    await rejectException(exceptionID, reason);
    setRefreshIndex((value) => value + 1);
  }

  async function handleRevoke(exceptionID: string) {
    await revokeException(exceptionID);
    setRefreshIndex((value) => value + 1);
  }

  return (
    <main className="app-shell">
      <header className="hero">
        <div>
          <p className="eyebrow">ChangeLock Security Dashboard</p>
          <h1>Audit Visibility, Approval Governance, and Operational Trends</h1>
          <p className="hero-copy">
            This dashboard renders the live PostgreSQL-backed reports, analytics, and exception governance APIs from the
            ChangeLock control plane.
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

      {!isHumanRole(role) && authStatus?.authenticated ? (
        <section className="panel status-banner">
          <div>
            <strong>Current token is not a human dashboard role.</strong>
            <p>Use a viewer, operator, or security_admin token for the UI. service_internal is reserved for backend validation calls.</p>
          </div>
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

      {activeTab !== "inventory" && activeTab !== "vulnerabilities" ? (
        <Filters
          filters={filters}
          tab={activeTab}
          onChange={(name, value) => setFilters((current) => ({ ...current, [name]: value }))}
          onRefresh={() => setRefreshIndex((value) => value + 1)}
          onReset={() => setFilters(initialFilters)}
        />
      ) : null}

      {activeTab === "analytics" ? (
        <section className="analytics-grid">
          <AnalyticsTrendsPanel trends={trends} loading={loading} />
          <TopViolatorsPanel data={topViolators} loading={loading} />
          <DriftStatsPanel data={driftStats} loading={loading} />
        </section>
      ) : null}

      {activeTab === "exceptions" ? (
        <>
          <section className="summary-grid">
            {Object.entries(exceptionReport?.status_counts || {}).map(([status, count]) => (
              <article className="summary-card" key={status}>
                <span className="summary-label">{status}</span>
                <strong className="summary-value">{count}</strong>
              </article>
            ))}
          </section>

          <section className="analytics-grid">
            <ExceptionRequestForm enabled={canRequest} submitting={requestSubmitting} onSubmit={handleRequestException} />
            <PendingExceptionsPanel
              pending={pendingExceptions}
              canApprove={canApprove}
              canRevoke={canApprove}
              loading={loading}
              onApprove={handleApprove}
              onReject={handleReject}
              onRevoke={handleRevoke}
            />
          </section>

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
        </>
      ) : null}

      {activeTab === "inventory" ? <SBOMInventoryPanel /> : null}

      {activeTab === "vulnerabilities" ? <VulnerabilityOpsPanel role={role} /> : null}

      {activeTab !== "analytics" && activeTab !== "exceptions" && activeTab !== "inventory" && activeTab !== "vulnerabilities" ? (
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
      ) : null}
    </main>
  );
}
