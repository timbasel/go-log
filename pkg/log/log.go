package log

import "io"

var defaultLogger = NewDefaultLogger()

// SetOutputs adds the provided io.Writers to the output of the default logger using the default formatter
func SetOutputs(output ...io.Writer) {
	defaultLogger.SetOutputs(output...)
}

// SetFormattedOutputs adds the provided io.Writer to the output of the default logger using the provided formatter
func SetFormattedOutputs(outputs map[io.Writer]Formatter) {
	defaultLogger.SetFormattedOutputs(outputs)
}

// SetDebugMode toggles if debug messages are written to the default loggers outputs
func SetDebugMode(state bool) {
	defaultLogger.SetDebugMode(state)
}

// ClearOutputs removes all outputs from the default logger
func ClearOutputs() {
	defaultLogger.ClearOutputs()
}

// BlacklistFunctions adds the provided function names to the default loggers debug output blacklist
func BlacklistFunctions(names ...string) {
	defaultLogger.BlacklistFunctions(names...)
}

// BlacklistPackages adds the provided package names to the default loggers debug output blacklist
func BlacklistPackages(names ...string) {
	defaultLogger.BlacklistPackages(names...)
}

// ClearBlacklist removes all entries from the default loggers blacklist
func ClearBlacklist() {
	defaultLogger.ClearBlacklist()
}

// WhitelistFunctions adds the provided function names to the default loggers debug output whitelist
func WhitelistFunctions(names ...string) {
	defaultLogger.WhitelistFunctions(names...)
}

// WhitelistPackages adds the provided package name to the default loggers debug output whitelist
func WhitelistPackages(names ...string) {
	defaultLogger.WhitelistPackages(names...)
}

// ClearWhitelist removes all entries from the loggers whitelist
func ClearWhitelist() {
	defaultLogger.ClearWhitelist()
}

// Error writes an error message to the default log
func Error(msg ...string) {
	defaultLogger.Error(msg...)
}

// Errorf writes a formatted error message to the default log
func Errorf(format string, arguments ...interface{}) {
	defaultLogger.Errorf(format, arguments...)
}

// Info writes an info message to the default log
func Info(msg ...string) {
	defaultLogger.Info(msg...)
}

// Infof writes a formatted error message to the default log
func Infof(format string, arguments ...interface{}) {
	defaultLogger.Infof(format, arguments...)
}

// Debug writes a debug message to the default log
func Debug(msg ...string) {
	defaultLogger.Debug(msg...)
}

// Debugf writes a formatted error message to the default log
func Debugf(format string, arguments ...interface{}) {
	defaultLogger.Debugf(format, arguments...)
}
