package errors

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

// BaseError is base error struct.
type BaseError struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`

	err   error
	base  error
	frame xerrors.Frame
}

// Is reports whether any error in err's chain matches target.
func Is(err, target error) bool {
	var e *BaseError
	var t *BaseError
	if errors.As(err, &e) {
		if errors.As(target, &t) {
			return errors.Is(e.base, t.base)
		}
		return errors.Is(e.base, target)
	}
	if errors.As(target, &t) {
		return errors.Is(err, t.base)
	}
	return errors.Is(err, target)
}

// As finds the first error in err's chain that matches target, and if so, sets
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Throw is constructor to create error object.
func Throw(err error) *BaseError {
	var bErr *BaseError

	switch {
	case errors.As(err, &bErr):
		return bErr
	default:
		errors.As(ErrInternalServerError.Wrap(err), &bErr)
		return bErr
	}
}

// New is constructor to create error object.
func New(status, code int, message string) *BaseError {
	const skip = 1

	return &BaseError{
		Status:  status,
		Code:    code,
		Message: message,
		frame:   xerrors.Caller(skip),
	}
}

// Error is error interface implementation.
func (e *BaseError) Error() string {
	return e.Message
}

// Format is error interface implementation.
func (e *BaseError) Format(f fmt.State, c rune) {
	xerrors.FormatError(e, f, c)
}

// FormatError is error interface implementation.
func (e *BaseError) FormatError(p xerrors.Printer) error {
	e.frame.Format(p)

	return e.Unwrap()
}

// Unwrap returns the result of calling the Unwrap method on err.
func (e *BaseError) Unwrap() error {
	return e.err
}

// Wrap is constructor to wrap error object.
func (e *BaseError) Wrap(target error) error {
	const skip = 1
	err := *e
	err.base = e
	err.err = target
	err.frame = xerrors.Caller(skip)

	return &err
}

// Trace is constructor to wrap error object.
func (e *BaseError) Trace() error {
	const skip = 1
	err := *e
	err.base = e
	err.frame = xerrors.Caller(skip)

	return &err
}
