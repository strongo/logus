package logus

type Severity int

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
