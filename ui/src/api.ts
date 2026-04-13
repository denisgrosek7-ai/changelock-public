import type { AuditHealth, EventFilters, EventsResponse, Summary, TabKey } from "./types";

const API_BASE_URL = (import.meta.env.VITE_API_BASE_URL || "/api").replace(/\/$/, "");
const API_TIMEOUT_MS = Number.parseInt(import.meta.env.VITE_API_TIMEOUT_MS || "8000", 10);

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
      headers: { Accept: "application/json" },
      cache: "no-store",
      signal: controller.signal,
    });

    if (!response.ok) {
      const contentType = response.headers.get("content-type") || "";
      if (contentType.includes("application/json")) {
        const payload = (await response.json()) as { error?: string };
        throw new Error(payload.error || `request failed with status ${response.status}`);
      }

      const payload = await response.text();
      throw new Error(payload || `request failed with status ${response.status}`);
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

export async function getHealth() {
  return fetchJSON<AuditHealth>("/health");
}

export async function getSummary(filters: Pick<EventFilters, "environment" | "tenant_id">) {
  return fetchJSON<Summary>("/v1/reports/summary", filters);
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

  return fetchJSON<EventsResponse>(path, params);
}

export function apiBaseURL() {
  return API_BASE_URL;
}
