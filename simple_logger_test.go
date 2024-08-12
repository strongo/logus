package logus

import (
	"context"
	"testing"
)

func TestNewSimpleLogger(t *testing.T) {
	tests := []struct {
		name     string
		severity Severity
		format   string
		args     []any
		want     string
	}{
		{
			name:   "debug",
			format: "debugging %s %d",
			args:   []any{"abc", 123},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestHandler()
			simpleLogger := NewSimpleLogger(NewComponentLogger("UnitTest"))
			ctx := context.Background()
			simpleLogger.Debugf(ctx, tt.format, tt.args...)
			assertSingleLogEntry(t, ctx, LogEntry{Severity: SeverityDebug, MessageFormat: tt.format, MessageArgs: tt.args})
		})
	}
}
