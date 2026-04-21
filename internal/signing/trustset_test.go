package signing

import (
	"context"
	"testing"
)

func TestTrustSetVerifyWithRetiredVerifyOnlyMember(t *testing.T) {
	oldRuntime, err := NewRuntime(Config{
		Mode:           ModeSoftware,
		Purposes:       map[string]struct{}{PurposeExceptions: {}},
		KeyID:          "old-key",
		Algorithm:      AlgorithmHMACSHA256,
		SoftwareSecret: "old-secret",
	}, ProviderOptions{})
	if err != nil {
		t.Fatalf("NewRuntime(old) error = %v", err)
	}
	newRuntime, err := NewRuntime(Config{
		Mode:           ModeSoftware,
		Purposes:       map[string]struct{}{PurposeExceptions: {}},
		KeyID:          "new-key",
		Algorithm:      AlgorithmHMACSHA256,
		SoftwareSecret: "new-secret",
	}, ProviderOptions{})
	if err != nil {
		t.Fatalf("NewRuntime(new) error = %v", err)
	}

	payload := []byte(`{"exception_id":"EX-1"}`)
	envelope, err := oldRuntime.Sign(context.Background(), PurposeExceptions, payload)
	if err != nil {
		t.Fatalf("oldRuntime.Sign() error = %v", err)
	}

	result, path, err := TrustSet{
		Members: []TrustSetMember{
			{MemberID: "active-new", Runtime: newRuntime, LifecycleState: KeyStateActive},
			{MemberID: "retired-old", Runtime: oldRuntime, LifecycleState: KeyStateRetiredVerifyOnly},
		},
	}.Verify(context.Background(), PurposeExceptions, payload, envelope)
	if err != nil {
		t.Fatalf("TrustSet.Verify() error = %v", err)
	}
	if result.State != StateVerified || path.MemberID != "retired-old" || path.LifecycleState != KeyStateRetiredVerifyOnly {
		t.Fatalf("expected retired verify-only member to preserve historical verification, got result=%#v path=%#v", result, path)
	}
}

func TestTrustSetRejectsRevokedMatchingMember(t *testing.T) {
	runtime, err := NewRuntime(Config{
		Mode:           ModeSoftware,
		Purposes:       map[string]struct{}{PurposeExceptions: {}},
		KeyID:          "revoked-key",
		Algorithm:      AlgorithmHMACSHA256,
		SoftwareSecret: "revoked-secret",
	}, ProviderOptions{})
	if err != nil {
		t.Fatalf("NewRuntime() error = %v", err)
	}
	payload := []byte(`{"exception_id":"EX-2"}`)
	envelope, err := runtime.Sign(context.Background(), PurposeExceptions, payload)
	if err != nil {
		t.Fatalf("runtime.Sign() error = %v", err)
	}

	result, _, err := TrustSet{
		Members: []TrustSetMember{
			{MemberID: "revoked", Runtime: runtime, LifecycleState: KeyStateRevoked},
		},
	}.Verify(context.Background(), PurposeExceptions, payload, envelope)
	if err != nil {
		t.Fatalf("TrustSet.Verify() error = %v", err)
	}
	if result.State != StateFailed {
		t.Fatalf("expected revoked trust-set member to fail verification, got %#v", result)
	}
}
