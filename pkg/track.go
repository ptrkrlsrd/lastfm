package lastfm

import "fmt"

type Track struct {
	Attr struct {
		Rank string `json:"rank"`
	} `json:"@attr"`
	Artist   Artist `json:"artist"`
	Duration string `json:"duration"`
	Image    []struct {
		Text string `json:"#text"`
		Size string `json:"size"`
	} `json:"image"`
	Mbid       string      `json:"mbid"`
	Name       string      `json:"name"`
	Playcount  string      `json:"playcount"`
	Streamable interface{} `json:"streamable"`
	URL        string      `json:"url"`
}

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

type RecentTrack struct {
	Album  SimpleAlbum  `json:"album"`
	Artist SimpleArtist `json:"artist"`
	Date   struct {
		Text string `json:"#text"`
		Uts  string `json:"uts"`
	} `json:"date"`
	Image []struct {
		Text string `json:"#text"`
		Size string `json:"size"`
	} `json:"image"`
	Mbid       string `json:"mbid"`
	Name       string `json:"name"`
	Streamable string `json:"streamable"`
	URL        string `json:"url"`
}

func (recentTrack RecentTrack) ToString() string {
	date := recentTrack.Date.Text
	if date == "" {
		date = "playing now"
	}
	return fmt.Sprintf("%s) %s: %s\n", date, recentTrack.Artist.Name, recentTrack.Name)
}
