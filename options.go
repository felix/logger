package logger

import (
	"io"
)

// An Option configures a logger
type Option func(*Logger) error

// SetFormatter sets the output formatter for the logger.
func SetFormatter(f MessageWriter) Option {
	return func(l *Logger) error {
		l.formatter = f
		return nil
	}
}

// SetOutput sets the output for the logger.
func SetOutput(w io.Writer) Option {
	return func(l *Logger) error {
		l.out = w
		return nil
	}
}

// SetLevel configures the minimum level to log.
func SetLevel(lvl Level) Option {
	return func(l *Logger) error {
		l.SetLevel(lvl)
		return nil
	}
}

// SetTimeFormat configures the format used for timestamps.
func SetTimeFormat(f string) Option {
	return func(l *Logger) error {
		l.timeFormat = f
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
