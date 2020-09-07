package teams

import (
	"context"
	"encoding/json"
	"github.com/angelokurtis/football-bets/bets/internal/http"
	"github.com/pkg/errors"
	"net/url"
	"os"
)

func Get(c context.Context, href string) ([]byte, error) {
	u := &url.URL{
		Scheme: "http",
		Host:   os.Getenv("TEAMS_ADDRESS"),
		Path:   href,
	}
	body, err := http.Get(c, u.String())
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetOne(c context.Context, href string) (*Team, error) {
	body, err := Get(c, href)
	if err != nil {
		return nil, err
	}
	var obj *Team
	if err := json.Unmarshal(body, &obj); err != nil {
		return nil, errors.WithStack(err)
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
