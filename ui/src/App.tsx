import { useEffect, useState } from "react";

import {
  APIError,
  acknowledgeIncident,
  acknowledgeRecommendation,
  addIncidentNote,
  apiBaseURL,
  apiTokenConfigured,
  assignRecommendation,
  commentRecommendation,
  downloadHandoffBundle,
  getCommandCenterNotifications,
  getCommandCenterSearch,
  getAnalyticsAnomalies,
  getAnalyticsDelta,
  getAnalyticsScorecards,
  getAnalyticsSegments,
  getForensicsDelta,
  getForensicsState,
  getForensicsTimeline,
  getForensicsVEXFlashback,
  getHandoffVerification,
  approveException,
  acceptRecommendation,
  assignIncident,
  getAuthStatus,
  getDriftStats,
  getExecutiveDefenseReport,
  getEvents,
  getExceptionReport,
  getExceptions,
  getFederationGlobalView,
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
  getHardeningActions,
  getHardeningPosture,
  getRecommendations,
  getRuntimeActiveStates,
  getRuntimeClosedLoopStatus,
  getRuntimeDriftFindings,
  getRuntimeDriftStatus,
  getRuntimeEnforcement,
  getRuntimeFindings,
  getRuntimeIntegrity,
  getSecurityTimeline,
  getRuntimeWorkloads,
  getSummary,
  getSystemicWeaknesses,
  getSyncStatus,
  getValidationHarnessRuns,
  getValidationHarnessScenarios,
  getValidationHarnessScore,
  getTopologyDelta,
  getTopologyGraph,
  getTopologyHeatmap,
  getTopViolators,
  getTrends,
  rejectException,
  rejectRecommendation,
  reopenIncident,
  resolveIncident,
  requestException,
  requestRecommendationApproval,
  revokeException,
  sealHandoff,
  verifyRecommendation,
  watchIncident,
  executeRecommendation,
  runForensicsReplay,
  runValidationHarness,
  runValidationHarnessWhatIf,
} from "./api";
import { AIInsightsPanel } from "./components/AIInsightsPanel";
import { AnalyticsInsightsPanel } from "./components/AnalyticsInsightsPanel";
import { AnalyticsTrendsPanel } from "./components/AnalyticsTrendsPanel";
import { CommandCenterPanel } from "./components/CommandCenterPanel";
import { DefensePosturePanel } from "./components/DefensePosturePanel";
import { DriftStatsPanel } from "./components/DriftStatsPanel";
import { EventDetails } from "./components/EventDetails";
import { EventsTable } from "./components/EventsTable";
import { ExceptionRequestForm } from "./components/ExceptionRequestForm";
import { Filters } from "./components/Filters";
import { FederationInsightsPanel } from "./components/FederationInsightsPanel";
import { ForensicsInsightsPanel } from "./components/ForensicsInsightsPanel";
import { HealthBadge } from "./components/HealthBadge";
import { IncidentWorkbench } from "./components/IncidentWorkbench";
import { OverviewDashboard } from "./components/OverviewDashboard";
import { PendingExceptionsPanel } from "./components/PendingExceptionsPanel";
import { RuntimeDriftPanel } from "./components/RuntimeDriftPanel";
import { RuntimeIntegrityPanel } from "./components/RuntimeIntegrityPanel";
import { SBOMInventoryPanel } from "./components/SBOMInventoryPanel";
import { SigningIdentityPanel } from "./components/SigningIdentityPanel";
import { ValidationHarnessPanel } from "./components/ValidationHarnessPanel";
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
  Recommendation,
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
  CommandCenterPersona,
  CommandCenterFocusTarget,
  CommandCenterNotificationsResponse,
  CommandCenterSearchResponse,
  DefensePostureState,
  DriftStatsResponse,
  EventFilters,
  FederationGlobalView,
  ExceptionReport,
  ExceptionRequestInput,
  HardeningExecutionRecord,
  HandoffSealResponse,
  PolicyException,
  RuntimeActiveState,
  RuntimeClosedLoopStatus,
  RuntimeDriftFinding,
  RuntimeDriftStatus,
  RuntimeEnforcementDecision,
  RuntimeIntegrityFinding,
  RuntimeIntegrityState,
  RuntimeWorkloadView,
  SecurityTimelineResponse,
  StoredEvent,
  Summary,
  SyncStatus,
  TabKey,
  ForensicReplayResponse,
  ForensicTimelineResponse,
  PointInTimeState,
  TimeDeltaResult,
  TopViolatorsResponse,
  TopologyBlastRadiusResponse,
  TopologyDeltaResponse,
  TopologyGraphResponse,
  TopologyHeatmapResponse,
  TrendsResponse,
  ValidationHarnessRun,
  ValidationHarnessScenario,
  ValidationHarnessScore,
  ValidationHarnessWhatIfResponse,
  VEXFlashbackResponse,
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
  { key: "forensics", label: "Forensics", description: "Point-in-time reconstruction, VEX flashback, timeline, and counterfactual replay." },
  { key: "federation", label: "Federation", description: "Cross-region proof reuse, trust decisions, policy sync, and anchor health." },
  { key: "validation", label: "Validation", description: "Controlled policy dry-runs, bounded chaos rehearsal, and what-if confidence checks." },
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

