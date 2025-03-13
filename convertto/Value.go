package convertto

func Value[T any](pointer *T) T {
	return *pointer
}
