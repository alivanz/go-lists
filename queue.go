package lists

import "sync/atomic"

type Queue[T any] struct {
	head atomic.Value
	tail atomic.Value
}

func (queue *Queue[T]) Push(v T) {
	n := &node[T]{v: v}
	for {
		tail := queue.tail.Load()
		if queue.tail.CompareAndSwap(tail, n) {
			if tail == nil {
				queue.head.Store(n)
			} else {
				tail.(*node[T]).next.Store(n)
			}
			return
		}
	}
}

func (queue *Queue[T]) Pop(v *T) bool {
	for {
		h := queue.head.Load()
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
		if queue.head.CompareAndSwap(head, next) {
			*v = head.v
			return true
		}
	}
}
