package gsync

import "sync"

// Pool is a generic sync.Pool
type Pool[T any] struct {
	p   sync.Pool
	New func() T
}

// Get returns a value from the pool
func (p *Pool[T]) Get() T {
	if p.p.New == nil {
		p.p.New = func() any { return p.New() }
	}
	return p.p.Get().(T)
}

// Put returns a value to the pool
func (p *Pool[T]) Put(v T) { p.p.Put(v) }
