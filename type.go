package scribe

const (
	// TypeNotification represents a notification entry
	TypeNotification = "notification"
	// TypeSuccess represents a success entry
	TypeSuccess = "success"
	// TypeWarning represents a warning entry
	TypeWarning = "warning"
	// TypeError represents an error entry
	TypeError = "error"
	// TypeDebug represents a debug entry
	TypeDebug = "debug"
)

// Type represents an entry type
type Type string

func (t Type) Verbosity() Verbosity {
	switch t {
	case TypeNotification:
		return VerbosityNotifications
	case TypeSuccess:
		return VerbositySuccesses
	case TypeWarning:
		return VerbosityWarnings
	case TypeError:
		return VerbosityErrors
	case TypeDebug:
		return VerbosityErrors

	default:
		return VerbosityAll
	}
}
