package logus

import (
	"context"
	"fmt"
	"log"
)

func NewStandardGoLogger() LogEntryHandler {
	return goLogger{}
}

var logPrintf = log.Printf

var _ LogEntryHandler = (*goLogger)(nil)

type goLogger struct {
}

func (s goLogger) Log(_ context.Context, entry LogEntry) error {
	var severity string
	if int(entry.Severity) < len(SeverityNames) {
		severity = SeverityNames[entry.Severity]
	} else {
		severity = fmt.Sprintf("SEVERITY%d", entry.Severity)
		if entry.Component == "" {
			logPrintf("WARNING: unknown log arg severity: %d", entry.Severity)
		} else {
			logPrintf("WARNING: %s: unknown log arg severity: %d", entry.Component, entry.Severity)
		}
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
		logPrintf("%s: Component=%s: %s", severity, entry.Component, message)
	}
	return nil
}
