package property

import "errors"

var (
	// ErrCannotAssign is invoked if receiver type does not match the stored value.
	ErrCannotAssign = errors.New("cannot assign the value to the provided receiver")
	// ErrNotAPointer is returned if provided receiver is not a pointer type.
	ErrNotAPointer = errors.New("receiver is not a pointer")
	// ErrNotAvailable means that requested property does not exist in the storage.
	ErrNotAvailable = errors.New("property is not available")

	// ErrImmutable is invoked if the value is not allowed to be changed.
	ErrImmutable = errors.New("value cannot be changed")
	// ErrNotAString is returned if provided value is not a string.
	ErrNotAString = errors.New("value is not a string")
	// ErrNotABoolean is returned if provided value is not a boolean.
	ErrNotABoolean = errors.New("value is not a boolean")
)
