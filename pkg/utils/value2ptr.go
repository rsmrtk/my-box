package utils

func ValueToPtr[T any](v T) *T {
	return &v
}
