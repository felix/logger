package internal

import (
	"testing"
	"time"
)

type stringer struct{}

func (s stringer) String() string {
	return "I am a stringer"
}

func TestToString(t *testing.T) {
	epoch := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	s := "string"
	b := true
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
		{nil, ""},
		{new(stringer), "I am a stringer"},
		{epoch, "-0001-11-30 00:00:00 +0000 UTC"},
		{struct{ string }{"test"}, "{test}"},
		// Pointers
		{&s, "string"},
		{&epoch, "-0001-11-30 00:00:00 +0000 UTC"},
		{&b, "true"},
	}

	for _, tt := range tests {
		actual := ToString(tt.in)
		if actual != tt.expected {
			t.Errorf("ToString(%v) => %s, expected %s", tt.in, actual, tt.expected)
		}
	}
}
