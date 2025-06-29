package lg

import (
	"log/slog"
	"time"
)

func Err(err error) slog.Attr {
	return slog.Any("error", err)
}

func String(key, value string) slog.Attr {
	return slog.String(key, value)
}

// Int64 returns an Attr for an int64.
func Int64(key string, value int64) slog.Attr {
	return slog.Int64(key, value)
}

// Int converts an int to an int64 and returns
// an Attr with that value.
func Int(key string, value int) slog.Attr {
	return Int64(key, int64(value))
}

// Uint64 returns an Attr for a uint64.
func Uint64(key string, value uint64) slog.Attr {
	return slog.Uint64(key, value)
}

// Float64 returns an Attr for a floating-point number.
func Float64(key string, value float64) slog.Attr {
	return slog.Float64(key, value)
}

// Bool returns an Attr for a bool.
func Bool(key string, value bool) slog.Attr {
	return slog.Bool(key, value)
}

// Time returns an Attr for a [time.Time].
// It discards the monotonic portion.
func Time(key string, value time.Time) slog.Attr {
	return slog.Time(key, value)
}

// Duration returns an Attr for a [time.Duration].
func Duration(key string, value time.Duration) slog.Attr {
	return slog.Duration(key, value)
}

// Group returns an Attr for a Group [Value].
// The first argument is the key; the remaining arguments
// are converted to Attrs as in [Logger.Log].
//
// Use Group to collect several key-value pairs under a single
// key on a log line, or as the result of LogValue
// in order to log a single value as multiple Attrs.
func Group(key string, values ...any) slog.Attr {
	return slog.Group(key, values...)
}

// Any returns an Attr for the supplied value.
// See [AnyValue] for how values are treated.
func Any(key string, value any) slog.Attr {
	return slog.Any(key, value)
}
