package lastfm

import (
	"encoding/json"
	"fmt"

	"github.com/ptrkrlsrd/lastfm/pkg/input"
	"github.com/ptrkrlsrd/lastfm/pkg/models"
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

func (client *Client) GetSimilarArtists(query string) (similarArtists models.SimilarArtists, err error) {
	var inModel input.SimilarArtists

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

func (client *Client) GetArtistInfo(query string) (artistInfo models.ArtistInfo, err error) {
	var inModel input.ArtistInfo

	url := fmt.Sprintf("%smethod=artist.getinfo&artist=%s&api_key=%s&format=json",
		baseURL, query, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return artistInfo, err
	}

	if err = json.Unmarshal(data, &inModel); err != nil {
		return artistInfo, err
	}

	return inModel.ArtistInfo, nil
}

func (client *Client) GetAlbumInfo(artist string, album string) (albumInfo models.AlbumInfo, err error) {
	var inModel input.AlbumInfo

	url := fmt.Sprintf("%smethod=album.getinfo&artist=%s&album=%s&api_key=%s&format=json",
		baseURL, artist, album, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return albumInfo, err
	}

	if err = json.Unmarshal(data, &inModel); err != nil {
		return albumInfo, err
	}

	return inModel.AlbumInfo, nil
}

func (client *Client) GetTopTracks(user string) (tracks []models.Track, err error) {
	var inModel input.TopTracks

	url := fmt.Sprintf("%smethod=user.gettoptracks&user=%s&api_key=%s&format=json",
		baseURL, user, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return tracks, err
	}

	if err = json.Unmarshal(data, &inModel); err != nil {
		return tracks, err
	}

	topTracksData := inModel.Data
	return topTracksData.Tracks, nil
}

func (client *Client) GetRecentTracks(user string) (tracks []models.RecentTrack, err error) {
	var inModel input.RecentTracks

	url := fmt.Sprintf("%smethod=user.getrecenttracks&user=%s&api_key=%s&format=json",
		baseURL, user, client.apiKey)

	data, err := unet.Fetch(url)
	if err != nil {
		return tracks, err
	}

	if err = json.Unmarshal(data, &inModel); err != nil {
		return tracks, err
	}

	recentTracks := inModel.RecentTracks
	return recentTracks.Tracks, nil
}
