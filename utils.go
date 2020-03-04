package scribe

import (
	"fmt"
	"path"
	"runtime"

	"github.com/hatchify/colors"
)

const dot = "● "

var (
	green  = colors.New(colors.FGGreen)
	yellow = colors.New(colors.FGYellow)
	red    = colors.New(colors.FGRed)

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
