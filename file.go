package scribe

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/Hatch1fy/cron"
	"github.com/hatchify/errors"
)

const (
	// ErrFileUnset is returned when a file is not set and attempting to be written to
	ErrFileUnset = errors.Error("cannot write, file is currently unset")
	// ErrRotationIntervalSet is returned when a rotation interval is attempting to be set when it's already been set
	ErrRotationIntervalSet = errors.Error("error setting rotation interval, value already set")
)

// NewFile will return a new Scribe
func NewFile(name, dir string) (fp *File, err error) {
	var f File
	f.name = name
	f.dir = dir

	if err = f.openFile(); err != nil {
		return
	}

	fp = &f
	return
}

// File manages a file output writer for a Scribe
type File struct {
	mux sync.RWMutex
	f   *os.File

	name string
	dir  string

	// Interval of how often file is rotated
	// Note: Rotation will not occur if duration is unset
	rotateInterval time.Duration
	// Maximum number of lines before file is rotated
	// Note: Line rotation will not occur if line capacity is unset
	lineCapacity int64

	// Current line count
	lineCount int64
}

func (f *File) openFile() (err error) {
	// Create filename from directory and service name
	filename := path.Join(f.dir, f.name+".log")

	// Open target file
	if f.f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0744); err != nil {
		err = fmt.Errorf("error opening scribe file for \"%s\": %v", filename, err)
		return
	}

	return
}

func (f *File) rotateFile() (err error) {
	if f.lineCount == 0 {
		return
	}

	originalName := f.f.Name()
	// Set the directory
	dir := filepath.Dir(originalName)
	// Set current file name
	name := filepath.Base(originalName)
	// Set new file name as the current name with a timestamp prefix
	newName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), name)
	// Set new filename (dir + name)
	newFilename := filepath.Join(dir, newName)

	if err = f.f.Close(); err != nil {
		return
	}

	if err = os.Rename(originalName, newFilename); err != nil {
		return
	}

	return f.openFile()
}

func (f *File) incrementCount() (err error) {
	if f.lineCount++; f.lineCapacity == 0 || f.lineCapacity > f.lineCount {
		return
	}

	return f.rotateFile()
}

// Write will write a new Scribe entry
func (f *File) Write(e *Entry) (err error) {
	f.mux.Lock()
	defer f.mux.Unlock()
	if f.f == nil {
		return ErrFileUnset
	}

	if err = json.NewEncoder(f.f).Encode(e); err != nil {
		return
	}

	if err = f.incrementCount(); err != nil {
		err = fmt.Errorf("error rotating file: %v", err)
		return
	}

	return
}

// Rotate will trigger a file to rotate
// Note: This can be used to manually trigger a rotation for any reason
func (f *File) Rotate() (err error) {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.rotateFile()
}

// SetLineCapacity will set the line capacity
func (f *File) SetLineCapacity(lines int64) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.lineCapacity = lines
}

// SetRotationInterval will set the rotation interval
func (f *File) SetRotationInterval(d time.Duration) (err error) {
	f.mux.Lock()
	defer f.mux.Unlock()

	if f.rotateInterval > 0 {
		return ErrRotationIntervalSet
	}

	f.rotateInterval = d
	cron.New(newErrorCatch(f.Rotate)).Every(f.rotateInterval)
	return
}
