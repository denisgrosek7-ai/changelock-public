Točka 8 / Val B implements repo and SDK integration contracts only.

- Točka 8 is not complete in Val B.
- `point_8_pass` is not allowed in Val B.
- Val B depends on the patched Točka 7 / Val E compatibility gate.
- Repo-local ChangeLock-as-code config is schema-bound, scope-bound, review-bound, and advisory/local unless governed otherwise.
- Repo-local config cannot override enterprise governance.
- Policy preview is advisory and cannot approve deployment.
- Local-to-CI continuity does not turn local advisory PASS-like output into CI PASS.
- SDK/API surfaces cannot mutate canonical evidence, approve deployment, certify trust, or bypass policy.
- Examples/templates are adoption helpers, not compliance certification or production approval.
- Actual production SDK runtime, repo config runtime/parser, plugin runtime, marketplace publishing, and Točka 9 remain out of scope.
