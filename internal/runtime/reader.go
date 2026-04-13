package runtime

import (
	"context"
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var ErrObservedStateUnavailable = errors.New("observed runtime state unavailable")

type WorkloadTarget struct {
	Namespace string
	Workload  string
}

type StateReader interface {
	ReadObservedWorkload(ctx context.Context, target WorkloadTarget) (ObservedWorkloadState, error)
}

type NoopReader struct{}

func (NoopReader) ReadObservedWorkload(_ context.Context, _ WorkloadTarget) (ObservedWorkloadState, error) {
	return ObservedWorkloadState{}, ErrObservedStateUnavailable
}

type FixtureReader struct {
	states map[string]ObservedWorkloadState
}

type fixtureDocument struct {
	Workloads []ObservedWorkloadState `yaml:"workloads"`
}

func NewFixtureReader(path string) (*FixtureReader, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read runtime fixture: %w", err)
	}

	var document fixtureDocument
	if err := yaml.Unmarshal(data, &document); err != nil {
		return nil, fmt.Errorf("decode runtime fixture: %w", err)
	}

	states := make(map[string]ObservedWorkloadState, len(document.Workloads))
	for _, workload := range document.Workloads {
		if workload.Namespace == "" || workload.Workload == "" {
			continue
		}
		states[targetKey(WorkloadTarget{Namespace: workload.Namespace, Workload: workload.Workload})] = workload
	}

	return &FixtureReader{states: states}, nil
}

func (r *FixtureReader) ReadObservedWorkload(_ context.Context, target WorkloadTarget) (ObservedWorkloadState, error) {
	if r == nil {
		return ObservedWorkloadState{}, ErrObservedStateUnavailable
	}

	state, ok := r.states[targetKey(target)]
	if !ok {
		return ObservedWorkloadState{}, fmt.Errorf("%w: %s/%s", ErrObservedStateUnavailable, target.Namespace, target.Workload)
	}
	return state, nil
}

func targetKey(target WorkloadTarget) string {
	return target.Namespace + "/" + target.Workload
}
