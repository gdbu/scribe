package scribe

import "os"

// NewStdout will create a new instance of Stdout
func NewStdout() *Basic {
	b := NewBasic(os.Stdout)
	b.SetSuffix("\n")
	return b
}
