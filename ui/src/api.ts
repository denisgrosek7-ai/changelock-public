import type {
  AdvisoryReadbackRef,
  AdvisoryReadbackResponse,
  AdvisoryShareGrant,
  DefenseGapAssessment,
  DecisionEvidenceEnvelope,
  ExecutiveDefenseReport,
  IncidentExport,
  IncidentPackage,
  InvestigationIncident,
  MetricIncidentDrilldown,
  PolicyReplayAssessment,
  Recommendation,
  RecommendationActionTemplate,
  SystemicWeaknessResponse,
} from "./incidents";
import type {
  ActiveWorkloadRef,
  AnalyticsAnomaliesResponse,
  AnalyticsAnomaly,
  AnalyticsComparisonContext,
  AnalyticsDeltaResponse,
  AnalyticsMetricDefinition,
  AnalyticsMetricTrend,
  AnalyticsScorecardCard,
  AnalyticsScorecardsResponse,
  AnalyticsSegmentCatalogItem,
  AnalyticsSegmentDelta,
  AnalyticsSegmentsResponse,
  AIInsightsResponse,
  AuditFinding,
  AuditHealth,
  AuditReport,
  AuthStatus,
  BreakGlassGuidance,
  DriftStatsResponse,
  EventFilters,
  EventsResponse,
  ExceptionActionResponse,
  ExceptionReport,
  ExceptionRequestInput,
  HandoffSealResponse,
  ExceptionsResponse,
  GuidanceGrouping,
  GuidanceItem,
  GuidanceResponse,
  GuidanceSummary,
  GuidanceVEXDraftSuggestion,
  SBOMComponent,
  SBOMComponentsResponse,
  SBOMDocument,
  SBOMImageResponse,
  PolicyException,
  PublishedTrustView,
  ReasonCount,
  PointInTimeState,
  TimeDeltaResult,
  ForensicTimelineResponse,
  VEXFlashbackResponse,
  ForensicReplayResponse,
  RuntimeDriftFinding,
  RuntimeActiveState,
  RuntimeActiveStatesResponse,
  RuntimeClosedLoopStatus,
  RuntimeDriftFindingsResponse,
  RuntimeDriftStatus,
  SigningIdentityFinding,
  SigningIdentityFindingsResponse,
  SigningIdentityObservation,
  SigningIdentityObservationsResponse,
  SigningIdentityPoliciesResponse,
  SigningIdentityPolicy,
  SigningIdentityStatus,
  StoredEvent,
  Summary,
  SyncStatus,
  TabKey,
  TrustBadge,
  TrustScoreMetric,
  TrustScorecard,
  TopViolator,
  TopViolatorsResponse,
  TopologyBlastRadiusResponse,
  TopologyContainmentOption,
  TopologyDeltaItem,
  TopologyDeltaResponse,
  TopologyEdge,
  TopologyGraphResponse,
  TopologyGraphSummary,
  TopologyGraphView,
  TopologyHeatmapResponse,
  TopologyNode,
  TopologyRiskPath,
  TopologyServicesResponse,
  TrendBucket,
  TrendsResponse,
  VEXCreateInput,
  VEXMatch,
  VEXStatement,
  VEXStatementActionResponse,
  VEXStatementsResponse,
  VEXStatusSummary,
  VulnerabilityBlastRadiusItem,
  VulnerabilityBlastRadiusResponse,
  VulnerabilityDecision,
  VulnerabilityDecisionInput,
  VulnerabilityDecisionsResponse,
  VulnerabilityFinding,
  VulnerabilityNetResponse,
  VulnerabilitiesResponse,
  VulnerabilityRescanResponse,
  VulnerabilityTimelineEntry,
  VulnerabilityTimelineResponse,
  VerificationResult,
  VerifierSummary,
  StandardsMapping,
  SealedManifest,
} from "./types";

type RuntimeConfig = {
  apiBaseUrl?: string;
  apiToken?: string;
  apiTimeoutMs?: string | number;
};

declare global {
  interface Window {
    __CHANGELOCK_CONFIG__?: RuntimeConfig;
  }
}

const runtimeConfig = window.__CHANGELOCK_CONFIG__ || {};
const API_BASE_URL = (runtimeConfig.apiBaseUrl || import.meta.env.VITE_API_BASE_URL || "/api").replace(/\/$/, "");
const API_TOKEN = String(runtimeConfig.apiToken || import.meta.env.VITE_API_TOKEN || "").trim();
const API_TIMEOUT_MS = Number.parseInt(String(runtimeConfig.apiTimeoutMs || import.meta.env.VITE_API_TIMEOUT_MS || "8000"), 10);

export class APIError extends Error {
  status: number;

  constructor(status: number, message: string) {
    super(message);
    this.name = "APIError";
    this.status = status;
  }
}

function buildURL(path: string, params?: Record<string, string | string[] | undefined>) {
  const url = new URL(`${API_BASE_URL}${path}`, window.location.origin);
  if (params) {
    for (const [key, value] of Object.entries(params)) {
      if (Array.isArray(value)) {
        value.filter(Boolean).forEach((item) => url.searchParams.append(key, item));
      } else if (value) {
        url.searchParams.set(key, value);
      }
    }
  }
  return url.toString();
}

async function fetchJSON<T>(
  path: string,
  options?: {
    method?: string;
    params?: Record<string, string | string[] | undefined>;
    body?: unknown;
  },
): Promise<T> {
  const controller = new AbortController();
  const timeoutID = window.setTimeout(() => controller.abort(), API_TIMEOUT_MS);

  try {
    const response = await fetch(buildURL(path, options?.params), {
      method: options?.method || "GET",
      headers: {
        Accept: "application/json",
        ...(options?.body !== undefined ? { "Content-Type": "application/json" } : {}),
        ...(API_TOKEN ? { Authorization: `Bearer ${API_TOKEN}` } : {}),
      },
      body: options?.body !== undefined ? JSON.stringify(options.body) : undefined,
      cache: "no-store",
      signal: controller.signal,
    });

    if (!response.ok) {
      const contentType = response.headers.get("content-type") || "";
      if (contentType.includes("application/json")) {
        const payload = (await response.json()) as { error?: string };
        throw new APIError(response.status, payload.error || `request failed with status ${response.status}`);
      }

      const payload = await response.text();
      throw new APIError(response.status, payload || `request failed with status ${response.status}`);
    }

    return response.json() as Promise<T>;
  } catch (error) {
    if (error instanceof DOMException && error.name === "AbortError") {
      throw new Error("Audit API request timed out.");
    }
    throw error;
  } finally {
    window.clearTimeout(timeoutID);
  }
}

async function fetchBlob(
  path: string,
  options?: {
    method?: string;
    params?: Record<string, string | string[] | undefined>;
    body?: unknown;
  },
): Promise<Blob> {
  const controller = new AbortController();
  const timeoutID = window.setTimeout(() => controller.abort(), API_TIMEOUT_MS);

  try {
    const response = await fetch(buildURL(path, options?.params), {
      method: options?.method || "GET",
      headers: {
        Accept: "application/octet-stream",
        ...(options?.body !== undefined ? { "Content-Type": "application/json" } : {}),
        ...(API_TOKEN ? { Authorization: `Bearer ${API_TOKEN}` } : {}),
      },
      body: options?.body !== undefined ? JSON.stringify(options.body) : undefined,
      cache: "no-store",
      signal: controller.signal,
    });

    if (!response.ok) {
      const contentType = response.headers.get("content-type") || "";
      if (contentType.includes("application/json")) {
        const payload = (await response.json()) as { error?: string };
        throw new APIError(response.status, payload.error || `request failed with status ${response.status}`);
      }

      const payload = await response.text();
      throw new APIError(response.status, payload || `request failed with status ${response.status}`);
    }

    return response.blob();
  } catch (error) {
    if (error instanceof DOMException && error.name === "AbortError") {
      throw new Error("Audit API request timed out.");
    }
    throw error;
  } finally {
    window.clearTimeout(timeoutID);
  }
}

function isRecord(value: unknown): value is Record<string, unknown> {
  return typeof value === "object" && value !== null && !Array.isArray(value);
}

function readString(value: unknown, field: string): string {
  if (typeof value !== "string") {
    throw new Error(`Audit API returned invalid ${field}.`);
  }
  return value;
}

function readOptionalString(value: unknown, field: string): string | undefined {
  if (value === undefined || value === null) {
    return undefined;
  }
  return readString(value, field);
}

function readNumber(value: unknown, field: string): number {
  if (typeof value !== "number" || Number.isNaN(value)) {
    throw new Error(`Audit API returned invalid ${field}.`);
  }
  return value;
}

function readBoolean(value: unknown, field: string): boolean {
  if (typeof value !== "boolean") {
    throw new Error(`Audit API returned invalid ${field}.`);
  }
  return value;
}

function readOptionalStringArray(value: unknown, field: string): string[] | undefined {
  if (value === undefined || value === null) {
    return undefined;
  }
  if (!Array.isArray(value) || !value.every((item) => typeof item === "string")) {
    throw new Error(`Audit API returned invalid ${field}.`);
  }
  return value;
}

function readOptionalArray(value: unknown, field: string): unknown[] {
  if (value === undefined || value === null) {
    return [];
  }
  if (!Array.isArray(value)) {
    throw new Error(`Audit API returned invalid ${field}.`);
  }
  return value;
}

function readOptionalRecord(value: unknown, field: string): Record<string, unknown> | undefined {
  if (value === undefined || value === null) {
    return undefined;
  }
  if (!isRecord(value)) {
    throw new Error(`Audit API returned invalid ${field}.`);
  }
  return value;
}

function parseReasonCount(value: unknown): ReasonCount {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid top_deny_reasons item.");
  }
  return {
    reason: readString(value.reason, "reason"),
    count: readNumber(value.count, "count"),
  };
}

function parseVerifierSummary(value: unknown): VerifierSummary | undefined {
  if (value === undefined || value === null) {
    return undefined;
  }
  if (!isRecord(value) || typeof value.signature_valid !== "boolean" || typeof value.attestation_valid !== "boolean") {
    throw new Error("Audit API returned invalid verifier_summary.");
  }
  return {
    signature_valid: value.signature_valid,
    attestation_valid: value.attestation_valid,
  };
}

function parseSummary(value: unknown): Summary {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid summary payload.");
  }
  const counts = readOptionalRecord(value.counts_by_event_type, "counts_by_event_type") || {};
  const countsByEventType: Record<string, number> = {};
  for (const [key, count] of Object.entries(counts)) {
    countsByEventType[key] = readNumber(count, `counts_by_event_type.${key}`);
  }

  const denyReasons = value.top_deny_reasons;
  if (!Array.isArray(denyReasons)) {
    throw new Error("Audit API returned invalid top_deny_reasons.");
  }

  return {
    total_events: readNumber(value.total_events, "total_events"),
    total_allow: readNumber(value.total_allow, "total_allow"),
    total_deny: readNumber(value.total_deny, "total_deny"),
    total_error: readNumber(value.total_error, "total_error"),
    counts_by_event_type: countsByEventType,
    top_deny_reasons: denyReasons.map(parseReasonCount),
    recent_runtime_drift_deny: readNumber(value.recent_runtime_drift_deny, "recent_runtime_drift_deny"),
  };
}

function parseStoredEvent(value: unknown): StoredEvent {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid event payload.");
  }

  const decision = readString(value.decision, "events[].decision");
  if (decision !== "ALLOW" && decision !== "DENY" && decision !== "ERROR") {
    throw new Error("Audit API returned invalid events[].decision.");
  }

  return {
    id: readNumber(value.id, "events[].id"),
    received_at: readString(value.received_at, "events[].received_at"),
    request_id: readOptionalString(value.request_id, "events[].request_id"),
    timestamp: readOptionalString(value.timestamp, "events[].timestamp"),
    component: readString(value.component, "events[].component"),
    event_type: readString(value.event_type, "events[].event_type"),
    actor: readOptionalString(value.actor, "events[].actor"),
    cluster_id: readOptionalString(value.cluster_id, "events[].cluster_id"),
    tenant_id: readOptionalString(value.tenant_id, "events[].tenant_id"),
    repo: readOptionalString(value.repo, "events[].repo"),
    branch: readOptionalString(value.branch, "events[].branch"),
    environment: readOptionalString(value.environment, "events[].environment"),
    namespace: readOptionalString(value.namespace, "events[].namespace"),
    workload: readOptionalString(value.workload, "events[].workload"),
    image: readOptionalString(value.image, "events[].image"),
    digest: readOptionalString(value.digest, "events[].digest"),
    cve_id: readOptionalString(value.cve_id, "events[].cve_id"),
    decision,
    reasons: readOptionalStringArray(value.reasons, "events[].reasons"),
    drift_result: readOptionalString(value.drift_result, "events[].drift_result"),
    drift_classes: readOptionalStringArray(value.drift_classes, "events[].drift_classes"),
    verifier_summary: parseVerifierSummary(value.verifier_summary),
    policy_version: readOptionalString(value.policy_version, "events[].policy_version"),
    policy_bundle_id: readOptionalString(value.policy_bundle_id, "events[].policy_bundle_id"),
    policy_bundle_hash: readOptionalString(value.policy_bundle_hash, "events[].policy_bundle_hash"),
    decision_hash: readOptionalString(value.decision_hash, "events[].decision_hash"),
    is_exception: typeof value.is_exception === "boolean" ? value.is_exception : undefined,
    exception_id: readOptionalString(value.exception_id, "events[].exception_id"),
    exception_type: readOptionalString(value.exception_type, "events[].exception_type"),
    exception_status: readOptionalString(value.exception_status, "events[].exception_status") as StoredEvent["exception_status"],
    exception_reason: readOptionalString(value.exception_reason, "events[].exception_reason"),
    exception_ticket_id: readOptionalString(value.exception_ticket_id, "events[].exception_ticket_id"),
    exception_requested_by: readOptionalString(value.exception_requested_by, "events[].exception_requested_by"),
    exception_requested_at: readOptionalString(value.exception_requested_at, "events[].exception_requested_at"),
    exception_approved_by: readOptionalString(value.exception_approved_by, "events[].exception_approved_by"),
    exception_approved_at: readOptionalString(value.exception_approved_at, "events[].exception_approved_at"),
    exception_rejected_by: readOptionalString(value.exception_rejected_by, "events[].exception_rejected_by"),
    exception_rejected_at: readOptionalString(value.exception_rejected_at, "events[].exception_rejected_at"),
    exception_rejection_reason: readOptionalString(value.exception_rejection_reason, "events[].exception_rejection_reason"),
    exception_expires_at: readOptionalString(value.exception_expires_at, "events[].exception_expires_at"),
    evidence: readOptionalRecord(value.evidence, "events[].evidence"),
    raw_event: readOptionalRecord(value.raw_event, "events[].raw_event"),
  };
}

function parseEventsResponse(value: unknown): EventsResponse {
  if (!isRecord(value) || !Array.isArray(value.events)) {
    throw new Error("Audit API returned invalid events response.");
  }
  return {
    events: value.events.map(parseStoredEvent),
  };
}

function parsePolicyException(value: unknown): PolicyException {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid exception payload.");
  }
  return {
    id: readNumber(value.id, "exceptions[].id"),
    exception_id: readString(value.exception_id, "exceptions[].exception_id"),
    exception_type: readString(value.exception_type, "exceptions[].exception_type"),
    status: readString(value.status, "exceptions[].status") as PolicyException["status"],
    tenant_id: readOptionalString(value.tenant_id, "exceptions[].tenant_id"),
    environment: readOptionalString(value.environment, "exceptions[].environment"),
    namespace: readOptionalString(value.namespace, "exceptions[].namespace"),
    repo: readOptionalString(value.repo, "exceptions[].repo"),
    image_digest: readOptionalString(value.image_digest, "exceptions[].image_digest"),
    cve_id: readOptionalString(value.cve_id, "exceptions[].cve_id"),
    reason: readString(value.reason, "exceptions[].reason"),
    ticket_id: readString(value.ticket_id, "exceptions[].ticket_id"),
    requested_by: readOptionalString(value.requested_by, "exceptions[].requested_by"),
    requested_at: readOptionalString(value.requested_at, "exceptions[].requested_at"),
    approved_by: readOptionalString(value.approved_by, "exceptions[].approved_by"),
    approved_at: readOptionalString(value.approved_at, "exceptions[].approved_at"),
    rejected_by: readOptionalString(value.rejected_by, "exceptions[].rejected_by"),
    rejected_at: readOptionalString(value.rejected_at, "exceptions[].rejected_at"),
    rejection_reason: readOptionalString(value.rejection_reason, "exceptions[].rejection_reason"),
    created_at: readString(value.created_at, "exceptions[].created_at"),
    expires_at: readString(value.expires_at, "exceptions[].expires_at"),
    active: readBoolean(value.active, "exceptions[].active"),
    last_updated_at: readOptionalString(value.last_updated_at, "exceptions[].last_updated_at"),
    metadata: readOptionalRecord(value.metadata, "exceptions[].metadata"),
  };
}

function parseExceptionsResponse(value: unknown): ExceptionsResponse {
  if (!isRecord(value) || !Array.isArray(value.exceptions)) {
    throw new Error("Audit API returned invalid exceptions response.");
  }
  return { exceptions: value.exceptions.map(parsePolicyException) };
}

function parseExceptionActionResponse(value: unknown): ExceptionActionResponse {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid exception action response.");
  }
  return {
    status: readString(value.status, "status"),
    exception: parsePolicyException(value.exception),
  };
}

function parseExceptionReport(value: unknown): ExceptionReport {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid exception report.");
  }
  const parseExceptionArray = (field: string) => {
    const raw = value[field];
    if (raw === undefined || raw === null) {
      return undefined;
    }
    if (!Array.isArray(raw)) {
      throw new Error(`Audit API returned invalid ${field}.`);
    }
    return raw.map(parsePolicyException);
  };
  const recentUsed = value.recent_used;
  if (!Array.isArray(recentUsed)) {
    throw new Error("Audit API returned invalid recent_used.");
  }
  const recentInactiveRaw = value.recent_inactive;
  if (!Array.isArray(recentInactiveRaw)) {
    throw new Error("Audit API returned invalid recent_inactive.");
  }

  const statusCountsRaw = readOptionalRecord(value.status_counts, "status_counts") || {};
  const statusCounts: Record<string, number> = {};
  for (const [key, count] of Object.entries(statusCountsRaw)) {
    statusCounts[key] = readNumber(count, `status_counts.${key}`);
  }

  return {
    active: parseExceptionArray("active") || [],
    pending: parseExceptionArray("pending"),
    rejected: parseExceptionArray("rejected"),
    revoked: parseExceptionArray("revoked"),
    expired: parseExceptionArray("expired"),
    recent_used: recentUsed.map(parseStoredEvent),
    recent_inactive: recentInactiveRaw.map(parsePolicyException),
    status_counts: statusCounts,
  };
}

function parseTrendBucket(value: unknown): TrendBucket {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid trends bucket.");
  }
  return {
    timestamp: readString(value.timestamp, "buckets[].timestamp"),
    allow_count: readNumber(value.allow_count, "buckets[].allow_count"),
    deny_count: readNumber(value.deny_count, "buckets[].deny_count"),
    error_count: readNumber(value.error_count, "buckets[].error_count"),
  };
}

function parseTrendsResponse(value: unknown): TrendsResponse {
  if (!isRecord(value) || !Array.isArray(value.buckets)) {
    throw new Error("Audit API returned invalid trends response.");
  }
  const totalsRaw = readOptionalRecord(value.totals, "totals") || {};
  const totals: Record<string, number> = {};
  for (const [key, count] of Object.entries(totalsRaw)) {
    totals[key] = readNumber(count, `totals.${key}`);
  }
  return {
    buckets: value.buckets.map(parseTrendBucket),
    totals,
    applied_filters: readOptionalRecord(value.applied_filters, "applied_filters") as Record<string, string> || {},
    metric_trends: readOptionalArray(value.metric_trends, "metric_trends").map(parseAnalyticsMetricTrend),
    comparison: value.comparison ? parseAnalyticsComparisonContext(value.comparison) : undefined,
    limitations: readOptionalStringArray(value.limitations, "limitations"),
  };
}

function parseAnalyticsComparisonContext(value: unknown): AnalyticsComparisonContext {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid analytics comparison context.");
  }
  return {
    window: readString(value.window, "comparison.window"),
    compare_to: readString(value.compare_to, "comparison.compare_to"),
    group_by: readString(value.group_by, "comparison.group_by"),
    current_start: readString(value.current_start, "comparison.current_start"),
    current_end: readString(value.current_end, "comparison.current_end"),
    baseline_start: readString(value.baseline_start, "comparison.baseline_start"),
    baseline_end: readString(value.baseline_end, "comparison.baseline_end"),
    applied_filters: (readOptionalRecord(value.applied_filters, "comparison.applied_filters") as Record<string, string>) || {},
  };
}

function parseAnalyticsMetricDefinition(value: unknown): AnalyticsMetricDefinition {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid analytics metric definition.");
  }
  return {
    key: readString(value.key, "definition.key"),
    label: readString(value.label, "definition.label"),
    metric_class: readString(value.metric_class, "definition.metric_class"),
    description: readString(value.description, "definition.description"),
    formula: readString(value.formula, "definition.formula"),
    grain: readString(value.grain, "definition.grain"),
    default_window: readString(value.default_window, "definition.default_window"),
    segments: readOptionalStringArray(value.segments, "definition.segments"),
    exclusions: readOptionalStringArray(value.exclusions, "definition.exclusions"),
    owner: readString(value.owner, "definition.owner"),
    interpretation: readString(value.interpretation, "definition.interpretation"),
  };
}

