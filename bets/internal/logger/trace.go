package logger

import (
	"context"
	"fmt"
	"net/http"
)

type trace struct {
	TraceId string
	SpanId  string
	Sampled string
}

func traceFromHeader(h http.Header) *trace {
	context.Background()
	return &trace{
		TraceId: h.Get("X-B3-TraceId"),
		SpanId:  h.Get("X-B3-SpanId"),
		Sampled: h.Get("X-B3-Sampled"),
	}
}

func (b *trace) String() string {
	if len(b.TraceId) > 0 || len(b.SpanId) > 0 || len(b.Sampled) > 0 {
		return fmt.Sprintf("[%s,%s,%s] ", b.TraceId, b.SpanId, b.Sampled)
	}
	return ""
}
