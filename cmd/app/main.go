package main

import (
	"fmt"

	"github.com/withoutforget/gods/conts"
)

func main() {
	var h = conts.NewHeap[int](func(a, b int) bool { return a < b })
	for i := range 5 {
		h.Push(4 - i)
	}
	for range 5 {
		fmt.Println(h.RemoveAt(2))
	}
}
