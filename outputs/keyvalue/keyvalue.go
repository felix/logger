package keyvalue

import (
	"bufio"
	"fmt"
	"github.com/felix/logger"
	"strings"
)

// KeyValueWriter implementation
type Writer struct{}

// New creates a new writer
func Writer() *Writer {
	return &Writer{}
}

// Write implements the logger.MessageWriter interface
func (kv Writer) Write(w *bufio.Writer, m logger.Message) {
	w.WriteString(m.Time)
	w.WriteByte(' ')
	w.WriteString(fmt.Sprintf("[%-5s]", strings.ToUpper(m.Level.String())))
	if m.Name != "" {
		w.WriteByte(' ')
		w.WriteString(m.Name)
		w.WriteByte(':')
	}

	for i := 0; i < len(m.Fields); i = i + 2 {
		w.WriteByte(' ')
		w.WriteString(toString(m.Fields[i]))
		w.WriteByte('=')
		s := toString(m.Fields[i+1])
		if strings.ContainsAny(s, " \t\n\r") {
			w.WriteByte('"')
			w.WriteString(s)
			w.WriteByte('"')
		} else {
			w.WriteString(s)
		}
	}
}
