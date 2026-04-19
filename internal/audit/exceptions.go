package audit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/signing"
)

const (
	ExceptionTypeBreakGlass   = "BREAK_GLASS"
	ExceptionTypeDigestBypass = "DIGEST_BYPASS"
	ExceptionTypeCVEWhitelist = "CVE_WHITELIST"

	ExceptionStatusPending  = "PENDING"
	ExceptionStatusApproved = "APPROVED"
	ExceptionStatusRejected = "REJECTED"
	ExceptionStatusRevoked  = "REVOKED"
	ExceptionStatusExpired  = "EXPIRED"

	ApprovalActionRequested        = "REQUESTED"
	ApprovalActionApproved         = "APPROVED"
	ApprovalActionRejected         = "REJECTED"
	ApprovalActionRevoked          = "REVOKED"
	ApprovalActionUsed             = "USED"
	ApprovalActionValidationFailed = "VALIDATION_FAILED"

	EventTypeExceptionRequested        = "exception_requested"
	EventTypeExceptionApproved         = "exception_approved"
	EventTypeExceptionRejected         = "exception_rejected"
	EventTypeExceptionRevoked          = "exception_revoked"
	EventTypeExceptionUsed             = "exception_used"
	EventTypeExceptionValidationFailed = "exception_validation_failed"
)

var (
	ErrInvalidException  = errors.New("invalid policy exception")
	ErrExceptionNotFound = errors.New("policy exception not found")
)

type PolicyException struct {
	ID                 int64             `json:"id"`
	ExceptionID        string            `json:"exception_id"`
	ExceptionType      string            `json:"exception_type"`
	Status             string            `json:"status"`
	TenantID           string            `json:"tenant_id,omitempty"`
	Environment        string            `json:"environment,omitempty"`
	Namespace          string            `json:"namespace,omitempty"`
	Repo               string            `json:"repo,omitempty"`
	ImageDigest        string            `json:"image_digest,omitempty"`
	CVEID              string            `json:"cve_id,omitempty"`
	Reason             string            `json:"reason"`
	TicketID           string            `json:"ticket_id"`
	RequestedBy        string            `json:"requested_by,omitempty"`
	RequestedAt        *time.Time        `json:"requested_at,omitempty"`
	ApprovedBy         string            `json:"approved_by,omitempty"`
	ApprovedAt         *time.Time        `json:"approved_at,omitempty"`
	RejectedBy         string            `json:"rejected_by,omitempty"`
	RejectedAt         *time.Time        `json:"rejected_at,omitempty"`
	RejectionReason    string            `json:"rejection_reason,omitempty"`
	CreatedAt          time.Time         `json:"created_at"`
	ExpiresAt          time.Time         `json:"expires_at"`
	Active             bool              `json:"active"`
	LastUpdatedAt      *time.Time        `json:"last_updated_at,omitempty"`
	Signature          *signing.Envelope `json:"signature,omitempty"`
	VerificationState  string            `json:"verification_state,omitempty"`
	VerificationReason string            `json:"verification_reason,omitempty"`
	Metadata           json.RawMessage   `json:"metadata,omitempty"`
}

type ApprovalLog struct {
	ID          int64           `json:"id"`
	ExceptionID string          `json:"exception_id"`
	Action      string          `json:"action"`
	Actor       string          `json:"actor"`
	ActorRole   string          `json:"actor_role,omitempty"`
	Reason      string          `json:"reason,omitempty"`
	CreatedAt   time.Time       `json:"created_at"`
	Metadata    json.RawMessage `json:"metadata,omitempty"`
}

type ExceptionCreateRequest struct {
	ExceptionID   string          `json:"exception_id"`
	ExceptionType string          `json:"exception_type"`
	TenantID      string          `json:"tenant_id,omitempty"`
	Environment   string          `json:"environment,omitempty"`
	Namespace     string          `json:"namespace,omitempty"`
	Repo          string          `json:"repo,omitempty"`
	ImageDigest   string          `json:"image_digest,omitempty"`
	CVEID         string          `json:"cve_id,omitempty"`
	Reason        string          `json:"reason"`
	TicketID      string          `json:"ticket_id"`
	ApprovedBy    string          `json:"approved_by,omitempty"`
	ExpiresAt     *time.Time      `json:"expires_at,omitempty"`
	TTLHours      int             `json:"ttl_hours,omitempty"`
	Metadata      json.RawMessage `json:"metadata,omitempty"`
}

