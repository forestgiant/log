package logger

import (
	"io"
	"os"
)

// Logger implements the Logger interface as well as some helper methods for leveled logging.
type Logger struct {
	Writer    io.Writer
	Formatter Formatter
	context   []interface{}
}

// Log the provided key value pairs
func (l *Logger) Log(keyvals ...interface{}) error {
	var writer io.Writer
	var formatter Formatter

	if l.Writer == nil {
		writer = os.Stdout
	} else {
		writer = l.Writer
	}

	if l.Formatter == nil {
		formatter = JSONFormatter{}
	} else {
		formatter = l.Formatter
	}

	return formatter.Format(writer, append(l.context, keyvals...)...)
}

// With returns a copy of the logger with the provided key-value pairs to the backing context
func (l Logger) With(keyvals ...interface{}) Logger {
	return Logger{
		Writer:    l.Writer,
		Formatter: l.Formatter,
		context:   append(l.context, keyvals...),
	}
}

// Debug logs the provided key value pairs, adding the debug log level
func (l *Logger) Debug(message string, keyvals ...interface{}) error {
	return l.Log(append([]interface{}{"level", "debug", "msg", message}, keyvals...)...)
}

// Info logs the provided key value pairs, adding the info log level
func (l *Logger) Info(message string, keyvals ...interface{}) error {
	return l.Log(append([]interface{}{"level", "info", "msg", message}, keyvals...)...)
}

// Emergency logs the provided key value pairs, adding the emergency log level
func (l *Logger) Emergency(message string, keyvals ...interface{}) error {
	return l.Log(append([]interface{}{"level", "emergency", "msg", message}, keyvals...)...)
}

// Alert logs the provided key value pairs, adding the alert log level
func (l *Logger) Alert(message string, keyvals ...interface{}) error {
	return l.Log(append([]interface{}{"level", "alert", "msg", message}, keyvals...)...)
}

// Critical logs the provided key value pairs, adding the critical log level
func (l *Logger) Critical(message string, keyvals ...interface{}) error {
	return l.Log(append([]interface{}{"level", "critical", "msg", message}, keyvals...)...)
}

// Error logs the provided key value pairs, adding the error log level
func (l *Logger) Error(message string, keyvals ...interface{}) error {
	return l.Log(append([]interface{}{"level", "error", "msg", message}, keyvals...)...)
}

// Notice logs the provided key value pairs, adding the notice log level
func (l *Logger) Notice(message string, keyvals ...interface{}) error {
	return l.Log(append([]interface{}{"level", "notice", "msg", message}, keyvals...)...)
}

// Warning logs the provided key value pairs, adding the warning log level
func (l *Logger) Warning(message string, keyvals ...interface{}) error {
	return l.Log(append([]interface{}{"level", "warning", "msg", message}, keyvals...)...)
}
