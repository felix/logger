package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"src.userspace.com.au/logger/message"
	"src.userspace.com.au/logger/writers/kv"
)

func TestLoggerOptions(t *testing.T) {
	l, err := New(Level(message.INFO))
	if err != nil {
		t.Errorf("New() failed: %s", err)
	}
	if !l.IsInfo() {
		t.Errorf("IsInfo() => %t, expected true", l.IsInfo())
	}

	err = Level(message.DEBUG)(l)
	if err != nil {
		t.Errorf("level option failed: %s", err)
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
	kv, err := kv.New(
		kv.SetOutput(&buf),
		kv.SetTimeFormat(""), // Ignore time
	)
	if err != nil {
		t.Fatal("failed to create keyvalue: ", err)
	}
	log, err := New(Writer(kv))
	if err != nil {
		t.Fatal("failed to create logger: ", err)
	}

	for i, tt := range tests {
		log.SetLevel(tt.min)
		log.LogAtLevel(tt.level, "test")
		actual := strings.TrimSpace(buf.String())
		buf.Reset()
		expected := fmt.Sprintf("[%s] test", tt.level)
		if (actual == expected) != tt.output {
			t.Errorf("invalid levelled output for test %d, got %q", i, actual)
		}

		// Test logging without a level
		log.Log("test")
		actual = strings.TrimSpace(buf.String())
		buf.Reset()
		if actual != "test" {
			t.Errorf("invalid Log output for test %d, got %q", i, actual)
		}
	}
}

func TestNamed(t *testing.T) {
	tests := []struct {
		existing string
		name     string
		expected string
	}{
		{"one", "two", "one.two"},
		{"", "two", "two"},
		{"one.one.one", "two", "one.one.one.two"},
	}
	var buf bytes.Buffer
	kv, err := kv.New(kv.SetOutput(&buf))
	if err != nil {
		t.Fatal("failed to create keyvalue: ", err)
	}
	log, err := New(Writer(kv))
	if err != nil {
		t.Fatal("failed to create logger: ", err)
	}

	for _, tt := range tests {
		log.name = tt.existing
		named := log.Named(tt.name)
		named.Error("test")

		actual := buf.String()
		buf.Reset()
		expected := fmt.Sprintf("[error] %s: test", tt.expected)
		if !strings.Contains(actual, expected) {
			t.Errorf("expected %q got %q", expected, actual)
		}
	}
}
