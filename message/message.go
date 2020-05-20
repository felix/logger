package message

import "time"

// Message type for implementors of the Writer interface.
type Message struct {
	// Optional logger name
	Name string
	// The time log() was called
	Time time.Time
	// The message content
	Content string
	// Optional fields for the logger
	Fields map[string]string
}

// Writer interface for writing messages.
type Writer interface {
	Write(Message)
}
