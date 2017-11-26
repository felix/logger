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

	for i := 0; i < len(m.Fields); i = i + 2 {
		io.WriteString(w, " ")
		io.WriteString(w, maybeQuote(logger.ToString(m.Fields[i])))
		io.WriteString(w, "=")
		s := logger.ToString(m.Fields[i+1])
		io.WriteString(w, maybeQuote(s))
	}
}

func maybeQuote(s string) string {
	if strings.ContainsAny(s, " \t\n\r") {
		return fmt.Sprintf("%q", s)
	}
	return s
}
