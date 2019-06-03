package log

import (
	"strings"

	"github.com/benbjohnson/clock"
	"github.com/gookit/color"
)

// DefaultFormatter formats log to human readable text with timestamp, caller and log level
type DefaultFormatter struct {
	ColorsDisabled    bool
	Colors            map[Level]color.Style
	TimestampDisabled bool
	TimestampLayout   string
	Clock             clock.Clock
	CallerDisabled    bool
}

// NewDefaultFormatter initializes a new DefaultFormatter
func NewDefaultFormatter() (f *DefaultFormatter) {
	return &DefaultFormatter{
		Colors: map[Level]color.Style{
			DebugLevel: color.New(color.FgLightCyan),
			InfoLevel:  color.New(color.FgLightWhite),
			ErrorLevel: color.New(color.FgLightRed),
		},
		TimestampLayout: "2006-01-02 15:04:05",
		Clock:           clock.New(),
	}
}

// Format formats a single log message
func (f *DefaultFormatter) Format(level Level, msg string) (formattedMsg string) {
	entry := &strings.Builder{}

	if !f.TimestampDisabled {
		entry.WriteString(getTimestamp(f.Clock, f.TimestampLayout))
		entry.WriteString(" ")
	}

	if !f.CallerDisabled {
		entry.WriteString("<")
		entry.WriteString(getCallersFullFunctionName())
		entry.WriteString("> ")
	}

	entry.WriteString(level.String())
	entry.WriteString(":\t")

	if f.ColorsDisabled {
		formattedMsg = entry.String() + msg + "\n"
	} else {
		formattedMsg = f.Colors[level].Render(entry.String()) + msg + "\n"
	}
	return formattedMsg
}