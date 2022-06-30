package lists

type node[T any] struct {
	next *node[T]
	v    T
}
