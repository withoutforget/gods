package main

import (
	"fmt"

	"github.com/withoutforget/gods/list"
)

func main() {
	fmt.Println("=== Queue Example (FIFO) ===")

	// Create queue
	queue := list.NewQueue[string]()

	// Push elements
	fmt.Println("\nEnqueuing elements...")
	queue.Push("First")
	queue.Push("Second")
	queue.Push("Third")
	queue.Push("Fourth")
	queue.Push("Fifth")

	fmt.Printf("Queue length: %d\n", queue.Len())
	fmt.Printf("Is empty: %v\n", queue.Empty())

	// Pop elements (FIFO - First In First Out)
	fmt.Println("\nDequeuing elements:")
	for !queue.Empty() {
		val := queue.Pop()
		fmt.Printf("  Dequeued: %s, Remaining: %d\n", val, queue.Len())
	}

	fmt.Printf("\nIs empty now: %v\n", queue.Empty())

	// Real-world example: task processor simulation
	fmt.Println("\n--- Task Processor Simulation ---")
	taskQueue := list.NewQueue[Task]()

	// Add tasks
	tasks := []Task{
		{ID: 1, Name: "Send email"},
		{ID: 2, Name: "Process payment"},
		{ID: 3, Name: "Generate report"},
		{ID: 4, Name: "Update database"},
	}

	fmt.Println("\nAdding tasks to queue:")
	for _, task := range tasks {
		taskQueue.Push(task)
		fmt.Printf("  Added: Task #%d - %s\n", task.ID, task.Name)
	}

	fmt.Println("\nProcessing tasks in order:")
	for !taskQueue.Empty() {
		task := taskQueue.Pop()
		fmt.Printf("  Processing: Task #%d - %s\n", task.ID, task.Name)
	}
}

type Task struct {
	ID   int
	Name string
}
