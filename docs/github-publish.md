# GitHub Public Publish Flow

ChangeLock now uses a split repository model:

- private repo:
  - full implementation
  - source of truth
- public repo:
  - docs-only buyer-facing material

You should push code only to the private repository.

## Source of truth for public content

Public-facing files now live under:

- `public-docs/`

That directory is the only content exported to the public GitHub repository.

It intentionally contains:

- `README.md`
- `LICENSE`
- public conceptual docs

It intentionally does **not** contain:

- backend code
- UI code
- Helm charts
- deploy manifests
- tests
- scripts other than the private export helper

## Automatic publish workflow

The private repository now contains:

- `.github/workflows/publish-public-docs.yml`
- `scripts/export_public_repo.sh`

Behavior:

1. you push changes to the private repo
2. if the push includes changes under `public-docs/`, the workflow runs on private `main`
3. the workflow exports only `public-docs/` content
4. it pushes that docs-only result to:
   - `denisgrosek7-ai/changelock-public`

This means daily work happens only in the private repo.

## One-time setup required

GitHub Actions in one repository cannot automatically push to another repository without explicit credentials.

Create a fine-grained GitHub token with:

- repository access:
  - `denisgrosek7-ai/changelock-public`
- permission:
  - `Contents: Read and write`

Store it as a private repo secret in `changelock-privat`:

- secret name:
  - `CHANGELOCK_PUBLIC_REPO_TOKEN`

Without that secret, the publish workflow will fail closed and the public repo will not be updated.

## Local dry-run of the export

You can preview exactly what would be published:

```bash
chmod +x scripts/export_public_repo.sh
./scripts/export_public_repo.sh /tmp/changelock-public-preview
```

The preview directory should contain only the docs-only public structure.

## Operational rule

Treat `public-docs/` as the buyer-facing/public surface.

If something should be public:
- edit it under `public-docs/`

If something should remain private:
- keep it outside `public-docs/`
