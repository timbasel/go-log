package log_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"
	log "github.com/timbasel/go-log/pkg"
)

func prepareTestJSONFormatter() (formatter *log.JSONFormatter) {
	formatter = log.NewJSONFormatter()
	clock := clock.NewMock()
	timestamp, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+00:00")
	clock.Set(timestamp)
	formatter.Clock = clock
	return formatter
}

func TestJSONFormatter(t *testing.T) {
	formatter := prepareTestJSONFormatter()

	msg := "this is a test message"
	timestamp := "2006-01-02T15:04:05Z"
	packageName := "github.com/timbasel/go-log/pkg_test"
	functionName := "TestJSONFormatter"
	expectedFormat := "{\"function\":\"%s\",\"level\":\"%s\",\"msg\":\"%s\",\"package\":\"%s\",\"time\":\"%s\"}\n"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, fmt.Sprintf(expectedFormat, functionName, "ERROR", msg, packageName, timestamp)},
		{log.InfoLevel, fmt.Sprintf(expectedFormat, functionName, "INFO", msg, packageName, timestamp)},
		{log.DebugLevel, fmt.Sprintf(expectedFormat, functionName, "DEBUG", msg, packageName, timestamp)},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, testCase.expected, formattedMsg)
	}
}

func TestJSONFormatterDisabledTimestamp(t *testing.T) {
	formatter := prepareTestJSONFormatter()
	formatter.TimestampDisabled = true

	msg := "this is a test message"
	packageName := "github.com/timbasel/go-log/pkg_test"
	functionName := "TestJSONFormatterDisabledTimestamp"
	expectedFormat := "{\"function\":\"%s\",\"level\":\"%s\",\"msg\":\"%s\",\"package\":\"%s\"}\n"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, fmt.Sprintf(expectedFormat, functionName, "ERROR", msg, packageName)},
		{log.InfoLevel, fmt.Sprintf(expectedFormat, functionName, "INFO", msg, packageName)},
		{log.DebugLevel, fmt.Sprintf(expectedFormat, functionName, "DEBUG", msg, packageName)},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, testCase.expected, formattedMsg)
	}
}

func TestJSONFormatterDisabledCaller(t *testing.T) {
	formatter := prepareTestJSONFormatter()
	formatter.CallerDisabled = true

	msg := "this is a test message"
	timestamp := "2006-01-02T15:04:05Z"
	expectedFormat := "{\"level\":\"%s\",\"msg\":\"%s\",\"time\":\"%s\"}\n"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, fmt.Sprintf(expectedFormat, "ERROR", msg, timestamp)},
		{log.InfoLevel, fmt.Sprintf(expectedFormat, "INFO", msg, timestamp)},
		{log.DebugLevel, fmt.Sprintf(expectedFormat, "DEBUG", msg, timestamp)},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, testCase.expected, formattedMsg)
	}
}

func TestJSONFormatterPrettyPrint(t *testing.T) {
	formatter := prepareTestJSONFormatter()
	formatter.PrettyPrint = true

	msg := "this is a test message"
	timestamp := "2006-01-02T15:04:05Z"
	packageName := "github.com/timbasel/go-log/pkg_test"
	functionName := "TestJSONFormatterPrettyPrint"
	expectedFormat := "{\n\t\"function\": \"%s\",\n\t\"level\": \"%s\",\n\t\"msg\": \"%s\",\n\t\"package\": \"%s\",\n\t\"time\": \"%s\"\n}\n"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, fmt.Sprintf(expectedFormat, functionName, "ERROR", msg, packageName, timestamp)},
		{log.InfoLevel, fmt.Sprintf(expectedFormat, functionName, "INFO", msg, packageName, timestamp)},
		{log.DebugLevel, fmt.Sprintf(expectedFormat, functionName, "DEBUG", msg, packageName, timestamp)},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, testCase.expected, formattedMsg)
	}
}
