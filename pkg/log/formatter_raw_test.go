package log_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timbasel/go-log/pkg/log"
)

func TestRawFormatter(t *testing.T) {
	formatter := log.NewRawFormatter()

	msg := "this is a test message"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.ErrorLevel, msg + "\n"},
		{log.InfoLevel, msg + "\n"},
		{log.DebugLevel, msg + "\n"},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, testCase.expected, formattedMsg)
	}
}
