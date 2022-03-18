package log

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/acarl005/stripansi"
	"github.com/benbjohnson/clock"
)

// JSONFormatter formats log to json objects with timestamp, function, package and log level
type JSONFormatter struct {
	ColorsDisabled    bool
	TimestampDisabled bool
	TimestampLayout   string
	Clock             clock.Clock
	CallerDisabled    bool
	PrettyPrint       bool
}

// NewJSONFormatter initializes a new JSONFormatter
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{
		ColorsDisabled:  true,
		TimestampLayout: time.RFC3339,
		Clock:           clock.New(),
	}
}

// Format formats a single log message
func (f *JSONFormatter) Format(level Level, msg string) string {
	entries := map[string]string{}

	if !f.TimestampDisabled {
		entries["time"] = getTimestamp(f.Clock, f.TimestampLayout)
	}

	if !f.CallerDisabled {
		entries["package"] = getCallersPackageName()
		entries["function"] = getCallersFunctionName()
	}

	entries["level"] = level.String()
	if f.CallerDisabled {
		entries["msg"] = stripansi.Strip(msg)
	} else {
		entries["msg"] = msg
	}

	buffer := &strings.Builder{}
	encoder := json.NewEncoder(buffer)
	if f.PrettyPrint {
		encoder.SetIndent("", "\t")
	}
	if err := encoder.Encode(entries); err != nil {
		return ""
	}
	return buffer.String()
}
