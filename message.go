package logger

import (
	"io"
)

// Message type for implementors of the MessageWriter interface.
type Message struct {
	// Optional logger name
	Name string
	// The time log() was called
	Time string
	// The log level
	Level Level
	// The message content
	Content string
	// Optional fields for the logger
	Fields []interface{}
	// Optional extras for this log message
	Extras []interface{}
}

// MessageWriter interface for writing messages.
type MessageWriter interface {
	Write(io.Writer, Message)
}
