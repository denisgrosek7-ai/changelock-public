# Runtime / Substrate Depth Expansion Plan

This document is a locked planning document for the next runtime-depth program after the bounded `Phase 8` close.

It is intentionally strict:

- it describes intended scope, controls, and review boundaries
- it does not claim current implementation
- it does not override the documentation truth policy
- it exists to keep kernel-near runtime work reviewable, technically bounded, and resistant to marketing overclaim

Before treating any statement here as implementation status, read [documentation-truth-policy.md](./documentation-truth-policy.md).

## 0. Purpose

This program moves ChangeLock from:

- substrate-aware baseline observability

to:

- a substrate-backed execution trust layer

where the system can correlate:

- process and exec state
- binary and file identity
- file, network, and selected syscall activity
- workload metadata
- node and execution-class context

into one evidence-traceable runtime model.

The goal is:

turn kernel-near execution signal into bounded, explainable, evidence-backed runtime truth without claiming absolute truth, generic memory-safety guarantees, or universal inline prevention authority.

## 1. Foundational Rule

This program must not create a new runtime shadow-truth store.

Substrate-backed signal must remain:

- evidence-backed
- versioned
- bounded by actual hook coverage
- explainable
- degradable when capabilities are missing

Every runtime result must stay explicitly separated between:

- `observed_fact`
- `derived_correlation`
- `enforcement_capable_signal`
- `unsupported_or_unknown`

No kernel-near signal may be described as absolute truth.

## 2. Scope

### In Scope

#### 2.1 Substrate Signal Plane

- tracepoint-based signal capture
- selected kprobe and kretprobe enrichment where justified
- verifier-safe eBPF programs
- security-hook aligned enforcement where technically supportable
- ring-buffer or map-backed event delivery
- userspace correlation over kernel-near signal

#### 2.2 Execution Attribution

- exec attribution
- process lineage
- thread and process correlation
- file access attribution
- socket and connect attribution
- workload, namespace, and node enrichment
- binary path and digest linkage where supportable
- image, signature, and attestation correlation where supportable

#### 2.3 Runtime Trust Scoring

- substrate confidence score
- low, medium, high, and critical drift banding
- stale, partial, and unsupported semantics
- false-positive budget discipline

#### 2.4 Enforcement Taxonomy

- `prevent`
- `observe`
- `contain`
- `terminate`
- `unsupported`

with explicit semantics about which actions truly block, which only contain or terminate, and which remain observational.

#### 2.5 Execution Class Model

- standard node
- hardened node
- confidential-capable node
- VM-backed node
- offline or air-gapped node

#### 2.6 Performance and Proof Discipline

- measured capture latency
- measured correlation latency
- measured enforcement-decision latency
- measured end-to-end event path overhead
- degraded-mode handling when signal or enrichment paths are unavailable

### Out of Scope

- generic memory-safety guarantees for all applications
- blanket buffer-overflow prevention claims
- absolute truth or invisible-defense claims
- universal inline blocking for every runtime attack path
- confidential-computing authority claims beyond declared evidence
- legal, regulatory, or insurer conclusions derived directly from runtime signal

## 3. Mandatory Architecture

### 3.1 Substrate Signal Layer

The kernel-near layer must stay:

- minimal
- verifier-safe
- explicit about hook coverage
- narrowly responsible for capture rather than broad business logic

Required signal families:

- exec lifecycle
- process lifecycle
- file activity
- network activity
- selected syscall or security-hook signal

### 3.2 Identity and Correlation Layer

Userspace correlation must be able to connect:

- PID and TID
- exec lineage
- namespace and workload metadata
- image digest and signature linkage
- attestation or provenance evidence where available
- node class
- policy profile
- prior runtime trust state

Correlation must distinguish:

- supported correlation
- partial correlation
- stale correlation
- unsupported correlation

### 3.3 Enforcement Layer

Enforcement may exist only where ChangeLock has enough:

- hook coverage
- confidence
- rollback discipline
- latency budget
- auditable action semantics

Every action must declare whether it is:

- truly preventive
- terminating after detection
- containing blast radius
- purely observational
- unsupported on the current substrate class

### 3.4 Evidence and Audit Layer

Every substrate-backed record must be mappable to:

- observed kernel event
- correlated runtime fact
- derived risk score
- enforcement decision
- evidence lineage
- confidence class
- freshness or staleness state

## 4. Mandatory Controls

This program must define all of the following before it can pass:

1. substrate support matrix by kernel, distro assumptions, cgroup mode, runtime assumptions, and execution class
2. execution attribution schema for exec, process, file, network, and selected syscall paths
3. binary and provenance correlation model
4. substrate confidence score model
5. false-positive budget model
6. enforcement taxonomy with explicit guarantees and non-guarantees
7. degraded-mode and unsupported-state model
8. measured performance budget model
9. replayable proof and benchmark methodology

