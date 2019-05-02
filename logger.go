package logger

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// Logger is a simple levelled logger.
type Logger struct {
	name       string
	min        Level
	fields     []interface{}
	timeFormat string
	formatter  MessageWriter
	out        io.Writer
}

// New creates a new logger instance
func New(opts ...Option) (*Logger, error) {
	l := &Logger{
		min:        ERROR,
		timeFormat: "2006-01-02T15:04:05.000Z0700",
		formatter:  new(KeyValue),
		out:        bufio.NewWriter(os.Stderr),
	}

	// Apply variadic options
	if err := l.Configure(opts...); err != nil {
		return nil, err
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
func (l *Logger) Log(lvl Level, msg string, args ...interface{}) {
	if l.min < lvl {
		return
	}

	// l.lock.Lock()
	// defer l.lock.Unlock()

	m := Message{
		Name:    l.name,
		Time:    time.Now().Format(l.timeFormat),
		Level:   lvl,
		Content: msg,
		Fields:  l.fields,
		Extras:  args,
	}

	l.formatter.Write(l.out, m)
	l.out.Write([]byte{'\n'})
}

// Error logs an error message.
func (l Logger) Error(msg string, args ...interface{}) { l.Log(ERROR, msg, args...) }

// Info logs an information message.
func (l Logger) Info(msg string, args ...interface{}) { l.Log(INFO, msg, args...) }

// Debug logs a debug message.
func (l Logger) Debug(msg string, args ...interface{}) { l.Log(DEBUG, msg, args...) }

// IsInfo determines the info status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func (l Logger) IsInfo() bool { return l.min >= INFO }

// IsDebug determines the debug status for a logger instance.
// Use this to conditionally execute blocks of code depending on the log verbosity.
func (l Logger) IsDebug() bool { return l.min >= DEBUG }

// SetLevel enables changing the minimum level for a logger instance.
func (l *Logger) SetLevel(lvl Level) { l.min = lvl }

// SetField enables changing the default fields for a logger instance.
func (l *Logger) SetField(k string, v interface{}) {
	l.fields = append(l.fields, k, v)
}

// SetName enables changing the name for a logger instance.
func (l *Logger) SetName(n string) {
	l.name = n
}

// GetNamed creates a new instance of a logger with a new name.
func (l Logger) GetNamed(n string) *Logger {
	var nl = l
	if nl.name != "" {
		nl.SetName(nl.name + "." + n)
	} else {
		nl.SetName(n)
	}
	return &nl
}

// ToString converts interface to string
func ToString(v interface{}) string {
	switch c := v.(type) {
	case string:
		return c
	case int:
		return strconv.FormatInt(int64(c), 10)
	case int64:
		return strconv.FormatInt(int64(c), 10)
	case int32:
		return strconv.FormatInt(int64(c), 10)
	case int16:
		return strconv.FormatInt(int64(c), 10)
	case int8:
		return strconv.FormatInt(int64(c), 10)
	case uint:
		return strconv.FormatUint(uint64(c), 10)
	case uint64:
		return strconv.FormatUint(uint64(c), 10)
	case uint32:
		return strconv.FormatUint(uint64(c), 10)
	case uint16:
		return strconv.FormatUint(uint64(c), 10)
	case uint8:
		return strconv.FormatUint(uint64(c), 10)
	case float32:
		return strconv.FormatFloat(float64(c), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(c, 'g', -1, 64)
	case bool:
		return strconv.FormatBool(c)
	default:
		return fmt.Sprintf("%v", c)
	}
}
