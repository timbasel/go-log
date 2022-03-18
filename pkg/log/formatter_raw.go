package log

import "github.com/acarl005/stripansi"

// RawFormatter formats log to just the message
type RawFormatter struct {
	ColorsDisabled bool
}

// NewRawFormatter initializes a new RawFormatter
func NewRawFormatter() (f *RawFormatter) {
	return &RawFormatter{}
}

// Format formats a single log message
func (f *RawFormatter) Format(level Level, msg string) string {
	if f.ColorsDisabled {
		return stripansi.Strip(msg) + "\n"
	} else {
		return msg + "\n"
	}
}
