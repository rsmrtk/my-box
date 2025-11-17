package utils

func NilToDefault[T any](v *T) T {
	if v != nil {
		return *v
	}
	return *new(T)
}
