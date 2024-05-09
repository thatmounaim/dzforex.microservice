package storage

import (
	"errors"
	"sync"
)

type MemoryStorer struct {
	data map[string]float32
	mu   sync.RWMutex
}

func NewMemoryStore() *MemoryStorer {
	return &MemoryStorer{
		data: make(map[string]float32),
		mu:   sync.RWMutex{},
	}
}
func (s *MemoryStorer) UpdateData(d map[string]float32) {
	for k, v := range d {
		s.data[k] = v
	}
}

func (s *MemoryStorer) GetAll() map[string]float32 {
	return s.data
}

func (s *MemoryStorer) Get(k string) (float32, error) {
	d, ok := s.data[k]
	if !ok {
		return 0.0, errors.New("not found")
	}

	return d, nil
}
