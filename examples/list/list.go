package main

import (
	"fmt"

	"github.com/withoutforget/gods/list"
)

func main() {
	// Создаём список
	list := list.NewList[int](0)

	// Добавляем элементы
	for i := 1; i <= 5; i++ {
		list.Append(i * 10)
	}
	fmt.Printf("Length: %d, Capacity: %d\n", list.Len(), list.Cap())

	// Прямой доступ по индексу
	fmt.Printf("Element at index 2: %d\n", *list.At(2))

	// Изменяем элемент
	*list.At(1) = 999

	// Итерация
	fmt.Print("All elements: ")
	for val := range list.All() {
		fmt.Printf("%d ", *val)
	}
	fmt.Println()

	// Итерация с индексами в обратном порядке
	fmt.Print("Reversed with indices: ")
	for idx, val := range list.AllIdxRev() {
		fmt.Printf("[%d]=%d ", idx, *val)
	}
	fmt.Println()

	// Удаление элемента
	list.Erase(2)
	fmt.Printf("After erasing index 2, length: %d\n", list.Len())

	// PopBack
	last := list.PopBack()
	fmt.Printf("Popped: %d, remaining length: %d\n", *last, list.Len())
}
