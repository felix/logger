package json

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"src.userspace.com.au/felix/logger/message"
)

func TestWriter(t *testing.T) {
	now := time.Now()
	var tests = []struct {
		in       message.Message
		expected map[string]interface{}
	}{
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.ERROR,
				Content: "msg",
			},
			expected: map[string]interface{}{"@level": "error", "@name": "test", "@message": "msg"},
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.ERROR,
				Content: "msg",
				Extras:  []interface{}{"one"},
			},
			expected: map[string]interface{}{"@level": "error", "@name": "test", "@message": "msg", "extra00": "one"},
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.ERROR,
				Content: "msg",
				Fields:  map[string]interface{}{"one": "1"},
			},
			expected: map[string]interface{}{"@level": "error", "@name": "test", "@message": "msg", "one": "1"},
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.ERROR,
				Content: "msg", Extras: []interface{}{"one", "1", "two", "2", "three", 3, "fo ur", "# 4"},
			},
			expected: map[string]interface{}{"@level": "error", "@name": "test", "@message": "msg", "one": "1", "two": "2", "three": float64(3), "fo ur": "# 4"},
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.DEBUG,
				Content: "msg",
				Extras:  []interface{}{"one"}, Fields: map[string]interface{}{"f1": "v1"},
			},
			expected: map[string]interface{}{"@level": "debug", "@name": "test", "@message": "msg", "f1": "v1", "extra00": "one"},
		},
	}

	var buf bytes.Buffer
	l, err := New(SetOutput(&buf))
	if err != nil {
		panic(err)
	}

	for _, tt := range tests {
		l.Write(tt.in)

		var raw map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &raw); err != nil {
			t.Fatal(err)
		}
		buf.Reset()

		// Ignore time
		delete(raw, "@time")

		if !reflect.DeepEqual(raw, tt.expected) {
			t.Errorf("got %q, expected %q\n", raw, tt.expected)
		}
	}
}
