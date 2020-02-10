package scribe

import "testing"

func TestScribe_New(t *testing.T) {
	s := New("Testing")
	s.Notification("Notification", nil)
	s.Success("Success", nil)
	s.Warning("Warning", nil)
	s.Error("Error", nil)
	s.Debug("Debug", nil)
}

func TestScribe_New_File(t *testing.T) {
	var (
		f   *File
		err error
	)

	if f, err = NewFile("test", "./"); err != nil {
		t.Fatalf("error creating file writer: %v", err)
	}

	f.SetLineCapacity(1)

	s := NewWithWriter(f, "Testing")
	s.Notification("Notification", nil)
	s.Success("Success", nil)
	s.Warning("Warning", nil)
	s.Error("Error", nil)
	s.Debug("Debug", nil)
}
