// log/syslog is not implemented for windows or plan9
// 	https://golang.org/src/log/syslog/syslog.go
// +build !windows,!plan9

package scribe

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
