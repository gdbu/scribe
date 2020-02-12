package scribe

import (
	"log/syslog"
)

// NewSyslog will create a new instance of Syslog
func NewSyslog(name, addr string) (sp *Syslog, err error) {
	var s Syslog
	if s.w, err = syslog.Dial("udp", addr, syslog.LOG_INFO, name); err != nil {
		return
	}

	s.Basic = NewBasic(s.w)
	sp = &s
	return
}

// Syslog is a Writer used to push scribe entries to std output
type Syslog struct {
	*Basic
	w *syslog.Writer
}

func (s *Syslog) Write(e *Entry) (err error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	msg := s.getMessage(e)

	switch e.Type {
	case TypeWarning:
		return s.w.Warning(msg)
	case TypeError:
		return s.w.Err(msg)
	case TypeDebug:
		return s.w.Debug(msg)

	default:
		// All the types not listed can be covered by the info log type
		return s.w.Info(msg)
	}
}

// Close will close an instance of a syslog writer
func (s *Syslog) Close() (err error) {
	return s.w.Close()
}
