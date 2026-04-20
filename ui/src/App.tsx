import { useEffect, useState } from "react";

import {
  APIError,
  acknowledgeIncident,
  addIncidentNote,
  apiBaseURL,
  apiTokenConfigured,
  getAnalyticsAnomalies,
  getAnalyticsDelta,
  getAnalyticsScorecards,
  getAnalyticsSegments,
  approveException,
  assignIncident,
  getAuthStatus,
  getDriftStats,
  getExecutiveDefenseReport,
  getEvents,
  getExceptionReport,
  getExceptions,
  getHealth,
  getIncidentDefenseGaps,
  getIncidentBlastRadius,
  getIncidentExport,
  getIncidentPolicyReplay,
  getIncidentPackage,
  getIncidents,
  getMetricDefenseGaps,
  getMetricBlastRadius,
  getMetricIncidents,
  getMetricPolicyReplay,
  getRuntimeActiveStates,
  getRuntimeClosedLoopStatus,
  getRuntimeDriftFindings,
  getRuntimeDriftStatus,
  getSummary,
  getSystemicWeaknesses,
  getSyncStatus,
  getTopologyDelta,
  getTopologyGraph,
  getTopologyHeatmap,
  getTopViolators,
  getTrends,
  rejectException,
  reopenIncident,
  resolveIncident,
  requestException,
  revokeException,
  watchIncident,
} from "./api";
import { AIInsightsPanel } from "./components/AIInsightsPanel";
import { AnalyticsInsightsPanel } from "./components/AnalyticsInsightsPanel";
import { AnalyticsTrendsPanel } from "./components/AnalyticsTrendsPanel";
import { DriftStatsPanel } from "./components/DriftStatsPanel";
import { EventDetails } from "./components/EventDetails";
import { EventsTable } from "./components/EventsTable";
import { ExceptionRequestForm } from "./components/ExceptionRequestForm";
import { Filters } from "./components/Filters";
import { HealthBadge } from "./components/HealthBadge";
import { IncidentWorkbench } from "./components/IncidentWorkbench";
import { OverviewDashboard } from "./components/OverviewDashboard";
import { PendingExceptionsPanel } from "./components/PendingExceptionsPanel";
import { RuntimeDriftPanel } from "./components/RuntimeDriftPanel";
import { SBOMInventoryPanel } from "./components/SBOMInventoryPanel";
import { SigningIdentityPanel } from "./components/SigningIdentityPanel";
import { TopologyInsightsPanel } from "./components/TopologyInsightsPanel";
import { TopViolatorsPanel } from "./components/TopViolatorsPanel";
import { TrustScorecardPanel } from "./components/TrustScorecardPanel";
import { VulnerabilityOpsPanel } from "./components/VulnerabilityOpsPanel";
import type {
  DefenseGapAssessment,
  ExecutiveDefenseReport,
  IncidentExport,
  IncidentPackage,
  PolicyReplayAssessment,
  IncidentReportAudience,
  InvestigationIncident,
  MetricIncidentDrilldown,
  SystemicWeaknessResponse,
} from "./incidents";
import type {
  AuditHealth,
  AuthStatus,
  AnalyticsAnomaliesResponse,
  AnalyticsDeltaResponse,
  AnalyticsScorecardsResponse,
  AnalyticsSegmentsResponse,
  DriftStatsResponse,
  EventFilters,
  ExceptionReport,
  ExceptionRequestInput,
  PolicyException,
  RuntimeActiveState,
  RuntimeClosedLoopStatus,
  RuntimeDriftFinding,
  RuntimeDriftStatus,
  StoredEvent,
  Summary,
  SyncStatus,
  TabKey,
  TopViolatorsResponse,
  TopologyBlastRadiusResponse,
  TopologyDeltaResponse,
  TopologyGraphResponse,
  TopologyHeatmapResponse,
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
  { key: "overview", label: "Posture", description: "Current posture, active incidents, and next operator actions." },
  { key: "events", label: "Investigations", description: "Evidence-backed event stream for drill-down and triage." },
  { key: "denies", label: "Denials", description: "Rejected operations and the trust reasons that blocked them." },
  { key: "runtime", label: "Runtime", description: "Drift findings and the workloads that need reconciliation." },
  { key: "analytics", label: "Governance", description: "Trends, violators, and control-plane operating pressure." },
  { key: "topology", label: "Topology", description: "Service-graph blast radius, drift, and containment guidance." },
  { key: "exceptions", label: "Exceptions", description: "Approval queue, status counts, and recent exception use." },
  { key: "inventory", label: "Components", description: "Investigate stored SBOM components by digest, package, or PURL." },
  { key: "vulnerabilities", label: "Vulnerabilities", description: "Active findings, blast radius, timelines, and VEX-lite decisions." },
  { key: "signing", label: "Signing Identities", description: "Authorized signers, observed identities, and workflow drift findings." },
  { key: "scorecard", label: "Trust Scorecard", description: "Measured hardening grade, audit findings, trust badges, and auditor-ready posture." },
  { key: "guidance", label: "AI Guidance", description: "Contextual grouping, bounded remediation guidance, VEX draft candidates, and break-glass reminders." },
];

