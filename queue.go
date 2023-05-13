package lists

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
}

func (queue *Queue[T]) Push(v T) {
	n := &node[T]{
		v: v,
	}
	for {
		tail := load(&queue.tail)
		if tail == nil {
			tail = new(node[T])
			if cas(&queue.tail, nil, tail) {
				store(&queue.head, tail)
			} else {
				tail = load(&queue.tail)
			}
		}
		if cas(&queue.tail, tail, n) {
			store(&tail.next, n)
			return
		}
	}
}

func (queue *Queue[T]) Pop(v *T) bool {
	for {
		head := load(&queue.head)
		if head == nil {
			return false
		}
		next := load(&head.next)
		if next == nil {
			return false
		}
		if cas(&queue.head, head, next) {
			*v = next.v
			return true
		}
	}
}
