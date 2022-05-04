package optionals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	x := Some(5)
	assert.True(t, x.Exists())
	assert.Equal(t, x.Get(), 5)
	assert.Equal(t, x.GetOrElse(4), 5)
	assert.Panics(t, func() {
		None[int]().Get()
	})
	assert.Equal(t, None[int]().GetOrElse(5), 5)
}

func TestFunc(t *testing.T) {
	y := func(some bool) Optional[string] {
		if some {
			return Some("present")
		} else {
			return None[string]()
		}
	}

	assert.Equal(t, y(true).GetOrElse("default"), "present")
	assert.Equal(t, y(false).GetOrElse("default"), "default")
	assert.Panics(t, func() {
		y(false).Get()
	})
}
