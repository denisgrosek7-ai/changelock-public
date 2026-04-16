# Deployment Modes

ChangeLock supports more than one operating posture.

## 1. Demo / evaluation

Use this mode when the goal is:
- product walkthroughs
- buyer evaluation
- internal demos
- early technical review

Characteristics:
- simpler setup
- lower-friction demo assumptions
- not intended as a production security posture

## 2. Standard production single-cluster

Use this mode when a team wants ChangeLock as a real deployment-security layer for one Kubernetes environment.

Characteristics:
- stronger auth and access controls
- production database and operational practices
- real trust and policy enforcement
- operator-owned deployment configuration

## 3. Production hub mode

Use this mode when ChangeLock acts as the primary governance and reporting authority across clusters.

Characteristics:
- central approvals and reporting
- machine-authenticated cluster interactions
- authoritative view of approved exceptions and operational status

## 4. Production spoke mode

Use this mode when a cluster enforces locally while consuming approved authority signals from a central hub.

Characteristics:
- local enforcement
- pull-based sync of approved exception state
- cluster identity and scoped trust
- deterministic behavior during hub outages

## 5. Higher-trust evidence mode

More advanced deployments may enable stronger evidence and verification behavior.

Characteristics:
- trust-sensitive evidence correlation
- explicit verification states
- stronger forensic value for approvals, sync snapshots, and deployment trust decisions

## Guiding principle

ChangeLock is designed so that demo convenience does not need to become a production default.
