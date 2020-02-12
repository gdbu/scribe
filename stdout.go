package scribe

import "os"

// NewStdout will create a new instance of Stdout
func NewStdout() *Stdout {
	var s Stdout
	s.Basic = NewBasic(os.Stdout)
	s.Basic.SetSuffix("\n")
	return &s
}

// Stdout is a Writer used to push scribe entries to std output
type Stdout struct {
	*Basic
}
