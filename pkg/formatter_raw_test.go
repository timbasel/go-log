package log_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	log "github.com/timbasel/go-log/pkg"
)

func TestRawFormatter(t *testing.T) {
	formatter := log.NewRawFormatter()

	msg := "this is a test message"

	testCases := []struct {
		level    log.Level
		expected string
	}{
		{log.Error, msg + "\n"},
		{log.Info, msg + "\n"},
		{log.Debug, msg + "\n"},
	}

	for _, testCase := range testCases {
		formattedMsg := formatter.Format(testCase.level, msg)

		assert.Equal(t, formattedMsg, testCase.expected)
	}
}
