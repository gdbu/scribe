package scribe

const (
	VerbosityAll       Verbosity = 0
	VerbositySuccesses Verbosity = 1 << iota
	VerbosityNotifications
	VerbosityWarnings
	VerbosityErrors
)

type Verbosity uint8
