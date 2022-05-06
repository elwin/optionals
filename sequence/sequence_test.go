package sequence

import (
	"strconv"
	"testing"

	"github.com/elwin/optionals/iterator"
	"github.com/stretchr/testify/assert"
)

func TestSequence(t *testing.T) {
	s1 := NewSequence[int](iterator.NewArrayIterator([]int{1, 2, 3}))
	s2 := Map(s1, func(t int) string {
		return strconv.Itoa(t)
	})
	assert.Equal(t, []string{"1", "2", "3"}, s2.ToSlice())
}

func TestXX(t *testing.T) {
	s1 := NewSequence[int](iterator.NewArrayIterator([]int{1, 2, 3}))
	s1.Map2()
}
