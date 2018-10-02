package models

import (
	"encoding/json"
)

type Artist struct {
	Name   string            `json:"name"`
	Mbid   string            `json:"mbid"`
	Match  string            `json:"match,omitempty"`
	URL    string            `json:"url"`
	Images map[string]string `json:"images,omitempty"`
}

func (artist *Artist) UnmarshalJSON(data []byte) error {
	var imgs = struct {
		Images []Image `json:"image,omitempty"`
	}{}

	if err := json.Unmarshal(data, &imgs); err != nil {
		return err
	}

	type Alias Artist
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(artist),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	artist.Images = TransformImages(imgs.Images)
	return nil
}

type SimilarArtists struct {
	Artists []Artist `json:"similar"`
	Query   string   `json:"input,omitempty"`
}

type ArtistInfo struct {
	Name   string            `json:"name"`
	Mbid   string            `json:"mbid"`
	URL    string            `json:"url"`
	Images map[string]string `json:"images,omitempty"`
	Ontour string            `json:"ontour"`
	Stats  struct {
		Listeners string `json:"listeners"`
		Playcount string `json:"playcount"`
	} `json:"stats"`
	Similar struct {
		Artists []Artist `json:"artist"`
	} `json:"similar"`
	Tags struct {
		Tag []Tag `json:"tag"`
	} `json:"tags"`
	Bio Bio `json:"bio"`
}

func (u *ArtistInfo) UnmarshalJSON(data []byte) error {
	type Alias ArtistInfo

	var imgs = struct {
		Images []Image `json:"image,omitempty"`
	}{}

	if err := json.Unmarshal(data, &imgs); err != nil {
		return nil
	}

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	u.Images = TransformImages(imgs.Images)
	return nil
}

func TransformImages(images []Image) map[string]string {
	var output = make(map[string]string)
	keys := []string{imageSmall,
		imageMedium,
		imageLarge,
		imageExtraLarge,
		imageMega}

	for i, v := range images {
		if i < len(keys) {
			key := keys[i]
			output[key] = v.URL
		}
	}

	return output
}

type SimpleArtist struct {
	Name string `json:"#text"`
	Mbid string `json:"mbid"`
}
