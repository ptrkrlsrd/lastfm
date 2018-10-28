package lastfm

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
