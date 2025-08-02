package logger

import (
	"log"
	"os"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type StdLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewStdLogger() *StdLogger {
	return &StdLogger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Ltime|log.Lshortfile),
	}
}

func (l *StdLogger) Info(msg string) {
	l.infoLogger.Println(msg)
}
func (l *StdLogger) Error(msg string) {
	l.errorLogger.Println(msg)
}
