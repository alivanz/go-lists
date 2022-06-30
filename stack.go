package lists

type Stack[T any] struct {
	head *node[T]
}

func (stack *Stack[T]) Push(v T) {
	node := &node[T]{
		v: v,
	}
	for {
		head := load(&stack.head)
		node.next = head
		if cas(&stack.head, head, node) {
			return
		}
	}
}

func (stack *Stack[T]) Pop(v *T) bool {
	for {
		head := load(&stack.head)
		if head == nil {
			return false
		}
		next := load(&head.next)
		if cas(&stack.head, head, next) {
			*v = head.v
			return true
		}
	}
}
