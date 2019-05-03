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
		in       []interface{}
		fields   map[string]interface{}
		expected map[string]interface{}
	}{
		{
			in:       []interface{}{"one"},
			expected: map[string]interface{}{"@level": "error", "@name": "test", "@message": "msg", "extra00": "one"},
		},
		{
			in:       []interface{}{"one", "1"},
			expected: map[string]interface{}{"@level": "error", "@name": "test", "@message": "msg", "one": "1"},
		},
		{
			in:       []interface{}{"one", "1", "two", "2", "three", 3, "fo ur", "# 4"},
			expected: map[string]interface{}{"@level": "error", "@name": "test", "@message": "msg", "one": "1", "two": "2", "three": float64(3), "fo ur": "# 4"},
		},
		{
			in:       []interface{}{"one"},
			fields:   map[string]interface{}{"f1": "v1"},
			expected: map[string]interface{}{"@level": "error", "@name": "test", "@message": "msg", "f1": "v1", "extra00": "one"},
		},
	}

	var buf bytes.Buffer
	writer, err := New(&buf)
	if err != nil {
		panic(err)
	}

	for _, tt := range tests {
		l, err := logger.New(logger.SetName("test"), logger.AddWriter(writer))
		if err != nil {
			t.Fatalf("New failed: %q", err)
		}

		for k, v := range tt.fields {
			l.SetField(k, v)
		}

		l.Error("msg", tt.in...)

		var raw map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &raw); err != nil {
			t.Fatal(err)
		}
		buf.Reset()

		// Ignore time
		delete(raw, "@time")

		if !reflect.DeepEqual(raw, tt.expected) {
			t.Errorf("Error(%q) => %q, expected %q\n", tt.in, raw, tt.expected)
		}
	}
}
