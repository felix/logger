package logger

type Level int

const (
	NoLevel Level = 0
	Debug   Level = 1
	Info    Level = 2
	Warn    Level = 3
	Error   Level = 4
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
