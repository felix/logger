package json

import (
	"encoding/json"
	"fmt"
	"io"

	"src.userspace.com.au/felix/logger"
)

// Writer implementation
type Writer struct{}

// New creates a new writer
func New() *Writer {
	return new(Writer)
}

// Write implements the logger.MessageWriter interface
func (w Writer) Write(lw io.Writer, m logger.Message) {
	vals := map[string]interface{}{
		"@name":    m.Name,
		"@level":   m.Level.String(),
		"@time":    m.Time,
		"@message": m.Content,
	}

	for i := 0; i < len(m.Fields); i = i + 2 {
		vals[logger.ToString(m.Fields[i])] = m.Fields[i+1]
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
				vals[logger.ToString(m.Extras[i])] = m.Extras[i+1]
			}
		}
	}

	if err := json.NewEncoder(lw).Encode(vals); err != nil {
		fmt.Fprintf(lw, "\"failed to encode JSON: %s\"", err)
	}
}
