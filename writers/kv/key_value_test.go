package kv

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"src.userspace.com.au/logger/message"
)

func TestWriter(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2019-05-03T13:38:29.987249+10:00")
	var tests = []struct {
		in       message.Message
		expected string
	}{
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Content: "msg",
			},
			expected: "2019-05-03T13:38:29.987+1000 test: msg",
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Content: "msg",
			},
			expected: "2019-05-03T13:38:29.987+1000 test: msg",
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Content: "msg",
				Fields:  map[string]string{"one": "1"},
			},
			expected: "2019-05-03T13:38:29.987+1000 test: msg one=1",
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Content: "msg",
			},
			expected: `2019-05-03T13:38:29.987+1000 test: msg`,
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Content: "msg",
				Fields:  map[string]string{"f1": "v1"},
			},
			expected: "2019-05-03T13:38:29.987+1000 test: msg f1=v1",
		},
	}

	var buf bytes.Buffer
	l, err := New(SetOutput(&buf))
	if err != nil {
		panic(err)
	}

	for _, tt := range tests {
		l.Write(tt.in)
		actual := strings.TrimSpace(buf.String())
		buf.Reset()

		if actual != tt.expected {
			t.Errorf("got %s, expected %s", actual, tt.expected)
		}
	}
}
