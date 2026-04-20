import { useEffect, useState } from "react";

import {
  buildIncidents,
  type DefenseGapAssessment,
  type ExecutiveDefenseReport,
  type IncidentExport,
  type IncidentImpactTone,
  type IncidentPackage,
  type PolicyReplayAssessment,
  type IncidentReportAudience,
  type InvestigationIncident,
  type MetricIncidentDrilldown,
} from "../incidents";
import type { StoredEvent } from "../types";
import { EventDetails } from "./EventDetails";
import { EventsTable } from "./EventsTable";

type Props = {
  incidents?: InvestigationIncident[];
  events: StoredEvent[];
  loading: boolean;
  error: string | null;
  role?: string;
  metricDrilldown?: MetricIncidentDrilldown | null;
  onClearMetricDrilldown?: () => void;
  onLoadExport?: (incidentID: string, audience: IncidentReportAudience) => Promise<IncidentExport>;
  onLoadPackage?: (incidentIDs: string[], audience: IncidentReportAudience) => Promise<IncidentPackage>;
  onLoadExecutiveReport?: (incidentIDs: string[], audience: IncidentReportAudience) => Promise<ExecutiveDefenseReport>;
  onLoadIncidentDefenseGaps?: (incidentID: string) => Promise<DefenseGapAssessment>;
  onLoadMetricDefenseGaps?: (metricKey: string) => Promise<DefenseGapAssessment>;
  onLoadIncidentPolicyReplay?: (incidentID: string) => Promise<PolicyReplayAssessment>;
  onLoadMetricPolicyReplay?: (metricKey: string) => Promise<PolicyReplayAssessment>;
  onAcknowledge?: (incidentID: string, summary?: string) => Promise<void>;
  onWatch?: (incidentID: string, summary?: string) => Promise<void>;
  onAssign?: (incidentID: string, owner: string, reason: string) => Promise<void>;
  onResolve?: (
    incidentID: string,
    input: {
      resolution_type: string;
      resolution_summary: string;
      resolution_details?: string;
      resolution_refs?: string[];
      follow_up_required?: boolean;
    },
  ) => Promise<void>;
  onReopen?: (incidentID: string, reason?: string) => Promise<void>;
  onAddNote?: (incidentID: string, note: string) => Promise<void>;
};

type ReportMode = "report" | "json";

function formatTimestamp(timestamp?: string) {
  if (!timestamp) {
    return "-";
  }
  return new Date(timestamp).toLocaleString();
}

function severityClass(value: InvestigationIncident["severity"]) {
  return value === "critical" || value === "high" ? "deny" : value === "medium" ? "warning" : "muted";
}

function statusClass(value: InvestigationIncident["status"]) {
  return value === "active" ? "deny" : value === "watch" ? "warning" : "allow";
}

function lifecycleClass(value: InvestigationIncident["state"]) {
  return value === "resolved" ? "allow" : value === "watching" ? "warning" : value === "reopened" ? "deny" : "muted";
}

function priorityClass(value: InvestigationIncident["priority"]) {
  return value === "critical" ? "deny" : value === "high" ? "warning" : value === "medium" ? "muted" : "allow";
}

function outcomeClass(value: "deny" | "allow" | "error" | "signal") {
  return value === "deny" ? "deny" : value === "allow" ? "allow" : value === "error" ? "critical" : "muted";
}

function impactClass(value: IncidentImpactTone) {
  return value === "critical" ? "critical" : value === "warning" ? "warning" : value === "allow" ? "allow" : "muted";
}

function confidenceClass(value: "high" | "medium" | "limited") {
  return value === "high" ? "allow" : value === "medium" ? "warning" : "muted";
}

function trendClass(value: "improving" | "watch" | "worsening") {
  return value === "improving" ? "allow" : value === "worsening" ? "deny" : "warning";
}

function shieldBandClass(value: "strong" | "watch" | "at_risk") {
  return value === "strong" ? "allow" : value === "watch" ? "warning" : "deny";
}

function humanizeKey(value: string) {
  return value.replace(/_/g, " ");
}

function describeScope(incident: InvestigationIncident) {
  const repos = incident.affectedRepos.length;
  const environments = incident.affectedEnvironments.length;
  const workloads = incident.affectedWorkloads.length;
  const formatCount = (count: number, singular: string, plural = `${singular}s`) => `${count} ${count === 1 ? singular : plural}`;
  return `${formatCount(repos, "repo")} · ${formatCount(environments, "env")} · ${formatCount(workloads, "workload")}`;
}

function firstOrDash(values: string[], fallback = "-") {
  return values.length > 0 ? values[0] : fallback;
}

function renderValueList(values: string[], emptyMessage: string, limit = values.length) {
  if (values.length === 0) {
    return <div className="summary-list-empty">{emptyMessage}</div>;
  }

  return (
    <ul className="summary-list summary-list--compact">
      {values.slice(0, limit).map((value) => (
        <li key={value}>
          <span>{value}</span>
        </li>
      ))}
    </ul>
  );
}

function renderChipList(values: string[], emptyMessage: string, keyPrefix: string, limit = values.length) {
  if (values.length === 0) {
    return <div className="summary-list-empty">{emptyMessage}</div>;
  }

  return (
    <div className="chip-row">
      {values.slice(0, limit).map((value) => (
        <span className="chip chip--muted" key={`${keyPrefix}-${value}`}>{value}</span>
      ))}
    </div>
  );
}

function renderDefenseGapActions(title: string, values: string[]) {
  return (
    <div>
      <span className="summary-label">{title}</span>
      {renderValueList(values, `No ${title.toLowerCase()} guidance attached.`, 4)}
    </div>
  );
}

