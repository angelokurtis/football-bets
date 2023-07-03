package otel

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewTransport() *otelhttp.Transport {
	return otelhttp.NewTransport(http.DefaultTransport)
}
