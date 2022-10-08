package lists

import "sync/atomic"

type node[T any] struct {
	next atomic.Value
	v    T
}
