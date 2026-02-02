package list

import (
	"iter"
)

type LinkedListNode[T any] struct {
	value T
	prev  *LinkedListNode[T]
	next  *LinkedListNode[T]
}

func (this *LinkedListNode[T]) GetValue() *T {
	return &this.value
}

type LinkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
	sz   int
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{}
}

func (this *LinkedList[T]) checkIdx(idx int) {
	if 0 <= idx && idx < this.sz {
		return
	}
	panic("LinkedList out of range")
}

func (this *LinkedList[T]) PushBack(value T) {
	var node = LinkedListNode[T]{value: value}
	if this.head == nil {
		this.head = &node
		this.tail = this.head
		this.sz += 1
		return
	}
	node.prev = this.tail
	this.tail.next = &node
	this.tail = &node
	this.sz += 1
}

func (this *LinkedList[T]) PushFront(value T) {
	var node = LinkedListNode[T]{value: value}
	if this.head == nil {
		this.head = &node
		this.tail = this.head
		this.sz += 1
		return
	}
	node.next = this.head
	this.head.prev = &node
	this.head = &node
	this.sz += 1
}

func (this *LinkedList[T]) PopBack() T {
	if this.head == this.tail {
		var tmp = this.tail.value
		this.head = nil
		this.tail = nil
		this.sz -= 1
		return tmp
	}
	var tmp = this.tail.value
	var prev = this.tail.prev
	prev.next = nil
	this.tail = prev
	this.sz -= 1
	return tmp
}

func (this *LinkedList[T]) PopFront() T {
	if this.head == this.tail {
		var tmp = this.head.value
		this.head = nil
		this.tail = nil
		this.sz -= 1

		return tmp
	}
	var tmp = this.head.value
	var next = this.head.next
	next.prev = nil
	this.head = next
	this.sz -= 1
	return tmp
}

func (this *LinkedList[T]) Len() int {
	return this.sz
}

func (this *LinkedList[T]) Empty() bool {
	return this.sz == 0
}

func (this *LinkedList[T]) At(idx int) *T {
	this.checkIdx(idx)
	return this.Get(idx)
}

// some optimzed for fastest access
func (this *LinkedList[T]) Get(idx int) *T {
	mid := this.sz / 2
	if idx < mid {
		curr := this.head
		for range idx {
			curr = curr.next
		}

		return &curr.value
	} else {
		curr := this.tail
		for v := this.sz - 1 - idx; v > 0; v-- {
			curr = curr.prev
		}

		return &curr.value
	}
}

func (this *LinkedList[T]) Erase(node *LinkedListNode[T]) {
	prev, next := node.prev, node.next
	if prev != nil {
		prev.next = next
	} else {
		this.head = next
	}
	if next != nil {
		next.prev = prev
	} else {
		this.tail = prev
	}
	node.prev = nil
	node.next = nil
}

func (this *LinkedList[T]) Front() *LinkedListNode[T] {
	return this.head
}
func (this *LinkedList[T]) Back() *LinkedListNode[T] {
	return this.tail
}

func (this *LinkedList[T]) AtNode(idx int) *LinkedListNode[T] {
	this.checkIdx(idx)
	return this.GetNode(idx)
}

func (this *LinkedList[T]) GetNode(idx int) *LinkedListNode[T] {
	mid := this.sz / 2
	if idx < mid {
		curr := this.head
		for range idx {
			curr = curr.next
		}

		return curr
	} else {
		curr := this.tail
		for v := this.sz - 1 - idx; v > 0; v-- {
			curr = curr.prev
		}

		return curr
	}
}

func (this *LinkedList[T]) All() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		curr := this.head
		for curr != nil {
			if !yield(&curr.value) {
				return
			}
			curr = curr.next
		}
	}
}

func (this *LinkedList[T]) AllIdx() iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		curr := this.head
		var i = 0
		for curr != nil {
			if !yield(i, &curr.value) {
				return
			}
			curr = curr.next
			i++
		}
	}
}

func (this *LinkedList[T]) AllRev() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		curr := this.tail
		for curr != nil {
			if !yield(&curr.value) {
				return
			}
			curr = curr.prev
		}
	}
}

func (this *LinkedList[T]) AllIdxRev() iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		curr := this.tail
		var i = this.sz - 1
		for curr != nil {
			if !yield(i, &curr.value) {
				return
			}
			curr = curr.prev
			i--
		}
	}
}
