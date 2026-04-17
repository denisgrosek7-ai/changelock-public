package signingidentity

import (
	"testing"
	"time"
)

func TestParseConfigRejectsInvalidEnforcementMode(t *testing.T) {
	_, err := ParseConfig(func(key string) string {
		if key == "CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT" {
			return "broken"
		}
		return ""
	})
	if err == nil {
		t.Fatal("expected config error")
	}
}

func TestEvaluateAuthorizesMatchingGitHubOIDCPolicy(t *testing.T) {
	cfg := Config{Enforcement: EnforcementEnforce}
	policy, err := NewPolicy(CreatePolicyRequest{
		ProviderType:   ProviderGitHubOIDC,
		Issuer:         "https://token.actions.githubusercontent.com",
		SignerIdentity: "https://github.com/my-org/acme-app/.github/workflows/build.yml@refs/heads/main",
		Subject:        "repo:my-org/acme-app",
		Repository:     "my-org/acme-app",
		Workflow:       ".github/workflows/build.yml",
		Ref:            "refs/heads/main",
	}, "admin", time.Now().UTC())
	if err != nil {
		t.Fatalf("NewPolicy() error = %v", err)
	}

	decision := Evaluate(cfg, []Policy{policy}, DecisionInput{
		Issuer:            "https://token.actions.githubusercontent.com",
		SignerIdentity:    "https://github.com/my-org/acme-app/.github/workflows/build.yml@refs/heads/main",
		Subject:           "repo:my-org/acme-app",
		Repository:        "my-org/acme-app",
		Workflow:          ".github/workflows/build.yml",
		Ref:               "refs/heads/main",
		TransparencyState: "verified",
	})
	if decision.Authorized != AuthorizationAuthorized || decision.ReasonCode != ReasonAuthorized || decision.Deny {
		t.Fatalf("unexpected decision %#v", decision)
	}
}

func TestEvaluateRequiresTransparencyWhenConfigured(t *testing.T) {
	cfg := Config{
		Enforcement:  EnforcementEnforce,
		RequireRekor: true,
	}
	decision := Evaluate(cfg, nil, DecisionInput{
		Issuer:         "https://token.actions.githubusercontent.com",
		SignerIdentity: "https://github.com/my-org/acme-app/.github/workflows/build.yml@refs/heads/main",
	})
	if decision.ReasonCode != ReasonTransparencyUnverified || !decision.Deny {
		t.Fatalf("expected transparency enforcement failure, got %#v", decision)
	}
}

func TestEvaluateDistrustedAfterCutoff(t *testing.T) {
	now := time.Now().UTC()
	policy, err := NewPolicy(CreatePolicyRequest{
		ProviderType:   ProviderGitHubOIDC,
		Issuer:         "https://token.actions.githubusercontent.com",
		SignerIdentity: "https://github.com/my-org/acme-app/.github/workflows/build.yml@refs/heads/main",
		Repository:     "my-org/acme-app",
	}, "admin", now)
	if err != nil {
		t.Fatalf("NewPolicy() error = %v", err)
	}
	cutoff := now.Add(-1 * time.Hour)
	policy.DistrustedAfter = &cutoff

	decision := Evaluate(Config{Enforcement: EnforcementEnforce}, []Policy{policy}, DecisionInput{
		Issuer:         policy.Issuer,
		SignerIdentity: policy.SignerIdentity,
		Repository:     policy.Repository,
		EvidenceAt:     timePointer(now),
	})
	if decision.ReasonCode != ReasonDistrustedAfterCutoff || !decision.Deny {
		t.Fatalf("expected distrust cutoff denial, got %#v", decision)
	}
}

func TestBuildWorkflowFindingsDetectsSigningCapableWorkflowWithoutPolicy(t *testing.T) {
	findings := BuildWorkflowFindings([]WorkflowDocument{
		{Path: ".github/workflows/build.yml", SigningCapable: true},
	}, nil, time.Now().UTC())
	if len(findings) != 1 || findings[0].Type != FindingWorkflowTokenNoPolicy {
		t.Fatalf("expected signing-capable workflow finding, got %#v", findings)
	}
}
