package scribe

import "testing"

func TestScribe_New(t *testing.T) {
	s := New("Testing")
	s.Notification("This is a notification log", nil)
	s.Success("This is a success log", nil)
	s.Warning("This is a warning log", nil)
	s.Error("This is an error log", nil)
	s.Debug("This is a debug log", nil)
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
	s.Notification("This is a notification log", nil)
	s.Success("This is a success log", nil)
	s.Warning("This is a warning log", nil)
	s.Error("This is an error log", nil)
	s.Debug("This is a debug log", nil)
}
