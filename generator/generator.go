package generator

type Generator[T any] struct {
	buf     T
	hasNext bool
	c       chan T
}

func (g Generator[T]) HasNext() bool {
	return g.hasNext
}

func (g *Generator[T]) Next() T {
	cur := g.buf
	g.buf, g.hasNext = <-g.c

	return cur
}

func Generate[T any](f func(yield func(T))) Generator[T] {
	c := make(chan T)

	go func() {
		f(func(x T) {
			c <- x
		})
		close(c)
	}()

	next, ok := <-c
	return Generator[T]{
		buf:     next,
		hasNext: ok,
		c:       c,
	}
}

func Slice[T any](generator Generator[T], n int) Generator[T] {
	return Generate[T](func(yield func(T)) {
		for n > 0 && generator.hasNext {
			yield(generator.Next())
			n--
		}

		return
	})
}

func List[T any](generator Generator[T]) []T {
	var out []T
	for generator.HasNext() {
		out = append(out, generator.Next())
	}

	return out
}
