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
	if got := GetSpanID(context.Background()); got != "" {
		t.Errorf("GetSpanID() = %s, want empty", got)
	}
}

func TestGetTraceID(t *testing.T) {
	const traceID = "trace123"
	ctx := context.WithValue(context.Background(), &traceIDKey, traceID)
	if got := GetTraceID(ctx); got != traceID {
		t.Errorf("GetTraceID() = %s, want %s", got, traceID)
	}
	if got := GetTraceID(context.Background()); got != "" {
		t.Errorf("GetTraceID() = %s, want empty", got)
	}
}

func TestGetLabels(t *testing.T) {
	labels := map[string]string{
		"l1": "v1",
		"l2": "v2",
	}
	ctx := WithLabels(context.Background(), labels)
	got := GetLabels(ctx)
	if len(got) != 2 || got["l1"] != "v1" || got["l2"] != "v2" {
		t.Errorf("GetLabels() = %v, want %v", got, labels)
	}
	if gotEmpty := GetLabels(context.Background()); gotEmpty != nil {
		t.Errorf("GetLabels() = %v, want nil", gotEmpty)
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
	actual, ok := ctx.Value(&labelsKey).(map[string]string)
	if !ok {
		t.Fatalf("WithLabels() = %T, want %T", actual, labels)
	}
	if len(actual) != len(labels) {
		t.Fatalf("WithLabels() = %d, want %d", len(actual), len(labels))
	}
	for k, v := range labels {
		if actual[k] != v {
			t.Errorf("WithLabels()[%s] = %s, want %s", k, actual[k], v)
		}
	}
}
