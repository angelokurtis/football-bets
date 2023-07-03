package otel

import "net/http"

func SpanFormatter(_ string, r *http.Request) string {
	return r.Method + " " + r.URL.Path
}
