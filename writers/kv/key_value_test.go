package kv

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/felix/logger/message"
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
				Level:   message.ERROR,
				Content: "msg",
			},
			expected: "2019-05-03T13:38:29.987+1000 [error] test: msg",
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.ERROR,
				Content: "msg",
				Extras:  []interface{}{"one"},
			},
			expected: "2019-05-03T13:38:29.987+1000 [error] test: msg extra00=one",
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.ERROR,
				Content: "msg",
				Fields:  map[string]interface{}{"one": "1"},
			},
			expected: "2019-05-03T13:38:29.987+1000 [error] test: msg one=1",
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.ERROR,
				Content: "msg", Extras: []interface{}{"one", "1", "two", "2", "three", 3, "fo ur", "# 4"},
			},
			expected: `2019-05-03T13:38:29.987+1000 [error] test: msg one=1 two=2 three=3 "fo ur"="# 4"`,
		},
		{
			in: message.Message{
				Time:    now,
				Name:    "test",
				Level:   message.DEBUG,
				Content: "msg",
				Fields:  map[string]interface{}{"f1": "v1"},
				Extras:  []interface{}{"one"},
			},
			expected: "2019-05-03T13:38:29.987+1000 [debug] test: msg f1=v1 extra00=one",
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
