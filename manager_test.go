package property

import (
	"testing"
)

var _ Manager = (*MapManager)(nil)

func Test_Properties(t *testing.T) {
	t.Run("Given option storage", func(t *testing.T) {
		manager := NewManager()
		t.Run("check if manager is able to save properties", func(t *testing.T) {
			manager.SetProperty("foo", "bar")
			manager.SetProperty("answer", int(42))
			if val, ok := manager.storage["foo"]; !ok || val != "bar" {
				t.Error("storage does not contain expected value")
			}
			if val, ok := manager.storage["answer"]; !ok || val != int(42) {
				t.Error("storage does not contain expected value")
			}
		})
		t.Run("check if storage is able to assign the value", func(t *testing.T) {
			var str string
			func() {
				if err := manager.GetProperty("bar", &str); err != ErrNotAvailable {
					t.Errorf("error %q was expected to be %q", err, ErrNotAvailable)
				}
			}()
			func() {
				if err := manager.GetProperty("foo", str); err != ErrNotAPointer {
					t.Errorf("error %q was expected to be %q", err, ErrNotAPointer)
				}
			}()
			var num int
			func() {
				if err := manager.GetProperty("foo", &num); err != ErrCannotAssign {
					t.Errorf("error %q was expected to be %q", err, ErrCannotAssign)
				}
			}()
			func() {
				if err := manager.GetProperty("foo", &str); err != nil {
					t.Error("should not return an error")
				}
				if str != "bar" {
					t.Errorf("string value %q was expected to be %q", str, "bar")
				}
				if err := manager.GetProperty("answer", &num); err != nil {
					t.Error("should not return an error")
				}
				if num != int(42) {
					t.Errorf("int value %d was expected to be %v", num, 42)
				}
			}()
		})
	})
}
