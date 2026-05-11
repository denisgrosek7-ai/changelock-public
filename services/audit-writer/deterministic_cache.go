package main

import (
	"encoding/json"
	"sync"
)

type cachedJSONValue struct {
	once    sync.Once
	payload []byte
}

func loadCachedJSON[T any](cache *cachedJSONValue, build func() T) T {
	cache.once.Do(func() {
		value := build()
		payload, err := json.Marshal(value)
		if err == nil {
			cache.payload = payload
		}
	})
	if len(cache.payload) == 0 {
		return build()
	}
	var value T
	if err := json.Unmarshal(cache.payload, &value); err != nil {
		return build()
	}
	return value
}
