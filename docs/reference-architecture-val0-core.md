# Reference Architecture Hardening Val 0 Core

Točka 6 / Val 0 establishes the fail-closed foundation for Reference Architecture Hardening.

Reference architecture in this program means a validated reference blueprint with measured conformance. It is not certification, not approval authority, and not canonical truth.

Val 0 implements:

- blueprint contract discipline
- blueprint family taxonomy
- matched, partially_matched, degraded, unsupported, drifted, superseded_reference, and unknown conformance semantics
- environment fit rules across topology, trust anchors, audit path, connectivity, residency, operator model, and verifier or partner access
- conformance evidence requirements with RFC3339 timestamp parsing and freshness checks
- compatibility and deprecation baseline handling
- a narrow Val 0 proofs surface that keeps `point_6_state` not complete

Key guardrails:

- blueprint alignment is bounded by explicit assumptions, capabilities, evidence scope, freshness, and caveats
- degraded and unsupported conditions remain explicit and must not silently become matched
- deprecated or superseded references do not pass as clean matched references
- redaction or caveat omission must not convert degraded or unsupported alignment into matched
- all reference architecture outputs remain projection-only over the canonical evidence spine

Supported blueprint families in Val 0:

- `enterprise_default`
- `high_assurance`
- `regulated_privacy_first`
- `sovereign_air_gapped`
- `performance_sensitive`
- `partner_msp_suitable`

Val 0 does not implement:

- real blueprint family contents beyond taxonomy and contract placeholders
- blueprint-as-code delivery packs
- Terraform, Helm, or deployment recipes
- resilience and scaling hardening packs
- dashboard UI or deviation alerting
- final reference architecture gate
- integrated closure
- `point_6_pass`

Later waves remain required:

- Val A: concrete blueprint families
- Val B: blueprint-as-code and conformance kit
- Val C: resilience and scaling hardening
- Val D: operational visibility and final reference architecture gate
- Val E: integrated reference architecture closure
