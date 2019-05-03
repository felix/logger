package json

import (
	"encoding/json"
	"fmt"
	"io"

	"src.userspace.com.au/felix/logger/internal"
	"src.userspace.com.au/felix/logger/message"
)

// Writer implementation
type Writer struct {
	writer io.Writer
}

// New creates a new writer
func New(w io.Writer) (*Writer, error) {
	return &Writer{writer: w}, nil
}

// Write implements the logger.Writer interface
func (w Writer) Write(m message.Message) {
	vals := map[string]interface{}{
		"@name":    m.Name,
		"@level":   m.Level.String(),
		"@time":    m.Time,
		"@message": m.Content,
	}

	for i := 0; i < len(m.Fields); i = i + 2 {
		vals[internal.ToString(m.Fields[i])] = m.Fields[i+1]
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
