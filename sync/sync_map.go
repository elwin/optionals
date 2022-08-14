package sync

import (
	"sync"

	"github.com/elwin/x/optionals"
)

type Map[a any, b any] struct {
	m sync.Map
}

func NewMap[a any, b any]() *Map[a, b] {
	return &Map[a, b]{}
}

func (s *Map[a, b]) Set(key a, val b) {
	s.m.Store(key, val)
}

func (s *Map[a, b]) Get(key a) optionals.Optional[b] {
	val, ok := s.m.Load(key)
	if !ok {
		return optionals.None[b]()
	}

	return optionals.Some(val.(b))
}

func (s *Map[a, b]) GetOrElse(key a, def b) b {
	return s.Get(key).GetOrElse(def)
}
