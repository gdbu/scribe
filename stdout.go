package scribe

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
)

const dot = "● "

var (
	green  = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)
	red    = color.New(color.FgRed)

	greenDot  = green.Sprint(dot)
	yellowDot = yellow.Sprint(dot)
	redDot    = red.Sprint(dot)
)

// NewStdout will create a new instance of Stdout
func NewStdout() *Stdout {
	var s Stdout
	s.typeMap = map[Type]string{
		TypeNotification: dot,
		TypeSuccess:      greenDot,
		TypeWarning:      yellowDot,
		TypeError:        redDot,
		TypeDebug:        dot,
	}

	return &s
}

// Stdout is a Writer used to push scribe entries to std output
type Stdout struct {
	mux sync.RWMutex

	typeMap map[Type]string
}

func (s *Stdout) getDataString(e *Entry) (out string) {
	if e.Data == nil {
		return
	}

	return fmt.Sprintf(" (%+v)", e.Data)
}

func (s *Stdout) Write(e *Entry) (err error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	prefix := s.typeMap[e.Type]
	data := s.getDataString(e)
	fmt.Println(prefix + e.Message + data)
	return
}

// SetTypePrefix will set the prefix string for a given Type
func (s *Stdout) SetTypePrefix(t Type, prefix string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.typeMap[t] = prefix
}
