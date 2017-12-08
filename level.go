package logger

import "strings"

// Level defines the logger output level
type Level int

const (
	// NoLevel is prior to being defined
	NoLevel Level = 0
	// Debug is for development
	Debug Level = 1
	// Info are for interesting runtime events
	Info Level = 2
	// Warn is for almost errors
	Warn Level = 3
	// Error is a runtime problem
	Error Level = 4
)

func (lvl Level) String() string {
	switch lvl {
	case 1:
		return "debug"
	case 2:
		return "info"
	case 3:
		return "warn"
	case 4:
		return "error"
	default:
		return "unknown"
	}
}

// LevelFromString helps select a level
func LevelFromString(l string) Level {
	switch strings.ToLower(l) {
	case "debug":
		return Debug
	case "warn":
		return Warn
	case "Info":
		return Info
	case "Error":
		return Error
	default:
		return NoLevel
	}
}
