package audit

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const defaultAuditRelativePath = "artifacts/audit/changelock-events.jsonl"

type Sink interface {
	Write(ctx context.Context, event Event) error
}

type Writer struct {
	sink Sink
	now  func() time.Time
}

func NewWriter(sink Sink) *Writer {
	if sink == nil {
		sink = NoopSink{}
	}

	return &Writer{
		sink: sink,
		now:  time.Now,
	}
}

func NewDefaultWriter() *Writer {
	return NewWriter(newDefaultSink())
}

func DefaultFilePath() string {
	if path := os.Getenv("CHANGELOCK_AUDIT_FILE"); path != "" {
		return path
	}
	return filepath.Clean(defaultAuditRelativePath)
}

func (w *Writer) Write(ctx context.Context, event Event) error {
	event = NormalizeEvent(event, w.now)
	return w.sink.Write(ctx, event)
}

func NewRequestID() string {
	var bytes [8]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		return time.Now().UTC().Format("20060102T150405.000000000")
	}
	return hex.EncodeToString(bytes[:])
}

func newDefaultSink() Sink {
	remoteURL := os.Getenv("AUDIT_WRITER_URL")
	if remoteURL == "" {
		remoteURL = os.Getenv("CHANGELOCK_AUDIT_WRITER_URL")
	}

	if remoteURL == "" {
		return NewFileSink(DefaultFilePath())
	}

	token := FirstNonEmpty(
		os.Getenv("CHANGELOCK_SYNC_TOKEN"),
		os.Getenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN"),
	)
	clusterID := strings.TrimSpace(os.Getenv("CHANGELOCK_CLUSTER_ID"))
	sinks := []Sink{NewHTTPSinkWithConfig(remoteURL, defaultHTTPSinkTimeout(), token, clusterID)}
	if filePath := os.Getenv("CHANGELOCK_AUDIT_FILE"); filePath != "" {
		sinks = append(sinks, NewFileSink(filePath))
	}
	return NewMultiSink(sinks...)
}

func defaultHTTPSinkTimeout() time.Duration {
	value := os.Getenv("AUDIT_WRITER_TIMEOUT")
	if value == "" {
		value = os.Getenv("CHANGELOCK_AUDIT_WRITER_TIMEOUT")
	}
	if value == "" {
		return 2 * time.Second
	}

	duration, err := time.ParseDuration(value)
	if err != nil || duration <= 0 {
		return 2 * time.Second
	}
	return duration
}
