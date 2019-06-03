package log

// The Formatter interface is used to implement custom Formatters
type Formatter interface {
	Format(level Level, msg string) (formattedMsg string)
}
