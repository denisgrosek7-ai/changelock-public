package audit

import (
	"encoding/json"
	"slices"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/signing"
)

func CanonicalExceptionEvidence(exception PolicyException) ([]byte, error) {
	type payload struct {
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
		RequestedBy   string          `json:"requested_by,omitempty"`
		RequestedAt   *time.Time      `json:"requested_at,omitempty"`
		ApprovedBy    string          `json:"approved_by,omitempty"`
		ApprovedAt    *time.Time      `json:"approved_at,omitempty"`
		CreatedAt     time.Time       `json:"created_at"`
		ExpiresAt     time.Time       `json:"expires_at"`
		Metadata      json.RawMessage `json:"metadata,omitempty"`
	}

	return json.Marshal(payload{
		ExceptionID:   strings.TrimSpace(exception.ExceptionID),
		ExceptionType: strings.TrimSpace(exception.ExceptionType),
		TenantID:      strings.TrimSpace(exception.TenantID),
		Environment:   strings.TrimSpace(exception.Environment),
		Namespace:     strings.TrimSpace(exception.Namespace),
		Repo:          strings.TrimSpace(exception.Repo),
		ImageDigest:   strings.TrimSpace(exception.ImageDigest),
		CVEID:         strings.TrimSpace(exception.CVEID),
		Reason:        strings.TrimSpace(exception.Reason),
		TicketID:      strings.TrimSpace(exception.TicketID),
		RequestedBy:   strings.TrimSpace(exception.RequestedBy),
		RequestedAt:   cloneTimePointer(exception.RequestedAt),
		ApprovedBy:    strings.TrimSpace(exception.ApprovedBy),
		ApprovedAt:    cloneTimePointer(exception.ApprovedAt),
		CreatedAt:     exception.CreatedAt.UTC(),
		ExpiresAt:     exception.ExpiresAt.UTC(),
		Metadata:      normalizeMetadata(exception.Metadata),
	})
}

func CanonicalExceptionSyncSnapshot(snapshot ExceptionSyncSnapshot) ([]byte, error) {
	type payload struct {
		ClusterID   string            `json:"cluster_id"`
		Revision    string            `json:"revision"`
		GeneratedAt time.Time         `json:"generated_at"`
		Exceptions  []SyncedException `json:"exceptions"`
	}

	normalized := make([]SyncedException, 0, len(snapshot.Exceptions))
	for _, exception := range snapshot.Exceptions {
		normalized = append(normalized, cloneSyncedException(exception))
	}
	return json.Marshal(payload{
		ClusterID:   strings.TrimSpace(snapshot.ClusterID),
		Revision:    strings.TrimSpace(snapshot.Revision),
		GeneratedAt: snapshot.GeneratedAt.UTC(),
		Exceptions:  normalized,
	})
}

func cloneSignatureEnvelope(envelope *signing.Envelope) *signing.Envelope {
	if envelope == nil {
		return nil
	}
	copy := *envelope
	copy.SignedAt = envelope.SignedAt.UTC()
	return &copy
}

func cloneSyncedException(exception SyncedException) SyncedException {
	if exception.Metadata != nil {
		exception.Metadata = slices.Clone(exception.Metadata)
	}
	exception.Signature = cloneSignatureEnvelope(exception.Signature)
	exception.RequestedAt = cloneTimePointer(exception.RequestedAt)
	exception.ApprovedAt = cloneTimePointer(exception.ApprovedAt)
	exception.LastUpdatedAt = cloneTimePointer(exception.LastUpdatedAt)
	exception.CreatedAt = exception.CreatedAt.UTC()
	exception.ExpiresAt = exception.ExpiresAt.UTC()
	return exception
}
