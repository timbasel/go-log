package log

import "io"

var defaultLogger = NewLogger()

// SetOutput adds the provided io.Writer to the output of the default logger using the default formatter
func SetOutput(output io.Writer) {
	defaultLogger.SetOutput(output)
}

// SetFormattedOutput adds the provided io.Writer to the output of the default logger using the provided formatter
func SetFormattedOutput(output io.Writer, formatter Formatter) {
	defaultLogger.SetFormattedOutput(output, formatter)
}

// ClearOutputs removes all outputs from the default logger
func ClearOutputs() {
	defaultLogger.ClearOutputs()
}

// Error writes an error message to the default log
func Error(msg ...string) {
	defaultLogger.Error(msg...)
}

// Info writes an info message to the default log
func Info(msg ...string) {
	defaultLogger.Info(msg...)
}

// Debug writes a debug message to the default log
func Debug(msg ...string) {
	defaultLogger.Debug(msg...)
}
