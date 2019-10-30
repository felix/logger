package logger

import (
	"os"
	"sync"
	"time"

	"src.userspace.com.au/logger/message"
	"src.userspace.com.au/logger/writers/kv"
)

// Logger is a simple levelled logger.
type Logger struct {
	name    string
	min     message.Level
	fields  map[string]interface{}
	writers []message.Writer
	lock    *sync.RWMutex
}

// New creates a new logger instance
func New(opts ...Option) (*Logger, error) {
	l := &Logger{
		min:     message.WARN,
		fields:  make(map[string]interface{}),
		lock:    new(sync.RWMutex),
		writers: []message.Writer{},
	}

	// Apply variadic options
	for _, opt := range opts {
		if err := opt(l); err != nil {
			return nil, err
		}
	}

	// Add default writer
	if len(l.writers) == 0 {
		kv, err := kv.New(kv.SetOutput(os.Stderr))
		if err != nil {
			return nil, err
		}
		l.writers = []message.Writer{kv}
	}
	return l, nil
}

// Log a message with no level.
func (l *Logger) Log(msg string, args ...interface{}) *Logger {
	l.LogAtLevel(message.NONE, msg, args...)
	return l
}

// LogAtLevel logs a message with a specified level.
func (l *Logger) LogAtLevel(lvl message.Level, msg string, args ...interface{}) *Logger {
	if l.min < lvl {
		return l
	}

	l.lock.RLock()
	defer l.lock.RUnlock()

	m := message.Message{
		Name:    l.name,
		Time:    time.Now(),
		Level:   lvl,
		Content: msg,
		Fields:  l.fields,
		Extras:  args,
	}

	for _, w := range l.writers {
		w.Write(m)
	}
	return l
}

// Error logs an error message.
func (l *Logger) Error(msg string, args ...interface{}) *Logger {
	return l.LogAtLevel(message.ERROR, msg, args...)
}

// Warn logs an information message.
func (l *Logger) Warn(msg string, args ...interface{}) *Logger {
	return l.LogAtLevel(message.WARN, msg, args...)
}

// Info logs an information message.
func (l *Logger) Info(msg string, args ...interface{}) *Logger {
	return l.LogAtLevel(message.INFO, msg, args...)
}

// Debug logs a debug message.
func (l *Logger) Debug(msg string, args ...interface{}) *Logger {
	return l.LogAtLevel(message.DEBUG, msg, args...)
}

// IsWarn determines the info status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func (l *Logger) IsWarn() bool { return l.min >= message.WARN }

// IsInfo determines the info status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func (l *Logger) IsInfo() bool { return l.min >= message.INFO }

// IsDebug determines the debug status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func (l *Logger) IsDebug() bool { return l.min >= message.DEBUG }

// SetLevelAsString enables changing the minimum level for a logger instance.
func (l *Logger) SetLevelAsString(lvl string) *Logger {
	l.SetLevel(message.Levels[lvl])
	return l
}

// SetLevel enables changing the minimum level for a logger instance.
func (l *Logger) SetLevel(lvl message.Level) *Logger {
	l.min = lvl
	return l
}

// Field enables changing the default fields for a logger instance.
func (l *Logger) Field(k string, v interface{}) *Logger {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.fields[k] = v
	return l
}

// Named creates a new instance of a logger with a new name.
func (l *Logger) Named(n string) *Logger {
	nl := *l
	if l.name != "" {
		nl.name = l.name + "." + n
	} else {
		nl.name = n
	}
	return &nl
}
