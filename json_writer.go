package logger

import (
	"bufio"
	"encoding/json"
)

type JSONWriter struct{}

func NewJSONWriter() *JSONWriter {
	return &JSONWriter{}
}

func (jw JSONWriter) Write(w *bufio.Writer, m Message) {
	vals := map[string]interface{}{
		"@name":  m.Name,
		"@level": m.Level.String(),
		"@time":  m.Time,
	}

	for i := 0; i < len(m.Fields); i = i + 2 {
		vals[toString(m.Fields[i])] = m.Fields[i+1]
	}

	err := json.NewEncoder(w).Encode(vals)
	if err != nil {
		panic(err)
	}
}
