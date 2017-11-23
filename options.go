package logger

import "io"

type Options struct {
	Name       string
	Level      Level
	Fields     []interface{}
	Output     io.Writer
	TimeFormat string
	Formatter  MessageWriter
}