function parseAnalyticsSegmentDelta(value: unknown): AnalyticsSegmentDelta {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid analytics segment delta.");
  }
  return {
    segment_key: readString(value.segment_key, "segments[].segment_key"),
    segment_label: readString(value.segment_label, "segments[].segment_label"),
    current_value: readNumber(value.current_value, "segments[].current_value"),
    baseline_value: readNumber(value.baseline_value, "segments[].baseline_value"),
    delta_value: readNumber(value.delta_value, "segments[].delta_value"),
    delta_percent: readNumber(value.delta_percent, "segments[].delta_percent"),
    direction: readString(value.direction, "segments[].direction"),
  };
}

function parseAnalyticsMetricTrend(value: unknown): AnalyticsMetricTrend {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid analytics metric trend.");
  }
  return {
    definition: parseAnalyticsMetricDefinition(value.definition),
    current_value: readNumber(value.current_value, "metric_trends[].current_value"),
    baseline_value: readNumber(value.baseline_value, "metric_trends[].baseline_value"),
    delta_value: readNumber(value.delta_value, "metric_trends[].delta_value"),
    delta_percent: readNumber(value.delta_percent, "metric_trends[].delta_percent"),
    direction: readString(value.direction, "metric_trends[].direction"),
    velocity: readString(value.velocity, "metric_trends[].velocity"),
    summary: readString(value.summary, "metric_trends[].summary"),
    segment_highlights: readOptionalArray(value.segment_highlights, "metric_trends[].segment_highlights").map(parseAnalyticsSegmentDelta),
    limitations: readOptionalStringArray(value.limitations, "metric_trends[].limitations"),
  };
}

function parseTopViolator(value: unknown): TopViolator {
  if (!isRecord(value) || !Array.isArray(value.top_reasons)) {
    throw new Error("Audit API returned invalid top-violator item.");
  }
  return {
    key: readString(value.key, "items[].key"),
    deny_count: readNumber(value.deny_count, "items[].deny_count"),
    top_reasons: value.top_reasons.map(parseReasonCount),
  };
}

function parseTopViolatorsResponse(value: unknown): TopViolatorsResponse {
  if (!isRecord(value) || !Array.isArray(value.items)) {
    throw new Error("Audit API returned invalid top-violators response.");
  }
  return {
    items: value.items.map(parseTopViolator),
    applied_filters: readOptionalRecord(value.applied_filters, "applied_filters") as Record<string, string> || {},
  };
}

function parseDriftStatsResponse(value: unknown): DriftStatsResponse {
  if (!isRecord(value) || !Array.isArray(value.top_drifted_workloads)) {
    throw new Error("Audit API returned invalid drift stats response.");
  }
  const countsRaw = readOptionalRecord(value.counts_by_drift_class, "counts_by_drift_class") || {};
  const countsByClass: Record<string, number> = {};
  for (const [key, count] of Object.entries(countsRaw)) {
    countsByClass[key] = readNumber(count, `counts_by_drift_class.${key}`);
  }
  return {
    total_runtime_drift_denies: readNumber(value.total_runtime_drift_denies, "total_runtime_drift_denies"),
    counts_by_drift_class: countsByClass,
    top_drifted_workloads: value.top_drifted_workloads.map((item) => {
      if (!isRecord(item)) {
        throw new Error("Audit API returned invalid top_drifted_workloads item.");
      }
      return {
        workload: readString(item.workload, "top_drifted_workloads[].workload"),
        namespace: readOptionalString(item.namespace, "top_drifted_workloads[].namespace"),
        tenant_id: readOptionalString(item.tenant_id, "top_drifted_workloads[].tenant_id"),
        environment: readOptionalString(item.environment, "top_drifted_workloads[].environment"),
        count: readNumber(item.count, "top_drifted_workloads[].count"),
      };
    }),
    mean_time_to_resolve_seconds:
      value.mean_time_to_resolve_seconds === undefined || value.mean_time_to_resolve_seconds === null
        ? null
        : readNumber(value.mean_time_to_resolve_seconds, "mean_time_to_resolve_seconds"),
    applied_filters: readOptionalRecord(value.applied_filters, "applied_filters") as Record<string, string> || {},
  };
}

function parseAnalyticsDeltaResponse(value: unknown): AnalyticsDeltaResponse {
  if (!isRecord(value) || !Array.isArray(value.segments)) {
    throw new Error("Audit API returned invalid analytics delta response.");
  }
  return {
    definition: parseAnalyticsMetricDefinition(value.definition),
    comparison: parseAnalyticsComparisonContext(value.comparison),
    segments: value.segments.map(parseAnalyticsSegmentDelta),
    summary: readString(value.summary, "summary"),
    limitations: readOptionalStringArray(value.limitations, "limitations"),
  };
}

function parseAnalyticsAnomaly(value: unknown): AnalyticsAnomaly {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid analytics anomaly.");
  }
  return {
    type: readString(value.type, "items[].type"),
    title: readString(value.title, "items[].title"),
    metric_key: readString(value.metric_key, "items[].metric_key"),
    reason: readString(value.reason, "items[].reason"),
    baseline: readString(value.baseline, "items[].baseline"),
    deviation: readString(value.deviation, "items[].deviation"),
    segment: readString(value.segment, "items[].segment"),
    severity: readString(value.severity, "items[].severity"),
    recommended_next_step: readString(value.recommended_next_step, "items[].recommended_next_step"),
    evidence_refs: readOptionalStringArray(value.evidence_refs, "items[].evidence_refs"),
    limitations: readOptionalStringArray(value.limitations, "items[].limitations"),
  };
}

function parseAnalyticsAnomaliesResponse(value: unknown): AnalyticsAnomaliesResponse {
  if (!isRecord(value) || !Array.isArray(value.items)) {
    throw new Error("Audit API returned invalid analytics anomalies response.");
  }
  return {
    comparison: parseAnalyticsComparisonContext(value.comparison),
    items: value.items.map(parseAnalyticsAnomaly),
    limitations: readOptionalStringArray(value.limitations, "limitations"),
  };
}

function parseAnalyticsScorecardCard(value: unknown): AnalyticsScorecardCard {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid analytics scorecard card.");
  }
  return {
    definition: parseAnalyticsMetricDefinition(value.definition),
    status: readString(value.status, "cards[].status"),
    current_value: readNumber(value.current_value, "cards[].current_value"),
    baseline_value: readNumber(value.baseline_value, "cards[].baseline_value"),
    delta_value: readNumber(value.delta_value, "cards[].delta_value"),
    delta_percent: readNumber(value.delta_percent, "cards[].delta_percent"),
    direction: readString(value.direction, "cards[].direction"),
    summary: readString(value.summary, "cards[].summary"),
  };
}

function parseAnalyticsScorecardsResponse(value: unknown): AnalyticsScorecardsResponse {
  if (!isRecord(value) || !Array.isArray(value.cards)) {
    throw new Error("Audit API returned invalid analytics scorecards response.");
  }
  return {
    comparison: parseAnalyticsComparisonContext(value.comparison),
    cards: value.cards.map(parseAnalyticsScorecardCard),
    limitations: readOptionalStringArray(value.limitations, "limitations"),
  };
}

function parseAnalyticsSegmentCatalogItem(value: unknown): AnalyticsSegmentCatalogItem {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid analytics segment catalog item.");
  }
  return {
    group: readString(value.group, "items[].group"),
    values: readOptionalStringArray(value.values, "items[].values") || [],
  };
}

function parseAnalyticsSegmentsResponse(value: unknown): AnalyticsSegmentsResponse {
  if (!isRecord(value) || !Array.isArray(value.items)) {
    throw new Error("Audit API returned invalid analytics segments response.");
  }
  return {
    comparison: parseAnalyticsComparisonContext(value.comparison),
    items: value.items.map(parseAnalyticsSegmentCatalogItem),
    limitations: readOptionalStringArray(value.limitations, "limitations"),
  };
}

function parseTopologyNode(value: unknown): TopologyNode {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid topology node.");
  }
  return {
    node_id: readString(value.node_id, "topology.node_id"),
    service: readString(value.service, "topology.service"),
    workload: readOptionalString(value.workload, "topology.workload"),
    namespace: readOptionalString(value.namespace, "topology.namespace"),
    cluster: readOptionalString(value.cluster, "topology.cluster"),
    environment: readOptionalString(value.environment, "topology.environment"),
    team: readOptionalString(value.team, "topology.team"),
    repo: readOptionalString(value.repo, "topology.repo"),
    artifact_digest: readOptionalString(value.artifact_digest, "topology.artifact_digest"),
    public_exposure: readBoolean(value.public_exposure, "topology.public_exposure"),
    sensitivity_class: readString(value.sensitivity_class, "topology.sensitivity_class"),
    node_risk_score: readNumber(value.node_risk_score, "topology.node_risk_score"),
    blast_radius_score: readNumber(value.blast_radius_score, "topology.blast_radius_score"),
    critical_reach_count: readNumber(value.critical_reach_count, "topology.critical_reach_count"),
    public_entry_flag: readBoolean(value.public_entry_flag, "topology.public_entry_flag"),
    sensitive_asset_reach_flag: readBoolean(value.sensitive_asset_reach_flag, "topology.sensitive_asset_reach_flag"),
    propagation_class: readString(value.propagation_class, "topology.propagation_class"),
    trust_boundary_crossings: readNumber(value.trust_boundary_crossings, "topology.trust_boundary_crossings"),
    last_seen: readString(value.last_seen, "topology.last_seen"),
    evidence_refs: readOptionalStringArray(value.evidence_refs, "topology.evidence_refs"),
  };
}

function parseTopologyEdge(value: unknown): TopologyEdge {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid topology edge.");
  }
  return {
    source: readString(value.source, "topology.edges[].source"),
    target: readString(value.target, "topology.edges[].target"),
    edge_type: readString(value.edge_type, "topology.edges[].edge_type"),
    connectivity_class: readString(value.connectivity_class, "topology.edges[].connectivity_class"),
    evidence_source: readString(value.evidence_source, "topology.edges[].evidence_source"),
    confidence: readString(value.confidence, "topology.edges[].confidence"),
    last_seen: readOptionalString(value.last_seen, "topology.edges[].last_seen"),
    environment_scope: readOptionalString(value.environment_scope, "topology.edges[].environment_scope"),
    evidence_refs: readOptionalStringArray(value.evidence_refs, "topology.edges[].evidence_refs"),
  };
}

function parseTopologyGraphView(value: unknown): TopologyGraphView {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid topology graph view.");
  }
  return {
    nodes: readOptionalArray(value.nodes, "topology.graph.nodes").map(parseTopologyNode),
    edges: readOptionalArray(value.edges, "topology.graph.edges").map(parseTopologyEdge),
  };
}

function parseTopologyGraphSummary(value: unknown): TopologyGraphSummary {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid topology graph summary.");
  }
  return {
    declared_nodes: readNumber(value.declared_nodes, "topology.summary.declared_nodes"),
    declared_edges: readNumber(value.declared_edges, "topology.summary.declared_edges"),
    observed_nodes: readNumber(value.observed_nodes, "topology.summary.observed_nodes"),
    observed_edges: readNumber(value.observed_edges, "topology.summary.observed_edges"),
    effective_nodes: readNumber(value.effective_nodes, "topology.summary.effective_nodes"),
    effective_edges: readNumber(value.effective_edges, "topology.summary.effective_edges"),
    public_entry_nodes: readNumber(value.public_entry_nodes, "topology.summary.public_entry_nodes"),
    critical_nodes: readNumber(value.critical_nodes, "topology.summary.critical_nodes"),
    high_blast_radius: readNumber(value.high_blast_radius, "topology.summary.high_blast_radius"),
  };
}

function parseTopologyServicesResponse(value: unknown): TopologyServicesResponse {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid topology services response.");
  }
  return {
    items: readOptionalArray(value.items, "topology.items").map(parseTopologyNode),
    applied_filters: (readOptionalRecord(value.applied_filters, "topology.applied_filters") as Record<string, string>) || {},
    limitations: readOptionalStringArray(value.limitations, "topology.limitations"),
  };
}

function parseTopologyHeatmapResponse(value: unknown): TopologyHeatmapResponse {
  return parseTopologyServicesResponse(value);
}

function parseTopologyRiskPath(value: unknown): TopologyRiskPath {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid topology risk path.");
  }
  return {
    nodes: readOptionalStringArray(value.nodes, "topology.top_risk_paths[].nodes") || [],
    edge_types: readOptionalStringArray(value.edge_types, "topology.top_risk_paths[].edge_types") || [],
    summary: readString(value.summary, "topology.top_risk_paths[].summary"),
  };
}

function parseTopologyContainmentOption(value: unknown): TopologyContainmentOption {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid topology containment option.");
  }
  return {
    option_id: readString(value.option_id, "topology.containment_options[].option_id"),
    title: readString(value.title, "topology.containment_options[].title"),
    summary: readString(value.summary, "topology.containment_options[].summary"),
    restriction_plan: readOptionalStringArray(value.restriction_plan, "topology.containment_options[].restriction_plan") || [],
    closed_edge_types: readOptionalStringArray(value.closed_edge_types, "topology.containment_options[].closed_edge_types") || [],
    estimated_score_reduction: readNumber(value.estimated_score_reduction, "topology.containment_options[].estimated_score_reduction"),
    approval_mode: readString(value.approval_mode, "topology.containment_options[].approval_mode"),
    evidence_refs: readOptionalStringArray(value.evidence_refs, "topology.containment_options[].evidence_refs"),
  };
}

function parseTopologyBlastRadiusResponse(value: unknown): TopologyBlastRadiusResponse {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid topology blast radius response.");
  }
  return {
    subject_ref: readString(value.subject_ref, "topology.subject_ref"),
    subject_type: readString(value.subject_type, "topology.subject_type"),
    affected_nodes: readOptionalArray(value.affected_nodes, "topology.affected_nodes").map(parseTopologyNode),
    primary_affected_node: value.primary_affected_node ? parseTopologyNode(value.primary_affected_node) : undefined,
    reachable_nodes: readOptionalArray(value.reachable_nodes, "topology.reachable_nodes").map(parseTopologyNode),
    critical_reach_count: readNumber(value.critical_reach_count, "topology.critical_reach_count"),
    blast_radius_score: readNumber(value.blast_radius_score, "topology.blast_radius_score"),
    trust_boundary_crossings: readNumber(value.trust_boundary_crossings, "topology.trust_boundary_crossings"),
    declared_edge_count: readNumber(value.declared_edge_count, "topology.declared_edge_count"),
    observed_edge_count: readNumber(value.observed_edge_count, "topology.observed_edge_count"),
    top_risk_paths: readOptionalArray(value.top_risk_paths, "topology.top_risk_paths").map(parseTopologyRiskPath),
    containment_options: readOptionalArray(value.containment_options, "topology.containment_options").map(parseTopologyContainmentOption),
    evidence_refs: readOptionalStringArray(value.evidence_refs, "topology.evidence_refs"),
    limitations: readOptionalStringArray(value.limitations, "topology.limitations"),
  };
}

function parseTopologyDeltaItem(value: unknown): TopologyDeltaItem {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid topology delta item.");
  }
  return {
    node_id: readString(value.node_id, "topology.delta.items[].node_id"),
    service: readString(value.service, "topology.delta.items[].service"),
    current_blast_radius_score: readNumber(value.current_blast_radius_score, "topology.delta.items[].current_blast_radius_score"),
    baseline_blast_radius_score: readNumber(value.baseline_blast_radius_score, "topology.delta.items[].baseline_blast_radius_score"),
    delta: readNumber(value.delta, "topology.delta.items[].delta"),
    edge_additions: readNumber(value.edge_additions, "topology.delta.items[].edge_additions"),
    critical_reach_delta: readNumber(value.critical_reach_delta, "topology.delta.items[].critical_reach_delta"),
    drift_signals: readOptionalStringArray(value.drift_signals, "topology.delta.items[].drift_signals"),
  };
}

function parseTopologyDeltaResponse(value: unknown): TopologyDeltaResponse {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid topology delta response.");
  }
  return {
    comparison: parseAnalyticsComparisonContext(value.comparison),
    items: readOptionalArray(value.items, "topology.delta.items").map(parseTopologyDeltaItem),
    limitations: readOptionalStringArray(value.limitations, "topology.delta.limitations"),
  };
}

function parseTopologyGraphResponse(value: unknown): TopologyGraphResponse {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid topology graph response.");
  }
  return {
    declared_graph: parseTopologyGraphView(value.declared_graph),
    observed_graph: parseTopologyGraphView(value.observed_graph),
    effective_graph: parseTopologyGraphView(value.effective_graph),
    summary: parseTopologyGraphSummary(value.summary),
    applied_filters: (readOptionalRecord(value.applied_filters, "topology.applied_filters") as Record<string, string>) || {},
    limitations: readOptionalStringArray(value.limitations, "topology.limitations"),
  };
}

function parseSealedManifestScope(value: unknown): SealedManifest["scope"] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid handoff manifest scope.");
  }
  return {
    audience: readString(value.audience, "handoff.scope.audience"),
    selection_mode: readString(value.selection_mode, "handoff.scope.selection_mode"),
    selection_summary: readString(value.selection_summary, "handoff.scope.selection_summary"),
    incident_count: readNumber(value.incident_count, "handoff.scope.incident_count"),
    incident_refs: readOptionalStringArray(value.incident_refs, "handoff.scope.incident_refs") || [],
    tenant_id: readOptionalString(value.tenant_id, "handoff.scope.tenant_id"),
    environment: readOptionalString(value.environment, "handoff.scope.environment"),
    repo: readOptionalString(value.repo, "handoff.scope.repo"),
  };
}

function parseSealedManifestRedaction(value: unknown): SealedManifest["redaction_profile"] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid handoff redaction profile.");
  }
  return {
    audience: readString(value.audience, "handoff.redaction_profile.audience"),
    profile_version: readString(value.profile_version, "handoff.redaction_profile.profile_version"),
    summary: readOptionalStringArray(value.summary, "handoff.redaction_profile.summary") || [],
  };
}

function parseSealedManifestArtifact(value: unknown): SealedManifest["artifacts"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid handoff artifact.");
  }
  return {
    path: readString(value.path, "handoff.artifact.path"),
    media_type: readString(value.media_type, "handoff.artifact.media_type"),
    sha256: readString(value.sha256, "handoff.artifact.sha256"),
    role: readString(value.role, "handoff.artifact.role"),
    advisory_only: value.advisory_only === undefined ? undefined : readBoolean(value.advisory_only, "handoff.artifact.advisory_only"),
  };
}

function parseSealedManifestReadbackRef(value: unknown): NonNullable<SealedManifest["readback_refs"]>[number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid handoff readback ref.");
  }
  return {
    resource_type: readString(value.resource_type, "handoff.readback_ref.resource_type"),
    resource_id: readOptionalString(value.resource_id, "handoff.readback_ref.resource_id"),
    evidence_hash: readString(value.evidence_hash, "handoff.readback_ref.evidence_hash"),
    resource_uri: readOptionalString(value.resource_uri, "handoff.readback_ref.resource_uri"),
  };
}

function parseSealedManifestForensicRef(value: unknown): NonNullable<SealedManifest["forensic_refs"]>[number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid handoff forensic ref.");
  }
  return {
    context_uri: readOptionalString(value.context_uri, "handoff.forensic_ref.context_uri"),
    context_type: readString(value.context_type, "handoff.forensic_ref.context_type"),
    timestamp: readString(value.timestamp, "handoff.forensic_ref.timestamp"),
    advisory_only: readBoolean(value.advisory_only, "handoff.forensic_ref.advisory_only"),
    counterfactual: value.counterfactual === undefined ? undefined : readBoolean(value.counterfactual, "handoff.forensic_ref.counterfactual"),
  };
}

function parseSealedManifest(value: unknown): SealedManifest {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid sealed manifest.");
  }
  return {
    package_id: readString(value.package_id, "handoff.manifest.package_id"),
    package_type: readString(value.package_type, "handoff.manifest.package_type"),
    schema_version: readString(value.schema_version, "handoff.manifest.schema_version"),
    created_at: readString(value.created_at, "handoff.manifest.created_at"),
    generator_identity: readString(value.generator_identity, "handoff.manifest.generator_identity"),
    scope: parseSealedManifestScope(value.scope),
    redaction_profile: parseSealedManifestRedaction(value.redaction_profile),
    artifacts: readOptionalArray(value.artifacts, "handoff.manifest.artifacts").map(parseSealedManifestArtifact),
    evidence_refs: readOptionalStringArray(value.evidence_refs, "handoff.manifest.evidence_refs") || [],
    readback_refs: readOptionalArray(value.readback_refs, "handoff.manifest.readback_refs").map(parseSealedManifestReadbackRef),
    forensic_refs: readOptionalArray(value.forensic_refs, "handoff.manifest.forensic_refs").map(parseSealedManifestForensicRef),
    root_hash: readString(value.root_hash, "handoff.manifest.root_hash"),
    limitations: readOptionalStringArray(value.limitations, "handoff.manifest.limitations"),
  };
}

