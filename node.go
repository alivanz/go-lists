package lists

type node[T any] struct {
	next Pointer[node[T]]
	v    T
}
