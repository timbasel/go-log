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

	outputs   map[io.Writer]Formatter
	debugMode bool
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

// Error writes an error message to the log
func (logger *Logger) Error(msg ...string) {
	logger.write(Error, strings.Join(msg, " "))
}

// Info writes an info message to the log
func (logger *Logger) Info(msg ...string) {
	logger.write(Info, strings.Join(msg, " "))
}

// Debug writes a debug message to the log
func (logger *Logger) Debug(msg ...string) {
	if logger.debugMode {
		logger.write(Debug, strings.Join(msg, " "))
	}
}

func (logger *Logger) write(level Level, msg string) {
	for writer, formatter := range logger.outputs {
		writer.Write([]byte(formatter.Format(level, msg)))
	}
}
