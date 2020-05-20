package logger

import (
	"src.userspace.com.au/logger/message"
)

var std *Logger

func init() {
	std, _ = New()
}

// Log a message.
func Log(args ...interface{}) { std.Log(args...) }

// Info is an alias for Log.
func Info(args ...interface{}) *Logger { return std.Log(args...) }

// Debug logs a debug message.
func Debug(args ...interface{}) { std.Debug(args...) }

// IsDebug determines the debug status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func IsDebug() bool { return std.IsDebug() }

// Field enables changing the default fields for a logger instance.
func Field(k string, v interface{}) *Logger { return std.Field(k, v) }

// Fields enables setting or changing the default fields for a logger instance.
func Fields(args ...interface{}) *Logger { return std.Fields(args...) }

// FieldMap enables setting or changing the default fields for a logger instance.
func FieldMap(f map[string]interface{}) *Logger { return std.FieldMap(f) }

// Named creates a new instance of a logger with a new name.
func Named(n string) *Logger { return std.Named(n) }

// SetWriter sets the writer for the default logger.
func SetWriter(w message.Writer) { std.writers = []message.Writer{w} }

// WithSource annotates the log message with the source that called logger.
func WithSource() *Logger { return std.WithSource() }
