package iterator

type Iterator[T any] interface {
	Next() T
	HasNext() bool
	ToSlice() []T
}

type ArrayIterator[T any] struct {
	elems []T
	pos   int
}

func (a *ArrayIterator[T]) Next() T {
	next := a.elems[a.pos]
	a.pos++
	return next
}

func (a *ArrayIterator[T]) HasNext() bool {
	return a.pos < len(a.elems)
}

func (a *ArrayIterator[T]) ToSlice() []T {
	var out []T
	for {
		if !a.HasNext() {
			return out
		}

		out = append(out, a.Next())
	}
}

func NewArrayIterator[T any](elems []T) *ArrayIterator[T] {
	return &ArrayIterator[T]{
		elems: elems,
		pos:   0,
	}
}
