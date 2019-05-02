package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	l, err := New(SetLevel(INFO))
	if err != nil {
		t.Errorf("New() failed: %s", err)
	}
	if !l.IsInfo() {
		t.Errorf("IsInfo() => %t, expected true", l.IsInfo())
	}

	err = l.Configure(SetLevel(DEBUG))
	if err != nil {
		t.Errorf("Configure() failed: %s", err)
	}
	if !l.IsDebug() {
		t.Errorf("IsDebug() => %t, expected true", l.IsDebug())
	}
}
