package tb

// Coalesce returns the first non-zero value from the given list of arguments.
// If none of the values are non-zero, a zero value for the given type is
// returned. Note that pointers are supported by this function, a nil pointer
// is a zero value.
func Coalesce[T comparable](items ...T) T {
	var zero T
	for _, item := range items {
		if item != zero {
			return item
		}
	}
	return zero
}

// IfNul returns the fallback value if the pointer argument is nil. If the
// pointer argument is non-nil, it is dereferenced, creating a shallow copy
// of the object.
func IfNil[T any](p *T, fallback T) T {
	if p != nil {
		return *p
	}
	return fallback
}

// ZeroIfNil returns a zero value for the given type if the pointer argument
// is nil. If the argument is non-nil, it is dereferenced, creating a shallow
// copy of the object.
func ZeroIfNil[T any](p *T) T {
	var zero T
	return IfNil(p, zero)
}

// If returns the second argument if the first argument is true, otherwise
// is returns the third argument. Equivalent to `first ? second : third` in other
// programming languages.
func If[T any](first bool, second T, third T) T {
	if first {
		return second
	}
	return third
}
