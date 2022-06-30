package lists

import (
	"sync/atomic"
	"unsafe"
)

func cas[T any](p **T, old, new *T) bool {
	return atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(p)), unsafe.Pointer(old), unsafe.Pointer(new))
}

func load[T any](p **T) *T {
	return (*T)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(p))))
}
