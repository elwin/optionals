package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	fib = func(yield func(int)) {
		x, y := 0, 1
		for {
			yield(x)
			x, y = y, x+y
		}
	}

	fibResults = []int{0, 1, 1, 2, 3, 5, 8, 13}
)

func TestGenerate(t *testing.T) {
	gen := Generate(fib)

	var out []int
	for gen.HasNext() {
		next := gen.Next()
		out = append(out, next)

		if next > 10 {
			return
		}
	}

	assert.Equal(t, fibResults, out)
}

func TestList(t *testing.T) {
	assert.Equal(t, fibResults, List(Slice(Generate(fib), 8)))
}
