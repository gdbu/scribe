package scribe

const (
	VerbosityNone Verbosity = 1 << iota
	VerbosityAll
	VerbositySuccesses
	VerbosityNotifications
	VerbosityWarnings
	VerbosityErrors
)

type Verbosity uint8
