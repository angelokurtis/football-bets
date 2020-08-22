package teams

import (
	"encoding/json"
	"net/url"
	"os"
)

import (
	"github.com/angelokurtis/football-bets/bets/internal/http"
)

func Get(href string, headers map[string][]string) ([]byte, error) {
	u := &url.URL{
		Scheme: "http",
		Host:   os.Getenv("TEAMS_ADDRESS"),
		Path:   href,
	}
	body, err := http.Get(u.String(), headers)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetOne(href string, headers map[string][]string) (*Team, error) {
	body, err := Get(href, headers)
	if err != nil {
		return nil, err
	}
	var obj *Team
	if err := json.Unmarshal(body, &obj); err != nil {
		return nil, err
	}
	return obj, nil
}

type Team struct {
	Links struct {
		Championships struct {
			Href string `json:"href"`
		} `json:"championships"`
		Matches struct {
			Href string `json:"href"`
		} `json:"matches"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Team struct {
			Href string `json:"href"`
		} `json:"team"`
	} `json:"_links"`
	ImageURL string `json:"image_url,omitempty"`
	Name     string `json:"name"`
}
