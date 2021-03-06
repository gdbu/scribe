package scribe

import (
	"fmt"
	"io"
)

var g *Basic = NewStdout()

const debugFmt = "%s:%d :: %s"

// New will return a new Scribe
func New(prefix string) *Scribe {
	return NewWithWriter(g, prefix)
}

// NewWithWriter will return a new Scribe with a provided Writer
func NewWithWriter(w Writer, prefix string) *Scribe {
	var s Scribe
	s.w = w
	s.v = VerbosityAll

	if len(prefix) > 0 {
		s.prefix = prefix + " :: "
	}

	return &s
}

// Scribe will create new scribe entries
type Scribe struct {
	w Writer
	v Verbosity

	prefix string
}

// new will append a new scribe Entry
func (s *Scribe) new(t Type, msg string, data interface{}) {
	switch {
	case s.v == VerbosityAll:
	case s.v == VerbosityNone:
		return
	case s.v&t.Verbosity() != 0:

	default:
		return
	}

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
func (s *Scribe) Notification(msg string) {
	s.new(TypeNotification, msg, nil)
}

// Notificationf will create a new notificaton entry with a format message
func (s *Scribe) Notificationf(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	s.Notification(msg)
}

// NotificationWithData will create a new notificaton entry with Data
func (s *Scribe) NotificationWithData(msg string, data interface{}) {
	s.new(TypeNotification, msg, data)
}

// NotificationWithDataf will create a new notificaton entry with Data and a format message
func (s *Scribe) NotificationWithDataf(msg string, data interface{}, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	s.NotificationWithData(msg, data)
}

// Success will create a new success entry
func (s *Scribe) Success(msg string) {
	s.new(TypeSuccess, msg, nil)
}

// Successf will create a new success entry with a format message
func (s *Scribe) Successf(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	s.Success(msg)
}

// SuccessWithData will create a new success entry with Data
func (s *Scribe) SuccessWithData(msg string, data interface{}) {
	s.new(TypeSuccess, msg, data)
}

// SuccessWithDataf will create a new success entry with Data and a format message
func (s *Scribe) SuccessWithDataf(msg string, data interface{}, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	s.SuccessWithData(msg, data)
}

// Warning will create a new warning entry
func (s *Scribe) Warning(msg string) {
	s.new(TypeWarning, msg, nil)
}

// Warningf will create a new warning entry with a format message
func (s *Scribe) Warningf(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	s.Warning(msg)
}

// WarningWithData will create a new warning entry with Data
func (s *Scribe) WarningWithData(msg string, data interface{}) {
	s.new(TypeWarning, msg, data)
}

// WarningWithDataf will create a new warning entry with Data and a format message
func (s *Scribe) WarningWithDataf(msg string, data interface{}, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	s.WarningWithData(msg, data)
}

// Error will create a new error entry
func (s *Scribe) Error(msg string) {
	s.new(TypeError, msg, nil)
}

// Errorf will create a new error entry with a format message
func (s *Scribe) Errorf(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	s.Error(msg)
}

// ErrorWithData will create a new error entry with Data
func (s *Scribe) ErrorWithData(msg string, data interface{}) {
	s.new(TypeError, msg, data)
}

// ErrorWithDataf will create a new error entry with Data and a format message
func (s *Scribe) ErrorWithDataf(msg string, data interface{}, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	s.ErrorWithData(msg, data)
}

// Debug will create a new debug entry
func (s *Scribe) Debug(msg string) {
	filename, lineNumber := getDebugVals()
	// Prepend the message with the caller's filename and line number
	msg = fmt.Sprintf(debugFmt, filename, lineNumber, msg)
	s.new(TypeDebug, msg, nil)
}

// Debugf will create a new debug entry with a format message
func (s *Scribe) Debugf(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	s.Debug(msg)
}

// DebugWithData will create a new debug entry with Data
func (s *Scribe) DebugWithData(msg string, data interface{}) {
	filename, lineNumber := getDebugVals()
	// Prepend the message with the caller's filename and line number
	msg = fmt.Sprintf(debugFmt, filename, lineNumber, msg)
	s.new(TypeDebug, msg, data)
}

// DebugWithDataf will create a new debug entry with Data and a format message
func (s *Scribe) DebugWithDataf(msg string, data interface{}, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	s.DebugWithData(msg, data)
}

// SetGlobalWriter will set the global writer
func SetGlobalWriter(w io.Writer) {
	g.setWriter(w)
}

// SetVerbosity will set the verbosity of a scribe instance
func SetVerbosity(s *Scribe, v Verbosity) {
	s.v = v
}
