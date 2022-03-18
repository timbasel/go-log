package log

import (
	"fmt"
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
			DebugLevel: color.New(color.BgGray, color.FgWhite),
			InfoLevel:  color.New(color.BgWhite, color.Black),
			ErrorLevel: color.New(color.BgRed, color.FgWhite),
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

	if f.ColorsDisabled {
		entry.WriteString(level.String())
	} else {
		entry.WriteString(f.Colors[level].Render(center(level.String(), 9)))
	}
	entry.WriteString(":\t")

	return entry.String() + msg + "\n"
}

func center(s string, w int) string {
	return fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w+len(s))/2, s))
}
