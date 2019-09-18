package logger

import (
	"src.userspace.com.au/felix/logger/message"
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

// Level configures the minimum level to log.
func Level(lvl message.Level) Option {
	return func(l *Logger) error {
		l.SetLevel(lvl)
		return nil
	}
}

// LevelAsString configures the minimum level to log.
func LevelAsString(lvl string) Option {
	return func(l *Logger) error {
		l.SetLevelAsString(lvl)
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
