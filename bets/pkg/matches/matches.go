package matches

import (
	"encoding/json"
	"fmt"
	"github.com/angelokurtis/football-bets/bets/internal/http"
	"math/rand"
	"time"
)

func Get(href string) ([]byte, error) {
	url := fmt.Sprintf("http://matches:8080%s", href)
	body, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetAll() ([]*Match, error) {
	body, err := Get("/matches")
	if err != nil {
		return nil, err
	}
	var obj *Response
	if err := json.Unmarshal(body, &obj); err != nil {
		return nil, err
	}
	return obj.Embedded.Matches, nil
}

func GetRandomly() (*Match, error) {
	matches, err := GetAll()
	if err != nil {
		return nil, err
	}
	if len(matches) == 0 {
		return nil, nil
	}
	matches = shuffle(matches)
	return matches[0], nil
}

func shuffle(matches []*Match) []*Match {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	res := make([]*Match, len(matches))
	n := len(matches)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(matches))
		res[i] = matches[randIndex]
		matches = append(matches[:randIndex], matches[randIndex+1:]...)
	}
	return res
}

type Response struct {
	Embedded struct {
		Matches []*Match `json:"matches"`
	} `json:"_embedded"`
}

type Match struct {
	Links struct {
		Championship struct {
			Href string `json:"href"`
		} `json:"championship"`
		Match struct {
			Href string `json:"href"`
		} `json:"match"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Date      string `json:"date"`
	ScoreAway struct {
		Links struct {
			Team struct {
				Href string `json:"href"`
			} `json:"team"`
		} `json:"_links"`
		Goals int `json:"goals"`
	} `json:"score_away"`
	ScoreHome struct {
		Links struct {
			Team struct {
				Href string `json:"href"`
			} `json:"team"`
		} `json:"_links"`
		Goals int `json:"goals"`
	} `json:"score_home"`
	Status string `json:"status"`
}
