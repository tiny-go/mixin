package property

// StringValidator checks if provided value is a string.
type StringValidator string

func (sv StringValidator) String() string {
	return string(sv)
}

// Validate is an actual validator func.
func (StringValidator) Validate(_ Manager, v interface{}) error {
	if _, ok := v.(string); ok {
		return nil
	}
	return ErrNotAString
}
