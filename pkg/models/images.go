package models

const (
	imageSmall      = "small"
	imageMedium     = "medium"
	imageLarge      = "large"
	imageExtraLarge = "extralarge"
	imageMega       = "mega"
)

// Image ..
type Image struct {
	URL  string `json:"#text,omitempty"`
	Size string `json:"size,omitempty"`
}

// Images ...
type Images struct {
	Mega       string `json:"mega"`
	ExtraLarge string `json:"extralarge"`
	Large      string `json:"large"`
	Medium     string `json:"medium"`
	Small      string `json:"small"`
}

// TransformImages ...
func (images *Images) TransformImages(imageSlice []Image) {
	for _, v := range imageSlice {
		switch v.Size {
		case imageMega:
			images.Mega = v.URL
		case imageExtraLarge:
			images.ExtraLarge = v.URL
		case imageLarge:
			images.Large = v.URL
		case imageMedium:
			images.Medium = v.URL
		case imageSmall:
			images.Small = v.URL
		}
	}
}
