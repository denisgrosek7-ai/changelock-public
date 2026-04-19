package audit

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/signing"
)

const (
	SyncModeDisabled = "disabled"
	SyncModeHub      = "hub"
	SyncModeSpoke    = "spoke"

	SyncFailModeLastKnownGood = "last-known-good"
	SyncFailModeDeny          = "deny"

	SyncHealthDisabled = "disabled"
	SyncHealthHealthy  = "healthy"
	SyncHealthStale    = "stale"
	SyncHealthError    = "error"
)

type SyncedException struct {
	ExceptionID   string            `json:"exception_id"`
	ExceptionType string            `json:"exception_type"`
	TenantID      string            `json:"tenant_id,omitempty"`
	Environment   string            `json:"environment,omitempty"`
	Namespace     string            `json:"namespace,omitempty"`
	Repo          string            `json:"repo,omitempty"`
	ImageDigest   string            `json:"image_digest,omitempty"`
	CVEID         string            `json:"cve_id,omitempty"`
	Reason        string            `json:"reason"`
	TicketID      string            `json:"ticket_id"`
	RequestedBy   string            `json:"requested_by,omitempty"`
	RequestedAt   *time.Time        `json:"requested_at,omitempty"`
	ApprovedBy    string            `json:"approved_by,omitempty"`
	ApprovedAt    *time.Time        `json:"approved_at,omitempty"`
	CreatedAt     time.Time         `json:"created_at"`
	ExpiresAt     time.Time         `json:"expires_at"`
	LastUpdatedAt *time.Time        `json:"last_updated_at,omitempty"`
	Signature     *signing.Envelope `json:"signature,omitempty"`
	Metadata      json.RawMessage   `json:"metadata,omitempty"`
}

type ExceptionSyncSnapshot struct {
	ClusterID   string            `json:"cluster_id"`
	Revision    string            `json:"revision"`
	GeneratedAt time.Time         `json:"generated_at"`
	Exceptions  []SyncedException `json:"exceptions"`
	Signature   *signing.Envelope `json:"signature,omitempty"`
}

type SyncStatus struct {
	SyncMode             string     `json:"sync_mode,omitempty"`
	Mode                 string     `json:"mode"`
	ClusterID            string     `json:"cluster_id,omitempty"`
	HubURL               string     `json:"hub_url,omitempty"`
	FailMode             string     `json:"fail_mode,omitempty"`
	Health               string     `json:"health"`
	CurrentRevision      string     `json:"current_revision,omitempty"`
	RevisionETag         string     `json:"revision_etag,omitempty"`
	LastSuccessfulSyncAt *time.Time `json:"last_successful_sync_at,omitempty"`
	LastAttemptAt        *time.Time `json:"last_attempt_at,omitempty"`
	LastError            string     `json:"last_error,omitempty"`
	CachePresent         bool       `json:"cache_present"`
	StaleAfterSeconds    int64      `json:"stale_after_seconds,omitempty"`
	SignerMode           string     `json:"signer_mode,omitempty"`
	VerificationState    string     `json:"verification_state,omitempty"`
	VerificationReason   string     `json:"verification_reason,omitempty"`
	Summary              string     `json:"summary,omitempty"`
}

func SyncedExceptionFromPolicyException(exception PolicyException) SyncedException {
	return SyncedException{
		ExceptionID:   exception.ExceptionID,
		ExceptionType: exception.ExceptionType,
		TenantID:      exception.TenantID,
		Environment:   exception.Environment,
		Namespace:     exception.Namespace,
		Repo:          exception.Repo,
		ImageDigest:   exception.ImageDigest,
		CVEID:         exception.CVEID,
		Reason:        exception.Reason,
		TicketID:      exception.TicketID,
		RequestedBy:   exception.RequestedBy,
		RequestedAt:   cloneTimePointer(exception.RequestedAt),
		ApprovedBy:    exception.ApprovedBy,
		ApprovedAt:    cloneTimePointer(exception.ApprovedAt),
		CreatedAt:     exception.CreatedAt.UTC(),
		ExpiresAt:     exception.ExpiresAt.UTC(),
		LastUpdatedAt: cloneTimePointer(exception.LastUpdatedAt),
		Signature:     cloneSignatureEnvelope(exception.Signature),
		Metadata:      normalizeMetadata(exception.Metadata),
	}
}

func (exception SyncedException) ToPolicyException(now time.Time, id int64) PolicyException {
	if now.IsZero() {
		now = time.Now().UTC()
	}
	createdAt := exception.CreatedAt.UTC()
	if createdAt.IsZero() {
		createdAt = now
	}
	approvedAt := cloneTimePointer(exception.ApprovedAt)
	if approvedAt == nil {
		approvedAt = timePointer(createdAt)
	}
	lastUpdatedAt := cloneTimePointer(exception.LastUpdatedAt)
	if lastUpdatedAt == nil {
		lastUpdatedAt = cloneTimePointer(approvedAt)
	}
	return PolicyException{
		ID:            id,
		ExceptionID:   strings.TrimSpace(exception.ExceptionID),
		ExceptionType: normalizeExceptionType(exception.ExceptionType),
		Status:        ExceptionStatusApproved,
		TenantID:      strings.TrimSpace(exception.TenantID),
		Environment:   strings.TrimSpace(exception.Environment),
		Namespace:     strings.TrimSpace(exception.Namespace),
		Repo:          strings.TrimSpace(exception.Repo),
		ImageDigest:   strings.TrimSpace(exception.ImageDigest),
		CVEID:         strings.TrimSpace(strings.ToUpper(exception.CVEID)),
		Reason:        strings.TrimSpace(exception.Reason),
		TicketID:      strings.TrimSpace(exception.TicketID),
		RequestedBy:   strings.TrimSpace(exception.RequestedBy),
		RequestedAt:   cloneTimePointer(exception.RequestedAt),
		ApprovedBy:    strings.TrimSpace(exception.ApprovedBy),
		ApprovedAt:    approvedAt,
		CreatedAt:     createdAt,
		ExpiresAt:     exception.ExpiresAt.UTC(),
		Active:        true,
		LastUpdatedAt: lastUpdatedAt,
		Signature:     cloneSignatureEnvelope(exception.Signature),
		Metadata:      normalizeMetadata(exception.Metadata),
	}
}

func ComputeExceptionSyncRevision(exceptions []SyncedException) string {
	normalized := make([]SyncedException, 0, len(exceptions))
	for _, exception := range exceptions {
		normalized = append(normalized, cloneSyncedException(exception))
	}
	sort.Slice(normalized, func(i, j int) bool {
		return normalized[i].ExceptionID < normalized[j].ExceptionID
	})
	payload, err := json.Marshal(normalized)
	if err != nil {
		return ""
	}
	sum := sha256.Sum256(payload)
	return hex.EncodeToString(sum[:])
}

func cloneTimePointer(value *time.Time) *time.Time {
	if value == nil {
		return nil
	}
	copy := value.UTC()
	return &copy
}
