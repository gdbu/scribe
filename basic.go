package scribe

import (
	"io"
	"sync"
)

// NewBasic will create a new instance of Basic
func NewBasic(w io.Writer) *Basic {
	var b Basic
	b.w = w
	b.typeMap = map[Type]string{
		TypeNotification: dot,
		TypeSuccess:      greenDot,
		TypeWarning:      yellowDot,
		TypeError:        redDot,
		TypeDebug:        dot,
	}

	return &b
}

// Basic is a Writer used to push scribe entries to std output
type Basic struct {
	mux sync.RWMutex
	w   io.Writer

	typeMap map[Type]string
	suffix  string
}

func (b *Basic) getMessage(e *Entry) (msg string) {
	prefix := b.typeMap[e.Type]
	return prefix + e.String() + b.suffix
}

func (b *Basic) Write(e *Entry) (err error) {
	b.mux.Lock()
	defer b.mux.Unlock()
	bs := []byte(b.getMessage(e))
	b.w.Write(bs)
	return
}

// SetTypePrefix will set the prefix string for a given Type
func (b *Basic) SetTypePrefix(t Type, prefix string) {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.typeMap[t] = prefix
}

// SetSuffix will set the message suffix string
func (b *Basic) SetSuffix(suffix string) {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.suffix = suffix
}

func (b *Basic) setWriter(w io.Writer) {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.w = w
}
