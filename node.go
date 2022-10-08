package lists

type node[T any] struct {
	next atomic.Value
	v    T
}
