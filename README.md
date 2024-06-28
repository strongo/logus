# Package: [github.com/strongo/logus](https://github.com/strongo/logus)

Package strongo/logus provides structured context aware logging for Go (golang)
that supports grouping of log entries using trace and slot IDs.

## Usage

```go
package foo

import "context"
import "github.com/strongo/logus"

// Bar demonstrates how to use logus logger 
func Bar(c context.Context) {

	logus.Debugf(c, "This is a debug message without trance ID, unless it was set outside")
  
	const traceID = "123"
	ct := logus.WithTraceID(c, traceID)
	logus.Infof(ct, "This is an info message with a traceID=%s", traceID)

	const slotID = "456"
	ct = logus.WithSlotID(ct, slotID)
	logus.Warningf(ct, "This is a warning message with same trace ID and additional slotID=%s", slotID)
  
	logus.Logf(c, logus.SeverityError, "This is an error log message without trace ID")
}

func init() {
	logus.AddLogEntryHandler(logus.StandardGoLogger())
	Bar(context.Background())
}

```

## Logus log entry handlers

- [logus/go_logger.go](./go_logger.go) - log entries to STDOUT & STDERR.
- [github.com/strongo/loguscloud](https://github.com/strongo/logusgcloud) - send log entries to Google Cloud Logging.
  Support grouping of log entries by request (using trace & span ID).

