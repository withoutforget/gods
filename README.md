# gods - Generic Data Structures for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/withoutforget/gods.svg)](https://pkg.go.dev/github.com/withoutforget/gods)

Pure Go implementation of common data structures using generics (Go 1.18+).

## Features

- 🚀 **Generic implementations** - Type-safe data structures using Go generics
- 📦 **Zero dependencies** - Pure Go, no external dependencies
- 🔄 **Iterator support** - All structures support Go 1.23+ iterators
- 🎯 **Simple API** - Clean and intuitive interfaces*
- ⚡ **Efficient** - Optimized implementations with good performance characteristics**

\*, \*\* I'm not sure if it's true

## Installation

```bash
go get github.com/withoutforget/gods
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/withoutforget/gods/list"
)

func main() {
    // Create a dynamic list
    l := list.NewList[int](0)
    l.Append(10)
    l.Append(20)
    l.Append(30)
    
    // Iterate using Go 1.23 iterators
    for val := range l.All() {
        fmt.Println(*val)
    }
}
```

## Data Structures

### Linear Structures (`list/`)

| Structure | Description | Key Operations |
|-----------|-------------|----------------|
| **List** | Dynamic array with automatic growth | `Append`, `At`, `Get`, `Erase` - O(1) amortized |
| **LinkedList** | Doubly-linked list | `PushBack`, `PushFront`, `PopBack`, `PopFront` - O(1) |
| **Stack** | LIFO stack backed by List | `Push`, `Pop` - O(1) amortized |
| **Queue** | FIFO queue backed by LinkedList | `Push`, `Pop` - O(1) |

### Hash Structures (`hash/`)

| Structure | Description | Key Operations |
|-----------|-------------|----------------|
| **HashMap** | Hash table with open addressing | `Set`, `Get`, `Delete` - O(1) average |
| **Set** | Unordered unique elements | `Add`, `Contains`, `Delete` - O(1) average |
| **LRU** | Least Recently Used cache | `Put`, `Get` - O(1) average |

### Tree Structures (`tree/`)

| Structure | Description | Key Operations |
|-----------|-------------|----------------|
| **BST** | Binary search tree with custom comparator | `Insert`, `Search`, `Delete` - O(log n) average |
| **Heap** | Binary min/max heap with custom comparator | `Push`, `Pop`, `Top` - O(log n) |

## Usage Examples

### List (Dynamic Array)

```go
import "github.com/withoutforget/gods/list"

// Create list
l := list.NewList[int](10) // initial capacity 10
l.Append(1)
l.Append(2)
l.Append(3)

// Access elements
val := l.At(0)    // returns *int
*val = 100        // modify in place

// Iterate
for val := range l.All() {
    fmt.Println(*val)
}

// With indices
for idx, val := range l.AllIdx() {
    fmt.Printf("[%d] = %d\n", idx, *val)
}

// Remove element
l.Erase(1)       // removes element at index 1
last := l.PopBack() // removes and returns last element
```

### LinkedList (Doubly-Linked)

```go
import "github.com/withoutforget/gods/list"

ll := list.NewLinkedList[string]()

// Add elements
ll.PushBack("world")
ll.PushFront("hello")  // ["hello", "world"]

// Access
first := ll.At(0)      // O(n) access by index
node := ll.Front()     // O(1) get first node
value := node.GetValue()

// Iterate forward
for val := range ll.All() {
    fmt.Println(*val)
}

// Iterate backward
for val := range ll.AllRev() {
    fmt.Println(*val)
}

// Remove
ll.PopFront()
ll.PopBack()
```

### Stack (LIFO)

```go
import "github.com/withoutforget/gods/list"

stack := list.NewStack[int]()

stack.Push(1)
stack.Push(2)
stack.Push(3)

top := stack.Pop()  // returns 3
fmt.Println(top)
fmt.Println(stack.Len())    // 2
fmt.Println(stack.Empty())  // false
```

### Queue (FIFO)

```go
import "github.com/withoutforget/gods/list"

queue := list.NewQueue[string]()

queue.Push("first")
queue.Push("second")
queue.Push("third")

first := queue.Pop()  // returns "first"
fmt.Println(first)
fmt.Println(queue.Len())    // 2
fmt.Println(queue.Empty())  // false
```

### HashMap

