package conts

type Queue[T any] struct {
	data LinkedList[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{data: NewLinkedList[T]()}
}

func (this *Queue[T]) Push(value T) {
	this.data.PushBack(value)
}

func (this *Queue[T]) Pop() T {
	return this.data.PopFront()
}

func (this *Queue[T]) Len() int {
	return this.data.Len()
}
func (this *Queue[T]) Empty() bool {
	return this.data.Empty()
}
