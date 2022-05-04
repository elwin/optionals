package optionals

// Optional may or may not hold an element
type Optional[T any] struct {
	element T
	exists  bool
}

// Some returns an optional containing the provided element
func Some[T any](element T) Optional[T] {
	return Optional[T]{
		element: element,
		exists:  true,
	}
}

// None returns an optional without any element
// Shall be called like `None[type]()`
func None[T any]() Optional[T] {
	return Optional[T]{}
}

// Exists returns whether the optional contains an element
func (o Optional[T]) Exists() bool {
	return o.exists
}

// Get returns the contained element if it exists and panics otherwise
func (o Optional[T]) Get() T {
	if !o.exists {
		panic("trying to get element from empty optional")
	}

	return o.element
}

// GetOrElse returns the contained element if it exists and def otherwise
func (o Optional[T]) GetOrElse(def T) T {
	if !o.exists {
		return def
	}

	return o.element
}
