package log_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	log "github.com/timbasel/go-log/pkg"
)

func prepareTestLogger() (logger *log.Logger, output *strings.Builder) {
	output = &strings.Builder{}
	logger = log.NewLogger()
	logger.ClearOutputs()
	logger.SetFormattedOutput(output, log.NewRawFormatter())
	return logger, output
}

func TestLogger(t *testing.T) {
	logger, output := prepareTestLogger()

	msg := "this is a test message"

	testCases := []struct {
		function func(msg ...string)
		expected string
	}{
		{logger.Error, msg + "\n"},
		{logger.Info, msg + "\n"},
		{logger.Debug, ""}, // should be empty because debug mode is disabled by default
	}

	for _, testCase := range testCases {
		testCase.function(msg)

		assert.Equal(t, output.String(), testCase.expected)

		output.Reset()
	}
}

func TestLoggerEnabledDebugMode(t *testing.T) {
	logger, output := prepareTestLogger()
	logger.SetDebugMode(true)

	msg := "this is a test message"

	testCases := []struct {
		function func(msg ...string)
		expected string
	}{
		{logger.Error, msg + "\n"},
		{logger.Info, msg + "\n"},
		{logger.Debug, msg + "\n"},
	}

	for _, testCase := range testCases {
		testCase.function(msg)

		assert.Equal(t, output.String(), testCase.expected)

		output.Reset()
	}
}

func TestLoggerMultipleOutputs(t *testing.T) {
	output1 := &strings.Builder{}
	output2 := &strings.Builder{}
	logger := log.NewLogger()
	msg := "this is a test message"
	logger.ClearOutputs()
	logger.SetOutput(output1)
	logger.SetOutput(output2)

	logger.Info(msg)

	assert.Contains(t, output1.String(), msg)
	assert.Contains(t, output2.String(), msg)
}
