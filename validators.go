package mixin

import "errors"

// StringValidator checks if ptovided value is a string.
type StringValidator string

func (sv StringValidator) String() string {
	return string(sv)
}

// Validate is an actual validator func.
func (StringValidator) Validate(v interface{}) error {
	if _, ok := v.(string); ok {
		return nil
	}
	return errors.New("value is not a string")
}
