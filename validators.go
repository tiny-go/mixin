package mixin

// ImmutableValidator prevents value from being changed.
type ImmutableValidator string

func (iv ImmutableValidator) String() string {
	return string(iv)
}

// Validate is an actual validator func. It allows to set the value, but it cannot
// be changed.
func (iv ImmutableValidator) Validate(m Mixin, _ interface{}) error {
	var recv interface{}
	if m.GetProperty(iv.String(), &recv) == ErrNotAvailable {
		return nil
	}
	return ErrImmutable
}

// StringValidator checks if provided value is a string.
type StringValidator string

func (sv StringValidator) String() string {
	return string(sv)
}

// Validate is an actual validator func.
func (StringValidator) Validate(_ Mixin, v interface{}) error {
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
func (BooleanValidator) Validate(_ Mixin, v interface{}) error {
	if _, ok := v.(bool); ok {
		return nil
	}
	return ErrNotABoolean
}
