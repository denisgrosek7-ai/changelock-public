# Reference Architecture Hardening Val A Core

Točka 6 / Val A implements the core blueprint family profile layer on top of the confirmed Val 0 blueprint discipline foundation.

These are validated reference blueprint profiles. They are bounded by assumptions, required capabilities, evidence requirements, degraded conditions, unsupported conditions, support boundaries, and advisory projection rules. They are not certified architectures.

Val A implements these core families:

- `enterprise_default`
- `high_assurance`
- `regulated_privacy_first`
- `sovereign_air_gapped`
- `performance_sensitive`
- `partner_msp_suitable`

Each family profile stays bounded:

- by the Val 0 contract and conformance evaluator
- by explicit environment fit and support boundaries
- by measured evidence requirements and caveats
- by degraded and unsupported conditions that remain visible
- by advisory `projection_only not_canonical_truth` language

Val A does not implement:

- blueprint-as-code delivery packs
- Terraform or Helm
- validation harness execution
- resilience or scaling scenario execution
- dashboard UI
- final reference architecture gate
- integrated closure
- `point_6_pass`

Additional family-specific guardrails in Val A:

- `high_assurance` is stricter than `enterprise_default`, but it does not require all workloads to run in enclaves
- `regulated_privacy_first` makes residency, redaction, export, and evidence custody assumptions explicit without claiming legal certification
- `sovereign_air_gapped` makes offline transfer and local trust boundaries explicit without implementing offline tooling
- `performance_sensitive` makes performance and audit write-path assumptions explicit without promising performance guarantees
- `partner_msp_suitable` preserves customer authority and blocks partner shadow truth or canonical authority claims

Točka 6 remains `not_complete` in Val A. Later waves remain required:

- Val B: blueprint-as-code and validation
- Val C: resilience and scaling hardening
- Val D: operational visibility and final reference architecture gate
- Val E: integrated reference architecture closure
