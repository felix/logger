package json

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"src.userspace.com.au/felix/logger"
)

func TestWriter(t *testing.T) {
	var tests = []struct {
		in  []interface{}
		out map[string]interface{}
	}{
		{
			in:  []interface{}{"one"},
			out: map[string]interface{}{"@level": "error", "@name": "test", "@message": "msg", "extra00": "one"},
		},
		{
			in:  []interface{}{"one", "1"},
			out: map[string]interface{}{"@level": "error", "@name": "test", "@message": "msg", "one": "1"},
		},
		{
			in:  []interface{}{"one", "1", "two", "2", "three", 3, "fo ur", "# 4"},
			out: map[string]interface{}{"@level": "error", "@name": "test", "@message": "msg", "one": "1", "two": "2", "three": float64(3), "fo ur": "# 4"},
		},
	}

	var buf bytes.Buffer

	for _, tt := range tests {
		l, err := logger.New(
			logger.SetName("test"),
			logger.SetOutput(&buf),
			logger.SetFormatter(New()),
		)
		if err != nil {
			t.Fatalf("New failed: %q", err)
		}

		l.Error("msg", tt.in...)

		var raw map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &raw); err != nil {
			t.Fatal(err)
		}
		buf.Reset()

		// Ignore time
		delete(raw, "@time")

		if !reflect.DeepEqual(raw, tt.out) {
			t.Errorf("Error(%q) => %q, expected %q\n", tt.in, raw, tt.out)
		}
	}
}
