package logus

import (
	"context"
	"errors"
	"testing"
)

type errorLogHandler struct{}

func (errorLogHandler) Log(ctx context.Context, entry LogEntry) error {
	return errors.New("mock log error")
}

func TestLoggerMethods(t *testing.T) {
	mockHandler := &testLogEntryHandler{}
	l := &logger{
		handlers: []LogEntryHandler{mockHandler},
	}
	ctx := context.Background()

	l.Debugf(ctx, "dbg")
	l.Defaultf(ctx, "def")
	l.Infof(ctx, "inf")
	l.Noticef(ctx, "not")
	l.Warningf(ctx, "wrn")
	l.Errorf(ctx, "err")
	l.Criticalf(ctx, "crt")
	l.Alertf(ctx, "alr")

	if len(mockHandler.entries) != 8 {
		t.Fatalf("expected 8 log entries, got %d", len(mockHandler.entries))
	}

	// Test handler returning error
	errHandler := errorLogHandler{}
	lErr := &logger{handlers: []LogEntryHandler{errHandler}}
	lErr.Log(ctx, LogEntry{Severity: SeverityInfo, MessageFormat: "test error handler"})

	// Test addLogEntryHandler
	t.Run("addLogEntryHandler", func(t *testing.T) {
		assertPanic := func(fn func()) {
			defer func() {
				if r := recover(); r == nil {
					t.Fatal("expected panic")
				}
			}()
			fn()
		}

		l2 := &logger{}
		assertPanic(func() { l2.addLogEntryHandler(nil) })

		h1 := &testLogEntryHandler{}
		l2.addLogEntryHandler(h1)
		// Duplicate handler should panic
		assertPanic(func() { l2.addLogEntryHandler(h1) })
	})
}
