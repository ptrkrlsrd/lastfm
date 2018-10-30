package lastfm

const (
	imageSmall      = "small"      // imageSmall The smallest image size
	imageMedium     = "medium"     // imageMedium A medium image size
	imageLarge      = "large"      // imageLarge A large image size
	imageExtraLarge = "extralarge" // imageExtraLarge A extra large image size
	imageMega       = "mega"       // imageMega The biggest image size
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

// TransformImages Transform images from the way LastFM handles images to the way handled here
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
