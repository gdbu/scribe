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
