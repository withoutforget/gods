package main

import (
	"fmt"

	"github.com/withoutforget/gods/list"
)

func main() {
	fmt.Println("=== LinkedList Example ===")

	// Create linked list
	ll := list.NewLinkedList[string]()

	// Add to back
	fmt.Println("\nAdding elements...")
	ll.PushBack("Alice")
	ll.PushBack("Bob")
	ll.PushBack("Charlie")

	// Add to front
	ll.PushFront("Zero")

	fmt.Printf("Length: %d\n", ll.Len())

	// Access by index (optimized: searches from nearest end)
	fmt.Printf("\nElement at index 0: %s\n", *ll.At(0))
	fmt.Printf("Element at index 2: %s\n", *ll.At(2))

	// Iterate forward
	fmt.Print("\nForward iteration: ")
	for val := range ll.All() {
		fmt.Printf("%s ", *val)
	}
	fmt.Println()

	// Iterate with indices
	fmt.Println("\nWith indices:")
	for idx, val := range ll.AllIdx() {
		fmt.Printf("  [%d] = %s\n", idx, *val)
	}

	// Iterate backward
	fmt.Print("\nBackward iteration: ")
	for val := range ll.AllRev() {
		fmt.Printf("%s ", *val)
	}
	fmt.Println()

	// Work with nodes directly
	fmt.Println("\nDirect node access:")
	frontNode := ll.Front()
	fmt.Printf("Front node value: %s\n", *frontNode.GetValue())

	backNode := ll.Back()
	fmt.Printf("Back node value: %s\n", *backNode.GetValue())

	// Erase specific node
	secondNode := ll.AtNode(1)
	fmt.Printf("\nErasing node at index 1 (value: %s)\n", *secondNode.GetValue())
	ll.Erase(secondNode)
	fmt.Printf("Length after erase: %d\n", ll.Len())

	// Pop operations
	first := ll.PopFront()
	last := ll.PopBack()
	fmt.Printf("\nPopped front: %s\n", first)
	fmt.Printf("Popped back: %s\n", last)
	fmt.Printf("Final length: %d\n", ll.Len())

	// Check if empty
	fmt.Printf("Is empty: %v\n", ll.Empty())
}
