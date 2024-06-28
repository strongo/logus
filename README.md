# Package: [github.com/strongo/logus](https://github.com/strongo/logus)

Package strongo/logus provides structured context aware logging for Go (golang)
that supports grouping of log entries using trace and span IDs.

## Usage

```go
package foo

import "context"
import "github.com/strongo/logus"

func init() {
  logus.AddLogEntryHandler(logus.StandardGoLogger())
  Bar(context.Background())
}

// Bar demonstrates how to use logus logger 
func Bar(c context.Context) {

	logus.Debugf(c, "This is a debug message without trace ID, unless it was set outside")
  
	const traceID = "123"
	ct := logus.WithTraceID(c, traceID)
	logus.Infof(ct, "This is an info message with a traceID=%s", traceID)

	const spanID = "456"
	ct = logus.WithSpanID(ct, spanID)
	logus.Warningf(ct, "A warning with same trace ID and additional spanID=%s", spanID)
  
	logus.Logf(c, logus.SeverityError, "This is an error log message without trace ID")
}
```

## Logus log entry handlers

- [logus/go_logger.go](./go_logger.go) - log entries to STDOUT & STDERR.
- [github.com/strongo/loguscloud](https://github.com/strongo/logusgcloud) - send log entries to Google Cloud Logging.
  Support grouping of log entries by request (using trace & span ID).