type ExceptionActionRequest struct {
	Reason   string          `json:"reason,omitempty"`
	Metadata json.RawMessage `json:"metadata,omitempty"`
}

type ExceptionFilter struct {
	Active        *bool
	Status        string
	ExceptionType string
	TenantID      string
	Environment   string
	Namespace     string
	Repo          string
	ImageDigest   string
	CVEID         string
	Limit         int
}

type ExceptionValidationRequest struct {
	ExceptionID   string `json:"exception_id"`
	ExceptionType string `json:"exception_type,omitempty"`
	TenantID      string `json:"tenant_id,omitempty"`
	Environment   string `json:"environment,omitempty"`
	Namespace     string `json:"namespace,omitempty"`
	Repo          string `json:"repo,omitempty"`
	ImageDigest   string `json:"image_digest,omitempty"`
	CVEID         string `json:"cve_id,omitempty"`
}

type ExceptionValidationResult struct {
	Valid              bool             `json:"valid"`
	Reason             string           `json:"reason,omitempty"`
	VerificationState  string           `json:"verification_state,omitempty"`
	VerificationReason string           `json:"verification_reason,omitempty"`
	Exception          *PolicyException `json:"exception,omitempty"`
}

type ExceptionReport struct {
	Active         []PolicyException `json:"active"`
	Pending        []PolicyException `json:"pending,omitempty"`
	Rejected       []PolicyException `json:"rejected,omitempty"`
	Revoked        []PolicyException `json:"revoked,omitempty"`
	Expired        []PolicyException `json:"expired,omitempty"`
	RecentUsed     []StoredEvent     `json:"recent_used"`
	RecentInactive []PolicyException `json:"recent_inactive"`
	StatusCounts   map[string]int64  `json:"status_counts,omitempty"`
}

type ExceptionValidator interface {
	Validate(ctx context.Context, request ExceptionValidationRequest) (ExceptionValidationResult, error)
}

func NormalizeExceptionCreateRequest(request ExceptionCreateRequest, now func() time.Time) (ExceptionCreateRequest, error) {
	if now == nil {
		now = time.Now
	}

	request.ExceptionID = strings.TrimSpace(request.ExceptionID)
	request.ExceptionType = normalizeExceptionType(request.ExceptionType)
	request.TenantID = strings.TrimSpace(request.TenantID)
	request.Environment = strings.TrimSpace(request.Environment)
	request.Namespace = strings.TrimSpace(request.Namespace)
	request.Repo = strings.TrimSpace(request.Repo)
	request.ImageDigest = strings.TrimSpace(request.ImageDigest)
	request.CVEID = strings.TrimSpace(strings.ToUpper(request.CVEID))
	request.Reason = strings.TrimSpace(request.Reason)
	request.TicketID = strings.TrimSpace(request.TicketID)
	request.ApprovedBy = strings.TrimSpace(request.ApprovedBy)
	request.Metadata = normalizeMetadata(request.Metadata)

	switch {
	case request.ExceptionID == "":
		return request, fmt.Errorf("%w: exception_id is required", ErrInvalidException)
	case request.ExceptionType == "":
		return request, fmt.Errorf("%w: exception_type is required", ErrInvalidException)
	case request.Reason == "":
		return request, fmt.Errorf("%w: reason is required", ErrInvalidException)
	case request.TicketID == "":
		return request, fmt.Errorf("%w: ticket_id is required", ErrInvalidException)
	}

	switch request.ExceptionType {
	case ExceptionTypeBreakGlass, ExceptionTypeDigestBypass, ExceptionTypeCVEWhitelist:
	default:
		return request, fmt.Errorf("%w: unsupported exception_type %q", ErrInvalidException, request.ExceptionType)
	}

	if request.ExpiresAt == nil {
		if request.TTLHours <= 0 {
			return request, fmt.Errorf("%w: expires_at or ttl_hours is required", ErrInvalidException)
		}
		expiresAt := now().UTC().Add(time.Duration(request.TTLHours) * time.Hour)
		request.ExpiresAt = &expiresAt
	}
	expiresAt := request.ExpiresAt.UTC()
	if !expiresAt.After(now().UTC()) {
		return request, fmt.Errorf("%w: expires_at must be in the future", ErrInvalidException)
	}
	request.ExpiresAt = &expiresAt

	if err := validateExceptionScope(request); err != nil {
		return request, err
	}

	return request, nil
}

