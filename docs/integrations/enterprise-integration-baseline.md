# Enterprise Integration Baseline

This document describes the bounded enterprise integration spine introduced in Wave 2B.

The goal is not to claim full outbound connector automation.
The goal is to expose stable, evidence-backed integration surfaces for identity clarity, ticket preparation, SIEM export, and evidence export.

## Scope

Wave 2B adds:

- identity and SSO introspection
- Jira and ServiceNow ticket-draft preparation
- normalized SIEM export
- bounded evidence export for incidents, recommendations, handoff packages, and validation runs

It does not add:

- direct remote ticket creation
- unbounded generic export jobs
- custom per-customer connector logic
- a new integration truth store

## Routes

- `GET /v1/integrations/identity`
- `GET /v1/integrations/tickets/catalog`
- `POST /v1/integrations/tickets/prepare`
- `GET /v1/integrations/siem/export`
- `POST /v1/integrations/evidence/export`

## Identity And SSO Clarity

`GET /v1/integrations/identity` returns:

- current authenticated actor
- configured auth mode
- identity provider class
- role and tenant claim semantics when OIDC is enabled
- approval and audit attribution semantics

This route explains the current auth model.
It does not prove live upstream IdP health.

## Ticket Baseline

`GET /v1/integrations/tickets/catalog` exposes the stable local contract for supported ticket systems.

`POST /v1/integrations/tickets/prepare` produces an outbound draft with:

- incident linkage
- recommendation linkage
- evidence refs
- validation refs
- handoff refs
- deeplinks back into ChangeLock
- approval-required signal when remediation still needs human review

The output is intentionally draft-only.
Remote Jira or ServiceNow writes remain outside this baseline.

## SIEM Export Baseline

`GET /v1/integrations/siem/export` emits a normalized event stream with:

- event type
- source component
- severity
- decision
- correlation id
- subject ref
- incident ref
- recommendation ref
- evidence refs

This is a stable export schema over canonical audit truth.
Downstream indexing, retention, and alert routing remain the responsibility of the receiving SIEM.

## Evidence Export Baseline

`POST /v1/integrations/evidence/export` exposes bounded export scopes for:

- incidents
- recommendations
- sealed handoff packages
- validation runs

The export distinguishes:

- sealed vs unsealed items
- canonical vs advisory surfaces
- stable API references vs downstream copied artifacts

This route is evidence-aware, but it does not itself freeze or re-seal the referenced objects.
