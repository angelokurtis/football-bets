package championships

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
