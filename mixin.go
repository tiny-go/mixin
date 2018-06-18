package mixin

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
)

// PropertyValidator knows the name of the property that should be validated and
// contains an actual validator func.
type PropertyValidator interface {
	fmt.Stringer
	Validate(Mixin, interface{}) error
}

// Mixin is responsible for the management of the custom object properties
// (implies an ability to save/retrieve the properties from the object by name).
type Mixin interface {
	// GetProperty should read the value of custom (user-defined) parameter and put to
	// the provided receiver (should be a pointer). It returns an error in case if
	// parameter does not exist or receiver has an invalid type.
	GetProperty(name string, recv interface{}) error
	// SetProperty should add a custom parameter to the object (or replace if exists).
	SetProperty(name string, value interface{}) error
}

// mixin is a PropertyManager implementation.
type mixin struct {
	mu         sync.Mutex
	storage    map[string]interface{}
	validators map[string][]func(Mixin, interface{}) error
}

// New creates a new Mixin.
func New(validators ...PropertyValidator) Mixin {
	mixin := &mixin{
		storage:    make(map[string]interface{}),
		validators: make(map[string][]func(Mixin, interface{}) error),
	}
	for _, validator := range validators {
		mixin.validators[validator.String()] = append(mixin.validators[validator.String()], validator.Validate)
	}
	return mixin
}

// GetProperty retrieves custom object property (if exists) and assignes it to the
// provided receiver (if possible).
func (m *mixin) GetProperty(name string, recv interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = ErrCannotAssign
		}
	}()
	// check argument type (should be a pointer)
	rv := reflect.ValueOf(recv)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return ErrNotAPointer
	}
	m.mu.Lock()
	val, ok := m.storage[name]
	m.mu.Unlock()
	if !ok {
		return ErrNotAvailable
	}
	rv.Elem().Set(reflect.ValueOf(val))
	return nil
}

// SetProperty stores custom object property.
func (m *mixin) SetProperty(name string, value interface{}) error {
	if validators, ok := m.validators[name]; ok {
		for _, validator := range validators {
			if err := validator(m, value); err != nil {
				return err
			}
		}
	}
	m.mu.Lock()
	m.storage[name] = value
	m.mu.Unlock()
	return nil
}

func (m *mixin) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.storage)
}
