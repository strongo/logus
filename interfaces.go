package logus

import "context"

// LogEntry hold data to be logged
type LogEntry struct {
	Severity      Severity
	MessageFormat string
	MessageArgs   []any
	Payload       any
}

// Logger defines logs dispatcher
type Logger interface {
	Log(ctx context.Context, entry LogEntry)
}

type SimpleLogger interface {
	Debugf(ctx context.Context, format string, args ...any)
	Infof(ctx context.Context, format string, args ...any)
	Errorf(ctx context.Context, format string, args ...any)
	Warningf(ctx context.Context, format string, args ...any)
	Criticalf(ctx context.Context, format string, args ...any)
}

// LogEntryHandler should implement persistence of logs entries
type LogEntryHandler interface {
	Log(ctx context.Context, entry LogEntry) error
}
