package conts

import (
	"iter"
	"unsafe"

	"github.com/withoutforget/gods/clibutils"
)

type List[T any] struct {
	data unsafe.Pointer
	sz   int
	cap  int
}

func NewList[T any](cap int) List[T] {
	if cap <= 0 {
		cap = 1
	}
	var this List[T]
	this.data = unsafe.Pointer(&make([]T, cap)[0])
	this.sz = 0
	this.cap = cap
	return this
}

func (this *List[T]) swap(rhs *List[T]) {
	this.data, rhs.data = clibutils.Swap(this.data, rhs.data)
	this.sz, rhs.sz = clibutils.Swap(this.sz, rhs.sz)
	this.cap, rhs.cap = clibutils.Swap(this.cap, rhs.cap)
}

func (this *List[T]) checkIdx(idx int) {
	if 0 <= idx && idx < this.sz {
		return
	}
	panic("List out of range")
}

func (this *List[T]) realloc(newCap int) {
	var tmp List[T] = NewList[T](newCap)
	for i := range this.Len() {
		tmp.Append(*this.At(i))
	}
	this.swap(&tmp)
}

func (this *List[T]) At(idx int) *T {
	this.checkIdx(idx)
	return clibutils.GetPtr[T](this.data, uintptr(idx))
}

func (this *List[T]) Get(idx int) *T {
	return clibutils.GetPtr[T](this.data, uintptr(idx))
}

func (this *List[T]) AtConst(idx int) T {
	this.checkIdx(idx)
	return *clibutils.GetPtr[T](this.data, uintptr(idx))
}

func (this *List[T]) GetConst(idx int) T {
	return *clibutils.GetPtr[T](this.data, uintptr(idx))
}

func (this *List[T]) Append(value T) {
	if this.sz == this.cap {
		if this.cap == 0 {
			this.realloc(2)
		} else {
			this.realloc(this.cap * 2)
		}
	}
	*clibutils.GetPtr[T](this.data, uintptr(this.sz)) = value
	this.sz += 1
}

func (this *List[T]) ShrinkToFit() {
	this.realloc(this.sz)
}
func (this *List[T]) PopBack() *T {
	if !this.Empty() {
		val := clibutils.GetPtr[T](this.data, uintptr(this.sz-1))
		this.sz -= 1
		return val
	}
	panic("Trying pop back empty")
}

func (this *List[T]) Len() int {
	return this.sz
}

func (this *List[T]) Empty() bool {
	return this.sz == 0
}

func (this *List[T]) Cap() int {
	return this.cap
}

func (this *List[T]) Erase(idx int) {
	this.checkIdx(idx)
	for i := idx; i+1 < this.sz; i++ {
		*clibutils.GetPtr[T](this.data, uintptr(i)) = *clibutils.GetPtr[T](this.data, uintptr(i+1))
	}
	this.sz -= 1
}

func (this *List[T]) All() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		for i := range this.sz {
			if !yield(clibutils.GetPtr[T](this.data, uintptr(i))) {
				return
			}
		}
	}
}

func (this *List[T]) AllIdx() iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		for i := range this.sz {
			if !yield(i, clibutils.GetPtr[T](this.data, uintptr(i))) {
				return
			}
		}
	}
}

func (this *List[T]) AllRev() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		for i := this.sz - 1; i >= 0; i-- {
			if !yield(clibutils.GetPtr[T](this.data, uintptr(i))) {
				return
			}
		}
	}
}

func (this *List[T]) AllIdxRev() iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		for i := this.sz - 1; i >= 0; i-- {
			if !yield(i, clibutils.GetPtr[T](this.data, uintptr(i))) {
				return
			}
		}
	}
}
