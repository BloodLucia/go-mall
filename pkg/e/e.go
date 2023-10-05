package e

import "net/http"

type Error struct {
	Code    int
	Reason  string
	Message string
	Err     error
	Stack   string
}

func New(code int, reason string) *Error {
	return &Error{Code: code, Reason: reason}
}

// Error implements error
func (e *Error) Error() string {
	return e.Message
}

// WithMsg with message
func (e *Error) WithMsg(message string) *Error {
	e.Message = message
	return e
}

// WithErr with error
func (e *Error) WithErr(err error) *Error {
	e.Err = err
	return e
}

// WithStack with stack
func (e *Error) WithStack() *Error {
	//e.Stack = LogStack(2, 0)
	return e
}

func BadRequest(reason string) *Error {
	return New(http.StatusBadRequest, reason)
}

func IsBadRequest(err *Error) bool {
	return err.Code == 400
}

func Unauthorized(reason string) *Error {
	return New(http.StatusUnauthorized, reason)
}

func IsUnauthorized(err *Error) bool {
	return err.Code == 401
}

func Forbidden(reason string) *Error {
	return New(http.StatusForbidden, reason)
}

func IsForbidden(err *Error) bool {
	return err.Code == 403
}

func NotFound(reason string) *Error {
	return New(http.StatusNotFound, reason)
}

func IsNotFound(err *Error) bool {
	return err.Code == 404
}

func InternalServer(reason string) *Error {
	return New(http.StatusInternalServerError, reason)
}

func IsInternalServer(err *Error) bool {
	return err.Code == 500
}

func GatewayTimeout(reason string) *Error {
	return New(http.StatusGatewayTimeout, reason)
}

func IsGatewayTimeout(err *Error) bool {
	return err.Code == 504
}
