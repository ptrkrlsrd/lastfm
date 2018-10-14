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
type Images map[string]string

// TransformImages ...
func (images *Images) TransformImages(imageSlice []Image) {
	var output = make(Images)

	keys := []string{imageSmall,
		imageMedium,
		imageLarge,
		imageExtraLarge,
		imageMega}

	for i, v := range imageSlice {
		if i < len(keys) {
			key := keys[i]
			output[key] = v.URL
		}
	}

	*images = output
}

// GetLargestImage ..
func (images Images) GetLargestImage() string {
	imgURL := images[imageExtraLarge]
	if imgURL != "" {
		return imgURL
	}

	sizes := []string{imageMega, imageExtraLarge, imageLarge, imageMedium}

	for _, v := range sizes {
		if images[v] != "" {
			return images[v]
		}
	}

	return imgURL
}
