package logus

import (
	"context"
	"testing"
)

func TestGetLogger(t *testing.T) {
	l := GetLogger()
	if l == nil {
		t.Errorf("Expected logger, got nil")
		return
	}
	ctx, _ := setupTestHandler()
	const message = "TestGetLogger"
	const severity = SeverityInfo
	l.Log(ctx, LogEntry{Severity: severity, MessageFormat: message})
	assertSingleLogEntry(t, ctx, LogEntry{Severity: severity, MessageFormat: message})
}

func TestGetSimpleLogger(t *testing.T) {
	l := GetSimpleLogger()
	if l == nil {
		t.Errorf("Expected logger, got nil")
		return
	}
	ctx, _ := setupTestHandler()
	const message = "TestGetLogger"
	l.Infof(ctx, message)
	assertSingleLogEntry(t, ctx, LogEntry{Severity: SeverityInfo, MessageFormat: message})
}

func setupTestHandler() (context.Context, *testLogEntryHandler) {
	testHandler := &testLogEntryHandler{}
	dispatcher.handlers = []LogEntryHandler{testHandler}
	return context.Background(), testHandler
}

func assertSingleLogEntry(t *testing.T, ctx context.Context, expected LogEntry) {
	testHandler := dispatcher.handlers[0].(*testLogEntryHandler)
	if count := len(testHandler.entries); count != 1 {
		t.Errorf("Expected 1 log arg, got %d", count)
	}
	actual := testHandler.entries[0]
	if actual.logEntry.Severity != expected.Severity {
		t.Errorf("Expected %v, got %v", expected.Severity, actual.logEntry.Severity)
		return
	}
	if actual.logEntry.MessageFormat != expected.MessageFormat {
		t.Errorf("Expected %q, got %q", expected.MessageFormat, actual.logEntry.MessageFormat)
		return
	}
	if len(actual.logEntry.MessageArgs) != len(expected.MessageArgs) {
		t.Errorf("Expected len(MessageArgs)=%d, got %d", len(expected.MessageArgs), len(actual.logEntry.MessageArgs))
		return
	}
	for i, expectedArg := range expected.MessageArgs {
		if actual.logEntry.MessageArgs[i] != expectedArg {
			t.Errorf("Expected MessageArgs[%d]=%v, got %v", i, expectedArg, actual.logEntry.MessageArgs[i])
		}
	}
	if actual.ctx != ctx {
		t.Errorf("Expected context does not match logged context")
	}
}

func TestLog(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	entry := LogEntry{
		Severity:      SeverityDebug,
		MessageFormat: "TestLog",
	}

	expected := entry // Copy input

	Log(ctx, entry)

	assertSingleLogEntry(t, ctx, expected)
}

func TestLogf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	severity := SeverityDebug
	format := "v1=%v"
	Logf(ctx, severity, format, 101)

	// Assert
	assertSingleLogEntry(t, ctx, LogEntry{Severity: severity, MessageFormat: format, MessageArgs: []any{101}})
}

func TestAddLogEntryHandler(t *testing.T) {
	// Setup
	dispatcher.handlers = make([]LogEntryHandler, 0, 1)

	// Perform
	handler := &testLogEntryHandler{}
	AddLogEntryHandler(handler)

	// Assert
	if len(dispatcher.handlers) != 1 {
		t.Errorf("Expected 1 handler, got %v", len(dispatcher.handlers))
	}
	if dispatcher.handlers[0] != handler {
		t.Errorf("Expected handler, got %v", dispatcher.handlers[0])
	}
}

func TestDebugf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Debugf(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, LogEntry{Severity: SeverityDebug, MessageFormat: "v1=%v", MessageArgs: []any{201}})
}

func TestDefaultf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Defaultf(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, LogEntry{Severity: SeverityDefault, MessageFormat: "v1=%v", MessageArgs: []any{201}})
}

func TestInfof(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Infof(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, LogEntry{Severity: SeverityInfo, MessageFormat: "v1=%v", MessageArgs: []any{201}})
}

func TestNoticef(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Noticef(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, LogEntry{Severity: SeverityNotice, MessageFormat: "v1=%v", MessageArgs: []any{201}})
}

func TestWarningf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Warningf(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, LogEntry{Severity: SeverityWarning, MessageFormat: "v1=%v", MessageArgs: []any{201}})
}
func TestErrorf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Errorf(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, LogEntry{Severity: SeverityError, MessageFormat: "v1=%v", MessageArgs: []any{201}})
}

func TestCriticalf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Criticalf(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, LogEntry{Severity: SeverityCritical, MessageFormat: "v1=%v", MessageArgs: []any{201}})
}

func TestAlertf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Alertf(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, LogEntry{Severity: SeverityAlert, MessageFormat: "v1=%v", MessageArgs: []any{201}})
}
