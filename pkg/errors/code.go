package errors

import "errors"
var (
	ErrInvalidConfig = errors.New("invalid config")
	ErrInvalidDBType = errors.New("invalid db type")
	ErrInvalidConnection = errors.New("invalid connection to db")
)