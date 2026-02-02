package conts

type Set[T comparable] struct {
	data HashMap[T, struct{}]
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{data: NewHashMap[T, struct{}]()}
}

func (this *Set[T]) Add(value T) {
	this.data.Set(value, struct{}{})
}
func (this *Set[T]) Contains(value T) bool {
	return this.data.Get(value) != nil
}
func (this *Set[T]) Delete(value T) {
	this.data.Delete(value)
}
