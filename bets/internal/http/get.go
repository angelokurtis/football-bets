package http

import (
	"github.com/angelokurtis/football-bets/bets/internal/log"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string, headers http.Header) ([]byte, error) {
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
func shouldBePropagated(header string) bool {
	return strings.HasPrefix(header, "X-B3-") || header == "Authorization"
}
