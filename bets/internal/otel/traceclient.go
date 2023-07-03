package otel

import (
	"fmt"
	"net/url"
	"os"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

func newTraceClient() (otlptrace.Client, error) {
	env, ok := os.LookupEnv(envExporterOTLPTracesEndpoint)
	if !ok {
		env = os.Getenv(envExporterOTLPEndpoint)
	}

	endpoint, err := url.Parse(env)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new trace client: %w", err)
	}

	opts := []otlptracegrpc.Option{otlptracegrpc.WithEndpoint(endpoint.Host)}
	if endpoint.Scheme != "https" {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}

	return otlptracegrpc.NewClient(opts...), nil
}
