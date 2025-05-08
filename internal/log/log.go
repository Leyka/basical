package log

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

type Logger struct {
	*slog.Logger
}

const (
	ErrorLabel = "error"
)

var logger *Logger

func init() {
	w := os.Stderr
	baseLogger := slog.New(tint.NewHandler(w, nil))

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
