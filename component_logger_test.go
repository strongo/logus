package logus

import (
	"testing"
)

func TestNewComponentLogger(t *testing.T) {
	const component = "test1"
	logger := NewComponentLogger(component)
	ctx, testLogHandler := setupTestHandler()

	const message = "test message"
	logger.Log(ctx, LogEntry{Severity: SeverityInfo, MessageFormat: message})
	assertSingleLogEntry(t, ctx, SeverityInfo, message)
	testLogHandler.entries[0].logEntry.Component = component
}
