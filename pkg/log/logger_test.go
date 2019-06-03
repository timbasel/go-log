package log_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timbasel/go-log/pkg/log"
)

func prepareTestLogger() (logger *log.Logger, output *strings.Builder) {
	output = &strings.Builder{}
	logger = log.NewLogger()
	logger.ClearOutputs()
	logger.SetFormattedOutputs(map[io.Writer]log.Formatter{output: log.NewRawFormatter()})
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

		assert.Equal(t, testCase.expected, output.String())

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

		assert.Equal(t, testCase.expected, output.String())

		output.Reset()
	}
}

func TestLoggerMultipleOutputs(t *testing.T) {
	logger, output1 := prepareTestLogger()
	output2 := &strings.Builder{}
	logger.SetFormattedOutputs(map[io.Writer]log.Formatter{output2: log.NewRawFormatter()})

	msg := "this is a test message"
	expected := msg + "\n"

	logger.Info(msg)

	assert.Equal(t, expected, output1.String())
	assert.Equal(t, expected, output2.String())
}

func TestLoggerBlacklist(t *testing.T) {
	logger, output := prepareTestLogger()
	logger.SetDebugMode(true)

	msg := "this is a test message"
	expected := msg + "\n"

	logger.BlacklistFunctions("TestLoggerBlacklist")
	logger.Debug(msg)
	assert.Equal(t, "", output.String())
	output.Reset()

	logger.ClearBlacklist()
	logger.BlacklistPackages("log_test")
	logger.Debug(msg)
	assert.Equal(t, "", output.String())
	output.Reset()

	logger.ClearBlacklist()
	logger.BlacklistPackages("github.com/timbasel/go-log/pkg/log_test")
	logger.Debug(msg)
	assert.Equal(t, "", output.String())
	output.Reset()

	logger.ClearBlacklist()
	logger.Debug(msg)
	assert.Equal(t, expected, output.String())
	output.Reset()
}

func TestLoggerWhitelist(t *testing.T) {
	logger, output := prepareTestLogger()
	logger.SetDebugMode(true)

	msg := "this is a test message"
	expected := msg + "\n"

	logger.Debug(msg)
	assert.Equal(t, expected, output.String())
	output.Reset()

	logger.WhitelistFunctions("TestLoggerWhitelist")
	logger.Debug(msg)
	assert.Equal(t, expected, output.String())
	output.Reset()

	logger.ClearWhitelist()
	logger.WhitelistFunctions("TestAnotherFunction")
	logger.Debug(msg)
	assert.Equal(t, "", output.String())
	output.Reset()

	logger.ClearWhitelist()
	logger.WhitelistPackages("github.com/timbasel/go-log/pkg/log_test")
	logger.Debug(msg)
	assert.Equal(t, expected, output.String())
	output.Reset()

	logger.ClearWhitelist()
	logger.WhitelistPackages("github.com/timbasel/go-log/pkg/another_pkg")
	logger.Debug(msg)
	assert.Equal(t, "", output.String())
	output.Reset()
}

func TestLoggerBlacklistAndWhitelist(t *testing.T) {
	logger, output := prepareTestLogger()
	logger.SetDebugMode(true)

	msg := "this is a test message"
	expected := msg + "\n"

	logger.Debug(msg)
	assert.Equal(t, expected, output.String())
	output.Reset()

	logger.WhitelistFunctions("TestLoggerBlacklistAndWhitelist")
	logger.BlacklistFunctions("TestLoggerBlacklistAndWhitelist")
	logger.Debug(msg)
	assert.Equal(t, "", output.String())
	output.Reset()

	logger.ClearBlacklist()
	logger.Debug(msg)
	assert.Equal(t, expected, output.String())
	output.Reset()

	logger.ClearBlacklist()
	logger.ClearWhitelist()
	logger.WhitelistPackages("github.com/timbasel/go-log/pkg/log")
	logger.BlacklistPackages("github.com/timbasel/go-log/pkg/log")
	logger.Debug(msg)
	assert.Equal(t, "", output.String())
	output.Reset()

	logger.ClearBlacklist()
	logger.Debug(msg)
	assert.Equal(t, expected, output.String())
	output.Reset()
}
