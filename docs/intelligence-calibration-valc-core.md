# Intelligence Calibration Val C Core

`Točka 5 / Val C` adds the bounded `Feedback, Suppression & Federated Tuning` layer on top of active `Val 0`, `Val A`, and `Val B` intelligence calibration proofs.

Val C is fail-closed on:
- `Val 0` calibration discipline foundation
- `Val A` reachability and VEX calibration
- `Val B` behavioral baseline and learning mode

Val C covers:
- structured operator feedback intake
- feedback review cockpit
- tuning proposal generation
- suppression safety application as candidate-only discipline
- suppression rollback / undo path
- local calibration change review
- federated signal weighting
- environment similarity gating
- local override discipline
- bounded propagation policy
- feedback / federated explanation payloads

Val C remains bounded:
- feedback does not mutate intelligence, suppress signals, or lower priority directly
- suppression stays candidate-only and cannot delete evidence, hide false-negative paths, or activate critical suppression
- federated signals remain advisory only, cannot override local evidence, and are gated by source quality plus environment similarity
- propagation remains disabled or advisory-only by default, requires redaction plus review, and cannot mutate remote calibration
- all outputs remain `projection_only` and `not_canonical_truth`

Val C does **not** complete `Točka 5`.

What remains for later waves:
- `Val D` defensive simulation harness and adversarial calibration validation
- `Val E` final calibration gate and integrated closure
