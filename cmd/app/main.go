package main

import (
	"fmt"

	"github.com/withoutforget/gods/conts"
)

func main() {
	b := conts.NewBST(func(a, b int) bool { return a < b })
	for i := range 10 {
		b.Insert(i)
	}
	for i := range 20 {
		fmt.Println(b.Search(i))
	}
}
