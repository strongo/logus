package logus

import (
	"context"
	"testing"
)

func TestNewSimpleLogger(t *testing.T) {
	setupTestHandler()
	sl := NewSimpleLogger(NewComponentLogger("UnitTest"))

	tests := []struct {
		name     string
		logFn    func(ctx context.Context, format string, args ...any)
		severity Severity
	}{
		{"debug", sl.Debugf, SeverityDebug},
		{"default", sl.Defaultf, SeverityDefault},
		{"info", sl.Infof, SeverityInfo},
		{"notice", sl.Noticef, SeverityNotice},
		{"warning", sl.Warningf, SeverityWarning},
		{"error", sl.Errorf, SeverityError},
		{"critical", sl.Criticalf, SeverityCritical},
		{"alert", sl.Alertf, SeverityAlert},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCtx, testHandler := setupTestHandler()
			format := "msg %s"
			args := []any{"arg"}
			tt.logFn(testCtx, format, args...)
			assertSingleLogEntry(t, testCtx, LogEntry{Severity: tt.severity, MessageFormat: format, MessageArgs: args, Component: "UnitTest"})
			_ = testHandler
		})
	}
}
