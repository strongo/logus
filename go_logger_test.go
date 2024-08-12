package logus

import (
	"context"
	"testing"
)

func Test_goLogger_Log(t *testing.T) {
	type args struct {
		entry LogEntry
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "debug", args: args{entry: LogEntry{Severity: SeverityDebug, MessageFormat: "debug message"}}},
		{name: "info", args: args{entry: LogEntry{Severity: SeverityInfo, MessageFormat: "informational message"}}},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := goLogger{}
			err := s.Log(ctx, tt.args.entry)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestStandardGoLogger(t *testing.T) {
	logEntryHandler := StandardGoLogger()
	if logEntryHandler == nil {
		t.Errorf("StandardGoLogger() = nil")
	}
	if _, ok := logEntryHandler.(LogEntryHandler); !ok {
		t.Errorf("StandardGoLogger() expected %T, got %T", &goLogger{}, logEntryHandler)
	}
}
