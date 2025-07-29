package wreck

import (
	"errors"
	"fmt"
)

// New creates a new error.
func New(msg string) *Error {
	return &Error{
		msg: msg,
	}
}

// Error is an error with attributes.
type Error struct {
	base *Error
	msg  string
	err  error
	args []any
}

// Error returns the internal error message.
func (e *Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %s", e.msg, e.err.Error())
	}
	return e.msg
}

// Message returns the public error message.
func (e *Error) Message() string {
	return e.msg
}

// Unwrap returns the error cause.
func (e *Error) Unwrap() error {
	return e.err
}

// Is reports whether the error matches the base error.
func (e *Error) Is(target error) bool {
	if base, ok := target.(*Error); ok {
		b := e.base
		for b != nil {
			if b == base {
				return true
			}
			b = b.base
		}
	}
	return false
}

// With returns a clone of the error with the specified key-value pair attributes.
func (e *Error) With(args ...any) *Error {
	return &Error{
		base: e,
		msg:  e.msg,
		err:  e.err,
		args: args,
	}
}

// New creates a new error from the base error.
func (e *Error) New(msg string, errs ...error) *Error {
	return &Error{
		base: e,
		msg:  msg,
		err:  errors.Join(errs...),
	}
}
