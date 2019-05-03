package message

// Message type for implementors of the Writer interface.
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

// Writer interface for writing messages.
type Writer interface {
	Write(Message)
}
