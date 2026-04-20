# Phase 8 Extended Surface: 8m-8w

## Purpose

This document preserves the historical extended `8m-8w` surface as review lineage.

It exists because the current repo grew through an extended late phase-8 tranche before the phase-9 overlays were named and split more explicitly.
Some of that historical surface now overlaps with current incident, readback, and intelligence routes.

Use this document as lineage guidance, not as a stronger truth source than code, routes, tests, or the main phase-9 surfaces.

## Current scope represented by this lineage document

The extended `8m-8w` surface maps to the parts of the repo that expanded from trust operations into:

- incident-centered operator views
- defense-gap style advisory assessment
- policy replay style operator explanation
- systemic weakness aggregation
- executive defense reporting
- readback and permalink-oriented evidence access

In the current repo, these surfaces live alongside phase-9 overlays and share implementation plumbing in `audit-writer`.

## Representative code surface

- [services/audit-writer/incidents.go](../../services/audit-writer/incidents.go)
- [services/audit-writer/readback.go](../../services/audit-writer/readback.go)
- [services/audit-writer/main.go](../../services/audit-writer/main.go)
- [services/audit-writer/recommendations.go](../../services/audit-writer/recommendations.go)

## Representative routes

- `GET /v1/incidents`
- `GET /v1/incidents/{id}/defense-gaps`
- `GET /v1/incidents/{id}/policy-replay`
- `GET /v1/ai/defense-gap-assessments`
- `GET /v1/ai/policy-replay`
- `GET /v1/ai/systemic-weaknesses`
- `GET /v1/ai/executive-defense-report`
- `POST /v1/readback/grants`
- `GET /v1/readback/defense-gap/{id}`
- `GET /v1/readback/policy-replay/{id}`
- `GET /v1/readback/systemic-weakness/{id}`

## Representative tests

- [services/audit-writer/main_test.go](../../services/audit-writer/main_test.go)
- [services/audit-writer/readback_test.go](../../services/audit-writer/readback_test.go)

## What this lineage document does not include

This document intentionally does not collapse later named phase-9 overlays back into phase 8.
It is not the canonical place to review:

- topology and blast-radius semantics
- time-travel forensics
- sealed handoff
- federation
- runtime integrity
- validation harnesses
- runtime hardening

Those are later layers and should be reviewed in their own current surfaces.
