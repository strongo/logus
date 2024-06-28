package logus

import (
	"context"
	"testing"
)

func TestGetSlotID(t *testing.T) {
	const slotID = "slot123"
	ctx := context.WithValue(context.Background(), &slotIDKey, slotID)
	if got := GetSlotID(ctx); got != slotID {
		t.Errorf("GetSlotID() = %s, want %s", got, slotID)
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

func TestWithSlotID(t *testing.T) {
	const slotID = "slot123"
	ctx := WithSlotID(context.Background(), slotID)
	if actual, ok := ctx.Value(&slotIDKey).(string); !ok {
		t.Errorf("WithSlotID() = %T, want %T", slotID, slotID)
	} else if actual != slotID {
		t.Errorf("WithSlotID() = %v, want %s", actual, slotID)
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
