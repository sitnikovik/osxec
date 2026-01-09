package code

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// ErrNotParsable is returned when an exit code cannot be parsed.
	ErrNotParsable = errors.New("exit code is not parsable")
	// ErrUnsupportedType is returned when an unsupported type is provided for exit code parsing.
	ErrUnsupportedType = errors.New("unsupported type for exit code")
)

// Code represents the exit code of a process.
type Code uint8

const (
	// Success indicates that the process completed successfully.
	//
	// This exit code is commonly used to indicate that a program has
	// executed without any errors.
	Success Code = 0
	// Failure indicates that the process failed.
	//
	// This exit code is commonly used to indicate a general error or failure
	// that does not fall under more specific categories.
	Failure Code = 1
)

// ParseCode parses the provided code value and returns it.
//
// The code parameter can be of type uint8, int, or string.
// It converts the provided value to a Code type accordingly.
//
// If the conversion fails, it returns an error and Failure code.
func ParseCode[T uint8 | int | string](code T) (Code, error) {
	switch v := any(code).(type) {
	case uint8:
		return Code(v), nil
	case int:
		if v < 0 || v > 255 {
			return Failure, fmt.Errorf(
				"%w from '%d': value out of range [0-255]",
				ErrNotParsable,
				v,
			)
		}
		return Code(v), nil
	case string:
		n, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return Failure, fmt.Errorf(
				"%w from '%s': %v",
				ErrNotParsable,
				v,
				err,
			)
		}
		return Code(uint8(n)), nil
	default:
		return Failure, ErrUnsupportedType
	}
}

// Int returns the integer representation of the Code.
func (c Code) Int() int {
	return int(c)
}

// Succeeded indicates whether the Code represents a success.
func (c Code) Succeeded() bool {
	return c.Equals(Success)
}

// Equals compares numeric exit code values by their numeric values.
//
// This ensures that parsed/constructed Codes with the same numeric value
// are considered equal even if they were produced differently.
func (c Code) Equals(cmp Code) bool {
	return c.Int() == cmp.Int()
}
