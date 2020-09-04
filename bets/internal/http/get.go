package http

import (
	"github.com/angelokurtis/football-bets/bets/internal/log"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string, headers http.Header) ([]byte, error) {
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for key := range headers {
		if shouldBePropagated(key) {
			val := headers.Get(key)
			req.Header.Add(key, val)
		}
	}

	log.Debugf("get req to %s", url)
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	status := &status{code: res.StatusCode}
	switch status.Group() {
	case ClientErrorResponse:
		return nil, errors.WithStack(NewClientError(method, url, res.StatusCode))
	case ServerErrorResponse:
		return nil, errors.WithStack(NewServerError(method, url, res.StatusCode))
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return body, nil
}

func shouldBePropagated(header string) bool {
	return strings.HasPrefix(header, "X-B3-") || header == "Authorization"
}
