package logus

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
)

func Test_goLogger_Log(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		arg      LogEntry
		expected string
	}{
		{
			name:     "debug_nil_context",
			ctx:      nil,
			arg:      LogEntry{Severity: SeverityDebug, MessageFormat: "debugging message"},
			expected: "DEBUG: debugging message",
		},
		{
			name:     "info_with_component",
			ctx:      context.Background(),
			arg:      LogEntry{Severity: SeverityInfo, Component: "UnitTest", MessageFormat: "informational message"},
			expected: "INFO: UnitTest: informational message"},
	}
	defer func() {
		logPrintf = log.Printf
	}()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var logPrintsCount int
			var logString string
			// Setup mocks
			logPrintf = func(format string, v ...interface{}) {
				logPrintsCount++
				logString = fmt.Sprintf(format, v...)
			}
			s := NewStandardGoLogger()
			err := s.Log(tt.ctx, tt.arg)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if logPrintsCount != 1 {
				t.Errorf("expected 1 log print, got %d", logPrintsCount)
			}
			if !strings.Contains(logString, tt.expected) {
				t.Errorf("expected log message to contain %q, got %q", tt.expected, logString)
			}
		})
	}
}