function parseHandoffSessionRecord(value: unknown): HandoffSealResponse["session"] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid handoff session.");
  }
  return {
    session_id: readString(value.session_id, "handoff.session.session_id"),
    package_id: readString(value.package_id, "handoff.session.package_id"),
    package_type: readString(value.package_type, "handoff.session.package_type"),
    scope_summary: readString(value.scope_summary, "handoff.session.scope_summary"),
    initiated_by: readString(value.initiated_by, "handoff.session.initiated_by"),
    initiated_at: readString(value.initiated_at, "handoff.session.initiated_at"),
    sign_mode: readString(value.sign_mode, "handoff.session.sign_mode"),
    co_sign_mode: readString(value.co_sign_mode, "handoff.session.co_sign_mode"),
    status: readString(value.status, "handoff.session.status"),
    final_bundle_ref: readString(value.final_bundle_ref, "handoff.session.final_bundle_ref"),
    manifest_hash: readString(value.manifest_hash, "handoff.session.manifest_hash"),
  };
}

function parseSealedBundleMetadata(value: unknown): HandoffSealResponse["bundle"] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid handoff bundle metadata.");
  }
  return {
    package_id: readString(value.package_id, "handoff.bundle.package_id"),
    bundle_path: readString(value.bundle_path, "handoff.bundle.bundle_path"),
    manifest_hash: readString(value.manifest_hash, "handoff.bundle.manifest_hash"),
    seal_status: readString(value.seal_status, "handoff.bundle.seal_status"),
    signature_count: readNumber(value.signature_count, "handoff.bundle.signature_count"),
    timestamp_status: readString(value.timestamp_status, "handoff.bundle.timestamp_status"),
    transparency_status: readString(value.transparency_status, "handoff.bundle.transparency_status"),
    verification_uri: readString(value.verification_uri, "handoff.bundle.verification_uri"),
    offline_verifier_present: readBoolean(value.offline_verifier_present, "handoff.bundle.offline_verifier_present"),
  };
}

function parseVerificationResult(value: unknown): VerificationResult {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid handoff verification.");
  }
  return {
    package_id: readString(value.package_id, "handoff.verification.package_id"),
    manifest_valid: readBoolean(value.manifest_valid, "handoff.verification.manifest_valid"),
    artifact_hashes_valid: readBoolean(value.artifact_hashes_valid, "handoff.verification.artifact_hashes_valid"),
    signatures_valid: readBoolean(value.signatures_valid, "handoff.verification.signatures_valid"),
    timestamp_valid: readBoolean(value.timestamp_valid, "handoff.verification.timestamp_valid"),
    transparency_valid: readBoolean(value.transparency_valid, "handoff.verification.transparency_valid"),
    signer_identities: readOptionalStringArray(value.signer_identities, "handoff.verification.signer_identities") || [],
    redaction_profile: readString(value.redaction_profile, "handoff.verification.redaction_profile"),
    overall_status: readString(value.overall_status, "handoff.verification.overall_status"),
    limitations: readOptionalStringArray(value.limitations, "handoff.verification.limitations"),
  };
}

function parseHandoffSealResponse(value: unknown): HandoffSealResponse {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid handoff response.");
  }
  return {
    package_id: readString(value.package_id, "handoff.package_id"),
    manifest: parseSealedManifest(value.manifest),
    session: parseHandoffSessionRecord(value.session),
    bundle: parseSealedBundleMetadata(value.bundle),
    verification: parseVerificationResult(value.verification),
    download_uri: readString(value.download_uri, "handoff.download_uri"),
    verification_uri: readString(value.verification_uri, "handoff.verification_uri"),
  };
}

function parseSBOMDocument(value: unknown): SBOMDocument {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid sbom document.");
  }
  return {
    id: readNumber(value.id, "document.id"),
    image_digest: readString(value.image_digest, "document.image_digest"),
    image_ref: readOptionalString(value.image_ref, "document.image_ref"),
    sbom_format: readString(value.sbom_format, "document.sbom_format"),
    source_ref: readOptionalString(value.source_ref, "document.source_ref"),
    sbom_hash: readOptionalString(value.sbom_hash, "document.sbom_hash"),
    created_at: readString(value.created_at, "document.created_at"),
  };
}

function parseSBOMComponent(value: unknown): SBOMComponent {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid sbom component.");
  }
  return {
    id: readNumber(value.id, "components[].id"),
    image_digest: readString(value.image_digest, "components[].image_digest"),
    component_name: readString(value.component_name, "components[].component_name"),
    component_version: readOptionalString(value.component_version, "components[].component_version"),
    component_type: readOptionalString(value.component_type, "components[].component_type"),
    license: readOptionalString(value.license, "components[].license"),
    purl: readOptionalString(value.purl, "components[].purl"),
    metadata: readOptionalRecord(value.metadata, "components[].metadata"),
    created_at: readString(value.created_at, "components[].created_at"),
  };
}

function parseSBOMImageResponse(value: unknown): SBOMImageResponse {
  if (!isRecord(value) || !Array.isArray(value.components)) {
    throw new Error("Audit API returned invalid sbom image response.");
  }
  return {
    document: parseSBOMDocument(value.document),
    component_count: readNumber(value.component_count, "component_count"),
    components: value.components.map(parseSBOMComponent),
  };
}

function parseSBOMComponentsResponse(value: unknown): SBOMComponentsResponse {
  if (!isRecord(value) || !Array.isArray(value.components)) {
    throw new Error("Audit API returned invalid sbom components response.");
  }
  return {
    components: value.components.map(parseSBOMComponent),
  };
}

function parseVulnerabilityDecision(value: unknown): VulnerabilityDecision {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid vulnerability decision.");
  }
  return {
    id: readNumber(value.id, "decision.id"),
    image_digest: readString(value.image_digest, "decision.image_digest"),
    cve_id: readString(value.cve_id, "decision.cve_id"),
    decision: readString(value.decision, "decision.decision") as VulnerabilityDecision["decision"],
    justification: readString(value.justification, "decision.justification"),
    decided_by: readString(value.decided_by, "decision.decided_by"),
    expires_at: readOptionalString(value.expires_at, "decision.expires_at"),
    active: readBoolean(value.active, "decision.active"),
    metadata: readOptionalRecord(value.metadata, "decision.metadata"),
    created_at: readString(value.created_at, "decision.created_at"),
    updated_at: readString(value.updated_at, "decision.updated_at"),
  };
}

function parseVEXMatch(value: unknown): VEXMatch {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid vex match.");
  }
  return {
    id: readNumber(value.id, "vex.id"),
    source_format: readString(value.source_format, "vex.source_format"),
    source_ref: readOptionalString(value.source_ref, "vex.source_ref"),
    vulnerability_id: readString(value.vulnerability_id, "vex.vulnerability_id"),
    status: readString(value.status, "vex.status") as VEXMatch["status"],
    justification: readOptionalString(value.justification, "vex.justification"),
    action_statement: readOptionalString(value.action_statement, "vex.action_statement"),
    impact_statement: readOptionalString(value.impact_statement, "vex.impact_statement"),
    fixed_version: readOptionalString(value.fixed_version, "vex.fixed_version"),
    created_by: readOptionalString(value.created_by, "vex.created_by"),
    updated_by: readOptionalString(value.updated_by, "vex.updated_by"),
    expires_at: readOptionalString(value.expires_at, "vex.expires_at"),
    created_at: readString(value.created_at, "vex.created_at"),
    updated_at: readString(value.updated_at, "vex.updated_at"),
  };
}

function parseVEXStatement(value: unknown): VEXStatement {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid vex statement.");
  }
  const scope = readOptionalRecord(value.scope, "statement.scope") || {};
  return {
    id: readNumber(value.id, "statement.id"),
    statement_key: readOptionalString(value.statement_key, "statement.statement_key"),
    source_format: readString(value.source_format, "statement.source_format"),
    source_ref: readOptionalString(value.source_ref, "statement.source_ref"),
    vulnerability_id: readString(value.vulnerability_id, "statement.vulnerability_id"),
    scope: {
      image_digest: readOptionalString(scope.image_digest, "statement.scope.image_digest"),
      package_name: readOptionalString(scope.package_name, "statement.scope.package_name"),
      purl: readOptionalString(scope.purl, "statement.scope.purl"),
      repo: readOptionalString(scope.repo, "statement.scope.repo"),
      workload: readOptionalString(scope.workload, "statement.scope.workload"),
      tenant_id: readOptionalString(scope.tenant_id, "statement.scope.tenant_id"),
      cluster_id: readOptionalString(scope.cluster_id, "statement.scope.cluster_id"),
      environment: readOptionalString(scope.environment, "statement.scope.environment"),
      namespace: readOptionalString(scope.namespace, "statement.scope.namespace"),
    },
    status: readString(value.status, "statement.status") as VEXStatement["status"],
    justification: readOptionalString(value.justification, "statement.justification"),
    action_statement: readOptionalString(value.action_statement, "statement.action_statement"),
    impact_statement: readOptionalString(value.impact_statement, "statement.impact_statement"),
    fixed_version: readOptionalString(value.fixed_version, "statement.fixed_version"),
    created_by: readOptionalString(value.created_by, "statement.created_by"),
    updated_by: readOptionalString(value.updated_by, "statement.updated_by"),
    expires_at: readOptionalString(value.expires_at, "statement.expires_at"),
    revoked_at: readOptionalString(value.revoked_at, "statement.revoked_at"),
    revoked_by: readOptionalString(value.revoked_by, "statement.revoked_by"),
    active: readBoolean(value.active, "statement.active"),
    metadata: readOptionalRecord(value.metadata, "statement.metadata"),
    created_at: readString(value.created_at, "statement.created_at"),
    updated_at: readString(value.updated_at, "statement.updated_at"),
  };
}

function parseVulnerabilityFinding(value: unknown): VulnerabilityFinding {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid vulnerability finding.");
  }
  return {
    id: readNumber(value.id, "findings[].id"),
    image_digest: readString(value.image_digest, "findings[].image_digest"),
    image_ref: readOptionalString(value.image_ref, "findings[].image_ref"),
    scan_run_id: readNumber(value.scan_run_id, "findings[].scan_run_id"),
    cve_id: readString(value.cve_id, "findings[].cve_id"),
    severity: readOptionalString(value.severity, "findings[].severity"),
    package_name: readOptionalString(value.package_name, "findings[].package_name"),
    package_version: readOptionalString(value.package_version, "findings[].package_version"),
    fixed_version: readOptionalString(value.fixed_version, "findings[].fixed_version"),
    purl: readOptionalString(value.purl, "findings[].purl"),
    status: readString(value.status, "findings[].status") as VulnerabilityFinding["status"],
    title: readOptionalString(value.title, "findings[].title"),
    description: readOptionalString(value.description, "findings[].description"),
    source: readOptionalString(value.source, "findings[].source"),
    metadata: readOptionalRecord(value.metadata, "findings[].metadata"),
    first_seen_at: readString(value.first_seen_at, "findings[].first_seen_at"),
    last_seen_at: readString(value.last_seen_at, "findings[].last_seen_at"),
    vex: value.vex ? parseVEXMatch(value.vex) : undefined,
    decision: value.decision ? parseVulnerabilityDecision(value.decision) : undefined,
  };
}

function parseActiveWorkload(value: unknown): ActiveWorkloadRef {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid workload reference.");
  }
  return {
    tenant_id: readOptionalString(value.tenant_id, "workloads[].tenant_id"),
    environment: readOptionalString(value.environment, "workloads[].environment"),
    namespace: readOptionalString(value.namespace, "workloads[].namespace"),
    workload: readOptionalString(value.workload, "workloads[].workload"),
    repo: readOptionalString(value.repo, "workloads[].repo"),
    image: readOptionalString(value.image, "workloads[].image"),
    digest: readOptionalString(value.digest, "workloads[].digest"),
  };
}

function parseVulnerabilitiesResponse(value: unknown): VulnerabilitiesResponse {
  if (!isRecord(value) || !Array.isArray(value.findings)) {
    throw new Error("Audit API returned invalid vulnerabilities response.");
  }
  return {
    findings: value.findings.map(parseVulnerabilityFinding),
  };
}

function parseBlastRadiusItem(value: unknown): VulnerabilityBlastRadiusItem {
  if (!isRecord(value) || !Array.isArray(value.findings) || !Array.isArray(value.workloads)) {
    throw new Error("Audit API returned invalid blast-radius item.");
  }
  return {
    image_digest: readString(value.image_digest, "items[].image_digest"),
    image_ref: readOptionalString(value.image_ref, "items[].image_ref"),
    findings: value.findings.map(parseVulnerabilityFinding),
    workloads: value.workloads.map(parseActiveWorkload),
  };
}

function parseVulnerabilityBlastRadiusResponse(value: unknown): VulnerabilityBlastRadiusResponse {
  if (!isRecord(value) || !Array.isArray(value.items)) {
    throw new Error("Audit API returned invalid blast-radius response.");
  }
  return {
    items: value.items.map(parseBlastRadiusItem),
    applied_filters: readOptionalRecord(value.applied_filters, "applied_filters") as Record<string, string> || {},
  };
}

function parseTimelineEntry(value: unknown): VulnerabilityTimelineEntry {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid timeline entry.");
  }
  return {
    image_digest: readString(value.image_digest, "items[].image_digest"),
    cve_id: readString(value.cve_id, "items[].cve_id"),
    package_name: readOptionalString(value.package_name, "items[].package_name"),
    package_version: readOptionalString(value.package_version, "items[].package_version"),
    severity: readOptionalString(value.severity, "items[].severity"),
    status: readString(value.status, "items[].status") as VulnerabilityTimelineEntry["status"],
    first_seen_at: readString(value.first_seen_at, "items[].first_seen_at"),
    last_seen_at: readString(value.last_seen_at, "items[].last_seen_at"),
    decision: value.decision ? parseVulnerabilityDecision(value.decision) : undefined,
  };
}

function parseVulnerabilityTimelineResponse(value: unknown): VulnerabilityTimelineResponse {
  if (!isRecord(value) || !Array.isArray(value.items)) {
    throw new Error("Audit API returned invalid vulnerability timeline response.");
  }
  return {
    items: value.items.map(parseTimelineEntry),
    applied_filters: readOptionalRecord(value.applied_filters, "applied_filters") as Record<string, string> || {},
  };
}

function parseVulnerabilityDecisionsResponse(value: unknown): VulnerabilityDecisionsResponse {
  if (!isRecord(value) || !Array.isArray(value.decisions)) {
    throw new Error("Audit API returned invalid vulnerability decisions response.");
  }
  return {
    decisions: value.decisions.map(parseVulnerabilityDecision),
  };
}

function parseVulnerabilityDecisionActionResponse(value: unknown): { status: string; decision: VulnerabilityDecision } {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid vulnerability decision action response.");
  }
  return {
    status: readString(value.status, "status"),
    decision: parseVulnerabilityDecision(value.decision),
  };
}

function parseVEXStatementsResponse(value: unknown): VEXStatementsResponse {
  if (!isRecord(value) || !Array.isArray(value.statements)) {
    throw new Error("Audit API returned invalid vex statements response.");
  }
  return {
    statements: value.statements.map(parseVEXStatement),
  };
}

function parseVEXStatementActionResponse(value: unknown): VEXStatementActionResponse {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid vex action response.");
  }
  return {
    status: readString(value.status, "status"),
    statement: parseVEXStatement(value.statement),
  };
}

function parseVEXStatusSummary(value: unknown): VEXStatusSummary {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid vex status.");
  }
  const counts = readOptionalRecord(value.counts_by_status, "counts_by_status") || {};
  const countsByStatus: Record<string, number> = {};
  for (const [key, count] of Object.entries(counts)) {
    countsByStatus[key] = readNumber(count, `counts_by_status.${key}`);
  }
  return {
    active_count: readNumber(value.active_count, "active_count"),
    expiring_count: readNumber(value.expiring_count, "expiring_count"),
    revoked_count: readNumber(value.revoked_count, "revoked_count"),
    counts_by_status: countsByStatus,
    applied_filters: readOptionalRecord(value.applied_filters, "applied_filters") as Record<string, string> | undefined,
  };
}

function parseSigningIdentityObservation(value: unknown): SigningIdentityObservation {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid signing identity observation.");
  }
  return {
    id: readString(value.id, "signing_identity.id"),
    provider_type: readOptionalString(value.provider_type, "signing_identity.provider_type"),
    issuer: readOptionalString(value.issuer, "signing_identity.issuer"),
    signer_identity: readOptionalString(value.signer_identity, "signing_identity.signer_identity"),
    subject: readOptionalString(value.subject, "signing_identity.subject"),
    repository: readOptionalString(value.repository, "signing_identity.repository"),
    workflow: readOptionalString(value.workflow, "signing_identity.workflow"),
    ref: readOptionalString(value.ref, "signing_identity.ref"),
    commit_sha: readOptionalString(value.commit_sha, "signing_identity.commit_sha"),
    image_digest: readOptionalString(value.image_digest, "signing_identity.image_digest"),
    tenant_id: readOptionalString(value.tenant_id, "signing_identity.tenant_id"),
    cluster_id: readOptionalString(value.cluster_id, "signing_identity.cluster_id"),
    environment: readOptionalString(value.environment, "signing_identity.environment"),
    first_seen_at: readOptionalString(value.first_seen_at, "signing_identity.first_seen_at"),
    last_seen_at: readOptionalString(value.last_seen_at, "signing_identity.last_seen_at"),
    event_count: readNumber(value.event_count, "signing_identity.event_count"),
    artifact_count: readNumber(value.artifact_count, "signing_identity.artifact_count"),
    verification_state: readOptionalString(value.verification_state, "signing_identity.verification_state"),
    authorized: readString(value.authorized, "signing_identity.authorized"),
    matched_policy_id: readOptionalString(value.matched_policy_id, "signing_identity.matched_policy_id"),
    distrusted_after: readOptionalString(value.distrusted_after, "signing_identity.distrusted_after"),
    reason_code: readOptionalString(value.reason_code, "signing_identity.reason_code"),
    reason_detail: readOptionalString(value.reason_detail, "signing_identity.reason_detail"),
  };
}

function parseSigningIdentityObservationsResponse(value: unknown): SigningIdentityObservationsResponse {
  if (!isRecord(value) || !Array.isArray(value.items)) {
    throw new Error("Audit API returned invalid signing identities response.");
  }
  return {
    items: value.items.map(parseSigningIdentityObservation),
  };
}

function parseSigningIdentityPolicy(value: unknown): SigningIdentityPolicy {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid signing identity policy.");
  }
  return {
    id: readString(value.id, "policy.id"),
    name: readOptionalString(value.name, "policy.name"),
    provider_type: readString(value.provider_type, "policy.provider_type"),
    issuer: readOptionalString(value.issuer, "policy.issuer"),
    signer_identity: readOptionalString(value.signer_identity, "policy.signer_identity"),
    subject: readOptionalString(value.subject, "policy.subject"),
    repository: readOptionalString(value.repository, "policy.repository"),
    workflow: readOptionalString(value.workflow, "policy.workflow"),
    ref: readOptionalString(value.ref, "policy.ref"),
    tenant_id: readOptionalString(value.tenant_id, "policy.tenant_id"),
    cluster_id: readOptionalString(value.cluster_id, "policy.cluster_id"),
    environment: readOptionalString(value.environment, "policy.environment"),
    enabled: readBoolean(value.enabled, "policy.enabled"),
    distrusted_after: readOptionalString(value.distrusted_after, "policy.distrusted_after"),
    distrust_reason: readOptionalString(value.distrust_reason, "policy.distrust_reason"),
    created_at: readString(value.created_at, "policy.created_at"),
    updated_at: readString(value.updated_at, "policy.updated_at"),
    created_by: readOptionalString(value.created_by, "policy.created_by"),
    updated_by: readOptionalString(value.updated_by, "policy.updated_by"),
  };
}

function parseSigningIdentityPoliciesResponse(value: unknown): SigningIdentityPoliciesResponse {
  if (!isRecord(value) || !Array.isArray(value.policies)) {
    throw new Error("Audit API returned invalid signing identity policies response.");
  }
  return {
    policies: value.policies.map(parseSigningIdentityPolicy),
  };
}

function parseSigningIdentityFinding(value: unknown): SigningIdentityFinding {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid signing identity finding.");
  }
  return {
    id: readString(value.id, "finding.id"),
    type: readString(value.type, "finding.type"),
    severity: readString(value.severity, "finding.severity"),
    repository: readOptionalString(value.repository, "finding.repository"),
    workflow: readOptionalString(value.workflow, "finding.workflow"),
    ref: readOptionalString(value.ref, "finding.ref"),
    policy_id: readOptionalString(value.policy_id, "finding.policy_id"),
    observation_id: readOptionalString(value.observation_id, "finding.observation_id"),
    reason: readString(value.reason, "finding.reason"),
    detected_at: readOptionalString(value.detected_at, "finding.detected_at"),
    advisory: readBoolean(value.advisory, "finding.advisory"),
  };
}

function parseSigningIdentityFindingsResponse(value: unknown): SigningIdentityFindingsResponse {
  if (!isRecord(value) || !Array.isArray(value.items)) {
    throw new Error("Audit API returned invalid signing identity findings response.");
  }
  return {
    items: value.items.map(parseSigningIdentityFinding),
  };
}

