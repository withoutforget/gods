package main

import (
	"fmt"

	"github.com/withoutforget/gods/tree"
)

func main() {
	fmt.Println("=== Binary Search Tree Example ===")

	// Create BST with comparator (less function)
	bst := tree.NewBST(func(a, b int) bool {
		return a < b
	})

	// Insert elements
	fmt.Println("\nInserting: 50, 30, 70, 20, 40, 60, 80")
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, val := range values {
		bst.Insert(val)
	}

	// Search for elements
	fmt.Println("\nSearching for elements:")
	searchVals := []int{30, 45, 70, 100}
	for _, val := range searchVals {
		found := bst.Search(val)
		fmt.Printf("  Search %d: %v\n", val, found)
	}

	// In-order traversal (sorted order)
	fmt.Print("\nIn-order traversal (sorted): ")
	for val := range bst.All() {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Delete elements
	fmt.Println("\nDeleting 30...")
	bst.Delete(30)

	fmt.Print("After delete: ")
	for val := range bst.All() {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Delete more elements
	fmt.Println("\nDeleting 50 (root)...")
	bst.Delete(50)

	fmt.Print("After delete: ")
	for val := range bst.All() {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Example with strings
	fmt.Println("\n--- String BST Example ---")
	stringBST := tree.NewBST(func(a, b string) bool {
		return a < b
	})

	words := []string{"dog", "cat", "elephant", "ant", "bear", "fox"}
	fmt.Println("\nInserting words:", words)
	for _, word := range words {
		stringBST.Insert(word)
	}

	fmt.Print("Alphabetical order: ")
	for word := range stringBST.All() {
		fmt.Printf("%s ", word)
	}
	fmt.Println()

	// Example: Finding median
	fmt.Println("\n--- Finding Median Example ---")
	numbers := []int{5, 3, 8, 1, 9, 2, 7, 4, 6}
	fmt.Println("Numbers:", numbers)

	numBST := tree.NewBST(func(a, b int) bool {
		return a < b
	})

	for _, num := range numbers {
		numBST.Insert(num)
	}

	sorted := []int{}
	for num := range numBST.All() {
		sorted = append(sorted, num)
	}

	fmt.Println("Sorted:", sorted)
	median := sorted[len(sorted)/2]
	fmt.Printf("Median: %d\n", median)

	// Demonstrating BST property
	fmt.Println("\n--- BST Property Verification ---")
	testBST := tree.NewBST(func(a, b int) bool {
		return a < b
	})

	insertOrder := []int{50, 30, 70, 20, 40, 60, 80}
	fmt.Println("Insert order:", insertOrder)
	for _, val := range insertOrder {
		testBST.Insert(val)
	}

	fmt.Println("\nBST maintains sorted order on traversal:")
	fmt.Print("  ")
	for val := range testBST.All() {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
