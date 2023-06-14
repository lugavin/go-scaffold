package errors

import "errors"

var (
	ErrInvalidParams = errors.New("params is invalid")
)

type ErrCode uint32

const (
	ParamsValidationErrCode ErrCode = 1 << iota
)

type AppError interface {
	error
	Cause() error
	Code() ErrCode
}

type appError struct {
	cause error
	code  ErrCode
	msg   string
}

func New(errCode ErrCode, errMsg string) *appError {
	return &appError{
		code: errCode,
		msg:  errMsg,
	}
}

func NewWithCause(errCode ErrCode, errMsg string, cause error) *appError {
	return &appError{
		cause: cause,
		code:  errCode,
		msg:   errMsg,
	}
}

func (e *appError) Error() string {
	if e.cause != nil {
		return e.cause.Error()
	}
	return e.msg
}

func (e *appError) Cause() error {
	return e.cause
}

func (e *appError) Code() ErrCode {
	return e.code
}

func (e *appError) Is(err error) bool {
	// Check, if our inner error is a direct match
	if errors.Is(errors.Unwrap(e), err) {
		return true
	}

	// Otherwise, we need to match using our error flags
	switch err {
	case ErrInvalidParams:
		return e.code&ParamsValidationErrCode != 0
	}

	return false
}
