package myapp

import (
	"fmt"
)

const (
	ERR_zero int = iota
	ERR_one
	ERR_two
	ERR_three
)

type AppError struct {
	error
	Code int
	Msg  string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[ERROR] %s(%d)", e.Msg, e.Code)
}
