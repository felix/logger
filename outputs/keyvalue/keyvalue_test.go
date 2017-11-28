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
			in:  []interface{}{"one"},
			out: "[INFO ] test: message=one",
		},
		{
			in:  []interface{}{"one", "two", "2"},
			out: "[INFO ] test: message=one two=2",
		},
		{
			in:  []interface{}{"one", "two", "2", "three", 3},
			out: "[INFO ] test: message=one two=2 three=3",
		},
		{
			in:  []interface{}{"one", "two", "2", "three", 3, "fo ur", "# 4"},
			out: "[INFO ] test: message=one two=2 three=3 \"fo ur\"=\"# 4\"",
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
