package lastfm

import "encoding/json"

type Artist struct {
	Name   string            `json:"name"`
	Mbid   string            `json:"mbid"`
	Match  string            `json:"match,omitempty"`
	URL    string            `json:"url"`
	Images map[string]string `json:"images,omitempty"`
}

func (artist *Artist) UnmarshalJSON(data []byte) error {
	imgs := struct {
		Images []Image `json:"image,omitempty"`
	}{}
	err := json.Unmarshal(data, &imgs)

	type Alias Artist
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(artist),
	}

	err = json.Unmarshal(data, &aux)

	artist.Images = TransformImages(imgs.Images)
	return err
}

type SimilarArtists struct {
	Artists []Artist `json:"similar"`
	Input   string   `json:"input,omitempty"`
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

	imgs := struct {
		Images []Image `json:"image,omitempty"`
	}{}

	err := json.Unmarshal(data, &imgs)
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	err = json.Unmarshal(data, &aux)
	u.Images = TransformImages(imgs.Images)
	return err
}

func TransformImages(images []Image) map[string]string {
	var output = make(map[string]string)
	keys := []string{imageSmall, imageMedium, imageLarge, imageExtraLarge, imageMega}
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
