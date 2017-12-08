package logger

import (
	"fmt"
	"io"
	"strings"
)

// DefaultWriter implementation
type DefaultWriter struct{}

// New creates a new writer
func NewDefaultWriter() *DefaultWriter {
	return &DefaultWriter{}
}

// Write implements the logger.MessageWriter interface
func (kv DefaultWriter) Write(w io.Writer, m Message) {
	prefix := fmt.Sprintf("%s [%-5s]", m.Time, strings.ToUpper(m.Level.String()))
	io.WriteString(w, prefix)
	if m.Name != "" {
		io.WriteString(w, " ")
		io.WriteString(w, m.Name)
		io.WriteString(w, ":")
	}

	offset := len(m.Fields) % 2
	if offset != 0 {
		io.WriteString(w, writeKV("message", m.Fields[0]))
	}

	for i := offset; i < len(m.Fields); i = i + 2 {
		io.WriteString(w, writeKV(m.Fields[i], m.Fields[i+1]))
	}
}

func writeKV(k, v interface{}) string {
	return fmt.Sprintf(
		" %s=%s",
		maybeQuote(ToString(k)),
		maybeQuote(ToString(v)),
	)
}

func maybeQuote(s string) string {
	if strings.ContainsAny(s, " \t\n\r") {
		return fmt.Sprintf("%q", s)
	}
	return s
}
