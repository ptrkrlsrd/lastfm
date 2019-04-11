package lastfm

import (
	"encoding/json"
	"fmt"

	"github.com/ptrkrlsrd/utilities/unet"
)

// AlbumInfo Info about an album
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

// UnmarshalJSON Custom unmarshal of JSON which transforms images correctly
func (albumInfo *AlbumInfo) UnmarshalJSON(data []byte) error {
	type Alias AlbumInfo
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(albumInfo),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var imgs struct {
		Images []Image `json:"image,omitempty"`
	}

	if err := json.Unmarshal(data, &imgs); err != nil {
		return err
	}

	albumInfo.Images.TransformImages(imgs.Images)
	return nil
}

// Summary ...
func (albumInfo *AlbumInfo) Summary() string {
	return fmt.Sprintf("Bio:\n\n%s", albumInfo.Wiki.Content)
}

type Album struct {
	Artist    Artist  `json:"artist"`
	Image     []Image `json:"image"`
	Images    Images  `json:"images,omitempty"`
	Mbid      string  `json:"mbid"`
	Name      string  `json:"name"`
	Playcount int64   `json:"playcount"`
	URL       string  `json:"url"`
}

// UnmarshalJSON Custom unmarshal of JSON which transforms images correctly
func (album *Album) UnmarshalJSON(data []byte) error {
	type Alias Album
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(album),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var imgs struct {
		Images []Image `json:"image,omitempty"`
	}

	if err := json.Unmarshal(data, &imgs); err != nil {
		return err
	}

	album.Images.TransformImages(imgs.Images)
	return nil
}

type TopAlbumResponse struct {
	Attr struct {
		Artist     string `json:"artist"`
		Page       string `json:"page"`
		PerPage    string `json:"perPage"`
		Total      string `json:"total"`
		TotalPages string `json:"totalPages"`
	} `json:"@attr"`
	Albums []Album `json:"album"`
}

type AlbumInfoResponse struct {
	Info AlbumInfo `json:"album"`
}

type TopAlbums struct {
	Data TopAlbumResponse `json:"topalbums"`
}

// SimpleAlbum A simple album struct containing just a Name and Mbid
type SimpleAlbum struct {
	Name string `json:"name"`
	Mbid string `json:"mbid"`
}

// UnmarshalJSON Custom unmarshal JSON
func (album *SimpleAlbum) UnmarshalJSON(data []byte) error {
	type Alias struct {
		Name string `json:"#text"`
		Mbid string `json:"mbid"`
	}

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(album),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	return nil
}

// GetAlbumInfo ...
func (client *Client) GetAlbumInfo(artist string, album string) (albumInfo AlbumInfo, err error) {
	var albumInfoResponse AlbumInfoResponse
	url := generateURL("album.getinfo", fmt.Sprintf("artist=%s&album=%s", artist, album), client.apiKey)
	data, err := unet.Fetch(url)

	if err != nil {
		return albumInfo, err
	}

	if err = json.Unmarshal(data, &albumInfoResponse); err != nil {
		return albumInfo, err
	}

	return albumInfoResponse.Info, nil
}

// GetAlbumInfoByID ...
func (client *Client) GetAlbumInfoByID(id string) (albums []Album, err error) {
	var topAlbums TopAlbums

	url := generateURL("album.gettopalbums", fmt.Sprintf("mbid=%s", id), client.apiKey)
	data, err := unet.Fetch(url)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &topAlbums); err != nil {
		return nil, err
	}

	return topAlbums.Data.Albums, nil
}
