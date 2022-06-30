[![Go Report Card](https://goreportcard.com/badge/github.com/alivanz/go-lists)](https://goreportcard.com/report/github.com/alivanz/go-lists)
[![Build Status](https://github.com/alivanz/go-lists/actions/workflows/test.yml/badge.svg)](https://github.com/alivanz/go-lists/actions)

# Introduction

lock-free / mutex-free and thread-safe lists written in Go

Supports 1.18 generics

# Example

## Queue

```go
func main() {
    var queue lists.Queue[int]
    // producer
    for i:=0; i<100; i++ {
        queue.Push(i)
    }
    // consumer
    var v int
    for queue.Pop(&v) {
        fmt.Println(v)
    }
    // 0
    // 1
    // 2
    // ...
    // 98
    // 99
}
```

## Stack
```go
func main() {
    var stack lists.Stack[int]
    // producer
    for i:=0; i<100; i++ {
        stack.Push(i)
    }
    // consumer
    var v int
    for stack.Pop(&v) {
        fmt.Println(v)
    }
    // 99
    // 98
    // ...
    // 1
    // 0
}
```