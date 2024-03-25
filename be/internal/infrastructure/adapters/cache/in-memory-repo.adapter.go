package cache

import (
	"sync"

	"github.com/PaoloEG/terrasense/internal/core/domain/errors"
)

type InMemoryRepo[T any] struct {
	data map[string]T
	mu   sync.RWMutex
}

func NewInMemoryRepo[T any](initData map[string]T) *InMemoryRepo[T] {
	return &InMemoryRepo[T]{
		data: initData,
	}
}

func (r *InMemoryRepo[T]) Save(id string, entity T) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[id] = entity
	return nil
}

func (r *InMemoryRepo[T]) Get(id string) (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if value, ok := r.data[id]; ok {
		return value, nil
	} else {
		return value, &errors.NotFoundError{Message: "Key " + id + " not found"}
	}
}

// func (r *InMemoryRepo[T]) Update(id string, entity T) error {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()
// 	r.data[id] = entity
// 	return nil
// }

// func (r *InMemoryRepo[T]) Delete(id string) error {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()
// 	delete(r.data, id)
// 	return nil
// }
