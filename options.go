package logger

import (
	"os"

	"src.userspace.com.au/logger/message"
)

// An Option configures a logger
type Option func(*Logger) error

// Writer add an output formatter for the logger.
func Writer(f message.Writer) Option {
	return func(l *Logger) error {
		l.writers = append(l.writers, f)
		return nil
	}
}

// ForceDebug sets debug.
func ForceDebug(b bool) Option {
	return func(l *Logger) error {
		l.debug = b
		return nil
	}
}

// DebugEnvVar sets debug if the envvar 'v' is not empty.
func DebugEnvVar(v string) Option {
	return func(l *Logger) error {
		l.debug = (os.Getenv(v) != "")
		return nil
	}
}

// Name configures the name of the logger.
func Name(n string) Option {
	return func(l *Logger) error {
		l.name = n
		return nil
	}
}
