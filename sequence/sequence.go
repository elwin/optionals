package sequence

import "github.com/elwin/optionals/iterator"

type Sequence[T any] struct {
	iterator.Iterator[T]
}

func NewSequence[T any](it iterator.Iterator[T]) *Sequence[T] {
	return &Sequence[T]{
		it,
	}
}

func Map[T any, S any](s *Sequence[T], f func(T) S) *Sequence[S] {
	return NewSequence[S](&MapIterator[T, S]{
		it: s,
		f:  f,
	})
}

type MapIterator[T any, S any] struct {
	it iterator.Iterator[T]
	f  func(T) S
}

func (m *MapIterator[T, S]) Next() S {
	return m.f(m.it.Next())
}

func (m *MapIterator[T, S]) HasNext() bool {
	return m.it.HasNext()
}

func (m *MapIterator[T, S]) ToSlice() []S {
	var out []S
	for {
		if !m.HasNext() {
			return out
		}

		out = append(out, m.Next())
	}
}
