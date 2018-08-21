package lastfm

const (
	imageSmall      = "small"
	imageMedium     = "medium"
	imageLarge      = "large"
	imageExtraLarge = "extralarge"
	imageMega       = "mega"
)

type Image struct {
	URL  string `json:"#text,omitempty"`
	Size string `json:"size,omitempty"`
}

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
