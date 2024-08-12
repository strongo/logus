package logus

import "context"

func NewSimpleLogger(logger Logger) SimpleLogger {
	return simpleLogger{logger: logger}
}

type simpleLogger struct {
	logger Logger
}

func (s simpleLogger) log(ctx context.Context, severity Severity, format string, args ...any) {
	logEntry := LogEntry{
		Severity:      severity,
		MessageFormat: format,
		MessageArgs:   args,
	}
	s.logger.Log(ctx, logEntry)
}

func (s simpleLogger) Debugf(ctx context.Context, format string, args ...any) {
	s.log(ctx, SeverityDebug, format, args...)
}

func (s simpleLogger) Defaultf(ctx context.Context, format string, args ...any) {
	s.log(ctx, SeverityDefault, format, args...)
}

func (s simpleLogger) Infof(ctx context.Context, format string, args ...any) {
	s.log(ctx, SeverityInfo, format, args...)
}

func (s simpleLogger) Noticef(ctx context.Context, format string, args ...any) {
	s.log(ctx, SeverityNotice, format, args...)
}

func (s simpleLogger) Errorf(ctx context.Context, format string, args ...any) {
	s.log(ctx, SeverityError, format, args...)
}

func (s simpleLogger) Warningf(ctx context.Context, format string, args ...any) {
	s.log(ctx, SeverityWarning, format, args...)
}

func (s simpleLogger) Criticalf(ctx context.Context, format string, args ...any) {
	s.log(ctx, SeverityCritical, format, args...)
}

func (s simpleLogger) Alertf(ctx context.Context, format string, args ...any) {
	s.log(ctx, SeverityAlert, format, args...)
}
