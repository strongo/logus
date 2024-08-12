package logus

import "context"

type componentLogger struct {
	component string
}

func NewComponentLogger(component string) Logger {
	return &componentLogger{component: component}
}

func (logger *componentLogger) Log(ctx context.Context, entry LogEntry) {
	if entry.Component == "" {
		entry.Component = logger.component
	} else {
		entry.Component = logger.component + "." + entry.Component
	}
	Log(ctx, entry)
}
