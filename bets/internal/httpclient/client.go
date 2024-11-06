package httpclient

import (
	"net/http"
	"net/url"
	"regexp"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func New() *http.Client {
	return &http.Client{Transport: newOpenTelemetryTransport()}
}

func newOpenTelemetryTransport() *otelhttp.Transport {
	return otelhttp.NewTransport(
		http.DefaultTransport,
		otelhttp.WithSpanNameFormatter(formatSpanName),
		otelhttp.WithMetricAttributesFn(extractMetricAttributes),
	)
}

var uuidRegex = regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}`)

// formatSpanName creates a span name based on the operation and HTTP request.
func formatSpanName(op string, req *http.Request) string {
	unescapedPath, err := url.PathUnescape(req.URL.Path)
	if err != nil {
		unescapedPath = req.URL.Path
	}

	formattedPath := sanitizeUUIDs(unescapedPath)

	return req.Method + " " + formattedPath
}

// extractMetricAttributes retrieves metric attributes from an HTTP request.
func extractMetricAttributes(req *http.Request) []attribute.KeyValue {
	unescapedPath, err := url.PathUnescape(req.URL.Path)
	if err != nil {
		unescapedPath = req.URL.Path
	}

	formattedPath := sanitizeUUIDs(unescapedPath)

	return []attribute.KeyValue{semconv.HTTPRoute(formattedPath)}
}

// sanitizeUUIDs replaces UUID patterns in the path with "{id}".
func sanitizeUUIDs(path string) string {
	return uuidRegex.ReplaceAllString(path, "{id}")
}
