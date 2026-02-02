package clibutils

import "unsafe"

func Swap[T any](a, b T) (T, T) {
	return b, a
}

func Sizeof[T any]() uintptr {
	return unsafe.Sizeof(*(*T)(nil))
}

func GetPtr[T any](array unsafe.Pointer, offset uintptr) *T {
	v := unsafe.Pointer(uintptr(array) + offset*Sizeof[T]())
	return (*T)(v)
}
