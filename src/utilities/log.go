package utilities

import "log"

const (
	info    = "INFO"
	err     = "ERROR"
	warning = "WARNING"
)

type Logger struct {
	log.Logger
}

func NewLogger(l log.Logger) Logger {
	return Logger{}
}

func (l Logger) Info(message string) {
	l.print(info, message)
}

func (l Logger) Error(message string) {
	l.print(err, message)
}

func (l Logger) Warning(message string) {
	l.print(warning, message)
}

func (l Logger) print(etype string, message string) {
	l.Printf("[%s] %s", etype, message)
}
