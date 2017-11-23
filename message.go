package logger

import (
	"bufio"
)

type Message struct {
	Name   string
	Time   string
	Level  Level
	Fields []interface{}
}

type MessageWriter interface {
	Write(*bufio.Writer, Message)
}
