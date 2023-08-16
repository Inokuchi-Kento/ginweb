/*
Package errors はエラー処理用です。
*/
package errors

import (
	"errors"
	"fmt"
	"ginweb/ent"
)

type serverError struct {
	errType ErrorType
	err     error
}

// Unwrap implements the errors.Wrapper interface.
func (e *serverError) Unwrap() error {
	return e.err
}

// Error implements the error interface.
func (e *serverError) Error() string {
	msg := fmt.Sprintf("%v %s ", e.errType.HTTPCode(), e.errType.Code())
	if e.err != nil {
		msg += e.err.Error()
	}

	return msg
}

func New(errType ErrorType, err error) *serverError {
	return &serverError{
		errType: errType,
		err:     err,
	}
}

func NewErrorf(errType ErrorType, format string, a ...interface{}) *serverError {
	return New(errType, fmt.Errorf(format, a...))
}

func ErrType(err error) ErrorType {
	var serverError *serverError
	if errors.As(err, &serverError) {
		return serverError.errType
	}
	if ent.IsNotFound(err) ||
		ent.IsNotLoaded(err) ||
		ent.IsNotSingular(err) ||
		ent.IsConstraintError(err) ||
		ent.IsValidationError(err) {
		return DatabaseError
	}
	return InternalError
}

func IsErrType(err *serverError, errType ErrorType) bool {
	return err.errType == errType
}
