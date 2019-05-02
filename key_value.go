package logger

import (
	"fmt"
	"io"
	"strings"
)

// KeyValue implementation.
type KeyValue struct{}

// Write implements the logger.MessageWriter interface.
func (kv KeyValue) Write(w io.Writer, m Message) {
	//fmt.Fprintf(w, "%s [%-5s] ", m.Time, m.Level)
	fmt.Fprintf(w, "%s [%s] ", m.Time, m.Level)
	if m.Name != "" {
		fmt.Fprintf(w, "%s: ", m.Name)
	}

	// Write message content first
	w.Write([]byte(m.Content))

	// Write fields before extras
	for i := 0; i < len(m.Fields); i = i + 2 {
		writeKV(w, m.Fields[i], m.Fields[i+1])
	}

	if len(m.Extras) > 0 {
		// Allow for an odd number of extras
		offset := len(m.Extras) % 2
		if offset != 0 {
			for k, v := range m.Extras {
				writeKV(w, fmt.Sprintf("extra%02d", k), v)
			}
		} else {
			for i := offset; i < len(m.Extras); i = i + 2 {
				writeKV(w, m.Extras[i], m.Extras[i+1])
			}
		}
	}
}

func writeKV(w io.Writer, k, v interface{}) (int, error) {
	return fmt.Fprintf(w, " %s=%s",
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
