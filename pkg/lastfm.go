package lastfm

import (
	"encoding/json"
	"fmt"

	"github.com/ptrkrlsrd/lastfm/pkg/models"
	"github.com/ptrkrlsrd/utilities/pkg/unet"
)

const (
	baseURL = "http://ws.audioscrobbler.com/2.0/?"
)

// Client ...
type Client struct {
	apiKey string
}

// NewClient ...
func NewClient(apiKey string) Client {
	return Client{apiKey: apiKey}
}

// GetSimilarArtists ...
func (client *Client) GetSimilarArtists(query string) (similarArtists models.SimilarArtists, err error) {
	var inModel struct {
		Data struct {
			Artists []models.Artist `json:"artist"`
			Query   struct {
				Artist string
			} `json:"@attr"`
		} `json:"similarartists"`
	}

	url := fmt.Sprintf("%smethod=artist.getsimilar&artist=%s&api_key=%s&limit=%d&format=json",
		baseURL, query, client.apiKey, 30)

	data, err := unet.Fetch(url)
	if err != nil {
		return similarArtists, err
	}

	if err := json.Unmarshal(data, &inModel); err != nil {
		return similarArtists, err
	}

	similarData := inModel.Data
	artists := similarData.Artists
	inputQuery := similarData.Query.Artist

	return models.SimilarArtists{Artists: artists, Query: inputQuery}, nil
}

// GetArtistInfo ...
func (client *Client) GetArtistInfo(query string) (artistInfo models.ArtistInfo, err error) {
	var url = fmt.Sprintf("%smethod=artist.getinfo&artist=%s&api_key=%s&format=json",
		baseURL, query, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return artistInfo, err
	}

	var inModel struct {
		Info models.ArtistInfo `json:"artist"`
	}

	if err = json.Unmarshal(data, &inModel); err != nil {
		return artistInfo, err
	}

	return inModel.Info, nil
}

// GetAlbumInfo ...
func (client *Client) GetAlbumInfo(artist string, album string) (albumInfo models.AlbumInfo, err error) {
	var inModel struct {
		Info models.AlbumInfo `json:"album"`
	}

	url := fmt.Sprintf("%smethod=album.getinfo&artist=%s&album=%s&api_key=%s&format=json",
		baseURL, artist, album, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return albumInfo, err
	}

	if err = json.Unmarshal(data, &inModel); err != nil {
		return albumInfo, err
	}

	return inModel.Info, nil
}

// GetTopTracks ...
func (client *Client) GetTopTracks(user string) (tracks []models.RecentTrack, err error) {
	var inModel struct {
		Tracks models.TopTracks `json:"toptracks"`
	}

	url := fmt.Sprintf("%smethod=user.gettoptracks&user=%s&api_key=%s&format=json",
		baseURL, user, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return tracks, err
	}

	if err = json.Unmarshal(data, &inModel); err != nil {
		return tracks, err
	}

	topTracksData := inModel.Tracks
	return topTracksData.Tracks, nil
}

// GetRecentTracks ...
func (client *Client) GetRecentTracks(user string) (tracks models.TopTracks, err error) {
	var inModel struct {
		Tracks models.TopTracks `json:"recentTracks"`
	}

	url := fmt.Sprintf("%smethod=user.getrecenttracks&user=%s&api_key=%s&format=json",
		baseURL, user, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return tracks, err
	}

	if err = json.Unmarshal(data, &inModel); err != nil {
		return tracks, err
	}

	recentTracks := inModel.Tracks
	if len(recentTracks.Tracks) == 0 {
		return tracks, fmt.Errorf("no results")
	}

	return recentTracks, nil
}
