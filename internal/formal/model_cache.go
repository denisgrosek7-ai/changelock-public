package formal

import (
	"encoding/json"
	"sync"
)

var (
	point12ValEFoundationModelOnce   sync.Once
	point12ValEFoundationModelCached Point12ValEFoundation

	point13ValEFoundationModelOnce   sync.Once
	point13ValEFoundationModelCached Point13ValEFoundation

	point14Val0DependencySnapshotModelOnce   sync.Once
	point14Val0DependencySnapshotModelCached Point14Val0DependencySnapshot
	point14Val0FoundationModelOnce           sync.Once
	point14Val0FoundationModelCached         Point14Val0Foundation

	point14ValADependencySnapshotModelOnce   sync.Once
	point14ValADependencySnapshotModelCached Point14ValADependencySnapshot
	point14ValAFoundationModelOnce           sync.Once
	point14ValAFoundationModelCached         Point14ValAFoundation

	point14ValBDependencySnapshotModelOnce   sync.Once
	point14ValBDependencySnapshotModelCached Point14ValBDependencySnapshot
	point14ValBFoundationModelOnce           sync.Once
	point14ValBFoundationModelCached         Point14ValBFoundation

	point14ValCDependencySnapshotModelOnce   sync.Once
	point14ValCDependencySnapshotModelCached Point14ValCDependencySnapshot
	point14ValCFoundationModelOnce           sync.Once
	point14ValCFoundationModelCached         Point14ValCFoundation

	point14ValDDependencySnapshotModelOnce   sync.Once
	point14ValDDependencySnapshotModelCached Point14ValDDependencySnapshot
	point14ValDFoundationModelOnce           sync.Once
	point14ValDFoundationModelCached         Point14ValDFoundation

	point14ValEDependencySnapshotModelOnce   sync.Once
	point14ValEDependencySnapshotModelCached Point14ValEDependencySnapshot
	point14ValEFoundationModelOnce           sync.Once
	point14ValEFoundationModelCached         Point14ValEFoundation
)

func cloneFormalModel[T any](value T) T {
	payload, err := json.Marshal(value)
	if err != nil {
		panic("formal model cache marshal failed: " + err.Error())
	}
	var cloned T
	if err := json.Unmarshal(payload, &cloned); err != nil {
		panic("formal model cache unmarshal failed: " + err.Error())
	}
	return cloned
}

func cachedFormalModel[T any](once *sync.Once, cached *T, build func() T) T {
	once.Do(func() {
		*cached = build()
	})
	return cloneFormalModel(*cached)
}
