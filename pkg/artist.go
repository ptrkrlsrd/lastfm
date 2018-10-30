package lastfm

import (
	"encoding/json"
)

// Artist An artist
type Artist struct {
	Name   string `json:"name"`
	Mbid   string `json:"mbid"`
	Match  string `json:"match,omitempty"`
	URL    string `json:"url"`
	Images Images `json:"images,omitempty"`
}

// UnmarshalJSON Unmarshal JSON and transform the images correctly
func (artist *Artist) UnmarshalJSON(data []byte) error {
	var imgs struct {
		Images []Image `json:"image,omitempty"`
	}

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

	artist.Images.TransformImages(imgs.Images)
	return nil
}

// SimilarArtists Similar artists and the users search query
type SimilarArtists struct {
	Artists []Artist `json:"similar"`
	Query   string   `json:"input,omitempty"`
}

// ArtistInfo Artist info
type ArtistInfo struct {
	Name    string `json:"name"`
	Mbid    string `json:"mbid"`
	URL     string `json:"url"`
	Images  Images `json:"images,omitempty"`
	Ontour  string `json:"ontour"`
	Stats   Stats  `json:"stats"`
	Similar struct {
		Artists []Artist `json:"artist"`
	} `json:"similar"`
	Tags Tags `json:"tags"`
	Bio  Bio  `json:"bio"`
}

// Summary ...
func (artistInfo *ArtistInfo) Summary() string {
	bioString := "Bio:\n\n" + artistInfo.Bio.Content + "\n"
	tagsString := "Tags:\n\n" + artistInfo.Tags.ToString("\n")

	return bioString + tagsString
}

// UnmarshalJSON ...
func (artistInfo *ArtistInfo) UnmarshalJSON(data []byte) error {
	var imgs struct {
		Images []Image `json:"image,omitempty"`
	}

	if err := json.Unmarshal(data, &imgs); err != nil {
		return nil
	}

	type Alias ArtistInfo
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(artistInfo),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	artistInfo.Images.TransformImages(imgs.Images)
	return nil
}

// SimpleArtist ...
type SimpleArtist struct {
	Name string `json:"#text"`
	Mbid string `json:"mbid"`
}