export function IncidentWorkbench({
  incidents: serverIncidents = [],
  events,
  loading,
  error,
  role,
  metricDrilldown,
  onClearMetricDrilldown,
  onLoadExport,
  onLoadPackage,
  onLoadExecutiveReport,
  onLoadIncidentDefenseGaps,
  onLoadMetricDefenseGaps,
  onLoadIncidentPolicyReplay,
  onLoadMetricPolicyReplay,
  onAcknowledge,
  onWatch,
  onAssign,
  onResolve,
  onReopen,
  onAddNote,
}: Props) {
  const incidents = serverIncidents.length > 0 ? serverIncidents : buildIncidents(events);
  const [selectedIncidentID, setSelectedIncidentID] = useState<string | null>(null);
  const selectedIncident = incidents.find((incident) => incident.id === selectedIncidentID) || incidents[0] || null;
  const [selectedEvent, setSelectedEvent] = useState<StoredEvent | null>(null);
  const [assignmentOwner, setAssignmentOwner] = useState("");
  const [assignmentReason, setAssignmentReason] = useState("");
  const [lifecycleSummary, setLifecycleSummary] = useState("");
  const [noteDraft, setNoteDraft] = useState("");
  const [resolutionType, setResolutionType] = useState("fixed");
  const [resolutionSummary, setResolutionSummary] = useState("");
  const [resolutionDetails, setResolutionDetails] = useState("");
  const [resolutionRefs, setResolutionRefs] = useState("");
  const [followUpRequired, setFollowUpRequired] = useState(false);
  const [actionSubmitting, setActionSubmitting] = useState(false);
  const [actionError, setActionError] = useState<string | null>(null);
  const [exportPayloads, setExportPayloads] = useState<Record<string, IncidentExport>>({});
  const [exportLoading, setExportLoading] = useState(false);
  const [exportError, setExportError] = useState<string | null>(null);
  const [reportMode, setReportMode] = useState<ReportMode>("report");
  const [reportAudience, setReportAudience] = useState<IncidentReportAudience>("internal");
  const [handoffMode, setHandoffMode] = useState(false);
  const [selectedPackageIDs, setSelectedPackageIDs] = useState<string[]>([]);
  const [packageAudience, setPackageAudience] = useState<IncidentReportAudience>("internal");
  const [packagePayload, setPackagePayload] = useState<IncidentPackage | null>(null);
  const [packageLoading, setPackageLoading] = useState(false);
  const [packageError, setPackageError] = useState<string | null>(null);
  const [packageHandoffMode, setPackageHandoffMode] = useState(false);
  const [executivePayload, setExecutivePayload] = useState<ExecutiveDefenseReport | null>(null);
  const [executiveLoading, setExecutiveLoading] = useState(false);
  const [executiveError, setExecutiveError] = useState<string | null>(null);
  const [incidentDefenseAssessments, setIncidentDefenseAssessments] = useState<Record<string, DefenseGapAssessment>>({});
  const [metricDefenseAssessments, setMetricDefenseAssessments] = useState<Record<string, DefenseGapAssessment>>({});
  const [incidentReplayAssessments, setIncidentReplayAssessments] = useState<Record<string, PolicyReplayAssessment>>({});
  const [metricReplayAssessments, setMetricReplayAssessments] = useState<Record<string, PolicyReplayAssessment>>({});
  const [defenseGapLoading, setDefenseGapLoading] = useState(false);
  const [metricDefenseGapLoading, setMetricDefenseGapLoading] = useState(false);
  const [incidentReplayLoading, setIncidentReplayLoading] = useState(false);
  const [metricReplayLoading, setMetricReplayLoading] = useState(false);
  const [defenseGapError, setDefenseGapError] = useState<string | null>(null);
  const [metricDefenseGapError, setMetricDefenseGapError] = useState<string | null>(null);
  const [incidentReplayError, setIncidentReplayError] = useState<string | null>(null);
  const [metricReplayError, setMetricReplayError] = useState<string | null>(null);
  const exportPayload = selectedIncident ? exportPayloads[`${selectedIncident.id}:${reportAudience}`] : undefined;
  const incidentDefenseGaps = selectedIncident ? incidentDefenseAssessments[selectedIncident.id] : undefined;
  const metricDefenseGaps = metricDrilldown ? metricDefenseAssessments[metricDrilldown.metricKey] : undefined;
  const incidentReplay = selectedIncident ? incidentReplayAssessments[selectedIncident.id] : undefined;
  const metricReplay = metricDrilldown ? metricReplayAssessments[metricDrilldown.metricKey] : undefined;
  const canManage = role === "operator" || role === "security_admin";
  const canResolve = role === "security_admin";

  useEffect(() => {
    if (!selectedIncidentID && incidents[0]) {
      setSelectedIncidentID(incidents[0].id);
      return;
    }
    if (selectedIncidentID && !incidents.some((incident) => incident.id === selectedIncidentID)) {
      setSelectedIncidentID(incidents[0]?.id || null);
    }
  }, [incidents, selectedIncidentID]);

  useEffect(() => {
    setSelectedPackageIDs((current) => {
      const available = new Set(incidents.map((incident) => incident.id));
      const next = current.filter((incidentID) => available.has(incidentID));
      if (next.length > 0) {
        return next;
      }
      return incidents[0] ? [incidents[0].id] : [];
    });
  }, [incidents]);

  useEffect(() => {
    if (!selectedIncident) {
      setSelectedEvent(null);
      return;
    }
    setSelectedEvent((current) => selectedIncident.events.find((event) => event.id === current?.id) || selectedIncident.events[0] || null);
  }, [selectedIncident]);

  useEffect(() => {
    if (!selectedIncident) {
      return;
    }
    setAssignmentOwner(selectedIncident.owner || selectedIncident.assignment.owner || "");
    setAssignmentReason(selectedIncident.assignment.reason || "");
    setLifecycleSummary("");
    setNoteDraft("");
    setResolutionType(selectedIncident.resolution.type || "fixed");
    setResolutionSummary(selectedIncident.resolution.summary || "");
    setResolutionDetails(selectedIncident.resolution.details || "");
    setResolutionRefs(selectedIncident.resolution.refs.join(", "));
    setFollowUpRequired(Boolean(selectedIncident.resolution.followUpRequired));
    setActionError(null);
    setExportError(null);
    setReportMode("report");
    setReportAudience("internal");
    setHandoffMode(false);
  }, [selectedIncident?.id]);

  useEffect(() => {
    const active = Boolean(exportPayload) && reportMode === "report" && handoffMode;
    document.body.classList.toggle("incident-handoff-print-active", active);
    return () => {
      document.body.classList.remove("incident-handoff-print-active");
    };
  }, [exportPayload, reportMode, handoffMode]);

  useEffect(() => {
    document.body.classList.toggle("incident-package-print-active", Boolean(packagePayload) && packageHandoffMode);
    return () => {
      document.body.classList.remove("incident-package-print-active");
    };
  }, [packagePayload, packageHandoffMode]);

  useEffect(() => {
    if (!selectedIncident || !onLoadIncidentDefenseGaps || incidentDefenseAssessments[selectedIncident.id]) {
      return;
    }
    let ignore = false;
    setDefenseGapLoading(true);
    setDefenseGapError(null);
    void onLoadIncidentDefenseGaps(selectedIncident.id)
      .then((payload) => {
        if (ignore) {
          return;
        }
        setIncidentDefenseAssessments((current) => ({ ...current, [selectedIncident.id]: payload }));
      })
      .catch((loadError) => {
        if (ignore) {
          return;
        }
        setDefenseGapError(loadError instanceof Error ? loadError.message : "Unable to load defense-gap assessment.");
      })
      .finally(() => {
        if (!ignore) {
          setDefenseGapLoading(false);
        }
      });
    return () => {
      ignore = true;
    };
  }, [incidentDefenseAssessments, onLoadIncidentDefenseGaps, selectedIncident]);

  useEffect(() => {
    if (!metricDrilldown?.metricKey || !onLoadMetricDefenseGaps || metricDefenseAssessments[metricDrilldown.metricKey]) {
      return;
    }
    let ignore = false;
    setMetricDefenseGapLoading(true);
    setMetricDefenseGapError(null);
    void onLoadMetricDefenseGaps(metricDrilldown.metricKey)
      .then((payload) => {
        if (ignore) {
          return;
        }
        setMetricDefenseAssessments((current) => ({ ...current, [metricDrilldown.metricKey]: payload }));
      })
      .catch((loadError) => {
        if (ignore) {
          return;
        }
        setMetricDefenseGapError(loadError instanceof Error ? loadError.message : "Unable to load metric defense-gap assessment.");
      })
      .finally(() => {
        if (!ignore) {
          setMetricDefenseGapLoading(false);
        }
      });
    return () => {
      ignore = true;
    };
  }, [metricDefenseAssessments, metricDrilldown, onLoadMetricDefenseGaps]);

  useEffect(() => {
    if (!selectedIncident || !onLoadIncidentPolicyReplay || incidentReplayAssessments[selectedIncident.id]) {
      return;
    }
    let ignore = false;
    setIncidentReplayLoading(true);
    setIncidentReplayError(null);
    void onLoadIncidentPolicyReplay(selectedIncident.id)
      .then((payload) => {
        if (ignore) {
          return;
        }
        setIncidentReplayAssessments((current) => ({ ...current, [selectedIncident.id]: payload }));
      })
      .catch((loadError) => {
        if (ignore) {
          return;
        }
        setIncidentReplayError(loadError instanceof Error ? loadError.message : "Unable to load policy replay assessment.");
      })
      .finally(() => {
        if (!ignore) {
          setIncidentReplayLoading(false);
        }
      });
    return () => {
      ignore = true;
    };
  }, [incidentReplayAssessments, onLoadIncidentPolicyReplay, selectedIncident]);

  useEffect(() => {
    if (!metricDrilldown?.metricKey || !onLoadMetricPolicyReplay || metricReplayAssessments[metricDrilldown.metricKey]) {
      return;
    }
    let ignore = false;
    setMetricReplayLoading(true);
    setMetricReplayError(null);
    void onLoadMetricPolicyReplay(metricDrilldown.metricKey)
      .then((payload) => {
        if (ignore) {
          return;
        }
        setMetricReplayAssessments((current) => ({ ...current, [metricDrilldown.metricKey]: payload }));
      })
      .catch((loadError) => {
        if (ignore) {
          return;
        }
        setMetricReplayError(loadError instanceof Error ? loadError.message : "Unable to load metric policy replay assessment.");
      })
      .finally(() => {
        if (!ignore) {
          setMetricReplayLoading(false);
        }
      });
    return () => {
      ignore = true;
    };
  }, [metricDrilldown, metricReplayAssessments, onLoadMetricPolicyReplay]);

  async function runAction(action: () => Promise<void>) {
    setActionSubmitting(true);
    setActionError(null);
    try {
      await action();
      setLifecycleSummary("");
      setNoteDraft("");
    } catch (submitError) {
      setActionError(submitError instanceof Error ? submitError.message : "Incident action failed.");
    } finally {
      setActionSubmitting(false);
    }
  }

  async function loadExport(incidentID: string, audience: IncidentReportAudience) {
    if (!onLoadExport) {
      return;
    }
    setExportLoading(true);
    setExportError(null);
    try {
      const payload = await onLoadExport(incidentID, audience);
      setExportPayloads((current) => ({ ...current, [`${incidentID}:${audience}`]: payload }));
    } catch (loadError) {
      setExportError(loadError instanceof Error ? loadError.message : "Unable to load incident export.");
    } finally {
      setExportLoading(false);
    }
  }

  async function loadPackage(incidentIDs: string[], audience: IncidentReportAudience) {
    if (!onLoadPackage) {
      return;
    }
    setPackageLoading(true);
    setPackageError(null);
    try {
      const payload = await onLoadPackage(incidentIDs, audience);
      setPackagePayload(payload);
    } catch (loadError) {
      setPackageError(loadError instanceof Error ? loadError.message : "Unable to load incident package.");
    } finally {
      setPackageLoading(false);
    }
  }

  async function loadExecutiveReport(incidentIDs: string[], audience: IncidentReportAudience) {
    if (!onLoadExecutiveReport) {
      return;
    }
    setExecutiveLoading(true);
    setExecutiveError(null);
    try {
      const payload = await onLoadExecutiveReport(incidentIDs, audience);
      setExecutivePayload(payload);
    } catch (loadError) {
      setExecutiveError(loadError instanceof Error ? loadError.message : "Unable to load executive defense report.");
    } finally {
      setExecutiveLoading(false);
    }
  }

  function downloadExport(payload: IncidentExport) {
    const blob = new Blob([JSON.stringify(payload, null, 2)], { type: "application/json" });
    const url = URL.createObjectURL(blob);
    const anchor = document.createElement("a");
    anchor.href = url;
    anchor.download = `${payload.incidentID.toLowerCase()}-${payload.audience}-case-export.json`;
    document.body.appendChild(anchor);
    anchor.click();
    anchor.remove();
    URL.revokeObjectURL(url);
  }

  function downloadPackage(payload: IncidentPackage) {
    const blob = new Blob([JSON.stringify(payload, null, 2)], { type: "application/json" });
    const url = URL.createObjectURL(blob);
    const anchor = document.createElement("a");
    anchor.href = url;
    anchor.download = `incident-package-${payload.audience}-${payload.selectionMode}.json`;
    document.body.appendChild(anchor);
    anchor.click();
    anchor.remove();
    URL.revokeObjectURL(url);
  }

  function downloadExecutiveReport(payload: ExecutiveDefenseReport) {
    const blob = new Blob([JSON.stringify(payload, null, 2)], { type: "application/json" });
    const url = URL.createObjectURL(blob);
    const anchor = document.createElement("a");
    anchor.href = url;
    anchor.download = `executive-defense-report-${payload.audience}-${payload.selectionMode}.json`;
    document.body.appendChild(anchor);
    anchor.click();
    anchor.remove();
    URL.revokeObjectURL(url);
  }

  function printHandoff() {
    window.print();
  }

  function togglePackageSelection(incidentID: string) {
    setSelectedPackageIDs((current) => (
      current.includes(incidentID)
        ? current.filter((value) => value !== incidentID)
        : [...current, incidentID]
    ));
  }

  if (loading) {
    return <section className="panel panel-empty">Grouping evidence-backed events into incidents…</section>;
  }

  if (error) {
    return <section className="panel panel-empty panel-error">Unable to build investigations. {error}</section>;
  }

  if (incidents.length === 0) {
    return <section className="panel panel-empty">No incident clusters matched the current investigation scope.</section>;
  }

  return (
    <section className="incident-grid">
      <section className="panel incident-list-panel">
        <div className="table-toolbar">
          <div>
            <span className="summary-label">Active investigations</span>
            <strong>{incidents.length}</strong>
          </div>
          <div className="chip-row">
            <span className="chip chip--muted">package selection {selectedPackageIDs.length}</span>
            <button
              type="button"
              className="button button-secondary"
              onClick={() => setSelectedPackageIDs(incidents.map((incident) => incident.id))}
            >
              Use all visible
            </button>
            <button
              type="button"
              className="button button-secondary"
              onClick={() => setSelectedPackageIDs(selectedIncident ? [selectedIncident.id] : [])}
            >
              Use current case
            </button>
          </div>
        </div>
        <div className="incident-list">
          {incidents.map((incident) => (
            <button
              key={incident.id}
              className={`incident-card ${selectedIncident?.id === incident.id ? "is-selected" : ""}`}
              onClick={() => setSelectedIncidentID(incident.id)}
            >
              <div className="incident-card__header">
                <span className={`chip chip--${severityClass(incident.severity)}`}>{incident.severity}</span>
                <span className={`chip chip--${priorityClass(incident.priority)}`}>priority {incident.priority}</span>
                <span className={`chip chip--${statusClass(incident.status)}`}>{incident.status}</span>
                {selectedPackageIDs.includes(incident.id) ? <span className="chip chip--allow">in package</span> : null}
              </div>
              <strong>{incident.title}</strong>
              <p>{incident.summary}</p>
              <div className="incident-card__meta">
                <span>{incident.eventCount} events</span>
                <span>{describeScope(incident)}</span>
              </div>
              <div className="chip-row">
                <span className="chip chip--muted">{incident.id}</span>
                <span className="chip chip--muted">{incident.category}</span>
              </div>
            </button>
          ))}
        </div>
      </section>

      <div className="incident-details-stack">
        <section className="panel incident-package-panel">
          <div className="details-header">
            <div>
              <span className="summary-label">Report package index</span>
              <h2>Multi-incident bundle</h2>
              <p>Build a derived package index over the current incident cases with one audience mode, aggregate summary, and package-level handoff output.</p>
            </div>
            <div className="chip-row">
              <span className="chip chip--muted">{selectedPackageIDs.length} selected</span>
              <span className="chip chip--muted">{incidents.length} visible</span>
            </div>
          </div>

          <div className="incident-evidence-grid">
            <div>
              <span className="summary-label">Package audience</span>
              <div className="chip-row">
                {(["internal", "auditor_safe", "customer_safe"] as IncidentReportAudience[]).map((audience) => (
                  <button
                    type="button"
                    key={`package-audience-${audience}`}
                    className={`button ${packageAudience === audience ? "" : "button-secondary"}`}
                    onClick={() => setPackageAudience(audience)}
                  >
                    {audience === "internal" ? "Internal" : audience === "auditor_safe" ? "Auditor-Safe" : "Customer-Safe"}
                  </button>
                ))}
              </div>
              <p className="incident-inline-copy">
                Current package selection stays incident-centric and reuses the same export/redaction rules as individual case reports.
              </p>
            </div>
            <div>
              <span className="summary-label">Selection controls</span>
              <div className="chip-row">
                <button
                  type="button"
                  className="button"
                  disabled={packageLoading || !onLoadPackage}
                  onClick={() => void loadPackage(selectedPackageIDs, packageAudience)}
                >
                  {packageLoading ? "Loading package…" : "Load selected incidents"}
                </button>
                <button
                  type="button"
                  className="button button-secondary"
                  disabled={packageLoading || !onLoadPackage}
                  onClick={() => void loadPackage([], packageAudience)}
                >
                  Load current filtered scope
                </button>
                <button
                  type="button"
                  className="button button-secondary"
                  onClick={() => setSelectedPackageIDs([])}
                >
                  Clear selection
                </button>
              </div>
              <div className="chip-row">
                <button
                  type="button"
                  className="button button-secondary"
                  disabled={!selectedIncident}
                  onClick={() => selectedIncident && togglePackageSelection(selectedIncident.id)}
                >
                  {selectedIncident && selectedPackageIDs.includes(selectedIncident.id) ? "Remove current case" : "Add current case"}
                </button>
                {packagePayload ? (
                  <button
                    type="button"
                    className="button button-secondary"
                    onClick={() => downloadPackage(packagePayload)}
                  >
                    Download package JSON
                  </button>
                ) : null}
                <button
                  type="button"
                  className="button button-secondary"
                  disabled={executiveLoading || !onLoadExecutiveReport}
                  onClick={() => void loadExecutiveReport(selectedPackageIDs, packageAudience)}
                >
                  {executiveLoading ? "Loading executive brief…" : "Load executive brief"}
                </button>
                <button
                  type="button"
                  className="button button-secondary"
                  disabled={executiveLoading || !onLoadExecutiveReport}
                  onClick={() => void loadExecutiveReport([], packageAudience)}
                >
                  Load scope brief
                </button>
                {executivePayload ? (
                  <button
                    type="button"
                    className="button button-secondary"
                    onClick={() => downloadExecutiveReport(executivePayload)}
                  >
                    Download executive JSON
                  </button>
                ) : null}
              </div>
              {packageError ? <p className="details-copy details-copy--error">{packageError}</p> : null}
              {executiveError ? <p className="details-copy details-copy--error">{executiveError}</p> : null}
            </div>
          </div>

          {selectedPackageIDs.length > 0 ? (
            <div className="chip-row">
              {selectedPackageIDs.map((incidentID) => (
                <button
                  type="button"
                  key={`selected-package-${incidentID}`}
                  className="chip chip--muted"
                  onClick={() => setSelectedIncidentID(incidentID)}
                >
                  {incidentID}
                </button>
              ))}
            </div>
          ) : (
            <div className="summary-list-empty">No explicit incident IDs selected. Loading the package will use the current filtered scope.</div>
          )}

          {executivePayload ? (
            <section className="incident-case-section incident-case-section--wide">
              <div className="incident-report-section__header">
                <span className="summary-label">Executive defense reporting</span>
                <strong>{executivePayload.boardPackage.headline}</strong>
              </div>
              <div className="chip-row">
                <span className={`chip chip--${shieldBandClass(executivePayload.shieldHealth.band)}`}>{executivePayload.shieldHealth.band}</span>
                <span className={`chip chip--${executivePayload.redacted ? "warning" : "allow"}`}>{executivePayload.redacted ? "redacted" : "internal"}</span>
                <span className="chip chip--muted">{executivePayload.audience}</span>
                <span className="chip chip--muted">{executivePayload.incidentCount} incidents</span>
              </div>
              <p>{executivePayload.boardPackage.narrative}</p>
              <p className="details-copy">{executivePayload.executiveSummary.whatMattersNow}</p>

              <div className="incident-evidence-grid">
                <div>
                  <span className="summary-label">Top risks</span>
                  {renderValueList(executivePayload.executiveSummary.topRisks, "No executive risk summary loaded.", 3)}
                </div>
                <div>
                  <span className="summary-label">Top improvements</span>
                  {renderValueList(executivePayload.executiveSummary.topImprovements, "No executive improvement priorities loaded.", 3)}
                </div>
              </div>

              <div className="summary-grid">
                <article className="summary-card">
                  <span className="summary-label">Shield health</span>
                  <strong className="summary-value">{executivePayload.shieldHealth.score}</strong>
                  <p>{executivePayload.shieldHealth.summary}</p>
                </article>
                {executivePayload.shieldHealth.components.map((component) => (
                  <article className="summary-card summary-card--compact" key={`shield-${component.key}`}>
                    <span className="summary-label">{component.label}</span>
                    <strong className="summary-value">{component.score}</strong>
                    <p>{component.summary}</p>
                  </article>
                ))}
              </div>

              <div className="incident-impact-list">
                {executivePayload.riskReductionTrends.map((trend) => (
                  <article className="incident-impact-card incident-defense-gap" key={`trend-${trend.key}`}>
                    <div className="incident-impact-card__header">
                      <strong>{trend.label}</strong>
                      <span className={`chip chip--${trendClass(trend.direction)}`}>{trend.direction}</span>
                    </div>
                    <p>{trend.value}</p>
                    <small>{trend.summary}</small>
                  </article>
                ))}
              </div>

              {executivePayload.strategicGaps.length > 0 ? (
                <div className="incident-impact-list">
                  {executivePayload.strategicGaps.map((gap) => (
                    <article className="incident-impact-card incident-defense-gap" key={`strategic-gap-${gap.id}`}>
                      <div className="incident-impact-card__header">
                        <strong>{gap.title}</strong>
                        <span className={`chip chip--${confidenceClass(gap.confidence)}`}>{gap.confidence}</span>
                      </div>
                      <p>{gap.summary}</p>
                      <small>{gap.investmentTarget}</small>
                    </article>
                  ))}
                </div>
              ) : null}

              <div className="incident-evidence-grid">
                <div>
                  <span className="summary-label">Investment priorities</span>
                  {renderValueList(executivePayload.boardPackage.investmentPriorities, "No investment priorities attached.", 4)}
                </div>
                <div>
                  <span className="summary-label">Next-quarter priorities</span>
                  {renderValueList(executivePayload.boardPackage.nextQuarterPriorities, "No next-quarter priorities attached.", 4)}
                </div>
              </div>

              <div className="incident-evidence-grid">
                <div>
                  <span className="summary-label">Business impact framing</span>
                  <p>{executivePayload.businessImpact.summary}</p>
                  {renderValueList(executivePayload.businessImpact.estimates.map((estimate) => `${estimate.label}: ${estimate.value}`), "No business-impact framing attached.", 4)}
                </div>
                <div>
                  <span className="summary-label">Limitations</span>
                  {renderValueList([...executivePayload.redactionSummary, ...executivePayload.limitations], "No executive limitations attached.", 8)}
                </div>
              </div>
            </section>
          ) : null}

          {packagePayload ? (
            <div className={`incident-package-preview ${packageHandoffMode ? "incident-package-preview--handoff" : ""}`}>
              <div className="chip-row">
                <span className={`chip chip--${packagePayload.redacted ? "warning" : "allow"}`}>
                  {packagePayload.redacted ? "Redacted package" : "Internal package"}
                </span>
                <span className="chip chip--muted">{packagePayload.audience}</span>
                <span className="chip chip--muted">{packagePayload.selectionMode === "explicit" ? "explicit selection" : "query-derived"}</span>
                <span className="chip chip--muted">{packagePayload.incidentCount} incidents</span>
              </div>

              <div className="chip-row">
                <button
                  type="button"
                  className={`button ${packageHandoffMode ? "" : "button-secondary"}`}
                  onClick={() => {
                    setHandoffMode(false);
                    setPackageHandoffMode((current) => !current);
                  }}
                >
                  {packageHandoffMode ? "Screen layout" : "Print-friendly"}
                </button>
                {packageHandoffMode ? (
                  <button type="button" className="button button-secondary" onClick={printHandoff}>
                    Print / Save PDF
                  </button>
                ) : null}
              </div>

              <div className={`incident-package-surface ${packageHandoffMode ? "incident-package-surface--handoff" : ""}`}>
                <section className="incident-report-section incident-report-section--hero">
                  <div className="incident-report-header">
                    <div>
                      <span className="summary-label">{packageHandoffMode ? "Package handoff" : "Package index"}</span>
                      <h3>{packagePayload.packageSummary}</h3>
                      <p>{packagePayload.selectionSummary}</p>
                    </div>
                    <div className="chip-row">
                      <span className={`chip chip--${packagePayload.redacted ? "warning" : "allow"}`}>
                        {packagePayload.redacted ? "Redacted" : "Internal canonical"}
                      </span>
                      <span className="chip chip--muted">{packagePayload.audience}</span>
                      {packageHandoffMode ? <span className="chip chip--muted">print-friendly</span> : null}
                    </div>
                  </div>
                  <div className="incident-report-header__meta">
                    <span>generated {formatTimestamp(packagePayload.generatedAt)}</span>
                    <span>{packagePayload.incidentCount} incidents</span>
                    <span>{packagePayload.selectionMode === "explicit" ? "explicit incident bundle" : "query-derived package"}</span>
                  </div>
                </section>

                <section className="incident-report-section">
                  <div className="incident-report-section__header">
                    <span className="summary-label">Aggregate summary</span>
                    <strong>Package-wide posture</strong>
                  </div>
                  <div className="summary-grid">
                    <article className="summary-card">
                      <span className="summary-label">Lifecycle mix</span>
                      <strong className="summary-value">{packagePayload.aggregate.byState.open || 0}</strong>
                      <p className="details-copy">
                        open {packagePayload.aggregate.byState.open || 0} · acknowledged {packagePayload.aggregate.byState.acknowledged || 0} · watching {packagePayload.aggregate.byState.watching || 0}
                      </p>
                    </article>
                    <article className="summary-card">
                      <span className="summary-label">Resolved / reopened</span>
                      <strong className="summary-value">{(packagePayload.aggregate.byState.resolved || 0) + (packagePayload.aggregate.byState.reopened || 0)}</strong>
                      <p className="details-copy">
                        resolved {packagePayload.aggregate.byState.resolved || 0} · reopened {packagePayload.aggregate.byState.reopened || 0}
                      </p>
                    </article>
                    <article className="summary-card">
                      <span className="summary-label">Severity mix</span>
                      <strong className="summary-value">{packagePayload.aggregate.bySeverity.high || 0}</strong>
                      <p className="details-copy">
                        critical {packagePayload.aggregate.bySeverity.critical || 0} · high {packagePayload.aggregate.bySeverity.high || 0} · medium {packagePayload.aggregate.bySeverity.medium || 0}
                      </p>
                    </article>
                    <article className="summary-card">
                      <span className="summary-label">Categories</span>
                      <strong className="summary-value">{Object.keys(packagePayload.aggregate.byCategory).length}</strong>
                      <p className="details-copy">{Object.entries(packagePayload.aggregate.byCategory).slice(0, 3).map(([key, count]) => `${key} ${count}`).join(" · ") || "No category breakdown recorded."}</p>
                    </article>
                  </div>
                </section>

                <section className="incident-report-section">
                  <div className="incident-report-section__header">
                    <span className="summary-label">Package intelligence</span>
                    <strong>Aggregated advisory package layer</strong>
                  </div>
                  <div className="chip-row">
                    <span className="chip chip--muted">advisory only</span>
                    <span className="chip chip--muted">generated {formatTimestamp(packagePayload.packageIntelligence.generatedAt)}</span>
                  </div>
                  <div className="incident-impact-list">
                    <article className="incident-impact-card incident-defense-gap">
                      <div className="incident-impact-card__header">
                        <strong>Top defense gaps</strong>
                      </div>
                      <p>{packagePayload.packageIntelligence.defenseGapSummary.rationale}</p>
                      {renderChipList(
                        packagePayload.packageIntelligence.defenseGapSummary.topGapTypes.map(humanizeKey),
                        "No dominant defense-gap types recorded for this package.",
                        "package-defense-gap-types",
                        4,
                      )}
                      <div className="chip-row">
                        {Object.entries(packagePayload.packageIntelligence.defenseGapSummary.confidenceMix).map(([key, value]) => (
                          <span className={`chip chip--${confidenceClass(key as "high" | "medium" | "limited")}`} key={`package-gap-confidence-${key}`}>
                            {key} {value}
                          </span>
                        ))}
                      </div>
                      {packagePayload.packageIntelligence.defenseGapSummary.topFindings.length > 0 ? (
                        <div className="summary-grid">
                          {packagePayload.packageIntelligence.defenseGapSummary.topFindings.slice(0, 2).map((gap) => (
                            <article className="summary-card summary-card--compact" key={`package-gap-${gap.gapType}`}>
                              <div className="overview-list-item__title">
                                <strong>{gap.title}</strong>
                                <span className={`chip chip--${confidenceClass(gap.confidence)}`}>{gap.confidence}</span>
                              </div>
                              <p>{gap.whyItMatters}</p>
                              {renderChipList(gap.relatedIncidentRefs, "No related incident refs attached.", `package-gap-related-${gap.gapType}`, 4)}
                            </article>
                          ))}
                        </div>
                      ) : null}
                    </article>

                    <article className="incident-impact-card incident-defense-gap">
                      <div className="incident-impact-card__header">
                        <strong>Replay summary</strong>
                      </div>
                      <p>{packagePayload.packageIntelligence.policyReplaySummary.shadowModeImpact}</p>
                      <div className="summary-grid">
                        <article className="summary-card summary-card--compact">
                          <span className="summary-label">Current outcome</span>
                          <strong className="summary-value">{packagePayload.packageIntelligence.policyReplaySummary.currentOutcome.blockingOrSurfacing}</strong>
                          <p>
                            blocking {packagePayload.packageIntelligence.policyReplaySummary.currentOutcome.blockingOrSurfacing} · monitoring {packagePayload.packageIntelligence.policyReplaySummary.currentOutcome.monitoringOnly} · resolved/reviewed {packagePayload.packageIntelligence.policyReplaySummary.currentOutcome.resolvedOrReviewed}
                          </p>
                        </article>
                        <article className="summary-card summary-card--compact">
                          <span className="summary-label">Proposed outcome</span>
                          <strong className="summary-value">{packagePayload.packageIntelligence.policyReplaySummary.delta.additionalRejections}</strong>
                          <p>
                            earlier denials {packagePayload.packageIntelligence.policyReplaySummary.proposedOutcome.earlierDenials} · evidence holds {packagePayload.packageIntelligence.policyReplaySummary.proposedOutcome.evidenceHolds} · narrower exceptions {packagePayload.packageIntelligence.policyReplaySummary.proposedOutcome.narrowerExceptions}
                          </p>
                        </article>
                        <article className="summary-card summary-card--compact">
                          <span className="summary-label">Blast radius</span>
                          <strong className="summary-value">{packagePayload.packageIntelligence.policyReplaySummary.blastRadius.incidentCount}</strong>
                          <p>
                            {packagePayload.packageIntelligence.policyReplaySummary.blastRadius.repoCount} repos · {packagePayload.packageIntelligence.policyReplaySummary.blastRadius.environmentCount} environments · {packagePayload.packageIntelligence.policyReplaySummary.blastRadius.workloadCount} workloads
                          </p>
                        </article>
                      </div>
                      {renderChipList(
                        packagePayload.packageIntelligence.policyReplaySummary.topCoverageGaps.map((gap) => humanizeKey(gap.gapType)),
                        "No dominant replay coverage gaps recorded for this package.",
                        "package-replay-gaps",
                        4,
                      )}
                    </article>

                    <article className="incident-impact-card incident-defense-gap">
                      <div className="incident-impact-card__header">
                        <strong>Systemic weakness</strong>
                      </div>
                      <p>{packagePayload.packageIntelligence.systemicWeaknessSummary.rootCauseHypothesis}</p>
                      <div className="chip-row">
                        <span className={`chip chip--${packagePayload.packageIntelligence.systemicWeaknessSummary.processFragility ? "warning" : "muted"}`}>
                          process fragility {packagePayload.packageIntelligence.systemicWeaknessSummary.processFragility ? "present" : "clear"}
                        </span>
                        <span className={`chip chip--${packagePayload.packageIntelligence.systemicWeaknessSummary.supplyChainBlindSpots ? "warning" : "muted"}`}>
                          supply-chain blind spots {packagePayload.packageIntelligence.systemicWeaknessSummary.supplyChainBlindSpots ? "present" : "clear"}
                        </span>
                      </div>
                      {packagePayload.packageIntelligence.systemicWeaknessSummary.topPatterns.length > 0 ? (
                        <div className="summary-grid">
                          {packagePayload.packageIntelligence.systemicWeaknessSummary.topPatterns.slice(0, 2).map((pattern) => (
                            <article className="summary-card summary-card--compact" key={`package-pattern-${pattern.patternKey}`}>
                              <div className="overview-list-item__title">
                                <strong>{pattern.title}</strong>
                                <span className={`chip chip--${priorityClass(pattern.priority)}`}>{pattern.priority}</span>
                              </div>
                              <small>{humanizeKey(pattern.patternKey)}</small>
                              {renderChipList(pattern.relatedIncidentRefs, "No related incidents attached.", `package-pattern-related-${pattern.patternKey}`, 4)}
                            </article>
                          ))}
                        </div>
                      ) : (
                        <div className="summary-list-empty">No dominant systemic weakness pattern is currently attached to this package.</div>
                      )}
                    </article>

                    <article className="incident-impact-card incident-defense-gap">
                      <div className="incident-impact-card__header">
                        <strong>Suggested package control move</strong>
                      </div>
                      <p>{packagePayload.packageIntelligence.recommendedActions.whyThisMattersNow}</p>
                      <div className="summary-grid">
                        <div>
                          <span className="summary-label">Immediate containment</span>
                          {renderValueList(packagePayload.packageIntelligence.recommendedActions.immediateContainment, "No immediate containment move recorded.", 3)}
                        </div>
                        <div>
                          <span className="summary-label">Near-term hardening</span>
                          {renderValueList(packagePayload.packageIntelligence.recommendedActions.nearTermHardening, "No near-term hardening move recorded.", 3)}
                        </div>
                        <div>
                          <span className="summary-label">Governance fix</span>
                          {renderValueList(packagePayload.packageIntelligence.recommendedActions.governanceFix, "No governance fix recorded.", 3)}
                        </div>
                      </div>
                    </article>
                  </div>
                </section>

                <section className="incident-report-section">
                  <div className="incident-report-section__header">
                    <span className="summary-label">Included cases</span>
                    <strong>{packagePayload.incidentRefs.length} linked incident reports</strong>
                  </div>
                  <div className="incident-package-table">
                    <div className="incident-package-table__row incident-package-table__row--header">
                      <span>Incident</span>
                      <span>State</span>
                      <span>Severity</span>
                      <span>Scope</span>
                      <span>Updated</span>
                    </div>
                    {packagePayload.incidents.map((item) => (
                      <button
                        type="button"
                        key={`package-item-${item.incidentID}`}
                        className="incident-package-table__row"
                        onClick={() => setSelectedIncidentID(item.incidentID)}
                      >
                        <span>
                          <strong>{item.incidentID}</strong>
                          <small>{item.title}</small>
                        </span>
                        <span>{item.state}</span>
                        <span>{item.severity} / {item.priority}</span>
                        <span>{item.scopeLabel || "-"}</span>
                        <span>{formatTimestamp(item.updatedAt || item.openedAt || item.resolvedAt)}</span>
                      </button>
                    ))}
                  </div>
                </section>

                <section className="incident-report-section">
                  <div className="incident-report-section__header">
                    <span className="summary-label">Limitations</span>
                  </div>
                  {renderValueList(
                    [
                      ...packagePayload.redactionSummary,
                      ...packagePayload.limitations,
                      ...packagePayload.packageIntelligence.defenseGapSummary.limitations,
                      ...packagePayload.packageIntelligence.policyReplaySummary.limitations,
                      ...packagePayload.packageIntelligence.systemicWeaknessSummary.limitations,
                    ],
                    "No explicit package limitations recorded.",
                    12,
                  )}
                </section>
              </div>
            </div>
          ) : null}
        </section>

        {selectedIncident ? (
          <section className="panel incident-detail-panel">
            {metricDrilldown ? (
              <section className="incident-drilldown-banner">
                <div>
                  <span className="summary-label">Metric drill-down</span>
                  <strong>{metricDrilldown.metricLabel}</strong>
                  <p>These incidents are linked to the selected trust metric through backend scorecard refs and preserved case lineage.</p>
                </div>
                {onClearMetricDrilldown ? (
                  <button type="button" className="button button-secondary" onClick={onClearMetricDrilldown}>
                    Clear metric focus
                  </button>
                ) : null}
              </section>
            ) : null}

            {metricDrilldown ? (
              <section className="incident-case-section incident-case-section--wide">
                <h3>Metric defense gaps</h3>
                {metricDefenseGapLoading ? (
                  <div className="summary-list-empty">Loading defense-gap assessment for the selected metric…</div>
                ) : metricDefenseGapError ? (
                  <p className="details-copy details-copy--error">{metricDefenseGapError}</p>
                ) : metricDefenseGaps ? (
                  <div className="incident-impact-list">
                    {metricDefenseGaps.defenseGaps.map((gap) => (
                      <article className="incident-impact-card incident-defense-gap" key={`metric-gap-${gap.gapType}`}>
                        <div className="incident-impact-card__header">
                          <strong>{gap.title}</strong>
                          <span className={`chip chip--${confidenceClass(gap.confidence)}`}>{gap.confidence}</span>
                        </div>
                        <p>{gap.whyItMatters}</p>
                        <div className="incident-evidence-grid">
                          {renderDefenseGapActions("Containment", gap.recommendedActions.containment)}
                          {renderDefenseGapActions("Hardening", gap.recommendedActions.hardening)}
                          {renderDefenseGapActions("Governance fix", gap.recommendedActions.governanceFix)}
                        </div>
                      </article>
                    ))}
                    <article className="incident-impact-card incident-defense-gap">
                      <div className="incident-impact-card__header">
                        <strong>Systemic pattern</strong>
                        <span className={`chip chip--${metricDefenseGaps.systemicPattern.present ? "warning" : "allow"}`}>
                          {metricDefenseGaps.systemicPattern.present ? "present" : "not concentrated"}
                        </span>
                      </div>
                      <p>{metricDefenseGaps.systemicPattern.summary}</p>
                      {renderChipList(metricDefenseGaps.systemicPattern.relatedIncidentRefs, "No related incidents linked beyond the current drill-down scope.", "metric-pattern", 6)}
                    </article>
                  </div>
                ) : (
                  <div className="summary-list-empty">No defense-gap assessment has been loaded for this metric yet.</div>
                )}
              </section>
            ) : null}

            {metricDrilldown ? (
              <section className="incident-case-section incident-case-section--wide">
                <h3>Policy replay and coverage gaps</h3>
                {metricReplayLoading ? (
                  <div className="summary-list-empty">Loading replay assessment for the selected metric…</div>
                ) : metricReplayError ? (
                  <p className="details-copy details-copy--error">{metricReplayError}</p>
                ) : metricReplay ? (
                  <div className="incident-impact-list">
                    <article className="incident-impact-card incident-defense-gap">
                      <div className="incident-impact-card__header">
                        <strong>Shadow-mode blast radius</strong>
                        <span className="chip chip--muted">{metricReplay.blastRadius.incidentCount} incidents</span>
                      </div>
                      <p>
                        {metricReplay.blastRadius.repoCount} repos · {metricReplay.blastRadius.environmentCount} environments · {metricReplay.blastRadius.workloadCount} workloads
                      </p>
                      {renderChipList(metricReplay.blastRadius.topScopes, "No dominant scope markers recorded.", "metric-replay-scope", 6)}
                    </article>
                    {metricReplay.coverageGaps.map((gap) => (
                      <article className="incident-impact-card incident-defense-gap" key={`metric-coverage-${gap.gapType}`}>
                        <div className="incident-impact-card__header">
                          <strong>{gap.title}</strong>
                          <span className={`chip chip--${confidenceClass(gap.confidence)}`}>{gap.confidence}</span>
                        </div>
                        <p>{gap.summary}</p>
                        <p className="incident-inline-copy">{gap.recommendedAction}</p>
                        {renderChipList(gap.evidenceRefs, "No evidence refs attached.", `metric-coverage-${gap.gapType}`, 5)}
                      </article>
                    ))}
                  </div>
                ) : (
                  <div className="summary-list-empty">No replay assessment has been loaded for this metric yet.</div>
                )}
              </section>
            ) : null}

            <div className="details-header">
              <div>
                <span className="summary-label">{selectedIncident.category}</span>
                <h2>{selectedIncident.title}</h2>
                <p>{selectedIncident.summary}</p>
              </div>
              <div className="chip-row">
                <span className={`chip chip--${severityClass(selectedIncident.severity)}`}>{selectedIncident.severity}</span>
                <span className={`chip chip--${priorityClass(selectedIncident.priority)}`}>priority {selectedIncident.priority}</span>
                <span className={`chip chip--${statusClass(selectedIncident.status)}`}>{selectedIncident.status}</span>
                <span className={`chip chip--${lifecycleClass(selectedIncident.state)}`}>{selectedIncident.state}</span>
                {selectedIncident.newActivityDetected ? <span className="chip chip--warning">new activity detected</span> : null}
              </div>
            </div>

            <p className="incident-case-copy">{selectedIncident.caseSummary}</p>

            <div className="incident-stat-grid">
              <article className="summary-card">
                <span className="summary-label">Blast Radius</span>
                <strong className="summary-value">{describeScope(selectedIncident)}</strong>
                <p className="details-copy">Primary repo: {firstOrDash(selectedIncident.affectedRepos)}</p>
              </article>
              <article className="summary-card">
                <span className="summary-label">Event Mix</span>
                <strong className="summary-value">
                  {selectedIncident.denyCount} / {selectedIncident.eventCount}
                </strong>
                <p className="details-copy">
                  {selectedIncident.denyCount} deny · {selectedIncident.errorCount} error · {selectedIncident.allowCount} allow
                </p>
              </article>
              <article className="summary-card">
                <span className="summary-label">Last Seen</span>
                <strong className="summary-value">{formatTimestamp(selectedIncident.lastSeenAt)}</strong>
                <p className="details-copy">First seen {formatTimestamp(selectedIncident.firstSeenAt)}</p>
              </article>
              <article className="summary-card">
                <span className="summary-label">Ownership</span>
                <strong className="summary-value">{selectedIncident.owner || selectedIncident.assignment.owner || "-"}</strong>
                <p className="details-copy">
                  {selectedIncident.assignment.reason
                    ? `${selectedIncident.assignment.reason} · ${formatTimestamp(selectedIncident.assignment.at)}`
                    : "No explicit owner assignment recorded yet."}
                </p>
              </article>
            </div>

            <div className="incident-case-grid">
              <section className="incident-case-section">
                <h3>What is happening</h3>
                <p>{selectedIncident.statusNarrative}</p>
              </section>
              <section className="incident-case-section">
                <h3>Likely cause</h3>
                <p>{selectedIncident.likelyCause}</p>
              </section>
              <section className="incident-case-section">
                <h3>Recommended action</h3>
                <p>{selectedIncident.recommendedAction}</p>
                <ul className="summary-list summary-list--compact">
                  {selectedIncident.remediationChecklist.map((item) => (
                    <li key={item}>
                      <span>{item}</span>
                    </li>
                  ))}
                </ul>
              </section>
              <section className="incident-case-section">
                <h3>Governance impact</h3>
                {selectedIncident.governanceImpacts.length > 0 ? (
                  <div className="incident-impact-list">
                    {selectedIncident.governanceImpacts.map((impact) => (
                      <article className="incident-impact-card" key={impact.id}>
                        <div className="incident-impact-card__header">
                          <strong>{impact.title}</strong>
                          <span className={`chip chip--${impactClass(impact.tone)}`}>{impact.tone}</span>
                        </div>
                        <p>{impact.detail}</p>
                      </article>
                    ))}
                  </div>
                ) : (
                  <div className="summary-list-empty">No additional governance impact signals were derived from this case.</div>
                )}
              </section>
              <section className="incident-case-section incident-case-section--wide">
                <h3>Defense gaps</h3>
                {defenseGapLoading ? (
                  <div className="summary-list-empty">Loading defense-gap assessment for this case…</div>
                ) : defenseGapError ? (
                  <p className="details-copy details-copy--error">{defenseGapError}</p>
                ) : incidentDefenseGaps ? (
                  <div className="incident-impact-list">
                    {incidentDefenseGaps.defenseGaps.map((gap) => (
                      <article className="incident-impact-card incident-defense-gap" key={`incident-gap-${gap.gapType}`}>
                        <div className="incident-impact-card__header">
                          <strong>{gap.title}</strong>
                          <span className={`chip chip--${confidenceClass(gap.confidence)}`}>{gap.confidence}</span>
                        </div>
                        <p>{gap.whyItMatters}</p>
                        <div className="chip-row">
                          {gap.evidenceRefs.slice(0, 5).map((ref) => (
                            <span className="chip chip--muted" key={`${gap.gapType}-${ref}`}>{ref}</span>
                          ))}
                        </div>
                        <div className="incident-evidence-grid">
                          {renderDefenseGapActions("Containment", gap.recommendedActions.containment)}
                          {renderDefenseGapActions("Hardening", gap.recommendedActions.hardening)}
                          {renderDefenseGapActions("Governance fix", gap.recommendedActions.governanceFix)}
                        </div>
                      </article>
                    ))}
                    <article className="incident-impact-card incident-defense-gap">
                      <div className="incident-impact-card__header">
                        <strong>Systemic weakness view</strong>
                        <span className={`chip chip--${incidentDefenseGaps.systemicPattern.present ? "warning" : "allow"}`}>
                          {incidentDefenseGaps.systemicPattern.present ? "repeated pattern" : "single-case dominant"}
                        </span>
                      </div>
                      <p>{incidentDefenseGaps.systemicPattern.summary}</p>
                      {renderChipList(incidentDefenseGaps.systemicPattern.relatedIncidentRefs, "No additional related incidents were linked from the current filtered scope.", "incident-pattern", 6)}
                    </article>
                  </div>
                ) : (
                  <div className="summary-list-empty">No defense-gap assessment has been loaded for this case yet.</div>
                )}
              </section>
              <section className="incident-case-section incident-case-section--wide">
                <h3>Policy replay and coverage gaps</h3>
                {incidentReplayLoading ? (
                  <div className="summary-list-empty">Loading replay assessment for this case…</div>
                ) : incidentReplayError ? (
                  <p className="details-copy details-copy--error">{incidentReplayError}</p>
                ) : incidentReplay ? (
                  <div className="incident-impact-list">
                    {incidentReplay.replayResults.map((result) => (
                      <article className="incident-impact-card incident-defense-gap" key={`incident-replay-${result.caseRef}`}>
                        <div className="incident-impact-card__header">
                          <strong>{result.title}</strong>
                          <span className={`chip chip--${confidenceClass(result.confidence)}`}>{result.confidence}</span>
                        </div>
                        <p><strong>Current outcome:</strong> {result.currentOutcome}</p>
                        <p><strong>Proposed outcome:</strong> {result.proposedOutcome}</p>
                        <p>{result.delta}</p>
                        {renderChipList(result.supportingEvidenceRefs, "No supporting evidence refs attached.", `incident-replay-${result.caseRef}`, 6)}
                      </article>
                    ))}
                    {incidentReplay.coverageGaps.map((gap) => (
                      <article className="incident-impact-card incident-defense-gap" key={`incident-coverage-${gap.gapType}`}>
                        <div className="incident-impact-card__header">
                          <strong>{gap.title}</strong>
                          <span className={`chip chip--${confidenceClass(gap.confidence)}`}>{gap.confidence}</span>
                        </div>
                        <p>{gap.summary}</p>
                        <p className="incident-inline-copy">{gap.recommendedAction}</p>
                        {renderChipList(gap.relatedIncidentRefs, "No related incidents attached.", `incident-coverage-related-${gap.gapType}`, 5)}
                      </article>
                    ))}
                    <article className="incident-impact-card incident-defense-gap">
                      <div className="incident-impact-card__header">
                        <strong>Replay blast radius</strong>
                        <span className="chip chip--muted">{incidentReplay.blastRadius.incidentCount} incident</span>
                      </div>
                      <p>
                        {incidentReplay.blastRadius.repoCount} repos · {incidentReplay.blastRadius.environmentCount} environments · {incidentReplay.blastRadius.workloadCount} workloads
                      </p>
                      {renderChipList(incidentReplay.blastRadius.topScopes, "No scope labels recorded.", "incident-replay-scope", 4)}
                    </article>
                  </div>
                ) : (
                  <div className="summary-list-empty">No replay assessment has been loaded for this case yet.</div>
                )}
              </section>
              <section className="incident-case-section">
                <h3>Metric context</h3>
                {selectedIncident.metricLinks.length > 0 ? (
                  <div className="incident-impact-list">
                    {selectedIncident.metricLinks.map((link) => (
                      <article className="incident-impact-card" key={link.metricKey}>
                        <div className="incident-impact-card__header">
                          <strong>{link.metricLabel}</strong>
                          <span className="chip chip--muted">impact {link.impactWeight}</span>
                        </div>
                        <p>{link.linkReason}</p>
                        <div className="chip-row">
                          {link.supportingRefs.slice(0, 5).map((ref) => (
                            <span className="chip chip--muted" key={`${link.metricKey}-${ref}`}>{ref}</span>
                          ))}
                        </div>
                      </article>
                    ))}
                  </div>
                ) : (
                  <div className="summary-list-empty">No scorecard metric linkage was attached to this incident.</div>
                )}
              </section>
              <section className="incident-case-section">
                <h3>Lifecycle and ownership</h3>
                <div className="incident-evidence-grid">
                  <div>
                    <span className="summary-label">Lifecycle</span>
                    <p className="incident-inline-copy">
                      {selectedIncident.state} · updated {formatTimestamp(selectedIncident.updatedAt || selectedIncident.lastActivityAt)}
                    </p>
                    <p className="incident-inline-copy">
                      Opened {formatTimestamp(selectedIncident.openedAt || selectedIncident.firstSeenAt)}
                    </p>
                  </div>
                  <div>
                    <span className="summary-label">Scope anchor</span>
                    <p className="incident-inline-copy">
                      {selectedIncident.scopeType || "-"} · {selectedIncident.scopeRef || "-"}
                    </p>
                    <p className="incident-inline-copy">
                      {selectedIncident.repository || selectedIncident.environment || selectedIncident.tenantID || "No primary scope anchor recorded."}
                    </p>
                  </div>
                  <div>
                    <span className="summary-label">Assignment</span>
                    <p className="incident-inline-copy">{selectedIncident.assignment.owner || "Unassigned"}</p>
                    <p className="incident-inline-copy">
                      {selectedIncident.assignment.reason || "No assignment rationale recorded."}
                      {selectedIncident.assignment.by ? ` · ${selectedIncident.assignment.by}` : ""}
                    </p>
                  </div>
                </div>
              </section>
              {canManage ? (
                <section className="incident-case-section incident-case-section--wide">
                  <h3>Case actions</h3>
                  <div className="incident-actions-grid">
                    <div className="incident-action-block">
                      <span className="summary-label">Lifecycle</span>
                      <input
                        className="input"
                        value={lifecycleSummary}
                        onChange={(event) => setLifecycleSummary(event.target.value)}
                        placeholder="Optional rationale for acknowledge, watch, or reopen"
                      />
                      <div className="chip-row">
                        <button className="button" disabled={actionSubmitting || !onAcknowledge} onClick={() => void runAction(() => onAcknowledge?.(selectedIncident.id, lifecycleSummary) || Promise.resolve())}>
                          Acknowledge
                        </button>
                        <button className="button" disabled={actionSubmitting || !onWatch} onClick={() => void runAction(() => onWatch?.(selectedIncident.id, lifecycleSummary) || Promise.resolve())}>
                          Watch
                        </button>
                        <button className="button button-secondary" disabled={actionSubmitting || !onReopen} onClick={() => void runAction(() => onReopen?.(selectedIncident.id, lifecycleSummary) || Promise.resolve())}>
                          Reopen
                        </button>
                      </div>
                    </div>

                    <div className="incident-action-block">
                      <span className="summary-label">Owner assignment</span>
                      <input className="input" value={assignmentOwner} onChange={(event) => setAssignmentOwner(event.target.value)} placeholder="owner" />
                      <input className="input" value={assignmentReason} onChange={(event) => setAssignmentReason(event.target.value)} placeholder="assignment reason" />
                      <button
                        className="button"
                        disabled={actionSubmitting || !onAssign || !assignmentOwner.trim() || !assignmentReason.trim()}
                        onClick={() => void runAction(() => onAssign?.(selectedIncident.id, assignmentOwner.trim(), assignmentReason.trim()) || Promise.resolve())}
                      >
                        Save owner
                      </button>
                    </div>

                    <div className="incident-action-block">
                      <span className="summary-label">Operator note</span>
                      <textarea
                        className="textarea"
                        value={noteDraft}
                        onChange={(event) => setNoteDraft(event.target.value)}
                        rows={4}
                        placeholder="Add bounded investigation context without changing derived truth."
                      />
                      <button
                        className="button"
                        disabled={actionSubmitting || !onAddNote || !noteDraft.trim()}
                        onClick={() => void runAction(() => onAddNote?.(selectedIncident.id, noteDraft.trim()) || Promise.resolve())}
                      >
                        Add note
                      </button>
                    </div>

                    {canResolve ? (
                      <div className="incident-action-block incident-action-block--wide">
                        <span className="summary-label">Structured resolution</span>
                        <div className="incident-resolution-grid">
                          <select className="input" value={resolutionType} onChange={(event) => setResolutionType(event.target.value)}>
                            <option value="fixed">fixed</option>
                            <option value="mitigated">mitigated</option>
                            <option value="accepted_risk">accepted_risk</option>
                            <option value="false_positive">false_positive</option>
                            <option value="duplicate">duplicate</option>
                            <option value="temporary_containment">temporary_containment</option>
                          </select>
                          <input
                            className="input"
                            value={resolutionSummary}
                            onChange={(event) => setResolutionSummary(event.target.value)}
                            placeholder="resolution summary"
                          />
                          <input
                            className="input"
                            value={resolutionRefs}
                            onChange={(event) => setResolutionRefs(event.target.value)}
                            placeholder="evidence refs, comma separated"
                          />
                        </div>
                        <textarea
                          className="textarea"
                          value={resolutionDetails}
                          onChange={(event) => setResolutionDetails(event.target.value)}
                          rows={4}
                          placeholder="resolution details"
                        />
                        <label className="checkbox">
                          <input type="checkbox" checked={followUpRequired} onChange={(event) => setFollowUpRequired(event.target.checked)} />
                          <span>Follow-up still required</span>
                        </label>
                        <button
                          className="button"
                          disabled={actionSubmitting || !onResolve || !resolutionSummary.trim()}
                          onClick={() =>
                            void runAction(() =>
                              onResolve?.(selectedIncident.id, {
                                resolution_type: resolutionType,
                                resolution_summary: resolutionSummary.trim(),
                                resolution_details: resolutionDetails.trim() || undefined,
                                resolution_refs: resolutionRefs.split(",").map((value) => value.trim()).filter(Boolean),
                                follow_up_required: followUpRequired,
                              }) || Promise.resolve(),
                            )
                          }
                        >
                          Resolve incident
                        </button>
                      </div>
                    ) : null}
                  </div>
                  {actionError ? <p className="details-copy details-copy--error">{actionError}</p> : null}
                </section>
              ) : null}
              <section className="incident-case-section incident-case-section--wide">
                <h3>Impacted scope</h3>
                <div className="incident-scope-grid">
                  <div>
                    <span className="summary-label">Repos and environments</span>
                    <div className="chip-row">
                      {selectedIncident.affectedRepos.slice(0, 6).map((repo) => (
                        <span className="chip chip--muted" key={repo}>{repo}</span>
                      ))}
                      {selectedIncident.affectedEnvironments.slice(0, 6).map((environment) => (
                        <span className="chip chip--muted" key={environment}>{environment}</span>
                      ))}
                      {selectedIncident.affectedTenants.slice(0, 4).map((tenant) => (
                        <span className="chip chip--muted" key={tenant}>{tenant}</span>
                      ))}
                    </div>
                  </div>
                  <div>
                    <span className="summary-label">Workloads and namespaces</span>
                    <div className="chip-row">
                      {selectedIncident.affectedWorkloads.slice(0, 6).map((workload) => (
                        <span className="chip chip--muted" key={workload}>{workload}</span>
                      ))}
                      {selectedIncident.affectedNamespaces.slice(0, 4).map((namespace) => (
                        <span className="chip chip--muted" key={namespace}>{namespace}</span>
                      ))}
                    </div>
                  </div>
                  <div>
                    <span className="summary-label">Images and control-plane components</span>
                    <div className="chip-row">
                      {selectedIncident.affectedImages.slice(0, 4).map((image) => (
                        <span className="chip chip--muted" key={image}>{image}</span>
                      ))}
                      {selectedIncident.affectedComponents.slice(0, 4).map((component) => (
                        <span className="chip chip--muted" key={component}>{component}</span>
                      ))}
                    </div>
                  </div>
                </div>
              </section>
            </div>

            <div className="incident-case-grid">
              <section className="incident-case-section incident-case-section--wide">
                <h3>Incident timeline</h3>
                <ol className="incident-timeline">
                  {selectedIncident.timeline.map((entry) => (
                    <li className="incident-timeline__item" key={entry.id}>
                      <div className="incident-timeline__meta">
                        <span className={`chip chip--${outcomeClass(entry.outcome)}`}>{entry.outcome}</span>
                        <strong>{entry.title}</strong>
                        <small>{formatTimestamp(entry.timestamp)}</small>
                      </div>
                      <p>{entry.summary}</p>
                      {entry.requestID ? <code>{entry.requestID}</code> : null}
                    </li>
                  ))}
                </ol>
              </section>
              <section className="incident-case-section">
                <h3>Reason stack</h3>
                {renderValueList(selectedIncident.relatedReasons, "No additional reason codes recorded.", 6)}
              </section>
              <section className="incident-case-section">
                <h3>Linked refs</h3>
                <div className="incident-evidence-grid">
                  <div>
                    <span className="summary-label">Findings</span>
                    {renderValueList(selectedIncident.findingRefs, "No finding refs attached.", 6)}
                  </div>
                  <div>
                    <span className="summary-label">Guidance and scorecard</span>
                    {renderValueList(
                      [...selectedIncident.guidanceRefs, ...selectedIncident.scorecardRefs],
                      "No guidance or scorecard refs attached.",
                      6,
                    )}
                  </div>
                </div>
              </section>
              <section className="incident-case-section incident-case-section--wide">
                <h3>Case export</h3>
                <div className="incident-evidence-grid">
                  <div>
                    <span className="summary-label">Export linkage</span>
                    <p className="incident-inline-copy">
                      Load the canonical case export to inspect the incident package with lifecycle, history, linked metrics, and evidence lineage.
                    </p>
                    <div className="chip-row">
                      {(["internal", "auditor_safe", "customer_safe"] as IncidentReportAudience[]).map((audience) => (
                        <button
                          type="button"
                          key={audience}
                          className={`button ${reportAudience === audience ? "" : "button-secondary"}`}
                          onClick={() => setReportAudience(audience)}
                        >
                          {audience === "internal" ? "Internal" : audience === "auditor_safe" ? "Auditor-Safe" : "Customer-Safe"}
                        </button>
                      ))}
                    </div>
                    <div className="chip-row">
                      <button
                        type="button"
                        className="button"
                        disabled={exportLoading || !onLoadExport}
                        onClick={() => void loadExport(selectedIncident.id, reportAudience)}
                      >
                        {exportLoading ? "Loading export…" : "Load current audience"}
                      </button>
                      {exportPayload ? (
                        <button
                          type="button"
                          className="button button-secondary"
                          onClick={() => downloadExport(exportPayload)}
                        >
                          Download JSON
                        </button>
                      ) : null}
                    </div>
                    {exportError ? <p className="details-copy details-copy--error">{exportError}</p> : null}
                  </div>
                  <div>
                    <span className="summary-label">Report mode</span>
                    {exportPayload ? (
                      <>
                        <div className="chip-row">
                          <span className={`chip chip--${exportPayload.redacted ? "warning" : "allow"}`}>
                            {exportPayload.redacted ? "Redacted" : "Internal canonical"}
                          </span>
                          <span className="chip chip--muted">{exportPayload.audience}</span>
                        </div>
                        <p className="incident-inline-copy">
                          {exportPayload.history.length} history entries · {exportPayload.relatedEventRefs.length} linked events
                        </p>
                        <p className="incident-inline-copy">
                          Generated {formatTimestamp(exportPayload.generatedAt)} · state {exportPayload.state}
                        </p>
                        <div className="chip-row">
                          <button
                            type="button"
                            className={`button ${reportMode === "report" ? "" : "button-secondary"}`}
                            onClick={() => setReportMode("report")}
                          >
                            Case Report
                          </button>
                          <button
                            type="button"
                            className={`button ${reportMode === "json" ? "" : "button-secondary"}`}
                            onClick={() => {
                              setReportMode("json");
                              setHandoffMode(false);
                            }}
                          >
                            {exportPayload.audience === "internal" ? "Canonical JSON" : "Redacted JSON"}
                          </button>
                        </div>
                        {reportMode === "report" ? (
                          <div className="chip-row">
                            <button
                              type="button"
                              className={`button ${handoffMode ? "" : "button-secondary"}`}
                              onClick={() => {
                                setPackageHandoffMode(false);
                                setHandoffMode((current) => !current);
                              }}
                            >
                              {handoffMode ? "Screen layout" : "Print-friendly"}
                            </button>
                            {handoffMode ? (
                              <button type="button" className="button button-secondary" onClick={printHandoff}>
                                Print / Save PDF
                              </button>
                            ) : null}
                          </div>
                        ) : null}
                      </>
                    ) : (
                      <div className="summary-list-empty">No export payload loaded for this case yet.</div>
                    )}
                  </div>
                </div>
                {exportPayload ? (
                  <div className={`incident-export-preview ${handoffMode && reportMode === "report" ? "incident-export-preview--handoff" : ""}`}>
                    {reportMode === "report" ? (
                      <div className={`incident-report-surface ${handoffMode ? "incident-report-surface--handoff" : ""}`}>
                        <section className="incident-report-section incident-report-section--hero">
                          <div className="incident-report-header">
                            <div>
                              <span className="summary-label">{handoffMode ? "Auditor handoff" : "Case report"}</span>
                              <h3>{exportPayload.title}</h3>
                              <p>{exportPayload.incidentID} · generated {formatTimestamp(exportPayload.generatedAt)}</p>
                            </div>
                            <div className="chip-row">
                              <span className={`chip chip--${exportPayload.redacted ? "warning" : "allow"}`}>
                                {exportPayload.redacted ? "Redacted" : "Internal canonical"}
                              </span>
                              <span className="chip chip--muted">{exportPayload.audience}</span>
                              {handoffMode ? <span className="chip chip--muted">print-friendly</span> : null}
                            </div>
                          </div>
                          <div className="incident-report-header__meta">
                            <span>state {exportPayload.state}</span>
                            <span>severity {exportPayload.severity}</span>
                            <span>priority {exportPayload.priority}</span>
                            {exportPayload.newActivityDetected ? <span>new activity detected</span> : null}
                          </div>
                        </section>

                        <section className="incident-report-section">
                          <div className="incident-report-section__header">
                            <span className="summary-label">Case header</span>
                            <strong>{exportPayload.incidentID}</strong>
                          </div>
                          {exportPayload.redacted ? (
                            <div className="chip-row">
                              <span className="chip chip--warning">Redacted report</span>
                              <span className="chip chip--muted">{exportPayload.audience}</span>
                            </div>
                          ) : null}
                          <div className="incident-evidence-grid">
                            <div>
                              <p className="incident-inline-copy"><strong>{exportPayload.title}</strong></p>
                              <p className="incident-inline-copy">{exportPayload.summary}</p>
                            </div>
                            <div>
                              <div className="chip-row">
                                <span className={`chip chip--${severityClass(exportPayload.severity)}`}>{exportPayload.severity}</span>
                                <span className={`chip chip--${priorityClass(exportPayload.priority)}`}>priority {exportPayload.priority}</span>
                                <span className={`chip chip--${lifecycleClass(exportPayload.state)}`}>{exportPayload.state}</span>
                                {exportPayload.newActivityDetected ? <span className="chip chip--warning">new activity detected</span> : null}
                              </div>
                              <p className="incident-inline-copy">Owner: {exportPayload.owner || "-"}</p>
                              <p className="incident-inline-copy">
                                Opened {formatTimestamp(exportPayload.openedAt)} · Updated {formatTimestamp(exportPayload.updatedAt)}
                              </p>
                            </div>
                          </div>
                        </section>

                        <section className="incident-report-section">
                          <div className="incident-report-section__header">
                            <span className="summary-label">Scope context</span>
                          </div>
                          <div className="incident-evidence-grid">
                            <div>
                              <p className="incident-inline-copy">{exportPayload.scopeType || "-"} · {exportPayload.scopeRef || "-"}</p>
                              <p className="incident-inline-copy">{exportPayload.repository || exportPayload.environment || exportPayload.tenantID || "No primary scope anchor recorded."}</p>
                            </div>
                            <div>
                              {renderChipList(
                                [exportPayload.tenantID, exportPayload.clusterID, exportPayload.environment, exportPayload.repository].filter(Boolean) as string[],
                                "No scoped context recorded.",
                                "report-scope",
                                8,
                              )}
                            </div>
                          </div>
                        </section>

                        <section className="incident-report-section">
                          <div className="incident-report-section__header">
                            <span className="summary-label">Verdict and metric context</span>
                          </div>
                          <div className="incident-evidence-grid">
                            <div>
                              <span className="summary-label">Reason codes</span>
                              {renderValueList(exportPayload.reasonCodes, "No reason codes attached.", 8)}
                            </div>
                            <div>
                              <span className="summary-label">Linked metrics</span>
                              {exportPayload.metricLinks.length > 0 ? (
                                <div className="incident-impact-list">
                                  {exportPayload.metricLinks.map((link) => (
                                    <article className="incident-impact-card" key={`report-${link.metricKey}`}>
                                      <div className="incident-impact-card__header">
                                        <strong>{link.metricLabel}</strong>
                                        <span className="chip chip--muted">impact {link.impactWeight}</span>
                                      </div>
                                      <p>{link.linkReason}</p>
                                      {renderChipList(link.supportingRefs, "No supporting refs attached.", `report-metric-${link.metricKey}`, 5)}
                                    </article>
                                  ))}
                                </div>
                              ) : (
                                <div className="summary-list-empty">No linked metric context attached.</div>
                              )}
                            </div>
                          </div>
                        </section>

                        <section className="incident-report-section">
                          <div className="incident-report-section__header">
                            <span className="summary-label">Evidence and governance</span>
                          </div>
                          <div className="incident-evidence-grid">
                            <div>
                              <span className="summary-label">Evidence refs</span>
                              {renderValueList(exportPayload.evidenceRefs, "No evidence refs attached.", 8)}
                              <span className="summary-label">Linked event refs</span>
                              {renderValueList(
                                exportPayload.relatedEventRefs.map((eventRef) => eventRef.requestID || `event:${eventRef.eventID}`),
                                "No linked event refs attached.",
                                8,
                              )}
                            </div>
                            <div>
                              <span className="summary-label">Governance impact</span>
                              {exportPayload.governanceImpacts.length > 0 ? (
                                <div className="incident-impact-list">
                                  {exportPayload.governanceImpacts.map((impact) => (
                                    <article className="incident-impact-card" key={`report-impact-${impact.id}`}>
                                      <div className="incident-impact-card__header">
                                        <strong>{impact.title}</strong>
                                        <span className={`chip chip--${impactClass(impact.tone)}`}>{impact.tone}</span>
                                      </div>
                                      <p>{impact.detail}</p>
                                    </article>
                                  ))}
                                </div>
                              ) : (
                                <div className="summary-list-empty">No governance impact attached.</div>
                              )}
                            </div>
                          </div>
                        </section>

                        <section className="incident-report-section">
                          <div className="incident-report-section__header">
                            <span className="summary-label">Resolution and operator context</span>
                          </div>
                          <div className="incident-evidence-grid">
                            <div>
                              <span className="summary-label">Resolution</span>
                              <p className="incident-inline-copy">{exportPayload.resolution.type || "-"} · {exportPayload.resolution.summary || "No structured resolution recorded."}</p>
                              {exportPayload.resolution.details ? <p className="incident-inline-copy">{exportPayload.resolution.details}</p> : null}
                              {exportPayload.resolution.refs.length > 0 ? renderValueList(exportPayload.resolution.refs, "No resolution refs attached.", 6) : null}
                            </div>
                            <div>
                              <span className="summary-label">History and notes</span>
                              <p className="incident-inline-copy">{exportPayload.history.length} history entries · {exportPayload.notes.length} operator notes</p>
                              {renderValueList(
                                exportPayload.history.slice(-5).map((entry) => `${entry.kind.replaceAll("_", " ")} · ${entry.summary}`),
                                "No lifecycle history attached.",
                                5,
                              )}
                            </div>
                          </div>
                        </section>

                        <section className="incident-report-section">
                          <div className="incident-report-section__header">
                            <span className="summary-label">Limitations</span>
                          </div>
                          {renderValueList(
                            [...exportPayload.redactionSummary, ...exportPayload.limitations],
                            "No explicit export limitations recorded.",
                            10,
                          )}
                        </section>
                      </div>
                    ) : (
                      <pre className="incident-export-json">{JSON.stringify(exportPayload, null, 2)}</pre>
                    )}
                  </div>
                ) : null}
              </section>
              <section className="incident-case-section">
                <h3>Evidence pack</h3>
                <div className="incident-evidence-grid">
                  <div>
                    <span className="summary-label">Request IDs</span>
                    {renderValueList(selectedIncident.evidencePack.requestIDs, "No request IDs captured.", 4)}
                  </div>
                  <div>
                    <span className="summary-label">Digests and bundles</span>
                    {renderValueList(
                      [...selectedIncident.evidencePack.digests, ...selectedIncident.evidencePack.bundles],
                      "No digest or bundle refs captured.",
                      6,
                    )}
                  </div>
                  <div>
                    <span className="summary-label">Exceptions and vulnerabilities</span>
                    {renderValueList(
                      [...selectedIncident.evidencePack.exceptions, ...selectedIncident.evidencePack.vulnerabilities],
                      "No exception or vulnerability refs captured.",
                      6,
                    )}
                  </div>
                  <div>
                    <span className="summary-label">Primary reason</span>
                    <p className="incident-inline-copy">{selectedIncident.primaryReason}</p>
                    {selectedIncident.evidenceRefs.length > 0 ? <code className="truncate">{firstOrDash(selectedIncident.evidenceRefs)}</code> : null}
                  </div>
                </div>
              </section>
              <section className="incident-case-section incident-case-section--wide">
                <h3>Resolution</h3>
                {selectedIncident.resolution.type || selectedIncident.resolutionSummary ? (
                  <div className="incident-evidence-grid">
                    <div>
                      <span className="summary-label">Resolution type</span>
                      <p className="incident-inline-copy">{selectedIncident.resolution.type || "-"}</p>
                    </div>
                    <div>
                      <span className="summary-label">Summary</span>
                      <p className="incident-inline-copy">{selectedIncident.resolution.summary || selectedIncident.resolutionSummary || "-"}</p>
                    </div>
                    <div>
                      <span className="summary-label">By and when</span>
                      <p className="incident-inline-copy">
                        {selectedIncident.resolution.by || "-"} · {formatTimestamp(selectedIncident.resolution.at || selectedIncident.resolvedAt)}
                      </p>
                    </div>
                    <div>
                      <span className="summary-label">Resolution refs</span>
                      {renderValueList(selectedIncident.resolution.refs, "No structured resolution refs attached.", 6)}
                    </div>
                  </div>
                ) : (
                  <div className="summary-list-empty">No structured resolution has been recorded for this incident yet.</div>
                )}
              </section>
              <section className="incident-case-section incident-case-section--wide">
                <h3>Operator notes and history</h3>
                <div className="incident-evidence-grid">
                  <div>
                    <span className="summary-label">Notes</span>
                    {selectedIncident.notes.length > 0 ? (
                      <ul className="summary-list summary-list--compact">
                        {selectedIncident.notes.slice(-6).map((note) => (
                          <li key={note.id}>
                            <span>{note.note}</span>
                            <small>{note.actor ? ` · ${note.actor}` : ""}{note.timestamp ? ` · ${formatTimestamp(note.timestamp)}` : ""}</small>
                          </li>
                        ))}
                      </ul>
                    ) : (
                      <div className="summary-list-empty">No operator notes recorded yet.</div>
                    )}
                  </div>
                  <div>
                    <span className="summary-label">History</span>
                    {selectedIncident.history.length > 0 ? (
                      <ul className="summary-list summary-list--compact">
                        {selectedIncident.history.slice(-8).map((entry) => (
                          <li key={entry.id}>
                            <span>{entry.kind.replaceAll("_", " ")} · {entry.summary}</span>
                            <small>{entry.actor ? ` · ${entry.actor}` : ""}{entry.timestamp ? ` · ${formatTimestamp(entry.timestamp)}` : ""}</small>
                          </li>
                        ))}
                      </ul>
                    ) : (
                      <div className="summary-list-empty">No lifecycle history recorded yet.</div>
                    )}
                  </div>
                </div>
              </section>
            </div>
          </section>
        ) : null}

        <section className="panel incident-events-panel">
          <div className="details-header">
            <div>
              <span className="summary-label">Underlying evidence</span>
              <h2>Related events and raw details</h2>
              <p>Use the case file above to understand the incident first. Then inspect the individual events and raw evidence records below.</p>
            </div>
          </div>
          <div className="content-grid">
            <EventsTable
              events={selectedIncident?.events || []}
              selectedEventID={selectedEvent?.id || null}
              onSelect={setSelectedEvent}
              loading={false}
              error={null}
              title="Related Events"
              emptyMessage="No raw events are attached to this investigation."
            />
            <EventDetails event={selectedEvent} />
          </div>
        </section>
      </div>
    </section>
  );
}
