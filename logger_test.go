package logger

import (
	"bytes"
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
			out: "[INFO ] test: message=one\n",
		},
		{
			in:  []interface{}{"one", "two", "2"},
			out: "[INFO ] test: message=one two=2\n",
		},
		{
			in:  []interface{}{"one", "two", "2", "three", 3},
			out: "[INFO ] test: message=one two=2 three=3\n",
		},
		{
			in:  []interface{}{"one", "two", "2", "three", 3, "fo ur", "# 4"},
			out: "[INFO ] test: message=one two=2 three=3 \"fo ur\"=\"# 4\"\n",
		},
	}

	for _, tt := range tests {
		var buf bytes.Buffer
		logger := New(&Options{
			Name:   "test",
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

func TestKeyValueWriterWithFields(t *testing.T) {
	var tests = []struct {
		in  []interface{}
		out string
	}{
		{
			in:  []interface{}{"one"},
			out: "[INFO ] test: message=one added=this\n",
		},
		{
			in:  []interface{}{"one", "two", "2"},
			out: "[INFO ] test: message=one two=2 added=this\n",
		},
		{
			in:  []interface{}{"one", "two", "2", "three", 3},
			out: "[INFO ] test: message=one two=2 three=3 added=this\n",
		},
		{
			in:  []interface{}{"one", "two", "2", "three", 3, "fo ur", "# 4"},
			out: "[INFO ] test: message=one two=2 three=3 \"fo ur\"=\"# 4\" added=this\n",
		},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		logger := New(&Options{
			Name:   "test",
			Output: &buf,
		}).WithFields("added", "this")

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

func TestLevels(t *testing.T) {
	logger := New(&Options{
		Name:  "test",
		Level: Debug,
	})

	if !logger.IsDebug() {
		t.Errorf("Level Debug check failed")
	}

	logger = New(&Options{
		Name:  "test",
		Level: Error,
	})

	if !logger.IsError() {
		t.Errorf("Level Error check failed")
	}
}
