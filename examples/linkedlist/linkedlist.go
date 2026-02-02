package main

import (
	"fmt"

	"github.com/withoutforget/gods/list"
)

func main() {
	// Создаём двусвязный список
	ll := list.NewLinkedList[string]()

	// Добавляем в конец
	ll.PushBack("Alice")
	ll.PushBack("Bob")
	ll.PushBack("Charlie")

	// Добавляем в начало
	ll.PushFront("Zero")

	fmt.Printf("Length: %d\n", ll.Len())

	// Доступ по индексу (оптимизирован: с начала или конца)
	fmt.Printf("Element at index 1: %s\n", *ll.At(1))

	// Итерация вперёд
	fmt.Print("Forward: ")
	for val := range ll.All() {
		fmt.Printf("%s ", *val)
	}
	fmt.Println()

	// Итерация назад
	fmt.Print("Backward: ")
	for val := range ll.AllRev() {
		fmt.Printf("%s ", *val)
	}
	fmt.Println()

	// Работа с нодами напрямую
	node := ll.Front()
	fmt.Printf("Front node value: %s\n", *node.GetValue())

	// Удаление ноды
	secondNode := ll.AtNode(1)
	ll.Erase(secondNode)
	fmt.Printf("After erasing second node, length: %d\n", ll.Len())

	// Pop операции
	first := ll.PopFront()
	last := ll.PopBack()
	fmt.Printf("Popped front: %s, back: %s\n", first, last)
	fmt.Printf("Remaining length: %d\n", ll.Len())
}
