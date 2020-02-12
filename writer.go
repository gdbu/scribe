package scribe

// Writer is an interface to write scribe entries
type Writer interface {
	Write(*Entry) error
}
