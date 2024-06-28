package logus

import "context"

type testLogEntryHandler struct {
	entries []struct {
		logEntry LogEntry
		ctx      context.Context
	}
}

func (h *testLogEntryHandler) Log(ctx context.Context, entry LogEntry) error {
	h.entries = append(h.entries, struct {
		logEntry LogEntry
		ctx      context.Context
	}{
		logEntry: entry,
		ctx:      ctx,
	})
	return nil
}
