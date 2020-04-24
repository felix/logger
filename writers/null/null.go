package null

import (
	"src.userspace.com.au/logger/message"
)

// Writer implementation.
type Writer struct {
}

// New creates a new Writer writer.
func New() *Writer {
	return &Writer{}
}

// Write implements the message.Writer interface.
func (w Writer) Write(m message.Message) {}
