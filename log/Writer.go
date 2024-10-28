package log

type Writer interface {
	Write(message string) error
}
