package convert

func ToValue[T any](pointer *T) T {
	return *pointer
}
