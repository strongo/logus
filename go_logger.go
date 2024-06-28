package logus

import (
	"context"
	"fmt"
	"log"
)

type goLogger struct{}

func (s goLogger) Log(_ context.Context, entry LogEntry) error {
	var severity string
	switch entry.Severity {
	case SeverityDefault:
		severity = "DEFAULT"
	case SeverityDebug:
		severity = "DEBUG"
	case SeverityInfo:
		severity = "INFO"
	case SeverityWarning:
		severity = "WARNING"
	case SeverityError:
		severity = "ERROR"
	case SeverityCritical:
		severity = "CRITICAL"
	default:
		severity = fmt.Sprintf("SEVERITY=%d", entry.Severity)
	}
	log.Printf("%s: %s", severity, entry.Message)
	return nil
}

var _ LogEntryHandler = (*goLogger)(nil)

func StandardGoLogger() LogEntryHandler {
	return goLogger{}
}
