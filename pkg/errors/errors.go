package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

type baseError struct {
	err   error
	frame xerrors.Frame
}

func (e *baseError) Error() string {
	return e.err.Error()
}

func (e *baseError) Format(f fmt.State, c rune) {
	xerrors.FormatError(e, f, c)
}

func (e *baseError) FormatError(p xerrors.Printer) error {
	e.frame.Format(p)

	return e.Unwrap()
}

// Unwrap returns the result of calling the Unwrap method on err.
func (e *baseError) Unwrap() error {
	return e.err
}

// Wrap is constructor to wrap error object.
func Wrap(err error) error {
	const skip = 1

	return &baseError{
		err:   err,
		frame: xerrors.Caller(skip),
	}
}
