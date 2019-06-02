package log

// The Level describes the severity of the log message
type Level int

const (
	// Debug level is used for messages for the developer debugging the application
	Debug Level = iota
	// Info level is used for messages for the user running the application
	Info
	// Error level is used in case the application encountered an unhandleable error and will terminate after messaging
	Error
)

// String returns the string name for a log level
func (level Level) String() string {
	return [...]string{"DEBUG", "INFO", "ERROR"}[level]
}
