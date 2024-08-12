package logus

import (
	"context"
	"fmt"
)

var dispatcher logger

func GetLogger() Logger {
	return &dispatcher
}

func GetSimpleLogger() SimpleLogger {
	return &dispatcher
}

func AddLogEntryHandler(handler LogEntryHandler) {
	dispatcher.addLogEntryHandler(handler)
}

func Log(ctx context.Context, entry LogEntry) {
	dispatcher.Log(ctx, entry)
}

func Logf(ctx context.Context, severity Severity, format string, args ...any) {
	logEntry := LogEntry{
		Severity:      severity,
		MessageFormat: fmt.Sprintf(format, args...),
	}
	Log(ctx, logEntry)
}

func Debugf(ctx context.Context, format string, args ...any) {
	Logf(ctx, SeverityDebug, format, args...)
}

func Defaultf(ctx context.Context, format string, args ...any) {
	Logf(ctx, SeverityDefault, format, args...)
}

func Infof(ctx context.Context, format string, args ...any) {
	Logf(ctx, SeverityInfo, format, args...)
}

func Noticef(ctx context.Context, format string, args ...any) {
	Logf(ctx, SeverityNotice, format, args...)
}

func Warningf(ctx context.Context, format string, args ...any) {
	Logf(ctx, SeverityWarning, format, args...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	Logf(ctx, SeverityError, format, args...)
}

func Criticalf(ctx context.Context, format string, args ...any) {
	Logf(ctx, SeverityCritical, format, args...)
}

func Alertf(ctx context.Context, format string, args ...any) {
	Logf(ctx, SeverityAlert, format, args...)
}
