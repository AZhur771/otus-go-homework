package logger

import (
	"log"
	"os"
)

type Level string

const (
	Debug Level = "debug"
	Info        = "info"
	Warn        = "warn"
	Error       = "error"
)

var (
	DebugLogger   = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)
	InfoLogger    = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	ErrorLogger   = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
)

type Logger struct {
	Level Level
}

func levelInSlice(l Level, list []Level) bool {
	for _, el := range list {
		if el == l {
			return true
		}
	}
	return false
}

func New(level Level) *Logger {
	if !levelInSlice(level, []Level{Debug, Info, Warn, Error}) {
		log.Fatalf("unknown log level: %s", level)
	}

	return &Logger{
		Level: level,
	}
}

func (l Logger) Debug(msg string) {
	if l.Level == Debug {
		DebugLogger.Println(msg)
	}
}

func (l Logger) Info(msg string) {
	if levelInSlice(l.Level, []Level{Debug, Info}) {
		InfoLogger.Println(msg)
	}
}

func (l Logger) Warn(msg string) {
	if levelInSlice(l.Level, []Level{Debug, Info, Warn}) {
		WarningLogger.Println(msg)
	}
}

func (l Logger) Error(msg string) {
	ErrorLogger.Println(msg)
}
