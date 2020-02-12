# Scribe [![GoDoc](https://godoc.org/github.com/hatchify/scribe?status.svg)](https://godoc.org/github.com/hatchify/scribe) ![Status](https://img.shields.io/badge/status-beta-yellow.svg)
Scribe is an output logger helper library a few simple features:

* Thread-safe
* Colored dot prefix for status (success, warning, error, debug)
* Debug output (with filename and line number)

![screenshot](https://github.com/hatchify/scribe/blob/master/screenshot.png?raw=true "Screenshot of scribe")

## Usage
The primary usage of scribe is utilizing the package-level logger. See below for examples of the available methods:

### Success
```go 
func ExampleSuccess() {
	s.Success("This is a success message!", nil)
}
```

### Warning
```go 
func ExampleWarning() {
	s.Warning("This is a warning message!", nil)
}
```

### Error
```go 
func ExampleError() {
	s.Error("This is an error message!", nil)
}
```

### Debug
```go 
func ExampleDebug() {
	s.Debug("This is a debug message!", nil)
}
```










