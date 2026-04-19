import type {
  AuditHealth,
  DriftStatsResponse,
  ExceptionReport,
  Summary,
  SyncStatus,
  TopViolatorsResponse,
  TrendsResponse,
} from "./types";

export type PostureLevel = "stable" | "degraded" | "at-risk";
export type IncidentSeverity = "critical" | "high" | "medium" | "low";
export type ActionPriority = "now" | "next" | "watch";
export type TrustMetricStatus = "verified" | "partial" | "gap" | "unknown";
export type GuidanceConfidence = "high" | "medium" | "limited";

export type OverviewMetric = {
  id: string;
  label: string;
  value: number;
  tone: "neutral" | "allow" | "deny" | "drift" | "warning";
  trend: string;
  context: string;
};

export type OverviewIncident = {
  id: string;
  severity: IncidentSeverity;
  title: string;
  summary: string;
  scope: string;
  evidence: string;
  action: string;
};

export type OverviewAction = {
  id: string;
  priority: ActionPriority;
  title: string;
  detail: string;
  source: string;
};

export type OverviewPosture = {
  level: PostureLevel;
  title: string;
  summary: string;
  changed: string;
  scope: string;
  reasons: string[];
};

export type OverviewTrustMetric = {
  id: string;
  label: string;
  score: number;
  weight: number;
  status: TrustMetricStatus;
  detail: string;
  defenseGap: string;
  rootWeakness: string;
  nextMove: string;
};

export type OverviewTrustSummary = {
  score: number;
  grade: "A" | "B" | "C" | "D" | "F";
  summary: string;
  derivedFrom: string;
  metrics: OverviewTrustMetric[];
  draggers: string[];
};

export type OverviewGuidanceItem = {
  id: string;
  title: string;
  confidence: GuidanceConfidence;
  summary: string;
  nextStep: string;
  source: string;
};

export type OverviewModel = {
  posture: OverviewPosture;
  trust: OverviewTrustSummary;
  guidance: OverviewGuidanceItem[];
  incidents: OverviewIncident[];
  actions: OverviewAction[];
  metrics: OverviewMetric[];
  denyDrivers: Array<{ label: string; count: number }>;
  blastRadius: Array<{ label: string; detail: string; count: number }>;
};

type Inputs = {
  health: AuditHealth | null;
  summary: Summary | null;
  trends: TrendsResponse | null;
  topViolators: TopViolatorsResponse | null;
  driftStats: DriftStatsResponse | null;
  exceptionReport: ExceptionReport | null;
  syncStatus: SyncStatus | null;
};

function humanizeReason(value: string) {
  return value.replace(/[_-]+/g, " ").trim();
}

function pluralize(value: number, singular: string, plural = `${singular}s`) {
  return `${value} ${value === 1 ? singular : plural}`;
}

function normalizeReason(summary: Summary | null, matcher: string) {
  if (!summary) {
    return 0;
  }
  const normalizedMatcher = matcher.trim().toLowerCase();
  return summary.top_deny_reasons.find((item) => item.reason.toLowerCase() === normalizedMatcher)?.count ?? 0;
}

function findReasonIncludes(summary: Summary | null, matcher: string) {
  if (!summary) {
    return 0;
  }
  const normalizedMatcher = matcher.trim().toLowerCase();
  return summary.top_deny_reasons
    .filter((item) => item.reason.toLowerCase().includes(normalizedMatcher))
    .reduce((total, item) => total + item.count, 0);
}

function latestTrendDelta(
  trends: TrendsResponse | null,
  key: "allow_count" | "deny_count" | "error_count",
) {
  if (!trends || trends.buckets.length < 2) {
    return null;
  }
  const current = trends.buckets[trends.buckets.length - 1]?.[key] ?? 0;
  const previous = trends.buckets[trends.buckets.length - 2]?.[key] ?? 0;
  return current - previous;
}

function formatDelta(delta: number | null, noun: string) {
  if (delta == null || delta === 0) {
    return `flat vs previous bucket`;
  }
  const direction = delta > 0 ? "up" : "down";
  return `${direction} ${Math.abs(delta)} ${noun} vs previous bucket`;
}

