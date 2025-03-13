package errors

func MustNil(err error) {
	if err != nil {
		panic(err)
	}
}