func NormalizeExceptionActionRequest(request ExceptionActionRequest) ExceptionActionRequest {
	request.Reason = strings.TrimSpace(request.Reason)
	request.Metadata = normalizeMetadata(request.Metadata)
	return request
}

func NormalizeExceptionFilter(filter ExceptionFilter) (ExceptionFilter, error) {
	filter.Status = normalizeExceptionStatus(filter.Status)
	filter.ExceptionType = normalizeExceptionType(filter.ExceptionType)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	filter.Environment = strings.TrimSpace(filter.Environment)
	filter.Namespace = strings.TrimSpace(filter.Namespace)
	filter.Repo = strings.TrimSpace(filter.Repo)
	filter.ImageDigest = strings.TrimSpace(filter.ImageDigest)
	filter.CVEID = strings.TrimSpace(strings.ToUpper(filter.CVEID))

	if filter.Status != "" {
		switch filter.Status {
		case ExceptionStatusPending, ExceptionStatusApproved, ExceptionStatusRejected, ExceptionStatusRevoked, ExceptionStatusExpired:
		default:
			return filter, fmt.Errorf("%w: unsupported status %q", ErrInvalidException, filter.Status)
		}
	}

	if filter.ExceptionType != "" {
		switch filter.ExceptionType {
		case ExceptionTypeBreakGlass, ExceptionTypeDigestBypass, ExceptionTypeCVEWhitelist:
		default:
			return filter, fmt.Errorf("%w: unsupported exception_type %q", ErrInvalidException, filter.ExceptionType)
		}
	}

	if filter.Limit <= 0 {
		filter.Limit = 50
	}
	if filter.Limit > 500 {
		filter.Limit = 500
	}

	return filter, nil
}

func NormalizeExceptionValidationRequest(request ExceptionValidationRequest) (ExceptionValidationRequest, error) {
	request.ExceptionID = strings.TrimSpace(request.ExceptionID)
	request.ExceptionType = normalizeExceptionType(request.ExceptionType)
	request.TenantID = strings.TrimSpace(request.TenantID)
	request.Environment = strings.TrimSpace(request.Environment)
	request.Namespace = strings.TrimSpace(request.Namespace)
	request.Repo = strings.TrimSpace(request.Repo)
	request.ImageDigest = strings.TrimSpace(request.ImageDigest)
	request.CVEID = strings.TrimSpace(strings.ToUpper(request.CVEID))

	if request.ExceptionID == "" {
		return request, fmt.Errorf("%w: exception_id is required", ErrInvalidException)
	}

	return request, nil
}

func NormalizeApprovalLog(log ApprovalLog) ApprovalLog {
	log.Action = normalizeApprovalAction(log.Action)
	log.Actor = strings.TrimSpace(log.Actor)
	log.ActorRole = strings.TrimSpace(log.ActorRole)
	log.Reason = strings.TrimSpace(log.Reason)
	log.Metadata = normalizeMetadata(log.Metadata)
	return log
}

func (exception PolicyException) EffectiveStatus(now time.Time) string {
	status := normalizeExceptionStatus(exception.Status)
	if status == "" {
		if exception.Active {
			status = ExceptionStatusApproved
		} else {
			status = ExceptionStatusRevoked
		}
	}

	switch status {
	case ExceptionStatusRejected, ExceptionStatusRevoked:
		return status
	case ExceptionStatusPending, ExceptionStatusApproved:
		if !exception.ExpiresAt.After(now.UTC()) {
			return ExceptionStatusExpired
		}
		if status == ExceptionStatusApproved && !exception.Active {
			return ExceptionStatusRevoked
		}
		return status
	case ExceptionStatusExpired:
		return ExceptionStatusExpired
	default:
		if !exception.ExpiresAt.After(now.UTC()) {
			return ExceptionStatusExpired
		}
		if exception.Active {
			return ExceptionStatusApproved
		}
		return ExceptionStatusRevoked
	}
}

func (exception PolicyException) IsCurrentlyActive(now time.Time) bool {
	return exception.EffectiveStatus(now) == ExceptionStatusApproved
}

