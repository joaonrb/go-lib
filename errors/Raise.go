package errors

func Raise(err error) {
	if err != nil {
		panic(err)
	}
}