function isHumanRole(role?: string) {
  return role === "viewer" || role === "operator" || role === "security_admin";
}

function isOptionalFeatureMissing(error: unknown) {
  return error instanceof APIError && error.status === 404;
}

export default function App() {
  const [activeTab, setActiveTab] = useState<TabKey>("overview");
  const [filters, setFilters] = useState<EventFilters>(initialFilters);
  const [summary, setSummary] = useState<Summary | null>(null);
  const [events, setEvents] = useState<StoredEvent[]>([]);
  const [selectedEvent, setSelectedEvent] = useState<StoredEvent | null>(null);
  const [health, setHealth] = useState<AuditHealth | null>(null);
  const [authStatus, setAuthStatus] = useState<AuthStatus | null>(null);
  const [syncStatus, setSyncStatus] = useState<SyncStatus | null>(null);
  const [trends, setTrends] = useState<TrendsResponse | null>(null);
  const [topViolators, setTopViolators] = useState<TopViolatorsResponse | null>(null);
  const [driftStats, setDriftStats] = useState<DriftStatsResponse | null>(null);
  const [analyticsDelta, setAnalyticsDelta] = useState<AnalyticsDeltaResponse | null>(null);
  const [analyticsAnomalies, setAnalyticsAnomalies] = useState<AnalyticsAnomaliesResponse | null>(null);
  const [analyticsScorecards, setAnalyticsScorecards] = useState<AnalyticsScorecardsResponse | null>(null);
  const [analyticsSegments, setAnalyticsSegments] = useState<AnalyticsSegmentsResponse | null>(null);
  const [topologyGraph, setTopologyGraph] = useState<TopologyGraphResponse | null>(null);
  const [topologyHeatmap, setTopologyHeatmap] = useState<TopologyHeatmapResponse | null>(null);
  const [topologyDelta, setTopologyDelta] = useState<TopologyDeltaResponse | null>(null);
  const [runtimeActiveStates, setRuntimeActiveStates] = useState<RuntimeActiveState[]>([]);
  const [runtimeClosedLoopStatus, setRuntimeClosedLoopStatus] = useState<RuntimeClosedLoopStatus | null>(null);
  const [runtimeDriftFindings, setRuntimeDriftFindings] = useState<RuntimeDriftFinding[]>([]);
  const [runtimeDriftStatus, setRuntimeDriftStatus] = useState<RuntimeDriftStatus | null>(null);
  const [exceptionReport, setExceptionReport] = useState<ExceptionReport | null>(null);
  const [incidents, setIncidents] = useState<InvestigationIncident[]>([]);
  const [systemicWeaknesses, setSystemicWeaknesses] = useState<SystemicWeaknessResponse | null>(null);
  const [executiveReport, setExecutiveReport] = useState<ExecutiveDefenseReport | null>(null);
  const [pendingExceptions, setPendingExceptions] = useState<PolicyException[]>([]);
  const [metricDrilldown, setMetricDrilldown] = useState<MetricIncidentDrilldown | null>(null);
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
          setSyncStatus(null);
          setSummary(null);
          setEvents([]);
          setSelectedEvent(null);
          setTrends(null);
          setTopViolators(null);
          setDriftStats(null);
          setAnalyticsDelta(null);
          setAnalyticsAnomalies(null);
          setAnalyticsScorecards(null);
          setAnalyticsSegments(null);
          setTopologyGraph(null);
          setTopologyHeatmap(null);
          setTopologyDelta(null);
          setRuntimeActiveStates([]);
          setRuntimeClosedLoopStatus(null);
          setRuntimeDriftFindings([]);
          setRuntimeDriftStatus(null);
          setExceptionReport(null);
          setIncidents([]);
          setExecutiveReport(null);
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
          const scopedTenantID = authResult.value.tenant_id || filters.tenant_id;
          const baseFilters = {
            tenant_id: scopedTenantID,
            environment: filters.environment,
            repo: filters.repo,
          };

          const promises: Array<Promise<void>> = [
            getSummary({ environment: filters.environment, tenant_id: scopedTenantID }).then(setSummary),
          ];
          if (isHumanRole(authResult.value.role)) {
            promises.push(
              getSyncStatus()
                .then(setSyncStatus)
                .catch((syncError) => {
                  if (isOptionalFeatureMissing(syncError)) {
                    setSyncStatus({
                      mode: "disabled",
                      sync_mode: "disabled",
                      health: "disabled",
                      cache_present: false,
                      summary: "Sync status endpoint is not enabled on this backend.",
                    });
                    return;
                  }
                  throw syncError;
                }),
            );
          } else {
            setSyncStatus(null);
          }

          if (activeTab === "overview") {
            promises.push(
              getEvents(activeTab, { ...filters, tenant_id: scopedTenantID }).then((response) => {
                setEvents(response.events);
                setSelectedEvent((current) => response.events.find((event) => event.id === current?.id) || response.events[0] || null);
              }),
            );
            promises.push(
              getTrends({
                window_days: "7",
                granularity: "day",
                ...baseFilters,
              }).then(setTrends).catch((trendsError) => {
                if (isOptionalFeatureMissing(trendsError)) {
                  setTrends(null);
                  return;
                }
                throw trendsError;
              }),
            );
            promises.push(
              getTopViolators({
                window_days: "7",
                limit: "5",
                dimension: "repo",
                ...baseFilters,
              }).then(setTopViolators).catch((violatorsError) => {
                if (isOptionalFeatureMissing(violatorsError)) {
                  setTopViolators(null);
                  return;
                }
                throw violatorsError;
              }),
            );
            promises.push(
              getDriftStats({
                window_days: "7",
                ...baseFilters,
              }).then(setDriftStats).catch((driftError) => {
                if (isOptionalFeatureMissing(driftError)) {
                  setDriftStats(null);
                  return;
                }
                throw driftError;
              }),
            );
            promises.push(
              getExceptionReport(baseFilters).then(setExceptionReport).catch((reportError) => {
                if (isOptionalFeatureMissing(reportError)) {
                  setExceptionReport(null);
                  return;
                }
                throw reportError;
              }),
            );
            promises.push(
              getSystemicWeaknesses({
                ...filters,
                tenant_id: scopedTenantID,
              }).then(setSystemicWeaknesses).catch((systemicError) => {
                if (isOptionalFeatureMissing(systemicError)) {
                  setSystemicWeaknesses(null);
                  return;
                }
                throw systemicError;
              }),
            );
            promises.push(
              getExecutiveDefenseReport({
                environment: filters.environment,
                tenant_id: scopedTenantID,
                repo: filters.repo,
              }, [], "internal").then(setExecutiveReport).catch((executiveError) => {
                if (isOptionalFeatureMissing(executiveError)) {
                  setExecutiveReport(null);
                  return;
                }
                throw executiveError;
              }),
            );
            setPendingExceptions([]);
            setIncidents([]);
          } else if (activeTab === "events") {
            promises.push(
              (metricDrilldown
                ? getMetricIncidents(metricDrilldown.metricKey, { ...filters, tenant_id: scopedTenantID }).then((response) => {
                    setMetricDrilldown(response);
                    setIncidents(response.incidents);
                    setEvents([]);
                    setSelectedEvent(null);
                  })
                : getIncidents({ ...filters, tenant_id: scopedTenantID }).then((response) => {
                    setIncidents(response);
                    setEvents([]);
                    setSelectedEvent(null);
                  })).catch((incidentError) => {
                if (isOptionalFeatureMissing(incidentError)) {
                  return getEvents(activeTab, { ...filters, tenant_id: scopedTenantID }).then((response) => {
                    setIncidents([]);
                    setEvents(response.events);
                    setSelectedEvent((current) => response.events.find((event) => event.id === current?.id) || response.events[0] || null);
                  });
                }
                throw incidentError;
              }),
            );
            setTrends(null);
            setTopViolators(null);
            setDriftStats(null);
            setExceptionReport(null);
            setSystemicWeaknesses(null);
            setExecutiveReport(null);
            setAnalyticsDelta(null);
            setAnalyticsAnomalies(null);
            setAnalyticsScorecards(null);
            setAnalyticsSegments(null);
            setTopologyGraph(null);
            setTopologyHeatmap(null);
            setTopologyDelta(null);
            setPendingExceptions([]);
          } else if (activeTab === "topology") {
            promises.push(
              getTopologyGraph({
                window: "28d",
                compare_to: "previous_window",
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                limit: filters.limit,
              }).then(setTopologyGraph),
            );
            promises.push(
              getTopologyHeatmap({
                window: "28d",
                compare_to: "previous_window",
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                limit: filters.limit,
              }).then(setTopologyHeatmap),
            );
            promises.push(
              getTopologyDelta({
                window: "28d",
                compare_to: "previous_window",
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                limit: filters.limit,
              }).then(setTopologyDelta),
            );
            setEvents([]);
            setSelectedEvent(null);
            setTrends(null);
            setTopViolators(null);
            setDriftStats(null);
            setExceptionReport(null);
            setSystemicWeaknesses(null);
            setExecutiveReport(null);
            setAnalyticsDelta(null);
            setAnalyticsAnomalies(null);
            setAnalyticsScorecards(null);
            setAnalyticsSegments(null);
            setTopologyGraph(null);
            setTopologyHeatmap(null);
            setTopologyDelta(null);
            setRuntimeActiveStates([]);
            setRuntimeClosedLoopStatus(null);
            setRuntimeDriftFindings([]);
            setRuntimeDriftStatus(null);
            setIncidents([]);
            setPendingExceptions([]);
          } else if (activeTab === "analytics") {
            promises.push(
              getTrends({
                window: "28d",
                window_days: "30",
                compare_to: "previous_window",
                group_by: "service",
                granularity: "day",
                ...baseFilters,
              }).then(setTrends),
            );
            promises.push(
              getAnalyticsDelta({
                window: "28d",
                compare_to: "previous_window",
                group_by: "service",
                metric: "policy_friction_rate",
                ...baseFilters,
              }).then(setAnalyticsDelta),
            );
            promises.push(
              getAnalyticsAnomalies({
                window: "28d",
                compare_to: "previous_window",
                group_by: "service",
                ...baseFilters,
              }).then(setAnalyticsAnomalies),
            );
            promises.push(
              getAnalyticsScorecards({
                window: "28d",
                compare_to: "previous_window",
                group_by: "service",
                ...baseFilters,
              }).then(setAnalyticsScorecards),
            );
            promises.push(
              getAnalyticsSegments({
                window: "28d",
                compare_to: "previous_window",
                group_by: "service",
                ...baseFilters,
              }).then(setAnalyticsSegments),
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
            setSystemicWeaknesses(null);
            setExecutiveReport(null);
            setRuntimeActiveStates([]);
            setRuntimeClosedLoopStatus(null);
            setRuntimeDriftFindings([]);
            setRuntimeDriftStatus(null);
            setExceptionReport(null);
            setSystemicWeaknesses(null);
            setExecutiveReport(null);
            setPendingExceptions([]);
          } else if (activeTab === "runtime") {
            promises.push(
              getRuntimeDriftFindings({
                tenant_id: scopedTenantID,
                limit: filters.limit,
              }).then((response) => setRuntimeDriftFindings(response.items)),
            );
            promises.push(
              getRuntimeDriftStatus({
                tenant_id: scopedTenantID,
                limit: filters.limit,
              }).then(setRuntimeDriftStatus),
            );
            promises.push(
              getRuntimeActiveStates({
                tenant_id: scopedTenantID,
                limit: filters.limit,
              }).then((response) => setRuntimeActiveStates(response.items)),
            );
            promises.push(
              getRuntimeClosedLoopStatus({
                tenant_id: scopedTenantID,
                limit: filters.limit,
              }).then(setRuntimeClosedLoopStatus),
            );
            promises.push(
              getEvents(activeTab, { ...filters, tenant_id: scopedTenantID }).then((response) => {
                setEvents(response.events);
                setSelectedEvent((current) => response.events.find((event) => event.id === current?.id) || response.events[0] || null);
              }),
            );
            setTrends(null);
            setTopViolators(null);
            setDriftStats(null);
            setAnalyticsDelta(null);
            setAnalyticsAnomalies(null);
            setAnalyticsScorecards(null);
            setAnalyticsSegments(null);
            setExceptionReport(null);
            setSystemicWeaknesses(null);
            setExecutiveReport(null);
            setPendingExceptions([]);
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
            setExecutiveReport(null);
            setAnalyticsDelta(null);
            setAnalyticsAnomalies(null);
            setAnalyticsScorecards(null);
            setAnalyticsSegments(null);
            setTopologyGraph(null);
            setTopologyHeatmap(null);
            setTopologyDelta(null);
          } else if (activeTab === "inventory" || activeTab === "vulnerabilities" || activeTab === "signing" || activeTab === "scorecard" || activeTab === "guidance") {
            setEvents([]);
            setSelectedEvent(null);
            setTrends(null);
            setTopViolators(null);
            setDriftStats(null);
            setAnalyticsDelta(null);
            setAnalyticsAnomalies(null);
            setAnalyticsScorecards(null);
            setAnalyticsSegments(null);
            setTopologyGraph(null);
            setTopologyHeatmap(null);
            setTopologyDelta(null);
            setRuntimeActiveStates([]);
            setRuntimeClosedLoopStatus(null);
            setRuntimeDriftFindings([]);
            setRuntimeDriftStatus(null);
            setExceptionReport(null);
            setSystemicWeaknesses(null);
            setExecutiveReport(null);
            setIncidents([]);
            setPendingExceptions([]);
          } else {
            promises.push(
              getEvents(activeTab, { ...filters, tenant_id: scopedTenantID }).then((response) => {
                setEvents(response.events);
                setSelectedEvent((current) => response.events.find((event) => event.id === current?.id) || response.events[0] || null);
              }),
            );
            setTrends(null);
            setTopViolators(null);
            setDriftStats(null);
            setAnalyticsDelta(null);
            setAnalyticsAnomalies(null);
            setAnalyticsScorecards(null);
            setAnalyticsSegments(null);
            setTopologyGraph(null);
            setTopologyHeatmap(null);
            setTopologyDelta(null);
            setRuntimeActiveStates([]);
            setRuntimeClosedLoopStatus(null);
            setRuntimeDriftFindings([]);
            setRuntimeDriftStatus(null);
            setExceptionReport(null);
            setSystemicWeaknesses(null);
            setExecutiveReport(null);
            setIncidents([]);
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
  }, [activeTab, filters.component, filters.decision, filters.environment, filters.limit, filters.repo, filters.tenant_id, metricDrilldown?.metricKey, refreshIndex]);

  const activeTabMeta = tabs.find((tab) => tab.key === activeTab) || tabs[0];
  const role = authStatus?.role;
  const enforcedTenantID = authStatus?.tenant_id || "";
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

  async function handleAcknowledgeIncident(incidentID: string, summary?: string) {
    await acknowledgeIncident(incidentID, summary);
    setRefreshIndex((value) => value + 1);
  }

  async function handleWatchIncident(incidentID: string, summary?: string) {
    await watchIncident(incidentID, summary);
    setRefreshIndex((value) => value + 1);
  }

  async function handleAssignIncident(incidentID: string, owner: string, reason: string) {
    await assignIncident(incidentID, owner, reason);
    setRefreshIndex((value) => value + 1);
  }

  async function handleResolveIncident(
    incidentID: string,
    input: {
      resolution_type: string;
      resolution_summary: string;
      resolution_details?: string;
      resolution_refs?: string[];
      follow_up_required?: boolean;
    },
  ) {
    await resolveIncident(incidentID, input);
    setRefreshIndex((value) => value + 1);
  }

  async function handleReopenIncident(incidentID: string, reason?: string) {
    await reopenIncident(incidentID, reason);
    setRefreshIndex((value) => value + 1);
  }

  async function handleIncidentNote(incidentID: string, note: string) {
    await addIncidentNote(incidentID, note);
    setRefreshIndex((value) => value + 1);
  }

  async function handleLoadIncidentExport(incidentID: string, audience: IncidentReportAudience): Promise<IncidentExport> {
    return getIncidentExport(incidentID, {
      environment: filters.environment,
      tenant_id: enforcedTenantID || filters.tenant_id,
      repo: filters.repo,
    }, audience);
  }

  async function handleLoadIncidentPackage(incidentIDs: string[], audience: IncidentReportAudience): Promise<IncidentPackage> {
    return getIncidentPackage({
      environment: filters.environment,
      tenant_id: enforcedTenantID || filters.tenant_id,
      repo: filters.repo,
      scorecard_ref: metricDrilldown?.metricKey,
    }, incidentIDs, audience);
  }

  async function handleLoadExecutiveReport(incidentIDs: string[], audience: IncidentReportAudience): Promise<ExecutiveDefenseReport> {
    return getExecutiveDefenseReport({
      environment: filters.environment,
      tenant_id: enforcedTenantID || filters.tenant_id,
      repo: filters.repo,
      scorecard_ref: metricDrilldown?.metricKey,
    }, incidentIDs, audience);
  }

  async function handleLoadIncidentDefenseGaps(incidentID: string): Promise<DefenseGapAssessment> {
    return getIncidentDefenseGaps(incidentID, {
      environment: filters.environment,
      tenant_id: enforcedTenantID || filters.tenant_id,
      repo: filters.repo,
    });
  }

  async function handleLoadMetricDefenseGaps(metricKey: string): Promise<DefenseGapAssessment> {
    return getMetricDefenseGaps(metricKey, {
      ...filters,
      tenant_id: enforcedTenantID || filters.tenant_id,
    });
  }

  async function handleLoadIncidentPolicyReplay(incidentID: string): Promise<PolicyReplayAssessment> {
    return getIncidentPolicyReplay(incidentID, {
      environment: filters.environment,
      tenant_id: enforcedTenantID || filters.tenant_id,
      repo: filters.repo,
    });
  }

  async function handleLoadMetricPolicyReplay(metricKey: string): Promise<PolicyReplayAssessment> {
    return getMetricPolicyReplay(metricKey, {
      ...filters,
      tenant_id: enforcedTenantID || filters.tenant_id,
    });
  }

  async function handleLoadIncidentBlastRadius(incidentID: string): Promise<TopologyBlastRadiusResponse> {
    return getIncidentBlastRadius(incidentID, {
      environment: filters.environment,
      tenant_id: enforcedTenantID || filters.tenant_id,
      repo: filters.repo,
    });
  }

  async function handleLoadMetricBlastRadius(metricKey: string): Promise<TopologyBlastRadiusResponse> {
    return getMetricBlastRadius(metricKey, {
      ...filters,
      tenant_id: enforcedTenantID || filters.tenant_id,
    });
  }

  return (
    <main className="app-shell">
      <header className="hero">
        <div>
          <p className="eyebrow">ChangeLock Security Dashboard</p>
          <h1>Security posture and active incidents</h1>
          <p className="hero-copy">
            Live ChangeLock audit signals, grouped into the operator questions that matter first: what changed, what is at
            risk, what is affected, and what to do next.
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
                ? `${authStatus.email || authStatus.subject}${authStatus.tenant_id ? ` · tenant ${authStatus.tenant_id}` : authStatus?.global_scope ? " · global scope" : ""}${authStatus.token_id ? ` (${authStatus.token_id})` : ""}`
                : apiTokenConfigured()
                  ? "Token configured"
                  : "No token configured"}
            </small>
          </div>
          <div className="panel health-panel">
            <span className="summary-label">Cross-Cluster Sync</span>
            <strong>{syncStatus ? `${syncStatus.mode} · ${syncStatus.health}` : "unavailable"}</strong>
            <small>
              {syncStatus
                ? `${syncStatus.cluster_id ? `${syncStatus.cluster_id} · ` : ""}${syncStatus.current_revision ? `rev ${syncStatus.current_revision.slice(0, 12)}` : "no revision"}${syncStatus.fail_mode ? ` · ${syncStatus.fail_mode}` : ""}${syncStatus.summary ? ` · ${syncStatus.summary}` : ""}`
                : "No sync status loaded"}
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

      {activeTab !== "overview" ? (
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
      ) : (
        <section className="panel overview-scope-bar">
          <div>
            <span className="summary-label">Posture workspace</span>
            <strong>{activeTabMeta.description}</strong>
            {lastLoadedAt ? <p className="tab-header__meta">Last refresh at {lastLoadedAt}</p> : null}
          </div>
          <button className="button" onClick={() => setRefreshIndex((value) => value + 1)}>
            Refresh posture
          </button>
        </section>
      )}

      {activeTab !== "inventory" && activeTab !== "vulnerabilities" && activeTab !== "signing" && activeTab !== "scorecard" && activeTab !== "guidance" ? (
        <Filters
          filters={filters}
          tab={activeTab}
          enforcedTenantID={enforcedTenantID || undefined}
          onChange={(name, value) => setFilters((current) => ({ ...current, [name]: value }))}
          onRefresh={() => setRefreshIndex((value) => value + 1)}
          onReset={() => setFilters({ ...initialFilters, tenant_id: enforcedTenantID })}
        />
      ) : null}

      {activeTab === "overview" ? (
        <>
          <OverviewDashboard
            health={health}
            summary={summary}
            trends={trends}
            topViolators={topViolators}
            driftStats={driftStats}
            exceptionReport={exceptionReport}
            systemicWeaknesses={systemicWeaknesses}
            executiveReport={executiveReport}
            syncStatus={syncStatus}
            loading={loading}
            onSelectTrustMetric={(metricKey, label) => {
              setMetricDrilldown({
                metricKey,
                metricLabel: label,
                incidents: [],
                limitations: [],
              });
              setActiveTab("events");
            }}
          />

          <section className="panel overview-evidence-header">
            <div>
              <span className="summary-label">Recent evidence-backed decisions</span>
              <h2>Investigation feed</h2>
              <p>Use the latest events and evidence payloads to validate root cause before widening scope or approving exceptions.</p>
            </div>
          </section>
        </>
      ) : null}

      {activeTab === "analytics" ? (
        <section className="analytics-grid">
          <AnalyticsInsightsPanel
            trends={trends}
            delta={analyticsDelta}
            anomalies={analyticsAnomalies}
            scorecards={analyticsScorecards}
            segments={analyticsSegments}
            loading={loading}
          />
          <AnalyticsTrendsPanel trends={trends} loading={loading} />
          <TopViolatorsPanel data={topViolators} loading={loading} />
          <DriftStatsPanel data={driftStats} loading={loading} />
        </section>
      ) : null}

      {activeTab === "topology" ? (
        <section className="analytics-grid">
          <TopologyInsightsPanel
            graph={topologyGraph}
            heatmap={topologyHeatmap}
            delta={topologyDelta}
            loading={loading}
          />
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
            <ExceptionRequestForm enabled={canRequest} submitting={requestSubmitting} enforcedTenantID={enforcedTenantID || undefined} onSubmit={handleRequestException} />
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

      {activeTab === "runtime" ? (
        <>
          <RuntimeDriftPanel
            findings={runtimeDriftFindings}
            status={runtimeDriftStatus}
            activeStates={runtimeActiveStates}
            closedLoopStatus={runtimeClosedLoopStatus}
            loading={loading}
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
        </>
      ) : null}

      {activeTab === "inventory" ? <SBOMInventoryPanel tenantID={enforcedTenantID || undefined} /> : null}

      {activeTab === "vulnerabilities" ? <VulnerabilityOpsPanel role={role} tenantID={enforcedTenantID || undefined} /> : null}
      {activeTab === "signing" ? <SigningIdentityPanel tenantID={enforcedTenantID || undefined} /> : null}
      {activeTab === "scorecard" ? <TrustScorecardPanel tenantID={enforcedTenantID || undefined} /> : null}
      {activeTab === "guidance" ? <AIInsightsPanel tenantID={enforcedTenantID || undefined} /> : null}

      {activeTab === "events" ? (
        <IncidentWorkbench
          incidents={incidents}
          events={events}
          loading={loading}
          error={error}
          role={role}
          metricDrilldown={metricDrilldown}
          onClearMetricDrilldown={() => setMetricDrilldown(null)}
          onLoadExport={handleLoadIncidentExport}
          onLoadPackage={handleLoadIncidentPackage}
          onLoadExecutiveReport={handleLoadExecutiveReport}
          onLoadIncidentDefenseGaps={handleLoadIncidentDefenseGaps}
          onLoadMetricDefenseGaps={handleLoadMetricDefenseGaps}
          onLoadIncidentPolicyReplay={handleLoadIncidentPolicyReplay}
          onLoadMetricPolicyReplay={handleLoadMetricPolicyReplay}
          onLoadIncidentBlastRadius={handleLoadIncidentBlastRadius}
          onLoadMetricBlastRadius={handleLoadMetricBlastRadius}
          onAcknowledge={handleAcknowledgeIncident}
          onWatch={handleWatchIncident}
          onAssign={handleAssignIncident}
          onResolve={handleResolveIncident}
          onReopen={handleReopenIncident}
          onAddNote={handleIncidentNote}
        />
      ) : null}

      {activeTab !== "analytics" && activeTab !== "topology" && activeTab !== "runtime" && activeTab !== "exceptions" && activeTab !== "inventory" && activeTab !== "vulnerabilities" && activeTab !== "signing" && activeTab !== "scorecard" && activeTab !== "guidance" ? (
        <section className="content-grid">
          <EventsTable
            events={events}
            selectedEventID={selectedEvent?.id || null}
            onSelect={setSelectedEvent}
            loading={loading}
            error={error}
            title={activeTab === "overview" ? "Recent Decisions" : "Recent Events"}
            emptyMessage={activeTab === "overview" ? "No recent evidence-backed decisions matched the current scope." : undefined}
          />
          <EventDetails event={selectedEvent} />
        </section>
      ) : null}
    </main>
  );
}
