const test = require("node:test");
const assert = require("node:assert/strict");

const { buildCLIArgs, buildGuidanceArgs, diagnosticText, extractDigestPinnedImage, normalizeDiagnostics, resolveWorkspacePath } = require("../extension");

test("extractDigestPinnedImage returns the first digest-pinned image", () => {
  const text = `
apiVersion: apps/v1
kind: Deployment
spec:
  template:
    spec:
      containers:
        - name: api
          image: ghcr.io/example/api:latest
        - name: worker
          image: ghcr.io/example/worker@sha256:abcd1234
`;
  assert.equal(extractDigestPinnedImage(text), "ghcr.io/example/worker@sha256:abcd1234");
});

test("buildCLIArgs reuses changelock-cli preflight flags", () => {
  const args = buildCLIArgs({
    files: ["/repo/deploy/app.yaml"],
    policyDir: "/repo/deploy/kyverno",
    bundleDir: "/repo/policies",
    scanner: "trivy",
    tenant: "acme",
    repository: "my-org/acme-app",
    apiUrl: "http://127.0.0.1:8094",
    offline: false,
    image: "ghcr.io/example/api@sha256:abcd",
  });
  assert.deepEqual(args, [
    "preflight",
    "--output",
    "json",
    "--file",
    "/repo/deploy/app.yaml",
    "--policy-dir",
    "/repo/deploy/kyverno",
    "--bundle-dir",
    "/repo/policies",
    "--scanner",
    "trivy",
    "--tenant",
    "acme",
    "--repository",
    "my-org/acme-app",
    "--api-url",
    "http://127.0.0.1:8094",
    "--image",
    "ghcr.io/example/api@sha256:abcd",
  ]);
});

test("buildGuidanceArgs reuses changelock-cli guidance command", () => {
  assert.deepEqual(buildGuidanceArgs("/tmp/changelock/result.json"), [
    "guidance",
    "--input",
    "/tmp/changelock/result.json",
    "--format",
    "markdown",
  ]);
});

test("normalizeDiagnostics preserves machine-readable diagnostics only", () => {
  const result = normalizeDiagnostics({
    diagnostics: [{ reason_code: "manifest_policy_violation" }, null, "bad"],
  });
  assert.deepEqual(result, [{ reason_code: "manifest_policy_violation" }]);
});

test("diagnosticText includes fix hints and docs refs", () => {
  const text = diagnosticText({
    summary: "Kyverno reported policy violations",
    fix_hint: "Review the deny output.",
    docs_ref: "docs/developer-preflight-cli.md",
    evaluation_state: "fail",
  });
  assert.match(text, /Kyverno reported policy violations/);
  assert.match(text, /Fix: Review the deny output\./);
  assert.match(text, /Docs: docs\/developer-preflight-cli.md/);
  assert.match(text, /State: fail/);
});

test("resolveWorkspacePath keeps absolute paths intact and resolves relative paths", () => {
  assert.equal(resolveWorkspacePath("/repo", "/tmp/policies"), "/tmp/policies");
  assert.equal(resolveWorkspacePath("/repo", "deploy/kyverno"), "/repo/deploy/kyverno");
});
