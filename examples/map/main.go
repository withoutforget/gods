package main

import (
	"fmt"

	"github.com/withoutforget/gods/hash"
)

func main() {
	fmt.Println("=== HashMap Example ===")

	// Create hash map
	m := hash.NewHashMap[string, int]()

	// Set key-value pairs
	fmt.Println("\nAdding key-value pairs...")
	m.Set("Alice", 30)
	m.Set("Bob", 25)
	m.Set("Charlie", 35)
	m.Set("Diana", 28)
	m.Set("Eve", 32)

	fmt.Printf("Map size: %d\n", m.Len())
	fmt.Printf("Load factor: %.2f\n", m.LoadFactor())

	// Get values
	fmt.Println("\nGetting values:")
	names := []string{"Alice", "Bob", "Frank"}
	for _, name := range names {
		age := m.Get(name)
		if age != nil {
			fmt.Printf("  %s: %d years old\n", name, *age)
		} else {
			fmt.Printf("  %s: not found\n", name)
		}
	}

	// Update existing value
	fmt.Println("\nUpdating Alice's age to 31...")
	m.Set("Alice", 31)
	fmt.Printf("Alice: %d\n", *m.Get("Alice"))

	// Iterate over keys
	fmt.Println("\nAll keys:")
	for key := range m.Keys() {
		fmt.Printf("  - %s\n", key)
	}

	// Iterate over key-value pairs
	fmt.Println("\nAll key-value pairs:")
	for key, val := range m.All() {
		fmt.Printf("  %s: %d\n", key, *val)
	}

	// Delete entry
	fmt.Println("\nDeleting Bob...")
	m.Delete("Bob")
	fmt.Printf("Map size after delete: %d\n", m.Len())

	// Verify deletion
	bob := m.Get("Bob")
	fmt.Printf("Bob exists: %v\n", bob != nil)

	// Real-world example: word frequency counter
	fmt.Println("\n--- Word Frequency Counter ---")
	text := "the quick brown fox jumps over the lazy dog the fox was quick"
	wordCount := countWords(text)

	fmt.Println("\nWord frequencies:")
	for word, count := range wordCount.All() {
		fmt.Printf("  '%s': %d\n", word, *count)
	}
}

func countWords(text string) hash.HashMap[string, int] {
	wordCount := hash.NewHashMap[string, int]()

	// Simple word splitting (spaces)
	word := ""
	for _, ch := range text + " " {
		if ch == ' ' {
			if word != "" {
				count := wordCount.Get(word)
				if count == nil {
					wordCount.Set(word, 1)
				} else {
					wordCount.Set(word, *count+1)
				}
				word = ""
			}
		} else {
			word += string(ch)
		}
	}

	return wordCount
}
