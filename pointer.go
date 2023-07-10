//go:build go1.8

package lists

import "sync/atomic"

// equivalent to atomic.Pointer
type Pointer[T any] struct {
	data atomic.Value
}

func (ptr *Pointer[T]) Load() *T {
	v, ok := ptr.data.Load().(*T)
	if !ok {
		return nil
	}
	return v
}

func (ptr *Pointer[T]) Store(v *T) {
	ptr.data.Store(v)
}

func (ptr *Pointer[T]) CompareAndSwap(old, new *T) bool {
	return ptr.data.CompareAndSwap(old, new)
}