function parseSigningIdentityStatus(value: unknown): SigningIdentityStatus {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid signing identity status.");
  }
  const counts = readOptionalRecord(value.counts_by_reason_code, "counts_by_reason_code") || {};
  const countsByReasonCode: Record<string, number> = {};
  for (const [key, count] of Object.entries(counts)) {
    countsByReasonCode[key] = readNumber(count, `counts_by_reason_code.${key}`);
  }
  return {
    enforcement_mode: readString(value.enforcement_mode, "enforcement_mode"),
    require_rekor: readBoolean(value.require_rekor, "require_rekor"),
    total_policies: readNumber(value.total_policies, "total_policies"),
    enabled_policies: readNumber(value.enabled_policies, "enabled_policies"),
    observed_identities: readNumber(value.observed_identities, "observed_identities"),
    authorized: readNumber(value.authorized, "authorized"),
    unauthorized: readNumber(value.unauthorized, "unauthorized"),
    unknown: readNumber(value.unknown, "unknown"),
    findings: readNumber(value.findings, "findings"),
    workflow_drift_findings: readNumber(value.workflow_drift_findings, "workflow_drift_findings"),
    counts_by_reason_code: countsByReasonCode,
  };
}

function parseVulnerabilityNetResponse(value: unknown): VulnerabilityNetResponse {
  if (!isRecord(value) || !Array.isArray(value.findings)) {
    throw new Error("Audit API returned invalid vulnerability net response.");
  }
  return {
    raw_count: readNumber(value.raw_count, "raw_count"),
    resolved_by_vex_count: readNumber(value.resolved_by_vex_count, "resolved_by_vex_count"),
    actionable_count: readNumber(value.actionable_count, "actionable_count"),
    under_investigation_count: readNumber(value.under_investigation_count, "under_investigation_count"),
    severity_threshold: readOptionalString(value.severity_threshold, "severity_threshold"),
    threshold_breached: readBoolean(value.threshold_breached, "threshold_breached"),
    findings: value.findings.map(parseVulnerabilityFinding),
    applied_filters: readOptionalRecord(value.applied_filters, "applied_filters") as Record<string, string> || {},
  };
}

function parseVulnerabilityRescanResponse(value: unknown): VulnerabilityRescanResponse {
  if (!isRecord(value) || !Array.isArray(value.scanned_digests)) {
    throw new Error("Audit API returned invalid vulnerability rescan response.");
  }
  if (!value.scanned_digests.every((item) => typeof item === "string")) {
    throw new Error("Audit API returned invalid scanned_digests.");
  }
  return {
    status: readString(value.status, "status"),
    scanned_digests: value.scanned_digests,
    scan_runs: readNumber(value.scan_runs, "scan_runs"),
  };
}

function parseTrustScoreMetric(value: unknown): TrustScoreMetric {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid trust metric.");
  }
  return {
    id: readString(value.id, "trust_metric.id"),
    name: readString(value.name, "trust_metric.name"),
    weight: readNumber(value.weight, "trust_metric.weight"),
    score: readNumber(value.score, "trust_metric.score"),
    status: readString(value.status, "trust_metric.status"),
    reason_code: readString(value.reason_code, "trust_metric.reason_code"),
    reason_detail: readOptionalString(value.reason_detail, "trust_metric.reason_detail"),
    evidence_refs: readOptionalStringArray(value.evidence_refs, "trust_metric.evidence_refs"),
    advisory_only: readBoolean(value.advisory_only, "trust_metric.advisory_only"),
    public_publishable: readBoolean(value.public_publishable, "trust_metric.public_publishable"),
    mapping_refs: readOptionalStringArray(value.mapping_refs, "trust_metric.mapping_refs"),
  };
}

function parseTrustScorecard(value: unknown): TrustScorecard {
  if (!isRecord(value) || !Array.isArray(value.metrics)) {
    throw new Error("Audit API returned invalid trust scorecard.");
  }
  return {
    id: readString(value.id, "scorecard.id"),
    scope_type: readString(value.scope_type, "scorecard.scope_type"),
    scope_ref: readString(value.scope_ref, "scorecard.scope_ref"),
    tenant_id: readOptionalString(value.tenant_id, "scorecard.tenant_id"),
    cluster_id: readOptionalString(value.cluster_id, "scorecard.cluster_id"),
    environment: readOptionalString(value.environment, "scorecard.environment"),
    repo: readOptionalString(value.repo, "scorecard.repo"),
    calculated_at: readString(value.calculated_at, "scorecard.calculated_at"),
    overall_grade: readString(value.overall_grade, "scorecard.overall_grade"),
    overall_score: readNumber(value.overall_score, "scorecard.overall_score"),
    signing_coverage: readNumber(value.signing_coverage, "scorecard.signing_coverage"),
    transparency_coverage: readNumber(value.transparency_coverage, "scorecard.transparency_coverage"),
    sbom_or_provenance_coverage: readNumber(value.sbom_or_provenance_coverage, "scorecard.sbom_or_provenance_coverage"),
    actionable_vulnerability_count: readNumber(value.actionable_vulnerability_count, "scorecard.actionable_vulnerability_count"),
    stale_exception_count: readNumber(value.stale_exception_count, "scorecard.stale_exception_count"),
    publication_mode: readString(value.publication_mode, "scorecard.publication_mode"),
    metrics: value.metrics.map(parseTrustScoreMetric),
    notes: readOptionalStringArray(value.notes, "scorecard.notes"),
  };
}

function parseTrustBadge(value: unknown): TrustBadge {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid trust badge.");
  }
  return {
    id: readString(value.id, "badge.id"),
    label: readString(value.label, "badge.label"),
    state: readString(value.state, "badge.state"),
    summary: readString(value.summary, "badge.summary"),
    public_publishable: readBoolean(value.public_publishable, "badge.public_publishable"),
    svg: readOptionalString(value.svg, "badge.svg"),
  };
}

function parseAuditFinding(value: unknown): AuditFinding {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid audit finding.");
  }
  return {
    id: readString(value.id, "audit_finding.id"),
    category: readString(value.category, "audit_finding.category"),
    severity: readString(value.severity, "audit_finding.severity"),
    status: readString(value.status, "audit_finding.status"),
    reason_code: readString(value.reason_code, "audit_finding.reason_code"),
    reason_detail: readOptionalString(value.reason_detail, "audit_finding.reason_detail"),
    scope_ref: readOptionalString(value.scope_ref, "audit_finding.scope_ref"),
    evidence_refs: readOptionalStringArray(value.evidence_refs, "audit_finding.evidence_refs"),
    advisory_only: readBoolean(value.advisory_only, "audit_finding.advisory_only"),
    public_publishable: readBoolean(value.public_publishable, "audit_finding.public_publishable"),
    detected_at: readString(value.detected_at, "audit_finding.detected_at"),
  };
}

function parseStandardsMapping(value: unknown): StandardsMapping {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid standards mapping.");
  }
  return {
    standard: readString(value.standard, "mapping.standard"),
    control: readString(value.control, "mapping.control"),
    status: readString(value.status, "mapping.status"),
    summary: readString(value.summary, "mapping.summary"),
    evidence_refs: readOptionalStringArray(value.evidence_refs, "mapping.evidence_refs"),
  };
}

function parsePublishedTrustView(value: unknown): PublishedTrustView {
  if (!isRecord(value) || !Array.isArray(value.badges) || !Array.isArray(value.metrics) || !Array.isArray(value.mapping)) {
    throw new Error("Audit API returned invalid published trust view.");
  }
  return {
    generated_at: readString(value.generated_at, "published_trust.generated_at"),
    scope_type: readString(value.scope_type, "published_trust.scope_type"),
    scope_ref: readString(value.scope_ref, "published_trust.scope_ref"),
    overall_grade: readString(value.overall_grade, "published_trust.overall_grade"),
    overall_score: readNumber(value.overall_score, "published_trust.overall_score"),
    badges: value.badges.map(parseTrustBadge),
    metrics: value.metrics.map(parseTrustScoreMetric),
    mapping: value.mapping.map(parseStandardsMapping),
    notes: readOptionalStringArray(value.notes, "published_trust.notes"),
  };
}

function parseAuditReport(value: unknown): AuditReport {
  if (!isRecord(value) || !Array.isArray(value.findings) || !Array.isArray(value.badges) || !Array.isArray(value.standards_mapping)) {
    throw new Error("Audit API returned invalid audit report.");
  }
  return {
    id: readString(value.id, "audit_report.id"),
    generated_at: readString(value.generated_at, "audit_report.generated_at"),
    scope_type: readString(value.scope_type, "audit_report.scope_type"),
    scope_ref: readString(value.scope_ref, "audit_report.scope_ref"),
    scorecard: parseTrustScorecard(value.scorecard),
    findings: value.findings.map(parseAuditFinding),
    badges: value.badges.map(parseTrustBadge),
    standards_mapping: value.standards_mapping.map(parseStandardsMapping),
    public_view: value.public_view ? parsePublishedTrustView(value.public_view) : undefined,
    limitations: readOptionalStringArray(value.limitations, "audit_report.limitations"),
    format: readOptionalString(value.format, "audit_report.format"),
    generated_by: readOptionalString(value.generated_by, "audit_report.generated_by"),
  };
}

function parseGuidanceGrouping(value: unknown): GuidanceGrouping {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid guidance grouping.");
  }
  return {
    key: readString(value.key, "guidance.grouping.key"),
    label: readString(value.label, "guidance.grouping.label"),
    category: readString(value.category, "guidance.grouping.category"),
    finding_count: readNumber(value.finding_count, "guidance.grouping.finding_count"),
    priority: readString(value.priority, "guidance.grouping.priority"),
    contextual_risk_score: readNumber(value.contextual_risk_score, "guidance.grouping.contextual_risk_score"),
    heuristic: readBoolean(value.heuristic, "guidance.grouping.heuristic"),
  };
}

function parseGuidanceVEXDraftSuggestion(value: unknown): GuidanceVEXDraftSuggestion {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid VEX draft suggestion.");
  }
  return {
    id: readString(value.id, "guidance.vex_draft.id"),
    candidate_status: readString(value.candidate_status, "guidance.vex_draft.candidate_status"),
    justification: readString(value.justification, "guidance.vex_draft.justification"),
    impact_statement: readString(value.impact_statement, "guidance.vex_draft.impact_statement"),
    missing_evidence: readOptionalStringArray(value.missing_evidence, "guidance.vex_draft.missing_evidence"),
    confidence: readString(value.confidence, "guidance.vex_draft.confidence"),
    confidence_basis: readOptionalString(value.confidence_basis, "guidance.vex_draft.confidence_basis"),
    advisory_only: readBoolean(value.advisory_only, "guidance.vex_draft.advisory_only"),
    requires_human_review: readBoolean(value.requires_human_review, "guidance.vex_draft.requires_human_review"),
    docs_refs: readOptionalStringArray(value.docs_refs, "guidance.vex_draft.docs_refs"),
  };
}

function parseBreakGlassGuidance(value: unknown): BreakGlassGuidance {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid break-glass guidance.");
  }
  return {
    scope_explanation: readString(value.scope_explanation, "guidance.break_glass.scope_explanation"),
    narrower_alternative: readOptionalString(value.narrower_alternative, "guidance.break_glass.narrower_alternative"),
    cleanup_reminders: readOptionalStringArray(value.cleanup_reminders, "guidance.break_glass.cleanup_reminders"),
    proposed_containment_steps: readOptionalStringArray(value.proposed_containment_steps, "guidance.break_glass.proposed_containment_steps"),
    confidence: readString(value.confidence, "guidance.break_glass.confidence"),
    confidence_basis: readOptionalString(value.confidence_basis, "guidance.break_glass.confidence_basis"),
    advisory_only: readBoolean(value.advisory_only, "guidance.break_glass.advisory_only"),
    requires_human_review: readBoolean(value.requires_human_review, "guidance.break_glass.requires_human_review"),
    docs_refs: readOptionalStringArray(value.docs_refs, "guidance.break_glass.docs_refs"),
  };
}

function parseGuidanceItem(value: unknown): GuidanceItem {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid guidance item.");
  }
  return {
    id: readString(value.id, "guidance.item.id"),
    category: readString(value.category, "guidance.item.category"),
    source_component: readOptionalString(value.source_component, "guidance.item.source_component"),
    grouping: parseGuidanceGrouping(value.grouping),
    related_reason_codes: readOptionalStringArray(value.related_reason_codes, "guidance.item.related_reason_codes"),
    finding_refs: readOptionalStringArray(value.finding_refs, "guidance.item.finding_refs"),
    evidence_refs: readOptionalStringArray(value.evidence_refs, "guidance.item.evidence_refs"),
    docs_refs: readOptionalStringArray(value.docs_refs, "guidance.item.docs_refs"),
    scope_type: readOptionalString(value.scope_type, "guidance.item.scope_type"),
    scope_ref: readOptionalString(value.scope_ref, "guidance.item.scope_ref"),
    tenant_id: readOptionalString(value.tenant_id, "guidance.item.tenant_id"),
    cluster_id: readOptionalString(value.cluster_id, "guidance.item.cluster_id"),
    environment: readOptionalString(value.environment, "guidance.item.environment"),
    repository: readOptionalString(value.repository, "guidance.item.repository"),
    severity: readOptionalString(value.severity, "guidance.item.severity"),
    priority: readString(value.priority, "guidance.item.priority"),
    confidence: readString(value.confidence, "guidance.item.confidence"),
    confidence_basis: readOptionalString(value.confidence_basis, "guidance.item.confidence_basis"),
    explanation: readOptionalString(value.explanation, "guidance.item.explanation"),
    recommendation_summary: readOptionalString(value.recommendation_summary, "guidance.item.recommendation_summary"),
    recommendation_steps: readOptionalStringArray(value.recommendation_steps, "guidance.item.recommendation_steps"),
    safer_alternative: readOptionalString(value.safer_alternative, "guidance.item.safer_alternative"),
    impact_summary: readOptionalString(value.impact_summary, "guidance.item.impact_summary"),
    data_limitations: readOptionalStringArray(value.data_limitations, "guidance.item.data_limitations"),
    advisory_only: readBoolean(value.advisory_only, "guidance.item.advisory_only"),
    requires_human_review: readBoolean(value.requires_human_review, "guidance.item.requires_human_review"),
    generated_at: readString(value.generated_at, "guidance.item.generated_at"),
    generated_by: readString(value.generated_by, "guidance.item.generated_by"),
    template_version: readOptionalString(value.template_version, "guidance.item.template_version"),
    heuristic: readBoolean(value.heuristic, "guidance.item.heuristic"),
    vex_draft: value.vex_draft ? parseGuidanceVEXDraftSuggestion(value.vex_draft) : undefined,
    break_glass_guidance: value.break_glass_guidance ? parseBreakGlassGuidance(value.break_glass_guidance) : undefined,
  };
}

function parseGuidanceSummary(value: unknown): GuidanceSummary {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid guidance summary.");
  }
  const countsByCategorySource = readOptionalRecord(value.counts_by_category, "guidance.summary.counts_by_category") || {};
  const countsByCategory: Record<string, number> = {};
  for (const [key, count] of Object.entries(countsByCategorySource)) {
    countsByCategory[key] = readNumber(count, `guidance.summary.counts_by_category.${key}`);
  }
  const countsByPrioritySource = readOptionalRecord(value.counts_by_priority, "guidance.summary.counts_by_priority") || {};
  const countsByPriority: Record<string, number> = {};
  for (const [key, count] of Object.entries(countsByPrioritySource)) {
    countsByPriority[key] = readNumber(count, `guidance.summary.counts_by_priority.${key}`);
  }
  return {
    total_items: readNumber(value.total_items, "guidance.summary.total_items"),
    counts_by_category: countsByCategory,
    counts_by_priority: countsByPriority,
    guidance_mode: readString(value.guidance_mode, "guidance.summary.guidance_mode"),
    ai_enabled: readBoolean(value.ai_enabled, "guidance.summary.ai_enabled"),
    deterministic_only: readBoolean(value.deterministic_only, "guidance.summary.deterministic_only"),
    limitations: readOptionalStringArray(value.limitations, "guidance.summary.limitations"),
  };
}

function parseGuidanceResponse(value: unknown): GuidanceResponse {
  if (!isRecord(value) || !Array.isArray(value.items)) {
    throw new Error("Audit API returned invalid guidance response.");
  }
  return {
    generated_at: readString(value.generated_at, "guidance.generated_at"),
    scope_type: readOptionalString(value.scope_type, "guidance.scope_type"),
    scope_ref: readOptionalString(value.scope_ref, "guidance.scope_ref"),
    tenant_id: readOptionalString(value.tenant_id, "guidance.tenant_id"),
    cluster_id: readOptionalString(value.cluster_id, "guidance.cluster_id"),
    environment: readOptionalString(value.environment, "guidance.environment"),
    repository: readOptionalString(value.repository, "guidance.repository"),
    items: value.items.map(parseGuidanceItem),
    summary: parseGuidanceSummary(value.summary),
  };
}

function parseAIInsightsResponse(value: unknown): AIInsightsResponse {
  if (!isRecord(value) || !Array.isArray(value.top_items)) {
    throw new Error("Audit API returned invalid AI insights response.");
  }
  return {
    summary: parseGuidanceSummary(value.summary),
    top_items: value.top_items.map(parseGuidanceItem),
  };
}

function parseAuthStatus(value: unknown): AuthStatus {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid auth status.");
  }
  return {
    authenticated: readBoolean(value.authenticated, "authenticated"),
    auth_mode: readString(value.auth_mode, "auth_mode"),
    subject: readOptionalString(value.subject, "subject"),
    role: readOptionalString(value.role, "role"),
    token_id: readOptionalString(value.token_id, "token_id"),
    identity_type: readOptionalString(value.identity_type, "identity_type"),
    email: readOptionalString(value.email, "email"),
    tenant_id: readOptionalString(value.tenant_id, "tenant_id"),
    global_scope: value.global_scope === undefined ? undefined : readBoolean(value.global_scope, "global_scope"),
  };
}

function parseSyncStatus(value: unknown): SyncStatus {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid sync status.");
  }
  return {
    sync_mode: readOptionalString(value.sync_mode, "sync_mode"),
    mode: readString(value.mode, "mode"),
    cluster_id: readOptionalString(value.cluster_id, "cluster_id"),
    hub_url: readOptionalString(value.hub_url, "hub_url"),
    fail_mode: readOptionalString(value.fail_mode, "fail_mode"),
    health: readString(value.health, "health"),
    current_revision: readOptionalString(value.current_revision, "current_revision"),
    revision_etag: readOptionalString(value.revision_etag, "revision_etag"),
    last_successful_sync_at: readOptionalString(value.last_successful_sync_at, "last_successful_sync_at"),
    last_attempt_at: readOptionalString(value.last_attempt_at, "last_attempt_at"),
    last_error: readOptionalString(value.last_error, "last_error"),
    cache_present: readBoolean(value.cache_present, "cache_present"),
    stale_after_seconds:
      value.stale_after_seconds === undefined ? undefined : readNumber(value.stale_after_seconds, "stale_after_seconds"),
    summary: readOptionalString(value.summary, "summary"),
  };
}

function parseIncidentImpact(value: unknown): InvestigationIncident["governanceImpacts"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid incident impact.");
  }
  return {
    id: readString(value.id, "governance_impacts[].id"),
    title: readString(value.title, "governance_impacts[].title"),
    detail: readString(value.detail, "governance_impacts[].detail"),
    tone: readString(value.tone, "governance_impacts[].tone") as InvestigationIncident["governanceImpacts"][number]["tone"],
  };
}

function parseIncidentTimelineEntry(value: unknown): InvestigationIncident["timeline"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid incident timeline entry.");
  }
  return {
    id: readString(value.id, "timeline[].id"),
    kind: readOptionalString(value.kind, "timeline[].kind"),
    timestamp: readOptionalString(value.timestamp, "timeline[].timestamp"),
    title: readString(value.title, "timeline[].title"),
    summary: readString(value.summary, "timeline[].summary"),
    eventType: readString(value.event_type, "timeline[].event_type"),
    outcome: readString(value.outcome, "timeline[].outcome") as InvestigationIncident["timeline"][number]["outcome"],
    requestID: readOptionalString(value.request_id, "timeline[].request_id"),
    actor: readOptionalString(value.actor, "timeline[].actor"),
  };
}

function parseIncidentNote(value: unknown): InvestigationIncident["notes"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid incident note.");
  }
  return {
    id: readString(value.id, "notes[].id"),
    note: readString(value.note, "notes[].note"),
    actor: readOptionalString(value.actor, "notes[].actor"),
    timestamp: readOptionalString(value.timestamp, "notes[].timestamp"),
  };
}

function parseIncidentHistoryEntry(value: unknown): InvestigationIncident["history"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid incident history entry.");
  }
  return {
    id: readString(value.id, "history[].id"),
    kind: readString(value.kind, "history[].kind"),
    timestamp: readOptionalString(value.timestamp, "history[].timestamp"),
    actor: readOptionalString(value.actor, "history[].actor"),
    summary: readString(value.summary, "history[].summary"),
    state: readOptionalString(value.state, "history[].state"),
    owner: readOptionalString(value.owner, "history[].owner"),
    note: readOptionalString(value.note, "history[].note"),
  };
}

function parseIncidentEvidencePack(value: unknown): InvestigationIncident["evidencePack"] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid incident evidence pack.");
  }
  return {
    requestIDs: readOptionalStringArray(value.request_ids, "evidence_pack.request_ids") || [],
    digests: readOptionalStringArray(value.digests, "evidence_pack.digests") || [],
    bundles: readOptionalStringArray(value.bundles, "evidence_pack.bundles") || [],
    exceptions: readOptionalStringArray(value.exceptions, "evidence_pack.exceptions") || [],
    vulnerabilities: readOptionalStringArray(value.vulnerabilities, "evidence_pack.vulnerabilities") || [],
  };
}

