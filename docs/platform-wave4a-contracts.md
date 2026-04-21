# Wave 4A Contracts

`Wave 4A` opens enterprise integration surfaces without introducing a new uncontrolled truth layer.

Implemented surfaces:

- `GET /v1/integrations/identity-fabric`
  - configured auth and role-binding model
  - tenant-to-business-role mapping
  - approver classes, break-glass treatment, and actor-lineage rules

- `GET /v1/integrations/itsm-lifecycle`
  - Jira / ServiceNow-ready lifecycle contract
  - incident, remediation, and approval ticket semantics
  - evidence linkage, closure discipline, and reassignment model

- `GET /v1/integrations/itsm-lifecycle/flows`
  - evidence-backed lifecycle projection for a specific incident
  - incident vs remediation vs approval flow distinction
  - closure blockers and operator override refs

- `GET /v1/integrations/siem-sync`
  - bounded outbound event export contract
  - inbound trust classes, severity normalization, action-mapping matrix, and policy-gated evaluate path

- `POST /v1/integrations/siem-sync/evaluate`
  - explainable external-signal evaluation
  - correlation-id preservation
  - trust label, workflow mapping, and safety-limit metadata
  - recommendation-only mapping without runtime execution

- `GET /v1/incidents/collaboration`
  - shared incident context
  - linked evidence, readback, handoff, and validation refs where present
  - remediation progress, closure blockers, and post-remediation verification state
  - audience-aware export variants and disclosure discipline

- `GET /v1/integrations/safety`
  - trusted vs advisory integration boundaries
  - bounded write permissions
  - degraded-mode and replay-safe behavior per connector

- `GET /v1/integrations/safety/health`
  - connector health visibility contract
  - replay-safe and bounded-write posture per connector
  - degraded behavior summary for operational review

Guardrails:

- identity integration describes configured bindings and actor lineage, but does not claim upstream IdP health
- ITSM lifecycle remains `draft_before_write_only` in this slice and does not claim live remote ticket mutation
- inbound SIEM/SOAR signals never execute runtime actions directly; they map only into bounded recommendation or review paths
- incident collaboration is derived from canonical incident and recommendation surfaces, not from a new mutable collaboration store
- integration safety keeps connectors from becoming a shadow control plane
