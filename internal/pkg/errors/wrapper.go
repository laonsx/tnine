package errors

import (
	stderrors "errors"
	"fmt"

	pkgerrors "github.com/pkg/errors"
)

func New(message string) error {
	return stderrors.New(message)
}

func Errorf(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}

func WithStack(err error) error {
	return pkgerrors.WithStack(err)
}

func Wrap(err error, message string) error {
	return pkgerrors.Wrapf(err, message)
}

func Wrapf(err error, format string, args ...any) error {
	return pkgerrors.Wrapf(err, format, args...)
}

func WithMessage(err error, message string) error {
	return pkgerrors.WithMessage(err, message)
}

func WithMessagef(err error, format string, args ...any) error {
	return pkgerrors.WithMessagef(err, format, args...)
}

func Cause(err error) error {
	return pkgerrors.Cause(err)
}

func Is(err, target error) bool {
	return stderrors.Is(err, target)
}

func As(err error, target any) bool {
	return stderrors.As(err, target)
}

func Unwrap(err error) error {
	return stderrors.Unwrap(err)
}

type Frame = pkgerrors.Frame

type StackTrace = pkgerrors.StackTrace
