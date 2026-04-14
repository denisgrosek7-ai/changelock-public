package audit

import (
	"context"
	"testing"
	"time"
)

func TestNormalizeExceptionCreateRequestAppliesTTL(t *testing.T) {
	base := time.Date(2026, 4, 14, 10, 0, 0, 0, time.UTC)

	request, err := NormalizeExceptionCreateRequest(ExceptionCreateRequest{
		ExceptionID:   "EX-1",
		ExceptionType: ExceptionTypeBreakGlass,
		Environment:   "prod",
		Reason:        "P0 fix",
		TicketID:      "INC-1",
		ApprovedBy:    "oncall@example.com",
		TTLHours:      2,
	}, func() time.Time { return base })
	if err != nil {
		t.Fatalf("NormalizeExceptionCreateRequest() error = %v", err)
	}
	if request.ExpiresAt == nil || !request.ExpiresAt.Equal(base.Add(2*time.Hour)) {
		t.Fatalf("unexpected expires_at %#v", request.ExpiresAt)
	}
}

func TestPolicyExceptionMatchesScopeAndExpiry(t *testing.T) {
	now := time.Date(2026, 4, 14, 10, 0, 0, 0, time.UTC)
	exception := PolicyException{
		ExceptionID:   "EX-1",
		ExceptionType: ExceptionTypeDigestBypass,
		Environment:   "prod",
		Namespace:     "acme-prod",
		ImageDigest:   "sha256:abc123",
		ExpiresAt:     now.Add(time.Hour),
		Active:        true,
	}

	if matched, reason := exception.Matches(ExceptionValidationRequest{
		ExceptionID: "EX-1",
		Environment: "prod",
		Namespace:   "acme-prod",
		ImageDigest: "sha256:abc123",
	}, now); !matched || reason != "" {
		t.Fatalf("expected match, got matched=%v reason=%q", matched, reason)
	}

	if matched, reason := exception.Matches(ExceptionValidationRequest{
		ExceptionID: "EX-1",
		Environment: "prod",
		Namespace:   "acme-prod",
		ImageDigest: "sha256:def456",
	}, now); matched || reason == "" {
		t.Fatalf("expected mismatch, got matched=%v reason=%q", matched, reason)
	}

	if matched, reason := exception.Matches(ExceptionValidationRequest{
		ExceptionID: "EX-1",
		Environment: "prod",
		Namespace:   "acme-prod",
		ImageDigest: "sha256:abc123",
	}, now.Add(2*time.Hour)); matched || reason != "exception is expired" {
		t.Fatalf("expected expiry mismatch, got matched=%v reason=%q", matched, reason)
	}
}

func TestMemoryStoreExceptionLifecycle(t *testing.T) {
	store := NewMemoryStore()
	base := time.Date(2026, 4, 14, 10, 0, 0, 0, time.UTC)
	store.now = func() time.Time { return base }

	exception, err := store.CreateException(context.Background(), ExceptionCreateRequest{
		ExceptionID:   "EX-2026-001",
		ExceptionType: ExceptionTypeBreakGlass,
		TenantID:      "acme",
		Environment:   "prod",
		Namespace:     "acme-prod",
		Reason:        "P0 production fix",
		TicketID:      "INC-1234",
		ApprovedBy:    "oncall@example.com",
		TTLHours:      1,
	})
	if err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}

	result, err := store.ValidateException(context.Background(), ExceptionValidationRequest{
		ExceptionID: "EX-2026-001",
		TenantID:    "acme",
		Environment: "prod",
		Namespace:   "acme-prod",
	})
	if err != nil {
		t.Fatalf("ValidateException() error = %v", err)
	}
	if !result.Valid || result.Exception == nil || result.Exception.ExceptionID != exception.ExceptionID {
		t.Fatalf("unexpected validation result %#v", result)
	}

	filtered, err := store.ListExceptions(context.Background(), ExceptionFilter{
		Environment: "prod",
		Limit:       10,
	})
	if err != nil {
		t.Fatalf("ListExceptions() error = %v", err)
	}
	if len(filtered) != 1 || filtered[0].ExceptionID != exception.ExceptionID {
		t.Fatalf("unexpected exceptions %#v", filtered)
	}

	revoked, err := store.RevokeException(context.Background(), exception.ExceptionID)
	if err != nil {
		t.Fatalf("RevokeException() error = %v", err)
	}
	if revoked.Active {
		t.Fatalf("expected exception to be inactive, got %#v", revoked)
	}

	afterRevoke, err := store.ValidateException(context.Background(), ExceptionValidationRequest{
		ExceptionID: exception.ExceptionID,
		TenantID:    "acme",
		Environment: "prod",
		Namespace:   "acme-prod",
	})
	if err != nil {
		t.Fatalf("ValidateException() after revoke error = %v", err)
	}
	if afterRevoke.Valid || afterRevoke.Reason != "exception is inactive" {
		t.Fatalf("unexpected post-revoke validation %#v", afterRevoke)
	}
}

func TestMemoryStoreExceptionReport(t *testing.T) {
	store := NewMemoryStore()
	base := time.Date(2026, 4, 14, 10, 0, 0, 0, time.UTC)
	store.now = func() time.Time { return base }

	activeException, err := store.CreateException(context.Background(), ExceptionCreateRequest{
		ExceptionID:   "EX-ACTIVE",
		ExceptionType: ExceptionTypeBreakGlass,
		Environment:   "prod",
		Namespace:     "acme-prod",
		Reason:        "P0 fix",
		TicketID:      "INC-1",
		ApprovedBy:    "oncall@example.com",
		TTLHours:      2,
	})
	if err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}
	if _, err := store.CreateException(context.Background(), ExceptionCreateRequest{
		ExceptionID:   "EX-INACTIVE",
		ExceptionType: ExceptionTypeBreakGlass,
		Environment:   "prod",
		Namespace:     "acme-prod",
		Reason:        "rollback",
		TicketID:      "INC-2",
		ApprovedBy:    "ops@example.com",
		TTLHours:      1,
	}); err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}
	if _, err := store.RevokeException(context.Background(), "EX-INACTIVE"); err != nil {
		t.Fatalf("RevokeException() error = %v", err)
	}
	if _, err := store.Ingest(context.Background(), Event{
		Component:           "deploy-gate",
		EventType:           EventTypeExceptionUsed,
		Decision:            DecisionAllow,
		Environment:         "prod",
		Namespace:           "acme-prod",
		Digest:              "sha256:abc123",
		IsException:         true,
		ExceptionID:         activeException.ExceptionID,
		ExceptionType:       activeException.ExceptionType,
		ExceptionReason:     activeException.Reason,
		ExceptionTicketID:   activeException.TicketID,
		ExceptionApprovedBy: activeException.ApprovedBy,
		ExceptionExpiresAt:  &activeException.ExpiresAt,
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	report, err := store.ExceptionReport(context.Background(), ExceptionFilter{
		Environment: "prod",
		Limit:       10,
	})
	if err != nil {
		t.Fatalf("ExceptionReport() error = %v", err)
	}
	if len(report.Active) != 1 || report.Active[0].ExceptionID != "EX-ACTIVE" {
		t.Fatalf("unexpected active exceptions %#v", report.Active)
	}
	if len(report.RecentInactive) != 1 || report.RecentInactive[0].ExceptionID != "EX-INACTIVE" {
		t.Fatalf("unexpected inactive exceptions %#v", report.RecentInactive)
	}
	if len(report.RecentUsed) != 1 || report.RecentUsed[0].ExceptionID != "EX-ACTIVE" {
		t.Fatalf("unexpected used events %#v", report.RecentUsed)
	}
}
