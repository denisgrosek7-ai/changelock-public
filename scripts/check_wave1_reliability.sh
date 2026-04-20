#!/bin/bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

export GOCACHE="${GOCACHE:-/tmp/changelock-gocache}"

go test ./services/deploy-gate -run 'TestAdmissionReview(AllowsTrustedWorkload|DecisionIsIdempotentForSameInput|DeniesWhenArtifactVerifierErrors|DeniesWhenVEXAwareLookupFails|DeniesUnauthorizedSignerWhenEnforced)'
go test ./services/audit-writer -run 'Test(HandoffSealDownloadAndVerify|HandoffCosignKeepsManifestHashBoundToSamePackage|HandoffVerifyReportsTimestampTransparencyAndArtifactFailures|ForensicsStateDeltaTimelineReplayAndFlashback|ForensicsStateReturnsLimitationsForEmptyScope|RuntimeIntegrityStateFindingsAndEnforcement|RuntimeIntegrityMarksTelemetryGapAsUnverifiable|StrictValidationRegressionChaosAndCompatibilityRuns|ReadyHandlerReturnsServiceUnavailableWhenStorePingFails)'
