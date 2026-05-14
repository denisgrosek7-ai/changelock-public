package operability

import (
	"encoding/json"
	"sync"
)

var (
	deploymentMultiTenantValEFoundationModelOnce   sync.Once
	deploymentMultiTenantValEFoundationModelCached DeploymentMultiTenantValEFoundation
)

func cloneOperabilityModel[T any](value T) T {
	payload, err := json.Marshal(value)
	if err != nil {
		panic("operability model cache marshal failed: " + err.Error())
	}
	var cloned T
	if err := json.Unmarshal(payload, &cloned); err != nil {
		panic("operability model cache unmarshal failed: " + err.Error())
	}
	return cloned
}

func cachedOperabilityModel[T any](once *sync.Once, cached *T, build func() T) T {
	once.Do(func() {
		*cached = build()
	})
	return cloneOperabilityModel(*cached)
}
