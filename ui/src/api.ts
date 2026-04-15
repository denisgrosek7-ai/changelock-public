import type {
  AuditHealth,
  AuthStatus,
  EventFilters,
  EventsResponse,
  ReasonCount,
  StoredEvent,
  Summary,
  TabKey,
  VerifierSummary,
} from "./types";

const API_BASE_URL = (import.meta.env.VITE_API_BASE_URL || "/api").replace(/\/$/, "");
const API_TOKEN = (import.meta.env.VITE_API_TOKEN || "").trim();
const API_TIMEOUT_MS = Number.parseInt(import.meta.env.VITE_API_TIMEOUT_MS || "8000", 10);

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

async function fetchJSON<T>(path: string, params?: Record<string, string | undefined>): Promise<T> {
  const controller = new AbortController();
  const timeoutID = window.setTimeout(() => controller.abort(), API_TIMEOUT_MS);

  try {
    const response = await fetch(buildURL(path, params), {
      headers: {
        Accept: "application/json",
        ...(API_TOKEN ? { Authorization: `Bearer ${API_TOKEN}` } : {}),
      },
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
    reason: readString(value.reason, "top_deny_reasons.reason"),
    count: readNumber(value.count, "top_deny_reasons.count"),
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
    exception_reason: readOptionalString(value.exception_reason, "events[].exception_reason"),
    exception_ticket_id: readOptionalString(value.exception_ticket_id, "events[].exception_ticket_id"),
    exception_approved_by: readOptionalString(value.exception_approved_by, "events[].exception_approved_by"),
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

export async function getHealth() {
  return fetchJSON<AuditHealth>("/health");
}

export async function getAuthStatus() {
  return fetchJSON<AuthStatus>("/v1/auth/me");
}

export async function getSummary(filters: Pick<EventFilters, "environment" | "tenant_id">) {
  return parseSummary(await fetchJSON<unknown>("/v1/reports/summary", filters));
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

  return parseEventsResponse(await fetchJSON<unknown>(path, params));
}

export function apiBaseURL() {
  return API_BASE_URL;
}

export function apiTokenConfigured() {
  return API_TOKEN.length > 0;
}
