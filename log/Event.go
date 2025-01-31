package log

import (
	"fmt"
	"time"

	"github.com/joaonrb/go-lib/convertto"
)

type Event struct {
	message      string
	args         []any
	level        Level
	createAt     time.Time
	labels       map[string]string
	logger       *Logger
	buildMessage *string
}

func newEvent(level Level, logger *Logger, message string, args ...any) *Event {
	return &Event{
		message:  message,
		args:     args,
		level:    level,
		createAt: time.Now(),
		labels:   map[string]string{},
		logger:   logger,
	}
}

func (event *Event) String() string {
	return fmt.Sprintf(
		"Event{Name: %s, Message: %s, Level: %s}",
		convertto.JSON(event.Name()).TryValue(),
		convertto.JSON(event.message).TryValue(),
		event.level.String(),
	)
}

func (event *Event) Name() string {
	return event.logger.config.Name
}

func (event *Event) Message() string {
	if event.buildMessage == nil {
		event.buildMessage = convertto.Pointer(fmt.Sprintf(event.message, event.args...))
	}
	return *event.buildMessage
}

func (event *Event) Level() Level {
	return event.level
}

func (event *Event) CreatedAt() time.Time {
	return event.createAt
}

func (event *Event) Labels(fn func(key, value string)) {
	for k, v := range event.labels {
		fn(k, v)
	}
}

func (event *Event) Label(key, value string) {
	event.labels[key] = value
}

func (event *Event) DeleteLabel(key string) {
	delete(event.labels, key)
}

func (event *Event) Done() {
	event.logger.Flush(event)
}
