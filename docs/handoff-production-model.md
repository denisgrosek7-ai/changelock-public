# Handoff Production Model

Wave 2C hardens sealed handoff into a bounded production transport model.

## What is included

- deterministic bundle assembly
- offline verification as a first-class integrity path
- explicit timestamp and transparency edge-state handling
- quality-gate reporting over stored package truth
- archive-integrity and long-term re-attest guidance
- signer-backend contract metadata for future KMS or HSM adapters

## What is not claimed

- permanent external timestamp availability
- permanent transparency availability
- live KMS or HSM integration in every deployment by default

## Production stance

Use `GET /v1/handoff/quality-gates` to review:

- active signer backend
- supported backend model contract
- offline verify support
- deterministic assembly status
- edge-state handling posture
- archive integrity posture

The quality-gate surface is evidence-backed and derived from stored handoff records.
