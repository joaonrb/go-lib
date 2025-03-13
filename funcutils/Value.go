package funcutils

func Value[T any](pointer *T) T {
	return *pointer
}
