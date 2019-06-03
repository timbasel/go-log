package log

// The Level describes the severity of the log message
type Level int

const (
	// DebugLevel is used for messages for the developer debugging the application
	DebugLevel Level = iota
	// InfoLevel is used for messages for the user running the application
	InfoLevel
	// ErrorLevel is used in case the application encountered an unhandleable error and will terminate after messaging
	ErrorLevel
)

// String returns the string name for a log level
func (level Level) String() string {
	return [...]string{"DEBUG", "INFO", "ERROR"}[level]
}
