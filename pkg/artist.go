package lastfm

import (
	"encoding/json"
	"fmt"

	"github.com/ptrkrlsrd/utilities/unet"
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

// Summary gives a summary of an ArtistInfo struct
func (artistInfo *ArtistInfo) Summary() string {
	bioString := artistInfo.Bio.Content
	tagsString := artistInfo.Tags.ToString("\n")

	return fmt.Sprintf("Bio:\n\n%s\nTags:\n\n%s\n", bioString, tagsString)
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
	Name string `json:"name"`
	Mbid string `json:"mbid"`
}

// UnmarshalJSON Custom unmarshal JSON
func (artist *SimpleArtist) UnmarshalJSON(data []byte) error {
	type Alias struct {
		Name string `json:"#text"`
		Mbid string `json:"mbid"`
	}

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(artist),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	return nil
}

type similarArtistsResponse struct {
	Data struct {
		Artists []Artist `json:"artist"`
		Query   struct {
			Artist string
		} `json:"@attr"`
	} `json:"similarartists"`
}

// GetSimilarArtists ...
func (client *Client) GetSimilarArtists(artist string) (similarArtists SimilarArtists, err error) {
	var lastfmAPIResponse similarArtistsResponse

	url := generateURL("artist.getsimilar", fmt.Sprintf("artist=%s", artist), client.apiKey)
	data, err := unet.Fetch(url)
	if err != nil {
		return similarArtists, err
	}

	if err := json.Unmarshal(data, &lastfmAPIResponse); err != nil {
		return similarArtists, err
	}

	similarData := lastfmAPIResponse.Data
	artists := similarData.Artists
	inputQuery := similarData.Query.Artist

	return SimilarArtists{Artists: artists, Query: inputQuery}, nil
}

// GetArtistInfo ...
func (client *Client) GetArtistInfo(artist string) (artistInfo ArtistInfo, err error) {
	url := generateURL("artist.getinfo", fmt.Sprintf("artist=%s", artist), client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return artistInfo, err
	}

	var lastfmAPIResponse struct {
		Info ArtistInfo `json:"artist"`
	}

	if err = json.Unmarshal(data, &lastfmAPIResponse); err != nil {
		return artistInfo, err
	}

	return lastfmAPIResponse.Info, nil
}

// GetTopTracksByArtist ...
func (client *Client) GetTopTracksByArtist(artist string) (artistInfo ArtistInfo, err error) {
	url := generateURL("artist.gettoptracks", fmt.Sprintf("artist=%s", artist), client.apiKey)
	data, err := unet.Fetch(url)
	if err != nil {
		return artistInfo, err
	}

	var lastfmAPIResponse struct {
		Info ArtistInfo `json:"artist"`
	}

	if err = json.Unmarshal(data, &lastfmAPIResponse); err != nil {
		return artistInfo, err
	}

	return lastfmAPIResponse.Info, nil
}

// GetArtistsTopAlbumsByName ...
func (client *Client) GetArtistsTopAlbumsByName(artist string) (topAlbums []Album, err error) {
	url := generateURL("artist.gettopalbums", fmt.Sprintf("artist=%s", artist), client.apiKey)
	return client.getArtistsTopAlbums(url)
}

// GetArtistsTopAlbumsByID ...
func (client *Client) GetArtistsTopAlbumsByID(artist string) (topAlbums []Album, err error) {
	url := generateURL("artist.gettopalbums", fmt.Sprintf("mbid=%s", artist), client.apiKey)
	return client.getArtistsTopAlbums(url)
}

func (client *Client) getArtistsTopAlbums(url string) (topAlbums []Album, err error) {
	data, err := unet.Fetch(url)
	if err != nil {
		return []Album{}, err
	}

	var lastfmAPIResponse struct {
		TopAlbumResponse TopAlbumResponse `json:"topalbums"`
	}

	if err = json.Unmarshal(data, &lastfmAPIResponse); err != nil {
		return []Album{}, err
	}
	return lastfmAPIResponse.TopAlbumResponse.Albums, nil
}
