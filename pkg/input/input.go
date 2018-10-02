package input

import "github.com/ptrkrlsrd/lastfm/pkg/models"

type SimilarArtists struct {
	Data struct {
		Artists []models.Artist `json:"artist"`
		Query   struct {
			Artist string `json:"artist"`
		} `json:"@attr"`
	} `json:"similarartists"`
}

type RecentTracks struct {
	RecentTracks models.RecentTracks `json:"recenttracks"`
}

type TopTracks struct {
	Data struct {
		Attr struct {
			Page       string `json:"page"`
			PerPage    string `json:"perPage"`
			Total      string `json:"total"`
			TotalPages string `json:"totalPages"`
			User       string `json:"user"`
		} `json:"@attr"`
		Tracks []models.Track `json:"track"`
	} `json:"toptracks"`
}

type AlbumInfo struct {
	AlbumInfo models.AlbumInfo `json:"album"`
}

type ArtistInfo struct {
	ArtistInfo models.ArtistInfo `json:"artist"`
}
