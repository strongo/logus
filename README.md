# Package: [github.com/strongo/logus](https://github.com/strongo/logus)

Package strongo/logus provides structured context aware logging for Go (golang)
that supports grouping of log entries using trace and span IDs.

<!-- dev-approach:v1 -->
## Our approach to development

We build with our own tooling:

- **[SpecScore](https://specscore.md)** — specify requirements as `SpecScore.md` artifacts
- **[SpecStudio](https://specscore.studio)** — author & manage specs across their lifecycle
- **[inGitDB](https://ingitdb.com)** — store structured data in Git where applicable
- **[DALgo](https://dalgo.io)** — data access layer for Go
- **[cover100.dev](https://cover100.dev)** — drive toward 100% test coverage
- **[DataTug](https://datatug.io)** — query & explore data
<!-- /dev-approach -->

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



## DataTug

This project is enhanced with [DataTug](https://datatug.app). See the [datatug](./datatug) directory for details.
