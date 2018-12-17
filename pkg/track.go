package lastfm

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
	Mbid       string  `json:"mbid"`
	Name       string  `json:"name"`
	Streamable string  `json:"streamable"`
	URL        string  `json:"url"`
}

// ToString ...
func (recentTrack RecentTrack) ToString() string {
	date := recentTrack.Date.Text
	format := "%s: %s"
	if date == "" {
		date = "now"
	}

	return fmt.Sprintf(format, recentTrack.Artist.Name, recentTrack.Name)
}
