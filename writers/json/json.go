package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"src.userspace.com.au/logger/message"
)

// Writer implementation
type Writer struct {
	timeFormat string
	writer     io.Writer
}

// New creates a new writer
func New(opts ...Option) (*Writer, error) {
	w := &Writer{
		timeFormat: "2006-01-02T15:04:05.000Z0700",
		writer:     os.Stdout,
	}

	for _, opt := range opts {
		if err := opt(w); err != nil {
			return nil, err
		}
	}
	return w, nil
}

// Write implements the logger.Writer interface
func (w Writer) Write(m message.Message) {
	vals := map[string]interface{}{
		"_name":    m.Name,
		"_time":    m.Time.Format(w.timeFormat),
		"_message": m.Content,
	}

	for k, v := range m.Fields {
		vals[k] = v
	}

	if err := json.NewEncoder(w.writer).Encode(vals); err != nil {
		fmt.Fprintf(w.writer, "\"failed to encode JSON: %s\"", err)
	}
}
