package lastfm

import (
	"encoding/json"
)

type Image struct {
	URL  string `json:"#text,omitempty"`
	Size string `json:"size,omitempty"`
}

type Artist struct {
	Name   string            `json:"name"`
	Mbid   string            `json:"mbid"`
	Match  string            `json:"match,omitempty"`
	URL    string            `json:"url"`
	Images map[string]string `json:"images,omitempty"`
}

func (artist *Artist) UnmarshalJSON(data []byte) error {
	imgs := struct {
		Images []Image `json:"image,omitempty"`
	}{}
	err := json.Unmarshal(data, &imgs)

	type Alias Artist
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(artist),
	}

	err = json.Unmarshal(data, &aux)
	images := map[string]string{}
	keys := []string{"small", "medium", "large", "extralarge", "mega"}
	for i, v := range imgs.Images {
		if i < len(keys) {
			key := keys[i]
			images[key] = v.URL
		}
	}

	artist.Images = images
	return err
}

type SimilarArtists struct {
	Artists []Artist `json:"similar"`
	Input   string   `json:"input,omitempty"`
}

type Tag struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Link struct {
	Text string `json:"#text,omitempty"`
	Rel  string `json:"rel,omitempty"`
	Href string `json:"href,omitempty"`
}

type Bio struct {
	Links struct {
		Link Link `json:"link,omitempty"`
	} `json:"links,omitempty"`
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

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
	Mbid       string `json:"mbid"`
	Name       string `json:"name"`
	Playcount  string `json:"playcount"`
	Streamable struct {
		Text      string `json:"#text"`
		Fulltrack string `json:"fulltrack"`
	} `json:"streamable"`
	URL string `json:"url"`
}

type ArtistInfo struct {
	Info struct {
		Name   string            `json:"name"`
		Mbid   string            `json:"mbid"`
		URL    string            `json:"url"`
		Images map[string]string `json:"images,omitempty"`
		Ontour string            `json:"ontour"`
		Stats  struct {
			Listeners string `json:"listeners"`
			Playcount string `json:"playcount"`
		} `json:"stats"`
		Similar struct {
			Artists []Artist `json:"artist"`
		} `json:"similar"`
		Tags struct {
			Tag []Tag `json:"tag"`
		} `json:"tags"`
		Bio Bio `json:"bio"`
	} `json:"artist"`
}

func (u *ArtistInfo) UnmarshalJSON(data []byte) error {
	type Alias ArtistInfo

	imgs := struct {
		Images []Image `json:"image,omitempty"`
	}{}

	err := json.Unmarshal(data, &imgs)
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	err = json.Unmarshal(data, &aux)
	images := map[string]string{}
	keys := []string{"small", "medium", "large", "extralarge", "mega"}
	for i, v := range imgs.Images {
		if i < len(keys) {
			key := keys[i]
			images[key] = v.URL
		}
	}

	u.Info.Images = images
	return err
}

type AlbumInfo struct {
	Info struct {
		Name      string            `json:"name"`
		Artist    string            `json:"artist"`
		Mbid      string            `json:"mbid"`
		URL       string            `json:"url"`
		Images    map[string]string `json:"images,omitempty"`
		Listeners string            `json:"listeners"`
		Playcount string            `json:"playcount"`
		Tracks    struct {
			Track []Track `json:"track"`
		} `json:"tracks"`
		Tags struct {
			Tag []Tag `json:"tag"`
		} `json:"tags"`
		Wiki Bio `json:"wiki"`
	} `json:"album"`
}

// UnmarshalJSON Unmarsh
func (u *AlbumInfo) UnmarshalJSON(data []byte) error {
	type Alias AlbumInfo
	imgs := struct {
		Info struct {
			Images []Image `json:"image,omitempty"`
		} `json:"album"`
	}{}
	err := json.Unmarshal(data, &imgs)

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	err = json.Unmarshal(data, &aux)
	images := map[string]string{}
	keys := []string{"small", "medium", "large", "extralarge", "mega"}
	for i, v := range imgs.Info.Images {
		if i < len(keys) {
			key := keys[i]
			images[key] = v.URL
		}
	}

	u.Info.Images = images
	return err
}

type InArtistInfo struct {
	Info struct {
		Name   string  `json:"name"`
		Mbid   string  `json:"mbid"`
		URL    string  `json:"url"`
		Images []Image `json:"image"`
		Ontour string  `json:"ontour"`
		Stats  struct {
			Listeners string `json:"listeners"`
			Playcount string `json:"playcount"`
		} `json:"stats"`
		Similar struct {
			Artists []Artist `json:"artist"`
		} `json:"similar"`
		Tags struct {
			Tag []Tag `json:"tag"`
		} `json:"tags"`
		Bio Bio `json:"bio"`
	} `json:"artist"`
}

type Toptracks struct {
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
}
