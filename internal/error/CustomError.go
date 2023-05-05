package error

import (
	"errors"
	"reflect"
)

type customError struct {
	StatusCode int
	Err        error
}

type Type customError

/*
	Implement error.
*/
func (customError Type) Error() string {
	return customError.Err.Error()
}

func newCustomError(statusCode int, err error) Type {
	return Type{
		StatusCode: statusCode,
		Err:        err,
	}
}

/*
	Check if an error is CustomError or not
*/
func IsInstanceOfCustomError(err interface{}) bool {
	if err == nil {
		return false
	}
	return reflect.TypeOf(err) == reflect.TypeOf(Type{})
}

/*
	Type Casting: Cast error type to customError type if it is possible.
	If not, panic an error.
*/
func Cast(err interface{}) Type {
	if !IsInstanceOfCustomError(err) {
		panic("Can not cast error to CustomError")
	}
	return err.(Type)
}

/*
	Panic a CustomError.
*/
func Throw(statusCode int, message string) {
	panic(newCustomError(statusCode, errors.New(message)))
}
