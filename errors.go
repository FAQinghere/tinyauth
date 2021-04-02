package tinyauth

import "fmt"

type ErrType int

const (
	ErrBadInput ErrType = iota
	ErrInternal
	ErrAuthFailed
)

type Error struct {
	Msg string
	Err error
	ErrType
}

func (l *Error) Error() string {
	return fmt.Sprintf("%s: %s", l.Msg, l.Err.Error())
}
func (l *Error) Unwrap() error { return l.Err }

func NewErrorBadInput(msg string, err error) *Error {
	return &Error{
		Msg:     msg,
		Err:     err,
		ErrType: ErrBadInput,
	}
}

func NewErrorInternal(msg string, err error) *Error {
	return &Error{
		Msg:     msg,
		Err:     err,
		ErrType: ErrInternal,
	}
}

func NewErrorAuthFailed(msg string, err error) *Error {
	return &Error{
		Msg:     msg,
		Err:     err,
		ErrType: ErrAuthFailed,
	}
}
