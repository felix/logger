package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestDefaultWriter(t *testing.T) {
	var tests = []struct {
		in  []interface{}
		out string
	}{
		{
			in:  []interface{}{"one"},
			out: "[INFO ] testlog: one",
		},
		{
			in:  []interface{}{"one", "two", "2"},
			out: "[INFO ] testlog: one two 2",
		},
		{
			in:  []interface{}{"one", "two", "2", "three", 3},
			out: "[INFO ] testlog: one two 2 three 3",
		},
		{
			in:  []interface{}{"one", map[string]string{"two": "2", "three": "3"}},
			out: "[INFO ] testlog: one two 2 three 3",
		},
	}

	for _, tt := range tests {
		var buf bytes.Buffer
		logger := New(&Options{
			Name:   "testlog",
			Output: &buf,
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

// Comparison with hclog
func BenchmarkLogger(b *testing.B) {
	b.Run("info with 10 pairs", func(b *testing.B) {
		var buf bytes.Buffer

		logger := New(&Options{
			Name:   "test",
			Output: &buf,
		})

		for i := 0; i < b.N; i++ {
			logger.Info("this is some message",
				"name", "foo",
				"what", "benchmarking yourself",
				"why", "to see what's slow",
				"k4", "value",
				"k5", "value",
				"k6", "value",
				"k7", "value",
				"k8", "value",
				"k9", "value",
				"k10", "value",
			)
		}
	})
}
