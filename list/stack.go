package list

type Stack[T any] struct {
	data List[T]
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{data: NewList[T](0)}
}

func (this *Stack[T]) Push(value T) {
	this.data.Append(value)
}

func (this *Stack[T]) Pop() T {
	return *this.data.PopBack()
}

func (this *Stack[T]) Len() int {
	return this.data.Len()
}
func (this *Stack[T]) Empty() bool {
	return this.data.Empty()
}
