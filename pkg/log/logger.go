package log

import (
	"io"
	"os"
	"strings"
	"sync"
)

// A Logger writes formatted messages to the set of outputs.
type Logger struct {
	mutex sync.Mutex

	outputs map[io.Writer]Formatter

	debugMode bool
	blacklist []string
	whitelist []string
}

// NewLogger initializes a new logger.
// By default the logger will print its output the stdout console using the default formatter
func NewLogger() (logger *Logger) {
	return &Logger{
		outputs: map[io.Writer]Formatter{
			os.Stdout: NewDefaultFormatter(),
		},
		debugMode: false,
	}
}

// SetOutput adds the provided io.Writer to the loggers outputs using the default formatter
func (logger *Logger) SetOutput(output io.Writer) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.outputs[output] = NewDefaultFormatter()
}

// SetFormattedOutput adds the provided io.Writer to the loggers outputs with the provided formatter
func (logger *Logger) SetFormattedOutput(output io.Writer, formatter Formatter) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.outputs[output] = formatter
}

// ClearOutputs removes all set outputs from the logger
func (logger *Logger) ClearOutputs() {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.outputs = map[io.Writer]Formatter{}
}

// SetDebugMode toggles if debug messages should be written to the log outputs
func (logger *Logger) SetDebugMode(state bool) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.debugMode = state
}

// SetBlacklist adds the provided function/package name entry to the loggers blacklist for debug output
func (logger *Logger) SetBlacklist(entry string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.blacklist = append(logger.blacklist, entry)
}

// ClearBlacklist removes all entries from the loggers blacklist
func (logger *Logger) ClearBlacklist() {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.blacklist = []string{}
}

// SetWhitelist adds the provided function/package name entry to the loggers whitelist for debug output
func (logger *Logger) SetWhitelist(entry string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.whitelist = append(logger.whitelist, entry)
}

// ClearWhitelist removes all entries from the loggers whitelist
func (logger *Logger) ClearWhitelist() {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.whitelist = []string{}
}

// Error writes an error message to the log
func (logger *Logger) Error(msg ...string) {
	logger.write(ErrorLevel, strings.Join(msg, " "))
}

// Info writes an info message to the log
func (logger *Logger) Info(msg ...string) {
	logger.write(InfoLevel, strings.Join(msg, " "))
}

// Debug writes a debug message to the log
func (logger *Logger) Debug(msg ...string) {
	if logger.debugMode {
		caller := getCallersFullFunctionName()
		if logger.isWhitelisted(caller) && !logger.isBlacklisted(caller) {
			logger.write(DebugLevel, strings.Join(msg, " "))
		}
	}
}

func (logger *Logger) write(level Level, msg string) {
	for writer, formatter := range logger.outputs {
		writer.Write([]byte(formatter.Format(level, msg)))
	}
}

func (logger *Logger) isBlacklisted(caller string) bool {
	for _, entry := range logger.blacklist {
		if strings.Contains(caller, entry) {
			return true
		}
	}
	return false
}

func (logger *Logger) isWhitelisted(caller string) bool {
	if len(logger.whitelist) < 1 {
		return true
	}

	for _, entry := range logger.whitelist {
		if strings.Contains(caller, entry) {
			return true
		}
	}
	return false
}