function parseIncidentAssignment(value: unknown): InvestigationIncident["assignment"] {
  if (value === undefined || value === null) {
    return {};
  }
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid incident assignment.");
  }
  return {
    owner: readOptionalString(value.owner, "assignment.owner"),
    at: readOptionalString(value.at, "assignment.at"),
    by: readOptionalString(value.by, "assignment.by"),
    reason: readOptionalString(value.reason, "assignment.reason"),
  };
}

function parseIncidentResolution(value: unknown): InvestigationIncident["resolution"] {
  if (value === undefined || value === null) {
    return { refs: [] };
  }
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid incident resolution.");
  }
  return {
    type: readOptionalString(value.type, "resolution.type"),
    summary: readOptionalString(value.summary, "resolution.summary"),
    details: readOptionalString(value.details, "resolution.details"),
    refs: readOptionalStringArray(value.refs, "resolution.refs") || [],
    by: readOptionalString(value.by, "resolution.by"),
    at: readOptionalString(value.at, "resolution.at"),
    followUpRequired:
      value.follow_up_required === undefined ? undefined : readBoolean(value.follow_up_required, "resolution.follow_up_required"),
  };
}

function parseIncidentMetricLink(value: unknown): InvestigationIncident["metricLinks"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid incident metric link.");
  }
  return {
    metricKey: readString(value.metric_key, "metric_links[].metric_key"),
    metricLabel: readString(value.metric_label, "metric_links[].metric_label"),
    linkReason: readString(value.link_reason, "metric_links[].link_reason"),
    supportingRefs: readOptionalStringArray(value.supporting_refs, "metric_links[].supporting_refs") || [],
    impactWeight: readNumber(value.impact_weight, "metric_links[].impact_weight"),
  };
}

function parseIncidentExportEventRef(value: unknown): IncidentExport["relatedEventRefs"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid incident export event ref.");
  }
  return {
    eventID: readNumber(value.event_id, "related_event_refs[].event_id"),
    requestID: readOptionalString(value.request_id, "related_event_refs[].request_id"),
    timestamp: readString(value.timestamp, "related_event_refs[].timestamp"),
    component: readString(value.component, "related_event_refs[].component"),
    eventType: readString(value.event_type, "related_event_refs[].event_type"),
    decision: readString(value.decision, "related_event_refs[].decision"),
    decisionHash: readOptionalString(value.decision_hash, "related_event_refs[].decision_hash"),
  };
}

function parseIncidentLifecycle(value: unknown): InvestigationIncident["lifecycle"] {
  if (value === undefined || value === null) {
    return undefined;
  }
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid incident lifecycle.");
  }
  return {
    state: readString(value.state, "lifecycle.state") as InvestigationIncident["state"],
    owner: readOptionalString(value.owner, "lifecycle.owner"),
    assignment: parseIncidentAssignment(value.assignment),
    resolution: parseIncidentResolution(value.resolution),
    resolutionSummary: readOptionalString(value.resolution_summary, "lifecycle.resolution_summary"),
    notes: Array.isArray(value.notes) ? value.notes.map(parseIncidentNote) : [],
    history: Array.isArray(value.history) ? value.history.map(parseIncidentHistoryEntry) : [],
    lastOperatorUpdateAt: readOptionalString(value.last_operator_update_at, "lifecycle.last_operator_update_at"),
    newActivityDetected:
      value.new_activity_detected === undefined ? false : readBoolean(value.new_activity_detected, "lifecycle.new_activity_detected"),
  };
}

function parseIncident(value: unknown): InvestigationIncident {
  if (!isRecord(value) || !Array.isArray(value.events) || !Array.isArray(value.governance_impacts) || !Array.isArray(value.timeline)) {
    throw new Error("Audit API returned invalid incident.");
  }
  const lifecycle = parseIncidentLifecycle(value.lifecycle);
  return {
    id: readString(value.id, "incidents[].id"),
    identityKey: readOptionalString(value.identity_key, "incidents[].identity_key"),
    categoryKey: readOptionalString(value.category_key, "incidents[].category_key"),
    title: readString(value.title, "incidents[].title"),
    category: readString(value.category, "incidents[].category"),
    severity: readString(value.severity, "incidents[].severity") as InvestigationIncident["severity"],
    priority: (readOptionalString(value.priority, "incidents[].priority") || readString(value.severity, "incidents[].severity")) as InvestigationIncident["priority"],
    state: (readOptionalString(value.state, "incidents[].state") || "open") as InvestigationIncident["state"],
    status: readString(value.status, "incidents[].status") as InvestigationIncident["status"],
    scopeType: readOptionalString(value.scope_type, "incidents[].scope_type"),
    scopeRef: readOptionalString(value.scope_ref, "incidents[].scope_ref"),
    tenantID: readOptionalString(value.tenant_id, "incidents[].tenant_id"),
    clusterID: readOptionalString(value.cluster_id, "incidents[].cluster_id"),
    environment: readOptionalString(value.environment, "incidents[].environment"),
    repository: readOptionalString(value.repository, "incidents[].repository"),
    summary: readString(value.summary, "incidents[].summary"),
    caseSummary: readString(value.case_summary, "incidents[].case_summary"),
    statusNarrative: readString(value.status_narrative, "incidents[].status_narrative"),
    likelyCause: readString(value.likely_cause, "incidents[].likely_cause"),
    recommendedAction: readString(value.recommended_action, "incidents[].recommended_action"),
    remediationChecklist: readOptionalStringArray(value.remediation_checklist, "incidents[].remediation_checklist") || [],
    eventCount: readNumber(value.event_count, "incidents[].event_count"),
    denyCount: readNumber(value.deny_count, "incidents[].deny_count"),
    allowCount: readNumber(value.allow_count, "incidents[].allow_count"),
    errorCount: readNumber(value.error_count, "incidents[].error_count"),
    openedAt: readOptionalString(value.opened_at, "incidents[].opened_at"),
    updatedAt: readOptionalString(value.updated_at, "incidents[].updated_at"),
    lastActivityAt: readOptionalString(value.last_activity_at, "incidents[].last_activity_at"),
    lastOperatorUpdateAt: readOptionalString(value.last_operator_update_at, "incidents[].last_operator_update_at"),
    resolvedAt: readOptionalString(value.resolved_at, "incidents[].resolved_at"),
    owner: readOptionalString(value.owner, "incidents[].owner"),
    assignment: parseIncidentAssignment(value.assignment),
    resolution: parseIncidentResolution(value.resolution),
    lifecycle,
    resolutionSummary: readOptionalString(value.resolution_summary, "incidents[].resolution_summary"),
    newActivityDetected:
      value.new_activity_detected === undefined ? lifecycle?.newActivityDetected || false : readBoolean(value.new_activity_detected, "incidents[].new_activity_detected"),
    notes: Array.isArray(value.notes) ? value.notes.map(parseIncidentNote) : lifecycle?.notes || [],
    history: Array.isArray(value.history) ? value.history.map(parseIncidentHistoryEntry) : lifecycle?.history || [],
    firstSeenAt: readOptionalString(value.first_seen_at, "incidents[].first_seen_at"),
    lastSeenAt: readOptionalString(value.last_seen_at, "incidents[].last_seen_at"),
    primaryReason: readString(value.primary_reason, "incidents[].primary_reason"),
    reasonCodes: readOptionalStringArray(value.reason_codes, "incidents[].reason_codes") || [],
    relatedReasons: readOptionalStringArray(value.related_reasons, "incidents[].related_reasons") || [],
    findingRefs: readOptionalStringArray(value.finding_refs, "incidents[].finding_refs") || [],
    guidanceRefs: readOptionalStringArray(value.guidance_refs, "incidents[].guidance_refs") || [],
    scorecardRefs: readOptionalStringArray(value.scorecard_refs, "incidents[].scorecard_refs") || [],
    metricLinks: Array.isArray(value.metric_links) ? value.metric_links.map(parseIncidentMetricLink) : [],
    affectedRepos: readOptionalStringArray(value.affected_repos, "incidents[].affected_repos") || [],
    affectedEnvironments: readOptionalStringArray(value.affected_environments, "incidents[].affected_environments") || [],
    affectedTenants: readOptionalStringArray(value.affected_tenants, "incidents[].affected_tenants") || [],
    affectedNamespaces: readOptionalStringArray(value.affected_namespaces, "incidents[].affected_namespaces") || [],
    affectedWorkloads: readOptionalStringArray(value.affected_workloads, "incidents[].affected_workloads") || [],
    affectedImages: readOptionalStringArray(value.affected_images, "incidents[].affected_images") || [],
    affectedComponents: readOptionalStringArray(value.affected_components, "incidents[].affected_components") || [],
    evidenceRefs: readOptionalStringArray(value.evidence_refs, "incidents[].evidence_refs") || [],
    evidencePack: parseIncidentEvidencePack(value.evidence_pack),
    governanceImpacts: value.governance_impacts.map(parseIncidentImpact),
    labels: readOptionalStringArray(value.labels, "incidents[].labels") || [],
    timeline: value.timeline.map(parseIncidentTimelineEntry),
    events: value.events.map(parseStoredEvent),
  };
}

function parseIncidentsResponse(value: unknown): InvestigationIncident[] {
  if (!isRecord(value) || !Array.isArray(value.incidents)) {
    throw new Error("Audit API returned invalid incidents response.");
  }
  return value.incidents.map(parseIncident);
}

function parseMetricIncidentsResponse(value: unknown): MetricIncidentDrilldown {
  if (!isRecord(value) || !Array.isArray(value.incidents)) {
    throw new Error("Audit API returned invalid metric drill-down response.");
  }
  return {
    metricKey: readString(value.metric_key, "metric_key"),
    metricLabel: readString(value.metric_label, "metric_label"),
    incidents: value.incidents.map(parseIncident),
    limitations: readOptionalStringArray(value.limitations, "limitations") || [],
  };
}

function parseIncidentExport(value: unknown): IncidentExport {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid incident export.");
  }
  return {
    generatedAt: readString(value.generated_at, "generated_at"),
    audience: (readOptionalString(value.audience, "audience") || "internal") as IncidentExport["audience"],
    redacted: value.redacted === undefined ? false : readBoolean(value.redacted, "redacted"),
    redactionSummary: readOptionalStringArray(value.redaction_summary, "redaction_summary") || [],
    incidentID: readString(value.incident_id, "incident_id"),
    identityKey: readOptionalString(value.identity_key, "identity_key"),
    title: readString(value.title, "title"),
    summary: readString(value.summary, "summary"),
    state: readString(value.state, "state") as IncidentExport["state"],
    severity: readString(value.severity, "severity") as IncidentExport["severity"],
    priority: readString(value.priority, "priority") as IncidentExport["priority"],
    owner: readOptionalString(value.owner, "owner"),
    openedAt: readOptionalString(value.opened_at, "opened_at"),
    updatedAt: readOptionalString(value.updated_at, "updated_at"),
    resolvedAt: readOptionalString(value.resolved_at, "resolved_at"),
    scopeType: readOptionalString(value.scope_type, "scope_type"),
    scopeRef: readOptionalString(value.scope_ref, "scope_ref"),
    tenantID: readOptionalString(value.tenant_id, "tenant_id"),
    clusterID: readOptionalString(value.cluster_id, "cluster_id"),
    environment: readOptionalString(value.environment, "environment"),
    repository: readOptionalString(value.repository, "repository"),
    governanceImpacts: Array.isArray(value.governance_impacts) ? value.governance_impacts.map(parseIncidentImpact) : [],
    reasonCodes: readOptionalStringArray(value.reason_codes, "reason_codes") || [],
    findingRefs: readOptionalStringArray(value.finding_refs, "finding_refs") || [],
    guidanceRefs: readOptionalStringArray(value.guidance_refs, "guidance_refs") || [],
    scorecardRefs: readOptionalStringArray(value.scorecard_refs, "scorecard_refs") || [],
    metricLinks: Array.isArray(value.metric_links) ? value.metric_links.map(parseIncidentMetricLink) : [],
    evidenceRefs: readOptionalStringArray(value.evidence_refs, "evidence_refs") || [],
    evidencePack: parseIncidentEvidencePack(value.evidence_pack),
    history: Array.isArray(value.history) ? value.history.map(parseIncidentHistoryEntry) : [],
    resolution: parseIncidentResolution(value.resolution),
    notes: Array.isArray(value.notes) ? value.notes.map(parseIncidentNote) : [],
    newActivityDetected:
      value.new_activity_detected === undefined ? false : readBoolean(value.new_activity_detected, "new_activity_detected"),
    relatedEventRefs: Array.isArray(value.related_event_refs) ? value.related_event_refs.map(parseIncidentExportEventRef) : [],
    limitations: readOptionalStringArray(value.limitations, "limitations") || [],
  };
}

function parseCountRecord(value: unknown, field: string): Record<string, number> {
  const record = readOptionalRecord(value, field) || {};
  const output: Record<string, number> = {};
  for (const [key, item] of Object.entries(record)) {
    output[key] = readNumber(item, `${field}.${key}`);
  }
  return output;
}

function parseIncidentPackageItem(value: unknown): IncidentPackage["incidents"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid incident package item.");
  }
  return {
    incidentID: readString(value.incident_id, "incidents[].incident_id"),
    title: readString(value.title, "incidents[].title"),
    summary: readString(value.summary, "incidents[].summary"),
    state: readString(value.state, "incidents[].state") as IncidentPackage["incidents"][number]["state"],
    severity: readString(value.severity, "incidents[].severity") as IncidentPackage["incidents"][number]["severity"],
    priority: readString(value.priority, "incidents[].priority") as IncidentPackage["incidents"][number]["priority"],
    category: readString(value.category, "incidents[].category"),
    scopeLabel: readOptionalString(value.scope_label, "incidents[].scope_label"),
    openedAt: readOptionalString(value.opened_at, "incidents[].opened_at"),
    updatedAt: readOptionalString(value.updated_at, "incidents[].updated_at"),
    resolvedAt: readOptionalString(value.resolved_at, "incidents[].resolved_at"),
  };
}

function parseIncidentPackage(value: unknown): IncidentPackage {
  if (!isRecord(value) || !Array.isArray(value.incidents) || !isRecord(value.aggregate) || !isRecord(value.package_intelligence)) {
    throw new Error("Audit API returned invalid incident package.");
  }
  return {
    generatedAt: readString(value.generated_at, "generated_at"),
    audience: (readOptionalString(value.audience, "audience") || "internal") as IncidentPackage["audience"],
    redacted: value.redacted === undefined ? false : readBoolean(value.redacted, "redacted"),
    redactionSummary: readOptionalStringArray(value.redaction_summary, "redaction_summary") || [],
    selectionMode: readString(value.selection_mode, "selection_mode") as IncidentPackage["selectionMode"],
    selectionSummary: readString(value.selection_summary, "selection_summary"),
    packageSummary: readString(value.package_summary, "package_summary"),
    incidentCount: readNumber(value.incident_count, "incident_count"),
    incidentRefs: readOptionalStringArray(value.incident_refs, "incident_refs") || [],
    aggregate: {
      byState: parseCountRecord((value.aggregate as Record<string, unknown>).by_state, "aggregate.by_state"),
      bySeverity: parseCountRecord((value.aggregate as Record<string, unknown>).by_severity, "aggregate.by_severity"),
      byCategory: parseCountRecord((value.aggregate as Record<string, unknown>).by_category, "aggregate.by_category"),
    },
    incidents: value.incidents.map(parseIncidentPackageItem),
    packageIntelligence: parsePackageIntelligence(value.package_intelligence),
    limitations: readOptionalStringArray(value.limitations, "limitations") || [],
  };
}

function parsePackageIntelligence(value: unknown): IncidentPackage["packageIntelligence"] {
  if (!isRecord(value) || !isRecord(value.defense_gap_summary) || !isRecord(value.policy_replay_summary) || !isRecord(value.systemic_weakness_summary) || !isRecord(value.recommended_actions)) {
    throw new Error("Audit API returned invalid package intelligence.");
  }
  const defenseGapSummary = value.defense_gap_summary as Record<string, unknown>;
  const policyReplaySummary = value.policy_replay_summary as Record<string, unknown>;
  const replayCurrentOutcome = readOptionalRecord(policyReplaySummary.current_outcome, "package_intelligence.policy_replay_summary.current_outcome") || {};
  const replayProposedOutcome = readOptionalRecord(policyReplaySummary.proposed_outcome, "package_intelligence.policy_replay_summary.proposed_outcome") || {};
  const replayDelta = readOptionalRecord(policyReplaySummary.delta, "package_intelligence.policy_replay_summary.delta") || {};
  const systemicWeaknessSummary = value.systemic_weakness_summary as Record<string, unknown>;
  const recommendedActions = value.recommended_actions as Record<string, unknown>;
  return {
    advisoryOnly: value.advisory_only === undefined ? true : readBoolean(value.advisory_only, "package_intelligence.advisory_only"),
    generatedAt: readString(value.generated_at, "package_intelligence.generated_at"),
    defenseGapSummary: {
      topGapTypes: readOptionalStringArray(defenseGapSummary.top_gap_types, "package_intelligence.defense_gap_summary.top_gap_types") || [],
      confidenceMix: parseCountRecord(defenseGapSummary.confidence_mix, "package_intelligence.defense_gap_summary.confidence_mix"),
      topFindings: readOptionalArray(defenseGapSummary.top_findings, "package_intelligence.defense_gap_summary.top_findings").map(parseDefenseGapFinding),
      rationale: readString(defenseGapSummary.rationale, "package_intelligence.defense_gap_summary.rationale"),
      limitations: readOptionalStringArray(defenseGapSummary.limitations, "package_intelligence.defense_gap_summary.limitations") || [],
    },
    policyReplaySummary: {
      currentOutcome: {
        blockingOrSurfacing: readNumber(replayCurrentOutcome.blocking_or_surfacing, "package_intelligence.policy_replay_summary.current_outcome.blocking_or_surfacing"),
        monitoringOnly: readNumber(replayCurrentOutcome.monitoring_only, "package_intelligence.policy_replay_summary.current_outcome.monitoring_only"),
        resolvedOrReviewed: readNumber(replayCurrentOutcome.resolved_or_reviewed, "package_intelligence.policy_replay_summary.current_outcome.resolved_or_reviewed"),
      },
      proposedOutcome: {
        earlierDenials: readNumber(replayProposedOutcome.earlier_denials, "package_intelligence.policy_replay_summary.proposed_outcome.earlier_denials"),
        evidenceHolds: readNumber(replayProposedOutcome.evidence_holds, "package_intelligence.policy_replay_summary.proposed_outcome.evidence_holds"),
        earlierContainment: readNumber(replayProposedOutcome.earlier_containment, "package_intelligence.policy_replay_summary.proposed_outcome.earlier_containment"),
        narrowerExceptions: readNumber(replayProposedOutcome.narrower_exceptions, "package_intelligence.policy_replay_summary.proposed_outcome.narrower_exceptions"),
      },
      delta: {
        additionalRejections: readNumber(replayDelta.additional_rejections, "package_intelligence.policy_replay_summary.delta.additional_rejections"),
        earlierContainmentPaths: readNumber(replayDelta.earlier_containment_paths, "package_intelligence.policy_replay_summary.delta.earlier_containment_paths"),
        impactedCases: readNumber(replayDelta.impacted_cases, "package_intelligence.policy_replay_summary.delta.impacted_cases"),
      },
      blastRadius: parseReplayBlastRadius(policyReplaySummary.blast_radius),
      topCoverageGaps: readOptionalArray(policyReplaySummary.top_coverage_gaps, "package_intelligence.policy_replay_summary.top_coverage_gaps").map(parseCoverageGapFinding),
      shadowModeImpact: readString(policyReplaySummary.shadow_mode_impact, "package_intelligence.policy_replay_summary.shadow_mode_impact"),
      limitations: readOptionalStringArray(policyReplaySummary.limitations, "package_intelligence.policy_replay_summary.limitations") || [],
    },
    systemicWeaknessSummary: {
      topPatterns: readOptionalArray(systemicWeaknessSummary.top_patterns, "package_intelligence.systemic_weakness_summary.top_patterns").map(parsePackageSystemicPattern),
      rootCauseHypothesis: readString(systemicWeaknessSummary.root_cause_hypothesis, "package_intelligence.systemic_weakness_summary.root_cause_hypothesis"),
      processFragility: readBoolean(systemicWeaknessSummary.process_fragility, "package_intelligence.systemic_weakness_summary.process_fragility"),
      supplyChainBlindSpots: readBoolean(systemicWeaknessSummary.supply_chain_blind_spots, "package_intelligence.systemic_weakness_summary.supply_chain_blind_spots"),
      executiveRecommendation: readString(systemicWeaknessSummary.executive_recommendation, "package_intelligence.systemic_weakness_summary.executive_recommendation"),
      limitations: readOptionalStringArray(systemicWeaknessSummary.limitations, "package_intelligence.systemic_weakness_summary.limitations") || [],
    },
    recommendedActions: {
      whyThisMattersNow: readString(recommendedActions.why_this_matters_now, "package_intelligence.recommended_actions.why_this_matters_now"),
      immediateContainment: readOptionalStringArray(recommendedActions.immediate_containment, "package_intelligence.recommended_actions.immediate_containment") || [],
      nearTermHardening: readOptionalStringArray(recommendedActions.near_term_hardening, "package_intelligence.recommended_actions.near_term_hardening") || [],
      governanceFix: readOptionalStringArray(recommendedActions.governance_fix, "package_intelligence.recommended_actions.governance_fix") || [],
    },
  };
}

