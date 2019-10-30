package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"src.userspace.com.au/logger/internal"
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

	if l := m.Level.String(); l != "" {
		vals["_level"] = m.Level.String()
	}

	for k, v := range m.Fields {
		vals[k] = v
	}

	if len(m.Extras) > 0 {
		// Allow for an odd number of extras
		offset := len(m.Extras) % 2
		if offset != 0 {
			for k, v := range m.Extras {
				vals[fmt.Sprintf("extra%02d", k)] = v
			}
		} else {
			for i := offset; i < len(m.Extras); i = i + 2 {
				vals[internal.ToString(m.Extras[i])] = m.Extras[i+1]
			}
		}
	}

	if err := json.NewEncoder(w.writer).Encode(vals); err != nil {
		fmt.Fprintf(w.writer, "\"failed to encode JSON: %s\"", err)
	}
}
