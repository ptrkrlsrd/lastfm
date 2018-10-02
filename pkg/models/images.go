package models

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