function parsePackageSystemicPattern(value: unknown): IncidentPackage["packageIntelligence"]["systemicWeaknessSummary"]["topPatterns"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid package systemic pattern.");
  }
  return {
    patternKey: readString(value.pattern_key, "package_intelligence.systemic_weakness_summary.top_patterns[].pattern_key"),
    title: readString(value.title, "package_intelligence.systemic_weakness_summary.top_patterns[].title"),
    priority: readString(value.priority, "package_intelligence.systemic_weakness_summary.top_patterns[].priority") as IncidentPackage["packageIntelligence"]["systemicWeaknessSummary"]["topPatterns"][number]["priority"],
    relatedIncidentRefs: readOptionalStringArray(value.related_incident_refs, "package_intelligence.systemic_weakness_summary.top_patterns[].related_incident_refs") || [],
    evidenceRefs: readOptionalStringArray(value.evidence_refs, "package_intelligence.systemic_weakness_summary.top_patterns[].evidence_refs") || [],
  };
}

function parseAdvisoryReadbackRef(value: unknown): AdvisoryReadbackRef {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid readback ref.");
  }
  return {
    resourceType: readString(value.resource_type, "readback.resource_type") as AdvisoryReadbackRef["resourceType"],
    resourceID: readString(value.resource_id, "readback.resource_id"),
    resourceURI: readString(value.resource_uri, "readback.resource_uri"),
    evidenceHash: readString(value.evidence_hash, "readback.evidence_hash"),
  };
}

function parseDecisionEvidenceEnvelope(value: unknown): DecisionEvidenceEnvelope {
  if (!isRecord(value) || !isRecord(value.verdict_context) || !isRecord(value.snapshot_refs)) {
    throw new Error("Audit API returned invalid decision evidence envelope.");
  }
  const verdictContext = value.verdict_context as Record<string, unknown>;
  const snapshotRefs = value.snapshot_refs as Record<string, unknown>;
  return {
    schemaVersion: readString(value.schema_version, "evidence_envelope.schema_version"),
    resourceType: readString(value.resource_type, "evidence_envelope.resource_type"),
    resourceID: readString(value.resource_id, "evidence_envelope.resource_id"),
    evidenceHash: readString(value.evidence_hash, "evidence_envelope.evidence_hash"),
    generatedAt: readString(value.generated_at, "evidence_envelope.generated_at"),
    subjectType: readString(value.subject_type, "evidence_envelope.subject_type"),
    subjectRef: readString(value.subject_ref, "evidence_envelope.subject_ref"),
    verdictContext: {
      summary: readOptionalString(verdictContext.summary, "evidence_envelope.verdict_context.summary"),
      currentOutcome: readOptionalString(verdictContext.current_outcome, "evidence_envelope.verdict_context.current_outcome"),
      proposedOutcome: readOptionalString(verdictContext.proposed_outcome, "evidence_envelope.verdict_context.proposed_outcome"),
      delta: readOptionalString(verdictContext.delta, "evidence_envelope.verdict_context.delta"),
      patternKey: readOptionalString(verdictContext.pattern_key, "evidence_envelope.verdict_context.pattern_key"),
      gapTypes: readOptionalStringArray(verdictContext.gap_types, "evidence_envelope.verdict_context.gap_types") || [],
    },
    snapshotRefs: {
      policySnapshotRef: readString(snapshotRefs.policy_snapshot_ref, "evidence_envelope.snapshot_refs.policy_snapshot_ref"),
      evaluatorInputHash: readString(snapshotRefs.evaluator_input_hash, "evidence_envelope.snapshot_refs.evaluator_input_hash"),
      evaluatorOutputHash: readString(snapshotRefs.evaluator_output_hash, "evidence_envelope.snapshot_refs.evaluator_output_hash"),
      evidenceRefs: readOptionalStringArray(snapshotRefs.evidence_refs, "evidence_envelope.snapshot_refs.evidence_refs") || [],
    },
    redactionProfileVersion: readString(value.redaction_profile_version, "evidence_envelope.redaction_profile_version"),
    projectionSchemaVersion: readString(value.projection_schema_version, "evidence_envelope.projection_schema_version"),
    advisoryOnly: value.advisory_only === undefined ? true : readBoolean(value.advisory_only, "evidence_envelope.advisory_only"),
    limitations: readOptionalStringArray(value.limitations, "evidence_envelope.limitations") || [],
  };
}

function parseAdvisoryReadbackResponse<TPayload>(
  value: unknown,
  parsePayload: (payload: unknown) => TPayload,
): AdvisoryReadbackResponse<TPayload> {
  if (!isRecord(value) || !isRecord(value.evidence_envelope)) {
    throw new Error("Audit API returned invalid readback response.");
  }
  return {
    resourceType: readString(value.resource_type, "resource_type"),
    resourceID: readString(value.resource_id, "resource_id"),
    permanentURI: readString(value.permanent_uri, "permanent_uri"),
    projectionAudience: (readOptionalString(value.projection_audience, "projection_audience") || "internal") as IncidentExport["audience"],
    advisoryOnly: value.advisory_only === undefined ? true : readBoolean(value.advisory_only, "advisory_only"),
    payloadSummary: readString(value.payload_summary, "payload_summary"),
    evidenceEnvelope: parseDecisionEvidenceEnvelope(value.evidence_envelope),
    payload: parsePayload(value.payload),
    limitations: readOptionalStringArray(value.limitations, "limitations") || [],
  };
}

function parseAdvisoryShareGrant(value: unknown): AdvisoryShareGrant {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid advisory share grant.");
  }
  return {
    grantID: readString(value.grant_id, "grant_id"),
    shareURL: readString(value.share_url, "share_url"),
    resourceType: readString(value.resource_type, "resource_type"),
    resourceID: readString(value.resource_id, "resource_id"),
    audience: (readOptionalString(value.audience, "audience") || "internal") as AdvisoryShareGrant["audience"],
    expiresAt: readString(value.expires_at, "expires_at"),
    purpose: readOptionalString(value.purpose, "purpose"),
  };
}

function parseRecommendationActionTemplate(value: unknown): RecommendationActionTemplate {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid recommendation action template.");
  }
  return {
    templateID: readString(value.template_id, "recommendation_action.template_id"),
    title: readString(value.title, "recommendation_action.title"),
    description: readString(value.description, "recommendation_action.description"),
    recommendationType: readString(value.recommendation_type, "recommendation_action.recommendation_type"),
    approvalMode: (readString(value.approval_mode, "recommendation_action.approval_mode") as RecommendationActionTemplate["approvalMode"]),
    requiredInputs: readOptionalStringArray(value.required_inputs, "recommendation_action.required_inputs") || [],
    allowedAudiences: (readOptionalStringArray(value.allowed_audiences, "recommendation_action.allowed_audiences") || []) as RecommendationActionTemplate["allowedAudiences"],
    idempotent: readBoolean(value.idempotent, "recommendation_action.idempotent"),
    cancelSemantics: readString(value.cancel_semantics, "recommendation_action.cancel_semantics"),
  };
}

function parseRecommendationOutcome(value: unknown): Recommendation["outcome"] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid recommendation outcome.");
  }
  return {
    status: readString(value.status, "recommendation.outcome.status") as Recommendation["outcome"]["status"],
    summary: readOptionalString(value.summary, "recommendation.outcome.summary"),
    verifiedAt: readOptionalString(value.verified_at, "recommendation.outcome.verified_at"),
  };
}

function parseRecommendationComment(value: unknown): Recommendation["comments"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid recommendation comment.");
  }
  return {
    id: readString(value.id, "recommendation.comments[].id"),
    comment: readString(value.comment, "recommendation.comments[].comment"),
    actor: readOptionalString(value.actor, "recommendation.comments[].actor"),
    timestamp: readOptionalString(value.timestamp, "recommendation.comments[].timestamp"),
  };
}

function parseRecommendationHistoryEntry(value: unknown): Recommendation["history"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid recommendation history entry.");
  }
  return {
    id: readString(value.id, "recommendation.history[].id"),
    eventType: readString(value.event_type, "recommendation.history[].event_type"),
    title: readString(value.title, "recommendation.history[].title"),
    summary: readString(value.summary, "recommendation.history[].summary"),
    actor: readOptionalString(value.actor, "recommendation.history[].actor"),
    timestamp: readOptionalString(value.timestamp, "recommendation.history[].timestamp"),
  };
}

function parseRecommendation(value: unknown): Recommendation {
  if (!isRecord(value) || !isRecord(value.action_template) || !isRecord(value.outcome)) {
    throw new Error("Audit API returned invalid recommendation.");
  }
  return {
    recommendationID: readString(value.recommendation_id, "recommendation.recommendation_id"),
    sourceType: readString(value.source_type, "recommendation.source_type"),
    sourceRef: readString(value.source_ref, "recommendation.source_ref"),
    subjectType: readString(value.subject_type, "recommendation.subject_type"),
    subjectRef: readString(value.subject_ref, "recommendation.subject_ref"),
    team: readOptionalString(value.team, "recommendation.team"),
    service: readOptionalString(value.service, "recommendation.service"),
    repo: readOptionalString(value.repo, "recommendation.repo"),
    environment: readOptionalString(value.environment, "recommendation.environment"),
    recommendationType: readString(value.recommendation_type, "recommendation.recommendation_type"),
    title: readString(value.title, "recommendation.title"),
    description: readString(value.description, "recommendation.description"),
    recommendedAction: readString(value.recommended_action, "recommendation.recommended_action"),
    rationale: readString(value.rationale, "recommendation.rationale"),
    evidenceRefs: readOptionalStringArray(value.evidence_refs, "recommendation.evidence_refs") || [],
    readbackRefs: readOptionalArray(value.readback_refs, "recommendation.readback_refs").map(parseAdvisoryReadbackRef),
    relatedIncidentRefs: readOptionalStringArray(value.related_incident_refs, "recommendation.related_incident_refs") || [],
    priorityBand: readString(value.priority_band, "recommendation.priority_band") as Recommendation["priorityBand"],
    impactScore: readNumber(value.impact_score, "recommendation.impact_score"),
    effortScore: readNumber(value.effort_score, "recommendation.effort_score"),
    confidenceScore: readNumber(value.confidence_score, "recommendation.confidence_score"),
    approvalMode: readString(value.approval_mode, "recommendation.approval_mode") as Recommendation["approvalMode"],
    status: readString(value.status, "recommendation.status") as Recommendation["status"],
    createdAt: readString(value.created_at, "recommendation.created_at"),
    expiresAt: readOptionalString(value.expires_at, "recommendation.expires_at"),
    supersededBy: readOptionalString(value.superseded_by, "recommendation.superseded_by"),
    verificationPlan: readOptionalStringArray(value.verification_plan, "recommendation.verification_plan") || [],
    feedbackSummary: readOptionalString(value.feedback_summary, "recommendation.feedback_summary"),
    actionTemplate: parseRecommendationActionTemplate(value.action_template),
    owner: readOptionalString(value.owner, "recommendation.owner"),
    comments: readOptionalArray(value.comments, "recommendation.comments").map(parseRecommendationComment),
    history: readOptionalArray(value.history, "recommendation.history").map(parseRecommendationHistoryEntry),
    outcome: parseRecommendationOutcome(value.outcome),
    advisoryOnly: value.advisory_only === undefined ? true : readBoolean(value.advisory_only, "recommendation.advisory_only"),
    limitations: readOptionalStringArray(value.limitations, "recommendation.limitations") || [],
  };
}

function parseReplayBlastRadius(value: unknown): PolicyReplayAssessment["blastRadius"] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid replay blast radius.");
  }
  return {
    incidentCount: readNumber(value.incident_count, "blast_radius.incident_count"),
    repoCount: readNumber(value.repo_count, "blast_radius.repo_count"),
    environmentCount: readNumber(value.environment_count, "blast_radius.environment_count"),
    workloadCount: readNumber(value.workload_count, "blast_radius.workload_count"),
    topScopes: readOptionalStringArray(value.top_scopes, "blast_radius.top_scopes") || [],
  };
}

function parseDefenseGapFinding(value: unknown): DefenseGapAssessment["defenseGaps"][number] {
  if (!isRecord(value) || !isRecord(value.recommended_actions)) {
    throw new Error("Audit API returned invalid defense gap finding.");
  }
  return {
    gapType: readString(value.gap_type, "defense_gaps[].gap_type"),
    title: readString(value.title, "defense_gaps[].title"),
    confidence: readString(value.confidence, "defense_gaps[].confidence") as DefenseGapAssessment["defenseGaps"][number]["confidence"],
    whyItMatters: readString(value.why_it_matters, "defense_gaps[].why_it_matters"),
    evidenceRefs: readOptionalStringArray(value.evidence_refs, "defense_gaps[].evidence_refs") || [],
    relatedIncidentRefs: readOptionalStringArray(value.related_incident_refs, "defense_gaps[].related_incident_refs") || [],
    recommendedActions: {
      containment: readOptionalStringArray((value.recommended_actions as Record<string, unknown>).containment, "defense_gaps[].recommended_actions.containment") || [],
      hardening: readOptionalStringArray((value.recommended_actions as Record<string, unknown>).hardening, "defense_gaps[].recommended_actions.hardening") || [],
      governanceFix: readOptionalStringArray((value.recommended_actions as Record<string, unknown>).governance_fix, "defense_gaps[].recommended_actions.governance_fix") || [],
    },
  };
}

function parseDefenseGapAssessment(value: unknown): DefenseGapAssessment {
  if (!isRecord(value) || !isRecord(value.systemic_pattern) || !Array.isArray(value.defense_gaps)) {
    throw new Error("Audit API returned invalid defense gap assessment.");
  }
  return {
    assessmentID: readString(value.assessment_id, "assessment_id"),
    subjectType: readString(value.subject_type, "subject_type") as DefenseGapAssessment["subjectType"],
    subjectRef: readString(value.subject_ref, "subject_ref"),
    generatedAt: readString(value.generated_at, "generated_at"),
    advisoryOnly: value.advisory_only === undefined ? true : readBoolean(value.advisory_only, "advisory_only"),
    defenseGaps: value.defense_gaps.map(parseDefenseGapFinding),
    systemicPattern: {
      present: readBoolean((value.systemic_pattern as Record<string, unknown>).present, "systemic_pattern.present"),
      patternKey: readOptionalString((value.systemic_pattern as Record<string, unknown>).pattern_key, "systemic_pattern.pattern_key"),
      summary: readString((value.systemic_pattern as Record<string, unknown>).summary, "systemic_pattern.summary"),
      relatedIncidentRefs: readOptionalStringArray((value.systemic_pattern as Record<string, unknown>).related_incident_refs, "systemic_pattern.related_incident_refs") || [],
    },
    readback: parseAdvisoryReadbackRef(value.readback),
    limitations: readOptionalStringArray(value.limitations, "limitations") || [],
  };
}

function parseCoverageGapFinding(value: unknown): PolicyReplayAssessment["coverageGaps"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid coverage gap finding.");
  }
  return {
    gapType: readString(value.gap_type, "coverage_gaps[].gap_type"),
    title: readString(value.title, "coverage_gaps[].title"),
    summary: readString(value.summary, "coverage_gaps[].summary"),
    confidence: readString(value.confidence, "coverage_gaps[].confidence") as PolicyReplayAssessment["coverageGaps"][number]["confidence"],
    evidenceRefs: readOptionalStringArray(value.evidence_refs, "coverage_gaps[].evidence_refs") || [],
    relatedIncidentRefs: readOptionalStringArray(value.related_incident_refs, "coverage_gaps[].related_incident_refs") || [],
    recommendedAction: readString(value.recommended_action, "coverage_gaps[].recommended_action"),
  };
}

function parsePolicyReplayResult(value: unknown): PolicyReplayAssessment["replayResults"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid policy replay result.");
  }
  return {
    caseRef: readString(value.case_ref, "replay_results[].case_ref"),
    title: readString(value.title, "replay_results[].title"),
    currentOutcome: readString(value.current_outcome, "replay_results[].current_outcome"),
    proposedOutcome: readString(value.proposed_outcome, "replay_results[].proposed_outcome"),
    delta: readString(value.delta, "replay_results[].delta"),
    supportingEvidenceRefs: readOptionalStringArray(value.supporting_evidence_refs, "replay_results[].supporting_evidence_refs") || [],
    confidence: readString(value.confidence, "replay_results[].confidence") as PolicyReplayAssessment["replayResults"][number]["confidence"],
    limitations: readOptionalStringArray(value.limitations, "replay_results[].limitations") || [],
  };
}

function parsePolicyReplayAssessment(value: unknown): PolicyReplayAssessment {
  if (!isRecord(value) || !isRecord(value.blast_radius) || !Array.isArray(value.replay_results) || !Array.isArray(value.coverage_gaps)) {
    throw new Error("Audit API returned invalid policy replay assessment.");
  }
  return {
    assessmentID: readString(value.assessment_id, "assessment_id"),
    subjectType: readString(value.subject_type, "subject_type") as PolicyReplayAssessment["subjectType"],
    subjectRef: readString(value.subject_ref, "subject_ref"),
    generatedAt: readString(value.generated_at, "generated_at"),
    advisoryOnly: value.advisory_only === undefined ? true : readBoolean(value.advisory_only, "advisory_only"),
    shadowMode: value.shadow_mode === undefined ? true : readBoolean(value.shadow_mode, "shadow_mode"),
    replayResults: value.replay_results.map(parsePolicyReplayResult),
    coverageGaps: value.coverage_gaps.map(parseCoverageGapFinding),
    blastRadius: {
      incidentCount: readNumber((value.blast_radius as Record<string, unknown>).incident_count, "blast_radius.incident_count"),
      repoCount: readNumber((value.blast_radius as Record<string, unknown>).repo_count, "blast_radius.repo_count"),
      environmentCount: readNumber((value.blast_radius as Record<string, unknown>).environment_count, "blast_radius.environment_count"),
      workloadCount: readNumber((value.blast_radius as Record<string, unknown>).workload_count, "blast_radius.workload_count"),
      topScopes: readOptionalStringArray((value.blast_radius as Record<string, unknown>).top_scopes, "blast_radius.top_scopes") || [],
    },
    readback: parseAdvisoryReadbackRef(value.readback),
    limitations: readOptionalStringArray(value.limitations, "limitations") || [],
  };
}

function parseSystemicWeakness(value: unknown): SystemicWeaknessResponse["weaknesses"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid systemic weakness.");
  }
  return {
    patternKey: readString(value.pattern_key, "weaknesses[].pattern_key"),
    title: readString(value.title, "weaknesses[].title"),
    priority: readString(value.priority, "weaknesses[].priority") as SystemicWeaknessResponse["weaknesses"][number]["priority"],
    summary: readString(value.summary, "weaknesses[].summary"),
    processFragility: readOptionalStringArray(value.process_fragility, "weaknesses[].process_fragility") || [],
    supplyChainBlindSpots: readOptionalStringArray(value.supply_chain_blind_spots, "weaknesses[].supply_chain_blind_spots") || [],
    rootCauseHypothesis: readString(value.root_cause_hypothesis, "weaknesses[].root_cause_hypothesis"),
    executiveRecommendation: readString(value.executive_recommendation, "weaknesses[].executive_recommendation"),
    relatedIncidentRefs: readOptionalStringArray(value.related_incident_refs, "weaknesses[].related_incident_refs") || [],
    evidenceRefs: readOptionalStringArray(value.evidence_refs, "weaknesses[].evidence_refs") || [],
    readback: parseAdvisoryReadbackRef(value.readback),
    limitations: readOptionalStringArray(value.limitations, "weaknesses[].limitations") || [],
  };
}

function parseSystemicWeaknessResponse(value: unknown): SystemicWeaknessResponse {
  if (!isRecord(value) || !Array.isArray(value.weaknesses)) {
    throw new Error("Audit API returned invalid systemic weakness response.");
  }
  return {
    generatedAt: readString(value.generated_at, "generated_at"),
    advisoryOnly: value.advisory_only === undefined ? true : readBoolean(value.advisory_only, "advisory_only"),
    scopeSummary: readString(value.scope_summary, "scope_summary"),
    weaknesses: value.weaknesses.map(parseSystemicWeakness),
    limitations: readOptionalStringArray(value.limitations, "limitations") || [],
  };
}

function parseExecutiveStrategicGap(value: unknown): ExecutiveDefenseReport["strategicGaps"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid executive strategic gap.");
  }
  return {
    id: readString(value.id, "strategic_gaps[].id"),
    title: readString(value.title, "strategic_gaps[].title"),
    summary: readString(value.summary, "strategic_gaps[].summary"),
    investmentTarget: readString(value.investment_target, "strategic_gaps[].investment_target"),
    confidence: readString(value.confidence, "strategic_gaps[].confidence") as ExecutiveDefenseReport["strategicGaps"][number]["confidence"],
    relatedIncidentRefs: readOptionalStringArray(value.related_incident_refs, "strategic_gaps[].related_incident_refs") || [],
    evidenceRefs: readOptionalStringArray(value.evidence_refs, "strategic_gaps[].evidence_refs") || [],
  };
}

