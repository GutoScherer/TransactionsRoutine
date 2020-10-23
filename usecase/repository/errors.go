package repository

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

var (
	// ErrRegisterNotFound represents a register not found error on repository
	ErrRegisterNotFound = errors.New(`register not found`)

	// ErrInvalidData represents a invalid data error on repository
	ErrInvalidData = errors.New(`invalid data`)

	// ErrDuplicatedEntry represents a duplicated entry error on repository
	ErrDuplicatedEntry = errors.New(`duplicated entry`)

	// ErrForeignKeyConstraint represents a duplicated entry error on repository
	ErrForeignKeyConstraint = errors.New(`foreign key constraint error`)
)

func buildRepositoryError(err *mysql.MySQLError) error {
	const (
		errDuplicatedEntry      = 1062
		errInvalidData          = 1406
		errForeignKeyConstraint = 1452
	)

	switch err.Number {
	case errInvalidData:
		return ErrInvalidData
	case errDuplicatedEntry:
		return ErrDuplicatedEntry
	case errForeignKeyConstraint:
		return ErrForeignKeyConstraint
	}

	return nil
}
