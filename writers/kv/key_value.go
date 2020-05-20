package kv

import (
	"fmt"
	"io"
	"os"
	"strings"

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
	fmt.Fprintf(w.writer, "%s ", m.Time.Format(w.timeFormat))
	if m.Name != "" {
		fmt.Fprintf(w.writer, "%s: ", m.Name)
	}

	// Write message content first
	w.writer.Write([]byte(m.Content))

	for k, v := range m.Fields {
		writeKV(w.writer, k, v)
	}
	w.writer.Write([]byte{'\n'})
}

func writeKV(w io.Writer, k, v string) (int, error) {
	return fmt.Fprintf(w, " %s=%s", maybeQuote(k), maybeQuote(v))
}

func maybeQuote(s string) string {
	if strings.ContainsAny(s, " \t\n\r") {
		return fmt.Sprintf("%q", s)
	}
	return s
}
