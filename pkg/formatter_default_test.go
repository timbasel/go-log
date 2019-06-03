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
	clock.Add(26*time.Hour + 4*time.Minute + 5*time.Second) // set timestamp to 1970-01-02T03:04:05+00:00
	formatter.Clock = clock
	return formatter
}

func TestDefaultFormatter(t *testing.T) {
	formatter := prepareTestDefaultFormatter()

	msg := "this is a test message"
	timestamp := "1970-01-02 03:04:05"
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
	timestamp := "1970-01-02 03:04:05"

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
	timestamp := "Fri Jan 02 03:04:05 1970"

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
