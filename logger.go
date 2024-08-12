package logus

import (
	"context"
	"fmt"
	"os"
)

var _ SimpleLogger = (*logger)(nil)

type logger struct {
	handlers []LogEntryHandler
}

func (d *logger) Debugf(ctx context.Context, format string, args ...any) {
	d.Log(ctx, LogEntry{Severity: SeverityDebug, MessageFormat: format, MessageArgs: args})
}

func (d *logger) Defaultf(ctx context.Context, format string, args ...any) {
	d.Log(ctx, LogEntry{Severity: SeverityDefault, MessageFormat: format, MessageArgs: args})
}

func (d *logger) Infof(ctx context.Context, format string, args ...any) {
	d.Log(ctx, LogEntry{Severity: SeverityInfo, MessageFormat: format, MessageArgs: args})
}

func (d *logger) Noticef(ctx context.Context, format string, args ...any) {
	d.Log(ctx, LogEntry{Severity: SeverityNotice, MessageFormat: format, MessageArgs: args})
}

func (d *logger) Warningf(ctx context.Context, format string, args ...any) {
	d.Log(ctx, LogEntry{Severity: SeverityWarning, MessageFormat: format, MessageArgs: args})
}

func (d *logger) Errorf(ctx context.Context, format string, args ...any) {
	d.Log(ctx, LogEntry{Severity: SeverityError, MessageFormat: format, MessageArgs: args})
}

func (d *logger) Criticalf(ctx context.Context, format string, args ...any) {
	d.Log(ctx, LogEntry{Severity: SeverityCritical, MessageFormat: format, MessageArgs: args})
}

func (d *logger) Alertf(ctx context.Context, format string, args ...any) {
	d.Log(ctx, LogEntry{Severity: SeverityAlert, MessageFormat: format, MessageArgs: args})
}

func (d *logger) Log(ctx context.Context, entry LogEntry) {
	for _, h := range d.handlers[:] {
		if err := h.Log(ctx, entry); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ERROR: failed to log entry to log handler %T: %v", h, err)
		}
	}
}

func (d *logger) addLogEntryHandler(handler LogEntryHandler) {
	d.handlers = append(d.handlers, handler)
}
