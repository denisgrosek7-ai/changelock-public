package audit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
)

const (
	ExceptionTypeBreakGlass   = "BREAK_GLASS"
	ExceptionTypeDigestBypass = "DIGEST_BYPASS"
	ExceptionTypeCVEWhitelist = "CVE_WHITELIST"

	EventTypeExceptionCreated          = "exception_created"
	EventTypeExceptionRevoked          = "exception_revoked"
	EventTypeExceptionUsed             = "exception_used"
	EventTypeExceptionValidationFailed = "exception_validation_failed"
)

var (
	ErrInvalidException  = errors.New("invalid policy exception")
	ErrExceptionNotFound = errors.New("policy exception not found")
)

type PolicyException struct {
	ID            int64           `json:"id"`
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
	ApprovedBy    string          `json:"approved_by"`
	CreatedAt     time.Time       `json:"created_at"`
	ExpiresAt     time.Time       `json:"expires_at"`
	Active        bool            `json:"active"`
	Metadata      json.RawMessage `json:"metadata,omitempty"`
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
	ApprovedBy    string          `json:"approved_by"`
	ExpiresAt     *time.Time      `json:"expires_at,omitempty"`
	TTLHours      int             `json:"ttl_hours,omitempty"`
	Metadata      json.RawMessage `json:"metadata,omitempty"`
}

type ExceptionFilter struct {
	Active        *bool
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
	Valid     bool             `json:"valid"`
	Reason    string           `json:"reason,omitempty"`
	Exception *PolicyException `json:"exception,omitempty"`
}

type ExceptionReport struct {
	Active         []PolicyException `json:"active"`
	RecentUsed     []StoredEvent     `json:"recent_used"`
	RecentInactive []PolicyException `json:"recent_inactive"`
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
	case request.ApprovedBy == "":
		return request, fmt.Errorf("%w: approved_by is required", ErrInvalidException)
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

func NormalizeExceptionFilter(filter ExceptionFilter) (ExceptionFilter, error) {
	filter.ExceptionType = normalizeExceptionType(filter.ExceptionType)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	filter.Environment = strings.TrimSpace(filter.Environment)
	filter.Namespace = strings.TrimSpace(filter.Namespace)
	filter.Repo = strings.TrimSpace(filter.Repo)
	filter.ImageDigest = strings.TrimSpace(filter.ImageDigest)
	filter.CVEID = strings.TrimSpace(strings.ToUpper(filter.CVEID))

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

func (exception PolicyException) IsCurrentlyActive(now time.Time) bool {
	return exception.Active && exception.ExpiresAt.After(now.UTC())
}

func (exception PolicyException) Matches(request ExceptionValidationRequest, now time.Time) (bool, string) {
	switch {
	case !exception.Active:
		return false, "exception is inactive"
	case !exception.ExpiresAt.After(now.UTC()):
		return false, "exception is expired"
	case request.ExceptionType != "" && exception.ExceptionType != request.ExceptionType:
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

func clonePolicyException(exception PolicyException) PolicyException {
	if exception.Metadata != nil {
		exception.Metadata = slices.Clone(exception.Metadata)
	}
	return exception
}

func normalizeExceptionType(value string) string {
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
