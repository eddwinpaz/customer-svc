package entity

import (
	"errors"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")

	ErrOpenFileLog = errors.New("Could not open log file")

	// ErrRecordNotFound returned
	ErrRecordNotFound = errors.New("record not found")

	ErrInvalidParameter = errors.New("invalid parameter")

	ErrNotSaved = errors.New("not saved")

	ErrNotUpdated = errors.New("not updated")

	ErrInvalidJSON = errors.New("invalid json")

	ErrInvalidCredentials = errors.New("invalid credentials")

	ErrMissingAuthorizationToken = errors.New("missing authorization token")

	ErrInvalidAuthorizationToken = errors.New("invalid authorization token")

	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("requested is not found")

	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("already exist")

	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")

	// ErrorNotFoundOnDB will throw if the given answer is not present on database
	ErrorNotFoundOnDB = errors.New("element not found on DB")

	// ErrorInvalidForm will throw if the given customer input is invalid for some
	// kind when present on delivery/web validation
	ErrorInvalidForm = errors.New("sent Form is not valid")

	// ErrorAlreadyExists will throw if the new customer email is in our database
	ErrorAlreadyExists = errors.New("email or Phone already exists")

	// ErrEmailExists email already exists error
	ErrEmailExists = errors.New("email already exists")

	// ErrPhoneExists phone already exists error
	ErrPhoneExists = errors.New("phone already exists")

	// ErrDatabaseError this happens when SQL fails
	ErrDatabaseError = errors.New("database Error, Try again")

	// ErrSQLError Internal Database Error
	ErrSQLError = errors.New("internal Database Error, Try Later")

	ErrDatabaseConnection = errors.New("database connection error")

	ErrCustomerNotDeleted = errors.New("customer not deleted")

	ErrStartingServer = errors.New("starting server error")

	//
)
