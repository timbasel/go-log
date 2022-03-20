package log_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"
	"github.com/timbasel/go-log/pkg/log"
)

func prepareTestCSVFormatter() *log.CSVFormatter {
	clock := clock.NewMock()
	timestamp, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+00:00")
	clock.Set(timestamp)

	formatter := log.NewCSVFormatter()
	formatter.Clock = clock
	return formatter
}

func TestCSVFormatter(t *testing.T) {
	formatter := prepareTestCSVFormatter()

	msg := "this is a test message, with a comma"
	timestamp := "2006-01-02T15:04:05Z"
	packageName := "github.com/timbasel/go-log/pkg/log_test"
	functionName := "TestCSVFormatter"
	expectedFormat := "%s,%s,%s,%s,\"%s\"\n"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, fmt.Sprintf(expectedFormat, timestamp, "ERROR", packageName, functionName, msg)},
		{log.InfoLevel, fmt.Sprintf(expectedFormat, timestamp, "INFO", packageName, functionName, msg)},
		{log.DebugLevel, fmt.Sprintf(expectedFormat, timestamp, "DEBUG", packageName, functionName, msg)},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, testCase.expected, formattedMsg)
	}
}

func TestCSVFormatterDisabledTimestamp(t *testing.T) {
	formatter := prepareTestCSVFormatter()
	formatter.TimestampDisabled = true

	msg := "this is a test message, with a comma"
	packageName := "github.com/timbasel/go-log/pkg/log_test"
	functionName := "TestCSVFormatterDisabledTimestamp"
	expectedFormat := "%s,%s,%s,\"%s\"\n"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, fmt.Sprintf(expectedFormat, "ERROR", packageName, functionName, msg)},
		{log.InfoLevel, fmt.Sprintf(expectedFormat, "INFO", packageName, functionName, msg)},
		{log.DebugLevel, fmt.Sprintf(expectedFormat, "DEBUG", packageName, functionName, msg)},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, testCase.expected, formattedMsg)
	}
}

func TestCSVFormatterDisabledCaller(t *testing.T) {
	formatter := prepareTestCSVFormatter()
	formatter.CallerDisabled = true

	msg := "this is a test message, with a comma"
	timestamp := "2006-01-02T15:04:05Z"
	expectedFormat := "%s,%s,\"%s\"\n"

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
