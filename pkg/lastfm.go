package lastfm

import (
	"encoding/json"
	"fmt"

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
func (client *Client) GetSimilarArtists(query string) (similarArtists SimilarArtists, err error) {
	var lastfmAPIResponse struct {
		Data struct {
			Artists []Artist `json:"artist"`
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

	if err := json.Unmarshal(data, &lastfmAPIResponse); err != nil {
		return similarArtists, err
	}

	similarData := lastfmAPIResponse.Data
	artists := similarData.Artists
	inputQuery := similarData.Query.Artist

	return SimilarArtists{Artists: artists, Query: inputQuery}, nil
}

// GetArtistInfo ...
func (client *Client) GetArtistInfo(query string) (artistInfo ArtistInfo, err error) {
	var url = fmt.Sprintf("%smethod=artist.getinfo&artist=%s&api_key=%s&format=json",
		baseURL, query, client.apiKey)

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

// GetAlbumInfo ...
func (client *Client) GetAlbumInfo(artist string, album string) (albumInfo AlbumInfo, err error) {
	var lastfmAPIResponse struct {
		Info AlbumInfo `json:"album"`
	}

	url := fmt.Sprintf("%smethod=album.getinfo&artist=%s&album=%s&api_key=%s&format=json",
		baseURL, artist, album, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return albumInfo, err
	}

	if err = json.Unmarshal(data, &lastfmAPIResponse); err != nil {
		return albumInfo, err
	}

	return lastfmAPIResponse.Info, nil
}

// GetAlbumInfo ...
func (client *Client) GetAlbumInfoByID(id string) (albumInfo AlbumInfo, err error) {
	var lastfmAPIResponse struct {
		Info AlbumInfo `json:"album"`
	}

	url := fmt.Sprintf("%smethod=album.getinfo&mbid=%s&api_key=%s&format=json",
		baseURL, id, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return albumInfo, err
	}

	if err = json.Unmarshal(data, &lastfmAPIResponse); err != nil {
		return albumInfo, err
	}

	return lastfmAPIResponse.Info, nil
}

// GetTopTracks ...
func (client *Client) GetTopTracks(user string) (tracks []RecentTrack, err error) {
	var lastfmAPIResponse struct {
		Tracks RecentTracks `json:"toptracks"`
	}

	url := fmt.Sprintf("%smethod=user.gettoptracks&user=%s&api_key=%s&format=json",
		baseURL, user, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return tracks, err
	}

	if err = json.Unmarshal(data, &lastfmAPIResponse); err != nil {
		return tracks, err
	}

	topTracksData := lastfmAPIResponse.Tracks
	return topTracksData.Tracks, nil
}

// GetRecentTracks ...
func (client *Client) GetRecentTracks(user string) (tracks RecentTracks, err error) {
	var lastfmAPIResponse struct {
		Tracks RecentTracks `json:"recentTracks"`
	}

	url := fmt.Sprintf("%smethod=user.getrecenttracks&user=%s&api_key=%s&format=json",
		baseURL, user, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return tracks, err
	}

	if err = json.Unmarshal(data, &lastfmAPIResponse); err != nil {
		return tracks, err
	}

	recentTracks := lastfmAPIResponse.Tracks
	if len(recentTracks.Tracks) == 0 {
		return tracks, fmt.Errorf("no results")
	}

	return recentTracks, nil
}
