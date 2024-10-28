package log

import (
	"log"
	"os"
)

type consoleWriter struct {
	logger *log.Logger
}

func ConsoleWriter() Writer {
	return &consoleWriter{logger: log.New(os.Stdin, "", 0)}
}

func (writer *consoleWriter) Write(message string) error {
	writer.logger.Println(message)
	return nil
}

func (*consoleWriter) String() string {
	return "ConsoleWriter{}"
}
