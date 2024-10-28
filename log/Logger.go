package log

import (
	"fmt"
	"os"
)

type Logger struct {
	config Config
}

func NewLogger(config Config) *Logger {
	return &Logger{
		config: config,
	}
}

func (logger *Logger) String() string {
	return fmt.Sprintf(
		"Logger{Config: %s}",
		logger.config,
	)
}

func (logger *Logger) Level() Level {
	return logger.config.Level
}

func (logger *Logger) Event(level Level, msg string, args ...any) *Event {
	return newEvent(level, logger, msg, args...)
}

func (logger *Logger) Print(msg string, args ...any) *Event {
	return logger.Event(LevelDefault, msg, args...)
}

func (logger *Logger) Debug(msg string, args ...any) *Event {
	return logger.Event(LevelDebug, msg, args...)
}

func (logger *Logger) Info(msg string, args ...any) *Event {
	return logger.Event(LevelInfo, msg, args...)
}

func (logger *Logger) Notice(msg string, args ...any) *Event {
	return logger.Event(LevelNotice, msg, args...)
}

func (logger *Logger) Warning(msg string, args ...any) *Event {
	return logger.Event(LevelWarning, msg, args...)
}

func (logger *Logger) Error(msg string, args ...any) *Event {
	return logger.Event(LevelError, msg, args...)
}

func (logger *Logger) Critical(msg string, args ...any) *Event {
	return logger.Event(LevelCritical, msg, args...)
}

func (logger *Logger) Alert(msg string, args ...any) *Event {
	return logger.Event(LevelAlert, msg, args...)
}

func (logger *Logger) Emergency(msg string, args ...any) *Event {
	return logger.Event(LevelEmergency, msg, args...)
}

func (logger *Logger) Flush(event *Event) {
	if event.level < logger.config.Level {
		return
	}
	var err error
	var msg string
	msg, err = logger.config.Formatter.Format(event)
	if err == nil {
		err = logger.config.Writer.Write(msg)
	}
	switch {
	case err != nil && logger.config.PanicOnError:
		panic(err)
	case err != nil:
		_, _ = fmt.Fprintf(os.Stderr, "Failed to flush event: %v", err)
	}

}
