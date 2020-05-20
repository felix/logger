package json

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"src.userspace.com.au/logger/message"
)

func TestWriter(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2019-05-03T13:38:29.987249+10:00")
	var tests = []struct {
		in       message.Message
		expected map[string]string
	}{
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Content: "msg",
			},
			expected: map[string]string{
				"_name": "test", "_message": "msg",
				"_time": "2019-05-03T13:38:29.987+1000",
			},
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Content: "msg",
			},
			expected: map[string]string{
				"_name": "test", "_message": "msg",
				"_time": "2019-05-03T13:38:29.987+1000",
			},
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Content: "msg",
				Fields:  map[string]string{"one": "1"},
			},
			expected: map[string]string{
				"_name": "test", "_message": "msg", "one": "1",
				"_time": "2019-05-03T13:38:29.987+1000",
			},
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Content: "msg",
			},
			expected: map[string]string{
				"_name": "test", "_message": "msg",
				"_time": "2019-05-03T13:38:29.987+1000",
			},
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Content: "msg",
				Fields:  map[string]string{"f1": "v1"},
			},
			expected: map[string]string{
				"_name": "test", "_message": "msg", "f1": "v1",
				"_time": "2019-05-03T13:38:29.987+1000",
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

		var raw map[string]string
		if err := json.Unmarshal(buf.Bytes(), &raw); err != nil {
			t.Fatal(err)
		}
		buf.Reset()

		if !reflect.DeepEqual(raw, tt.expected) {
			t.Errorf("got %q, expected %q\n", raw, tt.expected)
		}
	}
}
