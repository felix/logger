package logger

// Level defines the output level
type Level int

// Log levels
const (
	ERROR Level = iota // Wake someone up
	WARN               // Seomthing failed but don't wake anyone up
	INFO               // Good to know
	DEBUG              // Not for production
)

func (l Level) String() string {
	switch l {
	case 0:
		return "error"
	case 1:
		return "warn"
	case 2:
		return "info"
	case 3:
		return "debug"
	default:
		return "unknown"
	}
}

// Levels is a convenience for string -> level
var Levels = map[string]Level{
	"ERROR": ERROR,
	"WARN":  WARN,
	"INFO":  INFO,
	"DEBUG": DEBUG,
}
