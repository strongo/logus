package logus

import "context"

var traceIDKey = "traceID"

// WithTraceID adds trace ID to context
func WithTraceID(c context.Context, traceID string) context.Context {
	return context.WithValue(c, &traceIDKey, traceID)
}

// GetTraceID returns trace ID from context
func GetTraceID(c context.Context) string {
	if traceID, ok := c.Value(&traceIDKey).(string); ok {
		return traceID
	}
	return ""
}

var spanIDKey = "spanID"

// WithSpanID adds span ID to context
func WithSpanID(c context.Context, spanID string) context.Context {
	return context.WithValue(c, &spanIDKey, spanID)
}

// GetSpanID returns span ID from context
func GetSpanID(c context.Context) string {
	if spanID, ok := c.Value(&spanIDKey).(string); ok {
		return spanID
	}
	return ""
}

var labelsKey = "labels"

// WithLabels adds labels to context
func WithLabels(c context.Context, labels map[string]string) context.Context {
	return context.WithValue(c, &labelsKey, labels)
}

// GetLabels returns labels from context
func GetLabels(c context.Context) map[string]string {
	if labels, ok := c.Value(&labelsKey).(map[string]string); ok {
		return labels
	}
	return nil
}
