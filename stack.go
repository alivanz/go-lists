package lists

import "sync/atomic"

type Stack[T any] struct {
	head atomic.Value
}

func (stack *Stack[T]) Push(v T) {
	n := &node[T]{v: v}
	for {
		var head *node[T]
		h := stack.head.Load()
		if h != nil {
			head = h.(*node[T])
		}
		n.next.Store(head)
		if h == nil && stack.head.CompareAndSwap(nil, n) {
			return
		} else if stack.head.CompareAndSwap(head, n) {
			return
		}
	}
}

func (stack *Stack[T]) Pop(v *T) bool {
	for {
		h := stack.head.Load()
		if h == nil {
			return false
		}
		head := h.(*node[T])
		if head == nil {
			return false
		}
		n := head.next.Load()
		var next *node[T]
		if n != nil {
			next = n.(*node[T])
		}
		if stack.head.CompareAndSwap(head, next) {
			*v = head.v
			return true
		}
	}
}
