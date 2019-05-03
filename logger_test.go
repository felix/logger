package logger

import (
	"bytes"
	"strings"
	"testing"

	"src.userspace.com.au/felix/logger/message"
	"src.userspace.com.au/felix/logger/writers/kv"
)

func TestLoggerOptions(t *testing.T) {
	l, err := New(SetLevel(message.INFO))
	if err != nil {
		t.Errorf("New() failed: %s", err)
	}
	if !l.IsInfo() {
		t.Errorf("IsInfo() => %t, expected true", l.IsInfo())
	}

	err = l.Configure(SetLevel(message.DEBUG))
	if err != nil {
		t.Errorf("Configure() failed: %s", err)
	}
	if !l.IsDebug() {
		t.Errorf("IsDebug() => %t, expected true", l.IsDebug())
	}
}

func TestDefaultOutput(t *testing.T) {
	tests := []struct {
		msg      string
		names    []string
		fields   []interface{}
		extras   []interface{}
		level    message.Level
		minLevel message.Level
		expected string
	}{
		{
			msg:      "msg",
			level:    message.ERROR,
			minLevel: message.ERROR,
			expected: "[error] msg",
		},
		{
			msg: "msg", extras: []interface{}{"one", "two"},
			level:    message.ERROR,
			minLevel: message.ERROR,
			expected: "[error] msg one=two",
		},
		{
			msg: "msg", extras: []interface{}{"one", "two three"},
			level:    message.ERROR,
			minLevel: message.ERROR,
			expected: "[error] msg one=\"two three\"",
		},
		{
			// Odd number of extras
			msg: "msg", extras: []interface{}{"two"},
			level:    message.ERROR,
			minLevel: message.ERROR,
			expected: "[error] msg extra00=two",
		},
		{
			msg:      "msg",
			level:    message.INFO,
			minLevel: message.ERROR,
			expected: "",
		},
		{
			msg:      "msg",
			level:    message.INFO,
			minLevel: message.INFO,
			expected: "[info] msg",
		},
		{
			msg:      "msg",
			level:    message.INFO,
			minLevel: message.DEBUG,
			expected: "[info] msg",
		},
		{
			msg:      "msg",
			level:    message.DEBUG,
			minLevel: message.ERROR,
			expected: "",
		},
		{
			// Odd number of extras
			msg: "msg", extras: []interface{}{1},
			level:    message.ERROR,
			minLevel: message.ERROR,
			expected: "[error] msg extra00=1",
		},
		{
			msg: "msg", extras: []interface{}{"one", 2.23423},
			level:    message.ERROR,
			minLevel: message.ERROR,
			expected: "[error] msg one=2.23423",
		},
		{
			msg: "msg", extras: []interface{}{"one", 2.23423},
			fields:   []interface{}{"request", "1234"},
			level:    message.ERROR,
			minLevel: message.ERROR,
			expected: "[error] msg request=1234 one=2.23423",
		},
		{
			// Odd number of extras
			msg: "msg", extras: []interface{}{1},
			names:    []string{"one"},
			level:    message.ERROR,
			minLevel: message.ERROR,
			expected: "[error] one: msg extra00=1",
		},
		{
			// Odd number of extras
			msg: "msg", extras: []interface{}{1},
			names:    []string{"one", "two"},
			level:    message.ERROR,
			minLevel: message.ERROR,
			expected: "[error] one.two: msg extra00=1",
		},
	}
	var buf bytes.Buffer
	kv, err := kv.New(&buf)
	if err != nil {
		t.Fatal("failed to create keyvalue: ", err)
	}

	for _, tt := range tests {
		log, err := New(AddWriter(kv))
		if err != nil {
			t.Fatal("failed to create logger: ", err)
		}
		log.SetLevel(tt.minLevel)
		for _, n := range tt.names {
			log = log.GetNamed(n)
		}
		for i := 0; i < len(tt.fields); i = i + 2 {
			log.SetField(tt.fields[i].(string), tt.fields[i+1])
		}
		log.Log(tt.level, tt.msg, tt.extras...)
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
