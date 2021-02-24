package scribe

const (
	VerbosityAll       Verbosity = 0
	VerbositySuccesses Verbosity = 1 << iota
	VerbosityNotifications
	VerbosityWarnings
	VerbosityErrors
	VerbosityNone
)

type Verbosity uint8
