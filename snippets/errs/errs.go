// Package errs provides trpc error code type, which contains errcode errmsg.
// These definitions are multi-language universal.
package errs

import (
	"fmt"
)

var (
	traceable bool
	content   string
)

// trpc return code.
const (
	// RetOK means success.
	RetOK = 0

	// RetUnknown is the error code for unspecified errors.
	RetUnknown = 999
)

// ErrorType is the error code type, including framework error code and business error code.
const (
	ErrorTypeServer   = 1
	ErrorTypeBusiness = 2
	ErrorTypeClient   = 3 // The error code returned by the client call
	// represents the downstream framework error code.
)

const (
	// Success is the success prompt string.
	Success = "success"
)

// Error is the error code structure which contains error code type and error message.
type Error struct {
	Type int
	Code int32
	Msg  string
	Desc string
}

// Error implements the error interface and returns the error description.
func (e *Error) Error() string {
	if e == nil {
		return Success
	}
	switch e.Type {
	case ErrorTypeServer:
		return fmt.Sprintf("type:server, code:%d, msg:%s", e.Code, e.Msg)
	case ErrorTypeClient:
		return fmt.Sprintf("type:client, code:%d, msg:%s", e.Code, e.Msg)
	default:
		return fmt.Sprintf("type:business, code:%d, msg:%s", e.Code, e.Msg)
	}
}

// New creates an error, which defaults to the business error type to improve business development efficiency.
func New(code int, msg string) error {
	err := &Error{
		Type: ErrorTypeBusiness,
		Code: int32(code),
		Msg:  msg,
	}
	return err
}

// Newf creates an error, the default is the business error type, msg supports format strings.
func Newf(code int, format string, params ...interface{}) error {
	msg := fmt.Sprintf(format, params...)
	err := &Error{
		Type: ErrorTypeBusiness,
		Code: int32(code),
		Msg:  msg,
	}
	return err
}

// NewServerError creates a frame error.
func NewServerError(code int, msg string) error {
	err := &Error{
		Type: ErrorTypeServer,
		Code: int32(code),
		Msg:  msg,
		Desc: "server",
	}
	return err
}

// Code gets the error code through error.
func Code(e error) int {
	if e == nil {
		return 0
	}
	err, ok := e.(*Error)
	if !ok {
		return RetUnknown
	}
	if err == (*Error)(nil) {
		return 0
	}
	return int(err.Code)
}

// Msg gets error msg through error.
func Msg(e error) string {
	if e == nil {
		return Success
	}
	err, ok := e.(*Error)
	if !ok {
		return e.Error()
	}
	if err == (*Error)(nil) {
		return Success
	}
	return err.Msg
}
