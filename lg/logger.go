// Package lg implements wrapper for slog with extra methods
package lg

import (
	"context"
	"log/slog"
	"os"

	"github.com/umed/gib/must"
)

const LevelFatal = slog.LevelError + 1

const (
	Debug = "DEBUG"
	Info  = "INFO"
	Warn  = "WARN"
	Error = "ERROR"
	Fatal = "FATAL"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Fatal(msg string, args ...any)
}

type logger struct {
	*slog.Logger
}

var (
	DefaultLogger        = New()
	NopLogger     Logger = &logger{Logger: slog.New(slog.DiscardHandler)}
)

type Config struct {
	level string
}

type Option func(*Config)

func WithLevel(level string) Option {
	return func(c *Config) {
		c.level = level
	}
}

func New(opts ...Option) Logger {
	c := Config{
		level: Info,
	}
	for _, opt := range opts {
		opt(&c)
	}
	return &logger{
		Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:       MustParseLevel(c.level),
			ReplaceAttr: customReplaceAttr,
		})),
	}
}

func (l *logger) Fatal(msg string, args ...any) {
	l.Log(context.Background(), LevelFatal, msg, args...)
	os.Exit(1)
}

func ParseLevel(level string) (slog.Level, error) {
	var logLevel slog.Level
	err := logLevel.UnmarshalText([]byte(level))
	return logLevel, err
}

func MustParseLevel(level string) slog.Level {
	return must.Must(ParseLevel(level))
}

func customReplaceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey && a.Value.Any().(slog.Level) == LevelFatal {
		a.Value = slog.StringValue(Fatal)
	}
	return a
}
