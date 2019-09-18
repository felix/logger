package message

// Level defines the output level
type Level int

// Log levels
const (
	NONE  Level = iota // Always log it
	ERROR              // Wake someone up
	WARN               // Something failed but don't wake anyone up
	INFO               // Good to know
	DEBUG              // Not for production
)

func (l Level) String() string {
	switch l {
	case 0:
		return ""
	case 1:
		return "error"
	case 2:
		return "warn"
	case 3:
		return "info"
	case 4:
		return "debug"
	default:
		return "unknown"
	}
}

// Levels is a convenience for string -> level
var Levels = map[string]Level{
	"NONE":  NONE,
	"ERROR": ERROR,
	"WARN":  WARN,
	"INFO":  INFO,
	"DEBUG": DEBUG,
}
