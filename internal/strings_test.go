package internal

import (
	"testing"
)

func TestToString(t *testing.T) {
	tests := []struct {
		in       interface{}
		expected string
	}{
		{"string", "string"},
		{1, "1"},
		{0, "0"},
		{false, "false"},
		{int(-3), "-3"},
		{int8(3), "3"},
		{int16(30), "30"},
		{int32(30), "30"},
		{int64(30), "30"},
		{uint(3), "3"},
		{uint8(3), "3"},
		{uint16(30), "30"},
		{uint32(30), "30"},
		{uint64(30), "30"},
		{float32(3.0001), "3.0001"},
		{float64(3.0000001), "3.0000001"},
		{nil, "<nil>"},
	}

	for _, tt := range tests {
		actual := ToString(tt.in)
		if actual != tt.expected {
			t.Errorf("ToString(%v) => %s, expected %s", tt.in, actual, tt.expected)
		}
	}
}
