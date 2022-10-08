package lists

type Queue[T any] struct {
	head atomic.Value
	tail atomic.Value
}

func (queue *Queue[T]) Push(v T) {
	node := &node[T]{v: v}
	for {
		tail := queue.tail.Load().(node[T])
		if queue.tail.CompareAndSwap(tail, node) {
			if tail == nil {
				queue.head.Store(node)
			} else {
				tail.next.Store(node)
			}
			return
		}
	}
}

func (queue *Queue[T]) Pop(v *T) bool {
	for {
		head := queue.head.Load().(node[T])
		if head == nil {
			return false
		}
		next := head.next.Load().(node[T])
		if queue.head.CompareAndSwap(head, next) {
			*v = head.v
			return true
		}
	}
}
