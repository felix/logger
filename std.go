package logger

import (
	"src.userspace.com.au/felix/logger/message"
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

// SetLevelString enables changing the minimum level for a logger instance.
func SetLevelString(lvl string) { std.SetLevelString(lvl) }

// SetLevel enables changing the minimum level for a logger instance.
func SetLevel(lvl message.Level) { std.SetLevel(lvl) }

// SetField enables changing the default fields for a logger instance.
func SetField(k string, v interface{}) { std.SetField(k, v) }

// SetName enables changing the name for a logger instance.
func SetName(n string) { std.SetName(n) }

// GetNamed creates a new instance of a logger with a new name.
func GetNamed(n string) *Logger { return std.GetNamed(n) }
