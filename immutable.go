package property

// ImmutableValidator prevents value from being changed.
type ImmutableValidator string

func (iv ImmutableValidator) String() string {
	return string(iv)
}

// Validate is an actual validator func. It allows to set the value, but it cannot
// be changed.
func (iv ImmutableValidator) Validate(m Manager, _ interface{}) error {
	var recv interface{}
	if m.GetProperty(iv.String(), &recv) == ErrNotAvailable {
		return nil
	}
	return ErrImmutable
}
