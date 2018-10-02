package models

type Tag struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Link struct {
	Text string `json:"#text,omitempty"`
	Rel  string `json:"rel,omitempty"`
	Href string `json:"href,omitempty"`
}

type Bio struct {
	Links struct {
		Link Link `json:"link,omitempty"`
	} `json:"links,omitempty"`
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}
