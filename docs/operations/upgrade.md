# Upgrade

ChangeLock upgrades should preserve the shared PostgreSQL state and roll the stateless control-plane components safely.

Related docs:

- [Release Channels](release-channels.md)
- [Rollback Guide](rollback.md)
- [Compatibility Matrix](compatibility-matrix.md)
- [Support Boundaries](support.md)

## Recommended flow

1. Take a PostgreSQL backup before the upgrade.
2. Render the target manifests locally:
```bash
helm template changelock ./charts/changelock -n changelock-system -f ./charts/changelock/values-enterprise-example.yaml >/tmp/changelock-rendered.yaml
```
3. Review changes to:
- image tags
- auth secret references
- webhook TLS / CA bundle values
- replica counts and PDBs
- new feature env flags such as vuln-ops
4. Apply the upgrade:
```bash
helm upgrade --install changelock ./charts/changelock -n changelock-system -f ./charts/changelock/values-enterprise-example.yaml
```
5. Verify:
```bash
kubectl rollout status deployment/changelock-changelock-audit-writer
kubectl rollout status deployment/changelock-changelock-policy-engine
kubectl rollout status deployment/changelock-changelock-deploy-gate
kubectl get validatingwebhookconfiguration
```
6. Run the post-upgrade smoke checks in [go-live-checklist.md](go-live-checklist.md).

## Recovery from self-blocking admission

If a bad deploy-gate rollout blocks pod changes in an enforced namespace:

1. remove `changelock.io/enforce=enabled` from the affected namespace, or
2. upgrade the release with `deployGate.webhook.enabled=false`, then
3. restore the deploy-gate service and re-enable the label only after `/health` is green again

If the upgrade must be reversed instead of only unblocked, continue with the formal [rollback guide](rollback.md).