function defaultPersonaForRole(role?: string): CommandCenterPersona {
  if (role === "operator") {
    return "platform_operator";
  }
  if (role === "security_admin") {
    return "security_engineer";
  }
  return "developer";
}

function isTabKey(value: string | null): value is TabKey {
  return tabs.some((tab) => tab.key === value);
}

function readInitialNavigationState(): {
  tab: TabKey;
  query: string;
  lifecyclePhase: string;
  focusTarget: CommandCenterFocusTarget | null;
} {
  if (typeof window === "undefined") {
    return { tab: "overview", query: "", lifecyclePhase: "", focusTarget: null };
  }
  const params = new URLSearchParams(window.location.search);
  const tab = isTabKey(params.get("tab")) ? (params.get("tab") as TabKey) : "overview";
  const kind = params.get("focus_kind");
  const ref = params.get("focus_ref");
  const focusTab = isTabKey(params.get("focus_tab")) ? (params.get("focus_tab") as TabKey) : tab;
  return {
    tab,
    query: params.get("q") || "",
    lifecyclePhase: params.get("lifecycle_phase") || "",
    focusTarget:
      kind && ref
        ? {
            tab: focusTab,
            kind: kind as CommandCenterFocusTarget["kind"],
            ref,
            secondary_ref: params.get("focus_secondary_ref") || undefined,
            resource_uri: params.get("resource_uri") || undefined,
          }
        : null,
  };
}

function toDateTimeLocalValue(date: Date) {
  const year = date.getFullYear();
  const month = `${date.getMonth() + 1}`.padStart(2, "0");
  const day = `${date.getDate()}`.padStart(2, "0");
  const hours = `${date.getHours()}`.padStart(2, "0");
  const minutes = `${date.getMinutes()}`.padStart(2, "0");
  return `${year}-${month}-${day}T${hours}:${minutes}`;
}

function toForensicsTimestamp(value: string) {
  const parsed = value ? new Date(value) : new Date();
  if (Number.isNaN(parsed.getTime())) {
    return new Date().toISOString();
  }
  return parsed.toISOString();
}

function firstNonEmptyValue(...values: Array<string | null | undefined>) {
  return values.find((value) => typeof value === "string" && value.trim() !== "") || null;
}

