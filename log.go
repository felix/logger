package logger

import (
	"bufio"
	"os"
	"sync"
	"time"
)

type logger struct {
	name       string
	level      Level
	fields     []interface{}
	timeFormat string
	lock       *sync.Mutex
	formatter  MessageWriter
	out        *bufio.Writer
}

// New creates a new logger
func New(opts *Options) Logger {
	if opts == nil {
		opts = &Options{}
	}

	output := opts.Output
	if output == nil {
		output = os.Stderr
	}

	timeFormat := opts.TimeFormat
	if timeFormat == "" {
		timeFormat = DefaultTimeFormat
	}

	level := opts.Level
	if level == NoLevel {
		level = Info
	}

	l := logger{
		name:       opts.Name,
		lock:       new(sync.Mutex),
		level:      level,
		timeFormat: timeFormat,
		out:        bufio.NewWriter(output),
	}

	l.formatter = opts.Formatter
	if l.formatter == nil {
		l.formatter = NewDefaultWriter()
	}

	return &l
}

// Log is a generic logger function
func (l logger) Log(lvl Level, args ...interface{}) {
	if lvl < l.level {
		return
	}

	ts := time.Now()

	l.lock.Lock()
	defer l.lock.Unlock()

	// Place fields at the end
	args = append(args, l.fields...)

	msg := Message{
		Name:   l.name,
		Time:   ts.Format(l.timeFormat),
		Level:  lvl,
		Fields: make([]interface{}, 0),
	}

	// Allow for map arguments
	for _, f := range args {
		switch c := f.(type) {
		case map[string]string:
			for k, v := range c {
				msg.Fields = append(msg.Fields, k, v)
			}
		case map[string]int:
			for k, v := range c {
				msg.Fields = append(msg.Fields, k, v)
			}
		case map[int]string:
			for k, v := range c {
				msg.Fields = append(msg.Fields, k, v)
			}
		case map[string]interface{}:
			for k, v := range c {
				msg.Fields = append(msg.Fields, k, v)
			}
		default:
			msg.Fields = append(msg.Fields, c)
		}
	}

	l.formatter.Write(l.out, msg)
	l.out.WriteByte('\n')

	l.out.Flush()
}

// Convenience functions for logging at levels
func (l logger) Debug(args ...interface{}) { l.Log(Debug, args...) }
func (l logger) Warn(args ...interface{})  { l.Log(Warn, args...) }
func (l logger) Error(args ...interface{}) { l.Log(Error, args...) }
func (l logger) Info(args ...interface{})  { l.Log(Info, args...) }

// Test for current logging level
func (l logger) IsLevel(lvl Level) bool { return l.level <= lvl }
func (l logger) IsDebug() bool          { return l.IsLevel(Debug) }
func (l logger) IsInfo() bool           { return l.IsLevel(Info) }
func (l logger) IsWarn() bool           { return l.IsLevel(Warn) }
func (l logger) IsError() bool          { return l.IsLevel(Error) }

// WithFields sets the default fields for a new logger
func (l *logger) WithFields(args ...interface{}) Logger {
	var nl = *l
	nl.fields = append(nl.fields, args...)
	return &nl
}

// Named sets the name for a new logger
func (l *logger) Named(name string) Logger {
	var nl = *l
	if nl.name != "" {
		nl.name = nl.name + "." + name
	} else {
		nl.name = name
	}
	return &nl
}
