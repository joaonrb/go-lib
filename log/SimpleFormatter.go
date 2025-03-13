package log

import "fmt"

type simpleFormatter struct{}

func SimpleFormatter() Formatter {
	return &simpleFormatter{}
}

func (formatter *simpleFormatter) Format(event *Event) (string, error) {
	msg := fmt.Sprintf(
		"%s %s:%s - %s",
		event.CreatedAt().Format("2006-01-02 15:04:05"),
		event.Name(),
		event.Level().String(),
		event.Message(),
	)
	return msg, nil
}

func (*simpleFormatter) String() string {
	return "SimpleFormater{}"
}
