package httpclient

import (
	"net/http"
)

func New(transport http.RoundTripper) *http.Client {
	return &http.Client{Transport: transport}
}
