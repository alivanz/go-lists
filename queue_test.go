package lists

import (
	"sync"
	"testing"
)

func TestQueue(t *testing.T) {
	var queue Queue[int]
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
	var queue Queue[int]
	var wg sync.WaitGroup
	checks := make([]bool, 100)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		i := i
		go func() {
			defer wg.Done()
			queue.Push(i)
		}()
	}
	wg.Wait()
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			var v int
			if queue.Pop(&v) {
				checks[v] = true
			} else {
				t.Fail()
			}
		}()
	}
	wg.Wait()
	for i, v := range checks {
		if !v {
			t.Fatal(i)
		}
	}
}
