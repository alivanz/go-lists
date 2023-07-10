package lists

type Stack[T any] struct {
	head Pointer[node[T]]
}

func (stack *Stack[T]) Push(v T) {
	node := &node[T]{
		v: v,
	}
	for {
		head := stack.head.Load()
		node.next.Store(head)
		if stack.head.CompareAndSwap(head, node) {
			return
		}
	}
}

func (stack *Stack[T]) Pop(v *T) bool {
	for {
		head := stack.head.Load()
		if head == nil {
			return false
		}
		next := head.next.Load()
		if stack.head.CompareAndSwap(head, next) {
			*v = head.v
			return true
		}
	}
}
