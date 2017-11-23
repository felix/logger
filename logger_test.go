package logger

import (
	"bytes"
	//"fmt"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
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
		logger := New(&Options{
			Name:      "test",
			Output:    &buf,
			Formatter: NewKeyValueWriter(),
		})

		logger.Info(tt.in...)

		str := buf.String()

		//fmt.Printf("output => %s\n", str)

		// Chop timestamp
		dataIdx := strings.IndexByte(str, ' ')
		rest := str[dataIdx+1:]

		if rest != tt.out {
			t.Errorf("Info(%q) => %q, expected %q\n", tt.in, rest, tt.out)
		}
	}
}

func TestJSONWriter(t *testing.T) {
	var tests = []struct {
		in  []interface{}
		out map[string]interface{}
	}{
		{
			in:  []interface{}{"test message"},
			out: map[string]interface{}{"@level": "info", "@name": "test", "message": "test message"},
		},
		{
			in:  []interface{}{"test message", "name", "me"},
			out: map[string]interface{}{"@level": "info", "@name": "test", "message": "test message", "name": "me"},
		},
		{
			in:  []interface{}{"test message", "name", "me", "number", 2},
			out: map[string]interface{}{"@level": "info", "@name": "test", "message": "test message", "name": "me", "number": float64(2)},
		},
	}

	for _, tt := range tests {
		var buf bytes.Buffer
		logger := New(&Options{
			Name:      "test",
			Output:    &buf,
			Formatter: NewJSONWriter(),
		})

		logger.Info(tt.in...)

		b := buf.Bytes()

		var raw map[string]interface{}
		if err := json.Unmarshal(b, &raw); err != nil {
			t.Fatal(err)
		}

		delete(raw, "@time")

		if !cmp.Equal(raw, tt.out) {
			t.Errorf("Info(%q) => %q, expected %q\n", tt.in, raw, tt.out)
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
