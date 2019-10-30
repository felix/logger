package kv

import (
	"fmt"
	"io"
	"os"
	"strings"

	"src.userspace.com.au/logger/internal"
	"src.userspace.com.au/logger/message"
)

// Writer implementation.
type Writer struct {
	timeFormat string
	writer     io.Writer
}

// New creates a new Writer writer.
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

// Write implements the message.Writer interface.
func (w Writer) Write(m message.Message) {
	//fmt.Fprintf(w, "%s [%-5s] ", m.Time, m.Level)
	fmt.Fprintf(w.writer, "%s ", m.Time.Format(w.timeFormat))
	if l := m.Level.String(); l != "" {
		fmt.Fprintf(w.writer, "[%s] ", l)
	}
	if m.Name != "" {
		fmt.Fprintf(w.writer, "%s: ", m.Name)
	}

	// Write message content first
	w.writer.Write([]byte(m.Content))

	// Write fields before extras
	for k, v := range m.Fields {
		writeKV(w.writer, k, v)
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
	w.writer.Write([]byte{'\n'})
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
