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

// BlacklistFunction adds the provided function name to the loggers blacklist for debug output
func (logger *Logger) BlacklistFunction(name string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.blacklistFunctions = append(logger.blacklistFunctions, name)
}

// BlacklistPackage adds the provided package name to the loggers blacklist for debug output
func (logger *Logger) BlacklistPackage(name string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.blacklistPackages = append(logger.blacklistPackages, name)
}

// ClearBlacklist removes all entries from the loggers blacklist
func (logger *Logger) ClearBlacklist() {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.blacklistFunctions = []string{}
	logger.blacklistPackages = []string{}
}

// WhitelistFunction adds the provided function name to the loggers debug output whitelist
func (logger *Logger) WhitelistFunction(name string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.whitelistFunctions = append(logger.whitelistFunctions, name)
}

// WhitelistPackage adds the provided package name to the loggers whitelist for debug output
func (logger *Logger) WhitelistPackage(name string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.whitelistPackages = append(logger.whitelistPackages, name)
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
