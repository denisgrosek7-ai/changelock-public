import type { AuditHealth } from "../types";

type Props = {
  health: AuditHealth | null;
};

export function HealthBadge({ health }: Props) {
  const status = health?.status || "unknown";
  const tone = status === "ok" ? "allow" : "error";

  return <span className={`badge badge--${tone}`}>{status}</span>;
}
