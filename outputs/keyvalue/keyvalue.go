package keyvalue

import (
	"fmt"
	"github.com/felix/logger"
	"io"
	"strings"
)

// Writer implementation
type Writer struct{}

// New creates a new writer
func New() *Writer {
	return &Writer{}
}

// Write implements the logger.MessageWriter interface
func (kv Writer) Write(w io.Writer, m logger.Message) {
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
		maybeQuote(logger.ToString(k)),
		maybeQuote(logger.ToString(v)),
	)
}

func maybeQuote(s string) string {
	if strings.ContainsAny(s, " \t\n\r") {
		return fmt.Sprintf("%q", s)
	}
	return s
}
