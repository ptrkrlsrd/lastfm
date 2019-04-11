package lastfm

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/ptrkrlsrd/utilities/unet"
)

// Track ...
type Track struct {
	Attr struct {
		Rank string `json:"rank"`
	} `json:"@attr"`
	Artist     Artist      `json:"artist"`
	Duration   string      `json:"duration"`
	Image      []Image     `json:"image"`
	Mbid       string      `json:"mbid"`
	Name       string      `json:"name"`
	Playcount  string      `json:"playcount"`
	Streamable interface{} `json:"streamable"`
	URL        string      `json:"url"`
}

// ToString ..
func (track *Track) ToString() string {
	return fmt.Sprintf("%s: %s", track.Artist.Name, track.Name)
}

// RecentTracks RecentTracks
type RecentTracks struct {
	Attr struct {
		Page       string `json:"page"`
		PerPage    string `json:"perPage"`
		Total      string `json:"total"`
		TotalPages string `json:"totalPages"`
		User       string `json:"user"`
	} `json:"@attr"`
	Tracks []RecentTrack `json:"track"`
}

// RecentTrack ...
type RecentTrack struct {
	Album  SimpleAlbum  `json:"album"`
	Artist SimpleArtist `json:"artist"`
	Date   struct {
		Text string `json:"#text"`
		Uts  string `json:"uts"`
	} `json:"date"`
	Image      []Image `json:"image"`
	Images     Images  `json:"images,omitempty"`
	Mbid       string  `json:"mbid"`
	Name       string  `json:"name"`
	Streamable string  `json:"streamable"`
	URL        string  `json:"url"`
}

// UnmarshalJSON ...
func (recentTrack *RecentTrack) UnmarshalJSON(data []byte) error {
	var imgs struct {
		Images []Image `json:"image,omitempty"`
	}

	fmt.Println(string(data))

	if err := json.Unmarshal(data, &imgs); err != nil {
		return nil
	}

	type Alias RecentTrack
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(recentTrack),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	recentTrack.Images.TransformImages(imgs.Images)
	return nil
}

// ToString returns a track as a simple string
func (recentTrack RecentTrack) ToString() string {
	date := recentTrack.Date.Text
	format := "%s %s: %s"
	if date == "" {
		date = "Listening now"
	}

	date = fmt.Sprintf("%-18v", date)

	return fmt.Sprintf(format, date, recentTrack.Artist.Name, recentTrack.Name)
}

// ToColoredString returns a track as a colored string
func (recentTrack RecentTrack) ToColoredString() string {
	date := recentTrack.Date.Text
	if date == "" {
		date = "Listening now"
	}

	date = fmt.Sprintf("%-18v", date)

	d := color.New(color.FgWhite, color.Bold)
	return fmt.Sprintf("%s %s - %s", date, d.Sprint(recentTrack.Artist.Name), d.Sprint(recentTrack.Name))
}

// GetTopTracks ...
func (client *Client) GetTopTracks(user string) (tracks []RecentTrack, err error) {
	var lastfmAPIResponse struct {
		Tracks RecentTracks `json:"toptracks"`
	}

	url := generateURL("user.gettoptracks", fmt.Sprintf("user=%s", user), client.apiKey)
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

	url := generateURL("user.getrecenttracks", fmt.Sprintf("user=%s", user), client.apiKey)
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
