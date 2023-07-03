// Package matches provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package matches

// Championship defines model for Championship.
type Championship struct {
	Links *ChampionshipLinks `json:"_links,omitempty"`
	Name  *string            `json:"name,omitempty"`
	Year  *int64             `json:"year,omitempty"`
}

// ChampionshipLinks defines model for ChampionshipLinks.
type ChampionshipLinks struct {
	Championship *Link `json:"championship,omitempty"`
	Self         *Link `json:"self,omitempty"`
	Teams        *Link `json:"teams,omitempty"`
}

// Embedded defines model for Embedded.
type Embedded struct {
	Matches *[]Match `json:"matches,omitempty"`
}

// Link defines model for Link.
type Link struct {
	Href *string `json:"href,omitempty"`
}

// Match defines model for Match.
type Match struct {
	Links     *MatchLinks `json:"_links,omitempty"`
	Date      *string     `json:"date,omitempty"`
	ScoreAway *Score      `json:"score_away,omitempty"`
	ScoreHome *Score      `json:"score_home,omitempty"`
	Status    *string     `json:"status,omitempty"`
}

// MatchLinks defines model for MatchLinks.
type MatchLinks struct {
	Championship *Link `json:"championship,omitempty"`
	Match        *Link `json:"match,omitempty"`
	Self         *Link `json:"self,omitempty"`
}

// Matches defines model for Matches.
type Matches struct {
	Embedded *Embedded     `json:"_embedded,omitempty"`
	Links    *MatchesLinks `json:"_links,omitempty"`
}

// MatchesLinks defines model for MatchesLinks.
type MatchesLinks struct {
	Self *Link `json:"self,omitempty"`
}

// Score defines model for Score.
type Score struct {
	Links *ScoreLinks `json:"_links,omitempty"`
	Goals *int64      `json:"goals,omitempty"`
}

// ScoreLinks defines model for ScoreLinks.
type ScoreLinks struct {
	Team *Link `json:"team,omitempty"`
}
