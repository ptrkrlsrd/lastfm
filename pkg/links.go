package lastfm

// Link ..
type Link struct {
	Text string `json:"#text,omitempty"`
	Rel  string `json:"rel,omitempty"`
	Href string `json:"href,omitempty"`
}

// Links ...
type Links struct {
	Link Link `json:"link,omitempty"`
}
