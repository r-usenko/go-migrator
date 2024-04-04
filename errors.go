package migration

import "errors"

var (
	//ErrBeginTx                   = errors.New("invalid BeginTx instance")
	//ErrInvalidTypeOfDriver       = errors.New("driver must be struct or pointer of struct")
	//ErrInvalidSignatureOfBeginTx = errors.New("invalid signature of BeginTx method of driver")

	ErrInvalidTypeOfSource = errors.New("source must be struct or pointer of struct")
	ErrEmptySource         = errors.New("source hasn't have method for migrate")
	ErrInvalidSourceMethod = errors.New("invalid source method signature. must return error")
)
