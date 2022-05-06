package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayIterator(t *testing.T) {
	x := NewArrayIterator([]int{1, 2, 3})
	assert.True(t, x.HasNext())
	assert.Equal(t, 1, x.Next())
	assert.True(t, x.HasNext())
	assert.Equal(t, 2, x.Next())
	assert.True(t, x.HasNext())
	assert.Equal(t, 3, x.Next())
	assert.False(t, x.HasNext())
}

func TestArrayIterator_Empty(t *testing.T) {
	x := NewArrayIterator([]int{})
	assert.False(t, x.HasNext())
}

func TestArrayIterator_Slice(t *testing.T) {
	x := NewArrayIterator([]int{1, 2, 3})
	x.Next()
	assert.Equal(t, []int{2, 3}, x.ToSlice())
}
