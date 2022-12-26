package utilities

import "log"

const (
	info    = "INFO"
	err     = "ERROR"
	warning = "WARNING"
	fatal   = "FATAL"
)

type Logger struct {
	logger log.Logger
}

func NewLogger() Logger {
	return Logger{
		logger: *log.Default(),
	}
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

func (l Logger) PanicApp(message string) {
	l.printBreak(fatal, message)
}

func (l Logger) print(etype string, message string) {
	l.logger.Printf("[%s] %s", etype, message)
}

func (l Logger) printBreak(etype string, message string) {
	l.logger.Fatalf("[%s] %s", etype, message)
}
