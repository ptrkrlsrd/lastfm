package lastfm

import (
	"encoding/json"
	"fmt"

	"github.com/ptrkrlsrd/utilities/pkg/unet"
)

const (
	baseURL = "http://ws.audioscrobbler.com/2.0/?"
)

type Client struct {
	APIKey string
}

func NewClient(apiKey string) Client {
	client := Client{APIKey: apiKey}
	return client
}

func (client *Client) GetSimilarArtists(query string) (SimilarArtists, error) {
	url := fmt.Sprintf("%smethod=artist.getsimilar&artist=%s&api_key=%s&limit=%d&format=json",
		baseURL, query, client.APIKey, 30)

	var inModel = struct {
		Data struct {
			Artists []Artist `json:"artist"`
			Input   struct {
				Artist string `json:"artist"`
			} `json:"@attr"`
		} `json:"similarartists"`
	}{}

	data, err := unet.FetchData(url)
	if err != nil {
		return SimilarArtists{}, err
	}

	err = json.Unmarshal(data, &inModel)
	if err != nil {
		return SimilarArtists{}, err
	}

	similarartists := SimilarArtists{
		Artists: inModel.Data.Artists,
		Input:   inModel.Data.Input.Artist,
	}

	return similarartists, err
}

func (client *Client) GetArtistInfo(query string) (ArtistInfo, error) {
	var result ArtistInfo
	url := fmt.Sprintf("%smethod=artist.getinfo&artist=%s&api_key=%s&format=json",
		baseURL, query, client.APIKey)

	data, err := unet.FetchData(url)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

func (client *Client) GetAlbumInfo(artist string, album string) (AlbumInfo, error) {
	var result AlbumInfo
	url := fmt.Sprintf("%smethod=album.getinfo&artist=%s&album=%s&api_key=%s&format=json",
		baseURL, artist, album, client.APIKey)

	data, err := unet.FetchData(url)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}
