package logger

import (
	"bufio"
	"fmt"
	"strings"
)

// DefaultWriter implementation
type DefaultWriter struct{}

// NewDefaultWriter creates a new writer
func NewDefaultWriter() *DefaultWriter {
	return &DefaultWriter{}
}

// Write implements the logger.MessageWriter interface
func (dw DefaultWriter) Write(w *bufio.Writer, m Message) {
	w.WriteString(m.Time)
	w.WriteByte(' ')
	w.WriteString(fmt.Sprintf("[%-5s]", strings.ToUpper(m.Level.String())))
	if m.Name != "" {
		w.WriteByte(' ')
		w.WriteString(m.Name)
	}

	for _, f := range m.Fields {
		w.WriteByte(' ')
		w.WriteString(ToString(f))
	}
}
