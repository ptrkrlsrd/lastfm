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
	apiKey string
}

func NewClient(apiKey string) Client {
	return Client{apiKey: apiKey}
}

func (client *Client) GetSimilarArtists(query string) (SimilarArtists, error) {
	url := fmt.Sprintf("%smethod=artist.getsimilar&artist=%s&api_key=%s&limit=%d&format=json",
		baseURL, query, client.apiKey, 30)

	var inModel struct {
		Data struct {
			Artists []Artist `json:"artist"`
			Input   struct {
				Artist string `json:"artist"`
			} `json:"@attr"`
		} `json:"similarartists"`
	}

	data, err := unet.Fetch(url)
	if err != nil {
		return SimilarArtists{}, err
	}

	err = json.Unmarshal(data, &inModel)
	if err != nil {
		return SimilarArtists{}, err
	}

	return SimilarArtists{
		Artists: inModel.Data.Artists,
		Input:   inModel.Data.Input.Artist,
	}, err
}

func (client *Client) GetArtistInfo(query string) (ArtistInfo, error) {
	var inModel struct {
		ArtistInfo ArtistInfo `json:"artist"`
	}

	url := fmt.Sprintf("%smethod=artist.getinfo&artist=%s&api_key=%s&format=json",
		baseURL, query, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return ArtistInfo{}, err
	}

	err = json.Unmarshal(data, &inModel)
	return inModel.ArtistInfo, err
}

func (client *Client) GetAlbumInfo(artist string, album string) (AlbumInfo, error) {
	var inModel struct {
		AlbumInfo AlbumInfo `json:"album"`
	}

	url := fmt.Sprintf("%smethod=album.getinfo&artist=%s&album=%s&api_key=%s&format=json",
		baseURL, artist, album, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return AlbumInfo{}, err
	}

	err = json.Unmarshal(data, &inModel)
	return inModel.AlbumInfo, err
}

func (client *Client) GetTopTracks(user string) ([]Track, error) {
	var inModel = struct {
		Data struct {
			Attr struct {
				Page       string `json:"page"`
				PerPage    string `json:"perPage"`
				Total      string `json:"total"`
				TotalPages string `json:"totalPages"`
				User       string `json:"user"`
			} `json:"@attr"`
			Tracks []Track `json:"track"`
		} `json:"toptracks"`
	}{}

	url := fmt.Sprintf("%smethod=user.gettoptracks&user=%s&api_key=%s&format=json",
		baseURL, user, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return []Track{}, err
	}

	err = json.Unmarshal(data, &inModel)
	return inModel.Data.Tracks, err
}

func (client *Client) GetRecentTracks(user string) ([]RecentTrack, error) {
	var inModel struct {
		RecentTracks RecentTracks `json:"recenttracks"`
	}

	url := fmt.Sprintf("%smethod=user.getrecenttracks&user=%s&api_key=%s&format=json",
		baseURL, user, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return []RecentTrack{}, err
	}

	err = json.Unmarshal(data, &inModel)
	return inModel.RecentTracks.Tracks, err
}
