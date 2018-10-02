package models

import (
	"encoding/json"
)

type AlbumInfo struct {
	Name      string            `json:"name"`
	Artist    string            `json:"artist"`
	Mbid      string            `json:"mbid"`
	URL       string            `json:"url"`
	Images    map[string]string `json:"images,omitempty"`
	Listeners string            `json:"listeners"`
	Playcount string            `json:"playcount"`
	Tracks    struct {
		Track []Track `json:"track"`
	} `json:"tracks"`
	Tags struct {
		Tag []Tag `json:"tag"`
	} `json:"tags"`
	Wiki Bio `json:"wiki"`
}

// UnmarshalJSON UnmarshalJSON
func (u *AlbumInfo) UnmarshalJSON(data []byte) error {
	type Alias AlbumInfo
	var imgs = struct {
		Images []Image `json:"image,omitempty"`
	}{}

	err := json.Unmarshal(data, &imgs)
	if err != nil {
		return err
	}

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	err = json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	u.Images = TransformImages(imgs.Images)
	return nil
}

// SimpleAlbum ...
type SimpleAlbum struct {
	Name string `json:"#text"`
	Mbid string `json:"mbid"`
}
