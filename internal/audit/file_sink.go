package audit

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type FileSink struct {
	path string
	mu   sync.Mutex
}

func NewFileSink(path string) *FileSink {
	return &FileSink{path: path}
}

func (s *FileSink) Write(ctx context.Context, event Event) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}

	file, err := os.OpenFile(s.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	if _, err := file.Write(append(payload, '\n')); err != nil {
		return err
	}

	return nil
}
