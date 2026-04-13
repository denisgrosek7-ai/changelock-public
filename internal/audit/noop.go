package audit

import "context"

type NoopSink struct{}

func (NoopSink) Write(context.Context, Event) error {
	return nil
}
