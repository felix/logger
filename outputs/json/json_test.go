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
			in:  []interface{}{"one"},
			out: map[string]interface{}{"@level": "info", "@name": "test", "message": "one"},
		},
		{
			in:  []interface{}{"one", "two", "2"},
			out: map[string]interface{}{"@level": "info", "@name": "test", "message": "one", "two": "2"},
		},
		{
			in:  []interface{}{"one", "two", "2", "three", 3},
			out: map[string]interface{}{"@level": "info", "@name": "test", "message": "one", "two": "2", "three": float64(3)},
		},
		{
			in:  []interface{}{"one", "two", "2", "three", 3, "fo ur", "# 4"},
			out: map[string]interface{}{"@level": "info", "@name": "test", "message": "one", "two": "2", "three": float64(3), "fo ur": "# 4"},
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
