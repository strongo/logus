package logus

import "context"

// LogEntry hold data to be logged
type LogEntry struct {
	Severity Severity
	Message  string
	Payload  any
}

// Logger defines logs dispatcher
type Logger interface {
	Log(ctx context.Context, entry LogEntry)
}

// LogEntryHandler should implement persistence of logs entries
type LogEntryHandler interface {
	Log(ctx context.Context, entry LogEntry) error
}
