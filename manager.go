package property

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
)

// Validator knows the name of the property that should be validated and contains
// an actual validator func.
type Validator interface {
	fmt.Stringer
	Validate(Manager, interface{}) error
}

// Manager is responsible for the management of the custom object properties
// (implies an ability to save/retrieve the properties from the object by name).
type Manager interface {
	// GetProperty should read the value of custom (user-defined) parameter and put to
	// the provided receiver (should be a pointer). It returns an error in case if
	// parameter does not exist or receiver has an invalid type.
	GetProperty(name string, recv interface{}) error
	// SetProperty should add a custom parameter to the object (or replace if exists).
	SetProperty(name string, value interface{}) error
	// Range should call the provided func sequentially for each available property.
	// If func returns false, Range should stop the iteration.
	Range(func(property string, value interface{}) bool)
}

// MapManager is a PropertyManager implementation.
type MapManager struct {
	mu         sync.Mutex
	storage    map[string]interface{}
	validators map[string][]func(Manager, interface{}) error
}

// NewManager creates a new property manager.
func NewManager(validators ...Validator) *MapManager {
	manager := &MapManager{
		storage:    make(map[string]interface{}),
		validators: make(map[string][]func(Manager, interface{}) error),
	}
	for _, validator := range validators {
		manager.validators[validator.String()] = append(manager.validators[validator.String()], validator.Validate)
	}
	return manager
}

// GetProperty retrieves custom object property (if exists) and assignes it to the
// provided receiver (if possible).
func (m *MapManager) GetProperty(name string, recv interface{}) (err error) {
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
func (m *MapManager) SetProperty(name string, value interface{}) error {
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

// Range calls the provided func sequentially for each available property.
// If func returns false, Range stops the iteration.
func (m *MapManager) Range(f func(property string, value interface{}) bool) {
	for k, v := range m.storage {
		if !f(k, v) {
			break
		}
	}
}

// Value implements the database/sql/driver Valuer interface (needed for database
// drivers in order to store properties to DB).
func (m *MapManager) Value() (driver.Value, error) {
	return json.Marshal(m.storage)
}

// MarshalJSON implements custom JSON marshaller.
func (m *MapManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.storage)
}

// TODO: implement sql.Scanner interface
