package http

import (
	"context"
	"github.com/angelokurtis/football-bets/bets/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(c context.Context, url string) ([]byte, error) {
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	propagateContext(c, req)

	log := logger.New(c)
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

func propagateContext(c context.Context, req *http.Request) {
	if g, ok := c.(*gin.Context); ok {
		header := g.Request.Header
		for key := range header {
			if shouldBePropagated(key) {
				val := header.Get(key)
				req.Header.Add(key, val)
			}
		}
	}
}

func shouldBePropagated(header string) bool {
	return strings.HasPrefix(header, "X-B3-") || header == "Authorization"
}
