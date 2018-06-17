package mixin

// StringValidator checks if provided value is a string.
type StringValidator string

func (sv StringValidator) String() string {
	return string(sv)
}

// Validate is an actual validator func.
func (StringValidator) Validate(v interface{}) error {
	if _, ok := v.(string); ok {
		return nil
	}
	return ErrNotAString
}

// BooleanValidator checks if provided value is a boolean.
type BooleanValidator string

func (bv BooleanValidator) String() string {
	return string(bv)
}

// Validate is an actual validator func.
func (BooleanValidator) Validate(v interface{}) error {
	if _, ok := v.(bool); ok {
		return nil
	}
	return ErrNotABoolean
}
