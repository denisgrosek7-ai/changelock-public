# Measured Public Proof Expansion Val D Core

- GET `/v1/public/proof-expansion/vald/release-issuance-gate`
- GET `/v1/public/proof-expansion/vald/claim-lifecycle`
- GET `/v1/public/proof-expansion/vald/publication-decisions`
- GET `/v1/public/proof-expansion/vald/correction-workflow`
- GET `/v1/public/proof-expansion/vald/proofs`

This bounded `Val D` code slice adds:
- release-bound proof issuance decisions over existing sealed artifacts and verifier outputs
- explicit claim lifecycle visibility for `restricted`, `superseded`, `withdrawn`, and `claim_not_reissued` governance
- publication decisions that stay bounded to declared redaction tier, portal scope, and no-override automation posture
- correction workflow visibility for restriction, withdrawal, and supersession handling

This slice does not claim:
- a new mutable issuance database or shadow truth store
- fully automated external proof publication beyond bounded projection decisions
- final replay, redaction, signing, and compatibility signoff across all Point 2 dimensions

`Val D` remains fail-closed on active `Val C`.
