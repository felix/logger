package message

import (
	"testing"
)

func TestLevel(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"ERROR", "error"},
		{"DEBUG", "debug"},
		{"WARN", "warn"},
		{"INFO", "info"},
	}

	for _, tt := range tests {
		l := Levels[tt.in]
		actual := l.String()
		if actual != tt.expected {
			t.Errorf("got %s, expected %s", actual, tt.expected)
		}
	}
}
