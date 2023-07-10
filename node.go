package lists

import "sync/atomic"

type node[T any] struct {
	next atomic.Pointer[node[T]]
	v    T
}
