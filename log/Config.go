package log

import (
	"fmt"
	"github.com/joaonrb/go-lib/convertto"
)

type Config struct {
	Name         string
	Level        Level
	Formatter    Formatter
	Writer       Writer
	PanicOnError bool
}

func (config Config) String() string {
	return fmt.Sprintf(
		"Config{Name: %s, Level: %s, Formatter: %s, Writer: %s, PanicOnError: %t}",
		convertto.JSON(config.Name).MustValue(),
		config.Level,
		config.Formatter,
		config.Writer,
		config.PanicOnError,
	)
}