## 5. Prohibited Claims

This program is prohibited from:

1. describing substrate signal as absolute truth
2. describing kernel-near capture as invisible defense
3. claiming generic memory-integrity or buffer-overflow prevention
4. claiming universal inline blocking for all runtime threats
5. claiming hardware-rooted truth where only partial correlation exists
6. collapsing `observe`, `contain`, and `terminate` into one fake prevention claim
7. hiding unsupported or degraded states
8. publishing latency claims without measured methodology
9. creating a second runtime truth base outside the canonical evidence spine

## 6. Delivery Waves

### Val A - Substrate Observability Baseline

Must add:

- exec lifecycle tracing
- process lineage model
- basic file and network attribution
- workload and node enrichment
- substrate event schema
- stale, partial, and unsupported status model

`Val A` pass requires:

- every exec event gets a stable process identity record
- file and network events can be attributed to process and workload where supportable
- unsupported or degraded paths are explicit
- no shadow truth base is introduced

### Val B - Binary and Provenance Correlation

Must add:

- binary path and digest capture
- image and provenance correlation
- signature and attestation linkage
- process-to-image consistency checks
- drift semantics:
  - `expected`
  - `low_risk_drift`
  - `suspicious_drift`
  - `hard_mismatch`

`Val B` pass requires:

- ChangeLock can explain why a process is expected, drifted, or mismatched
- correlation is evidence-traceable
- unsupported correlation cases remain explicit

### Val C - Enforcement Taxonomy Baseline

Must add:

- observe-only policies
- sample or escalate mode
- selected deny or override paths where safe
- terminate or contain semantics
- approval and rollback discipline
- policy-to-hook mapping

`Val C` pass requires:

- each enforcement action declares what it truly guarantees
- each action has an explicit audit trail
- no action claims more than the underlying hook semantics support

### Val D - Execution Class Matrix

Must add:

- support matrix for all declared execution classes
- degraded and unsupported capability mapping by class
- signal coverage tracking by class
- overhead visibility by class

`Val D` pass requires:

- every execution class has a visible support matrix
- unsupported and degraded capabilities are explicit
- overhead and signal coverage are measured per class

### Val E - Performance and Proof Pack

Must add:

- p50, p95, and p99 measurement discipline
- capture latency metrics
- correlation latency metrics
- enforcement-decision latency metrics
- false-positive measurements
- replayable benchmark pack

`Val E` pass requires:

- no unmeasured latency claims remain
- benchmark methodology is replayable
- performance becomes a gate, not a narrative

## 7. Deliverables

### Technical Deliverables

- substrate event schema
- exec, process, file, and network correlation layer
- binary and provenance correlation layer
- substrate confidence score model
- enforcement taxonomy engine
- execution class support matrix
- degraded-mode handling model
- performance benchmark suite

### Operational Deliverables

- substrate support matrix document
- enforcement semantics guide
- false-positive budget guide
- degraded and unsupported behavior guide
- rollout and rollback guide

### Proof Deliverables

- exec attribution examples
- process-to-image consistency examples
- drift classification examples
- enforcement decision examples
- performance pack
- degraded-mode examples

## 8. Exit Criteria

This program is complete only if all of the following hold:

- ChangeLock can perform substrate-level exec, file, and network attribution where supportable
- binary and provenance correlation exists and is evidence-traceable
- enforcement taxonomy is explicit and audited
- execution class matrix exists
- unsupported and degraded states are formally defined
- performance budgets are defined and measured
- false-positive budget is defined and monitored
- no claim exceeds actual hook coverage or enforcement semantics

## 9. Pass and Fail Conditions

### Pass

This program is a pass when ChangeLock:

- has substrate-backed execution attribution
- has evidence-traceable process and provenance correlation
- can distinguish observe, contain, terminate, prevent, and unsupported paths cleanly
- has measured runtime budgets instead of performance marketing
- stays bounded by actual hook coverage and explainable evidence

### Fail

This program is a fail if any of the following occur:

- the system claims absolute runtime truth
- the system claims generic memory-safety it cannot provide
- unsupported or degraded states are hidden
- enforcement semantics are overstated
- performance claims are unmeasured
- substrate signal creates a second runtime truth base
- provenance correlation is heuristic-only and not evidence-traceable
- risk scoring or drift bands cannot be explained

## 10. Summary

The purpose of this program is not to turn ChangeLock into a universal kernel security oracle.

Its purpose is to make ChangeLock:

1. deeper at runtime execution attribution
2. more evidence-backed at substrate depth
3. clearer about what can be observed, correlated, enforced, or not supported
4. measurable in latency, overhead, and false-positive behavior

The correct outcome is:

- stronger technical moat
- stronger runtime evidence
- stronger explainability

without drifting into:

- absolute-truth rhetoric
- invisible-defense rhetoric
- generic memory-safety claims
- unsupported enforcement promises
