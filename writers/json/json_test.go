package json

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/felix/logger/message"
)

func TestWriter(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2019-05-03T13:38:29.987249+10:00")
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
			expected: map[string]interface{}{
				"@level": "error", "@name": "test", "@message": "msg",
				"@time": "2019-05-03T13:38:29.987249+10:00",
			},
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.ERROR,
				Content: "msg",
				Extras:  []interface{}{"one"},
			},
			expected: map[string]interface{}{
				"@level": "error", "@name": "test", "@message": "msg", "extra00": "one",
				"@time": "2019-05-03T13:38:29.987249+10:00",
			},
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.ERROR,
				Content: "msg",
				Fields:  map[string]interface{}{"one": "1"},
			},
			expected: map[string]interface{}{
				"@level": "error", "@name": "test", "@message": "msg", "one": "1",
				"@time": "2019-05-03T13:38:29.987249+10:00",
			},
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.ERROR,
				Content: "msg", Extras: []interface{}{"one", "1", "two", "2", "three", 3, "fo ur", "# 4"},
			},
			expected: map[string]interface{}{
				"@level": "error", "@name": "test", "@message": "msg", "one": "1", "two": "2", "three": float64(3), "fo ur": "# 4",
				"@time": "2019-05-03T13:38:29.987249+10:00",
			},
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.DEBUG,
				Content: "msg",
				Extras:  []interface{}{"one"}, Fields: map[string]interface{}{"f1": "v1"},
			},
			expected: map[string]interface{}{
				"@level": "debug", "@name": "test", "@message": "msg", "f1": "v1", "extra00": "one",
				"@time": "2019-05-03T13:38:29.987249+10:00",
			},
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

		if !reflect.DeepEqual(raw, tt.expected) {
			t.Errorf("got %q, expected %q\n", raw, tt.expected)
		}
	}
}