export default function App() {
  const initialNavigation = readInitialNavigationState();
  const [activeTab, setActiveTab] = useState<TabKey>(initialNavigation.tab);
  const [commandCenterPersona, setCommandCenterPersona] = useState<CommandCenterPersona>("developer");
  const [commandCenterQuery, setCommandCenterQuery] = useState(initialNavigation.query);
  const [commandCenterLifecyclePhase, setCommandCenterLifecyclePhase] = useState(initialNavigation.lifecyclePhase);
  const [commandCenterNotifications, setCommandCenterNotifications] = useState<CommandCenterNotificationsResponse | null>(null);
  const [commandCenterSearch, setCommandCenterSearch] = useState<CommandCenterSearchResponse | null>(null);
  const [commandCenterSearchLoading, setCommandCenterSearchLoading] = useState(false);
  const [focusTarget, setFocusTarget] = useState<CommandCenterFocusTarget | null>(initialNavigation.focusTarget);
  const [filters, setFilters] = useState<EventFilters>(initialFilters);
  const [summary, setSummary] = useState<Summary | null>(null);
  const [securityTimeline, setSecurityTimeline] = useState<SecurityTimelineResponse | null>(null);
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
  const [forensicsState, setForensicsState] = useState<PointInTimeState | null>(null);
  const [forensicsDelta, setForensicsDelta] = useState<TimeDeltaResult | null>(null);
  const [forensicsTimeline, setForensicsTimeline] = useState<ForensicTimelineResponse | null>(null);
  const [forensicsVEXFlashback, setForensicsVEXFlashback] = useState<VEXFlashbackResponse | null>(null);
  const [forensicsReplay, setForensicsReplay] = useState<ForensicReplayResponse | null>(null);
  const [federationView, setFederationView] = useState<FederationGlobalView | null>(null);
  const [validationScenarios, setValidationScenarios] = useState<ValidationHarnessScenario[]>([]);
  const [validationScore, setValidationScore] = useState<ValidationHarnessScore | null>(null);
  const [validationRuns, setValidationRuns] = useState<ValidationHarnessRun[]>([]);
  const [validationWhatIf, setValidationWhatIf] = useState<ValidationHarnessWhatIfResponse | null>(null);
  const [forensicsTimestamp, setForensicsTimestamp] = useState(() => toDateTimeLocalValue(new Date(Date.now() - 7 * 24 * 60 * 60 * 1000)));
  const [runtimeIntegrityStates, setRuntimeIntegrityStates] = useState<RuntimeIntegrityState[]>([]);
  const [runtimeWorkloads, setRuntimeWorkloads] = useState<RuntimeWorkloadView[]>([]);
  const [runtimeIntegrityFindings, setRuntimeIntegrityFindings] = useState<RuntimeIntegrityFinding[]>([]);
  const [runtimeEnforcement, setRuntimeEnforcement] = useState<RuntimeEnforcementDecision[]>([]);
  const [hardeningPosture, setHardeningPosture] = useState<DefensePostureState[]>([]);
  const [hardeningActions, setHardeningActions] = useState<HardeningExecutionRecord[]>([]);
  const [runtimeActiveStates, setRuntimeActiveStates] = useState<RuntimeActiveState[]>([]);
  const [runtimeClosedLoopStatus, setRuntimeClosedLoopStatus] = useState<RuntimeClosedLoopStatus | null>(null);
  const [runtimeDriftFindings, setRuntimeDriftFindings] = useState<RuntimeDriftFinding[]>([]);
  const [runtimeDriftStatus, setRuntimeDriftStatus] = useState<RuntimeDriftStatus | null>(null);
  const [exceptionReport, setExceptionReport] = useState<ExceptionReport | null>(null);
  const [incidents, setIncidents] = useState<InvestigationIncident[]>([]);
  const [systemicWeaknesses, setSystemicWeaknesses] = useState<SystemicWeaknessResponse | null>(null);
  const [executiveReport, setExecutiveReport] = useState<ExecutiveDefenseReport | null>(null);
  const [recommendations, setRecommendations] = useState<Recommendation[]>([]);
  const [pendingExceptions, setPendingExceptions] = useState<PolicyException[]>([]);
  const [metricDrilldown, setMetricDrilldown] = useState<MetricIncidentDrilldown | null>(null);
  const [loading, setLoading] = useState(true);
  const [requestSubmitting, setRequestSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [refreshIndex, setRefreshIndex] = useState(0);
  const [lastLoadedAt, setLastLoadedAt] = useState<string>("");

  function handleOpenTab(tab: TabKey) {
    setFocusTarget(null);
    setActiveTab(tab);
  }

  function handleOpenFocusTarget(target: CommandCenterFocusTarget) {
    setFocusTarget(target);
    setActiveTab(target.tab);
  }

  function handleSubmitCommandCenterSearch() {
    setCommandCenterSearchLoading(true);
    setRefreshIndex((current) => current + 1);
  }

  useEffect(() => {
    if (typeof window === "undefined") {
      return;
    }
    const params = new URLSearchParams(window.location.search);
    params.set("tab", activeTab);
    if (commandCenterQuery.trim() !== "") {
      params.set("q", commandCenterQuery.trim());
    } else {
      params.delete("q");
    }
    if (commandCenterLifecyclePhase.trim() !== "") {
      params.set("lifecycle_phase", commandCenterLifecyclePhase.trim());
    } else {
      params.delete("lifecycle_phase");
    }
    if (focusTarget) {
      params.set("focus_tab", focusTarget.tab);
      params.set("focus_kind", focusTarget.kind);
      params.set("focus_ref", focusTarget.ref);
      if (focusTarget.secondary_ref) {
        params.set("focus_secondary_ref", focusTarget.secondary_ref);
      } else {
        params.delete("focus_secondary_ref");
      }
      if (focusTarget.resource_uri) {
        params.set("resource_uri", focusTarget.resource_uri);
      } else {
        params.delete("resource_uri");
      }
    } else {
      params.delete("focus_tab");
      params.delete("focus_kind");
      params.delete("focus_ref");
      params.delete("focus_secondary_ref");
      params.delete("resource_uri");
    }
    const next = `${window.location.pathname}${params.toString() ? `?${params.toString()}` : ""}`;
    window.history.replaceState(null, "", next);
  }, [activeTab, commandCenterLifecyclePhase, commandCenterQuery, focusTarget]);

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
          setSecurityTimeline(null);
          setCommandCenterNotifications(null);
          setCommandCenterSearch(null);
          setCommandCenterSearchLoading(false);
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
          setForensicsState(null);
          setForensicsDelta(null);
          setForensicsTimeline(null);
          setForensicsVEXFlashback(null);
          setForensicsReplay(null);
          setFederationView(null);
          setRuntimeIntegrityStates([]);
          setRuntimeWorkloads([]);
          setRuntimeIntegrityFindings([]);
          setRuntimeEnforcement([]);
          setHardeningPosture([]);
          setHardeningActions([]);
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
          const forensicTimestampISO = toForensicsTimestamp(forensicsTimestamp);
          const forensicWindowEnd = new Date().toISOString();

          setForensicsState(null);
          setForensicsDelta(null);
          setForensicsTimeline(null);
          setForensicsVEXFlashback(null);
          setForensicsReplay(null);
          setFederationView(null);

          const promises: Array<Promise<void>> = [
            getSummary({ environment: filters.environment, tenant_id: scopedTenantID }).then(setSummary),
          ];
          if (activeTab !== "overview") {
            setCommandCenterNotifications(null);
            setCommandCenterSearchLoading(false);
          }
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
              getSecurityTimeline({
                component: filters.component,
                decision: filters.decision,
                environment: filters.environment,
                tenant_id: scopedTenantID,
                repo: filters.repo,
                limit: filters.limit,
                lifecycle_phase: commandCenterLifecyclePhase,
              }).then(setSecurityTimeline).catch((timelineError) => {
                if (isOptionalFeatureMissing(timelineError)) {
                  setSecurityTimeline(null);
                  return;
                }
                throw timelineError;
              }),
            );
            promises.push(
              getCommandCenterNotifications({
                environment: filters.environment,
                tenant_id: scopedTenantID,
                repo: filters.repo,
                limit: filters.limit,
                lifecycle_phase: commandCenterLifecyclePhase,
              }).then(setCommandCenterNotifications).catch((notificationsError) => {
                if (isOptionalFeatureMissing(notificationsError)) {
                  setCommandCenterNotifications(null);
                  return;
                }
                throw notificationsError;
              }),
            );
            if (commandCenterQuery.trim() !== "") {
              promises.push(
                getCommandCenterSearch({
                  q: commandCenterQuery.trim(),
                  environment: filters.environment,
                  tenant_id: scopedTenantID,
                  repo: filters.repo,
                  limit: filters.limit,
                })
                  .then(setCommandCenterSearch)
                  .finally(() => setCommandCenterSearchLoading(false)),
              );
            } else {
              setCommandCenterSearch(null);
              setCommandCenterSearchLoading(false);
            }
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
            promises.push(
              getRecommendations({
                environment: filters.environment,
                tenant_id: scopedTenantID,
                repo: filters.repo,
                limit: "6",
              }).then(setRecommendations).catch((recommendationError) => {
                if (isOptionalFeatureMissing(recommendationError)) {
                  setRecommendations([]);
                  return;
                }
                throw recommendationError;
              }),
            );
            setPendingExceptions([]);
            setIncidents([]);
          } else if (activeTab === "events") {
            setSecurityTimeline(null);
            setCommandCenterSearchLoading(false);
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
            setRecommendations([]);
            setPendingExceptions([]);
          } else if (activeTab === "topology") {
            setSecurityTimeline(null);
            setCommandCenterSearchLoading(false);
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
            setRecommendations([]);
            setRuntimeIntegrityStates([]);
            setRuntimeWorkloads([]);
            setRuntimeIntegrityFindings([]);
            setRuntimeEnforcement([]);
            setHardeningPosture([]);
            setHardeningActions([]);
            setRuntimeActiveStates([]);
            setRuntimeClosedLoopStatus(null);
            setRuntimeDriftFindings([]);
            setRuntimeDriftStatus(null);
            setIncidents([]);
            setPendingExceptions([]);
          } else if (activeTab === "forensics") {
            setSecurityTimeline(null);
            setCommandCenterSearchLoading(false);
            promises.push(
              getForensicsState({
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                timestamp: forensicTimestampISO,
                limit: filters.limit,
              }).then(setForensicsState),
            );
            promises.push(
              getForensicsDelta({
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                t1: forensicTimestampISO,
                t2: forensicWindowEnd,
                limit: filters.limit,
              }).then(setForensicsDelta),
            );
            promises.push(
              getForensicsTimeline({
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                t1: forensicTimestampISO,
                t2: forensicWindowEnd,
                limit: filters.limit,
              }).then(setForensicsTimeline),
            );
            promises.push(
              getForensicsVEXFlashback({
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                timestamp: forensicTimestampISO,
                limit: filters.limit,
              }).then(setForensicsVEXFlashback),
            );
            promises.push(
              runForensicsReplay(
                {
                  tenant_id: scopedTenantID,
                  environment: filters.environment,
                  repo: filters.repo,
                },
                {
                  timestamp: forensicTimestampISO,
                  replay_mode: "modern_full_stack_replay",
                },
              ).then(setForensicsReplay),
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
            setRecommendations([]);
            setRuntimeIntegrityStates([]);
            setRuntimeWorkloads([]);
            setRuntimeIntegrityFindings([]);
            setRuntimeEnforcement([]);
            setHardeningPosture([]);
            setHardeningActions([]);
            setRuntimeActiveStates([]);
            setRuntimeClosedLoopStatus(null);
            setRuntimeDriftFindings([]);
            setRuntimeDriftStatus(null);
            setIncidents([]);
            setPendingExceptions([]);
          } else if (activeTab === "federation") {
            setSecurityTimeline(null);
            promises.push(
              getFederationGlobalView().then(setFederationView),
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
            setForensicsState(null);
            setForensicsDelta(null);
            setForensicsTimeline(null);
            setForensicsVEXFlashback(null);
            setForensicsReplay(null);
            setRecommendations([]);
            setRuntimeActiveStates([]);
            setRuntimeClosedLoopStatus(null);
            setRuntimeDriftFindings([]);
            setRuntimeDriftStatus(null);
            setIncidents([]);
            setPendingExceptions([]);
          } else if (activeTab === "validation") {
            setSecurityTimeline(null);
            promises.push(
              getValidationHarnessScenarios().then(setValidationScenarios),
            );
            promises.push(
              getValidationHarnessScore({
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                limit: filters.limit,
              }).then(setValidationScore),
            );
            promises.push(
              getValidationHarnessRuns({
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                limit: filters.limit,
              }).then(setValidationRuns),
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
            setForensicsState(null);
            setForensicsDelta(null);
            setForensicsTimeline(null);
            setForensicsVEXFlashback(null);
            setForensicsReplay(null);
            setRecommendations([]);
            setRuntimeActiveStates([]);
            setRuntimeClosedLoopStatus(null);
            setRuntimeDriftFindings([]);
            setRuntimeDriftStatus(null);
            setIncidents([]);
            setPendingExceptions([]);
          } else if (activeTab === "analytics") {
            setSecurityTimeline(null);
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
            setRecommendations([]);
            setPendingExceptions([]);
          } else if (activeTab === "runtime") {
            setSecurityTimeline(null);
            promises.push(
              getRuntimeIntegrity({
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                limit: filters.limit,
              }).then((response) => setRuntimeIntegrityStates(response.items)),
            );
            promises.push(
              getRuntimeWorkloads({
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                limit: filters.limit,
              }).then((response) => setRuntimeWorkloads(response.items)),
            );
            promises.push(
              getRuntimeFindings({
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                limit: filters.limit,
              }).then((response) => setRuntimeIntegrityFindings(response.items)),
            );
            promises.push(
              getRuntimeEnforcement({
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                limit: filters.limit,
              }).then((response) => setRuntimeEnforcement(response.items)),
            );
            promises.push(
              getHardeningPosture({
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                limit: filters.limit,
              }).then(setHardeningPosture),
            );
            promises.push(
              getHardeningActions({
                tenant_id: scopedTenantID,
                environment: filters.environment,
                repo: filters.repo,
                limit: filters.limit,
              }).then(setHardeningActions),
            );
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
            setTopologyGraph(null);
            setTopologyHeatmap(null);
            setTopologyDelta(null);
            setForensicsState(null);
            setForensicsDelta(null);
            setForensicsTimeline(null);
            setForensicsVEXFlashback(null);
            setForensicsReplay(null);
            setFederationView(null);
            setRecommendations([]);
            setExceptionReport(null);
            setSystemicWeaknesses(null);
            setExecutiveReport(null);
            setPendingExceptions([]);
          } else if (activeTab === "exceptions") {
            setSecurityTimeline(null);
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
            setRecommendations([]);
          } else if (activeTab === "inventory" || activeTab === "vulnerabilities" || activeTab === "signing" || activeTab === "scorecard" || activeTab === "guidance") {
            setSecurityTimeline(null);
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
            setFederationView(null);
            setRuntimeActiveStates([]);
            setRuntimeClosedLoopStatus(null);
            setRuntimeDriftFindings([]);
            setRuntimeDriftStatus(null);
            setExceptionReport(null);
            setSystemicWeaknesses(null);
            setExecutiveReport(null);
            setRecommendations([]);
            setIncidents([]);
            setPendingExceptions([]);
          } else {
            setSecurityTimeline(null);
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
            setRecommendations([]);
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
  }, [activeTab, commandCenterLifecyclePhase, filters.component, filters.decision, filters.environment, filters.limit, filters.repo, filters.tenant_id, forensicsTimestamp, metricDrilldown?.metricKey, refreshIndex]);

  const activeTabMeta = tabs.find((tab) => tab.key === activeTab) || tabs[0];
  const role = authStatus?.role;
  const enforcedTenantID = authStatus?.tenant_id || "";
  const canRequest = role === "operator" || role === "security_admin";
  const canApprove = role === "security_admin";

  useEffect(() => {
    setCommandCenterPersona(defaultPersonaForRole(role));
  }, [role]);

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

  async function handleSealHandoff(input: {
    incidentIDs: string[];
    audience: IncidentReportAudience;
    includeForensics?: boolean;
    includeRuntime?: boolean;
    includeValidation?: boolean;
    includeRecommendations?: boolean;
    coSignMode?: string;
  }): Promise<HandoffSealResponse> {
    return sealHandoff({
      audience: input.audience,
      incident_ids: input.incidentIDs,
      include_forensics: input.includeForensics,
      include_runtime: input.includeRuntime,
      include_validation: input.includeValidation,
      include_recommendations: input.includeRecommendations,
      co_sign_mode: input.coSignMode,
    }, {
      environment: filters.environment,
      tenant_id: enforcedTenantID || filters.tenant_id,
      repo: filters.repo,
      scorecard_ref: metricDrilldown?.metricKey,
    });
  }

  async function handleGetHandoffVerification(packageID: string) {
    return getHandoffVerification(packageID);
  }

  async function handleDownloadHandoff(packageID: string) {
    const blob = await downloadHandoffBundle(packageID);
    const url = URL.createObjectURL(blob);
    const anchor = document.createElement("a");
    anchor.href = url;
    anchor.download = `${packageID}.safepkg`;
    document.body.appendChild(anchor);
    anchor.click();
    anchor.remove();
    URL.revokeObjectURL(url);
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

  async function handleLoadRecommendations(input: {
    incidentIDs?: string[];
    packageIncidentIDs?: string[];
    sourceType?: string;
    limit?: string;
  }): Promise<Recommendation[]> {
    return getRecommendations({
      environment: filters.environment,
      tenant_id: enforcedTenantID || filters.tenant_id,
      repo: filters.repo,
      limit: input.limit || filters.limit,
      source_type: input.sourceType,
    }, {
      incidentIDs: input.incidentIDs,
      packageIncidentIDs: input.packageIncidentIDs,
    });
  }

  async function handleAcknowledgeRecommendation(recommendationID: string) {
    await acknowledgeRecommendation(recommendationID);
    setRefreshIndex((value) => value + 1);
  }

  async function handleAcceptRecommendation(recommendationID: string) {
    await acceptRecommendation(recommendationID);
    setRefreshIndex((value) => value + 1);
  }

  async function handleRejectRecommendation(recommendationID: string, reason: string) {
    await rejectRecommendation(recommendationID, reason);
    setRefreshIndex((value) => value + 1);
  }

  async function handleExecuteRecommendation(recommendationID: string, input?: { template_id?: string; summary?: string }) {
    await executeRecommendation(recommendationID, input);
    setRefreshIndex((value) => value + 1);
  }

  async function handleVerifyRecommendation(recommendationID: string) {
    await verifyRecommendation(recommendationID);
    setRefreshIndex((value) => value + 1);
  }

  async function handleAssignRecommendation(recommendationID: string, owner: string, reason?: string) {
    await assignRecommendation(recommendationID, owner, reason);
    setRefreshIndex((value) => value + 1);
  }

  async function handleCommentRecommendation(recommendationID: string, comment: string) {
    await commentRecommendation(recommendationID, comment);
    setRefreshIndex((value) => value + 1);
  }

  async function handleRequestRecommendationApproval(recommendationID: string, summary?: string) {
    await requestRecommendationApproval(recommendationID, summary);
    setRefreshIndex((value) => value + 1);
  }

  async function handleRunValidationHarness() {
    const scopedTenantID = enforcedTenantID || filters.tenant_id;
    await runValidationHarness({
      tenant_id: scopedTenantID,
      environment: filters.environment,
      repo: filters.repo,
      limit: filters.limit,
    }, {
      mode: "policy_dry_run",
    });
    setRefreshIndex((value) => value + 1);
  }

  async function handleRunValidationWhatIf() {
    const scopedTenantID = enforcedTenantID || filters.tenant_id;
    const response = await runValidationHarnessWhatIf({
      tenant_id: scopedTenantID,
      environment: filters.environment,
      repo: filters.repo,
      limit: filters.limit,
    }, {
      inject_critical_vulnerability: true,
      rekor_unavailable: true,
      tighten_runtime_restrictions: true,
    });
    setValidationWhatIf(response);
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
            onClick={() => handleOpenTab(tab.key)}
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
          <CommandCenterPanel
            persona={commandCenterPersona}
            timeline={securityTimeline}
            notifications={commandCenterNotifications}
            recommendations={recommendations}
            executiveReport={executiveReport}
            loading={loading}
            searchQuery={commandCenterQuery}
            lifecyclePhase={commandCenterLifecyclePhase}
            searchResults={commandCenterSearch}
            searchLoading={commandCenterSearchLoading}
            focusTarget={focusTarget}
            onPersonaChange={setCommandCenterPersona}
            onOpenTab={handleOpenTab}
            onOpenTarget={handleOpenFocusTarget}
            onLifecyclePhaseChange={setCommandCenterLifecyclePhase}
            onSearchQueryChange={setCommandCenterQuery}
            onSearchSubmit={handleSubmitCommandCenterSearch}
          />

          <OverviewDashboard
            health={health}
            summary={summary}
            trends={trends}
            topViolators={topViolators}
            driftStats={driftStats}
            exceptionReport={exceptionReport}
            systemicWeaknesses={systemicWeaknesses}
            executiveReport={executiveReport}
            recommendations={recommendations}
            syncStatus={syncStatus}
            loading={loading}
            onSelectTrustMetric={(metricKey, label) => {
              setMetricDrilldown({
                metricKey,
                metricLabel: label,
                incidents: [],
                limitations: [],
              });
              handleOpenTab("events");
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

      {activeTab === "forensics" ? (
        <section className="analytics-grid">
          <ForensicsInsightsPanel
            state={forensicsState}
            delta={forensicsDelta}
            timeline={forensicsTimeline}
            flashback={forensicsVEXFlashback}
            replay={forensicsReplay}
            loading={loading}
            timestamp={forensicsTimestamp}
            onTimestampChange={setForensicsTimestamp}
          />
        </section>
      ) : null}

      {activeTab === "federation" ? (
        <section className="analytics-grid">
          <FederationInsightsPanel
            view={federationView}
            loading={loading}
            focusPeerID={focusTarget?.tab === "federation" && focusTarget.kind === "federation_peer" ? focusTarget.ref : null}
          />
        </section>
      ) : null}

      {activeTab === "validation" ? (
        <section className="analytics-grid">
          <ValidationHarnessPanel
            scenarios={validationScenarios}
            score={validationScore}
            runs={validationRuns}
            whatIf={validationWhatIf}
            loading={loading}
            focusRunID={focusTarget?.tab === "validation" && focusTarget.kind === "validation_run" ? focusTarget.ref : null}
            onRunHarness={handleRunValidationHarness}
            onRunWhatIf={handleRunValidationWhatIf}
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
          <RuntimeIntegrityPanel
            integrity={runtimeIntegrityStates}
            workloads={runtimeWorkloads}
            findings={runtimeIntegrityFindings}
            enforcement={runtimeEnforcement}
            loading={loading}
            focusSubjectRef={
              focusTarget?.tab === "runtime" && (focusTarget.kind === "runtime_subject" || focusTarget.kind === "runtime_finding" || focusTarget.kind === "hardening_execution")
                ? firstNonEmptyValue(focusTarget.secondary_ref, focusTarget.kind === "runtime_subject" ? focusTarget.ref : null)
                : null
            }
            focusFindingID={focusTarget?.tab === "runtime" && focusTarget.kind === "runtime_finding" ? focusTarget.ref : null}
          />
          <DefensePosturePanel
            posture={hardeningPosture}
            actions={hardeningActions}
            loading={loading}
            focusSubjectRef={
              focusTarget?.tab === "runtime" && (focusTarget.kind === "runtime_subject" || focusTarget.kind === "hardening_execution")
                ? firstNonEmptyValue(focusTarget.secondary_ref, focusTarget.kind === "runtime_subject" ? focusTarget.ref : null)
                : null
            }
            focusExecutionID={focusTarget?.tab === "runtime" && focusTarget.kind === "hardening_execution" ? focusTarget.ref : null}
          />
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
          refreshKey={refreshIndex}
          metricDrilldown={metricDrilldown}
          onClearMetricDrilldown={() => setMetricDrilldown(null)}
          onLoadExport={handleLoadIncidentExport}
          onLoadPackage={handleLoadIncidentPackage}
          onLoadExecutiveReport={handleLoadExecutiveReport}
          onSealHandoff={handleSealHandoff}
          onGetHandoffVerification={handleGetHandoffVerification}
          onDownloadHandoff={handleDownloadHandoff}
          onLoadIncidentDefenseGaps={handleLoadIncidentDefenseGaps}
          onLoadMetricDefenseGaps={handleLoadMetricDefenseGaps}
          onLoadIncidentPolicyReplay={handleLoadIncidentPolicyReplay}
          onLoadMetricPolicyReplay={handleLoadMetricPolicyReplay}
          onLoadIncidentBlastRadius={handleLoadIncidentBlastRadius}
          onLoadMetricBlastRadius={handleLoadMetricBlastRadius}
          onLoadRecommendations={handleLoadRecommendations}
          onAcknowledge={handleAcknowledgeIncident}
          onWatch={handleWatchIncident}
          onAssign={handleAssignIncident}
          onResolve={handleResolveIncident}
          onReopen={handleReopenIncident}
          onAddNote={handleIncidentNote}
          onAcknowledgeRecommendation={handleAcknowledgeRecommendation}
          onAcceptRecommendation={handleAcceptRecommendation}
          onRejectRecommendation={handleRejectRecommendation}
          onExecuteRecommendation={handleExecuteRecommendation}
          onVerifyRecommendation={handleVerifyRecommendation}
          onAssignRecommendation={handleAssignRecommendation}
          onCommentRecommendation={handleCommentRecommendation}
          onRequestRecommendationApproval={handleRequestRecommendationApproval}
          focusedIncidentID={focusTarget?.tab === "events" && focusTarget.kind === "incident" ? focusTarget.ref : null}
        />
      ) : null}

      {activeTab !== "analytics" && activeTab !== "topology" && activeTab !== "forensics" && activeTab !== "federation" && activeTab !== "runtime" && activeTab !== "exceptions" && activeTab !== "inventory" && activeTab !== "vulnerabilities" && activeTab !== "signing" && activeTab !== "scorecard" && activeTab !== "guidance" ? (
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
