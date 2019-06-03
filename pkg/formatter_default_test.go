package log_test

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"

	log "github.com/timbasel/go-log/pkg"
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
	caller := "pkg_test.TestDefaultFormatter"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, timestamp + " <" + caller + "> ERROR:\t" + msg + "\n"},
		{log.InfoLevel, timestamp + " <" + caller + "> INFO:\t" + msg + "\n"},
		{log.DebugLevel, timestamp + " <" + caller + "> DEBUG:\t" + msg + "\n"},
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
	caller := "pkg_test.TestDefaultFormatterDisabledTimestamp"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, "<" + caller + "> ERROR:\t" + msg + "\n"},
		{log.InfoLevel, "<" + caller + "> INFO:\t" + msg + "\n"},
		{log.DebugLevel, "<" + caller + "> DEBUG:\t" + msg + "\n"},
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

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, timestamp + " ERROR:\t" + msg + "\n"},
		{log.InfoLevel, timestamp + " INFO:\t" + msg + "\n"},
		{log.DebugLevel, timestamp + " DEBUG:\t" + msg + "\n"},
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
	caller := "pkg_test.TestDefaultFormatterCustomTimestamp"
	timestamp := "Mon Jan 02 15:04:05 2006"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, timestamp + " <" + caller + "> ERROR:\t" + msg + "\n"},
		{log.InfoLevel, timestamp + " <" + caller + "> INFO:\t" + msg + "\n"},
		{log.DebugLevel, timestamp + " <" + caller + "> DEBUG:\t" + msg + "\n"},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, testCase.expected, formattedMsg)
	}
}
