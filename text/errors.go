package text

import (
	"errors"
)

var (
	// ErrNotConnected is the not connected error.
	ErrNotConnected = errors.New("not connected")

	// ErrNoSuchFileOrDirectory is the no such file or directory error.
	ErrNoSuchFileOrDirectory = errors.New("no such file or directory")

	// ErrCannotIncludeDirectories is the cannot include directories error.
	ErrCannotIncludeDirectories = errors.New("cannot include directories")

	// ErrMissingDSN is the missing dsn error.
	ErrMissingDSN = errors.New("missing dsn")

	// ErrNoPreviousTransactionExists is the no previous transaction exists error.
	ErrNoPreviousTransactionExists = errors.New("no previous transaction exists")

	// ErrPreviousTransactionExists is the previous transaction exists error.
	ErrPreviousTransactionExists = errors.New("previous transaction exists")

	// ErrPasswordAttemptsExhausted is the exhausted password attempts error.
	ErrPasswordAttemptsExhausted = errors.New("password attempts exhuasted")

	// ErrSingleTransactionCannotBeUsedWithInteractiveMode is the single transaction cannot be used with interactive mode error.
	ErrSingleTransactionCannotBeUsedWithInteractiveMode = errors.New("--single-transaction cannot be used with interactive mode")

	// ErrNoEditorDefined is the no editor defined error.
	ErrNoEditorDefined = errors.New("no editor defined")

	// ErrUnknownCommand is the unknown command error.
	ErrUnknownCommand = errors.New("unknown command")

	// ErrMissingRequiredArgument is the missing required argument error.
	ErrMissingRequiredArgument = errors.New("missing required argument")

	// ErrDriverNotAvailable is the driver not available error.
	ErrDriverNotAvailable = errors.New("driver not available")

	// ErrPasswordNotSupportedByDriver is the password not supported by driver error.
	ErrPasswordNotSupportedByDriver = errors.New(`\password not supported by driver`)
)
