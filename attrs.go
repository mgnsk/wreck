package wreck

import (
	"errors"
	"log/slog"
)

// Value extracts a single error attribute.
func Value(err error, key string) (slog.Value, bool) {
	for _, a := range Attrs(err) {
		if a.Key == key {
			return a.Value, true
		}
	}

	return slog.Value{}, false
}

// Args extracts error attributes as key-value pairs.
func Args(err error) (args []any) {
	for _, a := range Attrs(err) {
		args = append(args, a.Key, a.Value.Any())
	}

	return args
}

// Attrs extracts error attributes.
func Attrs(err error) []slog.Attr {
	var args []any

	var werr *Error
	if errors.As(err, &werr) {
		for werr != nil {
			args = append(args, werr.args...)
			werr = werr.base
		}
	}

	if len(args) == 0 {
		return nil
	}

	return slog.Group("", args...).Value.Group()
}
