package infrastructure

import (
	"fmt"
	"log"
	"os"
)

// DefaultLogLevel is the log level to default to if given log level is incorrect
const DefaultLogLevel = 3

// Logger is an abstraction to do logging in the application
type Logger struct {
	logLevel int
}

// LogLevels are the defined log levels
var LogLevels = [6]string{"FATAL", "ERROR", "WARN", "INFO", "DEBUG", "TRACE"}

// DefaultLogger returns a logger configured at the default log level
func DefaultLogger() *Logger {
	return &Logger{
		logLevel: DefaultLogLevel,
	}
}

// NewLogger is the constructor
func NewLogger(logLevel int) *Logger {
	if logLevel < 0 || logLevel > len(LogLevels) {
		logLevel = DefaultLogLevel
		tempLogger := DefaultLogger()
		tempLogger.Warn(fmt.Sprintf("unknown log level configured, falling back to default log level [%d]", DefaultLogLevel))
	}

	return &Logger{
		logLevel: logLevel,
	}
}

func (l *Logger) log(level int, message string) {
	if level <= l.logLevel {
		log.Printf(`[%s] "%s"`, LogLevels[level], message)
	}
}

// Fatal logs at log level 0 and exits the application
// Used for known errors that can't be recovered from
func (l *Logger) Fatal(message string, exitCode int) {
	if exitCode <= 0 {
		exitCode = 1
	}

	l.log(0, message)
	os.Exit(exitCode)
}

// Error logs at log level 1
// Used for known errors that can be recovered from
func (l *Logger) Error(message string) {
	l.log(1, message)
}

// Warn logs at log level 2
// Used for deprecation info and config errors
func (l *Logger) Warn(message string) {
	l.log(2, message)
}

// Info logs at log level 3
// Used for info at request-response level
func (l *Logger) Info(message string) {
	l.log(3, message)
}

// Debug logs at log level 4
// Used to debugging information required in production
func (l *Logger) Debug(message string) {
	l.log(4, message)
}

// Trace logs at log level 5
func (l *Logger) Trace(message string) {
	l.log(5, message)
}
