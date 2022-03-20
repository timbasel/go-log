package log_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"

	"github.com/timbasel/go-log/pkg/log"
)

func prepareTestDefaultFormatter() (formatter *log.DefaultFormatter) {
	formatter = log.NewDefaultFormatter()
	formatter.ColorsDisabled = true
	clock := clock.NewMock()
	timestamp, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+01:00")
	clock.Set(timestamp)
	formatter.Clock = clock
	return formatter
}

func TestDefaultFormatter(t *testing.T) {
	formatter := prepareTestDefaultFormatter()

	msg := "this is a test message"
	timestamp := "2006-01-02 15:04:05"
	caller := "log_test.TestDefaultFormatter"
	expectedFormat := "%s %s <%s>: %s\n"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, fmt.Sprintf(expectedFormat, timestamp, "ERROR", caller, msg)},
		{log.InfoLevel, fmt.Sprintf(expectedFormat, timestamp, "INFO", caller, msg)},
		{log.DebugLevel, fmt.Sprintf(expectedFormat, timestamp, "DEBUG", caller, msg)},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, testCase.expected, formattedMsg)
	}
}

func TestDefaultFormatterDisabledTimestamp(t *testing.T) {
	formatter := prepareTestDefaultFormatter()
	formatter.TimestampDisabled = true

	msg := "this is a test message"
	caller := "log_test.TestDefaultFormatterDisabledTimestamp"
	expectedFormat := "%s <%s>: %s\n"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, fmt.Sprintf(expectedFormat, "ERROR", caller, msg)},
		{log.InfoLevel, fmt.Sprintf(expectedFormat, "INFO", caller, msg)},
		{log.DebugLevel, fmt.Sprintf(expectedFormat, "DEBUG", caller, msg)},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, testCase.expected, formattedMsg)
	}
}

func TestDefaultFormatterDisabledCaller(t *testing.T) {
	formatter := prepareTestDefaultFormatter()
	formatter.CallerDisabled = true

	msg := "this is a test message"
	timestamp := "2006-01-02 15:04:05"
	expectedFormat := "%s %s: %s\n"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, fmt.Sprintf(expectedFormat, timestamp, "ERROR", msg)},
		{log.InfoLevel, fmt.Sprintf(expectedFormat, timestamp, "INFO", msg)},
		{log.DebugLevel, fmt.Sprintf(expectedFormat, timestamp, "DEBUG", msg)},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, testCase.expected, formattedMsg)
	}
}

func TestDefaultFormatterCustomTimestamp(t *testing.T) {
	formatter := prepareTestDefaultFormatter()
	formatter.TimestampLayout = "Mon Jan 02 15:04:05 2006"

	msg := "this is a test message"
	caller := "log_test.TestDefaultFormatterCustomTimestamp"
	timestamp := "Mon Jan 02 15:04:05 2006"
	expectedFormat := "%s %s <%s>: %s\n"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, fmt.Sprintf(expectedFormat, timestamp, "ERROR", caller, msg)},
		{log.InfoLevel, fmt.Sprintf(expectedFormat, timestamp, "INFO", caller, msg)},
		{log.DebugLevel, fmt.Sprintf(expectedFormat, timestamp, "DEBUG", caller, msg)},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, testCase.expected, formattedMsg)
	}
}
