package lists

import (
	"sync"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()
	var v int
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	if !queue.Pop(&v) {
		t.Fatal()
	} else if v != 1 {
		t.Fatal(v)
	}
	if !queue.Pop(&v) {
		t.Fatal()
	} else if v != 2 {
		t.Fatal(v)
	}
	if !queue.Pop(&v) {
		t.Fatal()
	} else if v != 3 {
		t.Fatal(v)
	}
	if queue.Pop(&v) {
		t.Fatal(v)
	}
}

func TestQueueRace(t *testing.T) {
	queue := NewQueue[int]()
	var v int
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		i := i
		go func() {
			defer wg.Done()
			queue.Push(i)
		}()
	}
	wg.Wait()
	var count int
	for {
		if queue.Pop(&v) {
			count += 1
		} else {
			break
		}
	}
	if count != 100 {
		t.Fatal(count)
	}
}
