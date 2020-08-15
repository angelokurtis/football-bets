package championships

import (
	"encoding/json"
	"fmt"
	"github.com/angelokurtis/football-bets/bets/internal/http"
)

func Get(href string) ([]byte, error) {
	url := fmt.Sprintf("http://championships:8080%s", href)
	body, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetOne(href string) (*Championship, error) {
	body, err := Get(href)
	if err != nil {
		return nil, err
	}
	var obj *Championship
	if err := json.Unmarshal(body, &obj); err != nil {
		return nil, err
	}
	return obj, nil
}

type Championship struct {
	Links struct {
		Championship struct {
			Href string `json:"href"`
		} `json:"championship"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Teams struct {
			Href string `json:"href"`
		} `json:"teams"`
	} `json:"_links"`
	Name string `json:"name"`
	Year int    `json:"year"`
}
