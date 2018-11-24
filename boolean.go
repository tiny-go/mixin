package property

// BooleanValidator checks if provided value is a boolean.
type BooleanValidator string

func (bv BooleanValidator) String() string {
	return string(bv)
}

// Validate is an actual validator func.
func (BooleanValidator) Validate(_ Manager, v interface{}) error {
	if _, ok := v.(bool); ok {
		return nil
	}
	return ErrNotABoolean
}
