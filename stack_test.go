package lists

import (
	"sync"
	"testing"
)

func TestStack(t *testing.T) {
	var stack Stack[int]
	var v int
	stack.Unpop(1)
	stack.Unpop(2)
	stack.Unpop(3)
	if !stack.Pop(&v) {
		t.Fatal()
	} else if v != 3 {
		t.Fatal(v)
	}
	if !stack.Pop(&v) {
		t.Fatal()
	} else if v != 2 {
		t.Fatal(v)
	}
	if !stack.Pop(&v) {
		t.Fatal()
	} else if v != 1 {
		t.Fatal(v)
	}
	if stack.Pop(&v) {
		t.Fatal(v)
	}
}

func TestStackRace(t *testing.T) {
	var wg sync.WaitGroup
	var stack Stack[int]
	var v int
	wg.Add(100)
	for i := 0; i < 100; i++ {
		i := i
		go func() {
			defer wg.Done()
			stack.Unpop(i)
		}()
	}
	wg.Wait()
	var count int
	for {
		if stack.Pop(&v) {
			count += 1
		} else {
			break
		}
	}
}
