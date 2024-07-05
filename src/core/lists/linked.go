package lists

type Linked[T any] struct {
	root *Element[T]
	last *Element[T]
	size int
}

func NewLinked[T any]() *Linked[T] {
	return &Linked[T]{
		root: nil,
		last: nil,
		size: 0,
	}
}

func (l *Linked[T]) Add(value T) *Element[T] {
	var (
		element *Element[T] = nil
	)
	element = NewElement(value)
	l.size += 1
	if l.root != nil {
		l.last.SetNext(element)
		element.SetPrev(l.last)
		element.SetNext(nil)
		l.last = element
	} else {
		l.root = element
		l.root.SetPrev(nil)
		l.root.SetNext(nil)
	}
	return element
}

func (l *Linked[T]) Get(index int) *Element[T] {
	var (
		position int         = 0
		result   *Element[T] = nil
		element  *Element[T] = l.root
	)
	if index > l.size || index < 0 {
		return nil
	}
	for {
		if position == index {
			result = element
			break
		}
		element = element.Next()
		position += 1
	}
	return result
}

func (l *Linked[T]) First() *Element[T] {
	return l.root
}

func (l *Linked[T]) Last() *Element[T] {
	return l.last
}

func (l *Linked[T]) Size() int {
	return l.size
}

func (l *Linked[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *Linked[T]) IsNotEmpty() bool {
	return l.size > 0
}
