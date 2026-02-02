package main

import (
	"fmt"

	"github.com/withoutforget/gods/hash"
)

func main() {
	fmt.Println("=== LRU Cache Example ===")

	// Create LRU cache with capacity 3
	cache := hash.NewLRU[string, int](3)

	fmt.Println("\nCapacity: 3")
	fmt.Println("\nAdding items...")

	// Add items
	cache.Put("a", 1)
	fmt.Println("Put: a -> 1")
	printCache(cache)

	cache.Put("b", 2)
	fmt.Println("\nPut: b -> 2")
	printCache(cache)

	cache.Put("c", 3)
	fmt.Println("\nPut: c -> 3")
	printCache(cache)

	// Cache is full, next insert will evict least recently used
	cache.Put("d", 4)
	fmt.Println("\nPut: d -> 4 (evicts 'a' - least recently used)")
	printCache(cache)

	// Get makes an item recently used
	fmt.Println("\nGet: b")
	val := cache.Get("b")
	if val != nil {
		fmt.Printf("  Value: %d\n", *val)
	}
	printCache(cache)

	// Now 'c' is least recently used
	cache.Put("e", 5)
	fmt.Println("\nPut: e -> 5 (evicts 'c' - least recently used)")
	printCache(cache)

	// Try to get evicted item
	fmt.Println("\nTrying to get 'a' (was evicted):")
	valA := cache.Get("a")
	if valA == nil {
		fmt.Println("  Not found (evicted)")
	}

	// Real-world example: Page cache
	fmt.Println("\n--- Web Page Cache Simulation ---")
	pageCache := hash.NewLRU[string, string](3)

	pages := []string{
		"/home",
		"/about",
		"/contact",
		"/home",     // access again
		"/products", // evicts /about
		"/contact",  // access again
		"/blog",     // evicts /home
	}

	for i, page := range pages {
		fmt.Printf("\nStep %d: Accessing %s\n", i+1, page)

		cached := pageCache.Get(page)
		if cached != nil {
			fmt.Printf("  ✓ Cache hit: %s\n", *cached)
		} else {
			fmt.Printf("  ✗ Cache miss, loading from disk...\n")
			content := fmt.Sprintf("Content of %s", page)
			pageCache.Put(page, content)
			fmt.Printf("  Cached: %s\n", content)
		}

		fmt.Print("  Current cache: ")
		for key := range pageCache.All() {
			fmt.Printf("%s ", key)
		}
		fmt.Println()
	}
}

func printCache(cache *hash.LRU[string, int]) {
	fmt.Print("  Cache (MRU -> LRU): ")
	for key, val := range cache.All() {
		fmt.Printf("%s:%d ", key, val)
	}
	fmt.Println()
}
