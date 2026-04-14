package identity

import "testing"

func TestCanonicalFileSetHashDeterministic(t *testing.T) {
	filesA := map[string][]byte{
		"global/artifact-policy.yaml": []byte("metadata:\n  name: artifact\n"),
		"tenants/acme/tenant.yaml":    []byte("metadata:\n  name: acme\n"),
	}
	filesB := map[string][]byte{
		"tenants/acme/tenant.yaml":    []byte("metadata:\n  name: acme\n"),
		"global/artifact-policy.yaml": []byte("metadata:\n  name: artifact\n"),
	}

	hashA := CanonicalFileSetHash(filesA)
	hashB := CanonicalFileSetHash(filesB)
	if hashA == "" {
		t.Fatal("expected non-empty hash")
	}
	if hashA != hashB {
		t.Fatalf("expected deterministic hash, got %q and %q", hashA, hashB)
	}
}

func TestDecisionHashChangesWhenInputChanges(t *testing.T) {
	base := DecisionInput{
		PolicyBundleHash: "sha256:bundle",
		ImageDigest:      "sha256:image",
		RequestID:        "req-1",
		Decision:         "DENY",
		Component:        "deploy-gate",
		Repo:             "my-org/acme-app",
		Environment:      "prod",
	}

	hashA := DecisionHash(base)
	hashB := DecisionHash(base)
	if hashA != hashB {
		t.Fatalf("expected deterministic decision hash, got %q and %q", hashA, hashB)
	}

	base.Decision = "ALLOW"
	hashC := DecisionHash(base)
	if hashA == hashC {
		t.Fatalf("expected decision hash to change when input changes, got %q", hashA)
	}
}
