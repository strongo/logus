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

var slotIDKey = "slotID"

// WithSlotID adds slot ID to context
func WithSlotID(c context.Context, slotID string) context.Context {
	return context.WithValue(c, &slotIDKey, slotID)
}

// GetSlotID returns slot ID from context
func GetSlotID(c context.Context) string {
	if slotID, ok := c.Value(&slotIDKey).(string); ok {
		return slotID
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
