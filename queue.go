package lists

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
}

func (queue *Queue[T]) Push(v T) {
	node := &node[T]{
		v: v,
	}
	for {
		tail := load(&queue.tail)
		if cas(&queue.tail, tail, node) {
			if tail == nil {
				store(&queue.head, node)
			} else {
				store(&tail.next, node)
			}
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
		if cas(&queue.head, head, next) {
			*v = head.v
			return true
		}
	}
}
