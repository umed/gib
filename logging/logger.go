package logging

import (
	"context"
	"log/slog"
	"os"

	"github.com/umed/gib/util"
)

const LevelFatal = slog.LevelError + 1

type Logger struct {
	*slog.Logger
}

var DefaultLogger = NewLogger("debug")

func NewLogger(level string) *Logger {
	return &Logger{
		Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:       MustParseLevel(level),
			ReplaceAttr: customReplaceAttr,
		})),
	}
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.Log(context.Background(), LevelFatal, msg, args...)
	os.Exit(1)
}

func ParseLevel(level string) (slog.Level, error) {
	var logLevel slog.Level
	err := logLevel.UnmarshalText([]byte(level))
	return logLevel, err
}

func MustParseLevel(level string) slog.Level {
	return util.Must(ParseLevel(level))
}

func customReplaceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey && a.Value.Any().(slog.Level) == LevelFatal {
		a.Value = slog.StringValue("FATAL")
	}
	return a
}
