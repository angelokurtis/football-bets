package matches

import (
	"fmt"
	"net/url"
	"os"
)

func NewClientWithHTTPClient(httpClient HttpRequestDoer) (*ClientWithResponses, error) {
	u, err := url.Parse(os.Getenv("MATCHES_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to create Teams client: %w", err)
	}

	c, err := NewClientWithResponses(u.String(), WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("failed to create Teams client: %w", err)
	}

	return c, err
}