function severityFromCount(count: number): IncidentSeverity {
  if (count >= 40) return "critical";
  if (count >= 20) return "high";
  if (count >= 8) return "medium";
  return "low";
}

function topViolatorScope(topViolators: TopViolatorsResponse | null) {
  const item = topViolators?.items[0];
  if (!item) {
    return "current filtered scope";
  }
  return item.key;
}

function topDriftScope(driftStats: DriftStatsResponse | null) {
  const workload = driftStats?.top_drifted_workloads[0];
  if (!workload) {
    return "runtime workloads in current scope";
  }
  return workload.namespace ? `${workload.namespace}/${workload.workload}` : workload.workload;
}

function clampScore(value: number) {
  return Math.max(0, Math.min(100, Math.round(value)));
}

function gradeFromScore(score: number): OverviewTrustSummary["grade"] {
  if (score >= 90) return "A";
  if (score >= 80) return "B";
  if (score >= 70) return "C";
  if (score >= 60) return "D";
  return "F";
}

function metricStatus(score: number): TrustMetricStatus {
  if (score >= 90) return "verified";
  if (score >= 75) return "partial";
  if (score >= 55) return "gap";
  return "unknown";
}

function confidenceForCount(count: number): GuidanceConfidence {
  if (count >= 15) return "high";
  if (count >= 4) return "medium";
  return "limited";
}

