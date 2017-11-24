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

func (l logger) Log(lvl Level, args ...interface{}) {
	if lvl < l.level {
		return
	}

	ts := time.Now()

	l.lock.Lock()
	defer l.lock.Unlock()

	msg := Message{
		Name:   l.name,
		Time:   ts.Format(l.timeFormat),
		Level:  lvl,
		Fields: make([]interface{}, 0),
	}

	offset := 0
	if len(args)%2 != 0 {
		msg.Fields = append(msg.Fields, "message", args[0])
		offset = 1
	}
	for i := offset; i < len(args); i = i + 2 {
		msg.Fields = append(msg.Fields, ToString(args[i]), args[i+1])
	}

	l.formatter.Write(l.out, msg)

	l.out.Flush()
}

func (l logger) Debug(args ...interface{}) {
	l.Log(Debug, args...)
}

func (l logger) Warn(args ...interface{}) {
	l.Log(Warn, args...)
}

func (l logger) Error(args ...interface{}) {
	l.Log(Error, args...)
}

func (l logger) Info(args ...interface{}) {
	l.Log(Info, args...)
}

func (l logger) IsLevel(lvl Level) bool {
	return l.level <= lvl
}

func (l *logger) IsDebug() bool { return l.IsLevel(Debug) }
func (l *logger) IsInfo() bool  { return l.IsLevel(Info) }
func (l *logger) IsWarn() bool  { return l.IsLevel(Warn) }
func (l *logger) IsError() bool { return l.IsLevel(Error) }

func (l *logger) WithFields(args ...interface{}) Logger {
	var nl logger = *l
	nl.fields = append(nl.fields, args...)
	return &nl
}

func (l *logger) Named(name string) Logger {
	var nl logger = *l
	if nl.name != "" {
		nl.name = nl.name + "." + name
	} else {
		nl.name = name
	}
	return &nl
}
