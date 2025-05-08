package log

import (
	"io"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

const logsFilePath = "logs.json"

const (
	ErrorLabel = "error"
)

var logger *Logger
var (
	LogFile *os.File
	err     error
)

func init() {
	LogFile, err = os.OpenFile(logsFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	multi := io.MultiWriter(os.Stdout, LogFile)
	handler := slog.NewJSONHandler(multi, nil)
	baseLogger := slog.New(handler)

	logger = &Logger{baseLogger}
}

func GetLogger() *Logger {
	return logger
}

// Logs error then exits
func (l *Logger) Fatal(msg string, args ...any) {
	l.Logger.Error(msg, args...)
	os.Exit(1)
}
