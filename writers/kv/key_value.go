package kv

import (
	"fmt"
	"io"
	"strings"

	"src.userspace.com.au/felix/logger/internal"
	"src.userspace.com.au/felix/logger/message"
)

// Writer implementation.
type Writer struct {
	writer io.Writer
}

// New creates a new Writer writer.
func New(w io.Writer) (*Writer, error) {
	return &Writer{writer: w}, nil
}

// Write implements the message.Writer interface.
func (w Writer) Write(m message.Message) {
	//fmt.Fprintf(w, "%s [%-5s] ", m.Time, m.Level)
	fmt.Fprintf(w.writer, "%s [%s] ", m.Time, m.Level)
	if m.Name != "" {
		fmt.Fprintf(w.writer, "%s: ", m.Name)
	}

	// Write message content first
	w.writer.Write([]byte(m.Content))

	// Write fields before extras
	for i := 0; i < len(m.Fields); i = i + 2 {
		writeKV(w.writer, m.Fields[i], m.Fields[i+1])
	}

	if len(m.Extras) > 0 {
		// Allow for an odd number of extras
		offset := len(m.Extras) % 2
		if offset != 0 {
			for k, v := range m.Extras {
				writeKV(w.writer, fmt.Sprintf("extra%02d", k), v)
			}
		} else {
			for i := offset; i < len(m.Extras); i = i + 2 {
				writeKV(w.writer, m.Extras[i], m.Extras[i+1])
			}
		}
	}
}

func writeKV(w io.Writer, k, v interface{}) (int, error) {
	return fmt.Fprintf(w, " %s=%s",
		maybeQuote(internal.ToString(k)),
		maybeQuote(internal.ToString(v)),
	)
}

func maybeQuote(s string) string {
	if strings.ContainsAny(s, " \t\n\r") {
		return fmt.Sprintf("%q", s)
	}
	return s
}