export function buildOverviewModel({
  health,
  summary,
  trends,
  topViolators,
  driftStats,
  exceptionReport,
  syncStatus,
}: Inputs): OverviewModel {
  const workflowMismatch = normalizeReason(summary, "workflow mismatch");
  const imageDriftReason = normalizeReason(summary, "image drift");
  const signatureFailures = findReasonIncludes(summary, "signature verification failed");
  const digestPinningGaps = findReasonIncludes(summary, "digest-pinned");
  const driftDenies = summary?.recent_runtime_drift_deny ?? driftStats?.total_runtime_drift_denies ?? 0;
  const pendingExceptions = exceptionReport?.pending?.length ?? 0;
  const activeExceptions = exceptionReport?.active?.length ?? 0;
  const totalErrors = summary?.total_error ?? 0;
  const denyDelta = latestTrendDelta(trends, "deny_count");
  const allowDelta = latestTrendDelta(trends, "allow_count");
  const errorDelta = latestTrendDelta(trends, "error_count");
  const monitoredSignals =
    (summary?.counts_by_event_type.artifact_verification_result ?? 0)
    + (summary?.counts_by_event_type.deploy_gate_decision ?? 0)
    + (summary?.counts_by_event_type.runtime_drift_result ?? 0);

  let postureLevel: PostureLevel = "stable";
  if (health?.status !== "ok" || totalErrors > 0 || workflowMismatch >= 25 || imageDriftReason >= 25 || driftDenies >= 8) {
    postureLevel = "at-risk";
  } else if ((summary?.total_deny ?? 0) >= 20 || pendingExceptions > 0 || activeExceptions > 0 || driftDenies > 0) {
    postureLevel = "degraded";
  }

  const postureReasons: string[] = [];
  if (workflowMismatch > 0) {
    postureReasons.push(`${pluralize(workflowMismatch, "deny")} are dominated by workflow mismatch.`);
  }
  if (imageDriftReason > 0 || driftDenies > 0) {
    postureReasons.push(`${pluralize(Math.max(imageDriftReason, driftDenies), "runtime drift signal")} still need reconciliation.`);
  }
  if (signatureFailures + digestPinningGaps > 0) {
    postureReasons.push(`${pluralize(signatureFailures + digestPinningGaps, "artifact trust gap")} came from signature or digest pinning issues.`);
  }
  if (pendingExceptions > 0 || activeExceptions > 0) {
    postureReasons.push(`${pluralize(pendingExceptions + activeExceptions, "exception record")} still need review or cleanup.`);
  }
  if (totalErrors > 0) {
    postureReasons.push(`${pluralize(totalErrors, "control-plane error")} surfaced in the current time window.`);
  }
  if (postureReasons.length === 0) {
    postureReasons.push("Current denies look explainable and contained within normal control-plane enforcement.");
  }

  const posture: OverviewPosture = {
    level: postureLevel,
    title:
      postureLevel === "stable"
        ? "Current posture is stable."
        : postureLevel === "degraded"
          ? "Current posture is degraded."
          : "Current posture is at risk.",
    summary:
      postureLevel === "stable"
        ? "Controls are enforcing as expected and the current signal mix looks explainable."
        : postureLevel === "degraded"
          ? "The control plane is blocking bad changes, but active signals now need operator review."
          : "The control plane is protecting runtime, but current trust and drift signals require immediate investigation.",
    changed:
      denyDelta != null
        ? `Deny activity is ${denyDelta > 0 ? "rising" : denyDelta < 0 ? "cooling down" : "unchanged"} (${formatDelta(denyDelta, "denies")}).`
        : "No recent baseline comparison is available yet.",
    scope: `Primary affected scope: ${topViolatorScope(topViolators)} · runtime focus: ${topDriftScope(driftStats)}`,
    reasons: postureReasons,
  };

  const incidents: OverviewIncident[] = [];
  if (workflowMismatch > 0) {
    incidents.push({
      id: "workflow-mismatch",
      severity: severityFromCount(workflowMismatch),
      title: "Workflow trust drift",
      summary: `${pluralize(workflowMismatch, "denied decision")} cite workflow mismatch as the dominant reason.`,
      scope: topViolatorScope(topViolators),
      evidence: `Top deny reason: workflow mismatch · dominant violator: ${topViolatorScope(topViolators)}`,
      action: "Review the latest signer or workflow ref changes and compare them with the trusted workflow policy.",
    });
  }
  if (imageDriftReason > 0 || driftDenies > 0) {
    const impact = Math.max(imageDriftReason, driftDenies);
    incidents.push({
      id: "runtime-image-drift",
      severity: severityFromCount(impact),
      title: "Runtime image drift",
      summary: `${pluralize(impact, "drift-related signal")} indicate workloads diverging from approved images.`,
      scope: topDriftScope(driftStats),
      evidence: `Runtime drift (24h): ${driftDenies} · top workload: ${topDriftScope(driftStats)}`,
      action: "Inspect the highest-drift workload, compare against the last approved digest, and reconcile the parent controller spec first.",
    });
  }
  if (signatureFailures + digestPinningGaps > 0) {
    const impact = signatureFailures + digestPinningGaps;
    incidents.push({
      id: "artifact-trust-hygiene",
      severity: severityFromCount(impact),
      title: "Artifact trust hygiene gaps",
      summary: `${pluralize(impact, "deny")} came from signature verification or digest-pinning failures.`,
      scope: topViolatorScope(topViolators),
      evidence: `${signatureFailures} signature failures · ${digestPinningGaps} digest pinning issues`,
      action: "Fix digest pinning and artifact trust errors before widening exceptions or retrying the pipeline.",
    });
  }
  if (pendingExceptions > 0 || activeExceptions > 0) {
    const impact = pendingExceptions + activeExceptions;
    incidents.push({
      id: "exception-governance",
      severity: severityFromCount(impact),
      title: "Exception governance needs review",
      summary: `${pluralize(impact, "exception")} remain pending or active in the current scope.`,
      scope: exceptionReport?.active?.[0]?.repo || "current governance queue",
      evidence: `${pendingExceptions} pending · ${activeExceptions} active`,
      action: pendingExceptions > 0
        ? "Review pending exception requests and narrow or reject any over-broad scope."
        : "Revisit active exceptions, confirm evidence, and revoke anything no longer justified.",
    });
  }
  if (totalErrors > 0) {
    incidents.push({
      id: "control-plane-errors",
      severity: severityFromCount(totalErrors),
      title: "Control-plane errors surfaced",
      summary: `${pluralize(totalErrors, "error event")} indicate at least one failed decision path or backend issue.`,
      scope: "audit or verifier path",
      evidence: `Error count: ${totalErrors} · ${formatDelta(errorDelta, "errors")}`,
      action: "Open the latest error event and confirm audit-writer, verifier, and database health before dismissing the signal.",
    });
  }

  incidents.sort((left, right) => {
    const severityRank = { critical: 4, high: 3, medium: 2, low: 1 };
    return severityRank[right.severity] - severityRank[left.severity];
  });

  const actions: OverviewAction[] = incidents.slice(0, 4).map((incident, index) => ({
    id: incident.id,
    priority: index === 0 ? "now" : index === 1 ? "next" : "watch",
    title: incident.action,
    detail: `${incident.title} · ${incident.summary}`,
    source: incident.evidence,
  }));

  if (actions.length === 0) {
    actions.push({
      id: "monitor-posture",
      priority: "watch",
      title: "Keep current controls steady and watch for trend changes.",
      detail: "No concentrated risk cluster is visible right now. Continue reviewing recent evidence-backed decisions.",
      source: "Current posture remains stable across the loaded scope.",
    });
  }

  const metrics: OverviewMetric[] = [
    {
      id: "blocked",
      label: "Blocked before deploy",
      value: summary?.total_deny ?? 0,
      tone: "deny",
      trend: formatDelta(denyDelta, "denies"),
      context: workflowMismatch > 0
        ? `Dominant driver: workflow mismatch (${workflowMismatch}).`
        : "Blocked changes stayed explainable through policy and trust checks.",
    },
    {
      id: "verified",
      label: "Verified decisions",
      value: summary?.total_allow ?? 0,
      tone: "allow",
      trend: formatDelta(allowDelta, "allows"),
      context: "Allowed operations remained traceable through the audit store.",
    },
    {
      id: "signals",
      label: "Monitored signals",
      value: monitoredSignals,
      tone: "neutral",
      trend: summary ? `${summary.counts_by_event_type.runtime_drift_result ?? 0} runtime findings recorded` : "No signal baseline yet",
      context: "Admission, artifact verification, and runtime drift stay visible in one surface.",
    },
    {
      id: "runtime-drift",
      label: "Runtime drift (24h)",
      value: driftDenies,
      tone: driftDenies > 0 ? "drift" : "neutral",
      trend: driftStats?.top_drifted_workloads[0]
        ? `Highest pressure: ${topDriftScope(driftStats)}`
        : "No dominant workload yet",
      context: imageDriftReason > 0
        ? `Image drift remains the main runtime class (${imageDriftReason}).`
        : "Review parent-controller reconciliation before pod-level restarts.",
    },
    {
      id: "exceptions",
      label: "Exception queue",
      value: pendingExceptions + activeExceptions,
      tone: pendingExceptions > 0 ? "warning" : "neutral",
      trend: pendingExceptions > 0 ? `${pendingExceptions} pending for review` : `${activeExceptions} active in current scope`,
      context: "Break-glass and exception hygiene should stay visible, bounded, and explainable.",
    },
  ];

  const denyDrivers = (summary?.top_deny_reasons ?? []).slice(0, 4).map((item) => ({
    label: humanizeReason(item.reason),
    count: item.count,
  }));

  const blastRadius = [
    ...(topViolators?.items ?? []).slice(0, 3).map((item) => ({
      label: item.key,
      detail: item.top_reasons[0] ? humanizeReason(item.top_reasons[0].reason) : "deny-heavy scope",
      count: item.deny_count,
    })),
    ...(driftStats?.top_drifted_workloads ?? []).slice(0, 2).map((workload) => ({
      label: workload.namespace ? `${workload.namespace}/${workload.workload}` : workload.workload,
      detail: "runtime drift workload",
      count: workload.count,
    })),
  ].slice(0, 5);

  if (syncStatus && syncStatus.health === "disabled" && syncStatus.summary) {
    actions.push({
      id: "sync-disabled",
      priority: "watch",
      title: "Cross-cluster sync is not configured on this backend.",
      detail: "This is informational for the current deployment, not a live incident by itself.",
      source: syncStatus.summary,
    });
  }

  const trustMetrics: OverviewTrustMetric[] = [
    {
      id: "artifact-integrity",
      label: "Artifact integrity",
      weight: 30,
      score: clampScore(100 - (signatureFailures * 18) - (digestPinningGaps * 12)),
      status: "unknown",
      detail:
        signatureFailures + digestPinningGaps > 0
          ? `${signatureFailures} signature failures and ${digestPinningGaps} digest pinning gaps are pulling trust posture down.`
          : "No dominant signature or digest-pinning pressure is visible in the loaded scope.",
      defenseGap: "Artifact integrity gap",
      rootWeakness: "Immutable artifact identity is not being proven consistently enough across the current pipeline.",
      nextMove: "Repair digest pinning and signature verification before retrying deploy flow.",
    },
    {
      id: "workflow-governance",
      label: "Workflow and signer governance",
      weight: 25,
      score: clampScore(100 - (workflowMismatch * 2.2)),
      status: "unknown",
      detail:
        workflowMismatch > 0
          ? `${pluralize(workflowMismatch, "deny")} point to workflow trust drift or signer policy mismatch.`
          : "No concentrated workflow-mismatch signal is visible in the current window.",
      defenseGap: "Signing and workflow governance gap",
      rootWeakness: "Trusted workflow and signer boundaries are drifting away from the intended repo/ref path.",
      nextMove: "Compare the latest workflow ref and signer policy before widening any trust scope.",
    },
    {
      id: "runtime-hardening",
      label: "Runtime hardening",
      weight: 20,
      score: clampScore(100 - (Math.max(imageDriftReason, driftDenies) * 4)),
      status: "unknown",
      detail:
        imageDriftReason > 0 || driftDenies > 0
          ? `${pluralize(Math.max(imageDriftReason, driftDenies), "runtime drift signal")} still need reconciliation.`
          : "Runtime drift pressure looks contained in the current scope.",
      defenseGap: "Runtime exposure gap",
      rootWeakness: "Deploy-time trust is not enough once live workload state drifts away from approved digests.",
      nextMove: "Start with the highest-drift workload and reconcile the parent controller first.",
    },
    {
      id: "exception-hygiene",
      label: "Exception hygiene",
      weight: 10,
      score: clampScore(100 - (pendingExceptions * 12) - (activeExceptions * 6)),
      status: "unknown",
      detail:
        pendingExceptions + activeExceptions > 0
          ? `${pendingExceptions} pending and ${activeExceptions} active exceptions remain in governance scope.`
          : "No visible exception backlog is dragging posture down right now.",
      defenseGap: "Exception governance gap",
      rootWeakness: "Exception paths are carrying live risk and can become standing policy bypass debt.",
      nextMove: pendingExceptions > 0
        ? "Review pending exception scope first and reject or narrow any over-broad request."
        : "Revalidate active exception evidence and revoke anything no longer tied to a live case.",
    },
    {
      id: "control-plane-health",
      label: "Control-plane health",
      weight: 15,
      score: clampScore((health?.status === "ok" ? 100 : 70) - (totalErrors * 15)),
      status: "unknown",
      detail:
        totalErrors > 0
          ? `${pluralize(totalErrors, "error event")} surfaced in the decision path and should stay tied to the same review.`
          : health?.status === "ok"
            ? "Backend health is currently reporting cleanly."
            : "Health telemetry is incomplete or degraded in the current view.",
      defenseGap: "Policy coverage gap",
      rootWeakness: "Control-plane execution or policy coverage is unstable enough to weaken consistent decision quality.",
      nextMove: "Open the latest error path and confirm the failing control surface before retrying affected operations.",
    },
  ].map((metric) => ({
    ...metric,
    status: metricStatus(metric.score),
  }));

  const weightedScore = trustMetrics.reduce((total, metric) => total + (metric.score * metric.weight), 0) / 100;
  const trustScore = clampScore(weightedScore);
  const trustGrade = gradeFromScore(trustScore);
  const draggers = trustMetrics
    .filter((metric) => metric.status === "gap" || metric.status === "unknown")
    .sort((left, right) => left.score - right.score)
    .slice(0, 3)
    .map((metric) => metric.label);

  const trust: OverviewTrustSummary = {
    score: trustScore,
    grade: trustGrade,
    summary:
      trustGrade === "A" || trustGrade === "B"
        ? "Measured trust signals are mostly healthy, with only bounded posture drag visible in the current scope."
        : trustGrade === "C"
          ? "Trust posture is still explainable, but repeated governance or runtime signals are materially pulling the current score down."
          : "Trust posture needs active operator attention. Current drift, governance, or artifact signals are strong enough to drag the score materially down.",
    derivedFrom: "Derived from current posture signals: artifact integrity, workflow governance, runtime hardening, exception hygiene, and control-plane health.",
    metrics: trustMetrics,
    draggers,
  };

  const guidance: OverviewGuidanceItem[] = [];
  if (workflowMismatch > 0) {
    guidance.push({
      id: "workflow-guidance",
      title: "Treat workflow mismatch as the primary trust review path.",
      confidence: confidenceForCount(workflowMismatch),
      summary: `${pluralize(workflowMismatch, "deny")} already point to the same workflow-trust pattern, so this is likely the fastest posture improvement lever.`,
      nextStep: "Compare the current workflow ref and signer policy before widening any exception or trust scope.",
      source: "Derived from top deny reasons and dominant violator scope.",
    });
  }
  if (imageDriftReason > 0 || driftDenies > 0) {
    const count = Math.max(imageDriftReason, driftDenies);
    guidance.push({
      id: "runtime-guidance",
      title: "Reconcile parent controllers before touching pods.",
      confidence: confidenceForCount(count),
      summary: `${pluralize(count, "runtime drift signal")} suggest the safer path is controller-first reconciliation, not blind pod restarts.`,
      nextStep: `Start with ${topDriftScope(driftStats)} and confirm the approved digest still matches the desired state.`,
      source: "Derived from runtime drift counts and top drifted workload.",
    });
  }
  if (signatureFailures + digestPinningGaps > 0) {
    const count = signatureFailures + digestPinningGaps;
    guidance.push({
      id: "artifact-guidance",
      title: "Repair artifact identity before retrying deployments.",
      confidence: confidenceForCount(count),
      summary: `${pluralize(count, "artifact trust signal")} indicate that digest pinning or signature verification is still incomplete in the current pipeline.`,
      nextStep: "Pin the digest and confirm signature and attestation checks line up with the same immutable artifact before re-running.",
      source: "Derived from artifact verification and digest-pinning deny reasons.",
    });
  }
  if (pendingExceptions > 0 || activeExceptions > 0) {
    const count = pendingExceptions + activeExceptions;
    guidance.push({
      id: "exception-guidance",
      title: "Keep exception cleanup coupled to incident review.",
      confidence: confidenceForCount(count),
      summary: `${pluralize(count, "exception")} remain in scope, which means governance debt is now part of operational posture, not a side queue.`,
      nextStep: pendingExceptions > 0
        ? "Review pending exception scope first and reject or narrow any over-broad request."
        : "Revalidate active exception evidence and revoke any entry that no longer maps to a live incident.",
      source: "Derived from current exception queue and recent exception use.",
    });
  }
  if (guidance.length === 0) {
    guidance.push({
      id: "monitor-guidance",
      title: "Current posture does not show one dominant remediation path.",
      confidence: "limited",
      summary: "Signals in the loaded scope look relatively contained, so keep monitoring trend changes and new evidence-backed incidents.",
      nextStep: "Use the investigations feed below if a new cluster of denies or drift signals starts forming.",
      source: "Derived from the absence of concentrated deny, drift, or exception pressure.",
    });
  }

  return {
    posture,
    trust,
    guidance: guidance.slice(0, 3),
    incidents,
    actions,
    metrics,
    denyDrivers,
    blastRadius,
  };
}
