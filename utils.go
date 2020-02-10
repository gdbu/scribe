package scribe

import (
	"fmt"
	"path"
	"runtime"
)

func newErrorCatch(fn func() error) func() {
	return func() {
		var err error
		if err = fn(); err != nil {
			fmt.Println(err)
		}
	}
}

func getDebugVals() (filename string, lineNumber int) {
	_, filename, lineNumber, _ = runtime.Caller(2)
	filename = path.Clean(filename)
	return
}
