package inmemdb

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type db[T any] struct {
	data map[uuid.UUID]T
	mu   sync.RWMutex
}

func NewDB[T any]() *db[T] {
	return &db[T]{
		data: make(map[uuid.UUID]T, 2048),
	}
}

func (d *db[T]) Save(entity T) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	id := uuid.New()
	d.data[id] = entity
	return nil
}

func (d *db[T]) Get(id string) (T, error) {
	cid, err := uuid.Parse(id)
	if err != nil {
		var zero T
		return zero, err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()

	value, ok := d.data[cid]
	if !ok {
		var zero T
		return zero, ErrNotFound
	}
	return value, nil
}

var ErrNotFound = errors.New("not found")

func (d *db[T]) Delete(id string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	delete(d.data, uuid)
	return nil
}
