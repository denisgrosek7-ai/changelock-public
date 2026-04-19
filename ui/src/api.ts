import type {
  DefenseGapAssessment,
  IncidentExport,
  IncidentPackage,
  InvestigationIncident,
  MetricIncidentDrilldown,
  PolicyReplayAssessment,
  SystemicWeaknessResponse,
} from "./incidents";
import type {
  ActiveWorkloadRef,
  AuditHealth,
  AuthStatus,
  DriftStatsResponse,
  EventFilters,
  EventsResponse,
  ExceptionActionResponse,
  ExceptionReport,
  ExceptionRequestInput,
  ExceptionsResponse,
  SBOMComponent,
  SBOMComponentsResponse,
  SBOMDocument,
  SBOMImageResponse,
  PolicyException,
  ReasonCount,
  StoredEvent,
  Summary,
  SyncStatus,
  TabKey,
  TopViolator,
  TopViolatorsResponse,
  TrendBucket,
  TrendsResponse,
  VulnerabilityBlastRadiusItem,
  VulnerabilityBlastRadiusResponse,
  VulnerabilityDecision,
  VulnerabilityDecisionInput,
  VulnerabilityDecisionsResponse,
  VulnerabilityFinding,
  VulnerabilitiesResponse,
  VulnerabilityRescanResponse,
  VulnerabilityTimelineEntry,
  VulnerabilityTimelineResponse,
  VerifierSummary,
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
  if (!isRecord(value) || !Array.isArray(value.incidents) || !isRecord(value.aggregate)) {
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
    limitations: readOptionalStringArray(value.limitations, "limitations") || [],
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
  window_days?: string;
  granularity?: string;
  tenant_id?: string;
  environment?: string;
  repo?: string;
  event_type?: string;
}) {
  return parseTrendsResponse(await fetchJSON<unknown>("/v1/analytics/trends", { params: filters }));
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

export function apiBaseURL() {
  return API_BASE_URL;
}

export function apiTokenConfigured() {
  return API_TOKEN.length > 0;
}
