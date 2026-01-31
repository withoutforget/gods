package conts

import "github.com/withoutforget/gods/clibutils"

type Heap[T comparable] struct {
	data List[T]
	comp func(a, b T) bool // less
}

func NewHeap[T comparable](comparator func(a, b T) bool) Heap[T] {
	return Heap[T]{comp: comparator}
}

func (this *Heap[T]) Len() int {
	return this.data.Len()
}

// count returns 2 if two children
// return 1 if only left exist
// return 0 if both doesn't exist
func (this *Heap[T]) getChildren(idx int) (Left int, Right int, Count int) {
	left := 2*idx + 1
	right := left + 1
	var count = 0
	if left < this.Len() {
		count++
	}
	if right < this.Len() {
		count++
	}
	return left, right, count
}

func (this *Heap[T]) getParent(idx int) (Idx int, Result bool) {
	if idx == 0 {
		return 0, false
	}
	return (idx - 1) / 2, true
}
func (this *Heap[T]) siftDown(idx int) {
	l, r, c := this.getChildren(idx)
	if c == 0 {
		return
	}

	parent := this.data.Get(idx)
	var targetChild int

	if c == 1 {
		targetChild = l
	} else {
		lChild, rChild := this.data.Get(l), this.data.Get(r)
		if this.comp(*lChild, *rChild) {
			targetChild = l
		} else {
			targetChild = r
		}
	}

	childPtr := this.data.Get(targetChild)
	if this.comp(*childPtr, *parent) {
		*parent, *childPtr = clibutils.Swap(*parent, *childPtr)
		this.siftDown(targetChild)
	}
}

func (this *Heap[T]) siftUP(idx int) {
	p, res := this.getParent(idx)
	if !res {
		return
	}
	a, b := this.data.Get(idx), this.data.Get(p)
	if this.comp(*a, *b) {
		*a, *b = clibutils.Swap(*a, *b)
		this.siftUP(p)
	}
}

func (this *Heap[T]) Push(value T) {
	this.data.Append(value)
	this.siftUP(this.data.Len() - 1)
}

func (this *Heap[T]) RemoveAt(idx int) T {
	bIdx, eIdx := idx, this.data.Len()-1
	if bIdx > eIdx { // maybe we shouldn't panic?
		panic("Heap out of range")
	}
	if bIdx == eIdx {
		value := *this.data.PopBack()
		return value
	}
	a, b := this.data.Get(bIdx), this.data.Get(eIdx)
	*a, *b = clibutils.Swap(*a, *b)
	value := *this.data.PopBack()
	this.siftDown(0)
	return value
}

func (this *Heap[T]) Pop() T {
	return this.RemoveAt(0)
}

func (this *Heap[T]) Empty() bool {
	return this.data.Empty()
}

func (this *Heap[T]) Top() T {
	return *this.data.Get(this.data.Len() - 1)
}
