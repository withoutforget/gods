package main

import (
	"fmt"

	"github.com/withoutforget/gods/conts"
)

func main() {
	c := conts.NewLRU[int, int](3)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	for i := range 3 {
		fmt.Println(c.Get(i + 1))
	}
	c.Put(4, 4)
	for k, v := range c.All() {
		fmt.Println(k, v)
	}
}
