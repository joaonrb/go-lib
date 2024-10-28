package log

type Formatter interface {
	Format(event *Event) (string, error)
}
