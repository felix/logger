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

// LevelString configures the minimum level to log.
func LevelString(lvl string) Option {
	return func(l *Logger) error {
		l.SetLevelString(lvl)
		return nil
	}
}

// Field configures an initial field of the logger.
func Field(k string, v interface{}) Option {
	return func(l *Logger) error {
		l.SetField(k, v)
		return nil
	}
}

// Fields configures an initial set of fields of the logger.
func Fields(f map[string]interface{}) Option {
	return func(l *Logger) error {
		for k, v := range f {
			l.SetField(k, v)
		}
		return nil
	}
}

// Name configures the name of the logger.
func Name(n string) Option {
	return func(l *Logger) error {
		l.SetName(n)
		return nil
	}
}
