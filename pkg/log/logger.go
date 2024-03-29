package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// A Logger writes formatted messages to the set of outputs.
type Logger struct {
	mutex sync.Mutex

	outputs map[io.Writer]Formatter

	debugMode          bool
	blacklistFunctions []string
	blacklistPackages  []string
	whitelistFunctions []string
	whitelistPackages  []string
}

// NewLogger initializes a new empty logger with no outputs configured
func NewLogger() (logger *Logger) {
	return &Logger{
		outputs:            map[io.Writer]Formatter{},
		blacklistFunctions: []string{},
		blacklistPackages:  []string{},
		whitelistFunctions: []string{},
		whitelistPackages:  []string{},
	}
}

// NewDefaultLogger initializes a new logger configured to write its output to the stdout console using the default formatter
func NewDefaultLogger() (logger *Logger) {
	logger = NewLogger()
	logger.SetOutputs(os.Stdout)
	return logger
}

// SetOutputs adds the provided io.Writers to the loggers outputs using the default formatter
func (logger *Logger) SetOutputs(outputs ...io.Writer) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, output := range outputs {
		logger.outputs[output] = NewDefaultFormatter()
	}
}

// SetFormattedOutputs adds the provided io.Writers to the loggers outputs with the provided custom formatters
func (logger *Logger) SetFormattedOutputs(outputs map[io.Writer]Formatter) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for writer, formatter := range outputs {
		logger.outputs[writer] = formatter
	}
}

// ClearOutputs removes all set outputs from the logger
func (logger *Logger) ClearOutputs() {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.outputs = map[io.Writer]Formatter{}
}

// SetDebugMode toggles if debug messages are written to the loggers outputs
func (logger *Logger) SetDebugMode(state bool) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.debugMode = state
}

// BlacklistFunctions adds the provided function names to the loggers debug output blacklist
func (logger *Logger) BlacklistFunctions(names ...string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, name := range names {
		logger.blacklistFunctions = append(logger.blacklistFunctions, name)
	}
}

// BlacklistPackages adds the provided package names to the loggers debug output blacklist
func (logger *Logger) BlacklistPackages(names ...string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, name := range names {
		logger.blacklistPackages = append(logger.blacklistPackages, name)
	}
}

// ClearBlacklist removes all entries from the loggers blacklist
func (logger *Logger) ClearBlacklist() {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.blacklistFunctions = []string{}
	logger.blacklistPackages = []string{}
}

// WhitelistFunctions adds the provided function names to the loggers debug output whitelist
func (logger *Logger) WhitelistFunctions(names ...string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, name := range names {
		logger.whitelistFunctions = append(logger.whitelistFunctions, name)
	}
}

// WhitelistPackages adds the provided package name to the loggers debug output whitelist
func (logger *Logger) WhitelistPackages(names ...string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, name := range names {
		logger.whitelistPackages = append(logger.whitelistPackages, name)
	}
}

// ClearWhitelist removes all entries from the loggers whitelist
func (logger *Logger) ClearWhitelist() {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.whitelistFunctions = []string{}
	logger.whitelistPackages = []string{}
}

// Error writes an error message to the log
func (logger *Logger) Error(msg ...string) {
	logger.write(ErrorLevel, strings.Join(msg, " "))
}

// Errorf writes a formatted error message to the log
func (logger *Logger) Errorf(format string, arguments ...interface{}) {
	logger.write(ErrorLevel, fmt.Sprintf(format, arguments...))
}

// Info writes an info message to the log
func (logger *Logger) Info(msg ...string) {
	logger.write(InfoLevel, strings.Join(msg, " "))
}

// Infof writes a formatted info message to the log
func (logger *Logger) Infof(format string, arguments ...interface{}) {
	logger.write(InfoLevel, fmt.Sprintf(format, arguments...))
}

// Debug writes a debug message to the log
func (logger *Logger) Debug(msg ...string) {
	if logger.debugMode {
		if logger.isWhitelisted() && !logger.isBlacklisted() {
			logger.write(DebugLevel, strings.Join(msg, " "))
		}
	}
}

// Debugf writes a formatted debug message to the log
func (logger *Logger) Debugf(format string, arguments ...interface{}) {
	if logger.debugMode {
		if logger.isWhitelisted() && !logger.isBlacklisted() {
			logger.write(DebugLevel, fmt.Sprintf(format, arguments...))
		}
	}
}

// XDebug disables the Debug method
func (logger *Logger) XDebug(msg ...string) {}

// XDebugf disables the Debugf method
func (logger *Logger) XDebugf(format string, arguments ...interface{}) {}

func (logger *Logger) write(level Level, msg string) {
	for writer, formatter := range logger.outputs {
		writer.Write([]byte(formatter.Format(level, msg)))
	}
}

func (logger *Logger) isBlacklisted() bool {
	functionListed := searchFunctionList(logger.blacklistFunctions, getCallersFunctionName())
	packageListed := searchPackageList(logger.blacklistPackages, getCallersPackageName())
	return functionListed || packageListed
}

func (logger *Logger) isWhitelisted() bool {
	if len(logger.whitelistFunctions) < 1 && len(logger.whitelistPackages) < 1 {
		return true
	}

	functionListed := searchFunctionList(logger.whitelistFunctions, getCallersFunctionName())
	packageListed := searchPackageList(logger.whitelistPackages, getCallersPackageName())
	return functionListed || packageListed
}

func searchFunctionList(list []string, name string) bool {
	for _, item := range list {
		if item == name {
			return true
		}
	}
	return false
}

func searchPackageList(list []string, name string) bool {
	for _, item := range list {
		if strings.HasPrefix(name, item) || strings.HasSuffix(name, item) {
			return true
		}
	}
	return false
}
