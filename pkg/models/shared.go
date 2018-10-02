package models

// Tag ..
type Tag struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Tags ...
type Tags struct {
	Tag []Tag `json:"tag"`
}

// ToString ...
func (tags *Tags) ToString(delimiter string) string {
	var tagsString string

	for _, v := range tags.Tag {
		tagsString += v.Name + delimiter
	}

	return tagsString
}

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

// Bio ...
type Bio struct {
	Links     Links  `json:"links,omitempty"`
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

// Stats ...
type Stats struct {
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
}
