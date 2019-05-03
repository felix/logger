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

func TestLevels(t *testing.T) {
	tests := []struct {
		min    message.Level
		level  message.Level
		output bool
	}{
		// error
		{message.ERROR, message.ERROR, true},
		{message.ERROR, message.WARN, false},
		{message.ERROR, message.INFO, false},
		{message.ERROR, message.DEBUG, false},
		// warn
		{message.WARN, message.ERROR, true},
		{message.WARN, message.WARN, true},
		{message.WARN, message.INFO, false},
		{message.WARN, message.DEBUG, false},
		// info
		{message.INFO, message.ERROR, true},
		{message.INFO, message.WARN, true},
		{message.INFO, message.INFO, true},
		{message.INFO, message.DEBUG, false},
		// debug
		{message.DEBUG, message.ERROR, true},
		{message.DEBUG, message.WARN, true},
		{message.DEBUG, message.INFO, true},
		{message.DEBUG, message.DEBUG, true},
	}
	var buf bytes.Buffer
	kv, err := kv.New(kv.SetOutput(&buf))
	if err != nil {
		t.Fatal("failed to create keyvalue: ", err)
	}
	log, err := New(AddWriter(kv))
	if err != nil {
		t.Fatal("failed to create logger: ", err)
	}

	for i, tt := range tests {
		log.SetLevel(tt.min)
		log.Log(tt.level, "test")
		actual := strings.TrimSpace(buf.String())
		buf.Reset()
		if (len(actual) > 0) != tt.output {
			t.Errorf("invalid output for test %d", i)
		}
	}
}
