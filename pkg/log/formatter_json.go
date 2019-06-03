package log

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/benbjohnson/clock"
)

// JSONFormatter formats log to json objects with timestamp, function, package and log level
type JSONFormatter struct {
	TimestampDisabled bool
	TimestampLayout   string
	Clock             clock.Clock
	CallerDisabled    bool
	PrettyPrint       bool
}

// NewJSONFormatter initializes a new JSONFormatter
func NewJSONFormatter() (f *JSONFormatter) {
	return &JSONFormatter{
		TimestampLayout: time.RFC3339,
	}
}

// Format formats a single log message
func (f *JSONFormatter) Format(level Level, msg string) (formattedMsg string) {
	entries := map[string]string{}

	if !f.TimestampDisabled {
		entries["time"] = getTimestamp(f.Clock, f.TimestampLayout)
	}

	if !f.CallerDisabled {
		entries["package"] = getCallersPackageName()
		entries["function"] = getCallersFunctionName()
	}

	entries["level"] = level.String()
	entries["msg"] = msg

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