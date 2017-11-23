package logger

import (
	"encoding/json"
	"io"
)

type JSONWriter struct{}

func (jw JSONWriter) Write(w io.Writer, m Message) {
	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		panic(err)
	}
}
