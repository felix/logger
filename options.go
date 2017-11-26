package logger

import "io"

// DefaultTimeFormat unless specified by options
const DefaultTimeFormat = "2006-01-02T15:04:05.000Z0700"

// Options to configure the logger
type Options struct {
	Name       string
	Level      Level
	Fields     []interface{}
	Output     io.Writer
	TimeFormat string
	Formatter  MessageWriter
}
