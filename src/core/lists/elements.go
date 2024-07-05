package lists

type Element[T any] struct {
	value T
	prev  *Element[T]
	next  *Element[T]
}

func NewElement[T any](value T) *Element[T] {
	return &Element[T]{
		value: value,
		prev:  nil,
		next:  nil,
	}
}

func (e *Element[T]) Value() T {
	return e.value
}

func (e *Element[T]) SetValue(value T) {
	e.value = value
}

func (e *Element[T]) Prev() *Element[T] {
	return e.prev
}

func (e *Element[T]) SetPrev(element *Element[T]) {
	e.prev = element
}

func (e *Element[T]) Next() *Element[T] {
	return e.next
}

func (e *Element[T]) SetNext(element *Element[T]) {
	e.next = element
}
