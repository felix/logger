package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"src.userspace.com.au/logger/message"
	"src.userspace.com.au/logger/writers/kv"
)

// Logger is a simple logger with optional structured.
type Logger struct {
	debug   bool
	name    string
	fields  map[string]string
	writers []message.Writer
	lock    *sync.RWMutex
}

// New creates a new logger instance
func New(opts ...Option) (*Logger, error) {
	l := &Logger{
		debug:   (os.Getenv("DEBUG") != ""),
		fields:  make(map[string]string),
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

// Log a message.
func (l *Logger) Log(args ...interface{}) *Logger {
	return l.log(args...)
}

func (l *Logger) log(args ...interface{}) *Logger {
	if len(args) == 0 {
		return l
	}
	m := message.Message{
		Name:    l.name,
		Time:    time.Now(),
		Content: toString(args...),
		Fields:  l.fields,
	}

	for _, w := range l.writers {
		w.Write(m)
	}
	return l
}

// toString converts interface to string
func toString(args ...interface{}) string {
	var buf strings.Builder
	last := len(args)
	for i, a := range args {
		buf.WriteString(fmt.Sprint(a))
		if i < last-1 {
			buf.WriteByte(' ')
		}
	}
	return buf.String()
}

// Info is an alias for Log.
func (l *Logger) Info(args ...interface{}) *Logger {
	return l.log(args...)
}

// Debug logs a debug message.
func (l *Logger) Debug(args ...interface{}) *Logger {
	if l.debug {
		return l.log(args...)
	}
	return l
}

// IsDebug determines the debug status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func (l *Logger) IsDebug() bool { return l.debug }

// Field enables setting or changing the default fields for a logger instance.
func (l *Logger) Field(k string, v interface{}) *Logger {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.fields[k] = toString(v)
	return l
}

// Fields enables setting or changing the default fields for a logger instance.
func (l *Logger) Fields(args ...interface{}) *Logger {
	l.lock.Lock()
	defer l.lock.Unlock()

	if len(args)%2 != 0 {
		args = append(args, "")
	}
	for i := 0; i < len(args); i += 2 {
		l.fields[toString(args[i])] = toString(args[i+1])
	}
	return l
}

// FieldMap enables setting or changing the default fields for a logger instance.
func (l *Logger) FieldMap(f map[string]interface{}) *Logger {
	l.lock.Lock()
	defer l.lock.Unlock()

	for k, v := range f {
		l.fields[k] = toString(v)
	}
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
