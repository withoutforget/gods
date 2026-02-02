package main

import (
	"fmt"

	"github.com/withoutforget/gods/tree"
)

func main() {
	fmt.Println("=== Heap Example ===")

	// Create min heap
	fmt.Println("\n--- Min Heap ---")
	minHeap := tree.NewHeap(func(a, b int) bool {
		return a < b // less means min heap
	})

	// Push elements
	fmt.Println("\nPushing: 5, 3, 8, 1, 9, 2, 7")
	values := []int{5, 3, 8, 1, 9, 2, 7}
	for _, val := range values {
		minHeap.Push(val)
	}

	fmt.Printf("Heap size: %d\n", minHeap.Len())

	// Pop all elements (will come out in sorted order)
	fmt.Print("\nPopping all (ascending): ")
	for !minHeap.Empty() {
		min := minHeap.Pop()
		fmt.Printf("%d ", min)
	}
	fmt.Println()

	// Create max heap
	fmt.Println("\n--- Max Heap ---")
	maxHeap := tree.NewHeap(func(a, b int) bool {
		return a > b // greater means max heap
	})

	fmt.Println("\nPushing: 5, 3, 8, 1, 9, 2, 7")
	for _, val := range values {
		maxHeap.Push(val)
	}

	// Pop all elements (will come out in reverse sorted order)
	fmt.Print("\nPopping all (descending): ")
	for !maxHeap.Empty() {
		max := maxHeap.Pop()
		fmt.Printf("%d ", max)
	}
	fmt.Println()

	// Top element (peek without removing)
	fmt.Println("\n--- Top Element (Peek) ---")
	peekHeap := tree.NewHeap(func(a, b int) bool {
		return a < b
	})

	peekHeap.Push(10)
	peekHeap.Push(5)
	peekHeap.Push(15)

	fmt.Printf("Top element: %d\n", peekHeap.Top())
	fmt.Printf("Heap size: %d\n", peekHeap.Len())
	fmt.Println("(Top doesn't remove the element)")

	// Real-world example: Priority Queue for tasks
	fmt.Println("\n--- Priority Queue Example ---")

	type Task struct {
		Name     string
		Priority int
	}

	// Min heap for priority (lower number = higher priority)
	taskQueue := tree.NewHeap(func(a, b Task) bool {
		return a.Priority < b.Priority
	})

	tasks := []Task{
		{"Send email", 3},
		{"Fix critical bug", 1},
		{"Update docs", 5},
		{"Review PR", 2},
		{"Deploy to prod", 1},
	}

	fmt.Println("\nAdding tasks:")
	for _, task := range tasks {
		fmt.Printf("  %s (priority %d)\n", task.Name, task.Priority)
		taskQueue.Push(task)
	}

	fmt.Println("\nProcessing tasks by priority:")
	for !taskQueue.Empty() {
		task := taskQueue.Pop()
		fmt.Printf("  [Priority %d] %s\n", task.Priority, task.Name)
	}

	// Finding K largest elements
	fmt.Println("\n--- Finding K Largest Elements ---")
	numbers := []int{3, 1, 5, 12, 2, 11, 7, 8, 4, 6}
	k := 3

	fmt.Printf("Array: %v\n", numbers)
	fmt.Printf("Find %d largest elements\n", k)

	// Use min heap of size k
	kLargest := tree.NewHeap(func(a, b int) bool {
		return a < b
	})

	for _, num := range numbers {
		kLargest.Push(num)
		if kLargest.Len() > k {
			kLargest.Pop() // remove smallest
		}
	}

	result := []int{}
	for !kLargest.Empty() {
		result = append([]int{kLargest.Pop()}, result...)
	}

	fmt.Printf("%d largest elements: %v\n", k, result)

	// Heap sort demonstration
	fmt.Println("\n--- Heap Sort ---")
	unsorted := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Unsorted: %v\n", unsorted)

	sorted := heapSort(unsorted)
	fmt.Printf("Sorted: %v\n", sorted)
}

func heapSort(arr []int) []int {
	heap := tree.NewHeap(func(a, b int) bool {
		return a < b
	})

	// Push all elements
	for _, val := range arr {
		heap.Push(val)
	}

	// Pop all elements (comes out sorted)
	result := make([]int, 0, len(arr))
	for !heap.Empty() {
		result = append(result, heap.Pop())
	}

	return result
}
