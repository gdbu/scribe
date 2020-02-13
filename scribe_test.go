package scribe

import (
	"os"
	"testing"
)

func TestScribe_New(t *testing.T) {
	s := New("Testing")
	s.Notification("This is a notification log")
	s.Success("This is a success log")
	s.Warning("This is a warning log")
	s.Error("This is an error log")
	s.Debug("This is a debug log")
}

func TestScribe_New_File(t *testing.T) {
	var (
		f   *File
		err error
	)

	if err = os.MkdirAll("./test_data", 0744); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll("./test_data")

	if f, err = NewFile("test", "./test_data"); err != nil {
		t.Fatalf("error creating file writer: %v", err)
	}

	f.SetLineCapacity(1)

	s := NewWithWriter(f, "Testing")
	s.Notification("This is a notification log")
	s.Success("This is a success log")
	s.Warning("This is a warning log")
	s.Error("This is an error log")
	s.Debug("This is a debug log")
}

func TestScribe_New_Syslog(t *testing.T) {
	addr := os.Getenv("SCRIBE_SYSLOG_ADDR")
	if len(addr) == 0 {
		t.Fatal("no syslog address env found, please set SCRIBE_SYSLOG_ADDR")
	}

	var (
		w   *Syslog
		err error
	)

	if w, err = NewSyslog("test", addr); err != nil {
		t.Fatalf("error creating syslog writer: %v", err)
	}
	defer w.Close()

	s := NewWithWriter(w, "Testing")
	s.Notification("This is a notification log")
	s.Success("This is a success log")
	s.Warning("This is a warning log")
	s.Error("This is an error log")
	s.Debug("This is a debug log")
}