```go
import "github.com/withoutforget/gods/hash"

m := hash.NewHashMap[string, int]()

// Set values
m.Set("alice", 30)
m.Set("bob", 25)
m.Set("charlie", 35)

// Get value
age := m.Get("alice")
if age != nil {
    fmt.Println(*age)  // 30
}

// Delete
m.Delete("bob")

// Iterate over keys
for key := range m.Keys() {
    fmt.Println(key)
}

// Iterate over key-value pairs
for key, val := range m.All() {
    fmt.Printf("%s: %d\n", key, *val)
}

fmt.Println(m.Len())  // 2
```

### Set

```go
import "github.com/withoutforget/gods/hash"

set := hash.NewSet[int]()

set.Add(1)
set.Add(2)
set.Add(3)
set.Add(2)  // duplicate, ignored

fmt.Println(set.Contains(2))  // true
fmt.Println(set.Contains(5))  // false

set.Delete(2)
fmt.Println(set.Contains(2))  // false
```

### LRU Cache

```go
import "github.com/withoutforget/gods/hash"

cache := hash.NewLRU[string, int](3)  // capacity 3

cache.Put("a", 1)
cache.Put("b", 2)
cache.Put("c", 3)
cache.Put("d", 4)  // evicts "a" (least recently used)

val := cache.Get("b")
if val != nil {
    fmt.Println(*val)  // 2
}

// Iterate (most recent first)
for key, val := range cache.All() {
    fmt.Printf("%s: %d\n", key, val)
}
```

### Binary Search Tree

```go
import "github.com/withoutforget/gods/tree"

// Create BST with custom comparator (less function)
bst := tree.NewBST(func(a, b int) bool { 
    return a < b 
})

bst.Insert(5)
bst.Insert(3)
bst.Insert(7)
bst.Insert(1)
bst.Insert(9)

// Search
found := bst.Search(7)  // true
fmt.Println(found)

// Delete
bst.Delete(3)

// Iterate in-order (sorted)
for val := range bst.All() {
    fmt.Println(val)
}
```

### Heap (Priority Queue)

```go
import "github.com/withoutforget/gods/tree"

// Min heap
minHeap := tree.NewHeap(func(a, b int) bool { 
    return a < b 
})

minHeap.Push(5)
minHeap.Push(2)
minHeap.Push(8)
minHeap.Push(1)

min := minHeap.Pop()  // returns 1
fmt.Println(min)

top := minHeap.Top()  // peek at top (2)
fmt.Println(top)

// Max heap
maxHeap := tree.NewHeap(func(a, b int) bool { 
    return a > b 
})

maxHeap.Push(5)
maxHeap.Push(2)
maxHeap.Push(8)

max := maxHeap.Pop()  // returns 8
```

## Iterator Support

All data structures support Go 1.23+ iterators:

```go
// Single value iterator
for val := range list.All() {
    fmt.Println(*val)
}

// Index and value iterator
for idx, val := range list.AllIdx() {
    fmt.Printf("[%d] = %d\n", idx, *val)
}

// Reverse iteration
for val := range list.AllRev() {
    fmt.Println(*val)
}

// Key-value iterator (maps)
for key, val := range hashMap.All() {
    fmt.Printf("%s: %d\n", key, *val)
}
```

## Examples

See the [examples/](./examples) directory for complete, runnable examples:

- [List](./examples/list/main.go)
- [LinkedList](./examples/linkedlist/main.go)
- [Stack](./examples/stack/main.go)
- [Queue](./examples/queue/main.go)
- [HashMap](./examples/hashmap/main.go)
- [Set](./examples/set/main.go)
- [LRU](./examples/lru/main.go)
- [BST](./examples/bst/main.go)
- [Heap](./examples/heap/main.go)

Run any example:

```bash
go run examples/list/main.go
go run examples/hashmap/main.go
```

## Documentation

Full documentation available at [pkg.go.dev/github.com/withoutforget/gods](https://pkg.go.dev/github.com/withoutforget/gods)

Use `go doc` to view documentation locally:

```bash
go doc github.com/withoutforget/gods/list
go doc github.com/withoutforget/gods/list.List
go doc github.com/withoutforget/gods/hash.HashMap
```

## Performance Characteristics

| Operation | List | LinkedList | HashMap | BST | Heap |
|-----------|------|------------|---------|-----|------|
| Insert/Add | O(1)* | O(1) | O(1)* | O(log n)* | O(log n) |
| Delete | O(n) | O(1)** | O(1)* | O(log n)* | O(log n) |
| Search/Get | O(1) | O(n) | O(1)* | O(log n)* | O(1)*** |
| Iterate | O(n) | O(n) | O(n) | O(n) | - |

\* Amortized
\*\* If node is known
\*\*\* Top element only (peek)

<small>readme has been made by ai</small>