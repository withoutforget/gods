package main

import (
	"fmt"

	"github.com/withoutforget/gods/list"
)

func main() {
	fmt.Println("=== Stack Example (LIFO) ===")

	// Create stack
	stack := list.NewStack[int]()

	// Push elements
	fmt.Println("\nPushing elements: 10, 20, 30, 40, 50")
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	fmt.Printf("Stack length: %d\n", stack.Len())
	fmt.Printf("Is empty: %v\n", stack.Empty())

	// Pop elements (LIFO - Last In First Out)
	fmt.Println("\nPopping elements:")
	for !stack.Empty() {
		val := stack.Pop()
		fmt.Printf("  Popped: %d, Remaining: %d\n", val, stack.Len())
	}

	fmt.Printf("\nIs empty now: %v\n", stack.Empty())

	// Real-world example: balanced parentheses checker
	fmt.Println("\n--- Balanced Parentheses Checker ---")
	testCases := []string{
		"(())",
		"(()",
		"()()",
		"((()))",
		")(",
	}

	for _, test := range testCases {
		result := checkBalanced(test)
		fmt.Printf("%s -> %v\n", test, result)
	}
}

func checkBalanced(s string) bool {
	stack := list.NewStack[rune]()

	for _, ch := range s {
		if ch == '(' {
			stack.Push(ch)
		} else if ch == ')' {
			if stack.Empty() {
				return false
			}
			stack.Pop()
		}
	}

	return stack.Empty()
}
