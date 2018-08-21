package lastfm

import "encoding/json"

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

// UnmarshalJSON Unmarsh
func (u *AlbumInfo) UnmarshalJSON(data []byte) error {
	type Alias AlbumInfo
	imgs := struct {
		Info struct {
			Images []Image `json:"image,omitempty"`
		} `json:"album"`
	}{}
	err := json.Unmarshal(data, &imgs)

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	err = json.Unmarshal(data, &aux)
	u.Images = TransformImages(imgs.Info.Images)
	return err
}

type SimpleAlbum struct {
	Name string `json:"#text"`
	Mbid string `json:"mbid"`
}
