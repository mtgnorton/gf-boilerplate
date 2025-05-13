package st

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var noopSpanInstance = &noopSpan{}

// noopSpan 是一个不执行任何操作的 Span 实现
type noopSpan struct {
	trace.Span
}

func (s noopSpan) End(options ...trace.SpanEndOption) {
}
func (s noopSpan) SpanContext() trace.SpanContext {
	return trace.SpanContext{}
}
func (s noopSpan) SetStatus(code codes.Code, description string) {}

func (s noopSpan) IsRecording() bool {
	return false
}
func (s noopSpan) SetName(name string) {}

func (s noopSpan) SetAttributes(kv ...attribute.KeyValue) {}

func (s noopSpan) AddEvent(name string, options ...trace.EventOption) {}

func (s noopSpan) RecordError(err error, opts ...trace.EventOption) {}

func (s noopSpan) AddLink(link trace.Link) {}

func (s noopSpan) TracerProvider() trace.TracerProvider {
	return nil
}
