package logus

import (
	"context"
	"fmt"
	"os"
)

var dispatcher dispatchLogger

type dispatchLogger struct {
	handlers []LogEntryHandler
}

func (d *dispatchLogger) Log(ctx context.Context, entry LogEntry) {
	for _, h := range d.handlers[:] {
		if err := h.Log(ctx, entry); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ERROR: failed to log entry to log handler %T: %v", h, err)
		}
	}
}

func (d *dispatchLogger) AddLogEntryHandler(handler LogEntryHandler) {
	d.handlers = append(d.handlers, handler)
}
