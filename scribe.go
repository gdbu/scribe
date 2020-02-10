package scribe

import "fmt"

var stdout = NewStdout()

const debugFmt = "%s:%d :: %s"

// New will return a new Scribe
func New(prefix string) *Scribe {
	return NewWithWriter(stdout, prefix)
}

// NewWithWriter will return a new Scribe with a provided Writer
func NewWithWriter(w Writer, prefix string) *Scribe {
	var s Scribe
	s.w = w
	if len(prefix) > 0 {
		s.prefix = prefix + " :: "
	}

	return &s
}

// Scribe will create new scribe entries
type Scribe struct {
	w Writer

	prefix string
}

// new will append a new scribe Entry
func (s *Scribe) new(t Type, msg string, data interface{}) {
	// Prepend prefix to message
	msg = s.prefix + msg
	// Create new entry from provided values
	e := newEntry(t, msg, data)
	// Write entry to writer
	if err := s.w.Write(e); err != nil {
		fmt.Printf("error writing Scribe entry: %v\n", err)
	}
}

// Notification will create a new notificaton entry
func (s *Scribe) Notification(msg string, data interface{}) {
	s.new(TypeNotification, msg, data)
}

// Success will create a new success entry
func (s *Scribe) Success(msg string, data interface{}) {
	s.new(TypeSuccess, msg, data)
}

// Warning will create a new warning entry
func (s *Scribe) Warning(msg string, data interface{}) {
	s.new(TypeWarning, msg, data)
}

// Error will create a new error entry
func (s *Scribe) Error(msg string, data interface{}) {
	s.new(TypeError, msg, data)
}

// Debug will create a new debug entry
func (s *Scribe) Debug(msg string, data interface{}) {
	filename, lineNumber := getDebugVals()
	// Prepend the message with the caller's filename and line number
	msg = fmt.Sprintf(debugFmt, filename, lineNumber, msg)
	s.new(TypeDebug, msg, data)
}