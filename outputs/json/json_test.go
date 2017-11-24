package json

import (
	"bytes"
	"encoding/json"
	"github.com/felix/logger"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestWriter(t *testing.T) {
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
		logger := logger.New(&logger.Options{
			Name:      "test",
			Output:    &buf,
			Formatter: New(),
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
