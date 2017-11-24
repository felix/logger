package json

import (
	"bufio"
	"encoding/json"
	"github.com/felix/logger"
)

// Writer implementation
type Writer struct{}

// New creates a new writer
func New() *Writer {
	return &Writer{}
}

// Write implements the logger.MessageWriter interface
func (w Writer) Write(bw *bufio.Writer, m logger.Message) {
	vals := map[string]interface{}{
		"@name":  m.Name,
		"@level": m.Level.String(),
		"@time":  m.Time,
	}

	for i := 0; i < len(m.Fields); i = i + 2 {
		vals[m.Fields[i].(string)] = m.Fields[i+1]
	}

	err := json.NewEncoder(bw).Encode(vals)
	if err != nil {
		panic(err)
	}
}
