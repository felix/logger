package logger

import (
	"os"
	"sync"
	"time"

	"src.userspace.com.au/felix/logger/message"
	"src.userspace.com.au/felix/logger/writers/kv"
)

// Logger is a simple levelled logger.
type Logger struct {
	name       string
	min        message.Level
	fields     map[string]interface{}
	timeFormat string
	writers    []message.Writer
	lock       sync.Mutex
}

// New creates a new logger instance
func New(opts ...Option) (*Logger, error) {
	l := &Logger{
		min:        message.ERROR,
		fields:     make(map[string]interface{}),
		timeFormat: "2006-01-02T15:04:05.000Z0700",
	}

	// Apply variadic options
	if err := l.Configure(opts...); err != nil {
		return nil, err
	}

	// Add default writer
	if len(l.writers) == 0 {
		kv, err := kv.New(os.Stderr)
		if err != nil {
			return nil, err
		}
		l.writers = append(l.writers, kv)
	}
	return l, nil
}

// Configure applies settings to the logger.
func (l *Logger) Configure(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(l); err != nil {
			return err
		}
	}
	return nil
}

// Log for a logger instance
func (l *Logger) Log(lvl message.Level, msg string, args ...interface{}) {
	if l.min < lvl {
		return
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	m := message.Message{
		Name:    l.name,
		Time:    time.Now().Format(l.timeFormat),
		Level:   lvl,
		Content: msg,
		Fields:  l.fields,
		Extras:  args,
	}

	for _, w := range l.writers {
		w.Write(m)
	}
}

// Error logs an error message.
func (l *Logger) Error(msg string, args ...interface{}) { l.Log(message.ERROR, msg, args...) }

// Warn logs an information message.
func (l *Logger) Warn(msg string, args ...interface{}) { l.Log(message.WARN, msg, args...) }

// Info logs an information message.
func (l *Logger) Info(msg string, args ...interface{}) { l.Log(message.INFO, msg, args...) }

// Debug logs a debug message.
func (l *Logger) Debug(msg string, args ...interface{}) { l.Log(message.DEBUG, msg, args...) }

// IsWarn determines the info status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func (l *Logger) IsWarn() bool { return l.min >= message.WARN }

// IsInfo determines the info status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func (l *Logger) IsInfo() bool { return l.min >= message.INFO }

// IsDebug determines the debug status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func (l *Logger) IsDebug() bool { return l.min >= message.DEBUG }

// SetLevelString enables changing the minimum level for a logger instance.
func (l *Logger) SetLevelString(lvl string) {
	l.SetLevel(message.Levels[lvl])
}

// SetLevel enables changing the minimum level for a logger instance.
func (l *Logger) SetLevel(lvl message.Level) { l.min = lvl }

// SetField enables changing the default fields for a logger instance.
func (l *Logger) SetField(k string, v interface{}) {
	l.fields[k] = v
}

// SetName enables changing the name for a logger instance.
func (l *Logger) SetName(n string) {
	l.name = n
}

// GetNamed creates a new instance of a logger with a new name.
func (l *Logger) GetNamed(n string) *Logger {
	var nl = l
	if nl.name != "" {
		nl.SetName(nl.name + "." + n)
	} else {
		nl.SetName(n)
	}
	return nl
}
