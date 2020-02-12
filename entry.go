package scribe

import (
	"fmt"
	"time"
)

func newEntry(t Type, msg string, data interface{}) *Entry {
	var e Entry
	e.Timestamp = time.Now()
	e.Type = t
	e.Message = msg
	e.Data = data
	return &e
}

// Entry represents a scribe entry
type Entry struct {
	// Timestamp of entry
	Timestamp time.Time `json:"timestamp"`
	// Type of entry
	Type Type `json:"type"`
	// Message of entry
	Message string `json:"message"`
	// Any additional data associated with entry
	Data interface{} `json:"data"`
}

func (e *Entry) getDataString() (out string) {
	if e.Data == nil {
		return
	}

	return fmt.Sprintf(" (%+v)", e.Data)
}

func (e *Entry) String() string {
	data := e.getDataString()
	return e.Message + data
}
