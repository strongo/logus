package logus

import (
	"context"
	"testing"
)

func setupTestHandler() (context.Context, *testLogEntryHandler) {
	testHandler := &testLogEntryHandler{}
	dispatcher.handlers = []LogEntryHandler{testHandler}
	return context.Background(), testHandler
}

func assertSingleLogEntry(t *testing.T, ctx context.Context, severity Severity, message string) {
	testHandler := dispatcher.handlers[0].(*testLogEntryHandler)
	if count := len(testHandler.entries); count != 1 {
		t.Errorf("Expected 1 log entry, got %d", count)
	}
	logged := testHandler.entries[0]
	if logged.logEntry.Severity != severity {
		t.Errorf("Expected %v, got %v", severity, logged.logEntry.Severity)
	}
	if logged.ctx != ctx {
		t.Errorf("Expected context.Background(), got %v", logged.ctx)
	}
	if logged.logEntry.Message != message {
		t.Errorf("Expected %s, got %s", message, logged.logEntry.Message)
	}
}

func TestLog(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	entry := LogEntry{
		Severity: SeverityDebug,
		Message:  "TestLog",
	}
	Log(ctx, entry)

	// Assert
	assertSingleLogEntry(t, ctx, SeverityDebug, entry.Message)
}

func TestLogf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	severity := SeverityDebug
	format := "v1=%v"
	Logf(ctx, severity, format, 101)

	// Assert
	assertSingleLogEntry(t, ctx, severity, "v1=101")
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
	assertSingleLogEntry(t, ctx, SeverityDebug, "v1=201")
}

func TestDefaultf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Defaultf(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, SeverityDefault, "v1=201")
}

func TestInfof(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Infof(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, SeverityInfo, "v1=201")
}

func TestNoticef(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Noticef(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, SeverityNotice, "v1=201")
}

func TestWarningf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Warningf(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, SeverityWarning, "v1=201")
}
func TestErrorf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Errorf(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, SeverityError, "v1=201")
}

func TestCriticalf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Criticalf(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, SeverityCritical, "v1=201")
}

func TestAlertf(t *testing.T) {
	// Setup
	ctx, _ := setupTestHandler()

	// Perform
	Alertf(ctx, "v1=%v", 201)

	// Assert
	assertSingleLogEntry(t, ctx, SeverityAlert, "v1=201")
}
