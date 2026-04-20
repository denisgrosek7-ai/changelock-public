# Documentation Truth Policy

ChangeLock documentation is explanatory, not authoritative over implementation.

When there is any ambiguity about what is really implemented, use this precedence order:

1. code
2. routes and handlers
3. tests
4. merged review history
5. docs

## What each layer means

### 1. Code

Code is the primary implementation truth:

- handlers
- service logic
- control-flow decisions
- payload structures
- deterministic serialization and hashing logic

### 2. Routes and handlers

If the codebase exposes a route and the handler is wired, that route surface is stronger evidence than narrative docs alone.

### 3. Tests

Tests are the next strongest source for:

- intended semantics
- determinism expectations
- failure handling
- compatibility and regression behavior

### 4. Merged review history

Merged review history can clarify:

- why a surface exists
- what was considered in scope
- what guardrails were accepted

It does not override code or tests, but it is stronger than stale prose.

### 5. Docs

Docs explain scope, boundaries, and operator intent.
They should never be treated as proof that a capability exists if the code, routes, or tests do not support it.

## Documentation update rule

Whenever a phase boundary, route surface, schema contract, or operator support boundary changes:

- update the relevant docs in the same change when feasible
- if not feasible in the same change, record the documentation gap explicitly
- never silently use docs to expand product claims beyond code

## Review rule

For strict reviews:

- start from code
- confirm route wiring
- confirm tests
- use docs only as supporting explanation

## Why this policy exists

This prevents three common failures:

- docs drift that claims more than the code supports
- stale phase labels being treated as implementation truth
- review arguments based on prose instead of handlers, payloads, and tests
