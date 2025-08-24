// Package lg implements wrapper for slog with extra methods
package lg

import (
	"context"
	"io"
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
	Named(name string) Logger
	With(attrs ...any) Logger
}

type logger struct {
	*slog.Logger
	name string
}

var (
	DefaultLogger        = New()
	NopLogger     Logger = &logger{Logger: slog.New(slog.DiscardHandler)}
)

type Config struct {
	level            string
	addSource        bool
	jsonOutputFormat bool
	outputStream     io.Writer
}

type Option func(*Config)

func WithLevel(level string) Option {
	return func(c *Config) {
		c.level = level
	}
}

func WithSource() Option {
	return func(c *Config) {
		c.addSource = true
	}
}

func WithJSONOutputFormat() Option {
	return func(c *Config) {
		c.jsonOutputFormat = true
	}
}

func WithOutputWriter(w io.Writer) Option {
	return func(c *Config) {
		c.outputStream = w
	}
}

func New(opts ...Option) Logger {
	c := Config{
		level:        Info,
		outputStream: os.Stdout,
	}
	for _, opt := range opts {
		opt(&c)
	}
	handlerOptions := &slog.HandlerOptions{
		Level:       MustParseLevel(c.level),
		ReplaceAttr: customReplaceAttr,
		AddSource:   c.addSource,
	}
	if c.jsonOutputFormat {
		return &logger{
			Logger: slog.New(slog.NewJSONHandler(c.outputStream, handlerOptions)),
		}
	}
	return &logger{
		Logger: slog.New(slog.NewTextHandler(c.outputStream, handlerOptions)),
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

func (l *logger) Named(name string) Logger {
	return &logger{Logger: l.Logger.With(String("logger_name", l.name+"."+name)), name: name}
}

func (l *logger) With(attrs ...any) Logger {
	logger := *l
	logger.Logger = logger.Logger.With(attrs...)
	return &logger
}
