package log

import "io"

var globalLogger = NewDefaultLogger()

// SetOutputs adds the provided io.Writers to the output of the global logger using the global formatter
func SetOutputs(output ...io.Writer) {
	globalLogger.SetOutputs(output...)
}

// SetFormattedOutputs adds the provided io.Writer to the output of the global logger using the provided formatter
func SetFormattedOutputs(outputs map[io.Writer]Formatter) {
	globalLogger.SetFormattedOutputs(outputs)
}

// SetDebugMode toggles if debug messages are written to the global loggers outputs
func SetDebugMode(state bool) {
	globalLogger.SetDebugMode(state)
}

// ClearOutputs removes all outputs from the global logger
func ClearOutputs() {
	globalLogger.ClearOutputs()
}

// BlacklistFunctions adds the provided function names to the global loggers debug output blacklist
func BlacklistFunctions(names ...string) {
	globalLogger.BlacklistFunctions(names...)
}

// BlacklistPackages adds the provided package names to the global loggers debug output blacklist
func BlacklistPackages(names ...string) {
	globalLogger.BlacklistPackages(names...)
}

// ClearBlacklist removes all entries from the global loggers blacklist
func ClearBlacklist() {
	globalLogger.ClearBlacklist()
}

// WhitelistFunctions adds the provided function names to the global loggers debug output whitelist
func WhitelistFunctions(names ...string) {
	globalLogger.WhitelistFunctions(names...)
}

// WhitelistPackages adds the provided package name to the global loggers debug output whitelist
func WhitelistPackages(names ...string) {
	globalLogger.WhitelistPackages(names...)
}

// ClearWhitelist removes all entries from the global loggers whitelist
func ClearWhitelist() {
	globalLogger.ClearWhitelist()
}

// Error writes an error message to the global log
func Error(msg ...string) {
	globalLogger.Error(msg...)
}

// Errorf writes a formatted error message to the global log
func Errorf(format string, arguments ...interface{}) {
	globalLogger.Errorf(format, arguments...)
}

// Info writes an info message to the global log
func Info(msg ...string) {
	globalLogger.Info(msg...)
}

// Infof writes a formatted error message to the global log
func Infof(format string, arguments ...interface{}) {
	globalLogger.Infof(format, arguments...)
}

// Debug writes a debug message to the global log
func Debug(msg ...string) {
	globalLogger.Debug(msg...)
}

// Debugf writes a formatted error message to the global log
func Debugf(format string, arguments ...interface{}) {
	globalLogger.Debugf(format, arguments...)
}

// XDebug disables the Debug function
func XDebug(msg ...string) {}

// XDebugf disables the Debugf function
