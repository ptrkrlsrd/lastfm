package models

import (
	"encoding/json"
	"fmt"
)

// AlbumInfo ...
type AlbumInfo struct {
	Name      string `json:"name"`
	Artist    string `json:"artist"`
	Mbid      string `json:"mbid"`
	URL       string `json:"url"`
	Images    Images `json:"images,omitempty"`
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
	Tracks    struct {
		Track []Track `json:"track"`
	} `json:"tracks"`
	Tags struct {
		Tag []Tag `json:"tag"`
	} `json:"tags"`
	Wiki Bio `json:"wiki"`
}

// UnmarshalJSON UnmarshalJSON
func (albumInfo *AlbumInfo) UnmarshalJSON(data []byte) error {
	var imgs struct {
		Images []Image `json:"image,omitempty"`
	}

	if err := json.Unmarshal(data, &imgs); err != nil {
		return err
	}

	type Alias AlbumInfo
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(albumInfo),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	albumInfo.Images.TransformImages(imgs.Images)
	return nil
}

// Summary ...
func (albumInfo *AlbumInfo) Summary() string {
	return fmt.Sprintf("Bio:\n\n%s", albumInfo.Wiki.Content)
}

// SimpleAlbum ...
type SimpleAlbum struct {
	Name string `json:"#text"`
	Mbid string `json:"mbid"`
}
