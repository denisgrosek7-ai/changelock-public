Točka 8 / Val C implements the plugin and extensibility contracts only.

- Točka 8 is not complete in Val C.
- `point_8_pass` is not allowed in Val C.
- Val C depends on the patched Točka 7 / Val E compatibility gate and patched Val B compatibility behavior from `da7719b`.
- Plugins are advisory or developer-assist extensions, not approval, certification, canonical evidence, or production authorization authorities.
- Plugin manifest, capability, sandbox, lifecycle, diagnostics, performance, and trust-boundary rules are bounded and fail closed.
- Custom checks cannot override enterprise policy or approve deployment.
- Plugin diagnostics must preserve uncertainty, stale or partial states, production-only unknowns, and failure reasons.
- Sample plugin descriptors are examples only, not production-ready certified plugins.
- Actual plugin runtime, marketplace publishing, external plugin registry, production SDK runtime, and Točka 9 remain out of scope.
