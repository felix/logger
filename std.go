package logger

import (
	"src.userspace.com.au/logger/message"
)

var std *Logger

func init() {
	std, _ = New(Level(message.WARN))
}

// Error logs an error message.
func Error(msg string, args ...interface{}) { std.Error(msg, args...) }

// Warn logs an information message.
func Warn(msg string, args ...interface{}) { std.Warn(msg, args...) }

// Info logs an information message.
func Info(msg string, args ...interface{}) { std.Info(msg, args...) }

// Debug logs a debug message.
func Debug(msg string, args ...interface{}) { std.Debug(msg, args...) }

// IsWarn determines the info status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func IsWarn() bool { return std.IsWarn() }

// IsInfo determines the info status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func IsInfo() bool { return std.IsInfo() }

// IsDebug determines the debug status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func IsDebug() bool { return std.IsDebug() }

// SetLevelAsString enables changing the minimum level for a logger instance.
func SetLevelAsString(lvl string) { std.SetLevelAsString(lvl) }

// SetLevel enables changing the minimum level for a logger instance.
func SetLevel(lvl message.Level) { std.SetLevel(lvl) }

// Field enables changing the default fields for a logger instance.
func Field(k string, v interface{}) *Logger { return std.Field(k, v) }

// Named creates a new instance of a logger with a new name.
func Named(n string) *Logger { return std.Named(n) }

// SetWriter sets the writer for the default logger.
func SetWriter(w message.Writer) { std.writers = []message.Writer{w} }
