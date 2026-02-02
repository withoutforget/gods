package main

import (
	"fmt"

	"github.com/withoutforget/gods/list"
)

func main() {
	fmt.Println("=== List Example ===")

	// Create a list with initial capacity
	l := list.NewList[int](5)

	// Add elements
	fmt.Println("\nAdding elements...")
	for i := 1; i <= 10; i++ {
		l.Append(i * 10)
	}
	fmt.Printf("Length: %d, Capacity: %d\n", l.Len(), l.Cap())

	// Access by index
	fmt.Printf("\nElement at index 3: %d\n", *l.At(3))

	// Modify element
	*l.At(0) = 999
	fmt.Printf("Modified first element: %d\n", *l.At(0))

	// Iterate forward
	fmt.Print("\nAll elements (forward): ")
	for val := range l.All() {
		fmt.Printf("%d ", *val)
	}
	fmt.Println()

	// Iterate with indices
	fmt.Println("\nElements with indices:")
	for idx, val := range l.AllIdx() {
		fmt.Printf("  [%d] = %d\n", idx, *val)
	}

	// Iterate backward
	fmt.Print("\nAll elements (backward): ")
	for val := range l.AllRev() {
		fmt.Printf("%d ", *val)
	}
	fmt.Println()

	// Remove element
	fmt.Println("\nRemoving element at index 5...")
	l.Erase(5)
	fmt.Printf("Length after erase: %d\n", l.Len())

	// Pop from back
	last := l.PopBack()
	fmt.Printf("Popped from back: %d\n", *last)
	fmt.Printf("Final length: %d\n", l.Len())

	// Check if empty
	fmt.Printf("Is empty: %v\n", l.Empty())

	// Shrink to fit
	l.ShrinkToFit()
	fmt.Printf("After shrink - Length: %d, Capacity: %d\n", l.Len(), l.Cap())
}
