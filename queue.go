package lists

import "sync/atomic"

type Queue[T any] struct {
	head atomic.Pointer[node[T]]
	tail atomic.Pointer[node[T]]
}

func (queue *Queue[T]) Push(v T) {
	n := &node[T]{
		v: v,
	}
	for {
		tail := queue.tail.Load()
		if tail == nil {
			tail = new(node[T])
			if queue.tail.CompareAndSwap(nil, tail) {
				queue.head.Store(tail)
			} else {
				tail = queue.tail.Load()
			}
		}
		if queue.tail.CompareAndSwap(tail, n) {
			tail.next.Store(n)
			return
		}
	}
}

func (queue *Queue[T]) Pop(v *T) bool {
	for {
		head := queue.head.Load()
		if head == nil {
			return false
		}
		next := head.next.Load()
		if next == nil {
			return false
		}
		if queue.head.CompareAndSwap(head, next) {
			*v = next.v
			return true
		}
	}
}
