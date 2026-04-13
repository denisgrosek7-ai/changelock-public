package audit

import (
	"context"
	"errors"
)

type MultiSink struct {
	sinks []Sink
}

func NewMultiSink(sinks ...Sink) *MultiSink {
	filtered := make([]Sink, 0, len(sinks))
	for _, sink := range sinks {
		if sink != nil {
			filtered = append(filtered, sink)
		}
	}
	return &MultiSink{sinks: filtered}
}

func (s *MultiSink) Write(ctx context.Context, event Event) error {
	if s == nil || len(s.sinks) == 0 {
		return nil
	}

	var errs []error
	for _, sink := range s.sinks {
		if err := sink.Write(ctx, event); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}
