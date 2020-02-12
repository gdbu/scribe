package scribe

import (
	"fmt"
	"path"
	"runtime"

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
