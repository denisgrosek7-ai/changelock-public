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
	if afterRevoke.Valid || afterRevoke.Reason != "exception is revoked" {
		t.Fatalf("unexpected post-revoke validation %#v", afterRevoke)
	}
}

func TestMemoryStoreCVEWhitelistCreateAndValidate(t *testing.T) {
	store := NewMemoryStore()
	base := time.Date(2026, 4, 14, 10, 0, 0, 0, time.UTC)
	store.now = func() time.Time { return base }

	exception, err := store.CreateException(context.Background(), ExceptionCreateRequest{
		ExceptionID:   "EX-CVE-001",
		ExceptionType: ExceptionTypeCVEWhitelist,
		TenantID:      "acme",
		Environment:   "prod",
		Repo:          "my-org/acme-app",
		CVEID:         "cve-2026-1234",
		Reason:        "temporary CVE waiver",
		TicketID:      "SEC-1234",
		ApprovedBy:    "security@example.com",
		TTLHours:      2,
	})
	if err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}
	if exception.CVEID != "CVE-2026-1234" {
		t.Fatalf("expected normalized CVEID, got %#v", exception)
	}

	result, err := store.ValidateException(context.Background(), ExceptionValidationRequest{
		ExceptionID: "EX-CVE-001",
		TenantID:    "acme",
		Environment: "prod",
		Repo:        "my-org/acme-app",
		CVEID:       "cve-2026-1234",
	})
	if err != nil {
		t.Fatalf("ValidateException() error = %v", err)
	}
	if !result.Valid || result.Exception == nil || result.Exception.ExceptionID != "EX-CVE-001" {
		t.Fatalf("unexpected validation result %#v", result)
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

func TestMemoryStoreExceptionReportFiltersRecentUsedByCVEIDAndImageDigest(t *testing.T) {
	store := NewMemoryStore()
	base := time.Date(2026, 4, 14, 10, 0, 0, 0, time.UTC)
	store.now = func() time.Time { return base }

	cveException, err := store.CreateException(context.Background(), ExceptionCreateRequest{
		ExceptionID:   "EX-CVE",
		ExceptionType: ExceptionTypeCVEWhitelist,
		TenantID:      "acme",
		Environment:   "prod",
		Namespace:     "acme-prod",
		Repo:          "my-org/acme-app",
		CVEID:         "CVE-2026-0001",
		Reason:        "temporary waiver",
		TicketID:      "SEC-1",
		ApprovedBy:    "security@example.com",
		TTLHours:      2,
	})
	if err != nil {
		t.Fatalf("CreateException() cve error = %v", err)
	}
	digestException, err := store.CreateException(context.Background(), ExceptionCreateRequest{
		ExceptionID:   "EX-DIGEST",
		ExceptionType: ExceptionTypeDigestBypass,
		TenantID:      "acme",
		Environment:   "prod",
		Namespace:     "acme-prod",
		Repo:          "my-org/acme-app",
		ImageDigest:   "sha256:def456",
		Reason:        "digest bypass",
		TicketID:      "SEC-2",
		ApprovedBy:    "ops@example.com",
		TTLHours:      2,
	})
	if err != nil {
		t.Fatalf("CreateException() digest error = %v", err)
	}

	if _, err := store.Ingest(context.Background(), Event{
		Component:           "policy-engine",
		EventType:           EventTypeExceptionUsed,
		Decision:            DecisionAllow,
		TenantID:            "acme",
		Environment:         "prod",
		Namespace:           "acme-prod",
		Repo:                "my-org/acme-app",
		Digest:              "sha256:abc123",
		CVEID:               "CVE-2026-0001",
		IsException:         true,
		ExceptionID:         cveException.ExceptionID,
		ExceptionType:       cveException.ExceptionType,
		ExceptionReason:     cveException.Reason,
		ExceptionTicketID:   cveException.TicketID,
		ExceptionApprovedBy: cveException.ApprovedBy,
		ExceptionExpiresAt:  &cveException.ExpiresAt,
	}); err != nil {
		t.Fatalf("Ingest() cve event error = %v", err)
	}
	if _, err := store.Ingest(context.Background(), Event{
		Component:           "policy-engine",
		EventType:           EventTypeExceptionUsed,
		Decision:            DecisionAllow,
		TenantID:            "acme",
		Environment:         "prod",
		Namespace:           "acme-prod",
		Repo:                "my-org/acme-app",
		Digest:              "sha256:def456",
		IsException:         true,
		ExceptionID:         digestException.ExceptionID,
		ExceptionType:       digestException.ExceptionType,
		ExceptionReason:     digestException.Reason,
		ExceptionTicketID:   digestException.TicketID,
		ExceptionApprovedBy: digestException.ApprovedBy,
		ExceptionExpiresAt:  &digestException.ExpiresAt,
	}); err != nil {
		t.Fatalf("Ingest() digest event error = %v", err)
	}

	cveReport, err := store.ExceptionReport(context.Background(), ExceptionFilter{
		ExceptionType: ExceptionTypeCVEWhitelist,
		CVEID:         "CVE-2026-0001",
		Limit:         10,
	})
	if err != nil {
		t.Fatalf("ExceptionReport() cve error = %v", err)
	}
	if len(cveReport.Active) != 1 || cveReport.Active[0].ExceptionID != "EX-CVE" {
		t.Fatalf("unexpected cve active exceptions %#v", cveReport.Active)
	}
	if len(cveReport.RecentUsed) != 1 || cveReport.RecentUsed[0].ExceptionID != "EX-CVE" {
		t.Fatalf("unexpected cve used events %#v", cveReport.RecentUsed)
	}

	digestReport, err := store.ExceptionReport(context.Background(), ExceptionFilter{
		ImageDigest: "sha256:def456",
		Limit:       10,
	})
	if err != nil {
		t.Fatalf("ExceptionReport() digest error = %v", err)
	}
	if len(digestReport.Active) != 1 || digestReport.Active[0].ExceptionID != "EX-DIGEST" {
		t.Fatalf("unexpected digest active exceptions %#v", digestReport.Active)
	}
	if len(digestReport.RecentUsed) != 1 || digestReport.RecentUsed[0].ExceptionID != "EX-DIGEST" {
		t.Fatalf("unexpected digest used events %#v", digestReport.RecentUsed)
	}
}

func TestMemoryStoreApprovalLifecycleAndLogs(t *testing.T) {
	store := NewMemoryStore()
	base := time.Date(2026, 4, 15, 9, 0, 0, 0, time.UTC)
	store.now = func() time.Time { return base }

	pending, err := store.RequestException(context.Background(), ExceptionCreateRequest{
		ExceptionID:   "EX-PENDING-001",
		ExceptionType: ExceptionTypeBreakGlass,
		TenantID:      "acme",
		Environment:   "prod",
		Namespace:     "acme-prod",
		Reason:        "production restore",
		TicketID:      "INC-7000",
		TTLHours:      2,
	}, "demo-operator", "operator")
	if err != nil {
		t.Fatalf("RequestException() error = %v", err)
	}
	if pending.Status != ExceptionStatusPending || pending.RequestedBy != "demo-operator" {
		t.Fatalf("unexpected pending exception %#v", pending)
	}

	pendingResult, err := store.ValidateException(context.Background(), ExceptionValidationRequest{
		ExceptionID: "EX-PENDING-001",
		TenantID:    "acme",
		Environment: "prod",
		Namespace:   "acme-prod",
	})
	if err != nil {
		t.Fatalf("ValidateException() pending error = %v", err)
	}
	if pendingResult.Valid || pendingResult.Reason != "exception is pending approval" {
		t.Fatalf("unexpected pending validation %#v", pendingResult)
	}

	approved, err := store.ApproveException(context.Background(), "EX-PENDING-001", "demo-admin", "security_admin")
	if err != nil {
		t.Fatalf("ApproveException() error = %v", err)
	}
	if approved.Status != ExceptionStatusApproved || approved.ApprovedBy != "demo-admin" {
		t.Fatalf("unexpected approved exception %#v", approved)
	}

	validResult, err := store.ValidateException(context.Background(), ExceptionValidationRequest{
		ExceptionID: "EX-PENDING-001",
		TenantID:    "acme",
		Environment: "prod",
		Namespace:   "acme-prod",
	})
	if err != nil {
		t.Fatalf("ValidateException() approved error = %v", err)
	}
	if !validResult.Valid || validResult.Exception == nil || validResult.Exception.Status != ExceptionStatusApproved {
		t.Fatalf("unexpected approved validation %#v", validResult)
	}

	revoked, err := store.RevokeException(context.Background(), "EX-PENDING-001")
	if err != nil {
		t.Fatalf("RevokeException() error = %v", err)
	}
	if revoked.Status != ExceptionStatusRevoked {
		t.Fatalf("unexpected revoked exception %#v", revoked)
	}

	revokedResult, err := store.ValidateException(context.Background(), ExceptionValidationRequest{
		ExceptionID: "EX-PENDING-001",
		TenantID:    "acme",
		Environment: "prod",
		Namespace:   "acme-prod",
	})
	if err != nil {
		t.Fatalf("ValidateException() revoked error = %v", err)
	}
	if revokedResult.Valid || revokedResult.Reason != "exception is revoked" {
		t.Fatalf("unexpected revoked validation %#v", revokedResult)
	}

	rejectedPending, err := store.RequestException(context.Background(), ExceptionCreateRequest{
		ExceptionID:   "EX-PENDING-REJECT",
		ExceptionType: ExceptionTypeBreakGlass,
		TenantID:      "acme",
		Environment:   "prod",
		Namespace:     "acme-prod",
		Reason:        "unverified request",
		TicketID:      "INC-7001",
		TTLHours:      2,
	}, "demo-operator", "operator")
	if err != nil {
		t.Fatalf("RequestException() rejected pending error = %v", err)
	}
	if _, err := store.RejectException(context.Background(), rejectedPending.ExceptionID, "missing evidence", "demo-admin", "security_admin"); err != nil {
		t.Fatalf("RejectException() error = %v", err)
	}

	rejectedResult, err := store.ValidateException(context.Background(), ExceptionValidationRequest{
		ExceptionID: rejectedPending.ExceptionID,
		TenantID:    "acme",
		Environment: "prod",
		Namespace:   "acme-prod",
	})
	if err != nil {
		t.Fatalf("ValidateException() rejected error = %v", err)
	}
	if rejectedResult.Valid || rejectedResult.Reason != "exception is rejected" {
		t.Fatalf("unexpected rejected validation %#v", rejectedResult)
	}

	logs, err := store.ListApprovalLogs(context.Background(), "EX-PENDING-001", 10)
	if err != nil {
		t.Fatalf("ListApprovalLogs() error = %v", err)
	}
	actions := []string{}
	for _, log := range logs {
		actions = append(actions, log.Action)
	}
	for _, expected := range []string{ApprovalActionRequested, ApprovalActionApproved, ApprovalActionUsed, ApprovalActionRevoked, ApprovalActionValidationFailed} {
		found := false
		for _, action := range actions {
			if action == expected {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected approval action %q in %#v", expected, actions)
		}
	}
}

func TestMemoryStoreAnalytics(t *testing.T) {
	store := NewMemoryStore()
	base := time.Date(2026, 4, 15, 8, 0, 0, 0, time.UTC)
	current := base
	store.now = func() time.Time { return current }

	ingestAt := func(ts time.Time, event Event) {
		t.Helper()
		current = ts
		if _, err := store.Ingest(context.Background(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}

	ingestAt(base, Event{
		Component:   "deploy-gate",
		EventType:   EventTypeDeployGateDecision,
		Decision:    DecisionAllow,
		TenantID:    "acme",
		Environment: "prod",
		Repo:        "my-org/acme-app",
	})
	ingestAt(base.Add(30*time.Minute), Event{
		Component:   "deploy-gate",
		EventType:   EventTypeDeployGateDecision,
		Decision:    DecisionDeny,
		TenantID:    "acme",
		Environment: "prod",
		Repo:        "my-org/acme-app",
		Reasons:     []string{"workflow mismatch"},
	})
	ingestAt(base.Add(time.Hour), Event{
		Component:   "policy-engine",
		EventType:   EventTypePolicyDecision,
		Decision:    DecisionError,
		TenantID:    "acme",
		Environment: "prod",
		Repo:        "my-org/acme-app",
		Reasons:     []string{"policy load failed"},
	})
	ingestAt(base.Add(2*time.Hour), Event{
		Component:   "runtime-agent",
		EventType:   EventTypeRuntimeDriftResult,
		Decision:    DecisionDeny,
		TenantID:    "acme",
		Environment: "prod",
		Namespace:   "acme-prod",
		Workload:    "booking-api",
		Repo:        "my-org/acme-app",
		DriftResult: "image_drift",
	})
	ingestAt(base.Add(3*time.Hour), Event{
		Component:   "runtime-agent",
		EventType:   EventTypeRuntimeDriftResult,
		Decision:    DecisionAllow,
		TenantID:    "acme",
		Environment: "prod",
		Namespace:   "acme-prod",
		Workload:    "booking-api",
		Repo:        "my-org/acme-app",
		DriftResult: "no_drift",
	})

	trends, err := store.Trends(context.Background(), TrendsFilter{
		WindowDays:  30,
		Granularity: "day",
		TenantID:    "acme",
	})
	if err != nil {
		t.Fatalf("Trends() error = %v", err)
	}
	if len(trends.Buckets) == 0 || trends.Totals["allow"] != 2 || trends.Totals["deny"] != 2 || trends.Totals["error"] != 1 {
		t.Fatalf("unexpected trends %#v", trends)
	}

	topViolators, err := store.TopViolators(context.Background(), TopViolatorsFilter{
		WindowDays: 30,
		Dimension:  "repo",
		TenantID:   "acme",
	})
	if err != nil {
		t.Fatalf("TopViolators() error = %v", err)
	}
	if len(topViolators.Items) != 1 || topViolators.Items[0].Key != "my-org/acme-app" || topViolators.Items[0].DenyCount != 2 {
		t.Fatalf("unexpected top violators %#v", topViolators)
	}

	driftStats, err := store.DriftStats(context.Background(), DriftStatsFilter{
		WindowDays:  30,
		TenantID:    "acme",
		Environment: "prod",
		Repo:        "my-org/acme-app",
	})
	if err != nil {
		t.Fatalf("DriftStats() error = %v", err)
	}
	if driftStats.TotalRuntimeDriftDenies != 1 || driftStats.CountsByDriftClass["image_drift"] != 1 {
		t.Fatalf("unexpected drift stats %#v", driftStats)
	}
	if driftStats.MeanTimeToResolveSeconds == nil || *driftStats.MeanTimeToResolveSeconds != int64(time.Hour.Seconds()) {
		t.Fatalf("unexpected mttr %#v", driftStats)
	}
}
