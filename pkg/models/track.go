package models

import (
	"fmt"
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

// TopTracks ...
type TopTracks struct {
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
	Mbid       string  `json:"mbid"`
	Name       string  `json:"name"`
	Streamable string  `json:"streamable"`
	URL        string  `json:"url"`
}

// ToString ...
func (recentTrack RecentTrack) ToString() string {
	date := recentTrack.Date.Text
	if date == "" {
		date = "playing now"
	}
	return fmt.Sprintf("%s) %s: %s\n", date, recentTrack.Artist.Name, recentTrack.Name)
}
