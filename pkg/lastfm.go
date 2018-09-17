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
	var inModel SimilarArtistsInput

	url := fmt.Sprintf("%smethod=artist.getsimilar&artist=%s&api_key=%s&limit=%d&format=json",
		baseURL, query, client.apiKey, 30)

	data, err := unet.Fetch(url)
	if err != nil {
		return SimilarArtists{}, err
	}

	if err := json.Unmarshal(data, &inModel); err != nil {
		return SimilarArtists{}, err
	}

	similarData := inModel.Data
	artists := similarData.Artists
	input := similarData.Input.Artist

	return SimilarArtists{
		Artists: artists,
		Input:   input,
	}, nil
}

func (client *Client) GetArtistInfo(query string) (ArtistInfo, error) {
	var inModel ArtistInfoInput

	url := fmt.Sprintf("%smethod=artist.getinfo&artist=%s&api_key=%s&format=json",
		baseURL, query, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return ArtistInfo{}, err
	}

	if err = json.Unmarshal(data, &inModel); err != nil {
		return ArtistInfo{}, err
	}

	return inModel.ArtistInfo, nil
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

	if err = json.Unmarshal(data, &inModel); err != nil {
		return AlbumInfo{}, err
	}

	return inModel.AlbumInfo, nil
}

func (client *Client) GetTopTracks(user string) ([]Track, error) {
	var inModel TopTracksInput

	url := fmt.Sprintf("%smethod=user.gettoptracks&user=%s&api_key=%s&format=json",
		baseURL, user, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return []Track{}, err
	}

	if err = json.Unmarshal(data, &inModel); err != nil {
		return []Track{}, err
	}

	topTracksData := inModel.Data
	return topTracksData.Tracks, nil
}

func (client *Client) GetRecentTracks(user string) ([]RecentTrack, error) {
	var inModel RecentTracksInput

	url := fmt.Sprintf("%smethod=user.getrecenttracks&user=%s&api_key=%s&format=json",
		baseURL, user, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return []RecentTrack{}, err
	}

	if err = json.Unmarshal(data, &inModel); err != nil {
		return []RecentTrack{}, err
	}

	recentTracks := inModel.RecentTracks
	return recentTracks.Tracks, nil
}
