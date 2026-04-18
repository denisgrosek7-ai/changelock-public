package main

import (
	"context"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	runtimestate "github.com/denisgrosek/changelock/internal/runtime"
	"github.com/denisgrosek/changelock/internal/signingidentity"
)

type fakeControlPlane struct {
	desired    []audit.RuntimeDesiredStateView
	active     []audit.RuntimeActiveStateView
	net        audit.VulnerabilityNetResponse
	netErr     error
	identities []signingidentity.Observation
}

func (f fakeControlPlane) enabled() bool { return true }
func (f fakeControlPlane) desiredStates(context.Context) ([]audit.RuntimeDesiredStateView, error) {
	return f.desired, nil
}
func (f fakeControlPlane) activeStates(context.Context) ([]audit.RuntimeActiveStateView, error) {
	return f.active, nil
}
func (f fakeControlPlane) netVulnerabilities(context.Context, string, string, string, string) (audit.VulnerabilityNetResponse, error) {
	return f.net, f.netErr
}
func (f fakeControlPlane) signingIdentities(context.Context, string, string, string, string) ([]signingidentity.Observation, error) {
	return f.identities, nil
}

func TestRuntimeVEXQuarantineTriggersForNetActionableCritical(t *testing.T) {
	runtime := agentRuntime{
		config: runtimestate.SelfHealingConfig{
			RuntimeVEXQuarantine: true,
			RuntimeVEXSeverity:   "CRITICAL",
			RuntimeVEXRequireNet: true,
		},
		controlPlane: fakeControlPlane{
			net: audit.VulnerabilityNetResponse{
				ActionableCount:   2,
				ThresholdBreached: true,
			},
		},
	}
	outcome := runtime.runtimeVEXQuarantine(context.Background(), runtimestate.ApprovedWorkloadState{
		TenantID:  "acme",
		Namespace: "acme-prod",
	}, runtimestate.ComparisonResult{
		RunningDigest: "sha256:abc",
	}, false, "")
	if outcome == nil || outcome.Status != runtimestate.FindingStatusQuarantined {
		t.Fatalf("expected vex quarantine outcome, got %#v", outcome)
	}
}

func TestBuildActiveStateEventMarksProtectedTarget(t *testing.T) {
	desired := runtimestate.ApprovedWorkloadState{
		TenantID:                      "acme",
		ClusterID:                     "prod-eu",
		Namespace:                     "acme-prod",
		WorkloadKind:                  "Deployment",
		Workload:                      "booking-api",
		ServiceAccountName:            "booking-api",
		DesiredStateVerificationState: runtimestate.VerificationStateVerified,
	}
	observed := runtimestate.ObservedWorkloadState{
		ClusterID:          "prod-eu",
		Namespace:          "acme-prod",
		WorkloadKind:       "Deployment",
		Workload:           "booking-api",
		ServiceAccountName: "booking-api",
	}
	result := runtimestate.ComparisonResult{
		ClusterID:                     "prod-eu",
		Namespace:                     "acme-prod",
		WorkloadKind:                  "Deployment",
		Workload:                      "booking-api",
		Result:                        string(runtimestate.DriftClassImageDigest),
		Severity:                      runtimestate.DriftSeverityHigh,
		DesiredStateVerificationState: runtimestate.VerificationStateVerified,
	}
	outcome := &runtimestate.RemediationOutcome{Status: runtimestate.FindingStatusDetected}
	event := buildActiveStateEvent("scan-1", desired, observed, result, outcome, true, "protected namespace", "")
	if !event.ProtectedTarget || event.ProtectedReason != "protected namespace" {
		t.Fatalf("expected protected target metadata, got %#v", event)
	}
	if event.ReconciliationStatus != string(runtimestate.ReconciliationStatusDrift) {
		t.Fatalf("expected drift_detected status, got %#v", event)
	}
}

func TestRuntimeSignerIdentityQuarantineTriggersForUnauthorizedDigest(t *testing.T) {
	runtime := agentRuntime{
		identityConfig: signingidentity.Config{QuarantineOnDrift: true},
		controlPlane: fakeControlPlane{
			identities: []signingidentity.Observation{
				{
					ImageDigest:  "sha256:abc",
					Authorized:   signingidentity.AuthorizationUnauthorized,
					ReasonCode:   signingidentity.ReasonPolicyMissing,
					ReasonDetail: "no enabled signing identity policy matched the observed signer",
				},
			},
		},
	}
	outcome := runtime.runtimeSignerIdentityQuarantine(context.Background(), runtimestate.ApprovedWorkloadState{
		TenantID:  "acme",
		Namespace: "acme-prod",
	}, runtimestate.ComparisonResult{
		RunningDigest: "sha256:abc",
	}, false, "")
	if outcome == nil || outcome.Status != runtimestate.FindingStatusQuarantined {
		t.Fatalf("expected signer identity quarantine outcome, got %#v", outcome)
	}
}
