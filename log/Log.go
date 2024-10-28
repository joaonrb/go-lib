package log

import "os"

var dl = defaultLogger()

func defaultLogger() *Logger {
	config := Config{
		Name:         os.Args[0],
		Level:        LevelDefault,
		Formatter:    SimpleFormatter(),
		Writer:       ConsoleWriter(),
		PanicOnError: false,
	}
	return NewLogger(config)
}

func SetDefaultLogger(config Config) {
	dl = NewLogger(config)
}

func NewEvent(level Level, msg string, args ...any) *Event {
	return dl.Event(level, msg, args...)
}

func Print(msg string, args ...any) {
	dl.Print(msg, args...).Done()
}

func Debug(msg string, args ...any) {
	dl.Debug(msg, args...).Done()
}

func Info(msg string, args ...any) {
	dl.Info(msg, args...).Done()
}

func Notice(msg string, args ...any) {
	dl.Notice(msg, args...).Done()
}

func Warning(msg string, args ...any) {
	dl.Warning(msg, args...).Done()
}

func Error(msg string, args ...any) {
	dl.Error(msg, args...).Done()
}

func Critical(msg string, args ...any) {
	dl.Critical(msg, args...).Done()
}

func Alert(msg string, args ...any) {
	dl.Alert(msg, args...).Done()
}

func Emergency(msg string, args ...any) {
	dl.Emergency(msg, args...).Done()
}
