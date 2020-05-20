package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"src.userspace.com.au/logger/writers/kv"
)

func TestLoggerOptions(t *testing.T) {
	l, err := New()
	err = ForceDebug(true)(l)
	if err != nil {
		t.Errorf("level option failed: %s", err)
	}
	if !l.IsDebug() {
		t.Errorf("IsDebug() => %t, expected true", l.IsDebug())
	}
}

func TestLog(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected string
	}{
		{[]interface{}{"one", "two"}, "one two"},
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
		log.Log(tt.input...)

		actual := buf.String()
		buf.Reset()
		if !strings.Contains(actual, tt.expected) {
			t.Errorf("expected %q got %q", tt.expected, actual)
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
		named.Log("test")

		actual := buf.String()
		buf.Reset()
		expected := fmt.Sprintf("%s: test", tt.expected)
		if !strings.Contains(actual, expected) {
			t.Errorf("expected %q got %q", expected, actual)
		}
	}
}
