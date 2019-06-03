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

	debugMode          bool
	blacklistPackages  []string
	blacklistFunctions []string
	whitelistPackages  []string
	whitelistFunctions []string
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

// SetDebugMode toggles if debug messages should be written to the log outputs
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

// Info writes an info message to the log
func (logger *Logger) Info(msg ...string) {
	logger.write(InfoLevel, strings.Join(msg, " "))
}

// Debug writes a debug message to the log
func (logger *Logger) Debug(msg ...string) {
	if logger.debugMode {
		if logger.isWhitelisted() && !logger.isBlacklisted() {
			logger.write(DebugLevel, strings.Join(msg, " "))
		}
	}
}

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