function parseExecutiveRiskTrend(value: unknown): ExecutiveDefenseReport["riskReductionTrends"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid executive risk trend.");
  }
  return {
    key: readString(value.key, "risk_reduction_trends[].key"),
    label: readString(value.label, "risk_reduction_trends[].label"),
    direction: readString(value.direction, "risk_reduction_trends[].direction") as ExecutiveDefenseReport["riskReductionTrends"][number]["direction"],
    value: readString(value.value, "risk_reduction_trends[].value"),
    summary: readString(value.summary, "risk_reduction_trends[].summary"),
    evidenceRefs: readOptionalStringArray(value.evidence_refs, "risk_reduction_trends[].evidence_refs") || [],
  };
}

function parseExecutiveShieldHealthComponent(value: unknown): ExecutiveDefenseReport["shieldHealth"]["components"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid shield health component.");
  }
  return {
    key: readString(value.key, "shield_health.components[].key"),
    label: readString(value.label, "shield_health.components[].label"),
    score: readNumber(value.score, "shield_health.components[].score"),
    summary: readString(value.summary, "shield_health.components[].summary"),
    evidenceRefs: readOptionalStringArray(value.evidence_refs, "shield_health.components[].evidence_refs") || [],
  };
}

function parseExecutiveImpactEstimate(value: unknown): ExecutiveDefenseReport["businessImpact"]["estimates"][number] {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid executive impact estimate.");
  }
  return {
    key: readString(value.key, "business_impact.estimates[].key"),
    label: readString(value.label, "business_impact.estimates[].label"),
    value: readString(value.value, "business_impact.estimates[].value"),
    confidence: readString(value.confidence, "business_impact.estimates[].confidence") as ExecutiveDefenseReport["businessImpact"]["estimates"][number]["confidence"],
    summary: readString(value.summary, "business_impact.estimates[].summary"),
    assumptions: readOptionalStringArray(value.assumptions, "business_impact.estimates[].assumptions") || [],
  };
}

function parseExecutiveDefenseReport(value: unknown): ExecutiveDefenseReport {
  if (!isRecord(value) || !isRecord(value.executive_summary) || !Array.isArray(value.strategic_gaps) || !Array.isArray(value.risk_reduction_trends) || !isRecord(value.shield_health) || !Array.isArray((value.shield_health as Record<string, unknown>).components) || !isRecord(value.business_impact) || !Array.isArray((value.business_impact as Record<string, unknown>).estimates) || !isRecord(value.board_package)) {
    throw new Error("Audit API returned invalid executive defense report.");
  }
  return {
    generatedAt: readString(value.generated_at, "generated_at"),
    audience: (readOptionalString(value.audience, "audience") || "internal") as ExecutiveDefenseReport["audience"],
    redacted: value.redacted === undefined ? false : readBoolean(value.redacted, "redacted"),
    redactionSummary: readOptionalStringArray(value.redaction_summary, "redaction_summary") || [],
    advisoryOnly: value.advisory_only === undefined ? true : readBoolean(value.advisory_only, "advisory_only"),
    selectionMode: readString(value.selection_mode, "selection_mode") as ExecutiveDefenseReport["selectionMode"],
    scopeSummary: readString(value.scope_summary, "scope_summary"),
    incidentCount: readNumber(value.incident_count, "incident_count"),
    incidentRefs: readOptionalStringArray(value.incident_refs, "incident_refs") || [],
    executiveSummary: {
      topRisks: readOptionalStringArray((value.executive_summary as Record<string, unknown>).top_risks, "executive_summary.top_risks") || [],
      topImprovements: readOptionalStringArray((value.executive_summary as Record<string, unknown>).top_improvements, "executive_summary.top_improvements") || [],
      trendChange: readString((value.executive_summary as Record<string, unknown>).trend_change, "executive_summary.trend_change"),
      whatMattersNow: readString((value.executive_summary as Record<string, unknown>).what_matters_now, "executive_summary.what_matters_now"),
    },
    strategicGaps: value.strategic_gaps.map(parseExecutiveStrategicGap),
    riskReductionTrends: value.risk_reduction_trends.map(parseExecutiveRiskTrend),
    shieldHealth: {
      score: readNumber((value.shield_health as Record<string, unknown>).score, "shield_health.score"),
      band: readString((value.shield_health as Record<string, unknown>).band, "shield_health.band") as ExecutiveDefenseReport["shieldHealth"]["band"],
      summary: readString((value.shield_health as Record<string, unknown>).summary, "shield_health.summary"),
      components: ((value.shield_health as Record<string, unknown>).components as unknown[]).map(parseExecutiveShieldHealthComponent),
    },
    businessImpact: {
      summary: readString((value.business_impact as Record<string, unknown>).summary, "business_impact.summary"),
      estimates: ((value.business_impact as Record<string, unknown>).estimates as unknown[]).map(parseExecutiveImpactEstimate),
    },
    boardPackage: {
      headline: readString((value.board_package as Record<string, unknown>).headline, "board_package.headline"),
      narrative: readString((value.board_package as Record<string, unknown>).narrative, "board_package.narrative"),
      investmentPriorities: readOptionalStringArray((value.board_package as Record<string, unknown>).investment_priorities, "board_package.investment_priorities") || [],
      nextQuarterPriorities: readOptionalStringArray((value.board_package as Record<string, unknown>).next_quarter_priorities, "board_package.next_quarter_priorities") || [],
      packageSummary: readString((value.board_package as Record<string, unknown>).package_summary, "board_package.package_summary"),
    },
    limitations: readOptionalStringArray(value.limitations, "limitations") || [],
  };
}

function parseRuntimeDriftFinding(value: unknown): RuntimeDriftFinding {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid runtime drift finding.");
  }
  return {
    id: readString(value.id, "runtime_drift.id"),
    tenant_id: readOptionalString(value.tenant_id, "runtime_drift.tenant_id"),
    cluster_id: readOptionalString(value.cluster_id, "runtime_drift.cluster_id"),
    namespace: readString(value.namespace, "runtime_drift.namespace"),
    workload_kind: readString(value.workload_kind, "runtime_drift.workload_kind"),
    workload: readString(value.workload, "runtime_drift.workload"),
    service_account: readOptionalString(value.service_account, "runtime_drift.service_account"),
    drift_result: readString(value.drift_result, "runtime_drift.drift_result"),
    drift_classes: readOptionalStringArray(value.drift_classes, "runtime_drift.drift_classes"),
    drift_severity: readOptionalString(value.drift_severity, "runtime_drift.drift_severity"),
    remediation_mode: readOptionalString(value.remediation_mode, "runtime_drift.remediation_mode"),
    remediation_attempt:
      value.remediation_attempt === undefined ? undefined : readNumber(value.remediation_attempt, "runtime_drift.remediation_attempt"),
    remediable: readBoolean(value.remediable, "runtime_drift.remediable"),
    status: readString(value.status, "runtime_drift.status"),
    quarantine_reason: readOptionalString(value.quarantine_reason, "runtime_drift.quarantine_reason"),
    desired_state_verification_state: readOptionalString(
      value.desired_state_verification_state,
      "runtime_drift.desired_state_verification_state",
    ),
    detected_at: readString(value.detected_at, "runtime_drift.detected_at"),
    last_updated_at: readString(value.last_updated_at, "runtime_drift.last_updated_at"),
    last_event_type: readString(value.last_event_type, "runtime_drift.last_event_type"),
    reasons: readOptionalStringArray(value.reasons, "runtime_drift.reasons"),
    evidence: readOptionalRecord(value.evidence, "runtime_drift.evidence"),
  };
}

function parseRuntimeDriftFindingsResponse(value: unknown): RuntimeDriftFindingsResponse {
  if (!isRecord(value) || !Array.isArray(value.items)) {
    throw new Error("Audit API returned invalid runtime drift findings response.");
  }
  return {
    items: value.items.map(parseRuntimeDriftFinding),
  };
}

function parseRuntimeDriftStatus(value: unknown): RuntimeDriftStatus {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid runtime drift status.");
  }
  const countsBySeverity = readOptionalRecord(value.counts_by_severity, "counts_by_severity") || {};
  const countsByStatus = readOptionalRecord(value.counts_by_status, "counts_by_status") || {};
  return {
    total_findings: readNumber(value.total_findings, "total_findings"),
    active_findings: readNumber(value.active_findings, "active_findings"),
    quarantined: readNumber(value.quarantined, "quarantined"),
    failed: readNumber(value.failed, "failed"),
    remediated: readNumber(value.remediated, "remediated"),
    detected: readNumber(value.detected, "detected"),
    counts_by_severity: Object.fromEntries(
      Object.entries(countsBySeverity).map(([key, count]) => [key, readNumber(count, `counts_by_severity.${key}`)]),
    ),
    counts_by_status: Object.fromEntries(
      Object.entries(countsByStatus).map(([key, count]) => [key, readNumber(count, `counts_by_status.${key}`)]),
    ),
    last_detected_at: readOptionalString(value.last_detected_at, "last_detected_at"),
    last_updated_at: readOptionalString(value.last_updated_at, "last_updated_at"),
  };
}

function parseRuntimeActiveState(value: unknown): RuntimeActiveState {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid runtime active state.");
  }
  return {
    id: readString(value.id, "runtime_active_state.id"),
    tenant_id: readOptionalString(value.tenant_id, "runtime_active_state.tenant_id"),
    cluster_id: readOptionalString(value.cluster_id, "runtime_active_state.cluster_id"),
    namespace: readString(value.namespace, "runtime_active_state.namespace"),
    workload_kind: readString(value.workload_kind, "runtime_active_state.workload_kind"),
    workload: readString(value.workload, "runtime_active_state.workload"),
    service_account: readOptionalString(value.service_account, "runtime_active_state.service_account"),
    observed_digest: readOptionalString(value.observed_digest, "runtime_active_state.observed_digest"),
    approved_digest: readOptionalString(value.approved_digest, "runtime_active_state.approved_digest"),
    observed_config_hash: readOptionalString(value.observed_config_hash, "runtime_active_state.observed_config_hash"),
    expected_config_hash: readOptionalString(value.expected_config_hash, "runtime_active_state.expected_config_hash"),
    drift_result: readOptionalString(value.drift_result, "runtime_active_state.drift_result"),
    drift_classes: readOptionalStringArray(value.drift_classes, "runtime_active_state.drift_classes"),
    drift_severity: readOptionalString(value.drift_severity, "runtime_active_state.drift_severity"),
    reconciliation_status: readString(value.reconciliation_status, "runtime_active_state.reconciliation_status"),
    remediation_mode: readOptionalString(value.remediation_mode, "runtime_active_state.remediation_mode"),
    remediation_attempt:
      value.remediation_attempt === undefined ? undefined : readNumber(value.remediation_attempt, "runtime_active_state.remediation_attempt"),
    remediable: readBoolean(value.remediable, "runtime_active_state.remediable"),
    quarantine_reason: readOptionalString(value.quarantine_reason, "runtime_active_state.quarantine_reason"),
    quarantine_type: readOptionalString(value.quarantine_type, "runtime_active_state.quarantine_type"),
    protected_target:
      value.protected_target === undefined ? undefined : readBoolean(value.protected_target, "runtime_active_state.protected_target"),
    protected_reason: readOptionalString(value.protected_reason, "runtime_active_state.protected_reason"),
    desired_state_source_ref: readOptionalString(value.desired_state_source_ref, "runtime_active_state.desired_state_source_ref"),
    desired_state_approval_id: readOptionalString(value.desired_state_approval_id, "runtime_active_state.desired_state_approval_id"),
    desired_state_verification_state: readOptionalString(
      value.desired_state_verification_state,
      "runtime_active_state.desired_state_verification_state",
    ),
    last_error: readOptionalString(value.last_error, "runtime_active_state.last_error"),
    last_reconciled_at: readString(value.last_reconciled_at, "runtime_active_state.last_reconciled_at"),
    reasons: readOptionalStringArray(value.reasons, "runtime_active_state.reasons"),
    evidence: readOptionalRecord(value.evidence, "runtime_active_state.evidence"),
  };
}

function parseRuntimeActiveStatesResponse(value: unknown): RuntimeActiveStatesResponse {
  if (!isRecord(value) || !Array.isArray(value.items)) {
    throw new Error("Audit API returned invalid runtime active states response.");
  }
  return {
    items: value.items.map(parseRuntimeActiveState),
  };
}

function parseRuntimeClosedLoopStatus(value: unknown): RuntimeClosedLoopStatus {
  if (!isRecord(value)) {
    throw new Error("Audit API returned invalid runtime closed-loop status.");
  }
  const countsByStatus = readOptionalRecord(value.counts_by_status, "counts_by_status") || {};
  const countsByQuarantine = readOptionalRecord(value.counts_by_quarantine_type, "counts_by_quarantine_type") || {};
  return {
    total_targets: readNumber(value.total_targets, "total_targets"),
    in_sync: readNumber(value.in_sync, "in_sync"),
    drift_detected: readNumber(value.drift_detected, "drift_detected"),
    remediated: readNumber(value.remediated, "remediated"),
    failed: readNumber(value.failed, "failed"),
    quarantined: readNumber(value.quarantined, "quarantined"),
    protected_blocked: readNumber(value.protected_blocked, "protected_blocked"),
    counts_by_status: Object.fromEntries(
      Object.entries(countsByStatus).map(([key, count]) => [key, readNumber(count, `counts_by_status.${key}`)]),
    ),
    counts_by_quarantine_type: Object.fromEntries(
      Object.entries(countsByQuarantine).map(([key, count]) => [key, readNumber(count, `counts_by_quarantine_type.${key}`)]),
    ),
    last_reconciled_at: readOptionalString(value.last_reconciled_at, "last_reconciled_at"),
  };
}

export async function getHealth() {
  return fetchJSON<AuditHealth>("/health");
}

export async function getAuthStatus() {
  return parseAuthStatus(await fetchJSON<unknown>("/v1/auth/me"));
}

export async function getSyncStatus() {
  return parseSyncStatus(await fetchJSON<unknown>("/v1/sync/status"));
}

export async function getSummary(filters: Pick<EventFilters, "environment" | "tenant_id">) {
  return parseSummary(await fetchJSON<unknown>("/v1/reports/summary", { params: filters }));
}

export async function getEvents(tab: TabKey, filters: EventFilters) {
  const params: Record<string, string | undefined> = {
    component: filters.component,
    repo: filters.repo,
    environment: filters.environment,
    tenant_id: filters.tenant_id,
    limit: filters.limit,
  };

  if (tab === "events" && filters.decision) {
    params.decision = filters.decision;
  }

  const path =
    tab === "denies"
      ? "/v1/reports/denies"
      : tab === "runtime"
        ? "/v1/reports/runtime-drift"
        : "/v1/reports/events";

  return parseEventsResponse(await fetchJSON<unknown>(path, { params }));
}

export async function getIncidents(filters: EventFilters) {
  return parseIncidentsResponse(await fetchJSON<unknown>("/v1/incidents", {
    params: {
      decision: filters.decision,
      component: filters.component,
      repo: filters.repo,
      environment: filters.environment,
      tenant_id: filters.tenant_id,
      limit: filters.limit,
    },
  }));
}

export async function getMetricIncidents(metricKey: string, filters: EventFilters) {
  return parseMetricIncidentsResponse(await fetchJSON<unknown>(`/v1/scorecard/metrics/${encodeURIComponent(metricKey)}/incidents`, {
    params: {
      decision: filters.decision,
      component: filters.component,
      repo: filters.repo,
      environment: filters.environment,
      tenant_id: filters.tenant_id,
      limit: filters.limit,
    },
  }));
}

export async function getIncidentExport(
  incidentID: string,
  filters?: Pick<EventFilters, "environment" | "tenant_id" | "repo">,
  audience?: IncidentExport["audience"],
) {
  return parseIncidentExport(await fetchJSON<unknown>(`/v1/incidents/${encodeURIComponent(incidentID)}/export`, {
    params: {
      environment: filters?.environment,
      tenant_id: filters?.tenant_id,
      repo: filters?.repo,
      audience,
    },
  }));
}

export async function getIncidentPackage(
  filters: Pick<EventFilters, "environment" | "tenant_id" | "repo"> & {
    state?: string;
    severity?: string;
    category?: string;
    scorecard_ref?: string;
  },
  incidentIDs: string[],
  audience: IncidentPackage["audience"],
) {
  return parseIncidentPackage(await fetchJSON<unknown>("/v1/incidents/package", {
    params: {
      environment: filters.environment,
      tenant_id: filters.tenant_id,
      repo: filters.repo,
      state: filters.state,
      severity: filters.severity,
      category: filters.category,
      scorecard_ref: filters.scorecard_ref,
      audience,
      incident_id: incidentIDs,
    },
  }));
}

export async function getRecommendations(
  filters: Pick<EventFilters, "environment" | "tenant_id" | "repo" | "limit"> & {
    team?: string;
    service?: string;
    source_type?: string;
    subject_type?: string;
    recommendation_type?: string;
    status?: string;
  },
  options?: {
    incidentIDs?: string[];
    packageIncidentIDs?: string[];
  },
) {
  const payload = await fetchJSON<unknown>("/v1/recommendations", {
    params: {
      environment: filters.environment,
      tenant_id: filters.tenant_id,
      repo: filters.repo,
      limit: filters.limit,
      team: filters.team,
      service: filters.service,
      source_type: filters.source_type,
      subject_type: filters.subject_type,
      recommendation_type: filters.recommendation_type,
      status: filters.status,
      incident_id: options?.incidentIDs,
      package_incident_id: options?.packageIncidentIDs,
    },
  });
  if (!isRecord(payload) || !Array.isArray(payload.recommendations)) {
    throw new Error("Audit API returned invalid recommendations response.");
  }
  return payload.recommendations.map(parseRecommendation);
}

export async function getRecommendation(
  recommendationID: string,
  filters?: Pick<EventFilters, "environment" | "tenant_id" | "repo"> & {
    team?: string;
    service?: string;
  },
) {
  return parseRecommendation(await fetchJSON<unknown>(`/v1/recommendations/${encodeURIComponent(recommendationID)}`, {
    params: {
      environment: filters?.environment,
      tenant_id: filters?.tenant_id,
      repo: filters?.repo,
      team: filters?.team,
      service: filters?.service,
    },
  }));
}

export async function getRecommendationActions() {
  const payload = await fetchJSON<unknown>("/v1/recommendation-actions");
  if (!isRecord(payload) || !Array.isArray(payload.templates)) {
    throw new Error("Audit API returned invalid recommendation actions response.");
  }
  return payload.templates.map(parseRecommendationActionTemplate);
}

export async function getIncidentDefenseGaps(
  incidentID: string,
  filters?: Pick<EventFilters, "environment" | "tenant_id" | "repo">,
) {
  return parseDefenseGapAssessment(await fetchJSON<unknown>(`/v1/incidents/${encodeURIComponent(incidentID)}/defense-gaps`, {
    params: {
      environment: filters?.environment,
      tenant_id: filters?.tenant_id,
      repo: filters?.repo,
    },
  }));
}

export async function getMetricDefenseGaps(metricKey: string, filters: EventFilters) {
  return parseDefenseGapAssessment(await fetchJSON<unknown>(`/v1/scorecard/metrics/${encodeURIComponent(metricKey)}/defense-gaps`, {
    params: {
      decision: filters.decision,
      component: filters.component,
      repo: filters.repo,
      environment: filters.environment,
      tenant_id: filters.tenant_id,
      limit: filters.limit,
    },
  }));
}

export async function getIncidentPolicyReplay(
  incidentID: string,
  filters?: Pick<EventFilters, "environment" | "tenant_id" | "repo">,
) {
  return parsePolicyReplayAssessment(await fetchJSON<unknown>(`/v1/incidents/${encodeURIComponent(incidentID)}/policy-replay`, {
    params: {
      environment: filters?.environment,
      tenant_id: filters?.tenant_id,
      repo: filters?.repo,
    },
  }));
}

export async function getMetricPolicyReplay(metricKey: string, filters: EventFilters) {
  return parsePolicyReplayAssessment(await fetchJSON<unknown>(`/v1/scorecard/metrics/${encodeURIComponent(metricKey)}/policy-replay`, {
    params: {
      decision: filters.decision,
      component: filters.component,
      repo: filters.repo,
      environment: filters.environment,
      tenant_id: filters.tenant_id,
      limit: filters.limit,
    },
  }));
}

export async function getSystemicWeaknesses(filters: EventFilters & { scorecard_ref?: string }) {
  return parseSystemicWeaknessResponse(await fetchJSON<unknown>("/v1/ai/systemic-weaknesses", {
    params: {
      decision: filters.decision,
      component: filters.component,
      repo: filters.repo,
      environment: filters.environment,
      tenant_id: filters.tenant_id,
      limit: filters.limit,
      scorecard_ref: filters.scorecard_ref,
    },
  }));
}

export async function getDefenseGapReadback(resourceID: string) {
  return parseAdvisoryReadbackResponse(
    await fetchJSON<unknown>(`/v1/readback/defense-gap/${encodeURIComponent(resourceID)}`),
    parseDefenseGapAssessment,
  );
}

export async function getPolicyReplayReadback(resourceID: string) {
  return parseAdvisoryReadbackResponse(
    await fetchJSON<unknown>(`/v1/readback/policy-replay/${encodeURIComponent(resourceID)}`),
    parsePolicyReplayAssessment,
  );
}

export async function getSystemicWeaknessReadback(resourceID: string) {
  return parseAdvisoryReadbackResponse(
    await fetchJSON<unknown>(`/v1/readback/systemic-weakness/${encodeURIComponent(resourceID)}`),
    parseSystemicWeakness,
  );
}

