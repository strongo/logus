package logus

import (
	"context"
	"testing"
)

func TestGetSpanID(t *testing.T) {
	const spanID = "span123"
	ctx := context.WithValue(context.Background(), &spanIDKey, spanID)
	if got := GetSpanID(ctx); got != spanID {
		t.Errorf("GetSpanID() = %s, want %s", got, spanID)
	}
}

func TestGetTraceID(t *testing.T) {
	const traceID = "trace123"
	ctx := context.WithValue(context.Background(), &traceIDKey, traceID)
	if got := GetTraceID(ctx); got != traceID {
		t.Errorf("GetTraceID() = %s, want %s", got, traceID)
	}
}

func TestGetLabels(t *testing.T) {
	labels := map[string]string{
		"l1": "v1",
		"l2": "v2",
	}
	ctx := context.WithValue(context.Background(), &traceIDKey, labels)
	if got := GetLabels(ctx); &got == &labels {
		t.Errorf("GetTraceID() = %p, want %p", got, labels)
	}
}

func TestWithSpanID(t *testing.T) {
	const spanID = "span123"
	ctx := WithSpanID(context.Background(), spanID)
	if actual, ok := ctx.Value(&spanIDKey).(string); !ok {
		t.Errorf("WithSpanID() = %T, want %T", spanID, spanID)
	} else if actual != spanID {
		t.Errorf("WithSpanID() = %v, want %s", actual, spanID)
	}
}

func TestWithTraceID(t *testing.T) {
	const traceID = "trace123"
	ctx := WithTraceID(context.Background(), traceID)
	if actual, ok := ctx.Value(&traceIDKey).(string); !ok {
		t.Errorf("WithTraceID() = %T, want %T", actual, traceID)
	} else if actual != traceID {
		t.Errorf("WithTraceID() = %v, want %s", actual, traceID)
	}
}

func TestWithLabels(t *testing.T) {
	labels := map[string]string{
		"l1": "v1",
		"l2": "v2",
	}
	ctx := WithLabels(context.Background(), labels)
	if actual, ok := ctx.Value(&labelsKey).(map[string]string); !ok {
		t.Errorf("WithLabels() = %T, want %T", actual, labels)
	} else if &actual != &actual {
		t.Errorf("WithLabels() = %p, want %p", actual, labels)
	}
}
