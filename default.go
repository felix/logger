package logger

import (
	"fmt"
	"io"
	"strings"
)

// DefaultWriter implementation
type DefaultWriter struct{}

// NewDefaultWriter creates a new writer
func NewDefaultWriter() *DefaultWriter {
	return &DefaultWriter{}
}

// Write implements the logger.MessageWriter interface
func (dw DefaultWriter) Write(w io.Writer, m Message) {
	prefix := fmt.Sprintf("%s [%-5s]", m.Time, strings.ToUpper(m.Level.String()))
	io.WriteString(w, prefix)

	if m.Name != "" {
		io.WriteString(w, fmt.Sprintf(" %s:", m.Name))
	}

	for _, f := range m.Fields {
		io.WriteString(w, " ")
		io.WriteString(w, ToString(f))
	}
}
