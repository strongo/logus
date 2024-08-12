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
