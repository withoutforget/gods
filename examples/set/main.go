package main

import (
	"fmt"

	"github.com/withoutforget/gods/hash"
)

func main() {
	fmt.Println("=== Set Example ===")

	// Create set
	set := hash.NewSet[int]()

	// Add elements
	fmt.Println("\nAdding elements: 1, 2, 3, 4, 5")
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)
	set.Add(5)

	// Add duplicate (will be ignored)
	fmt.Println("Adding duplicate: 3")
	set.Add(3)

	// Check membership
	fmt.Println("\nChecking membership:")
	for _, num := range []int{2, 6, 4, 10} {
		exists := set.Contains(num)
		fmt.Printf("  Contains %d: %v\n", num, exists)
	}

	// Delete element
	fmt.Println("\nDeleting 3...")
	set.Delete(3)
	fmt.Printf("Contains 3 after delete: %v\n", set.Contains(3))

	// Real-world examples
	fmt.Println("\n--- Set Operations Example ---")

	// Example 1: Unique visitors
	visitors := hash.NewSet[string]()
	pageViews := []string{"Alice", "Bob", "Alice", "Charlie", "Bob", "Alice", "Diana"}

	fmt.Println("\nPage views:", pageViews)
	for _, visitor := range pageViews {
		visitors.Add(visitor)
	}
	fmt.Printf("Total page views: %d\n", len(pageViews))
	fmt.Printf("Unique visitors: (tracked via set)\n")

	// Example 2: Finding unique numbers
	fmt.Println("\n--- Finding Unique Numbers ---")
	numbers := []int{1, 2, 3, 2, 4, 5, 1, 6, 7, 3, 8}
	uniqueNums := findUnique(numbers)

	fmt.Println("Original:", numbers)
	fmt.Println("Unique:", uniqueNums)

	// Example 3: Removing duplicates from strings
	fmt.Println("\n--- Remove Duplicate Words ---")
	sentence := []string{"the", "quick", "brown", "fox", "the", "lazy", "dog", "quick"}
	unique := removeDuplicates(sentence)

	fmt.Println("Original:", sentence)
	fmt.Println("Without duplicates:", unique)
}

func findUnique(numbers []int) []int {
	set := hash.NewSet[int]()
	result := []int{}

	for _, num := range numbers {
		if !set.Contains(num) {
			set.Add(num)
			result = append(result, num)
		}
	}

	return result
}

func removeDuplicates(words []string) []string {
	set := hash.NewSet[string]()
	result := []string{}

	for _, word := range words {
		if !set.Contains(word) {
			set.Add(word)
			result = append(result, word)
		}
	}

	return result
}
