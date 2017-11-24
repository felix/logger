package keyvalue

import (
	"bytes"
	"github.com/felix/logger"
	"strings"
	"testing"
)

func TestKeyValueWriter(t *testing.T) {
	var tests = []struct {
		in  []interface{}
		out string
	}{
		{
			in:  []interface{}{"test message"},
			out: "[INFO ] test: message=\"test message\"",
		},
		{
			in:  []interface{}{"test message", "name", "me"},
			out: "[INFO ] test: message=\"test message\" name=me",
		},
		{
			in:  []interface{}{"test message", "name", "me", "number", 2},
			out: "[INFO ] test: message=\"test message\" name=me number=2",
		},
	}

	for _, tt := range tests {
		var buf bytes.Buffer
		logger := logger.New(&logger.Options{
			Name:      "test",
			Output:    &buf,
			Formatter: New(),
		})

		logger.Info(tt.in...)

		str := buf.String()

		// Chop timestamp
		dataIdx := strings.IndexByte(str, ' ')
		rest := str[dataIdx+1:]

		if rest != tt.out {
			t.Errorf("Info(%q) => %q, expected %q\n", tt.in, rest, tt.out)
		}
	}
}