package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestKeyValueOutput(t *testing.T) {
	tests := []struct {
		msg      string
		names    []string
		fields   []interface{}
		extras   []interface{}
		level    Level
		minLevel Level
		expected string
	}{
		{
			msg:      "msg",
			level:    ERROR,
			minLevel: ERROR,
			expected: "[error] msg",
		},
		{
			msg: "msg", extras: []interface{}{"one", "two"},
			level:    ERROR,
			minLevel: ERROR,
			expected: "[error] msg one=two",
		},
		{
			msg: "msg", extras: []interface{}{"one", "two three"},
			level:    ERROR,
			minLevel: ERROR,
			expected: "[error] msg one=\"two three\"",
		},
		{
			// Odd number of extras
			msg: "msg", extras: []interface{}{"two"},
			level:    ERROR,
			minLevel: ERROR,
			expected: "[error] msg extra00=two",
		},
		{
			msg:      "msg",
			level:    INFO,
			minLevel: ERROR,
			expected: "",
		},
		{
			msg:      "msg",
			level:    INFO,
			minLevel: INFO,
			expected: "[info] msg",
		},
		{
			msg:      "msg",
			level:    INFO,
			minLevel: DEBUG,
			expected: "[info] msg",
		},
		{
			msg:      "msg",
			level:    DEBUG,
			minLevel: ERROR,
			expected: "",
		},
		{
			// Odd number of extras
			msg: "msg", extras: []interface{}{1},
			level:    ERROR,
			minLevel: ERROR,
			expected: "[error] msg extra00=1",
		},
		{
			msg: "msg", extras: []interface{}{"one", 2.23423},
			level:    ERROR,
			minLevel: ERROR,
			expected: "[error] msg one=2.23423",
		},
		{
			msg: "msg", extras: []interface{}{"one", 2.23423},
			fields:   []interface{}{"request", "1234"},
			level:    ERROR,
			minLevel: ERROR,
			expected: "[error] msg request=1234 one=2.23423",
		},
		{
			// Odd number of extras
			msg: "msg", extras: []interface{}{1},
			names:    []string{"one"},
			level:    ERROR,
			minLevel: ERROR,
			expected: "[error] one: msg extra00=1",
		},
		{
			// Odd number of extras
			msg: "msg", extras: []interface{}{1},
			names:    []string{"one", "two"},
			level:    ERROR,
			minLevel: ERROR,
			expected: "[error] one.two: msg extra00=1",
		},
	}
	var buf bytes.Buffer

	for _, tt := range tests {
		logger, err := New(SetOutput(&buf))
		if err != nil {
			t.Fatal("failed to create logger: ", err)
		}
		logger.SetLevel(tt.minLevel)
		for _, n := range tt.names {
			logger = logger.GetNamed(n)
		}
		for i := 0; i < len(tt.fields); i = i + 2 {
			logger.SetField(tt.fields[i].(string), tt.fields[i+1])
		}
		logger.Log(tt.level, tt.msg, tt.extras...)
		actual := strings.TrimSpace(buf.String())
		buf.Reset()
		// Skip the timestamp etc.
		if idx := strings.IndexByte(actual, '['); idx > 0 {
			actual = actual[idx:]
		}
		if actual != tt.expected {
			t.Errorf("Log(%s, %v) => %s, expected %s", tt.level, tt.msg, actual, tt.expected)
		}
	}
}
