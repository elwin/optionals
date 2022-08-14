package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSynchronizedMap(t *testing.T) {
	m := NewMap[string, int]()
	m.Set("foo", 42)
	val, ok := m.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, 42, val)
}
