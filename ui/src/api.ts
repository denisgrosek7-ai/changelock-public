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

function buildURL(path: string, params?: Record<string, string | undefined>) {
  const url = new URL(`${API_BASE_URL}${path}`, window.location.origin);
  if (params) {
    for (const [key, value] of Object.entries(params)) {
      if (value) {
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
    params?: Record<string, string | undefined>;
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

export async function getHealth() {
  return fetchJSON<AuditHealth>("/health");
}

export async function getAuthStatus() {
  return parseAuthStatus(await fetchJSON<unknown>("/v1/auth/me"));
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
