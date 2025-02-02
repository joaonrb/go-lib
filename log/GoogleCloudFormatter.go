package log

import (
	"encoding/json"
	"fmt"
)

type googleCloudFormatter struct{}

func GoogleCloudFormatter() Formatter {
	return &googleCloudFormatter{}
}

func (formatter *googleCloudFormatter) Format(event *Event) (string, error) {
	entry := Entry{
		Message:   event.Message(),
		Severity:  event.Level().String(),
		Trace:     "",
		Component: event.Name(),
	}
	out, err := json.Marshal(entry)
	return string(out), err
}

func (formatter *googleCloudFormatter) String() string {
	return "GoogleCloudFormatter{}"
}

type Entry struct {
	Message   string `json:"message"`
	Severity  string `json:"severity,omitempty"`
	Trace     string `json:"logging.googleapis.com/trace,omitempty"`
	Component string `json:"component,omitempty"`
}

func (entry *Entry) String() string {
	return fmt.Sprintf(
		"GoogleCloudLogEntry{Message: \"%s\"}",
		entry.Message,
	)
}
