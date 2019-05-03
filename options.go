package logger

import (
	"src.userspace.com.au/felix/logger/message"
)

// An Option configures a logger
type Option func(*Logger) error

// AddWriter add an output formatter for the logger.
func AddWriter(f message.Writer) Option {
	return func(l *Logger) error {
		l.writers = append(l.writers, f)
		return nil
	}
}

// SetLevel configures the minimum level to log.
func SetLevel(lvl message.Level) Option {
	return func(l *Logger) error {
		l.SetLevel(lvl)
		return nil
	}
}

// SetField configures an initial field of the logger.
func SetField(k string, v interface{}) Option {
	return func(l *Logger) error {
		l.SetField(k, v)
		return nil
	}
}

// SetFields configures an initial set of fields of the logger.
func SetFields(f map[string]interface{}) Option {
	return func(l *Logger) error {
		for k, v := range f {
			l.SetField(k, v)
		}
		return nil
	}
}

// SetName configures the name of the logger.
func SetName(n string) Option {
	return func(l *Logger) error {
		l.SetName(n)
		return nil
	}
}
