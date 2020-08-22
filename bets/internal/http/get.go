package http

import (
	"github.com/angelokurtis/football-bets/bets/internal/log"
	"io/ioutil"
	"net/http"
)

var ignoredHeaders = []string{"Content-Length", "User-Agent", "Accept-Encoding", "Accept", "Connection"}

func Get(url string, headers http.Header) ([]byte, error) {
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	for key := range headers {
		if !shouldIgnore(key) {
			val := headers.Get(key)
			req.Header.Add(key, val)
			log.Debugf("set header '%s' = '%s'", key, val)
		}
	}

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

func shouldIgnore(e string) bool {
	for _, a := range ignoredHeaders {
		if a == e {
			return true
		}
	}
	return false
}
