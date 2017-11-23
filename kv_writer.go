package logger

import (
	"bufio"
	"fmt"
	"strings"
)

type KeyValueWriter struct{}

func (kv KeyValueWriter) Write(w *bufio.Writer, m Message) {
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
