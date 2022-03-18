package log

import (
	"encoding/csv"
	"strings"
	"time"

	"github.com/benbjohnson/clock"
)

// CSVFormatter formats the log to csv lines with timestamp, package, function, log level and message
type CSVFormatter struct {
	TimestampDisabled bool
	TimestampLayout   string
	Clock             clock.Clock
	CallerDisabled    bool
	PrettyPrint       bool
}

func NewCSVFormatter() *CSVFormatter {
	return &CSVFormatter{
		TimestampLayout: time.RFC3339,
		Clock:           clock.New(),
	}
}

func (f *CSVFormatter) Format(level Level, msg string) string {
	entries := []string{}

	if !f.TimestampDisabled {
		entries = append(entries, getTimestamp(f.Clock, f.TimestampLayout))
	}

	if !f.CallerDisabled {
		entries = append(entries, getCallersPackageName())
		entries = append(entries, getCallersFunctionName())
	}

	entries = append(entries, level.String())
	entries = append(entries, msg)

	buffer := &strings.Builder{}
	writer := csv.NewWriter(buffer)
	if err := writer.Write(entries); err != nil {
		return err.Error()
	}
	writer.Flush()
	return buffer.String()
}
