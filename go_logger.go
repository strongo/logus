package logus

import (
	"context"
	"fmt"
	"log"
)

type goLogger struct {
}

func (s goLogger) Log(_ context.Context, entry LogEntry) error {
	var severity string
	switch entry.Severity {
	case SeverityDebug:
		severity = "DEBUG"
	case SeverityDefault:
		severity = "DEFAULT"
	case SeverityInfo:
		severity = "INFO"
	case SeverityNotice:
		severity = "NOTICE"
	case SeverityWarning:
		severity = "WARNING"
	case SeverityError:
		severity = "ERROR"
	case SeverityCritical:
		severity = "CRITICAL"
	case SeverityAlert:
		severity = "ALERT"
	default:
		if entry.Component == "" {
			log.Printf("WARNING: unknown log arg severity: %d", entry.Severity)
		} else {
			log.Printf("WARNING: %s: unknown log arg severity: %d", entry.Component, entry.Severity)
		}
		severity = fmt.Sprintf("SEVERITY=%d", entry.Severity)
	}
	var message string
	if len(entry.MessageArgs) == 0 {
		message = entry.MessageFormat
	} else {
		message = fmt.Sprintf(entry.MessageFormat, entry.MessageArgs...)
	}
	if entry.Component == "" {
		logPrintf("%s: %s", severity, message)
	} else {
		logPrintf("%s: %s: %s", severity, entry.Component, message)
	}
	return nil
}

var logPrintf = log.Printf

var _ LogEntryHandler = (*goLogger)(nil)

func NewStandardGoLogger() LogEntryHandler {
	return goLogger{}
}