export async function createAdvisoryReadbackGrant(
  resourceType: "defense-gap" | "policy-replay" | "systemic-weakness",
  resourceID: string,
  audience: IncidentExport["audience"],
  purpose?: string,
  expiresInMinutes = 60,
) {
  return parseAdvisoryShareGrant(await fetchJSON<unknown>("/v1/readback/grants", {
    method: "POST",
    body: {
      resource_type: resourceType,
      resource_id: resourceID,
      audience,
      expires_in_minutes: expiresInMinutes,
      purpose,
    },
  }));
}

export async function getExecutiveDefenseReport(
  filters: Pick<EventFilters, "environment" | "tenant_id" | "repo"> & {
    state?: string;
    severity?: string;
    category?: string;
    scorecard_ref?: string;
  },
  incidentIDs: string[],
  audience: ExecutiveDefenseReport["audience"],
) {
  return parseExecutiveDefenseReport(await fetchJSON<unknown>("/v1/ai/executive-defense-report", {
    params: {
      environment: filters.environment,
      tenant_id: filters.tenant_id,
      repo: filters.repo,
      state: filters.state,
      severity: filters.severity,
      category: filters.category,
      scorecard_ref: filters.scorecard_ref,
      audience,
      incident_id: incidentIDs,
    },
  }));
}

export async function acknowledgeRecommendation(recommendationID: string) {
  return parseRecommendation(await fetchJSON<unknown>(`/v1/recommendations/${encodeURIComponent(recommendationID)}/acknowledge`, {
    method: "POST",
    body: {},
  }));
}

export async function acceptRecommendation(recommendationID: string) {
  return parseRecommendation(await fetchJSON<unknown>(`/v1/recommendations/${encodeURIComponent(recommendationID)}/accept`, {
    method: "POST",
    body: {},
  }));
}

export async function rejectRecommendation(recommendationID: string, reason: string) {
  return parseRecommendation(await fetchJSON<unknown>(`/v1/recommendations/${encodeURIComponent(recommendationID)}/reject`, {
    method: "POST",
    body: { reason },
  }));
}

export async function executeRecommendation(
  recommendationID: string,
  input?: {
    template_id?: string;
    summary?: string;
  },
) {
  return parseRecommendation(await fetchJSON<unknown>(`/v1/recommendations/${encodeURIComponent(recommendationID)}/execute`, {
    method: "POST",
    body: input || {},
  }));
}

export async function verifyRecommendation(recommendationID: string) {
  return parseRecommendation(await fetchJSON<unknown>(`/v1/recommendations/${encodeURIComponent(recommendationID)}/verify`, {
    method: "POST",
    body: {},
  }));
}

export async function assignRecommendation(recommendationID: string, owner: string, reason?: string) {
  return parseRecommendation(await fetchJSON<unknown>(`/v1/recommendations/${encodeURIComponent(recommendationID)}/assign`, {
    method: "POST",
    body: { owner, reason },
  }));
}

export async function commentRecommendation(recommendationID: string, comment: string) {
  return parseRecommendation(await fetchJSON<unknown>(`/v1/recommendations/${encodeURIComponent(recommendationID)}/comment`, {
    method: "POST",
    body: { comment },
  }));
}

export async function requestRecommendationApproval(recommendationID: string, summary?: string) {
  return parseRecommendation(await fetchJSON<unknown>(`/v1/recommendations/${encodeURIComponent(recommendationID)}/approval-request`, {
    method: "POST",
    body: summary ? { summary } : {},
  }));
}

export async function acknowledgeIncident(incidentID: string, summary?: string) {
  return parseIncident(await fetchJSON<unknown>(`/v1/incidents/${encodeURIComponent(incidentID)}/acknowledge`, {
    method: "POST",
    body: summary ? { summary } : {},
  }));
}

export async function watchIncident(incidentID: string, summary?: string) {
  return parseIncident(await fetchJSON<unknown>(`/v1/incidents/${encodeURIComponent(incidentID)}/watch`, {
    method: "POST",
    body: summary ? { summary } : {},
  }));
}

export async function assignIncident(incidentID: string, owner: string, reason: string) {
  return parseIncident(await fetchJSON<unknown>(`/v1/incidents/${encodeURIComponent(incidentID)}/assign`, {
    method: "POST",
    body: { owner, reason },
  }));
}

export async function resolveIncident(
  incidentID: string,
  input: {
    resolution_type: string;
    resolution_summary: string;
    resolution_details?: string;
    resolution_refs?: string[];
    follow_up_required?: boolean;
  },
) {
  return parseIncident(await fetchJSON<unknown>(`/v1/incidents/${encodeURIComponent(incidentID)}/resolve`, {
    method: "POST",
    body: input,
  }));
}

export async function reopenIncident(incidentID: string, reason?: string) {
  return parseIncident(await fetchJSON<unknown>(`/v1/incidents/${encodeURIComponent(incidentID)}/reopen`, {
    method: "POST",
    body: reason ? { reason } : {},
  }));
}

export async function addIncidentNote(incidentID: string, note: string) {
  return parseIncident(await fetchJSON<unknown>(`/v1/incidents/${encodeURIComponent(incidentID)}/notes`, {
    method: "POST",
    body: { note },
  }));
}

export async function getExceptionReport(filters: {
  status?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
}) {
  return parseExceptionReport(await fetchJSON<unknown>("/v1/reports/exceptions", { params: filters }));
}

export async function getExceptions(filters: {
  status?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
  limit?: string;
}) {
  return parseExceptionsResponse(await fetchJSON<unknown>("/v1/exceptions", { params: filters }));
}

export async function requestException(input: ExceptionRequestInput) {
  return parseExceptionActionResponse(await fetchJSON<unknown>("/v1/exceptions/request", { method: "POST", body: input }));
}

export async function approveException(exceptionID: string, reason?: string) {
  return parseExceptionActionResponse(
    await fetchJSON<unknown>(`/v1/exceptions/${encodeURIComponent(exceptionID)}/approve`, {
      method: "POST",
      body: { reason: reason || "" },
    }),
  );
}

export async function rejectException(exceptionID: string, reason: string) {
  return parseExceptionActionResponse(
    await fetchJSON<unknown>(`/v1/exceptions/${encodeURIComponent(exceptionID)}/reject`, {
      method: "POST",
      body: { reason },
    }),
  );
}

export async function revokeException(exceptionID: string) {
  return parseExceptionActionResponse(
    await fetchJSON<unknown>(`/v1/exceptions/${encodeURIComponent(exceptionID)}`, {
      method: "DELETE",
    }),
  );
}

export async function getTrends(filters: {
  window?: string;
  window_days?: string;
  compare_to?: string;
  group_by?: string;
  metric?: string;
  granularity?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
  service?: string;
  team?: string;
  subject?: string;
  event_type?: string;
}) {
  return parseTrendsResponse(await fetchJSON<unknown>("/v1/analytics/trends", { params: filters }));
}

export async function getAnalyticsDelta(filters: {
  window?: string;
  window_days?: string;
  compare_to?: string;
  group_by?: string;
  metric?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
  service?: string;
  team?: string;
  subject?: string;
}) {
  return parseAnalyticsDeltaResponse(await fetchJSON<unknown>("/v1/analytics/delta", { params: filters }));
}

export async function getAnalyticsAnomalies(filters: {
  window?: string;
  window_days?: string;
  compare_to?: string;
  group_by?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
  service?: string;
  team?: string;
  subject?: string;
}) {
  return parseAnalyticsAnomaliesResponse(await fetchJSON<unknown>("/v1/analytics/anomalies", { params: filters }));
}

export async function getAnalyticsScorecards(filters: {
  window?: string;
  window_days?: string;
  compare_to?: string;
  group_by?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
  service?: string;
  team?: string;
  subject?: string;
}) {
  return parseAnalyticsScorecardsResponse(await fetchJSON<unknown>("/v1/analytics/scorecards", { params: filters }));
}

export async function getAnalyticsSegments(filters: {
  window?: string;
  window_days?: string;
  compare_to?: string;
  group_by?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
  service?: string;
  team?: string;
  subject?: string;
}) {
  return parseAnalyticsSegmentsResponse(await fetchJSON<unknown>("/v1/analytics/segments", { params: filters }));
}

export async function getTopologyServices(filters: {
  window?: string;
  window_days?: string;
  compare_to?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
  namespace?: string;
  service?: string;
  workload?: string;
  limit?: string;
}) {
  return parseTopologyServicesResponse(await fetchJSON<unknown>("/v1/topology/services", { params: filters }));
}

export async function getTopologyGraph(filters: {
  window?: string;
  window_days?: string;
  compare_to?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
  namespace?: string;
  service?: string;
  workload?: string;
  limit?: string;
}) {
  return parseTopologyGraphResponse(await fetchJSON<unknown>("/v1/topology/graph", { params: filters }));
}

export async function getTopologyHeatmap(filters: {
  window?: string;
  window_days?: string;
  compare_to?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
  namespace?: string;
  service?: string;
  workload?: string;
  limit?: string;
}) {
  return parseTopologyHeatmapResponse(await fetchJSON<unknown>("/v1/topology/heatmap", { params: filters }));
}

export async function getTopologyDelta(filters: {
  window?: string;
  window_days?: string;
  compare_to?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
  namespace?: string;
  service?: string;
  workload?: string;
  limit?: string;
}) {
  return parseTopologyDeltaResponse(await fetchJSON<unknown>("/v1/topology/delta", { params: filters }));
}

export async function getForensicsState(filters: {
  tenant_id?: string;
  environment?: string;
  repo?: string;
  service?: string;
  workload?: string;
  incident_id?: string;
  image_digest?: string;
  cve_id?: string;
  timestamp?: string;
  limit?: string;
}) {
  return fetchJSON<PointInTimeState>("/v1/forensics/state", { params: filters });
}

export async function getForensicsDelta(filters: {
  tenant_id?: string;
  environment?: string;
  repo?: string;
  service?: string;
  workload?: string;
  incident_id?: string;
  image_digest?: string;
  cve_id?: string;
  t1?: string;
  t2?: string;
  limit?: string;
}) {
  return fetchJSON<TimeDeltaResult>("/v1/forensics/delta", { params: filters });
}

export async function getForensicsTimeline(filters: {
  tenant_id?: string;
  environment?: string;
  repo?: string;
  service?: string;
  workload?: string;
  incident_id?: string;
  image_digest?: string;
  cve_id?: string;
  t1?: string;
  t2?: string;
  limit?: string;
}) {
  return fetchJSON<ForensicTimelineResponse>("/v1/forensics/timeline", { params: filters });
}

export async function getForensicsVEXFlashback(filters: {
  tenant_id?: string;
  environment?: string;
  repo?: string;
  service?: string;
  workload?: string;
  incident_id?: string;
  image_digest?: string;
  cve_id?: string;
  timestamp?: string;
  limit?: string;
}) {
  return fetchJSON<VEXFlashbackResponse>("/v1/forensics/vex-flashback", { params: filters });
}

export async function runForensicsReplay(
  filters: {
    tenant_id?: string;
    environment?: string;
    repo?: string;
  },
  body: {
    timestamp: string;
    replay_mode: string;
    incident_id?: string;
    service?: string;
    workload?: string;
    image_digest?: string;
    cve_id?: string;
  },
) {
  return fetchJSON<ForensicReplayResponse>("/v1/forensics/replay", {
    method: "POST",
    params: filters,
    body,
  });
}

export async function getIncidentBlastRadius(
  incidentID: string,
  filters?: Pick<EventFilters, "environment" | "tenant_id" | "repo">,
) {
  return parseTopologyBlastRadiusResponse(await fetchJSON<unknown>(`/v1/incidents/${encodeURIComponent(incidentID)}/blast-radius`, {
    params: {
      environment: filters?.environment,
      tenant_id: filters?.tenant_id,
      repo: filters?.repo,
    },
  }));
}

export async function getMetricBlastRadius(metricKey: string, filters: EventFilters) {
  return parseTopologyBlastRadiusResponse(await fetchJSON<unknown>(`/v1/scorecard/metrics/${encodeURIComponent(metricKey)}/blast-radius`, {
    params: {
      decision: filters.decision,
      component: filters.component,
      repo: filters.repo,
      environment: filters.environment,
      tenant_id: filters.tenant_id,
      limit: filters.limit,
    },
  }));
}

export async function getTopViolators(filters: {
  window_days?: string;
  limit?: string;
  dimension?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
}) {
  return parseTopViolatorsResponse(await fetchJSON<unknown>("/v1/analytics/top-violators", { params: filters }));
}

export async function getDriftStats(filters: {
  window_days?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
}) {
  return parseDriftStatsResponse(await fetchJSON<unknown>("/v1/analytics/drift-stats", { params: filters }));
}

export async function searchSBOMComponents(filters: {
  component_name?: string;
  purl?: string;
  image_digest?: string;
  tenant_id?: string;
  limit?: string;
}) {
  return parseSBOMComponentsResponse(await fetchJSON<unknown>("/v1/sbom/components/search", { params: filters }));
}

export async function getSBOMImage(imageDigest: string, limit = 100, tenantID?: string) {
  return parseSBOMImageResponse(
    await fetchJSON<unknown>(`/v1/sbom/images/${encodeURIComponent(imageDigest)}`, {
      params: { limit: String(limit), tenant_id: tenantID },
    }),
  );
}

export async function getActiveVulnerabilities(filters: {
  severity?: string;
  cve_id?: string;
  image_digest?: string;
  component_name?: string;
  tenant_id?: string;
  environment?: string;
  include_suppressed?: string;
  limit?: string;
}) {
  return parseVulnerabilitiesResponse(await fetchJSON<unknown>("/v1/vulnerabilities/active", { params: filters }));
}

export async function getNetVulnerabilities(filters: {
  severity?: string;
  severity_threshold?: string;
  cve_id?: string;
  image_digest?: string;
  component_name?: string;
  tenant_id?: string;
  environment?: string;
  limit?: string;
}) {
  return parseVulnerabilityNetResponse(await fetchJSON<unknown>("/v1/vulnerabilities/net", { params: filters }));
}

export async function getVulnerabilityBlastRadius(filters: {
  cve_id?: string;
  component_name?: string;
  purl?: string;
  tenant_id?: string;
  limit?: string;
}) {
  return parseVulnerabilityBlastRadiusResponse(await fetchJSON<unknown>("/v1/vulnerabilities/blast-radius", { params: filters }));
}

export async function getVulnerabilityTimeline(filters: {
  image_digest: string;
  cve_id: string;
  tenant_id?: string;
  window_days?: string;
}) {
  return parseVulnerabilityTimelineResponse(await fetchJSON<unknown>("/v1/vulnerabilities/timeline", { params: filters }));
}

export async function getVulnerabilityDecisions(filters: {
  image_digest?: string;
  cve_id?: string;
  tenant_id?: string;
  active?: string;
  limit?: string;
}) {
  return parseVulnerabilityDecisionsResponse(await fetchJSON<unknown>("/v1/vulnerabilities/decisions", { params: filters }));
}

export async function getVEXStatements(filters: {
  vulnerability_id?: string;
  image_digest?: string;
  package_name?: string;
  tenant_id?: string;
  environment?: string;
  active?: string;
  limit?: string;
}) {
  return parseVEXStatementsResponse(await fetchJSON<unknown>("/v1/vex", { params: filters }));
}

export async function getVEXStatus(filters?: {
  vulnerability_id?: string;
  image_digest?: string;
  tenant_id?: string;
  cluster_id?: string;
}) {
  return parseVEXStatusSummary(await fetchJSON<unknown>("/v1/vex/status", { params: filters }));
}

export async function createVEXStatement(input: VEXCreateInput) {
  return parseVEXStatementActionResponse(
    await fetchJSON<unknown>("/v1/vex", {
      method: "POST",
      body: input,
    }),
  );
}

export async function revokeVEXStatement(statementID: number) {
  return parseVEXStatementActionResponse(
    await fetchJSON<unknown>(`/v1/vex/${encodeURIComponent(String(statementID))}/revoke`, {
      method: "POST",
    }),
  );
}

export async function createVulnerabilityDecision(input: VulnerabilityDecisionInput) {
  return parseVulnerabilityDecisionActionResponse(
    await fetchJSON<unknown>("/v1/vulnerabilities/decisions", {
      method: "POST",
      body: input,
    }),
  );
}

export async function deactivateVulnerabilityDecision(decisionID: number) {
  return parseVulnerabilityDecisionActionResponse(
    await fetchJSON<unknown>(`/v1/vulnerabilities/decisions/${decisionID}/deactivate`, {
      method: "POST",
    }),
  );
}

export async function rescanVulnerabilities(input?: { image_digest?: string; image_ref?: string }) {
  return parseVulnerabilityRescanResponse(
    await fetchJSON<unknown>("/v1/vulnerabilities/rescan", {
      method: "POST",
      body: input || {},
    }),
  );
}

export async function sealHandoff(
  input: {
    audience: string;
    incident_ids?: string[];
    include_forensics?: boolean;
    include_recommendations?: boolean;
    co_sign_mode?: string;
  },
  filters?: Record<string, string | string[] | undefined>,
) {
  return parseHandoffSealResponse(
    await fetchJSON<unknown>("/v1/handoff/seal", {
      method: "POST",
      params: filters,
      body: input,
    }),
  );
}

export async function getHandoff(packageID: string) {
  return parseHandoffSealResponse(await fetchJSON<unknown>(`/v1/handoff/${encodeURIComponent(packageID)}`));
}

export async function getHandoffManifest(packageID: string) {
  return parseSealedManifest(await fetchJSON<unknown>(`/v1/handoff/${encodeURIComponent(packageID)}/manifest`));
}

export async function getHandoffVerification(packageID: string) {
  return parseVerificationResult(await fetchJSON<unknown>(`/v1/handoff/${encodeURIComponent(packageID)}/verification`));
}

export async function cosignHandoff(packageID: string, signerRole: string) {
  return parseHandoffSealResponse(
    await fetchJSON<unknown>(`/v1/handoff/${encodeURIComponent(packageID)}/cosign`, {
      method: "POST",
      body: { signer_role: signerRole },
    }),
  );
}

export async function downloadHandoffBundle(packageID: string) {
  return fetchBlob(`/v1/handoff/${encodeURIComponent(packageID)}/download`);
}

export async function getRuntimeDriftFindings(filters: {
  cluster_id?: string;
  tenant_id?: string;
  namespace?: string;
  workload_kind?: string;
  workload?: string;
  severity?: string;
  status?: string;
  limit?: string;
}) {
  return parseRuntimeDriftFindingsResponse(await fetchJSON<unknown>("/v1/runtime/drift", { params: filters }));
}

export async function getRuntimeDriftStatus(filters: {
  cluster_id?: string;
  tenant_id?: string;
  namespace?: string;
  workload_kind?: string;
  workload?: string;
  severity?: string;
  status?: string;
  limit?: string;
}) {
  return parseRuntimeDriftStatus(await fetchJSON<unknown>("/v1/runtime/drift/status", { params: filters }));
}

export async function getRuntimeActiveStates(filters: {
  cluster_id?: string;
  tenant_id?: string;
  namespace?: string;
  workload_kind?: string;
  workload?: string;
  reconciliation_status?: string;
  quarantine_type?: string;
  limit?: string;
}) {
  return parseRuntimeActiveStatesResponse(await fetchJSON<unknown>("/v1/runtime/active-state", { params: filters }));
}

export async function getRuntimeClosedLoopStatus(filters: {
  cluster_id?: string;
  tenant_id?: string;
  namespace?: string;
  workload_kind?: string;
  workload?: string;
  reconciliation_status?: string;
  quarantine_type?: string;
  limit?: string;
}) {
  return parseRuntimeClosedLoopStatus(await fetchJSON<unknown>("/v1/runtime/closed-loop/status", { params: filters }));
}

export async function getSigningIdentityObservations(filters?: Record<string, string | undefined>) {
  return parseSigningIdentityObservationsResponse(await fetchJSON<unknown>("/v1/signing-identities", { params: filters }));
}

export async function getSigningIdentityPolicies() {
  return parseSigningIdentityPoliciesResponse(await fetchJSON<unknown>("/v1/signing-identities/policies"));
}

export async function getSigningIdentityFindings(filters?: Record<string, string | undefined>) {
  return parseSigningIdentityFindingsResponse(await fetchJSON<unknown>("/v1/signing-identities/findings", { params: filters }));
}

export async function getSigningIdentityStatus(filters?: Record<string, string | undefined>) {
  const payload = await fetchJSON<unknown>("/v1/signing-identities/status", { params: filters });
  if (!isRecord(payload) || !isRecord(payload.status)) {
    throw new Error("Audit API returned invalid signing identity status response.");
  }
  return parseSigningIdentityStatus(payload.status);
}

export async function createAuditReport(input?: {
  tenant_id?: string;
  cluster_id?: string;
  environment?: string;
  repo?: string;
  format?: string;
  include_public_view?: boolean;
}) {
  return parseAuditReport(
    await fetchJSON<unknown>("/v1/audit/reports", {
      method: "POST",
      body: input || {},
    }),
  );
}

export async function getAIGuidance(filters?: Record<string, string | undefined>) {
  return parseGuidanceResponse(await fetchJSON<unknown>("/v1/ai/guidance", { params: filters }));
}

export async function getAIInsights(filters?: Record<string, string | undefined>) {
  return parseAIInsightsResponse(await fetchJSON<unknown>("/v1/ai/insights", { params: filters }));
}

export function apiBaseURL() {
  return API_BASE_URL;
}

export function apiTokenConfigured() {
  return API_TOKEN.length > 0;
}
