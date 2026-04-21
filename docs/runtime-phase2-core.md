# Phase 2 Runtime Core Contracts

This bounded slice opens `Faza 2` by adding the first real execution-trust core surfaces:
- runtime substrate truth
- remote attestation verification
- trusted execution profile contracts
- response simulation and rollback drill proofs

## Added surfaces
- `GET /v1/runtime/substrate-truth`
- `POST /v1/runtime/substrate-truth`
- `GET /v1/runtime/trusted-execution-profiles`
- `POST /v1/runtime/attestation/verify`
- `GET /v1/runtime/attestation/verifications`
- `POST /v1/runtime/response/simulate`
- `POST /v1/runtime/response/rollback-drill`
- `GET /v1/runtime/phase2/proofs`

## What is now real in code
- runtime substrate truth is modeled as workload, process, node, and attestation binding in one evidence-backed record
- default trusted execution profiles exist for confidential, hardened, and crypto-hardening enforcement paths
- remote attestation verification is provider-aware and can gate credential release
- deploy-gate can deny a workload when a trusted execution profile mismatches the provided substrate or attestation evidence
- response simulation and rollback drill surfaces reuse the bounded runtime enforcement engine instead of inventing a second response authority
- Phase 2 proofs expose bounded artifacts for substrate truth, attestation verification, response simulation, and rollback drill evidence

## Boundaries
- this slice does not claim full kernel telemetry ownership, universal confidential provider support, or complete GitOps control-plane authority
- all Phase 2 evidence remains anchored in canonical audit events and does not introduce a new runtime shadow-truth store
