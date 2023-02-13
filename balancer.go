package submitter

import "sync/atomic"

// Balancer ...
type Balancer[T any] struct {
	elements []T
	next     uint32
}

// NewBalancer ...
func NewBalancer[T any](elements ...T) *Balancer[T] {
	return &Balancer[T]{elements: elements}
}

// All ...
func (b *Balancer[T]) All() []T {
	return b.elements
}

// Next ...
func (b *Balancer[T]) Next() T {
	if len(b.elements) == 1 {
		return b.elements[0]
	}
	n := atomic.AddUint32(&b.next, 1)
	i := (int(n) - 1) % len(b.elements)
	return b.elements[i]
}
