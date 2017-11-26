package logger

import (
	"io"
)

// Message type for writers
type Message struct {
	Name   string
	Time   string
	Level  Level
	Fields []interface{}
}

// MessageWriter interface for writing messages
type MessageWriter interface {
	Write(io.Writer, Message)
}
