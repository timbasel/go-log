package log

// RawFormatter formats log to just the message
type RawFormatter struct{}

// NewRawFormatter initializes a new RawFormatter
func NewRawFormatter() (f *RawFormatter) {
	return &RawFormatter{}
}

// Format formats a single log message
func (f *RawFormatter) Format(level Level, msg string) (formattedMsg string) {
	return msg + "\n"
}