func (exception PolicyException) Matches(request ExceptionValidationRequest, now time.Time) (bool, string) {
	switch status := exception.EffectiveStatus(now); status {
	case ExceptionStatusPending:
		return false, "exception is pending approval"
	case ExceptionStatusRejected:
		return false, "exception is rejected"
	case ExceptionStatusRevoked:
		return false, "exception is revoked"
	case ExceptionStatusExpired:
		return false, "exception is expired"
	case ExceptionStatusApproved:
	default:
		return false, "exception is inactive"
	}

	if request.ExceptionType != "" && exception.ExceptionType != request.ExceptionType {
		return false, "exception type does not match"
	}

	for _, candidate := range []struct {
		label    string
		expected string
		actual   string
	}{
		{label: "tenant_id", expected: exception.TenantID, actual: request.TenantID},
		{label: "environment", expected: exception.Environment, actual: request.Environment},
		{label: "namespace", expected: exception.Namespace, actual: request.Namespace},
		{label: "repo", expected: exception.Repo, actual: request.Repo},
		{label: "image_digest", expected: exception.ImageDigest, actual: request.ImageDigest},
		{label: "cve_id", expected: exception.CVEID, actual: request.CVEID},
	} {
		if candidate.expected == "" {
			continue
		}
		if candidate.actual == "" {
			return false, candidate.label + " is required for this exception scope"
		}
		if candidate.expected != candidate.actual {
			return false, candidate.label + " does not match"
		}
	}

	return true, ""
}

func (exception PolicyException) WithEffectiveStatus(now time.Time) PolicyException {
	copy := clonePolicyException(exception)
	copy.Status = copy.EffectiveStatus(now)
	return copy
}

func clonePolicyException(exception PolicyException) PolicyException {
	if exception.Metadata != nil {
		exception.Metadata = slices.Clone(exception.Metadata)
	}
	exception.Signature = cloneSignatureEnvelope(exception.Signature)
	if exception.RequestedAt != nil {
		requestedAt := exception.RequestedAt.UTC()
		exception.RequestedAt = &requestedAt
	}
	if exception.ApprovedAt != nil {
		approvedAt := exception.ApprovedAt.UTC()
		exception.ApprovedAt = &approvedAt
	}
	if exception.RejectedAt != nil {
		rejectedAt := exception.RejectedAt.UTC()
		exception.RejectedAt = &rejectedAt
	}
	if exception.LastUpdatedAt != nil {
		lastUpdatedAt := exception.LastUpdatedAt.UTC()
		exception.LastUpdatedAt = &lastUpdatedAt
	}
	return exception
}

func cloneApprovalLog(log ApprovalLog) ApprovalLog {
	if log.Metadata != nil {
		log.Metadata = slices.Clone(log.Metadata)
	}
	return log
}

func normalizeExceptionType(value string) string {
	return strings.ToUpper(strings.TrimSpace(value))
}

func normalizeExceptionStatus(value string) string {
	return strings.ToUpper(strings.TrimSpace(value))
}

func normalizeApprovalAction(value string) string {
	return strings.ToUpper(strings.TrimSpace(value))
}

func normalizeMetadata(value json.RawMessage) json.RawMessage {
	if len(value) == 0 {
		return json.RawMessage(`{}`)
	}
	if strings.TrimSpace(string(value)) == "" {
		return json.RawMessage(`{}`)
	}
	return slices.Clone(value)
}

func validateExceptionScope(request ExceptionCreateRequest) error {
	switch request.ExceptionType {
	case ExceptionTypeBreakGlass:
		if FirstNonEmpty(request.TenantID, request.Environment, request.Namespace, request.Repo, request.ImageDigest) == "" {
			return fmt.Errorf("%w: break-glass exceptions require tenant_id, environment, namespace, repo, or image_digest scope", ErrInvalidException)
		}
	case ExceptionTypeDigestBypass:
		if request.ImageDigest == "" {
			return fmt.Errorf("%w: digest bypass exceptions require image_digest", ErrInvalidException)
		}
	case ExceptionTypeCVEWhitelist:
		if request.CVEID == "" {
			return fmt.Errorf("%w: CVE whitelist exceptions require cve_id", ErrInvalidException)
		}
	}
	return nil
}
