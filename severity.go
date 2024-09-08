package logus

type Severity uint8

const (
	SeverityDebug Severity = iota
	SeverityDefault
	SeverityInfo
	SeverityNotice
	SeverityWarning
	SeverityError
	SeverityCritical
	SeverityAlert
)

var SeverityNames = []string{
	"DEBUG",
	"DEFAULT",
	"INFO",
	"NOTICE",
	"WARNING",
	"ERROR",
	"CRITICAL",
	"ALERT",
}
