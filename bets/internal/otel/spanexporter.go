package otel

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func newSpanExporter(ctx context.Context, client otlptrace.Client) (trace.SpanExporter, error) {
	tracesExporter := "otlp"
	if env, ok := os.LookupEnv(envTracesExporter); ok {
		tracesExporter = env
	}

	switch tracesExporter {
	case "none":
		return tracetest.NewNoopExporter(), nil
	case "otlp":
		exp, err := otlptrace.New(ctx, client)
		if err != nil {
			return nil, fmt.Errorf("failed to create a new OTLP trace exporter: %w", err)
		}

		return exp, nil
	default:
		return nil, fmt.Errorf("unrecognized value for traces exporter: %s", tracesExporter)
	}
}
