package httpclient

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/angelokurtis/football-bets/bets/internal/otel"
)

func New() *http.Client {
	return &http.Client{Transport: newOpenTelemetryTransport()}
}

func newOpenTelemetryTransport() *otelhttp.Transport {
	return otelhttp.NewTransport(
		http.DefaultTransport,
		otelhttp.WithSpanNameFormatter(otel.FormatSpanName),
		otelhttp.WithMetricAttributesFn(otel.ExtractMetricAttributes),
	)
}
